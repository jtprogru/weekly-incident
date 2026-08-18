// Command collect fetches every configured status feed and merges what it
// finds into the archive.
//
// It runs daily rather than weekly. Feed depth is finite and unpaginated —
// 50 incidents, which is 20 days for Cloudflare — so a weekly cadence would
// leave only two missed runs of slack before data is lost for good.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sort"
	"strings"
	"syscall"
	"time"

	"github.com/jtprogru/weekly-incident/internal/archive"
	"github.com/jtprogru/weekly-incident/internal/clock"
	"github.com/jtprogru/weekly-incident/internal/config"
	"github.com/jtprogru/weekly-incident/internal/model"
	"github.com/jtprogru/weekly-incident/internal/score"
	"github.com/jtprogru/weekly-incident/internal/source"
)

func main() {
	var (
		configPath  = flag.String("config", "config.yaml", "path to config.yaml")
		dataDir     = flag.String("data", "data", "archive root directory")
		summaryPath = flag.String("summary", "", "write a one-line run summary here (for the commit message)")
		warningPath = flag.String("warning", "", "write an archive-gap warning here, if there is one (for the alert)")
		nowFlag     = flag.String("now", "", "override the current time, RFC3339 (for reproducible runs)")
	)
	flag.Parse()

	log.SetFlags(0)
	if err := run(*configPath, *dataDir, *summaryPath, *warningPath, *nowFlag); err != nil {
		log.Fatalf("collect: %v", err)
	}
}

func run(configPath, dataDir, summaryPath, warningPath, nowFlag string) error {
	cfg, err := config.Load(configPath)
	if err != nil {
		return err
	}
	now, err := clock.Now(nowFlag)
	if err != nil {
		return err
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	arc := archive.New(dataDir)
	state, err := arc.LoadState()
	if err != nil {
		return err
	}

	// The gap is reported before anything else runs: it describes data that is
	// already unrecoverable, and nothing later in this run can reveal it.
	if warning := gapWarning(state, cfg, now); warning != "" {
		log.Print(warning)
		if warningPath != "" {
			if err := os.WriteFile(warningPath, []byte(warning+"\n"), 0o644); err != nil {
				return fmt.Errorf("write warning: %w", err)
			}
		}
	}

	fetcher := source.NewFetcher(cfg)
	var (
		incidents []model.Incident
		statuses  []model.SourceStatus
		failed    int
		seen      int
	)

	for _, sc := range cfg.Sources {
		src, err := source.New(sc)
		if err != nil {
			return fmt.Errorf("source %q: %w", sc.Vendor, err)
		}
		st := model.SourceStatus{Vendor: sc.Vendor, FetchedAt: now}

		res, err := src.Fetch(ctx, fetcher)
		switch {
		case err != nil:
			// One dead vendor is a line in the digest, not a failed run. The
			// threshold below decides when it stops being tolerable.
			st.OK = false
			st.Error = err.Error()
			failed++
			log.Printf("%-13s FAILED: %v", sc.Vendor, err)
		default:
			st.OK = true
			st.IncidentsSeen = len(res.Incidents)
			st.ParseErrors = res.ParseErrors
			incidents = append(incidents, res.Incidents...)
			seen += len(res.Incidents)
			state.Sources[sc.Vendor] = now
			// A healthy source says nothing. Thirteen "0 parse errors" lines
			// every day are noise that trains you to skip the log, and the
			// per-source counts are kept in index.json anyway. Only the
			// notable cases speak up.
			if res.ParseErrors > 0 {
				log.Printf("%-13s %d of %d records unparseable and dropped",
					sc.Vendor, res.ParseErrors, len(res.Incidents)+res.ParseErrors)
			}
		}
		statuses = append(statuses, st)

		if ctx.Err() != nil {
			return ctx.Err()
		}
	}

	stats, err := arc.Upsert(incidents, now)
	if err != nil {
		return err
	}
	if err := writeIndexes(arc, stats.Weeks, statuses, cfg, now); err != nil {
		return err
	}

	state.LastCollectAt = now
	if err := arc.SaveState(state); err != nil {
		return err
	}

	summary := fmt.Sprintf("collect: %d new, %d updated, %d seen, %d/%d sources ok%s",
		stats.Created, stats.Updated, seen, len(cfg.Sources)-failed, len(cfg.Sources), weekSuffix(stats.Weeks))
	log.Print(summary)
	if summaryPath != "" {
		if err := os.WriteFile(summaryPath, []byte(summary+"\n"), 0o644); err != nil {
			return fmt.Errorf("write summary: %w", err)
		}
	}

	// Losing more than half the sources points at this side of the wire, not at
	// nine vendors breaking at once.
	if share := float64(failed) / float64(len(cfg.Sources)); share > cfg.Collect.FailThreshold {
		return fmt.Errorf("%d of %d sources failed, over the %.0f%% threshold",
			failed, len(cfg.Sources), cfg.Collect.FailThreshold*100)
	}
	return nil
}

// writeIndexes rebuilds index.json for every week this run changed. A late
// postmortem lands in the week its incident began, so past weeks get rewritten
// too. Weeks the run did not change keep the index they already have, which is
// what stops a daily run from touching every week the feeds still reach.
func writeIndexes(arc *archive.Archive, weeks []archive.Week, statuses []model.SourceStatus, cfg *config.Config, now time.Time) error {
	for _, w := range weeks {
		incidents, err := arc.LoadWeek(w)
		if err != nil {
			return err
		}
		if _, err := arc.WriteIndex(w, score.BuildIndex(w, incidents, statuses, cfg, now)); err != nil {
			return err
		}
	}
	return nil
}

// gapWarning describes a hole the archive cannot fill, or returns empty when
// the run is on schedule. Nothing downstream can detect this after the fact:
// the missing incidents simply never appear.
func gapWarning(state archive.State, cfg *config.Config, now time.Time) string {
	if state.LastCollectAt.IsZero() {
		return ""
	}
	gap := now.Sub(state.LastCollectAt)
	if gap <= time.Duration(cfg.Collect.GapWarnDays)*24*time.Hour {
		return ""
	}
	return fmt.Sprintf("WARNING: %.0f days since the last collect (%s). Feeds hold about 50 "+
		"incidents and have no pagination, so anything that aged out in that window is gone for good.",
		gap.Hours()/24, state.LastCollectAt.Format(time.RFC3339))
}

// weekSuffix names the weeks the run changed. The first collect against empty
// storage backfills whatever the feeds still hold, which reached 72 weeks in
// practice, so long lists collapse to a range rather than filling the commit
// message.
func weekSuffix(weeks []archive.Week) string {
	if len(weeks) == 0 {
		return ""
	}
	ids := make([]string, len(weeks))
	for i, w := range weeks {
		ids[i] = w.String()
	}
	sort.Strings(ids)
	if len(ids) > 4 {
		return fmt.Sprintf(" (%d weeks, %s..%s)", len(ids), ids[0], ids[len(ids)-1])
	}
	return " (" + strings.Join(ids, ", ") + ")"
}
