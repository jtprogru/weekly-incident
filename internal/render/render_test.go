package render

import (
	"strings"
	"testing"
	"time"

	"github.com/jtprogru/weekly-incident/internal/model"
)

var (
	from = time.Date(2026, 8, 17, 0, 0, 0, 0, time.UTC)
	gen  = time.Date(2026, 8, 24, 7, 0, 0, 0, time.UTC)
)

func fixture() (model.WeekIndex, []model.Incident) {
	start := time.Date(2026, 8, 18, 3, 42, 0, 0, time.UTC)
	end := start.Add(4338 * time.Minute)
	dur := 4338
	short := 5

	big := model.Incident{
		Key: "aws/abc", Vendor: "aws", NativeID: "abc",
		Title: "Increased Packet loss", URL: "https://health.aws.amazon.com/health/status",
		Impact: model.ImpactUnknown, Status: model.StatusResolved,
		StartedAt: start, ResolvedAt: &end, DurationMinutes: &dur,
		Components: []string{"directconnect-eu-central-1"},
		Updates: []model.Update{
			{Status: model.StatusInvestigating, Body: "We are investigating.\n\nMore soon.", CreatedAt: start},
			{Status: model.StatusResolved, Body: strings.Repeat("long ", 400), CreatedAt: end},
		},
	}
	noise := model.Incident{
		Key: "cloudflare/x", Vendor: "cloudflare", NativeID: "x",
		Title: "Re-routed traffic | one datacenter", URL: "https://stspg.io/x",
		Impact: model.ImpactMinor, Status: model.StatusResolved,
		StartedAt: start, DurationMinutes: &short,
		Components: []string{},
	}

	idx := model.WeekIndex{
		Week: "2026-W34", From: from, To: from.AddDate(0, 0, 7), GeneratedAt: gen,
		Sources: []model.SourceStatus{
			{Vendor: "aws", OK: true, IncidentsSeen: 5},
			{Vendor: "datadog", OK: false, Error: "GET https://status.datadoghq.com: status 503"},
		},
		Incidents: []model.IndexEntry{
			{
				Key: big.Key, Vendor: big.Vendor, Title: big.Title, URL: big.URL,
				Impact: big.Impact, Status: big.Status, StartedAt: start, ResolvedAt: &end,
				DurationMinutes: &dur, Components: big.Components,
				Score:          model.Score{Value: 8676, Duration: dur, Components: 1, Weight: 2},
				AboveThreshold: true,
			},
			{
				Key: noise.Key, Vendor: noise.Vendor, Title: noise.Title, URL: noise.URL,
				Impact: noise.Impact, Status: noise.Status, StartedAt: start,
				DurationMinutes: &short, Components: noise.Components,
				Score:          model.Score{Value: 9, Duration: short, Components: 1, Weight: 1.8},
				AboveThreshold: false,
			},
		},
	}
	return idx, []model.Incident{big, noise}
}

func TestDigest(t *testing.T) {
	idx, incidents := fixture()
	out := Digest(idx, incidents)

	for _, want := range []string{
		"# Week 2026-W34",
		"2026-08-17 to 2026-08-23 (UTC)",
		"2 incidents collected, 1 above threshold",
		"## Above threshold",
		"[Increased Packet loss](https://health.aws.amazon.com/health/status)",
		"3d 0h 18m",
		"8676 = 4338 × 1 × 2.0",
		"## Timelines",
		"## Below threshold (1)",
		"## Sources",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("digest is missing %q", want)
		}
	}

	// A dead source has to be visible in the digest, not just in the logs.
	if !strings.Contains(out, "1 of 2 sources unavailable: datadog.") {
		t.Error("digest does not report the unavailable source in its summary")
	}
	if !strings.Contains(out, "FETCH FAILED") {
		t.Error("digest does not mark the failed source in the sources table")
	}

	// A pipe in a vendor title would otherwise split the markdown table.
	if strings.Contains(out, "Re-routed traffic | one") {
		t.Error("an unescaped pipe made it into the output")
	}
	if !strings.Contains(out, `Re-routed traffic \| one datacenter`) {
		t.Error("the pipe in the title was not escaped")
	}

	// Multi-paragraph bodies must not break the list they sit in.
	for _, line := range strings.Split(out, "\n") {
		if strings.HasPrefix(line, "- `") && strings.Contains(line, "We are investigating") {
			if !strings.Contains(line, "More soon.") {
				t.Error("a multi-paragraph update body was split across lines")
			}
		}
	}
}

func TestDigestWithNothingAboveThreshold(t *testing.T) {
	idx, incidents := fixture()
	for i := range idx.Incidents {
		idx.Incidents[i].AboveThreshold = false
	}
	out := Digest(idx, incidents)
	if !strings.Contains(out, "Nothing cleared the threshold this week.") {
		t.Error("a quiet week should say so plainly")
	}
	if strings.Contains(out, "## Timelines") {
		t.Error("no timelines section should appear when nothing cleared the threshold")
	}
}

func TestBriefTruncatesBodies(t *testing.T) {
	idx, incidents := fixture()
	out := Brief(idx, incidents, 5, 100)

	if !strings.Contains(out, "top 1 of 1 above threshold") {
		t.Error("brief header does not report the counts")
	}
	if strings.Contains(out, strings.Repeat("long ", 400)) {
		t.Error("brief did not truncate a long update body")
	}
	if !strings.Contains(out, "…") {
		t.Error("brief does not mark where it cut")
	}
}

func TestBriefRespectsTopN(t *testing.T) {
	idx, incidents := fixture()
	idx.Incidents[1].AboveThreshold = true

	out := Brief(idx, incidents, 1, 800)
	if !strings.Contains(out, "## Also above threshold (1)") {
		t.Error("brief does not list what it left out of the top N")
	}
}

