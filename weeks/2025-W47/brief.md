# Week 2025-W47 — top 2 of 2 above threshold

2025-11-17 to 2025-11-23 (UTC).

3 source records (3 saved revisions) → 3 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## 1. datadog — Dashboards Not Loading

- URL: https://stspg.io/bbdxgwpv272q
- Window: 2025-11-17 16:23 to 2025-11-17 17:20 UTC (56m)
- Impact: major, status: resolved
- Components: Web Application

<a id="ci-54a50f8d3a7c078291140bed"></a>

- Canonical ID: `ci-54a50f8d3a7c078291140bed` (1 source records).
  - [datadog/bg2rxpyn1hc3](../../data/2025/W47/datadog.json); vendor URL: https://stspg.io/bbdxgwpv272q
- Start time basis: `vendor_reported`.

  - `2025-11-17 16:23` investigating — We are investigating loading issues on the dashboard pages. As a result, some users might be getting errors when loading the web application. Please note that data processing and alerts are not affected by this incident.
  - `2025-11-17 16:40` identified — The issue has been identified and we taking measures to mitigate the issue, as well as working on a fix.
  - `2025-11-17 17:12` monitoring — The rollout with a fix is in progress, and we're no longer seeing errors, and are currently monitoring the incident and we are on the path of recovery.
  - `2025-11-17 17:20` resolved — All errors stopped as of 12:02ET. This incident has been resolved.

## 2. datadog — Web Application Not Loading

- URL: https://stspg.io/h30gbqqk6wgl
- Window: 2025-11-19 19:08 to 2025-11-19 19:44 UTC (36m)
- Impact: critical, status: resolved
- Components: Web Application

<a id="ci-8e0d54461822f0521802bb74"></a>

- Canonical ID: `ci-8e0d54461822f0521802bb74` (1 source records).
  - [datadog/cvdjtf81756n](../../data/2025/W47/datadog.json); vendor URL: https://stspg.io/h30gbqqk6wgl
- Start time basis: `vendor_reported`.

  - `2025-11-19 19:08` investigating — We are investigating loading issues on our web application. As a result, some users might be getting errors when loading the web application.
  - `2025-11-19 19:28` investigating — We continue investigating the issue with web application. Data processing and alerting remain operational.
  - `2025-11-19 19:37` monitoring — A fix has been implemented and we are monitoring the results.
  - `2025-11-19 19:44` resolved — This incident has been resolved as of 2:32PM ET.
