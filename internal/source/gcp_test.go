package source

import (
	"reflect"
	"testing"

	"github.com/jtprogru/weekly-incident/internal/model"
)

func TestGCPParse(t *testing.T) {
	g := &gcp{base: base{vendor: "gcp", url: "https://status.cloud.google.com/incidents.json"}}
	res, err := g.parse(golden(t, "gcp_incidents.json"))
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if res.ParseErrors != 0 {
		t.Errorf("ParseErrors = %d, want 0", res.ParseErrors)
	}
	if len(res.Incidents) != 2 {
		t.Fatalf("got %d incidents, want 2", len(res.Incidents))
	}

	in := res.Incidents[0]
	if in.Key != "gcp/3BvH3LVGcupoYqV6F4Nw" {
		t.Errorf("Key = %q", in.Key)
	}
	if in.URL != "https://status.cloud.google.com/incidents/3BvH3LVGcupoYqV6F4Nw" {
		t.Errorf("URL = %q, want the relative uri prefixed", in.URL)
	}
	// severity medium plus SERVICE_DISRUPTION folds to major.
	if in.Impact != model.ImpactMajor {
		t.Errorf("Impact = %q, want major", in.Impact)
	}
	if in.Status != model.StatusResolved {
		t.Errorf("Status = %q, want resolved (end is set)", in.Status)
	}
	// begin, not created: impact started 3.5 hours before Google opened the record.
	if want := mustTime(t, "2026-07-15T23:57:00Z"); !in.StartedAt.Equal(want) {
		t.Errorf("StartedAt = %s, want %s", in.StartedAt, want)
	}
	if in.DurationMinutes == nil || *in.DurationMinutes != 748 {
		t.Errorf("DurationMinutes = %v, want 748", in.DurationMinutes)
	}
	want := []string{"Bare Metal Solution", "Google Cloud NetApp Volumes", "VMWare engine"}
	if !reflect.DeepEqual(in.Components, want) {
		t.Errorf("Components = %v, want %v", in.Components, want)
	}
	if in.Title == "" || len([]rune(in.Title)) > gcpTitleLimit+1 {
		t.Errorf("Title = %q (%d runes), want a synthesized sentence within the limit", in.Title, len([]rune(in.Title)))
	}
	assertAscending(t, in.Updates)
}

func TestGCPImpact(t *testing.T) {
	cases := []struct {
		severity     string
		statusImpact string
		want         model.Impact
	}{
		{"high", "SERVICE_OUTAGE", model.ImpactCritical},
		{"medium", "SERVICE_DISRUPTION", model.ImpactMajor},
		{"low", "SERVICE_DISRUPTION", model.ImpactMinor},
		// An advisory entry is not an outage no matter what severity says.
		{"medium", "SERVICE_INFORMATION", model.ImpactNone},
		{"", "", model.ImpactUnknown},
	}
	for _, c := range cases {
		if got := gcpImpact(c.severity, c.statusImpact); got != c.want {
			t.Errorf("gcpImpact(%q, %q) = %q, want %q", c.severity, c.statusImpact, got, c.want)
		}
	}
}

func TestGCPUpdateStatus(t *testing.T) {
	cases := map[string]model.Status{
		"AVAILABLE":           model.StatusResolved,
		"SERVICE_DISRUPTION":  model.StatusInvestigating,
		"SERVICE_OUTAGE":      model.StatusInvestigating,
		"SERVICE_INFORMATION": model.StatusUnknown,
	}
	for in, want := range cases {
		if got := gcpUpdateStatus(in); got != want {
			t.Errorf("gcpUpdateStatus(%q) = %q, want %q", in, got, want)
		}
	}
}
