# Week 2026-W41 — top 5 of 34 above threshold

2026-10-05 to 2026-10-11 (UTC).

## 1. anthropic — Elevated errors on platform.claude.com

- URL: https://stspg.io/sjkw7njwf0mt
- Window: 2026-10-07 13:25 to not yet (5d 2h 51m (ongoing))
- Impact: major, status: monitoring
- Components: Claude Console (platform.claude.com)

  - `2026-10-07 13:25` investigating — We are investigating elevated errors affecting usage data on the platform.claude.com Usage page and the Admin API usage report. Core functionality is not affected. We will provide an update as soon as possible.
  - `2026-10-07 16:36` investigating — We are continuing to work on elevated errors affecting usage data on the platform.claude.com Usage and Rate limits pages and the Admin API usage report. A mitigation for the Usage page and the Admin API usage report is rolling out; once it is in place, the most recent usage may take about 30 minutes to appear there. Core functionality is not affected.
  - `2026-10-07 17:28` identified — We've identified the cause of errors loading usage data. The Admin API usage report has recovered, and fixes for the platform.claude.com Usage and Rate limits pages are in progress. The most recent 30 minutes of usage may be delayed. Core functionality is not affected.
  - `2026-10-07 21:44` monitoring — We've mitigated the errors loading usage data on platform.claude.com and in the Admin API usage report. While we work on a full fix, the most recent ~30 minutes of usage may take longer to appear on the platform.claude.com. API requests are not affected, and we will provide additional updates here.

## 2. grafana — IRM Unable to Send SMS to Thailand, Bangladesh, and Austrian Numbers

- URL: https://stspg.io/9l8htmqt82s2
- Window: 2026-10-05 16:17 to not yet (6d 23h 58m (ongoing))
- Impact: minor, status: identified
- Components: Incident Management and Response (IRM)

  - `2026-10-05 16:17` investigating — We are currently investigating an issue impacting all SMS attempts to Austria for IRM. All SMS attempts to Austria for IRM are currently blocked. We are working on a resolution, and an update will follow as soon as we have more info. Users should configure mobile push or phone calls as notifications.
  - `2026-10-05 21:47` identified — The issue has been identified. We are working on a resolution, but do not have an ETA at this time. We will provide another update as soon as we have more information to share. Users should configure mobile push or phone calls as notifications.
  - `2026-10-06 13:42` identified — This issue is also impacting numbers from Thailand and Bangladesh. The title of this status page has been updated to reflect this. Users should configure mobile push or phone calls as notifications. We are working on a resolution, and we will provide another update as soon as we have more information to share.
  - `2026-10-09 12:51` identified — We are continuing working on a resolution, and will provide another update as soon as we have more information to share. In the meantime we recommend, affected users to configure mobile push or phone calls as notifications.

## 3. cloudflare — Elevated Errors in China Network Colos

- URL: https://www.cloudflarestatus.com/incidents/cyxclg7nnhjy
- Window: 2026-10-10 04:13 to not yet (2d 12h 2m (ongoing))
- Impact: minor, status: identified
- Components: China Network

  - `2026-10-10 04:13` investigating — We are investigating a spike in errors observed in our China colos.
  - `2026-10-10 05:48` identified — The issue has been identified and a fix is being implemented.
  - `2026-10-12 04:23` identified — We are investigating reports of degraded performance for traffic traversing the public internet between certain cities and provinces in mainland China and origin servers located outside the region. The impact is being reported across multiple ISPs and manifests as slower speeds or intermittent timeouts. While this congestion occurs upstream of Cloudflare's network, our engineering team is actively working to mitigate the effects. Intermittent 524 timeout errors is expected to continue due to ongoing congestion. While there is no fixed timeline for full resolution, as there are dependencies outside our network, we are taking steps to reduce customer impact.

## 4. openai — Delayed Costs data in the Compliance API

- URL: https://status.openai.com/incidents/01M4GV1M2A3K38T73JCVCD9G0R
- Window: 2026-10-09 17:23 to not yet (2d 22h 52m (ongoing))
- Impact: minor, status: monitoring
- Components: —

  - `2026-10-09 17:23` identified — We have identified an issue causing delays in publishing organization cost data through the Costs API. API and ChatGPT usage are not affected, and other reporting channels remain available. We are working to restore normal data freshness.
  - `2026-10-09 17:27` identified — We have identified an issue causing delays in publishing Costs data through the Compliance API. The API Platform organization costs endpoint and ChatGPT usage are not affected, and other reporting channels remain available. We are working to restore normal data freshness.
  - `2026-10-09 18:46` identified — Some customers may continue to experience delayed Costs data through the Compliance API. ChatGPT usage and API Platform requests are not affected, and other reporting channels remain available. We have identified the cause and are continuing work to restore normal data freshness. We will provide another update as work progresses.
  - `2026-10-09 22:03` monitoring — We have implemented a mitigation, and delayed Costs data is now processing through the Compliance API. Some customers may continue to see incomplete Costs data while the backlog catches up. We expect this catch-up process to continue through Sunday, October 11. ChatGPT usage and API Platform requests remain unaffected. We are monitoring recovery.

## 5. slack — Slack for iOS closing unexpectedly when opening channels

