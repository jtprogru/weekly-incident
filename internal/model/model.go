// Package model defines the canonical incident types shared by every source
// parser, by the archive and by the renderers.
//
// Every field is stored exactly as the vendor reports it, in English. Times are
// normalized to UTC and serialized as RFC 3339. Nothing here is translated or
// summarized: the collector's job is to preserve, not to interpret.
package model

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// Impact is the severity a vendor assigns to an incident. Vendors that do not
// report severity in a machine-readable form get ImpactUnknown.
type Impact string

const (
	ImpactNone     Impact = "none"
	ImpactMinor    Impact = "minor"
	ImpactMajor    Impact = "major"
	ImpactCritical Impact = "critical"
	ImpactUnknown  Impact = "unknown"
)

// impactRank orders impacts for threshold comparisons. ImpactUnknown ranks
// below everything: Slack and AWS never report impact, and treating their
// incidents as severe by default would flood the digest.
var impactRank = map[Impact]int{
	ImpactUnknown:  -1,
	ImpactNone:     0,
	ImpactMinor:    1,
	ImpactMajor:    2,
	ImpactCritical: 3,
}

// AtLeast reports whether i is at least as severe as min.
func (i Impact) AtLeast(min Impact) bool {
	return impactRank[i] >= impactRank[min]
}

// ParseImpact maps a vendor-reported impact string onto the canonical set.
func ParseImpact(s string) Impact {
	switch Impact(s) {
	case ImpactNone, ImpactMinor, ImpactMajor, ImpactCritical:
		return Impact(s)
	default:
		return ImpactUnknown
	}
}

// Status is the lifecycle stage of an incident.
type Status string

const (
	StatusInvestigating Status = "investigating"
	StatusIdentified    Status = "identified"
	StatusMonitoring    Status = "monitoring"
	StatusResolved      Status = "resolved"
	StatusPostmortem    Status = "postmortem"
	StatusUnknown       Status = "unknown"
)

// ParseStatus maps a vendor-reported status string onto the canonical set.
func ParseStatus(s string) Status {
	switch Status(s) {
	case StatusInvestigating, StatusIdentified, StatusMonitoring, StatusResolved, StatusPostmortem:
		return Status(s)
	default:
		return StatusUnknown
	}
}

