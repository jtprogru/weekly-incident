# Week 2025-W42 — top 1 of 1 above threshold

2025-10-13 to 2025-10-19 (UTC).

1 source records (1 saved revisions) → 1 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## 1. datadog — Delayed AWS, GCP, Azure, SaaS integrations Metrics and Logs

- URL: https://stspg.io/x812ktyhkf0q
- Window: 2025-10-14 17:14 to 2025-10-14 19:19 UTC (2h 4m)
- Impact: critical, status: resolved
- Components: Metrics and Infra Monitoring

<a id="ci-73a428fa1e5d28803e4ea9f3"></a>

- Canonical ID: `ci-73a428fa1e5d28803e4ea9f3` (1 source records).
  - [datadog/twnwtk275knx](../../data/2025/W42/datadog.json); vendor URL: https://stspg.io/x812ktyhkf0q
- Start time basis: `vendor_reported`.

  - `2025-10-14 17:14` investigating — We are investigating increased latency processing AWS, GCP and Azure Metrics. As a result of this issue, some users may see delays or gaps in graphs that contain these metrics. To prevent spurious alerts, we have temporarily disabled monitors based on this data.
  - `2025-10-14 17:29` identified — The issue has been identified and a fix is being implemented.
  - `2025-10-14 17:43` monitoring — A fix has been implemented, and we are monitoring the results.
  - `2025-10-14 18:44` monitoring — Data flow has been restored for new incoming data. We are currently backfilling historical data.
  - `2025-10-14 19:19` resolved — This incident has been resolved.
