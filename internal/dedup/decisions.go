package dedup

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"slices"
	"sort"
	"strings"
	"time"

	"github.com/jtprogru/weekly-incident/internal/model"
	"github.com/jtprogru/weekly-incident/internal/storage"
	"gopkg.in/yaml.v3"
)

type Decision struct {
	ID             string            `yaml:"id" json:"id"`
	Type           string            `yaml:"type" json:"type"`
	Members        []string          `yaml:"members" json:"members"`
	Reason         string            `yaml:"reason" json:"reason"`
	Evidence       []string          `yaml:"evidence" json:"evidence"`
	Author         string            `yaml:"author" json:"author"`
	At             time.Time         `yaml:"at" json:"at"`
	Supersedes     string            `yaml:"supersedes,omitempty" json:"supersedes,omitempty"`
	Field          string            `yaml:"field,omitempty" json:"field,omitempty"`
	Value          any               `yaml:"value,omitempty" json:"value,omitempty"`
	ExpectedHashes map[string]string `yaml:"expected_hashes,omitempty" json:"expected_hashes,omitempty"`
}

type Decisions struct {
	SchemaVersion int        `yaml:"schema_version" json:"schema_version"`
	Decisions     []Decision `yaml:"decisions" json:"decisions"`
}

func LoadDecisions(path string) (Decisions, error) {
	var out Decisions
	b, err := os.ReadFile(path)
	if err != nil {
		return out, err
	}
	d := yaml.NewDecoder(bytes.NewReader(b))
	d.KnownFields(true)
	if err := d.Decode(&out); err != nil {
		return out, err
	}
	var extra any
	if err := d.Decode(&extra); !errors.Is(err, io.EOF) {
		return out, fmt.Errorf("decisions must contain one YAML document")
	}
	return out, nil
}

// Active checks the whole history before resolving supersession. Unknown
// members and conflicting replacement branches cannot disappear silently.
func (d Decisions) Active(records map[string]model.Incident) ([]Decision, error) {
	if d.SchemaVersion != 1 {
		return nil, fmt.Errorf("unsupported decisions schema %d", d.SchemaVersion)
	}
	byID := map[string]Decision{}
	for _, item := range d.Decisions {
		if item.ID == "" || strings.TrimSpace(item.Author) == "" || strings.TrimSpace(item.Reason) == "" || item.At.IsZero() || len(item.Evidence) == 0 {
			return nil, fmt.Errorf("decision %q requires id, author, reason, at and evidence", item.ID)
		}
		for _, evidence := range item.Evidence {
			if strings.TrimSpace(evidence) == "" {
				return nil, fmt.Errorf("empty evidence in %s", item.ID)
			}
		}
		if _, ok := byID[item.ID]; ok {
			return nil, fmt.Errorf("duplicate decision id %s", item.ID)
		}
		minMembers := 2
		switch item.Type {
		case "same_incident", "different_incident":
			if item.Field != "" || item.Value != nil || len(item.ExpectedHashes) > 0 {
				return nil, fmt.Errorf("unexpected override fields in %s", item.ID)
			}
		case "field_override":
			minMembers = 1
			if !slices.Contains([]string{"started_at", "resolved_at", "status", "impact", "components", "title", "source_revision", "duration_unknown"}, item.Field) {
				return nil, fmt.Errorf("unsupported override field %s", item.Field)
			}
			if len(item.ExpectedHashes) == 0 {
				return nil, fmt.Errorf("override %s must pin evidence hashes", item.ID)
			}
		default:
			return nil, fmt.Errorf("unsupported decision type %s", item.Type)
		}
		members := canonicalStrings(item.Members)
		if len(members) < minMembers || len(members) != len(item.Members) {
			return nil, fmt.Errorf("invalid members in %s", item.ID)
		}
		if item.Field == "source_revision" && len(members) != 1 {
			return nil, fmt.Errorf("source_revision requires exactly one source member in %s", item.ID)
		}
		vendor := ""
		for _, key := range members {
			in, ok := records[key]
			if !ok {
				return nil, fmt.Errorf("decision %s references unknown source %s", item.ID, key)
			}
			if vendor != "" && vendor != in.Vendor {
				return nil, fmt.Errorf("cross-vendor identity decision %s", item.ID)
			}
			vendor = in.Vendor
		}
		for key := range item.ExpectedHashes {
			if _, ok := records[key]; !ok {
				return nil, fmt.Errorf("override %s pins unknown source %s", item.ID, key)
			}
		}
		item.Members = members
		byID[item.ID] = item
	}
	superseded := map[string]string{}
	for id, item := range byID {
		if item.Supersedes == "" {
			continue
		}
		prior, ok := byID[item.Supersedes]
		if !ok || prior.At.After(item.At) {
			return nil, fmt.Errorf("invalid supersedes in %s", id)
		}
		if _, ok := superseded[item.Supersedes]; ok {
			return nil, fmt.Errorf("multiple replacements for %s", item.Supersedes)
		}
		superseded[item.Supersedes] = id
		seen := map[string]bool{id: true}
		for at := item.Supersedes; at != ""; at = byID[at].Supersedes {
			if seen[at] {
				return nil, fmt.Errorf("supersession cycle at %s", at)
			}
			seen[at] = true
			if _, ok := byID[at]; !ok {
				return nil, fmt.Errorf("unknown superseded decision %s", at)
			}
		}
	}
	out := []Decision{}
	for id, item := range byID {
		if _, ok := superseded[id]; !ok {
			out = append(out, item)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out, nil
}

func (d Decisions) Hash() (string, error) {
	d.Decisions = slices.Clone(d.Decisions)
	sort.Slice(d.Decisions, func(i, j int) bool { return d.Decisions[i].ID < d.Decisions[j].ID })
	return storage.Digest(d)
}
