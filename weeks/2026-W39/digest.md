# Week 2026-W39

2026-09-21 to 2026-09-27 (UTC). Generated 2026-09-28 14:41.

49 incidents collected, 18 above threshold. All 13 sources responded.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | aws | [Increased Error Rates](https://health.aws.amazon.com/health/status) | 2026-09-21 23:26 | 6d 15h 14m (ongoing) | unknown | ec2-us-east-1 | 19108 = 9554 × 1 × 2.0 |
| 2 | cloudflare | [Cloudflare One Clients are incorrectly challenged on some sites](https://www.cloudflarestatus.com/incidents/tl0vrgjy99r0) | 2026-09-22 19:23 | 5d 19h 18m (ongoing) | minor | Cloudflare One Client | 15044 = 8358 × 1 × 1.8 |
| 3 | cloudflare | [Network Performance Degradation — Asia-Pacific](https://www.cloudflarestatus.com/incidents/8wmvkv5jkf15) | 2026-09-23 18:42 | 4d 19h 58m (ongoing) | minor | Network | 12524 = 6958 × 1 × 1.8 |
| 4 | github | [Incident across several services](https://stspg.io/m1ps7yrhp4n8) | 2026-09-23 10:11 | 18h 44m | minor | API Requests | 1686 = 1124 × 1 × 1.5 |
| 5 | cloudflare | [Issues with Durable Objects](https://www.cloudflarestatus.com/incidents/gv29fd076q2t) | 2026-09-25 00:03 | 14h 45m | minor | Durable Objects | 1593 = 885 × 1 × 1.8 |
| 6 | cloudflare | [Intermittent authentication errors for API and R2](https://www.cloudflarestatus.com/incidents/jg6ys8sbd6ny) | 2026-09-23 09:08 | 11h 50m | minor | R2 | 1278 = 710 × 1 × 1.8 |
| 7 | grafana | [IRM Access Issues for a Small Group of Users](https://stspg.io/cqfnrz8bdkvz) | 2026-09-22 14:57 | 20h 10m | major | Incident Management and Response (IRM) | 1210 = 1210 × 1 × 1.0 |
| 8 | grafana | [Kubernetes Observability Billing & Usage Incorrect](https://stspg.io/4ntjl2yr75n4) | 2026-09-22 21:27 | 16h 38m | minor | Cost Management and Billing | 998 = 998 × 1 × 1.0 |
| 9 | anthropic | [Elevated errors for multiple models](https://stspg.io/90s4l9yywzbk) | 2026-09-22 00:57 | 1h 37m | major | claude.ai, Claude API (api.anthropic.com), Claude Code, Claude Cowork | 582 = 97 × 4 × 1.5 |
| 10 | cloudflare | [Browser isolation session data sync issues](https://www.cloudflarestatus.com/incidents/mltthh0jxxpz) | 2026-09-24 17:26 | 5h 11m | minor | Browser Isolation | 560 = 311 × 1 × 1.8 |
| 11 | cloudflare | [Delays Starting Cloudflare Workers Builds](https://www.cloudflarestatus.com/incidents/dwfzz0bwhmsq) | 2026-09-24 17:33 | 3h 18m | minor | Workers Builds | 356 = 198 × 1 × 1.8 |
| 12 | cloudflare | [Partial visibility of Workers and Durable Objects analytics for Fedramp High customers](https://www.cloudflarestatus.com/incidents/dcmwvjwfyhjg) | 2026-09-22 09:36 | 3h 15m | minor | Analytics | 351 = 195 × 1 × 1.8 |
| 13 | github | [Disruption with billing information updates](https://stspg.io/g2q4g5kv9frp) | 2026-09-24 16:51 | 3h 49m | minor | — | 344 = 229 × 1 × 1.5 |
| 14 | cloudflare | [Cloudflare Access updates delayed](https://www.cloudflarestatus.com/incidents/xbgbb80yw0jh) | 2026-09-24 15:20 | 3h 4m | minor | Access | 331 = 184 × 1 × 1.8 |
| 15 | stripe | [API call success rate drop](https://stspg.io/vfb8m896t1fk) | 2026-09-22 20:32 | 2h 14m | none | Stripe API, Stripe core components | 322 = 134 × 2 × 1.2 |
| 16 | cloudflare | [Elevated Errors with any / all in http_response_cache_settings](https://www.cloudflarestatus.com/incidents/f8wg5htd3zz4) | 2026-09-23 15:54 | 2h 13m | minor | Rules | 239 = 133 × 1 × 1.8 |
| 17 | openai | [Issues with Codex](https://status.openai.com/incidents/01M3DCNWMW57HYK8FJ5FBFPA39) | 2026-09-25 22:58 | 55m | critical | — | 82 = 55 × 1 × 1.5 |
| 18 | stripe | [Elevated API Errors](https://stspg.io/4ntbqnhzg7bb) | 2026-09-27 01:03 | 18m | major | Stripe API, Stripe core components | 43 = 18 × 2 × 1.2 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. aws — Increased Error Rates

- URL: https://health.aws.amazon.com/health/status
- Started: 2026-09-21 23:26 UTC
- Resolved: not yet
- Duration: 6d 15h 14m (ongoing)
- Impact: unknown, status: investigating
- Components: ec2-us-east-1

- `2026-09-21 23:26` **investigating** — We are investigating increased API error rates in the US-EAST-1 Region.
- `2026-09-21 23:44` **investigating** — We can confirm increased error rates affecting new instance launches, including elevated API error rates for RunInstances, CreateFleet, and related EC2 launch workflows in the US-EAST-1 Region. Customers attempting to launch new EC2 instances may experience launch failures or elevated API response latency. Our engineers are actively engaged and working towards recovering the issue. We are seeing initial signs of recovery. Existing instances remain unaffected. We recommend retrying failed RunInstances and CreateFleet API calls. We will provide an update by 5:30 PM PDT or sooner.
- `2026-09-22 00:24` **investigating** — Between 3:49 PM and 5:06 PM PDT, we experienced increased error rates affecting new instance launches in the US-EAST-1 Region. During this time, customers may have experienced elevated error rates when attempting to launch new EC2 instances through RunInstances, CreateFleet, and related EC2 launch workflows. AWS services that launch EC2 instances on behalf of customers, including Amazon ECS, AWS Fargate, Amazon EMR, amongst others were also impacted during this period. We identified the root cause as a recent deployment to a subsystem within the EC2 service responsible for processing instance launch requests. We initiated a rollback of the problematic deployment and began observing initial signs of recovery at 4:44 PM. Existing instances remain unaffected. The issue has been resolved and the service is operating normally.

### 2. cloudflare — Cloudflare One Clients are incorrectly challenged on some sites

- URL: https://www.cloudflarestatus.com/incidents/tl0vrgjy99r0
- Started: 2026-09-22 19:23 UTC
- Resolved: not yet
- Duration: 5d 19h 18m (ongoing)
- Impact: minor, status: identified
- Components: Cloudflare One Client

- `2026-09-22 19:23` **investigating** — Cloudflare is investigating an issue where Cloudflare One Client customers are incorrectly challenged while visiting some websites.
- `2026-09-23 11:02` **identified** — The issue has been identified and a fix is being implemented.

### 3. cloudflare — Network Performance Degradation — Asia-Pacific

- URL: https://www.cloudflarestatus.com/incidents/8wmvkv5jkf15
- Started: 2026-09-23 18:42 UTC
- Resolved: not yet
- Duration: 4d 19h 58m (ongoing)
- Impact: minor, status: identified
- Components: Network

- `2026-09-23 18:42` **investigating** — Since September 21, 2026 at 02:20 UTC, multiple subsea cable outages have caused congestion between Tokyo and Singapore datacenters. We have rerouted traffic to reduce impact and are working with third-party vendors to restore capacity.
- `2026-09-23 18:44` **identified** — The issue has been identified and a fix is being implemented.
- `2026-09-25 08:50` **identified** — Cloudflare continues to mitigate intermittent connectivity degradation for a small subset of customers caused by significant degradation of capacity in the region.

### 4. github — Incident across several services

- URL: https://stspg.io/m1ps7yrhp4n8
- Started: 2026-09-23 10:11 UTC
- Resolved: 2026-09-24 04:55 UTC
- Duration: 18h 44m
- Impact: minor, status: resolved
- Components: API Requests

- `2026-09-23 10:11` **investigating** — We are investigating reports of degraded performance for API Requests
- `2026-09-23 10:20` **investigating** — Database replicas have detached. We're working to restore the database replicas. Users may experience issues beyond creating organizations and a degraded experience with the GitHub API and Projects.
- `2026-09-23 10:57` **investigating** — Database replicas have been restored. Org creation and the GitHub API are no longer degraded.
- `2026-09-23 10:58` **investigating** — The degradation affecting API Requests has been mitigated. We are monitoring to ensure stability.
- `2026-09-23 11:38` **investigating** — We are seeing recovery for Projects. Users may experience stale search results for Projects while indexing catches up.
- `2026-09-23 13:35` **investigating** — Users may experience stale Project search results. We are working to increase indexing speed. All other services are available.
- `2026-09-23 17:01` **investigating** — Updates to issue labels may be delayed in being reflected in Projects by about ~10 minutes. We have added some capacity to work through the backlog more quickly, but it'll likely be a few hours to complete processing the full backlog of messages. All other services are available.
- `2026-09-23 17:30` **investigating** — We will post another update in approximately one hour to share our progress.
- `2026-09-23 18:42` **investigating** — Continuing to investigate the lag that may be experienced in issue labels being accurately reflected in Projects. We are working on alternate solutions to process the backlog of label updates.
- `2026-09-23 20:26` **investigating** — We are preparing to deploy a change that will mitigate the impact.
- `2026-09-23 21:39` **investigating** — Updates to issue labels may be delayed in being reflected in Projects. We are continuing to deploy a change that will accelerate processing of the backlog of label updates. All other services are available. We will provide another update within the next hour.
- `2026-09-24 00:16` **investigating** — We've deployed a change intended to accelerate processing of the backlog of issue label updates in Projects. A sizable backlog still remains and we continue working through it. All other services are operating normally. We will provide another update within the next hour.
- `2026-09-24 02:00` **investigating** — We are continuing to process the backlog of issue label updates for Projects. Users may still see delays before label changes are reflected in Projects. All other services are operating normally.
- `2026-09-24 03:14` **investigating** — We are continuing to process the backlog of issue label updates for Projects. Label changes may still be delayed. All other services are operating normally.
- `2026-09-24 04:55` **monitoring** — The degradation has been mitigated. We are monitoring to ensure stability.
- `2026-09-24 04:55` **resolved** — This incident has been resolved. Thank you for your patience and understanding as we addressed this issue. A detailed root cause analysis will be shared as soon as it is available.

### 5. cloudflare — Issues with Durable Objects

- URL: https://www.cloudflarestatus.com/incidents/gv29fd076q2t
- Started: 2026-09-25 00:03 UTC
- Resolved: 2026-09-25 14:49 UTC
- Duration: 14h 45m
- Impact: minor, status: resolved
- Components: Durable Objects

- `2026-09-25 00:03` **investigating** — Cloudflare is aware of and investigating an issue impacting some customers where new Workflows instances may be stuck in a Queued state. More updates to follow shortly.
- `2026-09-25 01:31` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-09-25 02:32` **investigating** — Cloudflare is investigating issues causing an increase of errors for Durable Objects .We are working to analyse and mitigate this problem. More updates to follow shortly.
- `2026-09-25 08:48` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-09-25 14:49` **resolved** — This incident has been resolved.

### 6. cloudflare — Intermittent authentication errors for API and R2

- URL: https://www.cloudflarestatus.com/incidents/jg6ys8sbd6ny
- Started: 2026-09-23 09:08 UTC
- Resolved: 2026-09-23 20:59 UTC
- Duration: 11h 50m
- Impact: minor, status: resolved
- Components: R2

- `2026-09-23 09:08` **identified** — A known issue affected authentication for a small percentage of requests to the API and R2. The issue was identified on Sep 22 at 13:30 UTC, and major impact was mitigated at 19:00 UTC. Our team is actively working to resolve the residual impact.
- `2026-09-23 20:09` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-09-23 20:59` **resolved** — This incident has been resolved.

### 7. grafana — IRM Access Issues for a Small Group of Users

- URL: https://stspg.io/cqfnrz8bdkvz
- Started: 2026-09-22 14:57 UTC
- Resolved: 2026-09-23 11:07 UTC
- Duration: 20h 10m
- Impact: major, status: resolved
- Components: Incident Management and Response (IRM)

- `2026-09-22 14:57` **investigating** — We are currently investigating an issue that may cause a small group of users to experience unexpected IRM access issues, or undelivered pages. Most users should not be affected.
- `2026-09-22 15:55` **identified** — The issue has been identified, and we are working on a fix.
- `2026-09-22 17:19` **monitoring** — A fix is currently being rolled out and will be deployed progressively over the next few hours. We appreciate your patience as this reaches all affected customers.
- `2026-09-22 18:57` **monitoring** — The fix has been deployed to the majority of deployments. We continue to monitor progress.
- `2026-09-23 11:07` **resolved** — Between 14:57 UTC on 22 September and 09:00 UTC on 23 September, a small group of users experienced unexpected access issues and undelivered pages in Grafana IRM. This has been fully resolved and all affected users should now have normal access. Thank you for your patience.

### 8. grafana — Kubernetes Observability Billing & Usage Incorrect

- URL: https://stspg.io/4ntjl2yr75n4
- Started: 2026-09-22 21:27 UTC
- Resolved: 2026-09-23 14:06 UTC
- Duration: 16h 38m
- Impact: minor, status: resolved
- Components: Cost Management and Billing

- `2026-09-22 21:27` **investigating** — Some customers may see a sudden spike in reported Kubernetes Observability usage even if they haven't activated the product. We are investigating the root cause, and will ensure the amount is corrected before any usage is invoiced.
- `2026-09-22 23:10` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-09-23 14:06` **resolved** — This incident has been resolved.

### 9. anthropic — Elevated errors for multiple models

- URL: https://stspg.io/90s4l9yywzbk
- Started: 2026-09-22 00:57 UTC
- Resolved: 2026-09-22 02:35 UTC
- Duration: 1h 37m
- Impact: major, status: resolved
- Components: claude.ai, Claude API (api.anthropic.com), Claude Code, Claude Cowork

- `2026-09-22 00:57` **investigating** — We are investigating elevated errors on requests to Claude Mythos 5.1, Claude Fable 5.1, and Claude Opus 5. We will provide an update as soon as possible.
- `2026-09-22 01:17` **identified** — We have identified the cause of elevated errors on requests to Claude Mythos 5.1, Claude Fable 5.1, and Claude Opus 5 and are working on a fix. We will provide an update as soon as possible.
- `2026-09-22 01:35` **identified** — We are continuing to work to resolve errors affecting some models. At this time, requests to Claude Fable 5 and 5.1 and Mythos 5 and 5.1 have returned to normal success rates. We are working to resolve remaining errors affecting Claude Opus 5, and will provide an additional update shortly.
- `2026-09-22 02:11` **monitoring** — We have seen success rates return to normal across affected models, and are monitoring closely to ensure no further issues.
- `2026-09-22 02:35` **resolved** — This issue has been resolved. Impact occurred from 5:50pm PT / 00:50 UTC to 7:10pm PT / 02:10 UTC.

### 10. cloudflare — Browser isolation session data sync issues

- URL: https://www.cloudflarestatus.com/incidents/mltthh0jxxpz
- Started: 2026-09-24 17:26 UTC
- Resolved: 2026-09-24 22:37 UTC
- Duration: 5h 11m
- Impact: minor, status: resolved
- Components: Browser Isolation

- `2026-09-24 17:26` **identified** — Cloudflare is investigating issues with Browser Isolation's session data sync. Customers that use Browser Isolation may observe errors on acquiring a browser stating that browser storage is unavailable, and they will need to sign into web apps again. They will continue to be able to use browsers during this time.
- `2026-09-24 17:28` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-09-24 22:37` **resolved** — This incident has been resolved.

### 11. cloudflare — Delays Starting Cloudflare Workers Builds

- URL: https://www.cloudflarestatus.com/incidents/dwfzz0bwhmsq
- Started: 2026-09-24 17:33 UTC
- Resolved: 2026-09-24 20:52 UTC
- Duration: 3h 18m
- Impact: minor, status: resolved
- Components: Workers Builds

- `2026-09-24 17:33` **monitoring** — Customers may experience delays starting Cloudflare Workers builds.
- `2026-09-24 19:54` **monitoring** — Queue levels & latency are returning to normal
- `2026-09-24 20:52` **resolved** — This incident has been resolved.

### 12. cloudflare — Partial visibility of Workers and Durable Objects analytics for Fedramp High customers

- URL: https://www.cloudflarestatus.com/incidents/dcmwvjwfyhjg
- Started: 2026-09-22 09:36 UTC
- Resolved: 2026-09-22 12:51 UTC
- Duration: 3h 15m
- Impact: minor, status: resolved
- Components: Analytics

- `2026-09-22 09:36` **identified** — Cloudflare is monitoring the restoration of Workers and Durable Objects analytics visibility for all Fedramp High customers. While the underlying data remains intact, affected datasets were not being correctly returned by our analytics dashboard and GraphQL API since June 11th. Services are currently recovering as we propagate the fix.
- `2026-09-22 10:08` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-09-22 12:51` **resolved** — This incident has been resolved.

### 13. github — Disruption with billing information updates

- URL: https://stspg.io/g2q4g5kv9frp
- Started: 2026-09-24 16:51 UTC
- Resolved: 2026-09-24 20:41 UTC
- Duration: 3h 49m
- Impact: minor, status: resolved
- Components: —

- `2026-09-24 16:51` **investigating** — We are investigating reports of impacted performance for some GitHub services.
- `2026-09-24 16:51` **investigating** — We are investigating failures on billing information updates. Customers may be unable to create or update their billing information in the meantime.
- `2026-09-24 17:40` **investigating** — We have identified the problem and are actively working on mitigation.
- `2026-09-24 18:22` **investigating** — We are continuing to work on mitigation. We will post another update in approximately one hour.
- `2026-09-24 20:10` **investigating** — We have applied the mitigation and are seeing signs of recovery. We are continuing to monitor the system closely.
- `2026-09-24 20:24` **monitoring** — The degradation has been mitigated. We are monitoring to ensure stability.
- `2026-09-24 20:41` **resolved** — This incident has been resolved. Thank you for your patience and understanding as we addressed this issue. A detailed root cause analysis will be shared as soon as it is available.

### 14. cloudflare — Cloudflare Access updates delayed

- URL: https://www.cloudflarestatus.com/incidents/xbgbb80yw0jh
- Started: 2026-09-24 15:20 UTC
- Resolved: 2026-09-24 18:25 UTC
- Duration: 3h 4m
- Impact: minor, status: resolved
- Components: Access

- `2026-09-24 15:20` **investigating** — Cloudflare engineering is investigating an issue causing a delay in configuration updates. This includes applications and policy updates. Authentication and policy enforcement are not impacted.
- `2026-09-24 15:31` **monitoring** — We have applied a fix and are actively monitoring.
- `2026-09-24 18:25` **resolved** — This incident has been resolved.

### 15. stripe — API call success rate drop

- URL: https://stspg.io/vfb8m896t1fk
- Started: 2026-09-22 20:32 UTC
- Resolved: 2026-09-22 22:47 UTC
- Duration: 2h 14m
- Impact: none, status: resolved
- Components: Stripe API, Stripe core components

- `2026-09-22 20:32` **identified** — We're currently investigating an issue affecting API accounts and capability call success rates. We've identified the root cause and are working on a mitigation. We will post further updates as we have more information.
- `2026-09-22 21:27` **monitoring** — API error rates are recovering as part of our mitigations. We're continuing to monitor until the issue is fully resolved. We will provide further updates once we have observed sustained recovery.
- `2026-09-22 22:47` **resolved** — From 19:13 to 22:19 UTC, we saw elevated error rates and response times with the API. This is now resolved.

### 16. cloudflare — Elevated Errors with any / all in http_response_cache_settings

- URL: https://www.cloudflarestatus.com/incidents/f8wg5htd3zz4
- Started: 2026-09-23 15:54 UTC
- Resolved: 2026-09-23 18:08 UTC
- Duration: 2h 13m
- Impact: minor, status: resolved
- Components: Rules

- `2026-09-23 15:54` **identified** — We have identified an issue causing unexpected errors when creating or updating configurations using any or all expressions in the http_response_cache_settings. This issue strictly affects API operations; existing active configurations and live traffic are not impacted. A fix is currently being deployed, and we will provide an update once the rollout is complete.
- `2026-09-23 18:08` **resolved** — This incident has been resolved.

### 17. openai — Issues with Codex

- URL: https://status.openai.com/incidents/01M3DCNWMW57HYK8FJ5FBFPA39
- Started: 2026-09-25 22:58 UTC
- Resolved: 2026-09-25 23:54 UTC
- Duration: 55m
- Impact: critical, status: resolved
- Components: —

- `2026-09-25 22:58` **identified** — We have identified that users are experiencing a codex outage. We have identified the internal issue and are heading towards a mitigation.
- `2026-09-25 23:03` **identified** — We have identified that users are experiencing elevated errors for the impacted services. We are working on implementing a mitigation.
- `2026-09-25 23:19` **identified** — Login via API key will unblock access at this time.
- `2026-09-25 23:34` **identified** — We have identified root cause and are moving towards mitigation.
- `2026-09-25 23:45` **monitoring** — We have applied the mitigation and are monitoring the recovery.
- `2026-09-25 23:54` **resolved** — All impacted services have now fully recovered.

### 18. stripe — Elevated API Errors

- URL: https://stspg.io/4ntbqnhzg7bb
- Started: 2026-09-27 01:03 UTC
- Resolved: 2026-09-27 01:21 UTC
- Duration: 18m
- Impact: major, status: resolved
- Components: Stripe API, Stripe core components

- `2026-09-27 01:03` **investigating** — Starting 00:32 UTC, we’ve seen elevated failures with the API. We’re investigating and will post an update as soon as possible.
- `2026-09-27 01:21` **resolved** — From 00:32 UTC to 01:01 UTC, we saw elevated API errors. This issue is now resolved.

## Below threshold (31)

- `cloudflare` [Intermittent 500 errors for backend services on Replicate](https://www.cloudflarestatus.com/incidents/2x4qt2q2ds9x) — 1h 54m
- `cloudflare` [Increased Errors for Durable Objects](https://www.cloudflarestatus.com/incidents/78c19w4xw4z2) — 1h 50m
- `stripe` [Elevated Klarna errors](https://stspg.io/4kvs6zwpn12h) — 1h 19m
- `cloudflare` [Edit Compression Rule Issue](https://www.cloudflarestatus.com/incidents/drq3pdx1f11y) — 1h 32m
- `cloudflare` [Network connectivity issues in LHR, London](https://www.cloudflarestatus.com/incidents/nhf2kr8q3nfw) — 1h 11m
- `cloudflare` [Bot Management Configuration Propagation Issues](https://www.cloudflarestatus.com/incidents/c2f7zj7fs8l8) — 1h 3m
- `circleci` [Delays in processing new signups and plan changes](https://stspg.io/qdgh114h7kwz) — 1h 52m
- `cloudflare` [Elevated error rates for Stream Live LL-HLS](https://www.cloudflarestatus.com/incidents/pvvq3811tpyq) — 1h 1m
- `cloudflare` [Network Performance Issues in San Jose (SJC-A)](https://www.cloudflarestatus.com/incidents/g4vvd0xrc2m8) — 1h 0m
- `openai` [Elevated Error Rates for ChatGPT Work, across Plus, Pro, Business, Enterprise and Education plans.](https://status.openai.com/incidents/01M34TAFM9WBGEBE3G1GBH4SE8) — 54m
- `stripe` [Elevated API errors](https://stspg.io/zxkgb0lnrxc8) — 33m
- `cloudflare` [Issues with 1.1.1.1 for Families](https://www.cloudflarestatus.com/incidents/rcjbfb0lyn65) — 43m
- `cloudflare` [Unable to start containers in Asia-Pacific](https://www.cloudflarestatus.com/incidents/rxmmx07ys7qg) — 41m
- `openai` [Elevated Error Rates on GPT-6 Astra Pro](https://status.openai.com/incidents/01M3ACKDE4GYY67FXXNRRSC0CE) — 46m
- `cloudflare` [Cloudflare Cloud Access Security Broker (CASB) Issues](https://www.cloudflarestatus.com/incidents/39dzwkhzcvv5) — 37m
- `openai` [Elevated Error Rates for ChatGPT across Plus and Pro plans.](https://status.openai.com/incidents/01M36RMC01ZFWQ861WYJKC4XE1) — 41m
- `cloudflare` [Elevated number of R2 503 errors in Australian Eastern Coast region](https://www.cloudflarestatus.com/incidents/x86v1qk8yp5v) — 17m
- `openai` [Increased error rate for Plus and Pro users.](https://status.openai.com/incidents/01M348TX2DX7KPM9X857N4APNV) — 39m
- `grafana` [Cloud Provider Observability - Hosted Azure metrics issues](https://stspg.io/y9pvmm8jfz0c) — 48m
- `circleci` [Delays starting Linux Machine jobs](https://stspg.io/8gpc503nlf95) — 47m
- `datadog` [Delayed Monitors Notifications](https://stspg.io/k1zlzpm9dfm2) — 46m
- `openai` [Mobile users unable to see Work Mode and the Model Picker](https://status.openai.com/incidents/01M389CDQ97B11QPMSRAAYSE5S) — 24m
- `datadog` [Users are unable to acknowledge, escalate, or resolve On-Call Pages](https://stspg.io/03h92q1l96z2) — 27m
- `cloudflare` [Unable to register certain domains](https://www.cloudflarestatus.com/incidents/lb5514cvngd7) — 11m
- `cloudflare` [Customers using BYOIP can have issues updating their BGP prefixes, including advertising or withdrawing prefixes.](https://www.cloudflarestatus.com/incidents/gwrzznc9f642) — 3m
- `cloudflare` [Increased Cache Failures](https://www.cloudflarestatus.com/incidents/mrqb4d6q0y26) — 0m
- `grafana` [Mimir write request errors](https://stspg.io/j6zm954jcg7m) — 0m
- `cloudflare` [Connectivity issues in Los Angeles (LAX)](https://www.cloudflarestatus.com/incidents/z84lp04f9sht) — 0m
- `cloudflare` [Network Performance Issues in Taipei](https://www.cloudflarestatus.com/incidents/zpkwd9fxsby5) — 0m
- `cloudflare` [Secondary DNS Update Delays](https://www.cloudflarestatus.com/incidents/7dt6f0chnclh) — 0m
- `cloudflare` [R2 Service Issues in Western North America (WNAM)](https://www.cloudflarestatus.com/incidents/6m2w2ls1gv75) — 0m

## Sources

| Vendor | Status | Incidents seen | Parse errors |
|--------|--------|----------------|--------------|
| anthropic | ok | 50 | 0 |
| atlassian | ok | 34 | 0 |
| aws | ok | 6 | 0 |
| circleci | ok | 50 | 0 |
| cloudflare | ok | 50 | 0 |
| datadog | ok | 50 | 0 |
| digitalocean | ok | 50 | 0 |
| gcp | ok | 6 | 0 |
| github | ok | 50 | 0 |
| grafana | ok | 50 | 0 |
| openai | ok | 25 | 0 |
| slack | ok | 35 | 0 |
| stripe | ok | 50 | 0 |

