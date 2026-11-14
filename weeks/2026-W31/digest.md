# Week 2026-W31

2026-07-27 to 2026-08-02 (UTC). Generated 2026-11-14 12:36.

51 canonical incidents, 19 above threshold. 2 of 14 sources unavailable: datadog, yandexcloud.

51 source records (53 saved revisions) → 51 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | grafana | [Issues with Billing/Usage Dashboard Metrics and Panels.](https://stspg.io/fz2tz9xsqx46) | 2026-07-30 13:15 | 2h 1m | major | AWS Australia - prod-ap-southeast-2: Ingestion, AWS Australia - prod-ap-southeast-2: Querying, AWS Brazil - prod-sa-east-1: Ingestion, AWS Brazil - prod-sa-east-1: Querying, AWS Canada - prod-ca-east-0: Ingestion, AWS Canada - prod-ca-east-0: Querying, AWS Germany - prod-eu-west-2: Ingestion, AWS Germany - prod-eu-west-2: Querying, AWS India - prod-ap-south-1: Ingestion, AWS India - prod-ap-south-1: Querying, AWS Ireland - prod-eu-west-6, AWS Japan - prod-ap-northeast-0: Ingestion, AWS Japan - prod-ap-northeast-0: Querying, AWS Singapore - prod-ap-southeast-1: Ingestion, AWS Singapore - prod-ap-southeast-1: Querying, AWS Sweden - prod-eu-north-0: Ingestion, AWS Sweden - prod-eu-north-0: Querying, AWS Switzerland - prod-eu-central-0, AWS UAE - prod-me-central-1: Ingestion, AWS UAE - prod-me-central-1: Querying, AWS US East - prod-us-east-0: Ingestion, AWS US East - prod-us-east-0: Querying, AWS US West - prod-us-west-0: Ingestion, AWS US West - prod-us-west-0: Querying, Azure Netherlands - prod-eu-west-3: Ingestion, Azure Netherlands - prod-eu-west-3: Querying, Azure US Central - prod-us-central-7: Ingestion, Azure US Central - prod-us-central-7: Querying, Azure US Central - us-central2: Ingestion, Azure US Central - us-central2: Querying, Azure US Central - us-central7: Ingestion, Federal Cloud - AWS US Gov West, GCP Australia - prod-au-southeast-0: Ingestion, GCP Australia - prod-au-southeast-0: Querying, GCP Belgium - prod-eu-west-0: Ingestion, GCP Belgium - prod-eu-west-0: Querying, GCP Brazil - prod-sa-east-0: Ingestion, GCP Brazil - prod-sa-east-0: Querying, GCP India - prod-ap-south-0: Ingestion, GCP India - prod-ap-south-0: Querying, GCP Singapore - prod-ap-southeast-0: Ingestion, GCP Singapore - prod-ap-southeast-0: Querying, GCP UK - prod-gb-south-0: Ingestion, GCP UK - prod-gb-south-0: Querying, GCP US Central - prod-us-central-0: Ingestion, GCP US Central - prod-us-central-0: Querying, GCP US Central - prod-us-central-5: Ingestion, GCP US Central - prod-us-central-5: Querying, GCS US - cortex-prod-04: Ingestion, GCS US - cortex-prod-04: Querying | 6050 = 121 × 50 × 1.0 |
| 2 | cloudflare | [Network Performance Issues in Hamburg, Germany](https://www.cloudflarestatus.com/incidents/3ywn8wy3kqh8) | 2026-07-31 15:20 | 23h 58m | minor | Hamburg, Germany - (HAM) | 2588 = 1438 × 1 × 1.8 |
| 3 | anthropic | [Elevated errors across many models](https://stspg.io/cpzpstxn0z6k) | 2026-07-30 05:57 | 4h 51m | major | Claude API (api.anthropic.com), Claude Code, Claude Cowork, claude.ai | 1746 = 291 × 4 × 1.5 |
| 4 | slack | [Feature Degradation Affecting the Receipt of Emails in Slack](https://slack-status.com/2026-07/6b1014d8bfb3afc0) | 2026-07-28 02:13 | 20h 55m | unknown | Apps/Integrations/APIs | 1506 = 1255 × 1 × 1.2 |
| 5 | digitalocean | [Agent Platform Requests Returning HTTP 500 Errors](https://stspg.io/54mh0lzk9733) | 2026-07-27 18:31 | 20h 52m | major | Agent Runtime | 1002 = 1252 × 1 × 0.8 |
| 6 | yandexcloud | [The problem is with the services](https://status.yandex.cloud/en/incidents/1901) | 2026-08-02 12:05 | 7h 54m | minor | Compute Cloud | 711 = 474 × 1 × 1.5 |
| 7 | grafana | [PDC Authentication Issues](https://stspg.io/p8c7l7pzd5f5) | 2026-07-30 14:45 | 2h 54m | major | AWS Germany - prod-eu-west-2, AWS US East - prod-us-east-0, GCP US Central - prod-us-central-0, GCP US East - prod-us-east-1 | 696 = 174 × 4 × 1.0 |
| 8 | digitalocean | [Response Degradation Impacting Kimi-K3](https://stspg.io/ynhlvpy6ctcm) | 2026-07-29 17:56 | 12h 19m | minor | Model Services | 591 = 739 × 1 × 0.8 |
| 9 | cloudflare | [Increased HTTP 5XX Errors in IAD](https://www.cloudflarestatus.com/incidents/l63k37vrcd9c) | 2026-07-31 18:30 | 4h 30m | none | — | 486 = 270 × 1 × 1.8 |
| 10 | cloudflare | [Possible Network Congestion in North America region](https://www.cloudflarestatus.com/incidents/z33kqrzksxqy) | 2026-07-30 05:48 | 4h 20m | minor | Network | 468 = 260 × 1 × 1.8 |
| 11 | yandexcloud | [The problem is with DataSphere services](https://status.yandex.cloud/en/incidents/1893) | 2026-07-28 03:51 | 4h 29m | major | DataSphere | 404 = 269 × 1 × 1.5 |
| 12 | cloudflare | [Possible Network Congestion between Singapore and Tokyo](https://www.cloudflarestatus.com/incidents/h1kxms9wcbwj) | 2026-07-29 08:29 | 3h 16m | minor | — | 353 = 196 × 1 × 1.8 |
| 13 | anthropic | [Elevated errors across all models](https://stspg.io/6xr5mmpjs1k3) | 2026-07-29 19:49 | 2h 46m | critical | — | 249 = 166 × 1 × 1.5 |
| 14 | cloudflare | [Increased HTTP 5xx Errors for us-east-1-aws](https://www.cloudflarestatus.com/incidents/s18kw61f2ht5) | 2026-07-31 02:01 | 2h 18m | minor | Network | 248 = 138 × 1 × 1.8 |
| 15 | grafana | [IRM Performance Degradation in EU Region](https://stspg.io/35m2lfvv6h87) | 2026-07-29 21:46 | 2h 52m | minor | Incident Management and Response (IRM) | 172 = 172 × 1 × 1.0 |
| 16 | github | [Incident with Actions](https://stspg.io/8nsh6820sff9) | 2026-07-29 15:26 | 34m | major | Actions | 51 = 34 × 1 × 1.5 |
| 17 | cloudflare | [Cloudflare Workers — Errors Deploying Workers Scripts](https://www.cloudflarestatus.com/incidents/z325n2dc14tk) | 2026-07-29 20:03 | 27m | major | Workers | 49 = 27 × 1 × 1.8 |
| 18 | datadog | [Synthetics Test Results Delayed](https://stspg.io/9trx2m1xsd4h) | 2026-07-27 19:05 | 46m | major | Synthetics | 46 = 46 × 1 × 1.0 |
| 19 | grafana | [Partial OTLP Write Outage in prod-us-east-3](https://stspg.io/05bt4v67qbyk) | 2026-07-28 19:40 | 3m | major | AWS US East - prod-us-east-3 | 3 = 3 × 1 × 1.0 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. grafana — Issues with Billing/Usage Dashboard Metrics and Panels.

- URL: https://stspg.io/fz2tz9xsqx46
- Started: 2026-07-30 13:15 UTC
- Resolved: 2026-07-30 15:17 UTC
- Duration: 2h 1m
- Impact: major, status: resolved
- Components: AWS Australia - prod-ap-southeast-2: Ingestion, AWS Australia - prod-ap-southeast-2: Querying, AWS Brazil - prod-sa-east-1: Ingestion, AWS Brazil - prod-sa-east-1: Querying, AWS Canada - prod-ca-east-0: Ingestion, AWS Canada - prod-ca-east-0: Querying, AWS Germany - prod-eu-west-2: Ingestion, AWS Germany - prod-eu-west-2: Querying, AWS India - prod-ap-south-1: Ingestion, AWS India - prod-ap-south-1: Querying, AWS Ireland - prod-eu-west-6, AWS Japan - prod-ap-northeast-0: Ingestion, AWS Japan - prod-ap-northeast-0: Querying, AWS Singapore - prod-ap-southeast-1: Ingestion, AWS Singapore - prod-ap-southeast-1: Querying, AWS Sweden - prod-eu-north-0: Ingestion, AWS Sweden - prod-eu-north-0: Querying, AWS Switzerland - prod-eu-central-0, AWS UAE - prod-me-central-1: Ingestion, AWS UAE - prod-me-central-1: Querying, AWS US East - prod-us-east-0: Ingestion, AWS US East - prod-us-east-0: Querying, AWS US West - prod-us-west-0: Ingestion, AWS US West - prod-us-west-0: Querying, Azure Netherlands - prod-eu-west-3: Ingestion, Azure Netherlands - prod-eu-west-3: Querying, Azure US Central - prod-us-central-7: Ingestion, Azure US Central - prod-us-central-7: Querying, Azure US Central - us-central2: Ingestion, Azure US Central - us-central2: Querying, Azure US Central - us-central7: Ingestion, Federal Cloud - AWS US Gov West, GCP Australia - prod-au-southeast-0: Ingestion, GCP Australia - prod-au-southeast-0: Querying, GCP Belgium - prod-eu-west-0: Ingestion, GCP Belgium - prod-eu-west-0: Querying, GCP Brazil - prod-sa-east-0: Ingestion, GCP Brazil - prod-sa-east-0: Querying, GCP India - prod-ap-south-0: Ingestion, GCP India - prod-ap-south-0: Querying, GCP Singapore - prod-ap-southeast-0: Ingestion, GCP Singapore - prod-ap-southeast-0: Querying, GCP UK - prod-gb-south-0: Ingestion, GCP UK - prod-gb-south-0: Querying, GCP US Central - prod-us-central-0: Ingestion, GCP US Central - prod-us-central-0: Querying, GCP US Central - prod-us-central-5: Ingestion, GCP US Central - prod-us-central-5: Querying, GCS US - cortex-prod-04: Ingestion, GCS US - cortex-prod-04: Querying
- Canonical incident: [ci-aa15d823569b32f896053338](#ci-aa15d823569b32f896053338)

- `2026-07-30 13:15` **investigating** — We are investigating an outage for the Billing / Usage dashboard metrics and panels. This appears to be partially affecting organizations using Grafana Cloud. We are working on identifying and resolving the issue
- `2026-07-30 13:50` **identified** — We have identified the issue, and are working on deploying a fix.
- `2026-07-30 14:34` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-07-30 15:17` **resolved** — This incident has been resolved.

### 2. cloudflare — Network Performance Issues in Hamburg, Germany

- URL: https://www.cloudflarestatus.com/incidents/3ywn8wy3kqh8
- Started: 2026-07-31 15:20 UTC
- Resolved: 2026-08-01 15:18 UTC
- Duration: 23h 58m
- Impact: minor, status: resolved
- Components: Hamburg, Germany - (HAM)
- Canonical incident: [ci-cc131d6ee61cf325905371da](#ci-cc131d6ee61cf325905371da)

- `2026-07-31 15:20` **identified** — Cloudflare is investigating issues with network performance in Hamburg, Germany (HAM). Customers routing through this location may experience request errors or failures. We have identified the problem and are working on a fix.
- `2026-07-31 16:58` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-08-01 15:18` **resolved** — This incident has been resolved.

### 3. anthropic — Elevated errors across many models

- URL: https://stspg.io/cpzpstxn0z6k
- Started: 2026-07-30 05:57 UTC
- Resolved: 2026-07-30 10:48 UTC
- Duration: 4h 51m
- Impact: major, status: resolved
- Components: Claude API (api.anthropic.com), Claude Code, Claude Cowork, claude.ai
- Canonical incident: [ci-d053ed79dc135551f56c34ab](#ci-d053ed79dc135551f56c34ab)

- `2026-07-30 05:57` **investigating** — We are currently investigating this issue.
- `2026-07-30 06:26` **identified** — All models except Opus 5 have recovered. We're working to recover Opus 5 as well.
- `2026-07-30 07:33` **identified** — As of 7:32 UTC / 0:32 PT, Opus 5 errors are back to baseline but we're seeing again Sonnet 5's errors up.
- `2026-07-30 08:06` **identified** — As of 8:05 UTC / 1:05 PT, Sonnet 5 errors are back to baseline, but we're seeing an increased error rate on Fable 5.
- `2026-07-30 08:33` **identified** — As of 01:33 PT / 08:33 UTC we are back to errors across all models.
- `2026-07-30 10:03` **monitoring** — As of 2:58 PT / 9:58 UTC the errors are back to baseline.
- `2026-07-30 10:23` **monitoring** — We are also investigating issues with Opus 4.7 and Opus 4.5.
- `2026-07-30 10:48` **resolved** — This incident has been resolved.

### 4. slack — Feature Degradation Affecting the Receipt of Emails in Slack

- URL: https://slack-status.com/2026-07/6b1014d8bfb3afc0
- Started: 2026-07-28 02:13 UTC
- Resolved: 2026-07-28 23:09 UTC
- Duration: 20h 55m
- Impact: unknown, status: resolved
- Components: Apps/Integrations/APIs
- Canonical incident: [ci-0626e37ed376c79ddab893c3](#ci-0626e37ed376c79ddab893c3)

- `2026-07-28 02:13` **unknown** — We're currently investigating an issue affecting users' ability to receive incoming emails into Slack, caused by some issues with our email ingestion pipeline. We're sorry for any inconvenience this is causing, our team is actively working on a fix and we'll share an update as soon as we have more information.
- `2026-07-28 03:25` **unknown** — We identified the cause of the issue preventing users from receiving incoming emails. We are deploying a fix and will share a further update once deployment and verification are complete.
- `2026-07-28 03:43` **unknown** — The fix has been successfully deployed, and validation is now complete. The issue preventing users from receiving incoming emails into Slack is fully resolved, and normal service functionality has been restored. We apologize for any disruptions to your day.

### 5. digitalocean — Agent Platform Requests Returning HTTP 500 Errors

- URL: https://stspg.io/54mh0lzk9733
- Started: 2026-07-27 18:31 UTC
- Resolved: 2026-07-28 15:24 UTC
- Duration: 20h 52m
- Impact: major, status: resolved
- Components: Agent Runtime
- Canonical incident: [ci-6c1b5d39356b2854e0b85c2c](#ci-6c1b5d39356b2854e0b85c2c)

- `2026-07-27 18:31` **investigating** — We are currently investigating this issue.
- `2026-07-27 19:44` **identified** — The issue has been identified and a fix is being implemented.
- `2026-07-28 15:24` **resolved** — This incident has been resolved.

### 6. yandexcloud — The problem is with the services

- URL: https://status.yandex.cloud/en/incidents/1901
- Started: 2026-08-02 12:05 UTC
- Resolved: 2026-08-02 20:00 UTC
- Duration: 7h 54m
- Impact: minor, status: resolved
- Components: Compute Cloud
- Canonical incident: [ci-d8ddfc68cb0cb901e10ad312](#ci-d8ddfc68cb0cb901e10ad312)

- `2026-08-02 08:03` **investigating** — Have you noticed a degradation in the performance of creating backups, images, and restoring disks from them
- `2026-08-03 11:28` **resolved** — At 23:00 Moscow time, the problem was fixed, and now the Cloud services are working correctly.

### 7. grafana — PDC Authentication Issues

- URL: https://stspg.io/p8c7l7pzd5f5
- Started: 2026-07-30 14:45 UTC
- Resolved: 2026-07-30 17:40 UTC
- Duration: 2h 54m
- Impact: major, status: resolved
- Components: AWS Germany - prod-eu-west-2, AWS US East - prod-us-east-0, GCP US Central - prod-us-central-0, GCP US East - prod-us-east-1
- Canonical incident: [ci-1c0a1a69b9c74833e3b9f8c6](#ci-1c0a1a69b9c74833e3b9f8c6)

- `2026-07-30 14:45` **investigating** — We are currently investigating an issue impacting PDC authentication. We will provide additional updates as they become available.
- `2026-07-30 14:51` **investigating** — We are continuing to investigate this issue.
- `2026-07-30 15:11` **investigating** — We are continuing to investigate this issue.
- `2026-07-30 15:56` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-07-30 17:40` **resolved** — This incident has been resolved.

### 8. digitalocean — Response Degradation Impacting Kimi-K3

- URL: https://stspg.io/ynhlvpy6ctcm
- Started: 2026-07-29 17:56 UTC
- Resolved: 2026-07-30 06:15 UTC
- Duration: 12h 19m
- Impact: minor, status: resolved
- Components: Model Services
- Canonical incident: [ci-32f3b2cfa3631f09150f2ba6](#ci-32f3b2cfa3631f09150f2ba6)

- `2026-07-29 17:56` **investigating** — We are currently investigating this issue.
- `2026-07-29 21:41` **investigating** — Our Engineering team is investigating reports of empty or broken responses impacting Kimi-K3. At this point, users may experience degraded response quality, specifically empty or broken outputs (including repeated characters), when querying the Kimi-K3 service. We apologize for the inconvenience and will share an update once we have more information.
- `2026-07-30 06:15` **resolved** — Our Engineering team has resolved the issue causing empty or broken responses from Kimi-K3. If you continue to experience any problems, please open a ticket with our Support team. We apologize for any inconvenience this may have caused.

### 9. cloudflare — Increased HTTP 5XX Errors in IAD

- URL: https://www.cloudflarestatus.com/incidents/l63k37vrcd9c
- Started: 2026-07-31 18:30 UTC
- Resolved: 2026-07-31 23:00 UTC
- Duration: 4h 30m
- Impact: none, status: resolved
- Components: —
- Canonical incident: [ci-206ee228089011fa6d50ef08](#ci-206ee228089011fa6d50ef08)

- `2026-08-01 01:05` **resolved** — Cloudflare identified and resolved an issue causing an increased level of 5XX errors in Ashburn, US (IAD) from 31st July 2026 18:45 UTC to 23:01 UTC.

### 10. cloudflare — Possible Network Congestion in North America region

- URL: https://www.cloudflarestatus.com/incidents/z33kqrzksxqy
- Started: 2026-07-30 05:48 UTC
- Resolved: 2026-07-30 10:08 UTC
- Duration: 4h 20m
- Impact: minor, status: resolved
- Components: Network
- Canonical incident: [ci-3cbe5ed366a4889cc769ece1](#ci-3cbe5ed366a4889cc769ece1)

- `2026-07-30 05:48` **investigating** — Cloudflare is observing possible network congestion in the North America region. Customers may see elevated latency and connectivity issues. We are working to mitigate impact to Internet users in this region.
- `2026-07-30 07:51` **monitoring** — Cloudflare identified that the Network degradation was due to maintenance conducted by fiber providers. The congestion has been alleviated by rerouting traffic. We are currently monitoring the situation and will follow up with more updates shortly.
- `2026-07-30 07:57` **monitoring** — We are continuing to monitor for any further issues.
- `2026-07-30 10:08` **resolved** — This incident has been resolved.

### 11. yandexcloud — The problem is with DataSphere services

- URL: https://status.yandex.cloud/en/incidents/1893
- Started: 2026-07-28 03:51 UTC
- Resolved: 2026-07-28 08:20 UTC
- Duration: 4h 29m
- Impact: major, status: resolved
- Components: DataSphere
- Canonical incident: [ci-d8376d26f24ac734f18fc303](#ci-d8376d26f24ac734f18fc303)

- `2026-07-28 06:29` **investigating** — We noticed a problem with opening projects in DataSphere. The developers are already aware of it and are fixing it.
- `2026-07-28 08:40` **resolved** — The service has been restored, and we no longer see any errors.

### 12. cloudflare — Possible Network Congestion between Singapore and Tokyo

- URL: https://www.cloudflarestatus.com/incidents/h1kxms9wcbwj
- Started: 2026-07-29 08:29 UTC
- Resolved: 2026-07-29 11:45 UTC
- Duration: 3h 16m
- Impact: minor, status: resolved
- Components: —
- Canonical incident: [ci-f3cd0c696ac0cde86d0e62fd](#ci-f3cd0c696ac0cde86d0e62fd)

- `2026-07-29 08:29` **investigating** — Cloudflare is observing possible network congestion between Singapore and Tokyo. We are working to mitigate impact to Internet users in this region.
- `2026-07-29 10:29` **identified** — The issue has been identified and a fix is being implemented.
- `2026-07-29 11:44` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-07-29 11:45` **resolved** — This incident has been resolved.

### 13. anthropic — Elevated errors across all models

- URL: https://stspg.io/6xr5mmpjs1k3
- Started: 2026-07-29 19:49 UTC
- Resolved: 2026-07-29 22:36 UTC
- Duration: 2h 46m
- Impact: critical, status: resolved
- Components: —
- Canonical incident: [ci-c9e33c2ec37f37b7b10a83ad](#ci-c9e33c2ec37f37b7b10a83ad)

- `2026-07-29 19:49` **investigating** — We are currently investigating this issue.
- `2026-07-29 20:33` **identified** — We have identified an issue resulting in elevated errors across multiple models, and are working to resolve this. We will provide an update as soon as possible.
- `2026-07-29 21:38` **identified** — We are continuing to work to fully resolve issues resulting in elevated requests and latency to Claude models. At this time, we are seeing recovery across most models. We will provide an additional update as soon as possible.
- `2026-07-29 22:20` **monitoring** — From 12:45 PT / 19:45 UTC through to 1:26 PT / 21:26 UTC, we saw elevated rates of errors across Claude models. We have seen recovery in success rates across all models and are monitoring closely to ensure no further issues.
- `2026-07-29 22:36` **resolved** — This issue has been resolved.

### 14. cloudflare — Increased HTTP 5xx Errors for us-east-1-aws

- URL: https://www.cloudflarestatus.com/incidents/s18kw61f2ht5
- Started: 2026-07-31 02:01 UTC
- Resolved: 2026-07-31 04:19 UTC
- Duration: 2h 18m
- Impact: minor, status: resolved
- Components: Network
- Canonical incident: [ci-738325c596fe150d64edc826](#ci-738325c596fe150d64edc826)

- `2026-07-31 02:01` **investigating** — Cloudflare is investigating an increased level of intermittent HTTP 5xx errors for customers using AWS region us-east-1-aws. We are working to analyse and mitigate this problem. More updates to follow shortly.
- `2026-07-31 02:34` **identified** — The issue has been identified and a fix is being implemented.
- `2026-07-31 04:08` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-07-31 04:19` **resolved** — This incident has been resolved.

### 15. grafana — IRM Performance Degradation in EU Region

- URL: https://stspg.io/35m2lfvv6h87
- Started: 2026-07-29 21:46 UTC
- Resolved: 2026-07-30 00:39 UTC
- Duration: 2h 52m
- Impact: minor, status: resolved
- Components: Incident Management and Response (IRM)
- Canonical incident: [ci-69620b632b63a889000f628f](#ci-69620b632b63a889000f628f)

- `2026-07-29 21:46` **investigating** — We are currently investigating an issue affecting Grafana IRM in our EU region. Users may experience a degraded or unresponsive IRM UI, reduced availability of the public API, and delays in notifications. Our engineering team is actively investigating the issue and working to restore normal service. We will provide another update as more information becomes available.
- `2026-07-29 23:07` **identified** — We have identified the underlying issue affecting Grafana IRM in our EU region and are actively working to restore normal service. Users may continue to experience a degraded or unresponsive IRM UI, reduced availability of the public API, and delays in notifications while mitigation efforts are underway. Our engineering team is working to restore full functionality as quickly as possible. We will provide another update as more information becomes available.
- `2026-07-30 00:39` **resolved** — The issue affecting Grafana IRM in the EU region has been resolved. The IRM UI, public API, and alert notification processing have been restored and are operating normally.

### 16. github — Incident with Actions

- URL: https://stspg.io/8nsh6820sff9
- Started: 2026-07-29 15:26 UTC
- Resolved: 2026-07-29 16:00 UTC
- Duration: 34m
- Impact: major, status: resolved
- Components: Actions
- Canonical incident: [ci-789023240211381935416cb7](#ci-789023240211381935416cb7)

- `2026-07-29 15:26` **investigating** — We are investigating reports of degraded availability for Actions
- `2026-07-29 15:34` **investigating** — We are investigating an issue affecting GitHub Actions. Some customers may experience timeouts or failures with runner registration and workflow runs may be delayed during startup. Our team is actively working to mitigate the impact by scaling capacity across additional infrastructure.
- `2026-07-29 15:40` **monitoring** — The degradation affecting Actions has been mitigated. We are monitoring to ensure stability.
- `2026-07-29 16:00` **resolved** — On July 29, 2026, from 14:51 UTC to 15:28 UTC, GitHub Actions experienced elevated REST API request timeouts and errors, failures registering runners, and delayed workflow run starts for customers whose traffic was served by a single infrastructure site. This was caused by an under-provisioned internal Actions service in that site: under increased load its instances ran out of memory and became unresponsive, and because Actions API requests wait synchronously on that service, requests routed through the affected site stalled and timed out. During the incident, approximately 2% of workflows were delayed. Requests served by other sites remained unaffected. Both standard and larger hosted runners routed through the affected site could see delayed job starts. <br /><br />The issue was mitigated by scaling out the runner-administration service in the affected site and increasing the replica count, which restored API availability and returned workflow run starts to normal. We are working to add horizontal autoscaling, memory-saturation alerting, and scaling-forecast monitoring for this service, along with responder playbooks, to reduce the likelihood of similar issues in the future.

### 17. cloudflare — Cloudflare Workers — Errors Deploying Workers Scripts

- URL: https://www.cloudflarestatus.com/incidents/z325n2dc14tk
- Started: 2026-07-29 20:03 UTC
- Resolved: 2026-07-29 20:31 UTC
- Duration: 27m
- Impact: major, status: resolved
- Components: Workers
- Canonical incident: [ci-60fe110deb5da1ad3190390c](#ci-60fe110deb5da1ad3190390c)

- `2026-07-29 20:03` **identified** — Cloudflare is investigating an issue where a subset of Workers script deployments are failing with errors. Customers may experience failures when attempting to deploy a Worker script via Wrangler or the Workers API. Workers that are already deployed and running in production are not affected. A fix is being prepared.
- `2026-07-29 20:04` **identified** — We are continuing to work on a fix for this issue.
- `2026-07-29 20:31` **resolved** — This incident has been resolved.

### 18. datadog — Synthetics Test Results Delayed

- URL: https://stspg.io/9trx2m1xsd4h
- Started: 2026-07-27 19:05 UTC
- Resolved: 2026-07-27 19:52 UTC
- Duration: 46m
- Impact: major, status: resolved
- Components: Synthetics
- Canonical incident: [ci-c9eac398ed12abcf135819be](#ci-c9eac398ed12abcf135819be)

- `2026-07-27 19:05` **investigating** — We are currenting investigating an issue running Synthetic tests.
- `2026-07-27 19:25` **investigating** — We're continuing to investigate an issue publishing Synthetic results. Synthetic tests are running and results are delayed.
- `2026-07-27 19:36` **identified** — We've identified the issue publishing Synthetic results. Synthetic tests are continuing to run with delayed results.
- `2026-07-27 19:46` **monitoring** — A fix has been implemented and we're monitoring the results.
- `2026-07-27 19:52` **resolved** — This incident has been resolved.

### 19. grafana — Partial OTLP Write Outage in prod-us-east-3

- URL: https://stspg.io/05bt4v67qbyk
- Started: 2026-07-28 19:40 UTC
- Resolved: 2026-07-28 19:43 UTC
- Duration: 3m
- Impact: major, status: resolved
- Components: AWS US East - prod-us-east-3
- Canonical incident: [ci-c542a4c1a90798d9b296f0f9](#ci-c542a4c1a90798d9b296f0f9)

- `2026-07-28 19:40` **investigating** — We are investigating an issue affecting OTLP ingestion in the prod-us-east-3 region. Customers may experience intermittent failures when writing telemetry to the OTLP endpoint. Based on current information, logs are confirmed to be affected, and metrics and traces may also be impacted. Our investigation indicates intermittent failures began around 08:00 UTC, with the primary period of impact occurring between 15:30 UTC and 17:45 UTC. We are continuing to investigate the scope and root cause of this issue and will provide additional updates as more information becomes available.
- `2026-07-28 19:43` **resolved** — The issue affecting OTLP ingestion in the prod-us-east-3 region has been resolved. Between 13:30 UTC and 17:45 UTC, some customers experienced intermittent failures when writing telemetry to the OTLP endpoint. Logs were confirmed to be affected, and metrics and traces may also have experienced intermittent ingestion failures. Our investigation has concluded, and the affected services have recovered.

## Below threshold (32)

- `cloudflare` [Cloudflare API Availability Reduced Availability](https://www.cloudflarestatus.com/incidents/pgcdxsjxkl0q) — 1h 10m
- `anthropic` [Elevated errors on Claude Opus 5](https://stspg.io/1ctlvqtq2z6k) — 1h 2m
- `cloudflare` [Page Shield Email Notification Issues](https://www.cloudflarestatus.com/incidents/qwd9dhqbcg0g) — 1h 36m
- `cloudflare` [Increased HTTP Errors in London](https://www.cloudflarestatus.com/incidents/k17p9vnmhkvp) — 1h 34m
- `anthropic` [Elevated errors on Claude Opus 5 and Haiku 4.5](https://stspg.io/996bl17pqbbx) — 56m
- `grafana` [Degraded Performance: Stack Provisioning Failures within certain reigons (PDC Setup)](https://stspg.io/gkrhww23994q) — 1h 22m
- `datadog` [Real user monitoring, distribution and custom processes metrics delayed](https://stspg.io/0y0jgdxjp7hh) — 1h 46m
- `anthropic` [Elevated errors on Claude Opus 5](https://stspg.io/b7yvdnh7g4k6) — 48m
- `anthropic` [Degraded performance on Claude Sonnet 5](https://stspg.io/5tnhyrw0vnfq) — 46m
- `anthropic` [Degraded performance on Claude Opus 4.8](https://stspg.io/5d61g084jy31) — 40m
- `github` [Incident with Copilot AI Model Providers](https://stspg.io/tdw1hjz3lj91) — 1h 44m
- `cloudflare` [Sensitive Data Detection Configuration Issue](https://www.cloudflarestatus.com/incidents/shq1h521dfl3) — 1h 14m
- `openai` [Elevated errors affecting ChatGPT conversations](https://status.openai.com/incidents/01KYSQ7PC81PV3PZE3ZEVFK6K5) — 1h 25m
- `openai` [Elevated error rates with the invalid_prompt error code](https://status.openai.com/incidents/01KYNNX2TQ3PV0KPJ9PTRRPX6C) — 1h 24m
- `openai` [Image generation unavailable in ChatGPT](https://status.openai.com/incidents/01KYJFG4S86K8FMTB4N6BDZS7F) — 1h 19m
- `github` [Degraded availability GPT 5.6 Luna](https://stspg.io/pj5lgggbc2rb) — 1h 13m
- `cloudflare` [Network Performance Issues in Singapore](https://www.cloudflarestatus.com/incidents/wcln78w0lx1h) — 1h 0m
- `github` [Copilot model Claude Fable 5 experiencing elevated errors](https://stspg.io/45g8r57kv3n9) — 1h 4m
- `stripe` [Elevated API errors](https://stspg.io/l0fsvp40d3hy) — 34m
- `stripe` [Elevated card authorization errors](https://stspg.io/ykd350r61yh5) — 1h 4m
- `grafana` [K6 - Cloud test-run issues](https://stspg.io/x4fmhx2pym34) — 1h 5m
- `github` [Incident with Copilot AI Model Providers](https://stspg.io/z58g1g529hzz) — 41m
- `openai` [Elevated Image Generation Error Rate in ChatGPT](https://status.openai.com/incidents/01KYMJND2XYYBDVCF8DKA5DY3N) — 35m
- `datadog` [Delayed Evaluation of Service Check monitors](https://stspg.io/803q340l4b6k) — 42m
- `openai` [Enterprise & Education Chat Errors](https://status.openai.com/incidents/01KYVPNEY9736Z04D3X1NSFFF9) — 24m
- `github` [Incident with GraphQL API Requests](https://stspg.io/vr201n49yl53) — 15m
- `openai` [We’re seeing elevated latency, timeouts, and interrupted streaming across gpt 5.1 mini model and gpt 4.1 mini](https://status.openai.com/incidents/01KYJC3TDHZVSQFFPTC8X9RAHC) — unknown
- `grafana` [Grafana Cloud non-billing usage metrics gaps in select regions](https://stspg.io/5db188nt4pmn) — 0m
- `grafana` [Partial Read Outage for Loki in prod-us-east-4](https://stspg.io/mm2jxybgmg52) — unknown
- `cloudflare` [Network Performance Issues in Istanbul](https://www.cloudflarestatus.com/incidents/pgbznnvc2vtw) — 0m
- `stripe` [Elevated API errors](https://stspg.io/sbldmtv3y0jq) — 0m
- `cloudflare` [RealtimeKit socket connection slowness and failed meeting joins](https://www.cloudflarestatus.com/incidents/7czwy7vt3phg) — 0m

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-aa15d823569b32f896053338"></a>

- Canonical ID: `ci-aa15d823569b32f896053338` (1 source records).
  - [grafana/79nwp08zs9vj](../../data/2026/W31/grafana.json); vendor URL: https://stspg.io/fz2tz9xsqx46
- Start time basis: `vendor_reported`.

<a id="ci-cc131d6ee61cf325905371da"></a>

- Canonical ID: `ci-cc131d6ee61cf325905371da` (1 source records).
  - [cloudflare/3ywn8wy3kqh8](../../data/2026/W31/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/3ywn8wy3kqh8
- Start time basis: `vendor_reported`.

<a id="ci-d053ed79dc135551f56c34ab"></a>

- Canonical ID: `ci-d053ed79dc135551f56c34ab` (1 source records).
  - [anthropic/fsh2zzzl2c4l](../../data/2026/W31/anthropic.json); vendor URL: https://stspg.io/cpzpstxn0z6k
- Start time basis: `vendor_reported`.

<a id="ci-0626e37ed376c79ddab893c3"></a>

- Canonical ID: `ci-0626e37ed376c79ddab893c3` (1 source records).
  - [slack/1574](../../data/2026/W31/slack.json); vendor URL: https://slack-status.com/2026-07/6b1014d8bfb3afc0
- Start time basis: `vendor_reported`.

<a id="ci-6c1b5d39356b2854e0b85c2c"></a>

- Canonical ID: `ci-6c1b5d39356b2854e0b85c2c` (1 source records).
  - [digitalocean/fcfplv03vbwt](../../data/2026/W31/digitalocean.json); vendor URL: https://stspg.io/54mh0lzk9733
- Start time basis: `vendor_reported`.

<a id="ci-d8ddfc68cb0cb901e10ad312"></a>

- Canonical ID: `ci-d8ddfc68cb0cb901e10ad312` (1 source records).
  - [yandexcloud/1901](../../data/2026/W31/yandexcloud.json); vendor URL: https://status.yandex.cloud/en/incidents/1901
- Start time basis: `vendor_reported`.

<a id="ci-1c0a1a69b9c74833e3b9f8c6"></a>

- Canonical ID: `ci-1c0a1a69b9c74833e3b9f8c6` (1 source records).
  - [grafana/z5hw9b0tmjzm](../../data/2026/W31/grafana.json); vendor URL: https://stspg.io/p8c7l7pzd5f5
- Start time basis: `vendor_reported`.

<a id="ci-4f196410da9583fa39830b53"></a>

- Canonical ID: `ci-4f196410da9583fa39830b53` (1 source records).
  - [cloudflare/pgcdxsjxkl0q](../../data/2026/W31/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/pgcdxsjxkl0q
- Start time basis: `vendor_reported`.

<a id="ci-32f3b2cfa3631f09150f2ba6"></a>

- Canonical ID: `ci-32f3b2cfa3631f09150f2ba6` (1 source records).
  - [digitalocean/rdfp2h4syxz8](../../data/2026/W31/digitalocean.json); vendor URL: https://stspg.io/ynhlvpy6ctcm
- Start time basis: `vendor_reported`.

<a id="ci-206ee228089011fa6d50ef08"></a>

- Canonical ID: `ci-206ee228089011fa6d50ef08` (1 source records).
  - [cloudflare/l63k37vrcd9c](../../data/2026/W31/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/l63k37vrcd9c
- Start time basis: `vendor_reported`.

<a id="ci-3cbe5ed366a4889cc769ece1"></a>

- Canonical ID: `ci-3cbe5ed366a4889cc769ece1` (1 source records).
  - [cloudflare/z33kqrzksxqy](../../data/2026/W31/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/z33kqrzksxqy
- Start time basis: `vendor_reported`.

<a id="ci-d8376d26f24ac734f18fc303"></a>

- Canonical ID: `ci-d8376d26f24ac734f18fc303` (1 source records).
  - [yandexcloud/1893](../../data/2026/W31/yandexcloud.json); vendor URL: https://status.yandex.cloud/en/incidents/1893
- Start time basis: `vendor_reported`.

<a id="ci-7cebba8dc39457209f212730"></a>

- Canonical ID: `ci-7cebba8dc39457209f212730` (1 source records).
  - [anthropic/mfdtrknpxghq](../../data/2026/W31/anthropic.json); vendor URL: https://stspg.io/1ctlvqtq2z6k
- Start time basis: `vendor_reported`.

<a id="ci-f3cd0c696ac0cde86d0e62fd"></a>

- Canonical ID: `ci-f3cd0c696ac0cde86d0e62fd` (1 source records).
  - [cloudflare/h1kxms9wcbwj](../../data/2026/W31/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/h1kxms9wcbwj
- Start time basis: `vendor_reported`.

<a id="ci-8ae088d1a8e23f9c5df2db97"></a>

- Canonical ID: `ci-8ae088d1a8e23f9c5df2db97` (1 source records).
  - [cloudflare/qwd9dhqbcg0g](../../data/2026/W31/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/qwd9dhqbcg0g
- Start time basis: `vendor_reported`.

<a id="ci-8141c5f942bf33b7ca49d49d"></a>

- Canonical ID: `ci-8141c5f942bf33b7ca49d49d` (1 source records).
  - [cloudflare/k17p9vnmhkvp](../../data/2026/W31/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/k17p9vnmhkvp
- Start time basis: `vendor_reported`.

<a id="ci-677fd70bb78baefc64c21e4e"></a>

- Canonical ID: `ci-677fd70bb78baefc64c21e4e` (1 source records).
  - [anthropic/rkk5x44tndw9](../../data/2026/W31/anthropic.json); vendor URL: https://stspg.io/996bl17pqbbx
- Start time basis: `vendor_reported`.

<a id="ci-d02133eed5026a1c0c67c55d"></a>

- Canonical ID: `ci-d02133eed5026a1c0c67c55d` (1 source records).
  - [grafana/gxb577ll1b7x](../../data/2026/W31/grafana.json); vendor URL: https://stspg.io/gkrhww23994q
- Start time basis: `vendor_reported`.

<a id="ci-57ebec0b7f0035008588b32a"></a>

- Canonical ID: `ci-57ebec0b7f0035008588b32a` (1 source records).
  - [datadog/2c2287dz638f](../../data/2026/W31/datadog.json); vendor URL: https://stspg.io/0y0jgdxjp7hh
- Start time basis: `vendor_reported`.

<a id="ci-beb52142c58182163427d6dc"></a>

- Canonical ID: `ci-beb52142c58182163427d6dc` (1 source records).
  - [anthropic/lhqp09kxq7pb](../../data/2026/W31/anthropic.json); vendor URL: https://stspg.io/b7yvdnh7g4k6
- Start time basis: `vendor_reported`.

<a id="ci-958c0996d705ce015305b45b"></a>

- Canonical ID: `ci-958c0996d705ce015305b45b` (1 source records).
  - [anthropic/jq4x54h69z76](../../data/2026/W31/anthropic.json); vendor URL: https://stspg.io/5tnhyrw0vnfq
- Start time basis: `vendor_reported`.

<a id="ci-c9e33c2ec37f37b7b10a83ad"></a>

- Canonical ID: `ci-c9e33c2ec37f37b7b10a83ad` (1 source records).
  - [anthropic/q2kg8n613kr3](../../data/2026/W31/anthropic.json); vendor URL: https://stspg.io/6xr5mmpjs1k3
- Start time basis: `vendor_reported`.

<a id="ci-738325c596fe150d64edc826"></a>

- Canonical ID: `ci-738325c596fe150d64edc826` (1 source records).
  - [cloudflare/s18kw61f2ht5](../../data/2026/W31/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/s18kw61f2ht5
- Start time basis: `vendor_reported`.

<a id="ci-87e5abffcd1918fb8c03a73d"></a>

- Canonical ID: `ci-87e5abffcd1918fb8c03a73d` (1 source records).
  - [anthropic/kqjy03gs895j](../../data/2026/W31/anthropic.json); vendor URL: https://stspg.io/5d61g084jy31
- Start time basis: `vendor_reported`.

<a id="ci-69620b632b63a889000f628f"></a>

- Canonical ID: `ci-69620b632b63a889000f628f` (1 source records).
  - [grafana/mbqs14vz2l29](../../data/2026/W31/grafana.json); vendor URL: https://stspg.io/35m2lfvv6h87
- Start time basis: `vendor_reported`.

<a id="ci-9fd7be4c0c91e6351239b446"></a>

- Canonical ID: `ci-9fd7be4c0c91e6351239b446` (1 source records).
  - [github/dsrfymph7my0](../../data/2026/W31/github.json); vendor URL: https://stspg.io/tdw1hjz3lj91
- Start time basis: `vendor_reported`.

<a id="ci-47b5f1657da36f33ba829b13"></a>

- Canonical ID: `ci-47b5f1657da36f33ba829b13` (1 source records).
  - [cloudflare/shq1h521dfl3](../../data/2026/W31/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/shq1h521dfl3
- Start time basis: `vendor_reported`.

<a id="ci-6c9b925d6c0d2d6bc9c0e7b4"></a>

- Canonical ID: `ci-6c9b925d6c0d2d6bc9c0e7b4` (1 source records).
  - [openai/01KYSQ7PC81PV3PZE3ZEVFK6K5](../../data/2026/W31/openai.json); vendor URL: https://status.openai.com/incidents/01KYSQ7PC81PV3PZE3ZEVFK6K5
- Start time basis: `vendor_reported`.

<a id="ci-ac01c3149ddf20139b887d53"></a>

- Canonical ID: `ci-ac01c3149ddf20139b887d53` (1 source records).
  - [openai/01KYNNX2TQ3PV0KPJ9PTRRPX6C](../../data/2026/W31/openai.json); vendor URL: https://status.openai.com/incidents/01KYNNX2TQ3PV0KPJ9PTRRPX6C
- Start time basis: `vendor_reported`.

<a id="ci-c144764bc18d498f98d98d7c"></a>

- Canonical ID: `ci-c144764bc18d498f98d98d7c` (1 source records).
  - [openai/01KYJFG4S86K8FMTB4N6BDZS7F](../../data/2026/W31/openai.json); vendor URL: https://status.openai.com/incidents/01KYJFG4S86K8FMTB4N6BDZS7F
- Start time basis: `vendor_reported`.

<a id="ci-b9b38b4963d5b0cdf5011f5b"></a>

- Canonical ID: `ci-b9b38b4963d5b0cdf5011f5b` (1 source records).
  - [github/kk183dslzdzd](../../data/2026/W31/github.json); vendor URL: https://stspg.io/pj5lgggbc2rb
- Start time basis: `vendor_reported`.

<a id="ci-277385acc2888eeafd280fdb"></a>

- Canonical ID: `ci-277385acc2888eeafd280fdb` (1 source records).
  - [cloudflare/wcln78w0lx1h](../../data/2026/W31/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/wcln78w0lx1h
- Start time basis: `vendor_reported`.

<a id="ci-a7c91a50ca60449a313c8edc"></a>

- Canonical ID: `ci-a7c91a50ca60449a313c8edc` (1 source records).
  - [github/9tpgqq1h4bs7](../../data/2026/W31/github.json); vendor URL: https://stspg.io/45g8r57kv3n9
- Start time basis: `vendor_reported`.

<a id="ci-f21c95f3f50ca1b5b43bec14"></a>

- Canonical ID: `ci-f21c95f3f50ca1b5b43bec14` (1 source records).
  - [stripe/hcmr28nnd39l](../../data/2026/W31/stripe.json); vendor URL: https://stspg.io/l0fsvp40d3hy
- Start time basis: `vendor_reported`.

<a id="ci-3abdd955fb1ca1b1bfcd0ba3"></a>

- Canonical ID: `ci-3abdd955fb1ca1b1bfcd0ba3` (1 source records).
  - [stripe/7pw1p0pw3b17](../../data/2026/W31/stripe.json); vendor URL: https://stspg.io/ykd350r61yh5
- Start time basis: `vendor_reported`.

<a id="ci-6264d53cb454512bd2ee123d"></a>

- Canonical ID: `ci-6264d53cb454512bd2ee123d` (1 source records).
  - [grafana/3rqk10z0jlr1](../../data/2026/W31/grafana.json); vendor URL: https://stspg.io/x4fmhx2pym34
- Start time basis: `vendor_reported`.

<a id="ci-021d15e2f8dcdf7e33f389b2"></a>

- Canonical ID: `ci-021d15e2f8dcdf7e33f389b2` (1 source records).
  - [github/sj1tzyrx599x](../../data/2026/W31/github.json); vendor URL: https://stspg.io/z58g1g529hzz
- Start time basis: `vendor_reported`.

<a id="ci-db8c7df70b6185d29184348c"></a>

- Canonical ID: `ci-db8c7df70b6185d29184348c` (1 source records).
  - [openai/01KYMJND2XYYBDVCF8DKA5DY3N](../../data/2026/W31/openai.json); vendor URL: https://status.openai.com/incidents/01KYMJND2XYYBDVCF8DKA5DY3N
- Start time basis: `vendor_reported`.

<a id="ci-789023240211381935416cb7"></a>

- Canonical ID: `ci-789023240211381935416cb7` (1 source records).
  - [github/75g5xmzptjqb](../../data/2026/W31/github.json); vendor URL: https://stspg.io/8nsh6820sff9
- Start time basis: `vendor_reported`.

<a id="ci-60fe110deb5da1ad3190390c"></a>

- Canonical ID: `ci-60fe110deb5da1ad3190390c` (1 source records).
  - [cloudflare/z325n2dc14tk](../../data/2026/W31/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/z325n2dc14tk
- Start time basis: `vendor_reported`.

<a id="ci-c9eac398ed12abcf135819be"></a>

- Canonical ID: `ci-c9eac398ed12abcf135819be` (1 source records).
  - [datadog/9kn7bmb25jh1](../../data/2026/W31/datadog.json); vendor URL: https://stspg.io/9trx2m1xsd4h
- Start time basis: `vendor_reported`.

<a id="ci-4a85ed2296b44db2d49b58a0"></a>

- Canonical ID: `ci-4a85ed2296b44db2d49b58a0` (1 source records).
  - [datadog/w3w4b0cchj1j](../../data/2026/W31/datadog.json); vendor URL: https://stspg.io/803q340l4b6k
- Start time basis: `vendor_reported`.

<a id="ci-4a1f257197e2da2dec105991"></a>

- Canonical ID: `ci-4a1f257197e2da2dec105991` (1 source records).
  - [openai/01KYVPNEY9736Z04D3X1NSFFF9](../../data/2026/W31/openai.json); vendor URL: https://status.openai.com/incidents/01KYVPNEY9736Z04D3X1NSFFF9
- Start time basis: `vendor_reported`.

<a id="ci-4136cd3bd2f1119de871ad0a"></a>

- Canonical ID: `ci-4136cd3bd2f1119de871ad0a` (1 source records).
  - [github/fmmsrcg5x638](../../data/2026/W31/github.json); vendor URL: https://stspg.io/vr201n49yl53
- Start time basis: `vendor_reported`.

<a id="ci-c542a4c1a90798d9b296f0f9"></a>

- Canonical ID: `ci-c542a4c1a90798d9b296f0f9` (1 source records).
  - [grafana/vp96psrp0tvn](../../data/2026/W31/grafana.json); vendor URL: https://stspg.io/05bt4v67qbyk
- Start time basis: `vendor_reported`.

<a id="ci-8937f20dab5f938aea9bf875"></a>

- Canonical ID: `ci-8937f20dab5f938aea9bf875` (1 source records).
  - [openai/01KYJC3TDHZVSQFFPTC8X9RAHC](../../data/2026/W31/openai.json); vendor URL: https://status.openai.com/incidents/01KYJC3TDHZVSQFFPTC8X9RAHC
- Start time basis: `vendor_reported`.
- Warning: Vendor resolved_at precedes started_at; duration is unknown, score uses zero duration; decision:legacy-time-openai-01KYJC3TDHZVSQFFPTC8X9RAHC

<a id="ci-af3564c09efdae23af8de429"></a>

- Canonical ID: `ci-af3564c09efdae23af8de429` (1 source records).
  - [grafana/6mdfbztbzctw](../../data/2026/W31/grafana.json); vendor URL: https://stspg.io/5db188nt4pmn
- Start time basis: `vendor_reported`.

<a id="ci-480b8fcfd4cf242d82fc6565"></a>

- Canonical ID: `ci-480b8fcfd4cf242d82fc6565` (1 source records).
  - [grafana/bq8dhxj55g7g](../../data/2026/W31/grafana.json); vendor URL: https://stspg.io/mm2jxybgmg52
- Start time basis: `vendor_reported`.
- Warning: Vendor resolved_at precedes started_at; duration is unknown, score uses zero duration; decision:legacy-time-grafana-bq8dhxj55g7g

<a id="ci-4e8bea227d9af55e1d268028"></a>

- Canonical ID: `ci-4e8bea227d9af55e1d268028` (1 source records).
  - [cloudflare/pgbznnvc2vtw](../../data/2026/W31/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/pgbznnvc2vtw
- Start time basis: `vendor_reported`.

<a id="ci-6cc05a1c0f5f08dbadca0af6"></a>

- Canonical ID: `ci-6cc05a1c0f5f08dbadca0af6` (1 source records).
  - [stripe/y8898n2yk19y](../../data/2026/W31/stripe.json); vendor URL: https://stspg.io/sbldmtv3y0jq
- Start time basis: `vendor_reported`.

<a id="ci-8c6a8fd652d92e15fed7c581"></a>

- Canonical ID: `ci-8c6a8fd652d92e15fed7c581` (1 source records).
  - [cloudflare/7czwy7vt3phg](../../data/2026/W31/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/7czwy7vt3phg
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
