// Package canonical builds reversible report views from source observations
// and a durable identity registry. It performs no network requests.
package canonical

import (
	"encoding/json"
	"fmt"
	"reflect"
	"slices"
	"sort"
	"strings"
	"time"

	"github.com/jtprogru/weekly-incident/internal/archive"
	"github.com/jtprogru/weekly-incident/internal/config"
	"github.com/jtprogru/weekly-incident/internal/dedup"
	"github.com/jtprogru/weekly-incident/internal/model"
	"github.com/jtprogru/weekly-incident/internal/source"
	"github.com/jtprogru/weekly-incident/internal/storage"
)

type Evidence struct {
	SourceKey string `json:"source_key"`
	Snapshot  string `json:"snapshot"`
	Revision  int    `json:"revision"`
}

type UpdateVersion struct {
	Evidence Evidence     `json:"evidence"`
	SeenAt   time.Time    `json:"seen_at"`
	Imported bool         `json:"imported"`
	Update   model.Update `json:"update"`
}

type UpdateHistory struct {
	ID       string          `json:"id"`
	Active   bool            `json:"active"`
	Selected *UpdateVersion  `json:"selected,omitempty"`
	Versions []UpdateVersion `json:"versions"`
}

type Incident struct {
	DurationUnknown bool                  `json:"duration_unknown,omitempty"`
	ID              string                `json:"canonical_id"`
	Members         []string              `json:"members"`
	Fields          model.Incident        `json:"fields"`
	Sources         []model.SourceLink    `json:"sources"`
	Reasons         []string              `json:"reasons"`
	StartQuality    string                `json:"start_quality"`
	FieldProvenance map[string][]Evidence `json:"field_provenance"`
	DecisionFields  map[string]string     `json:"field_decisions,omitempty"`
	UpdateHistory   []UpdateHistory       `json:"update_history"`
	Conflicts       []string              `json:"conflicts"`
	Warnings        []string              `json:"warnings"`
	SourceVersions  int                   `json:"source_versions"`
}

// Projection deliberately does not call Incident.Finalize: that would replace
// the canonical ID with a vendor-native key.
func (c Incident) Projection() model.Incident {
	in := c.Fields
	in.Key = c.ID
	in.Canonical = &model.CanonicalInfo{DurationUnknown: c.DurationUnknown, ID: c.ID, Sources: c.Sources, Reasons: c.Reasons, StartQuality: c.StartQuality, Warnings: c.Warnings}
	return in
}

func CurrentRecords(o archive.Observations) (map[string]model.Incident, error) {
	rows, err := o.Current()
	if err != nil {
		return nil, err
	}
	out := map[string]model.Incident{}
	for _, in := range rows {
		out[in.Key] = in
	}
	return out, nil
}

// MatchingRecords preserves historic AWS messages when the current feed has
// lost them. Fields remain current; this union is only input to identity rules.
func MatchingRecords(o archive.Observations, sources []config.Source) (map[string]model.Incident, error) {
	out, err := CurrentRecords(o)
	if err != nil {
		return nil, err
	}
	kinds := sourceKinds(sources)
	for key, in := range out {
		if kinds[in.Vendor] != config.KindAWSRSS {
			continue
		}
		updates := map[string]model.Update{}
		for _, r := range o.Manifest.Records[key].Revisions {
			for _, u := range source.IdentityUpdates(o.Snapshots[r.Hash].Record, config.KindAWSRSS) {
				updates[updateFingerprint(u)] = u
			}
		}
		in.Updates = []model.Update{}
		for _, id := range sortedKeys(updates) {
			in.Updates = append(in.Updates, updates[id])
		}
		out[key] = in
	}
	return out, nil
}

func Build(o archive.Observations, registry dedup.Registry, decisions dedup.Decisions, sources []config.Source) ([]Incident, error) {
	records, err := CurrentRecords(o)
	if err != nil {
		return nil, err
	}
	if err := registry.Validate(records); err != nil {
		return nil, err
	}
	active, err := decisions.Active(records)
	if err != nil {
		return nil, err
	}
	kinds := sourceKinds(sources)
	out := []Incident{}
	members := 0
	for _, g := range registry.Groups {
		c, err := buildGroup(o, g, active, kinds)
		if err != nil {
			return nil, err
		}
		out = append(out, c)
		members += len(g.Members)
	}
	if members != len(records) {
		return nil, fmt.Errorf("registry has %d memberships for %d sources; reconcile first", members, len(records))
	}
	return out, nil
}

