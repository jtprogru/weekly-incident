# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this is

A Go pipeline that collects incidents from thirteen vendor status pages into a git-versioned JSON archive (`data/`) and renders one Markdown fact sheet per ISO week (`weeks/`). No LLM, no database, no API keys. `README.md` describes the product; `PLAN.md` (Russian) holds the design and the reasoning behind every decision, including section 16 on what changed once the code met the live feeds.

## Commands

```bash
make test                       # go test -race ./...
make lint                       # go vet + golangci-lint (CI also enforces gofmt -l)
make cover                      # coverage summary
make collect                    # fetch all sources, merge into data/  (only network step)
make digest                     # render the week that just closed
make digest WEEK=2026-W34       # render a specific week
make testdata                   # re-capture golden feed responses from live endpoints
```

Single test: `go test ./internal/source -run TestAWSParse -v`. Package: `go test ./internal/archive/...`.

Both binaries take `-now` (RFC 3339) to pin the clock, plus `-config`, `-data`, and `-out`/`-summary`. Pinning the clock is what makes a run reproducible, and tests rely on it.

## Architecture

The pipeline is a straight line, and the two ends are deliberately disconnected:

`config.yaml` → `internal/source` (fetch + normalize) → `internal/archive` (upsert into `data/YYYY/Wnn/<vendor>.json`) → `internal/score` (rank, build `index.json`) → `internal/render` (`digest.md`, `brief.md`, `summary.txt`). `internal/clock` resolves `-now` for both commands.

`cmd/collect` is the only code that touches the network. `cmd/digest` never fetches, so any past week can be re-rendered with a different threshold or score. Keep it that way. `cmd/digest` prints the rendered week to **stdout** and everything else to stderr — the workflow reads the week from there, so do not add stdout chatter.

`internal/model` is the canonical vocabulary every other package speaks: `Incident`, `Update`, `Impact`, `Status`, `Score`, `IndexEntry`, `WeekIndex`. Parsers build `Incident` values, call `Finalize()` (derives `Key`, normalizes times to UTC, computes `DurationMinutes`, canonicalizes `Raw`), then `Validate()`, and skip the record on error rather than failing the whole feed.

## Invariants that are easy to break

**The archive is upsert-only.** Statuspage feeds hold ~50 incidents with no pagination (`?page=2` returns the same bytes). An incident missing from a feed fell out of the window; it was not retracted. Never delete on absence.

**Nothing may change on a re-run of the same data.** `archive.merge` compares the *serialized bytes* of old and new records and skips the write when they match; `Raw` is canonicalized (sorted keys, no whitespace) so a vendor reordering its JSON does not look like news; every slice is sorted with a total order. There is intentionally no "last seen" field — it would rewrite the whole archive daily and drown the git history.

The same rule covers the derived files, which carry a timestamp and would otherwise churn forever: `archive.WriteIndex` compares content while *ignoring* `GeneratedAt`, leaves the file alone when only the clock moved, and returns what is actually on disk so the digest quotes that timestamp rather than a fresh one; `cmd/digest.writeIfChanged` does the same for the three Markdown files. `TestUpsertIsIdempotent`, `TestUpsertSkipsUnchangedWeeks`, `TestWriteIndexIgnoresTheClock` and `TestAWSParseIsDeterministic` guard this.

**Incidents are routed by `StartedAt`**, so a run writes into past weeks (a late postmortem updates the week the incident began in). `collect` rebuilds `index.json` for every week it touched.

**`index.json` is fully derived** from the vendor files and rebuilt from scratch. `data/` and `weeks/` are written by the scheduled workflows and committed by them — do not hand-edit either; change the code or `config.yaml` and re-render.

**`ImpactUnknown` ranks below `none`.** Slack and AWS report no severity; treating them as severe would flood the digest. They clear the threshold on duration only.

**Threshold is a disjunction:** `duration >= min_duration_minutes` OR `impact >= min_impact`. A four-minute critical outage beats a two-day minor degradation. Nothing is ever dropped from `data/` — sub-threshold incidents collapse into a one-line tail in the digest.

## Adding a vendor

On the Statuspage v2 schema: one entry in `config.yaml` (`vendor`, `kind: statuspage`, `url`, `weight`). Nothing else.

On a new schema: a file in `internal/source`, a `Kind` constant in `internal/config`, a case in `source.New`, the kind added to `Validate`, a golden capture in `testdata/`, and a parse test. `weight` is editorial, not technical — how much the audience cares, not how reliable the vendor is.

## Tests and goldens

Parser tests run offline against captured feed responses in `testdata/`, one per source kind. When a vendor changes its schema and a test fails, re-capture with `make testdata` rather than editing the golden by hand — the point is that the tests describe the feeds as they really are. `scripts/fetch-testdata.sh` trims the captures, except the AWS feed, which stays whole because its grouping logic needs several chains.

## Endpoint and parser gotchas

`status.stripe.com` 404s; the live host is `www.stripestatus.com`. `status.anthropic.com` redirects to `status.claude.com`, which is configured directly. AWS still serves the undocumented `rss/all.rss`; the documented replacement is a current-state snapshot with no history.

The AWS feed emits *updates*, not incidents. The epoch in its `guid` is the item's own `pubDate`, not the event start, so chains are grouped by service key and cut at a `[RESOLVED]` headline or a gap wider than 24h. Its key is derived from the first visible update, which means a chain straddling the 50-item boundary can be re-keyed and duplicated — accepted knowingly.
