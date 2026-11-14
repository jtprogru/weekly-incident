# Week 2026-W30

2026-07-20 to 2026-07-26 (UTC). Generated 2026-11-14 12:36.

60 canonical incidents, 32 above threshold. 2 of 14 sources unavailable: datadog, yandexcloud.

60 source records (61 saved revisions) → 60 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | grafana | [Stacks using SCIM user provisioning are currently unable to log into Grafana](https://stspg.io/93c20d1dzpdp) | 2026-07-23 10:13 | 5h 24m | major | AWS Australia - prod-ap-southeast-2, AWS Australia - prod-au-southeast-1, AWS Brazil - prod-sa-east-1, AWS Canada - prod-ca-east-0, AWS Germany - prod-eu-west-2, AWS Germany - prod-eu-west-4, AWS India - prod-ap-south-1, AWS Ireland - prod-eu-west-6, AWS Japan - prod-ap-northeast-0, AWS Singapore - prod-ap-southeast-1, AWS Sweden - prod-eu-north-0, AWS Switzerland - prod-eu-central-0, AWS UAE - prod-me-central-1, AWS UK - prod-gb-south-1, AWS US East - prod-us-east-0, AWS US East - prod-us-east-2, AWS US West - prod-us-west-0, Azure Netherlands - prod-eu-west-3, Azure US Central - us-central2, Federal Cloud - AWS US Gov West, GCP Australia - prod-au-southeast-0, GCP Belgium - prod-eu-west-0, GCP Brazil - prod-sa-east-0, GCP India - prod-ap-south-0, GCP Singapore - prod-ap-southeast-0, GCP UK - prod-gb-south-0, GCP US Central - prod-us-central-0, GCP US Central - prod-us-central-3, GCP US Central - prod-us-central-4, GCP US East - prod-us-east-1, play.grafana.org | 10044 = 324 × 31 × 1.0 |
| 2 | anthropic | [Microsoft Office add-in availability](https://stspg.io/y411l9qwm1r8) | 2026-07-22 17:56 | 1d 21h 40m | none | — | 4110 = 2740 × 1 × 1.5 |
| 3 | openai | [Elevated errors affecting ChatGPT conversations](https://status.openai.com/incidents/01KYDN6YPS6ARY1EC9089N089G) | 2026-07-25 22:09 | 1d 18h 22m | minor | — | 3813 = 2542 × 1 × 1.5 |
| 4 | openai | [Elevated Error Rates](https://status.openai.com/incidents/01KY7SX5MYJ2BP51X5MXAPYX71) | 2026-07-23 15:36 | 1d 3h 57m | minor | — | 2516 = 1677 × 1 × 1.5 |
| 5 | grafana | [Cloud Log Exporter Unavailable](https://stspg.io/cjjmv6qvg1wl) | 2026-07-21 18:43 | 2h 22m | major | AWS Brazil - prod-sa-east-1, AWS Canada - prod-ca-east-0, AWS Germany - prod-eu-west-2, AWS India - prod-ap-south-1, AWS Japan - prod-ap-northeast-0, AWS Singapore - prod-ap-southeast-1, AWS Switzerland - prod-eu-central-0, AWS US West - prod-us-west-0, Azure Netherlands - prod-eu-west-3, Azure US Central - us-central7, GCP Belgium - prod-eu-west-0, GCP Brazil - prod-sa-east-0, GCP India - prod-ap-south-0, GCP Singapore - prod-ap-southeast-0, GCP UK - prod-gb-south-0 | 2130 = 142 × 15 × 1.0 |
| 6 | anthropic | [Elevated error rates for Opus 4.5](https://stspg.io/3hc27lgyzhcf) | 2026-07-20 07:03 | 5h 44m | minor | Claude API (api.anthropic.com), Claude Code, Claude Console (platform.claude.com), Claude Cowork | 2064 = 344 × 4 × 1.5 |
| 7 | anthropic | [Elevated errors on Haiku 4.5](https://stspg.io/twj0mkj40p93) | 2026-07-21 13:04 | 3h 33m | major | Claude API (api.anthropic.com), Claude Code, Claude Console (platform.claude.com), Claude Cowork, claude.ai | 1598 = 213 × 5 × 1.5 |
| 8 | anthropic | [Elevated errors on Haiku 4.5](https://stspg.io/lqg8z33rmjz9) | 2026-07-20 13:20 | 2h 20m | minor | Claude API (api.anthropic.com), Claude Code, Claude Console (platform.claude.com), Claude Cowork, claude.ai | 1050 = 140 × 5 × 1.5 |
| 9 | anthropic | [Elevated errors on several models](https://stspg.io/ltgjlwc77f49) | 2026-07-21 15:35 | 2h 43m | critical | Claude API (api.anthropic.com), Claude Code, Claude Cowork, claude.ai | 978 = 163 × 4 × 1.5 |
| 10 | anthropic | [Fable 5 requiring usage credits on Max plans](https://stspg.io/r596yvz5clzv) | 2026-07-20 07:35 | 8h 0m | minor | Claude Code | 720 = 480 × 1 × 1.5 |
| 11 | github | [Disruption with some GitHub services](https://stspg.io/j5c80shxqm53) | 2026-07-24 16:17 | 1h 19m | major | API Requests, Actions, Copilot, Issues, Pages, Pull Requests | 711 = 79 × 6 × 1.5 |
| 12 | anthropic | [Elevated errors for Opus 5](https://stspg.io/sx0f2mwcqw7h) | 2026-07-26 09:17 | 1h 27m | major | Claude API (api.anthropic.com), Claude Code, Claude Console (platform.claude.com), Claude Cowork, claude.ai | 652 = 87 × 5 × 1.5 |
| 13 | github | [Latency issues across a number of services](https://stspg.io/syhr80rth84z) | 2026-07-23 07:53 | 1h 45m | major | Actions, Issues, Pull Requests, Webhooks | 630 = 105 × 4 × 1.5 |
| 14 | grafana | [Intermittent write outage from 21:20-21:26, 21:54-21:56, and 22:22-22:23](https://stspg.io/xf7vsp1rv6yk) | 2026-07-23 23:00 | 6h 31m | none | — | 391 = 391 × 1 × 1.0 |
| 15 | anthropic | [Elevated errors for Mythos 5, Fable 5, Opus 5 and Claude Haiku 4.5](https://stspg.io/bkqqrx58hq88) | 2026-07-25 18:40 | 1h 4m | major | Claude API (api.anthropic.com), Claude Code, Claude Cowork, claude.ai | 384 = 64 × 4 × 1.5 |
| 16 | circleci | [Delays in starting jobs using the gen3 resource class](https://stspg.io/jjsrf04jn924) | 2026-07-22 19:57 | 6h 8m | minor | Machine Jobs | 368 = 368 × 1 × 1.0 |
| 17 | grafana | [Errors creating new Slack integration for Grafana IRM](https://stspg.io/w78r22dh774l) | 2026-07-24 09:17 | 4h 45m | major | Incident Management and Response (IRM) | 285 = 285 × 1 × 1.0 |
| 18 | openai | [Some users may experience elevated error rates in ChatGPT](https://status.openai.com/incidents/01KY81XMFAV8KXNJGVNEMKAK15) | 2026-07-23 17:56 | 2h 44m | minor | — | 246 = 164 × 1 × 1.5 |
| 19 | github | [Disruption with GPT 5.3 Codex](https://stspg.io/7ls47qxt8lpq) | 2026-07-20 16:03 | 2h 33m | minor | Copilot AI Model Providers | 230 = 153 × 1 × 1.5 |
| 20 | anthropic | [Elevated errors for Claude Fable 5, Claude Sonnet 5, Claude Haiku 4.5, and other models](https://stspg.io/4l8l9dtyx58j) | 2026-07-25 21:34 | 33m | major | Claude API (api.anthropic.com), Claude Code, Claude Cowork, claude.ai | 198 = 33 × 4 × 1.5 |
| 21 | slack | [Issue sending messages to channels via certain Workflows](https://slack-status.com/2026-07/e82ee127d8bed735) | 2026-07-24 08:22 | 2h 24m | unknown | Workflows | 173 = 144 × 1 × 1.2 |
| 22 | grafana | [Metrics Write Path Errors](https://stspg.io/l59vv2xfq9fx) | 2026-07-23 17:14 | 2h 33m | major | AWS US East - prod-us-east-0: Ingestion | 153 = 153 × 1 × 1.0 |
| 23 | grafana | [PDC Issues](https://stspg.io/vxd316w7jsys) | 2026-07-20 13:48 | 1h 15m | major | AWS Germany - prod-eu-west-2, AWS US East - prod-us-east-2 | 150 = 75 × 2 × 1.0 |
| 24 | digitalocean | [Gradient AI Serverless Inference Requests Timing Out for qwen3.5-397b-a17b](https://stspg.io/2fdnm4tns3f1) | 2026-07-22 16:50 | 2h 51m | minor | Inference | 137 = 171 × 1 × 0.8 |
| 25 | github | [Some SSH connections using deploy keys are failing](https://stspg.io/jczxclgm9znm) | 2026-07-21 10:31 | 1h 25m | critical | — | 128 = 85 × 1 × 1.5 |
| 26 | github | [Disruption with some GitHub services](https://stspg.io/2n2g6kzg67mh) | 2026-07-20 00:25 | 1h 20m | major | — | 120 = 80 × 1 × 1.5 |
| 27 | digitalocean | [Network Connectivity from India to NYC](https://stspg.io/y9g7qgc9qs6m) | 2026-07-24 10:36 | 2h 5m | minor | NYC1 | 100 = 125 × 1 × 0.8 |
| 28 | grafana | [Write Outage](https://stspg.io/z8r7rgprs19d) | 2026-07-24 17:02 | 1h 18m | critical | AWS US East - prod-us-east-3 | 78 = 78 × 1 × 1.0 |
| 29 | anthropic | [Service disruption on Claude services](https://stspg.io/d8v0g0cntyjx) | 2026-07-21 17:40 | 23m | critical | Claude Code, Claude Cowork | 69 = 23 × 2 × 1.5 |
| 30 | github | [Incident with Pull Requests](https://stspg.io/sm1tp7kfm4vj) | 2026-07-24 19:37 | 45m | critical | Pull Requests | 68 = 45 × 1 × 1.5 |
| 31 | github | [Actions run failures and delays](https://stspg.io/448g37mrq066) | 2026-07-25 12:31 | 41m | critical | Actions | 62 = 41 × 1 × 1.5 |
| 32 | github | [Several GPT models degraded](https://stspg.io/vh0xxw69dr6v) | 2026-07-25 09:42 | 28m | major | Copilot AI Model Providers | 42 = 28 × 1 × 1.5 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. grafana — Stacks using SCIM user provisioning are currently unable to log into Grafana

- URL: https://stspg.io/93c20d1dzpdp
- Started: 2026-07-23 10:13 UTC
- Resolved: 2026-07-23 15:37 UTC
- Duration: 5h 24m
- Impact: major, status: resolved
- Components: AWS Australia - prod-ap-southeast-2, AWS Australia - prod-au-southeast-1, AWS Brazil - prod-sa-east-1, AWS Canada - prod-ca-east-0, AWS Germany - prod-eu-west-2, AWS Germany - prod-eu-west-4, AWS India - prod-ap-south-1, AWS Ireland - prod-eu-west-6, AWS Japan - prod-ap-northeast-0, AWS Singapore - prod-ap-southeast-1, AWS Sweden - prod-eu-north-0, AWS Switzerland - prod-eu-central-0, AWS UAE - prod-me-central-1, AWS UK - prod-gb-south-1, AWS US East - prod-us-east-0, AWS US East - prod-us-east-2, AWS US West - prod-us-west-0, Azure Netherlands - prod-eu-west-3, Azure US Central - us-central2, Federal Cloud - AWS US Gov West, GCP Australia - prod-au-southeast-0, GCP Belgium - prod-eu-west-0, GCP Brazil - prod-sa-east-0, GCP India - prod-ap-south-0, GCP Singapore - prod-ap-southeast-0, GCP UK - prod-gb-south-0, GCP US Central - prod-us-central-0, GCP US Central - prod-us-central-3, GCP US Central - prod-us-central-4, GCP US East - prod-us-east-1, play.grafana.org
- Canonical incident: [ci-ecf9d6bdd2eba01d8cde5c68](#ci-ecf9d6bdd2eba01d8cde5c68)

- `2026-07-23 10:13` **investigating** — We are currently facing an issue where Stacks using SCIM user provisioning are currently unable to log into Grafana. We are currently investigating this issue and working on a fix.
- `2026-07-23 11:26` **identified** — The issue has been identified and we are currently working on the mitigations now. The scope is much more limited than initially thought only specific SCIM configurations are impacted.
- `2026-07-23 13:23` **identified** — We have applied a fix and are currently awaiting feedback from affected customers.
- `2026-07-23 15:37` **resolved** — This incident has been resolved.

### 2. anthropic — Microsoft Office add-in availability

- URL: https://stspg.io/y411l9qwm1r8
- Started: 2026-07-22 17:56 UTC
- Resolved: 2026-07-24 15:36 UTC
- Duration: 1d 21h 40m
- Impact: none, status: resolved
- Components: —
- Canonical incident: [ci-73994a4f2afecc8a2699b70b](#ci-73994a4f2afecc8a2699b70b)

- `2026-07-22 17:56` **identified** — Users who installed the following add-ins may experience issues accessing the add-in functionality: Claude by Anthropic for PowerPoint, Claude by Anthropic for Word, and Claude by Anthropic for Excel. To regain immediate access to Claude for Microsoft 365, install the add-in here: https://marketplace.microsoft.com/en-us/product/office/wa200010725
- `2026-07-24 15:36` **resolved** — This issue has been resolved.

### 3. openai — Elevated errors affecting ChatGPT conversations

- URL: https://status.openai.com/incidents/01KYDN6YPS6ARY1EC9089N089G
- Started: 2026-07-25 22:09 UTC
- Resolved: 2026-07-27 16:32 UTC
- Duration: 1d 18h 22m
- Impact: minor, status: resolved
- Components: —
- Canonical incident: [ci-1aa02afa33e6459e6e254fa7](#ci-1aa02afa33e6459e6e254fa7)

- `2026-07-25 22:09` **investigating** — We are investigating intermittent errors that may prevent some users from loading or continuing conversations in ChatGPT. The current period of impact began around 1:00 PM PT.
- `2026-07-25 23:16` **identified** — We have identified the source of intermittent errors affecting some ChatGPT conversations and are working to mitigate the issue. Error rates have decreased, but some users may continue to experience failures.
- `2026-07-25 23:57` **monitoring** — Mitigation has been implemented, and we are monitoring the recovery of ChatGPT conversations.
- `2026-07-27 16:32` **resolved** — All impacted services have now fully recovered.

### 4. openai — Elevated Error Rates

- URL: https://status.openai.com/incidents/01KY7SX5MYJ2BP51X5MXAPYX71
- Started: 2026-07-23 15:36 UTC
- Resolved: 2026-07-24 19:33 UTC
- Duration: 1d 3h 57m
- Impact: minor, status: resolved
- Components: —
- Canonical incident: [ci-0e74a35a4681977d515bc24d](#ci-0e74a35a4681977d515bc24d)

- `2026-07-23 15:36` **investigating** — We are investigating the issue for the listed services.
- `2026-07-23 15:56` **identified** — We have identified that users are experiencing elevated errors for the impacted services. We are working on implementing a mitigation.
- `2026-07-23 16:08` **identified** — We have identified that users are experiencing elevated errors for the impacted services. We are working on implementing a mitigation.
- `2026-07-23 16:48` **identified** — We have identified that users are experiencing elevated errors for the impacted services. We are working on implementing a mitigation with the downstream infrastructure provider.
- `2026-07-23 17:36` **identified** — We are continuing to work on implementing a mitigation.
- `2026-07-23 17:58` **identified** — We are continuing to work on implementing a mitigation.
- `2026-07-23 18:50` **monitoring** — We have applied the mitigation and are monitoring the recovery.
- `2026-07-23 19:37` **investigating** — We are investigating the issue for the listed services.
- `2026-07-23 20:22` **monitoring** — We have applied the mitigation and are monitoring the recovery.
- `2026-07-24 09:56` **investigating** — We are investigating the issue for the listed services.
- `2026-07-24 10:06` **investigating** — We are investigating the issue for the listed services.
- `2026-07-24 11:34` **monitoring** — We have applied the mitigation and are monitoring the recovery.
- `2026-07-24 19:33` **resolved** — All impacted services have now fully recovered.

### 5. grafana — Cloud Log Exporter Unavailable

- URL: https://stspg.io/cjjmv6qvg1wl
- Started: 2026-07-21 18:43 UTC
- Resolved: 2026-07-21 21:06 UTC
- Duration: 2h 22m
- Impact: major, status: resolved
- Components: AWS Brazil - prod-sa-east-1, AWS Canada - prod-ca-east-0, AWS Germany - prod-eu-west-2, AWS India - prod-ap-south-1, AWS Japan - prod-ap-northeast-0, AWS Singapore - prod-ap-southeast-1, AWS Switzerland - prod-eu-central-0, AWS US West - prod-us-west-0, Azure Netherlands - prod-eu-west-3, Azure US Central - us-central7, GCP Belgium - prod-eu-west-0, GCP Brazil - prod-sa-east-0, GCP India - prod-ap-south-0, GCP Singapore - prod-ap-southeast-0, GCP UK - prod-gb-south-0
- Canonical incident: [ci-4e1351d93fe3dc96ed255554](#ci-4e1351d93fe3dc96ed255554)

- `2026-07-21 18:43` **investigating** — Cloud Log Exporter is temporarily not available in the marked regions. We are currently investigating this issue and will update as soon as we have more info to share.
- `2026-07-21 20:04` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-07-21 21:06` **resolved** — This incident has been resolved.

### 6. anthropic — Elevated error rates for Opus 4.5

- URL: https://stspg.io/3hc27lgyzhcf
- Started: 2026-07-20 07:03 UTC
- Resolved: 2026-07-20 12:48 UTC
- Duration: 5h 44m
- Impact: minor, status: resolved
- Components: Claude API (api.anthropic.com), Claude Code, Claude Console (platform.claude.com), Claude Cowork
- Canonical incident: [ci-6e7020ec5d38bc05489c666c](#ci-6e7020ec5d38bc05489c666c)

- `2026-07-20 07:03` **investigating** — We are currently investigating this issue.
- `2026-07-20 07:57` **identified** — The issue has been identified and a fix is being implemented.
- `2026-07-20 10:05` **identified** — We are continuing to work on a fix for this issue.
- `2026-07-20 10:57` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-07-20 12:48` **resolved** — This incident has been resolved.

### 7. anthropic — Elevated errors on Haiku 4.5

- URL: https://stspg.io/twj0mkj40p93
- Started: 2026-07-21 13:04 UTC
- Resolved: 2026-07-21 16:38 UTC
- Duration: 3h 33m
- Impact: major, status: resolved
- Components: Claude API (api.anthropic.com), Claude Code, Claude Console (platform.claude.com), Claude Cowork, claude.ai
- Canonical incident: [ci-b254336524302a9d9c36cdbb](#ci-b254336524302a9d9c36cdbb)

- `2026-07-21 13:04` **investigating** — We are currently investigating this issue.
- `2026-07-21 13:14` **identified** — The issue has been identified and a fix is being implemented.
- `2026-07-21 13:22` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-07-21 14:44` **identified** — The issue has been identified and a fix is being implemented.
- `2026-07-21 15:30` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-07-21 16:38` **resolved** — This incident has been resolved.

### 8. anthropic — Elevated errors on Haiku 4.5

- URL: https://stspg.io/lqg8z33rmjz9
- Started: 2026-07-20 13:20 UTC
- Resolved: 2026-07-20 15:41 UTC
- Duration: 2h 20m
- Impact: minor, status: resolved
- Components: Claude API (api.anthropic.com), Claude Code, Claude Console (platform.claude.com), Claude Cowork, claude.ai
- Canonical incident: [ci-78ae2d5a9a030ef43adcb90e](#ci-78ae2d5a9a030ef43adcb90e)

- `2026-07-20 13:20` **investigating** — We are currently investigating this issue.
- `2026-07-20 15:09` **identified** — The issue has been identified and a fix is being implemented.
- `2026-07-20 15:33` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-07-20 15:41` **resolved** — This incident has been resolved.

### 9. anthropic — Elevated errors on several models

- URL: https://stspg.io/ltgjlwc77f49
- Started: 2026-07-21 15:35 UTC
- Resolved: 2026-07-21 18:18 UTC
- Duration: 2h 43m
- Impact: critical, status: resolved
- Components: Claude API (api.anthropic.com), Claude Code, Claude Cowork, claude.ai
- Canonical incident: [ci-f793d1cf9e56471fae5d5b28](#ci-f793d1cf9e56471fae5d5b28)

- `2026-07-21 15:35` **investigating** — We are currently investigating this issue.
- `2026-07-21 15:51` **identified** — The issue has been identified and a fix is being implemented.
- `2026-07-21 17:07` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-07-21 18:18` **resolved** — From 8:28am PT / 15:28 UTC through 9:26am PT / 16:26 UTC users experienced elevated errors on requests to multiple Claude models. This issue has been resolved.

### 10. anthropic — Fable 5 requiring usage credits on Max plans

- URL: https://stspg.io/r596yvz5clzv
- Started: 2026-07-20 07:35 UTC
- Resolved: 2026-07-20 15:35 UTC
- Duration: 8h 0m
- Impact: minor, status: resolved
- Components: Claude Code
- Canonical incident: [ci-c5227a9ce8abc98311bd9e27](#ci-c5227a9ce8abc98311bd9e27)

- `2026-07-20 07:35` **monitoring** — Some Claude Code Max plan users were incorrectly prompted to use usage credits to access Fable 5, which is included in their plan. If you were affected, please restart Claude Code.
- `2026-07-20 15:35` **resolved** — This incident has been resolved.

### 11. github — Disruption with some GitHub services

- URL: https://stspg.io/j5c80shxqm53
- Started: 2026-07-24 16:17 UTC
- Resolved: 2026-07-24 17:36 UTC
- Duration: 1h 19m
- Impact: major, status: resolved
- Components: API Requests, Actions, Copilot, Issues, Pages, Pull Requests
- Canonical incident: [ci-009cba4a402059471b65f57c](#ci-009cba4a402059471b65f57c)

- `2026-07-24 16:17` **investigating** — We are investigating reports of degraded performance for API Requests and Issues
- `2026-07-24 16:19` **investigating** — Actions is experiencing degraded performance. We are continuing to investigate.
- `2026-07-24 16:20` **investigating** — Pull Requests is experiencing degraded performance. We are continuing to investigate.
- `2026-07-24 16:22` **investigating** — We are investigating timeouts to some GitHub services
- `2026-07-24 16:26` **investigating** — Copilot is experiencing degraded performance. We are continuing to investigate.
- `2026-07-24 16:27` **investigating** — Pages is experiencing degraded performance. We are continuing to investigate.
- `2026-07-24 16:28` **investigating** — Actions is experiencing degraded availability. We are continuing to investigate.
- `2026-07-24 16:40` **investigating** — We have applied a mitigation and are monitoring for recovery
- `2026-07-24 16:41` **investigating** — Actions is experiencing degraded performance. We are continuing to investigate.
- `2026-07-24 17:16` **investigating** — The degradation affecting API Requests, Actions, Copilot, Issues, Pages and Pull Requests has been mitigated. We are monitoring to ensure stability.
- `2026-07-24 17:24` **investigating** — We are seeing recovery across all services
- `2026-07-24 17:36` **resolved** — On July 24th at 16:04 UTC, a loss of connectivity occurred in network paths in one of our three physical data center availability zones (AZs). This resulted in packet loss due to the remaining active paths becoming saturated. Our data centers use a leaf-spine switch fabric in each compute cage, and an aggregation layer interconnecting the spines from each cage within each AZ. The loss of connectivity affected links between one cage’s spine switches and the aggregation layer within that specific AZ. <br /><br />Workloads depending on compute resources in this cage became degraded due to packet loss, and exhibited intermittent errors: <br /><br />- Actions saw 10% of jobs fail during the impact window, and 5% of jobs succeeded but with delayed starts. <br />- 27% of GitHub issues interactions saw slow requests or timeouts. <br />- 4% of GitHub Copilot requests experienced errors, though most automatically retry. <br />- 4% of git push operations saw impacts during the affected window. <br />- Authentication requests saw increased latency during the affected window, but error rates, while elevated, were < 1% in all cases. <br /><br />We were able to mitigate the outage by re-routing affected connections to available fiber paths that were allocated for future capacity upgrades. Sufficient network capacity to eliminate packet loss was restored at 17:07, with most services showing full recovery by 17:16. All paths were restored and services healthy at 17:36. <br /><br />This incident affected 25% of available network interconnect capacity. Older cages utilize a 100Gbps network interface standard. To remove risk of reoccurrence, a planned upgrade to 400Gbps interfaces is being accelerated as much as possible, ensuring increased bandwidth available at all layers of the switch fabric for resiliency to path or device loss.

### 12. anthropic — Elevated errors for Opus 5

- URL: https://stspg.io/sx0f2mwcqw7h
- Started: 2026-07-26 09:17 UTC
- Resolved: 2026-07-26 10:44 UTC
- Duration: 1h 27m
- Impact: major, status: resolved
- Components: Claude API (api.anthropic.com), Claude Code, Claude Console (platform.claude.com), Claude Cowork, claude.ai
- Canonical incident: [ci-4b10c3920f80eefc8e3fd34a](#ci-4b10c3920f80eefc8e3fd34a)

- `2026-07-26 09:17` **investigating** — We are currently investigating this issue.
- `2026-07-26 09:45` **identified** — The issue has been identified and a fix is being implemented.
- `2026-07-26 10:28` **identified** — We are continuing to work on a fix for this issue.
- `2026-07-26 10:34` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-07-26 10:44` **resolved** — This incident has been resolved.

### 13. github — Latency issues across a number of services

- URL: https://stspg.io/syhr80rth84z
- Started: 2026-07-23 07:53 UTC
- Resolved: 2026-07-23 09:39 UTC
- Duration: 1h 45m
- Impact: major, status: resolved
- Components: Actions, Issues, Pull Requests, Webhooks
- Canonical incident: [ci-05543d9dec1918885cbe7394](#ci-05543d9dec1918885cbe7394)

- `2026-07-23 07:53` **investigating** — We are investigating reports of degraded availability for Actions, Issues and Webhooks
- `2026-07-23 08:25` **investigating** — Pull Requests is experiencing degraded performance. We are continuing to investigate.
- `2026-07-23 08:34` **investigating** — We're currently investigating latency across multiple services. This can show as Actions jobs taking longer to start, Issues search serving stale results, and other listed services being similarly impacted.
- `2026-07-23 09:18` **investigating** — The degradation affecting Issues has been mitigated. We are monitoring to ensure stability.
- `2026-07-23 09:19` **investigating** — The degradation affecting Actions has been mitigated. We are monitoring to ensure stability.
- `2026-07-23 09:22` **investigating** — We identified the source of latency affecting multiple services and applied a fix. Issues and Actions are recovering, and remaining affected services are seeing improvement as processing backlogs clear. We are actively monitoring recovery across all services.
- `2026-07-23 09:27` **investigating** — The degradation affecting Pull Requests has been mitigated. We are monitoring to ensure stability.
- `2026-07-23 09:35` **investigating** — Webhooks is operating normally.
- `2026-07-23 09:39` **monitoring** — The degradation has been mitigated. We are monitoring to ensure stability.
- `2026-07-23 09:39` **resolved** — On July 23, 2026, between 07:08 and 09:39 UTC, several services experienced delays: 8% of actions workflow runs experienced an average run start delay of 10 minutes, 5% of webhook deliveries exceeded SLO, and code scanning, repos, notifications, issues and pull requests experienced increased latency over the life of the incident. <br /><br />The root cause of the incident was a node of our background job processing system which did not recover after entering scheduled host maintenance. The incident was mitigated by identifying the problematic shard and restoring its correct state, after which queue backlogs drained and services recovered. <br /><br />To speed mitigation, we have added monitors for nodes in this unhealthy state after maintenance operations. To prevent future recurrence, we are adapting our lifecycle automation to verify host rejoin after a scheduled reboot.

### 14. grafana — Intermittent write outage from 21:20-21:26, 21:54-21:56, and 22:22-22:23

- URL: https://stspg.io/xf7vsp1rv6yk
- Started: 2026-07-23 23:00 UTC
- Resolved: 2026-07-24 05:32 UTC
- Duration: 6h 31m
- Impact: none, status: resolved
- Components: —
- Canonical incident: [ci-ca32c17732294caf3aa1db74](#ci-ca32c17732294caf3aa1db74)

- `2026-07-23 23:00` **monitoring** — status is health in prod-us-east-3.loki-prod-042 and we're continuing to monitor
- `2026-07-24 05:32` **resolved** — Writes are looking stable in the last 6h

### 15. anthropic — Elevated errors for Mythos 5, Fable 5, Opus 5 and Claude Haiku 4.5

- URL: https://stspg.io/bkqqrx58hq88
- Started: 2026-07-25 18:40 UTC
- Resolved: 2026-07-25 19:44 UTC
- Duration: 1h 4m
- Impact: major, status: resolved
- Components: Claude API (api.anthropic.com), Claude Code, Claude Cowork, claude.ai
- Canonical incident: [ci-3be400de7158658e04504c72](#ci-3be400de7158658e04504c72)

- `2026-07-25 18:40` **investigating** — We are investigating elevated errors on requests to Claude Mythos 5, Claude Fable 5, and Claude Haiku 4.5. We will provide an update as soon as possible.
- `2026-07-25 18:44` **identified** — The issue has been identified and a fix is being implemented.
- `2026-07-25 19:44` **resolved** — This incident has been resolved.

### 16. circleci — Delays in starting  jobs using the gen3 resource class

- URL: https://stspg.io/jjsrf04jn924
- Started: 2026-07-22 19:57 UTC
- Resolved: 2026-07-23 02:05 UTC
- Duration: 6h 8m
- Impact: minor, status: resolved
- Components: Machine Jobs
- Canonical incident: [ci-43f6b28bebc3d38ec09435cd](#ci-43f6b28bebc3d38ec09435cd)

- `2026-07-22 19:57` **investigating** — We are investigating delays in starting jobs using the gen3 resource class.
- `2026-07-22 20:43` **investigating** — We are continuing to investigate elevated wait times affecting customers using gen3 machine executors. Affected jobs may take longer than usual to start. We will provide another update as soon as we have more information to share.
- `2026-07-22 23:17` **investigating** — We are continuing to investigate elevated wait times affecting customers using gen3 machine executors. Affected jobs may take longer than usual to start. We will provide another update as soon as we have more information to share.
- `2026-07-23 00:31` **investigating** — We have identified the issue and are working with our provider to resolve the availability of gen 3 resources. Gen 1 and Gen 2 are both fully operational.
- `2026-07-23 02:05` **resolved** — Between 19:24 UTC on 22 July and 01:50 UTC on 23 July, customers using the gen3 preview machine (Linux VM) resource classes experienced delays and, in some cases, jobs failing to start. We have removed the gen3 resource classes while we address stability issues affecting them. Unfortunately, workflows that we were unable to provision compute capacity for during this period were cancelled. Affected customers should change their resource class from gen3 to gen2 and rerun those workflows. We thank you for your patience while our team worked to resolve this.

### 17. grafana — Errors creating new Slack integration for Grafana IRM

- URL: https://stspg.io/w78r22dh774l
- Started: 2026-07-24 09:17 UTC
- Resolved: 2026-07-24 14:03 UTC
- Duration: 4h 45m
- Impact: major, status: resolved
- Components: Incident Management and Response (IRM)
- Canonical incident: [ci-90ffd6bae1a9cfa743ada1dc](#ci-90ffd6bae1a9cfa743ada1dc)

- `2026-07-24 09:17` **investigating** — We are investigating a possible error creating Slack integrations. Will update the status soon.
- `2026-07-24 09:49` **identified** — We have identified the root cause of the problem, and we are working on a fix
- `2026-07-24 10:12` **identified** — We have verified the fix, and we are starting to roll it out
- `2026-07-24 14:03` **resolved** — This incident has been resolved.

### 18. openai — Some users may experience elevated error rates in ChatGPT

- URL: https://status.openai.com/incidents/01KY81XMFAV8KXNJGVNEMKAK15
- Started: 2026-07-23 17:56 UTC
- Resolved: 2026-07-23 20:41 UTC
- Duration: 2h 44m
- Impact: minor, status: resolved
- Components: —
- Canonical incident: [ci-5dbc328c60dadf215bffbdfe](#ci-5dbc328c60dadf215bffbdfe)

- `2026-07-23 17:56` **identified** — We have identified that users are experiencing elevated errors for the impacted services. We are working on implementing a mitigation.
- `2026-07-23 18:22` **monitoring** — We have applied the mitigation and are monitoring the recovery.
- `2026-07-23 20:41` **resolved** — All impacted services have now fully recovered.

### 19. github — Disruption with GPT 5.3 Codex

- URL: https://stspg.io/7ls47qxt8lpq
- Started: 2026-07-20 16:03 UTC
- Resolved: 2026-07-20 18:37 UTC
- Duration: 2h 33m
- Impact: minor, status: resolved
- Components: Copilot AI Model Providers
- Canonical incident: [ci-0cb621b41d6cdd86332a3a8a](#ci-0cb621b41d6cdd86332a3a8a)

- `2026-07-20 16:03` **investigating** — We are investigating reports of degraded performance for Copilot AI Model Providers
- `2026-07-20 16:04` **investigating** — We are investigating GPT Codex 5.3 performance.
- `2026-07-20 16:39` **investigating** — We identified that an upstream provider is returning errors for GPT 5.3 Codex requests, causing some users to experience failed or interrupted responses. Traffic is being automatically rerouted to mitigate impact. We are working with the upstream provider to resolve the underlying issue. We recommend using a different model at this time while we resolve the problem.
- `2026-07-20 17:21` **investigating** — The upstream provider continues to return intermittent errors for GPT 5.3 Codex requests. Some users may experience failed or interrupted responses when using this model. Traffic continues to be automatically rerouted to reduce impact, and we are actively working with the upstream provider on a resolution. In the meantime, selecting an alternative model will avoid disruption.
- `2026-07-20 18:25` **investigating** — We are seeing signs of sustained recovery from the upstream provider for GPT 5.3 Codex requests. Requests are succeeding as expected. We are awaiting confirmation from the provider that the issue is fully mitigated before resolving this incident.
- `2026-07-20 18:37` **resolved** — Between 06:39 and 18:11 UTC on July 20, 2026, the Copilot service experienced a degradation of the GPT 5.3 model due to an issue with our upstream provider. The upstream model provider returned intermittent errors for GPT 5.3 Codex requests, which caused some responses to fail. Auto mode requests that had selected GPT 5.3 Codex were also impacted. On average about 2% of GPT 5.3 Codex requests failed during this window. Copilot automatically routed eligible traffic away from the impacted provider to reduce customer impact. No other models were impacted.<br /><br />We worked with the upstream provider throughout the incident and confirmed sustained recovery before resolving.

### 20. anthropic — Elevated errors for Claude Fable 5, Claude Sonnet 5, Claude Haiku 4.5, and other models

- URL: https://stspg.io/4l8l9dtyx58j
- Started: 2026-07-25 21:34 UTC
- Resolved: 2026-07-25 22:08 UTC
- Duration: 33m
- Impact: major, status: resolved
- Components: Claude API (api.anthropic.com), Claude Code, Claude Cowork, claude.ai
- Canonical incident: [ci-d4546881179de53b400d9fcf](#ci-d4546881179de53b400d9fcf)

- `2026-07-25 21:34` **investigating** — We are investigating elevated errors on requests to Claude Fable 5, Claude Sonnet 5, Claude Haiku 4.5, and other Claude models. We will provide an update as soon as possible.
- `2026-07-25 22:08` **resolved** — This incident has been resolved.

### 21. slack — Issue sending messages to channels via certain Workflows

- URL: https://slack-status.com/2026-07/e82ee127d8bed735
- Started: 2026-07-24 08:22 UTC
- Resolved: 2026-07-24 10:47 UTC
- Duration: 2h 24m
- Impact: unknown, status: resolved
- Components: Workflows
- Canonical incident: [ci-95ab23ddb544dd23e055028b](#ci-95ab23ddb544dd23e055028b)

- `2026-07-24 08:22` **unknown** — We’re aware of an issue affecting Workflows that include 'refer to a message link'. Users with these workflows are unable to send a message to a channel, and will receive an ‘Input Validation Error’ when trying to perform this step. We’re actively investigating this issue and will provide an update in 30 minutes or sooner if more information becomes available.
- `2026-07-24 08:44` **unknown** — We've engaged our Worklow team to assist in the investigation. Initial analysis appears to show a conflict between specific variables. We’re continuing to work to understand whether this is the trigger of the issue. We’ll provide an update in 30 minutes or sooner if more information becomes available.
- `2026-07-24 09:08` **unknown** — The Workflow team is currently analyzing whether a recent change may have resulted in the conflict between variables that has been observed. It is expected to take some time to validate whether this is the case. We’ll provide an update in 60 minutes or sooner if more information becomes available.
- `2026-07-24 09:42` **unknown** — We’ve performed a test of reverting a recent change and this appears to have resolved the issue. We are in the process of performing a production wide revert of this change. This is expected to take some time to complete and we will validate that it is successful. We’ll provide an update in 2 hours or sooner if more information becomes available.
- `2026-07-24 10:47` **unknown** — We’ve completed the revert of the recent change and have validated that this has brought customers out of impact and can now successfully send a message to a channel using Workflows that include 'refer to a message link'. We apologize for how this incident may have affected you and your business.

### 22. grafana — Metrics Write Path Errors

- URL: https://stspg.io/l59vv2xfq9fx
- Started: 2026-07-23 17:14 UTC
- Resolved: 2026-07-23 19:47 UTC
- Duration: 2h 33m
- Impact: major, status: resolved
- Components: AWS US East - prod-us-east-0: Ingestion
- Canonical incident: [ci-47a857cf4f9be1b2995f839c](#ci-47a857cf4f9be1b2995f839c)

- `2026-07-23 17:14` **investigating** — We are currently investigating an elevated rate of error and latency in the impacted write path.
- `2026-07-23 17:37` **monitoring** — Error rates have dropped, and we are monitoring this issue for any recurrence.
- `2026-07-23 19:47` **resolved** — This incident has been resolved.

### 23. grafana — PDC Issues

- URL: https://stspg.io/vxd316w7jsys
- Started: 2026-07-20 13:48 UTC
- Resolved: 2026-07-20 15:03 UTC
- Duration: 1h 15m
- Impact: major, status: resolved
- Components: AWS Germany - prod-eu-west-2, AWS US East - prod-us-east-2
- Canonical incident: [ci-23106e38418c75e8a9d32b08](#ci-23106e38418c75e8a9d32b08)

- `2026-07-20 13:48` **investigating** — We are currently investigating an issue that is causing issues with PDC in the prod-eu-west-2 region. We will provide another update in 1-2 hours.
- `2026-07-20 14:13` **investigating** — We are continuing to investigate this issue.
- `2026-07-20 14:23` **monitoring** — A fix has been implemented, and we are observing recovery. We will continue to monitor the results.
- `2026-07-20 15:03` **resolved** — This incident has been resolved.

### 24. digitalocean — Gradient AI Serverless Inference Requests Timing Out for qwen3.5-397b-a17b

- URL: https://stspg.io/2fdnm4tns3f1
- Started: 2026-07-22 16:50 UTC
- Resolved: 2026-07-22 19:42 UTC
- Duration: 2h 51m
- Impact: minor, status: resolved
- Components: Inference
- Canonical incident: [ci-acb89c9774e962e45a689ae0](#ci-acb89c9774e962e45a689ae0)

- `2026-07-22 16:50` **investigating** — We are currently investigating this issue.
- `2026-07-22 19:42` **resolved** — This incident has been resolved.

### 25. github — Some SSH connections using deploy keys are failing

- URL: https://stspg.io/jczxclgm9znm
- Started: 2026-07-21 10:31 UTC
- Resolved: 2026-07-21 11:57 UTC
- Duration: 1h 25m
- Impact: critical, status: resolved
- Components: —
- Canonical incident: [ci-834f7edf166d94ffc85df08a](#ci-834f7edf166d94ffc85df08a)

- `2026-07-21 10:31` **investigating** — We are investigating reports of impacted performance for some GitHub services.
- `2026-07-21 11:04` **investigating** — We are investigating reports of intermittent SSH authentication failures affecting connections that use deploy keys. Customers may experience failed SSH connections when interacting with repositories via deploy keys. Our engineering team is actively investigating the root cause and working toward resolution.
- `2026-07-21 11:46` **investigating** — We have identified a recent code change as a potential cause of the SSH authentication failures affecting deploy key connections. Our engineering team is rolling back this change. Customers using deploy keys for SSH access to repositories may continue to experience intermittent connection failures until the fix is deployed.
- `2026-07-21 11:57` **resolved** — On July 21, 2026, between 07:41 UTC and 11:57 UTC, the SSH Authentication service was degraded and some SSH connections failed to authenticate. On average, 12.2% of SSH authentication requests failed, peaking at 15.7%. Both user RSA keys and deploy keys were impacted. This was due to a change in how our SSH service handled one public-key authentication method that caused the affected authentication attempts to be rejected as invalid. <br /><br />We mitigated the incident by reverting the change, after which SSH authentication returned to normal. <br /><br />We are working to expand our automated test coverage for our SSH public-key authentication flows to catch more edge cases and to improve observability and alerting on SSH authentication failures, to reduce our time to detection and mitigation of issues like this one in the future.

### 26. github — Disruption with some GitHub services

- URL: https://stspg.io/2n2g6kzg67mh
- Started: 2026-07-20 00:25 UTC
- Resolved: 2026-07-20 01:46 UTC
- Duration: 1h 20m
- Impact: major, status: resolved
- Components: —
- Canonical incident: [ci-38ea9948c6c84afed7535e2e](#ci-38ea9948c6c84afed7535e2e)

- `2026-07-20 00:25` **investigating** — We are investigating reports of impacted performance for some GitHub services.
- `2026-07-20 00:25` **monitoring** — Some Git LFS operations, and loading files through the API are failing. We are investigating.
- `2026-07-20 00:26` **monitoring** — Some Git LFS operations, and loading files through the API are failing. We are investigating.
- `2026-07-20 01:15` **investigating** — The Git LFS API issues have been identified as being related to a separate incident where Actions is experiencing degraded availability. We’re consolidating our investigation efforts under that incident.
- `2026-07-20 01:45` **investigating** — Git LFS API success rates have returned to normal. We will continue to monitor the service closely.
- `2026-07-20 01:46` **resolved** — Between July 19, 2026, at 23:05 UTC and July 20, 2026, at 03:55 UTC, Actions self-hosted and larger runners were unable to connect to GitHub. During this period, Actions jobs were delayed or failed when trying to acquire a runner. Jobs using standard and Mac hosted runners were not affected. Reconnection traffic from affected runners also increased load on GitHub APIs, resulting in 3-4 seconds of additional average request latency and elevated 5xx error rates.<br /><br />The incident was caused by a certificate lifecycle management failure in a subset of internal services, resulting in an SSL certificate expiration that disrupted runner connectivity. We restored service by rotating the affected certificate. Recovery began at 02:45 UTC. By 03:55 UTC, queued workflow backlog had been processed and workflow delay rates returned to normal.<br /><br />To prevent recurrence, we are strengthening certificate renewal automation, adding fallback expiry monitoring and alerting, and improving circuit-breaker protections during runner API disruptions to reduce the risk of cascading impact to other APIs.

### 27. digitalocean — Network Connectivity from India to NYC

- URL: https://stspg.io/y9g7qgc9qs6m
- Started: 2026-07-24 10:36 UTC
- Resolved: 2026-07-24 12:42 UTC
- Duration: 2h 5m
- Impact: minor, status: resolved
- Components: NYC1
- Canonical incident: [ci-96e507767d267cd3f27ed0a7](#ci-96e507767d267cd3f27ed0a7)

- `2026-07-24 10:36` **identified** — We have noticed ​​​​​​​network connectivity issues affecting some users connecting from India to resources in the NYC region. Based on our current investigation, users on Airtel and BSNL have reported increased latency, packet loss, or intermittent connectivity. We apologize for the inconvenience and will provide an update as soon as more information becomes available.
- `2026-07-24 11:45` **monitoring** — We are observing improvements in network connectivity for users in India accessing resources in our NYC region. Our team is monitoring the situation to ensure that the service has returned to normal operation and remain stable. We appreciate your patience and will provide an update once the issue is fully confirmed as resolved.
- `2026-07-24 12:42` **resolved** — Between 07:46 UTC & 10:24 UTC today, some users in India accessing resources in the NYC region may have experienced connectivity issues. The impact was limited to users using Airtel as their ISP. Our investigation determined that the issue was caused by an internal network issue within Airtel and was not related to DigitalOcean network. The connectivity issue now appears to be resolved, and users should no longer experience problems accessing their resources. We apologize for any inconvenience this may have caused. However, if you continue to experience any connectivity issues, please don't hesitate to open a support ticket for further investigation; we'll be happy to assist you further.

### 28. grafana — Write Outage

- URL: https://stspg.io/z8r7rgprs19d
- Started: 2026-07-24 17:02 UTC
- Resolved: 2026-07-24 18:20 UTC
- Duration: 1h 18m
- Impact: critical, status: resolved
- Components: AWS US East - prod-us-east-3
- Canonical incident: [ci-289396fa39e612576e90d1b7](#ci-289396fa39e612576e90d1b7)

- `2026-07-24 17:02` **investigating** — We are currently investigating a Write outage ongoing, starting at 16:30 UTC for the impacted region.
- `2026-07-24 17:18` **monitoring** — Healthy as of 17:00 UTC. We are continuing to monitor.
- `2026-07-24 18:20` **resolved** — This incident has been resolved.

### 29. anthropic — Service disruption on Claude services

- URL: https://stspg.io/d8v0g0cntyjx
- Started: 2026-07-21 17:40 UTC
- Resolved: 2026-07-21 18:03 UTC
- Duration: 23m
- Impact: critical, status: resolved
- Components: Claude Code, Claude Cowork
- Canonical incident: [ci-0052b1cb38327099cf156234](#ci-0052b1cb38327099cf156234)

- `2026-07-21 17:40` **identified** — We are investigating an outage affecting features such as document creation in claude.ai, Cowork Remote, Claude Code, Claude Code on the Web, Claude Tag, and Claude Design
- `2026-07-21 17:41` **identified** — We are continuing to work on a fix for this issue.
- `2026-07-21 17:46` **identified** — We are continuing to work on a fix for this issue.
- `2026-07-21 17:47` **identified** — We are continuing to work on a fix for this issue.
- `2026-07-21 17:51` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-07-21 18:03` **resolved** — This incident has been resolved.

### 30. github — Incident with Pull Requests

- URL: https://stspg.io/sm1tp7kfm4vj
- Started: 2026-07-24 19:37 UTC
- Resolved: 2026-07-24 20:23 UTC
- Duration: 45m
- Impact: critical, status: resolved
- Components: Pull Requests
- Canonical incident: [ci-2cce0527427160bd310e6cf1](#ci-2cce0527427160bd310e6cf1)

- `2026-07-24 19:37` **investigating** — We are investigating reports of degraded performance for Pull Requests
- `2026-07-24 19:40` **investigating** — We are investigating errors creating pull requests
- `2026-07-24 19:43` **investigating** — Pull Requests is experiencing degraded availability. We are continuing to investigate.
- `2026-07-24 19:59` **investigating** — We have applied a mitigation and are monitoring for recovery
- `2026-07-24 20:02` **monitoring** — The degradation affecting Pull Requests has been mitigated. We are monitoring to ensure stability.
- `2026-07-24 20:23` **resolved** — Between July 24, 19:17 UTC and July 24, 20:02 UTC, users were unable to create pull requests due to a database schema change. In total, 113,930 pull request creation attempts were impacted across 50,904 users, with an average error rate of 1.75% and a maximum error rate of 2.25% for all requests to Pull Requests service. Existing pull requests and other GitHub functionality were not affected. The issue was resolved by reverting the change to the affected database, upon which pull request creation immediately resumed.<br /><br />The root cause was related to a backfill workflow into the Vitess keyspace hosting Pull Request data. The backfill Vitess command encountered errors and increased VReplication lag, and the workflow was canceled at 19:17 UTC. The cancellation executed a misunderstood Vitess codepath that dropped the backing table to the target keyspace, leaving a non-existent reference that resulted in errors creating Pull Requests. The mitigation was executing a command to drop the vschema reference to the dropped table, allowing Pull Request creation to resume.<br /><br />We are adding stronger pre-flight validation to our tooling to prevent similar issues and expanding lower-environment support to provide better test coverage end-to-end before promoting them to production. We're also fixing our backfill migration tooling to protect from this specific codepath.

### 31. github — Actions run failures and delays

- URL: https://stspg.io/448g37mrq066
- Started: 2026-07-25 12:31 UTC
- Resolved: 2026-07-25 13:13 UTC
- Duration: 41m
- Impact: critical, status: resolved
- Components: Actions
- Canonical incident: [ci-9803cbc797dee5741614c403](#ci-9803cbc797dee5741614c403)

- `2026-07-25 12:31` **investigating** — We are investigating reports of degraded availability for Actions
- `2026-07-25 12:34` **investigating** — We are experiencing issues with GitHub Actions that are causing workflow run failures and delays for some users. Our engineering team is actively investigating the infrastructure issue and working to restore full functionality.
- `2026-07-25 12:58` **investigating** — We have applied a mitigation for the infrastructure issue affecting GitHub Actions. Workflow run failures and delays are improving but not yet fully resolved. Our engineering team continues to work on restoring full functionality across all affected infrastructure.
- `2026-07-25 12:59` **monitoring** — The degradation affecting Actions has been mitigated. We are monitoring to ensure stability.
- `2026-07-25 13:12` **monitoring** — We have seen recovery in GitHub Actions performance following our earlier mitigation. Workflow runs are processing normally, though jobs queued before 12:40 UTC may still experience failures and will need to be retried.
- `2026-07-25 13:13` **resolved** — Please refer to the combined summary in this related incident: https://www.githubstatus.com/incidents/s65j9gslmfm8

### 32. github — Several GPT models degraded

- URL: https://stspg.io/vh0xxw69dr6v
- Started: 2026-07-25 09:42 UTC
- Resolved: 2026-07-25 10:11 UTC
- Duration: 28m
- Impact: major, status: resolved
- Components: Copilot AI Model Providers
- Canonical incident: [ci-b4b6166e547786698f106abb](#ci-b4b6166e547786698f106abb)

- `2026-07-25 09:42` **investigating** — We are investigating reports of degraded availability for Copilot AI Model Providers
- `2026-07-25 09:48` **investigating** — We are experiencing degraded availability for the GPT-5.2, GPT-5.3-Codex, GPT-5.4, GPT-5.4 Mini, GPT-5.6 Sol, GPT-5.6 Terra, and GPT-5.6 Luna models in Copilot products and IDE surfaces. This is due to an issue with an upstream model provider. While we work with them to resolve the issue, we recommend choosing another model or selecting 'Auto' to continue using Copilot.
- `2026-07-25 10:04` **monitoring** — The degradation affecting Copilot AI Model Providers has been mitigated. We are monitoring to ensure stability.
- `2026-07-25 10:11` **resolved** — On July 25, 2026, between 09:07 and 10:04 UTC, the GPT-5.2, GPT-5.3-Codex, GPT-5.4, GPT-5.4 Mini, GPT-5.6 Sol, GPT-5.6 Terra and GPT-5.6 Luna models experienced degraded availability in GitHub Copilot products and IDE surfaces. Requests to these models had an average failure rate of 5.6%. Other Copilot models remained available as alternatives.<br /><br />The degradation was caused by an issue with an upstream model provider. Success rates returned to normal after the upstream issue was mitigated, and we continued monitoring before resolving the incident. We are working on improving the automated failover for the affected models to prevent similar incidents in the future.

## Below threshold (28)

- `anthropic` [Elevated errors on Opus 4.8 and Haiku 4.5](https://stspg.io/smvv9yf88dvb) — 1h 46m
- `anthropic` [Elevated errors on Haiku 4.5](https://stspg.io/s8jrbv6mhgjr) — 1h 41m
- `anthropic` [Elevated errors for Claude Opus 4.8](https://stspg.io/msvtz2sxr7nc) — 1h 34m
- `anthropic` [Elevated error rate on Opus 4.8](https://stspg.io/9lj1yq8c9yr4) — 1h 12m
- `anthropic` [Elevated errors across models](https://stspg.io/xg9zkq1s223z) — 1h 7m
- `anthropic` [Elevated error rates for Sonnet 5](https://stspg.io/qpyxcndzhjft) — 48m
- `anthropic` [Elevated errors across Fable 5](https://stspg.io/00dy61t4qg2t) — 47m
- `circleci` [Errors with GitHub APIs delaying workflows](https://stspg.io/96ffdz58cxdr) — 1h 9m
- `openai` [Elevated error rates](https://status.openai.com/incidents/01KYC921K145JTR1JK7DYKGWH1) — 1h 50m
- `openai` [Elevated Errors in Codex Review](https://status.openai.com/incidents/01KY9HYHCMK8AND23B4X17XQEK) — 1h 39m
- `github` [Disruption with actions hosted runners](https://stspg.io/hccgdw2k2b1q) — 1h 25m
- `stripe` [Elevated BLIK errors](https://stspg.io/pzxm22cjk88l) — 49m
- `anthropic` [Elevated errors for Claude Opus 4.1](https://stspg.io/40bs22jyqpnv) — 25m
- `slack` [Receiving reminders in Slack issue](https://slack-status.com/2026-07/df254c60ace250d9) — 1h 1m
- `circleci` [GitHub service degradation may impact customers using GitHub](https://stspg.io/8pvsx88khry9) — 1h 5m
- `digitalocean` [Llama-4-Maverick Model Availability](https://stspg.io/9pxmkrlffl58) — 1h 20m
- `datadog` [Latency with Pages and Monitor state changes](https://stspg.io/ljpc864wlwq2) — 1h 4m
- `grafana` [K6 - Cloud output test-runs are failing to fetch the script logs](https://stspg.io/hf1g7lk54bf4) — 57m
- `anthropic` [Service disruption on Claude Services](https://stspg.io/8ngvd23d61y9) — 18m
- `anthropic` [Elevated errors on Haiku 4.5](https://stspg.io/lnwsjc788wp9) — 17m
- `github` [Incident with Actions](https://stspg.io/vkt1mn9sny66) — 25m
- `stripe` [Elevated API errors](https://stspg.io/wtfd9x61syx3) — 14m
- `openai` [Elevated error rates](https://status.openai.com/incidents/01KYCGY017EG43XZS6GFVXA8VH) — 21m
- `openai` [Elevated API error rates and latency on gpt-image-2 model](https://status.openai.com/incidents/01KY8PGR54KXGBSFXHEJ3FM83F) — 13m
- `circleci` [Errors loading app.circleci.com](https://stspg.io/cw24pb57b6s9) — 13m
- `stripe` [Elevated BLIK Errors](https://stspg.io/g4vzl6b05kny) — 0m
- `github` [Incident With Blocked GitHub.com Traffic](https://stspg.io/dd1j7wg77hn9) — unknown
- `anthropic` [Sonnet 4.6 and Sonnet 5 errors elevated](https://stspg.io/ty4xl4702q4r) — unknown

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-ecf9d6bdd2eba01d8cde5c68"></a>

- Canonical ID: `ci-ecf9d6bdd2eba01d8cde5c68` (1 source records).
  - [grafana/gckqs9bstld1](../../data/2026/W30/grafana.json); vendor URL: https://stspg.io/93c20d1dzpdp
- Start time basis: `vendor_reported`.

<a id="ci-73994a4f2afecc8a2699b70b"></a>

- Canonical ID: `ci-73994a4f2afecc8a2699b70b` (1 source records).
  - [anthropic/yc7mr2gvl914](../../data/2026/W30/anthropic.json); vendor URL: https://stspg.io/y411l9qwm1r8
- Start time basis: `vendor_reported`.

<a id="ci-1aa02afa33e6459e6e254fa7"></a>

- Canonical ID: `ci-1aa02afa33e6459e6e254fa7` (1 source records).
  - [openai/01KYDN6YPS6ARY1EC9089N089G](../../data/2026/W30/openai.json); vendor URL: https://status.openai.com/incidents/01KYDN6YPS6ARY1EC9089N089G
- Start time basis: `vendor_reported`.

<a id="ci-0e74a35a4681977d515bc24d"></a>

- Canonical ID: `ci-0e74a35a4681977d515bc24d` (1 source records).
  - [openai/01KY7SX5MYJ2BP51X5MXAPYX71](../../data/2026/W30/openai.json); vendor URL: https://status.openai.com/incidents/01KY7SX5MYJ2BP51X5MXAPYX71
- Start time basis: `vendor_reported`.

<a id="ci-4e1351d93fe3dc96ed255554"></a>

- Canonical ID: `ci-4e1351d93fe3dc96ed255554` (1 source records).
  - [grafana/91n91dr0dg6x](../../data/2026/W30/grafana.json); vendor URL: https://stspg.io/cjjmv6qvg1wl
- Start time basis: `vendor_reported`.

<a id="ci-6e7020ec5d38bc05489c666c"></a>

- Canonical ID: `ci-6e7020ec5d38bc05489c666c` (1 source records).
  - [anthropic/hlt0yqp4lkww](../../data/2026/W30/anthropic.json); vendor URL: https://stspg.io/3hc27lgyzhcf
- Start time basis: `vendor_reported`.

<a id="ci-b254336524302a9d9c36cdbb"></a>

- Canonical ID: `ci-b254336524302a9d9c36cdbb` (1 source records).
  - [anthropic/j0y67nj6w93z](../../data/2026/W30/anthropic.json); vendor URL: https://stspg.io/twj0mkj40p93
- Start time basis: `vendor_reported`.

<a id="ci-78ae2d5a9a030ef43adcb90e"></a>

- Canonical ID: `ci-78ae2d5a9a030ef43adcb90e` (1 source records).
  - [anthropic/s39rkbp2xhpd](../../data/2026/W30/anthropic.json); vendor URL: https://stspg.io/lqg8z33rmjz9
- Start time basis: `vendor_reported`.

<a id="ci-f793d1cf9e56471fae5d5b28"></a>

- Canonical ID: `ci-f793d1cf9e56471fae5d5b28` (1 source records).
  - [anthropic/vcynh9cf33xp](../../data/2026/W30/anthropic.json); vendor URL: https://stspg.io/ltgjlwc77f49
- Start time basis: `vendor_reported`.

<a id="ci-825a7810739c35df64d1807c"></a>

- Canonical ID: `ci-825a7810739c35df64d1807c` (1 source records).
  - [anthropic/f7w8qkd8z57f](../../data/2026/W30/anthropic.json); vendor URL: https://stspg.io/smvv9yf88dvb
- Start time basis: `vendor_reported`.

<a id="ci-b66517e2a6d618c68de32f68"></a>

- Canonical ID: `ci-b66517e2a6d618c68de32f68` (1 source records).
  - [anthropic/v35b71rthzhp](../../data/2026/W30/anthropic.json); vendor URL: https://stspg.io/s8jrbv6mhgjr
- Start time basis: `vendor_reported`.

<a id="ci-c5227a9ce8abc98311bd9e27"></a>

- Canonical ID: `ci-c5227a9ce8abc98311bd9e27` (1 source records).
  - [anthropic/tnypgb2jbqnq](../../data/2026/W30/anthropic.json); vendor URL: https://stspg.io/r596yvz5clzv
- Start time basis: `vendor_reported`.

<a id="ci-009cba4a402059471b65f57c"></a>

- Canonical ID: `ci-009cba4a402059471b65f57c` (1 source records).
  - [github/yjysg0xrl67m](../../data/2026/W30/github.json); vendor URL: https://stspg.io/j5c80shxqm53
- Start time basis: `vendor_reported`.

<a id="ci-97c6893bf978af77d8a4bfc4"></a>

- Canonical ID: `ci-97c6893bf978af77d8a4bfc4` (1 source records).
  - [anthropic/4n0mmv2v41rf](../../data/2026/W30/anthropic.json); vendor URL: https://stspg.io/msvtz2sxr7nc
- Start time basis: `vendor_reported`.

<a id="ci-4b10c3920f80eefc8e3fd34a"></a>

- Canonical ID: `ci-4b10c3920f80eefc8e3fd34a` (1 source records).
  - [anthropic/zftg3gqkmv18](../../data/2026/W30/anthropic.json); vendor URL: https://stspg.io/sx0f2mwcqw7h
- Start time basis: `vendor_reported`.

<a id="ci-05543d9dec1918885cbe7394"></a>

- Canonical ID: `ci-05543d9dec1918885cbe7394` (1 source records).
  - [github/zq3c1jst2vkq](../../data/2026/W30/github.json); vendor URL: https://stspg.io/syhr80rth84z
- Start time basis: `vendor_reported`.

<a id="ci-2741927eb43012509c161e0f"></a>

- Canonical ID: `ci-2741927eb43012509c161e0f` (1 source records).
  - [anthropic/0lbk5m7mbqk5](../../data/2026/W30/anthropic.json); vendor URL: https://stspg.io/9lj1yq8c9yr4
- Start time basis: `vendor_reported`.

<a id="ci-2cbfcd7c32dfb0cf411c79d8"></a>

- Canonical ID: `ci-2cbfcd7c32dfb0cf411c79d8` (1 source records).
  - [anthropic/y65gwqbztbsp](../../data/2026/W30/anthropic.json); vendor URL: https://stspg.io/xg9zkq1s223z
- Start time basis: `vendor_reported`.

<a id="ci-ca32c17732294caf3aa1db74"></a>

- Canonical ID: `ci-ca32c17732294caf3aa1db74` (1 source records).
  - [grafana/rwz1qdd4hsjm](../../data/2026/W30/grafana.json); vendor URL: https://stspg.io/xf7vsp1rv6yk
- Start time basis: `vendor_reported`.

<a id="ci-3be400de7158658e04504c72"></a>

- Canonical ID: `ci-3be400de7158658e04504c72` (1 source records).
  - [anthropic/f1q2wvw4qj4j](../../data/2026/W30/anthropic.json); vendor URL: https://stspg.io/bkqqrx58hq88
- Start time basis: `vendor_reported`.

<a id="ci-43f6b28bebc3d38ec09435cd"></a>

- Canonical ID: `ci-43f6b28bebc3d38ec09435cd` (1 source records).
  - [circleci/l3hk1qv7hbmj](../../data/2026/W30/circleci.json); vendor URL: https://stspg.io/jjsrf04jn924
- Start time basis: `vendor_reported`.

<a id="ci-f12df746db0d4a7f7e6b070f"></a>

- Canonical ID: `ci-f12df746db0d4a7f7e6b070f` (1 source records).
  - [anthropic/6h7589wl6l45](../../data/2026/W30/anthropic.json); vendor URL: https://stspg.io/qpyxcndzhjft
- Start time basis: `vendor_reported`.

<a id="ci-d61375a8696a512f27c5f4c2"></a>

- Canonical ID: `ci-d61375a8696a512f27c5f4c2` (1 source records).
  - [anthropic/lw0sb1srtznh](../../data/2026/W30/anthropic.json); vendor URL: https://stspg.io/00dy61t4qg2t
- Start time basis: `vendor_reported`.

<a id="ci-90ffd6bae1a9cfa743ada1dc"></a>

- Canonical ID: `ci-90ffd6bae1a9cfa743ada1dc` (1 source records).
  - [grafana/9wj61fx62f6x](../../data/2026/W30/grafana.json); vendor URL: https://stspg.io/w78r22dh774l
- Start time basis: `vendor_reported`.

<a id="ci-5dbc328c60dadf215bffbdfe"></a>

- Canonical ID: `ci-5dbc328c60dadf215bffbdfe` (1 source records).
  - [openai/01KY81XMFAV8KXNJGVNEMKAK15](../../data/2026/W30/openai.json); vendor URL: https://status.openai.com/incidents/01KY81XMFAV8KXNJGVNEMKAK15
- Start time basis: `vendor_reported`.

<a id="ci-0cb621b41d6cdd86332a3a8a"></a>

- Canonical ID: `ci-0cb621b41d6cdd86332a3a8a` (1 source records).
  - [github/fd7j2mw8xw94](../../data/2026/W30/github.json); vendor URL: https://stspg.io/7ls47qxt8lpq
- Start time basis: `vendor_reported`.

<a id="ci-45a0ab5b92f94877a1a1874e"></a>

- Canonical ID: `ci-45a0ab5b92f94877a1a1874e` (1 source records).
  - [circleci/9lvbbbs9l87b](../../data/2026/W30/circleci.json); vendor URL: https://stspg.io/96ffdz58cxdr
- Start time basis: `vendor_reported`.

<a id="ci-d4546881179de53b400d9fcf"></a>

- Canonical ID: `ci-d4546881179de53b400d9fcf` (1 source records).
  - [anthropic/j6gv862066n7](../../data/2026/W30/anthropic.json); vendor URL: https://stspg.io/4l8l9dtyx58j
- Start time basis: `vendor_reported`.

<a id="ci-95ab23ddb544dd23e055028b"></a>

- Canonical ID: `ci-95ab23ddb544dd23e055028b` (1 source records).
  - [slack/1573](../../data/2026/W30/slack.json); vendor URL: https://slack-status.com/2026-07/e82ee127d8bed735
- Start time basis: `vendor_reported`.

<a id="ci-53fc37dd8f603fd118d91c2a"></a>

- Canonical ID: `ci-53fc37dd8f603fd118d91c2a` (1 source records).
  - [openai/01KYC921K145JTR1JK7DYKGWH1](../../data/2026/W30/openai.json); vendor URL: https://status.openai.com/incidents/01KYC921K145JTR1JK7DYKGWH1
- Start time basis: `vendor_reported`.

<a id="ci-47a857cf4f9be1b2995f839c"></a>

- Canonical ID: `ci-47a857cf4f9be1b2995f839c` (1 source records).
  - [grafana/txq1tkqmnxkx](../../data/2026/W30/grafana.json); vendor URL: https://stspg.io/l59vv2xfq9fx
- Start time basis: `vendor_reported`.

<a id="ci-23106e38418c75e8a9d32b08"></a>

- Canonical ID: `ci-23106e38418c75e8a9d32b08` (1 source records).
  - [grafana/k9x0hqq00sj9](../../data/2026/W30/grafana.json); vendor URL: https://stspg.io/vxd316w7jsys
- Start time basis: `vendor_reported`.

<a id="ci-6364bff573bd04b3dc063285"></a>

- Canonical ID: `ci-6364bff573bd04b3dc063285` (1 source records).
  - [openai/01KY9HYHCMK8AND23B4X17XQEK](../../data/2026/W30/openai.json); vendor URL: https://status.openai.com/incidents/01KY9HYHCMK8AND23B4X17XQEK
- Start time basis: `vendor_reported`.

<a id="ci-acb89c9774e962e45a689ae0"></a>

- Canonical ID: `ci-acb89c9774e962e45a689ae0` (1 source records).
  - [digitalocean/j6cjljphvgq7](../../data/2026/W30/digitalocean.json); vendor URL: https://stspg.io/2fdnm4tns3f1
- Start time basis: `vendor_reported`.

<a id="ci-834f7edf166d94ffc85df08a"></a>

- Canonical ID: `ci-834f7edf166d94ffc85df08a` (1 source records).
  - [github/g40zcbvchny4](../../data/2026/W30/github.json); vendor URL: https://stspg.io/jczxclgm9znm
- Start time basis: `vendor_reported`.

<a id="ci-71fa8bb666b57d5202f894b1"></a>

- Canonical ID: `ci-71fa8bb666b57d5202f894b1` (1 source records).
  - [github/20frdtvv3yg6](../../data/2026/W30/github.json); vendor URL: https://stspg.io/hccgdw2k2b1q
- Start time basis: `vendor_reported`.

<a id="ci-38ea9948c6c84afed7535e2e"></a>

- Canonical ID: `ci-38ea9948c6c84afed7535e2e` (1 source records).
  - [github/ph5nns5y4gxj](../../data/2026/W30/github.json); vendor URL: https://stspg.io/2n2g6kzg67mh
- Start time basis: `vendor_reported`.

<a id="ci-a8287c18822e0068b8cf01cd"></a>

- Canonical ID: `ci-a8287c18822e0068b8cf01cd` (1 source records).
  - [stripe/cm7qs1tjkw6c](../../data/2026/W30/stripe.json); vendor URL: https://stspg.io/pzxm22cjk88l
- Start time basis: `vendor_reported`.

<a id="ci-96e507767d267cd3f27ed0a7"></a>

- Canonical ID: `ci-96e507767d267cd3f27ed0a7` (1 source records).
  - [digitalocean/w53965shllfq](../../data/2026/W30/digitalocean.json); vendor URL: https://stspg.io/y9g7qgc9qs6m
- Start time basis: `vendor_reported`.

<a id="ci-289396fa39e612576e90d1b7"></a>

- Canonical ID: `ci-289396fa39e612576e90d1b7` (1 source records).
  - [grafana/mbf85jd64gh6](../../data/2026/W30/grafana.json); vendor URL: https://stspg.io/z8r7rgprs19d
- Start time basis: `vendor_reported`.

<a id="ci-69a2f3cdbf93291b8ecb824c"></a>

- Canonical ID: `ci-69a2f3cdbf93291b8ecb824c` (1 source records).
  - [anthropic/xvbzmkjbzh0x](../../data/2026/W30/anthropic.json); vendor URL: https://stspg.io/40bs22jyqpnv
- Start time basis: `vendor_reported`.

<a id="ci-69cbc5a717428af4aeb58694"></a>

- Canonical ID: `ci-69cbc5a717428af4aeb58694` (1 source records).
  - [slack/1572](../../data/2026/W30/slack.json); vendor URL: https://slack-status.com/2026-07/df254c60ace250d9
- Start time basis: `vendor_reported`.

<a id="ci-0052b1cb38327099cf156234"></a>

- Canonical ID: `ci-0052b1cb38327099cf156234` (1 source records).
  - [anthropic/4qsb0fp16tpb](../../data/2026/W30/anthropic.json); vendor URL: https://stspg.io/d8v0g0cntyjx
- Start time basis: `vendor_reported`.

<a id="ci-2cce0527427160bd310e6cf1"></a>

- Canonical ID: `ci-2cce0527427160bd310e6cf1` (1 source records).
  - [github/jxd617hfwfq8](../../data/2026/W30/github.json); vendor URL: https://stspg.io/sm1tp7kfm4vj
- Start time basis: `vendor_reported`.

<a id="ci-dbdcc0a857d3dc6e7767f005"></a>

- Canonical ID: `ci-dbdcc0a857d3dc6e7767f005` (1 source records).
  - [circleci/q17ysw9wn2pr](../../data/2026/W30/circleci.json); vendor URL: https://stspg.io/8pvsx88khry9
- Start time basis: `vendor_reported`.

<a id="ci-5c15923d0cbe4b7154e9404f"></a>

- Canonical ID: `ci-5c15923d0cbe4b7154e9404f` (1 source records).
  - [digitalocean/ym1lj8ww7cpz](../../data/2026/W30/digitalocean.json); vendor URL: https://stspg.io/9pxmkrlffl58
- Start time basis: `vendor_reported`.

<a id="ci-fe609f66dfa6fbfb0c77c9f9"></a>

- Canonical ID: `ci-fe609f66dfa6fbfb0c77c9f9` (1 source records).
  - [datadog/qvphgr1lgbzh](../../data/2026/W30/datadog.json); vendor URL: https://stspg.io/ljpc864wlwq2
- Start time basis: `vendor_reported`.

<a id="ci-9803cbc797dee5741614c403"></a>

- Canonical ID: `ci-9803cbc797dee5741614c403` (1 source records).
  - [github/pz7g535gbs6p](../../data/2026/W30/github.json); vendor URL: https://stspg.io/448g37mrq066
- Start time basis: `vendor_reported`.

<a id="ci-febede7b1859d07a0d87f1fd"></a>

- Canonical ID: `ci-febede7b1859d07a0d87f1fd` (1 source records).
  - [grafana/8vfwtqt5k7r4](../../data/2026/W30/grafana.json); vendor URL: https://stspg.io/hf1g7lk54bf4
- Start time basis: `vendor_reported`.

<a id="ci-ec535058ad3d33b5a04a6339"></a>

- Canonical ID: `ci-ec535058ad3d33b5a04a6339` (1 source records).
  - [anthropic/ng5shnqxj1r7](../../data/2026/W30/anthropic.json); vendor URL: https://stspg.io/8ngvd23d61y9
- Start time basis: `vendor_reported`.

<a id="ci-31a8f1e51b3878a7a9ba004c"></a>

- Canonical ID: `ci-31a8f1e51b3878a7a9ba004c` (1 source records).
  - [anthropic/db1mh1r2zrn5](../../data/2026/W30/anthropic.json); vendor URL: https://stspg.io/lnwsjc788wp9
- Start time basis: `vendor_reported`.

<a id="ci-b4b6166e547786698f106abb"></a>

- Canonical ID: `ci-b4b6166e547786698f106abb` (1 source records).
  - [github/vv9vvksmj4s9](../../data/2026/W30/github.json); vendor URL: https://stspg.io/vh0xxw69dr6v
- Start time basis: `vendor_reported`.

<a id="ci-ba71cf30c8ba4d7951128a2e"></a>

- Canonical ID: `ci-ba71cf30c8ba4d7951128a2e` (1 source records).
  - [github/s65j9gslmfm8](../../data/2026/W30/github.json); vendor URL: https://stspg.io/vkt1mn9sny66
- Start time basis: `vendor_reported`.

<a id="ci-3f61cab21fdb47e71889eb82"></a>

- Canonical ID: `ci-3f61cab21fdb47e71889eb82` (1 source records).
  - [stripe/s0nbrdrprmrt](../../data/2026/W30/stripe.json); vendor URL: https://stspg.io/wtfd9x61syx3
- Start time basis: `vendor_reported`.

<a id="ci-12a2d93695356aa521991049"></a>

- Canonical ID: `ci-12a2d93695356aa521991049` (1 source records).
  - [openai/01KYCGY017EG43XZS6GFVXA8VH](../../data/2026/W30/openai.json); vendor URL: https://status.openai.com/incidents/01KYCGY017EG43XZS6GFVXA8VH
- Start time basis: `vendor_reported`.

<a id="ci-dd84134b6a24813135a961c0"></a>

- Canonical ID: `ci-dd84134b6a24813135a961c0` (1 source records).
  - [openai/01KY8PGR54KXGBSFXHEJ3FM83F](../../data/2026/W30/openai.json); vendor URL: https://status.openai.com/incidents/01KY8PGR54KXGBSFXHEJ3FM83F
- Start time basis: `vendor_reported`.

<a id="ci-d6225ea0999176d1945d5310"></a>

- Canonical ID: `ci-d6225ea0999176d1945d5310` (1 source records).
  - [circleci/l6yh0qf4c20c](../../data/2026/W30/circleci.json); vendor URL: https://stspg.io/cw24pb57b6s9
- Start time basis: `vendor_reported`.

<a id="ci-537f425a078fbc6a0109667e"></a>

- Canonical ID: `ci-537f425a078fbc6a0109667e` (1 source records).
  - [stripe/xxkcvlgj8wxw](../../data/2026/W30/stripe.json); vendor URL: https://stspg.io/g4vzl6b05kny
- Start time basis: `vendor_reported`.

<a id="ci-234f9d724c7652a7742b54c1"></a>

- Canonical ID: `ci-234f9d724c7652a7742b54c1` (1 source records).
  - [github/594m87r8sw13](../../data/2026/W30/github.json); vendor URL: https://stspg.io/dd1j7wg77hn9
- Start time basis: `vendor_reported`.
- Warning: Vendor resolved_at precedes started_at; duration is unknown, score uses zero duration; decision:legacy-time-github-594m87r8sw13

<a id="ci-2093485de2e13e90ca10e8eb"></a>

- Canonical ID: `ci-2093485de2e13e90ca10e8eb` (1 source records).
  - [anthropic/1019wwb67615](../../data/2026/W30/anthropic.json); vendor URL: https://stspg.io/ty4xl4702q4r
- Start time basis: `vendor_reported`.
- Warning: Vendor resolved_at precedes started_at; duration is unknown, score uses zero duration; decision:legacy-time-anthropic-1019wwb67615

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
