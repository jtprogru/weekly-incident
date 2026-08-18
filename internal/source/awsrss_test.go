package source

import (
	"reflect"
	"testing"
	"time"

	"github.com/jtprogru/weekly-incident/internal/model"
)

func TestAWSParse(t *testing.T) {
	a := &awsRSS{base: base{vendor: "aws", url: "https://status.aws.amazon.com/rss/all.rss"}}
	res, err := a.parse(golden(t, "aws_all.rss"))
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if res.ParseErrors != 0 {
		t.Errorf("ParseErrors = %d, want 0", res.ParseErrors)
	}

	// The captured feed holds 50 update items across three service keys. Two of
	// those keys carry a March chain plus a stray April update, split by the
	// 1383-hour gap; the third is a single resolved chain. Five events in all —
	// the guid epoch would have produced fifty.
	if len(res.Incidents) != 5 {
		t.Fatalf("got %d incidents, want 5", len(res.Incidents))
	}

	dc, found := findChain(res.Incidents, "directconnect-eu-central-1", model.StatusResolved)
	if !found {
		t.Fatal("resolved directconnect-eu-central-1 chain not found")
	}
	if len(dc.Updates) != 20 {
		t.Errorf("got %d updates in the directconnect chain, want 20", len(dc.Updates))
	}
	if dc.Title != "Increased Packet loss" {
		t.Errorf("Title = %q, want the headline with its prefix and [RESOLVED] stripped", dc.Title)
	}
	if dc.Impact != model.ImpactUnknown {
		t.Errorf("Impact = %q, want unknown: AWS reports no severity", dc.Impact)
	}
	if want := mustTime(t, "2026-08-15T03:42:59Z"); !dc.StartedAt.Equal(want) {
		t.Errorf("StartedAt = %s, want %s", dc.StartedAt, want)
	}
	if dc.ResolvedAt == nil {
		t.Fatal("ResolvedAt is nil on a chain ending in [RESOLVED]")
	}
	if want := mustTime(t, "2026-08-18T04:01:19Z"); !dc.ResolvedAt.Equal(want) {
		t.Errorf("ResolvedAt = %s, want %s", dc.ResolvedAt, want)
	}
	if dc.DurationMinutes == nil || *dc.DurationMinutes != 4338 {
		t.Errorf("DurationMinutes = %v, want 4338", dc.DurationMinutes)
	}
	assertAscending(t, dc.Updates)
	if dc.URL != awsStatusURL {
		t.Errorf("URL = %q", dc.URL)
	}
}

func TestAWSParseIsDeterministic(t *testing.T) {
	// Grouping runs through a map, so the output order has to be sorted
	// explicitly or the archive would churn on every run.
	a := &awsRSS{base: base{vendor: "aws"}}
	first, err := a.parse(golden(t, "aws_all.rss"))
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	for range 5 {
		next, err := a.parse(golden(t, "aws_all.rss"))
		if err != nil {
			t.Fatalf("parse: %v", err)
		}
		if !reflect.DeepEqual(first, next) {
			t.Fatal("two parses of the same bytes differ")
		}
	}
}

func TestSplitAWSEvents(t *testing.T) {
	at := func(h int) time.Time { return time.Date(2026, 8, 1, h, 0, 0, 0, time.UTC) }
	entry := func(h int, title string) awsEntry {
		return awsEntry{item: awsItem{Title: title}, service: "svc", at: at(h)}
	}

	t.Run("gap splits", func(t *testing.T) {
		got := splitAWSEvents([]awsEntry{
			entry(0, "Service impact: x"),
			entry(3, "Service impact: x"),
			entry(3+25, "Service impact: y"), // 25h later: a different event
		})
		if len(got) != 2 {
			t.Fatalf("got %d events, want 2", len(got))
		}
		if len(got[0]) != 2 || len(got[1]) != 1 {
			t.Errorf("event sizes = %d, %d; want 2, 1", len(got[0]), len(got[1]))
		}
	})

	t.Run("resolution closes the event", func(t *testing.T) {
		got := splitAWSEvents([]awsEntry{
			entry(0, "Service impact: x"),
			entry(1, "Service is operating normally: [RESOLVED] x"),
			entry(2, "Service impact: y"), // one hour later, but the last one closed
		})
		if len(got) != 2 {
			t.Fatalf("got %d events, want 2", len(got))
		}
		if len(got[0]) != 2 || len(got[1]) != 1 {
			t.Errorf("event sizes = %d, %d; want 2, 1", len(got[0]), len(got[1]))
		}
	})

	t.Run("uninterrupted chain stays whole", func(t *testing.T) {
		got := splitAWSEvents([]awsEntry{
			entry(0, "Service impact: x"),
			entry(6, "Service impact: x"),
			entry(12, "Service impact: x"),
		})
		if len(got) != 1 {
			t.Fatalf("got %d events, want 1", len(got))
		}
	})
}

func TestParseAWSService(t *testing.T) {
	got, err := parseAWSService("https://status.aws.amazon.com/#directconnect-eu-central-1_1787025679")
	if err != nil {
		t.Fatalf("parseAWSService: %v", err)
	}
	if got != "directconnect-eu-central-1" {
		t.Errorf("service = %q", got)
	}
	for _, bad := range []string{
		"https://status.aws.amazon.com/",
		"https://status.aws.amazon.com/#noepoch",
		"https://status.aws.amazon.com/#svc_notanumber",
	} {
		if _, err := parseAWSService(bad); err == nil {
			t.Errorf("parseAWSService(%q) succeeded, want an error", bad)
		}
	}
}

func TestParseAWSTime(t *testing.T) {
	cases := map[string]string{
		"Mon, 17 Aug 2026 21:01:19 PDT":   "2026-08-18T04:01:19Z",
		"Sun, 01 Mar 2026 12:14:03 PST":   "2026-03-01T20:14:03Z",
		"Tue, 18 Aug 2026 11:19:00 UTC":   "2026-08-18T11:19:00Z",
		"Tue, 18 Aug 2026 11:19:00 +0000": "2026-08-18T11:19:00Z",
	}
	for in, want := range cases {
		got, err := parseAWSTime(in)
		if err != nil {
			t.Errorf("parseAWSTime(%q): %v", in, err)
			continue
		}
		if got.Format(time.RFC3339) != want {
			t.Errorf("parseAWSTime(%q) = %s, want %s", in, got.Format(time.RFC3339), want)
		}
	}
	// An unrecognized abbreviation must fail loudly rather than silently
	// pretending to be UTC and shifting an incident by hours.
	if _, err := parseAWSTime("Mon, 17 Aug 2026 21:01:19 XYZ"); err == nil {
		t.Error("unknown timezone accepted, want an error")
	}
}

func TestCleanAWSTitle(t *testing.T) {
	cases := map[string]string{
		"Service is operating normally: [RESOLVED] Increased Packet loss": "Increased Packet loss",
		"Service impact: Increased Packet loss":                           "Increased Packet loss",
		"Service disruption: Increased Error Rates":                       "Increased Error Rates",
		"Informational message: Elevated latency":                         "Elevated latency",
		"Something unprefixed":                                            "Something unprefixed",
	}
	for in, want := range cases {
		if got := cleanAWSTitle(in); got != want {
			t.Errorf("cleanAWSTitle(%q) = %q, want %q", in, got, want)
		}
	}
}

// findChain locates the incident built from one service's update chain.
func findChain(incidents []model.Incident, service string, status model.Status) (model.Incident, bool) {
	for _, in := range incidents {
		if in.Status == status && reflect.DeepEqual(in.Components, []string{service}) {
			return in, true
		}
	}
	return model.Incident{}, false
}
