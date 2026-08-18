// Command digest builds the weekly fact sheet from the archive.
//
// It never touches the network. Everything it needs was written by collect, so
// a digest can be regenerated for any past week, with a changed threshold or a
// changed score, without asking a single vendor for anything.
//
// Progress goes to stderr; the only thing on stdout is the ISO week that was
// rendered, so a caller can capture it without guessing from the filesystem.
package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/jtprogru/weekly-incident/internal/archive"
	"github.com/jtprogru/weekly-incident/internal/clock"
	"github.com/jtprogru/weekly-incident/internal/config"
	"github.com/jtprogru/weekly-incident/internal/model"
	"github.com/jtprogru/weekly-incident/internal/render"
	"github.com/jtprogru/weekly-incident/internal/score"
)

// summaryTopN is how many incidents the notification quotes. It is not
// BriefTopN: a Telegram message wants three lines, a pasteable brief wants
// enough context to write from.
const summaryTopN = 3

func main() {
	var (
		configPath = flag.String("config", "config.yaml", "path to config.yaml")
		dataDir    = flag.String("data", "data", "archive root directory")
		outDir     = flag.String("out", "weeks", "directory for the rendered week")
		weekFlag   = flag.String("week", "", "ISO week to render, e.g. 2026-W34 (default: the week before now)")
		nowFlag    = flag.String("now", "", "override the current time, RFC3339")
	)
	flag.Parse()

	log.SetFlags(0)
	week, err := run(*configPath, *dataDir, *outDir, *weekFlag, *nowFlag)
	if err != nil {
		log.Fatalf("digest: %v", err)
	}
	fmt.Println(week)
}

func run(configPath, dataDir, outDir, weekFlag, nowFlag string) (string, error) {
	cfg, err := config.Load(configPath)
	if err != nil {
		return "", err
	}
	now, err := clock.Now(nowFlag)
	if err != nil {
		return "", err
	}

	week, err := resolveWeek(weekFlag, now)
	if err != nil {
		return "", err
	}

	arc := archive.New(dataDir)
	incidents, err := arc.LoadWeek(week)
	if err != nil {
		return "", err
	}
	if len(incidents) == 0 {
		log.Printf("warning: no incidents archived for %s", week)
	}

	// Source health belongs to the collect run that fetched the data, so it is
	// carried over from whatever index that run left behind. WriteIndex returns
	// what is actually on disk: when nothing but the clock moved it keeps the
	// stored file, and the digest then quotes that file's timestamp instead of
	// inventing a fresh one.
	idx, err := arc.WriteIndex(week, score.BuildIndex(week, incidents, previousSources(arc, week), cfg, now))
	if err != nil {
		return "", err
	}

	dir := filepath.Join(outDir, week.String())
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return "", fmt.Errorf("mkdir %s: %w", dir, err)
	}
	for name, body := range map[string]string{
		"digest.md":   render.Digest(idx, incidents),
		"brief.md":    render.Brief(idx, incidents, cfg.Digest.BriefTopN, cfg.Digest.BriefBodyLimit),
		"summary.txt": summary(idx),
	} {
		if err := writeIfChanged(filepath.Join(dir, name), body); err != nil {
			return "", err
		}
	}

	above := 0
	for _, e := range idx.Incidents {
		if e.AboveThreshold {
			above++
		}
	}
	log.Printf("%s: %d incidents, %d above threshold -> %s", week, len(idx.Incidents), above, dir)
	return week.String(), nil
}

// writeIfChanged leaves an identical file untouched, so re-rendering a settled
// week produces no diff for the workflow to commit.
func writeIfChanged(path, body string) error {
	if old, err := os.ReadFile(path); err == nil && string(old) == body {
		return nil
	}
	if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
		return fmt.Errorf("write %s: %w", path, err)
	}
	return nil
}

// summary is the short text the workflow forwards to notiflow.
func summary(idx model.WeekIndex) string {
	var above []model.IndexEntry
	for _, e := range idx.Incidents {
		if e.AboveThreshold {
			above = append(above, e)
		}
	}

	s := fmt.Sprintf("Week %s: %d incidents, %d above threshold.\n", idx.Week, len(idx.Incidents), len(above))
	for i, e := range above {
		if i == summaryTopN {
			break
		}
		s += fmt.Sprintf("%d. %s — %s (%d min)\n", i+1, e.Vendor, e.Title, e.Score.Duration)
	}
	return s
}

// previousSources reuses the source health recorded by the last collect run. A
// missing index just means the week was never collected.
func previousSources(arc *archive.Archive, w archive.Week) []model.SourceStatus {
	idx, err := arc.ReadIndex(w)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			log.Printf("warning: could not read the existing index for %s: %v", w, err)
		}
		return nil
	}
	return idx.Sources
}

// resolveWeek defaults to the week that just closed: run on Monday morning, the
// digest covers the previous Monday through Sunday in full.
func resolveWeek(s string, now time.Time) (archive.Week, error) {
	if s == "" {
		return archive.WeekOf(now).Prev(), nil
	}
	return archive.ParseWeek(s)
}
