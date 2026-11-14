# Week 2026-W28 — top 5 of 10 above threshold

2026-07-06 to 2026-07-12 (UTC).

14 source records (14 saved revisions) → 14 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## 1. grafana — Issue with Dashboard Views Being Registered

- URL: https://stspg.io/y7gx5xr15rmm
- Window: 2026-07-10 15:57 to 2026-07-13 05:23 UTC (2d 13h 25m)
- Impact: minor, status: resolved
- Components: AWS Australia - prod-ap-southeast-2, AWS Australia - prod-au-southeast-1, AWS Brazil - prod-sa-east-1, AWS Canada - prod-ca-east-0, AWS Germany - prod-eu-west-2, AWS Germany - prod-eu-west-4, AWS India - prod-ap-south-1, AWS Ireland - prod-eu-west-6, AWS Japan - prod-ap-northeast-0, AWS Singapore - prod-ap-southeast-1, AWS Sweden - prod-eu-north-0, AWS Switzerland - prod-eu-central-0, AWS UAE - prod-me-central-1, AWS UK - prod-gb-south-1, AWS US East - prod-us-east-0, AWS US East - prod-us-east-2, AWS US West - prod-us-west-0, Azure Netherlands - prod-eu-west-3, Azure US Central - us-central2, Federal Cloud - AWS US Gov West, GCP Australia - prod-au-southeast-0, GCP Belgium - prod-eu-west-0, GCP Brazil - prod-sa-east-0, GCP India - prod-ap-south-0, GCP Singapore - prod-ap-southeast-0, GCP UK - prod-gb-south-0, GCP US Central - prod-us-central-0, GCP US Central - prod-us-central-3, GCP US Central - prod-us-central-4, GCP US East - prod-us-east-1, play.grafana.org

<a id="ci-c4bd5cd42566082c25028c4d"></a>

- Canonical ID: `ci-c4bd5cd42566082c25028c4d` (1 source records).
  - [grafana/rhrk2ck6ly0y](../../data/2026/W28/grafana.json); vendor URL: https://stspg.io/y7gx5xr15rmm
- Start time basis: `vendor_reported`.

  - `2026-07-10 15:57` investigating — For some dashboards, the Views/Error counts in the Dashboards/Folder list renders as `-` and never updates, even after dashboards are viewed repeatedly This is ultimately causing inaccurate or missing view counts.
  - `2026-07-10 21:40` investigating — We are continuing to investigate the underlying cause of this issue. At this time, the scope and impact remain unchanged, and we have no new information to share. We will provide another update as soon as more information becomes available.
  - `2026-07-13 05:21` investigating — We are continuing to investigate the underlying cause of this issue. At this time, the scope and impact remain unchanged, and we have no new information to share. We will provide another update as soon as more information becomes available.
  - `2026-07-13 05:23` resolved — At this stage, we are considering the incident resolved.

## 2. grafana — IRM Mobile App forcing some users to logout

- URL: https://stspg.io/lk4y3c36tmj2
- Window: 2026-07-10 20:27 to 2026-07-13 13:00 UTC (2d 16h 33m)
- Impact: minor, status: resolved
- Components: Incident Management and Response (IRM)

<a id="ci-d76e44fd370ab6c9c63313cc"></a>

- Canonical ID: `ci-d76e44fd370ab6c9c63313cc` (1 source records).
  - [grafana/dy7gggxhx3b1](../../data/2026/W28/grafana.json); vendor URL: https://stspg.io/lk4y3c36tmj2
- Start time basis: `vendor_reported`.

  - `2026-07-10 20:27` investigating — We are aware of an issue in the latest mobile app release that is causing some users to be signed out and asked to log in again. We have reproduced the behavior and are investigating and working on a fix. We will share another update as soon as we have more information.
  - `2026-07-10 20:37` investigating — Our investigation has narrowed the impact to version 2.39.4 of the Grafana mobile app on iOS. Users should avoid upgrading to this version until an updated release is available. As a precaution, we also recommend ensuring you have alternative notification methods configured (such as SMS, email, or phone calls) if you rely on mobile push notifications for alerting. We are actively working on a resolution and will provide another update as soon as more information is available.
  - `2026-07-10 22:17` identified — A fix has been submitted for review and will be released as Grafana Mobile v2.39.5 once it is approved. If your app is currently on v2.39.3, we recommend not upgrading to v2.39.4 and instead waiting for v2.39.5 to become available. Users who have already been signed out by this issue can log back into the app and continue using it. We will provide another update once v2.39.5 is available.
  - `2026-07-11 03:58` monitoring — A new iOS mobile app version (v2.39.5) has been released, which includes a fix for this issue. We are monitoring the rollout and its impact to ensure the issue has been fully resolved. If you continue to experience any problems after updating to the latest version, please let us know.
  - `2026-07-13 13:00` resolved — This incident has been resolved.

