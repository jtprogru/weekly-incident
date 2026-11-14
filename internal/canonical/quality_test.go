package canonical

import (
	"encoding/json"
	"path/filepath"
	"testing"
	"time"

	"github.com/jtprogru/weekly-incident/internal/archive"
	"github.com/jtprogru/weekly-incident/internal/config"
	"github.com/jtprogru/weekly-incident/internal/model"
	"github.com/jtprogru/weekly-incident/internal/score"
	"github.com/jtprogru/weekly-incident/internal/storage"
)

func TestUnknownDurationRequiresPinnedAcknowledgement(t *testing.T) {
	_, cfg, dp := setupAWS(t)
	root := filepath.Join(t.TempDir(), "data")
	cfg.Sources = []config.Source{{Vendor: "test", Kind: config.KindStatuspage, Weight: 1}}
	// A sub-minute negative interval must not disappear through integer rounding.
	end := testNow.Add(-time.Second)
	in := model.Incident{Vendor: "test", NativeID: "retroactive", Title: "Backdated report", StartedAt: testNow, ResolvedAt: &end, Status: model.StatusResolved, Impact: model.ImpactMinor, Raw: json.RawMessage(`{}`), Components: []string{"Global", "Global"}}
	in.Finalize()
	if _, err := archive.New(root).Upsert([]model.Incident{in}, testNow); err != nil {
		t.Fatal(err)
	}
	p, err := Prepare(root, cfg, dp, testNow, true, true)
	if err != nil {
		t.Fatal(err)
	}
	if len(p.Plan.Conflicts) != 1 {
		t.Fatal("negative interval was accepted")
	}
	extra := "    field: duration_unknown\n    value: true\n    expected_hashes:\n      '" + in.Key + "': '" + p.Plan.SourceHeads[in.Key] + "'\n"
	if err := storage.Write(dp, []byte("schema_version: 1\ndecisions:\n"+fmtDecision("unknown", "field_override", []string{in.Key}, extra))); err != nil {
		t.Fatal(err)
	}
	p, err = Prepare(root, cfg, dp, testNow, true, true)
	if err != nil {
		t.Fatal(err)
	}
	if len(p.Plan.Conflicts) != 0 || !p.Incidents[0].DurationUnknown || p.Incidents[0].Fields.DurationMinutes != nil {
		t.Fatal("unknown duration lost")
	}
	projection := p.Incidents[0].Projection()
	if duration, ongoing := projection.DurationAt(testNow.Add(time.Hour)); duration != 0 || ongoing {
		t.Fatal("closed event treated as ongoing")
	}
	if score.For(projection, 1, testNow).Components != 2 {
		t.Fatal("distinct components with equal display names collapsed")
	}
	if !projection.StartedAt.Equal(in.StartedAt) || !projection.ResolvedAt.Equal(end) {
		t.Fatal("vendor dates were fabricated")
	}
	if err := Apply(root, cfg, dp, p.Plan); err != nil {
		t.Fatal(err)
	}
}
