package source

import (
	"slices"
	"testing"

	"github.com/jtprogru/weekly-incident/internal/config"
	"github.com/jtprogru/weekly-incident/internal/dedup"
)

func TestAWSIdentityGUIDBackfill(t *testing.T) {
	a := &awsRSS{base: base{vendor: "aws"}}
	result, err := a.parse(golden(t, "aws_all.rss"))
	if err != nil {
		t.Fatal(err)
	}
	for _, in := range result.Incidents {
		original := slices.Clone(in.Updates)
		in.Updates = slices.Clone(in.Updates)
		for i := range in.Updates {
			if in.Updates[i].NativeID == "" {
				t.Fatal("parser lost GUID")
			}
			in.Updates[i].NativeID = ""
		}
		restored := IdentityUpdates(in, config.KindAWSRSS)
		for i, u := range restored {
			if u.NativeID != original[i].NativeID {
				t.Fatal("backfill differs from parser")
			}
		}
		if in.Updates[0].NativeID != "" {
			t.Fatal("backfill mutated source")
		}
	}
}

// Held-out transformations of a separately captured feed. These are explicit
// synthetic window truncations, not independently adjudicated real outages.
func TestAWSHeldOutWindowRules(t *testing.T) {
	a := &awsRSS{base: base{vendor: "aws"}}
	result, err := a.parse(golden(t, "aws_all.rss"))
	if err != nil {
		t.Fatal(err)
	}
	positives := 0
	negatives := 0
	for i, original := range result.Incidents {
		if len(original.Updates) >= 4 {
			short := original
			short.Updates = slices.Clone(original.Updates[len(original.Updates)/2:])
			short.StartedAt = short.Updates[0].CreatedAt
			if !dedup.AutoAWS(original, short) {
				t.Fatalf("missed held-out truncation %s", original.Key)
			}
			positives++
			otherRegion := short
			otherRegion.Components = []string{"another-region"}
			if dedup.AutoAWS(original, otherRegion) {
				t.Fatal("merged regions")
			}
			negatives++
		}
		for _, other := range result.Incidents[i+1:] {
			negatives++
			if dedup.AutoAWS(original, other) {
				t.Fatalf("merged separate captured chains %s %s", original.Key, other.Key)
			}
		}
	}
	if positives == 0 {
		t.Fatal("no held-out window examples")
	}
	t.Logf("held-out synthetic windows: %d positives, %d negatives; captured chains: %d", positives, negatives, len(result.Incidents))
}
