# Week 2025-W42

2025-10-13 to 2025-10-19 (UTC). Generated 2026-11-14 00:00.

1 canonical incidents, 1 above threshold. All 13 sources responded.

1 source records (1 saved revisions) → 1 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | datadog | [Delayed AWS, GCP, Azure, SaaS integrations Metrics and Logs](https://stspg.io/x812ktyhkf0q) | 2025-10-14 17:14 | 2h 4m | critical | Metrics and Infra Monitoring | 124 = 124 × 1 × 1.0 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. datadog — Delayed AWS, GCP, Azure, SaaS integrations Metrics and Logs

- URL: https://stspg.io/x812ktyhkf0q
- Started: 2025-10-14 17:14 UTC
- Resolved: 2025-10-14 19:19 UTC
- Duration: 2h 4m
- Impact: critical, status: resolved
- Components: Metrics and Infra Monitoring
- Canonical incident: [ci-73a428fa1e5d28803e4ea9f3](#ci-73a428fa1e5d28803e4ea9f3)

- `2025-10-14 17:14` **investigating** — We are investigating increased latency processing AWS, GCP and Azure Metrics. As a result of this issue, some users may see delays or gaps in graphs that contain these metrics. To prevent spurious alerts, we have temporarily disabled monitors based on this data.
- `2025-10-14 17:29` **identified** — The issue has been identified and a fix is being implemented.
- `2025-10-14 17:43` **monitoring** — A fix has been implemented, and we are monitoring the results.
- `2025-10-14 18:44` **monitoring** — Data flow has been restored for new incoming data. We are currently backfilling historical data.
- `2025-10-14 19:19` **resolved** — This incident has been resolved.

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-73a428fa1e5d28803e4ea9f3"></a>

- Canonical ID: `ci-73a428fa1e5d28803e4ea9f3` (1 source records).
  - [datadog/twnwtk275knx](../../data/2025/W42/datadog.json); vendor URL: https://stspg.io/x812ktyhkf0q
- Start time basis: `vendor_reported`.

## Sources

| Vendor | Status | Source records seen | Parse errors |
|--------|--------|----------------|--------------|
| anthropic | ok | 50 | 0 |
| atlassian | ok | 34 | 0 |
| aws | ok | 5 | 0 |
| circleci | ok | 50 | 0 |
| cloudflare | ok | 50 | 0 |
| datadog | ok | 50 | 0 |
| digitalocean | ok | 50 | 0 |
| gcp | ok | 4 | 0 |
| github | ok | 50 | 0 |
| grafana | ok | 50 | 0 |
| openai | ok | 25 | 0 |
| slack | ok | 43 | 0 |
| stripe | ok | 50 | 0 |
