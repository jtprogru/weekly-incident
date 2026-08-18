package archive

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/jtprogru/weekly-incident/internal/model"
)

var (
	runAt  = time.Date(2026, 8, 18, 6, 0, 0, 0, time.UTC)
	weekOf = time.Date(2026, 8, 18, 9, 30, 0, 0, time.UTC)
)

func incident(vendor, id string, started time.Time) model.Incident {
	in := model.Incident{
		Vendor:    vendor,
		NativeID:  id,
		Title:     "something broke",
		StartedAt: started,
		Raw:       json.RawMessage(`{"id":"` + id + `"}`),
	}
	in.Finalize()
	return in
}

func TestUpsertCreatesThenUpdates(t *testing.T) {
	a := New(t.TempDir())

	st, err := a.Upsert([]model.Incident{
		incident("github", "a", weekOf),
		incident("github", "b", weekOf),
		incident("aws", "c", weekOf),
	}, runAt)
	if err != nil {
		t.Fatalf("Upsert: %v", err)
	}
	if st.Created != 3 || st.Updated != 0 {
		t.Errorf("first run: created %d, updated %d; want 3 and 0", st.Created, st.Updated)
	}
	if len(st.Weeks) != 1 || st.Weeks[0].String() != "2026-W34" {
		t.Errorf("weeks touched = %v, want [2026-W34]", st.Weeks)
	}

	// Vendors get their own file, so one vendor's failure cannot corrupt
	// another's data and the diff stays readable.
	for _, vendor := range []string{"github", "aws"} {
		path := filepath.Join(a.Root, "2026", "W34", vendor+".json")
		if _, err := os.Stat(path); err != nil {
			t.Errorf("expected %s: %v", path, err)
		}
	}

	updated := incident("github", "a", weekOf)
	updated.Title = "something broke, then was explained"
	st, err = a.Upsert([]model.Incident{updated}, runAt.Add(24*time.Hour))
	if err != nil {
		t.Fatalf("Upsert: %v", err)
	}
	if st.Created != 0 || st.Updated != 1 {
		t.Errorf("second run: created %d, updated %d; want 0 and 1", st.Created, st.Updated)
	}

	got, err := a.LoadVendorWeek(WeekOf(weekOf), "github")
	if err != nil {
		t.Fatalf("LoadVendorWeek: %v", err)
	}
	if len(got) != 2 {
		t.Fatalf("got %d incidents, want 2", len(got))
	}
	for _, in := range got {
		if in.Key != "github/a" {
			continue
		}
		if in.Title != "something broke, then was explained" {
			t.Errorf("Title = %q, want the updated one", in.Title)
		}
		// FirstSeenAt records when this archive first saw the incident and must
		// survive every later update.
		if !in.FirstSeenAt.Equal(runAt) {
			t.Errorf("FirstSeenAt = %s, want the original %s", in.FirstSeenAt, runAt)
		}
	}
}

// TestUpsertIsIdempotent is the property the daily cron depends on: identical
// facts must produce identical bytes, or every run commits noise.
func TestUpsertIsIdempotent(t *testing.T) {
	incidents := []model.Incident{
		incident("github", "a", weekOf),
		incident("github", "b", weekOf.Add(-3*time.Hour)),
		incident("cloudflare", "c", weekOf),
	}

	a := New(t.TempDir())
	if _, err := a.Upsert(incidents, runAt); err != nil {
		t.Fatalf("Upsert: %v", err)
	}
	first := snapshot(t, a.Root)

	// A later run, on a different clock, seeing exactly the same feed content.
	for i := range 3 {
		if _, err := a.Upsert(incidents, runAt.Add(time.Duration(i+1)*24*time.Hour)); err != nil {
			t.Fatalf("Upsert: %v", err)
		}
		if got := snapshot(t, a.Root); got != first {
			t.Fatalf("run %d rewrote the archive with no new facts", i+2)
		}
	}
}

func TestUpsertRoutesByStartWeek(t *testing.T) {
	a := New(t.TempDir())
	// A postmortem published this week for an incident that began last week
	// belongs to last week's file.
	lastWeek := weekOf.AddDate(0, 0, -7)
	if _, err := a.Upsert([]model.Incident{
		incident("gcp", "old", lastWeek),
		incident("gcp", "new", weekOf),
	}, runAt); err != nil {
		t.Fatalf("Upsert: %v", err)
	}

	for week, wantKey := range map[string]string{"2026-W33": "gcp/old", "2026-W34": "gcp/new"} {
		w, err := ParseWeek(week)
		if err != nil {
			t.Fatal(err)
		}
		got, err := a.LoadVendorWeek(w, "gcp")
		if err != nil {
			t.Fatalf("LoadVendorWeek %s: %v", week, err)
		}
		if len(got) != 1 || got[0].Key != wantKey {
			t.Errorf("%s holds %v, want just %s", week, keys(got), wantKey)
		}
	}
}

func TestLoadWeekSkipsIndex(t *testing.T) {
	a := New(t.TempDir())
	w := WeekOf(weekOf)
	if _, err := a.Upsert([]model.Incident{incident("slack", "1", weekOf)}, runAt); err != nil {
		t.Fatalf("Upsert: %v", err)
	}
	if _, err := a.WriteIndex(w, model.WeekIndex{Week: w.String()}); err != nil {
		t.Fatalf("WriteIndex: %v", err)
	}

	got, err := a.LoadWeek(w)
	if err != nil {
		t.Fatalf("LoadWeek: %v", err)
	}
	if len(got) != 1 {
		t.Errorf("LoadWeek returned %d incidents, want 1: index.json must not be read as a vendor file", len(got))
	}
}

