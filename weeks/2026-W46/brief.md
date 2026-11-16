# Week 2026-W46 — top 5 of 35 above threshold

2026-11-09 to 2026-11-15 (UTC).

58 source records (91 saved revisions) → 58 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

> **2 of 14 sources unavailable: datadog, yandexcloud.**
> Their incidents are missing from this week, and the feeds are
> too shallow and unpaginated to backfill afterwards.
> - `datadog` — last successful fetch 2026-10-29 11:52, 18 days before this snapshot. Error: GET https://status.datadoghq.com/api/v2/incidents.json: status 403
> - `yandexcloud` — last successful fetch 2026-10-29 11:52, 18 days before this snapshot. Error: GET https://status.cloud.yandex.ru/api/incidents?lang=en&page=4: status 403

## 1. cloudflare — Changes to Cloudflare Workers are not producing Event Subscription events

- URL: https://www.cloudflarestatus.com/incidents/0ocqieaifg7h
- Window: 2026-11-10 01:53 UTC to now (6d 14h 41m, ongoing)
- Impact: minor, status: identified
- Components: Queues, Workers

<a id="ci-0a30f34c752fb49c0f4d211c"></a>

- Canonical ID: `ci-0a30f34c752fb49c0f4d211c` (1 source records).
  - [cloudflare/0ocqieaifg7h](../../data/2026/W46/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/0ocqieaifg7h
- Start time basis: `vendor_reported`.

  - `2026-11-10 01:53` investigating — Changes to Cloudflare Workers are temporarily not producing Event Subscription events. We have intentionally disabled this feature due to an on-going issue. We are actively working to restore it, but this, regrettably, may take a while. Changes to Workers themselves are working normally. Already deployed Workers are entirely unaffected.
  - `2026-11-11 05:05` identified — The issue has been identified and a fix is being implemented.

## 2. anthropic — Some users unable to sign up with Apple on the Claude iOS app

- URL: https://stspg.io/4858wf40dx89
- Window: 2026-11-14 07:11 UTC to now (2d 9h 23m, ongoing)
- Impact: major, status: investigating
- Components: claude.ai

<a id="ci-7d0c45c4220e872464de13e2"></a>

- Canonical ID: `ci-7d0c45c4220e872464de13e2` (1 source records).
  - [anthropic/ct4c2pwf2xpj](../../data/2026/W46/anthropic.json); vendor URL: https://stspg.io/4858wf40dx89
- Start time basis: `vendor_reported`.

  - `2026-11-14 07:11` investigating — We are investigating an issue with Sign in with Apple on the Claude iOS app, which is mainly affecting new accounts. New users can sign up with Google or email instead. We will provide an update as soon as possible.
  - `2026-11-14 20:12` investigating — We are continuing to investigate an issue with Sign in with Apple on the Claude iOS app, which is mainly affecting new accounts. New users can sign up with Google or email instead. We will provide another update as soon as possible.

## 3. openai — Delayed Costs Data in the Compliance API

- URL: https://status.openai.com/incidents/01M7ATRS6BMZ61E4JDHJGM79Y0
- Window: 2026-11-13 16:10 UTC to 2026-11-16 00:17 UTC (2d 8h 7m)
- Impact: minor, status: resolved
- Components: —

<a id="ci-f788538c36e9a228e97f389c"></a>

- Canonical ID: `ci-f788538c36e9a228e97f389c` (1 source records).
  - [openai/01M7ATRS6BMZ61E4JDHJGM79Y0](../../data/2026/W46/openai.json); vendor URL: https://status.openai.com/incidents/01M7ATRS6BMZ61E4JDHJGM79Y0
- Start time basis: `vendor_reported`.

  - `2026-11-13 16:10` identified — We have identified an issue causing delays in publishing Organization cost data through the Costs API. API and ChatGPT usage are not affected. We are working to restore normal data freshness.
  - `2026-11-13 18:56` monitoring — We’ve applied a mitigation and are monitoring the affected services to confirm recovery. Delays in cost data are still expected.
  - `2026-11-14 05:31` monitoring — We are monitoring the affected services to confirm recovery. Delays in cost data are still expected.
  - `2026-11-16 00:17` resolved — All affected services are operating normally again. We’re sorry for the disruption, and we appreciate your patience.

## 4. anthropic — Elevated errors affecting the Microsoft 365 connector

- URL: https://stspg.io/m0rbx38nmy1c
- Window: 2026-11-12 17:52 UTC to 2026-11-13 21:49 UTC (1d 3h 56m)
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

## 5. cloudflare — Replicate Model Registry Unavailable

- URL: https://www.cloudflarestatus.com/incidents/yu63fhh2peb7
- Window: 2026-11-11 04:33 UTC to 2026-11-13 01:44 UTC (1d 21h 10m)
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

## Also above threshold (30)

- `openai` Elevated errors affecting the Microsoft 365 connector (1d 5h 48m) — https://status.openai.com/incidents/01M78GAT3AK9P9W8Z230KHFVE8
- `cloudflare` Cloudflare is seeing an issue with Workers API availability (22h 6m) — https://www.cloudflarestatus.com/incidents/e7ktq32gphhe
- `cloudflare` Cloudflare is currently having an issue with the r8.im registry used for Replicate model publishing and builds is currently restricted (9h 56m) — https://www.cloudflarestatus.com/incidents/mc91s22odfvf
- `cloudflare` Changes made to Workers may be delayed or disrupted (9h 31m) — https://www.cloudflarestatus.com/incidents/1x0tbxpmmttu
- `slack` Trouble loading Slack and saving changes. (3h 34m) — https://slack-status.com/2026-11/b8bb993128ac94f5
- `grafana` Mimir query unexpected response with status code 500 (8h 31m) — https://stspg.io/fvnhy2nn4xq6
- `anthropic` Elevated errors on claude.ai, the Claude API, and Claude Code (2h 28m) — https://stspg.io/264r547qtkky
- `openai` Unable to start new Workspace Agent runs (9h 42m) — https://status.openai.com/incidents/01M76CW2D7FM7K7F8NJ4W0G08N
- `anthropic` Elevated errors on claude.ai, Claude Code, and platform.claude.com (1h 59m) — https://stspg.io/6kqyd0c5b1gr
- `aws` Elevated API issues and Launch Failures (5h 30m) — https://health.aws.amazon.com/health/status
- `cloudflare` Durable Objects increased error rates in North America (5h 5m) — https://www.cloudflarestatus.com/incidents/9zxr2bkqoct7
- `cloudflare` Network Performance Issues in Eastern North America (3h 37m) — https://www.cloudflarestatus.com/incidents/88hc76jnq2gn
- `grafana` Degraded performance for Grafana Cloud Metrics in prod-us-central-7 (1h 35m) — https://stspg.io/b9w7n973q0st
- `stripe` Elevated API Processing Errors (4h 45m) — https://stspg.io/t1vjtvqqncxf
- `stripe` Elevated PayNow errors (4h 37m) — https://stspg.io/fgz0c6z26b43
- `github` Incident with Webhooks and Actions (1h 50m) — https://stspg.io/4vdghm7tycm4
- `cloudflare` Replicate Issue (3h 3m) — https://www.cloudflarestatus.com/incidents/1na2mbfshgdh
- `github` Incident with Git Operations and Pull Requests (1h 44m) — https://stspg.io/16nsnk2th3rv
- `github` Increased timeouts due to surge in anonymous users (3h 10m) — https://stspg.io/jztczfplj6yb
- `aws` Increased Connectivity issues in US-EAST-1 (2h 15m) — https://health.aws.amazon.com/health/status
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
