package archive

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"time"

	"github.com/jtprogru/weekly-incident/internal/model"
)

func TestGlobalUpsertMovesSemanticWeekAndKeepsVersions(t *testing.T) {
	a := New(t.TempDir())
	in := incident("github", "a", time.Date(2026, 12, 31, 10, 0, 0, 0, time.UTC))
	oldWeek := WeekOf(in.StartedAt)
	if _, err := a.Upsert([]model.Incident{in}, runAt); err != nil {
		t.Fatal(err)
	}
	in.StartedAt = in.StartedAt.Add(7 * 24 * time.Hour)
	in.Finalize()
	newWeek := WeekOf(in.StartedAt)
	st, err := a.Upsert([]model.Incident{in}, runAt)
	if err != nil {
		t.Fatal(err)
	}
	if st.Created != 0 || st.Updated != 1 || len(st.Weeks) != 2 {
		t.Fatalf("stats %+v", st)
	}
	old, err := a.LoadWeek(oldWeek)
	if err != nil {
		t.Fatal(err)
	}
	current, err := a.LoadWeek(newWeek)
	if err != nil {
		t.Fatal(err)
	}
	if len(old) != 0 || len(current) != 1 {
		t.Fatalf("semantic weeks old/new %d/%d", len(old), len(current))
	}
	if _, err := os.Stat(a.vendorPath(newWeek, "github")); !os.IsNotExist(err) {
		t.Fatal("created a second physical record")
	}
	o, err := a.ReadObservations()
	if err != nil {
		t.Fatal(err)
	}
	if len(o.Manifest.Records[in.Key].Revisions) != 2 {
		t.Fatal("lost revision")
	}
	before := o.Manifest
	if _, err := a.Upsert([]model.Incident{in}, runAt.Add(time.Hour)); err != nil {
		t.Fatal(err)
	}
	o, err = a.ReadObservations()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(before, o.Manifest) {
		t.Fatal("polling created revision")
	}
}

func TestObservationABAAndTampering(t *testing.T) {
	a := New(t.TempDir())
	in := incident("github", "a", weekOf)
	for _, title := range []string{"A", "B", "A"} {
		in.Title = title
		if _, err := a.Upsert([]model.Incident{in}, runAt); err != nil {
			t.Fatal(err)
		}
	}
	o, err := a.ReadObservations()
	if err != nil {
		t.Fatal(err)
	}
	h := o.Manifest.Records[in.Key]
	if len(h.Revisions) != 3 || len(o.Snapshots) != 2 || h.Revisions[0].Hash != h.Revisions[2].Hash {
		t.Fatalf("history %+v", h)
	}
	path := filepath.Join(a.Root, "observations", in.Vendor, h.Revisions[0].Hash+".json")
	if err := os.WriteFile(path, []byte("{}"), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := a.ReadObservations(); err == nil {
		t.Fatal("accepted damaged snapshot")
	}
}

func TestLegacyRevisionOrderIsUnknown(t *testing.T) {
	a := New(t.TempDir())
	in := incident("github", "a", weekOf)
	if err := a.SaveVendorWeek(WeekOf(weekOf), "github", []model.Incident{in}); err != nil {
		t.Fatal(err)
	}
	in.Title = "other version"
	if err := a.SaveVendorWeek(WeekOf(weekOf).Prev(), "github", []model.Incident{in}); err != nil {
		t.Fatal(err)
	}
	inv, err := a.ReadAll()
	if err != nil {
		t.Fatal(err)
	}
	o, err := a.PrepareObservations(inv, runAt)
	if err != nil {
		t.Fatal(err)
	}
	if !o.Manifest.Records[in.Key].UnknownCurrent {
		t.Fatal("invented legacy ordering")
	}
	if err := a.WriteObservations(o); err != nil {
		t.Fatal(err)
	}
	again, err := a.PrepareObservations(inv, runAt.Add(time.Hour))
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(o.Manifest, again.Manifest) {
		t.Fatal("legacy import added duplicate revisions")
	}
	// A fresh response resolves unknown ordering without discarding either
	// imported payload or leaving stale compatibility copies behind.
	in.Title = "known current"
	if _, err := a.Upsert([]model.Incident{in}, runAt.Add(time.Hour)); err != nil {
		t.Fatal(err)
	}
	current, err := a.ReadObservations()
	if err != nil {
		t.Fatal(err)
	}
	if current.Manifest.Records[in.Key].UnknownCurrent || len(current.Manifest.Records[in.Key].Revisions) != 3 {
		t.Fatal("fresh fetch did not establish current revision")
	}
	rows, err := a.LoadWeek(WeekOf(in.StartedAt))
	if err != nil || len(rows) != 1 || rows[0].Title != in.Title {
		t.Fatalf("legacy projections: %v %v", rows, err)
	}
	inv, err = a.ReadAll()
	if err != nil {
		t.Fatal(err)
	}
	if _, err := a.PrepareObservations(inv, runAt.Add(2*time.Hour)); err != nil {
		t.Fatal(err)
	}
}

func TestRetryAfterProjectionFailureDoesNotInventRevisions(t *testing.T) {
	a := New(t.TempDir())
	in := incident("github", "a", weekOf)
	if _, err := a.Upsert([]model.Incident{in}, runAt); err != nil {
		t.Fatal(err)
	}
	o, err := a.ReadObservations()
	if err != nil {
		t.Fatal(err)
	}
	in.Title = "B"
	if err := o.Observe([]StoredIncident{{Path: WeekOf(weekOf).Dir() + "/github.json", Incident: in}}, runAt, false); err != nil {
		t.Fatal(err)
	}
	// Simulate a crash after the evidence manifest but before vendor projection.
	if err := a.WriteObservations(o); err != nil {
		t.Fatal(err)
	}
	if _, err := a.Upsert([]model.Incident{in}, runAt); err != nil {
		t.Fatal(err)
	}
	o, err = a.ReadObservations()
	if err != nil {
		t.Fatal(err)
	}
	if len(o.Manifest.Records[in.Key].Revisions) != 2 {
		t.Fatal("retry invented A/B revisions")
	}
}
