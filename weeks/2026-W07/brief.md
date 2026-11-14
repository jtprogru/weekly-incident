# Week 2026-W07 — top 1 of 1 above threshold

2026-02-09 to 2026-02-15 (UTC).

3 source records (3 saved revisions) → 3 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## 1. datadog — Delayed Traces in APM Trace Search

- URL: https://stspg.io/4znpw24vxbbg
- Window: 2026-02-11 20:56 to 2026-02-11 23:44 UTC (2h 48m)
- Impact: minor, status: resolved
- Components: APM, Monitors

<a id="ci-ed94679a568c57213addd81c"></a>

- Canonical ID: `ci-ed94679a568c57213addd81c` (1 source records).
  - [datadog/g5vw545yn7vt](../../data/2026/W07/datadog.json); vendor URL: https://stspg.io/4znpw24vxbbg
- Start time basis: `vendor_reported`.

  - `2026-02-11 20:56` investigating — We are investigating increased latency processing Traces in APM Trace Search. As a result of this issue, some users may see missing or delayed traces in APM Trace Search and delayed APM Trace Monitors since 21:30 UTC.
  - `2026-02-11 21:18` identified — Teams have identified the issue and are working to mitigate at this time
  - `2026-02-11 21:51` identified — Teams continue to work to mitigate the impact of this issue. At this time, APM Trace Processing, APM Trace Monitors are still delayed and Distribution Metrics monitors are delayed starting from 21:35 UTC
  - `2026-02-11 22:28` identified — Teams continue to work to mitigate the impact of this issue. At this time, APM Trace Processing, APM Trace Monitors are still delayed. Distribution Metrics monitors were delayed between 21:35 and 22:15 UTC
  - `2026-02-11 23:28` monitoring — A fix has been implemented and we are monitoring the results.
  - `2026-02-11 23:44` resolved — This incident has been resolved.
