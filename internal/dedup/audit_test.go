package dedup

import (
	"encoding/json"
	"os"
	"reflect"
	"slices"
	"strings"
	"testing"
	"time"

	"github.com/jtprogru/weekly-incident/internal/archive"
	"github.com/jtprogru/weekly-incident/internal/config"
	"github.com/jtprogru/weekly-incident/internal/model"
)

var fixedNow = time.Date(2026, 11, 14, 0, 0, 0, 0, time.UTC)

func TestAuditArchivedAWSWindow(t *testing.T) {
	b, err := os.ReadFile("../../testdata/dedup/aws-window.json")
	if err != nil {
		t.Fatal(err)
	}
	var incidents []model.Incident
	if err := json.Unmarshal(b, &incidents); err != nil {
		t.Fatal(err)
	}
	if len(incidents) != 2 {
		t.Fatal("expected the archived pair")
	}
	r := auditRecords(incidents, config.KindAWSRSS)
	if len(r.Candidates) != 1 {
		t.Fatalf("candidates = %d", len(r.Candidates))
	}
	c := r.Candidates[0]
	if len(c.SharedUpdates) != 16 || c.Level != "review_required" {
		t.Fatalf("shared = %d; level = %s", len(c.SharedUpdates), c.Level)
	}
	if !slices.Contains(c.Reasons, "update_set_containment") || slices.Contains(c.Reasons, "same_event_url") {
		t.Fatalf("reasons = %v", c.Reasons)
	}
	unique := map[string]bool{}
	for _, in := range incidents {
		for _, u := range in.Updates {
			unique[fingerprint(u.CreatedAt.Format(time.RFC3339Nano), normalizeBody(u.Body))] = true
		}
	}
	if len(unique) != 22 {
		t.Fatalf("fixture union has %d updates", len(unique))
	}
	if r.Statistics.SourceKeys != 2 || r.Statistics.RepeatedKeyGroups != 0 {
		t.Fatalf("statistics = %+v", r.Statistics)
	}
}

