package canonical

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"slices"
	"strings"
	"testing"
	"time"

	"github.com/jtprogru/weekly-incident/internal/archive"
	"github.com/jtprogru/weekly-incident/internal/config"
	"github.com/jtprogru/weekly-incident/internal/model"
	"github.com/jtprogru/weekly-incident/internal/storage"
)

var testNow = time.Date(2026, 11, 14, 0, 0, 0, 0, time.UTC)

func setupAWS(t *testing.T) (string, *config.Config, string) {
	t.Helper()
	root := t.TempDir()
	data := filepath.Join(root, "data")
	b, err := os.ReadFile("../../testdata/dedup/aws-window.json")
	if err != nil {
		t.Fatal(err)
	}
	if err := storage.Write(filepath.Join(data, "2026/W09/aws.json"), b); err != nil {
		t.Fatal(err)
	}
	dp := filepath.Join(root, "decisions.yaml")
	if err := storage.Write(dp, []byte("schema_version: 1\ndecisions: []\n")); err != nil {
		t.Fatal(err)
	}
	cfg, err := config.Parse([]byte("dedup:\n  mode: enforce\n  auto_aws: true\nsources:\n  - vendor: aws\n    kind: aws-rss\n    url: https://offline.invalid\n"))
	if err != nil {
		t.Fatal(err)
	}
	return data, cfg, dp
}

