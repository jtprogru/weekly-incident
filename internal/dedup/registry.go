package dedup

import (
	"fmt"
	"slices"
	"sort"
	"strings"
	"time"

	"github.com/jtprogru/weekly-incident/internal/config"
	"github.com/jtprogru/weekly-incident/internal/model"
	"github.com/jtprogru/weekly-incident/internal/storage"
)

const IdentityRulesVersion = "identity-v1"

type Group struct {
	ID      string   `json:"id"`
	Anchor  string   `json:"anchor"`
	Members []string `json:"members"`
	Reasons []string `json:"reasons"`
}

type Alias struct {
	ID      string   `json:"id"`
	Members []string `json:"original_members"`
	Targets []string `json:"targets"`
}

type Transition struct {
	Kind    string   `json:"kind"`
	From    []string `json:"from"`
	To      []string `json:"to"`
	Members []string `json:"members"`
}

type Operation struct {
	ID            string       `json:"id"`
	At            time.Time    `json:"at"`
	RulesVersion  string       `json:"rules_version"`
	DecisionsHash string       `json:"decisions_hash"`
	DecisionIDs   []string     `json:"decision_ids"`
	Transitions   []Transition `json:"transitions"`
}

type Registry struct {
	AutoLinks     [][2]string `json:"auto_links"`
	SchemaVersion int         `json:"schema_version"`
	RulesVersion  string      `json:"rules_version"`
	AutoAWS       bool        `json:"auto_aws"`
	DecisionsHash string      `json:"decisions_hash"`
	Groups        []Group     `json:"groups"`
	Aliases       []Alias     `json:"aliases"`
	Operations    []Operation `json:"operations"`
}

func EmptyRegistry() Registry {
	return Registry{SchemaVersion: 1, RulesVersion: IdentityRulesVersion, Groups: []Group{}, AutoLinks: [][2]string{}, Aliases: []Alias{}, Operations: []Operation{}}
}

func (r Registry) Validate(records map[string]model.Incident) error {
	if r.SchemaVersion != 1 {
		return fmt.Errorf("unsupported registry schema %d", r.SchemaVersion)
	}
	ids, members := map[string]bool{}, map[string]bool{}
	owner := map[string]string{}
	for _, g := range r.Groups {
		if g.ID == "" || ids[g.ID] || len(g.Members) == 0 || !slices.Contains(g.Members, g.Anchor) {
			return fmt.Errorf("invalid registry group %s", g.ID)
		}
		ids[g.ID] = true
		vendor := ""
		for _, key := range g.Members {
			in, ok := records[key]
			if !ok || members[key] {
				return fmt.Errorf("invalid membership for %s", key)
			}
			if vendor != "" && vendor != in.Vendor {
				return fmt.Errorf("cross-vendor group %s", g.ID)
			}
			vendor = in.Vendor
			members[key] = true
			owner[key] = g.ID
		}
	}
	for _, a := range r.Aliases {
		if a.ID == "" || ids[a.ID] || len(a.Members) == 0 || len(a.Targets) == 0 {
			return fmt.Errorf("invalid alias %s", a.ID)
		}
		ids[a.ID] = true
		for _, target := range a.Targets {
			if !slices.ContainsFunc(r.Groups, func(g Group) bool { return g.ID == target }) {
				return fmt.Errorf("alias %s has missing or indirect target %s", a.ID, target)
			}
		}
		targets := []string{}
		for _, key := range a.Members {
			if _, ok := records[key]; !ok {
				return fmt.Errorf("alias %s has unknown member %s", a.ID, key)
			}
			targets = append(targets, owner[key])
		}
		if !slices.Equal(canonicalStrings(targets), a.Targets) {
			return fmt.Errorf("alias %s targets disagree with its original membership", a.ID)
		}
	}
	links := map[[2]string]bool{}
	for _, pair := range r.AutoLinks {
		if pair[0] >= pair[1] || owner[pair[0]] == "" || owner[pair[0]] != owner[pair[1]] || links[pair] {
			return fmt.Errorf("invalid automatic identity link %v", pair)
		}
		links[pair] = true
	}
	operations := map[string]bool{}
	for _, op := range r.Operations {
		if op.ID == "" || operations[op.ID] || op.At.IsZero() || op.RulesVersion == "" || op.DecisionsHash == "" {
			return fmt.Errorf("invalid identity operation %s", op.ID)
		}
		operations[op.ID] = true
		for _, transition := range op.Transitions {
			if !slices.Contains([]string{"create", "merge", "split", "membership"}, transition.Kind) || len(transition.To) == 0 {
				return fmt.Errorf("invalid transition in %s", op.ID)
			}
			for _, id := range append(slices.Clone(transition.From), transition.To...) {
				if !ids[id] {
					return fmt.Errorf("unknown historical identity %s", id)
				}
			}
		}
	}
	return nil
}

