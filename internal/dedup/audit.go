package dedup

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"slices"
	"sort"
	"strings"
	"time"

	"github.com/jtprogru/weekly-incident/internal/archive"
	"github.com/jtprogru/weekly-incident/internal/config"
	"github.com/jtprogru/weekly-incident/internal/model"
)

const RulesVersion = "audit-v1"

// RecordRef identifies an occurrence, not just a source key: a repeated key in
// another week or even the same file must remain inspectable.
type RecordRef struct {
	Key        string     `json:"key"`
	Vendor     string     `json:"vendor"`
	NativeID   string     `json:"native_id"`
	Path       string     `json:"path"`
	Position   int        `json:"position"`
	Title      string     `json:"title"`
	URL        string     `json:"url"`
	StartedAt  time.Time  `json:"started_at"`
	ResolvedAt *time.Time `json:"resolved_at,omitempty"`
	Components []string   `json:"components"`
	Updates    int        `json:"updates"`
}

type UpdateMatch struct {
	CreatedAt   time.Time `json:"created_at"`
	Fingerprint string    `json:"fingerprint"`
	Body        string    `json:"body"`
}

type Candidate struct {
	ID              string        `json:"id"`
	Left            RecordRef     `json:"left"`
	Right           RecordRef     `json:"right"`
	Level           string        `json:"level"`
	Reasons         []string      `json:"reasons"`
	Cautions        []string      `json:"cautions"`
	SharedUpdateIDs []string      `json:"shared_update_ids"`
	SharedUpdates   []UpdateMatch `json:"shared_updates"`
}

type Statistics struct {
	Records           int `json:"archive_records"`
	SourceKeys        int `json:"distinct_source_keys"`
	Vendors           int `json:"vendors"`
	Files             int `json:"files"`
	RepeatedKeyGroups int `json:"repeated_key_groups"`
	CandidatePairs    int `json:"candidate_pairs"`
	ConflictedPairs   int `json:"conflicted_pairs"`
}

// Report's hashes describe the bytes actually inspected. RevisionLabel is
// optional caller-supplied context; it never replaces the file fingerprints.
type Report struct {
	SchemaVersion    int                 `json:"schema_version"`
	RulesVersion     string              `json:"rules_version"`
	GeneratedAt      time.Time           `json:"generated_at"`
	RevisionLabel    string              `json:"revision_label,omitempty"`
	InputFingerprint string              `json:"input_fingerprint"`
	ConfigSHA256     string              `json:"config_sha256"`
	Files            []archive.InputFile `json:"files"`
	Policies         []Policy            `json:"policies"`
	Statistics       Statistics          `json:"statistics"`
	Warnings         []string            `json:"warnings"`
	Candidates       []Candidate         `json:"candidates"`
}

type features struct {
	record  archive.StoredIncident
	policy  Policy
	url     string
	updates map[string]UpdateMatch
	ids     map[string][]model.Update
}

// Audit builds a deterministic, read-only report. Every different source key
// remains review_required, however convincing its overlap looks. This stage
// deliberately does not create groups or a unique-incident count.
func Audit(inv archive.Inventory, sources []config.Source, now time.Time) Report {
	r := Report{
		SchemaVersion: 1, RulesVersion: RulesVersion, GeneratedAt: now.UTC(),
		InputFingerprint: inv.Fingerprint, Files: slices.Clone(inv.Files),
		Policies: []Policy{}, Candidates: []Candidate{}, Warnings: []string{},
	}
	policies := make(map[string]Policy, len(sources))
	for _, s := range sources {
		policies[s.Vendor] = policyFor(s)
		r.Policies = append(r.Policies, policyFor(s))
	}
	sort.Slice(r.Policies, func(i, j int) bool { return r.Policies[i].Vendor < r.Policies[j].Vendor })
	records := slices.Clone(inv.Records)
	sort.Slice(records, func(i, j int) bool {
		a, b := records[i], records[j]
		if a.Incident.Key != b.Incident.Key {
			return a.Incident.Key < b.Incident.Key
		}
		if a.Path != b.Path {
			return a.Path < b.Path
		}
		return a.Position < b.Position
	})
	fs := make([]features, len(records))
	keys, vendors := map[string]int{}, map[string]bool{}
	for i, rec := range records {
		in := rec.Incident
		keys[in.Key]++
		vendors[in.Vendor] = true
		fs[i] = extract(rec, policies[in.Vendor])
	}
	for vendor := range vendors {
		if _, ok := policies[vendor]; !ok {
			r.Warnings = append(r.Warnings, "unknown source policy: "+vendor+"; only keys and exact update content are compared")
		}
	}
	if len(records) == 0 {
		r.Warnings = append(r.Warnings, "archive contains no incident records")
	}
	sort.Strings(r.Warnings)
	r.Statistics = Statistics{Records: len(records), SourceKeys: len(keys), Vendors: len(vendors), Files: len(inv.Files)}
	for _, count := range keys {
		if count > 1 {
			r.Statistics.RepeatedKeyGroups++
		}
	}
	for _, pair := range candidatePairs(fs) {
		c := compare(fs[pair[0]], fs[pair[1]])
		r.Candidates = append(r.Candidates, c)
		if c.Level == "conflict" {
			r.Statistics.ConflictedPairs++
		}
	}
	r.Statistics.CandidatePairs = len(r.Candidates)
	return r
}

