// Package archive stores source records and immutable observations. Vendor
// files retain their original ISO-week location; semantic week reads use the
// current StartedAt across the whole archive.
//
// The merge is upsert-only. Statuspage has no pagination and every feed holds a
// fixed number of incidents — 50, which is 62 days for GitHub but only 20 for
// Cloudflare — so an incident vanishing from a feed means it fell out of the
// window, not that it was retracted. Deleting on absence would quietly erase
// history that cannot be fetched again.
package archive

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/jtprogru/weekly-incident/internal/model"
	"github.com/jtprogru/weekly-incident/internal/storage"
)

// Archive is a directory of week folders, normally "data".
type Archive struct {
	Root string
}

// New opens an archive rooted at dir. The directory is created on first write.
func New(dir string) *Archive { return &Archive{Root: dir} }

// Stats summarizes what one Upsert changed.
type Stats struct {
	Created int
	Updated int
	Weeks   []Week
}

// Upsert merges incidents into their week files. now stamps FirstSeenAt on
// records the archive has not seen before; existing records keep the value they
// were first stored with.
//
// New records are routed by StartedAt. Existing keys retain their physical
// location even when the source corrects the start into another week.
func (a *Archive) Upsert(incidents []model.Incident, now time.Time) (Stats, error) {
	now = now.UTC()
	inv, err := a.ReadAll()
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		return Stats{}, err
	}
	obs, err := a.ReadObservations()
	if err == nil {
		missing := []StoredIncident{}
		for _, r := range inv.Records {
			if _, ok := obs.Manifest.Records[r.Incident.Key]; !ok {
				missing = append(missing, r)
			}
		}
		err = obs.Observe(missing, now, true)
	}
	if err != nil {
		return Stats{}, err
	}
	locations := map[string]StoredIncident{}
	occurrences := map[string][]StoredIncident{}
	for _, r := range inv.Records {
		occurrences[r.Incident.Key] = append(occurrences[r.Incident.Key], r)
		if _, ok := locations[r.Incident.Key]; !ok {
			locations[r.Incident.Key] = r
		}
	}
	grouped := map[string][]model.Incident{}
	changedWeeks := map[Week]bool{}
	fresh := []StoredIncident{}
	for _, in := range incidents {
		if err := in.Validate(); err != nil {
			return Stats{}, err
		}
		if !validVendor(in.Vendor) || in.Key != model.MakeKey(in.Vendor, in.NativeID) {
			return Stats{}, fmt.Errorf("invalid source key %q", in.Key)
		}
		path := filepath.Join(WeekOf(in.StartedAt).Dir(), in.Vendor+".json")
		if prev, ok := locations[in.Key]; ok {
			path = prev.Path
			in.FirstSeenAt = obs.Manifest.Records[in.Key].FirstSeenAt
			// Old archives may contain the same key at multiple physical paths.
			// Preserve those files, but refresh every projection from the known
			// current fetch; all former values were imported above as evidence.
			paths := map[string]bool{}
			for _, old := range occurrences[in.Key] {
				if !sameIncident(old.Incident, in) {
					changedWeeks[WeekOf(old.Incident.StartedAt)] = true
					changedWeeks[WeekOf(in.StartedAt)] = true
				}
				if old.Path != path && !paths[old.Path] {
					grouped[old.Path] = append(grouped[old.Path], in)
				}
				paths[old.Path] = true
			}
		} else {
			in.FirstSeenAt = now
			changedWeeks[WeekOf(in.StartedAt)] = true
		}
		grouped[path] = append(grouped[path], in)
		fresh = append(fresh, StoredIncident{Path: filepath.ToSlash(path), Position: 1, Incident: in})
	}
	if err := obs.Observe(fresh, now, false); err != nil {
		return Stats{}, err
	}
	// Persist evidence before updating the compatibility projection.
	if err := a.WriteObservations(obs); err != nil {
		return Stats{}, err
	}
	var st Stats
	for _, path := range sortedKeys(grouped) {
		var existing []model.Incident
		data, err := os.ReadFile(filepath.Join(a.Root, path))
		if err != nil && !errors.Is(err, fs.ErrNotExist) {
			return st, err
		}
		if err == nil {
			if err := json.Unmarshal(data, &existing); err != nil {
				return st, err
			}
		}
		merged, created, updated := merge(existing, grouped[path], now)
		if created == 0 && updated == 0 {
			continue
		}
		if err := writeJSON(filepath.Join(a.Root, path), merged); err != nil {
			return st, err
		}
		st.Created += created
		st.Updated += updated
	}
	for w := range changedWeeks {
		st.Weeks = append(st.Weeks, w)
	}
	sort.Slice(st.Weeks, func(i, j int) bool { return st.Weeks[i].Start().Before(st.Weeks[j].Start()) })
	return st, nil
}

