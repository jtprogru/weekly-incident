package source

import (
	"reflect"
	"strings"
	"testing"

	"github.com/jtprogru/weekly-incident/internal/model"
)

func TestSlackParse(t *testing.T) {
	s := &slack{base: base{vendor: "slack", url: "https://slack-status.com/api/v2.0.0/history"}}
	res, err := s.parse(golden(t, "slack_history.json"))
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if res.ParseErrors != 0 {
		t.Errorf("ParseErrors = %d, want 0", res.ParseErrors)
	}
	if len(res.Incidents) != 5 {
		t.Fatalf("got %d incidents, want 5", len(res.Incidents))
	}

	in := res.Incidents[0]
	if in.Key != "slack/1576" {
		t.Errorf("Key = %q, want the numeric id as a string", in.Key)
	}
	// Slack publishes no severity, so it can only ever clear the digest
	// threshold on duration.
	if in.Impact != model.ImpactUnknown {
		t.Errorf("Impact = %q, want unknown", in.Impact)
	}
	if in.Status != model.StatusInvestigating {
		t.Errorf("Status = %q, want investigating for an active incident", in.Status)
	}
	if in.ResolvedAt != nil {
		t.Errorf("ResolvedAt = %v, want nil while active", in.ResolvedAt)
	}
	if in.DurationMinutes != nil {
		t.Errorf("DurationMinutes = %v, want nil while active", in.DurationMinutes)
	}
	if want := mustTime(t, "2026-08-13T21:21:41Z"); !in.StartedAt.Equal(want) {
		t.Errorf("StartedAt = %s, want %s (the -07:00 offset normalized)", in.StartedAt, want)
	}
	if want := []string{"Messaging", "Workspace/Org Administration"}; !reflect.DeepEqual(in.Components, want) {
		t.Errorf("Components = %v, want %v", in.Components, want)
	}
	if len(in.Updates) != 6 {
		t.Fatalf("got %d updates, want 6", len(in.Updates))
	}
	assertAscending(t, in.Updates)
	for i, u := range in.Updates {
		if strings.Contains(u.Body, "<p>") || strings.Contains(u.Body, "</p>") {
			t.Errorf("update %d still carries HTML: %q", i, u.Body)
		}
	}
}

func TestSlackResolvedGetsDuration(t *testing.T) {
	body := []byte(`[{"id":1,"title":"t","type":"incident","status":"resolved",
		"url":"https://slack-status.com/x","date_created":"2026-08-01T00:00:00-07:00",
		"date_updated":"2026-08-01T02:30:00-07:00","services":["Messaging"],"notes":[]}]`)
	s := &slack{base: base{vendor: "slack"}}
	res, err := s.parse(body)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if len(res.Incidents) != 1 {
		t.Fatalf("got %d incidents, want 1", len(res.Incidents))
	}
	in := res.Incidents[0]
	if in.Status != model.StatusResolved {
		t.Errorf("Status = %q", in.Status)
	}
	if in.DurationMinutes == nil || *in.DurationMinutes != 150 {
		t.Errorf("DurationMinutes = %v, want 150", in.DurationMinutes)
	}
}