func TestAWSPlanApplyRebuildAndEvidence(t *testing.T) {
	root, cfg, dp := setupAWS(t)
	before, err := os.ReadFile(filepath.Join(root, "2026/W09/aws.json"))
	if err != nil {
		t.Fatal(err)
	}
	p, err := Prepare(root, cfg, dp, testNow, true, true)
	if err != nil {
		t.Fatal(err)
	}
	if p.Plan.Statistics.Incidents != 1 || p.Plan.Statistics.SourceRecords != 2 || len(p.Plan.Conflicts) != 0 {
		t.Fatalf("plan %+v conflicts %v", p.Plan.Statistics, p.Plan.Conflicts)
	}
	if len(p.Incidents[0].Fields.Updates) != 22 || len(p.Incidents[0].Sources) != 2 || p.Incidents[0].StartQuality != "first_observed_update" {
		t.Fatal("canonical evidence incomplete")
	}
	for _, u := range p.Incidents[0].Fields.Updates {
		if u.NativeID == "" {
			t.Fatal("historical RSS GUID not backfilled")
		}
	}
	if _, err := os.Stat(filepath.Join(root, "observations")); !os.IsNotExist(err) {
		t.Fatal("plan mutated archive")
	}
	b, err := storage.Encode(p.Plan)
	if err != nil {
		t.Fatal(err)
	}
	var fromJSON Plan
	if err := storage.Decode(b, &fromJSON); err != nil {
		t.Fatal(err)
	}
	if err := Apply(root, cfg, dp, fromJSON); err != nil {
		t.Fatal(err)
	}
	loaded, err := Load(root, cfg, dp)
	if err != nil {
		t.Fatal(err)
	}
	if loaded.Incidents[0].ID != p.Incidents[0].ID {
		t.Fatal("identity changed")
	}
	after, err := os.ReadFile(filepath.Join(root, "2026/W09/aws.json"))
	if err != nil {
		t.Fatal(err)
	}
	if string(before) != string(after) {
		t.Fatal("source archive changed on apply")
	}
	stable := hashTree(t, root)
	for range 2 {
		again, err := Prepare(root, cfg, dp, testNow, false, false)
		if err != nil {
			t.Fatal(err)
		}
		if err := Commit(root, again, cfg, nil, nil); err != nil {
			t.Fatal(err)
		}
	}
	if !reflect.DeepEqual(stable, hashTree(t, root)) {
		t.Fatal("rebuild changed identical output")
	}
	weeks := filepath.Join(filepath.Dir(root), "weeks")
	w := archive.Week{Year: 2026, Num: 9}
	if err := RenderWeek(root, weeks, loaded, cfg, w, testNow); err != nil {
		t.Fatal(err)
	}
	md, err := os.ReadFile(filepath.Join(weeks, w.String(), "digest.md"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(md), "2 source records") || !strings.Contains(string(md), loaded.Incidents[0].ID) {
		t.Fatal("render lost provenance")
	}
	if _, err := Load(root, cfg, dp); err != nil {
		t.Fatal(err)
	}
}

func TestPlanStalenessAndInterruptedBuildRecovery(t *testing.T) {
	root, cfg, dp := setupAWS(t)
	p, err := Prepare(root, cfg, dp, testNow, true, true)
	if err != nil {
		t.Fatal(err)
	}
	edited := p.Plan
	edited.Statistics.Incidents++
	if err := Apply(root, cfg, dp, edited); err == nil {
		t.Fatal("edited plan accepted")
	}
	if err := Apply(root, cfg, dp, p.Plan); err != nil {
		t.Fatal(err)
	}
	if err := Apply(root, cfg, dp, p.Plan); err == nil {
		t.Fatal("stale plan accepted")
	}
	path := filepath.Join(root, "2026/W09/index.json")
	if err := storage.Write(path, []byte("interrupted write")); err != nil {
		t.Fatal(err)
	}
	if _, err := Load(root, cfg, dp); err == nil {
		t.Fatal("corrupt derived artifact accepted")
	}
	repair, err := Prepare(root, cfg, dp, testNow, false, false)
	if err != nil {
		t.Fatal(err)
	}
	if err := Commit(root, repair, cfg, nil, nil); err != nil {
		t.Fatal(err)
	}
	if _, err := Load(root, cfg, dp); err != nil {
		t.Fatal(err)
	}
}

func TestAWSWindowShrinkAndNewResolution(t *testing.T) {
	root, cfg, dp := setupAWS(t)
	p, err := Prepare(root, cfg, dp, testNow, true, true)
	if err != nil {
		t.Fatal(err)
	}
	if err := Apply(root, cfg, dp, p.Plan); err != nil {
		t.Fatal(err)
	}
	key := p.Registry.Groups[0].Members[0]
	h := p.Observations.Manifest.Records[key]
	in := p.Observations.Snapshots[h.Revisions[len(h.Revisions)-1].Hash].Record
	in.Updates = slices.Clone(in.Updates[len(in.Updates)-2:])
	last := in.Updates[len(in.Updates)-1].CreatedAt.Add(time.Hour)
	in.Updates = append(in.Updates, model.Update{NativeID: "new-resolution", CreatedAt: last, Status: model.StatusResolved, Body: "The underlying service has fully recovered."})
	in.ResolvedAt = &last
	in.Status = model.StatusResolved
	in.Finalize()
	if _, err := archive.New(root).Upsert([]model.Incident{in}, testNow.Add(time.Hour)); err != nil {
		t.Fatal(err)
	}
	if _, err := Load(root, cfg, dp); err == nil {
		t.Fatal("stale canonical view accepted")
	}
	next, err := Prepare(root, cfg, dp, testNow.Add(time.Hour), true, false)
	if err != nil {
		t.Fatal(err)
	}
	if len(next.Incidents) != 1 || next.Incidents[0].ID != p.Incidents[0].ID || len(next.Incidents[0].Fields.Updates) != 23 || next.Incidents[0].Fields.ResolvedAt == nil {
		t.Fatalf("lost history/resolution: %+v", next.Plan.Statistics)
	}
	if err := Commit(root, next, cfg, nil, nil); err != nil {
		t.Fatal(err)
	}
}

func TestStableSourceCorrectionsAndWeekMigration(t *testing.T) {
	_, cfg, dp := setupAWS(t)
	// Use a fresh archive so this case contains a single stable source ID.
	root := filepath.Join(t.TempDir(), "data")
	cfg.Sources = []config.Source{{Vendor: "github", Kind: config.KindStatuspage, URL: "https://offline.invalid", Weight: 1}}
	in := model.Incident{Vendor: "github", NativeID: "one", Title: "incident", StartedAt: time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC), Raw: json.RawMessage(`{}`), Status: model.StatusInvestigating, Impact: model.ImpactMinor, Components: []string{"api"}, Updates: []model.Update{{NativeID: "update", CreatedAt: testNow, Body: "original", Status: model.StatusInvestigating}}}
	in.Finalize()
	a := archive.New(root)
	if _, err := a.Upsert([]model.Incident{in}, testNow); err != nil {
		t.Fatal(err)
	}
	p, err := Prepare(root, cfg, dp, testNow, true, true)
	if err != nil {
		t.Fatal(err)
	}
	if err := Apply(root, cfg, dp, p.Plan); err != nil {
		t.Fatal(err)
	}
	oldWeek := archive.WeekOf(in.StartedAt)
	in.StartedAt = in.StartedAt.Add(7 * 24 * time.Hour)
	in.Updates[0].Body = "corrected"
	in.Finalize()
	if _, err := a.Upsert([]model.Incident{in}, testNow); err != nil {
		t.Fatal(err)
	}
	next, err := Prepare(root, cfg, dp, testNow, true, false)
	if err != nil {
		t.Fatal(err)
	}
	if err := Commit(root, next, cfg, nil, nil); err != nil {
		t.Fatal(err)
	}
	if next.Incidents[0].ID != p.Incidents[0].ID || len(next.Incidents[0].Fields.Updates) != 1 || next.Incidents[0].Fields.Updates[0].Body != "corrected" || len(next.Incidents[0].UpdateHistory[0].Versions) != 2 {
		t.Fatal("update revisions or identity lost")
	}
	idx, err := a.ReadIndex(oldWeek)
	if err != nil {
		t.Fatal(err)
	}
	if len(idx.Incidents) != 0 {
		t.Fatal("old week kept incident")
	}
	if _, err := Load(root, cfg, dp); err != nil {
		t.Fatal(err)
	}
}

