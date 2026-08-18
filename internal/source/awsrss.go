package source

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/jtprogru/weekly-incident/internal/model"
)

// awsStatusURL is the only link worth recording. The feed's own <link> is
// always the status page root and carries no per-event information.
const awsStatusURL = "https://health.aws.amazon.com/health/status"

// awsEventGap splits one service's update stream into separate events.
//
// The threshold is picked from the shape of the real feed, measured 2026-08-18:
// inside an event AWS posts every 1 to 7 hours (the longest observed gap was
// 6.8h), while consecutive events on the same service key were 1383 hours
// apart. Anything between those two scales works; 24h sits comfortably in the
// middle.
const awsEventGap = 24 * time.Hour

// awsRSS parses status.aws.amazon.com/rss/all.rss.
//
// The feed emits updates, not incidents: one outage shows up as a chain of
// <item> entries sharing a service key. The guid looks like it encodes the
// event, but it does not — the epoch in it is the item's own pubDate, verified
// to the second across every item in the feed. So the only grouping signal the
// feed gives is the service key:
//
//	https://status.aws.amazon.com/#directconnect-eu-central-1_1787025679
//	                               └──── service key ────┘ └ this item ┘
//
// An event therefore ends either at a resolution headline, or at a gap longer
// than awsEventGap. Both are needed: AWS does not always post a resolution, and
// two of the three chains in the measured feed were closed by the gap alone.
//
// The documented replacement endpoint, health.aws.amazon.com/public/currentevents,
// is a current-state snapshot (three events when measured) with no history, so
// this legacy feed is the one that ships.
type awsRSS struct {
	base
}

type awsFeed struct {
	XMLName xml.Name  `xml:"rss"`
	Items   []awsItem `xml:"channel>item"`
}

type awsItem struct {
	Title       string `xml:"title"       json:"title"`
	Link        string `xml:"link"        json:"link"`
	PubDate     string `xml:"pubDate"     json:"pub_date"`
	GUID        string `xml:"guid"        json:"guid"`
	Description string `xml:"description" json:"description"`
}

// awsEntry is one parsed feed item, ready to be grouped.
type awsEntry struct {
	item    awsItem
	service string
	at      time.Time
}

func (a *awsRSS) Fetch(ctx context.Context, f *Fetcher) (Result, error) {
	b, err := f.Get(ctx, a.url)
	if err != nil {
		return Result{}, err
	}
	return a.parse(b)
}

func (a *awsRSS) parse(b []byte) (Result, error) {
	var feed awsFeed
	if err := xml.Unmarshal(b, &feed); err != nil {
		return Result{}, fmt.Errorf("%s: decode feed: %w", a.vendor, err)
	}

	var res Result
	byService := make(map[string][]awsEntry)
	for _, it := range feed.Items {
		service, err := parseAWSService(it.GUID)
		if err != nil {
			res.ParseErrors++
			continue
		}
		at, err := parseAWSTime(it.PubDate)
		if err != nil {
			res.ParseErrors++
			continue
		}
		byService[service] = append(byService[service], awsEntry{item: it, service: service, at: at})
	}

	services := make([]string, 0, len(byService))
	for s := range byService {
		services = append(services, s)
	}
	sort.Strings(services)

	for _, service := range services {
		entries := byService[service]
		sort.SliceStable(entries, func(i, j int) bool { return entries[i].at.Before(entries[j].at) })
		for _, chain := range splitAWSEvents(entries) {
			in, err := a.normalize(chain)
			if err != nil {
				res.ParseErrors++
				continue
			}
			res.Incidents = append(res.Incidents, in)
		}
	}

	sort.SliceStable(res.Incidents, func(i, j int) bool {
		if !res.Incidents[i].StartedAt.Equal(res.Incidents[j].StartedAt) {
			return res.Incidents[i].StartedAt.Before(res.Incidents[j].StartedAt)
		}
		return res.Incidents[i].NativeID < res.Incidents[j].NativeID
	})
	return res, nil
}

// splitAWSEvents cuts one service's chronologically sorted updates into events.
// A resolution headline closes the current event; so does a gap wider than
// awsEventGap.
func splitAWSEvents(entries []awsEntry) [][]awsEntry {
	var out [][]awsEntry
	var cur []awsEntry
	for i, e := range entries {
		if len(cur) > 0 {
			gap := e.at.Sub(entries[i-1].at)
			if gap > awsEventGap || isAWSResolved(entries[i-1].item.Title) {
				out = append(out, cur)
				cur = nil
			}
		}
		cur = append(cur, e)
	}
	if len(cur) > 0 {
		out = append(out, cur)
	}
	return out
}

