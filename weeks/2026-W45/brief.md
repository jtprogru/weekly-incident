# Week 2026-W45 — top 5 of 16 above threshold

2026-11-02 to 2026-11-08 (UTC).

45 source records (56 saved revisions) → 45 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## 1. cloudflare — Fedramp Customers have Delayed Purges

- URL: https://www.cloudflarestatus.com/incidents/wp49bp8bqo6c
- Window: 2026-11-06 17:47 to 2026-11-12 16:36 UTC (5d 22h 49m)
- Impact: minor, status: resolved
- Components: CDN Cache Purge

<a id="ci-b723223fb9dd7c34dba6ff7b"></a>

- Canonical ID: `ci-b723223fb9dd7c34dba6ff7b` (1 source records).
  - [cloudflare/wp49bp8bqo6c](../../data/2026/W45/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/wp49bp8bqo6c
- Start time basis: `vendor_reported`.

  - `2026-11-06 17:47` investigating — Cloudflare has identified a bug with our purge pipeline impacting fedramp customers. Stale content may be served until expired. Non-fedramp customers are not impacted.
  - `2026-11-06 17:59` identified — We are continuing to work on a fix for the issue.
  - `2026-11-12 16:36` resolved — This incident has been resolved.

## 2. openai — Elevated error rates with some U.S. business Mastercards

- URL: https://status.openai.com/incidents/01M6S2FV5TZWQ753NHRB7BF5K7
- Window: 2026-11-06 18:38 to 2026-11-10 16:06 UTC (3d 21h 27m)
- Impact: none, status: resolved
- Components: —

<a id="ci-b8f870d1369b4ec8be6fe40f"></a>

- Canonical ID: `ci-b8f870d1369b4ec8be6fe40f` (1 source records).
  - [openai/01M6S2FV5TZWQ753NHRB7BF5K7](../../data/2026/W45/openai.json); vendor URL: https://status.openai.com/incidents/01M6S2FV5TZWQ753NHRB7BF5K7
- Start time basis: `vendor_reported`.

  - `2026-11-06 18:38` identified — **Investigating — We’re seeing payment declines on some U.S. business Mastercards. Affected cards have numbers beginning with 544015, 546400, 543502, 543799, or 512998. Some cards beginning with 513989 are also affected.** **For API customers, this can prevent automatic recharges and one-time credit purchases. If your prepaid balance runs out, your API requests may fail.** **If you use one of these cards, we recommend adding a card from a different issuer before your next payment, even if you haven’t seen a decline. If an API credit purchase has already failed, use the new card to retry the purchase.** **We’re working with our payment partners and will share updates here.**
  - `2026-11-06 18:50` identified — Investigating — We’re seeing payment declines on some U.S. business Mastercards. Affected cards have numbers beginning with 544015, 546400, 543502, 543799, or 512998. Some cards beginning with 513989 are also affected. For API customers, this can prevent automatic recharges and one-time credit purchases. If your prepaid balance runs out, your API requests may fail. If you use one of these cards, we recommend adding a card from a different issuer before your next payment, even if you haven’t seen a decline. If an API credit purchase has already failed, use the new card to retry the purchase. We’re working with our payment partners and will share updates here.
  - `2026-11-07 01:29` monitoring — We’ve applied a mitigation and are monitoring the affected services to confirm recovery.
  - `2026-11-10 16:06` resolved — All affected services are operating normally again. We’re sorry for the disruption, and we appreciate your patience.

## 3. cloudflare — Spectrum Availability Issue in Phoenix

- URL: https://www.cloudflarestatus.com/incidents/c60k3g4vsoc1
- Window: 2026-11-07 00:59 to 2026-11-09 14:36 UTC (2d 13h 37m)
- Impact: minor, status: resolved
- Components: Spectrum

<a id="ci-5d9778b9959939de25c5186b"></a>

- Canonical ID: `ci-5d9778b9959939de25c5186b` (1 source records).
  - [cloudflare/c60k3g4vsoc1](../../data/2026/W45/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/c60k3g4vsoc1
- Start time basis: `vendor_reported`.

  - `2026-11-07 00:59` investigating — Cloudflare is investigating issues with Spectrum in Phoenix. This may be impacting the availability of existing Spectrum applications. We are working to understand the full impact and mitigate this problem. More updates to follow shortly.
  - `2026-11-07 02:51` monitoring — We are continuing to monitor for any further issues.
  - `2026-11-09 14:36` resolved — This incident has been resolved.

## 4. grafana — October usage report visualization issue

- URL: https://stspg.io/ksvgzjk4ctnl
- Window: 2026-11-03 22:00 to 2026-11-07 01:16 UTC (3d 3h 16m)
- Impact: minor, status: resolved
- Components: Cost Management and Billing

<a id="ci-ee5265daee2b7cfe1debc8fc"></a>

- Canonical ID: `ci-ee5265daee2b7cfe1debc8fc` (1 source records).
  - [grafana/sytnxyqfmbwg](../../data/2026/W45/grafana.json); vendor URL: https://stspg.io/ksvgzjk4ctnl
- Start time basis: `vendor_reported`.

  - `2026-11-03 22:00` identified — There was an issue with our billing metrics exporter that caused us to mis-report usage for October at month-end. This is causing October usage to be reported as 0 in our cost management app. We are working to backfill the correct data. Note that this did not affect billing, and is only a visualization issue.
  - `2026-11-04 22:11` identified — We are still working to backfill the correct data. Note that this did not affect billing, and is only a visualization issue.
  - `2026-11-06 21:41` identified — We are currently back filling the missing data. We will provide another update once the backfill has been completed.
  - `2026-11-07 01:16` resolved — This incident has been resolved.

## 5. grafana — IRM Unable to Voice Call Brazil Numbers

- URL: https://stspg.io/m1k738p976cl
- Window: 2026-11-02 00:36 to 2026-11-04 08:45 UTC (2d 8h 9m)
- Impact: major, status: resolved
- Components: Incident Management and Response (IRM)

<a id="ci-99928e6fdb4d472a21f59584"></a>

- Canonical ID: `ci-99928e6fdb4d472a21f59584` (1 source records).
  - [grafana/507h3wxc3mdh](../../data/2026/W45/grafana.json); vendor URL: https://stspg.io/m1k738p976cl
- Start time basis: `vendor_reported`.

  - `2026-11-02 00:36` investigating — We are currently investigating an issue impacting Voice calls to Brazil for IRM. We are working on a resolution, and an update will follow as soon as we have more info. Users should configure mobile push or SMS as notifications.
  - `2026-11-02 15:07` investigating — We are continuing to work on a resolution, and an update will follow as soon as we have more info. Users should configure mobile push or SMS as notifications.
  - `2026-11-02 17:11` investigating — Our efforts towards a resolution continues. We encourage usage of SMS and mobile push as alternative way in the meantime. Thank you for your patience.
  - `2026-11-03 16:12` investigating — We are starting to see signs of recovery therefore some users might see voice calls working fine for them, but the issue is not fully resolved yet. We are continuing our efforts towards a complete resolution and an update will follow as soon as we have more info. We encourage usage of SMS and mobile push as alternative notification mechanism in the meantime.
  - `2026-11-04 08:45` resolved — Brazil calls recovered at about 13:00 UTC on 3 November. At this point, we are marking the incident as resolved.

## Also above threshold (11)

- `cloudflare` Increased Latency and Errors for R2 and Downstream Services in Western North America (WNAM) (9h 15m) — https://www.cloudflarestatus.com/incidents/dr9s0uvzwk6n
- `cloudflare` Increased Errors for Durable Objects in Western Europe (19h 41m) — https://www.cloudflarestatus.com/incidents/oy8gxjb38yse
- `openai` ChatGPT Work and Dots Errors (13h 6m) — https://status.openai.com/incidents/01M6P50GPP1V98EF8MJ9SH393R
- `github` Incident with Pull Requests, Git Operations and Issues (2h 0m) — https://stspg.io/18blsgwnbyy6
- `cloudflare` Cloudflare Workers and Pages build errors (4h 51m) — https://www.cloudflarestatus.com/incidents/nbyxn84kvb6o
- `openai` File upload errors in ChatGPT (4h 52m) — https://status.openai.com/incidents/01M6FDZ89W4PKDS19K496J307G
- `aws` Elevated Launch Failures (2h 21m) — https://health.aws.amazon.com/health/status
- `cloudflare` Network Performance Issues in APJC (2h 30m) — https://www.cloudflarestatus.com/incidents/j25w3rpv45gs
- `cloudflare` Cloudflare Gateway DNS-over-TLS Connectivity Issues over IPv4 (2h 21m) — https://www.cloudflarestatus.com/incidents/3bv9locsg7ln
- `github` Incident with Actions and Pages (55m) — https://stspg.io/pjgptg3cfc6n
- `github` Incident with Actions (35m) — https://stspg.io/7pwv28hvkmz8
