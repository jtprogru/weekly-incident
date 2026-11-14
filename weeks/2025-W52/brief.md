# Week 2025-W52 — top 1 of 1 above threshold

2025-12-22 to 2025-12-28 (UTC).

2 source records (2 saved revisions) → 2 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## 1. stripe — Elevated iDeal errors

- URL: https://stspg.io/w20wf142qn51
- Window: 2025-12-22 15:50 to 2025-12-22 16:36 UTC (45m)
- Impact: major, status: resolved
- Components: Acquirers and payment methods

<a id="ci-1c54b942dc8277134a9db974"></a>

- Canonical ID: `ci-1c54b942dc8277134a9db974` (1 source records).
  - [stripe/fs326pl1tcn9](../../data/2025/W52/stripe.json); vendor URL: https://stspg.io/w20wf142qn51
- Start time basis: `vendor_reported`.

  - `2025-12-22 15:50` investigating — We're currently investigating a service disruption of the iDeal payment method that began at 15:33 UTC. This issue is external to Stripe and is causing increased error rates for transactions processed through the iDeal payment method. We are actively monitoring the situation and will provide updates as more information becomes available.
  - `2025-12-22 16:36` resolved — From 15:33 - 16:21 UTC, we saw elevated errors with the iDeal payment method. The issue was external to Stripe and is now resolved.