- URL: https://slack-status.com/2026-10/a0cfb0208a1d2032
- Window: 2026-10-09 15:35 to not yet (3d 0h 40m (ongoing))
- Impact: unknown, status: investigating
- Components: Connectivity

  - `2026-10-09 15:35` unknown — Users may experience iOS closing unexpectedly when opening channels. We're actively looking in to this, and we'll share another update as soon as we have more information. We appreciate your patience in the meantime.
  - `2026-10-09 16:20` unknown — We've developed a fix for the iOS app issue and are working to validate it before releasing it to customers. We'll share another update as soon as we have more information.
  - `2026-10-09 23:58` unknown — An updated version of the Slack iOS app that fixes the crash was approved and began rolling out gradually. Customers who download the updated version from the app store can now open channels without the app closing unexpectedly. We'll provide an update once additional information becomes available.

## Also above threshold (29)

- `grafana` Degraded Performance with Increased Latency for Grafana Cloud Metrics and Traces in GCP US Central (11h 33m) — https://stspg.io/cjbf6knm4jq4
- `cloudflare` Durable Objects Availability Issue in Eastern North America (5h 21m) — https://www.cloudflarestatus.com/incidents/z28czd65h29l
- `cloudflare` Network Performance Issues in Columbus (9h 2m) — https://www.cloudflarestatus.com/incidents/tt8f4ttcuonf
- `circleci` Delay in data appearing in UI and API (14h 35m) — https://stspg.io/hkfkwz292bz8
- `openai` Elevated Work Mode errors (19h 6m) — https://status.openai.com/incidents/01M45E0JFFJZ6EQ3BWYPW74R2Z
- `cloudflare` Cloudflare DNS Resolver 1.1.1.1 and Gateway DNS are having issues in SIN (4h 49m) — https://www.cloudflarestatus.com/incidents/4knp3pzs3qvx
- `gcp` Google Cloud Storage customers are experiencing elevated latencies and error rates when performing upload or write operations. (11h 11m) — https://status.cloud.google.com/incidents/mSFVt6PiP557bjRyYsut
- `cloudflare` API Shield JWT Validation Errors (11h 43m) — https://www.cloudflarestatus.com/incidents/ts295zqycqn7
- `cloudflare` Cloudflare API service issues (7h 19m) — https://www.cloudflarestatus.com/incidents/x8n12150xwf6
- `github` Several services are degraded (3h 43m) — https://stspg.io/w3rcssslq22w
- `github` Incident with Actions (3h 37m) — https://stspg.io/c11dc9nb1zdq
- `circleci` Delays in pipelines, workflows, UI data, and notifications (5h 18m) — https://stspg.io/s7tpg47srxdd
- `openai` Elevated errors in ChatGPT conversations (7h 2m) — https://status.openai.com/incidents/01M46G248GGHV9RSAJE82BZY1X
- `openai` Unable to create new ChatGPT Work threads in latest version of desktop app (6h 21m) — https://status.openai.com/incidents/01M4ADHMMBK7K92NJFJ0DXRYJM
- `openai` Elevated error rates in Codex Cloud (6h 2m) — https://status.openai.com/incidents/01M46KED6KQ6PR195KZYYVNY8W
- `github` Incident with Git Operations, Pull Requests and Actions (1h 10m) — https://stspg.io/96smrcth8bpg
- `openai` Admin Console Intermittently Inaccessible (5h 17m) — https://status.openai.com/incidents/01M46C47WYYD9NMZVGH83X3ZR1
- `stripe` Elevated Swish errors (6h 33m) — https://stspg.io/m7j171ppknbs
- `cloudflare` Network Performance Issues in Kochi (3h 0m) — https://www.cloudflarestatus.com/incidents/xczvsioz1un8
- `openai` Some ChatGPT content may be unavailable on Android (3h 26m) — https://status.openai.com/incidents/01M4H7FKXDA425TDNP2CRGEYA3
- `github` Incident with Git Operations, Issues, Actions and Pull Requests (47m) — https://stspg.io/4p8x0jnnsx7m
- `openai` Computer Use Is Experiencing Issues (2h 45m) — https://status.openai.com/incidents/01M493CZ0CB14KPC0EER83MNVB
- `circleci` Elevated wait times for machine jobs (3h 35m) — https://stspg.io/vcsrgfwdsy4z
- `openai` Errors processing large PDFs in the Responses API (2h 4m) — https://status.openai.com/incidents/01M48Y8E81QAFF9XKVGE118E6X
- `github` Disruption with some GitHub services (1h 44m) — https://stspg.io/pjcfc4c3pm7g
- `openai` Workspace admins may be unable to manage plugins (1h 38m) — https://status.openai.com/incidents/01M4HA7BCCXVHTR5DCYMSBQF2H
- `anthropic` Elevated errors loading usage data on platform.claude.com (1h 13m) — https://stspg.io/scdhhnplc2jn
- `anthropic` Elevated errors for Claude Opus 5.5 (18m) — https://stspg.io/bzc81gbv1b99
- `anthropic` Elevated errors for Claude Mythos 5.1 and Claude Fable 5.1 (15m) — https://stspg.io/wdqhxt8xt574

