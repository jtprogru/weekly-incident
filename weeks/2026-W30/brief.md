# Week 2026-W30 — top 5 of 32 above threshold

2026-07-20 to 2026-07-26 (UTC).

## 1. grafana — Stacks using SCIM user provisioning are currently unable to log into Grafana

- URL: https://stspg.io/93c20d1dzpdp
- Window: 2026-07-23 10:13 to 2026-07-23 15:37 UTC (5h 24m)
- Impact: major, status: resolved
- Components: AWS Australia - prod-ap-southeast-2, AWS Brazil - prod-sa-east-1, AWS Canada - prod-ca-east-0, AWS Germany - prod-eu-west-2, AWS Germany - prod-eu-west-4, AWS India - prod-ap-south-1, AWS Japan - prod-ap-northeast-0, AWS UAE - prod-me-central-1, AWS Singapore - prod-ap-southeast-1, AWS Sweden - prod-eu-north-0, AWS US East - prod-us-east-0, AWS US East - prod-us-east-2, AWS US West - prod-us-west-0, AWS Australia - prod-au-southeast-1, AWS UK - prod-gb-south-1, AWS Ireland - prod-eu-west-6, Azure US Central - us-central2, AWS Switzerland - prod-eu-central-0, Azure Netherlands - prod-eu-west-3, GCP Australia - prod-au-southeast-0, GCP Belgium - prod-eu-west-0, GCP Brazil - prod-sa-east-0, GCP India - prod-ap-south-0, GCP Singapore - prod-ap-southeast-0, GCP UK - prod-gb-south-0, GCP US Central - prod-us-central-0, GCP US Central - prod-us-central-3, GCP US Central - prod-us-central-4, GCP US East - prod-us-east-1, play.grafana.org, Federal Cloud - AWS US Gov West

  - `2026-07-23 10:13` investigating — We are currently facing an issue where Stacks using SCIM user provisioning are currently unable to log into Grafana. We are currently investigating this issue and working on a fix.
  - `2026-07-23 11:26` identified — The issue has been identified and we are currently working on the mitigations now. The scope is much more limited than initially thought only specific SCIM configurations are impacted.
  - `2026-07-23 13:23` identified — We have applied a fix and are currently awaiting feedback from affected customers.
  - `2026-07-23 15:37` resolved — This incident has been resolved.

## 2. anthropic — Microsoft Office add-in availability

- URL: https://stspg.io/y411l9qwm1r8
- Window: 2026-07-22 17:56 to 2026-07-24 15:36 UTC (1d 21h 40m)
- Impact: none, status: resolved
- Components: —

  - `2026-07-22 17:56` identified — Users who installed the following add-ins may experience issues accessing the add-in functionality: Claude by Anthropic for PowerPoint, Claude by Anthropic for Word, and Claude by Anthropic for Excel. To regain immediate access to Claude for Microsoft 365, install the add-in here: https://marketplace.microsoft.com/en-us/product/office/wa200010725
  - `2026-07-24 15:36` resolved — This issue has been resolved.

## 3. openai — Elevated errors affecting ChatGPT conversations

- URL: https://status.openai.com/incidents/01KYDN6YPS6ARY1EC9089N089G
- Window: 2026-07-25 22:09 to 2026-07-27 16:32 UTC (1d 18h 22m)
- Impact: minor, status: resolved
- Components: —

  - `2026-07-25 22:09` investigating — We are investigating intermittent errors that may prevent some users from loading or continuing conversations in ChatGPT. The current period of impact began around 1:00 PM PT.
  - `2026-07-25 23:16` identified — We have identified the source of intermittent errors affecting some ChatGPT conversations and are working to mitigate the issue. Error rates have decreased, but some users may continue to experience failures.
  - `2026-07-25 23:57` monitoring — Mitigation has been implemented, and we are monitoring the recovery of ChatGPT conversations.
  - `2026-07-27 16:32` resolved — All impacted services have now fully recovered.

## 4. openai — Elevated Error Rates

- URL: https://status.openai.com/incidents/01KY7SX5MYJ2BP51X5MXAPYX71
- Window: 2026-07-23 15:36 to 2026-07-24 19:33 UTC (1d 3h 57m)
- Impact: minor, status: resolved
- Components: —

  - `2026-07-23 15:36` investigating — We are investigating the issue for the listed services.
  - `2026-07-23 15:56` identified — We have identified that users are experiencing elevated errors for the impacted services. We are working on implementing a mitigation.
  - `2026-07-23 16:08` identified — We have identified that users are experiencing elevated errors for the impacted services. We are working on implementing a mitigation.
  - `2026-07-23 16:48` identified — We have identified that users are experiencing elevated errors for the impacted services. We are working on implementing a mitigation with the downstream infrastructure provider.
  - `2026-07-23 17:36` identified — We are continuing to work on implementing a mitigation.
  - `2026-07-23 17:58` identified — We are continuing to work on implementing a mitigation.
  - `2026-07-23 18:50` monitoring — We have applied the mitigation and are monitoring the recovery.
  - `2026-07-23 19:37` investigating — We are investigating the issue for the listed services.
  - `2026-07-23 20:22` monitoring — We have applied the mitigation and are monitoring the recovery.
  - `2026-07-24 09:56` investigating — We are investigating the issue for the listed services.
  - `2026-07-24 10:06` investigating — We are investigating the issue for the listed services.
  - `2026-07-24 11:34` monitoring — We have applied the mitigation and are monitoring the recovery.
  - `2026-07-24 19:33` resolved — All impacted services have now fully recovered.