## 3. grafana — PDC Degraded performance

- URL: https://stspg.io/ftjw34232ndk
- Window: 2026-07-09 11:37 to 2026-07-09 19:54 UTC (8h 17m)
- Impact: critical, status: resolved
- Components: AWS Australia - prod-au-southeast-1, AWS Germany - prod-eu-west-2, AWS Germany - prod-eu-west-4, AWS Singapore - prod-ap-southeast-1, AWS US East - prod-us-east-0, GCP US Central - prod-us-central-0

<a id="ci-6a5e89ec064af4ba534704c6"></a>

- Canonical ID: `ci-6a5e89ec064af4ba534704c6` (1 source records).
  - [grafana/c8vz3g703h4k](../../data/2026/W28/grafana.json); vendor URL: https://stspg.io/ftjw34232ndk
- Start time basis: `vendor_reported`.

  - `2026-07-09 11:37` investigating — We are currently facing performance degradation on PDC service hosted on Multiple clusters. Our Engineering Team is currently working on fixing the issue, we do apologize for any inconvenience.
  - `2026-07-09 12:36` investigating — We are still currently investigating the issue.
  - `2026-07-09 13:04` monitoring — A fix has been applied and we are seeing recovery. We will continue to monitor this.
  - `2026-07-09 14:10` monitoring — We are seeing recovery in some deployments, while others not just yet. We are continuing to monitor this incident.
  - `2026-07-09 14:44` monitoring — We are continuing to monitor for any further issues.
  - `2026-07-09 16:08` identified — We are continuing to see disruptions across multiple deployments, ranging from degraded performance to full outages for some customers. Additional resources have been engaged to mitigate this issue. We will post updates as they become available.
  - `2026-07-09 17:52` monitoring — We are now seeing recovery of PDC traffic to the affected customers. Our team will continue to monitor across shifts.
  - `2026-07-09 19:54` resolved — This incident has been resolved. Thank you for your patience.

## 4. github — Delays starting Actions runs

- URL: https://stspg.io/r6sj21r75r24
- Window: 2026-07-09 04:34 to 2026-07-09 13:52 UTC (9h 17m)
- Impact: critical, status: resolved
- Components: Actions, Pages

<a id="ci-27f783aabe4a50d497e4cd6a"></a>

- Canonical ID: `ci-27f783aabe4a50d497e4cd6a` (1 source records).
  - [github/cstx3v63mklm](../../data/2026/W28/github.json); vendor URL: https://stspg.io/r6sj21r75r24
