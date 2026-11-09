# Week 2026-W45

2026-11-02 to 2026-11-08 (UTC). Generated 2026-11-09 16:57.

45 incidents collected, 16 above threshold. 2 of 14 sources unavailable: datadog, yandexcloud.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | cloudflare | [Fedramp Customers have Delayed Purges](https://www.cloudflarestatus.com/incidents/wp49bp8bqo6c) | 2026-11-06 17:47 | 2d 23h 10m (ongoing) | minor | CDN Cache Purge | 7686 = 4270 × 1 × 1.8 |
| 2 | cloudflare | [Spectrum Availability Issue in Phoenix](https://www.cloudflarestatus.com/incidents/c60k3g4vsoc1) | 2026-11-07 00:59 | 2d 15h 58m (ongoing) | minor | Spectrum | 6908 = 3838 × 1 × 1.8 |
| 3 | openai | [Elevated error rates with some U.S. business Mastercards](https://status.openai.com/incidents/01M6S2FV5TZWQ753NHRB7BF5K7) | 2026-11-06 18:38 | 2d 22h 18m (ongoing) | none | — | 6327 = 4218 × 1 × 1.5 |
| 4 | grafana | [October usage report visualization issue](https://stspg.io/ksvgzjk4ctnl) | 2026-11-03 22:00 | 3d 3h 16m | minor | Cost Management and Billing | 4516 = 4516 × 1 × 1.0 |
| 5 | grafana | [IRM Unable to Voice Call Brazil Numbers](https://stspg.io/m1k738p976cl) | 2026-11-02 00:36 | 2d 8h 9m | major | Incident Management and Response (IRM) | 3369 = 3369 × 1 × 1.0 |
| 6 | cloudflare | [Increased Latency and Errors for R2 and Downstream Services in Western North America (WNAM)](https://www.cloudflarestatus.com/incidents/dr9s0uvzwk6n) | 2026-11-03 13:00 | 9h 15m | minor | Stream, Pipelines, R2 | 2997 = 555 × 3 × 1.8 |
| 7 | cloudflare | [Increased Errors for Durable Objects in Western Europe](https://www.cloudflarestatus.com/incidents/oy8gxjb38yse) | 2026-11-05 13:23 | 19h 41m | minor | Durable Objects | 2126 = 1181 × 1 × 1.8 |
| 8 | openai | [ChatGPT Work and Dots Errors](https://status.openai.com/incidents/01M6P50GPP1V98EF8MJ9SH393R) | 2026-11-05 15:25 | 13h 6m | minor | — | 1179 = 786 × 1 × 1.5 |
| 9 | github | [Incident with Pull Requests, Git Operations and Issues](https://stspg.io/18blsgwnbyy6) | 2026-11-04 09:49 | 2h 0m | major | Git Operations, API Requests, Issues, Pull Requests, Actions | 900 = 120 × 5 × 1.5 |
| 10 | cloudflare | [Cloudflare Workers and Pages build errors](https://www.cloudflarestatus.com/incidents/nbyxn84kvb6o) | 2026-11-03 14:46 | 4h 51m | minor | Workers Builds | 524 = 291 × 1 × 1.8 |
| 11 | openai | [File upload errors in ChatGPT](https://status.openai.com/incidents/01M6FDZ89W4PKDS19K496J307G) | 2026-11-03 00:46 | 4h 52m | minor | — | 438 = 292 × 1 × 1.5 |
| 12 | aws | [Elevated Launch Failures](https://health.aws.amazon.com/health/status) | 2026-11-06 22:51 | 2h 21m | unknown | ec2-us-west-2 | 282 = 141 × 1 × 2.0 |
| 13 | cloudflare | [Network Performance Issues in APJC](https://www.cloudflarestatus.com/incidents/j25w3rpv45gs) | 2026-11-03 03:26 | 2h 30m | minor | Network | 270 = 150 × 1 × 1.8 |
| 14 | cloudflare | [Cloudflare Gateway DNS-over-TLS Connectivity Issues over IPv4](https://www.cloudflarestatus.com/incidents/3bv9locsg7ln) | 2026-11-02 22:12 | 2h 21m | minor | Gateway | 254 = 141 × 1 × 1.8 |
| 15 | github | [Incident with Actions and Pages](https://stspg.io/pjgptg3cfc6n) | 2026-11-04 16:05 | 55m | major | Webhooks, Actions, Pages | 248 = 55 × 3 × 1.5 |
| 16 | github | [Incident with Actions](https://stspg.io/7pwv28hvkmz8) | 2026-11-04 17:40 | 35m | major | Actions | 52 = 35 × 1 × 1.5 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. cloudflare — Fedramp Customers have Delayed Purges

- URL: https://www.cloudflarestatus.com/incidents/wp49bp8bqo6c
- Started: 2026-11-06 17:47 UTC
- Resolved: not yet
- Duration: 2d 23h 10m (ongoing)
- Impact: minor, status: identified
- Components: CDN Cache Purge

- `2026-11-06 17:47` **investigating** — Cloudflare has identified a bug with our purge pipeline impacting fedramp customers. Stale content may be served until expired. Non-fedramp customers are not impacted.
- `2026-11-06 17:59` **identified** — We are continuing to work on a fix for the issue.

### 2. cloudflare — Spectrum Availability Issue in Phoenix

- URL: https://www.cloudflarestatus.com/incidents/c60k3g4vsoc1
- Started: 2026-11-07 00:59 UTC
- Resolved: not yet
- Duration: 2d 15h 58m (ongoing)
- Impact: minor, status: monitoring
- Components: Spectrum

- `2026-11-07 00:59` **investigating** — Cloudflare is investigating issues with Spectrum in Phoenix. This may be impacting the availability of existing Spectrum applications. We are working to understand the full impact and mitigate this problem. More updates to follow shortly.
- `2026-11-07 02:51` **monitoring** — We are continuing to monitor for any further issues.

### 3. openai — Elevated error rates with some U.S. business Mastercards

- URL: https://status.openai.com/incidents/01M6S2FV5TZWQ753NHRB7BF5K7
- Started: 2026-11-06 18:38 UTC
- Resolved: not yet
- Duration: 2d 22h 18m (ongoing)
- Impact: none, status: monitoring
- Components: —

- `2026-11-06 18:38` **identified** — **Investigating — We’re seeing payment declines on some U.S. business Mastercards. Affected cards have numbers beginning with 544015, 546400, 543502, 543799, or 512998. Some cards beginning with 513989 are also affected.** **For API customers, this can prevent automatic recharges and one-time credit purchases. If your prepaid balance runs out, your API requests may fail.** **If you use one of these cards, we recommend adding a card from a different issuer before your next payment, even if you haven’t seen a decline. If an API credit purchase has already failed, use the new card to retry the purchase.** **We’re working with our payment partners and will share updates here.**
- `2026-11-06 18:50` **identified** — Investigating — We’re seeing payment declines on some U.S. business Mastercards. Affected cards have numbers beginning with 544015, 546400, 543502, 543799, or 512998. Some cards beginning with 513989 are also affected. For API customers, this can prevent automatic recharges and one-time credit purchases. If your prepaid balance runs out, your API requests may fail. If you use one of these cards, we recommend adding a card from a different issuer before your next payment, even if you haven’t seen a decline. If an API credit purchase has already failed, use the new card to retry the purchase. We’re working with our payment partners and will share updates here.
- `2026-11-07 01:29` **monitoring** — We’ve applied a mitigation and are monitoring the affected services to confirm recovery.

### 4. grafana — October usage report visualization issue

- URL: https://stspg.io/ksvgzjk4ctnl
- Started: 2026-11-03 22:00 UTC
- Resolved: 2026-11-07 01:16 UTC
- Duration: 3d 3h 16m
- Impact: minor, status: resolved
- Components: Cost Management and Billing

- `2026-11-03 22:00` **identified** — There was an issue with our billing metrics exporter that caused us to mis-report usage for October at month-end. This is causing October usage to be reported as 0 in our cost management app. We are working to backfill the correct data. Note that this did not affect billing, and is only a visualization issue.
- `2026-11-04 22:11` **identified** — We are still working to backfill the correct data. Note that this did not affect billing, and is only a visualization issue.
- `2026-11-06 21:41` **identified** — We are currently back filling the missing data. We will provide another update once the backfill has been completed.
- `2026-11-07 01:16` **resolved** — This incident has been resolved.

### 5. grafana — IRM Unable to Voice Call Brazil Numbers

- URL: https://stspg.io/m1k738p976cl
- Started: 2026-11-02 00:36 UTC
- Resolved: 2026-11-04 08:45 UTC
- Duration: 2d 8h 9m
- Impact: major, status: resolved
- Components: Incident Management and Response (IRM)

- `2026-11-02 00:36` **investigating** — We are currently investigating an issue impacting Voice calls to Brazil for IRM. We are working on a resolution, and an update will follow as soon as we have more info. Users should configure mobile push or SMS as notifications.
- `2026-11-02 15:07` **investigating** — We are continuing to work on a resolution, and an update will follow as soon as we have more info. Users should configure mobile push or SMS as notifications.
- `2026-11-02 17:11` **investigating** — Our efforts towards a resolution continues. We encourage usage of SMS and mobile push as alternative way in the meantime. Thank you for your patience.
- `2026-11-03 16:12` **investigating** — We are starting to see signs of recovery therefore some users might see voice calls working fine for them, but the issue is not fully resolved yet. We are continuing our efforts towards a complete resolution and an update will follow as soon as we have more info. We encourage usage of SMS and mobile push as alternative notification mechanism in the meantime.
- `2026-11-04 08:45` **resolved** — Brazil calls recovered at about 13:00 UTC on 3 November. At this point, we are marking the incident as resolved.

### 6. cloudflare — Increased Latency and Errors for R2 and Downstream Services in Western North America (WNAM)

- URL: https://www.cloudflarestatus.com/incidents/dr9s0uvzwk6n
- Started: 2026-11-03 13:00 UTC
- Resolved: 2026-11-03 22:15 UTC
- Duration: 9h 15m
- Impact: minor, status: resolved
- Components: Stream, Pipelines, R2

- `2026-11-04 14:43` **resolved** — Cloudflare has resolved an issue affecting R2 and downstream services in the Western North America region. Some customers may have seen elevated latency and errors between approximately 13:00 and 22:15 UTC on November 3, 2026. This may also have affected downstream services such as Stream and Pipelines.

### 7. cloudflare — Increased Errors for Durable Objects in Western Europe

- URL: https://www.cloudflarestatus.com/incidents/oy8gxjb38yse
- Started: 2026-11-05 13:23 UTC
- Resolved: 2026-11-06 09:05 UTC
- Duration: 19h 41m
- Impact: minor, status: resolved
- Components: Durable Objects

- `2026-11-05 13:23` **monitoring** — Cloudflare investigated an issue that caused an increase of errors for Durable Objects in Western Europe. The issue has been resolved and we are monitoring the results.
- `2026-11-06 09:05` **resolved** — This incident has been resolved.

### 8. openai — ChatGPT Work and Dots Errors

- URL: https://status.openai.com/incidents/01M6P50GPP1V98EF8MJ9SH393R
- Started: 2026-11-05 15:25 UTC
- Resolved: 2026-11-06 04:31 UTC
- Duration: 13h 6m
- Impact: minor, status: resolved
- Components: —

- `2026-11-05 15:25` **investigating** — We’re investigating a potential issue for the listed services and checking how it affects users. We’ll share an update when we know more.
- `2026-11-05 15:30` **investigating** — We’re investigating a potential issue for the listed services and checking how it affects users. We’ll share an update when we know more.
- `2026-11-05 16:26` **investigating** — We’re investigating a potential issue for the listed services and checking how it affects users. We’ll share an update when we know more.
- `2026-11-05 16:57` **identified** — We have identified that users are experiencing elevated errors for the impacted services. We are working on implementing a mitigation.
- `2026-11-05 17:05` **monitoring** — We’ve applied a mitigation and are monitoring the affected services to confirm recovery.
- `2026-11-06 04:31` **resolved** — All affected services are operating normally again. We’re sorry for the disruption, and we appreciate your patience.

### 9. github — Incident with Pull Requests, Git Operations and Issues

- URL: https://stspg.io/18blsgwnbyy6
- Started: 2026-11-04 09:49 UTC
- Resolved: 2026-11-04 11:50 UTC
- Duration: 2h 0m
- Impact: major, status: resolved
- Components: Git Operations, API Requests, Issues, Pull Requests, Actions

- `2026-11-04 09:49` **investigating** — We are investigating reports of degraded performance for Git Operations, Issues and Pull Requests
- `2026-11-04 09:50` **investigating** — Actions is experiencing degraded performance. We are continuing to investigate.
- `2026-11-04 09:53` **investigating** — Pull Requests is experiencing degraded availability. We are continuing to investigate.
- `2026-11-04 10:38` **investigating** — We're investigating elevated errors and intermittent failures affecting GitHub including Pull Requests, Git Operations, and Issues. We'll continue to keep users updated as we work toward mitigation.
- `2026-11-04 11:11` **investigating** — API Requests is experiencing degraded performance. We are continuing to investigate.
- `2026-11-04 11:17` **investigating** — We're currently working on mitigating elevated errors and intermittent failures affecting GitHub, including Pull Requests, Git Operations, and Issues. Some requests may fail or be delayed while we work to restore service. We'll provide another update as more information becomes available.
- `2026-11-04 11:30` **investigating** — The degradation affecting API Requests, Git Operations, Issues and Pull Requests has been mitigated. We are monitoring to ensure stability.
- `2026-11-04 11:43` **monitoring** — The degradation affecting Actions has been mitigated. We are monitoring to ensure stability.
- `2026-11-04 11:50` **resolved** — This incident has been resolved. Thank you for your patience and understanding as we addressed this issue. A detailed root cause analysis will be shared as soon as it is available.

### 10. cloudflare — Cloudflare Workers and Pages build errors

- URL: https://www.cloudflarestatus.com/incidents/nbyxn84kvb6o
- Started: 2026-11-03 14:46 UTC
- Resolved: 2026-11-03 19:38 UTC
- Duration: 4h 51m
- Impact: minor, status: resolved
- Components: Workers Builds

- `2026-11-03 14:46` **investigating** — Cloudflare is aware of and investigating an issue with Cloudflare Workers and Pages build failures which potentially impacts multiple customers. Further detail will be provided as more information becomes available.
- `2026-11-03 15:15` **identified** — Customers may experience increased queue times for Workers Builders
- `2026-11-03 17:23` **monitoring** — Queue times are returning to normal. Customers with failed builds can "Retry" via the Dashboard or pushing a new commit.
- `2026-11-03 19:38` **resolved** — This incident has been resolved.

### 11. openai — File upload errors in ChatGPT

- URL: https://status.openai.com/incidents/01M6FDZ89W4PKDS19K496J307G
- Started: 2026-11-03 00:46 UTC
- Resolved: 2026-11-03 05:39 UTC
- Duration: 4h 52m
- Impact: minor, status: resolved
- Components: —

- `2026-11-03 00:46` **identified** — Some users are having trouble uploading files to ChatGPT or getting ChatGPT to read uploaded files. We are working on implementing a mitigation.
- `2026-11-03 00:58` **investigating** — We’re investigating a potential issue for the listed services and checking how it affects users. We’ll share an update when we know more.
- `2026-11-03 02:06` **identified** — We have identified that users are experiencing elevated errors for the impacted services. We are working on implementing a mitigation.
- `2026-11-03 02:51` **monitoring** — We’ve applied a mitigation and are monitoring the affected services to confirm recovery.
- `2026-11-03 05:39` **resolved** — All affected services are operating normally again. We’re sorry for the disruption, and we appreciate your patience.

### 12. aws — Elevated Launch Failures

- URL: https://health.aws.amazon.com/health/status
- Started: 2026-11-06 22:51 UTC
- Resolved: 2026-11-07 01:12 UTC
- Duration: 2h 21m
- Impact: unknown, status: resolved
- Components: ec2-us-west-2

- `2026-11-06 22:51` **investigating** — We are experiencing elevated launch failures in the US-WEST-2 Region.
- `2026-11-06 23:19` **investigating** — We are experiencing elevated launch failures in the US-WEST-2 Region caused by an issue with the EBS service. Services that depend on EBS or EC2 will experience failure rates of their own. We are working to mitigate the issues. We will provide you with another update by 4:00 PM PST or sooner if new information becomes available.
- `2026-11-07 00:08` **investigating** — We can confirm EC2 instance launch failures affecting all AZs (usw2-az1, usw2-az2, usw2-az3, usw2-az4) in the US-WEST-2 Region. We have identified the issue to be caused by a subsystem that supports EBS Volumes required during instance launch. Services that instantiate new instances, such as cross-region data copy for DataSync, within the US-WEST-2 Region as either the source or destination are also affected. We recommend that customers do not mutate running instances at this time that would result in the replacement or creation of a new EC2 instance in the Region. Running EC2 instances and already created EBS volumes continue to operate as normal. We are seeing early signs of recovery, and continue to take steps to restore launches within the Region. We will update this post by 5:30 PM PST or earlier if we have new information to share.
- `2026-11-07 01:12` **resolved** — Between 2:35 PM and 4:21 PM PST, we experienced elevated EC2 instance launch failures affecting all AZs (usw2-az1, usw2-az2, usw2-az3, usw2-az4) in the US-WEST-2 Region. EC2 Instance Launches, AWS services that instantiated new instances, as well as services such as DataSync or AWS Backup which require new instances to support data transfers to or from the US-WEST-2 Region were affected. Our engineers were automatically engaged and immediately began investigating the root cause. We identified the issue to be caused by a subsystem that supports EBS Volumes required during instance launch. We began remediation of that issue at 3:20 PM, and completed full recovery at 4:21 PM. We recommend customers retry requests that failed during this time. Running EC2 instances and already created EBS volumes were not impacted and continued to operate as normal. The issue has been resolved and all AWS services are now operating normally.

### 13. cloudflare — Network Performance Issues in APJC

- URL: https://www.cloudflarestatus.com/incidents/j25w3rpv45gs
- Started: 2026-11-03 03:26 UTC
- Resolved: 2026-11-03 05:56 UTC
- Duration: 2h 30m
- Impact: minor, status: resolved
- Components: Network

- `2026-11-03 04:44` **investigating** — Cloudflare is investigating issues with Network Performance in the APJC region. We are working to analyze and mitigate this problem. More updates to follow shortly.
- `2026-11-03 04:53` **identified** — The issue has been identified and a fix is being implemented.
- `2026-11-03 05:49` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-11-03 05:56` **resolved** — This incident has been resolved.

### 14. cloudflare — Cloudflare Gateway DNS-over-TLS Connectivity Issues over IPv4

- URL: https://www.cloudflarestatus.com/incidents/3bv9locsg7ln
- Started: 2026-11-02 22:12 UTC
- Resolved: 2026-11-03 00:34 UTC
- Duration: 2h 21m
- Impact: minor, status: resolved
- Components: Gateway

- `2026-11-02 22:12` **investigating** — Cloudflare is investigating an issue affecting Cloudflare Gateway DNS-over-TLS (DoT) connections over IPv4. Affected customers may experience connection timeouts and failed DNS resolution. This issue is limited to DoT over IPv4. DoT over IPv6, DNS-over-HTTPS (DoH), and standard DNS over UDP/TCP are not affected. We are working to identify the cause and restore connectivity. Further updates will be provided as more information becomes available.
- `2026-11-03 00:23` **monitoring** — Our team has been able to identify and apply an effective fix for this issue. We are currently monitoring the results.
- `2026-11-03 00:34` **resolved** — This incident has been resolved.

### 15. github — Incident with Actions and Pages

- URL: https://stspg.io/pjgptg3cfc6n
- Started: 2026-11-04 16:05 UTC
- Resolved: 2026-11-04 17:01 UTC
- Duration: 55m
- Impact: major, status: resolved
- Components: Webhooks, Actions, Pages

- `2026-11-04 16:05` **investigating** — We are investigating reports of degraded performance for Actions and Pages
- `2026-11-04 16:11` **investigating** — We are investigating reports of issues with Actions and Pages. We will continue to keep users updated on progress towards mitigation.
- `2026-11-04 16:12` **investigating** — Webhooks is experiencing degraded performance. We are continuing to investigate.
- `2026-11-04 16:17` **investigating** — Actions is experiencing degraded availability. We are continuing to investigate.
- `2026-11-04 16:34` **investigating** — We have identified an issue with events that might trigger Actions. We have applied mitigations and are monitoring for recovery at this point. Customers may observe delayed status updates for their jobs.
- `2026-11-04 16:45` **investigating** — The degradation affecting Webhooks has been mitigated. We are monitoring to ensure stability.
- `2026-11-04 16:48` **monitoring** — The degradation affecting Actions and Pages has been mitigated. We are monitoring to ensure stability.
- `2026-11-04 16:49` **monitoring** — The degradation has been mitigated. We are monitoring to ensure stability.
- `2026-11-04 17:01` **resolved** — On November 4, 2026, between 15:49 UTC and 16:48 UTC, and again between 17:20 UTC and 17:55 UTC, GitHub Actions, Pages, Webhooks, Code Scanning, and Copilot coding agent experienced delays. During this time, Actions workflow runs and job status updates were delayed by up to 25 minutes. Pages deployments dropped to about 4% of normal volume, webhook deliveries dropped to about 58% of normal volume, and Code Scanning analyses and Copilot coding agent sessions were delayed. Delayed events were processed after mitigation. <br /><br />This incident was caused by a configuration change to an internal event-routing service that unexpectedly paused delivery of several event streams. The change depended on feature flags that were not set up as intended, which the service did not handle gracefully. We mitigated the first period by rolling back the configuration change at 16:14 UTC. The code change itself was not reverted by the rollback, so a subsequent deployment reintroduced it. We mitigated the second period at 17:40 UTC by rolling back again, then processed queued events and cleared backlogs. <br /><br />We are taking several steps to prevent a repeat. The routing service will keep its existing behavior when a feature toggle is missing, so a misconfigured flag can no longer pause delivery. We are adding a pre-deploy check that each routing flag exists and is registered under the correct owning service, and we are making flag ownership less error-prone to set. We are also improving alerting so we detect paused event delivery sooner.

### 16. github — Incident with Actions

- URL: https://stspg.io/7pwv28hvkmz8
- Started: 2026-11-04 17:40 UTC
- Resolved: 2026-11-04 18:16 UTC
- Duration: 35m
- Impact: major, status: resolved
- Components: Actions

- `2026-11-04 17:40` **investigating** — We are investigating reports of degraded availability for Actions
- `2026-11-04 17:53` **investigating** — We have identified an issue with events that might trigger Actions. We have applied mitigations and are monitoring for recovery at this point. Customers may observe delayed status updates for their jobs.
- `2026-11-04 18:07` **monitoring** — The degradation affecting Actions has been mitigated. We are monitoring to ensure stability.
- `2026-11-04 18:16` **resolved** — Please refer to the combined summary in this related incident: https://www.githubstatus.com/incidents/xr3myzj7ms51

## Below threshold (29)

- `grafana` [GitHub returning 403 for all Git Sync repositories](https://stspg.io/c1s5cmykytmk) — 1h 57m
- `grafana` [Elevated Request Errors for Grafana Across Multiple Regions](https://stspg.io/jmzh1gddhwsr) — 1h 1m
- `cloudflare` [Network latency issues in NRT](https://www.cloudflarestatus.com/incidents/nbnize4e990z) — 1h 53m
- `github` [Degraded performance for multiple GitHub services](https://stspg.io/s5bs6jn41fwh) — 44m
- `cloudflare` [Network Performance Issues in Singapore](https://www.cloudflarestatus.com/incidents/3ehw20ozwlkr) — 1h 49m
- `grafana` [Increased query latency for Logs in GCP US Central (prod-us-central-0)](https://stspg.io/8w9r2df60c3j) — 1h 30m
- `github` [Increase in timeouts for some unauthenticated users](https://stspg.io/5yt31mfmyc6p) — 1h 56m
- `openai` [Elevated error rates across paid ChatGPT plans](https://status.openai.com/incidents/01M6MBCX427A7113A3T4FJ1N6Y) — 1h 54m
- `openai` [Elevated errors with image generation](https://status.openai.com/incidents/01M6XVDCKRT3PSB0YP8J30GTD5) — 1h 49m
- `anthropic` [Degraded performance for Claude Fable 5.1](https://stspg.io/3m3q1nwb142k) — 26m
- `github` [Incident with Git Operations](https://stspg.io/n478rfsspcmr) — 1h 43m
- `github` [Incident with API, Git, and Copilot requests](https://stspg.io/h1xqhrcn5gqn) — 30m
- `cloudflare` [Network disruption across multiple colocations in India](https://www.cloudflarestatus.com/incidents/1i6fnludjh1m) — 1h 11m
- `cloudflare` [Cloudflare Container Issues](https://www.cloudflarestatus.com/incidents/n5mvqjs26jg9) — 1h 10m
- `github` [Increased task-related errors for Copilot Mission Control](https://stspg.io/hj74883wgn0t) — 1h 16m
- `cloudflare` [Increased Durable Object Errors in the Dallas, TX US region](https://www.cloudflarestatus.com/incidents/7gixemnw9oei) — 1h 3m
- `cloudflare` [Wrangler Deployments, Pages, and Workers Builds failures](https://www.cloudflarestatus.com/incidents/wtsd415cavkc) — 21m
- `grafana` [Metrics Alert Rule Evaluation Failures in AWS Australia (prod-au-southeast-1)](https://stspg.io/yw5qq1khyhj1) — 1h 48m
- `cloudflare` [Worker Deployments Delayed](https://www.cloudflarestatus.com/incidents/7he1pmetw77p) — 48m
- `github` [Incident with Actions](https://stspg.io/lv4q3xm9r0rk) — 52m
- `cloudflare` [Elevated number of 5xx errors in HKG](https://www.cloudflarestatus.com/incidents/6c4ak19xodyt) — 12m
- `cloudflare` [Network Performance Issues in Chicago](https://www.cloudflarestatus.com/incidents/x9orajjmvvwa) — 17m
- `github` [Elevated errors on GPT-5.4 Nano](https://stspg.io/ns6pj9qy9glb) — 38m
- `circleci` [Increased wait time for Machine and Remote Docker jobs](https://stspg.io/pf820dqbcwhc) — 53m
- `cloudflare` [Workers Builds Issues](https://www.cloudflarestatus.com/incidents/yaa0witbl2me) — 28m
- `cloudflare` [Increased Errors for Durable Objects and Downstream Services in Eastern North America (ENAM)](https://www.cloudflarestatus.com/incidents/c3pjgxk8xnzn) — 2m
- `grafana` [Missing Metrics for Some Synthetic Monitoring Checks using the Montreal, North Virginia, and North California public probes](https://stspg.io/gg5pyd5dw55z) — 0m
- `grafana` [Grafana Cloud: k6 timeouts](https://stspg.io/lb883jzqhsdz) — 0m
- `grafana` [Synthetic Monitoring — Jakarta Public Probe Degradation](https://stspg.io/nhj64nnckg57) — 0m

## Sources

| Vendor | Status | Incidents seen | Parse errors |
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

