# Week 2025-W50 — top 2 of 2 above threshold

2025-12-08 to 2025-12-14 (UTC).

3 source records (3 saved revisions) → 3 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## 1. datadog — Delayed APM metric ingestion

- URL: https://stspg.io/qwpxql6b48pl
- Window: 2025-12-12 18:37 to 2025-12-12 21:52 UTC (3h 15m)
- Impact: major, status: resolved
- Components: Metrics and Infra Monitoring

<a id="ci-d42c11c45abe62d91a8e93fb"></a>

- Canonical ID: `ci-d42c11c45abe62d91a8e93fb` (1 source records).
  - [datadog/c4gh438m6pp3](../../data/2025/W50/datadog.json); vendor URL: https://stspg.io/qwpxql6b48pl
- Start time basis: `vendor_reported`.

  - `2025-12-12 18:37` investigating — We are currently investigating lag in ingesting apm metrics, which affects monitor evaluation.
  - `2025-12-12 18:55` investigating — We are continuing to investigate this issue.
  - `2025-12-12 19:26` investigating — We are currently investigating lag in ingesting apm and process metrics, which affects monitor evaluation and in some cases led to incorrect monitor alerts.
  - `2025-12-12 20:53` identified — We have identified the issue affecting ingestion delays in apm and process metrics and are working on recovery
  - `2025-12-12 21:52` resolved — All impact related to APM metrics has been resolved. A separate incident has been created to track the remaining impact in live process data.

## 2. datadog — Delayed Processes data

- URL: https://stspg.io/ppc0471xzd9w
- Window: 2025-12-12 21:49 to 2025-12-12 23:43 UTC (1h 54m)
- Impact: major, status: resolved
- Components: Metrics and Infra Monitoring

<a id="ci-afca22236f96448dec34a02e"></a>

- Canonical ID: `ci-afca22236f96448dec34a02e` (1 source records).
  - [datadog/rxrmt186707r](../../data/2025/W50/datadog.json); vendor URL: https://stspg.io/ppc0471xzd9w
- Start time basis: `vendor_reported`.

  - `2025-12-12 21:49` investigating — We are investigating increased latency processing Processes data. As a result of this issue, some users may see delays or gaps for data based on Process Monitoring. To prevent spurious alerts, we have temporarily disabled monitors based on this data.
  - `2025-12-12 21:53` investigating — We are continuing to investigate this issue.
  - `2025-12-12 22:41` identified — The issue has been identified and a fix is being implemented.
  - `2025-12-12 23:31` monitoring — A fix has been implemented and we are monitoring the results.
  - `2025-12-12 23:43` resolved — This incident has been resolved.
