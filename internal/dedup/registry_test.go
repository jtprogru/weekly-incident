package dedup

import (
	"reflect"
	"slices"
	"testing"
	"time"

	"github.com/jtprogru/weekly-incident/internal/config"
	"github.com/jtprogru/weekly-incident/internal/model"
)

func recordsFor(ids ...string) map[string]model.Incident {
	out := map[string]model.Incident{}
	for _, id := range ids {
		in := synthetic(id)
		out[in.Key] = in
	}
	return out
}
func decision(id, kind string, members ...string) Decision {
	return Decision{ID: id, Type: kind, Members: members, Author: "test fixture", Reason: "synthetic regression", Evidence: []string{"test fixture"}, At: fixedNow}
}
func decisions(ds ...Decision) Decisions { return Decisions{SchemaVersion: 1, Decisions: ds} }

func TestRegistryMergeSplitAndStableIdentity(t *testing.T) {
	records := recordsFor("a", "b", "c")
	r, err := Reconcile(EmptyRegistry(), records, nil, decisions(), false, fixedNow, true)
	if err != nil {
		t.Fatal(err)
	}
	oldIDs := map[string]string{}
	for _, g := range r.Groups {
		oldIDs[g.Anchor] = g.ID
	}
	merge := decision("merge", "same_incident", "test/a", "test/b")
	merged, err := Reconcile(r, records, nil, decisions(merge), false, fixedNow, true)
	if err != nil {
		t.Fatal(err)
	}
	if len(merged.Groups) != 2 || len(merged.Aliases) != 1 {
		t.Fatalf("merge %+v", merged)
	}
	for _, id := range []string{oldIDs["test/a"], oldIDs["test/b"]} {
		if _, err := merged.Resolve(id); err != nil {
			t.Fatal(err)
		}
	}
	split := decision("split", "different_incident", "test/a", "test/b")
	split.Supersedes = "merge"
	separated, err := Reconcile(merged, records, nil, decisions(merge, split), false, fixedNow, true)
	if err != nil {
		t.Fatal(err)
	}
	if len(separated.Groups) != 3 {
		t.Fatalf("split groups %d", len(separated.Groups))
	}
	var retained Group
	for _, g := range merged.Groups {
		if len(g.Members) == 2 {
			retained = g
		}
	}
	if !slices.ContainsFunc(separated.Groups, func(g Group) bool { return g.ID == retained.ID && g.Anchor == retained.Anchor }) {
		t.Fatal("split changed anchor identity")
	}
	repeat, err := Reconcile(separated, records, nil, decisions(merge, split), false, fixedNow.Add(time.Hour), false)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(separated, repeat) {
		t.Fatal("reconcile is not idempotent")
	}
	newRecord := synthetic("earlier")
	newRecord.StartedAt = fixedNow.Add(-365 * 24 * time.Hour)
	records[newRecord.Key] = newRecord
	more, err := Reconcile(repeat, records, nil, decisions(merge, split), false, fixedNow, false)
	if err != nil {
		t.Fatal(err)
	}
	for _, old := range separated.Groups {
		if !slices.ContainsFunc(more.Groups, func(g Group) bool { return g.ID == old.ID }) {
			t.Fatal("earlier incident changed assigned ID")
		}
	}
}

func TestGroupWideConstraintsAndDecisionValidation(t *testing.T) {
	records := recordsFor("a", "b", "c")
	for _, ds := range []Decisions{
		decisions(decision("ab", "same_incident", "test/a", "test/b"), decision("bc", "same_incident", "test/b", "test/c"), decision("ac", "different_incident", "test/a", "test/c")),
		decisions(decision("unknown", "same_incident", "test/a", "test/absent")),
		decisions(decision("id", "same_incident", "test/a", "test/b"), decision("id", "different_incident", "test/a", "test/b")),
	} {
		if _, err := Reconcile(EmptyRegistry(), records, nil, ds, false, fixedNow, true); err == nil {
			t.Fatal("invalid decisions accepted")
		}
	}
	a, b := decision("a", "same_incident", "test/a", "test/b"), decision("b", "different_incident", "test/a", "test/b")
	a.Supersedes = "b"
	b.Supersedes = "a"
	if _, err := decisions(a, b).Active(records); err == nil {
		t.Fatal("supersession cycle accepted")
	}
}