func extract(rec archive.StoredIncident, p Policy) features {
	f := features{record: rec, policy: p, url: eventURL(rec.Incident.URL, p), updates: map[string]UpdateMatch{}, ids: map[string][]model.Update{}}
	for _, u := range rec.Incident.Updates {
		body := normalizeBody(u.Body)
		if body != "" && !u.CreatedAt.IsZero() {
			key := fingerprint(u.CreatedAt.UTC().Format(time.RFC3339Nano), body)
			f.updates[key] = UpdateMatch{CreatedAt: u.CreatedAt.UTC(), Fingerprint: key, Body: body}
		}
		if p.UpdateIDs && u.NativeID != "" {
			f.ids[u.NativeID] = append(f.ids[u.NativeID], u)
		}
	}
	return f
}

// Inverted indexes avoid scanning every pair, without restricting evidence to
// one week. Vendor is part of every token, so cross-vendor collisions cannot
// become incident-identity evidence.
func candidatePairs(fs []features) [][2]int {
	type token struct{ vendor, kind, value string }
	index := map[token][]int{}
	pairs := map[[2]int]bool{}
	add := func(i int, kind, value string) {
		if value == "" {
			return
		}
		t := token{fs[i].record.Incident.Vendor, kind, value}
		for _, j := range index[t] {
			pairs[[2]int{j, i}] = true
		}
		index[t] = append(index[t], i)
	}
	for i, f := range fs {
		add(i, "key", f.record.Incident.Key)
		if f.policy.StableIncidentID {
			add(i, "source_id", f.record.Incident.NativeID)
		}
		add(i, "event_url", f.url)
		for key := range f.updates {
			add(i, "update", key)
		}
		for id := range f.ids {
			add(i, "update_id", id)
		}
	}
	out := make([][2]int, 0, len(pairs))
	for p := range pairs {
		out = append(out, p)
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i][0] != out[j][0] {
			return out[i][0] < out[j][0]
		}
		return out[i][1] < out[j][1]
	})
	return out
}

