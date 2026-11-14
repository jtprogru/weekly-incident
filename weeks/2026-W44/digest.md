# Week 2026-W44

2026-10-26 to 2026-11-01 (UTC). Generated 2026-11-14 12:36.

52 canonical incidents, 24 above threshold. 2 of 14 sources unavailable: datadog, yandexcloud.

52 source records (60 saved revisions) → 52 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | cloudflare | [Cloudflare Pages Upload Issue](https://www.cloudflarestatus.com/incidents/me060q7dbayp) | 2026-10-26 00:00 | 9d 6h 35m | minor | Dashboard, Pages | 48078 = 13355 × 2 × 1.8 |
| 2 | aws | [Increased invoke error rates](https://health.aws.amazon.com/health/status) | 2026-10-30 01:28 | 15d 11h 8m (ongoing) | unknown | apigateway-us-west-2 | 44536 = 22268 × 1 × 2.0 |
| 3 | cloudflare | [Network Performance Issues in Ashburn, IAD](https://www.cloudflarestatus.com/incidents/snth5ero95rz) | 2026-10-28 04:39 | 17h 39m | minor | Ashburn, United States - (IAD) | 1906 = 1059 × 1 × 1.8 |
| 4 | cloudflare | [Increased Errors for Durable Objects in the Eastern North America region](https://www.cloudflarestatus.com/incidents/klub12zy71b4) | 2026-10-28 00:17 | 3h 11m | minor | D1, Durable Objects, Queues, Workflows | 1375 = 191 × 4 × 1.8 |
| 5 | cloudflare | [Increased Durable Objects, D1 and Workflows Error Rates](https://www.cloudflarestatus.com/incidents/r4mz367vddi7) | 2026-10-29 14:16 | 2h 16m | minor | Artifacts, D1, Durable Objects, Workers Assets, Workflows | 1224 = 136 × 5 × 1.8 |
| 6 | datadog | [Metrics query are disrupted](https://stspg.io/d4klp5znv9c7) | 2026-10-28 16:10 | 2h 29m | major | APM, Bits AI, Cloud Cost Management, Cloud Network Monitoring, Metrics and Infra Monitoring, Monitors, Product Analytics, RUM | 1192 = 149 × 8 × 1.0 |
| 7 | cloudflare | [Increased Worker services errors](https://www.cloudflarestatus.com/incidents/6yfpsos7ksp3) | 2026-10-29 09:22 | 2h 26m | minor | Containers, D1, Workers Assets, Workflows | 1051 = 146 × 4 × 1.8 |
| 8 | anthropic | [Elevated errors affecting the Google Calendar connector](https://stspg.io/75qq31h8672d) | 2026-10-28 16:19 | 3h 46m | minor | Claude Code, Claude Cowork, claude.ai | 1017 = 226 × 3 × 1.5 |
| 9 | stripe | [Elevated Issuing errors from 17:13 - 20:41 UTC](https://stspg.io/kj9ch2nd498j) | 2026-10-28 17:39 | 4h 36m | minor | Banking-as-a-service, Stripe API | 662 = 276 × 2 × 1.2 |
| 10 | github | [Incident with Actions, Pages, and Webhooks](https://stspg.io/5cd2m4n59spv) | 2026-10-30 17:50 | 1h 32m | major | Actions, Pages, Webhooks | 414 = 92 × 3 × 1.5 |
| 11 | cloudflare | [Delayed, missed, and duplicated Workers Cron Triggers](https://www.cloudflarestatus.com/incidents/97y1vy5nsjrn) | 2026-10-26 21:48 | 3h 26m | minor | Workers | 371 = 206 × 1 × 1.8 |
| 12 | openai | [Elevated errors for GPT Image 2 API image edits](https://status.openai.com/incidents/01M66KSDH67V5BQ73PE8AD2W8S) | 2026-10-30 14:35 | 4h 1m | minor | — | 362 = 241 × 1 × 1.5 |
| 13 | cloudflare | [ERR_TIMED_OUT and ERR_EMPTY_RESPONSE being observed by Cloudflare Customers](https://www.cloudflarestatus.com/incidents/dk2obcyw2aew) | 2026-10-27 12:28 | 2h 43m | major | — | 293 = 163 × 1 × 1.8 |
| 14 | grafana | [Degraded Cloud Integrations](https://stspg.io/d7mh6vy0g5jh) | 2026-10-27 15:05 | 4h 41m | minor | Integrations | 281 = 281 × 1 × 1.0 |
| 15 | anthropic | [Elevated errors for Claude Sonnet 5.5](https://stspg.io/rql9h590npxd) | 2026-10-27 09:14 | 42m | major | Claude API (api.anthropic.com), Claude Code, Claude Cowork, claude.ai | 252 = 42 × 4 × 1.5 |
| 16 | anthropic | [Elevated errors for Claude Sonnet 5.5](https://stspg.io/39x7qh3sd7bd) | 2026-10-27 15:09 | 41m | major | Claude API (api.anthropic.com), Claude Code, Claude Cowork, claude.ai | 246 = 41 × 4 × 1.5 |
| 17 | github | [Disruption with Copilot Cloud services](https://stspg.io/22p3vp78z20q) | 2026-10-29 10:19 | 2h 29m | minor | — | 224 = 149 × 1 × 1.5 |
| 18 | github | [Incident with Dependabot](https://stspg.io/5bh6qnzn76tk) | 2026-10-30 20:59 | 2h 29m | minor | — | 224 = 149 × 1 × 1.5 |
| 19 | digitalocean | [Kubernetes Cluster Management Operations](https://stspg.io/b16t5y07v5l4) | 2026-10-27 22:14 | 3h 40m | major | Global | 176 = 220 × 1 × 0.8 |
| 20 | datadog | [Elevated Error Rates for Workflow Automation, App Builder, Bits Code, Bits Agent Builder](https://stspg.io/rq69kpmmw7h6) | 2026-10-26 17:26 | 56m | major | App Builder, Bits AI, Workflow Automation | 168 = 56 × 3 × 1.0 |
| 21 | stripe | [Elevated Instant Payouts failures](https://stspg.io/3gk09h1d44ts) | 2026-11-01 12:31 | 2h 5m | none | Stripe core components | 150 = 125 × 1 × 1.2 |
| 22 | github | [Disruption with some Copilot services](https://stspg.io/bwcyq078jzms) | 2026-10-28 14:12 | 1h 15m | critical | Copilot | 112 = 75 × 1 × 1.5 |
| 23 | github | [Disruption with GitHub services](https://stspg.io/6p81rf34z4nc) | 2026-10-30 14:31 | 31m | critical | Issues, Pages | 93 = 31 × 2 × 1.5 |
| 24 | grafana | [Frontend Observability Outage in prod-eu-west-2](https://stspg.io/dgmytp3hyf4l) | 2026-10-27 15:56 | 27m | critical | AWS Germany - prod-eu-west-2 | 27 = 27 × 1 × 1.0 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. cloudflare — Cloudflare Pages Upload Issue

- URL: https://www.cloudflarestatus.com/incidents/me060q7dbayp
- Started: 2026-10-26 00:00 UTC
- Resolved: 2026-11-04 06:35 UTC
- Duration: 9d 6h 35m
- Impact: minor, status: resolved
- Components: Dashboard, Pages
- Canonical incident: [ci-ac1206f83d1f52f837cc0aac](#ci-ac1206f83d1f52f837cc0aac)

- `2026-11-04 07:11` **resolved** — Between 2026/10/26 - 2026/11/4, some customers trying to upload static Pages assets from the Cloudflare Dashboard may have experienced issues creating a Pages project.

### 2. aws — Increased invoke error rates

- URL: https://health.aws.amazon.com/health/status
- Started: 2026-10-30 01:28 UTC
- Resolved: not yet
- Duration: 15d 11h 8m (ongoing)
- Impact: unknown, status: investigating
- Components: apigateway-us-west-2
- Canonical incident: [ci-d86c52c83cfb265a4c4667b8](#ci-d86c52c83cfb265a4c4667b8)

- `2026-10-30 01:28` **investigating** — Between 7:27 AM and 8:47 AM PDT, and between 10:47 AM and 4:15 PM PDT, we experienced intermittent invoke error rates and latencies for HTTP APIs in the US-WEST-2 Region. During this time, customers would have observed increased latency or errors when invoking HTTP APIs, with elevated errors occurring between 3:12 PM and 3:45 PM PDT. This represented the period of highest customer impact with the majority of affected customers experiencing errors or increased latencies during this window. Majority of retry attempts for affected requests were succeeding. We communicated to impacted customers via the Your Account view of the AWS Health Dashboard during both impact windows. We identified the root cause to be an issue with an internal subsystem and by 3:45 PM we applied the necessary mitigation actions to remediate the issue. By 4:00 PM, we saw initial signs of recovery and continued to monitor for sustained recovery. We confirmed full recovery at 4:15 PM. We recommend that customers retry failed requests. The issue is resolved and the service is operating normally.

### 3. cloudflare — Network Performance Issues in Ashburn, IAD

- URL: https://www.cloudflarestatus.com/incidents/snth5ero95rz
- Started: 2026-10-28 04:39 UTC
- Resolved: 2026-10-28 22:18 UTC
- Duration: 17h 39m
- Impact: minor, status: resolved
- Components: Ashburn, United States - (IAD)
- Canonical incident: [ci-0a026ea52a6b18844867b0a1](#ci-0a026ea52a6b18844867b0a1)

- `2026-10-28 04:39` **investigating** — Cloudflare is investigating network performance issues affecting our Ashburn (IAD) data center. Traffic routed through this location may experience elevated latency and/or 5xx errors. We are working to analyze and mitigate this problem. More updates to follow shortly.
- `2026-10-28 05:05` **monitoring** — We are continuing to monitor for any further issues.
- `2026-10-28 22:18` **resolved** — This incident has been resolved.

### 4. cloudflare — Increased Errors for Durable Objects in the Eastern North America region

- URL: https://www.cloudflarestatus.com/incidents/klub12zy71b4
- Started: 2026-10-28 00:17 UTC
- Resolved: 2026-10-28 03:28 UTC
- Duration: 3h 11m
- Impact: minor, status: resolved
- Components: D1, Durable Objects, Queues, Workflows
- Canonical incident: [ci-05e287e0555cc3e790b274a8](#ci-05e287e0555cc3e790b274a8)

- `2026-10-28 00:17` **investigating** — Cloudflare is aware of, and investigating an issue with Durable Objects and downstream services which potentially impacts multiple customers in the Eastern North America region. Durable Objects are experiencing an elevated level of overload errors in the aforementioned region. This may also affect downstream services such as D1, Queues and Workflows etc. We are currently investigating this issue.
- `2026-10-28 01:22` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-10-28 03:28` **resolved** — This incident has been resolved.

### 5. cloudflare — Increased Durable Objects, D1 and Workflows Error Rates

- URL: https://www.cloudflarestatus.com/incidents/r4mz367vddi7
- Started: 2026-10-29 14:16 UTC
- Resolved: 2026-10-29 16:33 UTC
- Duration: 2h 16m
- Impact: minor, status: resolved
- Components: Artifacts, D1, Durable Objects, Workers Assets, Workflows
- Canonical incident: [ci-fdbeb462c66148cf9f3c0500](#ci-fdbeb462c66148cf9f3c0500)

- `2026-10-29 14:16` **investigating** — Cloudflare is investigating an issue where customers are experiencing elevated errors for operations on Increased Durable Objects, D1 and Workflows.
- `2026-10-29 16:33` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-10-29 16:33` **resolved** — This incident has been resolved.

### 6. datadog — Metrics query are disrupted

- URL: https://stspg.io/d4klp5znv9c7
- Started: 2026-10-28 16:10 UTC
- Resolved: 2026-10-28 18:39 UTC
- Duration: 2h 29m
- Impact: major, status: resolved
- Components: APM, Bits AI, Cloud Cost Management, Cloud Network Monitoring, Metrics and Infra Monitoring, Monitors, Product Analytics, RUM
- Canonical incident: [ci-c387ca17fa2897162aa61ecd](#ci-c387ca17fa2897162aa61ecd)

- `2026-10-28 16:10` **investigating** — We are investigating metrics issues. As a result, some users might not be able to load dashboard or query metrics. Please note that data processing and alerts are not affected by this incident.
- `2026-10-28 16:40` **investigating** — We are investigating the issue.
- `2026-10-28 16:56` **investigating** — We are investigating the issue.
- `2026-10-28 17:26` **identified** — We identified the issue.
- `2026-10-28 18:18` **monitoring** — We are monitoring the issue.
- `2026-10-28 18:23` **monitoring** — We are monitoring the issue.
- `2026-10-28 18:39` **resolved** — The issue is now resolved.

### 7. cloudflare — Increased Worker services errors

- URL: https://www.cloudflarestatus.com/incidents/6yfpsos7ksp3
- Started: 2026-10-29 09:22 UTC
- Resolved: 2026-10-29 11:49 UTC
- Duration: 2h 26m
- Impact: minor, status: resolved
- Components: Containers, D1, Workers Assets, Workflows
- Canonical incident: [ci-33786c310a2330240a6ba8bd](#ci-33786c310a2330240a6ba8bd)

- `2026-10-29 09:22` **investigating** — Cloudflare has identified an issue where customers are experiencing elevated 500 errors impacting Durable Objects, D1, Workflows, Workers Assets uploads and Cloudchamber.
- `2026-10-29 10:51` **identified** — The issue has been identified and a fix is being implemented.
- `2026-10-29 11:37` **monitoring** — We are continuing to monitor for any further issues.
- `2026-10-29 11:49` **resolved** — This incident has been resolved.

### 8. anthropic — Elevated errors affecting the Google Calendar connector

- URL: https://stspg.io/75qq31h8672d
- Started: 2026-10-28 16:19 UTC
- Resolved: 2026-10-28 20:05 UTC
- Duration: 3h 46m
- Impact: minor, status: resolved
- Components: Claude Code, Claude Cowork, claude.ai
- Canonical incident: [ci-d0d7fba77aa64f80edfa9e7d](#ci-d0d7fba77aa64f80edfa9e7d)

- `2026-10-28 16:19` **investigating** — We are investigating elevated errors affecting the Google Calendar connector. Some requests to list or search calendar events have been failing since October 26. We will provide an update as soon as possible.
- `2026-10-28 17:59` **investigating** — We are continuing to investigate elevated errors affecting the Google Calendar connector, and are working with our upstream provider. We will provide another update as soon as we have more information.
- `2026-10-28 19:47` **monitoring** — We are seeing recovery for errors affecting the Google Calendar connector, and error rates have returned to normal levels. We are continuing to monitor with our upstream provider.
- `2026-10-28 20:05` **resolved** — The issue affecting the Google Calendar connector on claude.ai, Claude Code, and Claude Cowork has been resolved. Impact occurred from 04:07 PT / 11:07 UTC on October 26 to 11:52 PT / 18:52 UTC on October 28.

### 9. stripe — Elevated Issuing errors from 17:13 - 20:41 UTC

- URL: https://stspg.io/kj9ch2nd498j
- Started: 2026-10-28 17:39 UTC
- Resolved: 2026-10-28 22:15 UTC
- Duration: 4h 36m
- Impact: minor, status: resolved
- Components: Banking-as-a-service, Stripe API
- Canonical incident: [ci-9fbf6d62a2cbd6b5b53a4967](#ci-9fbf6d62a2cbd6b5b53a4967)

- `2026-10-28 17:39` **investigating** — We’re currently investigating elevated API errors beginning at 17:13 UTC. We’ll post an update soon.
- `2026-10-28 18:29` **identified** — We’re currently seeing elevated errors related to newly created Issuing cards beginning at 17:13 UTC. We’ve identified the cause, and are working on a fix.
- `2026-10-28 20:13` **identified** — Beginning at 17:13 UTC, we are experiencing elevated error rates for API requests involving Issuing cards that were created or updated between 17:13–17:54 UTC. We are currently working to correct the affected card data, and will update you soon.
- `2026-10-28 21:08` **identified** — From 17:15 - 20:41 UTC, Stripe Issuing cards that were created or updated between 17:15–17:54 UTC experienced elevated error rates for API requests. The majority of errors subsided by 19:32 UTC, with full resolution occurring at 20:41 UTC. We have corrected the affected card data, and this issue is now resolved.
- `2026-10-28 22:15` **resolved** — From 17:15 - 20:41 UTC, Stripe Issuing cards that were created or updated between 17:15–17:54 UTC experienced elevated error rates for API requests. The majority of errors subsided by 19:32 UTC, with full resolution occurring at 20:41 UTC. We have corrected the affected card data, and this issue is now resolved.

### 10. github — Incident with Actions, Pages, and Webhooks

- URL: https://stspg.io/5cd2m4n59spv
- Started: 2026-10-30 17:50 UTC
- Resolved: 2026-10-30 19:23 UTC
- Duration: 1h 32m
- Impact: major, status: resolved
- Components: Actions, Pages, Webhooks
- Canonical incident: [ci-7d4fd51d203cce058e9451e9](#ci-7d4fd51d203cce058e9451e9)

- `2026-10-30 17:50` **investigating** — We are investigating reports of degraded performance for Actions and Webhooks
- `2026-10-30 18:00` **investigating** — Pages is experiencing degraded performance. We are continuing to investigate.
- `2026-10-30 18:06` **investigating** — We continue to investigate delays affecting Actions workflows, Pages deployments, and webhook events, along with errors affecting code search and live updates. We have taken mitigation steps. Actions processing and Pages deployments are showing signs of recovery, but impact has not fully resolved. We are monitoring recovery and investigating the underlying cause.
- `2026-10-30 18:32` **investigating** — All affected services have recovered following mitigation. We continue to investigate root cause.
- `2026-10-30 18:32` **monitoring** — The degradation affecting Actions, Pages and Webhooks has been mitigated. We are monitoring to ensure stability.
- `2026-10-30 19:23` **resolved** — On October 30, 2026, from 17:50 to 18:32 UTC, Actions, Pages, and Webhooks experienced degraded performance. Actions workflow processing, Pages deployments, and webhook deliveries were delayed, and code search and live updates returned errors. Some Actions queues took additional time to drain after mitigation, and the incident was resolved at 19:23 UTC. <br /><br />The impact was triggered by maintenance on a cloud network gateway connecting an Actions service in the GitHub data center to Redis clusters in the Central US Azure region. The service's Redis client retried aggressively and redialed those connections. That produced a connection storm that saturated the gateway, with about 90% of connection attempts failing. Other services sharing the same network path were degraded as well. <br /><br />We mitigated the incident by removing the affected resources from rotation, which ended the connection storm and allowed the gateway to recover. We have already reduced the number of connection retries in the impacted service and are further improving retry behavior with jittered backoff, shorter connect timeouts, and retry budgets. We are also working with our infrastructure team to make GitHub more resilient to gateway maintenance.

### 11. cloudflare — Delayed, missed, and duplicated Workers Cron Triggers

- URL: https://www.cloudflarestatus.com/incidents/97y1vy5nsjrn
- Started: 2026-10-26 21:48 UTC
- Resolved: 2026-10-27 01:15 UTC
- Duration: 3h 26m
- Impact: minor, status: resolved
- Components: Workers
- Canonical incident: [ci-91adfb23683ebf8363d5726a](#ci-91adfb23683ebf8363d5726a)

- `2026-10-26 21:48` **investigating** — We are investigating reports that some Workers Cron Triggers are running late, being skipped, or running more than once. Affected customers may see scheduled Workers invocations that do not run on time or run more than once. We will post an update as we learn more.
- `2026-10-27 00:55` **monitoring** — We are continuing to monitor for any further issues.
- `2026-10-27 01:15` **resolved** — This incident has been resolved.

### 12. openai — Elevated errors for GPT Image 2 API image edits

- URL: https://status.openai.com/incidents/01M66KSDH67V5BQ73PE8AD2W8S
- Started: 2026-10-30 14:35 UTC
- Resolved: 2026-10-30 18:36 UTC
- Duration: 4h 1m
- Impact: minor, status: resolved
- Components: —
- Canonical incident: [ci-6e2d52ee775196cb2782fba0](#ci-6e2d52ee775196cb2782fba0)

- `2026-10-30 14:35` **identified** — We are experiencing elevated errors for some GPT Image 2 API image edit requests. We have identified the issue and are working on a fix
- `2026-10-30 16:06` **monitoring** — We’ve applied a mitigation and are monitoring the affected services to confirm recovery.
- `2026-10-30 18:36` **resolved** — All affected services are operating normally again. We’re sorry for the disruption, and we appreciate your patience.

### 13. cloudflare — ERR_TIMED_OUT and ERR_EMPTY_RESPONSE being observed by Cloudflare Customers

- URL: https://www.cloudflarestatus.com/incidents/dk2obcyw2aew
- Started: 2026-10-27 12:28 UTC
- Resolved: 2026-10-27 15:12 UTC
- Duration: 2h 43m
- Impact: major, status: resolved
- Components: —
- Canonical incident: [ci-58d2f990a87d08aa530c55b1](#ci-58d2f990a87d08aa530c55b1)

- `2026-10-27 12:28` **identified** — Cloudflare has identified an issue affecting customers whose traffic passes through an external vendor firewall. Customers may observe: ERR_TIMED_OUT on standard HTTPS websites ERR_EMPTY_RESPONSE on non-standard HTTPS ports (e.g., 8443) We're investigating to confirm the root cause and possible fixes.
- `2026-10-27 13:20` **identified** — Cloudflare has identified an issue affecting customers whose traffic passes through an external vendor firewall. Customers may observe: ERR_TIMED_OUT on standard HTTPS websites ERR_EMPTY_RESPONSE on non-standard HTTPS ports (e.g., 8443) We're investigating to confirm the root cause and possible fixes.
- `2026-10-27 14:38` **identified** — Cloudflare has identified an issue affecting customers whose traffic passes through an external vendor firewall. Customers may observe: ERR_TIMED_OUT on standard HTTPS websites ERR_EMPTY_RESPONSE on non-standard HTTPS ports (e.g., 8443) We're investigating to confirm the root cause and possible fixes.
- `2026-10-27 14:52` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-10-27 15:12` **resolved** — This incident has been resolved.

### 14. grafana — Degraded Cloud Integrations

- URL: https://stspg.io/d7mh6vy0g5jh
- Started: 2026-10-27 15:05 UTC
- Resolved: 2026-10-27 19:47 UTC
- Duration: 4h 41m
- Impact: minor, status: resolved
- Components: Integrations
- Canonical incident: [ci-61e652f02042e7dd81d50b04](#ci-61e652f02042e7dd81d50b04)

- `2026-10-27 15:05` **investigating** — We are currently investigating an issue impacting Cloud Integrations in prod-eu-west-3. We will provide another update when we have more to share.
- `2026-10-27 16:37` **investigating** — We are continuing to investigate this issue.
- `2026-10-27 17:16` **monitoring** — Things are no longer degraded. We are currently monitoring the issue
- `2026-10-27 19:47` **resolved** — This incident has been resolved.

### 15. anthropic — Elevated errors for Claude Sonnet 5.5

- URL: https://stspg.io/rql9h590npxd
- Started: 2026-10-27 09:14 UTC
- Resolved: 2026-10-27 09:56 UTC
- Duration: 42m
- Impact: major, status: resolved
- Components: Claude API (api.anthropic.com), Claude Code, Claude Cowork, claude.ai
- Canonical incident: [ci-786e978b5877fde448a171ad](#ci-786e978b5877fde448a171ad)

- `2026-10-27 09:14` **investigating** — We are investigating elevated errors on requests to Claude Sonnet 5.5. We will provide an update as soon as possible.
- `2026-10-27 09:38` **monitoring** — A fix has been deployed for the issue affecting Claude services and we are monitoring for recovery.
- `2026-10-27 09:56` **resolved** — The issue affecting Claude services has been resolved.

### 16. anthropic — Elevated errors for Claude Sonnet 5.5

- URL: https://stspg.io/39x7qh3sd7bd
- Started: 2026-10-27 15:09 UTC
- Resolved: 2026-10-27 15:50 UTC
- Duration: 41m
- Impact: major, status: resolved
- Components: Claude API (api.anthropic.com), Claude Code, Claude Cowork, claude.ai
- Canonical incident: [ci-abaadece8ac24fe3caa5b9e5](#ci-abaadece8ac24fe3caa5b9e5)

- `2026-10-27 15:09` **identified** — We are investigating elevated errors on requests to Claude Sonnet 5.5. We will provide an update as soon as possible.
- `2026-10-27 15:50` **resolved** — The issue affecting claude.ai, the Claude API, Claude Code, and Claude Cowork has been resolved. Impact occurred from 07:42 PT / 14:42 UTC to 08:17 PT / 15:17 UTC.

### 17. github — Disruption with Copilot Cloud services

- URL: https://stspg.io/22p3vp78z20q
- Started: 2026-10-29 10:19 UTC
- Resolved: 2026-10-29 12:49 UTC
- Duration: 2h 29m
- Impact: minor, status: resolved
- Components: —
- Canonical incident: [ci-b4638db777542c16ec149ad2](#ci-b4638db777542c16ec149ad2)

- `2026-10-29 10:19` **investigating** — We are investigating reports of impacted performance for some GitHub services.
- `2026-10-29 10:22` **investigating** — We are experiencing increased failures and latency on our agent sessions service, including creating and listing sessions. This can impact cloud based agent sessions on web, mobile, VS Code, CLI and the GitHub Copilot app.
- `2026-10-29 11:01` **investigating** — We continue to experience failure rates of 1-2% and increased latency on our agent sessions service, including creating and listing sessions. This can impact cloud based agent sessions on web, mobile, VS Code, CLI and the GitHub Copilot app.
- `2026-10-29 11:37` **investigating** — We continue to experience failure rates of 1-2% and increased latency on our agent sessions service, including creating and listing sessions. This can impact cloud based agent sessions on web, mobile, VS Code, CLI and the GitHub Copilot app. We are actively applying mitigations and monitoring for recovery.
- `2026-10-29 12:33` **monitoring** — The degradation has been mitigated. We are monitoring to ensure stability.
- `2026-10-29 12:49` **resolved** — On October 29, 2026, between 05:50 UTC and 12:40 UTC, Copilot agent sessions on GitHub.com and in VS Code were degraded. Live session updates often failed to load, and on average 2.14% of requests to the agent sessions service failed, with a peak of 8.61%. For about 10 minutes around 09:20 UTC, up to 40% of requests to list or create agent sessions failed. Copilot Chat on GitHub.com showed degraded agent session features during this time. The issue occurred when a large increase in task-list requests from a client experiment overloaded the cache that supports the agent sessions service. We mitigated the impact by increasing cache capacity, reducing nonessential load on the cache, and disabling the client experiment. We are moving rate limiting to separate infrastructure, reducing client polling, and improving our alerting so we can detect and respond to this type of overload faster.

### 18. github — Incident with Dependabot

- URL: https://stspg.io/5bh6qnzn76tk
- Started: 2026-10-30 20:59 UTC
- Resolved: 2026-10-30 23:29 UTC
- Duration: 2h 29m
- Impact: minor, status: resolved
- Components: —
- Canonical incident: [ci-0800b0801d93c1a197b5bd5f](#ci-0800b0801d93c1a197b5bd5f)

- `2026-10-30 20:59` **investigating** — We are investigating reports of impacted performance for some GitHub services.
- `2026-10-30 21:00` **investigating** — We are investigating elevated errors when viewing some Dependabot alerts.
- `2026-10-30 21:44` **investigating** — We are continuing to investigate elevated errors when viewing some Dependabot alerts. We’ll provide more updates as we learn more.
- `2026-10-30 22:32` **investigating** — We have mitigated impact and are seeing recovery. We are monitoring for full recovery.
- `2026-10-30 22:32` **monitoring** — The degradation has been mitigated. We are monitoring to ensure stability.
- `2026-10-30 23:29` **resolved** — On October 30, 2026, between 20:15 and 21:56 UTC, an advisory-publication-driven fan-out caused an unexpected saturation of network connections to the Dependabot alerts datastore. During this period we observed a high rate of write failures on this datastore; users saw elevated failure rates when dismissing or resolving Dependabot alerts, as well as failures when linking Dependabot update pull requests to Dependabot security alerts. We also observed delays in creating and updating security alerts that extended through November 2, 2026.<br /><br />The unexpected fan-out was caused by a missing job queue throttle after the queues for the fan-out work were migrated as part of system maintenance. We have restored the fan-out throttle and further isolated and throttled these processes to guard against further disruptions.

### 19. digitalocean — Kubernetes Cluster Management Operations

- URL: https://stspg.io/b16t5y07v5l4
- Started: 2026-10-27 22:14 UTC
- Resolved: 2026-10-28 01:55 UTC
- Duration: 3h 40m
- Impact: major, status: resolved
- Components: Global
- Canonical incident: [ci-ee5529e9e1b6d84d30de2ee4](#ci-ee5529e9e1b6d84d30de2ee4)

- `2026-10-27 22:14` **investigating** — Our Engineering team is currently investigating an issue affecting Managed Kubernetes cluster operations. During this time, Customers may experience delays or failures when attempting to create, delete, scale, or upgrade clusters, as well as with other management tasks. Existing workloads running on Managed Kubernetes are not affected and will continue to operate normally. Our Engineering team is actively investigating and treating this with the highest priority. We will provide another update as soon as more information becomes available.
- `2026-10-28 00:01` **investigating** — Our Engineering team continues to investigate an issue affecting Managed Kubernetes cluster operations. Customers may still experience delays or failures when creating, deleting, scaling, or upgrading clusters, across all regions. Existing workloads running on Managed Kubernetes are not affected and will continue to operate normally. We have no new information to share at this time. Our Engineering team is actively working on this issue and treating it with the highest priority. We will provide another update within the next hour, or sooner if more information becomes available.
- `2026-10-28 01:34` **monitoring** — Our Engineering team has implemented a fix for the issue affecting Managed Kubernetes cluster operations across all regions. Customers should now be able to create, delete, scale, and upgrade their clusters as expected. We are monitoring the situation to ensure all cluster management operations are functioning normally and the service remains stable. We will post another update once we confirm the issue is fully resolved.
- `2026-10-28 01:55` **resolved** — We have resolved the issue affecting Managed Kubernetes cluster operations across all regions. Cluster creation, deletion, scaling, and upgrade operations have returned to normal. We have monitored the environment and confirmed that services are operating normally. If you continue to experience any issues, please contact our Support team by opening a ticket here: https://cloudsupport.digitalocean.com/s/

### 20. datadog — Elevated Error Rates for Workflow Automation, App Builder, Bits Code, Bits Agent Builder

- URL: https://stspg.io/rq69kpmmw7h6
- Started: 2026-10-26 17:26 UTC
- Resolved: 2026-10-26 18:23 UTC
- Duration: 56m
- Impact: major, status: resolved
- Components: App Builder, Bits AI, Workflow Automation
- Canonical incident: [ci-f7298a1210434aa48a56af2a](#ci-f7298a1210434aa48a56af2a)

- `2026-10-26 17:26` **investigating** — We are currently investigating an issue affecting Workflow Automation, App Builder, Bits Code, and Agent Builder in our US1 datacenter. During this time, some users may experience: - Workflows containing JavaScript or Python actions failing to run, including workflows used for incident and on-call automation - App Builder apps that use JavaScript or Python failing, or parts of those apps not working as expected - Bits Code sessions failing to start or complete - Agent Builder agents with Bits Code enabled being unable to start code sessions or use bash tools Our engineering teams are actively working to resolve the issue, and we’ll follow up with more information shortly.
- `2026-10-26 17:30` **identified** — We have identified the underlying issue and are working on a fix.
- `2026-10-26 17:31` **identified** — We have identified the underlying issue and are working on a fix.
- `2026-10-26 17:40` **monitoring** — We have deployed a fix and we are monitoring the results. We will provide another update once the issue is fully resolved.
- `2026-10-26 18:23` **resolved** — The issue is now resolved.

### 21. stripe — Elevated Instant Payouts failures

- URL: https://stspg.io/3gk09h1d44ts
- Started: 2026-11-01 12:31 UTC
- Resolved: 2026-11-01 14:36 UTC
- Duration: 2h 5m
- Impact: none, status: resolved
- Components: Stripe core components
- Canonical incident: [ci-89fc61884dfc52f4f3fbce32](#ci-89fc61884dfc52f4f3fbce32)

- `2026-11-01 12:31` **investigating** — We're currently investigating elevated failure rates affecting Instant Payouts to cards in some countries. As of November 1, 2026 at 05:00 UTC, we've identified that a significant portion of failures are due to issuer-bank maintenance affecting payment networks. We're monitoring the situation closely and will provide further updates as the maintenance window progresses.
- `2026-11-01 12:50` **identified** — Instant Payouts to cards are experiencing elevated failure rates mainly in Canada and US. Since 05:00 UTC, a significant portion of these failures have been attributed to issuer-bank maintenance affecting payment networks. We are monitoring the situation and will provide updates as the maintenance window progresses.
- `2026-11-01 14:36` **resolved** — From 05:00–14:00 UTC on November 1, 2026, Instant Payouts to cards experienced elevated failure rates due to issuer-bank maintenance affecting payment networks, primarily impacting the US and Canada. This issue is now resolved, and payment processing has returned to normal.

### 22. github — Disruption with some Copilot services

- URL: https://stspg.io/bwcyq078jzms
- Started: 2026-10-28 14:12 UTC
- Resolved: 2026-10-28 15:28 UTC
- Duration: 1h 15m
- Impact: critical, status: resolved
- Components: Copilot
- Canonical incident: [ci-4ad2efd1b0495ea9aee3cc5f](#ci-4ad2efd1b0495ea9aee3cc5f)

- `2026-10-28 14:12` **investigating** — We are investigating reports of impacted performance for some GitHub services.
- `2026-10-28 14:15` **investigating** — We are experiencing issues with upstream model providers resulting in a degraded experience across Copilot services. The team is actively investigating the issue.
- `2026-10-28 14:33` **investigating** — We are experiencing degraded availability for the GPT-5.3-Codex, GPT-5.6 Luna, GPT-6.1 Sol, GPT-5.6 Terra, GPT-6 Luna, and MAI Code 1.1 Flash models in Copilot products and IDE surfaces. This is due to an issue with an upstream model provider. While we work with them to resolve the issue we recommend choosing another model to continue using Copilot.
- `2026-10-28 14:55` **investigating** — We are experiencing degraded availability for the GPT-5.3-Codex, GPT-5.6 Luna, GPT-6.1 Sol, GPT-5.6 Terra, GPT-6 Luna, and MAI Code 1.1 Flash models in Copilot products and IDE surfaces. We have identified the root cause and working to apply a mitigation. While we work to resolve the issue we recommend choosing another model to continue using Copilot.
- `2026-10-28 15:15` **investigating** — We have applied the mitigation and are seeing signs of recovery. We will continue to monitor and expect full resolution soon.
- `2026-10-28 15:24` **investigating** — Copilot is experiencing degraded performance. We are continuing to investigate.
- `2026-10-28 15:26` **monitoring** — The degradation affecting Copilot has been mitigated. We are monitoring to ensure stability.
- `2026-10-28 15:28` **resolved** — On October 28, 2026, between 13:52 and 15:24 UTC, Copilot users saw errors when using several GPT-family models and the MAI Code 1.1 Flash model. Affected experiences included Copilot Chat in IDEs and on GitHub.com, the Copilot CLI, Copilot Code Review, the Copilot Coding Agent, Code Scanning Autofix, and Code Quality. About 9.73% of requests to the affected models failed, with a peak of about 36.97% at the start of the incident.<br /><br />The errors happened because a credential that we use to authenticate with an upstream model provider became invalid. We mitigated the incident by replacing the credential and redeploying the service. We are strengthening how we manage credentials for upstream providers to prevent this type of failure, returning retryable errors when an upstream provider rejects a request, and improving our monitoring so that we detect and mitigate this type of problem faster in the future.

### 23. github — Disruption with GitHub services

- URL: https://stspg.io/6p81rf34z4nc
- Started: 2026-10-30 14:31 UTC
- Resolved: 2026-10-30 15:02 UTC
- Duration: 31m
- Impact: critical, status: resolved
- Components: Issues, Pages
- Canonical incident: [ci-301cc7d4f3fd40d9227fc6c8](#ci-301cc7d4f3fd40d9227fc6c8)

- `2026-10-30 14:31` **investigating** — We are investigating reports of impacted performance for some GitHub services.
- `2026-10-30 14:35` **investigating** — We are investigating degraded performance across multiple services. We will continue to keep users updated on progress towards mitigation.
- `2026-10-30 14:36` **investigating** — Issues is experiencing degraded performance. We are continuing to investigate.
- `2026-10-30 14:39` **investigating** — Pages is experiencing degraded performance. We are continuing to investigate.
- `2026-10-30 14:45` **investigating** — Issues is operating normally.
- `2026-10-30 14:48` **investigating** — A temporarily overloaded component triggered a burst of errors between 14:21 and 14:30 UTC across numerous services. Component health has returned to normal and we continue to monitor while investigating the root cause.<br />
- `2026-10-30 14:49` **monitoring** — The degradation affecting Pages has been mitigated. We are monitoring to ensure stability.
- `2026-10-30 15:02` **resolved** — This incident has been resolved. Thank you for your patience and understanding as we addressed this issue. A detailed root cause analysis will be shared as soon as it is available.

### 24. grafana — Frontend Observability Outage in prod-eu-west-2

- URL: https://stspg.io/dgmytp3hyf4l
- Started: 2026-10-27 15:56 UTC
- Resolved: 2026-10-27 16:24 UTC
- Duration: 27m
- Impact: critical, status: resolved
- Components: AWS Germany - prod-eu-west-2
- Canonical incident: [ci-61385bd384bb520ece77d186](#ci-61385bd384bb520ece77d186)

- `2026-10-27 15:56` **investigating** — We are currently investigating an issue impacting Frontend Observability since 15:20 UTC. We will provide another update as soon as we have more to share.
- `2026-10-27 16:24` **resolved** — This incident has been resolved.

## Below threshold (28)

- `cloudflare` [Increased Errors for Durable Objects in the Eastern North America region](https://www.cloudflarestatus.com/incidents/lmmf23klsrxq) — 1h 25m
- `anthropic` [Elevated errors affecting the Gmail connector](https://stspg.io/j86c100s8tqm) — 1h 48m
- `cloudflare` [Increased Errors for Durable Objects and Downstream Services in Eastern North America (ENAM)](https://www.cloudflarestatus.com/incidents/10m2bxjztzwh) — 50m
- `cloudflare` [Increased HTTP 429 Rate Limit Errors on Cloudflare Dashboard](https://www.cloudflarestatus.com/incidents/k2cd9e1fyskp) — 1h 39m
- `cloudflare` [Increased errors and latency for CNI 2.0 traffic in LAX](https://www.cloudflarestatus.com/incidents/oco7f885d2v4) — 1h 32m
- `cloudflare` [Cache Purging Delays Affecting Requests Routed Through EEUR](https://www.cloudflarestatus.com/incidents/8s6pn9rvof3l) — 1h 3m
- `cloudflare` [R2 Elevated Latency in Eastern Europe](https://www.cloudflarestatus.com/incidents/kp28d3kzgft1) — 1h 53m
- `cloudflare` [Issues with Durable Objects availability in Eastern Europe](https://www.cloudflarestatus.com/incidents/bc7e93q3eigp) — 1h 41m
- `openai` [Elevated errors with search in Dots, ChatGPT Work, and certain ChatGPT thinking experiences](https://status.openai.com/incidents/01M6754SR14HCQ3E9FX426E2N7) — 1h 35m
- `cloudflare` [We are seeing some customers having issues with Workers Builds](https://www.cloudflarestatus.com/incidents/2eno8w1fk27j) — 1h 16m
- `cloudflare` [Network Performance Issues in Queretaro, MX](https://www.cloudflarestatus.com/incidents/jajp50huvxtz) — 1h 10m
- `cloudflare` [Increased Errors for Browser Run](https://www.cloudflarestatus.com/incidents/2xiv8fdhk0bh) — 1h 7m
- `openai` [6.1 Sol Degraded Model Performance](https://status.openai.com/incidents/01M5YY4A3TQETEQ53DVSK8SW10) — 1h 17m
- `cloudflare` [Cloudflare Stream performance issues impacting certain Live Broadcasts and On-Demand Uploads](https://www.cloudflarestatus.com/incidents/mbq512v9oe4s) — 51m
- `cloudflare` [Cloudflare Workers build delays](https://www.cloudflarestatus.com/incidents/dv2vbq9kbgxn) — 50m
- `cloudflare` [Elevated Cloudflare Access API Errors](https://www.cloudflarestatus.com/incidents/4u4hsqqnr0hk) — 44m
- `cloudflare` [Network Performance Issues in Chicago (ORD)](https://www.cloudflarestatus.com/incidents/s5io0kegzlsz) — 42m
- `cloudflare` [Issues with AI Security for Apps](https://www.cloudflarestatus.com/incidents/03gc4o049emd) — 38m
- `cloudflare` [Network Error Log Processing Unavailable](https://www.cloudflarestatus.com/incidents/ktq18tcoqoo2) — 37m
- `circleci` [Delays for gen2 dedicated IP range jobs](https://stspg.io/t4pc57wmjnql) — 1h 3m
- `grafana` [ingestion errors in mimir-prod-66](https://stspg.io/42ggj9rqyyjm) — 1h 0m
- `cloudflare` [Increased latency and error rate for R2 in ENAM](https://www.cloudflarestatus.com/incidents/h6fgl47h6iht) — 10m
- `github` [Incident with Actions, Issues, API Requests and Pull Requests](https://stspg.io/0622ty903b4d) — 8m
- `cloudflare` [Some customers may notice their Pages and/or Workers builds are intermittently failing.](https://www.cloudflarestatus.com/incidents/j3llzsq3m17a) — 13m
- `openai` [Elevated error rates for most hosted Code Reviews and Security Reviews on private GitHub repositories](https://status.openai.com/incidents/01M5Y60NQGGN514EB3943V94EN) — 29m
- `cloudflare` [Cache purge delays in Qingdao Shandong, China](https://www.cloudflarestatus.com/incidents/5tlm9r9j9rpo) — 16m
- `openai` [Elevated error rates for Ads POST requests](https://status.openai.com/incidents/01M5WPZTT0385GMEV1353ZC4AM) — 11m
- `github` [Elevated 404s for GitHub App s2s tokens in git/API](https://stspg.io/4z5lmy7gqjk1) — unknown

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-ac1206f83d1f52f837cc0aac"></a>

- Canonical ID: `ci-ac1206f83d1f52f837cc0aac` (1 source records).
  - [cloudflare/me060q7dbayp](../../data/2026/W44/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/me060q7dbayp
- Start time basis: `vendor_reported`.

<a id="ci-d86c52c83cfb265a4c4667b8"></a>

- Canonical ID: `ci-d86c52c83cfb265a4c4667b8` (1 source records).
  - [aws/0fd29a8bd6906892](../../data/2026/W44/aws.json); vendor URL: https://health.aws.amazon.com/health/status
- Start time basis: `first_observed_update`.
- Warning: AWS onset is the first observed message, not a verified impact start

<a id="ci-0a026ea52a6b18844867b0a1"></a>

- Canonical ID: `ci-0a026ea52a6b18844867b0a1` (1 source records).
  - [cloudflare/snth5ero95rz](../../data/2026/W44/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/snth5ero95rz
- Start time basis: `vendor_reported`.

<a id="ci-05e287e0555cc3e790b274a8"></a>

- Canonical ID: `ci-05e287e0555cc3e790b274a8` (1 source records).
  - [cloudflare/klub12zy71b4](../../data/2026/W44/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/klub12zy71b4
- Start time basis: `vendor_reported`.

<a id="ci-fdbeb462c66148cf9f3c0500"></a>

- Canonical ID: `ci-fdbeb462c66148cf9f3c0500` (1 source records).
  - [cloudflare/r4mz367vddi7](../../data/2026/W44/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/r4mz367vddi7
- Start time basis: `vendor_reported`.

<a id="ci-c387ca17fa2897162aa61ecd"></a>

- Canonical ID: `ci-c387ca17fa2897162aa61ecd` (1 source records).
  - [datadog/z0b1h0ylmhm5](../../data/2026/W44/datadog.json); vendor URL: https://stspg.io/d4klp5znv9c7
- Start time basis: `vendor_reported`.

<a id="ci-33786c310a2330240a6ba8bd"></a>

- Canonical ID: `ci-33786c310a2330240a6ba8bd` (1 source records).
  - [cloudflare/6yfpsos7ksp3](../../data/2026/W44/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/6yfpsos7ksp3
- Start time basis: `vendor_reported`.

<a id="ci-d0d7fba77aa64f80edfa9e7d"></a>

- Canonical ID: `ci-d0d7fba77aa64f80edfa9e7d` (1 source records).
  - [anthropic/phg1y2yh07n2](../../data/2026/W44/anthropic.json); vendor URL: https://stspg.io/75qq31h8672d
- Start time basis: `vendor_reported`.

<a id="ci-9fbf6d62a2cbd6b5b53a4967"></a>

- Canonical ID: `ci-9fbf6d62a2cbd6b5b53a4967` (1 source records).
  - [stripe/k0m9vtfvkcjz](../../data/2026/W44/stripe.json); vendor URL: https://stspg.io/kj9ch2nd498j
- Start time basis: `vendor_reported`.

<a id="ci-9a7f8151a3440a23eb9c6b16"></a>

- Canonical ID: `ci-9a7f8151a3440a23eb9c6b16` (1 source records).
  - [cloudflare/lmmf23klsrxq](../../data/2026/W44/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/lmmf23klsrxq
- Start time basis: `vendor_reported`.

<a id="ci-b6ed5824d9ed86afb0fcfee0"></a>

- Canonical ID: `ci-b6ed5824d9ed86afb0fcfee0` (1 source records).
  - [anthropic/9ftx658xtd5z](../../data/2026/W44/anthropic.json); vendor URL: https://stspg.io/j86c100s8tqm
- Start time basis: `vendor_reported`.

<a id="ci-7d4fd51d203cce058e9451e9"></a>

- Canonical ID: `ci-7d4fd51d203cce058e9451e9` (1 source records).
  - [github/pwxsjgnvrk6b](../../data/2026/W44/github.json); vendor URL: https://stspg.io/5cd2m4n59spv
- Start time basis: `vendor_reported`.

<a id="ci-91adfb23683ebf8363d5726a"></a>

- Canonical ID: `ci-91adfb23683ebf8363d5726a` (1 source records).
  - [cloudflare/97y1vy5nsjrn](../../data/2026/W44/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/97y1vy5nsjrn
- Start time basis: `vendor_reported`.

<a id="ci-6e2d52ee775196cb2782fba0"></a>

- Canonical ID: `ci-6e2d52ee775196cb2782fba0` (1 source records).
  - [openai/01M66KSDH67V5BQ73PE8AD2W8S](../../data/2026/W44/openai.json); vendor URL: https://status.openai.com/incidents/01M66KSDH67V5BQ73PE8AD2W8S
- Start time basis: `vendor_reported`.

<a id="ci-a775a295344f822a5b55b1f4"></a>

- Canonical ID: `ci-a775a295344f822a5b55b1f4` (1 source records).
  - [cloudflare/10m2bxjztzwh](../../data/2026/W44/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/10m2bxjztzwh
- Start time basis: `vendor_reported`.

<a id="ci-02c4567f5ea9d068a54d85a6"></a>

- Canonical ID: `ci-02c4567f5ea9d068a54d85a6` (1 source records).
  - [cloudflare/k2cd9e1fyskp](../../data/2026/W44/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/k2cd9e1fyskp
- Start time basis: `vendor_reported`.

<a id="ci-de81578eb953a59591342afa"></a>

- Canonical ID: `ci-de81578eb953a59591342afa` (1 source records).
  - [cloudflare/oco7f885d2v4](../../data/2026/W44/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/oco7f885d2v4
- Start time basis: `vendor_reported`.

<a id="ci-58d2f990a87d08aa530c55b1"></a>

- Canonical ID: `ci-58d2f990a87d08aa530c55b1` (1 source records).
  - [cloudflare/dk2obcyw2aew](../../data/2026/W44/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/dk2obcyw2aew
- Start time basis: `vendor_reported`.

<a id="ci-61e652f02042e7dd81d50b04"></a>

- Canonical ID: `ci-61e652f02042e7dd81d50b04` (1 source records).
  - [grafana/jl41d84gp27h](../../data/2026/W44/grafana.json); vendor URL: https://stspg.io/d7mh6vy0g5jh
- Start time basis: `vendor_reported`.

<a id="ci-786e978b5877fde448a171ad"></a>

- Canonical ID: `ci-786e978b5877fde448a171ad` (1 source records).
  - [anthropic/0q89dylg9vjk](../../data/2026/W44/anthropic.json); vendor URL: https://stspg.io/rql9h590npxd
- Start time basis: `vendor_reported`.

<a id="ci-abaadece8ac24fe3caa5b9e5"></a>

- Canonical ID: `ci-abaadece8ac24fe3caa5b9e5` (1 source records).
  - [anthropic/lqhknfdjrcst](../../data/2026/W44/anthropic.json); vendor URL: https://stspg.io/39x7qh3sd7bd
- Start time basis: `vendor_reported`.

<a id="ci-0b3ca603b3dbf6efd01e6414"></a>

- Canonical ID: `ci-0b3ca603b3dbf6efd01e6414` (1 source records).
  - [cloudflare/8s6pn9rvof3l](../../data/2026/W44/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/8s6pn9rvof3l
- Start time basis: `vendor_reported`.

<a id="ci-b4638db777542c16ec149ad2"></a>

- Canonical ID: `ci-b4638db777542c16ec149ad2` (1 source records).
  - [github/rn76zdy367m2](../../data/2026/W44/github.json); vendor URL: https://stspg.io/22p3vp78z20q
- Start time basis: `vendor_reported`.

<a id="ci-0800b0801d93c1a197b5bd5f"></a>

- Canonical ID: `ci-0800b0801d93c1a197b5bd5f` (1 source records).
  - [github/lyyfm7j4pkgl](../../data/2026/W44/github.json); vendor URL: https://stspg.io/5bh6qnzn76tk
- Start time basis: `vendor_reported`.

<a id="ci-b4e01bec681ec53bcb92089e"></a>

- Canonical ID: `ci-b4e01bec681ec53bcb92089e` (1 source records).
  - [cloudflare/kp28d3kzgft1](../../data/2026/W44/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/kp28d3kzgft1
- Start time basis: `vendor_reported`.

<a id="ci-aec8a7b0f1260818b5afa8b2"></a>

- Canonical ID: `ci-aec8a7b0f1260818b5afa8b2` (1 source records).
  - [cloudflare/bc7e93q3eigp](../../data/2026/W44/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/bc7e93q3eigp
- Start time basis: `vendor_reported`.

<a id="ci-ee5529e9e1b6d84d30de2ee4"></a>

- Canonical ID: `ci-ee5529e9e1b6d84d30de2ee4` (1 source records).
  - [digitalocean/ckmpdk5z9xj9](../../data/2026/W44/digitalocean.json); vendor URL: https://stspg.io/b16t5y07v5l4
- Start time basis: `vendor_reported`.

<a id="ci-f7298a1210434aa48a56af2a"></a>

- Canonical ID: `ci-f7298a1210434aa48a56af2a` (1 source records).
  - [datadog/hnk297tpxhyk](../../data/2026/W44/datadog.json); vendor URL: https://stspg.io/rq69kpmmw7h6
- Start time basis: `vendor_reported`.

<a id="ci-89fc61884dfc52f4f3fbce32"></a>

- Canonical ID: `ci-89fc61884dfc52f4f3fbce32` (1 source records).
  - [stripe/kl8wz1n0z7m7](../../data/2026/W44/stripe.json); vendor URL: https://stspg.io/3gk09h1d44ts
- Start time basis: `vendor_reported`.

<a id="ci-45e9f79363f180cfde1439c1"></a>

- Canonical ID: `ci-45e9f79363f180cfde1439c1` (1 source records).
  - [openai/01M6754SR14HCQ3E9FX426E2N7](../../data/2026/W44/openai.json); vendor URL: https://status.openai.com/incidents/01M6754SR14HCQ3E9FX426E2N7
- Start time basis: `vendor_reported`.

<a id="ci-59ac344e7cd34241c232b2b3"></a>

- Canonical ID: `ci-59ac344e7cd34241c232b2b3` (1 source records).
  - [cloudflare/2eno8w1fk27j](../../data/2026/W44/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/2eno8w1fk27j
- Start time basis: `vendor_reported`.

<a id="ci-d5ebd458c845a56c7356f7d0"></a>

- Canonical ID: `ci-d5ebd458c845a56c7356f7d0` (1 source records).
  - [cloudflare/jajp50huvxtz](../../data/2026/W44/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/jajp50huvxtz
- Start time basis: `vendor_reported`.

<a id="ci-c9902a204ec57730df92852b"></a>

- Canonical ID: `ci-c9902a204ec57730df92852b` (1 source records).
  - [cloudflare/2xiv8fdhk0bh](../../data/2026/W44/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/2xiv8fdhk0bh
- Start time basis: `vendor_reported`.

<a id="ci-fe706913c855d401ef51edaf"></a>

- Canonical ID: `ci-fe706913c855d401ef51edaf` (1 source records).
  - [openai/01M5YY4A3TQETEQ53DVSK8SW10](../../data/2026/W44/openai.json); vendor URL: https://status.openai.com/incidents/01M5YY4A3TQETEQ53DVSK8SW10
- Start time basis: `vendor_reported`.

<a id="ci-4ad2efd1b0495ea9aee3cc5f"></a>

- Canonical ID: `ci-4ad2efd1b0495ea9aee3cc5f` (1 source records).
  - [github/l7ycmglmykgg](../../data/2026/W44/github.json); vendor URL: https://stspg.io/bwcyq078jzms
- Start time basis: `vendor_reported`.

<a id="ci-301cc7d4f3fd40d9227fc6c8"></a>

- Canonical ID: `ci-301cc7d4f3fd40d9227fc6c8` (1 source records).
  - [github/nxy2l0ssxm10](../../data/2026/W44/github.json); vendor URL: https://stspg.io/6p81rf34z4nc
- Start time basis: `vendor_reported`.

<a id="ci-4a34c2ec9e00d3603489e8b6"></a>

- Canonical ID: `ci-4a34c2ec9e00d3603489e8b6` (1 source records).
  - [cloudflare/mbq512v9oe4s](../../data/2026/W44/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/mbq512v9oe4s
- Start time basis: `vendor_reported`.

<a id="ci-753b9e93eab2160f5c1966d4"></a>

- Canonical ID: `ci-753b9e93eab2160f5c1966d4` (1 source records).
  - [cloudflare/dv2vbq9kbgxn](../../data/2026/W44/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/dv2vbq9kbgxn
- Start time basis: `vendor_reported`.

<a id="ci-d549b7600da800e13e20be49"></a>

- Canonical ID: `ci-d549b7600da800e13e20be49` (1 source records).
  - [cloudflare/4u4hsqqnr0hk](../../data/2026/W44/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/4u4hsqqnr0hk
- Start time basis: `vendor_reported`.

<a id="ci-0bf97609aa6d564738fee96e"></a>

- Canonical ID: `ci-0bf97609aa6d564738fee96e` (1 source records).
  - [cloudflare/s5io0kegzlsz](../../data/2026/W44/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/s5io0kegzlsz
- Start time basis: `vendor_reported`.

<a id="ci-aeb80215afb318b1a866c174"></a>

- Canonical ID: `ci-aeb80215afb318b1a866c174` (1 source records).
  - [cloudflare/03gc4o049emd](../../data/2026/W44/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/03gc4o049emd
- Start time basis: `vendor_reported`.

<a id="ci-467d8fd5a50d92527062d91a"></a>

- Canonical ID: `ci-467d8fd5a50d92527062d91a` (1 source records).
  - [cloudflare/ktq18tcoqoo2](../../data/2026/W44/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/ktq18tcoqoo2
- Start time basis: `vendor_reported`.

<a id="ci-6f306bb3ac0b303d6cfd9b15"></a>

- Canonical ID: `ci-6f306bb3ac0b303d6cfd9b15` (1 source records).
  - [circleci/8q9j33r2rssy](../../data/2026/W44/circleci.json); vendor URL: https://stspg.io/t4pc57wmjnql
- Start time basis: `vendor_reported`.

<a id="ci-c4a65e03ade80f253d75ab0f"></a>

- Canonical ID: `ci-c4a65e03ade80f253d75ab0f` (1 source records).
  - [grafana/wcvm832ndmhp](../../data/2026/W44/grafana.json); vendor URL: https://stspg.io/42ggj9rqyyjm
- Start time basis: `vendor_reported`.

<a id="ci-916357ea68d95b4d67dafef3"></a>

- Canonical ID: `ci-916357ea68d95b4d67dafef3` (1 source records).
  - [cloudflare/h6fgl47h6iht](../../data/2026/W44/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/h6fgl47h6iht
- Start time basis: `vendor_reported`.

<a id="ci-93aca4178d986501260a03d4"></a>

- Canonical ID: `ci-93aca4178d986501260a03d4` (1 source records).
  - [github/f1vvv30224wd](../../data/2026/W44/github.json); vendor URL: https://stspg.io/0622ty903b4d
- Start time basis: `vendor_reported`.

<a id="ci-142516150e151362b50b486f"></a>

- Canonical ID: `ci-142516150e151362b50b486f` (1 source records).
  - [cloudflare/j3llzsq3m17a](../../data/2026/W44/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/j3llzsq3m17a
- Start time basis: `vendor_reported`.

<a id="ci-dda285399c2b9bdb103698d1"></a>

- Canonical ID: `ci-dda285399c2b9bdb103698d1` (1 source records).
  - [openai/01M5Y60NQGGN514EB3943V94EN](../../data/2026/W44/openai.json); vendor URL: https://status.openai.com/incidents/01M5Y60NQGGN514EB3943V94EN
- Start time basis: `vendor_reported`.

<a id="ci-078a2b169acf2fb0138b752b"></a>

- Canonical ID: `ci-078a2b169acf2fb0138b752b` (1 source records).
  - [cloudflare/5tlm9r9j9rpo](../../data/2026/W44/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/5tlm9r9j9rpo
- Start time basis: `vendor_reported`.

<a id="ci-61385bd384bb520ece77d186"></a>

- Canonical ID: `ci-61385bd384bb520ece77d186` (1 source records).
  - [grafana/2b57xck2b1xf](../../data/2026/W44/grafana.json); vendor URL: https://stspg.io/dgmytp3hyf4l
- Start time basis: `vendor_reported`.

<a id="ci-9c1f4e002ad56c6b053b4f37"></a>

- Canonical ID: `ci-9c1f4e002ad56c6b053b4f37` (1 source records).
  - [openai/01M5WPZTT0385GMEV1353ZC4AM](../../data/2026/W44/openai.json); vendor URL: https://status.openai.com/incidents/01M5WPZTT0385GMEV1353ZC4AM
- Start time basis: `vendor_reported`.

<a id="ci-e2175d9923223b1264dff745"></a>

- Canonical ID: `ci-e2175d9923223b1264dff745` (1 source records).
  - [github/dnslv7xdl4fj](../../data/2026/W44/github.json); vendor URL: https://stspg.io/4z5lmy7gqjk1
- Start time basis: `vendor_reported`.
- Warning: Vendor resolved_at precedes started_at; duration is unknown, score uses zero duration; decision:legacy-time-github-dnslv7xdl4fj

## Sources

| Vendor | Status | Source records seen | Parse errors |
|--------|--------|----------------|--------------|
| anthropic | ok | 50 | 0 |
| atlassian | ok | 34 | 0 |
| aws | ok | 6 | 0 |
| circleci | ok | 50 | 0 |
| cloudflare | ok | 50 | 0 |
| datadog | FETCH FAILED: GET https://status.datadoghq.com/api/v2/incidents.json: status 403 | 0 | 0 |
| digitalocean | ok | 50 | 0 |
| gcp | ok | 8 | 0 |
| github | ok | 50 | 0 |
| grafana | ok | 50 | 0 |
| openai | ok | 25 | 0 |
| slack | ok | 28 | 0 |
| stripe | ok | 50 | 0 |
| yandexcloud | FETCH FAILED: GET https://status.cloud.yandex.ru/api/incidents?lang=en&page=4: status 403 | 0 | 0 |