## 5. grafana — Cloud Log Exporter Unavailable

- URL: https://stspg.io/cjjmv6qvg1wl
- Window: 2026-07-21 18:43 to 2026-07-21 21:06 UTC (2h 22m)
- Impact: major, status: resolved
- Components: AWS Brazil - prod-sa-east-1, AWS Canada - prod-ca-east-0, AWS Germany - prod-eu-west-2, AWS India - prod-ap-south-1, AWS Japan - prod-ap-northeast-0, AWS Singapore - prod-ap-southeast-1, AWS US West - prod-us-west-0, AWS Switzerland - prod-eu-central-0, Azure Netherlands - prod-eu-west-3, Azure US Central - us-central7, GCP Belgium - prod-eu-west-0, GCP Brazil - prod-sa-east-0, GCP India - prod-ap-south-0, GCP Singapore - prod-ap-southeast-0, GCP UK - prod-gb-south-0

  - `2026-07-21 18:43` investigating — Cloud Log Exporter is temporarily not available in the marked regions. We are currently investigating this issue and will update as soon as we have more info to share.
  - `2026-07-21 20:04` monitoring — A fix has been implemented and we are monitoring the results.
  - `2026-07-21 21:06` resolved — This incident has been resolved.

## Also above threshold (27)

- `anthropic` Elevated error rates for Opus 4.5 (5h 44m) — https://stspg.io/3hc27lgyzhcf
- `anthropic` Elevated errors on Haiku 4.5 (3h 33m) — https://stspg.io/twj0mkj40p93
- `anthropic` Elevated errors on Haiku 4.5 (2h 20m) — https://stspg.io/lqg8z33rmjz9
- `anthropic` Elevated errors on several models (2h 43m) — https://stspg.io/ltgjlwc77f49
- `anthropic` Fable 5 requiring usage credits on Max plans (8h 0m) — https://stspg.io/r596yvz5clzv
- `github` Disruption with some GitHub services (1h 19m) — https://stspg.io/j5c80shxqm53
- `anthropic` Elevated errors for Opus 5 (1h 27m) — https://stspg.io/sx0f2mwcqw7h
- `github` Latency issues across a number of services (1h 45m) — https://stspg.io/syhr80rth84z
- `grafana` Intermittent write outage from 21:20-21:26, 21:54-21:56, and 22:22-22:23 (6h 31m) — https://stspg.io/xf7vsp1rv6yk
- `anthropic` Elevated errors for Mythos 5, Fable 5, Opus 5 and Claude Haiku 4.5 (1h 4m) — https://stspg.io/bkqqrx58hq88
- `circleci` Delays in starting  jobs using the gen3 resource class (6h 8m) — https://stspg.io/jjsrf04jn924
- `grafana` Errors creating new Slack integration for Grafana IRM (4h 45m) — https://stspg.io/w78r22dh774l
- `openai` Some users may experience elevated error rates in ChatGPT (2h 44m) — https://status.openai.com/incidents/01KY81XMFAV8KXNJGVNEMKAK15
- `github` Disruption with GPT 5.3 Codex (2h 33m) — https://stspg.io/7ls47qxt8lpq
- `anthropic` Elevated errors for Claude Fable 5, Claude Sonnet 5, Claude Haiku 4.5, and other models (33m) — https://stspg.io/4l8l9dtyx58j
- `slack` Issue sending messages to channels via certain Workflows (2h 24m) — https://slack-status.com/2026-07/e82ee127d8bed735
- `grafana` Metrics Write Path Errors (2h 33m) — https://stspg.io/l59vv2xfq9fx
- `grafana` PDC Issues (1h 15m) — https://stspg.io/vxd316w7jsys
- `digitalocean` Gradient AI Serverless Inference Requests Timing Out for qwen3.5-397b-a17b (2h 51m) — https://stspg.io/2fdnm4tns3f1
- `github` Some SSH connections using deploy keys are failing (1h 25m) — https://stspg.io/jczxclgm9znm
- `github` Disruption with some GitHub services (1h 20m) — https://stspg.io/2n2g6kzg67mh
- `digitalocean` Network Connectivity from India to NYC (2h 5m) — https://stspg.io/y9g7qgc9qs6m
- `grafana` Write Outage (1h 18m) — https://stspg.io/z8r7rgprs19d
- `anthropic` Service disruption on Claude services (23m) — https://stspg.io/d8v0g0cntyjx
- `github` Incident with Pull Requests (45m) — https://stspg.io/sm1tp7kfm4vj
- `github` Actions run failures and delays (41m) — https://stspg.io/448g37mrq066
- `github` Several GPT models degraded (28m) — https://stspg.io/vh0xxw69dr6v

