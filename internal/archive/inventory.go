package archive

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/jtprogru/weekly-incident/internal/model"
)

// StoredIncident retains a physical occurrence, including repeated keys. An
// audit must not silently upsert duplicates before it can report them.
type StoredIncident struct {
	Path     string         `json:"path"`
	Position int            `json:"position"` // One-based JSON array position.
	Incident model.Incident `json:"incident"`
}

// InputFile identifies the exact bytes read, relative to the archive root.
type InputFile struct {
	Path   string `json:"path"`
	SHA256 string `json:"sha256"`
}

// Inventory is a read-only snapshot of vendor files, excluding derived data.
type Inventory struct {
	Weeks       []Week
	Records     []StoredIncident
	Files       []InputFile
	Fingerprint string
}

var yearDirectory = regexp.MustCompile(`^[0-9]{4}$`)
var weekDirectory = regexp.MustCompile(`^W[0-9]{2}$`)

// ReadAll loads only YYYY/Wnn/<vendor>.json, in lexical path order. A missing
// root or malformed vendor file is an error, never an apparently clean audit.
// It neither follows symlinks nor changes the archive.
func (a *Archive) ReadAll() (Inventory, error) {
	out := Inventory{Records: []StoredIncident{}, Files: []InputFile{}}
	years, err := os.ReadDir(a.Root)
	if err != nil {
		return out, fmt.Errorf("read archive: %w", err)
	}
	for _, year := range years {
		if !yearDirectory.MatchString(year.Name()) {
			continue
		}
		if !year.IsDir() {
			return out, fmt.Errorf("archive year %s is not a directory (symlinks are not supported)", year.Name())
		}
		weeks, err := os.ReadDir(filepath.Join(a.Root, year.Name()))
		if err != nil {
			return out, fmt.Errorf("read year %s: %w", year.Name(), err)
		}
		for _, week := range weeks {
			if !weekDirectory.MatchString(week.Name()) {
				continue
			}
			w, err := ParseWeek(year.Name() + "-" + week.Name())
			if err != nil || WeekOf(w.Start()) != w {
				return out, fmt.Errorf("invalid ISO week directory %s/%s", year.Name(), week.Name())
			}
			if !week.IsDir() {
				return out, fmt.Errorf("archive week %s is not a directory (symlinks are not supported)", w)
			}
			if err := a.readInventoryWeek(&out, w); err != nil {
				return out, err
			}
			out.Weeks = append(out.Weeks, w)
		}
	}
	// JSON framing makes the hash unambiguous even for unusual file names.
	manifest, err := json.Marshal(out.Files)
	if err != nil {
		return out, fmt.Errorf("encode input manifest: %w", err)
	}
	out.Fingerprint = digestBytes(manifest)
	return out, nil
}

func (a *Archive) readInventoryWeek(out *Inventory, w Week) error {
	entries, err := os.ReadDir(filepath.Join(a.Root, w.Dir()))
	if err != nil {
		return fmt.Errorf("read week %s: %w", w, err)
	}
	for _, entry := range entries {
		name := entry.Name()
		if name == indexFile || !strings.HasSuffix(name, ".json") {
			continue
		}
		rel := filepath.ToSlash(filepath.Join(w.Dir(), name))
		if !entry.Type().IsRegular() {
			return fmt.Errorf("archive file %s is not a regular file", rel)
		}
		b, err := os.ReadFile(filepath.Join(a.Root, filepath.FromSlash(rel)))
		if err != nil {
			return fmt.Errorf("read %s: %w", rel, err)
		}
		var incidents []model.Incident
		if err := json.Unmarshal(b, &incidents); err != nil {
			return fmt.Errorf("decode %s: %w", rel, err)
		}
		if incidents == nil {
			return fmt.Errorf("decode %s: expected an incident array, got null", rel)
		}
		vendor := strings.TrimSuffix(name, ".json")
		for i, in := range incidents {
			if err := in.Validate(); err != nil {
				return fmt.Errorf("%s record %d: %w", rel, i+1, err)
			}
			if in.Vendor != vendor || in.Key != model.MakeKey(in.Vendor, in.NativeID) {
				return fmt.Errorf("%s record %d: vendor/key disagrees with file or native_id", rel, i+1)
			}
			out.Records = append(out.Records, StoredIncident{Path: rel, Position: i + 1, Incident: in})
		}
		out.Files = append(out.Files, InputFile{Path: rel, SHA256: digestBytes(b)})
	}
	return nil
}

func digestBytes(b []byte) string {
	sum := sha256.Sum256(b)
	return hex.EncodeToString(sum[:])
}
