package source

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/jtprogru/weekly-incident/internal/model"
)

// gcpBaseURL prefixes the relative uri the feed returns ("incidents/<id>").
const gcpBaseURL = "https://status.cloud.google.com/"

// gcpTitleLimit caps the synthesized title. GCP ships a paragraph where a title
// belongs; the full text stays in Raw.
const gcpTitleLimit = 200

// gcp parses status.cloud.google.com/incidents.json. Google publishes only
// major incidents here — four in the half-year before 2026-08-18 — but each one
// arrives with a full postmortem, which is exactly the material this project
// exists for.
type gcp struct {
	base
}

type gcpIncident struct {
	ID           string `json:"id"`
	Number       string `json:"number"`
	Begin        string `json:"begin"`
	Created      string `json:"created"`
	End          string `json:"end"`
	ExternalDesc string `json:"external_desc"`
	StatusImpact string `json:"status_impact"`
	Severity     string `json:"severity"`
	ServiceKey   string `json:"service_key"`
	ServiceName  string `json:"service_name"`
	URI          string `json:"uri"`
	Updates      []struct {
		When   string `json:"when"`
		Text   string `json:"text"`
		Status string `json:"status"`
	} `json:"updates"`
	AffectedProducts []struct {
		Title string `json:"title"`
	} `json:"affected_products"`
}

func (g *gcp) Fetch(ctx context.Context, f *Fetcher) (Result, error) {
	b, err := f.Get(ctx, g.url)
	if err != nil {
		return Result{}, err
	}
	return g.parse(b)
}

func (g *gcp) parse(b []byte) (Result, error) {
	var raws []json.RawMessage
	if err := json.Unmarshal(b, &raws); err != nil {
		return Result{}, fmt.Errorf("%s: decode payload: %w", g.vendor, err)
	}

	var res Result
	for _, raw := range raws {
		in, err := g.normalize(raw)
		if err != nil {
			res.ParseErrors++
			continue
		}
		res.Incidents = append(res.Incidents, in)
	}
	return res, nil
}

func (g *gcp) normalize(raw json.RawMessage) (model.Incident, error) {
	var gi gcpIncident
	if err := json.Unmarshal(raw, &gi); err != nil {
		return model.Incident{}, err
	}
	if gi.ID == "" {
		return model.Incident{}, fmt.Errorf("gcp incident without id")
	}

	// begin is the onset of impact; created is when Google opened the record.
	started, err := time.Parse(time.RFC3339, gi.Begin)
	if err != nil {
		started, err = time.Parse(time.RFC3339, gi.Created)
		if err != nil {
			return model.Incident{}, fmt.Errorf("gcp %s: no parseable begin/created: %w", gi.ID, err)
		}
	}

	var resolved *time.Time
	status := model.StatusInvestigating
	if gi.End != "" {
		t, err := time.Parse(time.RFC3339, gi.End)
		if err != nil {
			return model.Incident{}, fmt.Errorf("gcp %s: parse end: %w", gi.ID, err)
		}
		resolved = &t
		status = model.StatusResolved
	}

	components := make([]string, 0, len(gi.AffectedProducts))
	for _, p := range gi.AffectedProducts {
		if p.Title != "" {
			components = append(components, p.Title)
		}
	}
	if len(components) == 0 && gi.ServiceName != "" {
		components = append(components, gi.ServiceName)
	}

	updates := make([]model.Update, 0, len(gi.Updates))
	for _, u := range gi.Updates {
		t, err := time.Parse(time.RFC3339, u.When)
		if err != nil {
			continue
		}
		updates = append(updates, model.Update{
			Status: gcpUpdateStatus(u.Status),
			// Update text is Markdown with escaped hashes; it is kept verbatim.
			// Un-escaping belongs to the renderer, not to the collector.
			Body:      strings.TrimSpace(u.Text),
			CreatedAt: t,
		})
	}
	sort.SliceStable(updates, func(i, j int) bool {
		return updates[i].CreatedAt.Before(updates[j].CreatedAt)
	})

	in := model.Incident{
		Vendor:     g.vendor,
		NativeID:   gi.ID,
		Title:      firstSentence(gi.ExternalDesc, gcpTitleLimit),
		URL:        gcpBaseURL + strings.TrimPrefix(gi.URI, "/"),
		Impact:     gcpImpact(gi.Severity, gi.StatusImpact),
		Status:     status,
		StartedAt:  started,
		ResolvedAt: resolved,
		Components: components,
		Updates:    updates,
		Raw:        raw,
	}
	in.Finalize()
	if err := in.Validate(); err != nil {
		return model.Incident{}, err
	}
	return in, nil
}

// gcpImpact folds Google's two severity axes into one. status_impact
// SERVICE_INFORMATION means the entry is advisory and outranks severity.
func gcpImpact(severity, statusImpact string) model.Impact {
	if statusImpact == "SERVICE_INFORMATION" {
		return model.ImpactNone
	}
	switch severity {
	case "high":
		return model.ImpactCritical
	case "medium":
		return model.ImpactMajor
	case "low":
		return model.ImpactMinor
	default:
		return model.ImpactUnknown
	}
}

// gcpUpdateStatus maps the three values seen in the feed. Anything else stays
// unknown rather than being guessed at.
func gcpUpdateStatus(s string) model.Status {
	switch s {
	case "AVAILABLE":
		return model.StatusResolved
	case "SERVICE_DISRUPTION", "SERVICE_OUTAGE":
		return model.StatusInvestigating
	default:
		return model.StatusUnknown
	}
}
