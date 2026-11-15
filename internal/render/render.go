// Package render turns a week of archived incidents into the two Markdown
// files a human works from.
//
// Everything here is English and deterministic. No network, no LLM, no
// translation: the collector's output has to be reproducible from the archive
// alone, and the Russian text of the actual post is written by hand elsewhere.
package render

import (
	"fmt"
	"html"
	"strings"
	"time"

	"github.com/jtprogru/weekly-incident/internal/model"
)

const (
	timeFormat = "2006-01-02 15:04"
	dateFormat = "2006-01-02"
)

// Digest renders the full fact sheet: a scored table of everything above the
// threshold, the timeline of each of those incidents, a one-line tail for the
// noise, and the health of every source.
func Digest(idx model.WeekIndex, incidents []model.Incident) string {
	incidentByKey := byKey(incidents)
	above, below := split(idx.Incidents)

	var b strings.Builder
	fmt.Fprintf(&b, "# Week %s\n\n", idx.Week)
	fmt.Fprintf(&b, "%s to %s (UTC). Generated %s.\n\n",
		idx.From.Format(dateFormat),
		idx.To.AddDate(0, 0, -1).Format(dateFormat),
		idx.GeneratedAt.Format(timeFormat))

	label := "incidents collected"
	if idx.Dedup != nil {
		label = "canonical incidents"
	}
	head := fmt.Sprintf("%d %s, %d above threshold.", len(idx.Incidents), label, len(above))
	if healthy := allSourcesAnswered(idx.Sources); healthy != "" {
		head += " " + healthy
	}
	fmt.Fprintf(&b, "%s\n\n", head)
	writeDedupCounts(&b, idx.Dedup)

	if w := unavailableWarning(idx.Sources, idx.GeneratedAt); w != "" {
		fmt.Fprintf(&b, "%s\n\n", w)
	}

	if w := parseErrorWarning(idx.Sources); w != "" {
		fmt.Fprintf(&b, "%s\n\n", w)
	}

	if len(above) == 0 {
		b.WriteString("Nothing cleared the threshold this week.\n\n")
	} else {
		b.WriteString("## Above threshold\n\n")
		b.WriteString("| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |\n")
		b.WriteString("|---|--------|----------|---------------|----------|--------|------------|-------|\n")
		for i, e := range above {
			fmt.Fprintf(&b, "| %d | %s | %s | %s | %s | %s | %s | %s |\n",
				i+1,
				e.Vendor,
				link(e.Title, e.URL),
				e.StartedAt.Format(timeFormat),
				duration(e),
				e.Impact,
				components(e.Components),
				scoreCell(e.Score))
		}
		b.WriteString("\nScore is duration in minutes, times component count, times vendor weight.\n\n")

		b.WriteString("## Timelines\n\n")
		for i, e := range above {
			writeTimeline(&b, i+1, e, incidentByKey[e.Key])
		}
	}

	if len(below) > 0 {
		fmt.Fprintf(&b, "## Below threshold (%d)\n\n", len(below))
		for _, e := range below {
			fmt.Fprintf(&b, "- `%s` %s — %s\n", e.Vendor, link(e.Title, e.URL), duration(e))
		}
		b.WriteString("\n")
	}

	writeCanonicalCatalogue(&b, idx.Incidents)
	writeSources(&b, idx.Sources)
	return strings.TrimRight(b.String(), "\n") + "\n"
}

// Brief renders the top incidents with their update bodies truncated, so the
// week fits in one paste. Nothing is summarized, only cut.
func Brief(idx model.WeekIndex, incidents []model.Incident, topN, bodyLimit int) string {
	incidentByKey := byKey(incidents)
	above, _ := split(idx.Incidents)

	top := above
	if len(top) > topN {
		top = top[:topN]
	}

	var b strings.Builder
	fmt.Fprintf(&b, "# Week %s — top %d of %d above threshold\n\n", idx.Week, len(top), len(above))
	fmt.Fprintf(&b, "%s to %s (UTC).\n\n", idx.From.Format(dateFormat), idx.To.AddDate(0, 0, -1).Format(dateFormat))
	writeDedupCounts(&b, idx.Dedup)

	if w := unavailableWarning(idx.Sources, idx.GeneratedAt); w != "" {
		fmt.Fprintf(&b, "%s\n\n", w)
	}

	for i, e := range top {
		in := incidentByKey[e.Key]
		fmt.Fprintf(&b, "## %d. %s — %s\n\n", i+1, e.Vendor, e.Title)
		fmt.Fprintf(&b, "- URL: %s\n", e.URL)
		fmt.Fprintf(&b, "- Window: %s\n", window(e))
		fmt.Fprintf(&b, "- Impact: %s, status: %s\n", e.Impact, e.Status)
		fmt.Fprintf(&b, "- Components: %s\n\n", components(e.Components))
		writeCanonicalInfo(&b, e.Canonical)
		for _, u := range in.Updates {
			fmt.Fprintf(&b, "  - `%s` %s — %s\n", u.CreatedAt.Format(timeFormat), u.Status, oneLine(truncate(u.Body, bodyLimit)))
		}
		b.WriteString("\n")
	}

	if rest := above[len(top):]; len(rest) > 0 {
		fmt.Fprintf(&b, "## Also above threshold (%d)\n\n", len(rest))
		for _, e := range rest {
			fmt.Fprintf(&b, "- `%s` %s (%s) — %s\n", e.Vendor, e.Title, duration(e), e.URL)
		}
		b.WriteString("\n")
	}
	return strings.TrimRight(b.String(), "\n") + "\n"
}