func TestAutomaticBridgeIsNotTransitive(t *testing.T) {
	records := recordsFor("a", "b", "c")
	// A is a subset of B, C is a disjoint subset of B. Pairwise connectedness
	// alone would merge two independent chains through B.
	u := func(n int) model.Update {
		return model.Update{CreatedAt: fixedNow.Add(time.Duration(n) * time.Minute), Body: []string{"first distinct report", "second distinct report", "third distinct report", "fourth distinct report"}[n], Status: model.StatusInvestigating}
	}
	a, b, c := records["test/a"], records["test/b"], records["test/c"]
	a.Updates = []model.Update{u(0), u(1)}
	b.Updates = []model.Update{u(0), u(1), u(2), u(3)}
	c.Updates = []model.Update{u(2), u(3)}
	records[a.Key], records[b.Key], records[c.Key] = a, b, c
	sources := []config.Source{{Vendor: "test", Kind: config.KindAWSRSS}}
	r, err := Reconcile(EmptyRegistry(), records, sources, decisions(), true, fixedNow, true)
	if err != nil {
		t.Fatal(err)
	}
	if len(r.Groups) != 3 {
		t.Fatal("ambiguous bridge auto-merged")
	}
	if !AutoAWS(a, b) || !AutoAWS(b, c) || AutoAWS(a, c) {
		t.Fatal("fixture does not exercise the bridge")
	}
	// Explicitly choosing A/B does not sweep C into that group.
	r, err = Reconcile(r, records, sources, decisions(decision("ab", "same_incident", a.Key, b.Key)), true, fixedNow, true)
	if err != nil {
		t.Fatal(err)
	}
	if len(r.Groups) != 2 {
		t.Fatal("group check failed")
	}
}

func TestRulesChangeNeedsPlan(t *testing.T) {
	records := recordsFor("a")
	r, err := Reconcile(EmptyRegistry(), records, nil, decisions(), false, fixedNow, true)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := Reconcile(r, records, nil, decisions(), true, fixedNow, false); err == nil {
		t.Fatal("silent rule change")
	}
	r.RulesVersion = "older-version"
	if _, err := Reconcile(r, records, nil, decisions(), false, fixedNow, false); err == nil {
		t.Fatal("silent version change")
	}
}

func TestAmbiguousAliasAfterSplit(t *testing.T) {
	records := recordsFor("a", "b", "c", "d")
	ab := decision("ab", "same_incident", "test/a", "test/b")
	cd := decision("cd", "same_incident", "test/c", "test/d")
	r, err := Reconcile(EmptyRegistry(), records, nil, decisions(ab, cd), false, fixedNow, true)
	if err != nil {
		t.Fatal(err)
	}
	all := decision("all", "same_incident", "test/a", "test/b", "test/c", "test/d")
	merged, err := Reconcile(r, records, nil, decisions(ab, cd, all), false, fixedNow, true)
	if err != nil {
		t.Fatal(err)
	}
	alias := merged.Aliases[0]
	notAll := decision("not-all", "different_incident", "test/a", "test/b", "test/c", "test/d")
	notAll.Supersedes = "all"
	notAB := decision("not-ab", "different_incident", "test/a", "test/b")
	notAB.Supersedes = "ab"
	notCD := decision("not-cd", "different_incident", "test/c", "test/d")
	notCD.Supersedes = "cd"
	split, err := Reconcile(merged, records, nil, decisions(ab, cd, all, notAll, notAB, notCD), false, fixedNow, true)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := split.Resolve(alias.ID); err == nil {
		t.Fatal("split alias silently chose a side")
	}
}
