// Package dedup audits source records without merging them. Its evidence is
// review material, not a claim about the number of unique real-world outages.
package dedup

import (
	"net/url"
	"regexp"
	"strings"

	"github.com/jtprogru/weekly-incident/internal/config"
)

// Policy records what the current parser preserves. Update IDs are compared
// only as candidate evidence: their reuse across incidents is not proof of
// identity. No policy in audit-v1 enables automatic merging.
type Policy struct {
	Vendor           string      `json:"vendor"`
	Kind             config.Kind `json:"kind"`
	StableIncidentID bool        `json:"stable_incident_id"`
	UpdateIDs        bool        `json:"update_ids_available"`
	EventURLs        bool        `json:"event_urls_available"`
}

func policyFor(s config.Source) Policy {
	p := Policy{Vendor: s.Vendor, Kind: s.Kind}
	switch s.Kind {
	case config.KindStatuspage, config.KindYandexCloud:
		p.StableIncidentID, p.UpdateIDs, p.EventURLs = true, true, true
	case config.KindSlack, config.KindGCP:
		p.StableIncidentID, p.EventURLs = true, true
	}
	return p
}

var slackEventPath = regexp.MustCompile(`^/[0-9]{4}-[0-9]{2}/[^/]+/?$`)

// eventURL accepts only incident-shaped links produced by the known adapters.
// Query parameters stay intact. Neither redirects nor http/https aliases are
// guessed, and a shared status-page root can never identify an incident.
func eventURL(raw string, p Policy) string {
	if !p.EventURLs {
		return ""
	}
	u, err := url.Parse(raw)
	if err != nil || u.Host == "" || u.User != nil || (u.Scheme != "https" && u.Scheme != "http") {
		return ""
	}
	host := strings.ToLower(u.Hostname())
	path := strings.TrimSuffix(u.Path, "/")
	i := strings.LastIndex(path, "/incidents/")
	incidentPath := i >= 0 && len(path[i+len("/incidents/"):]) > 0 && !strings.Contains(path[i+len("/incidents/"):], "/")
	valid := false
	switch p.Kind {
	case config.KindStatuspage:
		valid = incidentPath || (host == "stspg.io" && len(path) > 1 && !strings.Contains(path[1:], "/"))
	case config.KindSlack:
		valid = host == "slack-status.com" && slackEventPath.MatchString(u.Path)
	case config.KindGCP:
		valid = host == "status.cloud.google.com" && incidentPath
	case config.KindYandexCloud:
		valid = (host == "status.yandex.cloud" || host == "status.cloud.yandex.ru") && incidentPath
	}
	if !valid {
		return ""
	}
	u.Host = strings.ToLower(u.Host)
	u.Fragment, u.RawFragment = "", ""
	return u.String()
}
