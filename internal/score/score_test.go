package score

import (
	"testing"
	"time"

	"github.com/jtprogru/weekly-incident/internal/archive"
	"github.com/jtprogru/weekly-incident/internal/config"
	"github.com/jtprogru/weekly-incident/internal/model"
)

var now = time.Date(2026, 8, 24, 7, 0, 0, 0, time.UTC)

func resolved(vendor string, start time.Time, minutes int, components []string, impact model.Impact) model.Incident {
	end := start.Add(time.Duration(minutes) * time.Minute)
	in := model.Incident{
		Vendor: vendor, NativeID: "x" + vendor, Title: vendor + " incident",
		Impact: impact, Status: model.StatusResolved,
		StartedAt: start, ResolvedAt: &end, Components: components,
	}
	in.Finalize()
	return in
}

func TestFor(t *testing.T) {
	start := time.Date(2026, 8, 18, 0, 0, 0, 0, time.UTC)
	s := For(resolved("aws", start, 4338, []string{"a", "b"}, model.ImpactUnknown), 2.0, now)
	if s.Duration != 4338 || s.Components != 2 || s.Weight != 2.0 {
		t.Fatalf("multipliers = %+v", s)
	}
	if s.Value != 4338*2*2.0 {
		t.Errorf("Value = %v, want %v", s.Value, 4338*2*2.0)
	}
}

func TestForHandlesMissingMultipliers(t *testing.T) {
	start := time.Date(2026, 8, 18, 0, 0, 0, 0, time.UTC)
	// A vendor that lists no components has still broken something, so the
	// count floors at one instead of zeroing the whole score.
	s := For(resolved("slack", start, 60, nil, model.ImpactUnknown), 0, now)
	if s.Components != 1 {
		t.Errorf("Components = %d, want the floor of 1", s.Components)
	}
	if s.Weight != 1 {
		t.Errorf("Weight = %v, want the default 1", s.Weight)
	}
	if s.Value != 60 {
		t.Errorf("Value = %v, want 60", s.Value)
	}
}

func TestForMeasuresOngoingAgainstNow(t *testing.T) {
	in := model.Incident{
		Vendor: "aws", NativeID: "open", StartedAt: now.Add(-90 * time.Minute),
		Components: []string{"one"},
	}
	in.Finalize()
	if s := For(in, 1, now); s.Duration != 90 {
		t.Errorf("Duration = %d, want 90", s.Duration)
	}
}

func TestAboveThresholdIsADisjunction(t *testing.T) {
	d := config.Digest{MinDurationMinutes: 30, MinImpact: model.ImpactMajor}
	start := time.Date(2026, 8, 18, 0, 0, 0, 0, time.UTC)

	cases := []struct {
		name    string
		minutes int
		impact  model.Impact
		want    bool
	}{
		{"long and minor clears on duration", 240, model.ImpactMinor, true},
		{"short and critical clears on impact", 4, model.ImpactCritical, true},
		{"short and minor stays out", 4, model.ImpactMinor, false},
		{"exactly at the duration bound clears", 30, model.ImpactMinor, true},
		// Slack and AWS report no impact at all; they can only clear on
		// duration, which is what makes the impact rank of unknown matter.
		{"short and unknown stays out", 4, model.ImpactUnknown, false},
		{"long and unknown clears", 60, model.ImpactUnknown, true},
	}
	for _, c := range cases {
		in := resolved("v", start, c.minutes, []string{"a"}, c.impact)
		if got := AboveThreshold(in, d, now); got != c.want {
			t.Errorf("%s: got %v, want %v", c.name, got, c.want)
		}
	}
}

func TestBuildIndexSortsByScoreDescending(t *testing.T) {
	start := time.Date(2026, 8, 18, 0, 0, 0, 0, time.UTC)
	cfg := &config.Config{
		Digest: config.Digest{MinDurationMinutes: 30, MinImpact: model.ImpactMajor},
		Sources: []config.Source{
			{Vendor: "aws", Weight: 2.0},
			{Vendor: "github", Weight: 1.5},
			{Vendor: "digitalocean", Weight: 0.8},
		},
	}
	incidents := []model.Incident{
		resolved("digitalocean", start, 600, []string{"a"}, model.ImpactMinor), // 480
		resolved("aws", start, 300, []string{"a", "b"}, model.ImpactUnknown),   // 1200
		resolved("github", start, 10, []string{"a"}, model.ImpactMinor),        // 15, below
	}

	w := archive.WeekOf(start)
	idx := BuildIndex(w, incidents, []model.SourceStatus{{Vendor: "zzz"}, {Vendor: "aaa"}}, cfg, now)

	if idx.Week != "2026-W34" {
		t.Errorf("Week = %q", idx.Week)
	}
	if len(idx.Incidents) != 3 {
		t.Fatalf("got %d entries, want 3", len(idx.Incidents))
	}
	if idx.Incidents[0].Vendor != "aws" || idx.Incidents[1].Vendor != "digitalocean" {
		t.Errorf("order = %s, %s; want aws then digitalocean",
			idx.Incidents[0].Vendor, idx.Incidents[1].Vendor)
	}
	if idx.Incidents[2].AboveThreshold {
		t.Error("the 10-minute minor incident should be below the threshold")
	}
	if idx.Sources[0].Vendor != "aaa" {
		t.Error("sources are not sorted by vendor")
	}
	if !idx.From.Equal(w.Start()) || !idx.To.Equal(w.End()) {
		t.Errorf("window = %s..%s, want the ISO week bounds", idx.From, idx.To)
	}
}