// Update is a single message a vendor posted against an incident.
type Update struct {
	NativeID  string    `json:"native_id,omitempty"`
	Status    Status    `json:"status"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

// Incident is the canonical aggregate. Key is the only merge key in the
// project; Raw holds the vendor's original object so a normalization mistake
// found later can be repaired without re-crawling.
//
// There is deliberately no "last seen" field. It would change on every run and
// rewrite every file daily, turning the archive's git history into noise and
// making it impossible to tell a real update from a heartbeat.
type Incident struct {
	Key             string          `json:"key"`
	Vendor          string          `json:"vendor"`
	NativeID        string          `json:"native_id"`
	Title           string          `json:"title"`
	URL             string          `json:"url"`
	Impact          Impact          `json:"impact"`
	Status          Status          `json:"status"`
	StartedAt       time.Time       `json:"started_at"`
	ResolvedAt      *time.Time      `json:"resolved_at,omitempty"`
	DurationMinutes *int            `json:"duration_minutes,omitempty"`
	Components      []string        `json:"components"`
	Updates         []Update        `json:"updates"`
	FirstSeenAt     time.Time       `json:"first_seen_at"`
	Raw             json.RawMessage `json:"raw"`
}

// MakeKey builds the archive merge key for a vendor-native identifier pair.
func MakeKey(vendor, nativeID string) string {
	return vendor + "/" + nativeID
}

// Finalize fills the derived fields a parser should not have to compute by
// hand: the merge key and the resolved duration. It is called by every source
// parser before handing an incident to the archive.
//
// DurationMinutes stays nil while ResolvedAt is nil. "Ran for zero minutes" and
// "still running" are different facts and the score treats them differently.
func (in *Incident) Finalize() {
	in.Key = MakeKey(in.Vendor, in.NativeID)
	in.StartedAt = in.StartedAt.UTC()
	if in.ResolvedAt != nil {
		r := in.ResolvedAt.UTC()
		in.ResolvedAt = &r
		d := int(r.Sub(in.StartedAt).Minutes())
		if d < 0 {
			d = 0
		}
		in.DurationMinutes = &d
	} else {
		in.DurationMinutes = nil
	}
	for i := range in.Updates {
		in.Updates[i].CreatedAt = in.Updates[i].CreatedAt.UTC()
	}
	if in.Components == nil {
		in.Components = []string{}
	}
	if in.Updates == nil {
		in.Updates = []Update{}
	}
	if canonical, err := canonicalJSON(in.Raw); err == nil {
		in.Raw = canonical
	}
}

// canonicalJSON rewrites a vendor payload with sorted keys and no insignificant
// whitespace.
//
// Without this the archive would churn: a vendor that reorders its JSON keys
// between responses, or switches pretty-printing on, would rewrite every file
// it touches without a single fact having changed. Numbers are decoded as
// json.Number so large integers keep their literal form instead of turning into
// floats.
func canonicalJSON(b []byte) ([]byte, error) {
	if len(b) == 0 {
		return b, nil
	}
	dec := json.NewDecoder(bytes.NewReader(b))
	dec.UseNumber()
	var v any
	if err := dec.Decode(&v); err != nil {
		return nil, err
	}
	return json.Marshal(v)
}

// Validate reports the first structural problem that would make an incident
// useless downstream. Parsers call it to reject junk before it reaches disk.
func (in *Incident) Validate() error {
	switch {
	case in.Vendor == "":
		return fmt.Errorf("empty vendor")
	case in.NativeID == "":
		return fmt.Errorf("empty native_id")
	case in.StartedAt.IsZero():
		return fmt.Errorf("zero started_at")
	}
	return nil
}

// DurationAt returns the incident duration in minutes, measuring an unresolved
// incident against now. The second result reports whether the incident is still
// ongoing, which the renderers flag in the digest.
func (in *Incident) DurationAt(now time.Time) (int, bool) {
	if in.DurationMinutes != nil {
		return *in.DurationMinutes, false
	}
	d := int(now.UTC().Sub(in.StartedAt).Minutes())
	if d < 0 {
		d = 0
	}
	return d, true
}

// Score ranks an incident for the digest. Every multiplier is stored alongside
// the result: the ordering has to be explainable to a reader without reading
// the code.
type Score struct {
	Value      float64 `json:"value"`
	Duration   int     `json:"duration_minutes"`
	Components int     `json:"components"`
	Weight     float64 `json:"vendor_weight"`
}

// SourceStatus records how one vendor fetch went during a collect run.
type SourceStatus struct {
	Vendor        string    `json:"vendor"`
	OK            bool      `json:"ok"`
	Error         string    `json:"error,omitempty"`
	FetchedAt     time.Time `json:"fetched_at"`
	IncidentsSeen int       `json:"incidents_seen"`
	ParseErrors   int       `json:"parse_errors"`
}

// IndexEntry is the digest-facing projection of an incident: everything the
// renderers need without loading Raw.
type IndexEntry struct {
	Key             string     `json:"key"`
	Vendor          string     `json:"vendor"`
	Title           string     `json:"title"`
	URL             string     `json:"url"`
	Impact          Impact     `json:"impact"`
	Status          Status     `json:"status"`
	StartedAt       time.Time  `json:"started_at"`
	ResolvedAt      *time.Time `json:"resolved_at,omitempty"`
	DurationMinutes *int       `json:"duration_minutes,omitempty"`
	Ongoing         bool       `json:"ongoing"`
	Components      []string   `json:"components"`
	Score           Score      `json:"score"`
	AboveThreshold  bool       `json:"above_threshold"`
}

// WeekIndex is the per-week summary written to data/YYYY/Wnn/index.json. It is
// fully derived from the vendor files and rebuilt from scratch on every run.
type WeekIndex struct {
	Week        string         `json:"week"`
	From        time.Time      `json:"from"`
	To          time.Time      `json:"to"`
	GeneratedAt time.Time      `json:"generated_at"`
	Sources     []SourceStatus `json:"sources"`
	Incidents   []IndexEntry   `json:"incidents"`
}
