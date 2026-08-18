// Package score ranks incidents for the weekly digest and projects them into
// the week index.
//
// The score sorts and explains; it does not decide. Which incident becomes the
// story of the week is an editorial call made by a human reading the digest,
// so every multiplier is carried alongside the result and printed next to the
// row that it produced.
package score

import (
	"sort"
	"time"

	"github.com/jtprogru/weekly-incident/internal/archive"
	"github.com/jtprogru/weekly-incident/internal/config"
	"github.com/jtprogru/weekly-incident/internal/model"
)

// For scores one incident. An unresolved incident is measured against now, so
// an outage still running at digest time ranks by how long it has already run.
func For(in model.Incident, weight float64, now time.Time) model.Score {
	duration, _ := in.DurationAt(now)
	components := len(in.Components)
	if components < 1 {
		// A vendor that lists no components has still broken something.
		components = 1
	}
	if weight <= 0 {
		weight = 1
	}
	return model.Score{
		Value:      float64(duration) * float64(components) * weight,
		Duration:   duration,
		Components: components,
		Weight:     weight,
	}
}

// AboveThreshold reports whether an incident earns a full entry in the digest.
//
// The two conditions are a disjunction on purpose: a four-minute critical
// outage says more than a two-day minor degradation, and requiring both would
// hide the first.
func AboveThreshold(in model.Incident, d config.Digest, now time.Time) bool {
	duration, _ := in.DurationAt(now)
	if duration >= d.MinDurationMinutes {
		return true
	}
	return in.Impact.AtLeast(d.MinImpact)
}

// BuildIndex projects a week's incidents into the summary the renderers read,
// sorted by score, highest first.
func BuildIndex(w archive.Week, incidents []model.Incident, sources []model.SourceStatus, c *config.Config, now time.Time) model.WeekIndex {
	weights := c.Weights()

	entries := make([]model.IndexEntry, 0, len(incidents))
	for _, in := range incidents {
		_, ongoing := in.DurationAt(now)
		entries = append(entries, model.IndexEntry{
			Key:             in.Key,
			Vendor:          in.Vendor,
			Title:           in.Title,
			URL:             in.URL,
			Impact:          in.Impact,
			Status:          in.Status,
			StartedAt:       in.StartedAt,
			ResolvedAt:      in.ResolvedAt,
			DurationMinutes: in.DurationMinutes,
			Ongoing:         ongoing,
			Components:      in.Components,
			Score:           For(in, weights[in.Vendor], now),
			AboveThreshold:  AboveThreshold(in, c.Digest, now),
		})
	}
	SortEntries(entries)

	sorted := append([]model.SourceStatus(nil), sources...)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i].Vendor < sorted[j].Vendor })

	return model.WeekIndex{
		Week:        w.String(),
		From:        w.Start(),
		To:          w.End(),
		GeneratedAt: now.UTC(),
		Sources:     sorted,
		Incidents:   entries,
	}
}

// SortEntries orders entries by score descending, with start time and key as
// tiebreakers so the output never depends on input order.
func SortEntries(entries []model.IndexEntry) {
	sort.Slice(entries, func(i, j int) bool {
		if entries[i].Score.Value != entries[j].Score.Value {
			return entries[i].Score.Value > entries[j].Score.Value
		}
		if !entries[i].StartedAt.Equal(entries[j].StartedAt) {
			return entries[i].StartedAt.Before(entries[j].StartedAt)
		}
		return entries[i].Key < entries[j].Key
	})
}
