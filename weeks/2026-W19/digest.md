# Week 2026-W19

2026-05-04 to 2026-05-10 (UTC). Generated 2026-11-14 00:00.

13 canonical incidents, 7 above threshold. 2 of 14 sources unavailable: datadog, yandexcloud.

13 source records (13 saved revisions) → 13 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | slack | [EKM customers were previously experiencing issues with channel loading and message delivery](https://slack-status.com/2026-05/d284d808ed66c511) | 2026-05-08 18:31 | 3d 3h 16m | unknown | Apps/Integrations/APIs, Files, Messaging, Notifications | 21677 = 4516 × 4 × 1.2 |
| 2 | datadog | [Delayed Monitors Notifications](https://stspg.io/rbj54mc4nhgs) | 2026-05-08 00:06 | 14h 0m | critical | APM, CI Visibility, Error Tracking, Log Management, Metrics and Infra Monitoring, Monitors, RUM | 5880 = 840 × 7 × 1.0 |
| 3 | slack | [Slow behaviour in Slack](https://slack-status.com/2026-05/bedf971f90be70a6) | 2026-05-04 10:00 | 7h 37m | unknown | Workflows | 548 = 457 × 1 × 1.2 |
| 4 | slack | [Scheduled Grid Migrations are experiencing extended downtime](https://slack-status.com/2026-05/02b2d5b45d656512) | 2026-05-08 12:13 | 5h 30m | unknown | Connectivity | 396 = 330 × 1 × 1.2 |
| 5 | circleci | [GitHub incident with Pull Requests](https://stspg.io/pttscyh3qsv1) | 2026-05-06 17:54 | 3h 31m | none | GitHub Pull Requests | 211 = 211 × 1 × 1.0 |
| 6 | digitalocean | [Control Panel Errors - Unable to Enable 2FA and Google/GitHub](https://stspg.io/tyb995vv9ykq) | 2026-05-09 11:15 | 3h 27m | minor | Cloud Control Panel | 166 = 207 × 1 × 0.8 |
| 7 | circleci | [Some Pipelines are being lost](https://stspg.io/m852c4blmtsk) | 2026-05-06 19:42 | 30m | major | Pipelines & Workflows | 30 = 30 × 1 × 1.0 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. slack — EKM customers were previously experiencing issues with channel loading and message delivery

- URL: https://slack-status.com/2026-05/d284d808ed66c511
- Started: 2026-05-08 18:31 UTC
- Resolved: 2026-05-11 21:48 UTC
- Duration: 3d 3h 16m
- Impact: unknown, status: resolved
- Components: Apps/Integrations/APIs, Files, Messaging, Notifications
- Canonical incident: [ci-0dc1b47ec8d11c330ab2fbab](#ci-0dc1b47ec8d11c330ab2fbab)

- `2026-05-08 18:31` **unknown** — Between approximately 10:55 AM and 11:10 AM PDT, customers using Enterprise Key Management (EKM) may have experienced issues sending messages and loading some channels. Error rates have returned to normal levels and customers should now see expected performance. We are continuing to investigate the underlying cause and monitoring closely. We apologize for the disruption.
- `2026-05-08 18:44` **unknown** — We’re investigating this issue: Current status and actions being taken: We’re seeing renewed, elevated error rates and latency affecting customers using Enterprise Key Management (EKM). Affected customers may experience issues sending messages and loading some channels. Our team is actively investigating this issue. Estimated time to resolution: No ETA at this time. Scope: This is not affecting any specific region, but a percentage of Enterprise Key Management (EKM) customers may experience impact globally. Current impact to end users: Customers using Enterprise Key Management (EKM) may experience issues sending messages and loading some channels. Known workarounds: No known workarounds at this time. We'll continue to provide updates every 30 minutes until the impact is resolved for all users.
- `2026-05-08 19:20` **unknown** — We’re investigating this issue: Current status and actions being taken: We believe an issue with an upstream provider may be impacting some Slack services. We’re partnering with them to help with our investigation and are waiting on additional details from their analysis. In parallel, we’re actively reviewing the cache configuration to alleviate the impact. Estimated time to resolution: No ETA at this time. Scope: This is not affecting any specific region, but a percentage of Enterprise Key Management (EKM) customers may experience impact globally. Current impact to end users: Customers using Enterprise Key Management (EKM) may experience issues sending messages and loading some channels. Known workarounds: No known workarounds at this time. We'll continue to provide updates every 30 minutes until the impact is resolved for all users.
- `2026-05-08 19:54` **unknown** — We’re investigating this issue: Current status and actions being taken: We’ve deployed a fix that has improved our cache hit ratio and are observing recovery in our health metrics. In parallel, we're monitoring for updates from our upstream provider. Affected customers may already be seeing improvements, and we’ll continue to monitor the situation closely until it is fully resolved. Estimated time to resolution: No ETA at this time. Scope: This is not affecting any specific region, but a percentage of Enterprise Key Management customers may experience impact globally. Current impact to end users: Customers using Enterprise Key Management (EKM) may experience issues sending messages and loading channels. Additionally, they may experience delays with all messaging activity, including notifications and activity feeds. Known workarounds: No known workarounds at this time. We'll continue to provide updates every 30 minutes until the impact is resolved for all users.
- `2026-05-08 20:27` **unknown** — We’re investigating this issue: Current status and actions being taken: We're seeing consistent improvement in our health metrics while our upstream provider has begun implementing a fix on their side. Affected customers may already be seeing improvements, and we’ll continue to monitor the situation closely until it is fully resolved. Estimated time to resolution: No ETA at this time. Scope: This is not affecting any specific region, but a percentage of Enterprise Key Management customers may experience impact globally. Current impact to end users: Customers using Enterprise Key Management (EKM) may experience issues sending messages and loading channels. Additionally, they may experience delays with all messaging activity, including notifications and activity feeds. Known workarounds: No known workarounds at this time. We'll continue to provide updates every 30 minutes until the impact is resolved for all users.
- `2026-05-08 20:57` **unknown** — We’re investigating this issue: Current status and actions being taken: We continue to monitor the issue as our upstream provider continues their work toward a resolution. Affected customers may already be seeing improvements, and we’ll continue to monitor the situation closely until it is fully resolved. Estimated time to resolution: No ETA at this time. Scope: This is not affecting any specific region, but a percentage of Enterprise Key Management customers may experience impact globally. Current impact to end users: Customers using Enterprise Key Management (EKM) may experience issues sending messages and loading channels. Additionally, they may experience delays with all messaging activity, including notifications and activity feeds. Known workarounds: No known workarounds at this time. We'll continue to provide updates every 60 minutes until the impact is resolved for all users.
- `2026-05-08 21:58` **unknown** — We’re investigating this issue: Current status and actions being taken: Our upstream provider has implemented a fix and is continuing their investigation. In the meantime, we continue to observe improved health metrics. Affected customers may already be seeing improvements, and we’ll continue to monitor the situation closely until it is fully resolved Estimated time to resolution: No ETA at this time. Scope: This is not affecting any specific region, but a percentage of Enterprise Key Management customers may experience impact globally. Current impact to end users: Customers using Enterprise Key Management (EKM) may experience issues sending messages and loading channels. Additionally, they may experience delays with all messaging activity, including notifications and activity feeds. Known workarounds: No known workarounds at this time. We'll continue to provide updates every 60 minutes until the impact is resolved for all users.
- `2026-05-08 23:27` **unknown** — This issue is now resolved for all users. Current status and recent actions taken: Our upstream provider has successfully implemented a fix. Following a period of stable monitoring, all services have been restored. Customers should no longer experience issues accessing Slack. Scope: This issue didn’t affect any specific region, but a percentage of Enterprise Key Management customers may have experienced an impact globally. Previous impact to end users: Customers using Enterprise Key Management (EKM) may have experienced issues sending messages and loading channels. Additionally, they may have experienced delays with all messaging activity, including notifications and activity feeds. Required resolution steps: No additional steps are necessary. We apologize for any disruptions to your day.
- `2026-05-11 14:22` **unknown** — Current status and actions being taken: We’re investigating a recurrence of a previous incident causing elevated errors and high latency for EKM customers. As a precaution, EKM backfills have been paused to reduce load. Teams are evaluating pod scaling and assessing a high-volume load test workload contributing to EKGen traffic. A bridge is active with EKM and EKGen engineers engaged, and third-party infrastructure teams have been re-engaged as a potential contributing factor. Estimated time to resolution: Under investigation. No ETA available at this time. Scope: Customers using Enterprise Key Management (EKM). Current impact to end users: Customers using Enterprise Key Management (EKM) may experience issues sending messages and loading channels. Additionally, they may experience delays with all messaging activity, including notifications, Workflow, DMs, and activity feeds. Known workarounds: No known workarounds at this time. We'll continue to provide updates every 30 minutes until the impact is resolved for all users.
- `2026-05-11 15:11` **unknown** — Current status and actions being taken: We're continuing to investigate elevated error rates affecting a subset of users. We've identified additional load contributing to the issue and have taken steps to reduce it. Estimated time to resolution: Under investigation. No ETA available at this time. Scope: Customers using Enterprise Key Management (EKM) Current impact to end users: Customers using Enterprise Key Management (EKM) may experience issues sending messages and loading channels. Additionally, they may experience delays with all messaging activity, including notifications, Workflow, DMs, and activity feeds. Known workarounds: No known workarounds at this time. We'll provide another update in 30 minutes or as new information becomes available. We apologize for the disruption.
- `2026-05-11 16:05` **unknown** — Current status and actions being taken: We're continuing to investigate elevated error rates affecting a subset of users. We've identified additional load contributing to the issue and have taken steps to reduce it. We've also engaged our third-party vendor to assist in providing additional troubleshooting. Estimated time to resolution: Under investigation. No ETA available at this time. Scope: Customers using Enterprise Key Management (EKM) Current impact to end users: Customers using Enterprise Key Management (EKM) may experience issues sending messages and loading channels. Additionally, they may experience delays with all messaging activity, including notifications, Workflow, DMs, and activity feeds. Known workarounds: No known workarounds at this time. We'll provide another update in 30 minutes or as new information becomes available. We apologize for the disruption.
- `2026-05-11 16:38` **unknown** — Current status and actions being taken: We have identified contributing factors to the issue impacting EKM customers and have taken mitigating actions, including reducing the rate of KMS requests by approximately 50% and increasing cache capacity. We are continuing to monitor closely and are actively investigating the underlying cause in partnership with our third-party provider. Estimated time to resolution: The incident has been stable for over an hour with no new customer-facing impact detected. Scope: Customers using Enterprise Key Management (EKM) Current impact to end users: Customers using Enterprise Key Management (EKM) may have experienced issues sending messages and loading channels. Additionally, they may experience delays with all messaging activity, including notifications, Workflow, DMs, and activity feeds. Known workarounds: No known workarounds at this time. We'll provide another update by 9:50 AM PT. We apologize for the disruption.
- `2026-05-11 17:12` **unknown** — Our work on this issue is still ongoing. Current status and actions being taken: Our monitoring has shown no further instability and we have considered this incident now under control. We are continuing to monitor closely. We'll provide another update should we find any further information of note. We apologize for the disruption.
- `2026-05-11 19:02` **unknown** — Our work on this issue is still ongoing. Current status and actions being taken: Customer impact effectively stopped around 8:00AM PT and has remained stable. Our monitoring has shown no further instability and we have considered this incident under control. We are continuing to monitor closely and we'll update again when there's another meaningful update to share. We apologize for the disruption.
- `2026-05-11 21:48` **unknown** — Users who rely on Enterprise Key Management (EKM) may have experienced elevated latency and errors across messaging, channel loading, workflows, and file operations starting around 5:55 AM PDT on May 11, 2026. Impact was mitigated and the incident was brought under control at 9:16 AM PDT on May 11, 2026. We've deployed fixes addressing the root cause of the elevated encryption key request load. If you continue to experience trouble, please reload Slack using Command + Shift + R (Mac) or Ctrl + Shift + R (Windows/Linux). We apologize for the disruption and thank you for your patience.

### 2. datadog — Delayed Monitors Notifications

- URL: https://stspg.io/rbj54mc4nhgs
- Started: 2026-05-08 00:06 UTC
- Resolved: 2026-05-08 14:07 UTC
- Duration: 14h 0m
- Impact: critical, status: resolved
- Components: APM, CI Visibility, Error Tracking, Log Management, Metrics and Infra Monitoring, Monitors, RUM
- Canonical incident: [ci-5c354d99da28f970835e2b89](#ci-5c354d99da28f970835e2b89)

- `2026-05-08 00:06` **investigating** — We are investigating delays in Monitors Notifications, which began at 23:39 UTC.
- `2026-05-08 00:24` **investigating** — We are investigating increased latency across multiple products which began at 23:39 UTC. As a result of this issue, some users may see delays in data across the platform.
- `2026-05-08 00:47` **investigating** — Due to upstream provider issues, we are also continuing to see unavailability of telemetry data coming from AWS into Datadog. https://health.aws.amazon.com/health/status
- `2026-05-08 01:20` **identified** — We have identified upstream provider issues and are continuing to experience delays in processing data across multiple products. We are working on a fix. https://health.aws.amazon.com/health/status
- `2026-05-08 02:02` **identified** — We have identified upstream provider issues and are continuing to experience delays in processing data across multiple products. We are working on a fix. https://health.aws.amazon.com/health/status
- `2026-05-08 02:16` **identified** — We have identified due to upstream provider issues, we are continuing to see unavailability of telemetry data coming from AWS into Datadog. We are continuing to work on a fix. https://health.aws.amazon.com/health/status
- `2026-05-08 02:59` **identified** — We have identified due to upstream provider issues, we are continuing to see unavailability of telemetry data coming from AWS into Datadog. We are continuing to work on a fix. https://health.aws.amazon.com/health/status For metrics we are still seeing delays in distribution metrics. Counts, rates and gauge metrics are being processed normally for most customers
- `2026-05-08 04:13` **identified** — We have identified upstream provider issues and are continuing to experience delays in processing data across multiple products. We are continuing to work on a fix. https://health.aws.amazon.com/health/status For metrics, distribution metrics and point metrics are being processed normally for all customers.
- `2026-05-08 04:48` **identified** — We have identified upstream provider issues and are continuing to experience delays in processing data across multiple products. We are continuing to work on a fix. https://health.aws.amazon.com/health/status For metrics, distribution metrics and point metrics are being processed normally for all customers.
- `2026-05-08 05:30` **identified** — We have identified upstream provider issues and are continuing to experience delays in processing data across multiple products. We are continuing to work on a fix. https://health.aws.amazon.com/health/status For metrics, distribution metrics and point metrics are being processed normally for all customers.
- `2026-05-08 05:44` **identified** — We have identified upstream provider issues. Metrics, Logs, APM and RUM data are being processed normally for all customers. Alerting is also functional. There are still processing delays for CI visibility. We are monitoring the situation. https://health.aws.amazon.com/health/status
- `2026-05-08 06:21` **monitoring** — We have identified upstream provider issues. Metrics, Logs, APM, RUM and CI visibility data are being processed normally for all customers. Alerting is also functional. We are monitoring the situation. https://health.aws.amazon.com/health/status
- `2026-05-08 14:07` **resolved** — This incident has been resolved.

### 3. slack — Slow behaviour in Slack

- URL: https://slack-status.com/2026-05/bedf971f90be70a6
- Started: 2026-05-04 10:00 UTC
- Resolved: 2026-05-04 17:37 UTC
- Duration: 7h 37m
- Impact: unknown, status: resolved
- Components: Workflows
- Canonical incident: [ci-fc970ac075e258fe81586828](#ci-fc970ac075e258fe81586828)

- `2026-05-04 10:00` **unknown** — Some users are experiencing issues with workflows. Triggered workflows may fail to complete, with the first step not executing. We'll provide another update as soon as possible.
- `2026-05-04 11:43` **unknown** — The investigation identified elevated latency in EU backend services as the cause of workflows not completing. A subset of workflow trigger attempts is impacted, primarily, EU users. Pod restarts have stopped, and error rates are improving. A fix is being prepared and will be deployed shortly. We'll provide another update as soon as possible.
- `2026-05-04 13:08` **unknown** — A fix has been deployed to address the elevated latency in EU backend services. The team is monitoring the environment to ensure service stability and full recovery. We'll provide another update as soon as possible.
- `2026-05-04 17:37` **unknown** — This issue is now resolved for all users. Current status and recent actions taken: A fix has been deployed that resolved the underlying cause of this issue. Our health metrics have returned to normal levels and workflows are running as expected. Scope: Users in the EU region — primarily teams based in Central Europe — experienced elevated workflow failure rates. Previous impact to end users: Users may have experienced workflows failing to run, with the first step remaining stuck in an initialized state. Some users in the EU region may also have experienced general slowness when using Slack. Required resolution steps: No steps are required. Workflows should be running normally. If you continue to experience any trouble, please reload Slack using Command + Shift + R (Mac) or Ctrl + Shift + R (Windows/Linux). We apologize for any disruptions this may have caused to your day.

### 4. slack — Scheduled Grid Migrations are experiencing extended downtime

- URL: https://slack-status.com/2026-05/02b2d5b45d656512
- Started: 2026-05-08 12:13 UTC
- Resolved: 2026-05-08 17:44 UTC
- Duration: 5h 30m
- Impact: unknown, status: resolved
- Components: Connectivity
- Canonical incident: [ci-c8d23cd72fab87ac676bd3b4](#ci-c8d23cd72fab87ac676bd3b4)

- `2026-05-08 12:13` **unknown** — Users who are undergoing scheduled Grid Migrations are experiencing extended downtime. Downtime is expected during a Grid Migration however it's taking longer than usual. During the downtime users are unable to access Slack until the migration completes. We've identified an issue with Job Queues, which may be the cause and are investigating this further. The next update will be provided once additional information becomes available.
- `2026-05-08 17:44` **unknown** — This issue is now resolved. Current status and actions being taken: We identified database write throttling as the cause of the stalled grid migrations. We applied a drop rule for jobs that were stuck in a retry loop, which has restored normal migration processing. Scope: A small number of enterprise users undergoing grid migrations experienced extended delays in accessing Slack. Current impact to end users: Affected users may have experienced an inability to access Slack during their migration window. Known workarounds: No workarounds needed at this time. We apologize for how this incident affected you and your business.

### 5. circleci — GitHub incident with Pull Requests

- URL: https://stspg.io/pttscyh3qsv1
- Started: 2026-05-06 17:54 UTC
- Resolved: 2026-05-06 21:26 UTC
- Duration: 3h 31m
- Impact: none, status: resolved
- Components: GitHub Pull Requests
- Canonical incident: [ci-7e13c723b5a09b2249a61506](#ci-7e13c723b5a09b2249a61506)

- `2026-05-06 17:54` **monitoring** — GitHub is reporting a major incident for degraded pull request availability. This does not impact running CircleCI builds for newly created pull requests or push events.
- `2026-05-06 17:54` **monitoring** — We are continuing to monitor for any further issues.
- `2026-05-06 21:26` **resolved** — This incident has been resolved.

### 6. digitalocean — Control Panel Errors - Unable to Enable 2FA and Google/GitHub

- URL: https://stspg.io/tyb995vv9ykq
- Started: 2026-05-09 11:15 UTC
- Resolved: 2026-05-09 14:42 UTC
- Duration: 3h 27m
- Impact: minor, status: resolved
- Components: Cloud Control Panel
- Canonical incident: [ci-3403f15a18fedc3fe252dd01](#ci-3403f15a18fedc3fe252dd01)

- `2026-05-09 11:15` **investigating** — Our Engineering team is investigating an issue affecting the ability to enable Two-Factor Authentication (2FA) and Google/GitHub authentication through the Control Panel. During this time, users may encounter errors while enabling these authentication methods and could also experience issues accessing teams with secure sign-in enabled. We apologize for the inconvenience and will provide an update as soon as more information becomes available.
- `2026-05-09 13:51` **monitoring** — Our Engineering team has implemented necessary changes to address the issue affecting the ability to enable Two-Factor Authentication (2FA) and Google/GitHub authentication through the Control Panel. Our team is currently monitoring the situation to ensure stability. We appreciate your patience and will provide an update once the issue is fully confirmed as resolved.
- `2026-05-09 14:42` **resolved** — Between 5:35 and 14:25 UTC today, our Engineering team identified an issue that prevents enabling Two-Factor Authentication (2FA) and Google/GitHub authentication through the Control Panel. During this period, users might have encountered issues while enabling authentication methods and accessing teams with secure sign-in enabled. Our team has taken appropriate measures to address the issue. We can confirm that service has been restored and is now functioning normally. We regret the inconvenience caused. However, if you continue to experience any issues, please create a support ticket for further analysis.

### 7. circleci — Some Pipelines are being lost

- URL: https://stspg.io/m852c4blmtsk
- Started: 2026-05-06 19:42 UTC
- Resolved: 2026-05-06 20:13 UTC
- Duration: 30m
- Impact: major, status: resolved
- Components: Pipelines & Workflows
- Canonical incident: [ci-753245b59eda340e89bceebe](#ci-753245b59eda340e89bceebe)

- `2026-05-06 19:42` **investigating** — We are having an issue with some pipelines being lost. We are rolling back the affected service, and will share more information when we have it.
- `2026-05-06 19:43` **investigating** — We are continuing to investigate this issue.
- `2026-05-06 20:03` **monitoring** — The rollback was successful. Pipelines are flowing normally again. Monitoring for a bit.
- `2026-05-06 20:13` **resolved** — This incident is resolved. Thank you for your patience, and we apologize for any inconvenience.

## Below threshold (6)

- `digitalocean` [Multiple Services in NYC2](https://stspg.io/tb7f2ntnj8hm) — 1h 58m
- `digitalocean` [Let's Encrypt Outage Affecting Certificate Issuance and Managed Databases Operations](https://stspg.io/rw9m9n69q6tg) — 1h 8m
- `datadog` [Delays in APM-based Monitor Notifications](https://stspg.io/rww57bn61jkp) — 1h 38m
- `stripe` [Dashboard authentication errors](https://stspg.io/9559m3z4g1cq) — 45m
- `stripe` [Elevated Authentication Requests](https://stspg.io/yvt9gyxgydkb) — 39m
- `stripe` [Elevated API error rates from 18:15 - 18:34 UTC](https://stspg.io/vwp5wh5qbqlv) — 0m

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-0dc1b47ec8d11c330ab2fbab"></a>

- Canonical ID: `ci-0dc1b47ec8d11c330ab2fbab` (1 source records).
  - [slack/1551](../../data/2026/W19/slack.json); vendor URL: https://slack-status.com/2026-05/d284d808ed66c511
- Start time basis: `vendor_reported`.

<a id="ci-5c354d99da28f970835e2b89"></a>

- Canonical ID: `ci-5c354d99da28f970835e2b89` (1 source records).
  - [datadog/wjrkhvg9ydc1](../../data/2026/W19/datadog.json); vendor URL: https://stspg.io/rbj54mc4nhgs
- Start time basis: `vendor_reported`.

<a id="ci-487b87fdf7dd1f03fcc17179"></a>

- Canonical ID: `ci-487b87fdf7dd1f03fcc17179` (1 source records).
  - [digitalocean/h9t3sfjx36l3](../../data/2026/W19/digitalocean.json); vendor URL: https://stspg.io/tb7f2ntnj8hm
- Start time basis: `vendor_reported`.

<a id="ci-fc970ac075e258fe81586828"></a>

- Canonical ID: `ci-fc970ac075e258fe81586828` (1 source records).
  - [slack/1549](../../data/2026/W19/slack.json); vendor URL: https://slack-status.com/2026-05/bedf971f90be70a6
- Start time basis: `vendor_reported`.

<a id="ci-c8d23cd72fab87ac676bd3b4"></a>

- Canonical ID: `ci-c8d23cd72fab87ac676bd3b4` (1 source records).
  - [slack/1550](../../data/2026/W19/slack.json); vendor URL: https://slack-status.com/2026-05/02b2d5b45d656512
- Start time basis: `vendor_reported`.

<a id="ci-daf96497e756797cce502625"></a>

- Canonical ID: `ci-daf96497e756797cce502625` (1 source records).
  - [digitalocean/p3zyn7c4jyx7](../../data/2026/W19/digitalocean.json); vendor URL: https://stspg.io/rw9m9n69q6tg
- Start time basis: `vendor_reported`.

<a id="ci-7e13c723b5a09b2249a61506"></a>

- Canonical ID: `ci-7e13c723b5a09b2249a61506` (1 source records).
  - [circleci/d4gch9k0tcsq](../../data/2026/W19/circleci.json); vendor URL: https://stspg.io/pttscyh3qsv1
- Start time basis: `vendor_reported`.

<a id="ci-3403f15a18fedc3fe252dd01"></a>

- Canonical ID: `ci-3403f15a18fedc3fe252dd01` (1 source records).
  - [digitalocean/7hb39w1wzmwf](../../data/2026/W19/digitalocean.json); vendor URL: https://stspg.io/tyb995vv9ykq
- Start time basis: `vendor_reported`.

<a id="ci-abfa5e4d1bc9f1136cfa1f47"></a>

- Canonical ID: `ci-abfa5e4d1bc9f1136cfa1f47` (1 source records).
  - [datadog/s1mcpzcxhr9x](../../data/2026/W19/datadog.json); vendor URL: https://stspg.io/rww57bn61jkp
- Start time basis: `vendor_reported`.

<a id="ci-93543a2674deec8e9ac64f5d"></a>

- Canonical ID: `ci-93543a2674deec8e9ac64f5d` (1 source records).
  - [stripe/jf3qmzzgq9k3](../../data/2026/W19/stripe.json); vendor URL: https://stspg.io/9559m3z4g1cq
- Start time basis: `vendor_reported`.

<a id="ci-5d806369ab86015a7cd21935"></a>

- Canonical ID: `ci-5d806369ab86015a7cd21935` (1 source records).
  - [stripe/c1pk36f2kh25](../../data/2026/W19/stripe.json); vendor URL: https://stspg.io/yvt9gyxgydkb
- Start time basis: `vendor_reported`.

<a id="ci-753245b59eda340e89bceebe"></a>

- Canonical ID: `ci-753245b59eda340e89bceebe` (1 source records).
  - [circleci/68kk6515ztfp](../../data/2026/W19/circleci.json); vendor URL: https://stspg.io/m852c4blmtsk
- Start time basis: `vendor_reported`.

<a id="ci-baae7520879466714f047e70"></a>

- Canonical ID: `ci-baae7520879466714f047e70` (1 source records).
  - [stripe/s4ps5f7zf4zh](../../data/2026/W19/stripe.json); vendor URL: https://stspg.io/vwp5wh5qbqlv
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
| slack | ok | 30 | 0 |
| stripe | ok | 50 | 0 |
| yandexcloud | FETCH FAILED: GET https://status.cloud.yandex.ru/api/incidents?lang=en&page=4: status 403 | 0 | 0 |