// Synthetic cases intentionally challenge the known AWS pair rather than
// assuming that any shared wording identifies one incident.
func TestAuditEvidenceBoundaries(t *testing.T) {
	for _, tc := range []struct {
		name                   string
		kind                   config.Kind
		change                 func(*model.Incident, *model.Incident)
		pairs                  int
		level, reason, caution string
	}{
		{"exact-update", config.KindAWSRSS, func(a, b *model.Incident) {}, 1, "review_required", "exact_shared_updates", "single_shared_update_may_be_template"},
		{"different-region", config.KindAWSRSS, func(a, b *model.Incident) { b.Components = []string{"ec2-eu-west-1"} }, 1, "conflict", "exact_shared_updates", "disjoint_components"},
		{"recurrence-after-resolved", config.KindAWSRSS, func(a, b *model.Incident) {
			end := a.StartedAt.Add(time.Minute)
			a.ResolvedAt = &end
			b.StartedAt = end.Add(time.Hour)
		}, 1, "conflict", "exact_shared_updates", "non_overlapping_reported_windows"},
		{"only-title-and-components", config.KindAWSRSS, func(a, b *model.Incident) { b.Updates = nil }, 0, "", "", ""},
		{"same-body-different-time", config.KindAWSRSS, func(a, b *model.Incident) { b.Updates[0].CreatedAt = b.Updates[0].CreatedAt.Add(time.Hour) }, 0, "", "", ""},
		{"empty-body", config.KindAWSRSS, func(a, b *model.Incident) { a.Updates[0].Body = " "; b.Updates[0].Body = "\n" }, 0, "", "", ""},
		{"zero-update-time", config.KindAWSRSS, func(a, b *model.Incident) { a.Updates[0].CreatedAt = time.Time{}; b.Updates[0].CreatedAt = time.Time{} }, 0, "", "", ""},
		{"different-vendor", config.KindStatuspage, func(a, b *model.Incident) { b.Vendor = "elsewhere"; b.Finalize() }, 0, "", "", ""},
		{"same-event-url", config.KindStatuspage, func(a, b *model.Incident) { a.URL = "https://example.test/incidents/a"; b.URL = a.URL; b.Updates = nil }, 1, "review_required", "same_event_url", ""},
		{"same-statuspage-root", config.KindStatuspage, func(a, b *model.Incident) { a.URL = "https://example.test/"; b.URL = a.URL; b.Updates = nil }, 0, "", "", ""},
		{"edited-update-id", config.KindStatuspage, func(a, b *model.Incident) {
			a.Updates[0].NativeID = "u1"
			b.Updates[0].NativeID = "u1"
			b.Updates[0].Body = "corrected message"
		}, 1, "review_required", "shared_update_ids", "shared_update_id_has_different_content"},
		{"repeat-stable-key", config.KindStatuspage, func(a, b *model.Incident) { b.NativeID = a.NativeID; b.Finalize() }, 1, "exact_identity", "same_stable_source_id", ""},
		{"repeat-key-date-correction", config.KindStatuspage, func(a, b *model.Incident) {
			b.NativeID = a.NativeID
			b.StartedAt = a.StartedAt.Add(-7 * 24 * time.Hour)
			b.Finalize()
		}, 1, "conflict", "repeated_key", "repeated_key_content_differs_revision_order_unknown"},
		{"normalization", config.KindAWSRSS, func(a, b *model.Incident) { a.Updates[0].Body = " A\r\nB "; b.Updates[0].Body = "A\nB" }, 1, "review_required", "exact_shared_updates", ""},
		{"negation-is-significant", config.KindAWSRSS, func(a, b *model.Incident) {
			a.Updates[0].Body = "Service is affected"
			b.Updates[0].Body = "Service is not affected"
		}, 0, "", "", ""},
	} {
		t.Run(tc.name, func(t *testing.T) {
			a, b := synthetic("a"), synthetic("b")
			tc.change(&a, &b)
			r := auditRecords([]model.Incident{a, b}, tc.kind)
			if len(r.Candidates) != tc.pairs {
				t.Fatalf("candidates = %d, want %d", len(r.Candidates), tc.pairs)
			}
			if tc.pairs == 0 {
				return
			}
			c := r.Candidates[0]
			if c.Level != tc.level || !slices.Contains(c.Reasons, tc.reason) {
				t.Fatalf("candidate = %+v", c)
			}
			if tc.caution != "" && !slices.Contains(c.Cautions, tc.caution) {
				t.Fatalf("missing caution %s: %v", tc.caution, c.Cautions)
			}
		})
	}
}

func TestAuditPermutationAndNonMutation(t *testing.T) {
	in := []model.Incident{synthetic("a"), synthetic("b"), synthetic("c")}
	in[0].Components = []string{"two", "one", "two"}
	inv := inventory(in)
	sources := []config.Source{{Vendor: "test", Kind: config.KindStatuspage}, {Vendor: "other", Kind: config.KindAWSRSS}}
	before, err := json.Marshal(inv)
	if err != nil {
		t.Fatal(err)
	}
	want := Audit(inv, sources, fixedNow)
	after, err := json.Marshal(inv)
	if err != nil {
		t.Fatal(err)
	}
	if string(before) != string(after) {
		t.Fatal("audit mutated input")
	}
	slices.Reverse(inv.Records)
	slices.Reverse(sources)
	got := Audit(inv, sources, fixedNow)
	if !reflect.DeepEqual(got, want) {
		t.Fatal("input iteration order changed report")
	}
	if len(got.Candidates) != 3 {
		t.Fatalf("pair count = %d; pairs must not be merged into a group", len(got.Candidates))
	}
}

