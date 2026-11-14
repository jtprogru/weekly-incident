# Week 2026-W12

2026-03-16 to 2026-03-22 (UTC). Generated 2026-11-14 00:00.

5 canonical incidents, 2 above threshold. All 14 sources responded.

5 source records (5 saved revisions) → 5 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | stripe | [Elevated Pix errors](https://stspg.io/dzll5mrw1ryz) | 2026-03-22 16:02 | 19h 24m | minor | Acquirers and payment methods | 1397 = 1164 × 1 × 1.2 |
| 2 | datadog | [Intermittent "No Data" Status for Monitors](https://stspg.io/43kkcl1tqrrk) | 2026-03-19 00:54 | 1h 59m | major | Monitors | 119 = 119 × 1 × 1.0 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. stripe — Elevated Pix errors

- URL: https://stspg.io/dzll5mrw1ryz
- Started: 2026-03-22 16:02 UTC
- Resolved: 2026-03-23 11:27 UTC
- Duration: 19h 24m
- Impact: minor, status: resolved
- Components: Acquirers and payment methods
- Canonical incident: [ci-2a67f55dca2afd6cd9d25c1f](#ci-2a67f55dca2afd6cd9d25c1f)

- `2026-03-22 16:02` **identified** — We are currently observing elevated errors on Pix payments. Payments are likely affected due to degraded performance reported by Pix. We will continue to monitor the issue and an update will be provided as soon as more information becomes available.
- `2026-03-23 11:27` **resolved** — Between 14:13 UTC on the 22nd March and 01:40 UTC on the 23rd March, we observed a service disruption that impacted the Pix payment method. The issue was external to Stripe and is now resolved.

### 2. datadog — Intermittent "No Data" Status for Monitors

- URL: https://stspg.io/43kkcl1tqrrk
- Started: 2026-03-19 00:54 UTC
- Resolved: 2026-03-19 02:53 UTC
- Duration: 1h 59m
- Impact: major, status: resolved
- Components: Monitors
- Canonical incident: [ci-14dda82a474ebe0feb60f33d](#ci-14dda82a474ebe0feb60f33d)

- `2026-03-19 00:54` **investigating** — We are investigating an issue causing some metric monitors to intermittently report "No Data" status. This primarily affects monitors based on distribution metrics. Monitor alerting may be unreliable during this time. We are actively investigating and will provide updates as available.
- `2026-03-19 01:27` **investigating** — We are continuing to investigate the issue and will provide updates as available.
- `2026-03-19 02:04` **investigating** — We are continuing to investigate the issue and will provide updates as available.
- `2026-03-19 02:39` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-03-19 02:53` **resolved** — This incident has been resolved.

## Below threshold (3)

- `stripe` [Elevated errors on Klarna payments](https://stspg.io/0ggy0mssf939) — 1h 59m
- `circleci` [Insights API and Dashboard displaying stale data for reporting windows greater than 24 hours](https://stspg.io/hk09l57sybkn) — 50m
- `circleci` [Pipelines — Workflow Creation Errors and Missing GitHub Status Updates](https://stspg.io/n8qyfjwryytr) — unknown

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-2a67f55dca2afd6cd9d25c1f"></a>

- Canonical ID: `ci-2a67f55dca2afd6cd9d25c1f` (1 source records).
  - [stripe/gc6z4c7t1hdq](../../data/2026/W12/stripe.json); vendor URL: https://stspg.io/dzll5mrw1ryz
- Start time basis: `vendor_reported`.

<a id="ci-73545573a9a7172de10c7a09"></a>

- Canonical ID: `ci-73545573a9a7172de10c7a09` (1 source records).
  - [stripe/4y3y5flslh6p](../../data/2026/W12/stripe.json); vendor URL: https://stspg.io/0ggy0mssf939
- Start time basis: `vendor_reported`.

<a id="ci-14dda82a474ebe0feb60f33d"></a>

- Canonical ID: `ci-14dda82a474ebe0feb60f33d` (1 source records).
  - [datadog/j5zpb9phy9m6](../../data/2026/W12/datadog.json); vendor URL: https://stspg.io/43kkcl1tqrrk
- Start time basis: `vendor_reported`.

<a id="ci-3d462e59780bade8ffa4a173"></a>

- Canonical ID: `ci-3d462e59780bade8ffa4a173` (1 source records).
  - [circleci/01zmcr93tysy](../../data/2026/W12/circleci.json); vendor URL: https://stspg.io/hk09l57sybkn
- Start time basis: `vendor_reported`.

<a id="ci-d5531c2b79f49314f02b08ae"></a>

- Canonical ID: `ci-d5531c2b79f49314f02b08ae` (1 source records).
  - [circleci/g0b66jlry2kf](../../data/2026/W12/circleci.json); vendor URL: https://stspg.io/n8qyfjwryytr
- Start time basis: `vendor_reported`.
- Warning: Vendor resolved_at precedes started_at; duration is unknown, score uses zero duration; decision:legacy-time-circleci-g0b66jlry2kf

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
