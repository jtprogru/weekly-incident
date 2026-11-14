package canonical

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"slices"
	"sort"
	"time"

	"github.com/jtprogru/weekly-incident/internal/archive"
	"github.com/jtprogru/weekly-incident/internal/config"
	"github.com/jtprogru/weekly-incident/internal/dedup"
	"github.com/jtprogru/weekly-incident/internal/model"
	"github.com/jtprogru/weekly-incident/internal/render"
	"github.com/jtprogru/weekly-incident/internal/score"
	"github.com/jtprogru/weekly-incident/internal/storage"
)

type Inputs struct {
	Archive      string `json:"archive"`
	Weeks        string `json:"weeks"`
	Observations string `json:"observations"`
	Registry     string `json:"registry"`
	Decisions    string `json:"decisions"`
	Config       string `json:"config"`
}

type RowComparison struct {
	DurationMinutes *int        `json:"duration_minutes,omitempty"`
	DurationUnknown bool        `json:"duration_unknown,omitempty"`
	Key             string      `json:"key"`
	Week            string      `json:"week"`
	StartedAt       time.Time   `json:"started_at"`
	ResolvedAt      *time.Time  `json:"resolved_at,omitempty"`
	Components      []string    `json:"components"`
	Score           model.Score `json:"score"`
	AboveThreshold  bool        `json:"above_threshold"`
}

type Change struct {
	CanonicalID string          `json:"canonical_id"`
	Members     []string        `json:"members"`
	Reasons     []string        `json:"reasons"`
	Before      []RowComparison `json:"before"`
	After       RowComparison   `json:"after"`
}

type Plan struct {
	WeekComparisons []WeekComparison      `json:"week_comparisons"`
	SourceHeads     map[string]string     `json:"source_heads"`
	RevisionOptions map[string][]string   `json:"revision_options"`
	SchemaVersion   int                   `json:"schema_version"`
	RulesVersion    string                `json:"rules_version"`
	At              time.Time             `json:"at"`
	Inputs          Inputs                `json:"inputs"`
	RegistryHash    string                `json:"proposed_registry_hash"`
	CanonicalHash   string                `json:"proposed_canonical_hash"`
	Statistics      model.DedupStatistics `json:"statistics"`
	Changes         []Change              `json:"changes"`
	Conflicts       map[string][]string   `json:"conflicts"`
	Weeks           []string              `json:"weeks"`
}

type Manifest struct {
	SchemaVersion int                   `json:"schema_version"`
	RulesVersion  string                `json:"rules_version"`
	GeneratedAt   time.Time             `json:"generated_at"`
	Inputs        Inputs                `json:"inputs"`
	Artifacts     map[string]string     `json:"artifacts"`
	Statistics    model.DedupStatistics `json:"statistics"`
	Pending       []dedup.Candidate     `json:"pending"`
	Weeks         []string              `json:"weeks"`
}

type Prepared struct {
	Plan         Plan
	Inventory    archive.Inventory
	Observations archive.Observations
	Registry     dedup.Registry
	Decisions    dedup.Decisions
	Incidents    []Incident
	Pending      []dedup.Candidate
}

func ReadRegistry(root string) (dedup.Registry, error) {
	r := dedup.EmptyRegistry()
	err := storage.Read(filepath.Join(root, "identity/registry.json"), &r)
	if errors.Is(err, os.ErrNotExist) {
		return r, nil
	}
	return r, err
}

