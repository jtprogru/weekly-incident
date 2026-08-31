# Week 2026-W35 — top 5 of 14 above threshold

2026-08-24 to 2026-08-30 (UTC).

## 1. cloudflare — Incorrect geo location for some Cloudflare WARP users

- URL: https://www.cloudflarestatus.com/incidents/9g65dxfbcjln
- Window: 2026-08-27 18:46 to not yet (3d 19h 47m (ongoing))
- Impact: minor, status: identified
- Components: WARP

  - `2026-08-27 18:46` investigating — Cloudflare is investigating issues with Cloudflare WARP and Cloudflare Zero Trust users being incorrectly geolocated by some third party services.
  - `2026-08-27 18:56` identified — The issue has been identified and we are working with third party providers to fix it.
  - `2026-08-28 00:55` identified — We are continuing to work on a fix for this issue.
  - `2026-08-28 11:30` identified — We are continuing to work on a fix for this issue.
  - `2026-08-29 15:31` identified — We are continuing to work on a fix for this issue.
  - `2026-08-30 15:51` identified — We are continuing to work on a fix for this issue.

## 2. digitalocean — Cloud Control Panel and API

- URL: https://stspg.io/8l5c24cxh3hr
- Window: 2026-08-24 17:14 to 2026-08-25 19:04 UTC (1d 1h 49m)
- Impact: critical, status: resolved
- Components: API, Cloud Control Panel

  - `2026-08-24 17:14` investigating — Our Engineering team is investigating an issue impacting the DigitalOcean public API and Cloud Control Panel. At this time, users will see elevated 5xx errors for all API requests, including things like creates for various products. Users may also be unable to access cloud.digitalocean.com. We are actively investigating the root cause and will provide updates as soon as more information becomes available. We apologize for the inconvenience and appreciate your patience.
  - `2026-08-24 18:33` investigating — We are continuing to investigate the issue impacting the DigitalOcean public API and Cloud Control Panel. At this time, we are seeing some improvement in latency/error rates and some users may now be able to access the Cloud Control Panel. Serverless Inference users may currently experience 403 and 503 errors when making API requests or accessing affected services. We will continue to provide updates as more information becomes available. We apologize for the inconvenience and appreciate your patience.
  - `2026-08-24 19:14` monitoring — Our Engineering team has implemented mitigation measures to address the issue affecting the DigitalOcean public API and Cloud Control Panel. We are seeing requests succeed, and users should now be able to access the Cloud Control Panel and issue/receive API requests normally. We are currently monitoring the situation to ensure that the service has returned to normal operation and remains stable. We appreciate your patience and will provide an update once the issue is fully confirmed as resolved.
  - `2026-08-24 22:27` investigating — Our Engineering team is investigating a recurrence of the issue impacting the DigitalOcean public API and Cloud Control Panel. This issue does not impact our data plane (i.e. running droplets). Users may experience difficulties accessing the Cloud Control Panel, along with elevated latency and intermittent errors when accessing the public API or Cloud Control Plane. We are actively investigating the issue and working to restore normal service. We will continue to provide updates as more information becomes available. We apologize for the inconvenience and appreciate your patience.
  - `2026-08-25 00:54` monitoring — Our Engineering team has implemented mitigation measures to address the issue impacting the DigitalOcean public API and Cloud Control Panel. We are seeing improvements, and users should now be able to access the Cloud Control Panel and use the public API normally. We are continuing to monitor the service to ensure it remains stable and will provide further updates as more information becomes available. We appreciate your patience and understanding.
  - `2026-08-25 16:30` monitoring — Our Engineering team has continued to monitor the health of all services and has observed no recurrence of impact over the past 12+ hours. We will continue to monitor services for a short period to confirm stability before we mark this incident resolved. We understand the impact this incident had for DigitalOcean users and we will publish a postmortem on the status page to provide further details on root cause and mitigation.
  - `2026-08-25 19:04` resolved — Our Engineering team has confirmed all services remain stable and we will now resolve this incident. Further details on root cause and mitigation will be provided in a postmortem via this status page for users interested in learning more. Thank you for your patience throughout this incident. If you continue to experience any issues, please reach out to Support from within your account

## 3. cloudflare — DO and D1 are degraded in Hong Kong and Singapore

