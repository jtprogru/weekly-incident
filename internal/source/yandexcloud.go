package source

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/jtprogru/weekly-incident/internal/model"
)

// yandexCloudIncidentURL prefixes the incident id; the feed carries no link.
const yandexCloudIncidentURL = "https://status.yandex.cloud/en/incidents/"

// yandexCloud parses status.yandex.cloud/api/incidents, the undocumented
// endpoint behind Yandex Cloud's own status board.
//
// Depth is a number in the configured URL, not a loop here: page is
// cumulative, so page=N answers with the newest 10×(N+1) incidents in one
// response. The URL must not carry installation=ru. That filter hides every
// incident whose installations list is empty, and the zone-wide outages are
// exactly the ones filed that way.
type yandexCloud struct {
	base
}

type ycPayload struct {
	Items []json.RawMessage `json:"items"`
}

type ycIncident struct {
	ID                  json.Number `json:"id"`
	Title               string      `json:"title"`
	Status              string      `json:"status"`
	StartDate           *time.Time  `json:"startDate"`
	EndDate             *time.Time  `json:"endDate"`
	LevelID             int         `json:"levelId"`
	MaxLevelID          int         `json:"maxLevelId"`
	Report              string      `json:"report"`
	IsReportPublished   bool        `json:"isReportPublished"`
	ReportPublishedTime *time.Time  `json:"reportPublishedTime"`
	Services            []struct {
		Name string `json:"name"`
	} `json:"services"`
	Comments []struct {
		ID        json.Number `json:"id"`
		Type      string      `json:"type"`
		Content   string      `json:"content"`
		CreatedAt *time.Time  `json:"createdAt"`
	} `json:"comments"`
}

func (y *yandexCloud) Fetch(ctx context.Context, f *Fetcher) (Result, error) {
	b, err := f.Get(ctx, y.url)
	if err != nil {
		return Result{}, err
	}
	return y.parse(b)
}

func (y *yandexCloud) parse(b []byte) (Result, error) {
	var p ycPayload
	if err := json.Unmarshal(b, &p); err != nil {
		return Result{}, fmt.Errorf("%s: decode payload: %w", y.vendor, err)
	}
	// The window always holds the newest incidents, so a payload without the
	// array is not a quiet week: something other than the feed answered.
	if p.Items == nil {
		return Result{}, fmt.Errorf("%s: payload has no items", y.vendor)
	}

	var res Result
	for _, raw := range p.Items {
		in, err := y.normalize(raw)
		if err != nil {
			res.ParseErrors++
			continue
		}
		res.Incidents = append(res.Incidents, in)
	}
	return res, nil
}

