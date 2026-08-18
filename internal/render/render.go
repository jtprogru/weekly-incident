// Package render turns a week of archived incidents into the two Markdown
// files a human works from.
//
// Everything here is English and deterministic. No network, no LLM, no
// translation: the collector's output has to be reproducible from the archive
// alone, and the Russian text of the actual post is written by hand elsewhere.
package render

import (
	"fmt"
	"strings"

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

	fmt.Fprintf(&b, "%d incidents collected, %d above threshold. %s\n\n",
		len(idx.Incidents), len(above), sourceSummary(idx.Sources))

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

	writeSources(&b, idx.Sources)
	return b.String()
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

	for i, e := range top {
		in := incidentByKey[e.Key]
		fmt.Fprintf(&b, "## %d. %s — %s\n\n", i+1, e.Vendor, e.Title)
		fmt.Fprintf(&b, "- URL: %s\n", e.URL)
		fmt.Fprintf(&b, "- Window: %s to %s (%s)\n", e.StartedAt.Format(timeFormat), resolvedCell(e), duration(e))
		fmt.Fprintf(&b, "- Impact: %s, status: %s\n", e.Impact, e.Status)
		fmt.Fprintf(&b, "- Components: %s\n\n", components(e.Components))
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
	return b.String()
}

func writeTimeline(b *strings.Builder, n int, e model.IndexEntry, in model.Incident) {
	fmt.Fprintf(b, "### %d. %s — %s\n\n", n, e.Vendor, e.Title)
	fmt.Fprintf(b, "- URL: %s\n", e.URL)
	fmt.Fprintf(b, "- Started: %s UTC\n", e.StartedAt.Format(timeFormat))
	fmt.Fprintf(b, "- Resolved: %s\n", resolvedCell(e))
	fmt.Fprintf(b, "- Duration: %s\n", duration(e))
	fmt.Fprintf(b, "- Impact: %s, status: %s\n", e.Impact, e.Status)
	fmt.Fprintf(b, "- Components: %s\n", components(e.Components))
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
	b.WriteString("| Vendor | Status | Incidents seen | Parse errors |\n")
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

func sourceSummary(sources []model.SourceStatus) string {
	var failed []string
	for _, s := range sources {
		if !s.OK {
			failed = append(failed, s.Vendor)
		}
	}
	if len(failed) == 0 {
		return fmt.Sprintf("All %d sources responded.", len(sources))
	}
	return fmt.Sprintf("%d of %d sources unavailable: %s.",
		len(failed), len(sources), strings.Join(failed, ", "))
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
		return humanMinutes(e.Score.Duration) + " (ongoing)"
	}
	if e.DurationMinutes == nil {
		return "unknown"
	}
	return humanMinutes(*e.DurationMinutes)
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
		return "not yet"
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
