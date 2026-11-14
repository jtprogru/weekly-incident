# Week 2025-W52

2025-12-22 to 2025-12-28 (UTC). Generated 2026-11-14 00:00.

2 canonical incidents, 1 above threshold. All 13 sources responded.

2 source records (2 saved revisions) → 2 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | stripe | [Elevated iDeal errors](https://stspg.io/w20wf142qn51) | 2025-12-22 15:50 | 45m | major | Acquirers and payment methods | 54 = 45 × 1 × 1.2 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. stripe — Elevated iDeal errors

- URL: https://stspg.io/w20wf142qn51
- Started: 2025-12-22 15:50 UTC
- Resolved: 2025-12-22 16:36 UTC
- Duration: 45m
- Impact: major, status: resolved
- Components: Acquirers and payment methods
- Canonical incident: [ci-1c54b942dc8277134a9db974](#ci-1c54b942dc8277134a9db974)

- `2025-12-22 15:50` **investigating** — We're currently investigating a service disruption of the iDeal payment method that began at 15:33 UTC. This issue is external to Stripe and is causing increased error rates for transactions processed through the iDeal payment method. We are actively monitoring the situation and will provide updates as more information becomes available.
- `2025-12-22 16:36` **resolved** — From 15:33 - 16:21 UTC, we saw elevated errors with the iDeal payment method. The issue was external to Stripe and is now resolved.

## Below threshold (1)

- `stripe` [Elevated Swish errors](https://stspg.io/yc8n8zr8nvmn) — 44m

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-1c54b942dc8277134a9db974"></a>

- Canonical ID: `ci-1c54b942dc8277134a9db974` (1 source records).
  - [stripe/fs326pl1tcn9](../../data/2025/W52/stripe.json); vendor URL: https://stspg.io/w20wf142qn51
- Start time basis: `vendor_reported`.

<a id="ci-fafc7400a5e0d17a40ae002b"></a>

- Canonical ID: `ci-fafc7400a5e0d17a40ae002b` (1 source records).
  - [stripe/qtm1lh8g56l6](../../data/2025/W52/stripe.json); vendor URL: https://stspg.io/yc8n8zr8nvmn
- Start time basis: `vendor_reported`.

## Sources

| Vendor | Status | Source records seen | Parse errors |
|--------|--------|----------------|--------------|
| anthropic | ok | 50 | 0 |
| atlassian | ok | 34 | 0 |
| aws | ok | 4 | 0 |
| circleci | ok | 50 | 0 |
| cloudflare | ok | 50 | 0 |
| datadog | ok | 50 | 0 |
| digitalocean | ok | 50 | 0 |
| gcp | ok | 5 | 0 |
| github | ok | 50 | 0 |
| grafana | ok | 50 | 0 |
| openai | ok | 25 | 0 |
| slack | ok | 39 | 0 |
| stripe | ok | 50 | 0 |
