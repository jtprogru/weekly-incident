# Week 2026-W45 — top 5 of 16 above threshold

2026-11-02 to 2026-11-08 (UTC).

## 1. cloudflare — Fedramp Customers have Delayed Purges

- URL: https://www.cloudflarestatus.com/incidents/wp49bp8bqo6c
- Window: 2026-11-06 17:47 to not yet (2d 23h 10m (ongoing))
- Impact: minor, status: identified
- Components: CDN Cache Purge

  - `2026-11-06 17:47` investigating — Cloudflare has identified a bug with our purge pipeline impacting fedramp customers. Stale content may be served until expired. Non-fedramp customers are not impacted.
  - `2026-11-06 17:59` identified — We are continuing to work on a fix for the issue.

## 2. cloudflare — Spectrum Availability Issue in Phoenix

- URL: https://www.cloudflarestatus.com/incidents/c60k3g4vsoc1
- Window: 2026-11-07 00:59 to not yet (2d 15h 58m (ongoing))
- Impact: minor, status: monitoring
- Components: Spectrum

  - `2026-11-07 00:59` investigating — Cloudflare is investigating issues with Spectrum in Phoenix. This may be impacting the availability of existing Spectrum applications. We are working to understand the full impact and mitigate this problem. More updates to follow shortly.
  - `2026-11-07 02:51` monitoring — We are continuing to monitor for any further issues.

## 3. openai — Elevated error rates with some U.S. business Mastercards

- URL: https://status.openai.com/incidents/01M6S2FV5TZWQ753NHRB7BF5K7
- Window: 2026-11-06 18:38 to not yet (2d 22h 18m (ongoing))
- Impact: none, status: monitoring
- Components: —

  - `2026-11-06 18:38` identified — **Investigating — We’re seeing payment declines on some U.S. business Mastercards. Affected cards have numbers beginning with 544015, 546400, 543502, 543799, or 512998. Some cards beginning with 513989 are also affected.** **For API customers, this can prevent automatic recharges and one-time credit purchases. If your prepaid balance runs out, your API requests may fail.** **If you use one of these cards, we recommend adding a card from a different issuer before your next payment, even if you haven’t seen a decline. If an API credit purchase has already failed, use the new card to retry the purchase.** **We’re working with our payment partners and will share updates here.**
  - `2026-11-06 18:50` identified — Investigating — We’re seeing payment declines on some U.S. business Mastercards. Affected cards have numbers beginning with 544015, 546400, 543502, 543799, or 512998. Some cards beginning with 513989 are also affected. For API customers, this can prevent automatic recharges and one-time credit purchases. If your prepaid balance runs out, your API requests may fail. If you use one of these cards, we recommend adding a card from a different issuer before your next payment, even if you haven’t seen a decline. If an API credit purchase has already failed, use the new card to retry the purchase. We’re working with our payment partners and will share updates here.
  - `2026-11-07 01:29` monitoring — We’ve applied a mitigation and are monitoring the affected services to confirm recovery.

## 4. grafana — October usage report visualization issue

- URL: https://stspg.io/ksvgzjk4ctnl
- Window: 2026-11-03 22:00 to 2026-11-07 01:16 UTC (3d 3h 16m)
- Impact: minor, status: resolved
- Components: Cost Management and Billing

  - `2026-11-03 22:00` identified — There was an issue with our billing metrics exporter that caused us to mis-report usage for October at month-end. This is causing October usage to be reported as 0 in our cost management app. We are working to backfill the correct data. Note that this did not affect billing, and is only a visualization issue.
  - `2026-11-04 22:11` identified — We are still working to backfill the correct data. Note that this did not affect billing, and is only a visualization issue.
  - `2026-11-06 21:41` identified — We are currently back filling the missing data. We will provide another update once the backfill has been completed.
  - `2026-11-07 01:16` resolved — This incident has been resolved.

## 5. grafana — IRM Unable to Voice Call Brazil Numbers

- URL: https://stspg.io/m1k738p976cl
- Window: 2026-11-02 00:36 to 2026-11-04 08:45 UTC (2d 8h 9m)
- Impact: major, status: resolved
- Components: Incident Management and Response (IRM)

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

