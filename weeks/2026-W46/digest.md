# Week 2026-W46

2026-11-09 to 2026-11-15 (UTC). Generated 2026-11-16 16:35.

58 canonical incidents, 35 above threshold.

58 source records (91 saved revisions) → 58 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

> **2 of 14 sources unavailable: datadog, yandexcloud.**
> Their incidents are missing from this week, and the feeds are
> too shallow and unpaginated to backfill afterwards.
> - `datadog` — last successful fetch 2026-10-29 11:52, 18 days before this snapshot. Error: GET https://status.datadoghq.com/api/v2/incidents.json: status 403
> - `yandexcloud` — last successful fetch 2026-10-29 11:52, 18 days before this snapshot. Error: GET https://status.cloud.yandex.ru/api/incidents?lang=en&page=4: status 403

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | cloudflare | [Changes to Cloudflare Workers are not producing Event Subscription events](https://www.cloudflarestatus.com/incidents/0ocqieaifg7h) | 2026-11-10 01:53 | 6d 14h 41m (ongoing) | minor | Queues, Workers | 34276 = 9521 × 2 × 1.8 |
| 2 | anthropic | [Some users unable to sign up with Apple on the Claude iOS app](https://stspg.io/4858wf40dx89) | 2026-11-14 07:11 | 2d 9h 23m (ongoing) | major | claude.ai | 5164 = 3443 × 1 × 1.5 |
| 3 | openai | [Delayed Costs Data in the Compliance API](https://status.openai.com/incidents/01M7ATRS6BMZ61E4JDHJGM79Y0) | 2026-11-13 16:10 | 2d 8h 7m | minor | — | 5050 = 3367 × 1 × 1.5 |
| 4 | anthropic | [Elevated errors affecting the Microsoft 365 connector](https://stspg.io/m0rbx38nmy1c) | 2026-11-12 17:52 | 1d 3h 56m | major | Claude Cowork, claude.ai | 5028 = 1676 × 2 × 1.5 |
| 5 | cloudflare | [Replicate Model Registry Unavailable](https://www.cloudflarestatus.com/incidents/yu63fhh2peb7) | 2026-11-11 04:33 | 1d 21h 10m | minor | Replicate | 4878 = 2710 × 1 × 1.8 |
| 6 | openai | [Elevated errors affecting the Microsoft 365 connector](https://status.openai.com/incidents/01M78GAT3AK9P9W8Z230KHFVE8) | 2026-11-12 18:29 | 1d 5h 48m | minor | — | 2682 = 1788 × 1 × 1.5 |
| 7 | cloudflare | [Cloudflare is seeing an issue with Workers API availability](https://www.cloudflarestatus.com/incidents/e7ktq32gphhe) | 2026-11-13 15:31 | 22h 6m | minor | Workers | 2387 = 1326 × 1 × 1.8 |
| 8 | cloudflare | [Cloudflare is currently having an issue with the r8.im registry used for Replicate model publishing and builds is currently restricted](https://www.cloudflarestatus.com/incidents/mc91s22odfvf) | 2026-11-13 14:55 | 9h 56m | minor | Replicate | 1073 = 596 × 1 × 1.8 |
| 9 | cloudflare | [Changes made to Workers may be delayed or disrupted](https://www.cloudflarestatus.com/incidents/1x0tbxpmmttu) | 2026-11-09 16:15 | 9h 31m | minor | Workers | 1028 = 571 × 1 × 1.8 |
| 10 | slack | [Trouble loading Slack and saving changes.](https://slack-status.com/2026-11/b8bb993128ac94f5) | 2026-11-11 15:24 | 3h 34m | unknown | Apps/Integrations/APIs, Connectivity, Messaging, Workspace/Org Administration | 1027 = 214 × 4 × 1.2 |
| 11 | grafana | [Mimir query unexpected response with status code 500](https://stspg.io/fvnhy2nn4xq6) | 2026-11-10 03:38 | 8h 31m | minor | AWS US East - prod-us-east-0: Querying, GCP US Central - prod-us-central-0: Querying | 1022 = 511 × 2 × 1.0 |
| 12 | anthropic | [Elevated errors on claude.ai, the Claude API, and Claude Code](https://stspg.io/264r547qtkky) | 2026-11-13 13:09 | 2h 28m | major | Claude API (api.anthropic.com), Claude Code, Claude Cowork, claude.ai | 888 = 148 × 4 × 1.5 |
| 13 | openai | [Unable to start new Workspace Agent runs](https://status.openai.com/incidents/01M76CW2D7FM7K7F8NJ4W0G08N) | 2026-11-11 22:50 | 9h 42m | critical | — | 873 = 582 × 1 × 1.5 |
| 14 | anthropic | [Elevated errors on claude.ai, Claude Code, and platform.claude.com](https://stspg.io/6kqyd0c5b1gr) | 2026-11-13 09:58 | 1h 59m | major | Claude Code, Claude Console (platform.claude.com), Claude Cowork, claude.ai | 714 = 119 × 4 × 1.5 |
| 15 | aws | [Elevated API issues and Launch Failures](https://health.aws.amazon.com/health/status) | 2026-11-11 15:40 | 5h 30m | unknown | ec2-ap-northeast-1 | 660 = 330 × 1 × 2.0 |
| 16 | cloudflare | [Durable Objects increased error rates in North America](https://www.cloudflarestatus.com/incidents/9zxr2bkqoct7) | 2026-11-12 22:39 | 5h 5m | minor | Durable Objects | 549 = 305 × 1 × 1.8 |
| 17 | cloudflare | [Network Performance Issues in Eastern North America](https://www.cloudflarestatus.com/incidents/88hc76jnq2gn) | 2026-11-09 16:34 | 3h 37m | minor | Network | 391 = 217 × 1 × 1.8 |
| 18 | grafana | [Degraded performance for Grafana Cloud Metrics in prod-us-central-7](https://stspg.io/b9w7n973q0st) | 2026-11-11 20:37 | 1h 35m | major | Azure US Central - prod-us-central-7: Ingestion, Azure US Central - prod-us-central-7: Querying, Azure US Central - us-central7, Azure US Central - us-central7: Ingestion | 380 = 95 × 4 × 1.0 |
| 19 | stripe | [Elevated API Processing Errors](https://stspg.io/t1vjtvqqncxf) | 2026-11-14 02:50 | 4h 45m | none | Stripe API | 342 = 285 × 1 × 1.2 |
| 20 | stripe | [Elevated PayNow errors](https://stspg.io/fgz0c6z26b43) | 2026-11-14 17:46 | 4h 37m | none | Acquirers and payment methods | 332 = 277 × 1 × 1.2 |
| 21 | github | [Incident with Webhooks and Actions](https://stspg.io/4vdghm7tycm4) | 2026-11-09 10:00 | 1h 50m | major | Actions, Webhooks | 330 = 110 × 2 × 1.5 |
| 22 | cloudflare | [Replicate Issue](https://www.cloudflarestatus.com/incidents/1na2mbfshgdh) | 2026-11-12 08:30 | 3h 3m | minor | Replicate | 329 = 183 × 1 × 1.8 |
| 23 | github | [Incident with Git Operations and Pull Requests](https://stspg.io/16nsnk2th3rv) | 2026-11-10 16:10 | 1h 44m | critical | Actions, Pull Requests | 312 = 104 × 2 × 1.5 |
| 24 | github | [Increased timeouts due to surge in anonymous users](https://stspg.io/jztczfplj6yb) | 2026-11-09 05:59 | 3h 10m | minor | — | 285 = 190 × 1 × 1.5 |
| 25 | aws | [Increased Connectivity issues in US-EAST-1](https://health.aws.amazon.com/health/status) | 2026-11-15 12:38 | 2h 15m | unknown | multipleservices-us-east-1 | 270 = 135 × 1 × 2.0 |
| 26 | cloudflare | [SSL Certificate Provisioning Issue](https://www.cloudflarestatus.com/incidents/21vwp89eiyjx) | 2026-11-11 07:45 | 2h 22m | minor | SSL Certificate Provisioning | 256 = 142 × 1 × 1.8 |
| 27 | grafana | [Elevated query errors for Grafana Cloud Metrics in prod-us-east-3](https://stspg.io/k3cgsqwglmzn) | 2026-11-11 20:21 | 4h 13m | minor | — | 253 = 253 × 1 × 1.0 |
| 28 | cloudflare | [Network Performance Issues in St. Louis, USA](https://www.cloudflarestatus.com/incidents/1h9c85fz6bm3) | 2026-11-10 22:50 | 2h 20m | minor | St. Louis, United States - (STL) | 252 = 140 × 1 × 1.8 |
| 29 | github | [Incident with Pull Requests](https://stspg.io/qzj7qjdcr8jc) | 2026-11-13 15:47 | 2h 42m | major | Pull Requests | 243 = 162 × 1 × 1.5 |
| 30 | openai | [Incorrect cost data in Costs API exports](https://status.openai.com/incidents/01M71W79MAGEXEX93QTVTBYJ4K) | 2026-11-10 04:42 | 2h 20m | none | — | 210 = 140 × 1 × 1.5 |
| 31 | github | [Incident with Actions and Checks API](https://stspg.io/7yw63xl9ztrv) | 2026-11-09 18:12 | 46m | critical | Actions, Pages, Webhooks | 207 = 46 × 3 × 1.5 |
| 32 | github | [Impact across multiple GitHub products including Actions, Pages, and Copilot products.](https://stspg.io/jdjlq5zjyl48) | 2026-11-11 20:48 | 33m | critical | Actions, Copilot, Pages, Webhooks | 198 = 33 × 4 × 1.5 |
| 33 | github | [Disruption with Email Delivery](https://stspg.io/lr8v4v8t7gs2) | 2026-11-12 19:36 | 2h 7m | minor | — | 190 = 127 × 1 × 1.5 |
| 34 | github | [Degraded availability for GPT-5.3 Codex and GPT-5.6 models in Copilot](https://stspg.io/t5gqmprtrdf2) | 2026-11-09 23:16 | 49m | major | Copilot, Copilot AI Model Providers | 147 = 49 × 2 × 1.5 |
| 35 | anthropic | [Elevated purchasing errors on claude.ai](https://stspg.io/qqf00d9r7pm4) | 2026-11-14 02:54 | 1h 28m | major | claude.ai | 132 = 88 × 1 × 1.5 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. cloudflare — Changes to Cloudflare Workers are not producing Event Subscription events

- URL: https://www.cloudflarestatus.com/incidents/0ocqieaifg7h
- Started: 2026-11-10 01:53 UTC
- Resolved: ongoing
- Duration: 6d 14h 41m
- Impact: minor, status: identified
- Components: Queues, Workers
- Canonical incident: [ci-0a30f34c752fb49c0f4d211c](#ci-0a30f34c752fb49c0f4d211c)

- `2026-11-10 01:53` **investigating** — Changes to Cloudflare Workers are temporarily not producing Event Subscription events. We have intentionally disabled this feature due to an on-going issue. We are actively working to restore it, but this, regrettably, may take a while. Changes to Workers themselves are working normally. Already deployed Workers are entirely unaffected.
- `2026-11-11 05:05` **identified** — The issue has been identified and a fix is being implemented.

### 2. anthropic — Some users unable to sign up with Apple on the Claude iOS app

- URL: https://stspg.io/4858wf40dx89
- Started: 2026-11-14 07:11 UTC
- Resolved: ongoing
- Duration: 2d 9h 23m
- Impact: major, status: investigating
- Components: claude.ai
- Canonical incident: [ci-7d0c45c4220e872464de13e2](#ci-7d0c45c4220e872464de13e2)

- `2026-11-14 07:11` **investigating** — We are investigating an issue with Sign in with Apple on the Claude iOS app, which is mainly affecting new accounts. New users can sign up with Google or email instead. We will provide an update as soon as possible.
- `2026-11-14 20:12` **investigating** — We are continuing to investigate an issue with Sign in with Apple on the Claude iOS app, which is mainly affecting new accounts. New users can sign up with Google or email instead. We will provide another update as soon as possible.

### 3. openai — Delayed Costs Data in the Compliance API

- URL: https://status.openai.com/incidents/01M7ATRS6BMZ61E4JDHJGM79Y0
- Started: 2026-11-13 16:10 UTC
- Resolved: 2026-11-16 00:17 UTC
- Duration: 2d 8h 7m
- Impact: minor, status: resolved
- Components: —
- Canonical incident: [ci-f788538c36e9a228e97f389c](#ci-f788538c36e9a228e97f389c)

- `2026-11-13 16:10` **identified** — We have identified an issue causing delays in publishing Organization cost data through the Costs API. API and ChatGPT usage are not affected. We are working to restore normal data freshness.
- `2026-11-13 18:56` **monitoring** — We’ve applied a mitigation and are monitoring the affected services to confirm recovery. Delays in cost data are still expected.
- `2026-11-14 05:31` **monitoring** — We are monitoring the affected services to confirm recovery. Delays in cost data are still expected.
- `2026-11-16 00:17` **resolved** — All affected services are operating normally again. We’re sorry for the disruption, and we appreciate your patience.

### 4. anthropic — Elevated errors affecting the Microsoft 365 connector

- URL: https://stspg.io/m0rbx38nmy1c
- Started: 2026-11-12 17:52 UTC
- Resolved: 2026-11-13 21:49 UTC
- Duration: 1d 3h 56m
- Impact: major, status: resolved
- Components: Claude Cowork, claude.ai
- Canonical incident: [ci-3aedfad25e1c72a10aaae7ca](#ci-3aedfad25e1c72a10aaae7ca)

- `2026-11-12 17:52` **investigating** — We are investigating elevated errors and slow responses affecting the Microsoft 365 connector. Requests to mail, SharePoint, Teams chat and calendar through the connector may fail, and new connections to the connector may not complete. We will provide an update as soon as possible.
- `2026-11-12 18:34` **identified** — We have identified an issue with an upstream provider affecting the Microsoft 365 connector. Mail, calendar and search tools are failing for most requests, and new connections to the connector may not complete. We are in contact with them and will provide an update as soon as possible.
- `2026-11-12 20:28` **identified** — We are continuing to work with our upstream provider on this issue. Mail, calendar and search tools in the Microsoft 365 connector are still failing for most requests. Claude for Outlook is also affected: mail and calendar actions may fail, and some users may not be able to open the add-in. We will provide another update within the hour.
- `2026-11-12 21:18` **identified** — We are continuing to work with our upstream provider, who is restarting affected servers. Mail, calendar and search tools in the Microsoft 365 connector are still failing for most requests, and Claude for Outlook remains affected. We will provide another update as soon as we have more information.
- `2026-11-12 22:23` **identified** — This issue is caused by a Microsoft 365 service incident, and mail, calendar and search tools in the Microsoft 365 connector are still failing for most requests. Claude for Outlook remains affected. You can follow Microsoft's progress at https://status.cloud.microsoft. We will update this post when service is restored.
- `2026-11-13 15:40` **monitoring** — We're seeing signs of recovery for the Microsoft 365 connector. Claude for Outlook is also recovering. Some SharePoint search requests are still returning errors. Microsoft's incident remains open, so errors may return. We're continuing to monitor and will update this post as soon as we have more info.
- `2026-11-13 21:49` **resolved** — This incident has been resolved. From 08:03 PT / 16:03 UTC on November 12 to 08:40 PT / 16:40 UTC on November 13, mail, calendar and search tools in the Microsoft 365 connector and Claude for Outlook returned errors for many requests because of a Microsoft 365 service incident. Microsoft is still completing its recovery, so you may see occasional errors or delays.

### 5. cloudflare — Replicate Model Registry Unavailable

- URL: https://www.cloudflarestatus.com/incidents/yu63fhh2peb7
- Started: 2026-11-11 04:33 UTC
- Resolved: 2026-11-13 01:44 UTC
- Duration: 1d 21h 10m
- Impact: minor, status: resolved
- Components: Replicate
- Canonical incident: [ci-63dc50e88eca748e5be37fd9](#ci-63dc50e88eca748e5be37fd9)

- `2026-11-11 04:33` **investigating** — Cloudflare is investigating issues with Replicate input uploads. We are working to understand the full impact and mitigate this problem. More updates to follow shortly
- `2026-11-11 05:11` **identified** — The issue has been identified and a fix is being implemented.
- `2026-11-11 07:33` **identified** — We are continuing to work on a fix for the issue. Current impact: Input uploads are degraded. Model builds and pushes via r8.im are unavailable.
- `2026-11-11 23:50` **identified** — We are continuing to work on a fix for the issue. Current impact: Input uploads are degraded. Model builds and pushes via r8.im are unavailable.
- `2026-11-12 22:55` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-11-13 01:44` **resolved** — This incident has been resolved.

### 6. openai — Elevated errors affecting the Microsoft 365 connector

- URL: https://status.openai.com/incidents/01M78GAT3AK9P9W8Z230KHFVE8
- Started: 2026-11-12 18:29 UTC
- Resolved: 2026-11-14 00:18 UTC
- Duration: 1d 5h 48m
- Impact: minor, status: resolved
- Components: —
- Canonical incident: [ci-44d9765457f6f0b90ddc1a0d](#ci-44d9765457f6f0b90ddc1a0d)

- `2026-11-12 18:29` **investigating** — We’re investigating an issue with Microsoft connectors where new connections may be broken and tools are not working as expected. We’ll share an update when we know more.
- `2026-11-13 16:20` **monitoring** — We’re seeing signs of recovery for the Microsoft 365 connector and are continuing to monitor the affected services to confirm recovery.
- `2026-11-14 00:18` **resolved** — All affected services are operating normally again. We’re sorry for the disruption, and we appreciate your patience.

### 7. cloudflare — Cloudflare is seeing an issue with Workers API availability

- URL: https://www.cloudflarestatus.com/incidents/e7ktq32gphhe
- Started: 2026-11-13 15:31 UTC
- Resolved: 2026-11-14 13:37 UTC
- Duration: 22h 6m
- Impact: minor, status: resolved
- Components: Workers
- Canonical incident: [ci-18358cb788c1961aed2943be](#ci-18358cb788c1961aed2943be)

- `2026-11-13 15:31` **investigating** — Cloudflare is investigating an issue impacting the Cloudflare Workers API, customers may be seeing an increase in 5xx errors deploying Workers or doing other API actions.
- `2026-11-13 15:39` **identified** — The issue has been identified and a fix is being implemented.
- `2026-11-13 15:42` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-11-14 13:37` **resolved** — This incident has been resolved.

### 8. cloudflare — Cloudflare is currently having an issue with the r8.im registry used for Replicate model publishing and builds is currently restricted

- URL: https://www.cloudflarestatus.com/incidents/mc91s22odfvf
- Started: 2026-11-13 14:55 UTC
- Resolved: 2026-11-14 00:51 UTC
- Duration: 9h 56m
- Impact: minor, status: resolved
- Components: Replicate
- Canonical incident: [ci-81ecb242f7e327a5c980a3ec](#ci-81ecb242f7e327a5c980a3ec)

- `2026-11-13 14:55` **identified** — The issue has been identified and a fix is being implemented.
- `2026-11-13 23:34` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-11-14 00:51` **resolved** — This incident has been resolved.

### 9. cloudflare — Changes made to Workers may be delayed or disrupted

- URL: https://www.cloudflarestatus.com/incidents/1x0tbxpmmttu
- Started: 2026-11-09 16:15 UTC
- Resolved: 2026-11-10 01:46 UTC
- Duration: 9h 31m
- Impact: minor, status: resolved
- Components: Workers
- Canonical incident: [ci-f943958e573df67fc42f3633](#ci-f943958e573df67fc42f3633)

- `2026-11-09 16:15` **investigating** — Cloudflare is aware of and investigating issues related to creating, editing and deleting Workers. Impact is limited to changing Workers — Workers that are already deployed are unaffected.
- `2026-11-09 16:22` **identified** — We are continuing to work on a fix for the issue.
- `2026-11-09 16:35` **monitoring** — We are continuing to monitor for any further issues.
- `2026-11-10 01:46` **resolved** — This incident has been resolved.

### 10. slack — Trouble loading Slack and saving changes.

- URL: https://slack-status.com/2026-11/b8bb993128ac94f5
- Started: 2026-11-11 15:24 UTC
- Resolved: 2026-11-11 18:58 UTC
- Duration: 3h 34m
- Impact: unknown, status: resolved
- Components: Apps/Integrations/APIs, Connectivity, Messaging, Workspace/Org Administration
- Canonical incident: [ci-6319520b55ed6af07842a935](#ci-6319520b55ed6af07842a935)

- `2026-11-11 15:24` **unknown** — We're investigating a new incident. Some users may be experiencing delays when loading Slack and when using Slack features. We're working to understand the full impact to users. Our team is actively investigating the issue and is focused on restoring full performance as quickly as possible. We'll provide an update as soon as possible, and not longer than 30 minutes, with further details on the impact and actions being taken.
- `2026-11-11 15:29` **unknown** — We are investigating an issue where some customers may encounter various errors across Slack, including failures when loading the app, saving preferences or profile changes, and deactivating members. Our team is actively working to resolve the issue, and we'll share our next update in 30 minutes or as soon as more information is available.
- `2026-11-11 15:55` **unknown** — We’re continuing to work on the Slack issue described in our last update. Some users may still encounter errors when loading the app, saving preferences or profile changes, and deactivating members. We’ll provide an update in 30 minutes or sooner if additional information becomes available.
- `2026-11-11 16:16` **unknown** — We’ve identified that some users may also be unable to log in, create channels, install apps, or publish workflows. We’re investigating a recent database change as the likely trigger leading to a database storage capacity issue. We began freeing storage and removed unhealthy database hosts. Some users may continue to see errors. We’re working to prevent related databases from being affected. We’ll provide an update in 30 minutes or sooner if additional information becomes available.
- `2026-11-11 16:27` **unknown** — Since about 15:51 UTC, errors have decreased and many users can load Slack, log in, and save changes again. Some users may still encounter errors while we continue to stabilize our database systems. We’re also taking steps to prevent related databases from running out of storage. We’ll provide an update in 30 minutes or sooner if additional information becomes available.
- `2026-11-11 16:48` **unknown** — Errors have decreased since about 15:51 UTC, and most users should now be able to load Slack, log in, create channels, and save changes. We’re continuing to stabilize our database systems and taking steps to prevent the issue from returning. We’ll provide an update in 30 minutes or sooner if additional information becomes available.
- `2026-11-11 17:02` **unknown** — Error rates have been at normal levels as of 16:15 UTC, and users should be able to load Slack, log in, create channels, and save changes. We’re continuing to restore full database redundancy and working on configuration changes to prevent the issue from returning. We’ll provide an update in 30 minutes or sooner if additional information becomes available.
- `2026-11-11 17:31` **unknown** — We’re continuing to monitor our database systems closely and haven’t seen the errors return based on our latest checks. Work to prevent this issue from happening again is still in progress. We’ll provide an update in 60 minutes or sooner if additional information becomes available.
- `2026-11-11 18:28` **unknown** — Systems remain stable, and we haven’t seen the errors return. We’re continuing to restore additional database capacity and are completing a planned change to further strengthen the service. We’ll provide an update in 60 minutes or sooner if additional information becomes available.
- `2026-11-11 18:58` **unknown** — The issue affecting Slack has been resolved, and users should be able to load Slack, log in, create channels, and save changes as expected. Error rates returned to normal levels at 16:15 UTC and have remained stable since. We’ve also brought additional database capacity online to keep systems healthy. We apologize for any inconvenience this caused. We’re continuing to investigate the cause and will take steps to help prevent this from happening again.

### 11. grafana — Mimir query unexpected response with status code 500

- URL: https://stspg.io/fvnhy2nn4xq6
- Started: 2026-11-10 03:38 UTC
- Resolved: 2026-11-10 12:09 UTC
- Duration: 8h 31m
- Impact: minor, status: resolved
- Components: AWS US East - prod-us-east-0: Querying, GCP US Central - prod-us-central-0: Querying
- Canonical incident: [ci-de339f578e09069268438191](#ci-de339f578e09069268438191)

- `2026-11-10 03:38` **identified** — We’ve identified an issue querying older data, and are already seeing recovery
- `2026-11-10 03:39` **identified** — We’ve identified an issue querying older data, and are already seeing recovery.
- `2026-11-10 04:42` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-11-10 09:53` **monitoring** — Our mitigation remains in place and query behavior in AWS US East (OH) and GCP US Central continues to improve. Some customers may still see occasional, brief query errors in the meantime.
- `2026-11-10 11:18` **monitoring** — Our mitigation remains in place and query behavior in AWS US East (OH) and GCP US Central continues to improve. Some customers may still see occasional, brief query errors in the meantime.
- `2026-11-10 12:09` **resolved** — From approximately 03:20 UTC on 10 November, some Grafana Cloud Metrics customers in AWS US East (OH) and GCP US Central saw errors, or failed queries, when reading older data. Service has fully recovered and this incident is now resolved.

### 12. anthropic — Elevated errors on claude.ai, the Claude API, and Claude Code

- URL: https://stspg.io/264r547qtkky
- Started: 2026-11-13 13:09 UTC
- Resolved: 2026-11-13 15:38 UTC
- Duration: 2h 28m
- Impact: major, status: resolved
- Components: Claude API (api.anthropic.com), Claude Code, Claude Cowork, claude.ai
- Canonical incident: [ci-f925f2043c9c282f805d95c5](#ci-f925f2043c9c282f805d95c5)

- `2026-11-13 13:09` **investigating** — We are investigating reports of degraded performance affecting claude.ai, the Claude API, and Claude Code. We will provide an update as soon as possible.
- `2026-11-13 13:52` **investigating** — We are investigating reports of degraded performance affecting claude.ai, the Claude API, Claude Code, and connectors (including Claude Docs and Claude Sheets).
- `2026-11-13 14:03` **investigating** — We are investigating reports of degraded performance affecting a variety of surfaces including claude.ai, the Claude API, Claude Code, connectors (including Claude Docs and Claude Sheets), Claude Security and Claude Managed Agents.
- `2026-11-13 14:36` **investigating** — We're seeing recovery on claude.ai and the Claude API, though errors may return while we continue working on this. Connectors, including Docs and Sheets, remain degraded, as do Claude in Meetings, Claude Managed Agents, Claude Security and connector use in Claude Code. Mobile apps are affected in the same way as claude.ai. We'll share another update soon.
- `2026-11-13 15:33` **monitoring** — Errors on claude.ai, including the mobile apps, the Claude API and Claude Code have returned to normal levels, and we are continuing to monitor. Existing connectors, including Docs and Sheets, are working again, and Claude Managed Agents and Claude Meetings are recovering. Adding a new connector on claude.ai is temporarily unavailable. Scheduled tasks and routines can run more than two hours late and some runs have been skipped, and Claude Security scans, reviews and project setup remain delayed. We'll share another update soon.
- `2026-11-13 15:38` **resolved** — The issue affecting multiple surfaces has been resolved.

### 13. openai — Unable to start new Workspace Agent runs

- URL: https://status.openai.com/incidents/01M76CW2D7FM7K7F8NJ4W0G08N
- Started: 2026-11-11 22:50 UTC
- Resolved: 2026-11-12 08:32 UTC
- Duration: 9h 42m
- Impact: critical, status: resolved
- Components: —
- Canonical incident: [ci-9df1118b495b8ee1e137b867](#ci-9df1118b495b8ee1e137b867)

- `2026-11-11 22:50` **investigating** — We are investigating an issue preventing users from starting new Workspace Agent runs. We will share another update as more information becomes available.
- `2026-11-12 08:32` **resolved** — All affected services are operating normally again. We’re sorry for the disruption, and we appreciate your patience.

### 14. anthropic — Elevated errors on claude.ai, Claude Code, and platform.claude.com

- URL: https://stspg.io/6kqyd0c5b1gr
- Started: 2026-11-13 09:58 UTC
- Resolved: 2026-11-13 11:58 UTC
- Duration: 1h 59m
- Impact: major, status: resolved
- Components: Claude Code, Claude Console (platform.claude.com), Claude Cowork, claude.ai
- Canonical incident: [ci-6d42be3a336cbfe96168081e](#ci-6d42be3a336cbfe96168081e)

- `2026-11-13 09:58` **investigating** — We are investigating an issue with users authenticating to Claude.ai and Claude Code. We will provide an update as soon as possible.
- `2026-11-13 11:05` **monitoring** — A fix has been deployed for the issue affecting Claude services and we are monitoring for recovery.
- `2026-11-13 11:58` **resolved** — The issue affecting claude.ai and Claude Cowork has been resolved.

### 15. aws — Elevated API issues and Launch Failures

- URL: https://health.aws.amazon.com/health/status
- Started: 2026-11-11 15:40 UTC
- Resolved: 2026-11-11 21:11 UTC
- Duration: 5h 30m
- Impact: unknown, status: resolved
- Components: ec2-ap-northeast-1
- Canonical incident: [ci-7e816e7f0841048155f8126a](#ci-7e816e7f0841048155f8126a)

- `2026-11-11 15:40` **investigating** — AP-NORTHEAST-1 リージョンにおいて、API エラー率の上昇および起動失敗が発生しております。起動失敗は、単一のアベイラビリティーゾーン (apne1-az1) に集中して発生していると見受けられます。現在、根本原因の特定および問題の軽減に取り組んでいます。 | We are experiencing elevated API error rates and launch failures in the AP-NORTHEAST-1 Region. Launch failures appear to be concentrated in a single Availability Zone (apne1-az1). We are working to identify the root cause and mitigate the issue.
- `2026-11-11 16:40` **investigating** — 現在 AP-NORTHEAST-1 リージョンにおいて、API エラー率の上昇および起動失敗が発生しております。起動失敗は引き続き単一のアベイラビリティーゾーン (apne1-az1) に集中しているように見受けられますが、他のアベイラビリティーゾーンにおいても少数の起動失敗が発生しております。 影響を受けているアベイラビリティーゾーンでは、EC2 および EBS リソースの CloudWatch アラームが Insufficient Data 状態に移行しており、Insufficient Data でアラートを発生するよう設定されているアラームは、アラーム状態に移行する可能性があります。影響を受けているアベイラビリティーゾーンにおける DescribeInstanceStatus API は、Insufficient Data を返しております。稼働中のインスタンスに影響はございませんが、インスタンスの正常性は、お客様ご自身のヘルステレメトリーに基づいてご判断くださいますようお願いいたします。 さらに、ECS および Fargate のタスク起動失敗、ELB のプロビジョニング遅延、NAT Gateway の作成および関連付け解除のワークフローの失敗など、他サービスにも影響が発生しております。 根本原因は特定済みであり、現在問題の軽減に取り組んでおります。次回の更新は、日本時間午前 2 時 45 分までに、もしくは新たな情報が入り次第、お知らせいたします。 | We are experiencing elevated API error rates and launch failures in the AP-NORTHEAST-1 Region. While launch failures still appear to be concentrated in a single Availability Zone (apne1-az1), we are now seeing low levels of launch failures in other availability zones as well. Cloudwatch alarms for EC2 and EBS resources in the affected Availability Zone are transitioning into Insufficient Data state and alarms configured to alert off of Insufficient Data may transition to Alert state. The DescribeInstanceStatus API in the affected Availability Zone is returning Insufficient Data. Running instances are not affected, and you should rely on your own health telemetry to determine the health of your instance. In addition, we are seeing impact to other services such as ECS and Fargate Task Launch failures, ELB provisioning delays and Nat Gateway failures for Create and Disassociate workflows. We have identified the root cause and are working to mitigate the issue. We will provide you with another update by 9:45 AM PST or sooner as new information becomes available.
- `2026-11-11 17:38` **investigating** — AP-NORTHEAST-1 リージョンにおける API エラー率の上昇および起動失敗について、回復の兆候が見られております。EC2 および EBS の CloudWatch ヘルスデータは、現在正常に報告されております。引き続き復旧作業の進捗を監視し、次回の更新は日本時間午前 3 時までに、もしくは新たな情報が入り次第、お知らせいたします。 | We are seeing early signs of recovery for the elevated API error rates and launch failures in the AP-NORTHEAST-1 Region. EC2 and EBS CloudWatch health data is now reporting correctly. We will continue to monitor the progress of our recovery efforts and will provide you with another update by 10:00 AM PST or sooner as new information becomes available.
- `2026-11-11 17:59` **investigating** — AP-NORTHEAST-1 リージョンにおいて、引き続き広範囲で回復の兆候が見られております。現在も緩和策の実施を継続しており、次回の進捗の更新は日本時間午前 3 時 45 分までに、もしくは新たな情報が入り次第、お知らせいたします。| We continue to see broad signs of recovery in the AP-NORTHEAST-1 Region. Our mitigation efforts are still underway, and we will provide you with another update on our progress by 10:45 AM PST or sooner as new information becomes available.
- `2026-11-11 18:54` **investigating** — 現在引き続き完全な復旧に向けて取り組んでおります。単一のアベイラビリティーゾーン (apne1-az1) において、新規作成または変更されたインスタンスへの接続の問題が発生していることを確認しております。この問題は、新しいインスタンスの起動、新しいクラスターの作成、Auto Scaling グループ、ゲートウェイ、および VPC ゲートウェイエンドポイントなどのアクションに影響を与えております。既存のリソースへの接続およびインスタンス間の通信には影響はございません。この新しい問題の根本原因は特定済みであり、現在問題の軽減に取り組んでおります。次回の更新は、日本時間午前 5 時までに、もしくは新たな情報が入り次第、お知らせいたします。| We continue to work towards full recovery. We are observing some connectivity issues to newly created or modified instances in a single Availability Zone (apne1-az1). This affects actions such as new instance launches, new cluster creation, autoscaling groups, gateways, and vpc gateway endpoints. Connectivity to existing resources and Instance to Instance communication are not impacted. We have identified the root cause of this new issue and are working to mitigate the issue. We will provide you with another update by 12:00 PM PST or sooner as new information becomes available.
- `2026-11-11 19:54` **investigating** — 引き続き完全な復旧に向けて取り組んでおります。お客様は、他の 2 つの正常なアベイラビリティーゾーン (apne1-az2、apne1-az3) で新しいインスタンスを起動することができます。また、影響を受けているアベイラビリティーゾーンにおける既存のインスタンスへの接続に影響を与えるような変更は行わないことをお客様におすすめいたします。これには、ネットワークインターフェイスやゲートウェイのアタッチおよびデタッチ、VPC ピアリングの変更などのアクションが含まれます。単一のアベイラビリティーゾーン (apne1-az1) において、新規作成または変更されたインスタンスへの断続的な接続の問題の解決に引き続き取り組んでおります。この問題は、新しいインスタンスの起動、新しいクラスターの作成、Auto Scaling グループ、ゲートウェイ、および VPC ゲートウェイエンドポイントなどのアクションに影響を与えております。既存のリソースへの接続およびインスタンス間の通信には影響はございません。次回の更新は、日本時間午前 6 時までに、もしくは新たな情報が入り次第、お知らせいたします。| We continue to work towards full recovery. Customers can launch new instances in the other two healthy Availability Zones (apne1-az2, apne1-az3). We also recommend that customers do not make changes that would impact connectivity to existing instances in the affected Availability Zone. This includes actions such as attaching and detaching network interfaces or gateways as well as VPC Peerings changes. We are continuing to work on resolving the intermittent connectivity issues to newly created or modified instances in a single Availability Zone (apne1-az1). This affects actions such as new instance launches, new cluster creation, autoscaling groups, gateways, and vpc gateway endpoints. Connectivity to existing resources and Instance to Instance communication are not impacted. We will provide you with another update by 1:00 PM PST or sooner as new information becomes available.
- `2026-11-11 21:11` **resolved** — 日本時間 2026 年 11 月 11 日 午後 11 時 46 分から 2026 年 11 月 12 日 午前 5 時 31 分にかけて、AP-NORTHEAST-1 リージョンの主に単一のアベイラビリティーゾーン (apne1-az1) において、断続的な API エラー率の上昇、および新規起動された EC2 インスタンスにおける接続の問題が発生しておりました。エンジニアが自動的に対応に加わり、直ちに調査を開始いたしました。最初の影響期間は、日本時間 2026 年 11 月 11 日 午後 11 時 46 分から 2026 年 11 月 12 日 午前 2 時 46 分であり、主に API エラーおよび EC2 インスタンスの起動失敗として現れておりました。Auto Scaling、ECS、RDS など、ワークフローを EC2 API に依存している他の AWS サービスも、この問題の影響を受けておりました。既存の EC2 インスタンスには影響はございませんでした。この最初の影響期間中、お客様には API のスロットリングまたは 5XX レスポンスが発生していた可能性がございます。また、一部の EC2 RunInstances API の呼び出しは成功したものの、インスタンスが正常に起動せず、TERMINATED 状態に移行した場合がございました。EC2 および EBS のヘルスデータが欠損または遅延し、その結果 CloudWatch アラームが INSUFFICIENT_DATA 状態に移行しておりました。影響を受けたサービスにおいて失敗した API リクエストやワークフローにつきましては、再試行いただくことを推奨いたします。 日本時間 午前 2 時 46 分に API エラーを緩和いたしましたが、この時点ではすべての問題は緩和されておりませんでした。EC2 API の回復を監視する中で、新規起動された EC2 インスタンス、およびネットワーク設定が変更された既存のインスタンスの両方で、接続の問題が発生していることを確認いたしました。これらの接続の問題は、EC2 インスタンスのネットワークの変更を伝播する内部サブシステムにおける遅延によるものであり、根本原因に関連するものでございました。日本時間 午前 5 時 31 分までに、問題を完全に解決いたしました。この間、接続の問題の軽減と根本原因の完全な把握の両面から、複数の対策を並行して実施しておりました。ECS、EKS、ELB、NAT Gateway など、EC2 インスタンスを起動する、または EC2 インスタンスのネットワーク設定を変更する他の AWS サービスも、この問題の影響を受けておりました。 影響を受けたサービスの一覧は以下のとおりです。問題は解決され、すべてのサービスは正常に稼働しております。| Between 6:46 AM PST and 12:31 PM PST, we experienced periods of API error rates, as well as connectivity issues for newly launched EC2 instances, primarily affecting a single Availability Zone (apne1-az1) in the AP-NORTHEAST-1 Region. Engineers were automatically engaged and immediately began investigating. The first impact window was between 6:46 AM PST and 9:46 AM PST, and primarily manifested as API errors and failed EC2 instance launches. Other AWS Services that rely on the EC2 APIs for their workflows, such as Autoscaling, ECS, and RDS, were also affected by this issue. Existing EC2 instances were not affected. During this first impact period, customers may have experienced API throttling or 5XX Responses. Additionally, some EC2 RunInstance API calls may have succeeded, but instances did not successfully launch, instead transitioning to a TERMINATED state. EC2 and EBS health data was missing or delayed, which resulted in CloudWatch alarms transitioning to an INSUFFICENT_DATA state. We recommend customers retry failed API requests or workflows that failed for any of the affected services. At 9:46 AM, we mitigated the API errors, but had not yet mitigated all issues. While monitoring recovery of the EC2 APIs, we observed that both newly launched EC2 instances and existing instances where the network configuration was modified were experiencing connectivity issues. These connectivity issues were the result of delays in an internal subsytem responsible for propagating EC2 instance networking changes, and related to the underlying root cause. By 12:31 PM PST, we had fully resolved the issue. During this time, we were working on multiple parallel paths to both mitigate the connectivity issues and fully understand the root cause. Other AWS Services that launch EC2 instances or modify EC2 instance networking configuration, such as ECS, EKS, ELB, and NAT Gateway, were also affected by this issue. A list of impacted services is available below. The issue has been resolved and all services are operating normally.

### 16. cloudflare — Durable Objects increased error rates in North America

- URL: https://www.cloudflarestatus.com/incidents/9zxr2bkqoct7
- Started: 2026-11-12 22:39 UTC
- Resolved: 2026-11-13 03:44 UTC
- Duration: 5h 5m
- Impact: minor, status: resolved
- Components: Durable Objects
- Canonical incident: [ci-b925591a04b02f7511d9a1fc](#ci-b925591a04b02f7511d9a1fc)

- `2026-11-12 22:39` **monitoring** — Durable Objects saw an increased error rate in North American colos due to a sudden increase in traffic. Impact has subsided. We are now monitoring.
- `2026-11-13 03:44` **resolved** — This incident has been resolved.

### 17. cloudflare — Network Performance Issues in Eastern North America

- URL: https://www.cloudflarestatus.com/incidents/88hc76jnq2gn
- Started: 2026-11-09 16:34 UTC
- Resolved: 2026-11-09 20:12 UTC
- Duration: 3h 37m
- Impact: minor, status: resolved
- Components: Network
- Canonical incident: [ci-51c613491673f28530db53c3](#ci-51c613491673f28530db53c3)

- `2026-11-09 16:34` **investigating** — Cloudflare is investigating issues with network performance in Eastern North America. We are working to analyse and mitigate this problem. More updates to follow shortly.
- `2026-11-09 16:56` **identified** — We are continuing to work on a fix for the issue.
- `2026-11-09 17:21` **monitoring** — We are continuing to monitor for any further issues.
- `2026-11-09 20:12` **resolved** — This incident has been resolved.

### 18. grafana — Degraded performance for Grafana Cloud Metrics in prod-us-central-7

- URL: https://stspg.io/b9w7n973q0st
- Started: 2026-11-11 20:37 UTC
- Resolved: 2026-11-11 22:13 UTC
- Duration: 1h 35m
- Impact: major, status: resolved
- Components: Azure US Central - prod-us-central-7: Ingestion, Azure US Central - prod-us-central-7: Querying, Azure US Central - us-central7, Azure US Central - us-central7: Ingestion
- Canonical incident: [ci-74bfb35ba2e60e544c598c72](#ci-74bfb35ba2e60e544c598c72)

- `2026-11-11 20:37` **investigating** — We are investigating an issue affecting Grafana Cloud Metrics, Logs, and Traces in the prod-us-central-7 region. Some customers may see failed or delayed writes, query errors, and failed alerting and recording rule evaluations.
- `2026-11-11 21:24` **investigating** — We are continuing to investigate this issue.
- `2026-11-11 21:27` **investigating** — We are continuing to investigate this issue.
- `2026-11-11 22:07` **monitoring** — Traffic has recovered and we are monitoring the results.
- `2026-11-11 22:13` **resolved** — Increased latency on a cloud provider database resulted in a 50m write and rule evaluation outage, starting about 2020 UTC. Metrics rule evaluations were last to recover, around 2110 UTC.

### 19. stripe — Elevated API Processing Errors

- URL: https://stspg.io/t1vjtvqqncxf
- Started: 2026-11-14 02:50 UTC
- Resolved: 2026-11-14 07:35 UTC
- Duration: 4h 45m
- Impact: none, status: resolved
- Components: Stripe API
- Canonical incident: [ci-d84d19af354ef2d1dfbe25b2](#ci-d84d19af354ef2d1dfbe25b2)

- `2026-11-14 02:50` **investigating** — We're currently investigating elevated rate limits affecting some merchants. We will post further updates once we have more information.
- `2026-11-14 03:03` **investigating** — We are seeing elevated HTTP processing errors on the Stripe API. Our engineers are engaged and investigating.
- `2026-11-14 03:44` **monitoring** — From 02:34 to 03:01 UTC on November 14th, users experienced elevated Stripe API error rates. Error rates have returned to normal, and our engineers are monitoring and validating that the issue is fully resolved.
- `2026-11-14 07:35` **resolved** — From 02:34 to 03:01 UTC on November 14th, we saw increased API error rates for the Stripe API. This is now resolved.

### 20. stripe — Elevated PayNow errors

- URL: https://stspg.io/fgz0c6z26b43
- Started: 2026-11-14 17:46 UTC
- Resolved: 2026-11-14 22:24 UTC
- Duration: 4h 37m
- Impact: none, status: resolved
- Components: Acquirers and payment methods
- Canonical incident: [ci-f2e12f4c1cfbbc05d541133a](#ci-f2e12f4c1cfbbc05d541133a)

- `2026-11-14 17:46` **investigating** — As of 17:31 UTC, we've seen elevated errors on PayNow payments. We're continuing to investigate and will post further updates.
- `2026-11-14 22:24` **resolved** — From 16:20-21:40 UTC, we saw elevated errors with the PayNow payment method. The issue was external to Stripe and is now resolved.

### 21. github — Incident with Webhooks and Actions

- URL: https://stspg.io/4vdghm7tycm4
- Started: 2026-11-09 10:00 UTC
- Resolved: 2026-11-09 11:50 UTC
- Duration: 1h 50m
- Impact: major, status: resolved
- Components: Actions, Webhooks
- Canonical incident: [ci-f233a5eaca549a7cd81af490](#ci-f233a5eaca549a7cd81af490)

- `2026-11-09 10:00` **investigating** — We are investigating reports of degraded performance for Actions and Webhooks
- `2026-11-09 10:28` **investigating** — We're investigating delays in Actions runs start times and webhook deliveries.
- `2026-11-09 10:47` **investigating** — We've applied a mitigation and are monitoring for recovery.
- `2026-11-09 11:06` **investigating** — The degradation affecting Actions has been mitigated. We are monitoring to ensure stability.
- `2026-11-09 11:17` **investigating** — Webhooks is experiencing degraded availability. We are continuing to investigate.
- `2026-11-09 11:42` **investigating** — Webhooks is experiencing degraded performance. We are continuing to investigate.
- `2026-11-09 11:45` **monitoring** — The degradation affecting Webhooks has been mitigated. We are monitoring to ensure stability.
- `2026-11-09 11:50` **resolved** — On November 9, 2026, from 09:23 to 11:40 UTC, one of our job queue clusters was degraded, delaying webhook deliveries, Actions workflow runs, and Dependabot and dependency graph updates. Until mitigation at 10:37 UTC, an average of 26% of webhook deliveries were delayed, peaking at 30%, and the backlog cleared by 11:40 UTC. Some deliveries were delayed by up to 3.5 hours, and delays for a small share began as early as 01:00 UTC. Actions skips events older than 30 minutes, so some workflow runs did not start. <br /><br />The incident was caused by a degraded host that had slowed one of the cluster's Kafka brokers since November 7. The broker stopped reporting the metrics our alerts rely on, so we did not identify it until it stalled the cluster on November 9. <br /><br />We mitigated the incident by removing the broker from service. We have since added an alert that pages us within 30 minutes when a host slows a broker this way, so we can remove the broker before it affects customers.

### 22. cloudflare — Replicate Issue

- URL: https://www.cloudflarestatus.com/incidents/1na2mbfshgdh
- Started: 2026-11-12 08:30 UTC
- Resolved: 2026-11-12 11:34 UTC
- Duration: 3h 3m
- Impact: minor, status: resolved
- Components: Replicate
- Canonical incident: [ci-e77f3bfa3f702ce06c0d0386](#ci-e77f3bfa3f702ce06c0d0386)

- `2026-11-12 08:30` **investigating** — Cloudflare Replicate customers may see an increase of 500 errors for predictions. More updates to follow shortly.
- `2026-11-12 09:34` **identified** — The issue has been identified and a fix is being implemented.
- `2026-11-12 11:34` **resolved** — This incident has been resolved.

### 23. github — Incident with Git Operations and Pull Requests

- URL: https://stspg.io/16nsnk2th3rv
- Started: 2026-11-10 16:10 UTC
- Resolved: 2026-11-10 17:55 UTC
- Duration: 1h 44m
- Impact: critical, status: resolved
- Components: Actions, Pull Requests
- Canonical incident: [ci-682c2f892744899d4df5e977](#ci-682c2f892744899d4df5e977)

- `2026-11-10 16:10` **investigating** — We are investigating reports of degraded performance for Actions and Pull Requests
- `2026-11-10 16:26` **investigating** — Actions and Pull Requests are experiencing degraded availability. We are continuing to investigate.
- `2026-11-10 16:31` **investigating** — We're investigating elevated response times for Git operations and pull requests. GitHub Actions workflows that are triggered from Pull Requests are delayed. We're working to mitigate the impact and will share updates as we learn more.
- `2026-11-10 16:39` **investigating** — Actions is experiencing degraded availability. We are continuing to investigate.
- `2026-11-10 17:07` **investigating** — We continue to investigate elevated response times for Git operations and pull requests. Some pull-request processing and Actions workflows are improving, but delays and failures persist. We're monitoring the impact and working to restore normal service.
- `2026-11-10 17:12` **investigating** — Pull Requests is experiencing degraded performance. We are continuing to investigate.
- `2026-11-10 17:28` **investigating** — Actions is operating normally.
- `2026-11-10 17:46` **investigating** — We continue to investigate elevated response times for Git operations and pull requests. Actions workflows are recovering. Pull request processing has shown some improvement, but users continue to experience delays and failures. We’re monitoring the impact and working to restore normal service.
- `2026-11-10 17:55` **resolved** — On November 10, 2026, from approximately 15:00 UTC to 17:55 UTC, GitHub experienced degraded performance affecting repository pages, pull requests, and GitHub Actions. Customers may have seen slow or failed page loads and delays processing pull request updates and merges. Some pull requests continued to display an earlier source-branch commit after the broader service degradation ended. GitHub Actions workflows were delayed for approximately 1.8% of events, peaking at about 7% between 16:30 and 16:50 UTC. Delays were typically 1.5 to 3.5 minutes, up to 10 minutes. Approximately 0.09% of pull request events did not start their workflows and needed to be re-triggered. <br /><br />The degradation resulted from overload in the shared infrastructure that supports Git operations. Recent changes increased background processing and the work required to compute pull request merges, adding pressure to the same systems used by customer-facing requests. We mitigated the incident by disabling a recent merge-processing optimization, pausing background analysis, and temporarily limiting lower-priority traffic. The temporary limits also restricted access to commit pages for users who were not signed in. Background processing continued to catch up after service health recovered. <br /><br />We have reduced repeated repository-analysis work and are adding workload-level concurrency controls to better protect customer-facing requests. We are also improving our procedures for safely reducing background load and making pull request synchronization more resilient to temporary failures.

### 24. github — Increased timeouts due to surge in anonymous users

- URL: https://stspg.io/jztczfplj6yb
- Started: 2026-11-09 05:59 UTC
- Resolved: 2026-11-09 09:09 UTC
- Duration: 3h 10m
- Impact: minor, status: resolved
- Components: —
- Canonical incident: [ci-5523009b2922e2f5a4c3e5f5](#ci-5523009b2922e2f5a4c3e5f5)

- `2026-11-09 05:59` **investigating** — We are investigating reports of impacted performance for some GitHub services.
- `2026-11-09 06:30` **investigating** — We are debugging a spike in anonymous traffic. Some users who are not logged in may see errors on the site. CI jobs and release downloads may fail with a 504 Gateway Timeout status if they make anonymous calls.
- `2026-11-09 06:51` **monitoring** — The degradation has been mitigated. We are monitoring to ensure stability.
- `2026-11-09 07:03` **monitoring** — The problem with anonymous traffic receiving 504 status codes has been mitigated. We are continuing to observe for an abundance of caution.
- `2026-11-09 07:43` **investigating** — We have observed a repeated increase in 504 errors for anonymous and unauthenticated traffic, including users who are not logged in and release downloads. We are continuing to investigate and apply mitigations.
- `2026-11-09 08:16` **investigating** — We are continuing to observe high rates of 504 errors for anonymous requests. We are continuing to work on mitigation. Logged-in users and authenticated requests are unaffected.
- `2026-11-09 08:47` **investigating** — Error rates have declined again, and we’re preparing targeted traffic controls while monitoring for a recurrence. Logged-in users and authenticated requests have not been affected.
- `2026-11-09 09:05` **monitoring** — The degradation has been mitigated. We are monitoring to ensure stability.
- `2026-11-09 09:09` **monitoring** — We are resolving this incident as error rates have remained low and we have prepared additional protections.
- `2026-11-09 09:09` **resolved** — This incident has been resolved. Thank you for your patience and understanding as we addressed this issue. A detailed root cause analysis will be shared as soon as it is available.

### 25. aws — Increased Connectivity issues in US-EAST-1

- URL: https://health.aws.amazon.com/health/status
- Started: 2026-11-15 12:38 UTC
- Resolved: 2026-11-15 14:53 UTC
- Duration: 2h 15m
- Impact: unknown, status: resolved
- Components: multipleservices-us-east-1
- Canonical incident: [ci-d0f875d5ab8f69a2f190b653](#ci-d0f875d5ab8f69a2f190b653)

- `2026-11-15 12:38` **investigating** — We are investigating reports of Connectivity issues in the US-EAST-1 Region. IAM Identity center sign-in requests and AWS SignUp are also impacted at this time.
- `2026-11-15 13:15` **investigating** — We continue to investigate Connectivity issues in the US-EAST-1 Region. IAM Identity center sign-in requests and AWS SignUp are impacted at this time. We can confirm errors when attempted to authenticate when using identity center or when using services that rely on IAM Identity center. For Kiro users, logins from social sign in methods are not impacted at this time. There are no workarounds for impacted customers.
- `2026-11-15 13:22` **investigating** — We are seeing early signs of recovery for authentication via IAM Identity Center and SignUp.
- `2026-11-15 14:23` **investigating** — We can confirm recovery for most AWS Services. We have fully mitigated the issue authentication errors from IAM Identity Center that was impacting Kiro, Connect, and other services which rely on IAM Identity Center for authentication. We continue to investigate an issue that is affecting a subset of Network Load Balancers in a single cell, in a single Availability Zone (use1-az1) in the US-EAST-1 Region. For customers who need immediate resolution, we recommend creating a new NLB with targets in multiple Availability Zones. We are actively working on multiple parallel paths to mitigate the issue. We will provide an update by 7:30 AM, or sooner if we have additional information to share.
- `2026-11-15 14:32` **investigating** — We are seeing significant signs of recovery and are continuing to monitor for full recovery.
- `2026-11-15 14:53` **resolved** — Between 3:05 AM and 6:10 AM PST we experienced signup errors, authentication issues, and connectivity issues in the US-EAST-1 Region. Customers using IAM Identity Center for authentication into the AWS Console, or for authentication into other AWS Services such as Kiro, Connect, etc would have been affected during this time. Engineers were automatically engaged at 3:10 AM PST and immediately began investigating the root cause, while in parallel working to mitigate both impacts. At 5:20, we fully mitigated the issue affecting authentication, and most AWS Services observed full recovery at this time. A subset of customers using NLB in a single cell in use1-az1 continued to experience connectivity issues until we fully mitigated the underlying issue at 6:21 AM PST. We recommend customers retry any failed requests. The issue has been resolved and all AWS Services are operating normally.

### 26. cloudflare — SSL Certificate Provisioning Issue

- URL: https://www.cloudflarestatus.com/incidents/21vwp89eiyjx
- Started: 2026-11-11 07:45 UTC
- Resolved: 2026-11-11 10:07 UTC
- Duration: 2h 22m
- Impact: minor, status: resolved
- Components: SSL Certificate Provisioning
- Canonical incident: [ci-65f3ac868ce78e42cb1b5267](#ci-65f3ac868ce78e42cb1b5267)

- `2026-11-11 07:45` **investigating** — Cloudflare is investigating an issue deploying a subset of Custom certificates. Customers are advised to temporarily pause certificate redeployments until this incident is resolved SSL certificates already in production remain unaffected and are operating normally.
- `2026-11-11 08:17` **identified** — The issue has been identified and a fix is being implemented.
- `2026-11-11 09:01` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-11-11 10:07` **resolved** — This incident has been resolved.

### 27. grafana — Elevated query errors for Grafana Cloud Metrics in prod-us-east-3

- URL: https://stspg.io/k3cgsqwglmzn
- Started: 2026-11-11 20:21 UTC
- Resolved: 2026-11-12 00:35 UTC
- Duration: 4h 13m
- Impact: minor, status: resolved
- Components: —
- Canonical incident: [ci-84dd75fecfe0fa655d5a32e8](#ci-84dd75fecfe0fa655d5a32e8)

- `2026-11-11 20:21` **monitoring** — Between 17:21 and 17:41 UTC, a small percentage of metrics queries for some Grafana Cloud Metrics customers in the prod-us-east-3 region failed. Ingestion was not affected. Query errors have stopped, and we are monitoring the service to confirm it stays stable.
- `2026-11-12 00:35` **resolved** — This incident has been resolved.

### 28. cloudflare — Network Performance Issues in St. Louis, USA

- URL: https://www.cloudflarestatus.com/incidents/1h9c85fz6bm3
- Started: 2026-11-10 22:50 UTC
- Resolved: 2026-11-11 01:10 UTC
- Duration: 2h 20m
- Impact: minor, status: resolved
- Components: St. Louis, United States - (STL)
- Canonical incident: [ci-a65f27cd1c724c257a51f457](#ci-a65f27cd1c724c257a51f457)

- `2026-11-10 22:50` **monitoring** — Cloudflare has identified issues with network performance in St. Louis (STL), USA. A fix has been implemented and we're monitoring the results.
- `2026-11-11 01:10` **resolved** — This incident has been resolved.

### 29. github — Incident with Pull Requests

- URL: https://stspg.io/qzj7qjdcr8jc
- Started: 2026-11-13 15:47 UTC
- Resolved: 2026-11-13 18:29 UTC
- Duration: 2h 42m
- Impact: major, status: resolved
- Components: Pull Requests
- Canonical incident: [ci-6fd6e5add514573bdc00bf7e](#ci-6fd6e5add514573bdc00bf7e)

- `2026-11-13 15:47` **investigating** — We are investigating reports of degraded availability for Pull Requests
- `2026-11-13 15:49` **investigating** — We’re investigating an issue with Pull Requests where some attempts to merge are timing out.
- `2026-11-13 16:09` **investigating** — Pull Requests is experiencing degraded performance. We are continuing to investigate.
- `2026-11-13 16:12` **investigating** — We’ve taken mitigating action to resolve the problem of merge requests timing out. This issue initially impacted 3% of requests and is now down to less than 1%. We continue to monitor recovery.
- `2026-11-13 16:52` **investigating** — We’re continuing to investigate and monitor ~1% of merge requests timing out.
- `2026-11-13 17:28` **monitoring** — The degradation affecting Pull Requests has been mitigated. We are monitoring to ensure stability.
- `2026-11-13 17:46` **monitoring** — We continue to monitor the stability of merge requests.
- `2026-11-13 18:29` **resolved** — This incident has been resolved. Thank you for your patience and understanding as we addressed this issue. A detailed root cause analysis will be shared as soon as it is available.

### 30. openai — Incorrect cost data in Costs API exports

- URL: https://status.openai.com/incidents/01M71W79MAGEXEX93QTVTBYJ4K
- Started: 2026-11-10 04:42 UTC
- Resolved: 2026-11-10 07:02 UTC
- Duration: 2h 20m
- Impact: none, status: resolved
- Components: —
- Canonical incident: [ci-e83b177687e865bcd9a02619](#ci-e83b177687e865bcd9a02619)

- `2026-11-10 04:42` **monitoring** — We have identified an issue causing incorrect cost data in Costs API exports published between November 6 and 9. We are working to correct the affected records and will provide an update once recovery is confirmed.
- `2026-11-10 07:02` **resolved** — All affected services are operating normally again. We’re sorry for the disruption, and we appreciate your patience.

### 31. github — Incident with Actions and Checks API

- URL: https://stspg.io/7yw63xl9ztrv
- Started: 2026-11-09 18:12 UTC
- Resolved: 2026-11-09 18:59 UTC
- Duration: 46m
- Impact: critical, status: resolved
- Components: Actions, Pages, Webhooks
- Canonical incident: [ci-059009c46813833b5e969eac](#ci-059009c46813833b5e969eac)

- `2026-11-09 18:12` **investigating** — We are investigating reports of degraded availability for Actions
- `2026-11-09 18:14` **investigating** — Actions is experiencing degraded availability. We are continuing to investigate.
- `2026-11-09 18:23` **investigating** — Webhooks is experiencing degraded performance. We are continuing to investigate.
- `2026-11-09 18:24` **investigating** — We're investigating elevated errors affecting GitHub Actions and the Checks API. We'll share another update as we learn more.
- `2026-11-09 18:30` **investigating** — Pages is experiencing degraded performance. We are continuing to investigate.
- `2026-11-09 18:43` **monitoring** — The degradation affecting Actions, Pages and Webhooks has been mitigated. We are monitoring to ensure stability.
- `2026-11-09 18:59` **resolved** — On November 9, 2026, between 18:07 and 18:59 UTC, GitHub Actions experienced degraded performance. During this time, some customers saw workflow runs fail to start, update or report status correctly. An average of 28% of workflow runs were affected, peaking at 65% during the worst period of the incident. Once Actions services recovered, some jobs were left in a state that limited available concurrency until the runs were released. <br /><br />The incident occurred during a planned migration of a core Actions database from GitHub’s datacenter to a new cloud region. After the migration completed, we shifted Actions service traffic to the new region. The volume of requests exceeded the capacity and rate limits configured there, so a portion of requests were rejected and workflow updates failed. Dependent services, including Webhooks, Pages and Copilot code review, saw secondary impact. <br /><br />We mitigated impact by shifting some of the traffic back to the original region, allowing service traffic to remain within the configured limits. We then increased capacity and rate limits in the new region before gradually shifting traffic back. While this was a one-time migration event for this cluster that is now complete, we are improving how we validate capacity and limits before shifting traffic during remaining migrations, including shifting traffic in smaller steps. To address the stuck runs left after the incident, we are improving our ability to detect and mitigate these runs for future cases as well as readying an API to enable customers to address this themselves.

### 32. github — Impact across multiple GitHub products including Actions, Pages, and Copilot products.

- URL: https://stspg.io/jdjlq5zjyl48
- Started: 2026-11-11 20:48 UTC
- Resolved: 2026-11-11 21:21 UTC
- Duration: 33m
- Impact: critical, status: resolved
- Components: Actions, Copilot, Pages, Webhooks
- Canonical incident: [ci-f28b53b3942f32847c952628](#ci-f28b53b3942f32847c952628)

- `2026-11-11 20:48` **investigating** — We are investigating reports of degraded availability for Pages
- `2026-11-11 20:48` **investigating** — Actions is experiencing degraded availability. We are continuing to investigate.
- `2026-11-11 20:48` **investigating** — Webhooks is experiencing degraded performance. We are continuing to investigate.
- `2026-11-11 20:56` **investigating** — Pages is experiencing degraded availability. We are continuing to investigate.
- `2026-11-11 20:56` **investigating** — We’re investigating issues impacting multiple GitHub products including Actions, Pages, and Copilot products.
- `2026-11-11 20:57` **investigating** — Copilot is experiencing degraded availability. We are continuing to investigate.
- `2026-11-11 21:07` **monitoring** — The degradation affecting Actions, Copilot, Pages and Webhooks has been mitigated. We are monitoring to ensure stability.
- `2026-11-11 21:07` **monitoring** — We’re seeing recovery across all services and continuing to monitor. The backlog of Actions jobs are being processed. We’ll update when we see full recovery.
- `2026-11-11 21:21` **monitoring** — The backlog of Actions jobs has been fully processed. We’re observing full recovery across all impacted services.
- `2026-11-11 21:21` **resolved** — This incident has been resolved. Thank you for your patience and understanding as we addressed this issue. A detailed root cause analysis will be shared as soon as it is available.

### 33. github — Disruption with Email Delivery

- URL: https://stspg.io/lr8v4v8t7gs2
- Started: 2026-11-12 19:36 UTC
- Resolved: 2026-11-12 21:43 UTC
- Duration: 2h 7m
- Impact: minor, status: resolved
- Components: —
- Canonical incident: [ci-c3da26c11ab1eb6ad984db67](#ci-c3da26c11ab1eb6ad984db67)

- `2026-11-12 19:36` **investigating** — We are investigating reports of impacted performance for some GitHub services.
- `2026-11-12 19:43` **investigating** — We are seeing disruptions with email delivery for account sign-in and new account sign up. We are investigating.
- `2026-11-12 20:12` **investigating** — Users may be unable to sign-in, verify e-mail addresses, reset passwords, or receive e-mail notifications due to degradation with an upstream e-mail provider. We will continue to monitor and provide updates.
- `2026-11-12 20:19` **monitoring** — The degradation has been mitigated. We are monitoring to ensure stability.
- `2026-11-12 20:21` **monitoring** — We are seeing recovery for sign-in, email address verification, and other forms of e-mail communication previously impacted. We will continue to monitor.
- `2026-11-12 21:02` **monitoring** — We’re monitoring an email disruption related to an upstream provider degradation. Users whose email is hosted by the affected provider may experience delays in message delivery. Email service for users on other providers is currently operating normally. We’ll share updates as we learn more.
- `2026-11-12 21:43` **monitoring** — We’ve resolved the email-processing impact on our systems and are continuing to monitor. The upstream provider disruption may still cause delays for customers using that provider. Email delivery through other providers is operating normally.
- `2026-11-12 21:43` **resolved** — This incident has been resolved. Thank you for your patience and understanding as we addressed this issue. A detailed root cause analysis will be shared as soon as it is available.

### 34. github — Degraded availability for GPT-5.3 Codex and GPT-5.6 models in Copilot

- URL: https://stspg.io/t5gqmprtrdf2
- Started: 2026-11-09 23:16 UTC
- Resolved: 2026-11-10 00:05 UTC
- Duration: 49m
- Impact: major, status: resolved
- Components: Copilot, Copilot AI Model Providers
- Canonical incident: [ci-009f79df0e7feef1fc4151ef](#ci-009f79df0e7feef1fc4151ef)

- `2026-11-09 23:16` **investigating** — We are investigating reports of degraded availability for Copilot
- `2026-11-09 23:16` **investigating** — We are experiencing degraded availability for GPT-5.3 Codex, GPT-5.6 Luna, and GPT-5.6 Sol in Copilot products and IDE surfaces due to an issue with an upstream model provider. We recommend choosing another model or selecting Auto to continue using Copilot until the issue is resolved.
- `2026-11-09 23:22` **investigating** — Copilot AI Model Providers is experiencing degraded availability. We are continuing to investigate.
- `2026-11-09 23:22` **investigating** — Copilot is operating normally.
- `2026-11-09 23:46` **monitoring** — The degradation affecting Copilot AI Model Providers has been mitigated. We are monitoring to ensure stability.
- `2026-11-09 23:59` **monitoring** — The issue affecting GPT-5.3 Codex, GPT-5.6 Luna, and GPT-5.6 Sol in Copilot products and IDE surfaces has been mitigated. We are monitoring the service to ensure stability.
- `2026-11-10 00:05` **resolved** — This incident has been resolved. Thank you for your patience and understanding as we addressed this issue. A detailed root cause analysis will be shared as soon as it is available.

### 35. anthropic — Elevated purchasing errors on claude.ai

- URL: https://stspg.io/qqf00d9r7pm4
- Started: 2026-11-14 02:54 UTC
- Resolved: 2026-11-14 04:23 UTC
- Duration: 1h 28m
- Impact: major, status: resolved
- Components: claude.ai
- Canonical incident: [ci-36723f83e0c44d736b99ae6f](#ci-36723f83e0c44d736b99ae6f)

- `2026-11-14 02:54` **investigating** — We are investigating elevated errors affecting purchases, plan upgrades and cancellations on claude.ai. Core functionality is not affected. We will provide an update as soon as possible.
- `2026-11-14 03:09` **monitoring** — Purchases on claude.ai are working normally, and delayed plan changes and payments are now being applied. Some customers who recently paid may need to wait up to an hour for their plan to update. There is no need to pay again. We are monitoring for recovery.
- `2026-11-14 04:23` **resolved** — The issue affecting purchases and plan changes on claude.ai has been resolved, and delayed plan changes and payments have now been applied. Impact occurred from 18:34 PT on November 13 / 02:34 UTC on November 14 to 20:03 PT on November 13 / 04:03 UTC on November 14.

## Below threshold (23)

- `cloudflare` [Increased errors for R2 and Tiered Cache](https://www.cloudflarestatus.com/incidents/a643zcsfc71q) — 1h 15m
- `openai` [Sites: Publishing failing](https://status.openai.com/incidents/01M79GYA0GA772Y6D089AH5478) — 1h 56m
- `cloudflare` [Elevated number of errors in Bogotá](https://www.cloudflarestatus.com/incidents/sozoycnmog2m) — 1h 30m
- `cloudflare` [Elevated error rate for workers startup](https://www.cloudflarestatus.com/incidents/y03on2svrrxt) — 1h 22m
- `cloudflare` [Issues with 1.1.1.1 public resolver in Canada](https://www.cloudflarestatus.com/incidents/rikplz89rpml) — 1h 7m
- `openai` [Elevated errors with ChatGPT billing and Ads billing](https://status.openai.com/incidents/01M7C40W0N5MZHNW0YM09QV962) — 1h 16m
- `github` [Copilot experiencing degraded model availability for some GPT models](https://stspg.io/cppzsh0j19vy) — 1h 14m
- `openai` [Elevated ChatGPT Conversation Errors](https://status.openai.com/incidents/01M70RDA2GWD2WXE58ANY0WV5R) — 52m
- `cloudflare` [Account Analytics Filter Issue](https://www.cloudflarestatus.com/incidents/m644az5hr9xo) — 42m
- `cloudflare` [Workers deployments errors with D1 bindings](https://www.cloudflarestatus.com/incidents/3ptyvdaljq9b) — 21m
- `github` [Incident with Webhooks API](https://stspg.io/ns57pv7lmynl) — 50m
- `cloudflare` [Elevated error rates on R2 buckets in West North America](https://www.cloudflarestatus.com/incidents/dpao7idy3zei) — 39m
- `cloudflare` [Cloudflare Worker Builds Issue](https://www.cloudflarestatus.com/incidents/byzl30zyfc47) — 12m
- `openai` [Elevated Errors On API](https://status.openai.com/incidents/01M7193DWYRFWQM5W6E3VZ7TM7) — 42m
- `cloudflare` [DNSSEC Configuration Change Issues](https://www.cloudflarestatus.com/incidents/8odt36zmvm36) — 33m
- `cloudflare` [Prediction issue on OpenAI models via Replicate](https://www.cloudflarestatus.com/incidents/3en9cdyjn2sq) — 31m
- `openai` [GPT-image-2.5 and 2.6 Sunburst & Flare High error rate](https://status.openai.com/incidents/01M77FAXVFV338GY7G9VRENRS5) — 32m
- `cloudflare` [Network Performance Degraded](https://www.cloudflarestatus.com/incidents/804glbb3pq6t) — 25m
- `cloudflare` [Elevated number of 5xx errors in Hong Kong](https://www.cloudflarestatus.com/incidents/4qxszy04pdjw) — 15m
- `grafana` [Cloud Logs writes - Elevated latency and error rates in prod-eu-central-0](https://stspg.io/m2ykdxy6012r) — 17m
- `openai` [Elevated errors affecting logins, sign-ups, and ChatGPT features](https://status.openai.com/incidents/01M765MJQK7F85SC82RMTYN0V3) — 8m
- `github` [Incident with Webhooks](https://stspg.io/r0jrtx03yvll) — 4m
- `anthropic` [Elevated errors on Fable 52 and Mythos 52](https://stspg.io/pytn9v8p2zc7) — 0m

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-0a30f34c752fb49c0f4d211c"></a>

- Canonical ID: `ci-0a30f34c752fb49c0f4d211c` (1 source records).
  - [cloudflare/0ocqieaifg7h](../../data/2026/W46/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/0ocqieaifg7h
- Start time basis: `vendor_reported`.

<a id="ci-7d0c45c4220e872464de13e2"></a>

- Canonical ID: `ci-7d0c45c4220e872464de13e2` (1 source records).
  - [anthropic/ct4c2pwf2xpj](../../data/2026/W46/anthropic.json); vendor URL: https://stspg.io/4858wf40dx89
- Start time basis: `vendor_reported`.

<a id="ci-f788538c36e9a228e97f389c"></a>

- Canonical ID: `ci-f788538c36e9a228e97f389c` (1 source records).
  - [openai/01M7ATRS6BMZ61E4JDHJGM79Y0](../../data/2026/W46/openai.json); vendor URL: https://status.openai.com/incidents/01M7ATRS6BMZ61E4JDHJGM79Y0
- Start time basis: `vendor_reported`.

<a id="ci-3aedfad25e1c72a10aaae7ca"></a>

- Canonical ID: `ci-3aedfad25e1c72a10aaae7ca` (1 source records).
  - [anthropic/sqn3fg1l96dl](../../data/2026/W46/anthropic.json); vendor URL: https://stspg.io/m0rbx38nmy1c
- Start time basis: `vendor_reported`.

<a id="ci-63dc50e88eca748e5be37fd9"></a>

- Canonical ID: `ci-63dc50e88eca748e5be37fd9` (1 source records).
  - [cloudflare/yu63fhh2peb7](../../data/2026/W46/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/yu63fhh2peb7
- Start time basis: `vendor_reported`.

<a id="ci-44d9765457f6f0b90ddc1a0d"></a>

- Canonical ID: `ci-44d9765457f6f0b90ddc1a0d` (1 source records).
  - [openai/01M78GAT3AK9P9W8Z230KHFVE8](../../data/2026/W46/openai.json); vendor URL: https://status.openai.com/incidents/01M78GAT3AK9P9W8Z230KHFVE8
- Start time basis: `vendor_reported`.

<a id="ci-18358cb788c1961aed2943be"></a>

- Canonical ID: `ci-18358cb788c1961aed2943be` (1 source records).
  - [cloudflare/e7ktq32gphhe](../../data/2026/W46/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/e7ktq32gphhe
- Start time basis: `vendor_reported`.

<a id="ci-81ecb242f7e327a5c980a3ec"></a>

- Canonical ID: `ci-81ecb242f7e327a5c980a3ec` (1 source records).
  - [cloudflare/mc91s22odfvf](../../data/2026/W46/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/mc91s22odfvf
- Start time basis: `vendor_reported`.

<a id="ci-f943958e573df67fc42f3633"></a>

- Canonical ID: `ci-f943958e573df67fc42f3633` (1 source records).
  - [cloudflare/1x0tbxpmmttu](../../data/2026/W46/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/1x0tbxpmmttu
- Start time basis: `vendor_reported`.

<a id="ci-6319520b55ed6af07842a935"></a>

- Canonical ID: `ci-6319520b55ed6af07842a935` (1 source records).
  - [slack/1583](../../data/2026/W46/slack.json); vendor URL: https://slack-status.com/2026-11/b8bb993128ac94f5
- Start time basis: `vendor_reported`.

<a id="ci-de339f578e09069268438191"></a>

- Canonical ID: `ci-de339f578e09069268438191` (1 source records).
  - [grafana/f3cgznwspd0m](../../data/2026/W46/grafana.json); vendor URL: https://stspg.io/fvnhy2nn4xq6
- Start time basis: `vendor_reported`.

<a id="ci-f925f2043c9c282f805d95c5"></a>

- Canonical ID: `ci-f925f2043c9c282f805d95c5` (1 source records).
  - [anthropic/34wq7gxnxvvm](../../data/2026/W46/anthropic.json); vendor URL: https://stspg.io/264r547qtkky
- Start time basis: `vendor_reported`.

<a id="ci-9df1118b495b8ee1e137b867"></a>

- Canonical ID: `ci-9df1118b495b8ee1e137b867` (1 source records).
  - [openai/01M76CW2D7FM7K7F8NJ4W0G08N](../../data/2026/W46/openai.json); vendor URL: https://status.openai.com/incidents/01M76CW2D7FM7K7F8NJ4W0G08N
- Start time basis: `vendor_reported`.

<a id="ci-6d42be3a336cbfe96168081e"></a>

- Canonical ID: `ci-6d42be3a336cbfe96168081e` (1 source records).
  - [anthropic/nn11wk485zfn](../../data/2026/W46/anthropic.json); vendor URL: https://stspg.io/6kqyd0c5b1gr
- Start time basis: `vendor_reported`.

<a id="ci-7e816e7f0841048155f8126a"></a>

- Canonical ID: `ci-7e816e7f0841048155f8126a` (1 source records).
  - [aws/26f1425c833448f5](../../data/2026/W46/aws.json); vendor URL: https://health.aws.amazon.com/health/status
- Start time basis: `first_observed_update`.
- Warning: AWS onset is the first observed message, not a verified impact start

<a id="ci-b925591a04b02f7511d9a1fc"></a>

- Canonical ID: `ci-b925591a04b02f7511d9a1fc` (1 source records).
  - [cloudflare/9zxr2bkqoct7](../../data/2026/W46/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/9zxr2bkqoct7
- Start time basis: `vendor_reported`.

<a id="ci-1e0a441616183f3491a33a6a"></a>

- Canonical ID: `ci-1e0a441616183f3491a33a6a` (1 source records).
  - [cloudflare/a643zcsfc71q](../../data/2026/W46/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/a643zcsfc71q
- Start time basis: `vendor_reported`.

<a id="ci-51c613491673f28530db53c3"></a>

- Canonical ID: `ci-51c613491673f28530db53c3` (1 source records).
  - [cloudflare/88hc76jnq2gn](../../data/2026/W46/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/88hc76jnq2gn
- Start time basis: `vendor_reported`.

<a id="ci-74bfb35ba2e60e544c598c72"></a>

- Canonical ID: `ci-74bfb35ba2e60e544c598c72` (1 source records).
  - [grafana/x84fc9x32w3b](../../data/2026/W46/grafana.json); vendor URL: https://stspg.io/b9w7n973q0st
- Start time basis: `vendor_reported`.

<a id="ci-d84d19af354ef2d1dfbe25b2"></a>

- Canonical ID: `ci-d84d19af354ef2d1dfbe25b2` (1 source records).
  - [stripe/xls9d911cdv3](../../data/2026/W46/stripe.json); vendor URL: https://stspg.io/t1vjtvqqncxf
- Start time basis: `vendor_reported`.

<a id="ci-f2e12f4c1cfbbc05d541133a"></a>

- Canonical ID: `ci-f2e12f4c1cfbbc05d541133a` (1 source records).
  - [stripe/45253p0bm8dk](../../data/2026/W46/stripe.json); vendor URL: https://stspg.io/fgz0c6z26b43
- Start time basis: `vendor_reported`.

<a id="ci-f233a5eaca549a7cd81af490"></a>

- Canonical ID: `ci-f233a5eaca549a7cd81af490` (1 source records).
  - [github/xxb3l2bm76dr](../../data/2026/W46/github.json); vendor URL: https://stspg.io/4vdghm7tycm4
- Start time basis: `vendor_reported`.

<a id="ci-e77f3bfa3f702ce06c0d0386"></a>

- Canonical ID: `ci-e77f3bfa3f702ce06c0d0386` (1 source records).
  - [cloudflare/1na2mbfshgdh](../../data/2026/W46/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/1na2mbfshgdh
- Start time basis: `vendor_reported`.

<a id="ci-682c2f892744899d4df5e977"></a>

- Canonical ID: `ci-682c2f892744899d4df5e977` (1 source records).
  - [github/hnf21m9jkl3c](../../data/2026/W46/github.json); vendor URL: https://stspg.io/16nsnk2th3rv
- Start time basis: `vendor_reported`.

<a id="ci-5523009b2922e2f5a4c3e5f5"></a>

- Canonical ID: `ci-5523009b2922e2f5a4c3e5f5` (1 source records).
  - [github/c9px7cw2ghrs](../../data/2026/W46/github.json); vendor URL: https://stspg.io/jztczfplj6yb
- Start time basis: `vendor_reported`.

<a id="ci-d0f875d5ab8f69a2f190b653"></a>

- Canonical ID: `ci-d0f875d5ab8f69a2f190b653` (1 source records).
  - [aws/b48c38d5804e30f8](../../data/2026/W46/aws.json); vendor URL: https://health.aws.amazon.com/health/status
- Start time basis: `first_observed_update`.
- Warning: AWS onset is the first observed message, not a verified impact start

<a id="ci-65f3ac868ce78e42cb1b5267"></a>

- Canonical ID: `ci-65f3ac868ce78e42cb1b5267` (1 source records).
  - [cloudflare/21vwp89eiyjx](../../data/2026/W46/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/21vwp89eiyjx
- Start time basis: `vendor_reported`.

<a id="ci-84dd75fecfe0fa655d5a32e8"></a>

- Canonical ID: `ci-84dd75fecfe0fa655d5a32e8` (1 source records).
  - [grafana/6drrz4tkd35v](../../data/2026/W46/grafana.json); vendor URL: https://stspg.io/k3cgsqwglmzn
- Start time basis: `vendor_reported`.

<a id="ci-a65f27cd1c724c257a51f457"></a>

- Canonical ID: `ci-a65f27cd1c724c257a51f457` (1 source records).
  - [cloudflare/1h9c85fz6bm3](../../data/2026/W46/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/1h9c85fz6bm3
- Start time basis: `vendor_reported`.

<a id="ci-6fd6e5add514573bdc00bf7e"></a>

- Canonical ID: `ci-6fd6e5add514573bdc00bf7e` (1 source records).
  - [github/1xzmt89vv61l](../../data/2026/W46/github.json); vendor URL: https://stspg.io/qzj7qjdcr8jc
- Start time basis: `vendor_reported`.

<a id="ci-e83b177687e865bcd9a02619"></a>

- Canonical ID: `ci-e83b177687e865bcd9a02619` (1 source records).
  - [openai/01M71W79MAGEXEX93QTVTBYJ4K](../../data/2026/W46/openai.json); vendor URL: https://status.openai.com/incidents/01M71W79MAGEXEX93QTVTBYJ4K
- Start time basis: `vendor_reported`.

<a id="ci-059009c46813833b5e969eac"></a>

- Canonical ID: `ci-059009c46813833b5e969eac` (1 source records).
  - [github/tp8cghg7tlc4](../../data/2026/W46/github.json); vendor URL: https://stspg.io/7yw63xl9ztrv
- Start time basis: `vendor_reported`.

<a id="ci-f28b53b3942f32847c952628"></a>

- Canonical ID: `ci-f28b53b3942f32847c952628` (1 source records).
  - [github/dvy0t3mvh1lg](../../data/2026/W46/github.json); vendor URL: https://stspg.io/jdjlq5zjyl48
- Start time basis: `vendor_reported`.

<a id="ci-c3da26c11ab1eb6ad984db67"></a>

- Canonical ID: `ci-c3da26c11ab1eb6ad984db67` (1 source records).
  - [github/nwbjzl7nfrc2](../../data/2026/W46/github.json); vendor URL: https://stspg.io/lr8v4v8t7gs2
- Start time basis: `vendor_reported`.

<a id="ci-d18b752d79ac740ed30a728d"></a>

- Canonical ID: `ci-d18b752d79ac740ed30a728d` (1 source records).
  - [openai/01M79GYA0GA772Y6D089AH5478](../../data/2026/W46/openai.json); vendor URL: https://status.openai.com/incidents/01M79GYA0GA772Y6D089AH5478
- Start time basis: `vendor_reported`.

<a id="ci-231bb209f7ff855710b58aeb"></a>

- Canonical ID: `ci-231bb209f7ff855710b58aeb` (1 source records).
  - [cloudflare/sozoycnmog2m](../../data/2026/W46/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/sozoycnmog2m
- Start time basis: `vendor_reported`.

<a id="ci-4b793ef9d218b05b34c71b13"></a>

- Canonical ID: `ci-4b793ef9d218b05b34c71b13` (1 source records).
  - [cloudflare/y03on2svrrxt](../../data/2026/W46/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/y03on2svrrxt
- Start time basis: `vendor_reported`.

<a id="ci-009f79df0e7feef1fc4151ef"></a>

- Canonical ID: `ci-009f79df0e7feef1fc4151ef` (1 source records).
  - [github/6jm5k16vyhfg](../../data/2026/W46/github.json); vendor URL: https://stspg.io/t5gqmprtrdf2
- Start time basis: `vendor_reported`.

<a id="ci-36723f83e0c44d736b99ae6f"></a>

- Canonical ID: `ci-36723f83e0c44d736b99ae6f` (1 source records).
  - [anthropic/qrpz8k79qm7x](../../data/2026/W46/anthropic.json); vendor URL: https://stspg.io/qqf00d9r7pm4
- Start time basis: `vendor_reported`.

<a id="ci-545ee5ea1cf012eff0f5e52e"></a>

- Canonical ID: `ci-545ee5ea1cf012eff0f5e52e` (1 source records).
  - [cloudflare/rikplz89rpml](../../data/2026/W46/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/rikplz89rpml
- Start time basis: `vendor_reported`.

<a id="ci-27e0783b783ea4f120ce0751"></a>

- Canonical ID: `ci-27e0783b783ea4f120ce0751` (1 source records).
  - [openai/01M7C40W0N5MZHNW0YM09QV962](../../data/2026/W46/openai.json); vendor URL: https://status.openai.com/incidents/01M7C40W0N5MZHNW0YM09QV962
- Start time basis: `vendor_reported`.

<a id="ci-9dcba3a58b082bafeaf1c9f1"></a>

- Canonical ID: `ci-9dcba3a58b082bafeaf1c9f1` (1 source records).
  - [github/51drrqvsv66g](../../data/2026/W46/github.json); vendor URL: https://stspg.io/cppzsh0j19vy
- Start time basis: `vendor_reported`.

<a id="ci-085e9c0bd00f16d6f60f316b"></a>

- Canonical ID: `ci-085e9c0bd00f16d6f60f316b` (1 source records).
  - [openai/01M70RDA2GWD2WXE58ANY0WV5R](../../data/2026/W46/openai.json); vendor URL: https://status.openai.com/incidents/01M70RDA2GWD2WXE58ANY0WV5R
- Start time basis: `vendor_reported`.

<a id="ci-c782bb452430313d6746dcb6"></a>

- Canonical ID: `ci-c782bb452430313d6746dcb6` (1 source records).
  - [cloudflare/m644az5hr9xo](../../data/2026/W46/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/m644az5hr9xo
- Start time basis: `vendor_reported`.

<a id="ci-0bc57a5282dbed06d54c172d"></a>

- Canonical ID: `ci-0bc57a5282dbed06d54c172d` (1 source records).
  - [cloudflare/3ptyvdaljq9b](../../data/2026/W46/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/3ptyvdaljq9b
- Start time basis: `vendor_reported`.

<a id="ci-5d9a94f9a2f392888a273ef4"></a>

- Canonical ID: `ci-5d9a94f9a2f392888a273ef4` (1 source records).
  - [github/gl7yhd7jx00v](../../data/2026/W46/github.json); vendor URL: https://stspg.io/ns57pv7lmynl
- Start time basis: `vendor_reported`.

<a id="ci-4bd3f33f1a56ff7f5bf2a791"></a>

- Canonical ID: `ci-4bd3f33f1a56ff7f5bf2a791` (1 source records).
  - [cloudflare/dpao7idy3zei](../../data/2026/W46/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/dpao7idy3zei
- Start time basis: `vendor_reported`.

<a id="ci-a542698e8ef79b42ed30ac4b"></a>

- Canonical ID: `ci-a542698e8ef79b42ed30ac4b` (1 source records).
  - [cloudflare/byzl30zyfc47](../../data/2026/W46/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/byzl30zyfc47
- Start time basis: `vendor_reported`.

<a id="ci-05050ca5d82c16e99e346711"></a>

- Canonical ID: `ci-05050ca5d82c16e99e346711` (1 source records).
  - [openai/01M7193DWYRFWQM5W6E3VZ7TM7](../../data/2026/W46/openai.json); vendor URL: https://status.openai.com/incidents/01M7193DWYRFWQM5W6E3VZ7TM7
- Start time basis: `vendor_reported`.

<a id="ci-2eaa61e78bd57c2c35275c2e"></a>

- Canonical ID: `ci-2eaa61e78bd57c2c35275c2e` (1 source records).
  - [cloudflare/8odt36zmvm36](../../data/2026/W46/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/8odt36zmvm36
- Start time basis: `vendor_reported`.

<a id="ci-16c2528ba68d8729e392e2f0"></a>

- Canonical ID: `ci-16c2528ba68d8729e392e2f0` (1 source records).
  - [cloudflare/3en9cdyjn2sq](../../data/2026/W46/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/3en9cdyjn2sq
- Start time basis: `vendor_reported`.

<a id="ci-3c7ac93d1cdecb7452892719"></a>

- Canonical ID: `ci-3c7ac93d1cdecb7452892719` (1 source records).
  - [openai/01M77FAXVFV338GY7G9VRENRS5](../../data/2026/W46/openai.json); vendor URL: https://status.openai.com/incidents/01M77FAXVFV338GY7G9VRENRS5
- Start time basis: `vendor_reported`.

<a id="ci-a3fae2816fedff9457ec110b"></a>

- Canonical ID: `ci-a3fae2816fedff9457ec110b` (1 source records).
  - [cloudflare/804glbb3pq6t](../../data/2026/W46/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/804glbb3pq6t
- Start time basis: `vendor_reported`.

<a id="ci-b2ee6e4e118351868f92ce17"></a>

- Canonical ID: `ci-b2ee6e4e118351868f92ce17` (1 source records).
  - [cloudflare/4qxszy04pdjw](../../data/2026/W46/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/4qxszy04pdjw
- Start time basis: `vendor_reported`.

<a id="ci-d53a19f3286a26b44d7cad97"></a>

- Canonical ID: `ci-d53a19f3286a26b44d7cad97` (1 source records).
  - [grafana/d8rx6y4sf9k4](../../data/2026/W46/grafana.json); vendor URL: https://stspg.io/m2ykdxy6012r
- Start time basis: `vendor_reported`.

<a id="ci-976d023193357523d2275408"></a>

- Canonical ID: `ci-976d023193357523d2275408` (1 source records).
  - [openai/01M765MJQK7F85SC82RMTYN0V3](../../data/2026/W46/openai.json); vendor URL: https://status.openai.com/incidents/01M765MJQK7F85SC82RMTYN0V3
- Start time basis: `vendor_reported`.

<a id="ci-36a9da0e512d2aaeaa534958"></a>

- Canonical ID: `ci-36a9da0e512d2aaeaa534958` (1 source records).
  - [github/1h2wbly1y72t](../../data/2026/W46/github.json); vendor URL: https://stspg.io/r0jrtx03yvll
- Start time basis: `vendor_reported`.

<a id="ci-e71687e11453b3caf8c93475"></a>

- Canonical ID: `ci-e71687e11453b3caf8c93475` (1 source records).
  - [anthropic/pgz8r13pk4sm](../../data/2026/W46/anthropic.json); vendor URL: https://stspg.io/pytn9v8p2zc7
- Start time basis: `vendor_reported`.

## Sources

| Vendor | Status | Source records seen | Parse errors |
|--------|--------|----------------|--------------|
| anthropic | ok | 50 | 0 |
| atlassian | ok | 34 | 0 |
| aws | ok | 7 | 0 |
| circleci | ok | 50 | 0 |
| cloudflare | ok | 50 | 0 |
| datadog | FETCH FAILED: GET https://status.datadoghq.com/api/v2/incidents.json: status 403 | 0 | 0 |
| digitalocean | ok | 50 | 0 |
| gcp | ok | 8 | 0 |
| github | ok | 50 | 0 |
| grafana | ok | 50 | 0 |
| openai | ok | 25 | 0 |
| slack | ok | 26 | 0 |
| stripe | ok | 50 | 0 |
| yandexcloud | FETCH FAILED: GET https://status.cloud.yandex.ru/api/incidents?lang=en&page=4: status 403 | 0 | 0 |
