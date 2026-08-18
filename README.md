# weekly-incident

[![ci](https://github.com/jtprogru/weekly-incident/actions/workflows/ci.yml/badge.svg)](https://github.com/jtprogru/weekly-incident/actions/workflows/ci.yml)
[![collect](https://github.com/jtprogru/weekly-incident/actions/workflows/collect.yml/badge.svg)](https://github.com/jtprogru/weekly-incident/actions/workflows/collect.yml)
[![digest](https://github.com/jtprogru/weekly-incident/actions/workflows/digest.yml/badge.svg)](https://github.com/jtprogru/weekly-incident/actions/workflows/digest.yml)
[![Go](https://img.shields.io/badge/go-1.26-00ADD8?logo=go&logoColor=white)](go.mod)
[![License: MIT](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

Collects incidents from thirteen vendor status pages into a versioned JSON archive, and renders one Markdown fact sheet per ISO week.

It exists to feed a weekly write-up. The code gathers and ranks; a human decides what the story of the week is and writes it. There is no LLM anywhere in the pipeline, no network access in the render step, and no API key needed to run any of it.

## What it produces

```
data/2026/W34/github.json      one file per vendor per ISO week, upsert-only
data/2026/W34/index.json       scored summary of the week, fully derived
data/state.json                last successful collect, per source
weeks/2026-W34/digest.md       the fact sheet: scored table, timelines, source health
weeks/2026-W34/brief.md        top incidents with update bodies truncated
weeks/2026-W34/summary.txt     the Telegram notification, already in HTML
```

A rendered week is in [`weeks/2026-W33/digest.md`](weeks/2026-W33/digest.md); the same week as a Telegram message is [`summary.txt`](weeks/2026-W33/summary.txt).

Every incident keeps the vendor's original payload in its `raw` field. A normalization mistake found six months from now can be repaired from the archive without asking a single vendor for anything.

The first run backfilled everything the feeds still held: 511 incidents across 72 ISO weeks, the oldest dated 2019. Everything since then arrives one daily run at a time.

## Sources

Ten vendors serve the Atlassian Statuspage v2 schema and share one parser: GitHub, Cloudflare, Datadog, Atlassian, Stripe, OpenAI, Anthropic, Grafana, CircleCI, DigitalOcean. Three need their own: Slack, GCP, AWS.

Adding a vendor on the Statuspage schema is one entry in `config.yaml`. Adding one on a new schema is a file in `internal/source` and a case in the dispatcher.

A few endpoints are not where you would guess. `status.stripe.com` answers 404; the live host is `www.stripestatus.com`. `status.anthropic.com` redirects to `status.claude.com`. AWS still serves `status.aws.amazon.com/rss/all.rss` despite the documentation pointing at a replacement that only reports current state.

## Running it

Go 1.26 and nothing else. No database, no credentials, no services to start.

```bash
make collect              # fetch every source, merge into data/
make digest               # render the week that just closed
make digest WEEK=2026-W34 # render a specific week
make test lint            # go test -race; gofmt, go vet, golangci-lint
make cover build fmt      # the rest of the targets; make help lists them
```

`collect` is the only thing that touches the network. `digest` reads the archive alone, so a week can be re-rendered with a different threshold or a different score without re-fetching anything.

Both accept `-config`, `-data` and `-now` (RFC 3339); pinning the clock is what makes a run reproducible, and the tests rely on it. `digest` also takes `-week` and `-out`, and prints the rendered week to stdout with everything else on stderr, which is how the workflow learns which week it just built. `collect` takes `-summary` for the commit message and `-warning` for the archive-gap alert, both written only when there is something to say.

Parser tests run offline against feed responses captured in `testdata/`. When a vendor changes its schema and a test fails, re-capture with `make testdata` instead of editing the golden by hand.

## Scheduling

`collect.yml` runs daily at 06:00 UTC. Daily rather than weekly, because vendor feeds hold a fixed number of incidents and Statuspage has no pagination: `?page=2` returns the same bytes. Fifty incidents is 62 days of history for GitHub but only 20 for Cloudflare, so a weekly cadence would leave two missed runs of slack before data is lost permanently. The daily commit doubles as the repository activity that keeps GitHub from disabling the schedule after 60 idle days.

`digest.yml` runs Mondays at 07:00 UTC and covers the week that just closed. It also accepts a week by hand through `workflow_dispatch`.

Both write to `main` and share a `concurrency: archive` group; if a push still loses a race, the job resyncs onto the new tip and re-runs itself. Collect merges upsert-only and digest reads only the archive, so the second pass lands exactly what the first one would have.

`ci.yml` gates every push and pull request on `make lint` and `make test`, so the gate is `gofmt -l`, `go vet`, golangci-lint and `go test -race` — the same four checks a developer runs locally, from the same targets. The scheduled workflows call `make collect` and `make digest` for the same reason: the only thing CI adds is the git plumbing around them.

All three need `TELEGRAM_BOT_TOKEN` and `TELEGRAM_CHAT_ID` as repository secrets for [notiflow](https://github.com/jtprogru/notiflow) notifications: the week's summary on Monday, a failure alert from either job, and a separate alert when `collect` finds the archive going stale. That last one matters more than it looks — a gap in the archive is silent damage, and the incidents that aged out of the feeds are gone for good.

## The score

`duration_minutes × max(1, component_count) × vendor_weight`.

Every multiplier is printed next to the row it produced, so the ordering can be explained without reading the code. An unresolved incident is measured against render time and marked `ongoing`.

The score sorts and suggests; it does not decide. Picking the incident of the week is an editorial call.

An incident earns a full entry in the digest if it ran at least `min_duration_minutes` **or** its impact is at least `min_impact`. The disjunction is deliberate: a four-minute critical outage says more than a two-day minor degradation. Everything else collapses into a one-line tail, and nothing is ever dropped from `data/`.

## Known limitations

**Azure is absent.** It has no machine-readable history. Its RSS feed reports only active events and was empty when measured.

**Scheduled maintenances are not collected.** Different endpoint, different semantics, and almost never material for a write-up. A maintenance that turns into an outage gets a real incident from the vendor anyway.

**History before the first run is unrecoverable.** The backfill got whatever the feeds happened to hold that day, which is deep for GitHub and shallow for Cloudflare. Feed depth is finite and unpaginated, so a long gap between runs leaves a permanent hole; that is why `collect` warns when the archive goes stale.

**AWS and Slack use undocumented endpoints.** They can change shape or disappear without notice.

**AWS incidents have no stable identifier.** The key is derived from the first update visible in the feed, so a chain sitting on the boundary of the 50-item window can be re-keyed and land in the archive twice. Accepted knowingly: a duplicate is cheaper than complicating the merge for a feed that may vanish anyway.

**AWS start times are approximate.** The archive records the first update that reached the feed. The true onset appears only in prose inside the final summary ("Starting August 14 7:33 PM PDT"), which is not worth parsing with a regex.

**Slack and AWS report no severity.** Their incidents carry `impact: unknown` and can only clear the digest threshold on duration.

**GCP publishes sparsely.** Four incidents in the half-year before this was written, but each arrives with a full postmortem.

## Data and licensing

The code is MIT licensed. The incident data is collected from public vendor status pages; the wording of those updates belongs to the vendors who wrote it. What gets published elsewhere are the facts — who, what, when, how long, which components — with a link to the source. Direct quotes stay short and in quotation marks.

The crawler sends a `User-Agent` naming this repository, and makes thirteen requests a day.

`data/` and `weeks/` are written by the scheduled workflows and committed by them. Do not hand-edit either: change the code or `config.yaml` and re-render.

## Layout

```
cmd/collect      fetch and merge
cmd/digest       render a week from the archive
internal/model   canonical Incident, Update, Score, WeekIndex
internal/config  config.yaml
internal/source  statuspage.go, slack.go, gcp.go, awsrss.go
internal/archive ISO weeks, upsert, state
internal/clock   -now, so a run can be pinned to a moment
internal/score   ranking and the week index
internal/render  digest.md, brief.md, the Telegram notification
testdata         captured feed responses; refresh with make testdata
scripts          fetch-testdata.sh
.github/workflows ci, collect, digest
PLAN.md          the design, including the reasoning behind each decision
```
