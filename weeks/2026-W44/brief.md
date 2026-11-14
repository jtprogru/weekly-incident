# Week 2026-W44 — top 5 of 24 above threshold

2026-10-26 to 2026-11-01 (UTC).

52 source records (60 saved revisions) → 52 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## 1. cloudflare — Cloudflare Pages Upload Issue

- URL: https://www.cloudflarestatus.com/incidents/me060q7dbayp
- Window: 2026-10-26 00:00 to 2026-11-04 06:35 UTC (9d 6h 35m)
- Impact: minor, status: resolved
- Components: Dashboard, Pages

<a id="ci-ac1206f83d1f52f837cc0aac"></a>

- Canonical ID: `ci-ac1206f83d1f52f837cc0aac` (1 source records).
  - [cloudflare/me060q7dbayp](../../data/2026/W44/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/me060q7dbayp
- Start time basis: `vendor_reported`.

  - `2026-11-04 07:11` resolved — Between 2026/10/26 - 2026/11/4, some customers trying to upload static Pages assets from the Cloudflare Dashboard may have experienced issues creating a Pages project.

## 2. aws — Increased invoke error rates

- URL: https://health.aws.amazon.com/health/status
- Window: 2026-10-30 01:28 to not yet (15d 11h 8m (ongoing))
- Impact: unknown, status: investigating
- Components: apigateway-us-west-2

<a id="ci-d86c52c83cfb265a4c4667b8"></a>

- Canonical ID: `ci-d86c52c83cfb265a4c4667b8` (1 source records).
  - [aws/0fd29a8bd6906892](../../data/2026/W44/aws.json); vendor URL: https://health.aws.amazon.com/health/status
- Start time basis: `first_observed_update`.
- Warning: AWS onset is the first observed message, not a verified impact start

  - `2026-10-30 01:28` investigating — Between 7:27 AM and 8:47 AM PDT, and between 10:47 AM and 4:15 PM PDT, we experienced intermittent invoke error rates and latencies for HTTP APIs in the US-WEST-2 Region. During this time, customers would have observed increased latency or errors when invoking HTTP APIs, with elevated errors occurring between 3:12 PM and 3:45 PM PDT. This represented the period of highest customer impact with the majority of affected customers experiencing errors or increased latencies during this window. Majority of retry attempts for affected requests were succeeding. We communicated to impacted customers via the Your Account view of the AWS Health Dashboard during both impact windows. We identified the root cause to be an issue with an internal subsystem and by 3:45 PM we applied the necessary mitigatio…

## 3. cloudflare — Network Performance Issues in Ashburn, IAD

- URL: https://www.cloudflarestatus.com/incidents/snth5ero95rz
- Window: 2026-10-28 04:39 to 2026-10-28 22:18 UTC (17h 39m)
- Impact: minor, status: resolved
- Components: Ashburn, United States - (IAD)

<a id="ci-0a026ea52a6b18844867b0a1"></a>

- Canonical ID: `ci-0a026ea52a6b18844867b0a1` (1 source records).
  - [cloudflare/snth5ero95rz](../../data/2026/W44/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/snth5ero95rz
- Start time basis: `vendor_reported`.

  - `2026-10-28 04:39` investigating — Cloudflare is investigating network performance issues affecting our Ashburn (IAD) data center. Traffic routed through this location may experience elevated latency and/or 5xx errors. We are working to analyze and mitigate this problem. More updates to follow shortly.
  - `2026-10-28 05:05` monitoring — We are continuing to monitor for any further issues.
  - `2026-10-28 22:18` resolved — This incident has been resolved.

## 4. cloudflare — Increased Errors for Durable Objects in the Eastern North America region

- URL: https://www.cloudflarestatus.com/incidents/klub12zy71b4
- Window: 2026-10-28 00:17 to 2026-10-28 03:28 UTC (3h 11m)
- Impact: minor, status: resolved
- Components: D1, Durable Objects, Queues, Workflows

<a id="ci-05e287e0555cc3e790b274a8"></a>

- Canonical ID: `ci-05e287e0555cc3e790b274a8` (1 source records).
  - [cloudflare/klub12zy71b4](../../data/2026/W44/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/klub12zy71b4
- Start time basis: `vendor_reported`.

  - `2026-10-28 00:17` investigating — Cloudflare is aware of, and investigating an issue with Durable Objects and downstream services which potentially impacts multiple customers in the Eastern North America region. Durable Objects are experiencing an elevated level of overload errors in the aforementioned region. This may also affect downstream services such as D1, Queues and Workflows etc. We are currently investigating this issue.
  - `2026-10-28 01:22` monitoring — A fix has been implemented and we are monitoring the results.
  - `2026-10-28 03:28` resolved — This incident has been resolved.

## 5. cloudflare — Increased Durable Objects, D1 and Workflows Error Rates

- URL: https://www.cloudflarestatus.com/incidents/r4mz367vddi7
- Window: 2026-10-29 14:16 to 2026-10-29 16:33 UTC (2h 16m)
- Impact: minor, status: resolved
- Components: Artifacts, D1, Durable Objects, Workers Assets, Workflows

<a id="ci-fdbeb462c66148cf9f3c0500"></a>

- Canonical ID: `ci-fdbeb462c66148cf9f3c0500` (1 source records).
  - [cloudflare/r4mz367vddi7](../../data/2026/W44/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/r4mz367vddi7
- Start time basis: `vendor_reported`.

  - `2026-10-29 14:16` investigating — Cloudflare is investigating an issue where customers are experiencing elevated errors for operations on Increased Durable Objects, D1 and Workflows.
  - `2026-10-29 16:33` monitoring — A fix has been implemented and we are monitoring the results.
  - `2026-10-29 16:33` resolved — This incident has been resolved.

## Also above threshold (19)

- `datadog` Metrics query are disrupted (2h 29m) — https://stspg.io/d4klp5znv9c7
- `cloudflare` Increased Worker services errors (2h 26m) — https://www.cloudflarestatus.com/incidents/6yfpsos7ksp3
- `anthropic` Elevated errors affecting the Google Calendar connector (3h 46m) — https://stspg.io/75qq31h8672d
- `stripe` Elevated Issuing errors from 17:13 - 20:41 UTC (4h 36m) — https://stspg.io/kj9ch2nd498j
- `github` Incident with Actions, Pages, and Webhooks (1h 32m) — https://stspg.io/5cd2m4n59spv
- `cloudflare` Delayed, missed, and duplicated Workers Cron Triggers (3h 26m) — https://www.cloudflarestatus.com/incidents/97y1vy5nsjrn
- `openai` Elevated errors for GPT Image 2 API image edits (4h 1m) — https://status.openai.com/incidents/01M66KSDH67V5BQ73PE8AD2W8S
- `cloudflare` ERR_TIMED_OUT and ERR_EMPTY_RESPONSE being observed by Cloudflare Customers (2h 43m) — https://www.cloudflarestatus.com/incidents/dk2obcyw2aew
- `grafana` Degraded Cloud Integrations (4h 41m) — https://stspg.io/d7mh6vy0g5jh
- `anthropic` Elevated errors for Claude Sonnet 5.5 (42m) — https://stspg.io/rql9h590npxd
- `anthropic` Elevated errors for Claude Sonnet 5.5 (41m) — https://stspg.io/39x7qh3sd7bd
- `github` Disruption with Copilot Cloud services (2h 29m) — https://stspg.io/22p3vp78z20q
- `github` Incident with Dependabot (2h 29m) — https://stspg.io/5bh6qnzn76tk
- `digitalocean` Kubernetes Cluster Management Operations (3h 40m) — https://stspg.io/b16t5y07v5l4
- `datadog` Elevated Error Rates for Workflow Automation, App Builder, Bits Code, Bits Agent Builder (56m) — https://stspg.io/rq69kpmmw7h6
- `stripe` Elevated Instant Payouts failures (2h 5m) — https://stspg.io/3gk09h1d44ts
- `github` Disruption with some Copilot services (1h 15m) — https://stspg.io/bwcyq078jzms
- `github` Disruption with GitHub services (31m) — https://stspg.io/6p81rf34z4nc
- `grafana` Frontend Observability Outage in prod-eu-west-2 (27m) — https://stspg.io/dgmytp3hyf4l
