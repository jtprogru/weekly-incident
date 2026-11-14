# Week 2026-W12 — top 2 of 2 above threshold

2026-03-16 to 2026-03-22 (UTC).

5 source records (5 saved revisions) → 5 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## 1. stripe — Elevated Pix errors

- URL: https://stspg.io/dzll5mrw1ryz
- Window: 2026-03-22 16:02 to 2026-03-23 11:27 UTC (19h 24m)
- Impact: minor, status: resolved
- Components: Acquirers and payment methods

<a id="ci-2a67f55dca2afd6cd9d25c1f"></a>

- Canonical ID: `ci-2a67f55dca2afd6cd9d25c1f` (1 source records).
  - [stripe/gc6z4c7t1hdq](../../data/2026/W12/stripe.json); vendor URL: https://stspg.io/dzll5mrw1ryz
- Start time basis: `vendor_reported`.

  - `2026-03-22 16:02` identified — We are currently observing elevated errors on Pix payments. Payments are likely affected due to degraded performance reported by Pix. We will continue to monitor the issue and an update will be provided as soon as more information becomes available.
  - `2026-03-23 11:27` resolved — Between 14:13 UTC on the 22nd March and 01:40 UTC on the 23rd March, we observed a service disruption that impacted the Pix payment method. The issue was external to Stripe and is now resolved.

## 2. datadog — Intermittent "No Data" Status for Monitors

- URL: https://stspg.io/43kkcl1tqrrk
- Window: 2026-03-19 00:54 to 2026-03-19 02:53 UTC (1h 59m)
- Impact: major, status: resolved
- Components: Monitors

<a id="ci-14dda82a474ebe0feb60f33d"></a>

- Canonical ID: `ci-14dda82a474ebe0feb60f33d` (1 source records).
  - [datadog/j5zpb9phy9m6](../../data/2026/W12/datadog.json); vendor URL: https://stspg.io/43kkcl1tqrrk
- Start time basis: `vendor_reported`.

  - `2026-03-19 00:54` investigating — We are investigating an issue causing some metric monitors to intermittently report "No Data" status. This primarily affects monitors based on distribution metrics. Monitor alerting may be unreliable during this time. We are actively investigating and will provide updates as available.
  - `2026-03-19 01:27` investigating — We are continuing to investigate the issue and will provide updates as available.
  - `2026-03-19 02:04` investigating — We are continuing to investigate the issue and will provide updates as available.
  - `2026-03-19 02:39` monitoring — A fix has been implemented and we are monitoring the results.
  - `2026-03-19 02:53` resolved — This incident has been resolved.
