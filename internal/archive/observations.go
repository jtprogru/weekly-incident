package archive

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strings"
	"time"

	"github.com/jtprogru/weekly-incident/internal/model"
	"github.com/jtprogru/weekly-incident/internal/storage"
)

type Revision struct {
	Sequence int       `json:"sequence"`
	Hash     string    `json:"hash"`
	SeenAt   time.Time `json:"seen_at"`
	Imported bool      `json:"imported"`
	Path     string    `json:"path"`
}

type SourceHistory struct {
	Key            string     `json:"key"`
	Vendor         string     `json:"vendor"`
	FirstSeenAt    time.Time  `json:"first_seen_at"`
	UnknownCurrent bool       `json:"unknown_current"`
	Revisions      []Revision `json:"revisions"`
}

type ObservationManifest struct {
	SchemaVersion int                      `json:"schema_version"`
	Records       map[string]SourceHistory `json:"records"`
}

type Snapshot struct {
	SchemaVersion int            `json:"schema_version"`
	Record        model.Incident `json:"record"`
}

// Observations retains content-addressed snapshots plus revision order. Map
// keys are source keys (currently one configured feed per vendor).
type Observations struct {
	Manifest  ObservationManifest
	Snapshots map[string]Snapshot
}

func EmptyObservations() Observations {
	return Observations{Manifest: ObservationManifest{SchemaVersion: 1, Records: map[string]SourceHistory{}}, Snapshots: map[string]Snapshot{}}
}

func snapshotOf(in model.Incident) (Snapshot, string, error) {
	in.FirstSeenAt = time.Time{}
	// Normalize raw JSON without changing incident semantics or deriving dates.
	if len(in.Raw) > 0 {
		dec := json.NewDecoder(bytes.NewReader(in.Raw))
		dec.UseNumber()
		var v any
		if err := dec.Decode(&v); err != nil {
			return Snapshot{}, "", err
		}
		b, err := json.Marshal(v)
		if err != nil {
			return Snapshot{}, "", err
		}
		in.Raw = b
	}
	s := Snapshot{SchemaVersion: 1, Record: in}
	hash, err := storage.Digest(s)
	return s, hash, err
}

func validVendor(v string) bool {
	return v != "" && v != "." && v != ".." && !strings.ContainsAny(v, "/\\")
}

func (a *Archive) ReadObservations() (Observations, error) {
	o := EmptyObservations()
	err := storage.Read(filepath.Join(a.Root, "observations/manifest.json"), &o.Manifest)
	if errors.Is(err, os.ErrNotExist) {
		return o, nil
	}
	if err != nil {
		return o, err
	}
	if o.Manifest.SchemaVersion != 1 || o.Manifest.Records == nil {
		return o, fmt.Errorf("unsupported observation manifest")
	}
	for key, h := range o.Manifest.Records {
		if h.Key != key || !validVendor(h.Vendor) || len(h.Revisions) == 0 {
			return o, fmt.Errorf("invalid history %s", key)
		}
		for i, r := range h.Revisions {
			_, hashErr := hex.DecodeString(r.Hash)
			if r.Sequence != i+1 || len(r.Hash) != 64 || hashErr != nil || !filepath.IsLocal(r.Path) {
				return o, fmt.Errorf("invalid revision %s/%d", key, i+1)
			}
			var s Snapshot
			if cached, ok := o.Snapshots[r.Hash]; ok {
				s = cached
			} else {
				if err := storage.Read(filepath.Join(a.Root, "observations", h.Vendor, r.Hash+".json"), &s); err != nil {
					return o, err
				}
			}
			hash, err := storage.Digest(s)
			if err != nil || hash != r.Hash || s.SchemaVersion != 1 || s.Record.Key != key || s.Record.Vendor != h.Vendor {
				return o, fmt.Errorf("snapshot integrity mismatch for %s/%s", key, r.Hash)
			}
			o.Snapshots[r.Hash] = s
		}
	}
	return o, nil
}

