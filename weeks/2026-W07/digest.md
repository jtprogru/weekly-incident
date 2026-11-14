# Week 2026-W07

2026-02-09 to 2026-02-15 (UTC). Generated 2026-11-14 00:00.

3 canonical incidents, 1 above threshold. All 14 sources responded.

3 source records (3 saved revisions) → 3 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | datadog | [Delayed Traces in APM Trace Search](https://stspg.io/4znpw24vxbbg) | 2026-02-11 20:56 | 2h 48m | minor | APM, Monitors | 336 = 168 × 2 × 1.0 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. datadog — Delayed Traces in APM Trace Search

- URL: https://stspg.io/4znpw24vxbbg
- Started: 2026-02-11 20:56 UTC
- Resolved: 2026-02-11 23:44 UTC
- Duration: 2h 48m
- Impact: minor, status: resolved
- Components: APM, Monitors
- Canonical incident: [ci-ed94679a568c57213addd81c](#ci-ed94679a568c57213addd81c)

- `2026-02-11 20:56` **investigating** — We are investigating increased latency processing Traces in APM Trace Search. As a result of this issue, some users may see missing or delayed traces in APM Trace Search and delayed APM Trace Monitors since 21:30 UTC.
- `2026-02-11 21:18` **identified** — Teams have identified the issue and are working to mitigate at this time
- `2026-02-11 21:51` **identified** — Teams continue to work to mitigate the impact of this issue. At this time, APM Trace Processing, APM Trace Monitors are still delayed and Distribution Metrics monitors are delayed starting from 21:35 UTC
- `2026-02-11 22:28` **identified** — Teams continue to work to mitigate the impact of this issue. At this time, APM Trace Processing, APM Trace Monitors are still delayed. Distribution Metrics monitors were delayed between 21:35 and 22:15 UTC
- `2026-02-11 23:28` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-02-11 23:44` **resolved** — This incident has been resolved.

## Below threshold (2)

- `datadog` [Delayed Traces in APM](https://stspg.io/19l6tqc8m1b2) — 1h 10m
- `stripe` [Elevated latency for Visa EU transactions](https://stspg.io/y5frtt77svhm) — 1h 15m

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-ed94679a568c57213addd81c"></a>

- Canonical ID: `ci-ed94679a568c57213addd81c` (1 source records).
  - [datadog/g5vw545yn7vt](../../data/2026/W07/datadog.json); vendor URL: https://stspg.io/4znpw24vxbbg
- Start time basis: `vendor_reported`.

<a id="ci-cc58b52f1ce45d0c6577ce8e"></a>

- Canonical ID: `ci-cc58b52f1ce45d0c6577ce8e` (1 source records).
  - [datadog/vkqm21bfss9c](../../data/2026/W07/datadog.json); vendor URL: https://stspg.io/19l6tqc8m1b2
- Start time basis: `vendor_reported`.

<a id="ci-ed399702064c1ec5324eb991"></a>

- Canonical ID: `ci-ed399702064c1ec5324eb991` (1 source records).
  - [stripe/dtd84fm0pj99](../../data/2026/W07/stripe.json); vendor URL: https://stspg.io/y5frtt77svhm
- Start time basis: `vendor_reported`.

## Sources

| Vendor | Status | Source records seen | Parse errors |
|--------|--------|----------------|--------------|
| anthropic | ok | 50 | 0 |
| atlassian | ok | 34 | 0 |
| aws | ok | 6 | 0 |
| circleci | ok | 50 | 0 |
| cloudflare | ok | 50 | 0 |
| datadog | ok | 50 | 0 |
| digitalocean | ok | 50 | 0 |
| gcp | ok | 8 | 0 |
| github | ok | 50 | 0 |
| grafana | ok | 50 | 0 |
| openai | ok | 25 | 0 |
| slack | ok | 31 | 0 |
| stripe | ok | 50 | 0 |
| yandexcloud | ok | 50 | 0 |
