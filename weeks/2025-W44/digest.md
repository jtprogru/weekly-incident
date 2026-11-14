# Week 2025-W44

2025-10-27 to 2025-11-02 (UTC). Generated 2026-11-14 00:00.

1 canonical incidents, 1 above threshold. All 13 sources responded.

1 source records (1 saved revisions) → 1 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | datadog | [Metrics data ingestion delayed and monitor evaluations degraded](https://stspg.io/ch4x738bv5qx) | 2025-10-28 18:32 | 3h 29m | major | Metrics and Infra Monitoring, Monitors | 418 = 209 × 2 × 1.0 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. datadog — Metrics data ingestion delayed and monitor evaluations degraded

- URL: https://stspg.io/ch4x738bv5qx
- Started: 2025-10-28 18:32 UTC
- Resolved: 2025-10-28 22:01 UTC
- Duration: 3h 29m
- Impact: major, status: resolved
- Components: Metrics and Infra Monitoring, Monitors
- Canonical incident: [ci-edb91f28c8e38b57a84a3872](#ci-edb91f28c8e38b57a84a3872)

- `2025-10-28 18:32` **identified** — The issue has been identified and a fix is being implemented.
- `2025-10-28 22:01` **resolved** — This incident has been resolved.

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-edb91f28c8e38b57a84a3872"></a>

- Canonical ID: `ci-edb91f28c8e38b57a84a3872` (1 source records).
  - [datadog/01x71sjh4ymw](../../data/2025/W44/datadog.json); vendor URL: https://stspg.io/ch4x738bv5qx
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