func TestConflictsRequirePinnedOverrides(t *testing.T) {
	root, cfg, dp := setupAWS(t)
	cfg.Dedup.AutoAWS = false
	// Merge two stable-source versions with conflicting start times.
	cfg.Sources = []config.Source{{Vendor: "aws", Kind: config.KindStatuspage, URL: "https://offline.invalid", Weight: 1}}
	inv, err := archive.New(root).ReadAll()
	if err != nil {
		t.Fatal(err)
	}
	a, b := inv.Records[0].Incident, inv.Records[1].Incident
	write := func(text string) {
		t.Helper()
		if err := storage.Write(dp, []byte(text)); err != nil {
			t.Fatal(err)
		}
	}
	base := fmtDecision("merge", "same_incident", []string{a.Key, b.Key}, "")
	write("schema_version: 1\ndecisions:\n" + base)
	p, err := Prepare(root, cfg, dp, testNow, true, true)
	if err != nil {
		t.Fatal(err)
	}
	if len(p.Plan.Conflicts) == 0 {
		t.Fatal("conflicting times silently selected")
	}
	if err := Apply(root, cfg, dp, p.Plan); err == nil {
		t.Fatal("applied conflicting group")
	}
	hashes := map[string]string{}
	for key, h := range p.Observations.Manifest.Records {
		hashes[key] = h.Revisions[len(h.Revisions)-1].Hash
	}
	extra := "    field: started_at\n    value: '" + a.StartedAt.Format(time.RFC3339) + "'\n    expected_hashes:\n      '" + a.Key + "': '" + hashes[a.Key] + "'\n      '" + b.Key + "': '" + hashes[b.Key] + "'\n"
	write("schema_version: 1\ndecisions:\n" + base + fmtDecision("start", "field_override", []string{a.Key}, extra))
	p, err = Prepare(root, cfg, dp, testNow, true, true)
	if err != nil {
		t.Fatal(err)
	}
	if len(p.Plan.Conflicts) != 0 {
		t.Fatalf("conflicts %v", p.Plan.Conflicts)
	}
	if err := Apply(root, cfg, dp, p.Plan); err != nil {
		t.Fatal(err)
	}
	a.Title = "later edit"
	if _, err := archive.New(root).Upsert([]model.Incident{a}, testNow); err != nil {
		t.Fatal(err)
	}
	if _, err := Prepare(root, cfg, dp, testNow, true, false); err == nil {
		t.Fatal("stale override accepted")
	}
}

func fmtDecision(id, kind string, members []string, extra string) string {
	return "  - id: " + id + "\n    type: " + kind + "\n    members: ['" + strings.Join(members, "', '") + "']\n    author: test fixture\n    at: 2026-11-14T00:00:00Z\n    reason: synthetic test\n    evidence: ['fixture']\n" + extra
}

func hashTree(t *testing.T, root string) map[string]string {
	t.Helper()
	out := map[string]string{}
	err := filepath.WalkDir(root, func(path string, e os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if e.IsDir() {
			return nil
		}
		b, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		out[rel] = storage.Hash(b)
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	return out
}
