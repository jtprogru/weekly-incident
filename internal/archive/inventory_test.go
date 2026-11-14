package archive

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/jtprogru/weekly-incident/internal/model"
)

func TestReadAllPreservesOccurrencesAndIgnoresDerivedFiles(t *testing.T) {
	a := New(t.TempDir())
	in := incident("github", "same", weekOf)
	for _, week := range []Week{{2025, 52}, {2026, 34}} {
		if err := a.SaveVendorWeek(week, "github", []model.Incident{in, in}); err != nil {
			t.Fatal(err)
		}
	}
	for _, path := range []string{"state.json", "2026/W34/index.json", "identity/registry.json", "canonical/incidents.json", "observations/github/revision.json"} {
		writeInventoryFile(t, a.Root, path, []byte("not incident JSON"))
	}
	before, err := a.ReadAll()
	if err != nil {
		t.Fatal(err)
	}
	if len(before.Records) != 4 || len(before.Files) != 2 {
		t.Fatalf("records/files: %d/%d", len(before.Records), len(before.Files))
	}
	if before.Records[0].Path != "2025/W52/github.json" || before.Records[1].Position != 2 {
		t.Fatalf("lost occurrence provenance: %+v", before.Records)
	}
	after, err := a.ReadAll()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(before, after) {
		t.Fatal("read changed inventory")
	}
	writeInventoryFile(t, a.Root, "2026/W34/index.json", []byte("changed derived file"))
	after, err = a.ReadAll()
	if err != nil {
		t.Fatal(err)
	}
	if after.Fingerprint != before.Fingerprint {
		t.Fatal("derived files affected source fingerprint")
	}
	path := filepath.Join(a.Root, "2026/W34/github.json")
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	writeInventoryFile(t, a.Root, "2026/W34/github.json", append(b, '\n'))
	after, err = a.ReadAll()
	if err != nil {
		t.Fatal(err)
	}
	if after.Fingerprint == before.Fingerprint {
		t.Fatal("fingerprint did not reflect actual input bytes")
	}
}

func TestReadAllFailsOnInvalidInputs(t *testing.T) {
	valid, err := json.Marshal([]model.Incident{incident("github", "a", weekOf)})
	if err != nil {
		t.Fatal(err)
	}
	for _, tc := range []struct{ name, path, body, want string }{
		{"malformed", "2026/W34/github.json", "{", "decode"},
		{"null", "2026/W34/github.json", "null", "expected an incident array"},
		{"invalid-record", "2026/W34/github.json", "[{}]", "record 1"},
		{"wrong-vendor", "2026/W34/slack.json", string(valid), "vendor/key"},
		{"wrong-key", "2026/W34/github.json", strings.Replace(string(valid), `"key":"github/a"`, `"key":"wrong"`, 1), "vendor/key"},
		{"invalid-week", "2025/W53/github.json", string(valid), "invalid ISO week"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			a := New(t.TempDir())
			writeInventoryFile(t, a.Root, tc.path, []byte(tc.body))
			if _, err := a.ReadAll(); err == nil || !strings.Contains(err.Error(), tc.want) {
				t.Fatalf("error = %v; want %s", err, tc.want)
			}
		})
	}
	if _, err := New(filepath.Join(t.TempDir(), "missing")).ReadAll(); err == nil {
		t.Fatal("missing root accepted")
	}
}

func TestReadAllRejectsSymlinkVendorFile(t *testing.T) {
	root := t.TempDir()
	writeInventoryFile(t, root, "outside.json", []byte("[]"))
	dir := filepath.Join(root, "2026/W34")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(filepath.Join(root, "outside.json"), filepath.Join(dir, "github.json")); err != nil {
		t.Fatal(err)
	}
	if _, err := New(root).ReadAll(); err == nil || !strings.Contains(err.Error(), "not a regular file") {
		t.Fatalf("error = %v", err)
	}
}

func writeInventoryFile(t *testing.T, root, path string, body []byte) {
	t.Helper()
	path = filepath.Join(root, path)
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, body, 0o644); err != nil {
		t.Fatal(err)
	}
}