func TestAuditUnknownPolicyAndEmptyArchive(t *testing.T) {
	r := Audit(inventory([]model.Incident{synthetic("a"), synthetic("b")}), nil, fixedNow)
	if r.Statistics.Records != 2 || len(r.Warnings) != 1 || len(r.Candidates) != 1 {
		t.Fatalf("report = %+v", r)
	}
	empty := Audit(inventory(nil), nil, fixedNow)
	if len(empty.Warnings) != 1 || empty.Statistics.Records != 0 {
		t.Fatal("empty archive not marked")
	}
}

func TestEventURLPolicy(t *testing.T) {
	for _, tc := range []struct {
		kind  config.Kind
		raw   string
		valid bool
	}{
		{config.KindAWSRSS, "https://health.aws.amazon.com/health/status", false},
		{config.KindStatuspage, "https://status.example.test/", false},
		{config.KindStatuspage, "https://status.example.test/api/v2/incidents.json", false},
		{config.KindStatuspage, "https://status.example.test/incidents/abc", true},
		{config.KindStatuspage, "https://stspg.io/abc", true},
		{config.KindStatuspage, "javascript:alert(1)", false},
		{config.KindStatuspage, "https://user:password@example.test/incidents/a", false},
		{config.KindSlack, "https://slack-status.com/2026-02/abc", true},
		{config.KindGCP, "https://status.cloud.google.com/incidents/abc", true},
		{config.KindYandexCloud, "https://status.yandex.cloud/en/incidents/123", true},
	} {
		t.Run(string(tc.kind)+tc.raw, func(t *testing.T) {
			got := eventURL(tc.raw, policyFor(config.Source{Kind: tc.kind}))
			if (got != "") != tc.valid {
				t.Fatalf("eventURL = %q, valid = %t", got, tc.valid)
			}
		})
	}
	p := policyFor(config.Source{Kind: config.KindStatuspage})
	if eventURL("https://example.test/incidents/a?event=1", p) == eventURL("https://example.test/incidents/a?event=2", p) {
		t.Fatal("query parameters lost")
	}
}

func TestMarkdownEscapesVendorContent(t *testing.T) {
	a, b := synthetic("a"), synthetic("b")
	a.Title = "<script>alert(1)</script> | [click](javascript:bad)\n# spoofed"
	a.Updates[0].Body = strings.Repeat("Я", 700)
	b.Updates[0].Body = a.Updates[0].Body
	r := auditRecords([]model.Incident{a, b}, config.KindAWSRSS)
	md := Markdown(r)
	if strings.Contains(md, "<script>") || strings.Contains(md, "\n# spoofed") || strings.Contains(md, "[click](javascript:") {
		t.Fatal("unescaped vendor content")
	}
	if !strings.Contains(md, "сокращено") || len([]rune(r.Candidates[0].SharedUpdates[0].Body)) != 700 {
		t.Fatal("full JSON evidence lost or truncation not disclosed")
	}
}

func auditRecords(incidents []model.Incident, kind config.Kind) Report {
	sources := []config.Source{{Vendor: "test", Kind: kind}, {Vendor: "aws", Kind: kind}}
	return Audit(inventory(incidents), sources, fixedNow)
}

func inventory(incidents []model.Incident) archive.Inventory {
	inv := archive.Inventory{Records: []archive.StoredIncident{}, Files: []archive.InputFile{}}
	for i, in := range incidents {
		inv.Records = append(inv.Records, archive.StoredIncident{Path: "2026/W46/" + in.Vendor + ".json", Position: i + 1, Incident: in})
	}
	return inv
}

func synthetic(id string) model.Incident {
	in := model.Incident{Vendor: "test", NativeID: id, Title: "Increased error rates", URL: "https://health.aws.amazon.com/health/status",
		StartedAt: fixedNow.Add(-time.Hour), Status: model.StatusInvestigating, Impact: model.ImpactUnknown,
		Components: []string{"ec2-us-east-1"}, Updates: []model.Update{{CreatedAt: fixedNow, Body: "We are investigating increased API error rates in us-east-1.", Status: model.StatusInvestigating}}, Raw: json.RawMessage(`{}`)}
	in.Finalize()
	return in
}
