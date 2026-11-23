# Week 2026-W47 — top 5 of 24 above threshold

2026-11-16 to 2026-11-22 (UTC).

50 source records (110 saved revisions) → 50 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

> **2 of 14 sources unavailable: datadog, yandexcloud.**
> Their incidents are missing from this week, and the feeds are
> too shallow and unpaginated to backfill afterwards.
> - `datadog` — last successful fetch 2026-10-29 11:52, 25 days before this snapshot. Error: GET https://status.datadoghq.com/api/v2/incidents.json: status 403
> - `yandexcloud` — last successful fetch 2026-10-29 11:52, 25 days before this snapshot. Error: GET https://status.cloud.yandex.ru/api/incidents?lang=en&page=4: status 403

## 1. grafana — Plugin installs and updates from the plugin catalog may fail

- URL: https://stspg.io/g8t0nj4lct2j
- Window: 2026-11-20 11:52 UTC to 2026-11-20 14:33 UTC (2h 41m)
- Impact: minor, status: resolved
- Components: AWS Australia - prod-au-southeast-1, AWS Australia - prod-au-southeast-2, AWS Brazil - prod-sa-east-1, AWS Canada - prod-ca-east-0, AWS Canada prod-ca-central-1, AWS East (VA) prod-us-east-1, AWS Germany - prod-eu-west-2, AWS Germany - prod-eu-west-4, AWS Germany prod-eu-central-1, AWS India - prod-ap-south-1, AWS Indonesia - prod-ap-southeast-2, AWS Indonesia - prod-ap-southeast-3, AWS Ireland - prod-eu-west-6, AWS Ireland eu-west-1, AWS Japan - prod-ap-northeast-0, AWS Singapore - prod-ap-southeast-1, AWS Sweden - prod-eu-north-0, AWS Sweden eu-north-1, AWS Switzerland - prod-eu-central-0, AWS Switzerland eu-central-2, AWS UAE - prod-me-central-1, AWS UK - prod-gb-south-1, AWS US East (OH) - prod-us-east-0, AWS US East (OH) - prod-us-east-2, AWS US East (VA) - prod-us-east-3, AWS US West - prod-us-west-0, AWS West prod-us-west-2, Azure Malasia - azure-malaysiawest-0, Azure Netherlands - prod-eu-west-3, Azure UAE - prod-me-central-2, Azure US Central - us-central2, Federal Cloud - AWS US Gov West, GCP Australia - prod-au-southeast-0, GCP Belgium - prod-eu-west-0, GCP Brazil - prod-sa-east-0, GCP India - prod-ap-south-0, GCP Saudi Arabia - prod-me-central-0, GCP Singapore - prod-ap-southeast-0, GCP UK - prod-gb-south-0, GCP US Central - prod-us-central-0, GCP US Central - prod-us-central-3, GCP US Central - prod-us-central-4, GCP US Central - prod-us-central-7, GCP US East - prod-us-east-1, play.grafana.org

<a id="ci-f80b8064b3c18ea7c6686c26"></a>

- Canonical ID: `ci-f80b8064b3c18ea7c6686c26` (1 source records).
  - [grafana/x868z5qvdcnn](../../data/2026/W47/grafana.json); vendor URL: https://stspg.io/g8t0nj4lct2j
- Start time basis: `vendor_reported`.

  - `2026-11-20 11:52` investigating — We are investigating an issue where installing or updating plugins from the plugin catalog may fail for a small number of Grafana Cloud stacks across all Grafana Cloud Regions. Customers may see an error when they try to install or update a plugin. We are actively working on a resolution. Next update by 12:50 UTC.
  - `2026-11-20 12:02` identified — We are investigating an issue where installing or updating plugins from the plugin catalog may fail for a small number of Grafana Cloud stacks across several Grafana Cloud Regions (prod-ap-south-0, prod-ap-southeast-1, prod-ap-southeast-2, prod-ap-southeast-3, prod-eu-central-0, prod-eu-west-8, prod-me-central-0, prod-me-central-2). Customers may see an error when they try to install or update a plugin. The issue has been identified and we are actively working on a resolution. Next update by 13:00 UTC.
  - `2026-11-20 13:05` identified — We are investigating an issue where installing or updating plugins from the plugin catalog may fail for a small number of Grafana Cloud stacks across several Grafana Cloud Regions (prod-ap-south-0, prod-ap-southeast-1, prod-ap-southeast-2, prod-ap-southeast-3, prod-eu-central-0, prod-eu-west-8, prod-me-central-0, prod-me-central-2). Customers may see an error when they try to install or update a plugin. The issue has been identified and we are actively working on a resolution. Next update by 14:00 UTC.
  - `2026-11-20 13:53` identified — We are aware of an issue where installing or updating plugins from the plugin catalog may fail for a small number of Grafana Cloud stacks across several Grafana Cloud Regions (prod-ap-south-0, prod-ap-southeast-1, prod-ap-southeast-2, prod-ap-southeast-3, prod-eu-central-0, prod-eu-west-8, prod-me-central-0, prod-me-central-2). Customers may see an error when they try to install or update a plugin. The issue has been identified and we are actively working on a resolution. Next update by 15:00 UTC.
  - `2026-11-20 14:05` monitoring — A fix has been implemented and we are monitoring the results. Next update by 15:00 UTC.
  - `2026-11-20 14:33` resolved — This incident has been resolved, and plugin functionality has been restored.

