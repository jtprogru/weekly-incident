# Week 2026-W34

2026-08-17 to 2026-08-23 (UTC). Generated 2026-08-25 06:22.

46 incidents collected, 19 above threshold. All 13 sources responded.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | gcp | [We are investigating an issue where customers may experience timeouts, service degradations, errors, and elevated latencies across multiple products in the us-west1 region.](https://status.cloud.google.com/incidents/utF3FMFdQfwBzJcGG6vf) | 2026-08-20 15:40 | 3h 40m | critical | AlloyDB for PostgreSQL, Apigee Edge Public Cloud, Apigee X, Artifact Registry, BigQuery Data Transfer Service, Cloud Build, Cloud Data Fusion, Cloud Filestore, Cloud Key Management Service, Cloud Monitoring, Cloud Run, Contact Center AI Platform, Dataproc Metastore, Google App Engine, Google BigQuery, Google Cloud Bigtable, Google Cloud Composer, Google Cloud Dataflow, Google Cloud Dataproc, Google Cloud Pub/Sub, Google Cloud SQL, Google Cloud Storage, Google Compute Engine, Google Kubernetes Engine, Identity and Access Management, Managed Service for Apache Kafka, Persistent Disk | 11880 = 220 × 27 × 2.0 |
| 2 | github | [Incident with GitHub.com](https://stspg.io/y1fl26l6wpzr) | 2026-08-17 13:40 | 7h 35m | critical | Git Operations, Webhooks, API Requests, Issues, Pull Requests, Actions, Pages, Copilot | 5460 = 455 × 8 × 1.5 |
| 3 | anthropic | [Degraded performance for multiple models](https://stspg.io/tcsfmtc03xgm) | 2026-08-18 16:20 | 2h 41m | minor | claude.ai, Claude API (api.anthropic.com), Claude Code, Claude Cowork | 966 = 161 × 4 × 1.5 |
| 4 | digitalocean | [Managed Databases Creation](https://stspg.io/8ld5fpw01tg4) | 2026-08-23 05:11 | 9h 21m | minor | NYC1, NYC3 | 898 = 561 × 2 × 0.8 |
| 5 | github | [Intermittent failures creating agent tasks](https://stspg.io/py1yl5mnq89c) | 2026-08-20 14:43 | 9h 54m | critical | — | 891 = 594 × 1 × 1.5 |
| 6 | digitalocean | [Major Service Interruption in MKC1](https://stspg.io/gdzqj0v18x0r) | 2026-08-19 11:10 | 8h 24m | critical | MKC1, MKC1 | 806 = 504 × 2 × 0.8 |
| 7 | aws | [Increased Error Rates](https://health.aws.amazon.com/health/status) | 2026-08-19 15:15 | 3h 32m | unknown | ec2-eu-west-2 | 424 = 212 × 1 × 2.0 |
| 8 | cloudflare | [Missing R2 Audit Logs](https://www.cloudflarestatus.com/incidents/jn37dmb9s37p) | 2026-08-20 19:19 | 3h 40m | none | Audit Logs | 396 = 220 × 1 × 1.8 |
| 9 | cloudflare | [Network Performance Issues in India](https://www.cloudflarestatus.com/incidents/xh5t9jy2cl6x) | 2026-08-20 11:38 | 3h 36m | minor | Network | 389 = 216 × 1 × 1.8 |
| 10 | github | [Intermittent failures in runner group and runner-related permissions pages](https://stspg.io/98zqb1k9jh0x) | 2026-08-18 07:40 | 4h 2m | minor | — | 363 = 242 × 1 × 1.5 |
| 11 | openai | [Unexpected logouts for some ChatGPT web users](https://status.openai.com/incidents/01M0JXWD740S0Y50DWJZS7SH75) | 2026-08-21 19:48 | 3h 26m | minor | — | 309 = 206 × 1 × 1.5 |
| 12 | circleci | [GitHub Incidents impacting CircleCI functionality](https://stspg.io/mk4lvmlx939f) | 2026-08-17 14:04 | 5h 1m | major | — | 301 = 301 × 1 × 1.0 |
| 13 | cloudflare | [Lists Service Degraded](https://www.cloudflarestatus.com/incidents/wdp9xb0250qk) | 2026-08-17 19:56 | 2h 44m | minor | Lists | 295 = 164 × 1 × 1.8 |
| 14 | openai | [Elevated Errors for Thinking mode in ChatGPT](https://status.openai.com/incidents/01M0FQAR3NNH3ANVTQMBRD47DC) | 2026-08-20 13:56 | 2h 55m | minor | — | 262 = 175 × 1 × 1.5 |
| 15 | cloudflare | [Elevated error rates in Chicago (ORD) region.](https://www.cloudflarestatus.com/incidents/877bn0l3ptp8) | 2026-08-19 21:19 | 2h 3m | minor | Network | 221 = 123 × 1 × 1.8 |
| 16 | cloudflare | [Network Performance Issues in Asia Pacific Region](https://www.cloudflarestatus.com/incidents/xl112dfsfz6q) | 2026-08-21 17:39 | 2h 0m | none | Network | 216 = 120 × 1 × 1.8 |
| 17 | anthropic | [Elevated errors on requests to multiple models](https://stspg.io/w7yghz6hfr72) | 2026-08-20 19:16 | 26m | major | claude.ai, Claude API (api.anthropic.com), Claude Code, Claude Cowork | 156 = 26 × 4 × 1.5 |
| 18 | github | [Incident with Actions](https://stspg.io/q3ck88mqw08z) | 2026-08-18 09:36 | 46m | major | — | 69 = 46 × 1 × 1.5 |
| 19 | grafana | [Cloud Logs read path outage on eu-west-2](https://stspg.io/1g4dpsdmf2cf) | 2026-08-18 22:53 | 0m | critical | — | 0 = 0 × 1 × 1.0 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. gcp — We are investigating an issue where customers may experience timeouts, service degradations, errors, and elevated latencies across multiple products in the us-west1 region.

- URL: https://status.cloud.google.com/incidents/utF3FMFdQfwBzJcGG6vf
- Started: 2026-08-20 15:40 UTC
- Resolved: 2026-08-20 19:20 UTC
- Duration: 3h 40m
- Impact: critical, status: resolved
- Components: AlloyDB for PostgreSQL, Apigee Edge Public Cloud, Apigee X, Artifact Registry, BigQuery Data Transfer Service, Cloud Build, Cloud Data Fusion, Cloud Filestore, Cloud Key Management Service, Cloud Monitoring, Cloud Run, Contact Center AI Platform, Dataproc Metastore, Google App Engine, Google BigQuery, Google Cloud Bigtable, Google Cloud Composer, Google Cloud Dataflow, Google Cloud Dataproc, Google Cloud Pub/Sub, Google Cloud SQL, Google Cloud Storage, Google Compute Engine, Google Kubernetes Engine, Identity and Access Management, Managed Service for Apache Kafka, Persistent Disk

- `2026-08-20 16:44` **investigating** — **Summary** We are investigating an issue where customers may experience timeouts, service degradations, errors, and elevated latencies across multiple products in the us-west1 region. **Description** We are experiencing an issue with multiple products, beginning on Thursday, 2026-08-20 08:40 PDT. Our engineering team continues to investigate the issue. We will provide an update by Thursday, 2026-08-20 10:30 PDT with details. We apologize to all who are affected by the disruption. **Diagnosis / Customer Symptoms** Customers in us-west1 may experience timeouts, service degradations, errors, and elevated latencies across multiple products. **Workaround** None at this time.
- `2026-08-20 17:13` **investigating** — **Summary** We are investigating an issue where customers may experience timeouts, service degradations, errors, and elevated latencies across multiple products in the us-west1 region. **Description** Mitigation actions are implemented by our engineering teams, recovery trends have been observed across infrastructure layers. Active efforts remain underway to bring impacted cloud services back to full operation. We will provide an update by Thursday, 2026-08-20 10:45 PDT with details. **Diagnosis / Customer Symptoms** Customers in us-west1 may experience timeouts, service degradations, errors, and elevated latencies across multiple products. **Workaround** We recommend customers to failover to other regions where feasible.
- `2026-08-20 17:32` **investigating** — **Summary** We are investigating an issue where customers may experience timeouts, service degradations, errors, and elevated latencies across multiple products in the us-west1 region. **Description** Mitigation actions are implemented by our engineering teams, recovery trends have been observed across infrastructure layers. Active efforts remain underway to bring impacted cloud services back to full operation. We will provide an update by Thursday, 2026-08-20 11:00 PDT with details. **Diagnosis / Customer Symptoms** Customers in us-west1 may experience timeouts, service degradations, errors, and elevated latencies across multiple products. **Workaround** We recommend customers to failover to other regions where feasible.
- `2026-08-20 18:03` **investigating** — **Summary** The issue causing timeouts, degradations, errors, and latencies across multiple us-west1 products has been mitigated and we are working to recover all products. **Description** Mitigation actions have been completed by our engineering teams and we are seeing recovery from multiple products. We are continuing to work to recover the remaining products. We will provide an update by Thursday, 2026-08-20 11:45 PDT with details. **Diagnosis / Customer Symptoms** Customers in us-west1 may experience timeouts, service degradations, errors, and elevated latencies across multiple products. **Workaround** No workarounds needed at this time.
- `2026-08-20 18:53` **investigating** — **Summary** The issue causing timeouts, degradations, errors, and latencies across multiple us-west1 products has been mitigated and we are working to recover all products. **Description** Mitigation actions have been completed by our engineering teams and we are seeing recovery from multiple products. We are continuing to work to recover the remaining products. We will provide an update by Thursday, 2026-08-20 12:30 PDT with details. **Diagnosis / Customer Symptoms** Customers in us-west1 may experience timeouts, service degradations, errors, and elevated latencies across multiple products. **Workaround** No workarounds needed at this time.
- `2026-08-20 19:21` **investigating** — **Summary** The issue causing timeouts, degradations, errors, and latencies across multiple us-west1 products has been mitigated and we are working to recover all products. **Description** Mitigation actions have been completed by our engineering teams and we are seeing recovery from multiple products. We are continuing to work to recover the remaining products. We will provide an update by Thursday, 2026-08-20 13:30 PDT with details. **Diagnosis / Customer Symptoms** Customers in us-west1 may experience timeouts, service degradations, errors, and elevated latencies across multiple products. **Workaround** No workarounds needed at this time.
- `2026-08-20 19:37` **resolved** — **Summary** The issue causing timeouts, degradations, errors, and latencies across multiple us-west1 products has been mitigated. **Description** We have mitigated the issue impacting multiple products in our us-west1 region as of Thursday, 2026-08-20 10:22 PDT. Our engineering teams have restored capacity on a planned optical maintenance that caused unexpected congestion in Dalles, Oregon metro /us-west1 region. Our systems stabilized and services have recovered once capacity was restored. Our teams are continuing to monitor for any residual impact. We will publish an analysis of this incident once we have completed our internal investigation. We thank you for your patience while we worked on resolving the issue. **Diagnosis / Customer Symptoms** Customers in us-west1 may have experienced timeouts, service degradations, errors, and elevated latencies across multiple products. **Workaround** This issue is now mitigated.
- `2026-08-25 06:20` **resolved** — # Preliminary Incident Report We sincerely apologize for the disruption this incident caused to your business operations. Recognizing your reliance on Google Cloud, we express our sincere regrets for any operational impact experienced. Our engineering teams are actively addressing the underlying root cause to prevent future recurrences. Please note that the information provided herein reflects our current understanding as of the time of publication and remains subject to revision as the investigation progresses. A comprehensive Incident Report detailing preventive measures will be issued upon conclusion of our inquiry. If you have experienced impact outside of what is listed below, please reach out to Google Cloud Support using https://cloud.google.com/support. ## Date/Time of the Issue (All time US/Pacific) * Incident Start: 20 August 2026 08:00 * Incident End: 20 August 2026 10:22 * Duration: 2 hours, 22 minutes ## Summary On Thursday, 20 August 2026, multiple Google Cloud services encountered elevated latency, provisioning failures, increased error rates, and service degradations lasting for a duration of 2 hours and 22 minutes. ## Preliminary Root Cause The disruption originated during scheduled fiber optic maintenance, which unexpectedly compromised network capacity between data centers within the us-west1 region. Automated rerouting mechanisms failed to properly redistribute traffic to alternate capacity, resulting in network congestion as volumes exceeded available bandwidth in the affected area. This underlying network degradation subsequently impacted higher-level service components through request throttling, increased latency, and cascading retries. Consequently, both control plane operations and data plane requests failed to execute successfully across several dependent Google Cloud services, producing the observed latency and elevated error rates. ## Remediation Internal monitoring alerted Google engineers to the incident, confirming that multiple services across us-west1 were affected. To mitigate the immediate operational impact, engineers initiated traffic draining protocols, re-routing service workloads away from the degraded network infrastructure. Following the restoration of inter-campus network capacity and the resolution of the core issue, engineers systematically restored traffic to the us-west1 region and confirmed service normalization. Long-term engineering solutions are currently being finalized and assigned to address the root cause and strengthen system resilience regarding fiber maintenance procedures. ## Description of Impact On Thursday, August 20, between 08:00 and 10:22 US/Pacific, Google Cloud customers in the us-west1 region encountered elevated latency, provisioning failures, and increased error rates. * Proportionate Impact / Error Rates: The precise scope of impact—including the proportion of affected projects and applications—and detailed error metrics will be published in the comprehensive Root Cause Analysis, as the reduced capacity caused request throttling, latency spikes, and increased retry volume. * Scope Exclusion: Services and workloads operating in regions other than us-west1 remained fully operational and unaffected. **Affected Services and Features** The following services experienced elevated latencies and/or increased error rates, across their respective data planes and control planes: * AlloyDB * Apache Kafka * Apigee Edge Public Cloud * Apigee X * Artifact Registry * BigQuery & BigQuery Data Transfer Service * Cloud Build * Cloud Data Fusion * Cloud Dataflow * Cloud Filestore * Cloud Key Management Service (KMS) * Cloud Monitoring * Cloud Run * Cloud SQL * Contact Center AI Platform * Dataproc Metastore * Google App Engine * Google Cloud Bigtable * Google Cloud Pub/Sub * Google Cloud Storage (GCS) * Google Compute Engine (GCE) * Google Kubernetes Engine (GKE) * Identity and Access Management (IAM) * Managed Airflow (Cloud Composer) *Managed Service for Apache Spark (Dataproc) Persistent Disk

### 2. github — Incident with GitHub.com

- URL: https://stspg.io/y1fl26l6wpzr
- Started: 2026-08-17 13:40 UTC
- Resolved: 2026-08-17 21:15 UTC
- Duration: 7h 35m
- Impact: critical, status: resolved
- Components: Git Operations, Webhooks, API Requests, Issues, Pull Requests, Actions, Pages, Copilot

- `2026-08-17 13:40` **investigating** — We are investigating reports of impacted performance for some GitHub services.
- `2026-08-17 13:41` **investigating** — API Requests is experiencing degraded performance. We are continuing to investigate.
- `2026-08-17 13:42` **investigating** — Actions is experiencing degraded performance. We are continuing to investigate.
- `2026-08-17 13:44` **investigating** — Webhooks is experiencing degraded performance. We are continuing to investigate.
- `2026-08-17 13:45` **investigating** — We are seeing an approximate 20% error rate across numerous experiences including Pull Requests, Issues, and others. Investigations are currently under way and we will be posting updates as they become available
- `2026-08-17 13:46` **investigating** — Issues is experiencing degraded performance. We are continuing to investigate.
- `2026-08-17 13:58` **investigating** — Pull Requests is experiencing degraded performance. We are continuing to investigate.
- `2026-08-17 14:04` **investigating** — We are experiencing high error rates around 20% for web experiences and api traffic. Archive downloads and raw repository content downloads are experiencing an approximate 50% error rate. Investigations are on-going into the root cause, and updates will continue to be provided as we investigate.
- `2026-08-17 14:24` **investigating** — We are experiencing high error rates around 20% for web experiences and api traffic. Archive downloads and raw repository content downloads are experiencing an approximate 50% error rate. SAML and OIDC authentication, SCIM, and Team Sync are also impacted. Investigations are on-going and we will continue to provide updates as we discover more information.
- `2026-08-17 14:31` **investigating** — Copilot is experiencing degraded availability. We are continuing to investigate.
- `2026-08-17 14:45` **investigating** — Pull Requests is experiencing degraded availability. We are continuing to investigate.
- `2026-08-17 14:49` **investigating** — Issues is experiencing degraded availability. We are continuing to investigate.
- `2026-08-17 14:54` **investigating** — Pull Requests is experiencing degraded availability. We are continuing to investigate.
- `2026-08-17 14:58` **investigating** — Actions is experiencing degraded availability. We are continuing to investigate.
- `2026-08-17 14:58` **investigating** — We are experiencing high error rates around 20% for web experiences and api traffic. Archive downloads and raw repository content downloads are experiencing an approximate 50% error rate. SAML and OIDC authentication, SCIM, and Team Sync are also impacted. We are currently performing mitigations based on our investigation thus far and are monitoring for improvement.
- `2026-08-17 14:58` **investigating** — Webhooks is experiencing degraded availability. We are continuing to investigate.
- `2026-08-17 15:01` **investigating** — API Requests is experiencing degraded availability. We are continuing to investigate.
- `2026-08-17 15:10` **investigating** — Pages is experiencing degraded performance. We are continuing to investigate.
- `2026-08-17 15:21` **investigating** — Git Operations is experiencing degraded performance. We are continuing to investigate.
- `2026-08-17 15:40` **investigating** — Webhooks is experiencing degraded performance. We are continuing to investigate.
- `2026-08-17 15:42` **investigating** — We are experiencing high error rates around 20% for web experiences and api traffic. Archive downloads and raw repository content downloads are experiencing an approximate 50% error rate. SAML and OIDC authentication, SCIM, and Team Sync are also impacted. We are currently performing mitigations and will post updates as we progress.
- `2026-08-17 16:16` **investigating** — We are experiencing high error rates around 20% for web experiences and api traffic. Archive downloads and raw repository content downloads are experiencing an approximate 50% error rate. SAML and OIDC authentication, SCIM, and Team Sync are also impacted. We are still working to identify the root cause and will continue to post updates as we learn more and perform mitigation.
- `2026-08-17 16:36` **investigating** — We identified the problematic component and have taken corrective actions. There are strong signs of recovery but we are still working to completely restore service, with error rates still remaining slightly elevated. We will post further updates as recovery continues.
- `2026-08-17 16:59` **investigating** — The degradation affecting API Requests, Actions, Git Operations, Issues, Pages, Pull Requests and Webhooks has been mitigated. We are monitoring to ensure stability.
- `2026-08-17 17:30` **investigating** — Git Operations is experiencing degraded performance. We are continuing to investigate.
- `2026-08-17 17:34` **investigating** — We identified the problematic component and have taken corrective actions, but we are seeing residual impact across numerous services. We are continuing to apply additional mitigations and investigate the remaining impact.
- `2026-08-17 17:36` **investigating** — Issues is experiencing degraded performance. We are continuing to investigate.
- `2026-08-17 18:11` **investigating** — We identified the problematic component and have taken corrective actions, but we are seeing residual impact in the form of sporadic authentication failures. We are continuing to apply additional mitigations and investigate the remaining impact.
- `2026-08-17 18:23` **investigating** — The degradation affecting Git Operations has been mitigated. We are monitoring to ensure stability.
- `2026-08-17 18:48` **investigating** — API Requests is experiencing degraded availability. We are continuing to investigate.
- `2026-08-17 19:01` **investigating** — API Requests is operating normally.
- `2026-08-17 19:13` **investigating** — We are continuing to investigate sporadic authentication failures. We have partially disabled authentication token retries and have seen improvement, and we are monitoring impact before fully applying this mitigation.
- `2026-08-17 20:08` **investigating** — We are continuing to investigate sporadic failures affecting Copilot authentication in some applications. Copilot usage via the GitHub CLI and GitHub App are unaffected.
- `2026-08-17 20:22` **investigating** — Issues is operating normally.
- `2026-08-17 20:45` **investigating** — We are continuing to apply mitigations to address sporadic Copilot authentication failures in some applications. We expect full recovery within the next 30 minutes. Copilot usage via the GitHub CLI and GitHub App are unaffected.
- `2026-08-17 21:15` **resolved** — On August 17, 2026, from 13:28–21:15 UTC (7h 47m), GitHub.com experienced elevated errors and latency across Issues, Pull Requests, APIs, Actions, and Copilot. At peak, web/API error rates were approximately 20%, while archive and raw-content downloads reached approximately 50%. SAML/OIDC authentication, SCIM, and Team Sync were also affected, as well as Actions workflows in GHEC with Data Residency that depend on public workflow step definitions hosted on GitHub.com. Most services recovered by 16:36 UTC as our Central US datacenter recovered; Actions was degraded until approximately 18:03 UTC; and Copilot Token Service fully recovered by 21:02. <br /><br />Some of the failing traffic was moved from Central US to Northern Virginia where it was served successfully until the network failure in Central US was debugged and resolved. Delayed replies to a single internal endpoint triggered a latent retry bug in VS Code that amplified traffic by approximately 10x and caused delayed recovery for the Copilot Token Service. <br /><br />The immediate cause of the failure was network saturation on load balancers in Central US due to a new peak in traffic. Originally this was caused by an Istio sidecar pod reaching its concurrency limits and failing to auto scale correctly because of a misconfigured policy that watched host service but not sidecar limits. One failure cascaded to more and eventually four HAProxy nodes exhausted their flow limits, degrading the gateway auth path and causing widespread authentication latency and failures. The problem was worsened by optimistic retry logic which overloaded internal load balancers. Pausing HAProxy on those nodes simultaneously produced immediate broad recovery. <br /><br />The retry storm in Northern VA was fixed by 1) temporarily reducing gateway retry logic with a PR and 2) blocking inbound Copilot Token Service token requests at the load balancers with a 403, and then gradually ramping back up traffic per-site to allow callers to succeed. <br /><br />Residual Copilot authentication failures continued because client retry behavior amplified load: a failed token operation could generate many extra requests and enter a retry loop. Copilot Token Service traffic increased from a normal 7–9K RPS to 70–100K RPS. Reducing gateway authentication retries and blocking retry-triggering responses stabilized Copilot Token Service and completed recovery. <br /><br />Complicating factors that impeded recovery included a number of scraping attacks on codeload endpoints. <br /><br />To prevent recurrence, our follow-up actions include: <br /><br />- Correcting autoscaling policies to account for service-mesh sidecar concurrency and capacity. <br /><br />- Auditing Istio request, concurrency, and scaling limits across affected services. <br /><br />- Reviewing retry limits and backoff behavior across gateways and clients. <br /><br />- Addressing the VS Code retry behavior that amplified Copilot token traffic. <br /><br />- Improving load-balancer capacity monitoring and regional failover safeguards.

### 3. anthropic — Degraded performance for multiple models

- URL: https://stspg.io/tcsfmtc03xgm
- Started: 2026-08-18 16:20 UTC
- Resolved: 2026-08-18 19:01 UTC
- Duration: 2h 41m
- Impact: minor, status: resolved
- Components: claude.ai, Claude API (api.anthropic.com), Claude Code, Claude Cowork

- `2026-08-18 16:20` **investigating** — We are investigating reports of degraded performance affecting multiple models. We will provide an update as soon as possible.
- `2026-08-18 16:20` **investigating** — We are investigating elevated errors on requests to Claude Mythos 5, Claude Fable 5, Claude Opus 5, Claude Sonnet 5, Claude Haiku 4.5, and other Claude models. We will provide an update as soon as possible.
- `2026-08-18 17:12` **investigating** — We are investigating elevated errors on requests to Claude Opus 5. We will provide an update as soon as possible.
- `2026-08-18 18:26` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-08-18 19:01` **resolved** — The issue affecting Claude Opus 5 has been resolved. Impact occurred from 16:11 to 18:23 UTC.

### 4. digitalocean — Managed Databases Creation

- URL: https://stspg.io/8ld5fpw01tg4
- Started: 2026-08-23 05:11 UTC
- Resolved: 2026-08-23 14:33 UTC
- Duration: 9h 21m
- Impact: minor, status: resolved
- Components: NYC1, NYC3

- `2026-08-23 05:11` **investigating** — Our Engineering team is investigating an issue with our Managed Database product. At this time, users may experience errors while creating clusters, both via Cloud Control Panel and API requests. We apologize for the inconvenience and will share an update once we have more information.
- `2026-08-23 07:06` **investigating** — We continue to investigate an issue affecting DigitalOcean Managed Databases cluster creation across multiple regions. During this time, users may experience errors when creating Managed Database clusters via the Cloud Control Panel or API. Our engineering team is implementing a mitigation and working to restore normal provisioning. We will provide another update as soon as more information is available.
- `2026-08-23 11:30` **monitoring** — Our Engineering team has implemented the necessary fixes affecting Managed Database cluster creation in the NYC3 and NYC1 regions. Users should now be able to create new Managed Database clusters. We are currently monitoring the situation to ensure that the service has returned to normal operation and remains stable. We appreciate your patience and will provide an update once the issue is fully confirmed as resolved.
- `2026-08-23 14:33` **resolved** — Our Engineering team has resolved the issue preventing cluster creation on Managed Databases in the NYC3 and NYC1 regions. Customers can now create new Managed Databases Clusters normally. Our Engineering team has monitored the environment to ensure stability. We sincerely apologize for any inconvenience this disruption may have caused to your operations. If you continue to experience issues, please open a ticket with our Support team: https://cloudsupport.digitalocean.com/s/

### 5. github — Intermittent failures creating agent tasks

- URL: https://stspg.io/py1yl5mnq89c
- Started: 2026-08-20 14:43 UTC
- Resolved: 2026-08-21 00:37 UTC
- Duration: 9h 54m
- Impact: critical, status: resolved
- Components: —

- `2026-08-20 14:43` **investigating** — We are investigating reports of impacted performance for some GitHub services.
- `2026-08-20 14:51` **investigating** — Users may experience delays when starting tasks using Copilot Cloud Agent. We are actively investigating the issue and will provide updates as we learn more.
- `2026-08-20 15:01` **investigating** — We have identified the problematic component and are working to fail over to a healthy instance. Further updates will be provided as we perform mitigations.
- `2026-08-20 15:41` **investigating** — We are experiencing issues with Copilot Cloud Agent tasks, resulting in newly started tasks not properly displaying on-going progress. These Copilot Cloud Agent tasks are still being completed correctly but lack proper visibility. We are actively investigating the issue and will provide updates as we learn more.
- `2026-08-20 16:14` **investigating** — Users are experiencing delays when starting tasks using Copilot Cloud Agent and are not be able to see the status of these tasks. Copilot Cloud Agent tasks are still being completed. We have identified the cause of the issue and are putting mitigations in place to return service to normal levels. We will provide another update about the expected recovery time shortly.
- `2026-08-20 17:05` **investigating** — We are seeing signs of recovery for Copilot Cloud Agent task status visibility, but this recovery is slower than anticipated. We are pursuing additional mitigating measures to accelerate recovery.
- `2026-08-20 17:32` **investigating** — We are observing gradual recovery for Copilot Cloud Agent task status visibility, with session output delayed approximately 1 hour. We have taken additional steps to accelerate the recovery and are continuing to monitor the impact.
- `2026-08-20 18:04` **investigating** — We are continuing to observe gradual recovery for Copilot Cloud Agent task status visibility, with session output delayed by approximately 1 hour. We have taken additional steps to accelerate the recovery and are continuing to monitor the impact.
- `2026-08-20 18:45` **investigating** — We are continuing to observe gradual recovery for Copilot Cloud Agent task status visibility, with session output delayed by approximately 1 hour. We have taken additional steps to accelerate the recovery and expect this to take effect within the next hour.
- `2026-08-20 19:35` **investigating** — We are continuing to observe gradual recovery for Copilot Cloud Agent task status visibility. Session output continues to be delayed by approximately 1 hour as our remediation steps take effect.
- `2026-08-20 20:37` **investigating** — We are seeing gradual recovery in Copilot Cloud Agent task status visibility as we deploy a fix for the root cause. Session output remains delayed by approximately one hour while remediation continues.
- `2026-08-21 00:37` **resolved** — This incident has been resolved. Thank you for your patience and understanding as we addressed this issue. A detailed root cause analysis will be shared as soon as it is available.

### 6. digitalocean — Major Service Interruption in MKC1

- URL: https://stspg.io/gdzqj0v18x0r
- Started: 2026-08-19 11:10 UTC
- Resolved: 2026-08-19 19:34 UTC
- Duration: 8h 24m
- Impact: critical, status: resolved
- Components: MKC1, MKC1

- `2026-08-19 11:10` **investigating** — Our Engineering team is currently investigating an issue in the MKC1 region affecting multiple racks and nodes. During this time, customers may experience disruptions to GPU workloads, and Kubernetes (DOKS) worker nodes may enter a NotReady state. Our team is actively working to restore connectivity and bring impacted nodes back online. If you continue to experience issues, please open a support ticket from within your account.
- `2026-08-19 12:22` **identified** — We have identified the root cause of the issue in the MKC1 region affecting multiple racks and nodes. Our Engineering team is actively implementing remediation steps to restore connectivity and bring the impacted nodes back online. During this time, customers may continue to experience disruptions to GPU workloads and Serverless Inference. Kubernetes (DOKS) worker nodes may also remain in a NotReady state, and customers may be unable to reach Kubernetes API endpoints or perform cluster-management operations in affected clusters. We will provide another update as soon as we have more information.
- `2026-08-19 15:28` **identified** — Our Engineering team continues to work on the issue affecting the MKC1 region. We are actively working to restore connectivity and bring the impacted nodes back online. We will provide another update as soon as we have more information.
- `2026-08-19 18:59` **monitoring** — At this time, the regional control plane is fully healthy. CPU Droplets, Managed databases, Load balancers, Block storage and Spaces are operating normally. Droplet create, resize and other management operations are working. Kubernetes (DOKS) control planes are reachable. Our teams are continuing to work with the facility to restore all equipment. GPU Droplets in MKC1 remain offline or unreachable. DOKS GPU worker nodes may remain NotReady, and GPU-backed inference endpoints in this region may be unavailable. We will monitor the regional control plane for a short time and then resolve this incident. GPU customers will receive personalized updates with more information in lieu of this status page. If you have questions about your affected resources, contact support and reference this incident.
- `2026-08-19 19:34` **resolved** — Our Engineering team has confirmed the regional control plane is healthy and CPU Droplets, Managed databases, Load balancers, Block storage, Kubernetes (DOKS) and Spaces are operating normally. GPU Droplets continue to be impacted and our teams are working to restore all nodes. We will communicate with GPU Droplet customers separately via Slack and email with more information and regular updates. Thank you for your patience throughout this incident. If you continue to experience any issues, please open a support ticket from within your account.

### 7. aws — Increased Error Rates

- URL: https://health.aws.amazon.com/health/status
- Started: 2026-08-19 15:15 UTC
- Resolved: 2026-08-19 18:47 UTC
- Duration: 3h 32m
- Impact: unknown, status: resolved
- Components: ec2-eu-west-2

- `2026-08-19 15:15` **investigating** — We are investigating an issue that is impacting launching new EC2 instances and resources in a newly launched Availability Zone (euw2-az4) in the EU-WEST-2 Region. During this time, affected customers may experience issues when creating or modifying resources in the Region. Other AWS services may also be impacted. For immediate recovery, we recommend that customers use alternative Availability Zones (euw2-az1, euw2-az2, and euw2-az3) where applicable. Existing running instances and resources are not affected. We will provide another update by 10:00 AM PDT, or sooner if we have additional information to share.
- `2026-08-19 17:03` **investigating** — On August 18 we launched a new Availability Zone (euw2-az4) in the EU-WEST-2 Region. After the launch, we began experiencing errors launching EC2 instances in the new Availability Zone when a default subnet is not present. We can confirm that existing running instances and resources are not affected. Workflows that automatically get a list of Availability Zones in the Region via the DescribeAvailabilityZones API and then attempt to launch new instances or create resources in the new Availability Zone may encounter errors. For EC2 instance launch failures, we are taking mitigating steps to automatically create default subnets, where one is not already present, when an EC2 instance launch is targeting the new Availability Zone. For customers and workflows that require immediate remediation <a href="https://docs.aws.amazon.com/vpc/latest/userguide/work-with-default-vpc.html#create-default-subnet">you may create a default subnet</a> in the new Availability Zone. This will enable EC2 instance launches to successfully complete. For other resources, such as Lambda functions, where the new Availability Zone is currently not supported, we recommend customers update their workflows to exclude the newly launched Availability Zone and continue resource creation using the other Availability Zones in the Region. While we don't have an exact estimate for how long our mitigation efforts will take, we will keep you up to date on our progress and provide you with another update by 1:00 PM PDT or sooner as new information becomes available.
- `2026-08-19 18:47` **resolved** — Between August 18 5:00 PM and August 19 11:00 AM PDT, we experienced elevated errors launching EC2 instances in a newly launched Availability Zone (euw2-az4) in the EU-WEST-2 Region. After the new Availability Zone launch, we began experiencing errors when using a default VPC. We discovered the root cause of the issue on August 19 at 9:00 AM and began deploying a change to resolve the issue at 9:30 AM. While the change was underway, we began to see incremental improvements in new instance launches, with full recovery at 11:00 AM. Existing running instances and resources were not affected. Some regional services, such as Lambda functions or Aurora databases, were not available at the launch of the new Availability Zone and service availability will be added over time. Customers attempting to create resources before the services become available will see a message reporting that it is not supported in the Availability Zone. The issue has been resolved and the service is operating normally.

### 8. cloudflare — Missing R2 Audit Logs

- URL: https://www.cloudflarestatus.com/incidents/jn37dmb9s37p
- Started: 2026-08-20 19:19 UTC
- Resolved: 2026-08-20 22:59 UTC
- Duration: 3h 40m
- Impact: none, status: resolved
- Components: Audit Logs

- `2026-08-20 19:19` **identified** — We have identified an issue where audit logs for R2 are missing in the dashboard, we are working to resolve the issue.
- `2026-08-20 21:59` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-08-20 22:59` **resolved** — This incident has been resolved.

### 9. cloudflare — Network Performance Issues in India

- URL: https://www.cloudflarestatus.com/incidents/xh5t9jy2cl6x
- Started: 2026-08-20 11:38 UTC
- Resolved: 2026-08-20 15:14 UTC
- Duration: 3h 36m
- Impact: minor, status: resolved
- Components: Network

- `2026-08-20 11:38` **investigating** — Cloudflare is investigating issues with network performance in India. We are working to analyse and mitigate this problem. More updates to follow shortly.
- `2026-08-20 12:01` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-08-20 15:14` **resolved** — This incident has been resolved.

### 10. github — Intermittent failures in runner group and runner-related permissions pages

- URL: https://stspg.io/98zqb1k9jh0x
- Started: 2026-08-18 07:40 UTC
- Resolved: 2026-08-18 11:42 UTC
- Duration: 4h 2m
- Impact: minor, status: resolved
- Components: —

- `2026-08-18 07:40` **investigating** — We are investigating reports of impacted performance for some GitHub services.
- `2026-08-18 07:40` **monitoring** — We are investigating reports of failure to load runner groups and runner-related permissions for customers using larger runners.
- `2026-08-18 10:41` **monitoring** — We have identified the source of a communication issue between Actions services and are working toward mitigation. Customers may experience failure to load runner groups and runner-related permissions issues when using Larger Runners.
- `2026-08-18 11:24` **monitoring** — We have applied a mitigation and are seeing recovery signals. We will continue monitoring recovery and providing updates.
- `2026-08-18 11:42` **resolved** — On August 18, 2026, between 05:02 UTC and 11:30 UTC, customers were unable to view or manage Actions Runners and Runner Groups through the GitHub UI and API. <br /><br />The issue was caused by failures in backend requests reading runner and runner group data. The failures were caused by an expired authentication certificate unique to this service. The certificate had been rotated in KeyVault, but a step to enable use at runtime had been paused to prevent recurrence of previous incidents triggered by this operation. <br /><br />The impact was mitigated by completing the enablement of the new certificate in the backend system. We have added additional monitoring to this and other certificates. This service is also in the process of being replaced as part of our availability and scale work, bringing this authentication path and secret management in line with patterns across all GitHub services.

### 11. openai — Unexpected logouts for some ChatGPT web users

- URL: https://status.openai.com/incidents/01M0JXWD740S0Y50DWJZS7SH75
- Started: 2026-08-21 19:48 UTC
- Resolved: 2026-08-21 23:15 UTC
- Duration: 3h 26m
- Impact: minor, status: resolved
- Components: —

- `2026-08-21 19:48` **investigating** — We’re investigating an issue causing some ChatGPT web users to be unexpectedly logged out when refreshing the page.
- `2026-08-21 20:36` **monitoring** — We have applied a mitigation and are seeing recovery from the issue causing some ChatGPT web users to be unexpectedly logged out when refreshing the page. We’ll continue monitoring closely to confirm the recovery is sustained.
- `2026-08-21 23:15` **resolved** — The issue causing some ChatGPT web users to be unexpectedly logged out when refreshing the page has been resolved.

### 12. circleci — GitHub Incidents impacting CircleCI functionality

- URL: https://stspg.io/mk4lvmlx939f
- Started: 2026-08-17 14:04 UTC
- Resolved: 2026-08-17 19:05 UTC
- Duration: 5h 1m
- Impact: major, status: resolved
- Components: —

- `2026-08-17 14:04` **identified** — GitHub has reported an incident that is impacting CircleCI pipeline triggering and logging in: https://www.githubstatus.com/incidents/zkxwbgr0cnmx Jobs that are in flights are running, but status reporting to GitHub Pull Requests may fail.
- `2026-08-17 14:25` **identified** — We are continuing to see high error rates on GitHub APIs and reduced webhook traffic. We will continue to provide updates as more information becomes available.
- `2026-08-17 15:49` **identified** — We are continuing to see high error rates on GitHub's APIs and reduced webhook traffic. Customers using GitHub as their VCS may experience impact to their CircleCI platform experience as a result. This is directly linked to the incident GitHub is experiencing: https://www.githubstatus.com/incidents/zkxwbgr0cnmx We will provide updates as more information becomes available.
- `2026-08-17 16:22` **monitoring** — We have started to observe improved stability in GitHub APIs. We will continue to monitor as they recover.
- `2026-08-17 17:01` **monitoring** — GitHub API error rates and latencies appear to have recovered. We will continue to monitor.
- `2026-08-17 17:23` **monitoring** — Some of GitHub's APIs are still degraded, resulting in a small number of failures relating to commit status updates and hook processing. We will continue to do our best to mitigate these effects.
- `2026-08-17 19:05` **resolved** — GitHub's APIs appear to be operating normally.

### 13. cloudflare — Lists Service Degraded

- URL: https://www.cloudflarestatus.com/incidents/wdp9xb0250qk
- Started: 2026-08-17 19:56 UTC
- Resolved: 2026-08-17 22:41 UTC
- Duration: 2h 44m
- Impact: minor, status: resolved
- Components: Lists

- `2026-08-17 19:56` **identified** — Changes to URL Lists will experience delays before taking effect. Rulesets that contain lists may result in errors when editing them. Existing Lists continue to function normally.
- `2026-08-17 21:41` **identified** — We are continuing to work toward a fix for this issue. Changes to URL Lists will experience delays before taking effect. Existing Lists continue to function normally.
- `2026-08-17 22:17` **monitoring** — A fix has been implemented and we are currently monitoring the results.
- `2026-08-17 22:41` **resolved** — This incident is now resolved.

### 14. openai — Elevated Errors for Thinking mode in ChatGPT

- URL: https://status.openai.com/incidents/01M0FQAR3NNH3ANVTQMBRD47DC
- Started: 2026-08-20 13:56 UTC
- Resolved: 2026-08-20 16:51 UTC
- Duration: 2h 55m
- Impact: minor, status: resolved
- Components: —

- `2026-08-20 13:56` **investigating** — We are investigating an issue causing some users to experience errors when using Thinking in ChatGPT.
- `2026-08-20 14:12` **investigating** — We are investigating an issue causing some users to experience errors when using Thinking mode and Image Generation in ChatGPT.
- `2026-08-20 14:53` **monitoring** — We have implemented a mitigation for the issue for elevated errors for Thinking Mode and Image Generation in ChatGPT/ChatGPT Work
- `2026-08-20 16:51` **resolved** — All impacted services have now fully recovered.

### 15. cloudflare — Elevated error rates in Chicago (ORD) region.

- URL: https://www.cloudflarestatus.com/incidents/877bn0l3ptp8
- Started: 2026-08-19 21:19 UTC
- Resolved: 2026-08-19 23:23 UTC
- Duration: 2h 3m
- Impact: minor, status: resolved
- Components: Network

- `2026-08-19 21:19` **investigating** — We are investigating reports of elevated error rates in the Chicago (ORD) region.
- `2026-08-19 23:23` **resolved** — This incident has been resolved.

### 16. cloudflare — Network Performance Issues in Asia Pacific Region

- URL: https://www.cloudflarestatus.com/incidents/xl112dfsfz6q
- Started: 2026-08-21 17:39 UTC
- Resolved: 2026-08-21 19:39 UTC
- Duration: 2h 0m
- Impact: none, status: resolved
- Components: Network

- `2026-08-21 17:39` **investigating** — Cloudflare is investigating issues with Network Performance in Asia Pacific region. We are working to analyze and mitigate this problem. More updates to follow shortly.
- `2026-08-21 17:47` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-08-21 19:39` **resolved** — This incident has been resolved.

### 17. anthropic — Elevated errors on requests to multiple models

- URL: https://stspg.io/w7yghz6hfr72
- Started: 2026-08-20 19:16 UTC
- Resolved: 2026-08-20 19:42 UTC
- Duration: 26m
- Impact: major, status: resolved
- Components: claude.ai, Claude API (api.anthropic.com), Claude Code, Claude Cowork

- `2026-08-20 19:16` **investigating** — We are investigating elevated errors on requests to some Claude models. We will provide additional updates as soon as possible.
- `2026-08-20 19:42` **resolved** — This issue has been resolved.

### 18. github — Incident with Actions

- URL: https://stspg.io/q3ck88mqw08z
- Started: 2026-08-18 09:36 UTC
- Resolved: 2026-08-18 10:23 UTC
- Duration: 46m
- Impact: major, status: resolved
- Components: —

- `2026-08-18 09:36` **investigating** — We are investigating reports of impacted performance for some GitHub services.
- `2026-08-18 10:23` **resolved** — On August 18, 2026, between 05:02 UTC and 11:30 UTC, customers were unable to run jobs on Actions Larger Runners and were unable to view or manage Actions Runners and Runner Groups through the GitHub UI and API. <br /><br />These issues were caused by failures in backend requests resolving essential metadata for starting Larger Runner workflow runs and for reading runner and runner group data. The failures were caused by an expired authentication certificate unique to this service. The certificate had been rotated in KeyVault, but a step to enable use at runtime had been paused to prevent recurrence of previous incidents that had been triggered by this operation. <br /><br />We mitigated the issues by completing the enablement of the new certificate in the backend system. We have added additional monitoring to this and other certificates. The relevant service is also in the process of being replaced as part of our availability and scale work, bringing this authentication path and secret management in line with patterns across all GitHub services.

### 19. grafana — Cloud Logs read path outage on eu-west-2

- URL: https://stspg.io/1g4dpsdmf2cf
- Started: 2026-08-18 22:53 UTC
- Resolved: 2026-08-18 22:00 UTC
- Duration: 0m
- Impact: critical, status: resolved
- Components: —

- `2026-08-18 22:53` **resolved** — We observed an issue which impacted the following service: reads of Cloud Logs in the eu-west-2 region. Affected services would include Explore, dashboards, alert evaluation when querying Logs. The time of impact lasted from approximately 22h15 and 22h23 UTC, on loki-prod-012. This incident has since been resolved.

## Below threshold (27)

- `cloudflare` [Durable Objects and Downstream Service Errors](https://www.cloudflarestatus.com/incidents/f35vbpr3dv4h) — 59m
- `anthropic` [Degraded performance for Claude Opus 5, Claude Sonnet 5](https://stspg.io/dw9p8s0wg6hh) — 1h 32m
- `anthropic` [Degraded performance for Claude Opus 5 and Claude Haiku 4.5](https://stspg.io/zg9tfmjj1th2) — 1h 19m
- `cloudflare` [Increased Errors for Durable Objects](https://www.cloudflarestatus.com/incidents/w1d9976ls02m) — 1h 11m
- `cloudflare` [Network Performance Issues in Phoenix, Los Angeles](https://www.cloudflarestatus.com/incidents/f8ms50xkfn4t) — 1h 10m
- `cloudflare` [Cloudflare Dashboard Issues](https://www.cloudflarestatus.com/incidents/p19kmglclwy4) — 1h 59m
- `cloudflare` [Cloudflare Access Browser-based RDP Issue](https://www.cloudflarestatus.com/incidents/k6v45x988py7) — 1h 41m
- `digitalocean` [Container Registry and Spaces Accessibility](https://stspg.io/2b4phnbmtm49) — 1h 52m
- `anthropic` [Elevated errors on Google connectors](https://stspg.io/lmcgl9dpjxk0) — 28m
- `cloudflare` [Increased Errors for Durable Objects in the Hong Kong region](https://www.cloudflarestatus.com/incidents/mlr8smgngv5h) — 21m
- `cloudflare` [Elevated Errors R2 and Durable Objects](https://www.cloudflarestatus.com/incidents/rq2nwg85qqn4) — 41m
- `cloudflare` [Increased Errors for Durable Objects in the Hong Kong region](https://www.cloudflarestatus.com/incidents/14kntn62njng) — 20m
- `cloudflare` [Workers AI GLM 5.2 is unavailable](https://www.cloudflarestatus.com/incidents/46j9vvprj159) — 1h 8m
- `cloudflare` [Zero Trust Client Posture Failures](https://www.cloudflarestatus.com/incidents/mts6889b2wcx) — 50m
- `cloudflare` [Increased Errors for Durable Objects and Downstream Services in Indonesia](https://www.cloudflarestatus.com/incidents/76jq00tfhd4c) — 49m
- `openai` [Elevated Codex API authentication errors](https://status.openai.com/incidents/01M0G1RZER839AZXWMYKSZF3GR) — 53m
- `openai` [Chatgpt.com is down - all signups and logins are down as of right now](https://status.openai.com/incidents/01M0E7K87VJNMGW0QTMHPEQQ39) — 52m
- `aws` [Increased Error Rates](https://health.aws.amazon.com/health/status) — 37m
- `cloudflare` [Cloudflare Lists Bulk Actions failing](https://www.cloudflarestatus.com/incidents/2hxsrlmxlf5h) — 37m
- `cloudflare` [Network connectivity issues in Dubai (DXB)](https://www.cloudflarestatus.com/incidents/xdzz2t911t1w) — 26m
- `cloudflare` [CDNJS Elevated Errors](https://www.cloudflarestatus.com/incidents/pqcnjfxrycb2) — 18m
- `circleci` [Insight data is currently delayed](https://stspg.io/gwc69d86r5sy) — 30m
- `openai` [Elevated errors deploying Sites](https://status.openai.com/incidents/01M0B4WSV41BCFZ9VDWKSMQVSP) — 20m
- `cloudflare` [Network performance issue in Singapore](https://www.cloudflarestatus.com/incidents/5g1nlg4hc2lh) — 16m
- `cloudflare` [Network performance issue in Mumbai](https://www.cloudflarestatus.com/incidents/fwvfd2zzbvg8) — 14m
- `cloudflare` [MTLS Validation issues](https://www.cloudflarestatus.com/incidents/jzv93gkq85bj) — 8m
- `cloudflare` [Elevated network errors for traffic routed between North America origins and SIN (Singapore) colo](https://www.cloudflarestatus.com/incidents/1wwc0f6m1c21) — 0m

## Sources

| Vendor | Status | Incidents seen | Parse errors |
|--------|--------|----------------|--------------|
| anthropic | ok | 50 | 0 |
| atlassian | ok | 34 | 0 |
| aws | ok | 4 | 0 |
| circleci | ok | 50 | 0 |
| cloudflare | ok | 50 | 0 |
| datadog | ok | 50 | 0 |
| digitalocean | ok | 50 | 0 |
| gcp | ok | 5 | 0 |
| github | ok | 50 | 0 |
| grafana | ok | 50 | 0 |
| openai | ok | 25 | 0 |
| slack | ok | 41 | 0 |
| stripe | ok | 50 | 0 |

