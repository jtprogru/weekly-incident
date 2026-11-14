# Week 2023-W40 — top 1 of 1 above threshold

2023-10-02 to 2023-10-08 (UTC).

1 source records (1 saved revisions) → 1 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## 1. atlassian — For a limited number of customers, activation for new products on some tenants is blocked

- URL: https://stspg.io/9q0x4gy5g22l
- Window: 2023-10-04 14:28 to 2023-10-04 17:20 UTC (2h 52m)
- Impact: none, status: resolved
- Components: —

<a id="ci-fe90d35bc5ffacfd43566c60"></a>

- Canonical ID: `ci-fe90d35bc5ffacfd43566c60` (1 source records).
  - [atlassian/b2lqb2yxry16](../../data/2023/W40/atlassian.json); vendor URL: https://stspg.io/9q0x4gy5g22l
- Start time basis: `vendor_reported`.

  - `2023-10-04 14:28` identified — We identified an issue with the product provisioning service. We identified the issue and we are working on a fix. We will post an update in the next hour with estimated fix time
  - `2023-10-04 15:01` identified — We started fixing the affected tenants, we will post an update on the estimated time to resolve all tenants in the next hour
  - `2023-10-04 15:47` monitoring — We've solved 15% of the issues and have a plan to recover 100% exponentially within the hour
  - `2023-10-04 16:20` monitoring — We've solved 52% of the issues and are still on track to recover 100% within the original timeframe
  - `2023-10-04 17:01` monitoring — 99% of issues fixed and services are back to normal. Last 1% are manually begin fixed
  - `2023-10-04 17:20` resolved — 100% resolved and services running normally