func writeTimeline(b *strings.Builder, n int, e model.IndexEntry, in model.Incident) {
	fmt.Fprintf(b, "### %d. %s — %s\n\n", n, e.Vendor, e.Title)
	fmt.Fprintf(b, "- URL: %s\n", e.URL)
	fmt.Fprintf(b, "- Started: %s UTC\n", e.StartedAt.Format(timeFormat))
	fmt.Fprintf(b, "- Resolved: %s\n", resolvedCell(e))
	fmt.Fprintf(b, "- Duration: %s\n", plainDuration(e))
	fmt.Fprintf(b, "- Impact: %s, status: %s\n", e.Impact, e.Status)
	fmt.Fprintf(b, "- Components: %s\n", components(e.Components))
	if e.Canonical != nil {
		fmt.Fprintf(b, "- Canonical incident: [%s](#%s)\n", e.Canonical.ID, e.Canonical.ID)
	}
	if len(in.Updates) > 0 {
		b.WriteString("\n")
		for _, u := range in.Updates {
			fmt.Fprintf(b, "- `%s` **%s** — %s\n", u.CreatedAt.Format(timeFormat), u.Status, oneLine(u.Body))
		}
	}
	b.WriteString("\n")
}

func writeSources(b *strings.Builder, sources []model.SourceStatus) {
	if len(sources) == 0 {
		return
	}
	b.WriteString("## Sources\n\n")
	b.WriteString("| Vendor | Status | Source records seen | Parse errors |\n")
	b.WriteString("|--------|--------|----------------|--------------|\n")
	for _, s := range sources {
		status := "ok"
		if !s.OK {
			status = "FETCH FAILED: " + oneLine(s.Error)
		}
		fmt.Fprintf(b, "| %s | %s | %d | %d |\n", s.Vendor, status, s.IncidentsSeen, s.ParseErrors)
	}
	b.WriteString("\n")
}

// parseErrorWarning shouts when a feed stopped matching its parser. A silently
// growing count in the sources table is exactly the kind of thing that gets
// scrolled past, and it means incidents are being dropped.
func parseErrorWarning(sources []model.SourceStatus) string {
	var broken []string
	total := 0
	for _, s := range sources {
		if s.ParseErrors > 0 {
			broken = append(broken, fmt.Sprintf("%s (%d)", s.Vendor, s.ParseErrors))
			total += s.ParseErrors
		}
	}
	if total == 0 {
		return ""
	}
	return fmt.Sprintf("**Warning: %d record(s) could not be parsed and were dropped: %s.** "+
		"A vendor most likely changed its schema; re-capture the goldens with `make testdata` and check the parser.",
		total, strings.Join(broken, ", "))
}

// unavailableWarning shouts when a source did not answer, and says for how
// long. It sits above the week's content rather than in the sources table at
// the bottom, because a missing vendor is invisible by nature: nothing in the
// week says "datadog would have had four incidents here".
//
// The age is measured against the snapshot the digest quotes, never against the
// wall clock, so re-rendering a settled week leaves the note where it was.
func unavailableWarning(sources []model.SourceStatus, snapshot time.Time) string {
	var failed []string
	var lines []string
	for _, s := range sources {
		if s.OK {
			continue
		}
		failed = append(failed, s.Vendor)
		lines = append(lines, fmt.Sprintf("- `%s` — %s. Error: %s",
			s.Vendor, lastSuccessPhrase(s, snapshot), oneLine(s.Error)))
	}
	if len(failed) == 0 {
		return ""
	}

	var b strings.Builder
	fmt.Fprintf(&b, "> **%d of %d sources unavailable: %s.**\n>",
		len(failed), len(sources), strings.Join(failed, ", "))
	b.WriteString(" Their incidents are missing from this week, and the feeds are\n")
	b.WriteString("> too shallow and unpaginated to backfill afterwards.\n")
	for _, line := range lines {
		b.WriteString("> " + line + "\n")
	}
	return strings.TrimRight(b.String(), "\n")
}