func buildGroup(o archive.Observations, g dedup.Group, active []dedup.Decision, kinds map[string]config.Kind) (Incident, error) {
	c := Incident{ID: g.ID, Members: slices.Clone(g.Members), Sources: []model.SourceLink{}, Reasons: slices.Clone(g.Reasons), StartQuality: "vendor_reported", FieldProvenance: map[string][]Evidence{}, DecisionFields: map[string]string{}, UpdateHistory: []UpdateHistory{}, Conflicts: []string{}, Warnings: []string{}}
	selected := map[string]archive.Revision{}
	chosen := map[string]model.Incident{}
	overrides := map[string]dedup.Decision{}
	for _, d := range active {
		if d.Type != "field_override" || !slices.Contains(g.Members, d.Members[0]) {
			continue
		}
		for _, key := range d.Members {
			if !slices.Contains(g.Members, key) {
				return c, fmt.Errorf("override %s spans separate canonical groups", d.ID)
			}
		}
		for _, key := range g.Members {
			h := o.Manifest.Records[key]
			hash := h.Revisions[len(h.Revisions)-1].Hash
			if d.ExpectedHashes[key] != hash {
				return c, fmt.Errorf("stale override %s: evidence changed for %s", d.ID, key)
			}
		}
		name := d.Field
		if name == "source_revision" {
			name += "/" + d.Members[0]
		}
		if _, ok := overrides[name]; ok {
			return c, fmt.Errorf("conflicting field overrides for %s/%s", g.ID, name)
		}
		overrides[name] = d
	}
	for _, key := range g.Members {
		h := o.Manifest.Records[key]
		r := h.Revisions[len(h.Revisions)-1]
		if d, ok := overrides["source_revision/"+key]; ok {
			hash, ok := d.Value.(string)
			if !ok {
				return c, fmt.Errorf("source_revision requires a snapshot hash")
			}
			found := false
			for _, v := range h.Revisions {
				if v.Hash == hash {
					r = v
					found = true
				}
			}
			if !found {
				return c, fmt.Errorf("snapshot %s does not belong to %s", hash, key)
			}
			c.DecisionFields["source_revision/"+key] = d.ID
		} else if h.UnknownCurrent {
			c.Conflicts = append(c.Conflicts, "unknown_revision_order:"+key)
		}
		in := o.Snapshots[r.Hash].Record
		in.FirstSeenAt = h.FirstSeenAt
		selected[key] = r
		chosen[key] = in
		c.SourceVersions += len(h.Revisions)
		c.Sources = append(c.Sources, model.SourceLink{Key: key, Path: r.Path, URL: in.URL})
	}
	base := chosen[g.Members[0]]
	c.Fields = base
	c.Fields.Key = ""
	c.Fields.NativeID = ""
	c.Fields.Raw = nil
	c.Fields.Canonical = nil
	aws := kinds[base.Vendor] == config.KindAWSRSS
	if aws {
		c.StartQuality = "first_observed_update"
		c.Warnings = append(c.Warnings, "AWS onset is the first observed message, not a verified impact start")
	}
	fields := []string{"started_at", "resolved_at", "status", "impact", "components", "title"}
	for _, field := range fields {
		for _, key := range g.Members {
			if field == "title" && key != g.Members[0] {
				continue
			}
			r := selected[key]
			c.FieldProvenance[field] = append(c.FieldProvenance[field], Evidence{SourceKey: key, Snapshot: r.Hash, Revision: r.Sequence})
			in := chosen[key]
			if field == "title" {
				continue
			}
			if aws && (field == "started_at" || field == "resolved_at" || field == "status") {
				continue
			}
			if !reflect.DeepEqual(fieldValue(base, field), fieldValue(in, field)) {
				c.Conflicts = append(c.Conflicts, "field:"+field)
			}
		}
	}
	c.Fields.Updates, c.UpdateHistory, c.Conflicts = buildUpdates(o, g, selected, kinds, c.Conflicts)
	if aws && len(c.Fields.Updates) > 0 {
		c.Fields.StartedAt = c.Fields.Updates[0].CreatedAt
		last := c.Fields.Updates[len(c.Fields.Updates)-1]
		c.Fields.Status = last.Status
		c.Fields.ResolvedAt = nil
		if last.Status == model.StatusResolved {
			at := last.CreatedAt
			c.Fields.ResolvedAt = &at
		}
		for _, field := range []string{"started_at", "resolved_at", "status"} {
			c.FieldProvenance[field] = nil
			for _, h := range c.UpdateHistory {
				if h.Active && h.Selected != nil {
					c.FieldProvenance[field] = append(c.FieldProvenance[field], h.Selected.Evidence)
				}
			}
		}
	}
	for _, key := range g.Members {
		in := chosen[key]
		if c.Fields.FirstSeenAt.IsZero() || (!in.FirstSeenAt.IsZero() && in.FirstSeenAt.Before(c.Fields.FirstSeenAt)) {
			c.Fields.FirstSeenAt = in.FirstSeenAt
		}
	}
	for _, field := range sortedKeys(overrides) {
		d := overrides[field]
		if strings.HasPrefix(field, "source_revision/") {
			continue
		}
		if field == "duration_unknown" {
			continue
		}
		if err := applyField(&c.Fields, field, d.Value); err != nil {
			return c, fmt.Errorf("override %s: %w", d.ID, err)
		}
		c.DecisionFields[field] = d.ID
		if field == "started_at" {
			c.StartQuality = "editorial_override"
		}
		c.Conflicts = slices.DeleteFunc(c.Conflicts, func(s string) bool { return s == "field:"+field })
	}
	if d, ok := overrides["duration_unknown"]; ok {
		if d.Value != true || c.Fields.ResolvedAt == nil || !c.Fields.ResolvedAt.Before(c.Fields.StartedAt) {
			return c, fmt.Errorf("override %s requires true and contradictory resolved/start timestamps", d.ID)
		}
		c.DurationUnknown = true
		c.DecisionFields["duration_unknown"] = d.ID
		c.Warnings = append(c.Warnings, "Vendor resolved_at precedes started_at; duration is unknown, score uses zero duration; decision:"+d.ID)
	}
	// A display name is not a component ID: Statuspage can list multiple
	// distinct regional/product components named "Global".
	c.Fields.Components = sortedStrings(c.Fields.Components)
	c.Fields.DurationMinutes = nil
	if c.Fields.ResolvedAt != nil {
		duration := int(c.Fields.ResolvedAt.Sub(c.Fields.StartedAt).Minutes())
		if c.Fields.ResolvedAt.Before(c.Fields.StartedAt) {
			if !c.DurationUnknown {
				c.Conflicts = append(c.Conflicts, "resolved_before_started")
			}
		} else {
			c.Fields.DurationMinutes = &duration
		}
	}
	if c.Fields.StartedAt.IsZero() {
		c.Conflicts = append(c.Conflicts, "unknown_started_at")
	}
	c.Conflicts = uniqueStrings(c.Conflicts)
	return c, nil
}