- Start time basis: `vendor_reported`.

  - `2026-07-09 04:34` investigating — We are investigating reports of degraded performance for Actions
  - `2026-07-09 04:51` investigating — Approximately 5% of GitHub Actions runs on GitHub-hosted runners are experiencing run start delays exceeding 5 minutes. A small portion of these runs may fail after extended delays. We have identified the cause and are working on a mitigation.
  - `2026-07-09 06:01` investigating — We are continuing to work on a mitigation.
  - `2026-07-09 10:07` investigating — Actions is experiencing degraded availability. We are continuing to investigate.
  - `2026-07-09 10:15` investigating — Approximately 30% of GitHub Actions runs on GitHub-hosted runners are experiencing run start delays exceeding 5 minutes. A smaller percentage of those are exhausting retries and failing to start.
  - `2026-07-09 12:01` investigating — Approximately 30% of GitHub Actions runs on GitHub-hosted runners are experiencing run start delays exceeding 5 minutes. A smaller percentage of those are exhausting retries and failing to start.<br />This has caused some customers to exceed their hosted compute concurrency and experience increased impact.<br /><br />We are continuing to working on infrastructure mitigations. <br /><br />Next update in one hour.
  - `2026-07-09 12:36` investigating — Pages is experiencing degraded performance. We are continuing to investigate.
  - `2026-07-09 12:46` investigating — Actions is experiencing degraded availability. We are continuing to investigate.
  - `2026-07-09 12:54` investigating — We're seeing Actions and Pages recovery. <br /><br />For a period of approximate 20 minutes ~96% of GitHub Actions runs on GitHub-hosted runners were failing to start, but has now recovered and we are seeing jobs processing.<br /><br />GitHub pages builds were also failing during that period, but Pages are still accessible.<br /><br />We are continuing to monitor for full recovery.
  - `2026-07-09 13:16` investigating — Actions is experiencing degraded availability. We are continuing to investigate.
  - `2026-07-09 13:17` investigating — We are continuing to monitor slow recovery in Actions and Pages builds as the system works through the high volume of backlog.<br /><br />Customers may see a small rate of API and job failures as the system is recovering.<br /><br />Copilot Cloud Agent and Copilot Code Review also failed to start for approximately 30 minutes during this incident, and we are monitoring recovery.<br /><br />Pages were accessible throughout the incident.
  - `2026-07-09 13:39` monitoring — The degradation affecting Actions and Pages has been mitigated. We are monitoring to ensure stability.
  - `2026-07-09 13:41` monitoring — Actions, Pages builds, Copilot Cloud Agent, and Copilot Code review have all recovered and are mitigated.<br /><br />We are continuing to monitor to ensure full recovery, and investigating the health of the affected infrastructure.
  - `2026-07-09 13:52` resolved — On July 9, 2026, between 03:29 UTC and 13:39 UTC, GitHub Actions experienced delayed and failed job starts on GitHub-hosted runners. The incident was caused by an unhealthy state in a backend data service responsible for provisioning hosted runners, preventing runner acquisition for a subset of workloads. During most of the incident, approximately 8% of workflow runs on hosted runners were delayed by more than 5 minutes, while roughly 2% failed to start.<br /><br />At 13:39 UTC, we restored the health of the backend data replication system, allowing provisioning to recover and the accumulated workflow backlog to drain. Service performance then returned to expected levels. We are improving provisioning-service resiliency, workload distribution, and capacity balancing to reduce the likelihoo…

## 5. digitalocean — Kubernetes Deployments in NYC1

- URL: https://stspg.io/p8xctv4k5yjv
- Window: 2026-07-09 20:31 to 2026-07-11 01:27 UTC (1d 4h 56m)
- Impact: minor, status: resolved
- Components: NYC1

<a id="ci-1a942ff25a96e5821fc90b28"></a>

- Canonical ID: `ci-1a942ff25a96e5821fc90b28` (1 source records).
  - [digitalocean/2kxtxyvzzzk0](../../data/2026/W28/digitalocean.json); vendor URL: https://stspg.io/p8xctv4k5yjv
- Start time basis: `vendor_reported`.

  - `2026-07-09 20:31` investigating — Our Engineering team is investigating an issue affecting Kubernetes deployments in NYC1. Users may see intermittent DNS failures and NodeNotReady events from application workloads during this time. We apologize for the inconvenience, we'll share new information on this page as soon as it is available.
  - `2026-07-10 01:18` monitoring — The issue affecting Kubernetes deployments in NYC1 has subsided. Workloads should now be functioning normally. Our Engineering team is continuing to monitor the affected systems to confirm full resolution. We'll update this page if any further action is needed. We apologize for the inconvenience this may have caused.
  - `2026-07-11 01:27` resolved — The issue affecting Kubernetes deployments in the NYC1 region has been resolved. Our investigation found intermittent DNS timeouts affecting a small number of DOKS clusters, with affected worker nodes running on shared-CPU Droplets. This is a documented limitation for latency-sensitive cluster DNS workloads such as CoreDNS. The affected clusters are currently functional. To reduce the risk of recurrence, we recommend running CoreDNS on non-shared/dedicated CPU node pools and using sufficient CoreDNS replicas. If you continue to see DNS failures or NodeNotReady events, please open a support ticket so we can investigate that cluster specifically.

## Also above threshold (5)

- `grafana` Grafana rulers crash-looping on prometheus (5h 6m) — https://stspg.io/p3xgglcjcppk
- `github` Actions and Codespaces APIs experiencing partial failures (2h 2m) — https://stspg.io/rmlwm65yzzjj
- `grafana` Delayed ingestion and recording rule evaluation failures for Mimir in prod-ap-south-1 (2h 39m) — https://stspg.io/wf0b60558zyz
- `datadog` Delayed data streams data (1h 50m) — https://stspg.io/y5323mlt6x4g
- `grafana` Network Degredation in prod-us-central-0 (0m) — https://stspg.io/b7q606nsf9sy
