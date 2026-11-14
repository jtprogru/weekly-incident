# Week 2026-W27

2026-06-29 to 2026-07-05 (UTC). Generated 2026-11-14 00:00.

28 canonical incidents, 22 above threshold. All 14 sources responded.

28 source records (28 saved revisions) → 28 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | grafana | [Some Queries Failing](https://stspg.io/hmkbllwz66w2) | 2026-07-02 18:10 | 3h 25m | minor | AWS Australia - prod-ap-southeast-2, AWS Brazil - prod-sa-east-1, AWS Canada - prod-ca-east-0, AWS Germany - prod-eu-west-2, AWS India - prod-ap-south-1, AWS Ireland - prod-eu-west-6, AWS Japan - prod-ap-northeast-0, AWS Singapore - prod-ap-southeast-1, AWS Sweden - prod-eu-north-0, AWS Switzerland - prod-eu-central-0, AWS UAE - prod-me-central-1, AWS US East - prod-us-east-0, AWS US West - prod-us-west-0, Azure Netherlands - prod-eu-west-3, Azure US Central - us-central2, Azure US Central - us-central7, Federal Cloud - AWS US Gov West, GCP Australia - prod-au-southeast-0, GCP Belgium - prod-eu-west-0, GCP Brazil - prod-sa-east-0, GCP India - prod-ap-south-0, GCP Singapore - prod-ap-southeast-0, GCP UK - prod-gb-south-0, GCP US Central - prod-us-central-0, GCP US Central - prod-us-central-5 | 5125 = 205 × 25 × 1.0 |
| 2 | grafana | [Elevated Loki Query Bytes Reporting](https://stspg.io/wqbptt5wfqtg) | 2026-07-01 19:43 | 2h 17m | minor | AWS Australia - prod-ap-southeast-2, AWS Brazil - prod-sa-east-1, AWS Canada - prod-ca-east-0, AWS Germany - prod-eu-west-2, AWS India - prod-ap-south-1, AWS Ireland - prod-eu-west-6, AWS Japan - prod-ap-northeast-0, AWS Singapore - prod-ap-southeast-1, AWS Sweden - prod-eu-north-0, AWS Switzerland - prod-eu-central-0, AWS UAE - prod-me-central-1, AWS US East - prod-us-east-0, AWS US West - prod-us-west-0, Azure Netherlands - prod-eu-west-3, Azure US Central - us-central2, Azure US Central - us-central7, Federal Cloud - AWS US Gov West, GCP Australia - prod-au-southeast-0, GCP Belgium - prod-eu-west-0, GCP Brazil - prod-sa-east-0, GCP India - prod-ap-south-0, GCP Singapore - prod-ap-southeast-0, GCP UK - prod-gb-south-0, GCP US Central - prod-us-central-0, GCP US Central - prod-us-central-5 | 3425 = 137 × 25 × 1.0 |
| 3 | grafana | [Mimir read errors and high latency in prod-eu-west-0](https://stspg.io/h6ydsghh5vvs) | 2026-06-29 11:37 | 20h 43m | minor | GCP Belgium - prod-eu-west-0: Ingestion, GCP Belgium - prod-eu-west-0: Querying | 2486 = 1243 × 2 × 1.0 |
| 4 | digitalocean | [Agent Timeouts While Retrieving Data from Knowledge Bases](https://stspg.io/sp8b09xxqvjt) | 2026-07-01 06:25 | 21h 13m | none | Agent Runtime, Knowledge Bases | 2037 = 1273 × 2 × 0.8 |
| 5 | datadog | [Delayed data processing and errors in multiple products](https://stspg.io/5snql687z9c4) | 2026-06-30 17:51 | 2h 40m | major | APM, CI Visibility, Code Coverage, Database Monitoring, Error Tracking, Observability Pipelines, RUM, Synthetics, Test Optimization | 1440 = 160 × 9 × 1.0 |
| 6 | grafana | [Grafana Cloud IRM alert groups failing to produce alerts in us-east-3 region.](https://stspg.io/34p7wm05ydt0) | 2026-07-05 11:37 | 19h 59m | major | Incident Management and Response (IRM) | 1199 = 1199 × 1 × 1.0 |
| 7 | yandexcloud | [The problem is with SpeechKit speech recognition services.](https://status.yandex.cloud/en/incidents/1869) | 2026-06-30 21:00 | 12h 0m | minor | SpeechKit | 1080 = 720 × 1 × 1.5 |
| 8 | digitalocean | [Managed Database Clusters](https://stspg.io/rpkck6zr7ht5) | 2026-07-02 22:04 | 10h 15m | minor | Global, Global | 984 = 615 × 2 × 0.8 |
| 9 | stripe | [Elevated WeChat Pay errors](https://stspg.io/17c6mmm4lg2f) | 2026-07-02 15:07 | 12h 48m | minor | Acquirers and payment methods | 922 = 768 × 1 × 1.2 |
| 10 | digitalocean | [DNS Lookup Failures for Managed Database](https://stspg.io/r6yw4xc720r8) | 2026-07-01 11:35 | 4h 52m | minor | DNS, FRA1, FRA1 | 701 = 292 × 3 × 0.8 |
| 11 | circleci | [Some customers are unable to access their organizations](https://stspg.io/ty6ll6xj9c18) | 2026-06-29 10:23 | 6h 34m | minor | CircleCI UI | 394 = 394 × 1 × 1.0 |
| 12 | digitalocean | [Limited access to Deepseek V4 Pro model](https://stspg.io/wwsc9zft1lxd) | 2026-07-04 22:13 | 2h 39m | major | Inference, Model Services | 254 = 159 × 2 × 0.8 |
| 13 | github | [Delays in copilot budget limits resets for some users](https://stspg.io/jy3yhfvwxpjh) | 2026-07-01 10:51 | 2h 34m | minor | Copilot | 231 = 154 × 1 × 1.5 |
| 14 | grafana | [Loki and Frontend Observability - Major Outage in prod-us-central-0 region](https://stspg.io/3v5hcx1v7p32) | 2026-07-02 04:55 | 1h 53m | critical | GCP US Central - prod-us-central-0, GCP US Central - prod-us-central-0 | 226 = 113 × 2 × 1.0 |
| 15 | grafana | [Rule evaluation error on cluster prod-gb-south-0](https://stspg.io/qldkxzbdb7qp) | 2026-06-29 10:34 | 2h 38m | minor | GCP UK - prod-gb-south-0: Rule Evaluation | 158 = 158 × 1 × 1.0 |
| 16 | circleci | [Outage impacting web experience and pipelines starting](https://stspg.io/n86hwbb1whpc) | 2026-07-02 15:25 | 1h 18m | critical | CircleCI UI, Pipelines & Workflows | 156 = 78 × 2 × 1.0 |
| 17 | datadog | [Elevated Error Rates](https://stspg.io/h1vdjddg90hd) | 2026-07-01 18:31 | 2h 5m | none | — | 125 = 125 × 1 × 1.0 |
| 18 | grafana | [Confluent API Outage](https://stspg.io/b3bbl34zgmbw) | 2026-06-29 12:37 | 1h 49m | major | Integrations | 109 = 109 × 1 × 1.0 |
| 19 | digitalocean | [Droplet Resize](https://stspg.io/m6y2gmlz3jrv) | 2026-07-01 11:13 | 2h 14m | none | Cloud Control Panel | 107 = 134 × 1 × 0.8 |
| 20 | yandexcloud | [Partial unavailability of Virtual Private Cloud services](https://status.yandex.cloud/en/incidents/1865) | 2026-06-30 10:20 | 1h 6m | major | Virtual Private Cloud | 99 = 66 × 1 × 1.5 |
| 21 | datadog | [Errors Pulling Container Images](https://stspg.io/gs85vq0xmxcx) | 2026-07-03 00:25 | 1h 3m | critical | Container Registry | 63 = 63 × 1 × 1.0 |
| 22 | datadog | [Delayed Monitors Notifications](https://stspg.io/hsz1r8ww26gn) | 2026-06-30 10:35 | 43m | major | Monitors | 43 = 43 × 1 × 1.0 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. grafana — Some Queries Failing

- URL: https://stspg.io/hmkbllwz66w2
- Started: 2026-07-02 18:10 UTC
- Resolved: 2026-07-02 21:35 UTC
- Duration: 3h 25m
- Impact: minor, status: resolved
- Components: AWS Australia - prod-ap-southeast-2, AWS Brazil - prod-sa-east-1, AWS Canada - prod-ca-east-0, AWS Germany - prod-eu-west-2, AWS India - prod-ap-south-1, AWS Ireland - prod-eu-west-6, AWS Japan - prod-ap-northeast-0, AWS Singapore - prod-ap-southeast-1, AWS Sweden - prod-eu-north-0, AWS Switzerland - prod-eu-central-0, AWS UAE - prod-me-central-1, AWS US East - prod-us-east-0, AWS US West - prod-us-west-0, Azure Netherlands - prod-eu-west-3, Azure US Central - us-central2, Azure US Central - us-central7, Federal Cloud - AWS US Gov West, GCP Australia - prod-au-southeast-0, GCP Belgium - prod-eu-west-0, GCP Brazil - prod-sa-east-0, GCP India - prod-ap-south-0, GCP Singapore - prod-ap-southeast-0, GCP UK - prod-gb-south-0, GCP US Central - prod-us-central-0, GCP US Central - prod-us-central-5
- Canonical incident: [ci-b2091b3ec81242da991f3243](#ci-b2091b3ec81242da991f3243)

- `2026-07-02 18:10` **identified** — Queries with drop __error__, including log volume histogram queries (the queries that generate the histogram visualization in Grafana), are failing due to a bug with series limit checks. The Root cause has been identified, and we are working on a fix.
- `2026-07-02 20:16` **identified** — We are continuing to work on rolling back a PR responsible for this behavior. Once we have more information, we will share it here. Thank you for your patience.
- `2026-07-02 21:35` **resolved** — This incident has been resolved. Thank you for your patience.

### 2. grafana — Elevated Loki Query Bytes Reporting

- URL: https://stspg.io/wqbptt5wfqtg
- Started: 2026-07-01 19:43 UTC
- Resolved: 2026-07-01 22:00 UTC
- Duration: 2h 17m
- Impact: minor, status: resolved
- Components: AWS Australia - prod-ap-southeast-2, AWS Brazil - prod-sa-east-1, AWS Canada - prod-ca-east-0, AWS Germany - prod-eu-west-2, AWS India - prod-ap-south-1, AWS Ireland - prod-eu-west-6, AWS Japan - prod-ap-northeast-0, AWS Singapore - prod-ap-southeast-1, AWS Sweden - prod-eu-north-0, AWS Switzerland - prod-eu-central-0, AWS UAE - prod-me-central-1, AWS US East - prod-us-east-0, AWS US West - prod-us-west-0, Azure Netherlands - prod-eu-west-3, Azure US Central - us-central2, Azure US Central - us-central7, Federal Cloud - AWS US Gov West, GCP Australia - prod-au-southeast-0, GCP Belgium - prod-eu-west-0, GCP Brazil - prod-sa-east-0, GCP India - prod-ap-south-0, GCP Singapore - prod-ap-southeast-0, GCP UK - prod-gb-south-0, GCP US Central - prod-us-central-0, GCP US Central - prod-us-central-5
- Canonical incident: [ci-403c0d3369f9b580168e1f9e](#ci-403c0d3369f9b580168e1f9e)

- `2026-07-01 19:43` **investigating** — We are investigating an issue where some customers may see higher Loki query byte usage reported than was actually consumed. This affects usage reporting only; there is no impact to query execution or service availability. The issue began at approximately 13:20 UTC and is ongoing. We expect the issue to be resolved soon and will provide another update as more information becomes available.
- `2026-07-01 22:00` **resolved** — We’ve implemented a fix and can confirm the issue is fully resolved as of 20:25 UTC. Thank you for your patience.

### 3. grafana — Mimir read errors and high latency in prod-eu-west-0

- URL: https://stspg.io/h6ydsghh5vvs
- Started: 2026-06-29 11:37 UTC
- Resolved: 2026-06-30 08:20 UTC
- Duration: 20h 43m
- Impact: minor, status: resolved
- Components: GCP Belgium - prod-eu-west-0: Ingestion, GCP Belgium - prod-eu-west-0: Querying
- Canonical incident: [ci-2643670698537376487b6cc4](#ci-2643670698537376487b6cc4)

- `2026-06-29 11:37` **investigating** — We are currently investigating an issue with Mimir in prod-eu-west-0 we are seeing read errors and high latency. This incident is currently ongoing. The errors are recovering but we are currently looking into the route cause of this.
- `2026-06-29 13:52` **investigating** — The errors are recovering, and we are still looking into the root cause of this.
- `2026-06-29 15:37` **investigating** — The errors and latency have now recovered, we continue investigating the root cause.
- `2026-06-29 17:47` **identified** — We've identified a possible cause, and a mitigation is in place to prevent further occurrences.
- `2026-06-30 08:20` **resolved** — Since the mitigation has been applied, we have not seen the errors return. At this point, we are considering the incident resolved.

### 4. digitalocean — Agent Timeouts While Retrieving Data from Knowledge Bases

- URL: https://stspg.io/sp8b09xxqvjt
- Started: 2026-07-01 06:25 UTC
- Resolved: 2026-07-02 03:38 UTC
- Duration: 21h 13m
- Impact: none, status: resolved
- Components: Agent Runtime, Knowledge Bases
- Canonical incident: [ci-c016a22dbd3c939ab94c7e55](#ci-c016a22dbd3c939ab94c7e55)

- `2026-07-01 06:25` **investigating** — Our Engineering team is currently investigating an issue where Agents are experiencing timeouts while retrieving data from Knowledge Bases. As a result, affected Agents may fail to retrieve data and return the following error: "Failed to retrieve data from Knowledge base(s) - timeout" Please be assured that we are treating this as a high-priority issue and are actively working to mitigate it. We sincerely apologise for any inconvenience this may have caused and appreciate your patience and understanding. If you have any further questions, please create a support ticket so that we can investigate your specific case further.
- `2026-07-01 08:39` **identified** — Our Engineering team has identified the issue causing Agents to experience timeouts while retrieving data from Knowledge Bases. Please be assured that our Engineering team is actively working on a fix and is treating this issue with high priority. We sincerely apologise for any inconvenience this may have caused and appreciate your patience and understanding. If you have any further questions, please create a support ticket so that we can investigate your specific case further.
- `2026-07-02 03:38` **resolved** — This incident has been resolved, and our teams continue to monitor the results. If you experience any further issues, please contact support.

### 5. datadog — Delayed data processing and errors in multiple products

- URL: https://stspg.io/5snql687z9c4
- Started: 2026-06-30 17:51 UTC
- Resolved: 2026-06-30 20:32 UTC
- Duration: 2h 40m
- Impact: major, status: resolved
- Components: APM, CI Visibility, Code Coverage, Database Monitoring, Error Tracking, Observability Pipelines, RUM, Synthetics, Test Optimization
- Canonical incident: [ci-caba25cbe35e7151a8c9f079](#ci-caba25cbe35e7151a8c9f079)

- `2026-06-30 17:51` **investigating** — We are currently investigating an issue affecting CI Visibility, Test Optimization, Code Coverage, Software Delivery, Preemptive Alerting, Session Replay, RUM Explorer, Agent Observability, Database Monitoring, Event Correlation, Error Tracking, and all Security products in our US1 datacenter. As a result, some users may experience delays in data processing and errors when using affected products. Our engineering teams have identified the issue and are actively working on a resolution. Our engineering teams are actively working to resolve the issue, and we'll follow up with more information shortly. Please let us know if you have any questions or need anything else in the meantime.
- `2026-06-30 17:54` **investigating** — We are continuing to investigate this issue.
- `2026-06-30 18:28` **investigating** — We are continuing to investigate this issue.
- `2026-06-30 18:36` **investigating** — We are continuing to investigate this issue.
- `2026-06-30 18:57` **identified** — The issue has been identified and a fix is being implemented.
- `2026-06-30 19:06` **identified** — We are continuing to work on a fix for this issue.
- `2026-06-30 19:36` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-06-30 20:32` **resolved** — This incident has been resolved.

### 6. grafana — Grafana Cloud IRM alert groups failing to produce alerts in us-east-3 region.

- URL: https://stspg.io/34p7wm05ydt0
- Started: 2026-07-05 11:37 UTC
- Resolved: 2026-07-06 07:37 UTC
- Duration: 19h 59m
- Impact: major, status: resolved
- Components: Incident Management and Response (IRM)
- Canonical incident: [ci-ec8c7e9b66186864472ab8f8](#ci-ec8c7e9b66186864472ab8f8)

- `2026-07-05 11:37` **identified** — Some IRM alert groups in us-east-3 region may not be able to produce alerts and could be sluggish/misbehaving in general. The issue started around 01:00 UTC on July 3rd. The issue has been identified and the cause of the issue has been fixed. Our team is actively working on getting the service back to healthy state.
- `2026-07-05 13:38` **monitoring** — The work on bringing services back to healthy state is completed and alerts should be created correctly, without a delay now. We're keeping the incident open and monitoring for any potential hiccups that might occur.
- `2026-07-06 07:37` **resolved** — We haven't noticed any further issues in this region for alert group processing since yesterday. This incident i fully resolved.

### 7. yandexcloud — The problem is with SpeechKit speech recognition services.

- URL: https://status.yandex.cloud/en/incidents/1869
- Started: 2026-06-30 21:00 UTC
- Resolved: 2026-07-01 09:00 UTC
- Duration: 12h 0m
- Impact: minor, status: resolved
- Components: SpeechKit
- Canonical incident: [ci-a3f1398bc03b034c29d89bb7](#ci-a3f1398bc03b034c29d89bb7)

- `2026-07-01 08:44` **investigating** — We are experiencing problems with the availability of the SpeechKit speech recognition service via API v1. Currently, we are not recording any impact on users, we continue to investigate the reasons for unavailability and delays.
- `2026-07-03 10:06` **resolved** — We no longer see any problems with the service.

### 8. digitalocean — Managed Database Clusters

- URL: https://stspg.io/rpkck6zr7ht5
- Started: 2026-07-02 22:04 UTC
- Resolved: 2026-07-03 08:20 UTC
- Duration: 10h 15m
- Impact: minor, status: resolved
- Components: Global, Global
- Canonical incident: [ci-0895e0fe0250a38b16f680b8](#ci-0895e0fe0250a38b16f680b8)

- `2026-07-02 22:04` **investigating** — Our engineering team is investigating an issue affecting Managed Database Clusters. Currently, users may encounter delays when creating/scaling/forking and restoring Standard MySQL, Standard PostgreSQL, OpenSearch, Kafka, and Valkey clusters through the Cloud Control Panel or API. We apologize for the inconvenience and will provide further updates as soon as more information is available.
- `2026-07-03 02:28` **investigating** — Our engineering team is still investigating the issue affecting Managed Database Clusters. Users may still encounter delays creating/scaling/forking/restoring the aforementioned Managed Database Clusters. We are working to resolve this as soon as possible. we apologize for the inconvenience. We will post an update here once we have more information.
- `2026-07-03 06:08` **monitoring** — Our engineering team has implemented a fix to resolve the issue with Managed Database Clusters and is monitoring the situation. We will post an update as soon as the issue is fully resolved.
- `2026-07-03 08:20` **resolved** — Our engineering team has resolved the issue with Managed Database Clusters. All database systems should now be operating normally. If you continue to experience problems, please open a ticket with our support team. We apologize for any inconvenience.

### 9. stripe — Elevated WeChat Pay errors

- URL: https://stspg.io/17c6mmm4lg2f
- Started: 2026-07-02 15:07 UTC
- Resolved: 2026-07-03 03:56 UTC
- Duration: 12h 48m
- Impact: minor, status: resolved
- Components: Acquirers and payment methods
- Canonical incident: [ci-47c5671c7efbe982c022e118](#ci-47c5671c7efbe982c022e118)

- `2026-07-02 13:37` **investigating** — We're currently investigating a degradation of the WeChat Pay payment method that began at 7:10 UTC. This issue is external to Stripe and is causing increased error rates for transactions processed through the WeChat Pay payment method. We are actively monitoring the situation and will provide updates as more information becomes available.
- `2026-07-02 15:07` **identified** — We’re currently seeing a degradation of the WeChat Pay payment method. Our partner has identified the cause and is working on a fix.
- `2026-07-02 19:23` **identified** — WeChat Pay has confirmed they are addressing the issue and it will take some time to implement the fix. We are actively monitoring and will provide a further update as more information becomes available.
- `2026-07-03 03:56` **resolved** — From 07:10 July 2 to 01:20 July 3 UTC, we saw elevated errors which affected WeChat Pay payment method transactions. This issue has now been resolved by our partner.

### 10. digitalocean — DNS Lookup Failures for Managed Database

- URL: https://stspg.io/r6yw4xc720r8
- Started: 2026-07-01 11:35 UTC
- Resolved: 2026-07-01 16:28 UTC
- Duration: 4h 52m
- Impact: minor, status: resolved
- Components: DNS, FRA1, FRA1
- Canonical incident: [ci-ae8f5db0d97b26dd5773bf9f](#ci-ae8f5db0d97b26dd5773bf9f)

- `2026-07-01 11:35` **investigating** — As of 06:39 UTC, our Engineering team is investigating reports of intermittent DNS lookup failures for Managed Database hostnames from Managed Kubernetes, primarily affecting connections from Managed Kubernetes to Managed Database hostnames. At this point, customers in the FRA1 region may experience intermittent connectivity issues. We apologize for the inconvenience and will share an update once we have more information.
- `2026-07-01 12:15` **investigating** — We are continuing to investigate this issue.
- `2026-07-01 16:28` **resolved** — The intermittent DNS lookup failures between Managed Kubernetes and Managed Database hostnames in the FRA1 region have been resolved for now but please let us know if you see the issue again. Connectivity has remained stable, and all systems are operating normally. If you continue to experience any problems, please open a ticket with our support team. Thank you for your patience, and we apologize for any inconvenience.

### 11. circleci — Some customers are unable to access their organizations

- URL: https://stspg.io/ty6ll6xj9c18
- Started: 2026-06-29 10:23 UTC
- Resolved: 2026-06-29 16:58 UTC
- Duration: 6h 34m
- Impact: minor, status: resolved
- Components: CircleCI UI
- Canonical incident: [ci-c4349ca381003b5329f4cb25](#ci-c4349ca381003b5329f4cb25)

- `2026-06-29 10:23` **investigating** — We are investigating reports that some customers are unable to access their organizations in CircleCI. Those affected may see errors when loading their organizations. Our engineering team is actively investigating.
- `2026-06-29 11:33` **investigating** — We are continuing to investigate reports that some customers are unable to access their organizations in CircleCI. Those affected may see errors when loading their organizations. Our engineering team are still working to identify the cause. We will share another update as soon as we have more information.
- `2026-06-29 11:44` **investigating** — We are still investigating reports that some customers cannot access their organizations in CircleCI. Early findings suggest the problem may be related to GitHub authorization, where affected users receive a 403 from GitHub indicating the resource is SAML protected and the OAuth token may not be authorized. We have not yet confirmed this as the root cause. As a potential workaround, some users may be able to restore access by performing a re-authentication: https://support.circleci.com/hc/en-us/articles/360051228052-How-to-Perform-a-Re-Authentication We will post another update once we have more certainty.
- `2026-06-29 12:16` **investigating** — We are continuing to investigate reports that some customers cannot access their organizations in CircleCI. Our investigation continues to point to an authorization issue between CircleCI and GitHub. Some affected users have received a 403 from GitHub indicating the resource is SAML protected and the OAuth token may not be authorized. We previously shared re-authenticating as a potential workaround. We have since received reports that this didn't restore access for all affected customers, so we don't consider it a reliable resolution at this time. Our engineering team is continuing to investigate this to determine the cause and next steps. We will post another update as soon as we have more information. We apologize for the continued disruption and appreciate your patience.
- `2026-06-29 15:54` **identified** — We have identified the cause of errors accessing orgs. We are implementing a fix, and will monitor for recovery as the change rolls out.
- `2026-06-29 16:44` **monitoring** — Error rates viewing orgs are recovering. We will continue to monitor to verify the fix resolves the issues.
- `2026-06-29 16:58` **resolved** — This incident has been resolved.

### 12. digitalocean — Limited access to Deepseek V4 Pro model

- URL: https://stspg.io/wwsc9zft1lxd
- Started: 2026-07-04 22:13 UTC
- Resolved: 2026-07-05 00:53 UTC
- Duration: 2h 39m
- Impact: major, status: resolved
- Components: Inference, Model Services
- Canonical incident: [ci-1cbdc127ad43bcc8a001d1f3](#ci-1cbdc127ad43bcc8a001d1f3)

- `2026-07-04 22:13` **investigating** — Our Engineering team is investigating reports of an incident affecting the Deepseek V4 Pro model in Serverless Inference and Agent Platform. Users may experience errors when attempting to use this model, specifically receiving error 429 messages. We apologize for the inconvenience and are working to resolve the issue as soon as possible. We will provide an update once we have more information.
- `2026-07-04 23:18` **monitoring** — Our Engineering team has implemented a fix for the issue with the Deepseek V4 Pro model in Serverless Inference and Agent Platform. The model should now be operational, and users should no longer receive error 429 messages when attempting to use it. We are currently monitoring the situation to ensure the fix is successful and the model is functioning as expected. We will post an update if any further issues arise. If you continue to experience problems, please open a ticket with our Support team. We apologize for any inconvenience this may have caused.
- `2026-07-05 00:53` **resolved** — Our Engineering team has confirmed that the issue with the Deepseek V4 Pro model in Serverless Inference and Agent Platform has been fully resolved. The model is now operational, and users should be able to use it without experiencing any errors. If you continue to experience any problems, please open a ticket with our Support team. We apologize for any inconvenience this may have caused and appreciate your patience.

### 13. github — Delays in copilot budget limits resets for some users

- URL: https://stspg.io/jy3yhfvwxpjh
- Started: 2026-07-01 10:51 UTC
- Resolved: 2026-07-01 13:26 UTC
- Duration: 2h 34m
- Impact: minor, status: resolved
- Components: Copilot
- Canonical incident: [ci-7578b4f517535f283ae91dc6](#ci-7578b4f517535f283ae91dc6)

- `2026-07-01 10:51` **investigating** — We are investigating reports of degraded performance for Copilot
- `2026-07-01 11:44` **investigating** — We have identified the likely reason for the delays and are working on a solution.
- `2026-07-01 12:42` **investigating** — A fix for the delayed resets is currently being deployed.
- `2026-07-01 13:04` **monitoring** — The degradation affecting Copilot has been mitigated. We are monitoring to ensure stability.
- `2026-07-01 13:05` **monitoring** — The fix has been deployed globally and we are monitoring the results
- `2026-07-01 13:26` **resolved** — On July 1, 2026, between approximately 00:00 UTC and 13:04 UTC, some GitHub Copilot customers whose budget was exhausted before the monthly reset remained incorrectly blocked from paid Copilot usage after the new billing month began, even though their budgets had reset. Some budget changes also took longer than usual to apply. Only customers with an exhausted budget were affected, which limited the impact.<br /><br />This was caused by a caching issue at the monthly reset: for some users, a pre-reset "budget exhausted" status was re-saved and served even though their budget had reset, so they stayed blocked. We had built a safeguard ahead of the reset to prevent this, but it did not take effect because an internal configuration service did not load its settings correctly. We resolved the incident by deploying a change that discards the outdated status and recomputes access from current budget data independently of that configuration, and by working through the backlog of budget updates.<br /><br />To prevent recurrence, we are ensuring pre-reset status cannot survive the monthly budget reset, adding alerting for this failure mode, and increasing capacity to absorb the monthly surge of budget updates.

### 14. grafana — Loki and Frontend Observability - Major Outage in prod-us-central-0 region

- URL: https://stspg.io/3v5hcx1v7p32
- Started: 2026-07-02 04:55 UTC
- Resolved: 2026-07-02 06:49 UTC
- Duration: 1h 53m
- Impact: critical, status: resolved
- Components: GCP US Central - prod-us-central-0, GCP US Central - prod-us-central-0
- Canonical incident: [ci-bdd8200e6c7a6951257ff74b](#ci-bdd8200e6c7a6951257ff74b)

- `2026-07-02 04:55` **investigating** — We are currently investigating a major outage in Loki writes and Frontend Observability in the prod-us-central-0 region. Our Engineering team is investigating this and we will provide further updates as our investigation progresses.
- `2026-07-02 04:57` **investigating** — The AWS Logs integration in the same region is affected as well. We will provide further updates as our investigation progresses.
- `2026-07-02 06:49` **resolved** — This incident has been resolved by restarting the affected services.

### 15. grafana — Rule evaluation error on cluster prod-gb-south-0

- URL: https://stspg.io/qldkxzbdb7qp
- Started: 2026-06-29 10:34 UTC
- Resolved: 2026-06-29 13:12 UTC
- Duration: 2h 38m
- Impact: minor, status: resolved
- Components: GCP UK - prod-gb-south-0: Rule Evaluation
- Canonical incident: [ci-74caae68f2302c19902f85fc](#ci-74caae68f2302c19902f85fc)

- `2026-06-29 10:34` **investigating** — We are currently investigating Rule Evaluation errors on the cluster prod-gb-south-0 which is leading to error codes showing within the stacks. We are looking into the issue and will update accordingly.
- `2026-06-29 11:14` **monitoring** — A fix has been applied and we are currently monitoring results.
- `2026-06-29 13:12` **resolved** — This incident has been resolved.

### 16. circleci — Outage impacting web experience and pipelines starting

- URL: https://stspg.io/n86hwbb1whpc
- Started: 2026-07-02 15:25 UTC
- Resolved: 2026-07-02 16:44 UTC
- Duration: 1h 18m
- Impact: critical, status: postmortem
- Components: CircleCI UI, Pipelines & Workflows
- Canonical incident: [ci-6b67d170d143c0f52790ec63](#ci-6b67d170d143c0f52790ec63)

- `2026-07-02 15:25` **investigating** — We are currently investigating this issue.
- `2026-07-02 16:07` **investigating** — We are investigating an issue that may cause some customers to be redirected to a login page and prevents some pipelines from starting. Thank you for your patience as we continue our mitigation of this issue.
- `2026-07-02 16:24` **monitoring** — We have identified a source of increased system load which caused failures processing API requests. We have mitigated the issue and jobs are now running as expected and UI access has recovered. Some pipelines may have failed or failed to start. Resubmitting should correct these issues. We are continuing to monitor to ensure recovery.
- `2026-07-02 16:29` **monitoring** — Pipelines are continuing to run as normal. We will continue monitoring for recovery.
- `2026-07-02 16:44` **resolved** — Pipelines and UI access have fully recovered. Please contact customer support if you have issues running builds. This incident has been resolved.
- `2026-07-10 18:57` **postmortem** — ## Summary On July 2, 2026, from 15:12 UTC to 16:44 UTC, some CircleCI customers experienced failures starting pipelines and difficulty accessing the CircleCI UI, including an inability to log in. Customers triggering pipelines via the API may also have had pipelines fail to run. The incident originated from a combination of internal maintenance and customer-initiated project deletions running concurrently, which caused elevated load on a data service. The connection layer between our API and that data service was not configured to time out slow responses under the elevated load conditions seen during this incident. This eventually exhausted available threads and prevented the request routing layer from handling any other requests, resulting in customer-facing errors. The affected services were scaled up and the internal maintenance job was canceled to reduce pressure, allowing all affected functionality to return to normal. Customers whose jobs failed during this window may rerun affected jobs. A small number of customers may still see workflows that appear to be running; this is actively being corrected and does not impact billing. We thank you for your patience while our team worked on implementing a fix. ## What Happened \(all times UTC\) At 13:30 on July 2, 2026, as part of normal platform maintenance, our engineering team began a data deletion job, divided into several hundred individual maintenance tasks. At 14:40, unrelated to our platform maintenance, a significant number of customer project deletion requests were received and began executing. At 15:12, the combined load from the maintenance tasks and the customer-driven projection deletions caused an internal data service to slow down, and this began cascading into our API layer. Because connections between the API layer and the workflows data service were not configured to time out slow responses under elevated load conditions, the API-handling components became overwhelmed, went unresponsive, and began failing their health checks. This caused our request routing layer to lose available upstream requests, resulting in customers seeing errors when attempting to access the UI, to log in via GitHub or Bitbucket OAuth, or to start pipelines. At 15:16, our monitoring systems alerted and engineers were paged. The team investigated and confirmed the source of the elevated load. At 15:30 Engineers began to scale up the affected services and cancelled a long-running maintenance task to drain the load. The team noted signs of that the system was recovering around 16:12 and adjusted the status to monitoring at 16:24. At 16:44, the team confirmed that all customer-facing services had recovered to normal operating levels. Background load from the deletion job drained by approximately 17:00. Following the incident, we discovered a small number of workflows continued to inaccurately report that they were in a running state. This should be corrected within the next few days. In all cases, customers will only be billed for actual compute time used. ## Future Prevention and Process Improvement Our team is implementing the following improvements as a result of this incident: * After the incident was resolved we added timeout protections on connections between our API layer and internal data services, so that slow responses from a downstream service cannot exhaust API worker capacity. * We are actively investigating ways to enforce resource limits on background data tasks to prevent them from impacting customer-facing services. * We are improving observability and alerting on the affected data service so that resource pressure is detected earlier, before it reaches a level that impacts customers.

### 17. datadog — Elevated Error Rates

- URL: https://stspg.io/h1vdjddg90hd
- Started: 2026-07-01 18:31 UTC
- Resolved: 2026-07-01 20:37 UTC
- Duration: 2h 5m
- Impact: none, status: resolved
- Components: —
- Canonical incident: [ci-f8950744b2e269f1bc7bc6e6](#ci-f8950744b2e269f1bc7bc6e6)

- `2026-07-01 18:31` **investigating** — We are actively investigating elevated error rates across multiple Datadog products.
- `2026-07-01 18:57` **investigating** — We are investigating elevated error rates across multiple products including: Fleet Automation, Security Products, Cloudcraft, Cloud Network Monitoring, Serverless, Kubernetes Autoscaling, and Database Monitoring.
- `2026-07-01 19:10` **identified** — We have identified the issue of elevated error rates and are working through mitigation actions
- `2026-07-01 19:19` **monitoring** — We have deployed mitigation actions and are monitoring the recovery.
- `2026-07-01 19:43` **monitoring** — We are continuing to monitor the situation and are taking additional preventive measures.
- `2026-07-01 20:37` **resolved** — This incident has been resolved.

### 18. grafana — Confluent API Outage

- URL: https://stspg.io/b3bbl34zgmbw
- Started: 2026-06-29 12:37 UTC
- Resolved: 2026-06-29 14:27 UTC
- Duration: 1h 49m
- Impact: major, status: resolved
- Components: Integrations
- Canonical incident: [ci-9ab1ff96e9085f39bb114c00](#ci-9ab1ff96e9085f39bb114c00)

- `2026-06-29 12:37` **investigating** — We are investigating an issue affecting Confluent metrics ingestion across all regions. Due to an elevated error rate on the Confluent side, some metrics may not be ingested, resulting in potential data loss. We are actively investigating the issue and will provide updates as more information becomes available.
- `2026-06-29 14:27` **resolved** — This incident has been resolved.

### 19. digitalocean — Droplet Resize

- URL: https://stspg.io/m6y2gmlz3jrv
- Started: 2026-07-01 11:13 UTC
- Resolved: 2026-07-01 13:28 UTC
- Duration: 2h 14m
- Impact: none, status: resolved
- Components: Cloud Control Panel
- Canonical incident: [ci-9a50f2d38a2490f820de3406](#ci-9a50f2d38a2490f820de3406)

- `2026-07-01 11:13` **investigating** — Our Engineering team is investigating an issue affecting Droplet resizes using the "Downscale Anytime" option. At this time users may find this option unavailable or unresponsive in the Cloud Control Panel. As a workaround, resizes can be performed via the API while we work to resolve this issue. We apologize for the inconvenience and will share an update once we have more information.
- `2026-07-01 12:43` **monitoring** — Our Engineering team has identified the root cause of the issue affecting Droplet resizes using the "Downscale Anytime" option and has implemented a fix. Users should now be able to resize their Droplets using the "Downscale Anytime" option without experiencing any issues or errors. We are actively monitoring the situation to ensure the fix remains effective and will provide another update once the issue has been fully resolved.
- `2026-07-01 13:28` **resolved** — Our Engineering team has confirmed that the issue impacting Droplet resizes using the "Downscale Anytime" option has been fully resolved. Users can now resize their Droplets using the "Downscale Anytime" option without any issues. We appreciate your patience while we worked to resolve this issue. If you continue to experience any problems, please open a support ticket from your account so our team can investigate further.

### 20. yandexcloud — Partial unavailability of Virtual Private Cloud services

- URL: https://status.yandex.cloud/en/incidents/1865
- Started: 2026-06-30 10:20 UTC
- Resolved: 2026-06-30 11:26 UTC
- Duration: 1h 6m
- Impact: major, status: resolved
- Components: Virtual Private Cloud
- Canonical incident: [ci-03e6d1698f0d9223b38230ae](#ci-03e6d1698f0d9223b38230ae)

- `2026-06-30 11:05` **investigating** — We noticed a partial unavailability of Virtual Private Cloud services in the Cloud, which may cause new virtual machines to be created with a delay. The developers are already aware of it and are fixing it.
- `2026-06-30 11:41` **resolved** — By this point, the incident has been resolved and the impact has been removed.

### 21. datadog — Errors Pulling Container Images

- URL: https://stspg.io/gs85vq0xmxcx
- Started: 2026-07-03 00:25 UTC
- Resolved: 2026-07-03 01:29 UTC
- Duration: 1h 3m
- Impact: critical, status: resolved
- Components: Container Registry
- Canonical incident: [ci-4392f6b9104ac82ce862fe3d](#ci-4392f6b9104ac82ce862fe3d)

- `2026-07-03 00:25` **investigating** — We are investigating errors pulling Datadog container images from registry.datadoghq.com. As a result of this issue, some users may be unable to download or deploy the Datadog Agent, Cluster Agent, and other Datadog container images. Containers that are already running are not affected.
- `2026-07-03 00:55` **identified** — The issue has been identified and a fix is being implemented.
- `2026-07-03 01:21` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-07-03 01:29` **resolved** — This incident has been resolved.

### 22. datadog — Delayed Monitors Notifications

- URL: https://stspg.io/hsz1r8ww26gn
- Started: 2026-06-30 10:35 UTC
- Resolved: 2026-06-30 11:18 UTC
- Duration: 43m
- Impact: major, status: resolved
- Components: Monitors
- Canonical incident: [ci-e87283eef705d3197ae7093a](#ci-e87283eef705d3197ae7093a)

- `2026-06-30 10:35` **investigating** — We are investigating delays in Monitors Notifications, which began at 09:52 UTC.
- `2026-06-30 10:53` **identified** — We have identified the underlying issue and are working on a fix. It is important to note that no data has been lost, and notifications will be caught up once the service is operational again.
- `2026-06-30 11:11` **monitoring** — We have deployed a fix and we are monitoring the results. We will provide another update once the issue is fully resolved.
- `2026-06-30 11:18` **resolved** — This incident has been resolved.

## Below threshold (6)

- `github` [Incident with Pages](https://stspg.io/wgv67m39tbml) — 1h 31m
- `stripe` [Elevated P24 errors](https://stspg.io/9rxrgvlbpn8g) — 37m
- `digitalocean` [Droplet Actions in NYC3](https://stspg.io/1ggj0xnb3n4v) — 27m
- `github` [Disruption with some GitHub services - Signup Flow](https://stspg.io/vbk1zdbpmv6m) — 10m
- `yandexcloud` [Yandex Cloud Console Availability Issue](https://status.yandex.cloud/en/incidents/1861) — 8m
- `grafana` [Partial Write Outage for Grafana Cloud Logs in prod-eu-north-0.](https://stspg.io/lwn65pz9737z) — 0m

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-b2091b3ec81242da991f3243"></a>

- Canonical ID: `ci-b2091b3ec81242da991f3243` (1 source records).
  - [grafana/x0h2rsqpqk3w](../../data/2026/W27/grafana.json); vendor URL: https://stspg.io/hmkbllwz66w2
- Start time basis: `vendor_reported`.

<a id="ci-403c0d3369f9b580168e1f9e"></a>

- Canonical ID: `ci-403c0d3369f9b580168e1f9e` (1 source records).
  - [grafana/j5vgks57cphf](../../data/2026/W27/grafana.json); vendor URL: https://stspg.io/wqbptt5wfqtg
- Start time basis: `vendor_reported`.

<a id="ci-2643670698537376487b6cc4"></a>

- Canonical ID: `ci-2643670698537376487b6cc4` (1 source records).
  - [grafana/grtt5vk2dd3v](../../data/2026/W27/grafana.json); vendor URL: https://stspg.io/h6ydsghh5vvs
- Start time basis: `vendor_reported`.

<a id="ci-c016a22dbd3c939ab94c7e55"></a>

- Canonical ID: `ci-c016a22dbd3c939ab94c7e55` (1 source records).
  - [digitalocean/76l7xym7jd5m](../../data/2026/W27/digitalocean.json); vendor URL: https://stspg.io/sp8b09xxqvjt
- Start time basis: `vendor_reported`.

<a id="ci-caba25cbe35e7151a8c9f079"></a>

- Canonical ID: `ci-caba25cbe35e7151a8c9f079` (1 source records).
  - [datadog/71vvxf0rxxnz](../../data/2026/W27/datadog.json); vendor URL: https://stspg.io/5snql687z9c4
- Start time basis: `vendor_reported`.

<a id="ci-ec8c7e9b66186864472ab8f8"></a>

- Canonical ID: `ci-ec8c7e9b66186864472ab8f8` (1 source records).
  - [grafana/4m5tfwljjfxk](../../data/2026/W27/grafana.json); vendor URL: https://stspg.io/34p7wm05ydt0
- Start time basis: `vendor_reported`.

<a id="ci-a3f1398bc03b034c29d89bb7"></a>

- Canonical ID: `ci-a3f1398bc03b034c29d89bb7` (1 source records).
  - [yandexcloud/1869](../../data/2026/W27/yandexcloud.json); vendor URL: https://status.yandex.cloud/en/incidents/1869
- Start time basis: `vendor_reported`.

<a id="ci-0895e0fe0250a38b16f680b8"></a>

- Canonical ID: `ci-0895e0fe0250a38b16f680b8` (1 source records).
  - [digitalocean/rjfr3rbzxlx1](../../data/2026/W27/digitalocean.json); vendor URL: https://stspg.io/rpkck6zr7ht5
- Start time basis: `vendor_reported`.

<a id="ci-47c5671c7efbe982c022e118"></a>

- Canonical ID: `ci-47c5671c7efbe982c022e118` (1 source records).
  - [stripe/j5szf7f3wjp5](../../data/2026/W27/stripe.json); vendor URL: https://stspg.io/17c6mmm4lg2f
- Start time basis: `vendor_reported`.

<a id="ci-ae8f5db0d97b26dd5773bf9f"></a>

- Canonical ID: `ci-ae8f5db0d97b26dd5773bf9f` (1 source records).
  - [digitalocean/phwpbgwcmkjn](../../data/2026/W27/digitalocean.json); vendor URL: https://stspg.io/r6yw4xc720r8
- Start time basis: `vendor_reported`.

<a id="ci-c4349ca381003b5329f4cb25"></a>

- Canonical ID: `ci-c4349ca381003b5329f4cb25` (1 source records).
  - [circleci/njbxb52z1g1y](../../data/2026/W27/circleci.json); vendor URL: https://stspg.io/ty6ll6xj9c18
- Start time basis: `vendor_reported`.

<a id="ci-1cbdc127ad43bcc8a001d1f3"></a>

- Canonical ID: `ci-1cbdc127ad43bcc8a001d1f3` (1 source records).
  - [digitalocean/wvj4t3y98gs7](../../data/2026/W27/digitalocean.json); vendor URL: https://stspg.io/wwsc9zft1lxd
- Start time basis: `vendor_reported`.

<a id="ci-7578b4f517535f283ae91dc6"></a>

- Canonical ID: `ci-7578b4f517535f283ae91dc6` (1 source records).
  - [github/rl7f90w0n0gq](../../data/2026/W27/github.json); vendor URL: https://stspg.io/jy3yhfvwxpjh
- Start time basis: `vendor_reported`.

<a id="ci-bdd8200e6c7a6951257ff74b"></a>

- Canonical ID: `ci-bdd8200e6c7a6951257ff74b` (1 source records).
  - [grafana/gsdpkqp4kvlv](../../data/2026/W27/grafana.json); vendor URL: https://stspg.io/3v5hcx1v7p32
- Start time basis: `vendor_reported`.

<a id="ci-74caae68f2302c19902f85fc"></a>

- Canonical ID: `ci-74caae68f2302c19902f85fc` (1 source records).
  - [grafana/6fdsld9f0nsn](../../data/2026/W27/grafana.json); vendor URL: https://stspg.io/qldkxzbdb7qp
- Start time basis: `vendor_reported`.

<a id="ci-6b67d170d143c0f52790ec63"></a>

- Canonical ID: `ci-6b67d170d143c0f52790ec63` (1 source records).
  - [circleci/82j91239k876](../../data/2026/W27/circleci.json); vendor URL: https://stspg.io/n86hwbb1whpc
- Start time basis: `vendor_reported`.

<a id="ci-8e64d12e5dc9d5705e22b763"></a>

- Canonical ID: `ci-8e64d12e5dc9d5705e22b763` (1 source records).
  - [github/5bnwwg9tzd4q](../../data/2026/W27/github.json); vendor URL: https://stspg.io/wgv67m39tbml
- Start time basis: `vendor_reported`.

<a id="ci-f8950744b2e269f1bc7bc6e6"></a>

- Canonical ID: `ci-f8950744b2e269f1bc7bc6e6` (1 source records).
  - [datadog/smc2wgcmsx6r](../../data/2026/W27/datadog.json); vendor URL: https://stspg.io/h1vdjddg90hd
- Start time basis: `vendor_reported`.

<a id="ci-9ab1ff96e9085f39bb114c00"></a>

- Canonical ID: `ci-9ab1ff96e9085f39bb114c00` (1 source records).
  - [grafana/bsfk9nm4l61f](../../data/2026/W27/grafana.json); vendor URL: https://stspg.io/b3bbl34zgmbw
- Start time basis: `vendor_reported`.

<a id="ci-9a50f2d38a2490f820de3406"></a>

- Canonical ID: `ci-9a50f2d38a2490f820de3406` (1 source records).
  - [digitalocean/8mvnl6s4cmdt](../../data/2026/W27/digitalocean.json); vendor URL: https://stspg.io/m6y2gmlz3jrv
- Start time basis: `vendor_reported`.

<a id="ci-03e6d1698f0d9223b38230ae"></a>

- Canonical ID: `ci-03e6d1698f0d9223b38230ae` (1 source records).
  - [yandexcloud/1865](../../data/2026/W27/yandexcloud.json); vendor URL: https://status.yandex.cloud/en/incidents/1865
- Start time basis: `vendor_reported`.

<a id="ci-4392f6b9104ac82ce862fe3d"></a>

- Canonical ID: `ci-4392f6b9104ac82ce862fe3d` (1 source records).
  - [datadog/rpj3lwgx1s3s](../../data/2026/W27/datadog.json); vendor URL: https://stspg.io/gs85vq0xmxcx
- Start time basis: `vendor_reported`.

<a id="ci-d09b071c1b693c66719e7476"></a>

- Canonical ID: `ci-d09b071c1b693c66719e7476` (1 source records).
  - [stripe/wv134rlqz1ds](../../data/2026/W27/stripe.json); vendor URL: https://stspg.io/9rxrgvlbpn8g
- Start time basis: `vendor_reported`.

<a id="ci-0caf877c633677e5007804f9"></a>

- Canonical ID: `ci-0caf877c633677e5007804f9` (1 source records).
  - [digitalocean/7vmz1hmlzzrc](../../data/2026/W27/digitalocean.json); vendor URL: https://stspg.io/1ggj0xnb3n4v
- Start time basis: `vendor_reported`.

<a id="ci-e87283eef705d3197ae7093a"></a>

- Canonical ID: `ci-e87283eef705d3197ae7093a` (1 source records).
  - [datadog/41jdllgcbk9s](../../data/2026/W27/datadog.json); vendor URL: https://stspg.io/hsz1r8ww26gn
- Start time basis: `vendor_reported`.

<a id="ci-1cef651967c71ea82226d189"></a>

- Canonical ID: `ci-1cef651967c71ea82226d189` (1 source records).
  - [github/mc4yx9f8x1hk](../../data/2026/W27/github.json); vendor URL: https://stspg.io/vbk1zdbpmv6m
- Start time basis: `vendor_reported`.

<a id="ci-9a693a2d5dccf1fc1239cdcb"></a>

- Canonical ID: `ci-9a693a2d5dccf1fc1239cdcb` (1 source records).
  - [yandexcloud/1861](../../data/2026/W27/yandexcloud.json); vendor URL: https://status.yandex.cloud/en/incidents/1861
- Start time basis: `vendor_reported`.

<a id="ci-173ff1412f59f59df9f2d738"></a>

- Canonical ID: `ci-173ff1412f59f59df9f2d738` (1 source records).
  - [grafana/s8cx9hvdxd5v](../../data/2026/W27/grafana.json); vendor URL: https://stspg.io/lwn65pz9737z
- Start time basis: `vendor_reported`.

## Sources

| Vendor | Status | Source records seen | Parse errors |
|--------|--------|----------------|--------------|
| anthropic | ok | 50 | 0 |
| atlassian | ok | 34 | 0 |
| aws | ok | 6 | 0 |
| circleci | ok | 50 | 0 |
| cloudflare | ok | 50 | 0 |
| datadog | ok | 50 | 0 |
| digitalocean | ok | 50 | 0 |
| gcp | ok | 8 | 0 |
| github | ok | 50 | 0 |
| grafana | ok | 50 | 0 |
| openai | ok | 25 | 0 |
| slack | ok | 31 | 0 |
| stripe | ok | 50 | 0 |
| yandexcloud | ok | 50 | 0 |