func (a *awsRSS) normalize(chain []awsEntry) (model.Incident, error) {
	if len(chain) == 0 {
		return model.Incident{}, fmt.Errorf("empty aws chain")
	}
	first, last := chain[0], chain[len(chain)-1]

	updates := make([]model.Update, 0, len(chain))
	items := make([]awsItem, 0, len(chain))
	resolved := false
	for _, e := range chain {
		items = append(items, e.item)
		if isAWSResolved(e.item.Title) {
			resolved = true
		}
		updates = append(updates, model.Update{
			Status:    awsUpdateStatus(e.item.Title),
			Body:      strings.TrimSpace(e.item.Description),
			CreatedAt: e.at,
		})
	}

	// AWS publishes no stable event identifier, so the key is derived from the
	// service and the first update visible in the feed. See the known
	// limitations in README.md: a chain straddling the tail of the 50-item
	// window can be re-keyed and land in the archive twice.
	seed := first.service + "_" + strconv.FormatInt(first.at.Unix(), 10)
	sum := sha256.Sum256([]byte(seed))

	status := model.StatusInvestigating
	var resolvedAt *time.Time
	if resolved {
		status = model.StatusResolved
		t := last.at
		resolvedAt = &t
	}

	raw, err := json.Marshal(items)
	if err != nil {
		return model.Incident{}, fmt.Errorf("aws %s: marshal raw: %w", seed, err)
	}

	in := model.Incident{
		Vendor:   a.vendor,
		NativeID: hex.EncodeToString(sum[:])[:16],
		Title:    cleanAWSTitle(last.item.Title),
		URL:      awsStatusURL,
		// AWS reports no severity in any machine-readable form.
		Impact:     model.ImpactUnknown,
		Status:     status,
		StartedAt:  first.at,
		ResolvedAt: resolvedAt,
		Components: []string{first.service},
		Updates:    updates,
		Raw:        raw,
	}
	in.Finalize()
	if err := in.Validate(); err != nil {
		return model.Incident{}, err
	}
	return in, nil
}

// parseAWSService pulls the service key out of the guid fragment. The trailing
// "_<epoch>" is the item's own timestamp and carries no event information.
func parseAWSService(guid string) (string, error) {
	i := strings.LastIndex(guid, "#")
	if i < 0 {
		return "", fmt.Errorf("guid %q has no fragment", guid)
	}
	frag := guid[i+1:]
	j := strings.LastIndex(frag, "_")
	if j <= 0 {
		return "", fmt.Errorf("guid fragment %q is not <service>_<epoch>", frag)
	}
	if _, err := strconv.ParseInt(frag[j+1:], 10, 64); err != nil {
		return "", fmt.Errorf("guid fragment %q: bad epoch: %w", frag, err)
	}
	return frag[:j], nil
}

// awsZones maps the timezone abbreviations the feed uses onto numeric offsets.
// Go cannot resolve bare abbreviations reliably, so they are substituted before
// parsing. An unknown abbreviation is an error, never a silent UTC.
var awsZones = map[string]string{
	"UTC": "+0000", "GMT": "+0000",
	"PST": "-0800", "PDT": "-0700",
	"MST": "-0700", "MDT": "-0600",
	"CST": "-0600", "CDT": "-0500",
	"EST": "-0500", "EDT": "-0400",
}

// parseAWSTime reads "Mon, 17 Aug 2026 21:01:19 PDT".
func parseAWSTime(s string) (time.Time, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return time.Time{}, fmt.Errorf("empty pubDate")
	}
	if i := strings.LastIndex(s, " "); i > 0 {
		zone := s[i+1:]
		if off, ok := awsZones[zone]; ok {
			s = s[:i+1] + off
		} else if _, err := strconv.Atoi(zone); err != nil {
			return time.Time{}, fmt.Errorf("pubDate %q: unknown timezone %q", s, zone)
		}
	}
	t, err := time.Parse(time.RFC1123Z, s)
	if err != nil {
		return time.Time{}, fmt.Errorf("pubDate %q: %w", s, err)
	}
	return t.UTC(), nil
}

// awsTitlePrefixes are the fixed prefixes AWS puts on every headline.
var awsTitlePrefixes = []string{
	"Service is operating normally:",
	"Informational message:",
	"Service disruption:",
	"Service degradation:",
	"Service impact:",
}

func cleanAWSTitle(s string) string {
	s = strings.TrimSpace(s)
	for _, p := range awsTitlePrefixes {
		if strings.HasPrefix(s, p) {
			s = strings.TrimSpace(strings.TrimPrefix(s, p))
			break
		}
	}
	return strings.TrimSpace(strings.TrimPrefix(s, "[RESOLVED]"))
}

func isAWSResolved(title string) bool {
	t := strings.TrimSpace(title)
	return strings.Contains(t, "[RESOLVED]") || strings.HasPrefix(t, "Service is operating normally")
}

func awsUpdateStatus(title string) model.Status {
	if isAWSResolved(title) {
		return model.StatusResolved
	}
	return model.StatusInvestigating
}
