// Package source turns vendor status feeds into canonical incidents.
//
// Ten of the thirteen configured vendors share the Statuspage schema and go
// through one parser. Slack, GCP and AWS each need their own. A parser never
// aborts a whole feed over one malformed record: it skips the record, counts it
// in Result.ParseErrors and carries on, because the two undocumented feeds
// (AWS RSS, Slack history) can change shape without warning.
package source

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/jtprogru/weekly-incident/internal/config"
	"github.com/jtprogru/weekly-incident/internal/model"
)

// Result is what one vendor fetch produced.
type Result struct {
	Incidents   []model.Incident
	ParseErrors int
}

// Source fetches and normalizes one vendor's incidents.
type Source interface {
	Fetch(ctx context.Context, f *Fetcher) (Result, error)
}

// base is the state every parser needs and nothing more. Without it the same
// two fields were declared four times over.
type base struct {
	vendor string
	url    string
}

// New builds the parser for a configured source.
func New(s config.Source) (Source, error) {
	b := base{vendor: s.Vendor, url: s.URL}
	switch s.Kind {
	case config.KindStatuspage:
		return &statuspage{base: b}, nil
	case config.KindSlack:
		return &slack{base: b}, nil
	case config.KindGCP:
		return &gcp{base: b}, nil
	case config.KindAWSRSS:
		return &awsRSS{base: b}, nil
	default:
		return nil, fmt.Errorf("unknown source kind %q", s.Kind)
	}
}

// Fetcher performs HTTP GETs with the project's User-Agent and retry policy.
// Thirteen requests a day need no concurrency and no aggression.
type Fetcher struct {
	Client    *http.Client
	UserAgent string
	Retries   int
	Backoff   time.Duration
	// Sleep is swappable so tests do not wait on real backoff.
	Sleep func(time.Duration)
}

// NewFetcher builds a Fetcher from config.
func NewFetcher(c *config.Config) *Fetcher {
	return &Fetcher{
		Client:    &http.Client{Timeout: c.HTTP.Timeout.D()},
		UserAgent: c.UserAgent,
		Retries:   c.HTTP.Retries,
		Backoff:   c.HTTP.BackoffBase.D(),
		Sleep:     time.Sleep,
	}
}

// maxBody caps a response at 32 MiB. The largest feed measured was 800 KB, so
// anything near the cap means the endpoint stopped being a status feed.
const maxBody = 32 << 20

// Get fetches a URL, retrying on network errors and 5xx with exponential
// backoff. A 4xx is returned immediately: retrying a 404 only wastes time.
func (f *Fetcher) Get(ctx context.Context, url string) ([]byte, error) {
	attempts := f.Retries
	if attempts < 1 {
		attempts = 1
	}
	var lastErr error
	for i := range attempts {
		if i > 0 {
			delay := f.Backoff * (1 << (i - 1))
			if f.Sleep != nil {
				f.Sleep(delay)
			}
		}
		b, err := f.get(ctx, url)
		if err == nil {
			return b, nil
		}
		lastErr = err
		var he httpError
		if errors.As(err, &he) && he.Code < 500 {
			return nil, err
		}
		if ctx.Err() != nil {
			return nil, err
		}
	}
	return nil, fmt.Errorf("after %d attempts: %w", attempts, lastErr)
}

func (f *Fetcher) get(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", f.UserAgent)
	req.Header.Set("Accept", "application/json, application/rss+xml, application/xml;q=0.9, */*;q=0.8")

	resp, err := f.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, httpError{Code: resp.StatusCode, URL: url}
	}
	b, err := io.ReadAll(io.LimitReader(resp.Body, maxBody))
	if err != nil {
		return nil, fmt.Errorf("read body: %w", err)
	}
	return b, nil
}

type httpError struct {
	Code int
	URL  string
}

func (e httpError) Error() string { return fmt.Sprintf("GET %s: status %d", e.URL, e.Code) }