func fieldValue(in model.Incident, field string) any {
	switch field {
	case "started_at":
		return in.StartedAt.UTC()
	case "resolved_at":
		if in.ResolvedAt == nil {
			return nil
		}
		return in.ResolvedAt.UTC()
	case "status":
		return in.Status
	case "impact":
		return in.Impact
	case "components":
		return sortedStrings(in.Components)
	case "title":
		return in.Title
	}
	return nil
}

func applyField(in *model.Incident, field string, v any) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	switch field {
	case "started_at":
		err = json.Unmarshal(b, &in.StartedAt)
	case "resolved_at":
		err = json.Unmarshal(b, &in.ResolvedAt)
	case "status":
		err = json.Unmarshal(b, &in.Status)
		if err == nil && !slices.Contains([]model.Status{model.StatusInvestigating, model.StatusIdentified, model.StatusMonitoring, model.StatusResolved, model.StatusPostmortem, model.StatusUnknown}, in.Status) {
			err = fmt.Errorf("invalid status")
		}
	case "impact":
		err = json.Unmarshal(b, &in.Impact)
		if err == nil && !slices.Contains([]model.Impact{model.ImpactUnknown, model.ImpactNone, model.ImpactMinor, model.ImpactMajor, model.ImpactCritical}, in.Impact) {
			err = fmt.Errorf("invalid impact")
		}
	case "components":
		err = json.Unmarshal(b, &in.Components)
	case "title":
		err = json.Unmarshal(b, &in.Title)
	default:
		err = fmt.Errorf("unknown field %s", field)
	}
	return err
}

