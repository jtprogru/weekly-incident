# Week 2025-W47

2025-11-17 to 2025-11-23 (UTC). Generated 2026-11-14 00:00.

3 canonical incidents, 2 above threshold. All 13 sources responded.

3 source records (3 saved revisions) → 3 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | datadog | [Dashboards Not Loading](https://stspg.io/bbdxgwpv272q) | 2025-11-17 16:23 | 56m | major | Web Application | 56 = 56 × 1 × 1.0 |
| 2 | datadog | [Web Application Not Loading](https://stspg.io/h30gbqqk6wgl) | 2025-11-19 19:08 | 36m | critical | Web Application | 36 = 36 × 1 × 1.0 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. datadog — Dashboards Not Loading

- URL: https://stspg.io/bbdxgwpv272q
- Started: 2025-11-17 16:23 UTC
- Resolved: 2025-11-17 17:20 UTC
- Duration: 56m
- Impact: major, status: resolved
- Components: Web Application
- Canonical incident: [ci-54a50f8d3a7c078291140bed](#ci-54a50f8d3a7c078291140bed)

- `2025-11-17 16:23` **investigating** — We are investigating loading issues on the dashboard pages. As a result, some users might be getting errors when loading the web application. Please note that data processing and alerts are not affected by this incident.
- `2025-11-17 16:40` **identified** — The issue has been identified and we taking measures to mitigate the issue, as well as working on a fix.
- `2025-11-17 17:12` **monitoring** — The rollout with a fix is in progress, and we're no longer seeing errors, and are currently monitoring the incident and we are on the path of recovery.
- `2025-11-17 17:20` **resolved** — All errors stopped as of 12:02ET. This incident has been resolved.

### 2. datadog — Web Application Not Loading

- URL: https://stspg.io/h30gbqqk6wgl
- Started: 2025-11-19 19:08 UTC
- Resolved: 2025-11-19 19:44 UTC
- Duration: 36m
- Impact: critical, status: resolved
- Components: Web Application
- Canonical incident: [ci-8e0d54461822f0521802bb74](#ci-8e0d54461822f0521802bb74)

- `2025-11-19 19:08` **investigating** — We are investigating loading issues on our web application. As a result, some users might be getting errors when loading the web application.
- `2025-11-19 19:28` **investigating** — We continue investigating the issue with web application. Data processing and alerting remain operational.
- `2025-11-19 19:37` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2025-11-19 19:44` **resolved** — This incident has been resolved as of 2:32PM ET.

## Below threshold (1)

- `datadog` [Delayed Monitors Notifications](https://stspg.io/0r75b5ng1fc2) — 36m

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-54a50f8d3a7c078291140bed"></a>

- Canonical ID: `ci-54a50f8d3a7c078291140bed` (1 source records).
  - [datadog/bg2rxpyn1hc3](../../data/2025/W47/datadog.json); vendor URL: https://stspg.io/bbdxgwpv272q
- Start time basis: `vendor_reported`.

<a id="ci-7aad2e0629273bdb6b907bdd"></a>

- Canonical ID: `ci-7aad2e0629273bdb6b907bdd` (1 source records).
  - [datadog/29fsc4rhsscg](../../data/2025/W47/datadog.json); vendor URL: https://stspg.io/0r75b5ng1fc2
- Start time basis: `vendor_reported`.

<a id="ci-8e0d54461822f0521802bb74"></a>

- Canonical ID: `ci-8e0d54461822f0521802bb74` (1 source records).
  - [datadog/cvdjtf81756n](../../data/2025/W47/datadog.json); vendor URL: https://stspg.io/h30gbqqk6wgl
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
