package source

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/jtprogru/weekly-incident/internal/model"
)

// slack parses slack-status.com/api/v2.0.0/history, an undocumented endpoint
// that returns a flat array of incidents. It reports no severity, so every
// Slack incident carries ImpactUnknown and clears the digest threshold on
// duration alone.
type slack struct {
	base
}

type slackIncident struct {
	ID          json.Number `json:"id"`
	Title       string      `json:"title"`
	Type        string      `json:"type"`
	Status      string      `json:"status"`
	URL         string      `json:"url"`
	DateCreated string      `json:"date_created"`
	DateUpdated string      `json:"date_updated"`
	Services    []string    `json:"services"`
	Notes       []struct {
		Body        string `json:"body"`
		DateCreated string `json:"date_created"`
	} `json:"notes"`
}

func (s *slack) Fetch(ctx context.Context, f *Fetcher) (Result, error) {
	b, err := f.Get(ctx, s.url)
	if err != nil {
		return Result{}, err
	}
	return s.parse(b)
}

func (s *slack) parse(b []byte) (Result, error) {
	var raws []json.RawMessage
	if err := json.Unmarshal(b, &raws); err != nil {
		return Result{}, fmt.Errorf("%s: decode payload: %w", s.vendor, err)
	}

	var res Result
	for _, raw := range raws {
		in, err := s.normalize(raw)
		if err != nil {
			res.ParseErrors++
			continue
		}
		res.Incidents = append(res.Incidents, in)
	}
	return res, nil
}

func (s *slack) normalize(raw json.RawMessage) (model.Incident, error) {
	var si slackIncident
	if err := json.Unmarshal(raw, &si); err != nil {
		return model.Incident{}, err
	}

	id := si.ID.String()
	if id == "" {
		return model.Incident{}, fmt.Errorf("slack incident without id")
	}
	if _, err := strconv.ParseInt(id, 10, 64); err != nil {
		return model.Incident{}, fmt.Errorf("slack id %q is not numeric: %w", id, err)
	}

	started, err := time.Parse(time.RFC3339, si.DateCreated)
	if err != nil {
		return model.Incident{}, fmt.Errorf("slack %s: parse date_created: %w", id, err)
	}

	status := slackStatus(si.Status)
	var resolved *time.Time
	if status == model.StatusResolved && si.DateUpdated != "" {
		t, err := time.Parse(time.RFC3339, si.DateUpdated)
		if err != nil {
			return model.Incident{}, fmt.Errorf("slack %s: parse date_updated: %w", id, err)
		}
		resolved = &t
	}

	updates := make([]model.Update, 0, len(si.Notes))
	for _, n := range si.Notes {
		t, err := time.Parse(time.RFC3339, n.DateCreated)
		if err != nil {
			// One unreadable note must not cost the whole incident.
			continue
		}
		updates = append(updates, model.Update{
			// Notes carry no status of their own.
			Status:    model.StatusUnknown,
			Body:      stripHTML(n.Body),
			CreatedAt: t,
		})
	}
	sort.SliceStable(updates, func(i, j int) bool {
		return updates[i].CreatedAt.Before(updates[j].CreatedAt)
	})

	components := make([]string, 0, len(si.Services))
	for _, sv := range si.Services {
		if sv = strings.TrimSpace(sv); sv != "" {
			components = append(components, sv)
		}
	}

	in := model.Incident{
		Vendor:     s.vendor,
		NativeID:   id,
		Title:      strings.TrimSpace(si.Title),
		URL:        si.URL,
		Impact:     model.ImpactUnknown,
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

// slackStatus maps the two values the feed uses. Anything else is left unknown
// rather than guessed at.
func slackStatus(s string) model.Status {
	switch s {
	case "active":
		return model.StatusInvestigating
	case "resolved":
		return model.StatusResolved
	default:
		return model.StatusUnknown
	}
}