func buildUpdates(o archive.Observations, g dedup.Group, selected map[string]archive.Revision, kinds map[string]config.Kind, conflicts []string) ([]model.Update, []UpdateHistory, []string) {
	versions := map[string][]UpdateVersion{}
	current := map[string]map[string]UpdateVersion{}
	for _, key := range g.Members {
		h := o.Manifest.Records[key]
		kind := kinds[h.Vendor]
		bySource := map[string]UpdateVersion{}
		for _, r := range h.Revisions {
			in := o.Snapshots[r.Hash].Record
			for _, u := range source.IdentityUpdates(in, kind) {
				id := updateFingerprint(u)
				if u.NativeID != "" {
					id = key + "/" + u.NativeID
					if kind == config.KindAWSRSS {
						id = h.Vendor + "/" + u.NativeID
					}
				}
				v := UpdateVersion{Evidence: Evidence{SourceKey: key, Snapshot: r.Hash, Revision: r.Sequence}, SeenAt: r.SeenAt, Imported: r.Imported, Update: u}
				versions[id] = append(versions[id], v)
				if r.Sequence <= selected[key].Sequence && (kind == config.KindAWSRSS || r.Sequence == selected[key].Sequence) {
					bySource[id] = v
				}
			}
		}
		for id, v := range bySource {
			if current[id] == nil {
				current[id] = map[string]UpdateVersion{}
			}
			current[id][key] = v
		}
	}
	out := []model.Update{}
	histories := []UpdateHistory{}
	visible := map[string]bool{}
	for _, id := range sortedKeys(versions) {
		h := UpdateHistory{ID: id, Versions: versions[id]}
		candidates := current[id]
		keys := sortedKeys(candidates)
		if len(keys) > 0 {
			best := candidates[keys[0]]
			for _, key := range keys[1:] {
				v := candidates[key]
				if updateFingerprint(v.Update) == updateFingerprint(best.Update) && v.Update.Status == best.Update.Status {
					continue
				}
				if best.Imported && !v.Imported {
					best = v
					continue
				}
				if !best.Imported && v.Imported {
					continue
				}
				if !best.Imported && !v.Imported && !best.SeenAt.Equal(v.SeenAt) {
					if v.SeenAt.After(best.SeenAt) {
						best = v
					}
					continue
				}
				conflicts = append(conflicts, "update_revision_order:"+id)
			}
			h.Active = true
			h.Selected = &best
			fp := updateFingerprint(best.Update) + "/" + string(best.Update.Status)
			if !visible[fp] {
				visible[fp] = true
				out = append(out, best.Update)
			}
		}
		histories = append(histories, h)
	}
	sort.Slice(out, func(i, j int) bool {
		if !out[i].CreatedAt.Equal(out[j].CreatedAt) {
			return out[i].CreatedAt.Before(out[j].CreatedAt)
		}
		return updateFingerprint(out[i]) < updateFingerprint(out[j])
	})
	return out, histories, conflicts
}

func updateFingerprint(u model.Update) string {
	b, _ := json.Marshal([]string{u.CreatedAt.UTC().Format(time.RFC3339Nano), strings.TrimSpace(strings.ReplaceAll(u.Body, "\r\n", "\n"))})
	return storage.Hash(b)
}
func uniqueStrings(in []string) []string {
	return slices.Compact(sortedStrings(in))
}
func sortedStrings(in []string) []string {
	out := slices.Clone(in)
	sort.Strings(out)
	return out
}
func sortedKeys[V any](m map[string]V) []string {
	out := []string{}
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}
func sourceKinds(sources []config.Source) map[string]config.Kind {
	out := map[string]config.Kind{}
	for _, s := range sources {
		out[s.Vendor] = s.Kind
	}
	return out
}
