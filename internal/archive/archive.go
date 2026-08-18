// Package archive stores canonical incidents on disk, one JSON file per vendor
// per ISO week, and merges new observations into what is already there.
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
	"strings"
	"time"

	"github.com/jtprogru/weekly-incident/internal/model"
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
// Incidents are routed by StartedAt, so a run also writes into past weeks: a
// vendor that appends a postmortem two weeks late updates the file of the week
// the incident began in.
func (a *Archive) Upsert(incidents []model.Incident, now time.Time) (Stats, error) {
	now = now.UTC()

	type bucket struct {
		week   Week
		vendor string
	}
	grouped := make(map[bucket][]model.Incident)
	for _, in := range incidents {
		b := bucket{week: WeekOf(in.StartedAt), vendor: in.Vendor}
		grouped[b] = append(grouped[b], in)
	}

	buckets := make([]bucket, 0, len(grouped))
	for b := range grouped {
		buckets = append(buckets, b)
	}
	sort.Slice(buckets, func(i, j int) bool {
		if buckets[i].week != buckets[j].week {
			return buckets[i].week.Start().Before(buckets[j].week.Start())
		}
		return buckets[i].vendor < buckets[j].vendor
	})

	var st Stats
	weeks := make(map[Week]bool)
	for _, b := range buckets {
		existing, err := a.LoadVendorWeek(b.week, b.vendor)
		if err != nil {
			return st, err
		}
		merged, created, updated := merge(existing, grouped[b], now)
		// Seeing an incident again is not a change. Writing the file anyway
		// would rewrite every week the feeds still reach — 72 of them on the
		// first run here — on every single daily run.
		if created == 0 && updated == 0 {
			continue
		}
		if err := a.SaveVendorWeek(b.week, b.vendor, merged); err != nil {
			return st, err
		}
		st.Created += created
		st.Updated += updated
		weeks[b.week] = true
	}

	st.Weeks = make([]Week, 0, len(weeks))
	for w := range weeks {
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

// LoadWeek reads every vendor file for a week, in vendor order.
func (a *Archive) LoadWeek(w Week) ([]model.Incident, error) {
	dir := filepath.Join(a.Root, w.Dir())
	entries, err := os.ReadDir(dir)
	if errors.Is(err, fs.ErrNotExist) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("read week %s: %w", w, err)
	}

	names := make([]string, 0, len(entries))
	for _, e := range entries {
		name := e.Name()
		if e.IsDir() || !strings.HasSuffix(name, ".json") || name == indexFile {
			continue
		}
		names = append(names, strings.TrimSuffix(name, ".json"))
	}
	sort.Strings(names)

	var out []model.Incident
	for _, vendor := range names {
		in, err := a.LoadVendorWeek(w, vendor)
		if err != nil {
			return nil, err
		}
		out = append(out, in...)
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
func writeJSON(path string, v any) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return fmt.Errorf("mkdir %s: %w", filepath.Dir(path), err)
	}
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Errorf("encode %s: %w", path, err)
	}
	b = append(b, '\n')
	if err := os.WriteFile(path, b, 0o644); err != nil {
		return fmt.Errorf("write %s: %w", path, err)
	}
	return nil
}
