# Week 2023-W40

2023-10-02 to 2023-10-08 (UTC). Generated 2026-11-14 00:00.

1 canonical incidents, 1 above threshold. All 13 sources responded.

1 source records (1 saved revisions) → 1 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | atlassian | [For a limited number of customers, activation for new products on some tenants is blocked](https://stspg.io/9q0x4gy5g22l) | 2023-10-04 14:28 | 2h 52m | none | — | 172 = 172 × 1 × 1.0 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. atlassian — For a limited number of customers, activation for new products on some tenants is blocked

- URL: https://stspg.io/9q0x4gy5g22l
- Started: 2023-10-04 14:28 UTC
- Resolved: 2023-10-04 17:20 UTC
- Duration: 2h 52m
- Impact: none, status: resolved
- Components: —
- Canonical incident: [ci-fe90d35bc5ffacfd43566c60](#ci-fe90d35bc5ffacfd43566c60)

- `2023-10-04 14:28` **identified** — We identified an issue with the product provisioning service. We identified the issue and we are working on a fix. We will post an update in the next hour with estimated fix time
- `2023-10-04 15:01` **identified** — We started fixing the affected tenants, we will post an update on the estimated time to resolve all tenants in the next hour
- `2023-10-04 15:47` **monitoring** — We've solved 15% of the issues and have a plan to recover 100% exponentially within the hour
- `2023-10-04 16:20` **monitoring** — We've solved 52% of the issues and are still on track to recover 100% within the original timeframe
- `2023-10-04 17:01` **monitoring** — 99% of issues fixed and services are back to normal. Last 1% are manually begin fixed
- `2023-10-04 17:20` **resolved** — 100% resolved and services running normally

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-fe90d35bc5ffacfd43566c60"></a>

- Canonical ID: `ci-fe90d35bc5ffacfd43566c60` (1 source records).
  - [atlassian/b2lqb2yxry16](../../data/2023/W40/atlassian.json); vendor URL: https://stspg.io/9q0x4gy5g22l
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