func TestHumanMinutes(t *testing.T) {
	cases := map[int]string{
		0:    "0m",
		42:   "42m",
		90:   "1h 30m",
		1440: "1d 0h 0m",
		4338: "3d 0h 18m",
		-5:   "0m",
	}
	for in, want := range cases {
		if got := humanMinutes(in); got != want {
			t.Errorf("humanMinutes(%d) = %q, want %q", in, got, want)
		}
	}
}

func TestOngoingIsMarked(t *testing.T) {
	e := model.IndexEntry{Ongoing: true, Score: model.Score{Duration: 200}}
	if got := duration(e); got != "3h 20m (ongoing)" {
		t.Errorf("duration = %q, want the ongoing marker", got)
	}
}

func TestTruncateCutsOnRuneBoundary(t *testing.T) {
	// Vendor bodies are not always ASCII; slicing bytes would emit mojibake.
	s := "мониторинг восстановления"
	got := truncate(s, 10)
	if !strings.HasSuffix(got, "…") {
		t.Fatalf("truncate = %q, want a trailing ellipsis", got)
	}
	if r := []rune(strings.TrimSuffix(got, "…")); len(r) > 10 {
		t.Errorf("truncate kept %d runes, want at most 10", len(r))
	}
	if strings.ContainsRune(got, '�') {
		t.Error("truncate produced a replacement character")
	}
	if truncate("short", 100) != "short" {
		t.Error("truncate altered a string under the limit")
	}
}

func TestOneLine(t *testing.T) {
	if got := oneLine("a\r\nb\n\n  c  "); got != "a b c" {
		t.Errorf("oneLine = %q", got)
	}
}

func TestDigestWarnsAboutParseErrors(t *testing.T) {
	// A growing count in the sources table gets scrolled past; a dropped record
	// means the vendor changed its schema and incidents are going missing.
	idx, incidents := fixture()
	idx.Sources[0].ParseErrors = 3
	out := Digest(idx, incidents)

	if !strings.Contains(out, "3 record(s) could not be parsed") {
		t.Error("digest does not warn about dropped records")
	}
	if !strings.Contains(out, "aws (3)") {
		t.Error("the warning does not name the vendor")
	}
	if !strings.Contains(out, "make testdata") {
		t.Error("the warning does not say what to do about it")
	}

	clean, _ := fixture()
	if strings.Contains(Digest(clean, incidents), "could not be parsed") {
		t.Error("a clean run should carry no parse warning")
	}
}

func TestNotification(t *testing.T) {
	idx, _ := fixture()
	out := Notification(idx, 3)

	for _, want := range []string{
		"<b>2026-W34</b> — incidents of the week",
		"2026-08-17 → 2026-08-23 UTC",
		"<b>1 above threshold</b>",
		`<a href="https://health.aws.amazon.com/health/status">Increased Packet loss</a>`,
		"aws · 3d 0h 18m",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("notification is missing %q\n---\n%s", want, out)
		}
	}
	// A dead source has to be visible at a glance, not buried in the digest.
	if !strings.Contains(out, "<b>1/2 sources ok</b> (failed: datadog)") {
		t.Errorf("notification does not flag the failed source:\n%s", out)
	}
	// AWS reports no severity; printing "unknown" as if it meant something
	// would be worse than saying nothing.
	if strings.Contains(out, "unknown") {
		t.Error("notification prints the unknown impact")
	}
	if strings.Contains(out, "**") {
		t.Error("markdown leaked into a Telegram HTML message")
	}
}

func TestNotificationEscapesVendorText(t *testing.T) {
	// Vendor headlines are attacker-adjacent text: one bare < and Telegram
	// rejects the whole message as broken HTML.
	idx, _ := fixture()
	idx.Incidents[0].Title = `Errors in <script> & "quoted" mode`
	out := Notification(idx, 3)

	if strings.Contains(out, "<script>") {
		t.Error("a raw tag from a vendor title reached the message")
	}
	if !strings.Contains(out, "&lt;script&gt;") || !strings.Contains(out, "&amp;") {
		t.Errorf("vendor title was not escaped:\n%s", out)
	}
}

func TestNotificationTruncatesLongHeadlines(t *testing.T) {
	idx, _ := fixture()
	idx.Incidents[0].Title = strings.Repeat("very long headline ", 20)
	out := Notification(idx, 3)

	if !strings.Contains(out, "…") {
		t.Error("a long headline was not truncated")
	}
	for _, line := range strings.Split(out, "\n") {
		if len([]rune(line)) > 400 {
			t.Errorf("line of %d runes is too long for a chat message", len([]rune(line)))
		}
	}
}

func TestNotificationRespectsTopN(t *testing.T) {
	idx, _ := fixture()
	idx.Incidents[1].AboveThreshold = true

	out := Notification(idx, 1)
	if !strings.Contains(out, "…and 1 more above threshold.") {
		t.Errorf("notification does not say what it left out:\n%s", out)
	}
}

func TestNotificationOnAQuietWeek(t *testing.T) {
	idx, _ := fixture()
	for i := range idx.Incidents {
		idx.Incidents[i].AboveThreshold = false
	}
	out := Notification(idx, 3)
	if !strings.Contains(out, "Nothing cleared the threshold this week.") {
		t.Error("a quiet week should say so plainly")
	}
}

func TestNotificationFlagsDroppedRecords(t *testing.T) {
	idx, _ := fixture()
	idx.Sources[0].ParseErrors = 2
	if !strings.Contains(Notification(idx, 3), "2 record(s) dropped as unparseable") {
		t.Error("notification hides dropped records")
	}
}
