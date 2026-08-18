package source

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/jtprogru/weekly-incident/internal/config"
)

// testFetcher builds a Fetcher that does not actually sleep between retries.
func testFetcher(t *testing.T) *Fetcher {
	t.Helper()
	return &Fetcher{
		Client:    &http.Client{Timeout: 5 * time.Second},
		UserAgent: "weekly-incident/test",
		Retries:   3,
		Backoff:   time.Millisecond,
		Sleep:     func(time.Duration) {},
	}
}

func TestFetcherRetriesServerErrors(t *testing.T) {
	var calls int
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		if calls < 3 {
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		_, _ = w.Write([]byte(`{"ok":true}`))
	}))
	defer srv.Close()

	b, err := testFetcher(t).Get(context.Background(), srv.URL)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if string(b) != `{"ok":true}` {
		t.Errorf("body = %q", b)
	}
	if calls != 3 {
		t.Errorf("made %d calls, want 3", calls)
	}
}

func TestFetcherDoesNotRetryClientErrors(t *testing.T) {
	// status.stripe.com answers 404; retrying that three times only wastes the
	// run's time.
	var calls int
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		w.WriteHeader(http.StatusNotFound)
	}))
	defer srv.Close()

	if _, err := testFetcher(t).Get(context.Background(), srv.URL); err == nil {
		t.Fatal("Get succeeded on 404, want an error")
	}
	if calls != 1 {
		t.Errorf("made %d calls, want 1", calls)
	}
}

func TestFetcherGivesUpAfterRetries(t *testing.T) {
	var calls int
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer srv.Close()

	if _, err := testFetcher(t).Get(context.Background(), srv.URL); err == nil {
		t.Fatal("Get succeeded, want an error")
	}
	if calls != 3 {
		t.Errorf("made %d calls, want 3", calls)
	}
}

func TestFetcherSendsUserAgent(t *testing.T) {
	// A meaningful User-Agent is what separates polite collection from
	// scraping, and it is the thing that keeps the crawler unbanned.
	got := make(chan string, 1)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got <- r.Header.Get("User-Agent")
		_, _ = w.Write([]byte("{}"))
	}))
	defer srv.Close()

	if _, err := testFetcher(t).Get(context.Background(), srv.URL); err != nil {
		t.Fatalf("Get: %v", err)
	}
	if ua := <-got; ua != "weekly-incident/test" {
		t.Errorf("User-Agent = %q", ua)
	}
}

func TestNewDispatchesEveryKind(t *testing.T) {
	want := map[config.Kind]string{
		config.KindStatuspage: "*source.statuspage",
		config.KindSlack:      "*source.slack",
		config.KindGCP:        "*source.gcp",
		config.KindAWSRSS:     "*source.awsRSS",
	}
	for k, typeName := range want {
		s, err := New(config.Source{Vendor: "v", Kind: k, URL: "https://example.com"})
		if err != nil {
			t.Errorf("New(%q): %v", k, err)
			continue
		}
		if got := fmt.Sprintf("%T", s); got != typeName {
			t.Errorf("New(%q) built %s, want %s", k, got, typeName)
		}
	}
	if _, err := New(config.Source{Vendor: "v", Kind: "nope", URL: "https://example.com"}); err == nil {
		t.Error("New accepted an unknown kind")
	}
}
