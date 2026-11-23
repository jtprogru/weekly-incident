# Week 2026-W47

2026-11-16 to 2026-11-22 (UTC). Generated 2026-11-23 16:39.

50 canonical incidents, 24 above threshold.

50 source records (110 saved revisions) → 50 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

> **2 of 14 sources unavailable: datadog, yandexcloud.**
> Their incidents are missing from this week, and the feeds are
> too shallow and unpaginated to backfill afterwards.
> - `datadog` — last successful fetch 2026-10-29 11:52, 25 days before this snapshot. Error: GET https://status.datadoghq.com/api/v2/incidents.json: status 403
> - `yandexcloud` — last successful fetch 2026-10-29 11:52, 25 days before this snapshot. Error: GET https://status.cloud.yandex.ru/api/incidents?lang=en&page=4: status 403

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | grafana | [Plugin installs and updates from the plugin catalog may fail](https://stspg.io/g8t0nj4lct2j) | 2026-11-20 11:52 | 2h 41m | minor | AWS Australia - prod-au-southeast-1, AWS Australia - prod-au-southeast-2, AWS Brazil - prod-sa-east-1, AWS Canada - prod-ca-east-0, AWS Canada prod-ca-central-1, AWS East (VA) prod-us-east-1, AWS Germany - prod-eu-west-2, AWS Germany - prod-eu-west-4, AWS Germany prod-eu-central-1, AWS India - prod-ap-south-1, AWS Indonesia - prod-ap-southeast-2, AWS Indonesia - prod-ap-southeast-3, AWS Ireland - prod-eu-west-6, AWS Ireland eu-west-1, AWS Japan - prod-ap-northeast-0, AWS Singapore - prod-ap-southeast-1, AWS Sweden - prod-eu-north-0, AWS Sweden eu-north-1, AWS Switzerland - prod-eu-central-0, AWS Switzerland eu-central-2, AWS UAE - prod-me-central-1, AWS UK - prod-gb-south-1, AWS US East (OH) - prod-us-east-0, AWS US East (OH) - prod-us-east-2, AWS US East (VA) - prod-us-east-3, AWS US West - prod-us-west-0, AWS West prod-us-west-2, Azure Malasia - azure-malaysiawest-0, Azure Netherlands - prod-eu-west-3, Azure UAE - prod-me-central-2, Azure US Central - us-central2, Federal Cloud - AWS US Gov West, GCP Australia - prod-au-southeast-0, GCP Belgium - prod-eu-west-0, GCP Brazil - prod-sa-east-0, GCP India - prod-ap-south-0, GCP Saudi Arabia - prod-me-central-0, GCP Singapore - prod-ap-southeast-0, GCP UK - prod-gb-south-0, GCP US Central - prod-us-central-0, GCP US Central - prod-us-central-3, GCP US Central - prod-us-central-4, GCP US Central - prod-us-central-7, GCP US East - prod-us-east-1, play.grafana.org | 7245 = 161 × 45 × 1.0 |
| 2 | anthropic | [Claude Meetings auto-join limited on the Team plan](https://stspg.io/01qq3kzyyqzl) | 2026-11-16 15:10 | 2d 11h 4m | minor | claude.ai | 5316 = 3544 × 1 × 1.5 |
| 3 | openai | [Costs API data is delayed](https://status.openai.com/incidents/01M7N9VNR6X5SEC90M1YCDB36D) | 2026-11-17 17:46 | 23h 29m | minor | — | 2114 = 1409 × 1 × 1.5 |
| 4 | cloudflare | [Degraded connectivity between Europe and Asia Pacific](https://www.cloudflarestatus.com/incidents/f33c51cw0ojd) | 2026-11-18 22:26 | 4h 52m | minor | Cache Reserve, D1, R2 | 1577 = 292 × 3 × 1.8 |
| 5 | cloudflare | [Errors and stalled predictions on some models](https://www.cloudflarestatus.com/incidents/sv9tnfpjiojf) | 2026-11-17 15:02 | 12h 20m | minor | Replicate | 1332 = 740 × 1 × 1.8 |
| 6 | cloudflare | [Network Performance Issues in Washington, DC datacenter (IAD)](https://www.cloudflarestatus.com/incidents/oz6to8yacrxe) | 2026-11-20 20:35 | 11h 8m | minor | Network | 1202 = 668 × 1 × 1.8 |
| 7 | cloudflare | [Observability ingestion delay and query failures](https://www.cloudflarestatus.com/incidents/ljrvdciu7syf) | 2026-11-19 16:30 | 5h 25m | minor | Workers, Workers Observability | 1170 = 325 × 2 × 1.8 |
| 8 | cloudflare | [Network Performance Issues in Miami](https://www.cloudflarestatus.com/incidents/4lv21sf5fzhm) | 2026-11-20 05:00 | 9h 27m | minor | Miami, United States - (MIA) | 1021 = 567 × 1 × 1.8 |
| 9 | github | [Incident with Pull Requests](https://stspg.io/ljtn69jlt087) | 2026-11-16 11:59 | 4h 26m | minor | Actions, Pull Requests | 798 = 266 × 2 × 1.5 |
| 10 | cloudflare | [Issues with Durable Objects and Dependent Services](https://www.cloudflarestatus.com/incidents/ylj5po752g6h) | 2026-11-17 21:56 | 2h 21m | minor | Containers, Durable Objects, Stream | 761 = 141 × 3 × 1.8 |
| 11 | cloudflare | [Turnstile verification failures](https://www.cloudflarestatus.com/incidents/5831aw44ia7i) | 2026-11-17 10:24 | 6h 12m | minor | Turnstile | 670 = 372 × 1 × 1.8 |
| 12 | digitalocean | [Increased Error Rates in SFO3](https://stspg.io/bnb8ltd5ct7l) | 2026-11-19 04:37 | 13h 26m | minor | SFO3 | 645 = 806 × 1 × 0.8 |
| 13 | anthropic | [Elevated errors for Claude Sonnet 5.5 and Claude Fable 5.2](https://stspg.io/n18qjkd7dt4q) | 2026-11-16 07:49 | 1h 14m | major | Claude API (api.anthropic.com), Claude Code, Claude Cowork, claude.ai | 444 = 74 × 4 × 1.5 |
| 14 | stripe | [Elevated Apple Pay API errors](https://stspg.io/twy5139vhmjq) | 2026-11-17 17:14 | 2h 50m | none | Acquirers and payment methods, Stripe API | 408 = 170 × 2 × 1.2 |
| 15 | cloudflare | [Network Performance Issues in North America](https://www.cloudflarestatus.com/incidents/76lqmua6x3tk) | 2026-11-18 07:00 | 3h 0m | minor | Network | 324 = 180 × 1 × 1.8 |
| 16 | cloudflare | [Increased latency on Analytics API (GraphQL) queries](https://www.cloudflarestatus.com/incidents/9zpa7own49hb) | 2026-11-18 20:35 | 2h 47m | minor | — | 301 = 167 × 1 × 1.8 |
| 17 | anthropic | [Some users unable to sign in with Apple on claude.ai and the Claude iOS app](https://stspg.io/3tl6h35nnzwm) | 2026-11-17 18:55 | 2h 55m | major | claude.ai | 262 = 175 × 1 × 1.5 |
| 18 | cloudflare | [Cloudflare Access Issue](https://www.cloudflarestatus.com/incidents/sd2sya0mc6df) | 2026-11-19 06:48 | 2h 8m | minor | Access | 230 = 128 × 1 × 1.8 |
| 19 | cloudflare | [R2 elevated errors and timeouts for requests from South America](https://www.cloudflarestatus.com/incidents/gcn9xop0f7zm) | 2026-11-17 20:05 | 2h 3m | minor | R2 | 221 = 123 × 1 × 1.8 |
| 20 | circleci | [Jobs using IP ranges failing with infrastructure errors](https://stspg.io/m8bpg1brfjcw) | 2026-11-20 21:22 | 2h 26m | minor | Docker Jobs | 146 = 146 × 1 × 1.0 |
| 21 | anthropic | [Elevated errors for Claude Mythos 5.2 and Claude Fable 5.2](https://stspg.io/nlb5fptv28h5) | 2026-11-19 09:10 | 22m | major | Claude API (api.anthropic.com), Claude Code, Claude Cowork, claude.ai | 132 = 22 × 4 × 1.5 |
| 22 | digitalocean | [App Platform Accessibility in Cloud Panel](https://stspg.io/4ccggg9c514v) | 2026-11-16 19:10 | 2h 29m | none | — | 119 = 149 × 1 × 0.8 |
| 23 | openai | [Elevated error rates for gpt-6 Luna API requests](https://status.openai.com/incidents/01M7T8MRF9SDBDNQ62Z6V156ZZ) | 2026-11-19 16:01 | 56m | major | — | 84 = 56 × 1 × 1.5 |
| 24 | cloudflare | [Cloudflare Dashboard and Cloudflare API service issues in EU region](https://www.cloudflarestatus.com/incidents/iysf3fasf7wn) | 2026-11-20 18:24 | 28m | major | — | 50 = 28 × 1 × 1.8 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. grafana — Plugin installs and updates from the plugin catalog may fail

- URL: https://stspg.io/g8t0nj4lct2j
- Started: 2026-11-20 11:52 UTC
- Resolved: 2026-11-20 14:33 UTC
- Duration: 2h 41m
- Impact: minor, status: resolved
- Components: AWS Australia - prod-au-southeast-1, AWS Australia - prod-au-southeast-2, AWS Brazil - prod-sa-east-1, AWS Canada - prod-ca-east-0, AWS Canada prod-ca-central-1, AWS East (VA) prod-us-east-1, AWS Germany - prod-eu-west-2, AWS Germany - prod-eu-west-4, AWS Germany prod-eu-central-1, AWS India - prod-ap-south-1, AWS Indonesia - prod-ap-southeast-2, AWS Indonesia - prod-ap-southeast-3, AWS Ireland - prod-eu-west-6, AWS Ireland eu-west-1, AWS Japan - prod-ap-northeast-0, AWS Singapore - prod-ap-southeast-1, AWS Sweden - prod-eu-north-0, AWS Sweden eu-north-1, AWS Switzerland - prod-eu-central-0, AWS Switzerland eu-central-2, AWS UAE - prod-me-central-1, AWS UK - prod-gb-south-1, AWS US East (OH) - prod-us-east-0, AWS US East (OH) - prod-us-east-2, AWS US East (VA) - prod-us-east-3, AWS US West - prod-us-west-0, AWS West prod-us-west-2, Azure Malasia - azure-malaysiawest-0, Azure Netherlands - prod-eu-west-3, Azure UAE - prod-me-central-2, Azure US Central - us-central2, Federal Cloud - AWS US Gov West, GCP Australia - prod-au-southeast-0, GCP Belgium - prod-eu-west-0, GCP Brazil - prod-sa-east-0, GCP India - prod-ap-south-0, GCP Saudi Arabia - prod-me-central-0, GCP Singapore - prod-ap-southeast-0, GCP UK - prod-gb-south-0, GCP US Central - prod-us-central-0, GCP US Central - prod-us-central-3, GCP US Central - prod-us-central-4, GCP US Central - prod-us-central-7, GCP US East - prod-us-east-1, play.grafana.org
- Canonical incident: [ci-f80b8064b3c18ea7c6686c26](#ci-f80b8064b3c18ea7c6686c26)

- `2026-11-20 11:52` **investigating** — We are investigating an issue where installing or updating plugins from the plugin catalog may fail for a small number of Grafana Cloud stacks across all Grafana Cloud Regions. Customers may see an error when they try to install or update a plugin. We are actively working on a resolution. Next update by 12:50 UTC.
- `2026-11-20 12:02` **identified** — We are investigating an issue where installing or updating plugins from the plugin catalog may fail for a small number of Grafana Cloud stacks across several Grafana Cloud Regions (prod-ap-south-0, prod-ap-southeast-1, prod-ap-southeast-2, prod-ap-southeast-3, prod-eu-central-0, prod-eu-west-8, prod-me-central-0, prod-me-central-2). Customers may see an error when they try to install or update a plugin. The issue has been identified and we are actively working on a resolution. Next update by 13:00 UTC.
- `2026-11-20 13:05` **identified** — We are investigating an issue where installing or updating plugins from the plugin catalog may fail for a small number of Grafana Cloud stacks across several Grafana Cloud Regions (prod-ap-south-0, prod-ap-southeast-1, prod-ap-southeast-2, prod-ap-southeast-3, prod-eu-central-0, prod-eu-west-8, prod-me-central-0, prod-me-central-2). Customers may see an error when they try to install or update a plugin. The issue has been identified and we are actively working on a resolution. Next update by 14:00 UTC.
- `2026-11-20 13:53` **identified** — We are aware of an issue where installing or updating plugins from the plugin catalog may fail for a small number of Grafana Cloud stacks across several Grafana Cloud Regions (prod-ap-south-0, prod-ap-southeast-1, prod-ap-southeast-2, prod-ap-southeast-3, prod-eu-central-0, prod-eu-west-8, prod-me-central-0, prod-me-central-2). Customers may see an error when they try to install or update a plugin. The issue has been identified and we are actively working on a resolution. Next update by 15:00 UTC.
- `2026-11-20 14:05` **monitoring** — A fix has been implemented and we are monitoring the results. Next update by 15:00 UTC.
- `2026-11-20 14:33` **resolved** — This incident has been resolved, and plugin functionality has been restored.

### 2. anthropic — Claude Meetings auto-join limited on the Team plan

- URL: https://stspg.io/01qq3kzyyqzl
- Started: 2026-11-16 15:10 UTC
- Resolved: 2026-11-19 02:15 UTC
- Duration: 2d 11h 4m
- Impact: minor, status: resolved
- Components: claude.ai
- Canonical incident: [ci-235a56445707110c95d29f41](#ci-235a56445707110c95d29f41)

- `2026-11-16 15:10` **investigating** — We are investigating an issue preventing Claude from joining new calls in Claude Meetings (beta). Calls already in progress are not affected, and core functionality is not impacted. We will provide an update as soon as possible.
- `2026-11-16 15:12` **investigating** — We are continuing to investigate an issue preventing Claude from joining new calls in Claude Meetings (beta).
- `2026-11-16 16:27` **investigating** — We have temporarily paused Auto-join while we work on this. Adding Claude to a call by hand is also affected. Calls already in progress are not affected. We will provide another update as soon as possible.
- `2026-11-16 18:10` **investigating** — Claude may still be unable to join some calls in Claude Meetings (beta), and Auto-join remains paused. Calls already in progress and other Claude features are not affected. We expect this to take some time to fully resolve and will post an update when the status changes.
- `2026-11-17 02:02` **identified** — Claude Meetings (beta) saw higher than expected demand today, and Claude was unable to join many calls. We have put mitigations in place and are continuing to monitor and improve join reliability. Claude Meetings is now re-enabled for all Enterprise customers, though Claude may take a little longer than usual to join. Adding Claude to new calls and Auto-join remain paused for Team customers. We will share more updates as we expand capacity and begin restoring the feature for Team customers.
- `2026-11-17 23:39` **identified** — Claude Meetings (beta) is joining calls more reliably, and we are monitoring it. It remains enabled for all Enterprise customers. Team customers can again add Claude to new calls manually, and Auto-join is being restored gradually for Team customers.
- `2026-11-19 02:15` **resolved** — Claude Meetings (beta) is joining calls normally. Enterprise and Team customers can add Claude to calls. Auto-join is available for all Enterprise customers and some Team customers. We'll continue rolling out auto-join to additional Team customers in the coming days.

### 3. openai — Costs API data is delayed

- URL: https://status.openai.com/incidents/01M7N9VNR6X5SEC90M1YCDB36D
- Started: 2026-11-17 17:46 UTC
- Resolved: 2026-11-18 17:15 UTC
- Duration: 23h 29m
- Impact: minor, status: resolved
- Components: —
- Canonical incident: [ci-76934f15b2aaa10e3f4290ba](#ci-76934f15b2aaa10e3f4290ba)

- `2026-11-17 17:46` **investigating** — We have identified a delay in our Costs API. We are working to resolve the issue.
- `2026-11-17 18:41` **investigating** — We have identified the issue and we are working towards a full recovery.
- `2026-11-17 19:53` **monitoring** — We've pushed a fix and are monitoring the recovery.
- `2026-11-18 17:15` **resolved** — All affected services are operating normally again. We’re sorry for the disruption, and we appreciate your patience.
- `2026-11-19 02:42` **resolved** — All affected services are operating normally again. We’re sorry for the disruption, and we appreciate your patience.

### 4. cloudflare — Degraded connectivity between Europe and Asia Pacific

- URL: https://www.cloudflarestatus.com/incidents/f33c51cw0ojd
- Started: 2026-11-18 22:26 UTC
- Resolved: 2026-11-19 03:19 UTC
- Duration: 4h 52m
- Impact: minor, status: resolved
- Components: Cache Reserve, D1, R2
- Canonical incident: [ci-ed6c455e25909b9985a99948](#ci-ed6c455e25909b9985a99948)

- `2026-11-18 22:26` **investigating** — Cloudflare is investigating degraded connectivity between Europe and the Asia Pacific region. Customers may experience increased errors and latency for requests traversing this path, including requests involving R2, D1, and Cache Reserve. Other Cloudflare services that rely on Europe-to-Asia Pacific connectivity may also be affected. Network teams are working to mitigate the issue and restore normal connectivity. Further updates will be provided as more information becomes available.
- `2026-11-18 23:48` **identified** — The issue has been identified and a fix is being implemented.
- `2026-11-19 00:36` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-11-19 03:19` **resolved** — This incident has been resolved.

### 5. cloudflare — Errors and stalled predictions on some models

- URL: https://www.cloudflarestatus.com/incidents/sv9tnfpjiojf
- Started: 2026-11-17 15:02 UTC
- Resolved: 2026-11-18 03:23 UTC
- Duration: 12h 20m
- Impact: minor, status: resolved
- Components: Replicate
- Canonical incident: [ci-9574f530ec096ca780c018a2](#ci-9574f530ec096ca780c018a2)

- `2026-11-17 15:02` **investigating** — We are currently investigating this issue. Cloudflare is aware of, and investigating, an issue affecting predictions for some models on Replicate. Customers may see prediction requests return HTTP 500 errors, and some predictions may remain in a processing state without completing. Some API responses are also missing version information, which can cause validation errors in client libraries. We are working to understand the full impact and mitigate this problem.
- `2026-11-17 16:32` **identified** — The issue has been identified and a fix is being implemented. Current impact: prediction requests for some models may return HTTP 500 errors, and some predictions may stay in a processing state. Customers using the API may also see validation errors on responses.
- `2026-11-17 21:53` **identified** — The issue has been identified and a fix is being implemented. Current impact: predictions on OpenAI models are back to normal. Predictions on other models may still return HTTP 500 errors or stay in a processing state while we continue to work on a fix.
- `2026-11-18 02:56` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-11-18 03:23` **resolved** — This incident has been resolved.

### 6. cloudflare — Network Performance Issues in Washington, DC datacenter (IAD)

- URL: https://www.cloudflarestatus.com/incidents/oz6to8yacrxe
- Started: 2026-11-20 20:35 UTC
- Resolved: 2026-11-21 07:44 UTC
- Duration: 11h 8m
- Impact: minor, status: resolved
- Components: Network
- Canonical incident: [ci-dc68d8b187ee618247c8134d](#ci-dc68d8b187ee618247c8134d)

- `2026-11-20 20:35` **investigating** — Cloudflare is investigating issues with network performance in Washington, DC datacenter (IAD). We are working to analyze and mitigate this problem. More updates to follow shortly.
- `2026-11-20 20:36` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-11-21 07:44` **resolved** — This incident has been resolved.

### 7. cloudflare — Observability ingestion delay and query failures

- URL: https://www.cloudflarestatus.com/incidents/ljrvdciu7syf
- Started: 2026-11-19 16:30 UTC
- Resolved: 2026-11-19 21:55 UTC
- Duration: 5h 25m
- Impact: minor, status: resolved
- Components: Workers, Workers Observability
- Canonical incident: [ci-b6b7e2545b8368ef0ee80185](#ci-b6b7e2545b8368ef0ee80185)

- `2026-11-19 16:30` **identified** — Cloudflare has identified an issue impacting Observability ingestion and querying. Customers may be seeing delays in new data coming in and increased error rates querying Workers Observability and Cloudflare Traces data. We are working to analyze and mitigate this problem. More updates to follow shortly.
- `2026-11-19 21:23` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-11-19 21:55` **resolved** — This incident has been resolved.

### 8. cloudflare — Network Performance Issues in Miami

- URL: https://www.cloudflarestatus.com/incidents/4lv21sf5fzhm
- Started: 2026-11-20 05:00 UTC
- Resolved: 2026-11-20 14:28 UTC
- Duration: 9h 27m
- Impact: minor, status: resolved
- Components: Miami, United States - (MIA)
- Canonical incident: [ci-0630da281c5295644d89b2bc](#ci-0630da281c5295644d89b2bc)

- `2026-11-20 05:00` **investigating** — Cloudflare has identified issues with Network Performance in Miami. We are working to analyze and mitigate this problem. More updates to follow shortly.
- `2026-11-20 05:10` **identified** — The issue has been identified and a fix is being implemented.
- `2026-11-20 14:28` **resolved** — This incident has been resolved.

### 9. github — Incident with Pull Requests

- URL: https://stspg.io/ljtn69jlt087
- Started: 2026-11-16 11:59 UTC
- Resolved: 2026-11-16 16:26 UTC
- Duration: 4h 26m
- Impact: minor, status: resolved
- Components: Actions, Pull Requests
- Canonical incident: [ci-78b7d2af2e7959c04ad01870](#ci-78b7d2af2e7959c04ad01870)

- `2026-11-16 11:59` **investigating** — We are investigating reports of degraded performance for Pull Requests
- `2026-11-16 12:00` **investigating** — We’re investigating elevated timeouts when merging pull requests. Some users may be unable to merge pull requests successfully; timeouts are occurring at a higher rate than usual. We’re working to identify the cause and restore normal performance.
- `2026-11-16 12:35` **investigating** — We’re continuing to investigate elevated pull request merge timeouts. Approximately 2–3% of merges are timing out, compared with a baseline near 0.2%. We’ll share another update as we learn more.
- `2026-11-16 13:24` **investigating** — We’re continuing to investigate elevated pull request merge timeouts. Approximately 2–3% of merges are timing out, compared with a baseline near 0.2%. Mitigations so far have not restored normal performance, and our teams are continuing to investigate. We’ll share another update as we learn more.
- `2026-11-16 14:06` **investigating** — We’re continuing to investigate elevated pull request merge timeouts. Approximately 3% of merges are timing out, compared with a baseline near 0.2%. Mitigations so far have not materially improved performance, and our teams are working to reduce load and restore normal service. We’ll share another update as we learn more.
- `2026-11-16 15:03` **investigating** — We’re continuing to investigate elevated pull request merge timeouts. The latest reported rate is 4.75%, compared with a baseline near 0.2%, so some users may have difficulty merging pull requests. We’ll share another update as we learn more.
- `2026-11-16 15:33` **investigating** — Actions is experiencing degraded performance. We are continuing to investigate.
- `2026-11-16 15:46` **investigating** — We’re seeing early signs of improvement in pull request merge times, but recovery is not yet confirmed. The last confirmed timeout rate was 2.5%, compared with a baseline near 0.2%, so some users may still have difficulty merging pull requests. We’ll share another update as we learn more.
- `2026-11-16 15:55` **investigating** — We’re seeing early signs of improvement in pull request merge times & workflow start times, but recovery is not yet confirmed. Users may still have difficulty merging pull requests. We’re reducing load on affected systems and monitoring the impact. We’ll share another update as we learn more.
- `2026-11-16 16:07` **monitoring** — The degradation affecting Actions and Pull Requests has been mitigated. We are monitoring to ensure stability.
- `2026-11-16 16:16` **monitoring** — We’ve mitigated the issue affecting pull request merges and Actions workflow start times by reducing load on affected systems. Pull Requests and Actions have recovered, and we’re continuing to monitor their performance.
- `2026-11-16 16:26` **resolved** — On November 16, 2026, between 09:30 and 15:55 UTC, some pull request merges on GitHub timed out. On average, 2.9% of merge attempts timed out, with a peak of 7.1% at 15:17 UTC. Some GitHub Actions workflow runs triggered by pull requests were also delayed. Git operations such as push and fetch stayed available, with higher latency. This was a repeat of a shorter event on November 13, between 13:15 and 17:50 UTC, when 1.7% of merges timed out.<br /><br />Both events had the same cause: some of the storage servers that hold Git repository data reached their disk read bandwidth limits. A merge must be written to every copy of a repository, so one overloaded server slowed merges for every repository it hosts. Read load on these servers had been growing as part of a planned move of repositories to new infrastructure. A disk configuration setting on those servers also made them read far more data than necessary, which cut their usable bandwidth roughly in half. We mitigated the incident by correcting that setting and moving some load away from the busiest servers.<br /><br />The corrected setting is now applied permanently across our storage fleet. We are adding storage capacity in the affected region and alerting on per-server disk bandwidth so we catch saturation before it affects customers. Sustained merge timeouts will now page our on-call engineers directly. We are also adding capacity checks before moving repositories or traffic onto new infrastructure.

### 10. cloudflare — Issues with Durable Objects and Dependent Services

- URL: https://www.cloudflarestatus.com/incidents/ylj5po752g6h
- Started: 2026-11-17 21:56 UTC
- Resolved: 2026-11-18 00:17 UTC
- Duration: 2h 21m
- Impact: minor, status: resolved
- Components: Containers, Durable Objects, Stream
- Canonical incident: [ci-068fb03b92865405df02a4a7](#ci-068fb03b92865405df02a4a7)

- `2026-11-17 21:56` **investigating** — We are investigating reports of elevated timeout rates and performance degradation affecting Durable Objects. This may impact dependent services, including Stream and Containers. Our team is actively investigating, and further updates will be provided shortly.
- `2026-11-18 00:17` **resolved** — This incident has been resolved.

### 11. cloudflare — Turnstile verification failures

- URL: https://www.cloudflarestatus.com/incidents/5831aw44ia7i
- Started: 2026-11-17 10:24 UTC
- Resolved: 2026-11-17 16:37 UTC
- Duration: 6h 12m
- Impact: minor, status: resolved
- Components: Turnstile
- Canonical incident: [ci-9c9bde87f65773819425e821](#ci-9c9bde87f65773819425e821)

- `2026-11-17 10:24` **investigating** — We are investigating reports of Turnstile verification failures affecting some users with browser fingerprinting protection enabled, including Helium. We are assessing the impact and potential mitigations.
- `2026-11-17 10:42` **identified** — The issue has been identified and a fix is being implemented.
- `2026-11-17 16:37` **resolved** — This incident has been resolved.

### 12. digitalocean — Increased Error Rates in SFO3

- URL: https://stspg.io/bnb8ltd5ct7l
- Started: 2026-11-19 04:37 UTC
- Resolved: 2026-11-19 18:04 UTC
- Duration: 13h 26m
- Impact: minor, status: resolved
- Components: SFO3
- Canonical incident: [ci-cc949bb74d1345ff2e38f59b](#ci-cc949bb74d1345ff2e38f59b)

- `2026-11-19 04:37` **investigating** — As of 02:53 UTC, our Engineering team is investigating increased error rates impacting multiple services in the SFO3 region. At this point, users with resources in SFO3 may experience errors, timeouts, or delays when creating or managing Droplets, Volumes, Load Balancers, Reserved IPs, Cloud Firewalls, Managed Databases, and Managed Kubernetes clusters through the API or Cloud Control Panel. Some users may also experience connectivity issues with Load Balancers and with newly created or restarted Droplets. We apologize for the inconvenience and will share an update once we have more information.
- `2026-11-19 04:45` **monitoring** — Our Engineering team has mitigated the increased error rates impacting multiple services in the SFO3 region and is monitoring the situation. The impact was caused by a network routing issue that disrupted connectivity to one of our core database systems after a load balancer node became unavailable. As of 04:32 UTC, users should no longer experience errors or timeouts when creating or managing resources in SFO3. Operations that failed during the incident may need to be retried. We will post an update as soon as the issue is fully resolved.
- `2026-11-19 05:46` **monitoring** — Our Engineering team has mitigated the issue, and services in the SFO3 region are currently operating normally. The team is continuing to monitor the situation to ensure stability given the scope of the impact. We will continue monitoring and share another update once the situation is fully resolved. In the meantime, if you encounter any further errors or timeouts, please retry the affected operations. If the issue persists, please submit a ticket through your Cloud Support page so our team can assist you further: https://cloudsupport.digitalocean.com/s/createticket
- `2026-11-19 18:04` **resolved** — Our Engineering team has confirmed that the increased error rates in the SFO3 region are resolved. From 02:53 UTC to 04:32 UTC, users may have experienced errors, timeouts, or delays when creating or managing Droplets, Volumes, Load Balancers, Reserved IPs, Cloud Firewalls, Managed Databases, and Managed Kubernetes clusters through the API or Cloud Control Panel. Some users also experienced connectivity issues with Load Balancers and with newly created or restarted Droplets. The impact was caused by a network routing issue that disrupted connectivity to one of our core database systems after a load balancer node became unavailable. Services have been stable since mitigation. Operations that failed during the incident may need to be retried. We apologize for the inconvenience. If you continue to experience issues, please open a support ticket from your DigitalOcean account.

### 13. anthropic — Elevated errors for Claude Sonnet 5.5 and Claude Fable 5.2

- URL: https://stspg.io/n18qjkd7dt4q
- Started: 2026-11-16 07:49 UTC
- Resolved: 2026-11-16 09:04 UTC
- Duration: 1h 14m
- Impact: major, status: resolved
- Components: Claude API (api.anthropic.com), Claude Code, Claude Cowork, claude.ai
- Canonical incident: [ci-6f5074bea4248258a8a899af](#ci-6f5074bea4248258a8a899af)

- `2026-11-16 07:49` **investigating** — We are currently investigating this issue.
- `2026-11-16 08:02` **monitoring** — A fix has been deployed for Claude Sonnet 5.5 and other Claude models and we are monitoring for recovery.
- `2026-11-16 09:04` **resolved** — The issue affecting claude.ai, the Claude API, Claude Code, and Claude Cowork has been resolved. Impact occurred from 23:32 PT on November 15 / 07:32 UTC on November 16 to 00:45 PT / 08:45 UTC on November 16.

### 14. stripe — Elevated Apple Pay API errors

- URL: https://stspg.io/twy5139vhmjq
- Started: 2026-11-17 17:14 UTC
- Resolved: 2026-11-17 20:05 UTC
- Duration: 2h 50m
- Impact: none, status: resolved
- Components: Acquirers and payment methods, Stripe API
- Canonical incident: [ci-326a8282f949b4d0bca76b3f](#ci-326a8282f949b4d0bca76b3f)

- `2026-11-17 17:14` **investigating** — We’re currently investigating elevated error rates and response times with Apple Pay. We’ll post an update soon.
- `2026-11-17 18:04` **investigating** — We're continuing to investigate elevated error rates and response times with Apple Pay. We'll post updates as they become available.
- `2026-11-17 20:05` **resolved** — From 16:34 to 19:50 UTC, we saw increased Apple Pay failures due to an issue external to Stripe. This is now resolved.

### 15. cloudflare — Network Performance Issues in North America

- URL: https://www.cloudflarestatus.com/incidents/76lqmua6x3tk
- Started: 2026-11-18 07:00 UTC
- Resolved: 2026-11-18 10:00 UTC
- Duration: 3h 0m
- Impact: minor, status: resolved
- Components: Network
- Canonical incident: [ci-e95653f83490d9397a3121b6](#ci-e95653f83490d9397a3121b6)

- `2026-11-18 10:56` **resolved** — Between approximately 07:00 UTC and 10:00 UTC Cloudflare investigated issues with network performance in North America. This has since been resolved.

### 16. cloudflare — Increased latency on Analytics API (GraphQL) queries

- URL: https://www.cloudflarestatus.com/incidents/9zpa7own49hb
- Started: 2026-11-18 20:35 UTC
- Resolved: 2026-11-18 23:22 UTC
- Duration: 2h 47m
- Impact: minor, status: resolved
- Components: —
- Canonical incident: [ci-d994572f710fb0dc6f58db17](#ci-d994572f710fb0dc6f58db17)

- `2026-11-18 20:35` **identified** — Cloudflare has identified an issue where some customers may experience latency querying the GraphQL API. We are working to analyze and mitigate this problem. More updates to follow shortly.
- `2026-11-18 20:40` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-11-18 23:22` **resolved** — This incident has been resolved.

### 17. anthropic — Some users unable to sign in with Apple on claude.ai and the Claude iOS app

- URL: https://stspg.io/3tl6h35nnzwm
- Started: 2026-11-17 18:55 UTC
- Resolved: 2026-11-17 21:51 UTC
- Duration: 2h 55m
- Impact: major, status: resolved
- Components: claude.ai
- Canonical incident: [ci-9e386d50157e8d852e3e1ebd](#ci-9e386d50157e8d852e3e1ebd)

- `2026-11-17 18:55` **identified** — We are investigating an issue with Sign in with Apple on claude.ai and the Claude iOS app. Sign in with Google and email are not affected. We will provide an update as soon as possible.
- `2026-11-17 18:56` **identified** — We have identified the cause of the issue affecting claude.ai and are working on a fix. We will provide an update as soon as possible.
- `2026-11-17 21:51` **resolved** — The issue affecting Sign in with Apple on claude.ai and the Claude iOS app has been resolved. Impact occurred from 08:40 PT / 16:40 UTC to 12:20 PT / 20:20 UTC.

### 18. cloudflare — Cloudflare Access Issue

- URL: https://www.cloudflarestatus.com/incidents/sd2sya0mc6df
- Started: 2026-11-19 06:48 UTC
- Resolved: 2026-11-19 08:56 UTC
- Duration: 2h 8m
- Impact: minor, status: resolved
- Components: Access
- Canonical incident: [ci-2d0f4c7401e3a1dcc97c23c2](#ci-2d0f4c7401e3a1dcc97c23c2)

- `2026-11-19 06:48` **identified** — Cloudflare is investigating an issue with Cloudflare Access. A recent Cloudflare Access release removed the sub claim from service token JWTs, which can cause authentication failures for applications that strictly validate token claims. The change has been reverted and the fix is rolling out globally.
- `2026-11-19 08:32` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-11-19 08:56` **resolved** — This incident has been resolved.

### 19. cloudflare — R2 elevated errors and timeouts for requests from South America

- URL: https://www.cloudflarestatus.com/incidents/gcn9xop0f7zm
- Started: 2026-11-17 20:05 UTC
- Resolved: 2026-11-17 22:08 UTC
- Duration: 2h 3m
- Impact: minor, status: resolved
- Components: R2
- Canonical incident: [ci-d175690cd304cc3bc7279620](#ci-d175690cd304cc3bc7279620)

- `2026-11-17 20:05` **investigating** — We are investigating elevated errors and read timeouts affecting R2 requests originating from South America (impacting datacenters such as GRU, GIG, EZE, and SCL). Our team is actively working to identify and resolve the issue.
- `2026-11-17 20:13` **identified** — The issue has been identified and a fix is being implemented.
- `2026-11-17 20:53` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-11-17 22:08` **resolved** — This incident has been resolved.

### 20. circleci — Jobs using IP ranges failing with infrastructure errors

- URL: https://stspg.io/m8bpg1brfjcw
- Started: 2026-11-20 21:22 UTC
- Resolved: 2026-11-20 23:48 UTC
- Duration: 2h 26m
- Impact: minor, status: resolved
- Components: Docker Jobs
- Canonical incident: [ci-0ac3b5acc951515826915c90](#ci-0ac3b5acc951515826915c90)

- `2026-11-20 21:22` **identified** — Customers running Docker executor jobs with the IP ranges feature enabled are experiencing a higher rate of infrastructure failures. A fix is being rolled out. Customers who see an infrastructure failure on one of these jobs can rerun the failed job. We will provide another update as we have more information to share.
- `2026-11-20 21:58` **identified** — We are continuing to rollout a fix for IP range docker jobs. Customers may continue to retry jobs that use IP ranges that experience infra fails. We will provide an update within 20 minutes on the fix.
- `2026-11-20 22:19` **identified** — The fix for IP range docker jobs are rolling out across the fleet. We will monitor for recovery in infra fails.
- `2026-11-20 22:27` **monitoring** — A fix has been fully deployed, we are now monitoring for recovery.
- `2026-11-20 23:11` **monitoring** — We are continuing to monitor the rollout for recovery.
- `2026-11-20 23:48` **resolved** — We have seen recovery for IP range jobs for gen2 docker jobs. We appreciate your patience while we worked through the incident.

### 21. anthropic — Elevated errors for Claude Mythos 5.2 and Claude Fable 5.2

- URL: https://stspg.io/nlb5fptv28h5
- Started: 2026-11-19 09:10 UTC
- Resolved: 2026-11-19 09:33 UTC
- Duration: 22m
- Impact: major, status: resolved
- Components: Claude API (api.anthropic.com), Claude Code, Claude Cowork, claude.ai
- Canonical incident: [ci-0ef84a7a30c5464ef7fdc5fb](#ci-0ef84a7a30c5464ef7fdc5fb)

- `2026-11-19 09:10` **investigating** — We are investigating elevated errors on requests to Claude Mythos 5.2 and Claude Fable 5.2. We will provide an update as soon as possible.
- `2026-11-19 09:22` **monitoring** — A fix has been deployed for other Claude models and we are monitoring for recovery.
- `2026-11-19 09:33` **resolved** — The issue affecting claude.ai, the Claude API, Claude Code, and Claude Cowork has been resolved. Impact occurred from 00:56 PT / 08:56 UTC to 01:11 PT / 09:11 UTC.

### 22. digitalocean — App Platform Accessibility in Cloud Panel

- URL: https://stspg.io/4ccggg9c514v
- Started: 2026-11-16 19:10 UTC
- Resolved: 2026-11-16 21:40 UTC
- Duration: 2h 29m
- Impact: none, status: resolved
- Components: —
- Canonical incident: [ci-46eb7eb190a3535f323d85ea](#ci-46eb7eb190a3535f323d85ea)

- `2026-11-16 19:10` **resolved** — Between 16:40 UTC and 18:22 UTC, our Engineering team observed an increase in 4xx error rates with App Platform operations via the Cloud Control Panel. During this time, users may have experienced errors while attempting to view or update their apps on the Cloud Panel. This was a UI-only issue and apps continued to run throughout the duration of the incident. Our Engineering team has resolved this issue. As of 18:22 UTC, users should no longer be experiencing this issue and all services should now be functioning normally. We apologize for any inconvenience, and if you are still experiencing any problems or have additional questions, please open a Support ticket within your account.

### 23. openai — Elevated error rates for gpt-6 Luna API requests

- URL: https://status.openai.com/incidents/01M7T8MRF9SDBDNQ62Z6V156ZZ
- Started: 2026-11-19 16:01 UTC
- Resolved: 2026-11-19 16:58 UTC
- Duration: 56m
- Impact: major, status: resolved
- Components: —
- Canonical incident: [ci-eadeb781439c44861c99f3ec](#ci-eadeb781439c44861c99f3ec)

- `2026-11-19 16:01` **investigating** — We’re investigating a potential issue for the listed services and checking how it affects users. We’ll share an update when we know more.
- `2026-11-19 16:21` **monitoring** — We’ve applied a mitigation and are monitoring the affected services to confirm recovery.
- `2026-11-19 16:58` **resolved** — All affected services are operating normally again. We’re sorry for the disruption, and we appreciate your patience.

### 24. cloudflare — Cloudflare Dashboard and Cloudflare API service issues in EU region

- URL: https://www.cloudflarestatus.com/incidents/iysf3fasf7wn
- Started: 2026-11-20 18:24 UTC
- Resolved: 2026-11-20 18:53 UTC
- Duration: 28m
- Impact: major, status: resolved
- Components: —
- Canonical incident: [ci-35b755c4579076a68a1f64a6](#ci-35b755c4579076a68a1f64a6)

- `2026-11-20 18:24` **investigating** — Cloudflare is investigating issues with Cloudflare Dashboard and related APIs in EU region. These issues do not affect the serving of cached files via the Cloudflare CDN or other security features at the Cloudflare Edge. Customers using the Dashboard / Cloudflare APIs are impacted as requests might fail and/or errors may be displayed.
- `2026-11-20 18:38` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-11-20 18:53` **resolved** — Cloudflare has resolved API and dashboard issues in the EU region.

## Below threshold (26)

- `cloudflare` [Network disruption across multiple colocations in India](https://www.cloudflarestatus.com/incidents/njt9j3k2zdfw) — 42m
- `cloudflare` [Network Performance Issues in Paris](https://www.cloudflarestatus.com/incidents/k3r9i5tbre4i) — 1h 12m
- `cloudflare` [R2 and Durable Objects availability issues](https://www.cloudflarestatus.com/incidents/pfwuex7qix27) — 1h 42m
- `cloudflare` [Cloudflare One Gateway traffic unable to connect to Cloudflare domains that are regionalized](https://www.cloudflarestatus.com/incidents/tawn3zx7wh0z) — 1h 2m
- `cloudflare` [Network Performance Issues in Chicago datacenter (ORD)](https://www.cloudflarestatus.com/incidents/d2m4kw94ut4f) — 38m
- `circleci` [Delays and intermittent failures running Smarter Testing](https://stspg.io/g20zht6r8c4w) — 51m
- `cloudflare` [Errors creating predictions for some models](https://www.cloudflarestatus.com/incidents/69855o2zz6kl) — 1h 48m
- `cloudflare` [Cache Purging Delays in Amsterdam](https://www.cloudflarestatus.com/incidents/xsgbfncq7of7) — 1h 41m
- `openai` [Elevated conversation errors for Free and Go users](https://status.openai.com/incidents/01M7Q99XRGP8T4DWJN8DSZXA1X) — 1h 58m
- `cloudflare` [Cache Purging Delays in Milan](https://www.cloudflarestatus.com/incidents/akjr7qn5o622) — 1h 33m
- `cloudflare` [Cloudflare Access Issue](https://www.cloudflarestatus.com/incidents/0jozucwnpzy0) — 1h 29m
- `cloudflare` [R2 Downloads Stalling in Japan](https://www.cloudflarestatus.com/incidents/2wjvsj2kqt07) — 1h 17m
- `cloudflare` [Network Performance Issues in Ashburn, US](https://www.cloudflarestatus.com/incidents/3l6q71x7a37k) — 1h 7m
- `cloudflare` [Durable Objects and D1 – Increased Errors](https://www.cloudflarestatus.com/incidents/5pe5mb9fefi8) — 31m
- `cloudflare` [Elevated number of 52x origin connectivity errors in Seoul](https://www.cloudflarestatus.com/incidents/1nqok7v173qg) — 1h 0m
- `cloudflare` [Single-File Cache Purge Delays](https://www.cloudflarestatus.com/incidents/5siv904vn94q) — 1h 0m
- `github` [Disruption with some GitHub services](https://stspg.io/pvjz2g01trc7) — 1h 11m
- `cloudflare` [Network performance issues in Los Angeles (LAX) and Frankfurt (FRA) datacenters](https://www.cloudflarestatus.com/incidents/du63c25qjgjs) — 47m
- `grafana` [Increased Read Latency In prod-us-central-0 Region](https://stspg.io/p6bxmc4m9p1v) — 1h 18m
- `cloudflare` [Elevated number of 5xx and 499 errors](https://www.cloudflarestatus.com/incidents/1ip2m5rvk7vb) — 40m
- `github` [Degraded availability for GPT-6 Luna in Copilot](https://stspg.io/vv0kzmyd9wqs) — 40m
- `cloudflare` [Elevated 5xx error rates and latency in Mumbai](https://www.cloudflarestatus.com/incidents/pi6wgeh78wsn) — 33m
- `cloudflare` [Network Performance Issues in Montreal, Canada](https://www.cloudflarestatus.com/incidents/wscmdyvmpb5d) — 30m
- `github` [Degraded Availability for MAI Code 2 Flash in Copilot](https://stspg.io/3sbmvyy2ctyx) — 34m
- `anthropic` [Elevated errors affecting the Google Docs connector](https://stspg.io/qz7tfhb7sl68) — 31m
- `grafana` [Data Source Permission Errors](https://stspg.io/32w5gjms7lm2) — 0m

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-f80b8064b3c18ea7c6686c26"></a>

- Canonical ID: `ci-f80b8064b3c18ea7c6686c26` (1 source records).
  - [grafana/x868z5qvdcnn](../../data/2026/W47/grafana.json); vendor URL: https://stspg.io/g8t0nj4lct2j
- Start time basis: `vendor_reported`.

<a id="ci-235a56445707110c95d29f41"></a>

- Canonical ID: `ci-235a56445707110c95d29f41` (1 source records).
  - [anthropic/b5cq662qf2wm](../../data/2026/W47/anthropic.json); vendor URL: https://stspg.io/01qq3kzyyqzl
- Start time basis: `vendor_reported`.

<a id="ci-76934f15b2aaa10e3f4290ba"></a>

- Canonical ID: `ci-76934f15b2aaa10e3f4290ba` (1 source records).
  - [openai/01M7N9VNR6X5SEC90M1YCDB36D](../../data/2026/W47/openai.json); vendor URL: https://status.openai.com/incidents/01M7N9VNR6X5SEC90M1YCDB36D
- Start time basis: `vendor_reported`.

<a id="ci-ed6c455e25909b9985a99948"></a>

- Canonical ID: `ci-ed6c455e25909b9985a99948` (1 source records).
  - [cloudflare/f33c51cw0ojd](../../data/2026/W47/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/f33c51cw0ojd
- Start time basis: `vendor_reported`.

<a id="ci-9574f530ec096ca780c018a2"></a>

- Canonical ID: `ci-9574f530ec096ca780c018a2` (1 source records).
  - [cloudflare/sv9tnfpjiojf](../../data/2026/W47/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/sv9tnfpjiojf
- Start time basis: `vendor_reported`.

<a id="ci-dc68d8b187ee618247c8134d"></a>

- Canonical ID: `ci-dc68d8b187ee618247c8134d` (1 source records).
  - [cloudflare/oz6to8yacrxe](../../data/2026/W47/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/oz6to8yacrxe
- Start time basis: `vendor_reported`.

<a id="ci-b6b7e2545b8368ef0ee80185"></a>

- Canonical ID: `ci-b6b7e2545b8368ef0ee80185` (1 source records).
  - [cloudflare/ljrvdciu7syf](../../data/2026/W47/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/ljrvdciu7syf
- Start time basis: `vendor_reported`.

<a id="ci-0630da281c5295644d89b2bc"></a>

- Canonical ID: `ci-0630da281c5295644d89b2bc` (1 source records).
  - [cloudflare/4lv21sf5fzhm](../../data/2026/W47/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/4lv21sf5fzhm
- Start time basis: `vendor_reported`.

<a id="ci-78b7d2af2e7959c04ad01870"></a>

- Canonical ID: `ci-78b7d2af2e7959c04ad01870` (1 source records).
  - [github/2zfqmghmg3s3](../../data/2026/W47/github.json); vendor URL: https://stspg.io/ljtn69jlt087
- Start time basis: `vendor_reported`.

<a id="ci-068fb03b92865405df02a4a7"></a>

- Canonical ID: `ci-068fb03b92865405df02a4a7` (1 source records).
  - [cloudflare/ylj5po752g6h](../../data/2026/W47/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/ylj5po752g6h
- Start time basis: `vendor_reported`.

<a id="ci-9c9bde87f65773819425e821"></a>

- Canonical ID: `ci-9c9bde87f65773819425e821` (1 source records).
  - [cloudflare/5831aw44ia7i](../../data/2026/W47/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/5831aw44ia7i
- Start time basis: `vendor_reported`.

<a id="ci-cc949bb74d1345ff2e38f59b"></a>

- Canonical ID: `ci-cc949bb74d1345ff2e38f59b` (1 source records).
  - [digitalocean/bdr1xgf08wsr](../../data/2026/W47/digitalocean.json); vendor URL: https://stspg.io/bnb8ltd5ct7l
- Start time basis: `vendor_reported`.

<a id="ci-eda13e3957c841cf04012d54"></a>

- Canonical ID: `ci-eda13e3957c841cf04012d54` (1 source records).
  - [cloudflare/njt9j3k2zdfw](../../data/2026/W47/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/njt9j3k2zdfw
- Start time basis: `vendor_reported`.

<a id="ci-6f5074bea4248258a8a899af"></a>

- Canonical ID: `ci-6f5074bea4248258a8a899af` (1 source records).
  - [anthropic/r1b4f9bbnsx1](../../data/2026/W47/anthropic.json); vendor URL: https://stspg.io/n18qjkd7dt4q
- Start time basis: `vendor_reported`.

<a id="ci-326a8282f949b4d0bca76b3f"></a>

- Canonical ID: `ci-326a8282f949b4d0bca76b3f` (1 source records).
  - [stripe/pg7ynt93z5sq](../../data/2026/W47/stripe.json); vendor URL: https://stspg.io/twy5139vhmjq
- Start time basis: `vendor_reported`.

<a id="ci-d2800ec50d70d4865ca0091e"></a>

- Canonical ID: `ci-d2800ec50d70d4865ca0091e` (1 source records).
  - [cloudflare/k3r9i5tbre4i](../../data/2026/W47/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/k3r9i5tbre4i
- Start time basis: `vendor_reported`.

<a id="ci-5e62a9b42ce5a6b25fe697dc"></a>

- Canonical ID: `ci-5e62a9b42ce5a6b25fe697dc` (1 source records).
  - [cloudflare/pfwuex7qix27](../../data/2026/W47/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/pfwuex7qix27
- Start time basis: `vendor_reported`.

<a id="ci-b69a0a04e1c3b196224efa3e"></a>

- Canonical ID: `ci-b69a0a04e1c3b196224efa3e` (1 source records).
  - [cloudflare/tawn3zx7wh0z](../../data/2026/W47/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/tawn3zx7wh0z
- Start time basis: `vendor_reported`.

<a id="ci-e95653f83490d9397a3121b6"></a>

- Canonical ID: `ci-e95653f83490d9397a3121b6` (1 source records).
  - [cloudflare/76lqmua6x3tk](../../data/2026/W47/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/76lqmua6x3tk
- Start time basis: `vendor_reported`.

<a id="ci-d994572f710fb0dc6f58db17"></a>

- Canonical ID: `ci-d994572f710fb0dc6f58db17` (1 source records).
  - [cloudflare/9zpa7own49hb](../../data/2026/W47/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/9zpa7own49hb
- Start time basis: `vendor_reported`.

<a id="ci-1c9cdf577e7af89ad91ce2b0"></a>

- Canonical ID: `ci-1c9cdf577e7af89ad91ce2b0` (1 source records).
  - [cloudflare/d2m4kw94ut4f](../../data/2026/W47/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/d2m4kw94ut4f
- Start time basis: `vendor_reported`.

<a id="ci-9e386d50157e8d852e3e1ebd"></a>

- Canonical ID: `ci-9e386d50157e8d852e3e1ebd` (1 source records).
  - [anthropic/1lglzf469c0z](../../data/2026/W47/anthropic.json); vendor URL: https://stspg.io/3tl6h35nnzwm
- Start time basis: `vendor_reported`.

<a id="ci-2d0f4c7401e3a1dcc97c23c2"></a>

- Canonical ID: `ci-2d0f4c7401e3a1dcc97c23c2` (1 source records).
  - [cloudflare/sd2sya0mc6df](../../data/2026/W47/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/sd2sya0mc6df
- Start time basis: `vendor_reported`.

<a id="ci-d175690cd304cc3bc7279620"></a>

- Canonical ID: `ci-d175690cd304cc3bc7279620` (1 source records).
  - [cloudflare/gcn9xop0f7zm](../../data/2026/W47/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/gcn9xop0f7zm
- Start time basis: `vendor_reported`.

<a id="ci-e0fa6077501a06ff1050d939"></a>

- Canonical ID: `ci-e0fa6077501a06ff1050d939` (1 source records).
  - [circleci/s501b363k4hs](../../data/2026/W47/circleci.json); vendor URL: https://stspg.io/g20zht6r8c4w
- Start time basis: `vendor_reported`.

<a id="ci-38028a34c21a46885c5e6e09"></a>

- Canonical ID: `ci-38028a34c21a46885c5e6e09` (1 source records).
  - [cloudflare/69855o2zz6kl](../../data/2026/W47/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/69855o2zz6kl
- Start time basis: `vendor_reported`.

<a id="ci-56397c5ff27555054dc3f7e1"></a>

- Canonical ID: `ci-56397c5ff27555054dc3f7e1` (1 source records).
  - [cloudflare/xsgbfncq7of7](../../data/2026/W47/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/xsgbfncq7of7
- Start time basis: `vendor_reported`.

<a id="ci-cbad4ebdd0fcf5bf8e2c1602"></a>

- Canonical ID: `ci-cbad4ebdd0fcf5bf8e2c1602` (1 source records).
  - [openai/01M7Q99XRGP8T4DWJN8DSZXA1X](../../data/2026/W47/openai.json); vendor URL: https://status.openai.com/incidents/01M7Q99XRGP8T4DWJN8DSZXA1X
- Start time basis: `vendor_reported`.

<a id="ci-78d79423a88b61cf283b9728"></a>

- Canonical ID: `ci-78d79423a88b61cf283b9728` (1 source records).
  - [cloudflare/akjr7qn5o622](../../data/2026/W47/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/akjr7qn5o622
- Start time basis: `vendor_reported`.

<a id="ci-be077a670bb1f19ec3d3afc4"></a>

- Canonical ID: `ci-be077a670bb1f19ec3d3afc4` (1 source records).
  - [cloudflare/0jozucwnpzy0](../../data/2026/W47/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/0jozucwnpzy0
- Start time basis: `vendor_reported`.

<a id="ci-0ac3b5acc951515826915c90"></a>

- Canonical ID: `ci-0ac3b5acc951515826915c90` (1 source records).
  - [circleci/vlz24tmd3tpd](../../data/2026/W47/circleci.json); vendor URL: https://stspg.io/m8bpg1brfjcw
- Start time basis: `vendor_reported`.

<a id="ci-59aa64b298302d34b1112597"></a>

- Canonical ID: `ci-59aa64b298302d34b1112597` (1 source records).
  - [cloudflare/2wjvsj2kqt07](../../data/2026/W47/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/2wjvsj2kqt07
- Start time basis: `vendor_reported`.

<a id="ci-0ef84a7a30c5464ef7fdc5fb"></a>

- Canonical ID: `ci-0ef84a7a30c5464ef7fdc5fb` (1 source records).
  - [anthropic/c1kmrkmv5f3d](../../data/2026/W47/anthropic.json); vendor URL: https://stspg.io/nlb5fptv28h5
- Start time basis: `vendor_reported`.

<a id="ci-467c9bc1b353b62afa5210d9"></a>

- Canonical ID: `ci-467c9bc1b353b62afa5210d9` (1 source records).
  - [cloudflare/3l6q71x7a37k](../../data/2026/W47/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/3l6q71x7a37k
- Start time basis: `vendor_reported`.

<a id="ci-46eb7eb190a3535f323d85ea"></a>

- Canonical ID: `ci-46eb7eb190a3535f323d85ea` (1 source records).
  - [digitalocean/fdvp47jv22nl](../../data/2026/W47/digitalocean.json); vendor URL: https://stspg.io/4ccggg9c514v
- Start time basis: `vendor_reported`.

<a id="ci-b32f9b2dea89a7ebb570fbee"></a>

- Canonical ID: `ci-b32f9b2dea89a7ebb570fbee` (1 source records).
  - [cloudflare/5pe5mb9fefi8](../../data/2026/W47/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/5pe5mb9fefi8
- Start time basis: `vendor_reported`.

<a id="ci-0ff21d4640d2f770d62461a4"></a>

- Canonical ID: `ci-0ff21d4640d2f770d62461a4` (1 source records).
  - [cloudflare/1nqok7v173qg](../../data/2026/W47/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/1nqok7v173qg
- Start time basis: `vendor_reported`.

<a id="ci-bf581fe96f76c1f311de35b2"></a>

- Canonical ID: `ci-bf581fe96f76c1f311de35b2` (1 source records).
  - [cloudflare/5siv904vn94q](../../data/2026/W47/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/5siv904vn94q
- Start time basis: `vendor_reported`.

<a id="ci-8aa142091f1ba47f8b520d8b"></a>

- Canonical ID: `ci-8aa142091f1ba47f8b520d8b` (1 source records).
  - [github/m310qbkfm3ww](../../data/2026/W47/github.json); vendor URL: https://stspg.io/pvjz2g01trc7
- Start time basis: `vendor_reported`.

<a id="ci-4f57319de98ac6ddb725f3ab"></a>

- Canonical ID: `ci-4f57319de98ac6ddb725f3ab` (1 source records).
  - [cloudflare/du63c25qjgjs](../../data/2026/W47/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/du63c25qjgjs
- Start time basis: `vendor_reported`.

<a id="ci-eadeb781439c44861c99f3ec"></a>

- Canonical ID: `ci-eadeb781439c44861c99f3ec` (1 source records).
  - [openai/01M7T8MRF9SDBDNQ62Z6V156ZZ](../../data/2026/W47/openai.json); vendor URL: https://status.openai.com/incidents/01M7T8MRF9SDBDNQ62Z6V156ZZ
- Start time basis: `vendor_reported`.

<a id="ci-09d3dac6e2e96cb8789730cf"></a>

- Canonical ID: `ci-09d3dac6e2e96cb8789730cf` (1 source records).
  - [grafana/5n4xwywr0dmn](../../data/2026/W47/grafana.json); vendor URL: https://stspg.io/p6bxmc4m9p1v
- Start time basis: `vendor_reported`.

<a id="ci-90a8861e05e222a4b1c8bab7"></a>

- Canonical ID: `ci-90a8861e05e222a4b1c8bab7` (1 source records).
  - [cloudflare/1ip2m5rvk7vb](../../data/2026/W47/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/1ip2m5rvk7vb
- Start time basis: `vendor_reported`.

<a id="ci-d3263739af6388d82261a4b0"></a>

- Canonical ID: `ci-d3263739af6388d82261a4b0` (1 source records).
  - [github/gp7nhc88bwv9](../../data/2026/W47/github.json); vendor URL: https://stspg.io/vv0kzmyd9wqs
- Start time basis: `vendor_reported`.

<a id="ci-b0a47e07d3d726659fc2fd78"></a>

- Canonical ID: `ci-b0a47e07d3d726659fc2fd78` (1 source records).
  - [cloudflare/pi6wgeh78wsn](../../data/2026/W47/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/pi6wgeh78wsn
- Start time basis: `vendor_reported`.

<a id="ci-1eb588c115909e32b11c80f2"></a>

- Canonical ID: `ci-1eb588c115909e32b11c80f2` (1 source records).
  - [cloudflare/wscmdyvmpb5d](../../data/2026/W47/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/wscmdyvmpb5d
- Start time basis: `vendor_reported`.

<a id="ci-46d64af39ca9964f5096ad68"></a>

- Canonical ID: `ci-46d64af39ca9964f5096ad68` (1 source records).
  - [github/891rm8181xdf](../../data/2026/W47/github.json); vendor URL: https://stspg.io/3sbmvyy2ctyx
- Start time basis: `vendor_reported`.

<a id="ci-35b755c4579076a68a1f64a6"></a>

- Canonical ID: `ci-35b755c4579076a68a1f64a6` (1 source records).
  - [cloudflare/iysf3fasf7wn](../../data/2026/W47/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/iysf3fasf7wn
- Start time basis: `vendor_reported`.

<a id="ci-d59428cbc043fdf46acdf568"></a>

- Canonical ID: `ci-d59428cbc043fdf46acdf568` (1 source records).
  - [anthropic/6prv10gjdpz1](../../data/2026/W47/anthropic.json); vendor URL: https://stspg.io/qz7tfhb7sl68
- Start time basis: `vendor_reported`.

<a id="ci-8ce6a936273c3312a6f0ff4a"></a>

- Canonical ID: `ci-8ce6a936273c3312a6f0ff4a` (1 source records).
  - [grafana/wfkrzws94vfd](../../data/2026/W47/grafana.json); vendor URL: https://stspg.io/32w5gjms7lm2
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
| slack | ok | 26 | 0 |
| stripe | ok | 50 | 0 |
| yandexcloud | FETCH FAILED: GET https://status.cloud.yandex.ru/api/incidents?lang=en&page=4: status 403 | 0 | 0 |
