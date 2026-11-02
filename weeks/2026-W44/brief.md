# Week 2026-W44 — top 5 of 23 above threshold

2026-10-26 to 2026-11-01 (UTC).

## 1. aws — Increased invoke error rates

- URL: https://health.aws.amazon.com/health/status
- Window: 2026-10-30 01:28 to not yet (3d 15h 11m (ongoing))
- Impact: unknown, status: investigating
- Components: apigateway-us-west-2

  - `2026-10-30 01:28` investigating — Between 7:27 AM and 8:47 AM PDT, and between 10:47 AM and 4:15 PM PDT, we experienced intermittent invoke error rates and latencies for HTTP APIs in the US-WEST-2 Region. During this time, customers would have observed increased latency or errors when invoking HTTP APIs, with elevated errors occurring between 3:12 PM and 3:45 PM PDT. This represented the period of highest customer impact with the majority of affected customers experiencing errors or increased latencies during this window. Majority of retry attempts for affected requests were succeeding. We communicated to impacted customers via the Your Account view of the AWS Health Dashboard during both impact windows. We identified the root cause to be an issue with an internal subsystem and by 3:45 PM we applied the necessary mitigatio…

## 2. cloudflare — Network Performance Issues in Ashburn, IAD

- URL: https://www.cloudflarestatus.com/incidents/snth5ero95rz
- Window: 2026-10-28 04:39 to 2026-10-28 22:18 UTC (17h 39m)
- Impact: minor, status: resolved
- Components: Ashburn, United States - (IAD)

  - `2026-10-28 04:39` investigating — Cloudflare is investigating network performance issues affecting our Ashburn (IAD) data center. Traffic routed through this location may experience elevated latency and/or 5xx errors. We are working to analyze and mitigate this problem. More updates to follow shortly.
  - `2026-10-28 05:05` monitoring — We are continuing to monitor for any further issues.
  - `2026-10-28 22:18` resolved — This incident has been resolved.

## 3. cloudflare — Increased Errors for Durable Objects in the Eastern North America region

- URL: https://www.cloudflarestatus.com/incidents/klub12zy71b4
- Window: 2026-10-28 00:17 to 2026-10-28 03:28 UTC (3h 11m)
- Impact: minor, status: resolved
- Components: Durable Objects, D1, Queues, Workflows

  - `2026-10-28 00:17` investigating — Cloudflare is aware of, and investigating an issue with Durable Objects and downstream services which potentially impacts multiple customers in the Eastern North America region. Durable Objects are experiencing an elevated level of overload errors in the aforementioned region. This may also affect downstream services such as D1, Queues and Workflows etc. We are currently investigating this issue.
  - `2026-10-28 01:22` monitoring — A fix has been implemented and we are monitoring the results.
  - `2026-10-28 03:28` resolved — This incident has been resolved.

## 4. cloudflare — Increased Durable Objects, D1 and Workflows Error Rates

- URL: https://www.cloudflarestatus.com/incidents/r4mz367vddi7
- Window: 2026-10-29 14:16 to 2026-10-29 16:33 UTC (2h 16m)
- Impact: minor, status: resolved
- Components: Durable Objects, Artifacts, D1, Workers Assets, Workflows

  - `2026-10-29 14:16` investigating — Cloudflare is investigating an issue where customers are experiencing elevated errors for operations on Increased Durable Objects, D1 and Workflows.
  - `2026-10-29 16:33` monitoring — A fix has been implemented and we are monitoring the results.
  - `2026-10-29 16:33` resolved — This incident has been resolved.

## 5. datadog — Metrics query are disrupted

- URL: https://stspg.io/d4klp5znv9c7
- Window: 2026-10-28 16:10 to 2026-10-28 18:39 UTC (2h 29m)
- Impact: major, status: resolved
- Components: APM, Bits AI, Cloud Cost Management, Cloud Network Monitoring, Metrics and Infra Monitoring, Monitors, Product Analytics, RUM

  - `2026-10-28 16:10` investigating — We are investigating metrics issues. As a result, some users might not be able to load dashboard or query metrics. Please note that data processing and alerts are not affected by this incident.
  - `2026-10-28 16:40` investigating — We are investigating the issue.
  - `2026-10-28 16:56` investigating — We are investigating the issue.
  - `2026-10-28 17:26` identified — We identified the issue.
  - `2026-10-28 18:18` monitoring — We are monitoring the issue.
  - `2026-10-28 18:23` monitoring — We are monitoring the issue.
  - `2026-10-28 18:39` resolved — The issue is now resolved.

## Also above threshold (18)

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

