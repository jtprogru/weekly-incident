# Week 2026-W28

2026-07-06 to 2026-07-12 (UTC). Generated 2026-11-14 00:00.

14 canonical incidents, 10 above threshold. All 13 sources responded.

14 source records (14 saved revisions) → 14 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | grafana | [Issue with Dashboard Views Being Registered](https://stspg.io/y7gx5xr15rmm) | 2026-07-10 15:57 | 2d 13h 25m | minor | AWS Australia - prod-ap-southeast-2, AWS Australia - prod-au-southeast-1, AWS Brazil - prod-sa-east-1, AWS Canada - prod-ca-east-0, AWS Germany - prod-eu-west-2, AWS Germany - prod-eu-west-4, AWS India - prod-ap-south-1, AWS Ireland - prod-eu-west-6, AWS Japan - prod-ap-northeast-0, AWS Singapore - prod-ap-southeast-1, AWS Sweden - prod-eu-north-0, AWS Switzerland - prod-eu-central-0, AWS UAE - prod-me-central-1, AWS UK - prod-gb-south-1, AWS US East - prod-us-east-0, AWS US East - prod-us-east-2, AWS US West - prod-us-west-0, Azure Netherlands - prod-eu-west-3, Azure US Central - us-central2, Federal Cloud - AWS US Gov West, GCP Australia - prod-au-southeast-0, GCP Belgium - prod-eu-west-0, GCP Brazil - prod-sa-east-0, GCP India - prod-ap-south-0, GCP Singapore - prod-ap-southeast-0, GCP UK - prod-gb-south-0, GCP US Central - prod-us-central-0, GCP US Central - prod-us-central-3, GCP US Central - prod-us-central-4, GCP US East - prod-us-east-1, play.grafana.org | 114235 = 3685 × 31 × 1.0 |
| 2 | grafana | [IRM Mobile App forcing some users to logout](https://stspg.io/lk4y3c36tmj2) | 2026-07-10 20:27 | 2d 16h 33m | minor | Incident Management and Response (IRM) | 3873 = 3873 × 1 × 1.0 |
| 3 | grafana | [PDC Degraded performance](https://stspg.io/ftjw34232ndk) | 2026-07-09 11:37 | 8h 17m | critical | AWS Australia - prod-au-southeast-1, AWS Germany - prod-eu-west-2, AWS Germany - prod-eu-west-4, AWS Singapore - prod-ap-southeast-1, AWS US East - prod-us-east-0, GCP US Central - prod-us-central-0 | 2982 = 497 × 6 × 1.0 |
| 4 | github | [Delays starting Actions runs](https://stspg.io/r6sj21r75r24) | 2026-07-09 04:34 | 9h 17m | critical | Actions, Pages | 1671 = 557 × 2 × 1.5 |
| 5 | digitalocean | [Kubernetes Deployments in NYC1](https://stspg.io/p8xctv4k5yjv) | 2026-07-09 20:31 | 1d 4h 56m | minor | NYC1 | 1389 = 1736 × 1 × 0.8 |
| 6 | grafana | [Grafana rulers crash-looping on prometheus](https://stspg.io/p3xgglcjcppk) | 2026-07-08 12:17 | 5h 6m | minor | AWS US East - prod-us-east-0: Rule Evaluation, AWS US West - prod-us-west-0: Rule Evaluation | 612 = 306 × 2 × 1.0 |
| 7 | github | [Actions and Codespaces APIs experiencing partial failures](https://stspg.io/rmlwm65yzzjj) | 2026-07-07 14:14 | 2h 2m | major | Actions, Codespaces | 366 = 122 × 2 × 1.5 |
| 8 | grafana | [Delayed ingestion and recording rule evaluation failures for Mimir in prod-ap-south-1](https://stspg.io/wf0b60558zyz) | 2026-07-09 10:28 | 2h 39m | minor | AWS India - prod-ap-south-1: Alertmanager and Rules Configuration API, AWS India - prod-ap-south-1: Rule Evaluation | 318 = 159 × 2 × 1.0 |
| 9 | datadog | [Delayed data streams data](https://stspg.io/y5323mlt6x4g) | 2026-07-10 14:09 | 1h 50m | major | Data Streams Monitoring | 110 = 110 × 1 × 1.0 |
| 10 | grafana | [Network Degredation in prod-us-central-0](https://stspg.io/b7q606nsf9sy) | 2026-07-10 21:31 | 0m | major | — | 0 = 0 × 1 × 1.0 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. grafana — Issue with Dashboard Views Being Registered

- URL: https://stspg.io/y7gx5xr15rmm
- Started: 2026-07-10 15:57 UTC
- Resolved: 2026-07-13 05:23 UTC
- Duration: 2d 13h 25m
- Impact: minor, status: resolved
- Components: AWS Australia - prod-ap-southeast-2, AWS Australia - prod-au-southeast-1, AWS Brazil - prod-sa-east-1, AWS Canada - prod-ca-east-0, AWS Germany - prod-eu-west-2, AWS Germany - prod-eu-west-4, AWS India - prod-ap-south-1, AWS Ireland - prod-eu-west-6, AWS Japan - prod-ap-northeast-0, AWS Singapore - prod-ap-southeast-1, AWS Sweden - prod-eu-north-0, AWS Switzerland - prod-eu-central-0, AWS UAE - prod-me-central-1, AWS UK - prod-gb-south-1, AWS US East - prod-us-east-0, AWS US East - prod-us-east-2, AWS US West - prod-us-west-0, Azure Netherlands - prod-eu-west-3, Azure US Central - us-central2, Federal Cloud - AWS US Gov West, GCP Australia - prod-au-southeast-0, GCP Belgium - prod-eu-west-0, GCP Brazil - prod-sa-east-0, GCP India - prod-ap-south-0, GCP Singapore - prod-ap-southeast-0, GCP UK - prod-gb-south-0, GCP US Central - prod-us-central-0, GCP US Central - prod-us-central-3, GCP US Central - prod-us-central-4, GCP US East - prod-us-east-1, play.grafana.org
- Canonical incident: [ci-c4bd5cd42566082c25028c4d](#ci-c4bd5cd42566082c25028c4d)

- `2026-07-10 15:57` **investigating** — For some dashboards, the Views/Error counts in the Dashboards/Folder list renders as `-` and never updates, even after dashboards are viewed repeatedly This is ultimately causing inaccurate or missing view counts.
- `2026-07-10 21:40` **investigating** — We are continuing to investigate the underlying cause of this issue. At this time, the scope and impact remain unchanged, and we have no new information to share. We will provide another update as soon as more information becomes available.
- `2026-07-13 05:21` **investigating** — We are continuing to investigate the underlying cause of this issue. At this time, the scope and impact remain unchanged, and we have no new information to share. We will provide another update as soon as more information becomes available.
- `2026-07-13 05:23` **resolved** — At this stage, we are considering the incident resolved.

### 2. grafana — IRM Mobile App forcing some users to logout

- URL: https://stspg.io/lk4y3c36tmj2
- Started: 2026-07-10 20:27 UTC
- Resolved: 2026-07-13 13:00 UTC
- Duration: 2d 16h 33m
- Impact: minor, status: resolved
- Components: Incident Management and Response (IRM)
- Canonical incident: [ci-d76e44fd370ab6c9c63313cc](#ci-d76e44fd370ab6c9c63313cc)

- `2026-07-10 20:27` **investigating** — We are aware of an issue in the latest mobile app release that is causing some users to be signed out and asked to log in again. We have reproduced the behavior and are investigating and working on a fix. We will share another update as soon as we have more information.
- `2026-07-10 20:37` **investigating** — Our investigation has narrowed the impact to version 2.39.4 of the Grafana mobile app on iOS. Users should avoid upgrading to this version until an updated release is available. As a precaution, we also recommend ensuring you have alternative notification methods configured (such as SMS, email, or phone calls) if you rely on mobile push notifications for alerting. We are actively working on a resolution and will provide another update as soon as more information is available.
- `2026-07-10 22:17` **identified** — A fix has been submitted for review and will be released as Grafana Mobile v2.39.5 once it is approved. If your app is currently on v2.39.3, we recommend not upgrading to v2.39.4 and instead waiting for v2.39.5 to become available. Users who have already been signed out by this issue can log back into the app and continue using it. We will provide another update once v2.39.5 is available.
- `2026-07-11 03:58` **monitoring** — A new iOS mobile app version (v2.39.5) has been released, which includes a fix for this issue. We are monitoring the rollout and its impact to ensure the issue has been fully resolved. If you continue to experience any problems after updating to the latest version, please let us know.
- `2026-07-13 13:00` **resolved** — This incident has been resolved.

### 3. grafana — PDC Degraded performance

- URL: https://stspg.io/ftjw34232ndk
- Started: 2026-07-09 11:37 UTC
- Resolved: 2026-07-09 19:54 UTC
- Duration: 8h 17m
- Impact: critical, status: resolved
- Components: AWS Australia - prod-au-southeast-1, AWS Germany - prod-eu-west-2, AWS Germany - prod-eu-west-4, AWS Singapore - prod-ap-southeast-1, AWS US East - prod-us-east-0, GCP US Central - prod-us-central-0
- Canonical incident: [ci-6a5e89ec064af4ba534704c6](#ci-6a5e89ec064af4ba534704c6)

- `2026-07-09 11:37` **investigating** — We are currently facing performance degradation on PDC service hosted on Multiple clusters. Our Engineering Team is currently working on fixing the issue, we do apologize for any inconvenience.
- `2026-07-09 12:36` **investigating** — We are still currently investigating the issue.
- `2026-07-09 13:04` **monitoring** — A fix has been applied and we are seeing recovery. We will continue to monitor this.
- `2026-07-09 14:10` **monitoring** — We are seeing recovery in some deployments, while others not just yet. We are continuing to monitor this incident.
- `2026-07-09 14:44` **monitoring** — We are continuing to monitor for any further issues.
- `2026-07-09 16:08` **identified** — We are continuing to see disruptions across multiple deployments, ranging from degraded performance to full outages for some customers. Additional resources have been engaged to mitigate this issue. We will post updates as they become available.
- `2026-07-09 17:52` **monitoring** — We are now seeing recovery of PDC traffic to the affected customers. Our team will continue to monitor across shifts.
- `2026-07-09 19:54` **resolved** — This incident has been resolved. Thank you for your patience.

### 4. github — Delays starting Actions runs

- URL: https://stspg.io/r6sj21r75r24
- Started: 2026-07-09 04:34 UTC
- Resolved: 2026-07-09 13:52 UTC
- Duration: 9h 17m
- Impact: critical, status: resolved
- Components: Actions, Pages
- Canonical incident: [ci-27f783aabe4a50d497e4cd6a](#ci-27f783aabe4a50d497e4cd6a)

- `2026-07-09 04:34` **investigating** — We are investigating reports of degraded performance for Actions
- `2026-07-09 04:51` **investigating** — Approximately 5% of GitHub Actions runs on GitHub-hosted runners are experiencing run start delays exceeding 5 minutes. A small portion of these runs may fail after extended delays. We have identified the cause and are working on a mitigation.
- `2026-07-09 06:01` **investigating** — We are continuing to work on a mitigation.
- `2026-07-09 10:07` **investigating** — Actions is experiencing degraded availability. We are continuing to investigate.
- `2026-07-09 10:15` **investigating** — Approximately 30% of GitHub Actions runs on GitHub-hosted runners are experiencing run start delays exceeding 5 minutes. A smaller percentage of those are exhausting retries and failing to start.
- `2026-07-09 12:01` **investigating** — Approximately 30% of GitHub Actions runs on GitHub-hosted runners are experiencing run start delays exceeding 5 minutes. A smaller percentage of those are exhausting retries and failing to start.<br />This has caused some customers to exceed their hosted compute concurrency and experience increased impact.<br /><br />We are continuing to working on infrastructure mitigations. <br /><br />Next update in one hour.
- `2026-07-09 12:36` **investigating** — Pages is experiencing degraded performance. We are continuing to investigate.
- `2026-07-09 12:46` **investigating** — Actions is experiencing degraded availability. We are continuing to investigate.
- `2026-07-09 12:54` **investigating** — We're seeing Actions and Pages recovery. <br /><br />For a period of approximate 20 minutes ~96% of GitHub Actions runs on GitHub-hosted runners were failing to start, but has now recovered and we are seeing jobs processing.<br /><br />GitHub pages builds were also failing during that period, but Pages are still accessible.<br /><br />We are continuing to monitor for full recovery.
- `2026-07-09 13:16` **investigating** — Actions is experiencing degraded availability. We are continuing to investigate.
- `2026-07-09 13:17` **investigating** — We are continuing to monitor slow recovery in Actions and Pages builds as the system works through the high volume of backlog.<br /><br />Customers may see a small rate of API and job failures as the system is recovering.<br /><br />Copilot Cloud Agent and Copilot Code Review also failed to start for approximately 30 minutes during this incident, and we are monitoring recovery.<br /><br />Pages were accessible throughout the incident.
- `2026-07-09 13:39` **monitoring** — The degradation affecting Actions and Pages has been mitigated. We are monitoring to ensure stability.
- `2026-07-09 13:41` **monitoring** — Actions, Pages builds, Copilot Cloud Agent, and Copilot Code review have all recovered and are mitigated.<br /><br />We are continuing to monitor to ensure full recovery, and investigating the health of the affected infrastructure.
- `2026-07-09 13:52` **resolved** — On July 9, 2026, between 03:29 UTC and 13:39 UTC, GitHub Actions experienced delayed and failed job starts on GitHub-hosted runners. The incident was caused by an unhealthy state in a backend data service responsible for provisioning hosted runners, preventing runner acquisition for a subset of workloads. During most of the incident, approximately 8% of workflow runs on hosted runners were delayed by more than 5 minutes, while roughly 2% failed to start.<br /><br />At 13:39 UTC, we restored the health of the backend data replication system, allowing provisioning to recover and the accumulated workflow backlog to drain. Service performance then returned to expected levels. We are improving provisioning-service resiliency, workload distribution, and capacity balancing to reduce the likelihood and impact of similar incidents.

### 5. digitalocean — Kubernetes Deployments in NYC1

- URL: https://stspg.io/p8xctv4k5yjv
- Started: 2026-07-09 20:31 UTC
- Resolved: 2026-07-11 01:27 UTC
- Duration: 1d 4h 56m
- Impact: minor, status: resolved
- Components: NYC1
- Canonical incident: [ci-1a942ff25a96e5821fc90b28](#ci-1a942ff25a96e5821fc90b28)

- `2026-07-09 20:31` **investigating** — Our Engineering team is investigating an issue affecting Kubernetes deployments in NYC1. Users may see intermittent DNS failures and NodeNotReady events from application workloads during this time. We apologize for the inconvenience, we'll share new information on this page as soon as it is available.
- `2026-07-10 01:18` **monitoring** — The issue affecting Kubernetes deployments in NYC1 has subsided. Workloads should now be functioning normally. Our Engineering team is continuing to monitor the affected systems to confirm full resolution. We'll update this page if any further action is needed. We apologize for the inconvenience this may have caused.
- `2026-07-11 01:27` **resolved** — The issue affecting Kubernetes deployments in the NYC1 region has been resolved. Our investigation found intermittent DNS timeouts affecting a small number of DOKS clusters, with affected worker nodes running on shared-CPU Droplets. This is a documented limitation for latency-sensitive cluster DNS workloads such as CoreDNS. The affected clusters are currently functional. To reduce the risk of recurrence, we recommend running CoreDNS on non-shared/dedicated CPU node pools and using sufficient CoreDNS replicas. If you continue to see DNS failures or NodeNotReady events, please open a support ticket so we can investigate that cluster specifically.

### 6. grafana — Grafana rulers crash-looping on prometheus

- URL: https://stspg.io/p3xgglcjcppk
- Started: 2026-07-08 12:17 UTC
- Resolved: 2026-07-08 17:23 UTC
- Duration: 5h 6m
- Impact: minor, status: resolved
- Components: AWS US East - prod-us-east-0: Rule Evaluation, AWS US West - prod-us-west-0: Rule Evaluation
- Canonical incident: [ci-05999ebeaa673ddac57f300d](#ci-05999ebeaa673ddac57f300d)

- `2026-07-08 12:17` **investigating** — We are investigating issues with the grafana-ruler service on prod-us-east-2 and prod-us-west-0 which are causing periodic crash conditions. A code fix is currently being deployed to mitigate this
- `2026-07-08 14:23` **monitoring** — We are in the process of rolling out the fix. Our engineers are monitoring the progress.
- `2026-07-08 17:23` **resolved** — This incident has been resolved.

### 7. github — Actions and Codespaces APIs experiencing partial failures

- URL: https://stspg.io/rmlwm65yzzjj
- Started: 2026-07-07 14:14 UTC
- Resolved: 2026-07-07 16:17 UTC
- Duration: 2h 2m
- Impact: major, status: resolved
- Components: Actions, Codespaces
- Canonical incident: [ci-6ff912a7b00c20dabe200b9f](#ci-6ff912a7b00c20dabe200b9f)

- `2026-07-07 14:14` **investigating** — We are investigating reports of degraded performance for Actions and Codespaces
- `2026-07-07 14:14` **investigating** — Customers accessing the Actions and Codespaces REST APIs may see 500 class errors a small percentage of the time.<br /><br />Actions runs in progress are continuing successfully.
- `2026-07-07 14:32` **investigating** — Codespaces is experiencing degraded availability. We are continuing to investigate.
- `2026-07-07 14:47` **investigating** — Customers accessing the Actions and Codespaces REST APIs may see 500 class errors a percentage of the time. Retries may be successful.<br /><br />Actions runs and codespaces in progress are continuing successfully.<br /><br />We are continuing to investigate the source of these errors.
- `2026-07-07 15:06` **investigating** — Customers accessing Actions runners and Codespaces REST APIs continue to see 500 errors a percentage of the time. Retries may be successful.<br /><br />We continue to investigate the source of these errors.
- `2026-07-07 15:50` **investigating** — Customers will continue to see 500 errors for approximately 8% of Actions runner REST APIs and 13% of Codespaces REST APIs. Retries may be successful. We have identified a likely cause and are preparing a mitigation.
- `2026-07-07 16:00` **investigating** — We have rolled out the mitigation and are seeing recovery.
- `2026-07-07 16:02` **monitoring** — The degradation affecting Actions and Codespaces has been mitigated. We are monitoring to ensure stability.
- `2026-07-07 16:17` **resolved** — On July 7, 2026, between 14:01 UTC and 16:17 UTC the Actions and Codespaces REST APIs were degraded and returned intermittent 500-class errors for a percentage of requests. Error rates peaked at approximately 8% of Actions runner API requests and 13% of Codespaces API requests, though retries were frequently successful. In-progress Actions runs and Codespaces were not impacted and continued successfully. This was due to a recent change that did not deliver the expected performance and, under certain conditions, caused downstream errors.<br /><br />We mitigated the incident by rolling back the change, after which the affected services recovered.<br /><br />We are working to improve the resilience of our services to these conditions and to strengthen our monitoring to reduce our time to detection and mitigation of issues like this one in the future.

### 8. grafana — Delayed ingestion and recording rule evaluation failures for Mimir in prod-ap-south-1

- URL: https://stspg.io/wf0b60558zyz
- Started: 2026-07-09 10:28 UTC
- Resolved: 2026-07-09 13:07 UTC
- Duration: 2h 39m
- Impact: minor, status: resolved
- Components: AWS India - prod-ap-south-1: Alertmanager and Rules Configuration API, AWS India - prod-ap-south-1: Rule Evaluation
- Canonical incident: [ci-251300d8cca0059a06835088](#ci-251300d8cca0059a06835088)

- `2026-07-09 10:28` **investigating** — We are observing delayed ingestion and recording rule evaluation failures for Mimir in prod-ap-south-1. As of yet we have not noticed any customer impact however we are currently observing the cell.
- `2026-07-09 11:18` **monitoring** — A fix has been implemented. We are currently monitoring the results.
- `2026-07-09 13:07` **resolved** — This incident has been resolved.

### 9. datadog — Delayed data streams data

- URL: https://stspg.io/y5323mlt6x4g
- Started: 2026-07-10 14:09 UTC
- Resolved: 2026-07-10 15:59 UTC
- Duration: 1h 50m
- Impact: major, status: resolved
- Components: Data Streams Monitoring
- Canonical incident: [ci-04f70ed876fa502e03123828](#ci-04f70ed876fa502e03123828)

- `2026-07-10 14:09` **investigating** — We are currently investigating an issue affecting Data Streams Monitoring (DSM) in our US1 datacenter. As a result, data stream latency metrics may be delayed or not updating for some users. If you have monitors configured on those metrics, they may fail to alert as expected during this incident. We are not aware of any data loss at this time. Our engineering team is actively working on a fix, and we'll follow up with more information shortly.
- `2026-07-10 14:35` **investigating** — We are continuing to investigate the issue.
- `2026-07-10 15:18` **monitoring** — A fix has been implemented and we are monitoring the results. A small percentage of data sent between 6:23 UTC and 14:43 UTC is still being backfilled.
- `2026-07-10 15:59` **resolved** — This incident has been resolved.

### 10. grafana — Network Degredation in prod-us-central-0

- URL: https://stspg.io/b7q606nsf9sy
- Started: 2026-07-10 21:31 UTC
- Resolved: 2026-07-10 21:31 UTC
- Duration: 0m
- Impact: major, status: resolved
- Components: —
- Canonical incident: [ci-95692564778be6605aee742f](#ci-95692564778be6605aee742f)

- `2026-07-10 21:31` **resolved** — From approximately 20:24 UTC - 20:47 UTC a network issue in prod-us-central-0 isolated part of our infrastructure in one availability zone, temporarily cutting off connectivity to a subset of backend services. This caused elevated query latency and some delayed metric evaluations, along with related write-path errors in a logging subsystem. All systems recovered automatically once network connectivity was restored. Impacted customers may have had queries or rule evaluations fail during the incident, but ingestion was not impacted.

## Below threshold (4)

- `grafana` [Cannot access dashboard set up as a home page](https://stspg.io/c3qwt2n084hd) — 32m
- `grafana` [Delayed usage and billing data](https://stspg.io/x5s6smgnk95g) — 0m
- `grafana` [Loki Billing Data Lost](https://stspg.io/z0s7nh0tz4ss) — 0m
- `grafana` [Loki read path in prod-eu-west-2 was down](https://stspg.io/5740zjht5dcl) — 0m

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-c4bd5cd42566082c25028c4d"></a>

- Canonical ID: `ci-c4bd5cd42566082c25028c4d` (1 source records).
  - [grafana/rhrk2ck6ly0y](../../data/2026/W28/grafana.json); vendor URL: https://stspg.io/y7gx5xr15rmm
- Start time basis: `vendor_reported`.

<a id="ci-d76e44fd370ab6c9c63313cc"></a>

- Canonical ID: `ci-d76e44fd370ab6c9c63313cc` (1 source records).
  - [grafana/dy7gggxhx3b1](../../data/2026/W28/grafana.json); vendor URL: https://stspg.io/lk4y3c36tmj2
- Start time basis: `vendor_reported`.

<a id="ci-6a5e89ec064af4ba534704c6"></a>

- Canonical ID: `ci-6a5e89ec064af4ba534704c6` (1 source records).
  - [grafana/c8vz3g703h4k](../../data/2026/W28/grafana.json); vendor URL: https://stspg.io/ftjw34232ndk
- Start time basis: `vendor_reported`.

<a id="ci-27f783aabe4a50d497e4cd6a"></a>

- Canonical ID: `ci-27f783aabe4a50d497e4cd6a` (1 source records).
  - [github/cstx3v63mklm](../../data/2026/W28/github.json); vendor URL: https://stspg.io/r6sj21r75r24
- Start time basis: `vendor_reported`.

<a id="ci-1a942ff25a96e5821fc90b28"></a>

- Canonical ID: `ci-1a942ff25a96e5821fc90b28` (1 source records).
  - [digitalocean/2kxtxyvzzzk0](../../data/2026/W28/digitalocean.json); vendor URL: https://stspg.io/p8xctv4k5yjv
- Start time basis: `vendor_reported`.

<a id="ci-05999ebeaa673ddac57f300d"></a>

- Canonical ID: `ci-05999ebeaa673ddac57f300d` (1 source records).
  - [grafana/ftbs8q2gvqnx](../../data/2026/W28/grafana.json); vendor URL: https://stspg.io/p3xgglcjcppk
- Start time basis: `vendor_reported`.

<a id="ci-6ff912a7b00c20dabe200b9f"></a>

- Canonical ID: `ci-6ff912a7b00c20dabe200b9f` (1 source records).
  - [github/fz9sdc2q008p](../../data/2026/W28/github.json); vendor URL: https://stspg.io/rmlwm65yzzjj
- Start time basis: `vendor_reported`.

<a id="ci-251300d8cca0059a06835088"></a>

- Canonical ID: `ci-251300d8cca0059a06835088` (1 source records).
  - [grafana/tj6b68kpdblp](../../data/2026/W28/grafana.json); vendor URL: https://stspg.io/wf0b60558zyz
- Start time basis: `vendor_reported`.

<a id="ci-4ee178971c170bfe6f38b2c4"></a>

- Canonical ID: `ci-4ee178971c170bfe6f38b2c4` (1 source records).
  - [grafana/tszxz2dzv1f3](../../data/2026/W28/grafana.json); vendor URL: https://stspg.io/c3qwt2n084hd
- Start time basis: `vendor_reported`.

<a id="ci-04f70ed876fa502e03123828"></a>

- Canonical ID: `ci-04f70ed876fa502e03123828` (1 source records).
  - [datadog/9qt73ty9p9ph](../../data/2026/W28/datadog.json); vendor URL: https://stspg.io/y5323mlt6x4g
- Start time basis: `vendor_reported`.

<a id="ci-8d3fc4832d23693822625293"></a>

- Canonical ID: `ci-8d3fc4832d23693822625293` (1 source records).
  - [grafana/js4h6gdnlm77](../../data/2026/W28/grafana.json); vendor URL: https://stspg.io/x5s6smgnk95g
- Start time basis: `vendor_reported`.

<a id="ci-801ab4fca5960e68a7127d1d"></a>

- Canonical ID: `ci-801ab4fca5960e68a7127d1d` (1 source records).
  - [grafana/c628b5qrfvd9](../../data/2026/W28/grafana.json); vendor URL: https://stspg.io/z0s7nh0tz4ss
- Start time basis: `vendor_reported`.

<a id="ci-c77c0c6f7eed8bdc0855a84a"></a>

- Canonical ID: `ci-c77c0c6f7eed8bdc0855a84a` (1 source records).
  - [grafana/vj17mtb8b5b0](../../data/2026/W28/grafana.json); vendor URL: https://stspg.io/5740zjht5dcl
- Start time basis: `vendor_reported`.

<a id="ci-95692564778be6605aee742f"></a>

- Canonical ID: `ci-95692564778be6605aee742f` (1 source records).
  - [grafana/tdzys437r2h5](../../data/2026/W28/grafana.json); vendor URL: https://stspg.io/b7q606nsf9sy
- Start time basis: `vendor_reported`.

## Sources

| Vendor | Status | Source records seen | Parse errors |
|--------|--------|----------------|--------------|
| anthropic | ok | 50 | 0 |
| atlassian | ok | 34 | 0 |
| aws | ok | 4 | 0 |
| circleci | ok | 50 | 0 |
| cloudflare | ok | 50 | 0 |
| datadog | ok | 50 | 0 |
| digitalocean | ok | 50 | 0 |
| gcp | ok | 6 | 0 |
| github | ok | 50 | 0 |
| grafana | ok | 50 | 0 |
| openai | ok | 25 | 0 |
| slack | ok | 38 | 0 |
| stripe | ok | 50 | 0 |