func Prepare(root string, cfg *config.Config, decisionsPath string, now time.Time, reconcile bool, allowPolicyChange bool) (Prepared, error) {
	var p Prepared
	a := archive.New(root)
	inv, err := a.ReadAll()
	if err != nil {
		return p, err
	}
	originalObs, err := a.ReadObservations()
	if err != nil {
		return p, err
	}
	previous, err := ReadRegistry(root)
	if err != nil {
		return p, err
	}
	decisions, err := dedup.LoadDecisions(decisionsPath)
	if err != nil {
		return p, err
	}
	inputs, err := inputHashes(inv, originalObs.Manifest, previous, decisions, cfg)
	if err != nil {
		return p, err
	}
	o, err := a.PrepareObservations(inv, now)
	if err != nil {
		return p, err
	}
	records, err := MatchingRecords(o, cfg.Sources)
	if err != nil {
		return p, err
	}
	r := previous
	if reconcile {
		r, err = dedup.Reconcile(previous, records, cfg.Sources, decisions, cfg.Dedup.AutoAWS, now, allowPolicyChange)
		if err != nil {
			return p, err
		}
	} else {
		dh, err := decisions.Hash()
		if err != nil {
			return p, err
		}
		if len(r.Groups) == 0 || r.DecisionsHash != dh || r.RulesVersion != dedup.IdentityRulesVersion || r.AutoAWS != cfg.Dedup.AutoAWS {
			return p, fmt.Errorf("registry missing or policy changed; run dedup plan/apply")
		}
	}
	incidents, err := Build(o, r, decisions, cfg.Sources)
	if err != nil {
		return p, err
	}
	p = Prepared{Inventory: inv, Observations: o, Registry: r, Decisions: decisions, Incidents: incidents, Pending: []dedup.Candidate{}}
	owner := map[string]string{}
	for _, g := range r.Groups {
		for _, key := range g.Members {
			owner[key] = g.ID
		}
	}
	active, err := decisions.Active(records)
	if err != nil {
		return p, err
	}
	for _, c := range dedup.Audit(inv, cfg.Sources, now).Candidates {
		if owner[c.Left.Key] == owner[c.Right.Key] {
			continue
		}
		dismissed := false
		for _, d := range active {
			if d.Type == "different_incident" && slices.Contains(d.Members, c.Left.Key) && slices.Contains(d.Members, c.Right.Key) {
				dismissed = true
			}
		}
		if !dismissed {
			p.Pending = append(p.Pending, c)
		}
	}
	stats := statistics(incidents, p.Pending, dedup.IdentityRulesVersion)
	conflicts := map[string][]string{}
	for _, c := range incidents {
		if len(c.Conflicts) > 0 {
			conflicts[c.ID] = c.Conflicts
		}
	}
	rh, err := storage.Digest(r)
	if err != nil {
		return p, err
	}
	ch, err := storage.Digest(incidents)
	if err != nil {
		return p, err
	}
	weeks := map[string]bool{}
	for _, w := range inv.Weeks {
		weeks[w.String()] = true
	}
	for _, c := range incidents {
		weeks[archive.WeekOf(c.Fields.StartedAt).String()] = true
	}
	p.Plan = Plan{SchemaVersion: 1, RulesVersion: dedup.IdentityRulesVersion, At: now.UTC(), Inputs: inputs, RegistryHash: rh, CanonicalHash: ch, Statistics: stats, Changes: compareRows(inv, incidents, cfg, now), Conflicts: conflicts, Weeks: sortedKeys(weeks)}
	p.Plan.WeekComparisons = compareWeeks(inv, incidents, cfg, now, p.Plan.Weeks)
	p.Plan.SourceHeads = map[string]string{}
	p.Plan.RevisionOptions = map[string][]string{}
	for key, h := range o.Manifest.Records {
		p.Plan.SourceHeads[key] = h.Revisions[len(h.Revisions)-1].Hash
		if h.UnknownCurrent {
			for _, r := range h.Revisions {
				p.Plan.RevisionOptions[key] = append(p.Plan.RevisionOptions[key], r.Hash)
			}
		}
	}
	return p, nil
}