// Observe does not write. Initial imports with different values for the same
// source key are explicitly unordered. A later fetch supplies a known current
// value, while all imported evidence remains available.
func (o *Observations) Observe(records []StoredIncident, now time.Time, imported bool) error {
	grouped := map[string][]StoredIncident{}
	for _, r := range records {
		if err := r.Incident.Validate(); err != nil {
			return err
		}
		if !validVendor(r.Incident.Vendor) || r.Incident.Key != model.MakeKey(r.Incident.Vendor, r.Incident.NativeID) {
			return fmt.Errorf("invalid source key %q", r.Incident.Key)
		}
		grouped[r.Incident.Key] = append(grouped[r.Incident.Key], r)
	}
	for _, key := range sortedKeys(grouped) {
		rows := grouped[key]
		h, exists := o.Manifest.Records[key]
		if !exists {
			h = SourceHistory{Key: key, Vendor: rows[0].Incident.Vendor, FirstSeenAt: rows[0].Incident.FirstSeenAt, Revisions: []Revision{}}
		}
		hashes := map[string]StoredIncident{}
		for _, row := range rows {
			s, hash, err := snapshotOf(row.Incident)
			if err != nil {
				return err
			}
			if prev, ok := o.Snapshots[hash]; ok {
				p, err := storage.Digest(prev)
				if err != nil || p != hash {
					return fmt.Errorf("snapshot collision %s", hash)
				}
			}
			o.Snapshots[hash] = s
			if prev, ok := hashes[hash]; !ok || row.Path < prev.Path {
				hashes[hash] = row
			}
			if h.FirstSeenAt.IsZero() || (!row.Incident.FirstSeenAt.IsZero() && row.Incident.FirstSeenAt.Before(h.FirstSeenAt)) {
				h.FirstSeenAt = row.Incident.FirstSeenAt
			}
		}
		if len(hashes) > 1 && !imported {
			return fmt.Errorf("multiple current values for source %s", key)
		}
		if len(hashes) > 1 {
			h.UnknownCurrent = true
		}
		for _, hash := range sortedKeys(hashes) {
			if len(h.Revisions) > 0 && h.Revisions[len(h.Revisions)-1].Hash == hash && !h.UnknownCurrent {
				continue
			}
			// Re-reading an unordered legacy file must not add another revision.
			if imported && slices.ContainsFunc(h.Revisions, func(r Revision) bool { return r.Hash == hash }) {
				continue
			}
			h.Revisions = append(h.Revisions, Revision{Sequence: len(h.Revisions) + 1, Hash: hash, SeenAt: now.UTC(), Imported: imported, Path: hashes[hash].Path})
		}
		if !imported {
			h.UnknownCurrent = false
		}
		o.Manifest.Records[key] = h
	}
	return nil
}

// PrepareObservations imports a legacy archive or observes changes in an
// existing one. Stored snapshots not present in today's feed are retained.
func (a *Archive) PrepareObservations(inv Inventory, now time.Time) (Observations, error) {
	o, err := a.ReadObservations()
	if err != nil {
		return o, err
	}
	// Import only unregistered sources. Existing observations are the durable
	// history; an older compatibility projection after a crash is not a new
	// revision. Collect repairs that projection on its next fetch.
	missing := []StoredIncident{}
	for _, row := range inv.Records {
		h, exists := o.Manifest.Records[row.Incident.Key]
		if !exists {
			missing = append(missing, row)
			continue
		}
		_, hash, err := snapshotOf(row.Incident)
		if err != nil {
			return o, err
		}
		if h.UnknownCurrent {
			if !slices.ContainsFunc(h.Revisions, func(r Revision) bool { return r.Hash == hash }) {
				return o, fmt.Errorf("unregistered revision for %s", row.Incident.Key)
			}
		} else if h.Revisions[len(h.Revisions)-1].Hash != hash {
			return o, fmt.Errorf("source projection differs from history for %s; collect again to recover", row.Incident.Key)
		}
	}
	if err := o.Observe(missing, now, true); err != nil {
		return o, err
	}
	return o, nil
}

func (a *Archive) WriteObservations(o Observations) error {
	for _, hash := range sortedKeys(o.Snapshots) {
		s := o.Snapshots[hash]
		path := filepath.Join(a.Root, "observations", s.Record.Vendor, hash+".json")
		b, err := storage.Encode(s)
		if err != nil {
			return err
		}
		if old, err := os.ReadFile(path); err == nil && !bytes.Equal(old, b) {
			return fmt.Errorf("immutable snapshot changed: %s", path)
		}
		if err := storage.Write(path, b); err != nil {
			return err
		}
	}
	return storage.WriteJSON(filepath.Join(a.Root, "observations/manifest.json"), o.Manifest)
}

func (o Observations) Current() ([]model.Incident, error) {
	out := make([]model.Incident, 0, len(o.Manifest.Records))
	for _, key := range sortedKeys(o.Manifest.Records) {
		h := o.Manifest.Records[key]
		if len(h.Revisions) == 0 {
			return nil, fmt.Errorf("empty history %s", key)
		}
		in := o.Snapshots[h.Revisions[len(h.Revisions)-1].Hash].Record
		in.FirstSeenAt = h.FirstSeenAt
		out = append(out, in)
	}
	return out, nil
}

func sortedKeys[V any](m map[string]V) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}