func (r Registry) Resolve(id string) (string, error) {
	for _, g := range r.Groups {
		if g.ID == id {
			return id, nil
		}
	}
	for _, a := range r.Aliases {
		if a.ID == id {
			if len(a.Targets) != 1 {
				return "", fmt.Errorf("ambiguous alias %s: %v", id, a.Targets)
			}
			return a.Targets[0], nil
		}
	}
	return "", fmt.Errorf("unknown canonical id %s", id)
}

// Reconcile retains prior automatic identities, rebuilds manual constraints,
// and proposes only complete, unambiguous automatic groups. A manual split
// conservatively breaks the affected old automatic group before reconsidering
// it; it never arbitrarily assigns a bridge record to one side.
func Reconcile(previous Registry, records map[string]model.Incident, sources []config.Source, decisions Decisions, autoAWS bool, now time.Time, allowPolicyChange bool) (Registry, error) {
	if err := previous.Validate(records); err != nil {
		return Registry{}, err
	}
	active, err := decisions.Active(records)
	if err != nil {
		return Registry{}, err
	}
	dh, err := decisions.Hash()
	if err != nil {
		return Registry{}, err
	}
	if !allowPolicyChange && previous.DecisionsHash == "" {
		return Registry{}, fmt.Errorf("initialize identity with dedup plan/apply before enabling collect")
	}
	if len(previous.Groups) > 0 && !allowPolicyChange && (previous.RulesVersion != IdentityRulesVersion || previous.DecisionsHash != dh || previous.AutoAWS != autoAWS) {
		return Registry{}, fmt.Errorf("identity policy changed; run dedup plan/apply")
	}
	keys := []string{}
	for key := range records {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	parent := map[string]string{}
	for _, key := range keys {
		parent[key] = key
	}
	var root func(string) string
	root = func(k string) string {
		if parent[k] != k {
			parent[k] = root(parent[k])
		}
		return parent[k]
	}
	join := func(a, b string) {
		a, b = root(a), root(b)
		if a > b {
			a, b = b, a
		}
		parent[b] = a
	}
	forbidden := map[[2]string]bool{}
	for _, d := range active {
		if d.Type == "different_incident" {
			for i, a := range d.Members {
				for _, b := range d.Members[i+1:] {
					forbidden[pairKey(a, b)] = true
				}
			}
		}
	}
	for _, d := range active {
		if d.Type == "same_incident" {
			for _, key := range d.Members[1:] {
				join(d.Members[0], key)
			}
		}
	}
	for p := range forbidden {
		if root(p[0]) == root(p[1]) {
			return Registry{}, fmt.Errorf("same/different decisions conflict for %s and %s", p[0], p[1])
		}
	}
	// Existing auto groups are explicit evidence, independent of a shrinking feed.
	oldAuto := map[[2]string]bool{}
	for _, g := range previous.Groups {
		if !slices.Contains(g.Reasons, "aws-containment-v1") {
			continue
		}
		blocked := false
		for i, a := range g.Members {
			for _, b := range g.Members[i+1:] {
				if forbidden[pairKey(a, b)] {
					blocked = true
				}
			}
		}
		if !blocked {
			for i, a := range g.Members {
				for _, b := range g.Members[i+1:] {
					if slices.Contains(previous.AutoLinks, pairKey(a, b)) {
						oldAuto[pairKey(a, b)] = true
					}
				}
			}
		}
	}
	kinds := map[string]config.Kind{}
	for _, s := range sources {
		kinds[s.Vendor] = s.Kind
	}
	eligible := map[[2]string]bool{}
	for i, a := range keys {
		for _, b := range keys[i+1:] {
			p := pairKey(a, b)
			if forbidden[p] {
				continue
			}
			if oldAuto[p] || (autoAWS && kinds[records[a].Vendor] == config.KindAWSRSS && records[a].Vendor == records[b].Vendor && AutoAWS(records[a], records[b])) {
				eligible[p] = true
			}
		}
	}
	// Operate on whole manual groups. Automatic edges between two groups require
	// every cross pair to be compatible; connected components must be cliques.
	manual := map[string][]string{}
	for _, key := range keys {
		manual[root(key)] = append(manual[root(key)], key)
	}
	mkeys := []string{}
	for k := range manual {
		mkeys = append(mkeys, k)
	}
	sort.Strings(mkeys)
	adj := map[string][]string{}
	for i, a := range mkeys {
		for _, b := range mkeys[i+1:] {
			ok := true
			for _, x := range manual[a] {
				for _, y := range manual[b] {
					if !eligible[pairKey(x, y)] {
						ok = false
					}
				}
			}
			if ok {
				adj[a] = append(adj[a], b)
				adj[b] = append(adj[b], a)
			}
		}
	}
	visited := map[string]bool{}
	for _, start := range mkeys {
		if visited[start] {
			continue
		}
		component := []string{}
		queue := []string{start}
		visited[start] = true
		for len(queue) > 0 {
			x := queue[0]
			queue = queue[1:]
			component = append(component, x)
			for _, y := range adj[x] {
				if !visited[y] {
					visited[y] = true
					queue = append(queue, y)
				}
			}
		}
		clique := true
		for _, x := range component {
			if len(adj[x]) != len(component)-1 {
				clique = false
			}
		}
		if clique {
			for _, x := range component[1:] {
				join(component[0], x)
			}
		} else {
			// Keep already confirmed automatic subgroups, leaving new ambiguous
			// vertices separate. The negative decisions above take precedence.
			for _, g := range previous.Groups {
				if !slices.Contains(g.Reasons, "aws-containment-v1") {
					continue
				}
				ok := true
				for i, a := range g.Members {
					for _, b := range g.Members[i+1:] {
						if !oldAuto[pairKey(a, b)] {
							ok = false
						}
					}
				}
				if ok {
					for _, x := range g.Members[1:] {
						join(g.Members[0], x)
					}
				}
			}
		}
	}
	for p := range forbidden {
		if root(p[0]) == root(p[1]) {
			return Registry{}, fmt.Errorf("group constraint violated: %v", p)
		}
	}
	parts := map[string][]string{}
	for _, key := range keys {
		parts[root(key)] = append(parts[root(key)], key)
	}
	groups := []Group{}
	issued := map[string]bool{}
	for _, g := range previous.Groups {
		issued[g.ID] = true
	}
	for _, a := range previous.Aliases {
		issued[a.ID] = true
	}
	partKeys := []string{}
	for key := range parts {
		partKeys = append(partKeys, key)
	}
	sort.Strings(partKeys)
	for _, key := range partKeys {
		members := parts[key]
		g := Group{Anchor: members[0], Members: members, Reasons: []string{}}
		for _, old := range previous.Groups {
			if slices.Contains(members, old.Anchor) && (g.ID == "" || old.ID < g.ID) {
				g.ID = old.ID
				g.Anchor = old.Anchor
			}
		}
		if g.ID == "" {
			for n := 0; ; n++ {
				candidate := "ci-" + fingerprint("canonical-v1", g.Anchor, fmt.Sprint(n))[:24]
				if !issued[candidate] {
					g.ID = candidate
					issued[candidate] = true
					break
				}
			}
		}
		for _, d := range active {
			if d.Type == "same_incident" && slices.Contains(members, d.Members[0]) {
				g.Reasons = append(g.Reasons, "decision:"+d.ID)
			}
		}
		for i, a := range members {
			for _, b := range members[i+1:] {
				if eligible[pairKey(a, b)] {
					g.Reasons = append(g.Reasons, "aws-containment-v1")
				}
			}
		}
		g.Reasons = canonicalStrings(g.Reasons)
		groups = append(groups, g)
	}
	sort.Slice(groups, func(i, j int) bool { return groups[i].ID < groups[j].ID })
	owner := map[string]string{}
	for _, g := range groups {
		for _, key := range g.Members {
			owner[key] = g.ID
		}
	}
	aliases := slices.Clone(previous.Aliases)
	for _, old := range previous.Groups {
		if !slices.ContainsFunc(groups, func(g Group) bool { return g.ID == old.ID }) {
			aliases = append(aliases, Alias{ID: old.ID, Members: slices.Clone(old.Members)})
		}
	}
	for i := range aliases {
		targets := []string{}
		for _, key := range aliases[i].Members {
			targets = append(targets, owner[key])
		}
		aliases[i].Targets = canonicalStrings(targets)
	}
	sort.Slice(aliases, func(i, j int) bool { return aliases[i].ID < aliases[j].ID })
	autoLinks := [][2]string{}
	for pair := range eligible {
		if owner[pair[0]] == owner[pair[1]] {
			autoLinks = append(autoLinks, pair)
		}
	}
	sort.Slice(autoLinks, func(i, j int) bool {
		if autoLinks[i][0] != autoLinks[j][0] {
			return autoLinks[i][0] < autoLinks[j][0]
		}
		return autoLinks[i][1] < autoLinks[j][1]
	})
	next := Registry{AutoLinks: autoLinks, SchemaVersion: 1, RulesVersion: IdentityRulesVersion, AutoAWS: autoAWS, DecisionsHash: dh, Groups: groups, Aliases: aliases, Operations: slices.Clone(previous.Operations)}
	before, _ := storage.Digest(struct {
		Groups    []Group
		Aliases   []Alias
		Decisions string
		Rules     string
		Auto      bool
		Links     [][2]string
	}{previous.Groups, previous.Aliases, previous.DecisionsHash, previous.RulesVersion, previous.AutoAWS, previous.AutoLinks})
	after, _ := storage.Digest(struct {
		Groups    []Group
		Aliases   []Alias
		Decisions string
		Rules     string
		Auto      bool
		Links     [][2]string
	}{next.Groups, next.Aliases, next.DecisionsHash, next.RulesVersion, next.AutoAWS, next.AutoLinks})
	if before != after {
		changes := []Transition{}
		oldOwner := map[string]string{}
		for _, g := range previous.Groups {
			for _, key := range g.Members {
				oldOwner[key] = g.ID
			}
		}
		for _, old := range previous.Groups {
			targets := []string{}
			for _, key := range old.Members {
				targets = append(targets, owner[key])
			}
			targets = canonicalStrings(targets)
			if len(targets) > 1 {
				changes = append(changes, Transition{Kind: "split", From: []string{old.ID}, To: targets, Members: old.Members})
			}
		}
		for _, g := range groups {
			from := []string{}
			for _, key := range g.Members {
				from = append(from, oldOwner[key])
			}
			from = canonicalStrings(from)
			kind := "membership"
			if len(from) == 0 {
				kind = "create"
			} else if len(from) > 1 {
				kind = "merge"
			}
			if len(from) != 1 || from[0] != g.ID || !slices.ContainsFunc(previous.Groups, func(old Group) bool { return old.ID == g.ID && slices.Equal(old.Members, g.Members) }) {
				changes = append(changes, Transition{Kind: kind, From: from, To: []string{g.ID}, Members: g.Members})
			}
		}
		decisionIDs := []string{}
		for _, d := range active {
			decisionIDs = append(decisionIDs, d.ID)
		}
		next.Operations = append(next.Operations, Operation{ID: fingerprint(before, after, fmt.Sprint(len(previous.Operations))), At: now.UTC(), RulesVersion: IdentityRulesVersion, DecisionsHash: dh, DecisionIDs: decisionIDs, Transitions: changes})
	}
	return next, next.Validate(records)
}

func pairKey(a, b string) [2]string {
	if a > b {
		a, b = b, a
	}
	return [2]string{a, b}
}

// AutoAWS is deliberately narrower than candidate generation: exact service,
// containment, at least two distinct non-template messages, and no closure
// followed by another active chain. No fuzzy text or inferred start is used.
func AutoAWS(a, b model.Incident) bool {
	if a.Vendor != b.Vendor || len(a.Components) != 1 || a.Components[0] == "" || !slices.Equal(a.Components, b.Components) {
		return false
	}
	if (a.ResolvedAt != nil && a.ResolvedAt.Before(b.StartedAt)) || (b.ResolvedAt != nil && b.ResolvedAt.Before(a.StartedAt)) {
		return false
	}
	sets := func(in model.Incident) map[string]string {
		out := map[string]string{}
		for _, u := range in.Updates {
			body := normalizeBody(u.Body)
			low := strings.ToLower(strings.Trim(body, " .!"))
			if body == "" || u.CreatedAt.IsZero() || slices.Contains([]string{"investigating", "we are investigating", "service is operating normally", "resolved", "monitoring"}, low) {
				continue
			}
			out[fingerprint(u.CreatedAt.UTC().Format(time.RFC3339Nano), body)] = body
		}
		return out
	}
	x, y := sets(a), sets(b)
	shared := map[string]bool{}
	count := 0
	for k, body := range x {
		if _, ok := y[k]; ok {
			count++
			shared[body] = true
		}
	}
	if len(shared) < 2 || (count != len(x) && count != len(y)) {
		return false
	}
	for _, in := range []model.Incident{a, b} {
		updates := slices.Clone(in.Updates)
		sort.Slice(updates, func(i, j int) bool { return updates[i].CreatedAt.Before(updates[j].CreatedAt) })
		closed := false
		for _, u := range updates {
			if closed && (u.Status == model.StatusInvestigating || u.Status == model.StatusIdentified) {
				return false
			}
			if u.Status == model.StatusResolved {
				closed = true
			}
		}
	}
	return true
}