// lastSuccessPhrase renders when a source last answered, relative to the
// snapshot. A source that has never answered during a recorded run says so
// rather than inventing a date.
func lastSuccessPhrase(s model.SourceStatus, snapshot time.Time) string {
	if s.LastSuccessAt == nil || s.LastSuccessAt.IsZero() {
		return "no successful fetch recorded in this snapshot"
	}
	at := s.LastSuccessAt.UTC()
	switch days := int(snapshot.UTC().Sub(at).Hours() / 24); {
	case days < 1:
		return fmt.Sprintf("last successful fetch %s", at.Format(timeFormat))
	case days == 1:
		return fmt.Sprintf("last successful fetch %s, 1 day before this snapshot", at.Format(timeFormat))
	default:
		return fmt.Sprintf("last successful fetch %s, %d days before this snapshot", at.Format(timeFormat), days)
	}
}

// allSourcesAnswered says the quiet thing when every feed responded, so a
// healthy week still records that the check happened.
func allSourcesAnswered(sources []model.SourceStatus) string {
	if len(sources) == 0 {
		return ""
	}
	for _, s := range sources {
		if !s.OK {
			return ""
		}
	}
	return fmt.Sprintf("All %d sources responded.", len(sources))
}

func split(entries []model.IndexEntry) (above, below []model.IndexEntry) {
	for _, e := range entries {
		if e.AboveThreshold {
			above = append(above, e)
		} else {
			below = append(below, e)
		}
	}
	return above, below
}

func byKey(incidents []model.Incident) map[string]model.Incident {
	m := make(map[string]model.Incident, len(incidents))
	for _, in := range incidents {
		m[in.Key] = in
	}
	return m
}

// duration renders the entry's length, marking one that is still running.
func duration(e model.IndexEntry) string {
	if e.Ongoing {
		return plainDuration(e) + " (ongoing)"
	}
	return plainDuration(e)
}

// plainDuration renders the length alone: the ongoing marker belongs to the
// field that reports whether the incident ended, and repeating it in both
// places produced "to not yet (4d 10h 42m (ongoing))".
func plainDuration(e model.IndexEntry) string {
	if e.Ongoing {
		return humanMinutes(e.Score.Duration)
	}
	if e.DurationMinutes == nil {
		return "unknown"
	}
	return humanMinutes(*e.DurationMinutes)
}

// window renders one incident's span in one piece: a start, an end or "now",
// and the length. It exists because splitting the span across separate
// resolved/duration fields read as "to not yet" once the incident was ongoing.
func window(e model.IndexEntry) string {
	start := e.StartedAt.Format(timeFormat) + " UTC"
	end, mark := "now", ", ongoing"
	if e.ResolvedAt != nil {
		end, mark = e.ResolvedAt.Format(timeFormat)+" UTC", ""
	}
	return fmt.Sprintf("%s to %s (%s%s)", start, end, plainDuration(e), mark)
}

// humanMinutes renders 4338 as "3d 0h 18m" and 42 as "42m".
func humanMinutes(m int) string {
	if m < 0 {
		m = 0
	}
	d, rem := m/1440, m%1440
	h, min := rem/60, rem%60
	switch {
	case d > 0:
		return fmt.Sprintf("%dd %dh %dm", d, h, min)
	case h > 0:
		return fmt.Sprintf("%dh %dm", h, min)
	default:
		return fmt.Sprintf("%dm", min)
	}
}

func resolvedCell(e model.IndexEntry) string {
	if e.ResolvedAt == nil {
		return "ongoing"
	}
	return e.ResolvedAt.Format(timeFormat) + " UTC"
}

func components(cs []string) string {
	if len(cs) == 0 {
		return "—"
	}
	return strings.Join(cs, ", ")
}

// scoreCell shows the arithmetic, not just the result: a reader has to be able
// to see why one incident outranked another.
func scoreCell(s model.Score) string {
	return fmt.Sprintf("%.0f = %d × %d × %.1f", s.Value, s.Duration, s.Components, s.Weight)
}

func link(title, url string) string {
	title = escapePipes(oneLine(title))
	if url == "" {
		return title
	}
	return "[" + title + "](" + url + ")"
}

// escapePipes keeps a vendor title with a pipe in it from breaking the table.
func escapePipes(s string) string { return strings.ReplaceAll(s, "|", "\\|") }

