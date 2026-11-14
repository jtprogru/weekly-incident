package render

import (
	"fmt"
	"strings"

	"github.com/jtprogru/weekly-incident/internal/model"
)

func writeDedupCounts(b *strings.Builder, s *model.DedupStatistics) {
	if s == nil {
		return
	}
	fmt.Fprintf(b, "%d source records (%d saved revisions) → %d canonical incidents; %d repeated records grouped; %d candidate pairs pending review. Rules: `%s`. Counts refer to collected vendor reports, not all real-world outages.\n\n", s.SourceRecords, s.SourceVersions, s.Incidents, s.MergedRecords, s.PendingPairs, s.RulesVersion)
}

func writeCanonicalCatalogue(b *strings.Builder, entries []model.IndexEntry) {
	if len(entries) == 0 || entries[0].Canonical == nil {
		return
	}
	b.WriteString("## Identity and evidence\n\nSource record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.\n\n")
	for _, e := range entries {
		writeCanonicalInfo(b, e.Canonical)
	}
}

func writeCanonicalInfo(b *strings.Builder, c *model.CanonicalInfo) {
	if c == nil {
		return
	}
	fmt.Fprintf(b, "<a id=\"%s\"></a>\n\n- Canonical ID: `%s` (%d source records).\n", c.ID, c.ID, len(c.Sources))
	for _, s := range c.Sources {
		fmt.Fprintf(b, "  - %s; vendor URL: %s\n", link(s.Key, "../../data/"+s.Path), s.URL)
	}
	if len(c.Reasons) > 0 {
		fmt.Fprintf(b, "- Identity evidence: %s.\n", strings.Join(c.Reasons, ", "))
	}
	fmt.Fprintf(b, "- Start time basis: `%s`.\n", c.StartQuality)
	for _, w := range c.Warnings {
		fmt.Fprintf(b, "- Warning: %s\n", oneLine(w))
	}
	b.WriteString("\n")
}
