# Week 2026-W35

2026-08-24 to 2026-08-30 (UTC). Generated 2026-08-31 14:34.

36 incidents collected, 14 above threshold. All 13 sources responded.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | cloudflare | [Incorrect geo location for some Cloudflare WARP users](https://www.cloudflarestatus.com/incidents/9g65dxfbcjln) | 2026-08-27 18:46 | 3d 19h 47m (ongoing) | minor | WARP | 9913 = 5507 × 1 × 1.8 |
| 2 | digitalocean | [Cloud Control Panel and API](https://stspg.io/8l5c24cxh3hr) | 2026-08-24 17:14 | 1d 1h 49m | critical | API, Cloud Control Panel | 2478 = 1549 × 2 × 0.8 |
| 3 | cloudflare | [DO and D1 are degraded in Hong Kong and Singapore](https://www.cloudflarestatus.com/incidents/ksy4vxmbkfdk) | 2026-08-24 03:53 | 4h 34m | minor | Hong Kong - (HKG), D1, Durable Objects, Singapore, Singapore - (SIN), R2 | 2466 = 274 × 5 × 1.8 |
| 4 | cloudflare | [Increased Latency](https://www.cloudflarestatus.com/incidents/d68x0kzjwhqf) | 2026-08-24 20:38 | 18h 14m | minor | CDN/Cache | 1969 = 1094 × 1 × 1.8 |
| 5 | github | [Disruption with GitHub Billing](https://stspg.io/cpgst4hly3mh) | 2026-08-26 23:37 | 20h 6m | minor | — | 1809 = 1206 × 1 × 1.5 |
| 6 | anthropic | [Elevated errors for multiple models](https://stspg.io/5kzdprb1xsbx) | 2026-08-24 05:06 | 3h 24m | major | claude.ai, Claude API (api.anthropic.com), Claude Code, Claude Cowork | 1224 = 204 × 4 × 1.5 |
| 7 | cloudflare | [Increased HTTP 5xx Errors in Singapore](https://www.cloudflarestatus.com/incidents/7rlsppn13h62) | 2026-08-28 17:36 | 7h 6m | minor | Singapore, Singapore - (SIN) | 767 = 426 × 1 × 1.8 |
| 8 | anthropic | [Elevated errors on Claude Code and Claude Cowork](https://stspg.io/6rfby2z7stc2) | 2026-08-28 17:22 | 2h 59m | major | Claude Code, Claude Cowork | 537 = 179 × 2 × 1.5 |
| 9 | cloudflare | [Turnstile Challenge Issues](https://www.cloudflarestatus.com/incidents/bthgp8mprw38) | 2026-08-27 14:40 | 4h 48m | minor | Turnstile | 518 = 288 × 1 × 1.8 |
| 10 | github | [Incident with Actions](https://stspg.io/pg14nv9m3095) | 2026-08-26 15:11 | 2h 49m | critical | Actions, Pages | 507 = 169 × 2 × 1.5 |
| 11 | openai | [Users may experience an increase in error rates in Workspace Agents and ChatGPT Work on Web and Mobile](https://status.openai.com/incidents/01M1263309PD7VYV8EF95W2M2P) | 2026-08-27 18:00 | 4h 14m | minor | — | 381 = 254 × 1 × 1.5 |
| 12 | github | [Incident with Copilot AI Model Providers](https://stspg.io/573gc8p4328b) | 2026-08-27 10:04 | 2h 8m | critical | Copilot AI Model Providers | 192 = 128 × 1 × 1.5 |
| 13 | grafana | [Partial Logs Write Outage](https://stspg.io/zc28d5qtcwdf) | 2026-08-28 17:12 | 1h 29m | major | AWS US East - prod-us-east-3 | 89 = 89 × 1 × 1.0 |
| 14 | github | [Elevated errors on Fable 5 due to upstream provider](https://stspg.io/4dz9rv14dz14) | 2026-08-24 07:12 | 46m | major | Copilot AI Model Providers | 69 = 46 × 1 × 1.5 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. cloudflare — Incorrect geo location for some Cloudflare WARP users

- URL: https://www.cloudflarestatus.com/incidents/9g65dxfbcjln
- Started: 2026-08-27 18:46 UTC
- Resolved: not yet
- Duration: 3d 19h 47m (ongoing)
- Impact: minor, status: identified
- Components: WARP

- `2026-08-27 18:46` **investigating** — Cloudflare is investigating issues with Cloudflare WARP and Cloudflare Zero Trust users being incorrectly geolocated by some third party services.
- `2026-08-27 18:56` **identified** — The issue has been identified and we are working with third party providers to fix it.
- `2026-08-28 00:55` **identified** — We are continuing to work on a fix for this issue.
- `2026-08-28 11:30` **identified** — We are continuing to work on a fix for this issue.
- `2026-08-29 15:31` **identified** — We are continuing to work on a fix for this issue.
- `2026-08-30 15:51` **identified** — We are continuing to work on a fix for this issue.

### 2. digitalocean — Cloud Control Panel and API

- URL: https://stspg.io/8l5c24cxh3hr
- Started: 2026-08-24 17:14 UTC
- Resolved: 2026-08-25 19:04 UTC
- Duration: 1d 1h 49m
- Impact: critical, status: resolved
- Components: API, Cloud Control Panel

- `2026-08-24 17:14` **investigating** — Our Engineering team is investigating an issue impacting the DigitalOcean public API and Cloud Control Panel. At this time, users will see elevated 5xx errors for all API requests, including things like creates for various products. Users may also be unable to access cloud.digitalocean.com. We are actively investigating the root cause and will provide updates as soon as more information becomes available. We apologize for the inconvenience and appreciate your patience.
- `2026-08-24 18:33` **investigating** — We are continuing to investigate the issue impacting the DigitalOcean public API and Cloud Control Panel. At this time, we are seeing some improvement in latency/error rates and some users may now be able to access the Cloud Control Panel. Serverless Inference users may currently experience 403 and 503 errors when making API requests or accessing affected services. We will continue to provide updates as more information becomes available. We apologize for the inconvenience and appreciate your patience.
- `2026-08-24 19:14` **monitoring** — Our Engineering team has implemented mitigation measures to address the issue affecting the DigitalOcean public API and Cloud Control Panel. We are seeing requests succeed, and users should now be able to access the Cloud Control Panel and issue/receive API requests normally. We are currently monitoring the situation to ensure that the service has returned to normal operation and remains stable. We appreciate your patience and will provide an update once the issue is fully confirmed as resolved.
- `2026-08-24 22:27` **investigating** — Our Engineering team is investigating a recurrence of the issue impacting the DigitalOcean public API and Cloud Control Panel. This issue does not impact our data plane (i.e. running droplets). Users may experience difficulties accessing the Cloud Control Panel, along with elevated latency and intermittent errors when accessing the public API or Cloud Control Plane. We are actively investigating the issue and working to restore normal service. We will continue to provide updates as more information becomes available. We apologize for the inconvenience and appreciate your patience.
- `2026-08-25 00:54` **monitoring** — Our Engineering team has implemented mitigation measures to address the issue impacting the DigitalOcean public API and Cloud Control Panel. We are seeing improvements, and users should now be able to access the Cloud Control Panel and use the public API normally. We are continuing to monitor the service to ensure it remains stable and will provide further updates as more information becomes available. We appreciate your patience and understanding.
- `2026-08-25 16:30` **monitoring** — Our Engineering team has continued to monitor the health of all services and has observed no recurrence of impact over the past 12+ hours. We will continue to monitor services for a short period to confirm stability before we mark this incident resolved. We understand the impact this incident had for DigitalOcean users and we will publish a postmortem on the status page to provide further details on root cause and mitigation.
- `2026-08-25 19:04` **resolved** — Our Engineering team has confirmed all services remain stable and we will now resolve this incident. Further details on root cause and mitigation will be provided in a postmortem via this status page for users interested in learning more. Thank you for your patience throughout this incident. If you continue to experience any issues, please reach out to Support from within your account

### 3. cloudflare — DO and D1 are degraded in Hong Kong and Singapore

- URL: https://www.cloudflarestatus.com/incidents/ksy4vxmbkfdk
- Started: 2026-08-24 03:53 UTC
- Resolved: 2026-08-24 08:28 UTC
- Duration: 4h 34m
- Impact: minor, status: resolved
- Components: Hong Kong - (HKG), D1, Durable Objects, Singapore, Singapore - (SIN), R2

- `2026-08-24 03:53` **investigating** — Some DOs and D1 databases may be experiencing issues in Hong Kong and Singapore. We are working to analyse and mitigate this problem. More updates to follow shortly.
- `2026-08-24 06:42` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-08-24 08:28` **resolved** — This incident has been resolved.

### 4. cloudflare — Increased Latency

- URL: https://www.cloudflarestatus.com/incidents/d68x0kzjwhqf
- Started: 2026-08-24 20:38 UTC
- Resolved: 2026-08-25 14:52 UTC
- Duration: 18h 14m
- Impact: minor, status: resolved
- Components: CDN/Cache

- `2026-08-24 20:38` **investigating** — Cloudflare is investigating an issue which may result in latency increases for a subset of requests.
- `2026-08-24 22:24` **identified** — We have identified the cause of this issue and are implementing a fix.
- `2026-08-25 08:50` **identified** — We are continuing to work on a fix for this issue.
- `2026-08-25 14:37` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-08-25 14:52` **resolved** — This incident has been resolved.

### 5. github — Disruption with GitHub Billing

- URL: https://stspg.io/cpgst4hly3mh
- Started: 2026-08-26 23:37 UTC
- Resolved: 2026-08-27 19:44 UTC
- Duration: 20h 6m
- Impact: minor, status: resolved
- Components: —

- `2026-08-26 23:37` **investigating** — We are investigating reports of impacted performance for some GitHub services.
- `2026-08-26 23:42` **investigating** — We are currently investigating increased errors with billing services. Customers may observe failed billing budget page loads, and users of the Copilot CLI may observe failures starting or continuing sessions.
- `2026-08-27 00:31` **investigating** — We've applied a mitigation to unblock Copilot usage and have observed recovery for this particular impact. We're continuing to investigate and apply mitigations for the billing page disruption while monitoring to ensure Copilot remains recovered.
- `2026-08-27 01:35` **investigating** — We are continuing to monitor the mitigation that we have applied for the billing page disruption.
- `2026-08-27 14:49` **investigating** — Our mitigation is still holding as we continue to investigate to find the root cause.
- `2026-08-27 16:20` **investigating** — Our mitigation continues to hold, and service conditions remain stable. We are continuing to investigate the concentrated workload responsible for the issue and are preparing additional preventative improvements. We have not identified a material change in customer impact since the previous update. We will provide another update as the investigation progresses.
- `2026-08-27 17:58` **investigating** — No material change since the previous update. Service conditions remain stable following the mitigation, and we have not observed any further customer impact. We are actively monitoring the service while implementing targeted fixes to address the underlying root cause.
- `2026-08-27 19:44` **resolved** — This incident has been resolved. Thank you for your patience and understanding as we addressed this issue. A detailed root cause analysis will be shared as soon as it is available.

### 6. anthropic — Elevated errors for multiple models

- URL: https://stspg.io/5kzdprb1xsbx
- Started: 2026-08-24 05:06 UTC
- Resolved: 2026-08-24 08:30 UTC
- Duration: 3h 24m
- Impact: major, status: resolved
- Components: claude.ai, Claude API (api.anthropic.com), Claude Code, Claude Cowork

- `2026-08-24 05:06` **investigating** — We are investigating elevated errors on requests to Claude Mythos 5, Claude Fable 5, Claude Opus 5, and Claude Opus 4.8. We will provide an update as soon as possible.
- `2026-08-24 05:27` **identified** — We have identified the cause of elevated errors on requests to Claude Mythos 5, Claude Fable 5, Claude Opus 5, and other Claude models and are working on a fix. We will provide an update as soon as possible.
- `2026-08-24 06:42` **identified** — We are continuing to work to resolve issues causing elevated requests on multiple models. We will provide an update as soon as possible.
- `2026-08-24 07:47` **identified** — At this time, we have seen errors stabilize on Opus 5 and Fable 5, and are working to fully resolve success rates on all affected models. We will provide an update as soon as possible.
- `2026-08-24 08:30` **resolved** — This issue has been resolved. From 9:50pm PT / 04:50 UTC through 00:36am PT / 07:36 UTC, users saw elevated errors on requests to Claude models, including Claude Opus 5 and Fable 5.

### 7. cloudflare — Increased HTTP 5xx Errors in Singapore

- URL: https://www.cloudflarestatus.com/incidents/7rlsppn13h62
- Started: 2026-08-28 17:36 UTC
- Resolved: 2026-08-29 00:43 UTC
- Duration: 7h 6m
- Impact: minor, status: resolved
- Components: Singapore, Singapore - (SIN)

- `2026-08-28 17:36` **investigating** — Cloudflare is investigating an increased level of HTTP 5xx errors in Singapore. We are working to analyse and mitigate this problem. More updates to follow shortly.
- `2026-08-28 17:39` **identified** — The issue has been identified and a fix is being implemented.
- `2026-08-28 17:40` **identified** — We are continuing to work on a fix for this issue.
- `2026-08-28 17:58` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-08-29 00:43` **resolved** — This incident has been resolved.

### 8. anthropic — Elevated errors on Claude Code and Claude Cowork

- URL: https://stspg.io/6rfby2z7stc2
- Started: 2026-08-28 17:22 UTC
- Resolved: 2026-08-28 20:21 UTC
- Duration: 2h 59m
- Impact: major, status: resolved
- Components: Claude Code, Claude Cowork

- `2026-08-28 17:22` **identified** — We have identified an issue with an upstream cloud provider affecting Claude Cowork and Claude Code on the web. Some sessions may fail to start or disconnect mid-task; affected sessions can be retried. We are in contact with the provider and will provide an update as soon as possible.
- `2026-08-28 18:21` **monitoring** — A mitigation has been applied for the issue affecting Claude Cowork and Claude Code on the web, and we are monitoring for recovery. Sessions that disconnected earlier can be retried
- `2026-08-28 20:21` **resolved** — The issue affecting Claude Cowork and Claude Code on the web has been resolved.

### 9. cloudflare — Turnstile Challenge Issues

- URL: https://www.cloudflarestatus.com/incidents/bthgp8mprw38
- Started: 2026-08-27 14:40 UTC
- Resolved: 2026-08-27 19:29 UTC
- Duration: 4h 48m
- Impact: minor, status: resolved
- Components: Turnstile

- `2026-08-27 14:40` **investigating** — Cloudflare is investigating a potential issue with the Turnstile challenge platform. Users may experience failed challenges solve attempts. Further details will be provided as more information becomes available.
- `2026-08-27 19:29` **resolved** — This incident has been resolved.

### 10. github — Incident with Actions

- URL: https://stspg.io/pg14nv9m3095
- Started: 2026-08-26 15:11 UTC
- Resolved: 2026-08-26 18:01 UTC
- Duration: 2h 49m
- Impact: critical, status: resolved
- Components: Actions, Pages

- `2026-08-26 15:11` **investigating** — We are investigating reports of degraded availability for Actions
- `2026-08-26 15:12` **investigating** — Pages is experiencing degraded performance. We are continuing to investigate.
- `2026-08-26 15:23` **investigating** — We've identified an issue with a database primary and are failing over to a replica immediately
- `2026-08-26 15:48` **investigating** — primary failover briefly improved performance but did not fully mitigate, we've throttled inbound traffic and are investigating upstream Vitess issues
- `2026-08-26 16:14` **investigating** — We believe we've identified and addressed the issue and are ramping traffic back up slowly to ensure it doesn't recur. Some customers will continue to see delays as we ramp up.
- `2026-08-26 16:49` **investigating** — Pages is operating normally.
- `2026-08-26 16:50` **investigating** — We are continuing to observe recovery and delayed queues are burning down. Some customers will continue to see increased delays until all throttled work has been completed - we expect this within the next hour.
- `2026-08-26 17:32` **investigating** — We are continuing to observe recovery and expect actions inbound queues to be back to normal in <30min. Work will continue to flow through the system subject to per-customer concurrency limits.
- `2026-08-26 17:54` **monitoring** — The degradation affecting Actions has been mitigated. We are monitoring to ensure stability.
- `2026-08-26 18:00` **monitoring** — All inbound queues have recovered and Actions is operating as expected. 3.7% of jobs assigned to larger runners during the early stage of this incident are stuck waiting for runner assignment. Those will be canceled within the hour. Other runners are successfully processing all new jobs.
- `2026-08-26 18:01` **resolved** — On August 26, 2026 from 15:02 to 15:45 UTC, Actions jobs failed to start. The following 2 hours until 17:40 UTC, Actions runs were delayed starting by more than 5 minutes as the system caught up with delayed load. This impact was triggered by saturation of writes to the database primary used by the service processing triggers for Actions workflows. The primary was failed over, but the system did not fully recover. The saturation was caused by growing daily peak load combined with an upstream issue in GitHub’s event processing infrastructure, https://www.githubstatus.com/incidents/hcbtzksccj2f, which caused burst amplification of already-high load. Downstream throttles that were later used to recover were set ~10% too high to protect the system. <br /><br />At 15:45 UTC, throttling combined with service restarts recovered the service’s core health. Those throttles were gradually raised between 15:54 and 17:22 to restore full webhook processing for Actions runs. This ramp was deliberately slow to ensure we did not re-overwhelm the system given our original throttling was now known to be incorrectly set. The queue of webhook events was fully burned down at 17:40 UTC. <br /><br />3.7% of larger-runner jobs, along with some scale-set self-hosted jobs, remained stuck in queued or “waiting for runner” state. We deployed a change to force-revoke jobs in this state, and they transitioned to failed at 18:40 UTC, about 50 minutes after incident mitigation. Releasing these jobs also freed hosted concurrency for larger-runner jobs. <br /><br />Customers using concurrency groups saw longer impact due to a separate issue where runners assigned to a subset of jobs disconnected before the force-revoke mitigation was deployed, which prevented runner acquisition from progressing and left jobs in a waiting-for-runner state. This was resolved at 01:00 UTC on August 27. <br /><br />Some runs triggered during the 15:02-15:45 UTC incident window encountered a bug that left them showing as queued even after service recovery. In the backend, these runs had already failed and will automatically move to canceled state 24 hours after creation. As follow-up, we are fixing the root cause of this queued state and improving our ability to bulk-cancel affected runs. <br /><br />Several changes to improve the general scalability of this part of Actions were already complete and deploying to production. Rollout of those changes will be complete within the next 24 hours. Further work to improve scale, resiliency, and more graceful degradation of Actions workflows are in flight. We are also taking a repair item to accelerate clearing of stuck queued or waiting jobs in similar future cases.

### 11. openai — Users may experience an increase in error rates in Workspace Agents and ChatGPT Work on Web and Mobile

- URL: https://status.openai.com/incidents/01M1263309PD7VYV8EF95W2M2P
- Started: 2026-08-27 18:00 UTC
- Resolved: 2026-08-27 22:14 UTC
- Duration: 4h 14m
- Impact: minor, status: resolved
- Components: —

- `2026-08-27 18:00` **identified** — We have identified that users are experiencing elevated errors for the impacted services. We are working on implementing a mitigation.
- `2026-08-27 18:16` **identified** — We have identified that users are experiencing elevated errors for the impacted services. We are working on implementing a mitigation.
- `2026-08-27 19:07` **monitoring** — We have applied the mitigation and are monitoring the recovery.
- `2026-08-27 22:14` **resolved** — All impacted services have now fully recovered.

### 12. github — Incident with Copilot AI Model Providers

- URL: https://stspg.io/573gc8p4328b
- Started: 2026-08-27 10:04 UTC
- Resolved: 2026-08-27 12:12 UTC
- Duration: 2h 8m
- Impact: critical, status: resolved
- Components: Copilot AI Model Providers

- `2026-08-27 10:04` **investigating** — We are investigating reports of degraded availability for Copilot AI Model Providers
- `2026-08-27 10:43` **investigating** — We are experiencing degraded availability for the Kimi K3 model in Copilot products and IDE surfaces. This is due to an issue with an upstream model provider. While we work with them to resolve the issue, we recommend choosing another model or selecting 'Auto' to continue using Copilot.
- `2026-08-27 11:58` **investigating** — Copilot AI Model Providers is experiencing degraded performance. We are continuing to investigate.
- `2026-08-27 12:12` **investigating** — The issues with our upstream model provider have been mitigated, and Kimi K3 is once again available in Copilot products and IDE surfaces.<br />We will continue monitoring to ensure stability.
- `2026-08-27 12:12` **resolved** — This incident has been resolved. Thank you for your patience and understanding as we addressed this issue. A detailed root cause analysis will be shared as soon as it is available.

### 13. grafana — Partial Logs Write Outage

- URL: https://stspg.io/zc28d5qtcwdf
- Started: 2026-08-28 17:12 UTC
- Resolved: 2026-08-28 18:42 UTC
- Duration: 1h 29m
- Impact: major, status: resolved
- Components: AWS US East - prod-us-east-3

- `2026-08-28 17:12` **identified** — We’ve identified the cause of an issue impacting Log writes. Our team is currently implementing a fix.
- `2026-08-28 17:55` **monitoring** — This incident also had impact on Frontend Observability for the same region. A fix has been deployed, and things have stablized. We will continue to monitor on our end.
- `2026-08-28 18:42` **resolved** — This incident has been resolved.

### 14. github — Elevated errors on Fable 5 due to upstream provider

- URL: https://stspg.io/4dz9rv14dz14
- Started: 2026-08-24 07:12 UTC
- Resolved: 2026-08-24 07:58 UTC
- Duration: 46m
- Impact: major, status: resolved
- Components: Copilot AI Model Providers

- `2026-08-24 07:12` **investigating** — We are investigating reports of degraded availability for Copilot AI Model Providers
- `2026-08-24 07:12` **investigating** — We are experiencing degraded availability for the Fable model in Copilot products and IDE surfaces. This is due to an issue with the upstream model provider. While we work with them to resolve the issue, we recommend choosing another model or selecting 'Auto' to continue using Copilot.
- `2026-08-24 07:58` **resolved** — On August 24th, 2026, between approximately 06:35 and 07:25 UTC, the Copilot service experienced a degradation of the Claude Fable 5 model due to an issue with our upstream provider. Users encountered elevated error rates when using Claude Fable 5, with requests sometimes failing mid-response. No other models were impacted.<br /><br />The issue was resolved by a mitigation put in place by our provider. GitHub is working with our provider to further improve the resiliency of the service to prevent similar incidents in the future.

## Below threshold (22)

- `cloudflare` [Durable Objects and Downstream Service Errors](https://www.cloudflarestatus.com/incidents/rzc0p0rk680b) — 1h 21m
- `github` [Incident with Actions and Pull Requests](https://stspg.io/nrbwjftcz72d) — 1h 30m
- `cloudflare` [Increased HTTP 5xx Errors in Chicago](https://www.cloudflarestatus.com/incidents/9v68b4k90mry) — 1h 6m
- `stripe` [Elevated Swish errors](https://stspg.io/jcy642t9yszg) — 1h 25m
- `stripe` [Elevated Swish errors](https://stspg.io/1yfx7lx0ybwq) — 1h 16m
- `cloudflare` [Workers Builds are Degraded](https://www.cloudflarestatus.com/incidents/yj2zk6n6rz1c) — 1h 38m
- `cloudflare` [Realtimekit APIs failing due to breakage in DB connectivity](https://www.cloudflarestatus.com/incidents/kv7g91gcddsn) — 1h 0m
- `grafana` [Some Grafana UI features may be unavailable or reverting to legacy behaviour](https://stspg.io/lj9v7hbqc7gp) — 1h 42m
- `github` [Disruption with some GitHub services](https://stspg.io/4s4bvnc0frrj) — 58m
- `cloudflare` [Cloudflare MCP Portal tool invocation logging instability](https://www.cloudflarestatus.com/incidents/gjw0tnhh4f3f) — 38m
- `anthropic` [Issues logging into Claude.ai](https://stspg.io/6211zbpptv0y) — 14m
- `stripe` [Elevated card declines for some users in Mexico](https://stspg.io/bf2rqyjxyqkv) — 48m
- `grafana` [Elevated error rates affecting metrics writes in prod-us-central-0](https://stspg.io/75v8n623y4jn) — 56m
- `github` [Actions delays in starting runs](https://stspg.io/v6ysclb9vcbd) — 37m
- `circleci` [Intermittent errors when viewing plan usage UIs or calling plan usage APIs](https://stspg.io/961zy5b9fnjk) — 26m
- `grafana` [Incident Management unavailable in US Central](https://stspg.io/kdhyw5xk6r56) — 23m
- `anthropic` [Errors logging into Claude.ai](https://stspg.io/kg59rclpfzsz) — 3m
- `github` [Degraded Git Operations over SSH](https://stspg.io/9nls6g9ln8p4) — 0m
- `cloudflare` [Network performance issues in Los Angeles](https://www.cloudflarestatus.com/incidents/jfhwhy4n6199) — 0m
- `circleci` [GitHub Login Disruption](https://stspg.io/ns4jz57ynx52) — 0m
- `grafana` [Mimir Writes Incident in prod-us-central-0](https://stspg.io/cz0zxwg004tf) — 0m
- `cloudflare` [Network Issues in Montréal](https://www.cloudflarestatus.com/incidents/z557s71pzcj1) — 0m

## Sources

| Vendor | Status | Incidents seen | Parse errors |
|--------|--------|----------------|--------------|
| anthropic | ok | 50 | 0 |
| atlassian | ok | 34 | 0 |
| aws | ok | 4 | 0 |
| circleci | ok | 50 | 0 |
| cloudflare | ok | 50 | 0 |
| datadog | ok | 50 | 0 |
| digitalocean | ok | 50 | 0 |
| gcp | ok | 5 | 0 |
| github | ok | 50 | 0 |
| grafana | ok | 50 | 0 |
| openai | ok | 25 | 0 |
| slack | ok | 39 | 0 |
| stripe | ok | 50 | 0 |