func inputHashes(inv archive.Inventory, obs archive.ObservationManifest, r dedup.Registry, d dedup.Decisions, cfg *config.Config) (Inputs, error) {
	out := Inputs{Archive: inv.Fingerprint}
	var err error
	out.Weeks, err = storage.Digest(inv.Weeks)
	if err != nil {
		return out, err
	}
	out.Observations, err = storage.Digest(obs)
	if err != nil {
		return out, err
	}
	out.Registry, err = storage.Digest(r)
	if err != nil {
		return out, err
	}
	out.Decisions, err = d.Hash()
	if err != nil {
		return out, err
	}
	out.Config, err = storage.Digest(cfg)
	return out, err
}

func comparison(in model.Incident, cfg *config.Config, now time.Time) RowComparison {
	return RowComparison{DurationMinutes: in.DurationMinutes, DurationUnknown: in.Canonical != nil && in.Canonical.DurationUnknown, Key: in.Key, Week: archive.WeekOf(in.StartedAt).String(), StartedAt: in.StartedAt, ResolvedAt: in.ResolvedAt, Components: sortedStrings(in.Components), Score: score.For(in, cfg.Weights()[in.Vendor], now), AboveThreshold: score.AboveThreshold(in, cfg.Digest, now)}
}

func compareRows(inv archive.Inventory, incidents []Incident, cfg *config.Config, now time.Time) []Change {
	byKey := map[string][]model.Incident{}
	for _, row := range inv.Records {
		byKey[row.Incident.Key] = append(byKey[row.Incident.Key], row.Incident)
	}
	out := []Change{}
	for _, c := range incidents {
		after := comparison(c.Projection(), cfg, now)
		before := []RowComparison{}
		for _, key := range c.Members {
			for _, in := range byKey[key] {
				before = append(before, comparison(in, cfg, now))
			}
		}
		changed := len(before) != 1
		if len(before) == 1 {
			copyAfter := after
			copyAfter.Key = before[0].Key
			changed = !reflect.DeepEqual(before[0], copyAfter)
		}
		if changed {
			reasons := slices.Clone(c.Reasons)
			for _, field := range sortedKeys(c.DecisionFields) {
				reasons = append(reasons, "decision:"+c.DecisionFields[field])
			}
			out = append(out, Change{CanonicalID: c.ID, Members: c.Members, Reasons: reasons, Before: before, After: after})
		}
	}
	return out
}

func statistics(incidents []Incident, pending []dedup.Candidate, version string) model.DedupStatistics {
	s := model.DedupStatistics{RulesVersion: version, Incidents: len(incidents), PendingPairs: len(pending)}
	for _, c := range incidents {
		s.SourceRecords += len(c.Members)
		s.SourceVersions += c.SourceVersions
		if len(c.Conflicts) > 0 {
			s.ConflictedGroups++
		}
	}
	s.MergedRecords = s.SourceRecords - s.Incidents
	return s
}

func (p Prepared) Week(w archive.Week) ([]model.Incident, model.DedupStatistics) {
	cs := []Incident{}
	in := []model.Incident{}
	members := map[string]bool{}
	for _, c := range p.Incidents {
		if w.Contains(c.Fields.StartedAt) {
			cs = append(cs, c)
			in = append(in, c.Projection())
			for _, k := range c.Members {
				members[k] = true
			}
		}
	}
	pending := []dedup.Candidate{}
	for _, c := range p.Pending {
		if members[c.Left.Key] || members[c.Right.Key] {
			pending = append(pending, c)
		}
	}
	return in, statistics(cs, pending, dedup.IdentityRulesVersion)
}

