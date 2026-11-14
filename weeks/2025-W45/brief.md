# Week 2025-W45 — top 1 of 1 above threshold

2025-11-03 to 2025-11-09 (UTC).

1 source records (1 saved revisions) → 1 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## 1. datadog — Delayed Metrics for APM and distribution metrics

- URL: https://stspg.io/f3zkd8tmhdxm
- Window: 2025-11-05 15:42 to 2025-11-05 17:57 UTC (2h 14m)
- Impact: minor, status: resolved
- Components: APM, Metrics and Infra Monitoring, Monitors

<a id="ci-0dbd002ac57f9e8f98271d38"></a>

- Canonical ID: `ci-0dbd002ac57f9e8f98271d38` (1 source records).
  - [datadog/tw86ncfv8q78](../../data/2025/W45/datadog.json); vendor URL: https://stspg.io/f3zkd8tmhdxm
- Start time basis: `vendor_reported`.

  - `2025-11-05 15:42` identified — We are investigating increased latency processing Metrics from APM and distribution metrics. As a result of this issue, some users may see delays or gaps for metrics on graphs. To prevent spurious alerts, we have temporarily disabled monitors based on this data.
  - `2025-11-05 16:25` identified — We have identified the root cause, and scaled up the processing to catch up with the lag.
  - `2025-11-05 16:59` identified — The distribution metrics should be back to processing live, without latency. APM metrics are still being delayed, and we're actively working at getting it back to live.
  - `2025-11-05 17:57` resolved — APM metrics are now processing live.