// merge upserts fresh into existing by Key, preserving FirstSeenAt.
func merge(existing, fresh []model.Incident, now time.Time) (out []model.Incident, created, updated int) {
	byKey := make(map[string]model.Incident, len(existing)+len(fresh))
	for _, in := range existing {
		byKey[in.Key] = in
	}
	for _, in := range fresh {
		if prev, ok := byKey[in.Key]; ok {
			in.FirstSeenAt = prev.FirstSeenAt
			if sameIncident(prev, in) {
				continue
			}
			byKey[in.Key] = in
			updated++
			continue
		}
		in.FirstSeenAt = now
		byKey[in.Key] = in
		created++
	}

	out = make([]model.Incident, 0, len(byKey))
	for _, in := range byKey {
		out = append(out, in)
	}
	sortIncidents(out)
	return out, created, updated
}

// sameIncident reports whether two records would serialize identically.
//
// Comparing the encoded bytes rather than the structs is deliberate: it asks
// exactly the question that matters, which is whether writing this record would
// change the file. time.Time values that mean the same instant can differ
// structurally, so a struct comparison would report phantom changes.
func sameIncident(a, b model.Incident) bool {
	ab, err := json.Marshal(a)
	if err != nil {
		return false
	}
	bb, err := json.Marshal(b)
	if err != nil {
		return false
	}
	return bytes.Equal(ab, bb)
}

// sortIncidents imposes the archive's file order: oldest first, key as the
// tiebreaker. Without a total order the map iteration above would reshuffle
// every file on every run.
func sortIncidents(in []model.Incident) {
	sort.Slice(in, func(i, j int) bool {
		if !in[i].StartedAt.Equal(in[j].StartedAt) {
			return in[i].StartedAt.Before(in[j].StartedAt)
		}
		return in[i].Key < in[j].Key
	})
}

func (a *Archive) vendorPath(w Week, vendor string) string {
	return filepath.Join(a.Root, w.Dir(), vendor+".json")
}

// LoadVendorWeek reads one vendor's file. A missing file is not an error: it
// simply means nothing was recorded for that vendor that week.
func (a *Archive) LoadVendorWeek(w Week, vendor string) ([]model.Incident, error) {
	b, err := os.ReadFile(a.vendorPath(w, vendor))
	if errors.Is(err, fs.ErrNotExist) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("read %s %s: %w", w, vendor, err)
	}
	var out []model.Incident
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, fmt.Errorf("decode %s %s: %w", w, vendor, err)
	}
	return out, nil
}

// SaveVendorWeek writes one vendor's file.
func (a *Archive) SaveVendorWeek(w Week, vendor string, incidents []model.Incident) error {
	sortIncidents(incidents)
	return writeJSON(a.vendorPath(w, vendor), incidents)
}