- URL: https://www.cloudflarestatus.com/incidents/ksy4vxmbkfdk
- Window: 2026-08-24 03:53 to 2026-08-24 08:28 UTC (4h 34m)
- Impact: minor, status: resolved
- Components: Hong Kong - (HKG), D1, Durable Objects, Singapore, Singapore - (SIN), R2

  - `2026-08-24 03:53` investigating — Some DOs and D1 databases may be experiencing issues in Hong Kong and Singapore. We are working to analyse and mitigate this problem. More updates to follow shortly.
  - `2026-08-24 06:42` monitoring — A fix has been implemented and we are monitoring the results.
  - `2026-08-24 08:28` resolved — This incident has been resolved.

## 4. cloudflare — Increased Latency

- URL: https://www.cloudflarestatus.com/incidents/d68x0kzjwhqf
- Window: 2026-08-24 20:38 to 2026-08-25 14:52 UTC (18h 14m)
- Impact: minor, status: resolved
- Components: CDN/Cache

  - `2026-08-24 20:38` investigating — Cloudflare is investigating an issue which may result in latency increases for a subset of requests.
  - `2026-08-24 22:24` identified — We have identified the cause of this issue and are implementing a fix.
  - `2026-08-25 08:50` identified — We are continuing to work on a fix for this issue.
  - `2026-08-25 14:37` monitoring — A fix has been implemented and we are monitoring the results.
  - `2026-08-25 14:52` resolved — This incident has been resolved.

## 5. github — Disruption with GitHub Billing

- URL: https://stspg.io/cpgst4hly3mh
- Window: 2026-08-26 23:37 to 2026-08-27 19:44 UTC (20h 6m)
- Impact: minor, status: resolved
- Components: —

  - `2026-08-26 23:37` investigating — We are investigating reports of impacted performance for some GitHub services.
  - `2026-08-26 23:42` investigating — We are currently investigating increased errors with billing services. Customers may observe failed billing budget page loads, and users of the Copilot CLI may observe failures starting or continuing sessions.
  - `2026-08-27 00:31` investigating — We've applied a mitigation to unblock Copilot usage and have observed recovery for this particular impact. We're continuing to investigate and apply mitigations for the billing page disruption while monitoring to ensure Copilot remains recovered.
  - `2026-08-27 01:35` investigating — We are continuing to monitor the mitigation that we have applied for the billing page disruption.
  - `2026-08-27 14:49` investigating — Our mitigation is still holding as we continue to investigate to find the root cause.
  - `2026-08-27 16:20` investigating — Our mitigation continues to hold, and service conditions remain stable. We are continuing to investigate the concentrated workload responsible for the issue and are preparing additional preventative improvements. We have not identified a material change in customer impact since the previous update. We will provide another update as the investigation progresses.
  - `2026-08-27 17:58` investigating — No material change since the previous update. Service conditions remain stable following the mitigation, and we have not observed any further customer impact. We are actively monitoring the service while implementing targeted fixes to address the underlying root cause.
  - `2026-08-27 19:44` resolved — This incident has been resolved. Thank you for your patience and understanding as we addressed this issue. A detailed root cause analysis will be shared as soon as it is available.

## Also above threshold (9)

- `anthropic` Elevated errors for multiple models (3h 24m) — https://stspg.io/5kzdprb1xsbx
- `cloudflare` Increased HTTP 5xx Errors in Singapore (7h 6m) — https://www.cloudflarestatus.com/incidents/7rlsppn13h62
- `anthropic` Elevated errors on Claude Code and Claude Cowork (2h 59m) — https://stspg.io/6rfby2z7stc2
- `cloudflare` Turnstile Challenge Issues (4h 48m) — https://www.cloudflarestatus.com/incidents/bthgp8mprw38
- `github` Incident with Actions (2h 49m) — https://stspg.io/pg14nv9m3095
- `openai` Users may experience an increase in error rates in Workspace Agents and ChatGPT Work on Web and Mobile (4h 14m) — https://status.openai.com/incidents/01M1263309PD7VYV8EF95W2M2P
- `github` Incident with Copilot AI Model Providers (2h 8m) — https://stspg.io/573gc8p4328b
- `grafana` Partial Logs Write Outage (1h 29m) — https://stspg.io/zc28d5qtcwdf
- `github` Elevated errors on Fable 5 due to upstream provider (46m) — https://stspg.io/4dz9rv14dz14

