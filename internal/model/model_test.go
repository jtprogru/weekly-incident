package model

import (
	"testing"
	"time"
)

func TestFinalizeComputesDuration(t *testing.T) {
	start := time.Date(2026, 8, 18, 7, 40, 0, 0, time.UTC)
	end := start.Add(242 * time.Minute)

	in := Incident{Vendor: "github", NativeID: "abc", StartedAt: start, ResolvedAt: &end}
	in.Finalize()

	if in.Key != "github/abc" {
		t.Errorf("Key = %q", in.Key)
	}
	if in.DurationMinutes == nil || *in.DurationMinutes != 242 {
		t.Errorf("DurationMinutes = %v, want 242", in.DurationMinutes)
	}
	// Nil slices would serialize as null and break every consumer that expects
	// an array.
	if in.Components == nil || in.Updates == nil {
		t.Error("Finalize left a nil slice")
	}
}

func TestFinalizeLeavesOngoingDurationNil(t *testing.T) {
	// "Ran for zero minutes" and "still running" are different facts.
	in := Incident{Vendor: "slack", NativeID: "1", StartedAt: time.Now().UTC()}
	in.Finalize()
	if in.DurationMinutes != nil {
		t.Errorf("DurationMinutes = %v, want nil", in.DurationMinutes)
	}
}

func TestDurationAtMeasuresOngoingAgainstNow(t *testing.T) {
	start := time.Date(2026, 8, 18, 0, 0, 0, 0, time.UTC)
	now := start.Add(90 * time.Minute)

	in := Incident{Vendor: "aws", NativeID: "x", StartedAt: start}
	in.Finalize()
	d, ongoing := in.DurationAt(now)
	if !ongoing {
		t.Error("ongoing = false, want true")
	}
	if d != 90 {
		t.Errorf("duration = %d, want 90", d)
	}

	end := start.Add(30 * time.Minute)
	in.ResolvedAt = &end
	in.Finalize()
	d, ongoing = in.DurationAt(now)
	if ongoing {
		t.Error("ongoing = true on a resolved incident")
	}
	if d != 30 {
		t.Errorf("duration = %d, want 30", d)
	}
}

func TestValidate(t *testing.T) {
	now := time.Now().UTC()
	ok := Incident{Vendor: "v", NativeID: "1", StartedAt: now}
	if err := ok.Validate(); err != nil {
		t.Errorf("Validate on a good incident: %v", err)
	}
	for name, in := range map[string]Incident{
		"no vendor":    {NativeID: "1", StartedAt: now},
		"no native id": {Vendor: "v", StartedAt: now},
		"no start":     {Vendor: "v", NativeID: "1"},
	} {
		if err := in.Validate(); err == nil {
			t.Errorf("Validate accepted an incident with %s", name)
		}
	}
}

func TestImpactAtLeast(t *testing.T) {
	if !ImpactCritical.AtLeast(ImpactMajor) {
		t.Error("critical should clear a major threshold")
	}
	if ImpactMinor.AtLeast(ImpactMajor) {
		t.Error("minor should not clear a major threshold")
	}
	// Slack and AWS never report impact. Treating unknown as severe would flood
	// the digest with every short blip those two publish.
	if ImpactUnknown.AtLeast(ImpactNone) {
		t.Error("unknown should rank below none")
	}
}

func TestParseImpactAndStatus(t *testing.T) {
	if got := ParseImpact("critical"); got != ImpactCritical {
		t.Errorf("ParseImpact(critical) = %q", got)
	}
	if got := ParseImpact("catastrophic"); got != ImpactUnknown {
		t.Errorf("ParseImpact of an unknown value = %q, want unknown", got)
	}
	if got := ParseStatus("postmortem"); got != StatusPostmortem {
		t.Errorf("ParseStatus(postmortem) = %q", got)
	}
	if got := ParseStatus("whatever"); got != StatusUnknown {
		t.Errorf("ParseStatus of an unknown value = %q, want unknown", got)
	}
}
