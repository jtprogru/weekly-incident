# Week 2026-W46 — top 5 of 33 above threshold

2026-11-09 to 2026-11-15 (UTC).

54 source records (67 saved revisions) → 54 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## 1. cloudflare — Changes to Cloudflare Workers are not producing Event Subscription events

- URL: https://www.cloudflarestatus.com/incidents/0ocqieaifg7h
- Window: 2026-11-10 01:53 to not yet (4d 10h 42m (ongoing))
- Impact: minor, status: identified
- Components: Queues, Workers

<a id="ci-0a30f34c752fb49c0f4d211c"></a>

- Canonical ID: `ci-0a30f34c752fb49c0f4d211c` (1 source records).
  - [cloudflare/0ocqieaifg7h](../../data/2026/W46/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/0ocqieaifg7h
- Start time basis: `vendor_reported`.

  - `2026-11-10 01:53` investigating — Changes to Cloudflare Workers are temporarily not producing Event Subscription events. We have intentionally disabled this feature due to an on-going issue. We are actively working to restore it, but this, regrettably, may take a while. Changes to Workers themselves are working normally. Already deployed Workers are entirely unaffected.
  - `2026-11-11 05:05` identified — The issue has been identified and a fix is being implemented.

## 2. anthropic — Elevated errors affecting the Microsoft 365 connector

- URL: https://stspg.io/m0rbx38nmy1c
- Window: 2026-11-12 17:52 to 2026-11-13 21:49 UTC (1d 3h 56m)
- Impact: major, status: resolved
- Components: Claude Cowork, claude.ai

<a id="ci-3aedfad25e1c72a10aaae7ca"></a>

- Canonical ID: `ci-3aedfad25e1c72a10aaae7ca` (1 source records).
  - [anthropic/sqn3fg1l96dl](../../data/2026/W46/anthropic.json); vendor URL: https://stspg.io/m0rbx38nmy1c
- Start time basis: `vendor_reported`.

  - `2026-11-12 17:52` investigating — We are investigating elevated errors and slow responses affecting the Microsoft 365 connector. Requests to mail, SharePoint, Teams chat and calendar through the connector may fail, and new connections to the connector may not complete. We will provide an update as soon as possible.
  - `2026-11-12 18:34` identified — We have identified an issue with an upstream provider affecting the Microsoft 365 connector. Mail, calendar and search tools are failing for most requests, and new connections to the connector may not complete. We are in contact with them and will provide an update as soon as possible.
  - `2026-11-12 20:28` identified — We are continuing to work with our upstream provider on this issue. Mail, calendar and search tools in the Microsoft 365 connector are still failing for most requests. Claude for Outlook is also affected: mail and calendar actions may fail, and some users may not be able to open the add-in. We will provide another update within the hour.
  - `2026-11-12 21:18` identified — We are continuing to work with our upstream provider, who is restarting affected servers. Mail, calendar and search tools in the Microsoft 365 connector are still failing for most requests, and Claude for Outlook remains affected. We will provide another update as soon as we have more information.
  - `2026-11-12 22:23` identified — This issue is caused by a Microsoft 365 service incident, and mail, calendar and search tools in the Microsoft 365 connector are still failing for most requests. Claude for Outlook remains affected. You can follow Microsoft's progress at https://status.cloud.microsoft. We will update this post when service is restored.
  - `2026-11-13 15:40` monitoring — We're seeing signs of recovery for the Microsoft 365 connector. Claude for Outlook is also recovering. Some SharePoint search requests are still returning errors. Microsoft's incident remains open, so errors may return. We're continuing to monitor and will update this post as soon as we have more info.
  - `2026-11-13 21:49` resolved — This incident has been resolved. From 08:03 PT / 16:03 UTC on November 12 to 08:40 PT / 16:40 UTC on November 13, mail, calendar and search tools in the Microsoft 365 connector and Claude for Outlook returned errors for many requests because of a Microsoft 365 service incident. Microsoft is still completing its recovery, so you may see occasional errors or delays.

## 3. cloudflare — Replicate Model Registry Unavailable

- URL: https://www.cloudflarestatus.com/incidents/yu63fhh2peb7
- Window: 2026-11-11 04:33 to 2026-11-13 01:44 UTC (1d 21h 10m)
- Impact: minor, status: resolved
- Components: Replicate

<a id="ci-63dc50e88eca748e5be37fd9"></a>

- Canonical ID: `ci-63dc50e88eca748e5be37fd9` (1 source records).
  - [cloudflare/yu63fhh2peb7](../../data/2026/W46/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/yu63fhh2peb7
- Start time basis: `vendor_reported`.

  - `2026-11-11 04:33` investigating — Cloudflare is investigating issues with Replicate input uploads. We are working to understand the full impact and mitigate this problem. More updates to follow shortly
  - `2026-11-11 05:11` identified — The issue has been identified and a fix is being implemented.
  - `2026-11-11 07:33` identified — We are continuing to work on a fix for the issue. Current impact: Input uploads are degraded. Model builds and pushes via r8.im are unavailable.
  - `2026-11-11 23:50` identified — We are continuing to work on a fix for the issue. Current impact: Input uploads are degraded. Model builds and pushes via r8.im are unavailable.
  - `2026-11-12 22:55` monitoring — A fix has been implemented and we are monitoring the results.
  - `2026-11-13 01:44` resolved — This incident has been resolved.

## 4. openai — Elevated errors affecting the Microsoft 365 connector

- URL: https://status.openai.com/incidents/01M78GAT3AK9P9W8Z230KHFVE8
- Window: 2026-11-12 18:29 to 2026-11-14 00:18 UTC (1d 5h 48m)
- Impact: minor, status: resolved
- Components: —

<a id="ci-44d9765457f6f0b90ddc1a0d"></a>

- Canonical ID: `ci-44d9765457f6f0b90ddc1a0d` (1 source records).
  - [openai/01M78GAT3AK9P9W8Z230KHFVE8](../../data/2026/W46/openai.json); vendor URL: https://status.openai.com/incidents/01M78GAT3AK9P9W8Z230KHFVE8
- Start time basis: `vendor_reported`.

  - `2026-11-12 18:29` investigating — We’re investigating an issue with Microsoft connectors where new connections may be broken and tools are not working as expected. We’ll share an update when we know more.
  - `2026-11-13 16:20` monitoring — We’re seeing signs of recovery for the Microsoft 365 connector and are continuing to monitor the affected services to confirm recovery.
  - `2026-11-14 00:18` resolved — All affected services are operating normally again. We’re sorry for the disruption, and we appreciate your patience.

## 5. cloudflare — Cloudflare is seeing an issue with Workers API availability

- URL: https://www.cloudflarestatus.com/incidents/e7ktq32gphhe
- Window: 2026-11-13 15:31 to not yet (21h 5m (ongoing))
- Impact: minor, status: monitoring
- Components: Workers

<a id="ci-18358cb788c1961aed2943be"></a>

- Canonical ID: `ci-18358cb788c1961aed2943be` (1 source records).
  - [cloudflare/e7ktq32gphhe](../../data/2026/W46/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/e7ktq32gphhe
- Start time basis: `vendor_reported`.

  - `2026-11-13 15:31` investigating — Cloudflare is investigating an issue impacting the Cloudflare Workers API, customers may be seeing an increase in 5xx errors deploying Workers or doing other API actions.
  - `2026-11-13 15:39` identified — The issue has been identified and a fix is being implemented.
  - `2026-11-13 15:42` monitoring — A fix has been implemented and we are monitoring the results.

## Also above threshold (28)

- `openai` Delayed Costs Data in the Compliance API (20h 26m (ongoing)) — https://status.openai.com/incidents/01M7ATRS6BMZ61E4JDHJGM79Y0
- `cloudflare` Cloudflare is currently having an issue with the r8.im registry used for Replicate model publishing and builds is currently restricted (9h 56m) — https://www.cloudflarestatus.com/incidents/mc91s22odfvf
- `cloudflare` Changes made to Workers may be delayed or disrupted (9h 31m) — https://www.cloudflarestatus.com/incidents/1x0tbxpmmttu
- `slack` Trouble loading Slack and saving changes. (3h 34m) — https://slack-status.com/2026-11/b8bb993128ac94f5
- `grafana` Mimir query unexpected response with status code 500 (8h 31m) — https://stspg.io/fvnhy2nn4xq6
- `anthropic` Elevated errors on claude.ai, the Claude API, and Claude Code (2h 28m) — https://stspg.io/264r547qtkky
- `openai` Unable to start new Workspace Agent runs (9h 42m) — https://status.openai.com/incidents/01M76CW2D7FM7K7F8NJ4W0G08N
- `anthropic` Elevated errors on claude.ai, Claude Code, and platform.claude.com (1h 59m) — https://stspg.io/6kqyd0c5b1gr
- `aws` Elevated API issues and Launch Failures (5h 30m) — https://health.aws.amazon.com/health/status
- `cloudflare` Durable Objects increased error rates in North America (5h 5m) — https://www.cloudflarestatus.com/incidents/9zxr2bkqoct7
- `anthropic` Some users unable to sign up with Apple on the Claude iOS app (5h 25m (ongoing)) — https://stspg.io/4858wf40dx89
- `cloudflare` Network Performance Issues in Eastern North America (3h 37m) — https://www.cloudflarestatus.com/incidents/88hc76jnq2gn
- `grafana` Degraded performance for Grafana Cloud Metrics in prod-us-central-7 (1h 35m) — https://stspg.io/b9w7n973q0st
- `stripe` Elevated API Processing Errors (4h 45m) — https://stspg.io/t1vjtvqqncxf
- `github` Incident with Webhooks and Actions (1h 50m) — https://stspg.io/4vdghm7tycm4
- `cloudflare` Replicate Issue (3h 3m) — https://www.cloudflarestatus.com/incidents/1na2mbfshgdh
- `github` Incident with Git Operations and Pull Requests (1h 44m) — https://stspg.io/16nsnk2th3rv
- `github` Increased timeouts due to surge in anonymous users (3h 10m) — https://stspg.io/jztczfplj6yb
- `cloudflare` SSL Certificate Provisioning Issue (2h 22m) — https://www.cloudflarestatus.com/incidents/21vwp89eiyjx
- `grafana` Elevated query errors for Grafana Cloud Metrics in prod-us-east-3 (4h 13m) — https://stspg.io/k3cgsqwglmzn
- `cloudflare` Network Performance Issues in St. Louis, USA (2h 20m) — https://www.cloudflarestatus.com/incidents/1h9c85fz6bm3
- `github` Incident with Pull Requests (2h 42m) — https://stspg.io/qzj7qjdcr8jc
- `openai` Incorrect cost data in Costs API exports (2h 20m) — https://status.openai.com/incidents/01M71W79MAGEXEX93QTVTBYJ4K
- `github` Incident with Actions and Checks API (46m) — https://stspg.io/7yw63xl9ztrv
- `github` Impact across multiple GitHub products including Actions, Pages, and Copilot products. (33m) — https://stspg.io/jdjlq5zjyl48
- `github` Disruption with Email Delivery (2h 7m) — https://stspg.io/lr8v4v8t7gs2
- `github` Degraded availability for GPT-5.3 Codex and GPT-5.6 models in Copilot (49m) — https://stspg.io/t5gqmprtrdf2
- `anthropic` Elevated purchasing errors on claude.ai (1h 28m) — https://stspg.io/qqf00d9r7pm4
