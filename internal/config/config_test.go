package config

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/jtprogru/weekly-incident/internal/model"
)

// TestLoadShippedConfig guards the file the workflows actually run with.
func TestLoadShippedConfig(t *testing.T) {
	c, err := Load(filepath.Join("..", "..", "config.yaml"))
	if err != nil {
		t.Fatalf("load config.yaml: %v", err)
	}
	if len(c.Sources) != 13 {
		t.Errorf("got %d sources, want 13", len(c.Sources))
	}

	byKind := map[Kind]int{}
	for _, s := range c.Sources {
		byKind[s.Kind]++
		if s.Weight <= 0 {
			t.Errorf("source %q has weight %v", s.Vendor, s.Weight)
		}
	}
	// Ten of the thirteen vendors share the Statuspage schema; the other three
	// each need their own parser.
	if byKind[KindStatuspage] != 10 {
		t.Errorf("got %d statuspage sources, want 10", byKind[KindStatuspage])
	}
	for _, k := range []Kind{KindSlack, KindGCP, KindAWSRSS} {
		if byKind[k] != 1 {
			t.Errorf("got %d %q sources, want 1", byKind[k], k)
		}
	}
	if c.HTTP.Timeout.D() != 30*time.Second {
		t.Errorf("http.timeout = %v", c.HTTP.Timeout.D())
	}
	if c.Digest.MinImpact != model.ImpactMajor {
		t.Errorf("digest.min_impact = %q", c.Digest.MinImpact)
	}
	if w := c.Weights()["aws"]; w != 2.0 {
		t.Errorf("aws weight = %v, want 2.0", w)
	}
}

func TestDefaults(t *testing.T) {
	c := writeTemp(t, `
sources:
  - vendor: github
    kind: statuspage
    url: https://example.com/api/v2/incidents.json
`)
	if c.HTTP.Retries != 3 {
		t.Errorf("retries = %d, want the default 3", c.HTTP.Retries)
	}
	if c.Collect.FailThreshold != 0.5 {
		t.Errorf("fail_threshold = %v, want the default 0.5", c.Collect.FailThreshold)
	}
	if c.Sources[0].Weight != 1 {
		t.Errorf("weight = %v, want the default 1", c.Sources[0].Weight)
	}
}

func TestValidateRejects(t *testing.T) {
	cases := map[string]string{
		"no sources": `sources: []`,
		"unknown kind": `
sources:
  - vendor: x
    kind: carrier-pigeon
    url: https://example.com`,
		"duplicate vendor": `
sources:
  - vendor: x
    kind: statuspage
    url: https://example.com/a
  - vendor: x
    kind: statuspage
    url: https://example.com/b`,
		"empty url": `
sources:
  - vendor: x
    kind: statuspage
    url: ""`,
		"impossible threshold": `
collect:
  fail_threshold: 1.5
sources:
  - vendor: x
    kind: statuspage
    url: https://example.com`,
		"unknown min_impact": `
digest:
  min_impact: catastrophic
sources:
  - vendor: x
    kind: statuspage
    url: https://example.com`,
	}
	for name, body := range cases {
		path := filepath.Join(t.TempDir(), "config.yaml")
		if err := os.WriteFile(path, []byte(body), 0o600); err != nil {
			t.Fatal(err)
		}
		if _, err := Load(path); err == nil {
			t.Errorf("%s: Load succeeded, want an error", name)
		}
	}
}

func TestDurationParsing(t *testing.T) {
	c := writeTemp(t, `
http:
  timeout: 45s
  backoff_base: 1500ms
sources:
  - vendor: x
    kind: statuspage
    url: https://example.com
`)
	if c.HTTP.Timeout.D() != 45*time.Second {
		t.Errorf("timeout = %v", c.HTTP.Timeout.D())
	}
	if c.HTTP.BackoffBase.D() != 1500*time.Millisecond {
		t.Errorf("backoff_base = %v", c.HTTP.BackoffBase.D())
	}
}

func writeTemp(t *testing.T, body string) *Config {
	t.Helper()
	path := filepath.Join(t.TempDir(), "config.yaml")
	if err := os.WriteFile(path, []byte(body), 0o600); err != nil {
		t.Fatal(err)
	}
	c, err := Load(path)
	if err != nil {
		t.Fatalf("load: %v", err)
	}
	return c
}