func (y *yandexCloud) normalize(raw json.RawMessage) (model.Incident, error) {
	var yi ycIncident
	if err := json.Unmarshal(raw, &yi); err != nil {
		return model.Incident{}, err
	}

	id := yi.ID.String()
	if id == "" {
		return model.Incident{}, fmt.Errorf("yandex cloud incident without id")
	}
	if yi.StartDate == nil {
		return model.Incident{}, fmt.Errorf("yandex cloud %s: no startDate", id)
	}

	updates := make([]model.Update, 0, len(yi.Comments)+1)
	var lastComment *time.Time
	for _, c := range yi.Comments {
		if c.CreatedAt == nil {
			continue
		}
		// createdAt, not updatedAt: the board edits a comment in place, and
		// the edit must not move it along the timeline.
		updates = append(updates, model.Update{
			NativeID:  c.ID.String(),
			Status:    ycCommentStatus(c.Type),
			Body:      stripHTML(c.Content),
			CreatedAt: *c.CreatedAt,
		})
		if lastComment == nil || c.CreatedAt.After(*lastComment) {
			lastComment = c.CreatedAt
		}
	}

	status := ycStatus(yi.Status)
	resolved := yi.EndDate
	// The board can close an incident without an end date, as it did when one
	// outage was folded into a larger one. Left alone, a resolved incident
	// would count as ongoing forever, so the last word the vendor said on it
	// stands in for the end.
	if status == model.StatusResolved && resolved == nil {
		resolved = lastComment
	}

	// A published report is the postmortem. It becomes the final update and
	// moves the incident to the same status Statuspage uses for it.
	if yi.IsReportPublished && yi.Report != "" && yi.ReportPublishedTime != nil {
		updates = append(updates, model.Update{
			Status:    model.StatusPostmortem,
			Body:      stripHTML(yi.Report),
			CreatedAt: *yi.ReportPublishedTime,
		})
		if status == model.StatusResolved {
			status = model.StatusPostmortem
		}
	}
	// The feed lists comments newest first; the archive stores them in the
	// order they happened.
	sort.SliceStable(updates, func(i, j int) bool {
		return updates[i].CreatedAt.Before(updates[j].CreatedAt)
	})

	components := make([]string, 0, len(yi.Services))
	for _, s := range yi.Services {
		if name := strings.TrimSpace(s.Name); name != "" {
			components = append(components, name)
		}
	}

	// levelId is where the incident stands now and drops as it winds down;
	// maxLevelId is its peak, which is what the board itself shows once an
	// incident is resolved.
	level := yi.MaxLevelID
	if level == 0 {
		level = yi.LevelID
	}

	stripped, err := ycStripIcons(raw)
	if err != nil {
		return model.Incident{}, err
	}

	in := model.Incident{
		Vendor:     y.vendor,
		NativeID:   id,
		Title:      strings.TrimSpace(yi.Title),
		URL:        yandexCloudIncidentURL + id,
		Impact:     ycImpact(level),
		Status:     status,
		StartedAt:  *yi.StartDate,
		ResolvedAt: resolved,
		Components: components,
		Updates:    updates,
		Raw:        stripped,
	}
	in.Finalize()
	if err := in.Validate(); err != nil {
		return model.Incident{}, err
	}
	return in, nil
}

// ycStripIcons drops the inline SVG icon every affected service carries. It is
// decoration, it made up three quarters of a measured response, and a redesign
// on the vendor's side would rewrite the archive without a single fact having
// changed. Everything else in the payload is kept.
func ycStripIcons(raw json.RawMessage) (json.RawMessage, error) {
	dec := json.NewDecoder(bytes.NewReader(raw))
	dec.UseNumber()
	var v map[string]any
	if err := dec.Decode(&v); err != nil {
		return nil, err
	}
	services, _ := v["services"].([]any)
	for _, s := range services {
		if m, ok := s.(map[string]any); ok {
			delete(m, "icon")
		}
	}
	return json.Marshal(v)
}

// ycImpact maps the board's level ids. Only 1 (Minor) and 2 (Unavailable) have
// ever been published, across 489 incidents since 2020; 3 (Medium) and 4
// (Critical) come from the enum in the board's frontend. Unavailable is the top
// level in actual use and covers anything from one service to a whole zone, so
// it lands on major and critical is left to the level that carries the name.
func ycImpact(level int) model.Impact {
	switch level {
	case 1:
		return model.ImpactMinor
	case 2, 3:
		return model.ImpactMajor
	case 4:
		return model.ImpactCritical
	default:
		return model.ImpactUnknown
	}
}

// ycStatus maps the two values the feed uses. Anything else is left unknown
// rather than guessed at.
func ycStatus(s string) model.Status {
	switch s {
	case "open":
		return model.StatusInvestigating
	case "resolved":
		return model.StatusResolved
	default:
		return model.StatusUnknown
	}
}

// ycCommentStatus maps a comment type. "update" marks progress without naming
// a stage, so it stays unknown along with anything unseen.
func ycCommentStatus(t string) model.Status {
	switch t {
	case "investigation":
		return model.StatusInvestigating
	case "resolved":
		return model.StatusResolved
	default:
		return model.StatusUnknown
	}
}
