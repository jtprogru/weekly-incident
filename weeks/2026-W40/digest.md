# Week 2026-W40

2026-09-28 to 2026-10-04 (UTC). Generated 2026-11-14 12:36.

42 canonical incidents, 24 above threshold. 2 of 14 sources unavailable: datadog, yandexcloud.

42 source records (45 saved revisions) → 42 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | cloudflare | [Issues with Workers Build failing to start](https://www.cloudflarestatus.com/incidents/q2bck6r39bzk) | 2026-10-03 05:54 | 2d 1h 33m | minor | Workers Builds | 5351 = 2973 × 1 × 1.8 |
| 2 | cloudflare | [Network Performance Issues in Madrid](https://www.cloudflarestatus.com/incidents/fs7dz54l226c) | 2026-09-30 10:31 | 1d 19h 11m | minor | — | 4664 = 2591 × 1 × 1.8 |
| 3 | openai | [Elevated errors in ChatGPT Space Pages](https://status.openai.com/incidents/01M3R1T4FCR1337FBBYY9K2RPB) | 2026-09-30 02:20 | 1d 21h 24m | minor | — | 4086 = 2724 × 1 × 1.5 |
| 4 | cloudflare | [Network Performance Issues in Los Angeles (LAX)](https://www.cloudflarestatus.com/incidents/mq69fw5ykdwb) | 2026-09-30 13:49 | 5h 29m | minor | Containers, Durable Objects, Hyperdrive, R2, Stream, Vectorize | 3553 = 329 × 6 × 1.8 |
| 5 | cloudflare | [Cloudflare Workers build delays](https://www.cloudflarestatus.com/incidents/4m03612lprhr) | 2026-09-30 15:33 | 1d 5h 54m | minor | Workers Builds | 3229 = 1794 × 1 × 1.8 |
| 6 | yandexcloud | [The problem is with the services](https://status.yandex.cloud/en/incidents/2049) | 2026-09-29 12:04 | 23h 25m | minor | Monium | 2108 = 1405 × 1 × 1.5 |
| 7 | anthropic | [Delayed credits on the Claude Platform](https://stspg.io/jg26d0skzh27) | 2026-10-01 16:20 | 6h 15m | minor | Claude API (api.anthropic.com), Claude Console (platform.claude.com) | 1125 = 375 × 2 × 1.5 |
| 8 | anthropic | [Elevated errors on claude.ai, Claude Code, Claude Cowork and the Claude API](https://stspg.io/br61xzj05pp5) | 2026-09-29 14:21 | 2h 5m | major | Claude API (api.anthropic.com), Claude Code, Claude Console (platform.claude.com), Claude Cowork, claude.ai | 938 = 125 × 5 × 1.5 |
| 9 | openai | [Elevated error rates for ChatGPT Pro and Plus users](https://status.openai.com/incidents/01M3RDNY37CAPQGRKPBVMNRT4K) | 2026-09-30 05:47 | 7h 35m | minor | — | 682 = 455 × 1 × 1.5 |
| 10 | openai | [Support available via email](https://status.openai.com/incidents/01M3Q4HMNP8SZ27DTZKW6F5BJ4) | 2026-09-29 17:49 | 5h 29m | none | — | 494 = 329 × 1 × 1.5 |
| 11 | cloudflare | [Membership Permission Change Delays](https://www.cloudflarestatus.com/incidents/2h896759fg2f) | 2026-09-30 13:36 | 4h 34m | minor | API | 493 = 274 × 1 × 1.8 |
| 12 | openai | [Elevated errors across ChatGPT, Codex, and the API including the Agents API](https://status.openai.com/incidents/01M3Q4RK1SM4EMK445GGPG7C0N) | 2026-09-29 17:52 | 5h 21m | minor | — | 482 = 321 × 1 × 1.5 |
| 13 | slack | [Free Plan Customers Are Experiencing Message Failures](https://slack-status.com/2026-10/6c28dfb7b670dc02) | 2026-10-01 12:39 | 6h 20m | unknown | Messaging | 456 = 380 × 1 × 1.2 |
| 14 | grafana | [Elevated Write Latency and Errors for Grafana Cloud in GCP Saudi Arabia](https://stspg.io/hwjx5xthjp51) | 2026-10-03 13:53 | 5h 49m | none | — | 349 = 349 × 1 × 1.0 |
| 15 | aws | [Elevated packet loss](https://health.aws.amazon.com/health/status) | 2026-10-04 16:14 | 2h 41m | unknown | internetconnectivity-eu-south-2 | 322 = 161 × 1 × 2.0 |
| 16 | openai | [Elevated latency for some API requests](https://status.openai.com/incidents/01M3SMF1Q0TDCQVNYSXYG37QMR) | 2026-09-30 17:05 | 3h 30m | minor | — | 315 = 210 × 1 × 1.5 |
| 17 | github | [Actions Job Delays](https://stspg.io/kdqxfjn5qg6q) | 2026-10-01 14:47 | 3h 9m | minor | Actions | 284 = 189 × 1 × 1.5 |
| 18 | grafana | [Elevated Write Latency and Errors for Grafana Cloud Logs in Azure Netherlands](https://stspg.io/4wvrl2q1vm5f) | 2026-10-02 10:03 | 4h 0m | minor | Azure Netherlands - prod-eu-west-3 | 240 = 240 × 1 × 1.0 |
| 19 | grafana | [Elevated Latency and Errors for Grafana Cloud Logs in AWS US East (VA)](https://stspg.io/tpz4xkmrd835) | 2026-09-30 12:40 | 2h 58m | minor | AWS US East (VA) - prod-us-east-3 | 178 = 178 × 1 × 1.0 |
| 20 | datadog | [Delayed RUM sessions](https://stspg.io/gjz5frd5dyxl) | 2026-10-02 12:46 | 2h 17m | none | RUM | 137 = 137 × 1 × 1.0 |
| 21 | grafana | [Grafana Access Issues](https://stspg.io/tblzdh8qw6qb) | 2026-09-28 17:32 | 1h 49m | critical | AWS US East (VA) - prod-us-east-3 | 109 = 109 × 1 × 1.0 |
| 22 | openai | [Issues with login, signup, and ads](https://status.openai.com/incidents/01M3WCKTAN41RZ01SBRRTY4AYM) | 2026-10-01 18:46 | 53m | major | — | 80 = 53 × 1 × 1.5 |
| 23 | github | [Copilot Code Review is unable to complete reviews](https://stspg.io/phqvtmdm3y2m) | 2026-09-28 21:16 | 51m | critical | — | 76 = 51 × 1 × 1.5 |
| 24 | cloudflare | [Cloudflare CDN experiencing increase in errors.](https://www.cloudflarestatus.com/incidents/2xnmsnv8yv5x) | 2026-10-02 21:46 | 11m | critical | CDN Cache Purge, CDN/Cache | 40 = 11 × 2 × 1.8 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. cloudflare — Issues with Workers Build failing to start

- URL: https://www.cloudflarestatus.com/incidents/q2bck6r39bzk
- Started: 2026-10-03 05:54 UTC
- Resolved: 2026-10-05 07:28 UTC
- Duration: 2d 1h 33m
- Impact: minor, status: resolved
- Components: Workers Builds
- Canonical incident: [ci-528695273cb7c48118823556](#ci-528695273cb7c48118823556)

- `2026-10-03 05:54` **investigating** — Cloudflare is aware of and investigating an issue with Cloudflare Workers build failing to start. We are working to analyse and mitigate this problem. More updates to follow shortly.
- `2026-10-03 08:54` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-10-05 07:28` **resolved** — This incident has been resolved.

### 2. cloudflare — Network Performance Issues in Madrid

- URL: https://www.cloudflarestatus.com/incidents/fs7dz54l226c
- Started: 2026-09-30 10:31 UTC
- Resolved: 2026-10-02 05:42 UTC
- Duration: 1d 19h 11m
- Impact: minor, status: resolved
- Components: —
- Canonical incident: [ci-27ad44d3f4581eff490632ab](#ci-27ad44d3f4581eff490632ab)

- `2026-09-30 10:31` **investigating** — Cloudflare is investigating issues with network performance in Madrid (MAD). We are working to analyse and mitigate this problem. More updates to follow shortly.
- `2026-09-30 11:31` **identified** — The issue has been identified and a fix is being implemented.
- `2026-09-30 14:06` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-10-02 05:42` **resolved** — This incident has been resolved.

### 3. openai — Elevated errors in ChatGPT Space Pages

- URL: https://status.openai.com/incidents/01M3R1T4FCR1337FBBYY9K2RPB
- Started: 2026-09-30 02:20 UTC
- Resolved: 2026-10-01 23:44 UTC
- Duration: 1d 21h 24m
- Impact: minor, status: resolved
- Components: —
- Canonical incident: [ci-121d3799dc075ecfb7547960](#ci-121d3799dc075ecfb7547960)

- `2026-09-30 02:20` **identified** — We have identified that users may experience errors creating or interacting with Pages, including failures when using Page tools or connecting to live Page sessions. We are working on implementing a mitigation.
- `2026-09-30 06:03` **monitoring** — We have applied the mitigation and are monitoring the recovery.
- `2026-09-30 10:34` **monitoring** — All impacted services are recovered, we are continuing to monitor for any reoccurrence.
- `2026-09-30 15:00` **monitoring** — We have applied mitigations and are monitoring the recovery.
- `2026-10-01 16:51` **monitoring** — We are continuing to monitor our mitigiations. Some users may see slow or timed out results, especially for creating new pages or spaces.
- `2026-10-01 23:44` **resolved** — All impacted services have now fully recovered.

### 4. cloudflare — Network Performance Issues in Los Angeles (LAX)

- URL: https://www.cloudflarestatus.com/incidents/mq69fw5ykdwb
- Started: 2026-09-30 13:49 UTC
- Resolved: 2026-09-30 19:18 UTC
- Duration: 5h 29m
- Impact: minor, status: resolved
- Components: Containers, Durable Objects, Hyperdrive, R2, Stream, Vectorize
- Canonical incident: [ci-35a7c79177802e86306763e8](#ci-35a7c79177802e86306763e8)

- `2026-09-30 13:49` **investigating** — Cloudflare is investigating possible network congestion in Los Angeles (LAX) affecting Vectorize, Durable Objects, R2, Hyperdrive, Stream, and Cloudchamber. We are working to mitigate the impact to Internet users in the region.
- `2026-09-30 14:39` **investigating** — We are continuing to investigate this issue.
- `2026-09-30 19:18` **resolved** — This incident has been resolved.

### 5. cloudflare — Cloudflare Workers build delays

- URL: https://www.cloudflarestatus.com/incidents/4m03612lprhr
- Started: 2026-09-30 15:33 UTC
- Resolved: 2026-10-01 21:27 UTC
- Duration: 1d 5h 54m
- Impact: minor, status: resolved
- Components: Workers Builds
- Canonical incident: [ci-07a9384d13c66f040ca883f0](#ci-07a9384d13c66f040ca883f0)

- `2026-09-30 15:33` **investigating** — Cloudflare is aware of and investigating an issue with Cloudflare Workers build delays which potentially impacts multiple customers. Further detail will be provided as more information becomes available.
- `2026-09-30 17:42` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-10-01 21:27` **resolved** — This incident has been resolved.

### 6. yandexcloud — The problem is with the services

- URL: https://status.yandex.cloud/en/incidents/2049
- Started: 2026-09-29 12:04 UTC
- Resolved: 2026-09-30 11:30 UTC
- Duration: 23h 25m
- Impact: minor, status: resolved
- Components: Monium
- Canonical incident: [ci-c87399d3ddd1811f6a63b26b](#ci-c87399d3ddd1811f6a63b26b)

- `2026-09-29 12:05` **investigating** — We noticed a partial unavailability of services in the Cloud. The developers are already aware of it and are fixing it.
- `2026-09-30 13:37` **resolved** — At 15:30 Moscow time, the problem was resolved. If you still have questions about cloud resources, please contact technical support.

### 7. anthropic — Delayed credits on the Claude Platform

- URL: https://stspg.io/jg26d0skzh27
- Started: 2026-10-01 16:20 UTC
- Resolved: 2026-10-01 22:36 UTC
- Duration: 6h 15m
- Impact: minor, status: resolved
- Components: Claude API (api.anthropic.com), Claude Console (platform.claude.com)
- Canonical incident: [ci-27b4aeed894770dc5dda3175](#ci-27b4aeed894770dc5dda3175)

- `2026-10-01 16:20` **investigating** — We are investigating delays in credit purchases being reflected in account balances, resulting in some requests failing for insufficient credit balance. We will provide an additional update as soon as possible.
- `2026-10-01 18:23` **monitoring** — We have resolved the issue causing delays in credits being applied after purchase. We are monitoring closely to ensure no further issues.
- `2026-10-01 22:36` **resolved** — This issue has been resolved.

### 8. anthropic — Elevated errors on claude.ai, Claude Code, Claude Cowork and the Claude API

- URL: https://stspg.io/br61xzj05pp5
- Started: 2026-09-29 14:21 UTC
- Resolved: 2026-09-29 16:27 UTC
- Duration: 2h 5m
- Impact: major, status: resolved
- Components: Claude API (api.anthropic.com), Claude Code, Claude Console (platform.claude.com), Claude Cowork, claude.ai
- Canonical incident: [ci-4cd0d4eda962e504bab95d1c](#ci-4cd0d4eda962e504bab95d1c)

- `2026-09-29 14:21` **investigating** — We are investigating elevated error rates affecting Claude.ai (including the desktop and mobile apps), Claude Code and Claude Cowork. Users may see failed requests, errors loading or sending conversations, or be asked to sign in again; retrying may succeed. We will provide an update as soon as possible.
- `2026-09-29 14:28` **investigating** — We are investigating elevated error rates affecting Claude.ai (including the desktop and mobile apps), Claude Code and Claude Cowork. Some requests to the Claude API are also returning errors. We will provide an update as soon as possible.
- `2026-09-29 14:30` **investigating** — We are investigating reports of degraded performance affecting claude.ai, platform.claude.com, the Claude API, Claude Code, and Claude Cowork. We will provide an update as soon as possible.
- `2026-09-29 14:41` **investigating** — We have applied a mitigation and error rates have dropped substantially. Some sign-in attempts, account actions and Claude Code and Cowork sessions are still failing. We are continuing to work on a full fix.
- `2026-09-29 15:00` **investigating** — The errors that began at 14:00 UTC were mitigated at 14:36 UTC, and existing conversations are mostly working again. A second issue is preventing many users from signing in (single sign-on and Sign in with Apple are unavailable), starting new chats, voice conversations and Claude Code or Cowork sessions, making purchases, and uploading files. If you're signed in, please don't sign out. We're working on a fix and will post another update within 30 minutes.
- `2026-09-29 15:11` **monitoring** — Most services have recovered as of 14:59 UTC. Signing in, starting new chats, voice conversations, Claude Code and Cowork sessions, purchases and file uploads are working again. Some messages sent between 14:00 and 14:59 UTC may not have been saved. We are monitoring closely and will post another update within 30 minutes.
- `2026-09-29 15:30` **monitoring** — All services have been operating normally since 14:59 UTC. We are continuing to monitor.
- `2026-09-29 16:27` **resolved** — The issue affecting Claude.ai, the Claude desktop and mobile apps, Claude Code, Claude Cowork, and the Claude API has been resolved, and services have been operating normally since recovery. Impact occurred from 07:00 PT / 14:00 UTC to 07:59 PT / 14:59 UTC.

### 9. openai — Elevated error rates for ChatGPT Pro and Plus users

- URL: https://status.openai.com/incidents/01M3RDNY37CAPQGRKPBVMNRT4K
- Started: 2026-09-30 05:47 UTC
- Resolved: 2026-09-30 13:23 UTC
- Duration: 7h 35m
- Impact: minor, status: resolved
- Components: —
- Canonical incident: [ci-c4adfe6c70c9a39106c138db](#ci-c4adfe6c70c9a39106c138db)

- `2026-09-30 05:47` **identified** — We have identified that users are experiencing elevated errors for the impacted services. We are working on implementing a mitigation.
- `2026-09-30 05:56` **monitoring** — We have applied the mitigation and are monitoring the recovery.
- `2026-09-30 10:36` **monitoring** — All impacted services are recovered, we are continuing to monitor for any reoccurrence.
- `2026-09-30 13:23` **resolved** — All impacted services have now fully recovered.

### 10. openai — Support available via email

- URL: https://status.openai.com/incidents/01M3Q4HMNP8SZ27DTZKW6F5BJ4
- Started: 2026-09-29 17:49 UTC
- Resolved: 2026-09-29 23:18 UTC
- Duration: 5h 29m
- Impact: none, status: resolved
- Components: —
- Canonical incident: [ci-894814d25db2374e93595aa8](#ci-894814d25db2374e93595aa8)

- `2026-09-29 17:49` **investigating** — Our Help Center and support chat are currently unavailable. If you need assistance, you can still reach our support team by emailing [support@openai.com](mailto:support@openai.com "support@openai.com"). We apologize for the inconvenience.
- `2026-09-29 17:55` **investigating** — Our Help Center is available, but support chat remains unavailable. If you need assistance, you can browse the Help Center or email [support@openai.com](mailto:support@openai.com "support@openai.com").
- `2026-09-29 18:50` **investigating** — Our Help Center is available, but support chat remains unavailable. If you need assistance, you can browse the Help Center or email [support@openai.com](mailto:support@openai.com "support@openai.com").
- `2026-09-29 21:50` **monitoring** — We have applied the mitigation and are monitoring the recovery.
- `2026-09-29 23:18` **resolved** — The issue affecting support chat has been resolved. You can now contact our support team through the Help Center as usual. Thank you for your patience while we worked to restore service.

### 11. cloudflare — Membership Permission Change Delays

- URL: https://www.cloudflarestatus.com/incidents/2h896759fg2f
- Started: 2026-09-30 13:36 UTC
- Resolved: 2026-09-30 18:10 UTC
- Duration: 4h 34m
- Impact: minor, status: resolved
- Components: API
- Canonical incident: [ci-18dbd480aebb7dcec27ba93e](#ci-18dbd480aebb7dcec27ba93e)

- `2026-09-30 13:36` **investigating** — Cloudflare is investigating delays in applying changes to account member roles and permissions. Recently updated permissions may not be immediately reflected in some Cloudflare services. We are working to mitigate this issue.
- `2026-09-30 18:04` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-09-30 18:10` **resolved** — This incident has been resolved.

### 12. openai — Elevated errors across ChatGPT, Codex, and the API including the Agents API

- URL: https://status.openai.com/incidents/01M3Q4RK1SM4EMK445GGPG7C0N
- Started: 2026-09-29 17:52 UTC
- Resolved: 2026-09-29 23:14 UTC
- Duration: 5h 21m
- Impact: minor, status: resolved
- Components: —
- Canonical incident: [ci-662a5e44fee5b87f45b57285](#ci-662a5e44fee5b87f45b57285)

- `2026-09-29 17:52` **investigating** — We’re investigating elevated errors affecting ChatGPT, Codex, and the API. Some users may experience failed requests, difficulty logging in or signing up, and tasks that do not complete. We’ll share updates as we learn more.
- `2026-09-29 18:24` **investigating** — We’re investigating elevated errors affecting ChatGPT, Codex, and the API including the Agents API. Some users may experience failed requests, difficulty logging in or signing up, and tasks that do not complete. We’ll share updates as we learn more.
- `2026-09-29 19:08` **investigating** — We’re still investigating elevated errors affecting ChatGPT, Codex, and the API including the Agents API. Some users may experience failed requests, difficulty logging in or signing up, and tasks that do not complete. We’ll share updates as we learn more.
- `2026-09-29 19:39` **monitoring** — We have applied the mitigation and are monitoring the recovery.
- `2026-09-29 21:09` **monitoring** — We have applied the mitigation and are monitoring the recovery.
- `2026-09-29 22:03` **monitoring** — We have applied the mitigation and are monitoring the recovery.
- `2026-09-29 22:47` **monitoring** — We have applied the mitigation and are monitoring the recovery.
- `2026-09-29 23:14` **resolved** — The issue causing elevated errors across ChatGPT, Codex, and the API has been resolved, and all affected services have recovered. Thank you for your patience while we worked to restore service. The detailed Root Cause Analysis (RCA) will be published in the next 5 business days.
- `2026-10-06 08:26` **resolved** — All impacted services have now fully recovered.

### 13. slack — Free Plan Customers Are Experiencing Message Failures

- URL: https://slack-status.com/2026-10/6c28dfb7b670dc02
- Started: 2026-10-01 12:39 UTC
- Resolved: 2026-10-01 18:59 UTC
- Duration: 6h 20m
- Impact: unknown, status: resolved
- Components: Messaging
- Canonical incident: [ci-bb236aaa386a0f05c78e557f](#ci-bb236aaa386a0f05c78e557f)

- `2026-10-01 12:39` **unknown** — We're aware that some customers on our free plan are experiencing message-sending limit blocks in an unusual spike pattern, which is preventing them from sending messages. We're investigating why this message limit was triggered unexpectedly for some teams. If you're affected, you may be temporarily unable to send messages until this is resolved. We'll provide an update in 30 minutes or sooner if additional information becomes available.
- `2026-10-01 12:56` **unknown** — We identified that a third-party app is stuck in a loop that has triggered billions of messages. This is causing free plan users to exceed messaging limits. We apologize for the inconvenience and will provide another update as soon as we have more details to share.
- `2026-10-01 13:58` **unknown** — A feature toggle ramp-up caused an app integration to repeatedly post messages in a processing loop. This pushed free plan customers over the message volume threshold, triggering the sending block. We're working to restore access and are actively working on remediation options for impacted Workspaces. We apologize for the inconvenience and will provide another update as soon as we have more details to share.
- `2026-10-01 14:31` **unknown** — We’ve identified a potential fix for this issue and we’re working on getting this out to users as quickly as possible. We’ll be back once we have another update to share. Thank you for your continued patience as we work towards a solution.
- `2026-10-01 18:52` **unknown** — Deployment of the fix is underway to restore access for impacted free plan workspaces, and we are actively monitoring the environment. More information will follow as we progress through this stage of remediation.
- `2026-10-01 18:59` **unknown** — A messaging integration briefly caused some free-plan workspaces to exceed their daily messaging limit, blocking them from sending messages. We took action to lift the sending block for all affected free-plan workspaces, and have confirmed messaging is now restored. We apologize for how this incident affected you and your business. We will undertake a full investigation of the incident, establishing the technical trigger, the underlying cause, and preventive action to avoid a repeat in the future.

### 14. grafana — Elevated Write Latency and Errors for Grafana Cloud in GCP Saudi Arabia

- URL: https://stspg.io/hwjx5xthjp51
- Started: 2026-10-03 13:53 UTC
- Resolved: 2026-10-03 19:43 UTC
- Duration: 5h 49m
- Impact: none, status: resolved
- Components: —
- Canonical incident: [ci-1ccfeeb922c4e91ce38014bc](#ci-1ccfeeb922c4e91ce38014bc)

- `2026-10-03 13:53` **investigating** — Between 12:44 and 13:15 UTC, Grafana Cloud experienced elevated write latency and errors across multiple products, including Logs, Metrics, and Traces. Services have recovered and the underlying issue is being investigated and actively monitored.
- `2026-10-03 13:54` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-10-03 15:07` **monitoring** — The issue is currently resolved but we are continuing to investigate and monitor things at this time.
- `2026-10-03 19:43` **resolved** — We've identified the issue as related to an increase in latency and errors rates with our cloud provider's infrastructure. We are working with the provider on a root cause analysis. The service has recovered and is operating normally.

### 15. aws — Elevated packet loss

- URL: https://health.aws.amazon.com/health/status
- Started: 2026-10-04 16:14 UTC
- Resolved: 2026-10-04 18:56 UTC
- Duration: 2h 41m
- Impact: unknown, status: resolved
- Components: internetconnectivity-eu-south-2
- Canonical incident: [ci-dbfa60ace7c4652365f5108f](#ci-dbfa60ace7c4652365f5108f)

- `2026-10-04 16:14` **investigating** — We are investigating elevated packet loss in the EU-SOUTH-2 Region.
- `2026-10-04 16:36` **investigating** — We are experiencing elevated packet loss with a single Availability Zone (eus2-az1) in the EU-SOUTH-2 Region. This issue is also resulting in elevated API latencies and error rates for other workflows and AWS Services. Engineers were automatically engaged and are working to identify the root cause and to mitigate the issue. We will provide you with another update by 10:15 AM PDT or sooner as new information becomes available.
- `2026-10-04 17:06` **investigating** — Our efforts to mitigate the elevated packet loss with a single Availability Zone (eus2-az1) in the EU-SOUTH-2 Region is still underway. We are seeing some signs of improvement as a result of multiple parallel efforts. We continue to work to confirm root cause and fully mitigate the issue. While we currently do not have an estimate for how long our actions will take, we will keep you informed of our progress and provide you with another update by 11:00 AM PDT or sooner as new information becomes available.
- `2026-10-04 17:56` **investigating** — We have identified the root cause and are currently focusing our mitigation efforts to drive full recovery. While we don't have an exact estimate for how long our current efforts will take, we do expect to see incremental improvements as our efforts proceed. We will provide you with another update by 11:45 AM PDT, or sooner if new information becomes available.
- `2026-10-04 18:45` **investigating** — We are seeing early signs of recovery. We will provide you with another update by 12:00 PM PDT.
- `2026-10-04 18:56` **resolved** — Between 8:48 AM and 11:28 AM PDT, we experienced elevated packet loss within a single Availability Zone (eus2-az1) in the EU-SOUTH-2 Region. This issue resulted in elevated API latencies and error rates for multiple workflows and AWS Services. We were automatically alerted to the issue at 8:54 AM, and immediately began work to identify the root cause while taking multiple parallel paths to mitigate the issue. By 10:00 AM, our mitigation efforts had led to some improvements. By 10:38 AM, we identified the root cause as related to a networking configuration change that we were able to revert the change which led to full mitigation by 11:28 AM. The issue has been resolved and the service is operating normally.

### 16. openai — Elevated latency for some API requests

- URL: https://status.openai.com/incidents/01M3SMF1Q0TDCQVNYSXYG37QMR
- Started: 2026-09-30 17:05 UTC
- Resolved: 2026-09-30 20:35 UTC
- Duration: 3h 30m
- Impact: minor, status: resolved
- Components: —
- Canonical incident: [ci-07719c8b3d8c2c332a478c63](#ci-07719c8b3d8c2c332a478c63)

- `2026-09-30 17:05` **identified** — We have identified that users are experiencing elevated latency for the impacted services. We are working on implementing a mitigation.
- `2026-09-30 17:13` **identified** — We have identified that users are experiencing elevated latency for the impacted services. We are working on implementing a mitigation.
- `2026-09-30 17:51` **identified** — We have implemented an initial mitigation and are seeing improvements in API latency. We are continuing to investigate the scope of the issue and apply additional mitigations as needed. Some requests to the Responses and Chat Completions APIs may still experience elevated latency or timeouts.
- `2026-09-30 18:21` **monitoring** — We have applied the mitigation and are monitoring the recovery.
- `2026-09-30 20:35` **resolved** — All impacted services have now fully recovered.

### 17. github — Actions Job Delays

- URL: https://stspg.io/kdqxfjn5qg6q
- Started: 2026-10-01 14:47 UTC
- Resolved: 2026-10-01 17:56 UTC
- Duration: 3h 9m
- Impact: minor, status: resolved
- Components: Actions
- Canonical incident: [ci-431893aa76e4e2782fc58053](#ci-431893aa76e4e2782fc58053)

- `2026-10-01 14:47` **investigating** — We are investigating reports of degraded performance for Actions
- `2026-10-01 14:54` **investigating** — We have identified the cause of increased Actions run start delays on Ubuntu runners and are actively deploying a fix. Customers may continue to experience intermittent delays while the mitigation rolls out and service metrics return to normal.
- `2026-10-01 15:02` **investigating** — Run-start delays on Ubuntu runners have been resolved. We are investigating run-start delays on Windows runners and will share more information as it becomes available.
- `2026-10-01 15:10` **monitoring** — The degradation affecting Actions has been mitigated. We are monitoring to ensure stability.
- `2026-10-01 15:20` **investigating** — Actions is experiencing degraded performance. We are continuing to investigate.
- `2026-10-01 15:29` **investigating** — We are seeing a reoccurrence in run-start delays, and are continuing to investigate to issue. Customers will potentially experience delays of up to ten minutes.
- `2026-10-01 16:10` **investigating** — We have identified an issue with our upstream provider which is causing Actions requests to 429 which is creating the delays. We have escalated to the owning team and are investigating how to mitigate the 429s.
- `2026-10-01 16:48` **investigating** — We are currently applying a mitigation and anticipate recovery within thirty minutes.
- `2026-10-01 17:50` **investigating** — GitHub Actions experienced degraded performance for some hosted runners due to throttling within an upstream Azure dependency. Service capacity has recovered, and we are continuing to monitor while working with Azure on the underlying condition.
- `2026-10-01 17:56` **resolved** — On October 1, 2026, GitHub Actions experienced two periods of delayed and failed job starts on GitHub-hosted runners, beginning at approximately 14:15 UTC. Across the 14:15–17:56 UTC incident-response window, approximately 0.68% of workflow runs experienced start delays, and 0.196% of jobs failed to start. <br /><br />A routing change at our cloud provider concentrated API traffic on a single backend, causing an increase in rate-limit errors. Incorrect retry behavior in the provider’s inventory service added load to the affected APIs, compounding the rate limiting. This disrupted runner provisioning and reduced available capacity in affected regions. The resulting increase in demand in unaffected regions exposed a concurrency limit in our job-assignment system, delaying jobs even where runners were available. <br /><br />During the response, we increased job-assignment concurrency, scaled runner capacity, and worked with our cloud provider to reduce API load by excluding our runner subscriptions from its inventory service. By approximately 16:15 UTC, jobs were starting and executing as expected. We continued mitigation and monitoring until the incident was resolved at 17:56 UTC. <br /><br />Our cloud provider has since increased resource-management API rate limits for our workloads. To reduce the likelihood and impact of similar incidents, we are also adding more granular rate-limiting alerts and updating our job-assignment system to handle increased demand across fewer available regions.

### 18. grafana — Elevated Write Latency and Errors for Grafana Cloud Logs in Azure Netherlands

- URL: https://stspg.io/4wvrl2q1vm5f
- Started: 2026-10-02 10:03 UTC
- Resolved: 2026-10-02 14:03 UTC
- Duration: 4h 0m
- Impact: minor, status: resolved
- Components: Azure Netherlands - prod-eu-west-3
- Canonical incident: [ci-607fdcf14a52aff5e91bfd70](#ci-607fdcf14a52aff5e91bfd70)

- `2026-10-02 11:03` **investigating** — We are investigating elevated write latency and intermittent errors affecting Grafana Cloud Logs in Azure Netherlands. Customers may see delayed log ingestion and some failed write requests. Next update by 11:30 UTC.
- `2026-10-02 11:06` **investigating** — We are continuing to investigate this issue. Next update by 11:30 UTC.
- `2026-10-02 14:02` **identified** — We have identified the issue as related to a cloud provider dependency affecting the write path for Grafana Cloud Logs in Azure Netherlands. Latency has recovered to normal levels. We are continuing to monitor before fully closing this incident.
- `2026-10-02 14:03` **resolved** — Between 10:03 and 12:37 UTC, Grafana Cloud Logs in Azure Netherlands experienced elevated write latency and intermittent errors due to a cloud provider dependency issue. Service has fully recovered and is operating normally. Thank you for your patience.

### 19. grafana — Elevated Latency and Errors for Grafana Cloud Logs in AWS US East (VA)

- URL: https://stspg.io/tpz4xkmrd835
- Started: 2026-09-30 12:40 UTC
- Resolved: 2026-09-30 15:39 UTC
- Duration: 2h 58m
- Impact: minor, status: resolved
- Components: AWS US East (VA) - prod-us-east-3
- Canonical incident: [ci-183e2ffae08987f9f41809c7](#ci-183e2ffae08987f9f41809c7)

- `2026-09-30 12:40` **investigating** — We are investigating elevated latency and a small number of failed requests for Grafana Cloud Logs in AWS US East (VA). Some customers writing logs to this region may see delayed ingestion or occasional write errors. We are actively investigating and will share more as we learn it. Next update by 13:38 UTC.
- `2026-09-30 13:40` **investigating** — We are continuing to investigate this issue.
- `2026-09-30 13:43` **identified** — We have identified a contributing factor to the elevated latency and are rolling out a mitigation to increase write capacity in AWS US East (VA). We continue to investigate a related, longer-running pattern. Next update by 14:40 UTC.
- `2026-09-30 14:31` **identified** — An additional brief recurrence affected a small number of requests after our mitigation was deployed, though impact was smaller than earlier today. We continue to monitor and investigate a related, longer-running pattern. Next update by 15:40 UTC.
- `2026-09-30 15:39` **resolved** — Between 12:09 and 15:39 UTC on September 30, Grafana Cloud Logs in AWS US East (VA) experienced intermittent elevated latency and a small number of failed requests. Impact has returned to normal levels and all services are operating as expected. Thank you for your patience.

### 20. datadog — Delayed RUM sessions

- URL: https://stspg.io/gjz5frd5dyxl
- Started: 2026-10-02 12:46 UTC
- Resolved: 2026-10-02 15:03 UTC
- Duration: 2h 17m
- Impact: none, status: resolved
- Components: RUM
- Canonical incident: [ci-7a41712fe63460445d7134fa](#ci-7a41712fe63460445d7134fa)

- `2026-10-02 12:46` **investigating** — We are investigating increased latency processing RUM sessions. As a result of this issue, some users may see gaps or delays in RUM graphs as well as empty or partial query results on RUM Sessions, RUM Analytics, and RUM Application pages since Oct 2, 2026, 11:42 AM UTC To prevent false monitor alerts due to delayed data, monitors affected by the delay will not notify and will automatically resume once current data is available. All other monitors will operate normally.
- `2026-10-02 13:23` **monitoring** — The issue has been identified we're seeing signs of recovery and we are monitoring the issue.
- `2026-10-02 13:47` **monitoring** — Latency has now resolved and monitors evaluation is being back to normal. We are still monitoring the issue for now.
- `2026-10-02 15:03` **resolved** — This incident has been resolved.

### 21. grafana — Grafana Access Issues

- URL: https://stspg.io/tblzdh8qw6qb
- Started: 2026-09-28 17:32 UTC
- Resolved: 2026-09-28 19:21 UTC
- Duration: 1h 49m
- Impact: critical, status: resolved
- Components: AWS US East (VA) - prod-us-east-3
- Canonical incident: [ci-3b6ba69789cecdb4accf23ce](#ci-3b6ba69789cecdb4accf23ce)

- `2026-09-28 17:32` **investigating** — We are currently investigating an issue causing issues accessing the platform.
- `2026-09-28 18:00` **investigating** — We are continuing to investigate the issue, and will provide updates once we have more to share.
- `2026-09-28 18:38` **identified** — We have identified the issue, and put a mitigation in place. We continue to work towards a resolution
- `2026-09-28 19:21` **resolved** — This incident has been resolved.

### 22. openai — Issues with login, signup, and ads

- URL: https://status.openai.com/incidents/01M3WCKTAN41RZ01SBRRTY4AYM
- Started: 2026-10-01 18:46 UTC
- Resolved: 2026-10-01 19:39 UTC
- Duration: 53m
- Impact: major, status: resolved
- Components: —
- Canonical incident: [ci-11252a6149f2d714ef40fdb4](#ci-11252a6149f2d714ef40fdb4)

- `2026-10-01 18:46` **investigating** — We are investigating an issue affecting login and signup. We will share further updates as we learn more.
- `2026-10-01 19:06` **investigating** — We are investigating issues affecting login, signup, and ads
- `2026-10-01 19:17` **identified** — We have identified that users are experiencing login issues for the impacted services. We are working on implementing a mitigation.
- `2026-10-01 19:26` **monitoring** — We have applied the mitigation and are monitoring the recovery.
- `2026-10-01 19:39` **resolved** — All impacted services have now fully recovered.

### 23. github — Copilot Code Review is unable to complete reviews

- URL: https://stspg.io/phqvtmdm3y2m
- Started: 2026-09-28 21:16 UTC
- Resolved: 2026-09-28 22:08 UTC
- Duration: 51m
- Impact: critical, status: resolved
- Components: —
- Canonical incident: [ci-55eca2f1a8c0e00061c37085](#ci-55eca2f1a8c0e00061c37085)

- `2026-09-28 21:16` **investigating** — We are investigating reports of impacted performance for some GitHub services.
- `2026-09-28 21:23` **investigating** — Revert of the impacted change is in flight, expect full recovery once the deployment is done.
- `2026-09-28 22:08` **investigating** — Copilot Code review has recovered, for reviews that were previously not successful you can re-request review from your PR.
- `2026-09-28 22:08` **resolved** — On September 28, 2026, between 19:23 UTC and 22:08 UTC the Copilot Code Review service was degraded and pull request reviews did not complete in all environments. On average, approximately 70% of requested reviews did not complete. Reviews requested during this window were not automatically retried, and users needed to re-request a review from their pull request. This was due to a change in a dependent service that stopped delivering data Copilot Code Review needs to complete a review. Review completion is an inherently delayed signal, because each review job takes time to start and finish, so it took longer to detect the impact.<br /><br />We mitigated the incident by reverting the change.<br /><br />We are working to add a faster signal for reviews that fail to complete, add deployment validation that confirms a Copilot Code Review completes end to end, improve how we track internal dependencies between services, and make our deployment pipeline more reliable to reduce our time to detection and mitigation of issues like this one in the future.

### 24. cloudflare — Cloudflare CDN experiencing increase in errors.

- URL: https://www.cloudflarestatus.com/incidents/2xnmsnv8yv5x
- Started: 2026-10-02 21:46 UTC
- Resolved: 2026-10-02 21:57 UTC
- Duration: 11m
- Impact: critical, status: resolved
- Components: CDN Cache Purge, CDN/Cache
- Canonical incident: [ci-f39b9f6fa3da1e2121b651e0](#ci-f39b9f6fa3da1e2121b651e0)

- `2026-10-02 21:46` **monitoring** — Cloudflare is investigating an increase in errors for CDN configuration changes and purge calls specifically issued from the dashboard. Purges sent via API were unaffected. CDN configuration changes across the dashboard, API and others saw impact.
- `2026-10-02 21:57` **resolved** — This incident has been resolved.

## Below threshold (18)

- `datadog` [Delayed Events](https://stspg.io/8fknzl9hmvs7) — 48m
- `stripe` [Visa authorization errors](https://stspg.io/t5mhk2s9lbb1) — 1h 39m
- `circleci` [Delays starting Gen 2 Docker Jobs](https://stspg.io/4jm2frmlcq2q) — 1h 39m
- `openai` [Issues with data analysis and file creation in ChatGPT](https://status.openai.com/incidents/01M3YNNVTSCXS8YM2V1YMHMVSE) — 48m
- `cloudflare` [OTP Deliverability Errors](https://www.cloudflarestatus.com/incidents/vcstgt8gqkyk) — 39m
- `grafana` [Degraded data source queries](https://stspg.io/zf8d8bx0228f) — 13m
- `cloudflare` [Increased Errors for Durable Objects and R2](https://www.cloudflarestatus.com/incidents/hyv5p7wn1c8v) — 18m
- `cloudflare` [Replicate Flux Hotswap Models Stuck](https://www.cloudflarestatus.com/incidents/smn1d93zd3hr) — 17m
- `github` [Elevated request latency](https://stspg.io/cz0ngh11q0tf) — 20m
- `stripe` [Elevated API errors](https://stspg.io/0cknvjm6mzq6) — 24m
- `grafana` [Elevated Log Ingestion Errors for Grafana Cloud Logs in Azure Netherlands (EU)](https://stspg.io/scrxwgzkx24t) — 0m
- `cloudflare` [Network Performance Issues in Eastern North America](https://www.cloudflarestatus.com/incidents/32sg73xwyq89) — 0m
- `cloudflare` [WARP connectivity in Newark, NJ](https://www.cloudflarestatus.com/incidents/0dq5fbp2xwbj) — 0m
- `github` [[Retroactive] Actions workflow run failures after deployment gate approvals](https://stspg.io/f0yg8jk2gm80) — unknown
- `cloudflare` [Network Performance Issues in Columbus colo in Ohio](https://www.cloudflarestatus.com/incidents/n2645m15vwc9) — 0m
- `cloudflare` [Intermittent network connectivity issues in Mumbai](https://www.cloudflarestatus.com/incidents/4c842mwy6z4l) — 0m
- `cloudflare` [Network Performance Issues - Seoul (ICN)](https://www.cloudflarestatus.com/incidents/0kj5cjjq93xn) — unknown
- `digitalocean` [A Network Disruption in BLR1 Region Impacted Multiple Products (Resolved)](https://stspg.io/kj9rd8tnd1wt) — 0m

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-528695273cb7c48118823556"></a>

- Canonical ID: `ci-528695273cb7c48118823556` (1 source records).
  - [cloudflare/q2bck6r39bzk](../../data/2026/W40/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/q2bck6r39bzk
- Start time basis: `vendor_reported`.

<a id="ci-27ad44d3f4581eff490632ab"></a>

- Canonical ID: `ci-27ad44d3f4581eff490632ab` (1 source records).
  - [cloudflare/fs7dz54l226c](../../data/2026/W40/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/fs7dz54l226c
- Start time basis: `vendor_reported`.

<a id="ci-121d3799dc075ecfb7547960"></a>

- Canonical ID: `ci-121d3799dc075ecfb7547960` (1 source records).
  - [openai/01M3R1T4FCR1337FBBYY9K2RPB](../../data/2026/W40/openai.json); vendor URL: https://status.openai.com/incidents/01M3R1T4FCR1337FBBYY9K2RPB
- Start time basis: `vendor_reported`.

<a id="ci-35a7c79177802e86306763e8"></a>

- Canonical ID: `ci-35a7c79177802e86306763e8` (1 source records).
  - [cloudflare/mq69fw5ykdwb](../../data/2026/W40/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/mq69fw5ykdwb
- Start time basis: `vendor_reported`.

<a id="ci-07a9384d13c66f040ca883f0"></a>

- Canonical ID: `ci-07a9384d13c66f040ca883f0` (1 source records).
  - [cloudflare/4m03612lprhr](../../data/2026/W40/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/4m03612lprhr
- Start time basis: `vendor_reported`.

<a id="ci-c87399d3ddd1811f6a63b26b"></a>

- Canonical ID: `ci-c87399d3ddd1811f6a63b26b` (1 source records).
  - [yandexcloud/2049](../../data/2026/W40/yandexcloud.json); vendor URL: https://status.yandex.cloud/en/incidents/2049
- Start time basis: `vendor_reported`.

<a id="ci-27b4aeed894770dc5dda3175"></a>

- Canonical ID: `ci-27b4aeed894770dc5dda3175` (1 source records).
  - [anthropic/k0h22tsvnydg](../../data/2026/W40/anthropic.json); vendor URL: https://stspg.io/jg26d0skzh27
- Start time basis: `vendor_reported`.

<a id="ci-4cd0d4eda962e504bab95d1c"></a>

- Canonical ID: `ci-4cd0d4eda962e504bab95d1c` (1 source records).
  - [anthropic/4xvtc2gnq73l](../../data/2026/W40/anthropic.json); vendor URL: https://stspg.io/br61xzj05pp5
- Start time basis: `vendor_reported`.

<a id="ci-c4adfe6c70c9a39106c138db"></a>

- Canonical ID: `ci-c4adfe6c70c9a39106c138db` (1 source records).
  - [openai/01M3RDNY37CAPQGRKPBVMNRT4K](../../data/2026/W40/openai.json); vendor URL: https://status.openai.com/incidents/01M3RDNY37CAPQGRKPBVMNRT4K
- Start time basis: `vendor_reported`.

<a id="ci-894814d25db2374e93595aa8"></a>

- Canonical ID: `ci-894814d25db2374e93595aa8` (1 source records).
  - [openai/01M3Q4HMNP8SZ27DTZKW6F5BJ4](../../data/2026/W40/openai.json); vendor URL: https://status.openai.com/incidents/01M3Q4HMNP8SZ27DTZKW6F5BJ4
- Start time basis: `vendor_reported`.

<a id="ci-18dbd480aebb7dcec27ba93e"></a>

- Canonical ID: `ci-18dbd480aebb7dcec27ba93e` (1 source records).
  - [cloudflare/2h896759fg2f](../../data/2026/W40/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/2h896759fg2f
- Start time basis: `vendor_reported`.

<a id="ci-662a5e44fee5b87f45b57285"></a>

- Canonical ID: `ci-662a5e44fee5b87f45b57285` (1 source records).
  - [openai/01M3Q4RK1SM4EMK445GGPG7C0N](../../data/2026/W40/openai.json); vendor URL: https://status.openai.com/incidents/01M3Q4RK1SM4EMK445GGPG7C0N
- Start time basis: `vendor_reported`.

<a id="ci-bb236aaa386a0f05c78e557f"></a>

- Canonical ID: `ci-bb236aaa386a0f05c78e557f` (1 source records).
  - [slack/1579](../../data/2026/W40/slack.json); vendor URL: https://slack-status.com/2026-10/6c28dfb7b670dc02
- Start time basis: `vendor_reported`.

<a id="ci-1ccfeeb922c4e91ce38014bc"></a>

- Canonical ID: `ci-1ccfeeb922c4e91ce38014bc` (1 source records).
  - [grafana/1hsrzg6rs2bb](../../data/2026/W40/grafana.json); vendor URL: https://stspg.io/hwjx5xthjp51
- Start time basis: `vendor_reported`.

<a id="ci-dbfa60ace7c4652365f5108f"></a>

- Canonical ID: `ci-dbfa60ace7c4652365f5108f` (1 source records).
  - [aws/604a6eabb58b32e8](../../data/2026/W40/aws.json); vendor URL: https://health.aws.amazon.com/health/status
- Start time basis: `first_observed_update`.
- Warning: AWS onset is the first observed message, not a verified impact start

<a id="ci-07719c8b3d8c2c332a478c63"></a>

- Canonical ID: `ci-07719c8b3d8c2c332a478c63` (1 source records).
  - [openai/01M3SMF1Q0TDCQVNYSXYG37QMR](../../data/2026/W40/openai.json); vendor URL: https://status.openai.com/incidents/01M3SMF1Q0TDCQVNYSXYG37QMR
- Start time basis: `vendor_reported`.

<a id="ci-431893aa76e4e2782fc58053"></a>

- Canonical ID: `ci-431893aa76e4e2782fc58053` (1 source records).
  - [github/2dpbcq5j165n](../../data/2026/W40/github.json); vendor URL: https://stspg.io/kdqxfjn5qg6q
- Start time basis: `vendor_reported`.

<a id="ci-93dbdc5c27148b77e96326b1"></a>

- Canonical ID: `ci-93dbdc5c27148b77e96326b1` (1 source records).
  - [datadog/41hx12qtkl7x](../../data/2026/W40/datadog.json); vendor URL: https://stspg.io/8fknzl9hmvs7
- Start time basis: `vendor_reported`.

<a id="ci-607fdcf14a52aff5e91bfd70"></a>

- Canonical ID: `ci-607fdcf14a52aff5e91bfd70` (1 source records).
  - [grafana/9kg64q3pygpx](../../data/2026/W40/grafana.json); vendor URL: https://stspg.io/4wvrl2q1vm5f
- Start time basis: `vendor_reported`.

<a id="ci-183e2ffae08987f9f41809c7"></a>

- Canonical ID: `ci-183e2ffae08987f9f41809c7` (1 source records).
  - [grafana/pypb5l7bbsn6](../../data/2026/W40/grafana.json); vendor URL: https://stspg.io/tpz4xkmrd835
- Start time basis: `vendor_reported`.

<a id="ci-7a41712fe63460445d7134fa"></a>

- Canonical ID: `ci-7a41712fe63460445d7134fa` (1 source records).
  - [datadog/yrkfn68nnnnr](../../data/2026/W40/datadog.json); vendor URL: https://stspg.io/gjz5frd5dyxl
- Start time basis: `vendor_reported`.

<a id="ci-7a0fbcc0b5d0293fc9b043ff"></a>

- Canonical ID: `ci-7a0fbcc0b5d0293fc9b043ff` (1 source records).
  - [stripe/drn27663z90v](../../data/2026/W40/stripe.json); vendor URL: https://stspg.io/t5mhk2s9lbb1
- Start time basis: `vendor_reported`.

<a id="ci-3b6ba69789cecdb4accf23ce"></a>

- Canonical ID: `ci-3b6ba69789cecdb4accf23ce` (1 source records).
  - [grafana/2grcbfvdxsf7](../../data/2026/W40/grafana.json); vendor URL: https://stspg.io/tblzdh8qw6qb
- Start time basis: `vendor_reported`.

<a id="ci-1110ad397ec159f0d30d946f"></a>

- Canonical ID: `ci-1110ad397ec159f0d30d946f` (1 source records).
  - [circleci/22mrv46g6n7g](../../data/2026/W40/circleci.json); vendor URL: https://stspg.io/4jm2frmlcq2q
- Start time basis: `vendor_reported`.

<a id="ci-11252a6149f2d714ef40fdb4"></a>

- Canonical ID: `ci-11252a6149f2d714ef40fdb4` (1 source records).
  - [openai/01M3WCKTAN41RZ01SBRRTY4AYM](../../data/2026/W40/openai.json); vendor URL: https://status.openai.com/incidents/01M3WCKTAN41RZ01SBRRTY4AYM
- Start time basis: `vendor_reported`.

<a id="ci-55eca2f1a8c0e00061c37085"></a>

- Canonical ID: `ci-55eca2f1a8c0e00061c37085` (1 source records).
  - [github/169pg72gg9jr](../../data/2026/W40/github.json); vendor URL: https://stspg.io/phqvtmdm3y2m
- Start time basis: `vendor_reported`.

<a id="ci-bfcb4d79a790459424bca9ed"></a>

- Canonical ID: `ci-bfcb4d79a790459424bca9ed` (1 source records).
  - [openai/01M3YNNVTSCXS8YM2V1YMHMVSE](../../data/2026/W40/openai.json); vendor URL: https://status.openai.com/incidents/01M3YNNVTSCXS8YM2V1YMHMVSE
- Start time basis: `vendor_reported`.

<a id="ci-b987b00a0a06a03498a3ae21"></a>

- Canonical ID: `ci-b987b00a0a06a03498a3ae21` (1 source records).
  - [cloudflare/vcstgt8gqkyk](../../data/2026/W40/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/vcstgt8gqkyk
- Start time basis: `vendor_reported`.

<a id="ci-2e8442e7a8ebbdfb6dbbc235"></a>

- Canonical ID: `ci-2e8442e7a8ebbdfb6dbbc235` (1 source records).
  - [grafana/9tzsfy13vrfj](../../data/2026/W40/grafana.json); vendor URL: https://stspg.io/zf8d8bx0228f
- Start time basis: `vendor_reported`.

<a id="ci-6dfefcb7172aab9378f94fe1"></a>

- Canonical ID: `ci-6dfefcb7172aab9378f94fe1` (1 source records).
  - [cloudflare/hyv5p7wn1c8v](../../data/2026/W40/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/hyv5p7wn1c8v
- Start time basis: `vendor_reported`.

<a id="ci-f39b9f6fa3da1e2121b651e0"></a>

- Canonical ID: `ci-f39b9f6fa3da1e2121b651e0` (1 source records).
  - [cloudflare/2xnmsnv8yv5x](../../data/2026/W40/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/2xnmsnv8yv5x
- Start time basis: `vendor_reported`.

<a id="ci-fd44ebff80a0390ee9f3abc8"></a>

- Canonical ID: `ci-fd44ebff80a0390ee9f3abc8` (1 source records).
  - [cloudflare/smn1d93zd3hr](../../data/2026/W40/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/smn1d93zd3hr
- Start time basis: `vendor_reported`.

<a id="ci-009af033d155ed325a33ea38"></a>

- Canonical ID: `ci-009af033d155ed325a33ea38` (1 source records).
  - [github/c8466lzvnmsv](../../data/2026/W40/github.json); vendor URL: https://stspg.io/cz0ngh11q0tf
- Start time basis: `vendor_reported`.

<a id="ci-ff88d294537cb1136fe2ddc3"></a>

- Canonical ID: `ci-ff88d294537cb1136fe2ddc3` (1 source records).
  - [stripe/f45340rc0qm0](../../data/2026/W40/stripe.json); vendor URL: https://stspg.io/0cknvjm6mzq6
- Start time basis: `vendor_reported`.

<a id="ci-e9bab7335aebe3ded937cfc1"></a>

- Canonical ID: `ci-e9bab7335aebe3ded937cfc1` (1 source records).
  - [grafana/zz75y4xkqrk1](../../data/2026/W40/grafana.json); vendor URL: https://stspg.io/scrxwgzkx24t
- Start time basis: `vendor_reported`.

<a id="ci-753d1d3fcfb9eeff8c30ba8d"></a>

- Canonical ID: `ci-753d1d3fcfb9eeff8c30ba8d` (1 source records).
  - [cloudflare/32sg73xwyq89](../../data/2026/W40/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/32sg73xwyq89
- Start time basis: `vendor_reported`.

<a id="ci-2b361c5e360a5e18caf4269b"></a>

- Canonical ID: `ci-2b361c5e360a5e18caf4269b` (1 source records).
  - [cloudflare/0dq5fbp2xwbj](../../data/2026/W40/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/0dq5fbp2xwbj
- Start time basis: `vendor_reported`.

<a id="ci-3da83e55f3d3237dde6de18f"></a>

- Canonical ID: `ci-3da83e55f3d3237dde6de18f` (1 source records).
  - [github/dqn46wtvbdzv](../../data/2026/W40/github.json); vendor URL: https://stspg.io/f0yg8jk2gm80
- Start time basis: `vendor_reported`.
- Warning: Vendor resolved_at precedes started_at; duration is unknown, score uses zero duration; decision:legacy-time-github-dqn46wtvbdzv

<a id="ci-2323353f27f9e46330a51ff2"></a>

- Canonical ID: `ci-2323353f27f9e46330a51ff2` (1 source records).
  - [cloudflare/n2645m15vwc9](../../data/2026/W40/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/n2645m15vwc9
- Start time basis: `vendor_reported`.

<a id="ci-c8d43343b9d2e7e41447f074"></a>

- Canonical ID: `ci-c8d43343b9d2e7e41447f074` (1 source records).
  - [cloudflare/4c842mwy6z4l](../../data/2026/W40/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/4c842mwy6z4l
- Start time basis: `vendor_reported`.

<a id="ci-653eb6af8f15b06196359ad4"></a>

- Canonical ID: `ci-653eb6af8f15b06196359ad4` (1 source records).
  - [cloudflare/0kj5cjjq93xn](../../data/2026/W40/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/0kj5cjjq93xn
- Start time basis: `vendor_reported`.
- Warning: Vendor resolved_at precedes started_at; duration is unknown, score uses zero duration; decision:legacy-time-cloudflare-0kj5cjjq93xn

<a id="ci-8a091363f474ea7b8a674caf"></a>

- Canonical ID: `ci-8a091363f474ea7b8a674caf` (1 source records).
  - [digitalocean/6z6dt3t1bhb9](../../data/2026/W40/digitalocean.json); vendor URL: https://stspg.io/kj9rd8tnd1wt
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