// Commit writes only after every canonical group is valid. The manifest is
// the last file: readers verify every artifact and reject interrupted builds.
func Commit(root string, p Prepared, cfg *config.Config, statuses []model.SourceStatus, touched []archive.Week) error {
	if len(p.Plan.Conflicts) > 0 {
		return fmt.Errorf("%d canonical groups have unresolved conflicts; inspect the plan", len(p.Plan.Conflicts))
	}
	a := archive.New(root)
	artifacts := map[string][]byte{}
	put := func(path string, v any) error {
		b, err := storage.Encode(v)
		if err == nil {
			artifacts[path] = b
		}
		return err
	}
	if err := put("identity/registry.json", p.Registry); err != nil {
		return err
	}
	if err := put("canonical/incidents.json", p.Incidents); err != nil {
		return err
	}
	for _, id := range p.Plan.Weeks {
		w, err := archive.ParseWeek(id)
		if err != nil {
			return err
		}
		in, stats := p.Week(w)
		previous, readErr := a.ReadIndex(w)
		health := previous.Sources
		if len(statuses) > 0 && (readErr != nil || len(health) == 0 || slices.Contains(touched, w)) {
			health = statuses
		}
		idx := score.BuildIndex(w, in, health, cfg, p.Plan.At)
		idx.Dedup = &stats
		// Keep the original timestamp when only the build clock moved.
		if readErr == nil {
			before, after := previous, idx
			before.GeneratedAt = time.Time{}
			after.GeneratedAt = time.Time{}
			if reflect.DeepEqual(before, after) {
				idx.GeneratedAt = previous.GeneratedAt
			}
		}
		if err := put(filepath.ToSlash(filepath.Join(w.Dir(), "index.json")), idx); err != nil {
			return err
		}
	}
	if err := a.WriteObservations(p.Observations); err != nil {
		return err
	}
	// Generating new empty weeks may change inventory directories; hash the
	// semantic week set after writes when publishing the final manifest.
	for _, path := range sortedKeys(artifacts) {
		if err := storage.Write(filepath.Join(root, path), artifacts[path]); err != nil {
			return err
		}
	}
	inv, err := a.ReadAll()
	if err != nil {
		return err
	}
	inputs, err := inputHashes(inv, p.Observations.Manifest, p.Registry, p.Decisions, cfg)
	if err != nil {
		return err
	}
	m := Manifest{SchemaVersion: 1, RulesVersion: dedup.IdentityRulesVersion, GeneratedAt: p.Plan.At, Inputs: inputs, Artifacts: map[string]string{}, Statistics: p.Plan.Statistics, Pending: p.Pending, Weeks: p.Plan.Weeks}
	for path, b := range artifacts {
		m.Artifacts[path] = storage.Hash(b)
	}
	var previous Manifest
	if storage.Read(filepath.Join(root, "canonical/manifest.json"), &previous) == nil {
		before, after := previous, m
		before.GeneratedAt = time.Time{}
		after.GeneratedAt = time.Time{}
		if reflect.DeepEqual(before, after) {
			m.GeneratedAt = previous.GeneratedAt
		}
	}
	return storage.WriteJSON(filepath.Join(root, "canonical/manifest.json"), m)
}

// Apply re-evaluates both inputs and proposed outputs; editing a plan cannot
// smuggle in membership changes and stale plans cannot overwrite a newer run.
func Apply(root string, cfg *config.Config, decisionsPath string, plan Plan) error {
	if plan.SchemaVersion != 1 || plan.RulesVersion != dedup.IdentityRulesVersion {
		return fmt.Errorf("unsupported plan")
	}
	p, err := Prepare(root, cfg, decisionsPath, plan.At, true, true)
	if err != nil {
		return err
	}
	if !reflect.DeepEqual(p.Plan, plan) {
		return fmt.Errorf("stale or edited plan: inputs/proposals changed; run dedup plan again")
	}
	return Commit(root, p, cfg, nil, nil)
}