## 2. anthropic — Claude Meetings auto-join limited on the Team plan

- URL: https://stspg.io/01qq3kzyyqzl
- Window: 2026-11-16 15:10 UTC to 2026-11-19 02:15 UTC (2d 11h 4m)
- Impact: minor, status: resolved
- Components: claude.ai

<a id="ci-235a56445707110c95d29f41"></a>

- Canonical ID: `ci-235a56445707110c95d29f41` (1 source records).
  - [anthropic/b5cq662qf2wm](../../data/2026/W47/anthropic.json); vendor URL: https://stspg.io/01qq3kzyyqzl
- Start time basis: `vendor_reported`.

  - `2026-11-16 15:10` investigating — We are investigating an issue preventing Claude from joining new calls in Claude Meetings (beta). Calls already in progress are not affected, and core functionality is not impacted. We will provide an update as soon as possible.
  - `2026-11-16 15:12` investigating — We are continuing to investigate an issue preventing Claude from joining new calls in Claude Meetings (beta).
  - `2026-11-16 16:27` investigating — We have temporarily paused Auto-join while we work on this. Adding Claude to a call by hand is also affected. Calls already in progress are not affected. We will provide another update as soon as possible.
  - `2026-11-16 18:10` investigating — Claude may still be unable to join some calls in Claude Meetings (beta), and Auto-join remains paused. Calls already in progress and other Claude features are not affected. We expect this to take some time to fully resolve and will post an update when the status changes.
  - `2026-11-17 02:02` identified — Claude Meetings (beta) saw higher than expected demand today, and Claude was unable to join many calls. We have put mitigations in place and are continuing to monitor and improve join reliability. Claude Meetings is now re-enabled for all Enterprise customers, though Claude may take a little longer than usual to join. Adding Claude to new calls and Auto-join remain paused for Team customers. We will share more updates as we expand capacity and begin restoring the feature for Team customers.
  - `2026-11-17 23:39` identified — Claude Meetings (beta) is joining calls more reliably, and we are monitoring it. It remains enabled for all Enterprise customers. Team customers can again add Claude to new calls manually, and Auto-join is being restored gradually for Team customers.
  - `2026-11-19 02:15` resolved — Claude Meetings (beta) is joining calls normally. Enterprise and Team customers can add Claude to calls. Auto-join is available for all Enterprise customers and some Team customers. We'll continue rolling out auto-join to additional Team customers in the coming days.

## 3. openai — Costs API data is delayed

- URL: https://status.openai.com/incidents/01M7N9VNR6X5SEC90M1YCDB36D
- Window: 2026-11-17 17:46 UTC to 2026-11-18 17:15 UTC (23h 29m)
- Impact: minor, status: resolved
- Components: —

<a id="ci-76934f15b2aaa10e3f4290ba"></a>

- Canonical ID: `ci-76934f15b2aaa10e3f4290ba` (1 source records).
  - [openai/01M7N9VNR6X5SEC90M1YCDB36D](../../data/2026/W47/openai.json); vendor URL: https://status.openai.com/incidents/01M7N9VNR6X5SEC90M1YCDB36D
- Start time basis: `vendor_reported`.

  - `2026-11-17 17:46` investigating — We have identified a delay in our Costs API. We are working to resolve the issue.
  - `2026-11-17 18:41` investigating — We have identified the issue and we are working towards a full recovery.
  - `2026-11-17 19:53` monitoring — We've pushed a fix and are monitoring the recovery.
  - `2026-11-18 17:15` resolved — All affected services are operating normally again. We’re sorry for the disruption, and we appreciate your patience.
  - `2026-11-19 02:42` resolved — All affected services are operating normally again. We’re sorry for the disruption, and we appreciate your patience.

## 4. cloudflare — Degraded connectivity between Europe and Asia Pacific

- URL: https://www.cloudflarestatus.com/incidents/f33c51cw0ojd
- Window: 2026-11-18 22:26 UTC to 2026-11-19 03:19 UTC (4h 52m)
- Impact: minor, status: resolved
- Components: Cache Reserve, D1, R2

<a id="ci-ed6c455e25909b9985a99948"></a>