func TestLoadMissingWeekIsEmpty(t *testing.T) {
	a := New(t.TempDir())
	got, err := a.LoadWeek(WeekOf(weekOf))
	if err != nil {
		t.Errorf("LoadWeek on an empty archive: %v", err)
	}
	if len(got) != 0 {
		t.Errorf("got %d incidents from an empty archive", len(got))
	}
}

func TestState(t *testing.T) {
	a := New(t.TempDir())

	s, err := a.LoadState()
	if err != nil {
		t.Fatalf("LoadState on a new archive: %v", err)
	}
	if !s.LastCollectAt.IsZero() || s.Sources == nil {
		t.Errorf("fresh state = %+v, want a zero time and an initialized map", s)
	}

	s.LastCollectAt = runAt
	s.Sources["github"] = runAt
	if err := a.SaveState(s); err != nil {
		t.Fatalf("SaveState: %v", err)
	}

	back, err := a.LoadState()
	if err != nil {
		t.Fatalf("LoadState: %v", err)
	}
	if !back.LastCollectAt.Equal(runAt) || !back.Sources["github"].Equal(runAt) {
		t.Errorf("state did not round-trip: %+v", back)
	}
}

func TestFilesEndWithNewline(t *testing.T) {
	a := New(t.TempDir())
	if _, err := a.Upsert([]model.Incident{incident("github", "a", weekOf)}, runAt); err != nil {
		t.Fatalf("Upsert: %v", err)
	}
	b, err := os.ReadFile(filepath.Join(a.Root, "2026", "W34", "github.json"))
	if err != nil {
		t.Fatal(err)
	}
	if len(b) == 0 || b[len(b)-1] != '\n' {
		t.Error("archive file does not end with a newline")
	}
}

// snapshot concatenates every file in the archive so two runs can be compared
// byte for byte.
func snapshot(t *testing.T, root string) string {
	t.Helper()
	var out string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}
		b, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		out += path + "\n" + string(b) + "\n"
		return nil
	})
	if err != nil {
		t.Fatalf("snapshot: %v", err)
	}
	return out
}

func keys(in []model.Incident) []string {
	out := make([]string, len(in))
	for i, v := range in {
		out[i] = v.Key
	}
	return out
}

// TestUpsertSkipsUnchangedWeeks is what keeps the daily cron from rewriting the
// entire archive. The feeds reach back years, so a run that reports every week
// it saw as "touched" would churn dozens of files a day for nothing.
func TestUpsertSkipsUnchangedWeeks(t *testing.T) {
	a := New(t.TempDir())
	incidents := []model.Incident{incident("github", "a", weekOf)}

	if _, err := a.Upsert(incidents, runAt); err != nil {
		t.Fatalf("Upsert: %v", err)
	}
	path := filepath.Join(a.Root, "2026", "W34", "github.json")
	before, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}

	st, err := a.Upsert(incidents, runAt.Add(24*time.Hour))
	if err != nil {
		t.Fatalf("Upsert: %v", err)
	}
	if st.Created != 0 || st.Updated != 0 {
		t.Errorf("re-seeing the same incident reported %d new and %d updated, want 0 and 0",
			st.Created, st.Updated)
	}
	if len(st.Weeks) != 0 {
		t.Errorf("weeks touched = %v, want none", st.Weeks)
	}

	after, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	if !after.ModTime().Equal(before.ModTime()) {
		t.Error("the vendor file was rewritten despite nothing changing")
	}
}

// TestWriteIndexIgnoresTheClock guards the second half of the no-churn rule.
// index.json carries a generation timestamp, so writing it unconditionally
// would commit a diff of nothing but that timestamp every time a settled week
// is re-rendered.
func TestWriteIndexIgnoresTheClock(t *testing.T) {
	a := New(t.TempDir())
	w := WeekOf(weekOf)
	idx := model.WeekIndex{
		Week:        w.String(),
		From:        w.Start(),
		To:          w.End(),
		GeneratedAt: runAt,
		Incidents:   []model.IndexEntry{{Key: "github/a", Vendor: "github"}},
	}

	stored, err := a.WriteIndex(w, idx)
	if err != nil {
		t.Fatalf("WriteIndex: %v", err)
	}
	if !stored.GeneratedAt.Equal(runAt) {
		t.Fatalf("first write stored %s, want %s", stored.GeneratedAt, runAt)
	}
	path := filepath.Join(a.Root, w.Dir(), "index.json")
	before, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}

	later := idx
	later.GeneratedAt = runAt.Add(48 * time.Hour)
	stored, err = a.WriteIndex(w, later)
	if err != nil {
		t.Fatalf("WriteIndex: %v", err)
	}
	if !stored.GeneratedAt.Equal(runAt) {
		t.Errorf("stored GeneratedAt = %s, want the original %s: the caller must quote the file on disk",
			stored.GeneratedAt, runAt)
	}
	after, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	if !after.ModTime().Equal(before.ModTime()) {
		t.Error("index.json was rewritten when only its timestamp changed")
	}

	// A real change still lands.
	changed := later
	changed.Incidents = append(changed.Incidents, model.IndexEntry{Key: "aws/b", Vendor: "aws"})
	stored, err = a.WriteIndex(w, changed)
	if err != nil {
		t.Fatalf("WriteIndex: %v", err)
	}
	if len(stored.Incidents) != 2 || !stored.GeneratedAt.Equal(later.GeneratedAt) {
		t.Error("a genuine change did not reach the file")
	}
}
