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

// statuspage parses the Atlassian Statuspage v2 schema, which ten of the
// thirteen configured vendors serve verbatim.
//
// The feed has no pagination: ?page=N returns the same bytes, and history.json
// does not exist. Depth is whatever fits in the vendor's incident window, which
// ranged from 20 days (Cloudflare) to 62 days (GitHub) when measured on
// 2026-08-18. That is why the archive merges instead of rebuilding a window.
type statuspage struct {
	base
}

type spPayload struct {
	Page struct {
		URL string `json:"url"`
	} `json:"page"`
	Incidents []json.RawMessage `json:"incidents"`
}

type spIncident struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	Shortlink  string     `json:"shortlink"`
	Impact     string     `json:"impact"`
	Status     string     `json:"status"`
	CreatedAt  *time.Time `json:"created_at"`
	StartedAt  *time.Time `json:"started_at"`
	ResolvedAt *time.Time `json:"resolved_at"`
	Components []struct {
		Name string `json:"name"`
	} `json:"components"`
	IncidentUpdates []struct {
		ID        string     `json:"id"`
		Status    string     `json:"status"`
		Body      string     `json:"body"`
		CreatedAt *time.Time `json:"created_at"`
	} `json:"incident_updates"`
}

func (s *statuspage) Fetch(ctx context.Context, f *Fetcher) (Result, error) {
	b, err := f.Get(ctx, s.url)
	if err != nil {
		return Result{}, err
	}
	return s.parse(b)
}

func (s *statuspage) parse(b []byte) (Result, error) {
	var p spPayload
	if err := json.Unmarshal(b, &p); err != nil {
		return Result{}, fmt.Errorf("%s: decode payload: %w", s.vendor, err)
	}

	var res Result
	for _, raw := range p.Incidents {
		in, err := s.normalize(raw, p.Page.URL)
		if err != nil {
			res.ParseErrors++
			continue
		}
		res.Incidents = append(res.Incidents, in)
	}
	return res, nil
}

func (s *statuspage) normalize(raw json.RawMessage, pageURL string) (model.Incident, error) {
	var sp spIncident
	if err := json.Unmarshal(raw, &sp); err != nil {
		return model.Incident{}, err
	}

	// started_at is the real onset; created_at is when the vendor opened the
	// record. Prefer the former, fall back to the latter.
	started := sp.StartedAt
	if started == nil {
		started = sp.CreatedAt
	}
	if started == nil {
		return model.Incident{}, fmt.Errorf("incident %q: no started_at or created_at", sp.ID)
	}

	components := make([]string, 0, len(sp.Components))
	for _, c := range sp.Components {
		if c.Name != "" {
			components = append(components, c.Name)
		}
	}

	updates := make([]model.Update, 0, len(sp.IncidentUpdates))
	for _, u := range sp.IncidentUpdates {
		if u.CreatedAt == nil {
			continue
		}
		updates = append(updates, model.Update{
			NativeID:  u.ID,
			Status:    model.ParseStatus(u.Status),
			Body:      strings.TrimSpace(u.Body),
			CreatedAt: *u.CreatedAt,
		})
	}
	// Statuspage returns updates newest first; the archive stores them in the
	// order they happened.
	sort.SliceStable(updates, func(i, j int) bool {
		return updates[i].CreatedAt.Before(updates[j].CreatedAt)
	})

	in := model.Incident{
		Vendor:     s.vendor,
		NativeID:   sp.ID,
		Title:      strings.TrimSpace(sp.Name),
		URL:        spURL(sp, pageURL),
		Impact:     model.ParseImpact(sp.Impact),
		Status:     model.ParseStatus(sp.Status),
		StartedAt:  *started,
		ResolvedAt: sp.ResolvedAt,
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

// spURL picks the incident link. shortlink is present on most vendors but not
// all, so fall back to the canonical page path.
func spURL(sp spIncident, pageURL string) string {
	if sp.Shortlink != "" {
		return sp.Shortlink
	}
	if pageURL != "" {
		return strings.TrimSuffix(pageURL, "/") + "/incidents/" + sp.ID
	}
	return ""
}