// Load validates source fingerprints and all derived artifacts before a digest
// is allowed to render. It never makes an identity decision or changes a file.
func Load(root string, cfg *config.Config, decisionsPath string) (Prepared, error) {
	var p Prepared
	var m Manifest
	if err := storage.Read(filepath.Join(root, "canonical/manifest.json"), &m); err != nil {
		return p, fmt.Errorf("canonical view unavailable; run dedup plan/apply or rebuild: %w", err)
	}
	if m.SchemaVersion != 1 || m.RulesVersion != dedup.IdentityRulesVersion {
		return p, fmt.Errorf("unsupported canonical manifest")
	}
	inv, err := archive.New(root).ReadAll()
	if err != nil {
		return p, err
	}
	o, err := archive.New(root).ReadObservations()
	if err != nil {
		return p, err
	}
	r, err := ReadRegistry(root)
	if err != nil {
		return p, err
	}
	d, err := dedup.LoadDecisions(decisionsPath)
	if err != nil {
		return p, err
	}
	inputs, err := inputHashes(inv, o.Manifest, r, d, cfg)
	if err != nil {
		return p, err
	}
	if inputs != m.Inputs {
		return p, fmt.Errorf("canonical inputs changed; run dedup plan/apply or rebuild")
	}
	if m.Artifacts["identity/registry.json"] == "" || m.Artifacts["canonical/incidents.json"] == "" {
		return p, fmt.Errorf("incomplete canonical manifest")
	}
	for path, want := range m.Artifacts {
		if !filepath.IsLocal(path) {
			return p, fmt.Errorf("invalid artifact path %s", path)
		}
		b, err := os.ReadFile(filepath.Join(root, path))
		if err != nil || storage.Hash(b) != want {
			return p, fmt.Errorf("canonical artifact mismatch %s; run dedup rebuild", path)
		}
	}
	if err := storage.Read(filepath.Join(root, "canonical/incidents.json"), &p.Incidents); err != nil {
		return p, err
	}
	records, err := CurrentRecords(o)
	if err != nil {
		return p, err
	}
	if err := r.Validate(records); err != nil {
		return p, err
	}
	if _, err := d.Active(records); err != nil {
		return p, err
	}
	p.Inventory = inv
	p.Observations = o
	p.Registry = r
	p.Decisions = d
	p.Pending = m.Pending
	p.Plan = Plan{SchemaVersion: 1, RulesVersion: dedup.IdentityRulesVersion, At: m.GeneratedAt, Inputs: m.Inputs, Statistics: m.Statistics, Weeks: m.Weeks, Conflicts: map[string][]string{}}
	return p, nil
}

func RenderWeek(root, out string, p Prepared, cfg *config.Config, w archive.Week, now time.Time) error {
	in, stats := p.Week(w)
	a := archive.New(root)
	old, _ := a.ReadIndex(w)
	idx := score.BuildIndex(w, in, old.Sources, cfg, now)
	idx.Dedup = &stats
	idx, err := a.WriteIndex(w, idx)
	if err != nil {
		return err
	}
	// Maintain the integrity manifest when an ongoing score changes with time.
	var manifest Manifest
	if err := storage.Read(filepath.Join(root, "canonical/manifest.json"), &manifest); err != nil {
		return err
	}
	path := filepath.ToSlash(filepath.Join(w.Dir(), "index.json"))
	b, err := os.ReadFile(filepath.Join(root, path))
	if err != nil {
		return err
	}
	manifest.Artifacts[path] = storage.Hash(b)
	if !slices.Contains(manifest.Weeks, w.String()) {
		manifest.Weeks = append(manifest.Weeks, w.String())
		sort.Strings(manifest.Weeks)
	}
	inv, err := a.ReadAll()
	if err != nil {
		return err
	}
	manifest.Inputs.Weeks, err = storage.Digest(inv.Weeks)
	if err != nil {
		return err
	}
	if err := storage.WriteJSON(filepath.Join(root, "canonical/manifest.json"), manifest); err != nil {
		return err
	}
	for name, body := range map[string]string{"digest.md": render.Digest(idx, in), "brief.md": render.Brief(idx, in, cfg.Digest.BriefTopN, cfg.Digest.BriefBodyLimit), "summary.txt": render.Notification(idx, 3)} {
		if err := storage.Write(filepath.Join(out, w.String(), name), []byte(body)); err != nil {
			return err
		}
	}
	return nil
}
