package source

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/jtprogru/weekly-incident/internal/model"
)

// golden reads a captured feed response from the repository's testdata dir.
func golden(t *testing.T, name string) []byte {
	t.Helper()
	b, err := os.ReadFile(filepath.Join("..", "..", "testdata", name))
	if err != nil {
		t.Fatalf("read golden %s: %v", name, err)
	}
	return b
}

func mustTime(t *testing.T, s string) time.Time {
	t.Helper()
	v, err := time.Parse(time.RFC3339Nano, s)
	if err != nil {
		t.Fatalf("parse time %q: %v", s, err)
	}
	return v.UTC()
}

func TestStatuspageParse(t *testing.T) {
	s := &statuspage{base: base{vendor: "github", url: "https://www.githubstatus.com/api/v2/incidents.json"}}
	res, err := s.parse(golden(t, "statuspage_github.json"))
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
	if in.Key != "github/bmpybhnrky3x" {
		t.Errorf("Key = %q", in.Key)
	}
	if in.Title != "Intermittent failures in runner group and runner-related permissions pages" {
		t.Errorf("Title = %q", in.Title)
	}
	if in.URL != "https://stspg.io/98zqb1k9jh0x" {
		t.Errorf("URL = %q, want the shortlink", in.URL)
	}
	if in.Impact != model.ImpactMinor {
		t.Errorf("Impact = %q, want minor", in.Impact)
	}
	if in.Status != model.StatusResolved {
		t.Errorf("Status = %q, want resolved", in.Status)
	}

	// started_at wins over created_at: the two differ by 6ms in this record.
	if want := mustTime(t, "2026-08-18T07:40:35.670Z"); !in.StartedAt.Equal(want) {
		t.Errorf("StartedAt = %s, want %s", in.StartedAt, want)
	}
	if in.ResolvedAt == nil {
		t.Fatal("ResolvedAt is nil on a resolved incident")
	}
	if want := mustTime(t, "2026-08-18T11:42:59.304Z"); !in.ResolvedAt.Equal(want) {
		t.Errorf("ResolvedAt = %s, want %s", in.ResolvedAt, want)
	}
	if in.DurationMinutes == nil || *in.DurationMinutes != 242 {
		t.Errorf("DurationMinutes = %v, want 242", in.DurationMinutes)
	}

	if len(in.Updates) != 5 {
		t.Fatalf("got %d updates, want 5", len(in.Updates))
	}
	// The feed lists updates newest first; the archive stores them in order.
	if want := mustTime(t, "2026-08-18T07:40:35.712Z"); !in.Updates[0].CreatedAt.Equal(want) {
		t.Errorf("first update at %s, want the oldest (%s)", in.Updates[0].CreatedAt, want)
	}
	assertAscending(t, in.Updates)

	if in.Components == nil {
		t.Error("Components is nil, want an empty slice so the JSON stays [] not null")
	}
	if len(in.Raw) == 0 {
		t.Error("Raw is empty")
	}
}

func TestStatuspageURLFallback(t *testing.T) {
	in := spIncident{ID: "abc123"}
	if got := spURL(in, "https://status.example.com/"); got != "https://status.example.com/incidents/abc123" {
		t.Errorf("spURL without shortlink = %q", got)
	}
	in.Shortlink = "https://stspg.io/x"
	if got := spURL(in, "https://status.example.com"); got != "https://stspg.io/x" {
		t.Errorf("spURL with shortlink = %q", got)
	}
}

func TestStatuspageSkipsUnparseableIncident(t *testing.T) {
	// An incident with neither started_at nor created_at is unusable, but it
	// must not take the rest of the feed down with it.
	body := []byte(`{"page":{"url":"https://x"},"incidents":[
		{"id":"ok","name":"fine","status":"resolved","impact":"minor","created_at":"2026-08-18T07:00:00Z"},
		{"id":"broken","name":"no timestamps"}
	]}`)
	s := &statuspage{base: base{vendor: "test"}}
	res, err := s.parse(body)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if len(res.Incidents) != 1 {
		t.Errorf("got %d incidents, want 1", len(res.Incidents))
	}
	if res.ParseErrors != 1 {
		t.Errorf("ParseErrors = %d, want 1", res.ParseErrors)
	}
}

func assertAscending(t *testing.T, us []model.Update) {
	t.Helper()
	for i := 1; i < len(us); i++ {
		if us[i].CreatedAt.Before(us[i-1].CreatedAt) {
			t.Errorf("updates out of order at %d: %s before %s", i, us[i].CreatedAt, us[i-1].CreatedAt)
		}
	}
}
