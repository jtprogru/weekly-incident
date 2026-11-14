# Week 2026-W45

2026-11-02 to 2026-11-08 (UTC). Generated 2026-11-14 12:36.

45 canonical incidents, 16 above threshold. 2 of 14 sources unavailable: datadog, yandexcloud.

45 source records (56 saved revisions) → 45 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | cloudflare | [Fedramp Customers have Delayed Purges](https://www.cloudflarestatus.com/incidents/wp49bp8bqo6c) | 2026-11-06 17:47 | 5d 22h 49m | minor | CDN Cache Purge | 15424 = 8569 × 1 × 1.8 |
| 2 | openai | [Elevated error rates with some U.S. business Mastercards](https://status.openai.com/incidents/01M6S2FV5TZWQ753NHRB7BF5K7) | 2026-11-06 18:38 | 3d 21h 27m | none | — | 8410 = 5607 × 1 × 1.5 |
| 3 | cloudflare | [Spectrum Availability Issue in Phoenix](https://www.cloudflarestatus.com/incidents/c60k3g4vsoc1) | 2026-11-07 00:59 | 2d 13h 37m | minor | Spectrum | 6655 = 3697 × 1 × 1.8 |
| 4 | grafana | [October usage report visualization issue](https://stspg.io/ksvgzjk4ctnl) | 2026-11-03 22:00 | 3d 3h 16m | minor | Cost Management and Billing | 4516 = 4516 × 1 × 1.0 |
| 5 | grafana | [IRM Unable to Voice Call Brazil Numbers](https://stspg.io/m1k738p976cl) | 2026-11-02 00:36 | 2d 8h 9m | major | Incident Management and Response (IRM) | 3369 = 3369 × 1 × 1.0 |
| 6 | cloudflare | [Increased Latency and Errors for R2 and Downstream Services in Western North America (WNAM)](https://www.cloudflarestatus.com/incidents/dr9s0uvzwk6n) | 2026-11-03 13:00 | 9h 15m | minor | Pipelines, R2, Stream | 2997 = 555 × 3 × 1.8 |
| 7 | cloudflare | [Increased Errors for Durable Objects in Western Europe](https://www.cloudflarestatus.com/incidents/oy8gxjb38yse) | 2026-11-05 13:23 | 19h 41m | minor | Durable Objects | 2126 = 1181 × 1 × 1.8 |
| 8 | openai | [ChatGPT Work and Dots Errors](https://status.openai.com/incidents/01M6P50GPP1V98EF8MJ9SH393R) | 2026-11-05 15:25 | 13h 6m | minor | — | 1179 = 786 × 1 × 1.5 |
| 9 | github | [Incident with Pull Requests, Git Operations and Issues](https://stspg.io/18blsgwnbyy6) | 2026-11-04 09:49 | 2h 0m | major | API Requests, Actions, Git Operations, Issues, Pull Requests | 900 = 120 × 5 × 1.5 |
| 10 | cloudflare | [Cloudflare Workers and Pages build errors](https://www.cloudflarestatus.com/incidents/nbyxn84kvb6o) | 2026-11-03 14:46 | 4h 51m | minor | Workers Builds | 524 = 291 × 1 × 1.8 |
| 11 | openai | [File upload errors in ChatGPT](https://status.openai.com/incidents/01M6FDZ89W4PKDS19K496J307G) | 2026-11-03 00:46 | 4h 52m | minor | — | 438 = 292 × 1 × 1.5 |
| 12 | aws | [Elevated Launch Failures](https://health.aws.amazon.com/health/status) | 2026-11-06 22:51 | 2h 21m | unknown | ec2-us-west-2 | 282 = 141 × 1 × 2.0 |
| 13 | cloudflare | [Network Performance Issues in APJC](https://www.cloudflarestatus.com/incidents/j25w3rpv45gs) | 2026-11-03 03:26 | 2h 30m | minor | Network | 270 = 150 × 1 × 1.8 |
| 14 | cloudflare | [Cloudflare Gateway DNS-over-TLS Connectivity Issues over IPv4](https://www.cloudflarestatus.com/incidents/3bv9locsg7ln) | 2026-11-02 22:12 | 2h 21m | minor | Gateway | 254 = 141 × 1 × 1.8 |
| 15 | github | [Incident with Actions and Pages](https://stspg.io/pjgptg3cfc6n) | 2026-11-04 16:05 | 55m | major | Actions, Pages, Webhooks | 248 = 55 × 3 × 1.5 |
| 16 | github | [Incident with Actions](https://stspg.io/7pwv28hvkmz8) | 2026-11-04 17:40 | 35m | major | Actions | 52 = 35 × 1 × 1.5 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. cloudflare — Fedramp Customers have Delayed Purges

- URL: https://www.cloudflarestatus.com/incidents/wp49bp8bqo6c
- Started: 2026-11-06 17:47 UTC
- Resolved: 2026-11-12 16:36 UTC
- Duration: 5d 22h 49m
- Impact: minor, status: resolved
- Components: CDN Cache Purge
- Canonical incident: [ci-b723223fb9dd7c34dba6ff7b](#ci-b723223fb9dd7c34dba6ff7b)

- `2026-11-06 17:47` **investigating** — Cloudflare has identified a bug with our purge pipeline impacting fedramp customers. Stale content may be served until expired. Non-fedramp customers are not impacted.
- `2026-11-06 17:59` **identified** — We are continuing to work on a fix for the issue.
- `2026-11-12 16:36` **resolved** — This incident has been resolved.

### 2. openai — Elevated error rates with some U.S. business Mastercards

- URL: https://status.openai.com/incidents/01M6S2FV5TZWQ753NHRB7BF5K7
- Started: 2026-11-06 18:38 UTC
- Resolved: 2026-11-10 16:06 UTC
- Duration: 3d 21h 27m
- Impact: none, status: resolved
- Components: —
- Canonical incident: [ci-b8f870d1369b4ec8be6fe40f](#ci-b8f870d1369b4ec8be6fe40f)

- `2026-11-06 18:38` **identified** — **Investigating — We’re seeing payment declines on some U.S. business Mastercards. Affected cards have numbers beginning with 544015, 546400, 543502, 543799, or 512998. Some cards beginning with 513989 are also affected.** **For API customers, this can prevent automatic recharges and one-time credit purchases. If your prepaid balance runs out, your API requests may fail.** **If you use one of these cards, we recommend adding a card from a different issuer before your next payment, even if you haven’t seen a decline. If an API credit purchase has already failed, use the new card to retry the purchase.** **We’re working with our payment partners and will share updates here.**
- `2026-11-06 18:50` **identified** — Investigating — We’re seeing payment declines on some U.S. business Mastercards. Affected cards have numbers beginning with 544015, 546400, 543502, 543799, or 512998. Some cards beginning with 513989 are also affected. For API customers, this can prevent automatic recharges and one-time credit purchases. If your prepaid balance runs out, your API requests may fail. If you use one of these cards, we recommend adding a card from a different issuer before your next payment, even if you haven’t seen a decline. If an API credit purchase has already failed, use the new card to retry the purchase. We’re working with our payment partners and will share updates here.
- `2026-11-07 01:29` **monitoring** — We’ve applied a mitigation and are monitoring the affected services to confirm recovery.
- `2026-11-10 16:06` **resolved** — All affected services are operating normally again. We’re sorry for the disruption, and we appreciate your patience.

### 3. cloudflare — Spectrum Availability Issue in Phoenix

- URL: https://www.cloudflarestatus.com/incidents/c60k3g4vsoc1
- Started: 2026-11-07 00:59 UTC
- Resolved: 2026-11-09 14:36 UTC
- Duration: 2d 13h 37m
- Impact: minor, status: resolved
- Components: Spectrum
- Canonical incident: [ci-5d9778b9959939de25c5186b](#ci-5d9778b9959939de25c5186b)

- `2026-11-07 00:59` **investigating** — Cloudflare is investigating issues with Spectrum in Phoenix. This may be impacting the availability of existing Spectrum applications. We are working to understand the full impact and mitigate this problem. More updates to follow shortly.
- `2026-11-07 02:51` **monitoring** — We are continuing to monitor for any further issues.
- `2026-11-09 14:36` **resolved** — This incident has been resolved.

### 4. grafana — October usage report visualization issue

- URL: https://stspg.io/ksvgzjk4ctnl
- Started: 2026-11-03 22:00 UTC
- Resolved: 2026-11-07 01:16 UTC
- Duration: 3d 3h 16m
- Impact: minor, status: resolved
- Components: Cost Management and Billing
- Canonical incident: [ci-ee5265daee2b7cfe1debc8fc](#ci-ee5265daee2b7cfe1debc8fc)

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
- Canonical incident: [ci-99928e6fdb4d472a21f59584](#ci-99928e6fdb4d472a21f59584)

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
- Components: Pipelines, R2, Stream
- Canonical incident: [ci-b60995842eae776d1ce00d68](#ci-b60995842eae776d1ce00d68)

- `2026-11-04 14:43` **resolved** — Cloudflare has resolved an issue affecting R2 and downstream services in the Western North America region. Some customers may have seen elevated latency and errors between approximately 13:00 and 22:15 UTC on November 3, 2026. This may also have affected downstream services such as Stream and Pipelines.

### 7. cloudflare — Increased Errors for Durable Objects in Western Europe

- URL: https://www.cloudflarestatus.com/incidents/oy8gxjb38yse
- Started: 2026-11-05 13:23 UTC
- Resolved: 2026-11-06 09:05 UTC
- Duration: 19h 41m
- Impact: minor, status: resolved
- Components: Durable Objects
- Canonical incident: [ci-e499d5c3e8c5b725290504ee](#ci-e499d5c3e8c5b725290504ee)

- `2026-11-05 13:23` **monitoring** — Cloudflare investigated an issue that caused an increase of errors for Durable Objects in Western Europe. The issue has been resolved and we are monitoring the results.
- `2026-11-06 09:05` **resolved** — This incident has been resolved.

### 8. openai — ChatGPT Work and Dots Errors

- URL: https://status.openai.com/incidents/01M6P50GPP1V98EF8MJ9SH393R
- Started: 2026-11-05 15:25 UTC
- Resolved: 2026-11-06 04:31 UTC
- Duration: 13h 6m
- Impact: minor, status: resolved
- Components: —
- Canonical incident: [ci-623153a47de6f89772e68e0b](#ci-623153a47de6f89772e68e0b)

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
- Components: API Requests, Actions, Git Operations, Issues, Pull Requests
- Canonical incident: [ci-ddf101ecae8d81a3f4cb4014](#ci-ddf101ecae8d81a3f4cb4014)

- `2026-11-04 09:49` **investigating** — We are investigating reports of degraded performance for Git Operations, Issues and Pull Requests
- `2026-11-04 09:50` **investigating** — Actions is experiencing degraded performance. We are continuing to investigate.
- `2026-11-04 09:53` **investigating** — Pull Requests is experiencing degraded availability. We are continuing to investigate.
- `2026-11-04 10:38` **investigating** — We're investigating elevated errors and intermittent failures affecting GitHub including Pull Requests, Git Operations, and Issues. We'll continue to keep users updated as we work toward mitigation.
- `2026-11-04 11:11` **investigating** — API Requests is experiencing degraded performance. We are continuing to investigate.
- `2026-11-04 11:17` **investigating** — We're currently working on mitigating elevated errors and intermittent failures affecting GitHub, including Pull Requests, Git Operations, and Issues. Some requests may fail or be delayed while we work to restore service. We'll provide another update as more information becomes available.
- `2026-11-04 11:30` **investigating** — The degradation affecting API Requests, Git Operations, Issues and Pull Requests has been mitigated. We are monitoring to ensure stability.
- `2026-11-04 11:43` **monitoring** — The degradation affecting Actions has been mitigated. We are monitoring to ensure stability.
- `2026-11-04 11:50` **resolved** — On November 4, 2026, beginning at approximately 09:35 UTC, customers experienced intermittent errors and delays across several GitHub services, including Pull Requests, Issues, APIs, Git operations, and Actions.<br />The disruption was caused by a shared database reaching capacity. We restored service by switching to another database host. Recovery began around 11:24 UTC, and the remaining Actions impact was mitigated by 11:43 UTC. The incident was resolved at 11:50 UTC.<br />We have since made changes to reduce database load, increased capacity, and added monitoring to help identify similar problems earlier. We are also improving our recovery procedures.

### 10. cloudflare — Cloudflare Workers and Pages build errors

- URL: https://www.cloudflarestatus.com/incidents/nbyxn84kvb6o
- Started: 2026-11-03 14:46 UTC
- Resolved: 2026-11-03 19:38 UTC
- Duration: 4h 51m
- Impact: minor, status: resolved
- Components: Workers Builds
- Canonical incident: [ci-b80cfc68de6eff3faceb1511](#ci-b80cfc68de6eff3faceb1511)

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
- Canonical incident: [ci-1c3deab43555cfcfd330441d](#ci-1c3deab43555cfcfd330441d)

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
- Canonical incident: [ci-bb759c63588fd449137d73d1](#ci-bb759c63588fd449137d73d1)

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
- Canonical incident: [ci-bf30522b3a80535ebdde0905](#ci-bf30522b3a80535ebdde0905)

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
- Canonical incident: [ci-3437bb2a93ced4ea53af24d8](#ci-3437bb2a93ced4ea53af24d8)

- `2026-11-02 22:12` **investigating** — Cloudflare is investigating an issue affecting Cloudflare Gateway DNS-over-TLS (DoT) connections over IPv4. Affected customers may experience connection timeouts and failed DNS resolution. This issue is limited to DoT over IPv4. DoT over IPv6, DNS-over-HTTPS (DoH), and standard DNS over UDP/TCP are not affected. We are working to identify the cause and restore connectivity. Further updates will be provided as more information becomes available.
- `2026-11-03 00:23` **monitoring** — Our team has been able to identify and apply an effective fix for this issue. We are currently monitoring the results.
- `2026-11-03 00:34` **resolved** — This incident has been resolved.

### 15. github — Incident with Actions and Pages

- URL: https://stspg.io/pjgptg3cfc6n
- Started: 2026-11-04 16:05 UTC
- Resolved: 2026-11-04 17:01 UTC
- Duration: 55m
- Impact: major, status: resolved
- Components: Actions, Pages, Webhooks
- Canonical incident: [ci-00167d4a31079e96ed44c453](#ci-00167d4a31079e96ed44c453)

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
- Canonical incident: [ci-dab5664e56a8af24f1a552a9](#ci-dab5664e56a8af24f1a552a9)

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
- `grafana` [Missing Metrics for Some Synthetic Monitoring Checks using the Montreal, North Virginia, and North California public probes](https://stspg.io/gg5pyd5dw55z) — unknown
- `grafana` [Grafana Cloud: k6 timeouts](https://stspg.io/lb883jzqhsdz) — 0m
- `grafana` [Synthetic Monitoring — Jakarta Public Probe Degradation](https://stspg.io/nhj64nnckg57) — 0m

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-b723223fb9dd7c34dba6ff7b"></a>

- Canonical ID: `ci-b723223fb9dd7c34dba6ff7b` (1 source records).
  - [cloudflare/wp49bp8bqo6c](../../data/2026/W45/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/wp49bp8bqo6c
- Start time basis: `vendor_reported`.

<a id="ci-b8f870d1369b4ec8be6fe40f"></a>

- Canonical ID: `ci-b8f870d1369b4ec8be6fe40f` (1 source records).
  - [openai/01M6S2FV5TZWQ753NHRB7BF5K7](../../data/2026/W45/openai.json); vendor URL: https://status.openai.com/incidents/01M6S2FV5TZWQ753NHRB7BF5K7
- Start time basis: `vendor_reported`.

<a id="ci-5d9778b9959939de25c5186b"></a>

- Canonical ID: `ci-5d9778b9959939de25c5186b` (1 source records).
  - [cloudflare/c60k3g4vsoc1](../../data/2026/W45/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/c60k3g4vsoc1
- Start time basis: `vendor_reported`.

<a id="ci-551982cdcaee5a7c4f444676"></a>

- Canonical ID: `ci-551982cdcaee5a7c4f444676` (1 source records).
  - [grafana/djhnhnnbh6tg](../../data/2026/W45/grafana.json); vendor URL: https://stspg.io/c1s5cmykytmk
- Start time basis: `vendor_reported`.

<a id="ci-ee5265daee2b7cfe1debc8fc"></a>

- Canonical ID: `ci-ee5265daee2b7cfe1debc8fc` (1 source records).
  - [grafana/sytnxyqfmbwg](../../data/2026/W45/grafana.json); vendor URL: https://stspg.io/ksvgzjk4ctnl
- Start time basis: `vendor_reported`.

<a id="ci-99928e6fdb4d472a21f59584"></a>

- Canonical ID: `ci-99928e6fdb4d472a21f59584` (1 source records).
  - [grafana/507h3wxc3mdh](../../data/2026/W45/grafana.json); vendor URL: https://stspg.io/m1k738p976cl
- Start time basis: `vendor_reported`.

<a id="ci-b60995842eae776d1ce00d68"></a>

- Canonical ID: `ci-b60995842eae776d1ce00d68` (1 source records).
  - [cloudflare/dr9s0uvzwk6n](../../data/2026/W45/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/dr9s0uvzwk6n
- Start time basis: `vendor_reported`.

<a id="ci-108f9492555e05beae982c75"></a>

- Canonical ID: `ci-108f9492555e05beae982c75` (1 source records).
  - [grafana/qcknj94tk03k](../../data/2026/W45/grafana.json); vendor URL: https://stspg.io/jmzh1gddhwsr
- Start time basis: `vendor_reported`.

<a id="ci-e499d5c3e8c5b725290504ee"></a>

- Canonical ID: `ci-e499d5c3e8c5b725290504ee` (1 source records).
  - [cloudflare/oy8gxjb38yse](../../data/2026/W45/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/oy8gxjb38yse
- Start time basis: `vendor_reported`.

<a id="ci-623153a47de6f89772e68e0b"></a>

- Canonical ID: `ci-623153a47de6f89772e68e0b` (1 source records).
  - [openai/01M6P50GPP1V98EF8MJ9SH393R](../../data/2026/W45/openai.json); vendor URL: https://status.openai.com/incidents/01M6P50GPP1V98EF8MJ9SH393R
- Start time basis: `vendor_reported`.

<a id="ci-ddf101ecae8d81a3f4cb4014"></a>

- Canonical ID: `ci-ddf101ecae8d81a3f4cb4014` (1 source records).
  - [github/smtdswd1h2w4](../../data/2026/W45/github.json); vendor URL: https://stspg.io/18blsgwnbyy6
- Start time basis: `vendor_reported`.

<a id="ci-b80cfc68de6eff3faceb1511"></a>

- Canonical ID: `ci-b80cfc68de6eff3faceb1511` (1 source records).
  - [cloudflare/nbyxn84kvb6o](../../data/2026/W45/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/nbyxn84kvb6o
- Start time basis: `vendor_reported`.

<a id="ci-1c3deab43555cfcfd330441d"></a>

- Canonical ID: `ci-1c3deab43555cfcfd330441d` (1 source records).
  - [openai/01M6FDZ89W4PKDS19K496J307G](../../data/2026/W45/openai.json); vendor URL: https://status.openai.com/incidents/01M6FDZ89W4PKDS19K496J307G
- Start time basis: `vendor_reported`.

<a id="ci-51ef4b49854b1f7c143be39e"></a>

- Canonical ID: `ci-51ef4b49854b1f7c143be39e` (1 source records).
  - [cloudflare/nbnize4e990z](../../data/2026/W45/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/nbnize4e990z
- Start time basis: `vendor_reported`.

<a id="ci-bb759c63588fd449137d73d1"></a>

- Canonical ID: `ci-bb759c63588fd449137d73d1` (1 source records).
  - [aws/3480f26e19f4b0fc](../../data/2026/W45/aws.json); vendor URL: https://health.aws.amazon.com/health/status
- Start time basis: `first_observed_update`.
- Warning: AWS onset is the first observed message, not a verified impact start

<a id="ci-bf30522b3a80535ebdde0905"></a>

- Canonical ID: `ci-bf30522b3a80535ebdde0905` (1 source records).
  - [cloudflare/j25w3rpv45gs](../../data/2026/W45/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/j25w3rpv45gs
- Start time basis: `vendor_reported`.

<a id="ci-7ff6575f6a05f497e5a926fe"></a>

- Canonical ID: `ci-7ff6575f6a05f497e5a926fe` (1 source records).
  - [github/pfbmpxxfq1cy](../../data/2026/W45/github.json); vendor URL: https://stspg.io/s5bs6jn41fwh
- Start time basis: `vendor_reported`.

<a id="ci-3437bb2a93ced4ea53af24d8"></a>

- Canonical ID: `ci-3437bb2a93ced4ea53af24d8` (1 source records).
  - [cloudflare/3bv9locsg7ln](../../data/2026/W45/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/3bv9locsg7ln
- Start time basis: `vendor_reported`.

<a id="ci-00167d4a31079e96ed44c453"></a>

- Canonical ID: `ci-00167d4a31079e96ed44c453` (1 source records).
  - [github/xr3myzj7ms51](../../data/2026/W45/github.json); vendor URL: https://stspg.io/pjgptg3cfc6n
- Start time basis: `vendor_reported`.

<a id="ci-8cc3bcff23856130d93b4f06"></a>

- Canonical ID: `ci-8cc3bcff23856130d93b4f06` (1 source records).
  - [cloudflare/3ehw20ozwlkr](../../data/2026/W45/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/3ehw20ozwlkr
- Start time basis: `vendor_reported`.

<a id="ci-cedf7e51509b2e7cf946c927"></a>

- Canonical ID: `ci-cedf7e51509b2e7cf946c927` (1 source records).
  - [grafana/zvqy74jln4zq](../../data/2026/W45/grafana.json); vendor URL: https://stspg.io/8w9r2df60c3j
- Start time basis: `vendor_reported`.

<a id="ci-ef7ecb64af5022f2a53cbf03"></a>

- Canonical ID: `ci-ef7ecb64af5022f2a53cbf03` (1 source records).
  - [github/pyynvrm1vskw](../../data/2026/W45/github.json); vendor URL: https://stspg.io/5yt31mfmyc6p
- Start time basis: `vendor_reported`.

<a id="ci-8adf8097ddd1da977b9eb104"></a>

- Canonical ID: `ci-8adf8097ddd1da977b9eb104` (1 source records).
  - [openai/01M6MBCX427A7113A3T4FJ1N6Y](../../data/2026/W45/openai.json); vendor URL: https://status.openai.com/incidents/01M6MBCX427A7113A3T4FJ1N6Y
- Start time basis: `vendor_reported`.

<a id="ci-9c3b787ab94aa76a1fb32048"></a>

- Canonical ID: `ci-9c3b787ab94aa76a1fb32048` (1 source records).
  - [openai/01M6XVDCKRT3PSB0YP8J30GTD5](../../data/2026/W45/openai.json); vendor URL: https://status.openai.com/incidents/01M6XVDCKRT3PSB0YP8J30GTD5
- Start time basis: `vendor_reported`.

<a id="ci-3fbb3295bac39c5a85b7aaab"></a>

- Canonical ID: `ci-3fbb3295bac39c5a85b7aaab` (1 source records).
  - [anthropic/kw2lxj1th04c](../../data/2026/W45/anthropic.json); vendor URL: https://stspg.io/3m3q1nwb142k
- Start time basis: `vendor_reported`.

<a id="ci-0f8eb98c1846e36363b15f62"></a>

- Canonical ID: `ci-0f8eb98c1846e36363b15f62` (1 source records).
  - [github/13b6lyl8d7bz](../../data/2026/W45/github.json); vendor URL: https://stspg.io/n478rfsspcmr
- Start time basis: `vendor_reported`.

<a id="ci-0d60c6c5d5a6e66135fa0062"></a>

- Canonical ID: `ci-0d60c6c5d5a6e66135fa0062` (1 source records).
  - [github/rxlj89wpk7px](../../data/2026/W45/github.json); vendor URL: https://stspg.io/h1xqhrcn5gqn
- Start time basis: `vendor_reported`.

<a id="ci-818179eab73280f1eaa97869"></a>

- Canonical ID: `ci-818179eab73280f1eaa97869` (1 source records).
  - [cloudflare/1i6fnludjh1m](../../data/2026/W45/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/1i6fnludjh1m
- Start time basis: `vendor_reported`.

<a id="ci-7cd851e970c06849c5507961"></a>

- Canonical ID: `ci-7cd851e970c06849c5507961` (1 source records).
  - [cloudflare/n5mvqjs26jg9](../../data/2026/W45/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/n5mvqjs26jg9
- Start time basis: `vendor_reported`.

<a id="ci-e87475178b234263b524038a"></a>

- Canonical ID: `ci-e87475178b234263b524038a` (1 source records).
  - [github/tc28c0wzrnpz](../../data/2026/W45/github.json); vendor URL: https://stspg.io/hj74883wgn0t
- Start time basis: `vendor_reported`.

<a id="ci-0d5971d3da60d247293bcf2e"></a>

- Canonical ID: `ci-0d5971d3da60d247293bcf2e` (1 source records).
  - [cloudflare/7gixemnw9oei](../../data/2026/W45/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/7gixemnw9oei
- Start time basis: `vendor_reported`.

<a id="ci-1c68c823bd889463ee239084"></a>

- Canonical ID: `ci-1c68c823bd889463ee239084` (1 source records).
  - [cloudflare/wtsd415cavkc](../../data/2026/W45/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/wtsd415cavkc
- Start time basis: `vendor_reported`.

<a id="ci-fb2d15f04e841ac0c59afca7"></a>

- Canonical ID: `ci-fb2d15f04e841ac0c59afca7` (1 source records).
  - [grafana/x7ddtqs4j7ty](../../data/2026/W45/grafana.json); vendor URL: https://stspg.io/yw5qq1khyhj1
- Start time basis: `vendor_reported`.

<a id="ci-edc6a3ee98c7677737678fee"></a>

- Canonical ID: `ci-edc6a3ee98c7677737678fee` (1 source records).
  - [cloudflare/7he1pmetw77p](../../data/2026/W45/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/7he1pmetw77p
- Start time basis: `vendor_reported`.

<a id="ci-4393029cbcc9726beef81eac"></a>

- Canonical ID: `ci-4393029cbcc9726beef81eac` (1 source records).
  - [github/t89g2q1800fr](../../data/2026/W45/github.json); vendor URL: https://stspg.io/lv4q3xm9r0rk
- Start time basis: `vendor_reported`.

<a id="ci-124cce6c54809c232ed91d85"></a>

- Canonical ID: `ci-124cce6c54809c232ed91d85` (1 source records).
  - [cloudflare/6c4ak19xodyt](../../data/2026/W45/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/6c4ak19xodyt
- Start time basis: `vendor_reported`.

<a id="ci-4e521b6710355b18e095bae7"></a>

- Canonical ID: `ci-4e521b6710355b18e095bae7` (1 source records).
  - [cloudflare/x9orajjmvvwa](../../data/2026/W45/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/x9orajjmvvwa
- Start time basis: `vendor_reported`.

<a id="ci-4347342463f26262b1f8ce2e"></a>

- Canonical ID: `ci-4347342463f26262b1f8ce2e` (1 source records).
  - [github/db4x8hh362t8](../../data/2026/W45/github.json); vendor URL: https://stspg.io/ns6pj9qy9glb
- Start time basis: `vendor_reported`.

<a id="ci-e93638bfb39d39dc20f0db5f"></a>

- Canonical ID: `ci-e93638bfb39d39dc20f0db5f` (1 source records).
  - [circleci/qqh50c0z6mn3](../../data/2026/W45/circleci.json); vendor URL: https://stspg.io/pf820dqbcwhc
- Start time basis: `vendor_reported`.

<a id="ci-dab5664e56a8af24f1a552a9"></a>

- Canonical ID: `ci-dab5664e56a8af24f1a552a9` (1 source records).
  - [github/h02n9572mx9y](../../data/2026/W45/github.json); vendor URL: https://stspg.io/7pwv28hvkmz8
- Start time basis: `vendor_reported`.

<a id="ci-e8e6f657c9134c9588762018"></a>

- Canonical ID: `ci-e8e6f657c9134c9588762018` (1 source records).
  - [cloudflare/yaa0witbl2me](../../data/2026/W45/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/yaa0witbl2me
- Start time basis: `vendor_reported`.

<a id="ci-c7fe1346874fe3901184c44f"></a>

- Canonical ID: `ci-c7fe1346874fe3901184c44f` (1 source records).
  - [cloudflare/c3pjgxk8xnzn](../../data/2026/W45/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/c3pjgxk8xnzn
- Start time basis: `vendor_reported`.

<a id="ci-fe7e58ae023f72bd7d3c88c3"></a>

- Canonical ID: `ci-fe7e58ae023f72bd7d3c88c3` (1 source records).
  - [grafana/xtcvvd5p4v73](../../data/2026/W45/grafana.json); vendor URL: https://stspg.io/gg5pyd5dw55z
- Start time basis: `vendor_reported`.
- Warning: Vendor resolved_at precedes started_at; duration is unknown, score uses zero duration; decision:legacy-time-grafana-xtcvvd5p4v73

<a id="ci-8c1e4bd5cf9adc474892658d"></a>

- Canonical ID: `ci-8c1e4bd5cf9adc474892658d` (1 source records).
  - [grafana/2l051sx0k394](../../data/2026/W45/grafana.json); vendor URL: https://stspg.io/lb883jzqhsdz
- Start time basis: `vendor_reported`.

<a id="ci-2adb2cb00d45081b2dd07173"></a>

- Canonical ID: `ci-2adb2cb00d45081b2dd07173` (1 source records).
  - [grafana/g167fqm7tf73](../../data/2026/W45/grafana.json); vendor URL: https://stspg.io/nhj64nnckg57
- Start time basis: `vendor_reported`.

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