// oneLine flattens a multi-paragraph body so it survives inside a list item or
// a table cell.
func oneLine(s string) string {
	s = strings.ReplaceAll(s, "\r\n", "\n")
	s = strings.ReplaceAll(s, "\n", " ")
	return strings.Join(strings.Fields(s), " ")
}

// truncate cuts on a rune boundary and marks the cut.
func truncate(s string, limit int) string {
	if limit <= 0 {
		return s
	}
	r := []rune(s)
	if len(r) <= limit {
		return s
	}
	return strings.TrimSpace(string(r[:limit])) + "…"
}

// notifyTitleLimit keeps one headline to a single glance in a chat client.
const notifyTitleLimit = 90

// Notification renders the Telegram body for a finished week.
//
// The markup is Telegram HTML, which needs only &, < and > escaped. MarkdownV2
// would demand escaping a dozen punctuation marks that vendor titles are full
// of — dots, parens, dashes — and a single missed one makes Telegram reject the
// whole message.
//
// The link to the full digest is not built here: it needs the repository path,
// which only the workflow knows.
func Notification(idx model.WeekIndex, topN int) string {
	above, _ := split(idx.Incidents)

	var b strings.Builder
	fmt.Fprintf(&b, "<b>%s</b> — incidents of the week\n", esc(idx.Week))
	fmt.Fprintf(&b, "%s → %s UTC\n\n",
		idx.From.Format(dateFormat), idx.To.AddDate(0, 0, -1).Format(dateFormat))
	fmt.Fprintf(&b, "%d collected · <b>%d above threshold</b> · %s\n",
		len(idx.Incidents), len(above), sourceHealth(idx.Sources, idx.GeneratedAt))
	if idx.Dedup != nil {
		fmt.Fprintf(&b, "%d source records → %d canonical incidents · %d candidate pairs pending review\n", idx.Dedup.SourceRecords, idx.Dedup.Incidents, idx.Dedup.PendingPairs)
	}

	if w := parseErrorSummary(idx.Sources); w != "" {
		fmt.Fprintf(&b, "%s\n", w)
	}

	if len(above) == 0 {
		b.WriteString("\nNothing cleared the threshold this week.\n")
		return b.String()
	}

	top := above
	if len(top) > topN {
		top = top[:topN]
	}
	for i, e := range top {
		fmt.Fprintf(&b, "\n<b>%d.</b> %s\n", i+1, notifyLink(e))
		fmt.Fprintf(&b, "%s · %s%s\n", esc(e.Vendor), esc(duration(e)), impactSuffix(e))
	}
	if rest := len(above) - len(top); rest > 0 {
		fmt.Fprintf(&b, "\n…and %d more above threshold.\n", rest)
	}
	return b.String()
}

// notifyLink renders one headline as an anchor, or as bold text when the
// vendor gave no URL.
func notifyLink(e model.IndexEntry) string {
	title := esc(truncate(oneLine(e.Title), notifyTitleLimit))
	if e.URL == "" {
		return "<b>" + title + "</b>"
	}
	return `<a href="` + esc(e.URL) + `">` + title + `</a>`
}

// impactSuffix appends the severity, unless the vendor never reports one.
func impactSuffix(e model.IndexEntry) string {
	if e.Impact == model.ImpactUnknown {
		return ""
	}
	return " · " + esc(string(e.Impact))
}

// sourceHealth says how the fetch went in a handful of characters. A dead
// source carries its age, because "datadog" alone is easy to read past and
// "datadog silent 16d" is not.
func sourceHealth(sources []model.SourceStatus, snapshot time.Time) string {
	var failed []string
	for _, s := range sources {
		if s.OK {
			continue
		}
		note := s.Vendor
		if s.LastSuccessAt != nil && !s.LastSuccessAt.IsZero() {
			note += fmt.Sprintf(" silent %dd", int(snapshot.UTC().Sub(s.LastSuccessAt.UTC()).Hours()/24))
		}
		failed = append(failed, note)
	}
	if len(failed) == 0 {
		return fmt.Sprintf("%d/%d sources ok", len(sources), len(sources))
	}
	return fmt.Sprintf("<b>%d/%d sources ok</b> (failed: %s)",
		len(sources)-len(failed), len(sources), esc(strings.Join(failed, ", ")))
}

// parseErrorSummary is the one-line form of the digest's parse warning.
func parseErrorSummary(sources []model.SourceStatus) string {
	total := 0
	for _, s := range sources {
		total += s.ParseErrors
	}
	if total == 0 {
		return ""
	}
	return fmt.Sprintf("<b>%d record(s) dropped as unparseable</b> — a vendor likely changed its schema.", total)
}

func esc(s string) string { return html.EscapeString(s) }
