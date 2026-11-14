# Week 2026-W37

2026-09-07 to 2026-09-13 (UTC). Generated 2026-11-14 12:36.

46 canonical incidents, 25 above threshold. 2 of 14 sources unavailable: datadog, yandexcloud.

46 source records (49 saved revisions) → 46 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | cloudflare | [Hyperdrive Degraded Performance](https://www.cloudflarestatus.com/incidents/adw7e1heupa2) | 2026-09-11 00:00 | 24d 0h 0m | none | Hyperdrive | 62208 = 34560 × 1 × 1.8 |
| 2 | grafana | [Issues with Geomap Tiles](https://stspg.io/lr44wpz3m0tz) | 2026-09-08 18:05 | 4h 36m | minor | AWS Australia - prod-au-southeast-1, AWS Australia - prod-au-southeast-2, AWS Brazil - prod-sa-east-1, AWS Canada - prod-ca-east-0, AWS Germany - prod-eu-west-2, AWS Germany - prod-eu-west-4, AWS India - prod-ap-south-1, AWS Ireland - prod-eu-west-6, AWS Japan - prod-ap-northeast-0, AWS Singapore - prod-ap-southeast-1, AWS Sweden - prod-eu-north-0, AWS Switzerland - prod-eu-central-0, AWS UAE - prod-me-central-1, AWS UK - prod-gb-south-1, AWS US East (OH) - prod-us-east-0, AWS US East (OH) - prod-us-east-2, AWS US East (VA) - prod-us-east-3, AWS US West - prod-us-west-0, Azure Netherlands - prod-eu-west-3, Azure US Central - us-central2, Federal Cloud - AWS US Gov West, GCP Australia - prod-au-southeast-0, GCP Belgium - prod-eu-west-0, GCP Brazil - prod-sa-east-0, GCP India - prod-ap-south-0, GCP Saudi Arabia - prod-me-central-0, GCP Singapore - prod-ap-southeast-0, GCP UK - prod-gb-south-0, GCP US Central - prod-us-central-0, GCP US Central - prod-us-central-3, GCP US Central - prod-us-central-4, GCP US East - prod-us-east-1, play.grafana.org | 9108 = 276 × 33 × 1.0 |
| 3 | anthropic | [Degraded functionality for Claude Cowork on Windows](https://stspg.io/5yp3rhhztm30) | 2026-09-10 15:54 | 4d 3h 20m | major | Claude Cowork | 8940 = 5960 × 1 × 1.5 |
| 4 | cloudflare | [Workers Cron Triggers degraded](https://www.cloudflarestatus.com/incidents/sjs8s0q2x4hw) | 2026-09-09 19:17 | 2d 8h 16m | minor | Workers | 6077 = 3376 × 1 × 1.8 |
| 5 | cloudflare | [Cloudflare Tunnel Availability Issues](https://www.cloudflarestatus.com/incidents/k7l4hzmthl3k) | 2026-09-11 21:55 | 11h 1m | minor | Tunnel | 1190 = 661 × 1 × 1.8 |
| 6 | openai | [Unable to open shared ChatGPT Project using direct link](https://status.openai.com/incidents/01M24HYC9SX2Q4471EVMY6EXQE) | 2026-09-10 02:21 | 12h 6m | minor | — | 1089 = 726 × 1 × 1.5 |
| 7 | cloudflare | [Browser Isolation Session Initialization Failures](https://www.cloudflarestatus.com/incidents/z16209cfb1xv) | 2026-09-07 14:25 | 4h 49m | minor | Browser Isolation, Gateway | 1040 = 289 × 2 × 1.8 |
| 8 | cloudflare | [DNS Update Delays](https://www.cloudflarestatus.com/incidents/nhfcmzd1pv4p) | 2026-09-10 13:58 | 8h 24m | minor | DNS Updates | 907 = 504 × 1 × 1.8 |
| 9 | openai | [1% of ChatGPT Work (mobile/web) turns are failing for existing threads](https://status.openai.com/incidents/01M28MEQWTQJDRCRPFD9FWQ0H3) | 2026-09-11 16:22 | 9h 34m | minor | — | 861 = 574 × 1 × 1.5 |
| 10 | stripe | [Elevated Interac errors](https://stspg.io/j5s6qs75yrjz) | 2026-09-13 04:45 | 11h 4m | none | Acquirers and payment methods | 797 = 664 × 1 × 1.2 |
| 11 | openai | [Elevated errors for image generation](https://status.openai.com/incidents/01M20PYYYGRT9303VHAPA7YNT2) | 2026-09-08 14:32 | 7h 26m | minor | — | 669 = 446 × 1 × 1.5 |
| 12 | github | [Incident with several GitHub Services](https://stspg.io/8f3xch4y0v2k) | 2026-09-13 09:16 | 1h 28m | critical | API Requests, Actions, Issues, Pages, Pull Requests | 660 = 88 × 5 × 1.5 |
| 13 | slack | [Some customers experiencing blurry image previews and download issues](https://slack-status.com/2026-09/f171324a097abb9a) | 2026-09-12 11:04 | 8h 55m | unknown | Files | 642 = 535 × 1 × 1.2 |
| 14 | cloudflare | [Cloudflare Workers AI increased errors](https://www.cloudflarestatus.com/incidents/jsvvrtk5hvym) | 2026-09-09 14:25 | 5h 44m | minor | Workers AI | 619 = 344 × 1 × 1.8 |
| 15 | cloudflare | [Increased Workflow Instance Creation Errors](https://www.cloudflarestatus.com/incidents/np4n61ckjdw7) | 2026-09-07 05:12 | 2h 23m | minor | Seattle, WA, United States - (SEA), Workflows | 515 = 143 × 2 × 1.8 |
| 16 | openai | [File uploads are delayed or failing](https://status.openai.com/incidents/01M20QBACYZMK2PRCJQCBSWTYJ) | 2026-09-08 14:39 | 4h 36m | minor | — | 414 = 276 × 1 × 1.5 |
| 17 | cloudflare | [Issues creating and updating Cloud Connector Rules](https://www.cloudflarestatus.com/incidents/t4gby3dqzbg1) | 2026-09-12 08:06 | 2h 36m | minor | Cloud Connector | 281 = 156 × 1 × 1.8 |
| 18 | cloudflare | [Network Performance Issues with Taiwan datacenter (TPE)](https://www.cloudflarestatus.com/incidents/2njywrxj8qjf) | 2026-09-11 19:35 | 2h 28m | minor | Network | 266 = 148 × 1 × 1.8 |
| 19 | grafana | [High Latency in prod-us-east-2](https://stspg.io/f5m752d4l87f) | 2026-09-08 16:14 | 4h 19m | minor | AWS US East (OH) - prod-us-east-2 | 259 = 259 × 1 × 1.0 |
| 20 | openai | [Elevated errors for ChatGPT users in Europe](https://status.openai.com/incidents/01M27Q9GH7WVKNA61ZR3RDMSY5) | 2026-09-11 07:53 | 2h 23m | minor | — | 214 = 143 × 1 × 1.5 |
| 21 | yandexcloud | [Slowing down operations with access bindings](https://status.yandex.cloud/en/incidents/2005) | 2026-09-09 12:00 | 2h 22m | minor | Identity and Access Management | 213 = 142 × 1 × 1.5 |
| 22 | openai | [Elevated errors for GPT-5.6 Sol on the API](https://status.openai.com/incidents/01M27VK1VN54RQJGJEHB5JTQJJ) | 2026-09-11 09:08 | 2h 13m | minor | — | 200 = 133 × 1 × 1.5 |
| 23 | slack | [Trouble loading Workspace-level Apps & Workflows settings page is returning blank pages](https://slack-status.com/2026-09/c8ec111fea759b78) | 2026-09-10 00:34 | 2h 39m | unknown | Apps/Integrations/APIs | 191 = 159 × 1 × 1.2 |
| 24 | grafana | [Investigating elevated database load in AWS Germany](https://stspg.io/j1r7zw96mgh5) | 2026-09-09 11:27 | 2h 24m | minor | AWS Germany - prod-eu-west-4 | 144 = 144 × 1 × 1.0 |
| 25 | anthropic | [Elevated errors for Claude Mythos 5.1 and Claude Fable 5.1](https://stspg.io/n9n8kfcxtkfm) | 2026-09-11 14:19 | 17m | major | Claude API (api.anthropic.com), Claude Code, Claude Cowork, claude.ai | 102 = 17 × 4 × 1.5 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. cloudflare — Hyperdrive Degraded Performance

- URL: https://www.cloudflarestatus.com/incidents/adw7e1heupa2
- Started: 2026-09-11 00:00 UTC
- Resolved: 2026-10-05 00:00 UTC
- Duration: 24d 0h 0m
- Impact: none, status: resolved
- Components: Hyperdrive
- Canonical incident: [ci-74bd65ede6e8fc278492a918](#ci-74bd65ede6e8fc278492a918)

- `2026-10-07 23:56` **resolved** — Cloudflare identified an issue that affected Hyperdrive performance between 11th September and 5th October 2026. During this time, customers may have experienced elevated latency. This incident has been resolved.

### 2. grafana — Issues with Geomap Tiles

- URL: https://stspg.io/lr44wpz3m0tz
- Started: 2026-09-08 18:05 UTC
- Resolved: 2026-09-08 22:41 UTC
- Duration: 4h 36m
- Impact: minor, status: resolved
- Components: AWS Australia - prod-au-southeast-1, AWS Australia - prod-au-southeast-2, AWS Brazil - prod-sa-east-1, AWS Canada - prod-ca-east-0, AWS Germany - prod-eu-west-2, AWS Germany - prod-eu-west-4, AWS India - prod-ap-south-1, AWS Ireland - prod-eu-west-6, AWS Japan - prod-ap-northeast-0, AWS Singapore - prod-ap-southeast-1, AWS Sweden - prod-eu-north-0, AWS Switzerland - prod-eu-central-0, AWS UAE - prod-me-central-1, AWS UK - prod-gb-south-1, AWS US East (OH) - prod-us-east-0, AWS US East (OH) - prod-us-east-2, AWS US East (VA) - prod-us-east-3, AWS US West - prod-us-west-0, Azure Netherlands - prod-eu-west-3, Azure US Central - us-central2, Federal Cloud - AWS US Gov West, GCP Australia - prod-au-southeast-0, GCP Belgium - prod-eu-west-0, GCP Brazil - prod-sa-east-0, GCP India - prod-ap-south-0, GCP Saudi Arabia - prod-me-central-0, GCP Singapore - prod-ap-southeast-0, GCP UK - prod-gb-south-0, GCP US Central - prod-us-central-0, GCP US Central - prod-us-central-3, GCP US Central - prod-us-central-4, GCP US East - prod-us-east-1, play.grafana.org
- Canonical incident: [ci-8302dd689d389b0179980548](#ci-8302dd689d389b0179980548)

- `2026-09-08 18:05` **identified** — Some geomap tiles are displaying the “api key required". The issue has been identified, and our team is working on a fix. As a workaround, you can open the Geomap panel, navigate to the Map layers section, and change the Basemap layer from CARTO, or "Default base layer," to OpenStreetMap. The watermark should disappear immediately. Alternatively, if you would prefer to retain a dark-styled map similar to the current CARTO appearance, you can select MapLibre as the basemap and configure the URL as: https://tiles.openfreemap.org/styles/dark This option also renders without the watermark, does not require an API key, and provides a visual style similar to the dark CARTO basemap.
- `2026-09-08 22:41` **resolved** — This incident has been resolved.

### 3. anthropic — Degraded functionality for Claude Cowork on Windows

- URL: https://stspg.io/5yp3rhhztm30
- Started: 2026-09-10 15:54 UTC
- Resolved: 2026-09-14 19:14 UTC
- Duration: 4d 3h 20m
- Impact: major, status: resolved
- Components: Claude Cowork
- Canonical incident: [ci-43a9a49695dad394c58e93af](#ci-43a9a49695dad394c58e93af)

- `2026-09-10 15:54` **identified** — A Windows update released September 8 has left Claude Cowork on Windows unable to run local commands, because its workspace can no longer reach your computer's drive. For most users, chat and file reading and editing still work. There's no in-app workaround yet, and restarting or reinstalling won't help. Microsoft has developed a fix and is working to release it. We'll post an update when we have more to share.
- `2026-09-14 19:14` **resolved** — This issue has been resolved. Microsoft has released a September 14 Windows update that fixes it: https://support.microsoft.com/en-us/servicing/os/windows-11/2026/09/kb5129195-windows-11-24h2-25h2-security-update

### 4. cloudflare — Workers Cron Triggers degraded

- URL: https://www.cloudflarestatus.com/incidents/sjs8s0q2x4hw
- Started: 2026-09-09 19:17 UTC
- Resolved: 2026-09-12 03:34 UTC
- Duration: 2d 8h 16m
- Impact: minor, status: resolved
- Components: Workers
- Canonical incident: [ci-2ed8c7d7199fd963f73eaa14](#ci-2ed8c7d7199fd963f73eaa14)

- `2026-09-09 19:17` **investigating** — Workers Cron Triggers may not execute or may be delayed in executing. Updates to Workers Cron Triggers may take some time to take effect.
- `2026-09-09 19:18` **identified** — The issue has been identified and a fix is being implemented.
- `2026-09-12 02:25` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-09-12 03:34` **resolved** — This incident has been resolved.

### 5. cloudflare — Cloudflare Tunnel Availability Issues

- URL: https://www.cloudflarestatus.com/incidents/k7l4hzmthl3k
- Started: 2026-09-11 21:55 UTC
- Resolved: 2026-09-12 08:57 UTC
- Duration: 11h 1m
- Impact: minor, status: resolved
- Components: Tunnel
- Canonical incident: [ci-04ef0ea10c5ad910cb8e2ee1](#ci-04ef0ea10c5ad910cb8e2ee1)

- `2026-09-11 21:55` **investigating** — Cloudflare is investigating issues with Cloudflare Tunnel. More updates to follow shortly.
- `2026-09-11 23:15` **identified** — The issue has been identified and a fix is being implemented.
- `2026-09-12 03:49` **identified** — Mitigation is reducing the affected traffic, and recovery is progressing. Some existing Cloudflare Tunnel connections may continue to experience intermittent issues while this work completes. We continue to monitor progress.
- `2026-09-12 08:42` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-09-12 08:57` **resolved** — This incident has been resolved.

### 6. openai — Unable to open shared ChatGPT Project using direct link

- URL: https://status.openai.com/incidents/01M24HYC9SX2Q4471EVMY6EXQE
- Started: 2026-09-10 02:21 UTC
- Resolved: 2026-09-10 14:28 UTC
- Duration: 12h 6m
- Impact: minor, status: resolved
- Components: —
- Canonical incident: [ci-32bae5afda3ee87d4028bf97](#ci-32bae5afda3ee87d4028bf97)

- `2026-09-10 02:21` **investigating** — We are investigating the issue for the listed services.
- `2026-09-10 06:01` **identified** — We are still implementing the mitigation.
- `2026-09-10 09:37` **identified** — We are still implementing the mitigation.
- `2026-09-10 12:06` **monitoring** — We have applied the mitigation and are monitoring the recovery.
- `2026-09-10 14:28` **resolved** — All impacted services have now fully recovered.

### 7. cloudflare — Browser Isolation Session Initialization Failures

- URL: https://www.cloudflarestatus.com/incidents/z16209cfb1xv
- Started: 2026-09-07 14:25 UTC
- Resolved: 2026-09-07 19:14 UTC
- Duration: 4h 49m
- Impact: minor, status: resolved
- Components: Browser Isolation, Gateway
- Canonical incident: [ci-15e56813f70800815118200d](#ci-15e56813f70800815118200d)

- `2026-09-07 14:25` **investigating** — We are investigating reports of customers encountering a white screen when attempting to access websites behind Browser Isolation. This is caused by Browser Isolation session initialization failing to complete. Our engineering team is actively working to identify and resolve the root cause.
- `2026-09-07 14:59` **identified** — The issue has been identified and a fix is being implemented.
- `2026-09-07 15:11` **monitoring** — A fix has been implemented and we are monitoring the results. Customers still experiencing the issue may attempt to close all their Browser Isolation tabs for 2 minutes before starting a new session.
- `2026-09-07 19:14` **resolved** — This incident has been resolved.

### 8. cloudflare — DNS Update Delays

- URL: https://www.cloudflarestatus.com/incidents/nhfcmzd1pv4p
- Started: 2026-09-10 13:58 UTC
- Resolved: 2026-09-10 22:22 UTC
- Duration: 8h 24m
- Impact: minor, status: resolved
- Components: DNS Updates
- Canonical incident: [ci-a43b656dcadcf900e2ff9834](#ci-a43b656dcadcf900e2ff9834)

- `2026-09-10 13:58` **investigating** — Cloudflare is investigating delays in propagation of updates to DNS records. More updates to follow shortly.
- `2026-09-10 14:17` **identified** — The issue has been identified and a fix is being implemented.
- `2026-09-10 22:16` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-09-10 22:22` **resolved** — This incident has been resolved.

### 9. openai — 1% of ChatGPT Work (mobile/web) turns are failing for existing threads

- URL: https://status.openai.com/incidents/01M28MEQWTQJDRCRPFD9FWQ0H3
- Started: 2026-09-11 16:22 UTC
- Resolved: 2026-09-12 01:57 UTC
- Duration: 9h 34m
- Impact: minor, status: resolved
- Components: —
- Canonical incident: [ci-a69b1ce8c5f0ff9750fd3368](#ci-a69b1ce8c5f0ff9750fd3368)

- `2026-09-11 16:22` **investigating** — We have identified that users are experiencing elevated errors for the impacted services. We are working on implementing a mitigation.
- `2026-09-11 17:16` **identified** — We have identified that users are experiencing elevated errors for the impacted services. We are working on implementing a mitigation.
- `2026-09-11 17:17` **investigating** — We are investigating the issue for the listed services.
- `2026-09-11 21:51` **monitoring** — We have applied the mitigation and are monitoring the recovery.
- `2026-09-12 01:38` **monitoring** — We have applied the mitigation and are monitoring the recovery.
- `2026-09-12 01:57` **resolved** — All impacted services have now fully recovered.

### 10. stripe — Elevated Interac errors

- URL: https://stspg.io/j5s6qs75yrjz
- Started: 2026-09-13 04:45 UTC
- Resolved: 2026-09-13 15:49 UTC
- Duration: 11h 4m
- Impact: none, status: resolved
- Components: Acquirers and payment methods
- Canonical incident: [ci-8070d9450f8b693463537ceb](#ci-8070d9450f8b693463537ceb)

- `2026-09-13 04:45` **investigating** — Since 03:46 UTC, we’ve observed elevated error rates and increased response times with Interac payments, affecting in-person transactions in Canada. We have reached out to Interac and will share updates as soon as they are available.
- `2026-09-13 05:17` **monitoring** — We have observed recovery since 04:50 UTC but have been advised by Interac that ongoing maintenance will occur until 15:00 UTC. During this period, there may be intermittent errors and increased response times with Interac payments. We are monitoring closely and will share updates as soon as they are available.
- `2026-09-13 15:49` **resolved** — The maintenance at Interac has concluded successfully. We’ve observed reduced error rates and improved response times with Interac payments in Canada. The issue is now resolved.

### 11. openai — Elevated errors for image generation

- URL: https://status.openai.com/incidents/01M20PYYYGRT9303VHAPA7YNT2
- Started: 2026-09-08 14:32 UTC
- Resolved: 2026-09-08 21:59 UTC
- Duration: 7h 26m
- Impact: minor, status: resolved
- Components: —
- Canonical incident: [ci-70138b680b5e92b13729def9](#ci-70138b680b5e92b13729def9)

- `2026-09-08 14:32` **investigating** — We’re investigating elevated errors affecting image generation in ChatGPT and the Images API. Some requests may fail. We’ll share another update as we learn more.
- `2026-09-08 15:05` **investigating** — We’re continuing to investigate elevated errors affecting image generation in ChatGPT and the Images API. We’ve made changes aimed at reducing failures and are assessing their effect. Some requests may still fail. We’ll share another update as we learn more.
- `2026-09-08 15:38` **investigating** — > We’re continuing to investigate elevated errors affecting image generation in ChatGPT and the Images API. We saw an increase in failed Images API requests after our last update. We’re applying mitigations and assessing their effect. Some image requests may still fail. We’ll share another update as we learn more.
- `2026-09-08 16:35` **identified** — We have identified the issue affecting image generation in ChatGPT and the Images API. Some image-generation requests may continue to fail.
- `2026-09-08 19:11` **monitoring** — We have applied the mitigation and are monitoring the recovery.
- `2026-09-08 21:59` **resolved** — All impacted services have now fully recovered.

### 12. github — Incident with several GitHub Services

- URL: https://stspg.io/8f3xch4y0v2k
- Started: 2026-09-13 09:16 UTC
- Resolved: 2026-09-13 10:44 UTC
- Duration: 1h 28m
- Impact: critical, status: resolved
- Components: API Requests, Actions, Issues, Pages, Pull Requests
- Canonical incident: [ci-b76104eccbb3efb2c5838b24](#ci-b76104eccbb3efb2c5838b24)

- `2026-09-13 09:16` **investigating** — We are investigating reports of degraded availability for API Requests, Issues, Pages and Pull Requests
- `2026-09-13 09:25` **investigating** — Actions is experiencing degraded performance. We are continuing to investigate.
- `2026-09-13 09:36` **investigating** — We're seeing increased database replication delays on collab which is causing increased error rates in authorization endpoints and follow-on increased error rates across the system - we are investigating
- `2026-09-13 10:26` **investigating** — We have reduced load on this cluster with internal load-shedding and are seeing signs of recovery but continue to monitor
- `2026-09-13 10:28` **investigating** — Pull Requests is experiencing degraded performance. We are continuing to investigate.
- `2026-09-13 10:44` **resolved** — On September 13, 2026, between 08:43 and 10:44 UTC, GitHub experienced degraded availability across approximately 28 services, including Issues, Pull Requests, Actions, Codespaces, Pages, Notifications, Code Scanning, Git LFS, and new account signup. At peak, 8.8% of requests to create GitHub App installation access tokens failed. Token issuance for Actions workflows was also affected, impacting approximately 4% of workflows during the incident time frame. Creating issues through the web interface failed for about 96% of attempts, and signup failures were above 90%. <br /> <br />The cause was an internal data-cleanup job that began writing to a shared database cluster at 07:33 UTC. That cluster stores permission data read on nearly every authenticated request. The safeguard that was pacing the background job watched only one health signal — how far the database replicas were lagging — and that signal stayed low the whole time. It did not account for the load building on the primary itself, so the job kept writing while the primary quietly ran toward its limit. <br /><br />When the primary ran out of available connections, requests that needed it could not complete. First, there was no quick timeout on these database calls, so request handlers waited on the stalled database instead of failing fast, and the shared request-handling capacity degraded into site-wide errors. Second, a retry loop around token creation kept re-sending the writes that were already failing, which held the database saturated rather than letting it recover. <br /><br />Monitoring declared the incident at 08:50 UTC, but due to the broad impact and amplification from token creation, it took time to identify the source of the load. First responders mitigated by shedding internal load and pausing the job, and all services recovered by 10:44 UTC. <br /><br />To prevent recurrence, we are rate-limiting background jobs against shared, customer-serving databases by default, and adding automatic pausing and paging on primary-server load rather than replication lag alone. We are also surfacing running background work directly alongside database health signals so responders can see and pause it without leaving those dashboards, bounding retries in the token-issuing path, and adding request-level timeouts so one unhealthy database cannot consume shared web server capacity. In addition, we are breaking apart this database cluster to remove the single point of failure. We will be moving various service-specific data, including the authorization data, out of this shared cluster in the next two weeks.

### 13. slack — Some customers experiencing blurry image previews and download issues

- URL: https://slack-status.com/2026-09/f171324a097abb9a
- Started: 2026-09-12 11:04 UTC
- Resolved: 2026-09-12 20:00 UTC
- Duration: 8h 55m
- Impact: unknown, status: resolved
- Components: Files
- Canonical incident: [ci-80ac55f198b839ccd2197bc6](#ci-80ac55f198b839ccd2197bc6)

- `2026-09-12 11:04` **unknown** — We're aware there's an issue where a subset of customers are experiencing blurry images which can't be downloaded, this is happening across all clients. We're actively working to understand the cause, scope, fix, or other details.
- `2026-09-12 11:29` **unknown** — We've identified a likely technical trigger connected to a recent image-rendering service change. Current status and actions being taken: We have identified a likely technical trigger — a recent thumbnail-sharpening change to the image-rendering service. We're working to validate this, and are evaluating the rollback of the change as a possible fix. Estimated time to resolution: We are still confirming the technical trigger before a fix can be scoped. Scope: Affects a subset of customers across desktop, web, and mobile clients Current impact to end users: Images uploaded before roughly 9:30 PM GMT on Sep 11 appear blurry in previews and fail to download. Newly uploaded images are unaffected. Known workarounds: None identified yet. We appreciate your patience as we sort this out.
- `2026-09-12 12:52` **unknown** — We have rolled back the code change that caused the issue in the production environment. Validation is in progress. Current status and actions being taken: We confirmed the technical trigger (a thumbnail-sharpening code change) and completed a rollback in the production environment after successfully testing the fix in the test environment. We are performing additional validations to confirm the fix was successful. Estimated time to resolution: We do not have and estimated time frame at the moment,. Scope: Affects a subset of customers across desktop, web, and mobile clients; Current impact to end users: Images uploaded before roughly 9:30 PM GMT on Sep 11 appear blurry in previews and fail to download. Newly uploaded images are unaffected. We appreciate your patience as we sort this out.
- `2026-09-12 14:10` **unknown** — The earlier rollback did not resolve the issue, and investigation has been redirected to a file-serving problem rather than the thumbnail-generation change. Current status and actions being taken: The rollback of the thumbnail-sharpening change did not fix the issue. We have revised the technical trigger assessment to a file-serving problem, where the system returns near-empty successful responses instead of full image data. We are investigating file storage and access-policy changes, and gathering file details from affected customers to pinpoint the cause. Estimated time to resolution: No confirmed estimate yet; the team is still narrowing down the cause following the pivot to file serving. Scope: Affects a subset of customers across desktop, web, and mobile clients. Current impact to end users: Images uploaded before 9:30 PM GMT on Sep 11 appear blurry in previews and fail to download. Newly uploaded images are unaffected. The earlier rollback did not change this — previously affected images remained blurry and undownloadable. We appreciate your patience as we sort this out.
- `2026-09-12 15:24` **unknown** — We have confirmed the technical trigger and are preparing a fix Current status and actions being taken: We confirmed the technical trigger — a code change that caused older file records to lose their stored file names, breaking retrieval of existing images. The team is preparing a targeted fix to restore file names correctly when reading older files. The affected files are expected to recover automatically once the fix is deployed. Estimated time to resolution: We don't have an estimate yet. Scope: Affects a subset of customers across desktop, web, and mobile clients. Current impact to end users: Images uploaded before roughly 9:30 PM GMT on Sep 11 appear blurry in previews and fail to download. Newly uploaded images are unaffected. Previously affected images remain blurry and undownloadable. We appreciate your patience as we sort this out.
- `2026-09-12 20:00` **unknown** — Some customers experienced blurry images that couldn't be downloaded across all clients, affecting files uploaded before a recent file storage update. We deployed a fix restoring proper file access, and confirmed it resolved the issue for affected customers. We apologize for how this incident affected you and your business. We will undertake a full investigation of the incident, establishing the technical trigger, the underlying cause, and preventive action to avoid a repeat in the future.

### 14. cloudflare — Cloudflare Workers AI increased errors

- URL: https://www.cloudflarestatus.com/incidents/jsvvrtk5hvym
- Started: 2026-09-09 14:25 UTC
- Resolved: 2026-09-09 20:09 UTC
- Duration: 5h 44m
- Impact: minor, status: resolved
- Components: Workers AI
- Canonical incident: [ci-cc41854833958870b7fce062](#ci-cc41854833958870b7fce062)

- `2026-09-09 14:25` **investigating** — Cloudflare is aware of, and investigating, increased Workers AI errors when attempting to use GLM 5.3. Updates to follow.
- `2026-09-09 15:02` **identified** — The issue has been identified and a fix is being implemented.
- `2026-09-09 20:09` **resolved** — This incident has been resolved.

### 15. cloudflare — Increased Workflow Instance Creation Errors

- URL: https://www.cloudflarestatus.com/incidents/np4n61ckjdw7
- Started: 2026-09-07 05:12 UTC
- Resolved: 2026-09-07 07:35 UTC
- Duration: 2h 23m
- Impact: minor, status: resolved
- Components: Seattle, WA, United States - (SEA), Workflows
- Canonical incident: [ci-fa3fea16abc442a67734ed99](#ci-fa3fea16abc442a67734ed99)

- `2026-09-07 05:12` **investigating** — Cloudflare is aware of, and investigating an issue which potentially impacts multiple customers. Further detail will be provided as more information becomes available.
- `2026-09-07 07:03` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-09-07 07:35` **resolved** — This incident has been resolved.

### 16. openai — File uploads are delayed or failing

- URL: https://status.openai.com/incidents/01M20QBACYZMK2PRCJQCBSWTYJ
- Started: 2026-09-08 14:39 UTC
- Resolved: 2026-09-08 19:16 UTC
- Duration: 4h 36m
- Impact: minor, status: resolved
- Components: —
- Canonical incident: [ci-44680b434e693b65054736fd](#ci-44680b434e693b65054736fd)

- `2026-09-08 14:39` **investigating** — Some file uploads are failing to process. Successful uploads may take longer than usual to become available. We’re investigating and will share another update as we learn more.
- `2026-09-08 14:47` **investigating** — We are investigating the issue for the listed services.
- `2026-09-08 15:37` **investigating** — We’re continuing to investigate file-upload processing failures and delays in ChatGPT. Some files may appear uploaded but remain unavailable while processing; other uploads may fail. We have not confirmed recovery and will share another update as we learn more.
- `2026-09-08 16:29` **identified** — We have identified the issue causing file-upload processing failures and delays in ChatGPT. Some files may remain unavailable while processing, and some uploads may fail.
- `2026-09-08 17:48` **monitoring** — We have applied the mitigation and are monitoring the recovery.
- `2026-09-08 19:16` **resolved** — All impacted services have now fully recovered.

### 17. cloudflare — Issues creating and updating Cloud Connector Rules

- URL: https://www.cloudflarestatus.com/incidents/t4gby3dqzbg1
- Started: 2026-09-12 08:06 UTC
- Resolved: 2026-09-12 10:42 UTC
- Duration: 2h 36m
- Impact: minor, status: resolved
- Components: Cloud Connector
- Canonical incident: [ci-7ae694cca04d1da5510133da](#ci-7ae694cca04d1da5510133da)

- `2026-09-12 08:06` **investigating** — Cloudflare is investigating failures in updating and creating Cloud Connector Rules. We are working to mitigate this problem. More updates to follow shortly.
- `2026-09-12 08:11` **identified** — The issue has been identified and a fix is being implemented.
- `2026-09-12 10:10` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-09-12 10:42` **resolved** — This incident has been resolved.

### 18. cloudflare — Network Performance Issues with Taiwan datacenter (TPE)

- URL: https://www.cloudflarestatus.com/incidents/2njywrxj8qjf
- Started: 2026-09-11 19:35 UTC
- Resolved: 2026-09-11 22:03 UTC
- Duration: 2h 28m
- Impact: minor, status: resolved
- Components: Network
- Canonical incident: [ci-ece347690603aa8f7ebdb2d4](#ci-ece347690603aa8f7ebdb2d4)

- `2026-09-11 19:35` **investigating** — Cloudflare is investigating issues with Network Performance in Taiwan datacenter (TPE) We are working to analyze and mitigate this problem. More updates to follow shortly.
- `2026-09-11 22:03` **resolved** — This incident has been resolved.

### 19. grafana — High Latency in prod-us-east-2

- URL: https://stspg.io/f5m752d4l87f
- Started: 2026-09-08 16:14 UTC
- Resolved: 2026-09-08 20:34 UTC
- Duration: 4h 19m
- Impact: minor, status: resolved
- Components: AWS US East (OH) - prod-us-east-2
- Canonical incident: [ci-5e2264e6ecacb808cb82d4aa](#ci-5e2264e6ecacb808cb82d4aa)

- `2026-09-08 16:14` **investigating** — We are currently investigating an issue causing high latency for deployments in the prod-us-east-2 cluster.
- `2026-09-08 17:04` **monitoring** — We have observed latency return to normal levels, and we will continue to monitor this incident.
- `2026-09-08 20:34` **resolved** — This incident has been resolved.

### 20. openai — Elevated errors for ChatGPT users in Europe

- URL: https://status.openai.com/incidents/01M27Q9GH7WVKNA61ZR3RDMSY5
- Started: 2026-09-11 07:53 UTC
- Resolved: 2026-09-11 10:16 UTC
- Duration: 2h 23m
- Impact: minor, status: resolved
- Components: —
- Canonical incident: [ci-2fa2291e2b955996aa7c565a](#ci-2fa2291e2b955996aa7c565a)

- `2026-09-11 07:53` **identified** — We have identified that users are experiencing elevated errors for the impacted services. We are working on implementing a mitigation.
- `2026-09-11 08:37` **monitoring** — We have applied the mitigation and are monitoring the recovery.
- `2026-09-11 10:16` **resolved** — All impacted services have now fully recovered.

### 21. yandexcloud — Slowing down operations with access bindings

- URL: https://status.yandex.cloud/en/incidents/2005
- Started: 2026-09-09 12:00 UTC
- Resolved: 2026-09-09 14:22 UTC
- Duration: 2h 22m
- Impact: minor, status: resolved
- Components: Identity and Access Management
- Canonical incident: [ci-041ec95f4fc9cc6d67d57f8f](#ci-041ec95f4fc9cc6d67d57f8f)

- `2026-09-09 14:28` **investigating** — We have noticed a slowdown in operations with access bindings in the Cloud. The developers are already aware of it and are fixing it.
- `2026-09-09 14:39` **resolved** — At 17:22 Moscow time, the problem was fixed. We are investigating the causes of its occurrence.

### 22. openai — Elevated errors for GPT-5.6 Sol on the API

- URL: https://status.openai.com/incidents/01M27VK1VN54RQJGJEHB5JTQJJ
- Started: 2026-09-11 09:08 UTC
- Resolved: 2026-09-11 11:21 UTC
- Duration: 2h 13m
- Impact: minor, status: resolved
- Components: —
- Canonical incident: [ci-f025dc68ef6163f688fbc921](#ci-f025dc68ef6163f688fbc921)

- `2026-09-11 09:08` **identified** — We have identified that users are experiencing elevated errors for the impacted services. We are working on implementing a mitigation.
- `2026-09-11 10:09` **monitoring** — We have brought additional capacity online and are seeing improvements for some GPT-5.6 Sol API traffic. We continue to mitigate elevated errors affecting Chat Completions and Responses and are monitoring recovery
- `2026-09-11 11:21` **resolved** — All impacted services have now fully recovered.

### 23. slack — Trouble loading Workspace-level Apps & Workflows settings page is returning blank pages

- URL: https://slack-status.com/2026-09/c8ec111fea759b78
- Started: 2026-09-10 00:34 UTC
- Resolved: 2026-09-10 03:14 UTC
- Duration: 2h 39m
- Impact: unknown, status: resolved
- Components: Apps/Integrations/APIs
- Canonical incident: [ci-33521a347717a88a6c06e68b](#ci-33521a347717a88a6c06e68b)

- `2026-09-10 00:34` **unknown** — We’re aware of an issue affecting Workspace-level Apps & Workflows settings page. Users are seeing a blank page affecting these admin features. We’re actively investigating this issue and will provide an update in 30 minutes or sooner if more information becomes available.
- `2026-09-10 00:58` **unknown** — We've engaged our engineering team to assist in the investigation. We’re continuing to work on understanding the issue. We’ll provide an update in 30 minutes or sooner if more information becomes available.
- `2026-09-10 01:29` **unknown** — We've identified the cause of this issue, and our engineers are testing a fix. We'll provide an update in 30 minutes or sooner if additional information becomes available.
- `2026-09-10 01:37` **unknown** — The fix has been validated in our test environment and is now being deployed. We will provide an update once the deployment is complete, or sooner if additional information becomes available.
- `2026-09-10 03:14` **unknown** — This incident has been resolved. The engineering team has confirmed that the fix has been successfully deployed. We apologize for any disruptions this may have caused to your day.

### 24. grafana — Investigating elevated database load in AWS Germany

- URL: https://stspg.io/j1r7zw96mgh5
- Started: 2026-09-09 11:27 UTC
- Resolved: 2026-09-09 13:51 UTC
- Duration: 2h 24m
- Impact: minor, status: resolved
- Components: AWS Germany - prod-eu-west-4
- Canonical incident: [ci-d1d0d7778d4922fce64ad009](#ci-d1d0d7778d4922fce64ad009)

- `2026-09-09 11:27` **investigating** — We are investigating elevated database load affecting infrastructure supporting a subset of Grafana Cloud instances in the AWS Germany region (prod-eu-west-4). Affected users may experience slow responses when accessing Grafana. We will provide an update as more information becomes available.
- `2026-09-09 12:44` **monitoring** — A fix has been implemented and we have been seeing recovery to the affected region. We will continue to monitor.
- `2026-09-09 13:51` **resolved** — This incident has been resolved.

### 25. anthropic — Elevated errors for Claude Mythos 5.1 and Claude Fable 5.1

- URL: https://stspg.io/n9n8kfcxtkfm
- Started: 2026-09-11 14:19 UTC
- Resolved: 2026-09-11 14:37 UTC
- Duration: 17m
- Impact: major, status: resolved
- Components: Claude API (api.anthropic.com), Claude Code, Claude Cowork, claude.ai
- Canonical incident: [ci-e93379c888820ffa838fd3c9](#ci-e93379c888820ffa838fd3c9)

- `2026-09-11 14:19` **investigating** — We are investigating elevated errors on requests to Claude Fable 5.1 and Claude Mythos 5.1. We will provide an update as soon as possible.
- `2026-09-11 14:21` **identified** — We have identified the cause of elevated errors on requests to Claude Mythos 5.1 and Claude Fable 5.1 and are working on a fix. We will provide an update as soon as possible.
- `2026-09-11 14:28` **monitoring** — A fix has been deployed for Claude Mythos 5.1 and Claude Fable 5.1 and we are monitoring for recovery.
- `2026-09-11 14:28` **monitoring** — A fix has been deployed for the issue affecting Claude services and we are monitoring for recovery.
- `2026-09-11 14:37` **resolved** — The issue affecting Claude services has been resolved. Impact occurred from 13:57 to 14:20 UTC

## Below threshold (21)

- `cloudflare` [Cache Reserve Issue in Paris](https://www.cloudflarestatus.com/incidents/qz292kvy220g) — 1h 56m
- `anthropic` [Elevated latency on the Claude API for some users](https://stspg.io/hfs7749ykl14) — 1h 40m
- `cloudflare` [SSL Certificate Provisioning Delays for Let's Encrypt](https://www.cloudflarestatus.com/incidents/3r3qcgbb0714) — 1h 44m
- `cloudflare` [Issues rendering route configuration table in Networking > Routes UI](https://www.cloudflarestatus.com/incidents/kjfs3vqkxkgs) — 1h 40m
- `openai` [Delays in customer support responses](https://status.openai.com/incidents/01M26E3PPAYJMMRHX2MZJCTHMF) — 1h 55m
- `openai` [Elevated errors affecting ChatGPT Work](https://status.openai.com/incidents/01M26DX2PP1AFHGJK5E4M6S2J6) — 1h 40m
- `cloudflare` [Issues Editing Bulk Redirects](https://www.cloudflarestatus.com/incidents/92j1xmlkscw3) — 1h 15m
- `cloudflare` [Cloudflare Account Member Invite Issue](https://www.cloudflarestatus.com/incidents/pjr4m9q1xxb1) — 1h 7m
- `openai` [Increased Error Rate For Pro and Plus Plan Conversations](https://status.openai.com/incidents/01M23BNX5V1XT9P3NP0PX56FTN) — 1h 8m
- `cloudflare` [Performance degradation affecting TURN service](https://www.cloudflarestatus.com/incidents/pgn7y7hgd25l) — 49m
- `cloudflare` [Intermittent issues accessing the Dashboard on Firefox and Safari](https://www.cloudflarestatus.com/incidents/hy0gsndpkypz) — 46m
- `stripe` [Elevated Interac errors](https://stspg.io/7wcb364xld9l) — 34m
- `digitalocean` [Power Degradation impacting GPU Performance in ATL1](https://stspg.io/g4s2358yl9tt) — 1h 38m
- `openai` [Elevated errors for ChatGPT users in Europe](https://status.openai.com/incidents/01M2DV6788RRSDZ5K98S5HN0A8) — 51m
- `openai` [Codex GitHub Review and Pull Request Failures](https://status.openai.com/incidents/01M2D1CZG0862JB4JHVW56QF89) — 27m
- `openai` [Investigating unexpected usage limit resets](https://status.openai.com/incidents/01M23KG62KKK448RN434CQ64Z8) — 24m
- `circleci` [Issues loading pipelines for a small number of projects](https://stspg.io/l0tsrtv1tkbl) — 1m
- `cloudflare` [Errors in Unified Billing for some requests OpenAI models](https://www.cloudflarestatus.com/incidents/kq11rg5y66f2) — 0m
- `cloudflare` [Increased HTTP Errors](https://www.cloudflarestatus.com/incidents/dmc0gzyv8gcs) — 0m
- `cloudflare` [Issues with Workers VPC hostname route resolution on 2026-09-11](https://www.cloudflarestatus.com/incidents/9mcdtk1lphgl) — 0m
- `grafana` [Failures Submitting Self-Serve Configuration Changes for Grafana Cloud Logs in AWS Sweden (prod-eu-north-0)](https://stspg.io/gcg65gc0vtw8) — 0m

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-74bd65ede6e8fc278492a918"></a>

- Canonical ID: `ci-74bd65ede6e8fc278492a918` (1 source records).
  - [cloudflare/adw7e1heupa2](../../data/2026/W37/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/adw7e1heupa2
- Start time basis: `vendor_reported`.

<a id="ci-8302dd689d389b0179980548"></a>

- Canonical ID: `ci-8302dd689d389b0179980548` (1 source records).
  - [grafana/nm6vck9btxjb](../../data/2026/W37/grafana.json); vendor URL: https://stspg.io/lr44wpz3m0tz
- Start time basis: `vendor_reported`.

<a id="ci-43a9a49695dad394c58e93af"></a>

- Canonical ID: `ci-43a9a49695dad394c58e93af` (1 source records).
  - [anthropic/r1pqn1kb4hvk](../../data/2026/W37/anthropic.json); vendor URL: https://stspg.io/5yp3rhhztm30
- Start time basis: `vendor_reported`.

<a id="ci-2ed8c7d7199fd963f73eaa14"></a>

- Canonical ID: `ci-2ed8c7d7199fd963f73eaa14` (1 source records).
  - [cloudflare/sjs8s0q2x4hw](../../data/2026/W37/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/sjs8s0q2x4hw
- Start time basis: `vendor_reported`.

<a id="ci-04ef0ea10c5ad910cb8e2ee1"></a>

- Canonical ID: `ci-04ef0ea10c5ad910cb8e2ee1` (1 source records).
  - [cloudflare/k7l4hzmthl3k](../../data/2026/W37/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/k7l4hzmthl3k
- Start time basis: `vendor_reported`.

<a id="ci-32bae5afda3ee87d4028bf97"></a>

- Canonical ID: `ci-32bae5afda3ee87d4028bf97` (1 source records).
  - [openai/01M24HYC9SX2Q4471EVMY6EXQE](../../data/2026/W37/openai.json); vendor URL: https://status.openai.com/incidents/01M24HYC9SX2Q4471EVMY6EXQE
- Start time basis: `vendor_reported`.

<a id="ci-15e56813f70800815118200d"></a>

- Canonical ID: `ci-15e56813f70800815118200d` (1 source records).
  - [cloudflare/z16209cfb1xv](../../data/2026/W37/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/z16209cfb1xv
- Start time basis: `vendor_reported`.

<a id="ci-a43b656dcadcf900e2ff9834"></a>

- Canonical ID: `ci-a43b656dcadcf900e2ff9834` (1 source records).
  - [cloudflare/nhfcmzd1pv4p](../../data/2026/W37/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/nhfcmzd1pv4p
- Start time basis: `vendor_reported`.

<a id="ci-a69b1ce8c5f0ff9750fd3368"></a>

- Canonical ID: `ci-a69b1ce8c5f0ff9750fd3368` (1 source records).
  - [openai/01M28MEQWTQJDRCRPFD9FWQ0H3](../../data/2026/W37/openai.json); vendor URL: https://status.openai.com/incidents/01M28MEQWTQJDRCRPFD9FWQ0H3
- Start time basis: `vendor_reported`.

<a id="ci-8070d9450f8b693463537ceb"></a>

- Canonical ID: `ci-8070d9450f8b693463537ceb` (1 source records).
  - [stripe/hvy6s6zmp1mn](../../data/2026/W37/stripe.json); vendor URL: https://stspg.io/j5s6qs75yrjz
- Start time basis: `vendor_reported`.

<a id="ci-70138b680b5e92b13729def9"></a>

- Canonical ID: `ci-70138b680b5e92b13729def9` (1 source records).
  - [openai/01M20PYYYGRT9303VHAPA7YNT2](../../data/2026/W37/openai.json); vendor URL: https://status.openai.com/incidents/01M20PYYYGRT9303VHAPA7YNT2
- Start time basis: `vendor_reported`.

<a id="ci-b76104eccbb3efb2c5838b24"></a>

- Canonical ID: `ci-b76104eccbb3efb2c5838b24` (1 source records).
  - [github/0rn90wk115q9](../../data/2026/W37/github.json); vendor URL: https://stspg.io/8f3xch4y0v2k
- Start time basis: `vendor_reported`.

<a id="ci-80ac55f198b839ccd2197bc6"></a>

- Canonical ID: `ci-80ac55f198b839ccd2197bc6` (1 source records).
  - [slack/1578](../../data/2026/W37/slack.json); vendor URL: https://slack-status.com/2026-09/f171324a097abb9a
- Start time basis: `vendor_reported`.

<a id="ci-cc41854833958870b7fce062"></a>

- Canonical ID: `ci-cc41854833958870b7fce062` (1 source records).
  - [cloudflare/jsvvrtk5hvym](../../data/2026/W37/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/jsvvrtk5hvym
- Start time basis: `vendor_reported`.

<a id="ci-fa3fea16abc442a67734ed99"></a>

- Canonical ID: `ci-fa3fea16abc442a67734ed99` (1 source records).
  - [cloudflare/np4n61ckjdw7](../../data/2026/W37/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/np4n61ckjdw7
- Start time basis: `vendor_reported`.

<a id="ci-0def7dbcbd3e39560c2067b8"></a>

- Canonical ID: `ci-0def7dbcbd3e39560c2067b8` (1 source records).
  - [cloudflare/qz292kvy220g](../../data/2026/W37/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/qz292kvy220g
- Start time basis: `vendor_reported`.

<a id="ci-44680b434e693b65054736fd"></a>

- Canonical ID: `ci-44680b434e693b65054736fd` (1 source records).
  - [openai/01M20QBACYZMK2PRCJQCBSWTYJ](../../data/2026/W37/openai.json); vendor URL: https://status.openai.com/incidents/01M20QBACYZMK2PRCJQCBSWTYJ
- Start time basis: `vendor_reported`.

<a id="ci-479419e10dc071606b21b11d"></a>

- Canonical ID: `ci-479419e10dc071606b21b11d` (1 source records).
  - [anthropic/2pt83vlkk7x7](../../data/2026/W37/anthropic.json); vendor URL: https://stspg.io/hfs7749ykl14
- Start time basis: `vendor_reported`.

<a id="ci-7ae694cca04d1da5510133da"></a>

- Canonical ID: `ci-7ae694cca04d1da5510133da` (1 source records).
  - [cloudflare/t4gby3dqzbg1](../../data/2026/W37/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/t4gby3dqzbg1
- Start time basis: `vendor_reported`.

<a id="ci-ece347690603aa8f7ebdb2d4"></a>

- Canonical ID: `ci-ece347690603aa8f7ebdb2d4` (1 source records).
  - [cloudflare/2njywrxj8qjf](../../data/2026/W37/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/2njywrxj8qjf
- Start time basis: `vendor_reported`.

<a id="ci-5e2264e6ecacb808cb82d4aa"></a>

- Canonical ID: `ci-5e2264e6ecacb808cb82d4aa` (1 source records).
  - [grafana/20y0bl1tryqv](../../data/2026/W37/grafana.json); vendor URL: https://stspg.io/f5m752d4l87f
- Start time basis: `vendor_reported`.

<a id="ci-2fa2291e2b955996aa7c565a"></a>

- Canonical ID: `ci-2fa2291e2b955996aa7c565a` (1 source records).
  - [openai/01M27Q9GH7WVKNA61ZR3RDMSY5](../../data/2026/W37/openai.json); vendor URL: https://status.openai.com/incidents/01M27Q9GH7WVKNA61ZR3RDMSY5
- Start time basis: `vendor_reported`.

<a id="ci-041ec95f4fc9cc6d67d57f8f"></a>

- Canonical ID: `ci-041ec95f4fc9cc6d67d57f8f` (1 source records).
  - [yandexcloud/2005](../../data/2026/W37/yandexcloud.json); vendor URL: https://status.yandex.cloud/en/incidents/2005
- Start time basis: `vendor_reported`.

<a id="ci-f025dc68ef6163f688fbc921"></a>

- Canonical ID: `ci-f025dc68ef6163f688fbc921` (1 source records).
  - [openai/01M27VK1VN54RQJGJEHB5JTQJJ](../../data/2026/W37/openai.json); vendor URL: https://status.openai.com/incidents/01M27VK1VN54RQJGJEHB5JTQJJ
- Start time basis: `vendor_reported`.

<a id="ci-33521a347717a88a6c06e68b"></a>

- Canonical ID: `ci-33521a347717a88a6c06e68b` (1 source records).
  - [slack/1577](../../data/2026/W37/slack.json); vendor URL: https://slack-status.com/2026-09/c8ec111fea759b78
- Start time basis: `vendor_reported`.

<a id="ci-bdd5a7c5b1f4f279b9991113"></a>

- Canonical ID: `ci-bdd5a7c5b1f4f279b9991113` (1 source records).
  - [cloudflare/3r3qcgbb0714](../../data/2026/W37/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/3r3qcgbb0714
- Start time basis: `vendor_reported`.

<a id="ci-1af7d0aaf968c48551580870"></a>

- Canonical ID: `ci-1af7d0aaf968c48551580870` (1 source records).
  - [cloudflare/kjfs3vqkxkgs](../../data/2026/W37/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/kjfs3vqkxkgs
- Start time basis: `vendor_reported`.

<a id="ci-ba00d282682b590aa1cd8249"></a>

- Canonical ID: `ci-ba00d282682b590aa1cd8249` (1 source records).
  - [openai/01M26E3PPAYJMMRHX2MZJCTHMF](../../data/2026/W37/openai.json); vendor URL: https://status.openai.com/incidents/01M26E3PPAYJMMRHX2MZJCTHMF
- Start time basis: `vendor_reported`.

<a id="ci-065783e837beeb05c1e97d67"></a>

- Canonical ID: `ci-065783e837beeb05c1e97d67` (1 source records).
  - [openai/01M26DX2PP1AFHGJK5E4M6S2J6](../../data/2026/W37/openai.json); vendor URL: https://status.openai.com/incidents/01M26DX2PP1AFHGJK5E4M6S2J6
- Start time basis: `vendor_reported`.

<a id="ci-d1d0d7778d4922fce64ad009"></a>

- Canonical ID: `ci-d1d0d7778d4922fce64ad009` (1 source records).
  - [grafana/25953qg7xkdj](../../data/2026/W37/grafana.json); vendor URL: https://stspg.io/j1r7zw96mgh5
- Start time basis: `vendor_reported`.

<a id="ci-3f699c92bcc52a98224d2143"></a>

- Canonical ID: `ci-3f699c92bcc52a98224d2143` (1 source records).
  - [cloudflare/92j1xmlkscw3](../../data/2026/W37/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/92j1xmlkscw3
- Start time basis: `vendor_reported`.

<a id="ci-380881d38a65508e27962773"></a>

- Canonical ID: `ci-380881d38a65508e27962773` (1 source records).
  - [cloudflare/pjr4m9q1xxb1](../../data/2026/W37/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/pjr4m9q1xxb1
- Start time basis: `vendor_reported`.

<a id="ci-d10d658ebd6fdcb0396fe539"></a>

- Canonical ID: `ci-d10d658ebd6fdcb0396fe539` (1 source records).
  - [openai/01M23BNX5V1XT9P3NP0PX56FTN](../../data/2026/W37/openai.json); vendor URL: https://status.openai.com/incidents/01M23BNX5V1XT9P3NP0PX56FTN
- Start time basis: `vendor_reported`.

<a id="ci-e93379c888820ffa838fd3c9"></a>

- Canonical ID: `ci-e93379c888820ffa838fd3c9` (1 source records).
  - [anthropic/t33dncr5ydvl](../../data/2026/W37/anthropic.json); vendor URL: https://stspg.io/n9n8kfcxtkfm
- Start time basis: `vendor_reported`.

<a id="ci-26d32a31f92dbbf7be9ff4c7"></a>

- Canonical ID: `ci-26d32a31f92dbbf7be9ff4c7` (1 source records).
  - [cloudflare/pgn7y7hgd25l](../../data/2026/W37/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/pgn7y7hgd25l
- Start time basis: `vendor_reported`.

<a id="ci-6fa3d10f8821372be78b7f7b"></a>

- Canonical ID: `ci-6fa3d10f8821372be78b7f7b` (1 source records).
  - [cloudflare/hy0gsndpkypz](../../data/2026/W37/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/hy0gsndpkypz
- Start time basis: `vendor_reported`.

<a id="ci-621a3350dce43441a6fb34b9"></a>

- Canonical ID: `ci-621a3350dce43441a6fb34b9` (1 source records).
  - [stripe/51n7sdp7zl92](../../data/2026/W37/stripe.json); vendor URL: https://stspg.io/7wcb364xld9l
- Start time basis: `vendor_reported`.

<a id="ci-0bd5011d5c0b9b8ed7fd7844"></a>

- Canonical ID: `ci-0bd5011d5c0b9b8ed7fd7844` (1 source records).
  - [digitalocean/1t16vfgq3zpx](../../data/2026/W37/digitalocean.json); vendor URL: https://stspg.io/g4s2358yl9tt
- Start time basis: `vendor_reported`.

<a id="ci-02ed43c5e286001db63bda1d"></a>

- Canonical ID: `ci-02ed43c5e286001db63bda1d` (1 source records).
  - [openai/01M2DV6788RRSDZ5K98S5HN0A8](../../data/2026/W37/openai.json); vendor URL: https://status.openai.com/incidents/01M2DV6788RRSDZ5K98S5HN0A8
- Start time basis: `vendor_reported`.

<a id="ci-a033a11042dfe5c5cc6f8678"></a>

- Canonical ID: `ci-a033a11042dfe5c5cc6f8678` (1 source records).
  - [openai/01M2D1CZG0862JB4JHVW56QF89](../../data/2026/W37/openai.json); vendor URL: https://status.openai.com/incidents/01M2D1CZG0862JB4JHVW56QF89
- Start time basis: `vendor_reported`.

<a id="ci-c2c826bca3eeffabb8564763"></a>

- Canonical ID: `ci-c2c826bca3eeffabb8564763` (1 source records).
  - [openai/01M23KG62KKK448RN434CQ64Z8](../../data/2026/W37/openai.json); vendor URL: https://status.openai.com/incidents/01M23KG62KKK448RN434CQ64Z8
- Start time basis: `vendor_reported`.

<a id="ci-b28e98652deb3dee1f6895be"></a>

- Canonical ID: `ci-b28e98652deb3dee1f6895be` (1 source records).
  - [circleci/f2w7c18std4n](../../data/2026/W37/circleci.json); vendor URL: https://stspg.io/l0tsrtv1tkbl
- Start time basis: `vendor_reported`.

<a id="ci-b830f3dad2f45af1349b45b9"></a>

- Canonical ID: `ci-b830f3dad2f45af1349b45b9` (1 source records).
  - [cloudflare/kq11rg5y66f2](../../data/2026/W37/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/kq11rg5y66f2
- Start time basis: `vendor_reported`.

<a id="ci-068182968029b9a7195b734d"></a>

- Canonical ID: `ci-068182968029b9a7195b734d` (1 source records).
  - [cloudflare/dmc0gzyv8gcs](../../data/2026/W37/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/dmc0gzyv8gcs
- Start time basis: `vendor_reported`.

<a id="ci-f4e89b346a013c0cef993a28"></a>

- Canonical ID: `ci-f4e89b346a013c0cef993a28` (1 source records).
  - [cloudflare/9mcdtk1lphgl](../../data/2026/W37/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/9mcdtk1lphgl
- Start time basis: `vendor_reported`.

<a id="ci-ebce3eeb700f8e222488f75c"></a>

- Canonical ID: `ci-ebce3eeb700f8e222488f75c` (1 source records).
  - [grafana/n1b3mc363979](../../data/2026/W37/grafana.json); vendor URL: https://stspg.io/gcg65gc0vtw8
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
