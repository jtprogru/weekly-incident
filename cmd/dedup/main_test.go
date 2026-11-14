package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/jtprogru/weekly-incident/internal/dedup"
)

const pinnedTime = "2026-11-14T00:00:00Z"

func TestAuditCommandOfflineAndIdempotent(t *testing.T) {
	root, data, cfg := commandFixture(t)
	out := filepath.Join(root, "reports")
	before := treeBytes(t, data)
	weeksBefore := treeBytes(t, filepath.Join(root, "weeks"))
	args := []string{"audit", "-data", data, "-config", cfg, "-out", out, "-now", pinnedTime, "-revision", "test-fixture"}
	var stderr bytes.Buffer
	if err := run(args, &stderr); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(stderr.String(), "2 records, 1 vendors, 1 candidate pairs") {
		t.Fatalf("stderr = %s", stderr.String())
	}
	reports := treeBytes(t, out)
	var report dedup.Report
	if err := json.Unmarshal([]byte(reports["audit.json"]), &report); err != nil {
		t.Fatal(err)
	}
	if len(report.Candidates[0].SharedUpdates) != 16 || report.ConfigSHA256 == "" || report.InputFingerprint == "" || report.RevisionLabel != "test-fixture" {
		t.Fatalf("unexpected report: %+v", report.Statistics)
	}
	if report.GeneratedAt.Format(time.RFC3339) != pinnedTime {
		t.Fatal("-now ignored")
	}
	oldTime := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	for _, name := range []string{"audit.json", "audit.md"} {
		if err := os.Chtimes(filepath.Join(out, name), oldTime, oldTime); err != nil {
			t.Fatal(err)
		}
	}
	if err := run(args, &stderr); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(reports, treeBytes(t, out)) {
		t.Fatal("repeat changed report bytes")
	}
	for _, name := range []string{"audit.json", "audit.md"} {
		info, err := os.Stat(filepath.Join(out, name))
		if err != nil {
			t.Fatal(err)
		}
		if !info.ModTime().Equal(oldTime) {
			t.Fatal("identical report was rewritten")
		}
	}
	if !reflect.DeepEqual(before, treeBytes(t, data)) || !reflect.DeepEqual(weeksBefore, treeBytes(t, filepath.Join(root, "weeks"))) {
		t.Fatal("audit modified archive or weeks")
	}
}

func TestAuditRefusesProtectedOutputs(t *testing.T) {
	root, data, cfg := commandFixture(t)
	alias := filepath.Join(root, "alias")
	if err := os.Symlink(data, alias); err != nil {
		t.Fatal(err)
	}
	for _, out := range []string{data, filepath.Join(data, "report"), root, filepath.Join(root, "weeks"), filepath.Join(alias, "report")} {
		t.Run(out, func(t *testing.T) {
			before := treeBytes(t, data)
			var stderr bytes.Buffer
			err := run([]string{"audit", "-data", data, "-config", cfg, "-out", out, "-now", pinnedTime}, &stderr)
			if err == nil || !strings.Contains(err.Error(), "overlaps protected") {
				t.Fatalf("error = %v", err)
			}
			if !reflect.DeepEqual(before, treeBytes(t, data)) {
				t.Fatal("archive changed")
			}
		})
	}
}

func TestAuditErrorsDoNotProduceReports(t *testing.T) {
	for _, tc := range []struct {
		name    string
		args    []string
		corrupt bool
	}{
		{"missing-out", []string{"audit"}, false},
		{"unknown-command", []string{"apply"}, false},
		{"invalid-clock", []string{"audit", "-now", "bad-time"}, false},
		{"unknown-flag", []string{"audit", "-no-such-flag"}, false},
		{"trailing-argument", []string{"audit", "extra"}, false},
		{"bad-vendor-json", []string{"audit"}, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root, data, cfg := commandFixture(t)
			out := filepath.Join(root, "reports")
			args := append([]string{}, tc.args...)
			if tc.name != "missing-out" {
				args = append(args, "-out", out)
			}
			args = append(args, "-data", data, "-config", cfg)
			if tc.corrupt {
				if err := os.WriteFile(filepath.Join(data, "2026/W09/aws.json"), []byte("{"), 0o644); err != nil {
					t.Fatal(err)
				}
			}
			var stderr bytes.Buffer
			if err := run(args, &stderr); err == nil {
				t.Fatal("expected error")
			}
			if _, err := os.Stat(out); !os.IsNotExist(err) {
				t.Fatalf("report created on invalid input: %v", err)
			}
		})
	}
}

func TestAuditRefusesSymlinkReport(t *testing.T) {
	root, data, cfg := commandFixture(t)
	out := filepath.Join(root, "reports")
	if err := os.Mkdir(out, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(filepath.Join(data, "2026/W09/aws.json"), filepath.Join(out, "audit.md")); err != nil {
		t.Fatal(err)
	}
	before := treeBytes(t, data)
	var stderr bytes.Buffer
	err := run([]string{"audit", "-data", data, "-config", cfg, "-out", out}, &stderr)
	if err == nil || !strings.Contains(err.Error(), "not a regular file") {
		t.Fatalf("error = %v", err)
	}
	if _, err := os.Stat(filepath.Join(out, "audit.json")); !os.IsNotExist(err) {
		t.Fatal("wrote one report before checking all destinations")
	}
	if !reflect.DeepEqual(before, treeBytes(t, data)) {
		t.Fatal("archive changed")
	}
}

func TestAuditRefusesOutputContainingConfig(t *testing.T) {
	root, data, cfg := commandFixture(t)
	out := filepath.Join(root, "reports")
	if err := os.Mkdir(out, 0o755); err != nil {
		t.Fatal(err)
	}
	b, err := os.ReadFile(cfg)
	if err != nil {
		t.Fatal(err)
	}
	input := filepath.Join(out, "audit.json")
	if err := os.WriteFile(input, b, 0o644); err != nil {
		t.Fatal(err)
	}
	var stderr bytes.Buffer
	if err := run([]string{"audit", "-data", data, "-config", input, "-out", out}, &stderr); err == nil {
		t.Fatal("accepted overlapping config input")
	}
	after, err := os.ReadFile(input)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(b, after) {
		t.Fatal("config overwritten")
	}
}

func TestAuditHelp(t *testing.T) {
	var stderr bytes.Buffer
	if err := run([]string{"audit", "-h"}, &stderr); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(stderr.String(), "-out") {
		t.Fatal("help missing")
	}
}

func commandFixture(t *testing.T) (root, data, cfg string) {
	t.Helper()
	root = t.TempDir()
	data = filepath.Join(root, "data")
	dir := filepath.Join(data, "2026/W09")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	b, err := os.ReadFile("../../testdata/dedup/aws-window.json")
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "aws.json"), b, 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.Mkdir(filepath.Join(root, "weeks"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "weeks/digest.md"), []byte("unchanged"), 0o644); err != nil {
		t.Fatal(err)
	}
	cfg = filepath.Join(root, "config.yaml")
	// A deliberately unusable URL proves that auditing needs no live endpoint.
	if err := os.WriteFile(cfg, []byte("sources:\n  - vendor: aws\n    kind: aws-rss\n    url: https://offline.invalid/feed\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	return root, data, cfg
}

func treeBytes(t *testing.T, root string) map[string]string {
	t.Helper()
	out := map[string]string{}
	err := filepath.WalkDir(root, func(path string, entry os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() {
			return nil
		}
		b, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		out[rel] = string(b)
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	return out
}
