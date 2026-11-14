# Week 2026-W38

2026-09-14 to 2026-09-20 (UTC). Generated 2026-11-14 12:36.

57 canonical incidents, 21 above threshold. 2 of 14 sources unavailable: datadog, yandexcloud.

57 source records (59 saved revisions) → 57 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | aws | [Increased Error Rates](https://health.aws.amazon.com/health/status) | 2026-09-15 10:22 | 60d 2h 14m (ongoing) | unknown | multipleservices-me-south-1 | 173068 = 86534 × 1 × 2.0 |
| 2 | aws | [Increased Error Rates](https://health.aws.amazon.com/health/status) | 2026-09-15 10:27 | 60d 2h 9m (ongoing) | unknown | multipleservices-me-central-1 | 173058 = 86529 × 1 × 2.0 |
| 3 | cloudflare | [Cloudflare Queues issues](https://www.cloudflarestatus.com/incidents/s4x7n3nfqnyy) | 2026-09-14 20:22 | 1d 2h 8m | minor | Queues | 2822 = 1568 × 1 × 1.8 |
| 4 | cloudflare | [Replicate Pruna issue](https://www.cloudflarestatus.com/incidents/823rbdm2k25x) | 2026-09-17 01:09 | 20h 42m | minor | Replicate | 2236 = 1242 × 1 × 1.8 |
| 5 | cloudflare | [Some Replicate models are failing to start](https://www.cloudflarestatus.com/incidents/19g1m7tsvncw) | 2026-09-14 22:25 | 15h 40m | minor | Replicate | 1692 = 940 × 1 × 1.8 |
| 6 | cloudflare | [Network Performance Issues in Los Angeles](https://www.cloudflarestatus.com/incidents/qc2v4gxy3gfw) | 2026-09-16 02:03 | 14h 20m | minor | Los Angeles, CA, United States - (LAX) | 1548 = 860 × 1 × 1.8 |
| 7 | stripe | [Elevated errors on Google Pay payments](https://stspg.io/j7b7wsrfq7hm) | 2026-09-16 10:05 | 7h 46m | none | Acquirers and payment methods, Stripe core components | 1118 = 466 × 2 × 1.2 |
| 8 | openai | [Elevated errors affecting Work Mode in ChatGPT](https://status.openai.com/incidents/01M2GA8XTS6VB3QCDEGZ0HNAQ5) | 2026-09-14 15:58 | 11h 12m | minor | — | 1008 = 672 × 1 × 1.5 |
| 9 | github | [Degradation with Gemini 3.8 Flash](https://stspg.io/hsjy12f57b24) | 2026-09-16 07:20 | 10h 27m | major | Copilot AI Model Providers | 940 = 627 × 1 × 1.5 |
| 10 | openai | [Overbilling for OpenAI-hosted containers in the Agent API](https://status.openai.com/incidents/01M2VA7X37P1ASADSNZ1CG4N4D) | 2026-09-18 22:29 | 9h 22m | none | — | 843 = 562 × 1 × 1.5 |
| 11 | cloudflare | [Support Helpdesk Availability Issues](https://www.cloudflarestatus.com/incidents/zlyf7lvpp05s) | 2026-09-16 08:31 | 6h 26m | minor | Support Site | 695 = 386 × 1 × 1.8 |
| 12 | circleci | [Customers may experience delays in UI updates](https://stspg.io/40g5hyq2p6j4) | 2026-09-16 14:55 | 4h 34m | minor | CircleCI API, CircleCI UI | 548 = 274 × 2 × 1.0 |
| 13 | digitalocean | [Support Portal Service Disruption](https://stspg.io/n6hr8jd64qjl) | 2026-09-16 09:48 | 11h 21m | minor | Support Center | 545 = 681 × 1 × 0.8 |
| 14 | anthropic | [Intermittent error spikes for Claude Mythos 5.1 and Claude Fable 5.1](https://stspg.io/x1qpycpg6ln7) | 2026-09-15 09:48 | 1h 26m | major | Claude API (api.anthropic.com), Claude Code, Claude Cowork, claude.ai | 516 = 86 × 4 × 1.5 |
| 15 | stripe | [Stripe Support Systems issue](https://stspg.io/1pqrkshz4vdd) | 2026-09-16 10:05 | 7h 4m | none | Stripe core components | 509 = 424 × 1 × 1.2 |
| 16 | cloudflare | [Load Balancing Analytics and Notification Delays](https://www.cloudflarestatus.com/incidents/rnmq02lhv5w1) | 2026-09-14 23:03 | 3h 23m | minor | Load Balancing and Monitoring | 365 = 203 × 1 × 1.8 |
| 17 | cloudflare | [Issues with 1.1.1.1 public resolver on .tn queries](https://www.cloudflarestatus.com/incidents/9tgkg3ww89kf) | 2026-09-16 15:08 | 3h 11m | minor | Recursive DNS | 344 = 191 × 1 × 1.8 |
| 18 | openai | [Delayed support responses](https://status.openai.com/incidents/01M2V76GEPJB0HQGRA0QERRZ3T) | 2026-09-18 21:36 | 3h 33m | none | — | 320 = 213 × 1 × 1.5 |
| 19 | circleci | [Customers may experience delays in UI updates](https://stspg.io/0djh5p3ll96t) | 2026-09-15 14:42 | 2h 20m | minor | CircleCI UI | 140 = 140 × 1 × 1.0 |
| 20 | openai | [Elevated errors affecting ChatGPT Work mode](https://status.openai.com/incidents/01M2RC0WHGTX8EHDH7D60JJ7FC) | 2026-09-17 19:03 | 1h 21m | major | — | 122 = 81 × 1 × 1.5 |
| 21 | openai | [Elevated errors with gpt-image-2.5-flare](https://status.openai.com/incidents/01M2KQNE5C42NEZPX6V01NHH5W) | 2026-09-15 23:50 | 40m | major | — | 60 = 40 × 1 × 1.5 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. aws — Increased Error Rates

- URL: https://health.aws.amazon.com/health/status
- Started: 2026-09-15 10:22 UTC
- Resolved: not yet
- Duration: 60d 2h 14m (ongoing)
- Impact: unknown, status: investigating
- Components: multipleservices-me-south-1
- Canonical incident: [ci-a614ea8d00b0233b6712f617](#ci-a614ea8d00b0233b6712f617)

- `2026-09-15 10:22` **investigating** — We are providing an update on the disruption affecting the Middle East (Bahrain) (me-south-1) Region. The damage to our infrastructure spanned multiple Availability Zones and exceeded what our regional and multi-AZ services are designed to withstand. After a thorough assessment, we have determined that we are unable to restore access to the resources and data hosted exclusively in this Region. When the first Availability Zone was damaged in March, we began recommending that customers migrate their workloads to other Regions, and most did so before the Region became unavailable following the disruption of a second Availability Zone in April. Since then, we have supported the remaining customers in re-establishing their operations in alternate Regions, using backups where available or implementing alternative solutions to mitigate the impact of inaccessible data. In parallel, we assessed all affected infrastructure and exhausted every option for restoring data and resources that had not been migrated before the Region became unavailable. We remain committed to supporting our customers in Bahrain in the long term and will share a further update in early 2027. We have notified the relevant authorities and continue to work with them toward that goal.

### 2. aws — Increased Error Rates

- URL: https://health.aws.amazon.com/health/status
- Started: 2026-09-15 10:27 UTC
- Resolved: not yet
- Duration: 60d 2h 9m (ongoing)
- Impact: unknown, status: investigating
- Components: multipleservices-me-central-1
- Canonical incident: [ci-c45727debcd2b7b6318e58b4](#ci-c45727debcd2b7b6318e58b4)

- `2026-09-15 10:27` **investigating** — We are providing an update on the disruption affecting the Middle East (UAE) (me-central-1) Region. After a thorough assessment, we have determined that we are unable to restore access to the resources and data hosted exclusively in the mec1-az2 Availability Zone. We continue to work on recovering regional resources, as well as zonal resources hosted in the other affected Availability Zones (mec1-az1 and mec1-az3). Since the disruption began in March, most customers have been able to re-establish their operations in other Regions by restoring backups or copying data that remained accessible. AWS Support remains available to help customers who need assistance moving their applications to alternate Regions. We remain committed to supporting our customers in the UAE. We are working on replacing the affected infrastructure and will provide an update on the restoration of our services in the coming months. We have notified the relevant authorities and continue to work with them toward that goal.

### 3. cloudflare — Cloudflare Queues issues

- URL: https://www.cloudflarestatus.com/incidents/s4x7n3nfqnyy
- Started: 2026-09-14 20:22 UTC
- Resolved: 2026-09-15 22:31 UTC
- Duration: 1d 2h 8m
- Impact: minor, status: resolved
- Components: Queues
- Canonical incident: [ci-60e2b81e45fe81a886c6b491](#ci-60e2b81e45fe81a886c6b491)

- `2026-09-14 20:22` **investigating** — Cloudflare is aware of investigating issues with our Cloudflare Queues product.
- `2026-09-15 02:01` **identified** — We are investigating an issue where messages published to some queues are not being delivered to their consumers, resulting in delivery delays. Publishing to queues is unaffected, and messages are being received and safely retained — no messages are being lost. We have identified the root cause and are actively working on a fix. Affected messages will be delivered automatically once the issue is resolved.
- `2026-09-15 15:14` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-09-15 22:31` **resolved** — This incident has been resolved.

### 4. cloudflare — Replicate Pruna issue

- URL: https://www.cloudflarestatus.com/incidents/823rbdm2k25x
- Started: 2026-09-17 01:09 UTC
- Resolved: 2026-09-17 21:51 UTC
- Duration: 20h 42m
- Impact: minor, status: resolved
- Components: Replicate
- Canonical incident: [ci-243cd6c9a2a4398abcb56bab](#ci-243cd6c9a2a4398abcb56bab)

- `2026-09-17 01:09` **identified** — Some Pruna models may not be able to scale up and are therefore either experiencing long queue times or entirely unavailable. We are working with Pruna to resolve the issue.
- `2026-09-17 21:51` **resolved** — This incident has been resolved.

### 5. cloudflare — Some Replicate models are failing to start

- URL: https://www.cloudflarestatus.com/incidents/19g1m7tsvncw
- Started: 2026-09-14 22:25 UTC
- Resolved: 2026-09-15 14:06 UTC
- Duration: 15h 40m
- Impact: minor, status: resolved
- Components: Replicate
- Canonical incident: [ci-e20f71e79ef76e681160d37c](#ci-e20f71e79ef76e681160d37c)

- `2026-09-14 22:25` **investigating** — We are investigating an issue which is preventing Replicate models from failing to start new replicas.
- `2026-09-15 00:03` **identified** — Some third party models as well as https://replicate.com/qwen/qwen-edit-multiangle are affected and are unable to scale out. Please use alternate models while we work with providers to publish fixed versions.
- `2026-09-15 14:06` **resolved** — This incident has been resolved.

### 6. cloudflare — Network Performance Issues in Los Angeles

- URL: https://www.cloudflarestatus.com/incidents/qc2v4gxy3gfw
- Started: 2026-09-16 02:03 UTC
- Resolved: 2026-09-16 16:23 UTC
- Duration: 14h 20m
- Impact: minor, status: resolved
- Components: Los Angeles, CA, United States - (LAX)
- Canonical incident: [ci-ce6b8f7fa297ceee9c565401](#ci-ce6b8f7fa297ceee9c565401)

- `2026-09-16 02:03` **investigating** — Cloudflare is investigating an increased level of HTTP 5XX errors. We are working to analyse and mitigate this problem. More updates to follow shortly.
- `2026-09-16 03:44` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-09-16 16:23` **resolved** — This incident has been resolved.

### 7. stripe — Elevated errors on Google Pay payments

- URL: https://stspg.io/j7b7wsrfq7hm
- Started: 2026-09-16 10:05 UTC
- Resolved: 2026-09-16 17:52 UTC
- Duration: 7h 46m
- Impact: none, status: resolved
- Components: Acquirers and payment methods, Stripe core components
- Canonical incident: [ci-2a3bb12ed420845fa99943af](#ci-2a3bb12ed420845fa99943af)

- `2026-09-16 17:13` **monitoring** — We're currently investigating elevated errors on Google Pay payments. We'll post an update soon.
- `2026-09-16 17:20` **monitoring** — We've confirmed with the upstream partner that the Google Pay issue is recovering. We are monitoring and will provide an update when the issue is resolved.
- `2026-09-16 17:52` **resolved** — The issue with Google Pay has been resolved.

### 8. openai — Elevated errors affecting Work Mode in ChatGPT

- URL: https://status.openai.com/incidents/01M2GA8XTS6VB3QCDEGZ0HNAQ5
- Started: 2026-09-14 15:58 UTC
- Resolved: 2026-09-15 03:11 UTC
- Duration: 11h 12m
- Impact: minor, status: resolved
- Components: —
- Canonical incident: [ci-a9c5ae06af7af55dc4664414](#ci-a9c5ae06af7af55dc4664414)

- `2026-09-14 15:58` **monitoring** — Earlier today, some ChatGPT Plus users using Work Mode experienced errors starting or resuming tasks, or had limited access to workspace tools and files. We have applied a mitigation and service has recovered. We are monitoring the results.
- `2026-09-14 19:42` **monitoring** — We have applied additional mitigations for Work Mode. We are monitoring to confirm that workspace tools and files are functioning normally for ChatGPT Plus users.
- `2026-09-15 03:11` **resolved** — All impacted services have now fully recovered.

### 9. github — Degradation with Gemini 3.8 Flash

- URL: https://stspg.io/hsjy12f57b24
- Started: 2026-09-16 07:20 UTC
- Resolved: 2026-09-16 17:48 UTC
- Duration: 10h 27m
- Impact: major, status: resolved
- Components: Copilot AI Model Providers
- Canonical incident: [ci-a6fbf40baf34265e93af7396](#ci-a6fbf40baf34265e93af7396)

- `2026-09-16 07:21` **investigating** — We are investigating reports of degraded availability for Copilot AI Model Providers
- `2026-09-16 08:28` **investigating** — Copilot AI Model Providers is experiencing degraded performance. We are continuing to investigate.
- `2026-09-16 10:58` **investigating** — Copilot AI Model Providers is experiencing degraded performance. We are continuing to investigate.
- `2026-09-16 11:45` **monitoring** — The degradation affecting Copilot AI Model Providers has been mitigated. We are monitoring to ensure stability.
- `2026-09-16 17:48` **resolved** — On September 16, 2026, between 04:40 and 11:45 UTC, the Gemini 3.8 Flash model in GitHub Copilot experienced degraded availability. Requests to this model failed at an average rate of 6.4%, and the impact was highest during peak traffic hours. Other Copilot models were not affected. Users could continue to work with a different model or with 'Auto'.<br /><br />The degradation was caused due to a capacity issue with an upstream model provider. Failure rates returned to normal as the daily traffic peak passed. We monitored the model until it was healthy and resolved the incident at 17:48 UTC. We are working to make our systems resilient to cover peak demand for all the Copilot models.

### 10. openai — Overbilling for OpenAI-hosted containers in the Agent API

- URL: https://status.openai.com/incidents/01M2VA7X37P1ASADSNZ1CG4N4D
- Started: 2026-09-18 22:29 UTC
- Resolved: 2026-09-19 07:52 UTC
- Duration: 9h 22m
- Impact: none, status: resolved
- Components: —
- Canonical incident: [ci-e5b71cc5285362366ab15010](#ci-e5b71cc5285362366ab15010)

- `2026-09-18 22:29` **investigating** — We’re investigating an issue causing higher-than-expected charges for OpenAI-hosted containers in the Agent API. We’re working on a fix and preparing refunds for affected customers. We’ll share further updates as we make progress.
- `2026-09-18 23:22` **investigating** — We’re continuing to work on a fix for the container overbilling issue. We’re also reviewing affected usage to identify impacted customers and calculate refunds.
- `2026-09-19 01:08` **identified** — We are still implementing the mitigation.
- `2026-09-19 05:00` **monitoring** — We have applied the mitigation, and any new sessions would not run into this issue. We are monitoring the recovery.
- `2026-09-19 07:52` **resolved** — All impacted services have now fully recovered.

### 11. cloudflare — Support Helpdesk Availability Issues

- URL: https://www.cloudflarestatus.com/incidents/zlyf7lvpp05s
- Started: 2026-09-16 08:31 UTC
- Resolved: 2026-09-16 14:58 UTC
- Duration: 6h 26m
- Impact: minor, status: resolved
- Components: Support Site
- Canonical incident: [ci-e6ed2602ec74014c0b662cb2](#ci-e6ed2602ec74014c0b662cb2)

- `2026-09-16 08:31` **investigating** — Our helpdesk provider is currently experiencing issues, and as such customers might encounter errors while submitting or viewing tickets. Responses on customer inquiries may be delayed during this time.
- `2026-09-16 08:32` **investigating** — We are continuing to investigate this issue.
- `2026-09-16 09:11` **identified** — Support portal and live chat support are still unavailable due to the ongoing helpdesk provider outage. Customers may still experience issues submitting or viewing tickets/chats, and overall response times may be delayed. Phone support remains fully operational for urgent inquiries.
- `2026-09-16 10:21` **identified** — Support portal and live chat support are still unavailable due to the ongoing helpdesk provider outage. We highly recommend waiting for the incident resolution to submit or update cases. If you have an urgent issue, you can reach out to us on the emergency phone support line but please note that responses might be heavily delayed.
- `2026-09-16 11:43` **identified** — Support portal and live chat support remain unavailable while our helpdesk provider works toward a resolution. We strongly advise holding non-urgent cases until the incident is resolved. Please note that emergency phone support is reserved for P1 and emergency issues. Response times on this backup channel may still be heavily delayed.
- `2026-09-16 13:34` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-09-16 14:58` **resolved** — This incident has been resolved.

### 12. circleci — Customers may experience delays in UI updates

- URL: https://stspg.io/40g5hyq2p6j4
- Started: 2026-09-16 14:55 UTC
- Resolved: 2026-09-16 19:30 UTC
- Duration: 4h 34m
- Impact: minor, status: resolved
- Components: CircleCI API, CircleCI UI
- Canonical incident: [ci-a958b1ad2db72731629f9b5d](#ci-a958b1ad2db72731629f9b5d)

- `2026-09-16 14:55` **investigating** — We are investigating a delay in pipeline data becoming available. Pipelines are still running, but data about them may be delayed in the UI or API.
- `2026-09-16 15:29` **identified** — We have identified the issue and are working on a remediation. Thank you for your patience.
- `2026-09-16 17:34` **identified** — Pipeline data is now available as normal, with the exception of test results, which may still be delayed. We are continuing to work on this, thank you for your patience.
- `2026-09-16 18:22` **monitoring** — All pipeline data is now available as normal. We will continue to monitor recovery.
- `2026-09-16 19:30` **resolved** — This incident has been resolved.

### 13. digitalocean — Support Portal Service Disruption

- URL: https://stspg.io/n6hr8jd64qjl
- Started: 2026-09-16 09:48 UTC
- Resolved: 2026-09-16 21:09 UTC
- Duration: 11h 21m
- Impact: minor, status: resolved
- Components: Support Center
- Canonical incident: [ci-ce6cf37576f9fb2ca18f2510](#ci-ce6cf37576f9fb2ca18f2510)

- `2026-09-16 09:48` **investigating** — Our Support Portal is currently experiencing issues due to a third-party service disruption impacting our systems. During this time, customers may be unable to access the Support Portal or create new support tickets for their concerns. Existing support requests may also experience delayed responses from our Support team. As a workaround if you're unable to access the Support Portal or need to create a new support ticket, please submit your request via Support Contact form: https://www.digitalocean.com/company/contact/support. We apologize for the inconvenience and appreciate your patience while we work to resolve the issue. We'll provide another update as soon as we have more information.
- `2026-09-16 14:07` **monitoring** — The Support Portal and ticket creation are now accessible. Our upstream provider has applied a fix, and we've verified ticket creation and review are working as expected. We'll continue monitoring for the next 2 hours to confirm full stability and will post a final update once the incident is closed. If you're still unable to access the portal, please use our Support Contact form as a backup: https://www.digitalocean.com/company/contact/support We apologize for the disruption and appreciate your patience.
- `2026-09-16 21:09` **resolved** — From 08:06 UTC to 13:41 UTC, users may have experienced issues accessing the Support Portal, creating new support tickets, or reviewing existing support requests. Our Engineering team has confirmed the full resolution of the issue affecting the Support Portal following a fix from our upstream provider. We have verified that the Support Portal, ticket creation, and ticket review are now operating normally. If you continue to experience problems, please open a ticket with our support team. We apologize for any inconvenience and appreciate your patience.

### 14. anthropic — Intermittent error spikes for Claude Mythos 5.1 and Claude Fable 5.1

- URL: https://stspg.io/x1qpycpg6ln7
- Started: 2026-09-15 09:48 UTC
- Resolved: 2026-09-15 11:14 UTC
- Duration: 1h 26m
- Impact: major, status: resolved
- Components: Claude API (api.anthropic.com), Claude Code, Claude Cowork, claude.ai
- Canonical incident: [ci-ab8c0405cb91c1c91e3740a0](#ci-ab8c0405cb91c1c91e3740a0)

- `2026-09-15 09:48` **investigating** — We are investigating elevated errors on requests to Claude Mythos 5.1 and Claude Fable 5.1. We will provide an update as soon as possible.
- `2026-09-15 09:52` **investigating** — We are investigating elevated errors on requests to Claude Mythos 5.1 and Claude Fable 5.1. We will provide an update as soon as possible.
- `2026-09-15 10:50` **monitoring** — A fix has been deployed for Claude Mythos 5.1 and Claude Fable 5.1 and we are monitoring to ensure it has resolved the issue.
- `2026-09-15 11:14` **resolved** — A fix has been deployed for Claude Mythos 5.1 and Claude Fable 5.1 and we are monitoring to ensure it has resolved the issue.

### 15. stripe — Stripe Support Systems issue

- URL: https://stspg.io/1pqrkshz4vdd
- Started: 2026-09-16 10:05 UTC
- Resolved: 2026-09-16 17:09 UTC
- Duration: 7h 4m
- Impact: none, status: resolved
- Components: Stripe core components
- Canonical incident: [ci-0713bb8655f3b77f53b13502](#ci-0713bb8655f3b77f53b13502)

- `2026-09-16 10:05` **investigating** — We're currently investigating issues with our support systems due to a downstream partner issue. We've currently disabled live support channels and are monitoring the affected service provider. We will provide further updates once we have observed recovery.
- `2026-09-16 17:09` **resolved** — The partner issue affecting our support systems has been resolved. Stripe's live support channels have been restored and are operating normally. This issue has been resolved.

### 16. cloudflare — Load Balancing Analytics and Notification Delays

- URL: https://www.cloudflarestatus.com/incidents/rnmq02lhv5w1
- Started: 2026-09-14 23:03 UTC
- Resolved: 2026-09-15 02:27 UTC
- Duration: 3h 23m
- Impact: minor, status: resolved
- Components: Load Balancing and Monitoring
- Canonical incident: [ci-122400070332da9d34a294bc](#ci-122400070332da9d34a294bc)

- `2026-09-14 23:03` **investigating** — Cloudflare is investigating delays affecting analytics and notifications for Load Balancing health monitoring events. Load Balancing steering remains fully operational and is not impacted. We are working to resolve the issue.
- `2026-09-15 00:37` **monitoring** — Load Balancing health events analytics and notifications are now up to date. Events and notifications between 22:02 and 22:45 UTC were skipped.
- `2026-09-15 00:39` **monitoring** — We are continuing to monitor for any further issues.
- `2026-09-15 02:27` **resolved** — This incident has been resolved.

### 17. cloudflare — Issues with 1.1.1.1 public resolver on .tn queries

- URL: https://www.cloudflarestatus.com/incidents/9tgkg3ww89kf
- Started: 2026-09-16 15:08 UTC
- Resolved: 2026-09-16 18:19 UTC
- Duration: 3h 11m
- Impact: minor, status: resolved
- Components: Recursive DNS
- Canonical incident: [ci-071ecf6c7031ab0047d7c0eb](#ci-071ecf6c7031ab0047d7c0eb)

- `2026-09-16 15:08` **investigating** — Cloudflare is investigating an issue that potentially impacts users resolving queries for .tn TLD domains via the 1.1.1.1 public resolver. Further detail will be provided as more information becomes available.
- `2026-09-16 15:23` **identified** — The issue has been identified and a fix is being implemented.
- `2026-09-16 18:19` **resolved** — This incident has been resolved.

### 18. openai — Delayed support responses

- URL: https://status.openai.com/incidents/01M2V76GEPJB0HQGRA0QERRZ3T
- Started: 2026-09-18 21:36 UTC
- Resolved: 2026-09-19 01:10 UTC
- Duration: 3h 33m
- Impact: none, status: resolved
- Components: —
- Canonical incident: [ci-fa0deb7cf9036fe4e7326eef](#ci-fa0deb7cf9036fe4e7326eef)

- `2026-09-18 21:36` **identified** — Some users are experiencing delayed responses from OpenAI Support by email and chat. We’re investigating and will share updates as we learn more.
- `2026-09-18 21:59` **monitoring** — We have applied the mitigation and are monitoring the recovery. Some support responses by email and chat may still be delayed while we work through affected cases.
- `2026-09-19 01:10` **resolved** — All impacted services have now fully recovered.

### 19. circleci — Customers may experience delays in UI updates

- URL: https://stspg.io/0djh5p3ll96t
- Started: 2026-09-15 14:42 UTC
- Resolved: 2026-09-15 17:02 UTC
- Duration: 2h 20m
- Impact: minor, status: resolved
- Components: CircleCI UI
- Canonical incident: [ci-93e516397983f06663722786](#ci-93e516397983f06663722786)

- `2026-09-15 14:42` **investigating** — We are currently seeing a delay in pipeline data surfacing in the UI.
- `2026-09-15 15:14` **identified** — We are starting to see some recovery.
- `2026-09-15 17:02` **resolved** — The incident impacting pipeline data in the UI has now been resolved. We thank you for your patience while our engineers worked on reducing the delays.

### 20. openai — Elevated errors affecting ChatGPT Work mode

- URL: https://status.openai.com/incidents/01M2RC0WHGTX8EHDH7D60JJ7FC
- Started: 2026-09-17 19:03 UTC
- Resolved: 2026-09-17 20:24 UTC
- Duration: 1h 21m
- Impact: major, status: resolved
- Components: —
- Canonical incident: [ci-ec2943a36638dc057faf3eb0](#ci-ec2943a36638dc057faf3eb0)

- `2026-09-17 19:03` **investigating** — We are investigating the issue for the listed services. Some users may experience failures when starting or running Work tasks.
- `2026-09-17 19:38` **investigating** — We are investigating the issue for the listed services.
- `2026-09-17 19:57` **identified** — We are still implementing the mitigation.
- `2026-09-17 20:24` **resolved** — All impacted services have now fully recovered.

### 21. openai — Elevated errors with gpt-image-2.5-flare

- URL: https://status.openai.com/incidents/01M2KQNE5C42NEZPX6V01NHH5W
- Started: 2026-09-15 23:50 UTC
- Resolved: 2026-09-16 00:30 UTC
- Duration: 40m
- Impact: major, status: resolved
- Components: —
- Canonical incident: [ci-e078411b88d02e5f6c2951e9](#ci-e078411b88d02e5f6c2951e9)

- `2026-09-15 23:50` **investigating** — We are experiencing elevated error rates for image generation and editing requests using gpt-image-2.5-flare in the API. Our team is working to mitigate the issue. We will provide another update as more information becomes available.
- `2026-09-15 23:58` **monitoring** — We have applied the mitigation and are monitoring the recovery.
- `2026-09-16 00:30` **resolved** — All impacted services have now fully recovered.

## Below threshold (36)

- `grafana` [Increased Execution Time for Browser Checks in Synthetic Monitoring](https://stspg.io/8gz39t6rmtrv) — 1h 16m
- `cloudflare` [Increased Errors for Durable Objects and Downstream Services in Asia Pacific (APAC) Region](https://www.cloudflarestatus.com/incidents/n3n98pb4x9f4) — 1h 16m
- `cloudflare` [Cloudflare Dashboard Workers KV Issue](https://www.cloudflarestatus.com/incidents/jds6b1j6dj8q) — 1h 33m
- `yandexcloud` [Partial network unavailability](https://status.yandex.cloud/en/incidents/2023) — 1h 16m
- `cloudflare` [Network Performance Issues in Ashburn, VA, United States](https://www.cloudflarestatus.com/incidents/rhmpng7klcqw) — 1h 28m
- `github` [Disruption with some GitHub services](https://stspg.io/6pn81yd1x4hz) — 1h 30m
- `stripe` [Elevated declines on Visa](https://stspg.io/rshfm81kk6pl) — 1h 43m
- `yandexcloud` [The problem is with the Object Storage services](https://status.yandex.cloud/en/incidents/2021) — 1h 16m
- `openai` [Elevated error rates for Codex and ChatGPT Work](https://status.openai.com/incidents/01M2EWYR55J47M2BPG9WC76VEG) — 1h 10m
- `github` [Incident with Pull Requests](https://stspg.io/cx8xcm6k8f9s) — 1h 8m
- `datadog` [Delayed events triggered by emails](https://stspg.io/txhsm142w1cz) — 1h 35m
- `github` [Actions Larger Runner Jobs for some customers may be slow to start](https://stspg.io/f2b6fd4l83wy) — 55m
- `openai` [Elevated errors in ChatGPT Work](https://status.openai.com/incidents/01M2NRHTYAM3K3M93X1D068MCY) — 50m
- `github` [Disruption with some GitHub services](https://stspg.io/tbh60vgxfxj6) — 49m
- `github` [Elevated rate of errors for OpenAI models provided by Copilot](https://stspg.io/t938cy7649sn) — 49m
- `cloudflare` [Network latency affecting Singapore](https://www.cloudflarestatus.com/incidents/pj1qcy7lhm11) — 19m
- `openai` [Elevated errors from GPT-5.6 and GPT-5.6 Instant on paid plans](https://status.openai.com/incidents/01M2HYH64J2CZMZ63HEV2FCQFA) — 41m
- `circleci` [Increased wait times for macOS jobs](https://stspg.io/d885vgdcmp7p) — 1h 0m
- `anthropic` [Issues with Google Play subscriptions](https://stspg.io/wzwyxh4x9t39) — 32m
- `stripe` [Stripe Support Systems issue](https://stspg.io/h6zj34k6gg57) — 38m
- `openai` [We are seeing elevated error rates across API models](https://status.openai.com/incidents/01M2RMCS2HVBXGBFKEZ9RZR4FA) — 21m
- `cloudflare` [Email Notifications Delayed](https://www.cloudflarestatus.com/incidents/d4kgs6xmq9tl) — 17m
- `openai` [Degraded Performance affecting Agents API](https://status.openai.com/incidents/01M2H3J1D6Y7RHAP49GRWGJAY0) — 18m
- `openai` [SSO sign-in and SCIM provisioning issues](https://status.openai.com/incidents/01M2VBZB1RYSMXHZNRA25ZJ36X) — 18m
- `openai` [Ads Manager login issues](https://status.openai.com/incidents/01M2JZR57HRQTKQGTCWG31M3AS) — 1m
- `grafana` [CloudWatch request timeouts in prod-us-central-0](https://stspg.io/fbswxlld304t) — unknown
- `cloudflare` [Increased 5xx errors and latency for Los Angeles (LAX) Sept 14 22:15 - Sept 14 23:10 UTC](https://www.cloudflarestatus.com/incidents/95c8hvlcvxmj) — 0m
- `grafana` [Intermittent Metric Write Errors in GCP US Central (prod-us-central-0)](https://stspg.io/xh6frqpk1kxq) — unknown
- `cloudflare` [Network Performance Issues in Ashburn](https://www.cloudflarestatus.com/incidents/1k5ckxlq77kp) — 0m
- `cloudflare` [Increased HTTP 522 Errors in St. Louis](https://www.cloudflarestatus.com/incidents/wq8py6ffq6y7) — 0m
- `cloudflare` [Hyperdrive Elevated Origin Connection Failure Rates](https://www.cloudflarestatus.com/incidents/k4crxsx8xk36) — 0m
- `grafana` [Elevated Latency Managing Cloud Provider Integrations in GCP US Central](https://stspg.io/wrn11j3vwkpy) — unknown
- `cloudflare` [Increased HTTP Errors in GIG (Rio de Janeiro)](https://www.cloudflarestatus.com/incidents/83kl8stzw5zn) — 0m
- `cloudflare` [Network Performance Issues in Vancouver, Canada](https://www.cloudflarestatus.com/incidents/l3t86mb1xdxc) — 0m
- `cloudflare` [Network Performance Issues in Ashburn](https://www.cloudflarestatus.com/incidents/3czg2lc9y1t8) — 0m
- `cloudflare` [Elevated errors in Ashburn, VA (IAD)](https://www.cloudflarestatus.com/incidents/3rjsfq4q663b) — 0m

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-a614ea8d00b0233b6712f617"></a>

- Canonical ID: `ci-a614ea8d00b0233b6712f617` (1 source records).
  - [aws/33f3c8e8b99e2991](../../data/2026/W38/aws.json); vendor URL: https://health.aws.amazon.com/health/status
- Start time basis: `first_observed_update`.
- Warning: AWS onset is the first observed message, not a verified impact start

<a id="ci-c45727debcd2b7b6318e58b4"></a>

- Canonical ID: `ci-c45727debcd2b7b6318e58b4` (1 source records).
  - [aws/430be6afc66bab86](../../data/2026/W38/aws.json); vendor URL: https://health.aws.amazon.com/health/status
- Start time basis: `first_observed_update`.
- Warning: AWS onset is the first observed message, not a verified impact start

<a id="ci-60e2b81e45fe81a886c6b491"></a>

- Canonical ID: `ci-60e2b81e45fe81a886c6b491` (1 source records).
  - [cloudflare/s4x7n3nfqnyy](../../data/2026/W38/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/s4x7n3nfqnyy
- Start time basis: `vendor_reported`.

<a id="ci-243cd6c9a2a4398abcb56bab"></a>

- Canonical ID: `ci-243cd6c9a2a4398abcb56bab` (1 source records).
  - [cloudflare/823rbdm2k25x](../../data/2026/W38/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/823rbdm2k25x
- Start time basis: `vendor_reported`.

<a id="ci-615047b229dac127754e95df"></a>

- Canonical ID: `ci-615047b229dac127754e95df` (1 source records).
  - [grafana/7pq12nvfybqf](../../data/2026/W38/grafana.json); vendor URL: https://stspg.io/8gz39t6rmtrv
- Start time basis: `vendor_reported`.

<a id="ci-e20f71e79ef76e681160d37c"></a>

- Canonical ID: `ci-e20f71e79ef76e681160d37c` (1 source records).
  - [cloudflare/19g1m7tsvncw](../../data/2026/W38/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/19g1m7tsvncw
- Start time basis: `vendor_reported`.

<a id="ci-ce6b8f7fa297ceee9c565401"></a>

- Canonical ID: `ci-ce6b8f7fa297ceee9c565401` (1 source records).
  - [cloudflare/qc2v4gxy3gfw](../../data/2026/W38/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/qc2v4gxy3gfw
- Start time basis: `vendor_reported`.

<a id="ci-2a3bb12ed420845fa99943af"></a>

- Canonical ID: `ci-2a3bb12ed420845fa99943af` (1 source records).
  - [stripe/zmpdfkzbkjcl](../../data/2026/W38/stripe.json); vendor URL: https://stspg.io/j7b7wsrfq7hm
- Start time basis: `vendor_reported`.

<a id="ci-a9c5ae06af7af55dc4664414"></a>

- Canonical ID: `ci-a9c5ae06af7af55dc4664414` (1 source records).
  - [openai/01M2GA8XTS6VB3QCDEGZ0HNAQ5](../../data/2026/W38/openai.json); vendor URL: https://status.openai.com/incidents/01M2GA8XTS6VB3QCDEGZ0HNAQ5
- Start time basis: `vendor_reported`.

<a id="ci-96099a4d91fdeba1f8551e42"></a>

- Canonical ID: `ci-96099a4d91fdeba1f8551e42` (1 source records).
  - [cloudflare/n3n98pb4x9f4](../../data/2026/W38/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/n3n98pb4x9f4
- Start time basis: `vendor_reported`.

<a id="ci-a6fbf40baf34265e93af7396"></a>

- Canonical ID: `ci-a6fbf40baf34265e93af7396` (1 source records).
  - [github/nlxnbqnkdzdl](../../data/2026/W38/github.json); vendor URL: https://stspg.io/hsjy12f57b24
- Start time basis: `vendor_reported`.

<a id="ci-e5b71cc5285362366ab15010"></a>

- Canonical ID: `ci-e5b71cc5285362366ab15010` (1 source records).
  - [openai/01M2VA7X37P1ASADSNZ1CG4N4D](../../data/2026/W38/openai.json); vendor URL: https://status.openai.com/incidents/01M2VA7X37P1ASADSNZ1CG4N4D
- Start time basis: `vendor_reported`.

<a id="ci-e6ed2602ec74014c0b662cb2"></a>

- Canonical ID: `ci-e6ed2602ec74014c0b662cb2` (1 source records).
  - [cloudflare/zlyf7lvpp05s](../../data/2026/W38/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/zlyf7lvpp05s
- Start time basis: `vendor_reported`.

<a id="ci-a958b1ad2db72731629f9b5d"></a>

- Canonical ID: `ci-a958b1ad2db72731629f9b5d` (1 source records).
  - [circleci/3lnp36dwl73z](../../data/2026/W38/circleci.json); vendor URL: https://stspg.io/40g5hyq2p6j4
- Start time basis: `vendor_reported`.

<a id="ci-ce6cf37576f9fb2ca18f2510"></a>

- Canonical ID: `ci-ce6cf37576f9fb2ca18f2510` (1 source records).
  - [digitalocean/grb2v0td1nq5](../../data/2026/W38/digitalocean.json); vendor URL: https://stspg.io/n6hr8jd64qjl
- Start time basis: `vendor_reported`.

<a id="ci-ab8c0405cb91c1c91e3740a0"></a>

- Canonical ID: `ci-ab8c0405cb91c1c91e3740a0` (1 source records).
  - [anthropic/6304r9jjhj34](../../data/2026/W38/anthropic.json); vendor URL: https://stspg.io/x1qpycpg6ln7
- Start time basis: `vendor_reported`.

<a id="ci-0713bb8655f3b77f53b13502"></a>

- Canonical ID: `ci-0713bb8655f3b77f53b13502` (1 source records).
  - [stripe/43xfzqymv1dt](../../data/2026/W38/stripe.json); vendor URL: https://stspg.io/1pqrkshz4vdd
- Start time basis: `vendor_reported`.

<a id="ci-122400070332da9d34a294bc"></a>

- Canonical ID: `ci-122400070332da9d34a294bc` (1 source records).
  - [cloudflare/rnmq02lhv5w1](../../data/2026/W38/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/rnmq02lhv5w1
- Start time basis: `vendor_reported`.

<a id="ci-071ecf6c7031ab0047d7c0eb"></a>

- Canonical ID: `ci-071ecf6c7031ab0047d7c0eb` (1 source records).
  - [cloudflare/9tgkg3ww89kf](../../data/2026/W38/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/9tgkg3ww89kf
- Start time basis: `vendor_reported`.

<a id="ci-2fd5bfea2d9414d844ef4fce"></a>

- Canonical ID: `ci-2fd5bfea2d9414d844ef4fce` (1 source records).
  - [cloudflare/jds6b1j6dj8q](../../data/2026/W38/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/jds6b1j6dj8q
- Start time basis: `vendor_reported`.

<a id="ci-fa0deb7cf9036fe4e7326eef"></a>

- Canonical ID: `ci-fa0deb7cf9036fe4e7326eef` (1 source records).
  - [openai/01M2V76GEPJB0HQGRA0QERRZ3T](../../data/2026/W38/openai.json); vendor URL: https://status.openai.com/incidents/01M2V76GEPJB0HQGRA0QERRZ3T
- Start time basis: `vendor_reported`.

<a id="ci-9c230576af734d678c80b5cc"></a>

- Canonical ID: `ci-9c230576af734d678c80b5cc` (1 source records).
  - [yandexcloud/2023](../../data/2026/W38/yandexcloud.json); vendor URL: https://status.yandex.cloud/en/incidents/2023
- Start time basis: `vendor_reported`.

<a id="ci-a6e4bfcd702c6e3110148dac"></a>

- Canonical ID: `ci-a6e4bfcd702c6e3110148dac` (1 source records).
  - [cloudflare/rhmpng7klcqw](../../data/2026/W38/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/rhmpng7klcqw
- Start time basis: `vendor_reported`.

<a id="ci-93e516397983f06663722786"></a>

- Canonical ID: `ci-93e516397983f06663722786` (1 source records).
  - [circleci/ynpn1xvjsx86](../../data/2026/W38/circleci.json); vendor URL: https://stspg.io/0djh5p3ll96t
- Start time basis: `vendor_reported`.

<a id="ci-0cc595e1161392ee2f3c27b0"></a>

- Canonical ID: `ci-0cc595e1161392ee2f3c27b0` (1 source records).
  - [github/qwdwmtqghpk5](../../data/2026/W38/github.json); vendor URL: https://stspg.io/6pn81yd1x4hz
- Start time basis: `vendor_reported`.

<a id="ci-13a6e0272878d2a3e55dd338"></a>

- Canonical ID: `ci-13a6e0272878d2a3e55dd338` (1 source records).
  - [stripe/7h016hnd562f](../../data/2026/W38/stripe.json); vendor URL: https://stspg.io/rshfm81kk6pl
- Start time basis: `vendor_reported`.

<a id="ci-ec2943a36638dc057faf3eb0"></a>

- Canonical ID: `ci-ec2943a36638dc057faf3eb0` (1 source records).
  - [openai/01M2RC0WHGTX8EHDH7D60JJ7FC](../../data/2026/W38/openai.json); vendor URL: https://status.openai.com/incidents/01M2RC0WHGTX8EHDH7D60JJ7FC
- Start time basis: `vendor_reported`.

<a id="ci-2791a0df10e12308bc0e6ce5"></a>

- Canonical ID: `ci-2791a0df10e12308bc0e6ce5` (1 source records).
  - [yandexcloud/2021](../../data/2026/W38/yandexcloud.json); vendor URL: https://status.yandex.cloud/en/incidents/2021
- Start time basis: `vendor_reported`.

<a id="ci-273bb50a9e2f65246e6f944e"></a>

- Canonical ID: `ci-273bb50a9e2f65246e6f944e` (1 source records).
  - [openai/01M2EWYR55J47M2BPG9WC76VEG](../../data/2026/W38/openai.json); vendor URL: https://status.openai.com/incidents/01M2EWYR55J47M2BPG9WC76VEG
- Start time basis: `vendor_reported`.

<a id="ci-ac593db770dc9e1b3b31a8c1"></a>

- Canonical ID: `ci-ac593db770dc9e1b3b31a8c1` (1 source records).
  - [github/f6yrxnz5f7bs](../../data/2026/W38/github.json); vendor URL: https://stspg.io/cx8xcm6k8f9s
- Start time basis: `vendor_reported`.

<a id="ci-f2148be613263c21ebcc4243"></a>

- Canonical ID: `ci-f2148be613263c21ebcc4243` (1 source records).
  - [datadog/fsrqsbtpkp06](../../data/2026/W38/datadog.json); vendor URL: https://stspg.io/txhsm142w1cz
- Start time basis: `vendor_reported`.

<a id="ci-56b576df66fbe9242f90853a"></a>

- Canonical ID: `ci-56b576df66fbe9242f90853a` (1 source records).
  - [github/kzh5j0mt7qwc](../../data/2026/W38/github.json); vendor URL: https://stspg.io/f2b6fd4l83wy
- Start time basis: `vendor_reported`.

<a id="ci-dbe4c619cdda3c86fc9e1b23"></a>

- Canonical ID: `ci-dbe4c619cdda3c86fc9e1b23` (1 source records).
  - [openai/01M2NRHTYAM3K3M93X1D068MCY](../../data/2026/W38/openai.json); vendor URL: https://status.openai.com/incidents/01M2NRHTYAM3K3M93X1D068MCY
- Start time basis: `vendor_reported`.

<a id="ci-2dc2162f125ba3d8ae0ee0c3"></a>

- Canonical ID: `ci-2dc2162f125ba3d8ae0ee0c3` (1 source records).
  - [github/bk7zgdcq7s9t](../../data/2026/W38/github.json); vendor URL: https://stspg.io/tbh60vgxfxj6
- Start time basis: `vendor_reported`.

<a id="ci-ccb2aa60acee7d31ae649b16"></a>

- Canonical ID: `ci-ccb2aa60acee7d31ae649b16` (1 source records).
  - [github/nvy2q96mcmyg](../../data/2026/W38/github.json); vendor URL: https://stspg.io/t938cy7649sn
- Start time basis: `vendor_reported`.

<a id="ci-e9d8f8a53379caa67b237c23"></a>

- Canonical ID: `ci-e9d8f8a53379caa67b237c23` (1 source records).
  - [cloudflare/pj1qcy7lhm11](../../data/2026/W38/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/pj1qcy7lhm11
- Start time basis: `vendor_reported`.

<a id="ci-95c2317488c4381d89f99df6"></a>

- Canonical ID: `ci-95c2317488c4381d89f99df6` (1 source records).
  - [openai/01M2HYH64J2CZMZ63HEV2FCQFA](../../data/2026/W38/openai.json); vendor URL: https://status.openai.com/incidents/01M2HYH64J2CZMZ63HEV2FCQFA
- Start time basis: `vendor_reported`.

<a id="ci-a92144b1c20590764613d4d9"></a>

- Canonical ID: `ci-a92144b1c20590764613d4d9` (1 source records).
  - [circleci/q7z4zxl8mf24](../../data/2026/W38/circleci.json); vendor URL: https://stspg.io/d885vgdcmp7p
- Start time basis: `vendor_reported`.

<a id="ci-e078411b88d02e5f6c2951e9"></a>

- Canonical ID: `ci-e078411b88d02e5f6c2951e9` (1 source records).
  - [openai/01M2KQNE5C42NEZPX6V01NHH5W](../../data/2026/W38/openai.json); vendor URL: https://status.openai.com/incidents/01M2KQNE5C42NEZPX6V01NHH5W
- Start time basis: `vendor_reported`.

<a id="ci-14305937b4296e9c28a39a89"></a>

- Canonical ID: `ci-14305937b4296e9c28a39a89` (1 source records).
  - [anthropic/t8fpw9vcshl9](../../data/2026/W38/anthropic.json); vendor URL: https://stspg.io/wzwyxh4x9t39
- Start time basis: `vendor_reported`.

<a id="ci-eb703a624f6e9f3198cdab00"></a>

- Canonical ID: `ci-eb703a624f6e9f3198cdab00` (1 source records).
  - [stripe/qhc2pwg3n7t2](../../data/2026/W38/stripe.json); vendor URL: https://stspg.io/h6zj34k6gg57
- Start time basis: `vendor_reported`.

<a id="ci-32d121e4b0d5bbffded74044"></a>

- Canonical ID: `ci-32d121e4b0d5bbffded74044` (1 source records).
  - [openai/01M2RMCS2HVBXGBFKEZ9RZR4FA](../../data/2026/W38/openai.json); vendor URL: https://status.openai.com/incidents/01M2RMCS2HVBXGBFKEZ9RZR4FA
- Start time basis: `vendor_reported`.

<a id="ci-a6e9636246c8ab5086034beb"></a>

- Canonical ID: `ci-a6e9636246c8ab5086034beb` (1 source records).
  - [cloudflare/d4kgs6xmq9tl](../../data/2026/W38/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/d4kgs6xmq9tl
- Start time basis: `vendor_reported`.

<a id="ci-b5b7017bbeaf59872b329afa"></a>

- Canonical ID: `ci-b5b7017bbeaf59872b329afa` (1 source records).
  - [openai/01M2H3J1D6Y7RHAP49GRWGJAY0](../../data/2026/W38/openai.json); vendor URL: https://status.openai.com/incidents/01M2H3J1D6Y7RHAP49GRWGJAY0
- Start time basis: `vendor_reported`.

<a id="ci-6ffc049b1230b18ff9dcd804"></a>

- Canonical ID: `ci-6ffc049b1230b18ff9dcd804` (1 source records).
  - [openai/01M2VBZB1RYSMXHZNRA25ZJ36X](../../data/2026/W38/openai.json); vendor URL: https://status.openai.com/incidents/01M2VBZB1RYSMXHZNRA25ZJ36X
- Start time basis: `vendor_reported`.

<a id="ci-5a7ad323a768de7f5084b580"></a>

- Canonical ID: `ci-5a7ad323a768de7f5084b580` (1 source records).
  - [openai/01M2JZR57HRQTKQGTCWG31M3AS](../../data/2026/W38/openai.json); vendor URL: https://status.openai.com/incidents/01M2JZR57HRQTKQGTCWG31M3AS
- Start time basis: `vendor_reported`.

<a id="ci-38b70457f333d34861514ff0"></a>

- Canonical ID: `ci-38b70457f333d34861514ff0` (1 source records).
  - [grafana/jg705wgvd4bc](../../data/2026/W38/grafana.json); vendor URL: https://stspg.io/fbswxlld304t
- Start time basis: `vendor_reported`.
- Warning: Vendor resolved_at precedes started_at; duration is unknown, score uses zero duration; decision:legacy-time-grafana-jg705wgvd4bc

<a id="ci-ae2929f80707a32ef03a45a1"></a>

- Canonical ID: `ci-ae2929f80707a32ef03a45a1` (1 source records).
  - [cloudflare/95c8hvlcvxmj](../../data/2026/W38/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/95c8hvlcvxmj
- Start time basis: `vendor_reported`.

<a id="ci-7ec05373a47f681173570222"></a>

- Canonical ID: `ci-7ec05373a47f681173570222` (1 source records).
  - [grafana/jjyvkl193v79](../../data/2026/W38/grafana.json); vendor URL: https://stspg.io/xh6frqpk1kxq
- Start time basis: `vendor_reported`.
- Warning: Vendor resolved_at precedes started_at; duration is unknown, score uses zero duration; decision:legacy-time-grafana-jjyvkl193v79

<a id="ci-bdee58bcfe4ff817d4148598"></a>

- Canonical ID: `ci-bdee58bcfe4ff817d4148598` (1 source records).
  - [cloudflare/1k5ckxlq77kp](../../data/2026/W38/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/1k5ckxlq77kp
- Start time basis: `vendor_reported`.

<a id="ci-3db4afcbdac4279be2d1b885"></a>

- Canonical ID: `ci-3db4afcbdac4279be2d1b885` (1 source records).
  - [cloudflare/wq8py6ffq6y7](../../data/2026/W38/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/wq8py6ffq6y7
- Start time basis: `vendor_reported`.

<a id="ci-63c053527dd468b1cda8b364"></a>

- Canonical ID: `ci-63c053527dd468b1cda8b364` (1 source records).
  - [cloudflare/k4crxsx8xk36](../../data/2026/W38/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/k4crxsx8xk36
- Start time basis: `vendor_reported`.

<a id="ci-80b1967cf5ec6e6e4e5b8357"></a>

- Canonical ID: `ci-80b1967cf5ec6e6e4e5b8357` (1 source records).
  - [grafana/33y25k7rvqz7](../../data/2026/W38/grafana.json); vendor URL: https://stspg.io/wrn11j3vwkpy
- Start time basis: `vendor_reported`.
- Warning: Vendor resolved_at precedes started_at; duration is unknown, score uses zero duration; decision:legacy-time-grafana-33y25k7rvqz7

<a id="ci-f5c7f30640389ceda3b03cc5"></a>

- Canonical ID: `ci-f5c7f30640389ceda3b03cc5` (1 source records).
  - [cloudflare/83kl8stzw5zn](../../data/2026/W38/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/83kl8stzw5zn
- Start time basis: `vendor_reported`.

<a id="ci-e533f6278885222c911793dd"></a>

- Canonical ID: `ci-e533f6278885222c911793dd` (1 source records).
  - [cloudflare/l3t86mb1xdxc](../../data/2026/W38/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/l3t86mb1xdxc
- Start time basis: `vendor_reported`.

<a id="ci-1bf09ee3444b0d45eddc6458"></a>

- Canonical ID: `ci-1bf09ee3444b0d45eddc6458` (1 source records).
  - [cloudflare/3czg2lc9y1t8](../../data/2026/W38/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/3czg2lc9y1t8
- Start time basis: `vendor_reported`.

<a id="ci-faa6e8eb91281a0e0e089535"></a>

- Canonical ID: `ci-faa6e8eb91281a0e0e089535` (1 source records).
  - [cloudflare/3rjsfq4q663b](../../data/2026/W38/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/3rjsfq4q663b
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