// LoadWeek selects semantic membership from the complete archive, independent
// of physical paths. Legacy copies of an identical source key count once.
func (a *Archive) LoadWeek(w Week) ([]model.Incident, error) {
	inv, err := a.ReadAll()
	if errors.Is(err, fs.ErrNotExist) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var out []model.Incident
	seen := map[string]model.Incident{}
	for _, r := range inv.Records {
		if old, ok := seen[r.Incident.Key]; ok {
			_, oldHash, oldErr := snapshotOf(old)
			_, newHash, newErr := snapshotOf(r.Incident)
			if oldErr != nil || newErr != nil || oldHash != newHash {
				return nil, fmt.Errorf("conflicting records for %s", r.Incident.Key)
			}
			continue
		}
		seen[r.Incident.Key] = r.Incident
		if w.Contains(r.Incident.StartedAt) {
			out = append(out, r.Incident)
		}
	}
	sortIncidents(out)
	return out, nil
}

const indexFile = "index.json"

// WriteIndex stores the week summary the renderers read and returns whatever
// ended up on disk.
//
// The file is left alone when nothing but GeneratedAt would change. Without
// that guard a re-render of an unchanged week rewrites index.json with a new
// timestamp, and the workflow commits a diff containing no facts — exactly the
// churn the archive is built to avoid. The caller gets the stored index back so
// the rendered digest quotes the timestamp that is actually on disk.
func (a *Archive) WriteIndex(w Week, idx model.WeekIndex) (model.WeekIndex, error) {
	if prev, err := a.ReadIndex(w); err == nil && sameIndexContent(prev, idx) {
		return prev, nil
	}
	if err := writeJSON(filepath.Join(a.Root, w.Dir(), indexFile), idx); err != nil {
		return model.WeekIndex{}, err
	}
	return idx, nil
}

// sameIndexContent compares two summaries while ignoring when they were built.
func sameIndexContent(a, b model.WeekIndex) bool {
	a.GeneratedAt = time.Time{}
	b.GeneratedAt = time.Time{}
	ab, err := json.Marshal(a)
	if err != nil {
		return false
	}
	bb, err := json.Marshal(b)
	if err != nil {
		return false
	}
	return bytes.Equal(ab, bb)
}

// ReadIndex loads the week summary.
func (a *Archive) ReadIndex(w Week) (model.WeekIndex, error) {
	var idx model.WeekIndex
	b, err := os.ReadFile(filepath.Join(a.Root, w.Dir(), indexFile))
	if err != nil {
		return idx, fmt.Errorf("read index %s: %w", w, err)
	}
	if err := json.Unmarshal(b, &idx); err != nil {
		return idx, fmt.Errorf("decode index %s: %w", w, err)
	}
	return idx, nil
}

// State is the run-to-run bookkeeping kept at the archive root.
//
// It is what makes gap detection possible: feed depth is finite and there is no
// pagination, so a collector that stays broken for long enough loses data
// permanently and nothing in the archive itself would reveal the hole.
type State struct {
	LastCollectAt time.Time            `json:"last_collect_at"`
	Sources       map[string]time.Time `json:"source_last_success"`
}

const stateFile = "state.json"

// LoadState reads state.json, returning a zero State when the archive is new.
func (a *Archive) LoadState() (State, error) {
	var s State
	b, err := os.ReadFile(filepath.Join(a.Root, stateFile))
	if errors.Is(err, fs.ErrNotExist) {
		return State{Sources: map[string]time.Time{}}, nil
	}
	if err != nil {
		return s, fmt.Errorf("read state: %w", err)
	}
	if err := json.Unmarshal(b, &s); err != nil {
		return s, fmt.Errorf("decode state: %w", err)
	}
	if s.Sources == nil {
		s.Sources = map[string]time.Time{}
	}
	return s, nil
}

// SaveState writes state.json.
func (a *Archive) SaveState(s State) error {
	return writeJSON(filepath.Join(a.Root, stateFile), s)
}

// writeJSON writes indented JSON with a trailing newline, creating parents as
// needed. Indented and key-sorted so that a diff of the archive is readable and
// so that identical facts produce identical bytes.
func writeJSON(path string, v any) error { return storage.WriteJSON(path, v) }