- Canonical ID: `ci-ed6c455e25909b9985a99948` (1 source records).
  - [cloudflare/f33c51cw0ojd](../../data/2026/W47/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/f33c51cw0ojd
- Start time basis: `vendor_reported`.

  - `2026-11-18 22:26` investigating — Cloudflare is investigating degraded connectivity between Europe and the Asia Pacific region. Customers may experience increased errors and latency for requests traversing this path, including requests involving R2, D1, and Cache Reserve. Other Cloudflare services that rely on Europe-to-Asia Pacific connectivity may also be affected. Network teams are working to mitigate the issue and restore normal connectivity. Further updates will be provided as more information becomes available.
  - `2026-11-18 23:48` identified — The issue has been identified and a fix is being implemented.
  - `2026-11-19 00:36` monitoring — A fix has been implemented and we are monitoring the results.
  - `2026-11-19 03:19` resolved — This incident has been resolved.

## 5. cloudflare — Errors and stalled predictions on some models

- URL: https://www.cloudflarestatus.com/incidents/sv9tnfpjiojf
- Window: 2026-11-17 15:02 UTC to 2026-11-18 03:23 UTC (12h 20m)
- Impact: minor, status: resolved
- Components: Replicate

<a id="ci-9574f530ec096ca780c018a2"></a>

- Canonical ID: `ci-9574f530ec096ca780c018a2` (1 source records).
  - [cloudflare/sv9tnfpjiojf](../../data/2026/W47/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/sv9tnfpjiojf
- Start time basis: `vendor_reported`.

  - `2026-11-17 15:02` investigating — We are currently investigating this issue. Cloudflare is aware of, and investigating, an issue affecting predictions for some models on Replicate. Customers may see prediction requests return HTTP 500 errors, and some predictions may remain in a processing state without completing. Some API responses are also missing version information, which can cause validation errors in client libraries. We are working to understand the full impact and mitigate this problem.
  - `2026-11-17 16:32` identified — The issue has been identified and a fix is being implemented. Current impact: prediction requests for some models may return HTTP 500 errors, and some predictions may stay in a processing state. Customers using the API may also see validation errors on responses.
  - `2026-11-17 21:53` identified — The issue has been identified and a fix is being implemented. Current impact: predictions on OpenAI models are back to normal. Predictions on other models may still return HTTP 500 errors or stay in a processing state while we continue to work on a fix.
  - `2026-11-18 02:56` monitoring — A fix has been implemented and we are monitoring the results.
  - `2026-11-18 03:23` resolved — This incident has been resolved.

## Also above threshold (19)

- `cloudflare` Network Performance Issues in Washington, DC datacenter (IAD) (11h 8m) — https://www.cloudflarestatus.com/incidents/oz6to8yacrxe
- `cloudflare` Observability ingestion delay and query failures (5h 25m) — https://www.cloudflarestatus.com/incidents/ljrvdciu7syf
- `cloudflare` Network Performance Issues in Miami (9h 27m) — https://www.cloudflarestatus.com/incidents/4lv21sf5fzhm
- `github` Incident with Pull Requests (4h 26m) — https://stspg.io/ljtn69jlt087
- `cloudflare` Issues with Durable Objects and Dependent Services (2h 21m) — https://www.cloudflarestatus.com/incidents/ylj5po752g6h
- `cloudflare` Turnstile verification failures (6h 12m) — https://www.cloudflarestatus.com/incidents/5831aw44ia7i
- `digitalocean` Increased Error Rates in SFO3 (13h 26m) — https://stspg.io/bnb8ltd5ct7l
- `anthropic` Elevated errors for Claude Sonnet 5.5 and Claude Fable 5.2 (1h 14m) — https://stspg.io/n18qjkd7dt4q
- `stripe` Elevated Apple Pay API errors (2h 50m) — https://stspg.io/twy5139vhmjq
- `cloudflare` Network Performance Issues in North America (3h 0m) — https://www.cloudflarestatus.com/incidents/76lqmua6x3tk
- `cloudflare` Increased latency on Analytics API (GraphQL) queries (2h 47m) — https://www.cloudflarestatus.com/incidents/9zpa7own49hb
- `anthropic` Some users unable to sign in with Apple on claude.ai and the Claude iOS app (2h 55m) — https://stspg.io/3tl6h35nnzwm
- `cloudflare` Cloudflare Access Issue (2h 8m) — https://www.cloudflarestatus.com/incidents/sd2sya0mc6df
- `cloudflare` R2 elevated errors and timeouts for requests from South America (2h 3m) — https://www.cloudflarestatus.com/incidents/gcn9xop0f7zm
- `circleci` Jobs using IP ranges failing with infrastructure errors (2h 26m) — https://stspg.io/m8bpg1brfjcw
- `anthropic` Elevated errors for Claude Mythos 5.2 and Claude Fable 5.2 (22m) — https://stspg.io/nlb5fptv28h5
- `digitalocean` App Platform Accessibility in Cloud Panel (2h 29m) — https://stspg.io/4ccggg9c514v
- `openai` Elevated error rates for gpt-6 Luna API requests (56m) — https://status.openai.com/incidents/01M7T8MRF9SDBDNQ62Z6V156ZZ
- `cloudflare` Cloudflare Dashboard and Cloudflare API service issues in EU region (28m) — https://www.cloudflarestatus.com/incidents/iysf3fasf7wn
