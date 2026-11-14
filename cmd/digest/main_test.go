package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/jtprogru/weekly-incident/internal/archive"
	"github.com/jtprogru/weekly-incident/internal/canonical"
	"github.com/jtprogru/weekly-incident/internal/config"
	"github.com/jtprogru/weekly-incident/internal/storage"
)

func TestShadowThenEnforceAndStaleManifest(t *testing.T) {
	root := t.TempDir()
	data := filepath.Join(root, "data")
	out := filepath.Join(root, "weeks")
	cp := filepath.Join(root, "config.yaml")
	dp := filepath.Join(root, "decisions.yaml")
	fixture, err := os.ReadFile("../../testdata/dedup/aws-window.json")
	if err != nil {
		t.Fatal(err)
	}
	write := func(path string, b []byte) {
		t.Helper()
		if err := storage.Write(path, b); err != nil {
			t.Fatal(err)
		}
	}
	write(filepath.Join(data, "2026/W09/aws.json"), fixture)
	write(dp, []byte("schema_version: 1\ndecisions: []\n"))
	configText := "dedup:\n  mode: shadow\n  auto_aws: true\n  decisions: decisions.yaml\nsources:\n  - vendor: aws\n    kind: aws-rss\n    url: https://offline.invalid\n"
	write(cp, []byte(configText))
	nowText := "2026-11-14T00:00:00Z"
	now, _ := time.Parse(time.RFC3339, nowText)
	week, err := run(cp, data, out, "2026-W09", nowText)
	if err != nil || week != "2026-W09" {
		t.Fatalf("shadow: %s %v", week, err)
	}
	w, _ := archive.ParseWeek(week)
	idx, err := archive.New(data).ReadIndex(w)
	if err != nil {
		t.Fatal(err)
	}
	if len(idx.Incidents) != 2 {
		t.Fatal("shadow changed legacy membership")
	}
	var shadow canonical.Plan
	if err := storage.Read(filepath.Join(out, "shadow-dedup.json"), &shadow); err != nil {
		t.Fatal(err)
	}
	if shadow.Statistics.Incidents != 1 {
		t.Fatal("shadow did not propose merge")
	}
	if _, err := os.Stat(filepath.Join(data, "identity/registry.json")); !os.IsNotExist(err) {
		t.Fatal("shadow saved identities")
	}
	write(cp, []byte(strings.Replace(configText, "mode: shadow", "mode: enforce", 1)))
	cfg, err := config.Load(cp)
	if err != nil {
		t.Fatal(err)
	}
	p, err := canonical.Prepare(data, cfg, dp, now, true, true)
	if err != nil {
		t.Fatal(err)
	}
	if err := canonical.Apply(data, cfg, dp, p.Plan); err != nil {
		t.Fatal(err)
	}
	for range 2 {
		if week, err := run(cp, data, out, "2026-W09", nowText); err != nil || week != "2026-W09" {
			t.Fatalf("enforce: %s %v", week, err)
		}
	}
	idx, err = archive.New(data).ReadIndex(w)
	if err != nil {
		t.Fatal(err)
	}
	if len(idx.Incidents) != 1 || idx.Incidents[0].Key != p.Incidents[0].ID || idx.Incidents[0].Score.Value != p.Plan.Changes[0].After.Score.Value {
		t.Fatal("digest disagrees with canonical plan")
	}
	digestPath := filepath.Join(out, week, "digest.md")
	before, err := os.ReadFile(digestPath)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Count(string(before), "<a id=\"") != 1 || !strings.Contains(string(before), "aws/9fb26c8e8bd2a9c4") || !strings.Contains(string(before), "aws/6a8fc17b2bdd3e4d") {
		t.Fatal("lost evidence/anchors")
	}
	write(filepath.Join(data, "canonical/incidents.json"), []byte("[]\n"))
	if _, err := run(cp, data, out, "2026-W09", nowText); err == nil {
		t.Fatal("digest accepted damaged canonical artifact")
	}
	after, err := os.ReadFile(digestPath)
	if err != nil {
		t.Fatal(err)
	}
	if string(before) != string(after) {
		t.Fatal("failed digest overwrote prior report")
	}
	source, err := os.ReadFile(filepath.Join(data, "2026/W09/aws.json"))
	if err != nil {
		t.Fatal(err)
	}
	if string(source) != string(fixture) {
		t.Fatal("digest changed sources")
	}
}
