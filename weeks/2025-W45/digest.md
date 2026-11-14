# Week 2025-W45

2025-11-03 to 2025-11-09 (UTC). Generated 2026-11-14 00:00.

1 canonical incidents, 1 above threshold. All 13 sources responded.

1 source records (1 saved revisions) → 1 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | datadog | [Delayed Metrics for APM and distribution metrics](https://stspg.io/f3zkd8tmhdxm) | 2025-11-05 15:42 | 2h 14m | minor | APM, Metrics and Infra Monitoring, Monitors | 402 = 134 × 3 × 1.0 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. datadog — Delayed Metrics for APM and distribution metrics

- URL: https://stspg.io/f3zkd8tmhdxm
- Started: 2025-11-05 15:42 UTC
- Resolved: 2025-11-05 17:57 UTC
- Duration: 2h 14m
- Impact: minor, status: resolved
- Components: APM, Metrics and Infra Monitoring, Monitors
- Canonical incident: [ci-0dbd002ac57f9e8f98271d38](#ci-0dbd002ac57f9e8f98271d38)

- `2025-11-05 15:42` **identified** — We are investigating increased latency processing Metrics from APM and distribution metrics. As a result of this issue, some users may see delays or gaps for metrics on graphs. To prevent spurious alerts, we have temporarily disabled monitors based on this data.
- `2025-11-05 16:25` **identified** — We have identified the root cause, and scaled up the processing to catch up with the lag.
- `2025-11-05 16:59` **identified** — The distribution metrics should be back to processing live, without latency. APM metrics are still being delayed, and we're actively working at getting it back to live.
- `2025-11-05 17:57` **resolved** — APM metrics are now processing live.

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-0dbd002ac57f9e8f98271d38"></a>

- Canonical ID: `ci-0dbd002ac57f9e8f98271d38` (1 source records).
  - [datadog/tw86ncfv8q78](../../data/2025/W45/datadog.json); vendor URL: https://stspg.io/f3zkd8tmhdxm
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
| gcp | ok | 6 | 0 |
| github | ok | 50 | 0 |
| grafana | ok | 50 | 0 |
| openai | ok | 25 | 0 |
| slack | ok | 36 | 0 |
| stripe | ok | 50 | 0 |
