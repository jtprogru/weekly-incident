package canonical

import (
	"time"

	"github.com/jtprogru/weekly-incident/internal/archive"
	"github.com/jtprogru/weekly-incident/internal/config"
	"github.com/jtprogru/weekly-incident/internal/model"
	"github.com/jtprogru/weekly-incident/internal/score"
)

// WeekComparison captures ranking changes at the same clock and configuration.
// Before keys name source records; After keys name canonical incidents.
type WeekComparison struct {
	Week               string   `json:"week"`
	SourceRecords      int      `json:"source_records"`
	CanonicalIncidents int      `json:"canonical_incidents"`
	AboveBefore        int      `json:"above_threshold_before"`
	AboveAfter         int      `json:"above_threshold_after"`
	TopBefore          []string `json:"top_before"`
	TopAfter           []string `json:"top_after"`
}

func compareWeeks(inv archive.Inventory, incidents []Incident, cfg *config.Config, now time.Time, weeks []string) []WeekComparison {
	before, after := map[string][]model.Incident{}, map[string][]model.Incident{}
	for _, r := range inv.Records {
		w := archive.WeekOf(r.Incident.StartedAt).String()
		before[w] = append(before[w], r.Incident)
	}
	for _, c := range incidents {
		w := archive.WeekOf(c.Fields.StartedAt).String()
		after[w] = append(after[w], c.Projection())
	}
	out := []WeekComparison{}
	for _, id := range weeks {
		w, _ := archive.ParseWeek(id)
		b, a := score.BuildIndex(w, before[id], nil, cfg, now), score.BuildIndex(w, after[id], nil, cfg, now)
		c := WeekComparison{Week: id, SourceRecords: len(b.Incidents), CanonicalIncidents: len(a.Incidents), TopBefore: []string{}, TopAfter: []string{}}
		for _, e := range b.Incidents {
			if e.AboveThreshold {
				c.AboveBefore++
				if len(c.TopBefore) < cfg.Digest.BriefTopN {
					c.TopBefore = append(c.TopBefore, e.Key)
				}
			}
		}
		for _, e := range a.Incidents {
			if e.AboveThreshold {
				c.AboveAfter++
				if len(c.TopAfter) < cfg.Digest.BriefTopN {
					c.TopAfter = append(c.TopAfter, e.Key)
				}
			}
		}
		out = append(out, c)
	}
	return out
}
