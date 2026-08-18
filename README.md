# weekly-incident

Collects incidents from thirteen vendor status pages into a versioned JSON archive, and renders one Markdown fact sheet per ISO week.

It exists to feed a weekly write-up. The code gathers and ranks; a human decides what the story of the week is and writes it. There is no LLM anywhere in the pipeline, no network access in the render step, and no API key needed to run any of it.

## What it produces

```
data/2026/W34/github.json      one file per vendor per ISO week, upsert-only
data/2026/W34/index.json       scored summary of the week, fully derived
data/state.json                last successful collect, per source
weeks/2026-W34/digest.md       the fact sheet: scored table, timelines, source health
weeks/2026-W34/brief.md        top incidents with update bodies truncated
weeks/2026-W34/summary.txt     three lines, forwarded to Telegram by the workflow
```

Every incident keeps the vendor's original payload in its `raw` field. A normalization mistake found six months from now can be repaired from the archive without asking a single vendor for anything.

## Sources

Ten vendors serve the Atlassian Statuspage v2 schema and share one parser: GitHub, Cloudflare, Datadog, Atlassian, Stripe, OpenAI, Anthropic, Grafana, CircleCI, DigitalOcean. Three need their own: Slack, GCP, AWS.

Adding a vendor on the Statuspage schema is one entry in `config.yaml`. Adding one on a new schema is a file in `internal/source` and a case in the dispatcher.

A few endpoints are not where you would guess. `status.stripe.com` answers 404; the live host is `www.stripestatus.com`. `status.anthropic.com` redirects to `status.claude.com`. AWS still serves `status.aws.amazon.com/rss/all.rss` despite the documentation pointing at a replacement that only reports current state.

## Running it

```bash
make collect              # fetch every source, merge into data/
make digest               # render the week that just closed
make digest WEEK=2026-W34 # render a specific week
make test lint
```

`collect` is the only thing that touches the network. `digest` reads the archive alone, so a week can be re-rendered with a different threshold or a different score without re-fetching anything.

Both accept `-now` (RFC 3339) to pin the clock, which is what makes a run reproducible.

## Scheduling

`collect.yml` runs daily at 06:00 UTC. Daily rather than weekly, because vendor feeds hold a fixed number of incidents and Statuspage has no pagination: `?page=2` returns the same bytes. Fifty incidents is 62 days of history for GitHub but only 20 for Cloudflare, so a weekly cadence would leave two missed runs of slack before data is lost permanently. The daily commit doubles as the repository activity that keeps GitHub from disabling the schedule after 60 idle days.

`digest.yml` runs Mondays at 07:00 UTC and covers the week that just closed.

Both need `TELEGRAM_BOT_TOKEN` and `TELEGRAM_CHAT_ID` as repository secrets for [notiflow](https://github.com/jtprogru/notiflow) notifications.

## The score

`duration_minutes × max(1, component_count) × vendor_weight`.

Every multiplier is printed next to the row it produced, so the ordering can be explained without reading the code. An unresolved incident is measured against render time and marked `ongoing`.

The score sorts and suggests; it does not decide. Picking the incident of the week is an editorial call.

An incident earns a full entry in the digest if it ran at least `min_duration_minutes` **or** its impact is at least `min_impact`. The disjunction is deliberate: a four-minute critical outage says more than a two-day minor degradation. Everything else collapses into a one-line tail, and nothing is ever dropped from `data/`.

## Known limitations

**Azure is absent.** It has no machine-readable history. Its RSS feed reports only active events and was empty when measured.

**Scheduled maintenances are not collected.** Different endpoint, different semantics, and almost never material for a write-up. A maintenance that turns into an outage gets a real incident from the vendor anyway.

**History before the first run is unrecoverable.** Feed depth is finite and unpaginated. A long gap between runs leaves a permanent hole, which is why `collect` warns when the archive goes stale.

**AWS and Slack use undocumented endpoints.** They can change shape or disappear without notice.

**AWS incidents have no stable identifier.** The key is derived from the first update visible in the feed, so a chain sitting on the boundary of the 50-item window can be re-keyed and land in the archive twice. Accepted knowingly: a duplicate is cheaper than complicating the merge for a feed that may vanish anyway.

**AWS start times are approximate.** The archive records the first update that reached the feed. The true onset appears only in prose inside the final summary ("Starting August 14 7:33 PM PDT"), which is not worth parsing with a regex.

**Slack and AWS report no severity.** Their incidents carry `impact: unknown` and can only clear the digest threshold on duration.

**GCP publishes sparsely.** Four incidents in the half-year before this was written, but each arrives with a full postmortem.

## Data and licensing

The code is MIT licensed. The incident data is collected from public vendor status pages; the wording of those updates belongs to the vendors who wrote it. What gets published elsewhere are the facts — who, what, when, how long, which components — with a link to the source. Direct quotes stay short and in quotation marks.

The crawler sends a `User-Agent` naming this repository, and makes thirteen requests a day.

## Layout

```
cmd/collect      fetch and merge
cmd/digest       render a week from the archive
internal/model   canonical Incident, Update, Score, WeekIndex
internal/config  config.yaml
internal/source  statuspage.go, slack.go, gcp.go, awsrss.go
internal/archive ISO weeks, upsert, state
internal/score   ranking and the week index
internal/render  digest.md, brief.md
testdata         captured feed responses; refresh with make testdata
PLAN.md          the design, including the reasoning behind each decision
```
