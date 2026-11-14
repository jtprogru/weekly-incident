# Week 2026-W06

2026-02-02 to 2026-02-08 (UTC). Generated 2026-11-14 00:00.

2 canonical incidents, 1 above threshold. All 14 sources responded.

2 source records (2 saved revisions) → 2 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | datadog | [Delays in Monitor Evaluations](https://stspg.io/kzxxld69ftmr) | 2026-02-05 15:24 | 2h 18m | major | APM, Application Security Management, Cloud SIEM, Incident Response, Log Management, Metrics and Infra Monitoring, Monitors, RUM, Synthetics | 1242 = 138 × 9 × 1.0 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. datadog — Delays in Monitor Evaluations

- URL: https://stspg.io/kzxxld69ftmr
- Started: 2026-02-05 15:24 UTC
- Resolved: 2026-02-05 17:43 UTC
- Duration: 2h 18m
- Impact: major, status: resolved
- Components: APM, Application Security Management, Cloud SIEM, Incident Response, Log Management, Metrics and Infra Monitoring, Monitors, RUM, Synthetics
- Canonical incident: [ci-57f904136b32c968081e263b](#ci-57f904136b32c968081e263b)

- `2026-02-05 15:24` **investigating** — We are currently investigating this issue.
- `2026-02-05 16:04` **investigating** — We are continuing to investigate this issue.
- `2026-02-05 16:10` **investigating** — We are continuing to investigate this issue.
- `2026-02-05 16:11` **investigating** — We are continuing to investigate this issue.
- `2026-02-05 16:28` **investigating** — We are continuing to investigate this issue.
- `2026-02-05 16:31` **identified** — The issue has been identified and a fix is being implemented.
- `2026-02-05 16:50` **identified** — We are continuing to work on a fix for this issue.
- `2026-02-05 17:01` **identified** — We are observing recovery for the vast majority of monitors and are continuing to work on a full fix for this issue.
- `2026-02-05 17:23` **monitoring** — We have observed full recovery of monitors and will continue to monitor
- `2026-02-05 17:43` **resolved** — This incident has been resolved.

## Below threshold (1)

- `stripe` [Elevated errors on 3DS transactions for Commonwealth Bank of Australia](https://stspg.io/m7f5tchrrryz) — 0m

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-57f904136b32c968081e263b"></a>

- Canonical ID: `ci-57f904136b32c968081e263b` (1 source records).
  - [datadog/g545g04n9xkr](../../data/2026/W06/datadog.json); vendor URL: https://stspg.io/kzxxld69ftmr
- Start time basis: `vendor_reported`.

<a id="ci-0efb91243799e89715b664cd"></a>

- Canonical ID: `ci-0efb91243799e89715b664cd` (1 source records).
  - [stripe/flm1bpq06yxf](../../data/2026/W06/stripe.json); vendor URL: https://stspg.io/m7f5tchrrryz
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
