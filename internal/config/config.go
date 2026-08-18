// Package config loads config.yaml: the source list, the fetch policy and the
// digest thresholds. The file is versioned alongside the code and nothing in it
// is overridable at runtime; the notification secrets live in the workflow, not
// in Go.
package config

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"

	"github.com/jtprogru/weekly-incident/internal/model"
)

// Kind identifies which parser handles a source. Ten of the thirteen vendors
// are KindStatuspage; the other three each need their own.
type Kind string

const (
	KindStatuspage Kind = "statuspage"
	KindSlack      Kind = "slack"
	KindGCP        Kind = "gcp"
	KindAWSRSS     Kind = "aws-rss"
)

// Duration wraps time.Duration so durations can be written as "30s" in YAML,
// which yaml.v3 does not support out of the box.
type Duration time.Duration

// UnmarshalYAML parses a Go duration string.
func (d *Duration) UnmarshalYAML(n *yaml.Node) error {
	var s string
	if err := n.Decode(&s); err != nil {
		return err
	}
	v, err := time.ParseDuration(s)
	if err != nil {
		return fmt.Errorf("parse duration %q: %w", s, err)
	}
	*d = Duration(v)
	return nil
}

// D returns the wrapped duration.
func (d Duration) D() time.Duration { return time.Duration(d) }

// Source is one vendor entry. Weight is editorial, not technical: an AWS outage
// is more interesting to the audience than a DigitalOcean one of equal length.
type Source struct {
	Vendor string  `yaml:"vendor"`
	Kind   Kind    `yaml:"kind"`
	URL    string  `yaml:"url"`
	Weight float64 `yaml:"weight"`
}

// HTTP is the fetch policy applied to every source.
type HTTP struct {
	Timeout     Duration `yaml:"timeout"`
	Retries     int      `yaml:"retries"`
	BackoffBase Duration `yaml:"backoff_base"`
}

// Collect governs the daily run.
type Collect struct {
	// FailThreshold is the fraction of unreachable sources above which the job
	// goes red. Below it, a dead vendor is a note in the digest, not a failure.
	FailThreshold float64 `yaml:"fail_threshold"`
	// GapWarnDays is how stale the archive may get before the run warns about
	// probable data loss. Feed depth is finite and there is no pagination, so a
	// long gap is unrecoverable.
	GapWarnDays int `yaml:"gap_warn_days"`
}

// Digest governs the weekly render.
type Digest struct {
	// An incident clears the threshold if it ran at least MinDurationMinutes OR
	// its impact is at least MinImpact. Disjunction, not conjunction: a short
	// critical outage is more interesting than a long minor one.
	MinDurationMinutes int          `yaml:"min_duration_minutes"`
	MinImpact          model.Impact `yaml:"min_impact"`
	BriefTopN          int          `yaml:"brief_top_n"`
	BriefBodyLimit     int          `yaml:"brief_body_limit"`
}

// Config is the whole file.
type Config struct {
	UserAgent string   `yaml:"user_agent"`
	HTTP      HTTP     `yaml:"http"`
	Collect   Collect  `yaml:"collect"`
	Digest    Digest   `yaml:"digest"`
	Sources   []Source `yaml:"sources"`
}

// Load reads and validates config.yaml.
func Load(path string) (*Config, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}
	var c Config
	if err := yaml.Unmarshal(b, &c); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}
	c.applyDefaults()
	if err := c.Validate(); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}
	return &c, nil
}

func (c *Config) applyDefaults() {
	if c.UserAgent == "" {
		c.UserAgent = "weekly-incident/0.1 (+https://github.com/jtprogru/weekly-incident)"
	}
	if c.HTTP.Timeout == 0 {
		c.HTTP.Timeout = Duration(30 * time.Second)
	}
	if c.HTTP.Retries == 0 {
		c.HTTP.Retries = 3
	}
	if c.HTTP.BackoffBase == 0 {
		c.HTTP.BackoffBase = Duration(2 * time.Second)
	}
	if c.Collect.FailThreshold == 0 {
		c.Collect.FailThreshold = 0.5
	}
	if c.Collect.GapWarnDays == 0 {
		c.Collect.GapWarnDays = 14
	}
	if c.Digest.MinImpact == "" {
		c.Digest.MinImpact = model.ImpactMajor
	}
	if c.Digest.BriefTopN == 0 {
		c.Digest.BriefTopN = 5
	}
	if c.Digest.BriefBodyLimit == 0 {
		c.Digest.BriefBodyLimit = 800
	}
	for i := range c.Sources {
		if c.Sources[i].Weight == 0 {
			c.Sources[i].Weight = 1
		}
	}
}

// Validate rejects a config that would fail confusingly at fetch time.
func (c *Config) Validate() error {
	if len(c.Sources) == 0 {
		return fmt.Errorf("no sources configured")
	}
	seen := make(map[string]bool, len(c.Sources))
	for _, s := range c.Sources {
		switch {
		case s.Vendor == "":
			return fmt.Errorf("source with empty vendor")
		case seen[s.Vendor]:
			return fmt.Errorf("duplicate vendor %q", s.Vendor)
		case s.URL == "":
			return fmt.Errorf("source %q: empty url", s.Vendor)
		}
		switch s.Kind {
		case KindStatuspage, KindSlack, KindGCP, KindAWSRSS:
		default:
			return fmt.Errorf("source %q: unknown kind %q", s.Vendor, s.Kind)
		}
		seen[s.Vendor] = true
	}
	if c.Collect.FailThreshold <= 0 || c.Collect.FailThreshold > 1 {
		return fmt.Errorf("collect.fail_threshold must be in (0,1]")
	}
	switch c.Digest.MinImpact {
	case model.ImpactNone, model.ImpactMinor, model.ImpactMajor, model.ImpactCritical:
	default:
		return fmt.Errorf("digest.min_impact %q is not a real impact level", c.Digest.MinImpact)
	}
	return nil
}

// Weights maps vendor to editorial weight, for the scorer.
func (c *Config) Weights() map[string]float64 {
	m := make(map[string]float64, len(c.Sources))
	for _, s := range c.Sources {
		m[s.Vendor] = s.Weight
	}
	return m
}
