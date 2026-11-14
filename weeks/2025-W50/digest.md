# Week 2025-W50

2025-12-08 to 2025-12-14 (UTC). Generated 2026-11-14 00:00.

3 canonical incidents, 2 above threshold. All 14 sources responded.

3 source records (3 saved revisions) → 3 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | datadog | [Delayed APM metric ingestion](https://stspg.io/qwpxql6b48pl) | 2025-12-12 18:37 | 3h 15m | major | Metrics and Infra Monitoring | 195 = 195 × 1 × 1.0 |
| 2 | datadog | [Delayed Processes data](https://stspg.io/ppc0471xzd9w) | 2025-12-12 21:49 | 1h 54m | major | Metrics and Infra Monitoring | 114 = 114 × 1 × 1.0 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. datadog — Delayed APM metric ingestion

- URL: https://stspg.io/qwpxql6b48pl
- Started: 2025-12-12 18:37 UTC
- Resolved: 2025-12-12 21:52 UTC
- Duration: 3h 15m
- Impact: major, status: resolved
- Components: Metrics and Infra Monitoring
- Canonical incident: [ci-d42c11c45abe62d91a8e93fb](#ci-d42c11c45abe62d91a8e93fb)

- `2025-12-12 18:37` **investigating** — We are currently investigating lag in ingesting apm metrics, which affects monitor evaluation.
- `2025-12-12 18:55` **investigating** — We are continuing to investigate this issue.
- `2025-12-12 19:26` **investigating** — We are currently investigating lag in ingesting apm and process metrics, which affects monitor evaluation and in some cases led to incorrect monitor alerts.
- `2025-12-12 20:53` **identified** — We have identified the issue affecting ingestion delays in apm and process metrics and are working on recovery
- `2025-12-12 21:52` **resolved** — All impact related to APM metrics has been resolved. A separate incident has been created to track the remaining impact in live process data.

### 2. datadog — Delayed Processes data

- URL: https://stspg.io/ppc0471xzd9w
- Started: 2025-12-12 21:49 UTC
- Resolved: 2025-12-12 23:43 UTC
- Duration: 1h 54m
- Impact: major, status: resolved
- Components: Metrics and Infra Monitoring
- Canonical incident: [ci-afca22236f96448dec34a02e](#ci-afca22236f96448dec34a02e)

- `2025-12-12 21:49` **investigating** — We are investigating increased latency processing Processes data. As a result of this issue, some users may see delays or gaps for data based on Process Monitoring. To prevent spurious alerts, we have temporarily disabled monitors based on this data.
- `2025-12-12 21:53` **investigating** — We are continuing to investigate this issue.
- `2025-12-12 22:41` **identified** — The issue has been identified and a fix is being implemented.
- `2025-12-12 23:31` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2025-12-12 23:43` **resolved** — This incident has been resolved.

## Below threshold (1)

- `datadog` [Metrics data ingestion delayed and monitor evaluations degraded](https://stspg.io/5hqhsx3549dz) — 52m

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-d42c11c45abe62d91a8e93fb"></a>

- Canonical ID: `ci-d42c11c45abe62d91a8e93fb` (1 source records).
  - [datadog/c4gh438m6pp3](../../data/2025/W50/datadog.json); vendor URL: https://stspg.io/qwpxql6b48pl
- Start time basis: `vendor_reported`.

<a id="ci-afca22236f96448dec34a02e"></a>

- Canonical ID: `ci-afca22236f96448dec34a02e` (1 source records).
  - [datadog/rxrmt186707r](../../data/2025/W50/datadog.json); vendor URL: https://stspg.io/ppc0471xzd9w
- Start time basis: `vendor_reported`.

<a id="ci-9e58ced657109922cb72822a"></a>

- Canonical ID: `ci-9e58ced657109922cb72822a` (1 source records).
  - [datadog/xj87f6z9mm3w](../../data/2025/W50/datadog.json); vendor URL: https://stspg.io/5hqhsx3549dz
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
