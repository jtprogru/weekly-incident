package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/jtprogru/weekly-incident/internal/archive"
	"github.com/jtprogru/weekly-incident/internal/canonical"
	"github.com/jtprogru/weekly-incident/internal/config"
	"github.com/jtprogru/weekly-incident/internal/storage"
)

func TestCollectRetainsSourcesWhenCanonicalFails(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprint(w, `{"page":{"url":"https://example.test"},"incidents":[{"id":"one","name":"outage","status":"resolved","impact":"major","started_at":"2026-11-12T00:00:00Z","resolved_at":"2026-11-12T01:00:00Z","incident_updates":[{"id":"message","created_at":"2026-11-12T01:00:00Z","status":"resolved","body":"recovered"}]}]}`)
	}))
	defer server.Close()
	root := t.TempDir()
	data := filepath.Join(root, "data")
	cp := filepath.Join(root, "config.yaml")
	dp := filepath.Join(root, "decisions.yaml")
	result := filepath.Join(root, "result.json")
	if err := storage.Write(cp, []byte("dedup:\n  mode: enforce\n  decisions: decisions.yaml\nsources:\n  - vendor: github\n    kind: statuspage\n    url: "+server.URL+"\n")); err != nil {
		t.Fatal(err)
	}
	if err := storage.Write(dp, []byte("schema_version: 1\ndecisions: []\n")); err != nil {
		t.Fatal(err)
	}
	err := runWithResult(cp, data, "", "", "2026-11-14T00:00:00Z", result)
	if err == nil {
		t.Fatal("expected uninitialized registry failure")
	}
	var rr RunResult
	if err := storage.Read(result, &rr); err != nil {
		t.Fatal(err)
	}
	if !rr.SourceSaved || rr.CanonicalSaved {
		t.Fatalf("result %+v", rr)
	}
	a := archive.New(data)
	inv, err := a.ReadAll()
	if err != nil {
		t.Fatal(err)
	}
	if len(inv.Records) != 1 {
		t.Fatal("source lost")
	}
	state, err := a.LoadState()
	if err != nil {
		t.Fatal(err)
	}
	if state.LastCollectAt.IsZero() {
		t.Fatal("fetch state lost")
	}
	if _, err := a.ReadObservations(); err != nil {
		t.Fatal(err)
	}
	cfg, err := config.Load(cp)
	if err != nil {
		t.Fatal(err)
	}
	now := time.Date(2026, 11, 14, 0, 0, 0, 0, time.UTC)
	p, err := canonical.Prepare(data, cfg, dp, now, true, true)
	if err != nil {
		t.Fatal(err)
	}
	if err := canonical.Apply(data, cfg, dp, p.Plan); err != nil {
		t.Fatal(err)
	}
	if err := runWithResult(cp, data, "", "", "2026-11-14T00:00:00Z", result); err != nil {
		t.Fatal(err)
	}
	if err := storage.Read(result, &rr); err != nil {
		t.Fatal(err)
	}
	if !rr.CanonicalSaved {
		t.Fatal("successful canonical save not recorded")
	}
	if _, err := canonical.Load(data, cfg, dp); err != nil {
		t.Fatal(err)
	}
	// A later early failure must clear the previous success marker.
	if err := os.WriteFile(cp, []byte("invalid config"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := runWithResult(cp, data, "", "", "2026-11-14T00:00:00Z", result); err == nil {
		t.Fatal("bad config accepted")
	}
	if err := storage.Read(result, &rr); err != nil {
		t.Fatal(err)
	}
	if rr.SourceSaved {
		t.Fatal("stale success marker")
	}
}

func TestCollectRunRejectsInvalidConfig(t *testing.T) {
	if err := run(filepath.Join(t.TempDir(), "missing"), t.TempDir(), "", "", ""); err == nil {
		t.Fatal("missing config accepted")
	}
}