func compare(a, b features) Candidate {
	x, y := a.record.Incident, b.record.Incident
	c := Candidate{Left: reference(a.record), Right: reference(b.record), Level: "review_required",
		Reasons: []string{}, Cautions: []string{}, SharedUpdateIDs: []string{}, SharedUpdates: []UpdateMatch{}}
	c.ID = fingerprint(c.Left.Key, c.Left.Path, fmt.Sprint(c.Left.Position), c.Right.Key, c.Right.Path, fmt.Sprint(c.Right.Position))
	if x.Key == y.Key {
		c.Reasons = append(c.Reasons, "repeated_key")
	}
	if a.policy.StableIncidentID && x.NativeID == y.NativeID {
		c.Reasons = append(c.Reasons, "same_stable_source_id")
		c.Level = "exact_identity"
	}
	if a.url != "" && a.url == b.url {
		c.Reasons = append(c.Reasons, "same_event_url")
	}
	for id, left := range a.ids {
		if right, ok := b.ids[id]; ok {
			c.SharedUpdateIDs = append(c.SharedUpdateIDs, id)
			if !sameUpdateVersions(left, right) {
				c.Cautions = append(c.Cautions, "shared_update_id_has_different_content")
			}
		}
	}
	if len(c.SharedUpdateIDs) > 0 {
		c.Reasons = append(c.Reasons, "shared_update_ids")
		c.Cautions = append(c.Cautions, "update_id_scope_requires_review")
	}
	for key, u := range a.updates {
		if _, ok := b.updates[key]; ok {
			c.SharedUpdates = append(c.SharedUpdates, u)
		}
	}
	if len(c.SharedUpdates) > 0 {
		c.Reasons = append(c.Reasons, "exact_shared_updates")
	}
	if len(c.SharedUpdates) == 1 {
		c.Cautions = append(c.Cautions, "single_shared_update_may_be_template")
	}
	if len(c.SharedUpdates) > 0 && (len(c.SharedUpdates) == len(a.updates) || len(c.SharedUpdates) == len(b.updates)) {
		c.Reasons = append(c.Reasons, "update_set_containment")
	}
	ac, bc := canonicalStrings(x.Components), canonicalStrings(y.Components)
	if len(ac) == 0 || len(bc) == 0 {
		c.Cautions = append(c.Cautions, "components_unknown")
	} else if !intersects(ac, bc) {
		c.Cautions = append(c.Cautions, "disjoint_components")
		c.Level = "conflict"
	}
	if (x.ResolvedAt != nil && x.ResolvedAt.Before(y.StartedAt)) || (y.ResolvedAt != nil && y.ResolvedAt.Before(x.StartedAt)) {
		c.Cautions = append(c.Cautions, "non_overlapping_reported_windows")
		c.Level = "conflict"
	}
	if x.Key == y.Key && !sameRecordContent(x, y) {
		c.Cautions = append(c.Cautions, "repeated_key_content_differs_revision_order_unknown")
		c.Level = "conflict"
	}
	if a.policy.Kind == config.KindAWSRSS {
		c.Cautions = append(c.Cautions, "aws_start_is_first_visible_update")
	}
	c.Reasons, c.Cautions = canonicalStrings(c.Reasons), canonicalStrings(c.Cautions)
	sort.Strings(c.SharedUpdateIDs)
	sort.Slice(c.SharedUpdates, func(i, j int) bool {
		if !c.SharedUpdates[i].CreatedAt.Equal(c.SharedUpdates[j].CreatedAt) {
			return c.SharedUpdates[i].CreatedAt.Before(c.SharedUpdates[j].CreatedAt)
		}
		return c.SharedUpdates[i].Fingerprint < c.SharedUpdates[j].Fingerprint
	})
	return c
}

func reference(r archive.StoredIncident) RecordRef {
	in := r.Incident
	return RecordRef{Key: in.Key, Vendor: in.Vendor, NativeID: in.NativeID, Path: r.Path, Position: r.Position,
		Title: in.Title, URL: in.URL, StartedAt: in.StartedAt.UTC(), ResolvedAt: in.ResolvedAt,
		Components: canonicalStrings(in.Components), Updates: len(in.Updates)}
}

func normalizeBody(body string) string {
	return strings.TrimSpace(strings.ReplaceAll(body, "\r\n", "\n"))
}

func fingerprint(parts ...string) string {
	// Marshaling strings cannot fail; array framing avoids separator collisions.
	b, _ := json.Marshal(parts)
	sum := sha256.Sum256(b)
	return hex.EncodeToString(sum[:])
}

func canonicalStrings(in []string) []string {
	out := make([]string, 0, len(in))
	for _, s := range in {
		if s != "" {
			out = append(out, s)
		}
	}
	sort.Strings(out)
	return slices.Compact(out)
}

func intersects(a, b []string) bool {
	for _, s := range a {
		if slices.Contains(b, s) {
			return true
		}
	}
	return false
}

func sameUpdateVersions(a, b []model.Update) bool {
	encode := func(updates []model.Update) []string {
		out := make([]string, 0, len(updates))
		for _, u := range updates {
			out = append(out, fingerprint(u.CreatedAt.UTC().Format(time.RFC3339Nano), normalizeBody(u.Body), string(u.Status)))
		}
		return canonicalStrings(out)
	}
	return slices.Equal(encode(a), encode(b))
}

func sameRecordContent(a, b model.Incident) bool {
	// FirstSeenAt is local bookkeeping, not a revision ordering signal.
	a.FirstSeenAt, b.FirstSeenAt = time.Time{}, time.Time{}
	x, err := json.Marshal(a)
	if err != nil {
		return false
	}
	y, err := json.Marshal(b)
	return err == nil && string(x) == string(y)
}
