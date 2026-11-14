# Week 2026-W29

2026-07-13 to 2026-07-19 (UTC). Generated 2026-11-14 12:36.

38 canonical incidents, 26 above threshold. 2 of 14 sources unavailable: datadog, yandexcloud.

38 source records (40 saved revisions) → 38 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | grafana | [Issue with Dashboard Views Being Registered](https://stspg.io/0gjtnfjby0jy) | 2026-07-14 08:10 | 6d 2h 5m | minor | AWS Australia - prod-ap-southeast-2, AWS Australia - prod-au-southeast-1, AWS Brazil - prod-sa-east-1, AWS Canada - prod-ca-east-0, AWS Germany - prod-eu-west-2, AWS Germany - prod-eu-west-4, AWS India - prod-ap-south-1, AWS Ireland - prod-eu-west-6, AWS Japan - prod-ap-northeast-0, AWS Singapore - prod-ap-southeast-1, AWS Sweden - prod-eu-north-0, AWS Switzerland - prod-eu-central-0, AWS UAE - prod-me-central-1, AWS UK - prod-gb-south-1, AWS US East - prod-us-east-0, AWS US East - prod-us-east-2, AWS US West - prod-us-west-0, Azure Netherlands - prod-eu-west-3, Azure US Central - us-central2, Federal Cloud - AWS US Gov West, GCP Australia - prod-au-southeast-0, GCP Belgium - prod-eu-west-0, GCP Brazil - prod-sa-east-0, GCP India - prod-ap-south-0, GCP Singapore - prod-ap-southeast-0, GCP UK - prod-gb-south-0, GCP US Central - prod-us-central-0, GCP US Central - prod-us-central-3, GCP US Central - prod-us-central-4, GCP US East - prod-us-east-1, play.grafana.org | 271715 = 8765 × 31 × 1.0 |
| 2 | grafana | [Grafana Cloud login issues for users with role of None](https://stspg.io/9z649qh8csv8) | 2026-07-17 18:22 | 18h 52m | minor | AWS Australia - prod-ap-southeast-2, AWS Australia - prod-au-southeast-1, AWS Brazil - prod-sa-east-1, AWS Canada - prod-ca-east-0, AWS Germany - prod-eu-west-2, AWS Germany - prod-eu-west-4, AWS India - prod-ap-south-1, AWS Ireland - prod-eu-west-6, AWS Japan - prod-ap-northeast-0, AWS Singapore - prod-ap-southeast-1, AWS Sweden - prod-eu-north-0, AWS Switzerland - prod-eu-central-0, AWS UAE - prod-me-central-1, AWS UK - prod-gb-south-1, AWS US East - prod-us-east-0, AWS US East - prod-us-east-2, AWS US West - prod-us-west-0, Azure Netherlands - prod-eu-west-3, Azure US Central - us-central2, Federal Cloud - AWS US Gov West, GCP Australia - prod-au-southeast-0, GCP Belgium - prod-eu-west-0, GCP Brazil - prod-sa-east-0, GCP India - prod-ap-south-0, GCP Singapore - prod-ap-southeast-0, GCP UK - prod-gb-south-0, GCP US Central - prod-us-central-0, GCP US Central - prod-us-central-3, GCP US Central - prod-us-central-4, GCP US East - prod-us-east-1, play.grafana.org | 35092 = 1132 × 31 × 1.0 |
| 3 | gcp | [Google Cloud VMware Engine (GCVE), Google Cloud NetApp Volumes, and Bare Metal Solutions (BMS) services are experiencing a service outage in europe-west4-a due to a cooling failure.](https://status.cloud.google.com/incidents/3BvH3LVGcupoYqV6F4Nw) | 2026-07-15 23:57 | 12h 28m | major | Bare Metal Solution, Google Cloud NetApp Volumes, VMWare engine | 4488 = 748 × 3 × 2.0 |
| 4 | digitalocean | [Serverless Inference - Gemma4 Latency Issues Causing Timeouts & Slow Responses](https://stspg.io/rzs04wclgpfq) | 2026-07-13 19:04 | 2d 2h 42m | minor | Inference | 2434 = 3042 × 1 × 0.8 |
| 5 | anthropic | [Elevated errors for Claude Opus 4.7](https://stspg.io/8fyh12n7xnnh) | 2026-07-16 08:58 | 4h 32m | minor | Claude API (api.anthropic.com), Claude Code, Claude Console (platform.claude.com), Claude Cowork, claude.ai | 2040 = 272 × 5 × 1.5 |
| 6 | anthropic | [Elevated errors on Sonnet 5 and Haiku 4.5](https://stspg.io/bfnl00fvzpzp) | 2026-07-17 06:47 | 5h 33m | minor | Claude API (api.anthropic.com), Claude Code, Claude Cowork, claude.ai | 1998 = 333 × 4 × 1.5 |
| 7 | github | [Incident with GitHub Actions](https://stspg.io/w8d77c7t94zf) | 2026-07-19 23:34 | 5h 9m | critical | API Requests, Actions, Issues, Pages | 1854 = 309 × 4 × 1.5 |
| 8 | circleci | [Issues with network routing for Mac jobs](https://stspg.io/vg35mbl8lxpg) | 2026-07-13 21:52 | 1d 1h 26m | minor | macOS Jobs | 1526 = 1526 × 1 × 1.0 |
| 9 | digitalocean | [Block Storage Volume NYC1, NYC3, SGP1, SYD1 and BLR1](https://stspg.io/wv68wmnz55ln) | 2026-07-19 09:51 | 5h 41m | minor | BLR1, NYC1, NYC3, SGP1, SYD1 | 1364 = 341 × 5 × 0.8 |
| 10 | grafana | [Partial Outage in prod-eu-west-2](https://stspg.io/r9ysx33l8xbw) | 2026-07-15 22:24 | 1h 12m | major | AWS Germany - prod-eu-west-2, AWS Germany - prod-eu-west-2, AWS Germany - prod-eu-west-2, AWS Germany - prod-eu-west-2, AWS Germany - prod-eu-west-2, AWS Germany - prod-eu-west-2, AWS Germany - prod-eu-west-2, AWS Germany - prod-eu-west-2, AWS Germany - prod-eu-west-2: API, AWS Germany - prod-eu-west-2: Alertmanager, AWS Germany - prod-eu-west-2: Alertmanager and Rules Configuration API, AWS Germany - prod-eu-west-2: Ingestion, AWS Germany - prod-eu-west-2: Ingestion, AWS Germany - prod-eu-west-2: Metrics Generator, AWS Germany - prod-eu-west-2: Public Probes, AWS Germany - prod-eu-west-2: Querying, AWS Germany - prod-eu-west-2: Querying, AWS Germany - prod-eu-west-2: Rule Evaluation | 1296 = 72 × 18 × 1.0 |
| 11 | gcp | [Google Cloud VMware Engine (GCVE) Stretched Cluster customers are experiencing zonal outages impacting network connectivity across multiple regions.](https://status.cloud.google.com/incidents/T8gmtofFSTGT5tbhyciF) | 2026-07-14 17:00 | 10h 40m | major | VMWare engine | 1280 = 640 × 1 × 2.0 |
| 12 | grafana | [Some Reports of Grafana Not Loading.](https://stspg.io/ddf067fl4z8h) | 2026-07-15 18:07 | 2h 53m | major | AWS US East - prod-us-east-0, GCP US Central - prod-us-central-0, GCP US Central - prod-us-central-3, GCP US Central - prod-us-central-4 | 692 = 173 × 4 × 1.0 |
| 13 | grafana | [Delayed Aggregated Metrics (prod-us-central-0)](https://stspg.io/56wgtyyjvnt7) | 2026-07-15 19:19 | 8h 35m | minor | GCP US Central - prod-us-central-0: Ingestion | 515 = 515 × 1 × 1.0 |
| 14 | anthropic | [Elevated errors across Fable 5](https://stspg.io/jv85qb4xwc4x) | 2026-07-17 18:32 | 1h 11m | major | Claude API (api.anthropic.com), Claude Code, Claude Cowork, claude.ai | 426 = 71 × 4 × 1.5 |
| 15 | anthropic | [Elevated errors for multiple models](https://stspg.io/bdw3y9kd205m) | 2026-07-16 18:36 | 4h 16m | major | — | 384 = 256 × 1 × 1.5 |
| 16 | digitalocean | [Reserved IP routing in TOR1](https://stspg.io/x0s0pghvdcgv) | 2026-07-16 13:39 | 3h 44m | minor | Reserved IP, TOR1 | 358 = 224 × 2 × 0.8 |
| 17 | yandexcloud | [The problem is with the UI for the Managed Service for YDB](https://status.yandex.cloud/en/incidents/1883) | 2026-07-13 12:40 | 3h 30m | minor | Managed Service for YDB | 315 = 210 × 1 × 1.5 |
| 18 | github | [Disruption with some GitHub services](https://stspg.io/5c43kg01skdw) | 2026-07-16 09:13 | 3h 7m | critical | — | 280 = 187 × 1 × 1.5 |
| 19 | grafana | [Mimir Partial Write Outage](https://stspg.io/j81jwzycv215) | 2026-07-14 17:09 | 1h 53m | major | GCP US Central - prod-us-central-0: Ingestion, GCP US Central - prod-us-central-0: Querying | 226 = 113 × 2 × 1.0 |
| 20 | anthropic | [Enterprise SSO sign-in failures](https://stspg.io/2kbn7kkr22fx) | 2026-07-16 09:23 | 1h 3m | major | Claude Console (platform.claude.com), claude.ai | 189 = 63 × 2 × 1.5 |
| 21 | circleci | [Login issues for some Bitbucket users](https://stspg.io/lzl0fmm6xmts) | 2026-07-14 02:49 | 2h 25m | minor | Pipelines & Workflows | 145 = 145 × 1 × 1.0 |
| 22 | digitalocean | [Networking in BLR1](https://stspg.io/8yc90q76srsl) | 2026-07-19 06:57 | 2h 54m | minor | BLR1 | 139 = 174 × 1 × 0.8 |
| 23 | github | [Degraded REST API Availability](https://stspg.io/pzvjh0wyg401) | 2026-07-16 22:51 | 1h 22m | major | API Requests | 123 = 82 × 1 × 1.5 |
| 24 | grafana | [Adaptive Metrics aggregation delay in eu-west-0 region.](https://stspg.io/t147yktp9h8k) | 2026-07-16 12:56 | 2h 2m | minor | GCP Belgium - prod-eu-west-0: Ingestion | 122 = 122 × 1 × 1.0 |
| 25 | anthropic | [Elevated error rates for Sonnet 5](https://stspg.io/zfp1zk8gg7p3) | 2026-07-17 18:12 | 14m | major | Claude API (api.anthropic.com), Claude Code, Claude Console (platform.claude.com), Claude Cowork, claude.ai | 105 = 14 × 5 × 1.5 |
| 26 | github | [Actions runs are experiencing failures to start](https://stspg.io/t0tfsmbt70w7) | 2026-07-13 13:32 | 21m | major | Actions, Pages | 63 = 21 × 2 × 1.5 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. grafana — Issue with Dashboard Views Being Registered

- URL: https://stspg.io/0gjtnfjby0jy
- Started: 2026-07-14 08:10 UTC
- Resolved: 2026-07-20 10:15 UTC
- Duration: 6d 2h 5m
- Impact: minor, status: resolved
- Components: AWS Australia - prod-ap-southeast-2, AWS Australia - prod-au-southeast-1, AWS Brazil - prod-sa-east-1, AWS Canada - prod-ca-east-0, AWS Germany - prod-eu-west-2, AWS Germany - prod-eu-west-4, AWS India - prod-ap-south-1, AWS Ireland - prod-eu-west-6, AWS Japan - prod-ap-northeast-0, AWS Singapore - prod-ap-southeast-1, AWS Sweden - prod-eu-north-0, AWS Switzerland - prod-eu-central-0, AWS UAE - prod-me-central-1, AWS UK - prod-gb-south-1, AWS US East - prod-us-east-0, AWS US East - prod-us-east-2, AWS US West - prod-us-west-0, Azure Netherlands - prod-eu-west-3, Azure US Central - us-central2, Federal Cloud - AWS US Gov West, GCP Australia - prod-au-southeast-0, GCP Belgium - prod-eu-west-0, GCP Brazil - prod-sa-east-0, GCP India - prod-ap-south-0, GCP Singapore - prod-ap-southeast-0, GCP UK - prod-gb-south-0, GCP US Central - prod-us-central-0, GCP US Central - prod-us-central-3, GCP US Central - prod-us-central-4, GCP US East - prod-us-east-1, play.grafana.org
- Canonical incident: [ci-fa2aba4d84844a02ce489e4a](#ci-fa2aba4d84844a02ce489e4a)

- `2026-07-14 08:10` **investigating** — This is related to https://status.grafana.com/incidents/rhrk2ck6ly0y which was resolved by mistake. For some dashboards, the Views/Error counts in the Dashboards/Folder list renders as `-` and never updates, even after dashboards are viewed repeatedly. This is ultimately causing inaccurate or missing view counts. We are continuing to investigate the underlying cause of this issue. At this time, the scope and impact remain unchanged, and we have no new information to share. We will provide another update as soon as more information becomes available.
- `2026-07-14 14:27` **identified** — We've identified the cause of the issue impacting dashboard view and error counts in the Dashboards and Folder list. Our team is currently working on implementing a fix and validating the solution. We will provide another update within the next 24 hours, or sooner if we have additional information to share.
- `2026-07-15 12:05` **identified** — We continue to work on resolving this issue. We'll provide another update within the next 24 hours, or sooner if additional information becomes available.
- `2026-07-15 20:52` **identified** — We continue to work on resolving this issue. We'll provide another update within the next 24 hours, or sooner if additional information becomes available.
- `2026-07-16 13:34` **identified** — We continue to work on resolving the issue impacting dashboard view and error counts. There are no new updates to share at this time. Our next update will be provided within 24 hours.
- `2026-07-16 21:27` **identified** — We continue to work on resolving the issue impacting dashboard view and error counts. There are no new updates to share at this time. Our next update will be provided within 24 hours.
- `2026-07-20 08:17` **identified** — We are currently working on a fix related to this incident. There are no new updates to share at this time.
- `2026-07-20 10:15` **resolved** — A fix has been implemented and dashboard view and error counts in the Dashboards and Folder list are updating as expected. Thank you for your patience while we worked to address this issue.

### 2. grafana — Grafana Cloud login issues for users with role of None

- URL: https://stspg.io/9z649qh8csv8
- Started: 2026-07-17 18:22 UTC
- Resolved: 2026-07-18 13:15 UTC
- Duration: 18h 52m
- Impact: minor, status: resolved
- Components: AWS Australia - prod-ap-southeast-2, AWS Australia - prod-au-southeast-1, AWS Brazil - prod-sa-east-1, AWS Canada - prod-ca-east-0, AWS Germany - prod-eu-west-2, AWS Germany - prod-eu-west-4, AWS India - prod-ap-south-1, AWS Ireland - prod-eu-west-6, AWS Japan - prod-ap-northeast-0, AWS Singapore - prod-ap-southeast-1, AWS Sweden - prod-eu-north-0, AWS Switzerland - prod-eu-central-0, AWS UAE - prod-me-central-1, AWS UK - prod-gb-south-1, AWS US East - prod-us-east-0, AWS US East - prod-us-east-2, AWS US West - prod-us-west-0, Azure Netherlands - prod-eu-west-3, Azure US Central - us-central2, Federal Cloud - AWS US Gov West, GCP Australia - prod-au-southeast-0, GCP Belgium - prod-eu-west-0, GCP Brazil - prod-sa-east-0, GCP India - prod-ap-south-0, GCP Singapore - prod-ap-southeast-0, GCP UK - prod-gb-south-0, GCP US Central - prod-us-central-0, GCP US Central - prod-us-central-3, GCP US Central - prod-us-central-4, GCP US East - prod-us-east-1, play.grafana.org
- Canonical incident: [ci-6e0762625d6fd8ea5d1979c9](#ci-6e0762625d6fd8ea5d1979c9)

- `2026-07-17 18:22` **investigating** — We’re currently investigating an issue with user logins with the None role in Grafana Cloud. Our team is actively working to identify the cause. Thank you for your patience.
- `2026-07-17 19:28` **identified** — We’ve identified the cause of the issue impacting users with the role of None from logging in. Our team is currently implementing a fix.
- `2026-07-17 21:33` **identified** — We are in the process of implementing a fix and are monitoring the rollout. Thank you for your patience.
- `2026-07-18 13:15` **resolved** — Users with the None role authenticating via Grafana.com should be able to access their stacks again

### 3. gcp — Google Cloud VMware Engine (GCVE), Google Cloud NetApp Volumes, and Bare Metal Solutions (BMS) services are experiencing a service outage in europe-west4-a due to a cooling failure.

- URL: https://status.cloud.google.com/incidents/3BvH3LVGcupoYqV6F4Nw
- Started: 2026-07-15 23:57 UTC
- Resolved: 2026-07-16 12:25 UTC
- Duration: 12h 28m
- Impact: major, status: resolved
- Components: Bare Metal Solution, Google Cloud NetApp Volumes, VMWare engine
- Canonical incident: [ci-668272b64c22371acd86e65c](#ci-668272b64c22371acd86e65c)

- `2026-07-16 03:30` **investigating** — **Summary** Google Cloud VMware Engine (GCVE), Google Cloud NetApp Volumes, and Bare Metal Solution (BMS) services are experiencing temperature alerts in europe-west4-a. **Description** **GCVE:** Our engineering teams are currently working with data center facility teams to address the rising temperatures. A shutdown has already occurred for some private clouds, and we expect that others will shut down as well. **We recommend shutting down your VMware workload VMs for precaution so that recovery is safer. A shutdown of all systems started at 19:00 PDT \[4:00 Local Time (GMT+2)\]** We will notify you as soon as your Private Cloud has been powered back up and it is safe to restart the VMs. **BMS:** Some Bare Metal Solution servers are experiencing high temperatures and we are in the process of shutting them down. Customers may be unable to access their systems until we resolve the issue. **Google Cloud NetApp Volumes** Our engineering team has proactively shut down ONTAP clusters to protect the data stored on them. Customers may experience unavailability of Cloud NetApp Volumes until this issue is resolved. There is no estimated time for resolution. We will provide an update by Wednesday, 2026-07-15 21:30 PDT with current details. We apologize to all who are affected by the disruption. **Customer Symptoms** **GCVE:** Customers may have already observed a loss of connectivity to their private cloud. As a protective measure for hardware integrity and a recovery, we are shutting down all Private Clouds in this zone. **BMS:** Customers may have already lost connectivity to their BMS machines. As a protective measure for hardware integrity and a recovery, we are shutting down servers in this zone. **Google Cloud NetApp Volumes** Customers may have already lost connectivity to ONTAP clusters, or will lose connectivity to them shortly, as we proactively shut them down for their protection. Only Standard, Premium, and Extreme service levels are affected. **Workaround** There are no workarounds available for this issue at this time. Customers with multi-regional deployments are advised to route traffic to an alternate site.
- `2026-07-16 04:17` **investigating** — **Summary** Google Cloud VMware Engine (GCVE), Google Cloud NetApp Volumes, and Bare Metal Solution (BMS) services are experiencing temperature alerts in europe-west4-a **Description** The datacenter serving europe-west4-a for GCVE, BMS, and NetApp has experienced a power failure, which subsequently caused a cooling failure. Google has proactively turned down workloads in order to protect customer data from any risks posed by running infrastructure in a high temperature environment. The datacenter operator team is working diligently to bring cooling and power back online so we can safely restart affected workloads. **GCVE:** Our engineering teams have shut down private clouds, and we expect that others will shut down as well. We will notify customers as soon as their private cloud has been powered back up and it is safe to restart the VMs. **BMS:** Some Bare Metal Solution servers are experiencing high temperatures and we are in the process of shutting them down. Customers may be unable to access their systems until we resolve the issue. **Google Cloud NetApp Volumes** Our engineering team has proactively shut down ONTAP clusters to protect the data stored on them. Customers may experience unavailability of Cloud NetApp Volumes until this issue is resolved. There is no estimated time for resolution. We will provide an update by Wednesday, 2026-07-15 22:30 PDT with current details. We apologize to all who are affected by the disruption. **Customer Symptoms** **GCVE:** Customers would have already observed a loss of connectivity to their private cloud. As a protective measure for hardware integrity and a recovery, we have shut down all private clouds in this zone. **BMS:** Customers may have already lost connectivity to their BMS machines. As a protective measure for hardware integrity and a recovery, we are shutting down some servers in this zone. **Google Cloud NetApp Volumes** Customers would have already lost connectivity to ONTAP clusters, as we proactively shut them down for their protection. Only Standard, Premium, and Extreme service levels are affected. **Workaround** There are no workarounds available for this issue at this time. Customers with multi-regional deployments are advised to route traffic to an alternate site.
- `2026-07-16 04:41` **investigating** — **Summary** Google Cloud VMware Engine (GCVE), Google Cloud NetApp Volumes, and Bare Metal Solution (BMS) services are experiencing a service outage in europe-west4-a due to a cooling failure **Description** The datacenter serving europe-west4-a for GCVE, BMS, and NetApp has experienced a power failure, which subsequently caused a cooling failure. Google has proactively turned down workloads in order to protect customer data from any risks posed by running infrastructure in a high temperature environment. The datacenter operator team is working diligently to bring cooling and power back online so we can safely restart affected workloads. Some cooling has been restored, and we are working to restore affected workloads as available capacity allows. **GCVE:** Cooling has been restored at the site. As a precautionary measure, our engineering teams had shut down private clouds to protect these clouds from any damage. We are beginning service restoration and will notify customers as soon as their private cloud has been powered back up and it is safe to restart the VMs. **BMS:** Some Bare Metal Solution servers experienced high temperatures and we proactively shut them down. We are starting to bring the servers back online, and will provide an update once they are ready. Customers may be unable to access these systems until service is restored. **Google Cloud NetApp Volumes:** Our engineering team has proactively shut down ONTAP clusters to protect the data stored on them. Customers may experience unavailability of Cloud NetApp Volumes until this issue is resolved. There is no estimated time for resolution. We will provide an update by Wednesday, 2026-07-15 23:30 PDT with current details. We apologize to all who are affected by the disruption. **Customer Symptoms** **GCVE:** Customers would have already observed a loss of connectivity to their private cloud. As a protective measure for hardware integrity and recovery, we have shut down all private clouds in this zone. **BMS:** Customers may have lost connectivity to their BMS machines. As a protective measure for hardware integrity and recovery, we shut down some servers in this zone, and are actively restoring them. Customers may be unable to access these systems until restoration is complete. **Google Cloud NetApp Volumes:** Customers would have already lost connectivity to ONTAP clusters, as we proactively shut them down to protect these clusters from any damage. Only Standard, Premium, and Extreme service levels are affected. **Workaround** There are no workarounds available for this issue at this time. However, customers with multi-regional deployments are advised to route traffic to an alternate site.
- `2026-07-16 06:29` **investigating** — **Summary** Google Cloud VMware Engine (GCVE), Google Cloud NetApp Volumes, and Bare Metal Solution (BMS) services are experiencing a service outage in europe-west4-a due to a cooling failure. **Description** The datacenter serving europe-west4-a for GCVE, BMS, and NetApp has experienced a power failure, which subsequently caused a cooling failure. Google proactively turned down workloads in order to protect customer data from any risks posed by running infrastructure in a high temperature environment. Cooling and power have been restored at the site. The datacenter operator team is now working to safely restart affected workloads. **GCVE:** As a precautionary measure, our engineering teams had shut down private clouds to protect these clouds from any damage. Service restoration is still in progress. We will notify customers as soon as their private cloud has been powered back up and it is safe to restart the VMs. **BMS:** We are continuing to restore servers and monitor their health. Customers may be unable to access these systems until service restoration is complete. **Google Cloud NetApp Volumes:** Our engineering team had proactively shut down ONTAP clusters to protect the data stored on them. Customers may experience unavailability of Cloud NetApp Volumes until this issue is resolved. We will provide an update by Thursday, 2026-07-16 00:30 PDT with current details. **Customer Symptoms** **GCVE:** Customers would have already observed a loss of connectivity to their private cloud. As a protective measure for hardware integrity and recovery, we have shut down all private clouds in this zone. **BMS:** Customers may have lost connectivity to their servers. As a protective measure for hardware integrity and recovery, we shut down some servers in this zone. We are continuing to restore servers and monitor their health. Customers may be unable to access these systems until service restoration is complete. **Google Cloud NetApp Volumes:** Customers would have already lost connectivity to ONTAP clusters, as we proactively shut them down to protect these clusters from any damage. Only Standard, Premium, and Extreme service levels are affected. **Workaround** There are no workarounds available for this issue at this time. However, customers with multi-regional deployments are advised to route traffic to an alternate site.
- `2026-07-16 07:39` **investigating** — **Summary** Google Cloud VMware Engine (GCVE), Google Cloud NetApp Volumes, and Bare Metal Solution (BMS) services are experiencing a service outage in europe-west4-a due to a cooling failure. **Description** The datacenter serving europe-west4-a for GCVE, BMS, and NetApp has experienced a power failure, which subsequently caused a cooling failure. Google proactively turned down workloads in order to protect customer data from any risks posed by running infrastructure in a high temperature environment. Cooling and power have been restored at the site. The datacenter operator team is now working to safely restart affected workloads. **GCVE:** As a precautionary measure, our engineering teams had shut down private clouds to protect these clouds from any damage. Network fabric has been restored, and Private Cloud restoration remains underway. We will notify affected customers once their Private Clouds are restored to allow for VM operational verification. **BMS:** We are continuing to restore servers and monitor their health. Customers may be unable to access these systems until service restoration is complete. **Google Cloud NetApp Volumes:** Our engineering team had proactively shut down ONTAP clusters to protect the data stored on them. Recovery is currently in progress. Customers may experience unavailability of Cloud NetApp Volumes until this issue is resolved. We will provide an update by Thursday, 2026-07-16 01:30 PDT with current details. **Customer Symptoms** **GCVE:** Customers would have already observed a loss of connectivity to their private cloud. As a protective measure for hardware integrity and recovery, we have shut down all private clouds in this zone. **BMS:** Customers may have lost connectivity to their servers. As a protective measure for hardware integrity and recovery, we shut down some servers in this zone. We are continuing to restore servers and monitor their health. Customers may be unable to access these systems until service restoration is complete. **Google Cloud NetApp Volumes:** Customers would have already lost connectivity to ONTAP clusters, as we proactively shut them down to protect these clusters from any damage. Only Standard, Premium, and Extreme service levels are affected. **Workaround** There are no workarounds available for this issue at this time. However, customers with multi-regional deployments are advised to route traffic to an alternate site.
- `2026-07-16 08:32` **investigating** — **Summary** Google Cloud VMware Engine (GCVE), Google Cloud NetApp Volumes, and Bare Metal Solution (BMS) services experienced a service outage in europe-west4-a due to a cooling failure. **Description** The datacenter serving europe-west4-a for GCVE, BMS, and NetApp has experienced a power failure, which subsequently caused a cooling failure. Google proactively turned down workloads in order to protect customer data from any risks posed by running infrastructure in a high temperature environment. Cooling and power have been restored at the site. The datacenter operator team is now validating site stability. **GCVE:** Network fabric has been restored, and Private Cloud restoration remains underway. We are notifying the affected customers as their Private Clouds are restored to allow for VM operational verification. **BMS:** Server restoration is largely complete, with most systems now online and reporting a healthy status. We are actively working through the remaining validation efforts to achieve full recovery. **Google Cloud NetApp Volumes:** Service restoration is complete, and all affected servers are now operational. We continue to monitor these for full service recovery. If customers are still experiencing impact from this issue, please contact Support, and we will work with you to resolve any residual impact. We will provide an update by Thursday, 2026-07-16 02:30 PDT with current details. **Customer Symptoms** **GCVE:** Customers would have observed a loss of connectivity to their private cloud. As a protective measure for hardware integrity and recovery, we had shut down all Private Clouds in this zone. We are in the process of bringing all restored Private Clouds online. **BMS:** Customers may have lost connectivity to their servers. As a protective measure for hardware integrity and recovery, we had shut down some servers in this zone. We are continuing to restore servers and monitor their health. **Google Cloud NetApp Volumes:** Customers would have lost connectivity to ONTAP clusters, as they were proactively shut down to protect these clusters from any damage. Only Standard, Premium, and Extreme service levels were affected. Customers should now be able to access their volumes again. **Workaround** There are no workarounds available for this issue at this time. However, customers with multi-regional deployments are advised to route traffic to an alternate site.
- `2026-07-16 09:21` **investigating** — **Summary** Google Cloud VMware Engine (GCVE), Google Cloud NetApp Volumes, and Bare Metal Solution (BMS) services experienced a service outage in europe-west4-a due to a cooling failure. **Description** The datacenter serving europe-west4-a for GCVE, BMS, and NetApp has experienced a power failure, which subsequently caused a cooling failure. Google proactively turned down workloads in order to protect customer data from any risks posed by running infrastructure in a high temperature environment. Cooling and power have been restored at the site. The datacenter operator team is now validating site stability. **GCVE:** Network fabric has been restored, and Private Cloud restoration remains underway. We are notifying the affected customers as their Private Clouds are restored to allow for VM operational verification. **BMS:** Server restoration is largely complete, with most systems now online and reporting a healthy status. We are actively working through the remaining validation efforts to achieve full recovery. **Google Cloud NetApp Volumes:** Service restoration is complete, and all affected servers are now operational. We continue to monitor these for full service recovery. If customers are still experiencing impact from this issue, please contact Support, and we will work with you to resolve any residual impact. We will provide an update by Thursday, 2026-07-16 03:30 PDT with current details. **Customer Symptoms** **GCVE:** Customers would have observed a loss of connectivity to their private cloud. As a protective measure for hardware integrity and recovery, we had shut down all Private Clouds in this zone. We are in the process of bringing all restored Private Clouds online. **BMS:** Customers may have lost connectivity to their servers. As a protective measure for hardware integrity and recovery, we had shut down some servers in this zone. We are continuing to restore servers and monitor their health. **Google Cloud NetApp Volumes:** Customers would have lost connectivity to ONTAP clusters, as they were proactively shut down to protect these clusters from any damage. Only Standard, Premium, and Extreme service levels were affected. Customers should now be able to access their volumes again. **Workaround** None at this time.
- `2026-07-16 10:14` **investigating** — **Summary** Google Cloud VMware Engine (GCVE), Google Cloud NetApp Volumes, and Bare Metal Solution (BMS) services experienced a service outage in europe-west4-a due to a cooling failure. **Description** The datacenter serving europe-west4-a for GCVE, BMS, and NetApp has experienced a power failure, which subsequently caused a cooling failure. Google proactively turned down workloads in order to protect customer data from any risks posed by running infrastructure in a high temperature environment. Cooling and power have been restored at the site. The issue with GCVE and Google Cloud NetApp Volumes has been resolved for all affected users. We continue to recover the impact for BMS at the moment. **GCVE [Mitigated]:** Service restoration is complete across all the Private Clouds. Customers have been informed about this restoration and requested to verify that their VMs are fully operational. **Google Cloud NetApp Volumes [Mitigated]:** Service restoration is complete, and all affected servers are now operational. We continue to monitor these for full service recovery. **BMS:** Server restoration is largely complete, with most systems now online and reporting a healthy status. We are actively working through the remaining validation efforts to achieve full recovery. For Google Cloud NetApp Volumes and GCVE, if customers are still experiencing impact from this issue, please contact Support, and we will work with you to resolve any residual impact. We will provide the BMS restoration status update by Thursday, 2026-07-16 04:30 PDT with current details. **Customer Symptoms** The issue is now resolved for GCVE and Google Cloud NetApp Volumes. We continue to recover the impact for BMS at the moment. **GCVE [Mitigated]:** Customers would have observed a loss of connectivity to their private cloud. **Google Cloud NetApp Volumes [Mitigated]:** Customers would have lost connectivity to ONTAP clusters, as they were proactively shut down to protect these clusters from any damage. Only Standard, Premium, and Extreme service levels were affected. Customers should now be able to access their volumes again. **BMS:** Customers may have lost connectivity to their servers. As a protective measure for hardware integrity and recovery, we had shut down some servers in this zone. We are continuing to restore servers and monitor their health. **Workaround** None at this time.
- `2026-07-16 11:39` **investigating** — **Summary** Google Cloud VMware Engine (GCVE), Google Cloud NetApp Volumes, and Bare Metal Solution (BMS) services experienced a service outage in europe-west4-a due to a cooling failure. **Description** The datacenter serving europe-west4-a for GCVE, BMS, and NetApp has experienced a power failure, which subsequently caused a cooling failure. Google proactively turned down workloads in order to protect customer data from any risks posed by running infrastructure in a high temperature environment. Cooling and power have been restored at the site. The issue with GCVE and Google Cloud NetApp Volumes has been resolved for all affected users. We continue to recover the impact for BMS at the moment. **GCVE [Mitigated]:** Service restoration is complete across all the Private Clouds. Customers have been informed about this restoration and requested to verify that their VMs are fully operational. **Google Cloud NetApp Volumes [Mitigated]:** Service restoration is complete, and all affected servers are now operational. We continue to monitor these for full service recovery. **BMS:** Server restoration is largely complete, with most systems now online and reporting a healthy status. We are actively working through the remaining validation efforts to achieve full recovery. For Google Cloud NetApp Volumes and GCVE, if customers are still experiencing impact from this issue, please contact Support, and we will work with you to resolve any residual impact. We will provide the BMS restoration status update by Thursday, 2026-07-16 05:30 PDT with current details. **Customer Symptoms** The issue is now resolved for GCVE and Google Cloud NetApp Volumes. We continue to recover the impact for BMS at the moment. **GCVE [Mitigated]:** Customers would have observed a loss of connectivity to their private cloud. **Google Cloud NetApp Volumes [Mitigated]:** Customers would have lost connectivity to ONTAP clusters, as they were proactively shut down to protect these clusters from any damage. Only Standard, Premium, and Extreme service levels were affected. Customers should now be able to access their volumes again. **BMS:** Customers may have lost connectivity to their servers. As a protective measure for hardware integrity and recovery, we had shut down some servers in this zone. We are continuing to restore servers and monitor their health. **Workaround** None at this time.
- `2026-07-16 12:25` **investigating** — **Summary** Google Cloud VMware Engine (GCVE), Google Cloud NetApp Volumes, and Bare Metal Solution (BMS) services experienced a service outage in europe-west4-a due to a cooling failure. **Description** The datacenter serving europe-west4-a for GCVE, BMS, and NetApp has experienced a power failure, which subsequently caused a cooling failure. Google proactively turned down workloads in order to protect customer data from any risks posed by running infrastructure in a high temperature environment. Cooling and power have been restored at the site. The issue with GCVE and Google Cloud NetApp Volumes has been resolved for all affected users as of Thursday, 2026-07-16 02:31 PDT and 01:00 PDT respectively. The issue with BMS is largely resolved and is believed to be affecting a very small number of customers. Our Engineering will continue to recover this residual impact directly with the impacted customers. If you have questions or are still impacted, please open a case with the Support Team and we will work with you until this issue is resolved. **GCVE:** Service restoration is complete across all the Private Clouds. Customers have been informed about this restoration and requested to verify that their VMs are fully operational. **Google Cloud NetApp Volumes:** Service restoration is complete, and all affected servers are now operational. **BMS:** Service restoration is largely complete. We continue working with a very small number of remaining affected customers to achieve full recovery. We thank you for your patience while we continue to fully resolve the issue. **Customer Symptoms** The issue is now resolved for GCVE and Google Cloud NetApp Volumes. The issue with BMS is largely resolved and our engineering will continue to recover this residual impact directly with the impacted customers. **GCVE:** Customers would have observed a loss of connectivity to their private cloud. **Google Cloud NetApp Volumes:** Customers would have lost connectivity to ONTAP clusters, as they were proactively shut down to protect these clusters from any damage. Only Standard, Premium, and Extreme service levels were affected. **BMS:** Customers may have lost connectivity to their servers. As a protective measure for hardware integrity and recovery, we had shut down some servers in this zone. **Workaround** None at this time.
- `2026-07-18 00:22` **resolved** — # Preliminary Incident Report We sincerely apologize for the disruption this incident caused to your business. We know how much you rely on Google Cloud, and we regret the impact on your productivity. Please note, this information is based on our best knowledge at the time of posting and is subject to change as our investigation continues. A final Incident Report with preventative actions will be posted once our investigation is complete. If you have experienced impact outside of what is listed below, please reach out to Google Cloud Support using https://cloud.google.com/support. ## Date/Time of the Issue (All time US/Pacific) **Google Cloud VMware Engine Impact:** - Start: 15 July 2026 17:09 - End: 16 July 2026 2:33 - Duration: 9 Hours, 24 Minutes **Google Cloud NetApp Volumes Impact:** - Start: 15 July 2026 16:39 - End: 16 July 2026 01:10 - Duration: 8 Hours 31 Minutes **Bare Metal Solution Impact:** - Start: 15 July 2026 18:37 - End: 16 July 2026 07:34 - Duration: 12 Hours 57 Minutes ## Summary On Wednesday, 15 July 2026, Google Cloud VMware Engine, Bare Metal Solution, and Google Cloud NetApp Volumes experienced service interruptions for a total duration of 14 hours, 55 minutes. We are taking immediate steps to ensure this doesn’t happen again. ## Preliminary Root Cause A regional data center hosting services in europe-west4-a experienced a loss of utility power and subsequent cooling capacity. The sequence of events leading to customer impact was as follows: - An electrical fault occurred on the utility grid upstream of the data center, disrupting the electrical distribution gear and cooling equipment. - The loss of cooling infrastructure resulted in ambient temperatures rising rapidly within the affected data halls. Host servers, storage clusters, and network switches were shut down to prevent equipment damages due to the extreme heat. - The high temperatures in the data hall resulted in disruption to customer workloads and control plane operations across the affected services. Google engineers have begun a full root cause analysis and we will provide additional information once it is available. ## Remediation Google engineering teams were alerted to the issue via automated temperature and hardware unreachability telemetry starting at 16:39 US/Pacific. Engineers collaborated with the third-party facility provider to safely restore primary utility power and cooling systems, returning ambient room temperatures to safe operational levels. With the environment stabilized, field support engineers and remote teams systematically booted and verified the server hosts, storage nodes, and network fabric in a controlled sequence. Automated recovery playbooks were executed to restore underlying network switches and routing infrastructure, allowing storage appliances and compute clusters to be brought back online and verified for health. The underlying infrastructure has been fully recovered, and normal operations have been successfully restored. ## Description of Impact On 15 July 2026 from 16:39 to 16 July 2026 07:34 US/Pacific, customers experienced service disruptions across several products in europe-west4-a zone: - **Google Cloud VMware Engine:** Customers lost connectivity to their private clouds and workloads as foundational hosts and network switches were powered off for thermal protection. - **Bare Metal Solution:** Customers experienced connectivity loss to their database servers and storage appliances as the underlying hardware was powered down. - **Google Cloud NetApp Volumes:** Customers in the STANDARD, PREMIUM, and EXTREME service levels were unable to access their storage volumes. Control plane operations, including the creation of new storage pools, volumes, and backups, experienced failures in the affected region.
- `2026-07-25 13:16` **resolved** — ## \# Incident Report ## \#\# Summary On Wednesday, 15 July 2026, Google Cloud VMware Engine (GCVE), Bare Metal Solution (BMS), and Google Cloud NetApp Volumes (GCNV) experienced service interruptions for a total duration of 14 hours, 55 minutes. The root cause for this outage is a 3ms voltage drop in the power feed from the utility provider and subsequent utility breaker protective action. This incident was mitigated by the data center provider rectifying the failed systems followed by the Google engineering team restoring the services. ## \#\# Root Cause A regional data center hosting services in europe-west4-a experienced an upstream voltage transient affecting both utility power feeds A and B. During this event, the utility breakers on both A and B feeds tripped and initiated the transfer to the back-up power source DRUPS (Diesel Rotary Uninterruptible Power Supply). The transition of side B to DRUPS system was successful without any power interruption. The DRUPS back-up power system for side A failed to take over the facility load due to electrical component failures. The DRUPS failure to take over load initiated automatic transfer of the affected 3 rows from feed A to redundant power feed B. Rows 1 and 2 transferred to power feed B successfully. Row 3 failed to transfer to power feed B and experienced complete power loss due to an overload protection breaker trip. The root cause of the failed transfer on the affected row was identified as load deployment discrepancy which is under further investigation by Google engineering team. This resulted in loss of both redundant power feeds to the single row. During the voltage transient event, the server data hall experienced an increased temperature due to a cooling system failure. The chiller controller dropped offline during the voltage transient event, failing to signal the chilled water distribution pumps to restart and ultimately causing the chiller system A to shut down. The redundant source was not available due to known ongoing construction work at the facility. This resulted in the data hall temperatures reaching 44°C and subsequent shutdown of the affected data hall machines. Google engineering teams initiated machine shutdown procedures for the remainder of reachable devices as part of the cooling emergency shutdown process. A timeline of events during the incident is provided below. ### DataCenter Events | Timestamp(PST) | Event Description | | :---- | :---- | | 07-15-2026 16:24 | An electrical fault occurred on the utility grid upstream of the data center, disrupting the electrical distribution. | | 07-15-2026 16:24 | DRUPS back-up power system for side A failed to take over the data hall load. | | 07-15-2026 16:24 | The chiller controller dropped offline during the voltage transient event causing distribution pumps to stop and unable to restart. | | 07-15-2026 16:37 | Rows 1 and 2 transferred to power feed B successfully. Row 3 failed to transfer to power feed B and lost power. | | 07-15-2026 18:29 | Data hall temperatures reached 44C and crossed the safe machine operating threshold. | | 07-15-2026 19:55 | Machine shutdown procedures implemented. | | 07-15-2026 21:05 | Cooling system fully recovered and temperatures in the data hall returned to normal operating range. | | 07-15-2026 21:24 | Notification about cooling recovery sent by provider | | 07-15-2026 21:46 | Notification about power recovery sent by provider | ### GCVE Timelines | Timestamp(PST) | Event Description | | :---- | :---- | | 07-15-2026 16:30 | GCVE service power monitoring alert received | | 07-15-2026 16:41 | First symptom detected, servers reporting redundancy power feed failure. | | 07-15-2026 17:05 | Switch temperature \>60°C alerts triggered. | | 07-15-2026 17:09 | Start of user impact, first prober alert failure for north south traffic for customers. | | 07-15-2026 19:00 | Server / Network devices shutdown initiated. | | 07-15-2026 19:55 | All reachable server/network devices were shut down. | | 07-15-2026 21:24 | Cooling system fully recovered and temperatures in the datahall were returning to normal operating range. | | 07-15-2026 21:56 | Network restoration started with reachable hydra devices. | | 07-15-2026 22:40 | Onsite technician arrived, to recover the console servers. | | 07-15-2026 23:10 | Console server reboot/recovery complete. | | 07-16-2026 00:33 | Network recovery for all Placement Groups complete. | | 07-16-2026 02:36 | Incident mitigated, after recovering all the customer PCs are fully healthy. | This outage impacted 24 private clouds of 20 distinct customers in europe-west4-a region. ### Bare Metal Solutions (BMS) Timelines | Timestamp(PST) | Event Description | | :---- | :---- | | 07-15-2026 16:41 | First symptom detected, servers and storage reporting redundancy power feed failure | | 07-15-2026 17:52 | Automated monitoring detected rising temperatures | | 07-15-2026 18:37 | Four Netapp SAN storage nodes failed and were shut down due to overheating, resulting in storage availability issue | | 07-15-2026 18:37 | Customer impact started, first server shutdown detected | | 07-15-2026 18:59 | Reserve and buffer servers started to get powered off | | 07-15-2026 20:43 | Drop in the temperature has been detected | | 07-15-2026 20:52 | Four Netapp SAN storage nodes recovered, mitigating storage availability issue | | 07-15-2026 21:05 | The cooling system fully recovered and temperatures in the datahall were returning to normal operating range. | | 07-15-2026 23:23 | First customer server reboot started | | 07-16-2026 06:32 | Last customer server reboot started | | 07-16-2026 08:42 | Incident mitigated. All impacted customer servers were rebooted and confirmed in healthy state. | This outage impacted 9 distinct BMS customers in europe-west4 ### NetApp Timelines | Timestamp(PST) | Event Description | | :---- | :---- | | 07-15-2026 16:32 | Netapp facility remote monitoring alert on high temp and network switches switches failure received | | 07-15-2026 17:32 | All 6 clusters failed and were shut down automatically due to high temperature | | 07-15-2026 21:05 | Cooling capacity restored | | 07-15-2026 22:08 | Cluster recovery process initiated for all 6 clusters. 1st cluster recovered | | 07-15-2026 23:55 | All 6 Netapp Clusters and all customers recovered and validated as healthy & operational | ## \#\# Remediation and Prevention Google engineers were alerted to the infrastructure power loss on Wednesday, 15 July 2026 at 16:30 US/Pacific and subsequent increase in data hall temperatures at 17:44 US/Pacific via our automated monitoring and telemetry, and immediately began investigating. During the voltage transient event, the chillers maintained power but shut down because the distribution pumps failed to restart, resulting in no water flow. The pumps were manually switched from auto to hand mode to restore circulation and cooling. The facility team deployed an interim portable UPS to support the chiller controllers and prevent localized controller power loss. Upon confirmation of utility power restoration, the facility returned to utility power via automatic switching with the exception of two Remote Power Panels (RPPs), which required manual verification before restoration of power to affected server row. To provide full resiliency the facility team aligned a swing DRUPS unit to restore the redundant configuration on the A feed and is working on repairing the faulty DRUPS. To remediate the power loss at Row 3, once the electrical fault was confirmed to be fully isolated, teams reset the tripped breakers, successfully restoring both redundant power to the impacted server racks. Google is committed to preventing a repeat of this issue in the future and is completing the following actions: * Detailed investigation by the engineering team of the sequence of events that prevented transfer of load to DRUPS system and identified system improvements **\[ETA Aug 2026\].** * Determine and resolve the cause of overload conditions that caused power loss to Row 3 **\[ETA Aug 2026\].** * Perform investigation of chiller pump control system redundancy setup to ensure cooling system resiliency **\[ETA Sep 2026\].** * Implement improvements to facility system monitoring alerts for power events and configure early alert thresholds for cooling excursion detection. **\[ETA Aug 2026\].** * The GCVE engineering team is working to improve the automated triggering of shutdown during quickly evolving thermal runoff situations **\[ETA Oct 2026\]**. * The BMS engineering team is working with partner teams to update the incident classification and SLO definitions to improve personnel availability during major datacenter events **\[ETA Sep 2026\]**. * The GCNV engineering team is collaborate with the data center team to review incident handling SLA and runbooks, identify potential future occurrences and create playbooks to address/mitigate them, establish monthly joint emergency drill, and review if the Netapp cluster recovery process can be expedited **\[ETA Aug 2026\].** ## \#\# Detailed Description of Impact On 15 July 2026 from 16:39 to 16 July 2026 07:34 US/Pacific, customers experienced service disruptions across several products in europe-west4-a zone: * **Google Cloud VMware Engine**: Customers lost connectivity to their private clouds and workloads as foundational hosts and network switches were powered off for thermal protection. * **Bare Metal Solution**: Customers experienced connectivity loss to their database servers and storage appliances as the underlying hardware was powered down. * **Google Cloud NetApp Volumes**: Customers in the STANDARD, PREMIUM, and EXTREME service levels were unable to access their storage volumes. Control plane operations, including the creation of new storage pools, volumes, and backups, experienced failures in the affected region.

### 4. digitalocean — Serverless Inference - Gemma4 Latency Issues Causing Timeouts & Slow Responses

- URL: https://stspg.io/rzs04wclgpfq
- Started: 2026-07-13 19:04 UTC
- Resolved: 2026-07-15 21:46 UTC
- Duration: 2d 2h 42m
- Impact: minor, status: resolved
- Components: Inference
- Canonical incident: [ci-6dcd181a465685cc632e7a19](#ci-6dcd181a465685cc632e7a19)

- `2026-07-13 19:04` **identified** — We are currently experiencing an issue where customers using the Gemma 4 model on our Serverless Inference and Dedicated Inference platforms may experience severe latency or request timeouts. Our engineering team has identified a backend configuration issue as the root cause, which is temporarily degrading performance. We are actively working on a fix to restore normal response times and will provide another update as soon as the mitigation is in place
- `2026-07-13 22:24` **monitoring** — A fix has been deployed to resolve the backend configuration issue. We are closely monitoring system performance to ensure full recovery and normal response times for all Gemma 4 inference workflows.
- `2026-07-14 11:07` **identified** — We are currently experiencing an issue affecting customers using the Gemma 4 model on our Serverless Inference platform. Customers may experience significantly increased latency or request timeouts. Our Engineering team has identified a backend configuration issue as the root cause, which is temporarily impacting model performance. Please be assured that our Engineering team is actively working on a fix and is treating this issue with high priority. We sincerely apologise for any inconvenience this may have caused and appreciate your patience and understanding. If you have any further questions, please create a support ticket so that we can investigate your specific case further.
- `2026-07-15 10:06` **monitoring** — A fix has been deployed to resolve the issue. We are closely monitoring system performance to ensure full recovery and normal response times for all Gemma 4 inference workflows.
- `2026-07-15 21:46` **resolved** — The deployed fix has successfully restored full functionality, and our monitoring shows that system performance has completely stabilized. Response times for all Gemma 4 inference workflows have returned to normal baseline levels. We will continue to track platform stability moving forward to ensure long-term reliability. We apologize for any disruption this may have caused to your workflows and appreciate your patience throughout the recovery process.

### 5. anthropic — Elevated errors for Claude Opus 4.7

- URL: https://stspg.io/8fyh12n7xnnh
- Started: 2026-07-16 08:58 UTC
- Resolved: 2026-07-16 13:30 UTC
- Duration: 4h 32m
- Impact: minor, status: resolved
- Components: Claude API (api.anthropic.com), Claude Code, Claude Console (platform.claude.com), Claude Cowork, claude.ai
- Canonical incident: [ci-686212956fe2a481e1d2b0d3](#ci-686212956fe2a481e1d2b0d3)

- `2026-07-16 08:58` **investigating** — We are currently investigating this issue.
- `2026-07-16 13:30` **resolved** — This incident has been resolved.

### 6. anthropic — Elevated errors on Sonnet 5 and Haiku 4.5

- URL: https://stspg.io/bfnl00fvzpzp
- Started: 2026-07-17 06:47 UTC
- Resolved: 2026-07-17 12:21 UTC
- Duration: 5h 33m
- Impact: minor, status: resolved
- Components: Claude API (api.anthropic.com), Claude Code, Claude Cowork, claude.ai
- Canonical incident: [ci-f287bf87d3757080ff23133f](#ci-f287bf87d3757080ff23133f)

- `2026-07-17 06:47` **investigating** — We are currently investigating this issue.
- `2026-07-17 07:03` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-07-17 07:10` **identified** — The issue has been identified and a fix is being implemented.
- `2026-07-17 12:21` **resolved** — This incident has been resolved.

### 7. github — Incident with GitHub Actions

- URL: https://stspg.io/w8d77c7t94zf
- Started: 2026-07-19 23:34 UTC
- Resolved: 2026-07-20 04:44 UTC
- Duration: 5h 9m
- Impact: critical, status: resolved
- Components: API Requests, Actions, Issues, Pages
- Canonical incident: [ci-bd1ab603180908b002793ec5](#ci-bd1ab603180908b002793ec5)

- `2026-07-19 23:34` **investigating** — We are investigating reports of degraded performance for Actions
- `2026-07-19 23:37` **investigating** — We are investigating degraded availability for GitHub Actions on github.com and in GHEC DR stamps. New workflows may delay or fail to start, and ongoing runs may fail. We will provide more information as soon as we can.
- `2026-07-19 23:57` **investigating** — We have identified the cause of failures in GitHub Actions and are working to restore service.
- `2026-07-20 00:07` **investigating** — API Requests is experiencing degraded availability. We are continuing to investigate.
- `2026-07-20 00:20` **investigating** — Actions is experiencing degraded availability. We are continuing to investigate.
- `2026-07-20 00:48` **investigating** — Actions is experiencing degraded performance. We are continuing to investigate.
- `2026-07-20 00:49` **investigating** — Actions and Pages are experiencing degraded performance. We are continuing to investigate.
- `2026-07-20 00:52` **investigating** — Issues is experiencing degraded performance. We are continuing to investigate.
- `2026-07-20 01:11` **investigating** — We continue to work on mitigative efforts to restore Actions workflow runners, and have observed that the extended downtime has started to cause knock-on effects to other services.<br /><br />A separate incident was opened before we understood they were related. We will continue to post updates on this incident.
- `2026-07-20 01:37` **investigating** — Pages is operating normally.
- `2026-07-20 02:20` **investigating** — Issues is operating normally.
- `2026-07-20 02:23` **investigating** — API Requests is experiencing degraded performance. We are continuing to investigate.
- `2026-07-20 02:43` **investigating** — Issues, API Requests, and Pages have recovered. We are continuing to work on restoring GitHub Actions jobs using self-hosted or larger-hosted runners.
- `2026-07-20 03:03` **investigating** — API Requests is operating normally.
- `2026-07-20 03:27` **investigating** — We’re seeing recovery across all impacted Actions runners. The team is continuing to monitor for global recovery.
- `2026-07-20 03:34` **monitoring** — The degradation affecting Actions has been mitigated. We are monitoring to ensure stability.
- `2026-07-20 04:43` **monitoring** — Actions has fully recovered, and we will continue to monitor the platform to ensure stability.
- `2026-07-20 04:44` **resolved** — Between July 19, 2026, at 23:05 UTC and July 20, 2026, at 03:55 UTC, Actions self-hosted and larger runners were unable to connect to GitHub. During this period, Actions jobs were delayed or failed when trying to acquire a runner. Jobs using standard and Mac hosted runners were not affected. Reconnection traffic from affected runners also increased load on GitHub APIs, resulting in 3-4 seconds of additional average request latency and elevated 5xx error rates. <br /><br />The incident was caused by a certificate lifecycle management failure in a subset of internal services, resulting in an SSL certificate expiration that disrupted runner connectivity. We restored service by rotating the affected certificate. Recovery began at 02:45 UTC. By 03:55 UTC, queued workflow backlog had been processed and workflow delay rates returned to normal.<br /><br />To prevent recurrence, we are strengthening certificate renewal automation, adding fallback expiry monitoring and alerting, and improving circuit-breaker protections during runner API disruptions to reduce the risk of cascading impact to other APIs.

### 8. circleci — Issues with network routing for Mac jobs

- URL: https://stspg.io/vg35mbl8lxpg
- Started: 2026-07-13 21:52 UTC
- Resolved: 2026-07-14 23:18 UTC
- Duration: 1d 1h 26m
- Impact: minor, status: postmortem
- Components: macOS Jobs
- Canonical incident: [ci-d9d5a770feb641c6d6579b62](#ci-d9d5a770feb641c6d6579b62)

- `2026-07-13 21:52` **investigating** — We're received some reports of intermittent network issues between our mac infrastructure and a VCS provider again. We are investigating and will provide updates.
- `2026-07-13 22:21` **investigating** — We're still investigating these intermittent network issues. We'll update again shortly, we appreciate your patience.
- `2026-07-13 23:13` **investigating** — We've been able to reproduce the issue, but it still appears to be very much intermittent. We're continuing to work with our Mac infrastructure provider to help diagnose the issue. We'll update again soon.
- `2026-07-13 23:35` **monitoring** — We've identified the issue. Some customers may experience jobs that hang or fail to progress during Github fetch steps due to intermittent networking issues. Our engineers are actively monitoring. In the meantime, we've seen success with increasing the no_output_timeout to 15 - 20 min to allow more time for Github fetches. This can be done by following this community guide: https://support.circleci.com/hc/en-us/articles/360007188574-Build-has-Hit-Timeout-Limit
- `2026-07-14 13:59` **identified** — Unfortunately there are still some customers that may be experiencing jobs that hang or fail to progress during Github fetch steps due to ongoing intermittent networking issues with our Mac infrastructure provider. Our engineers are actively working with them to resolve this. We appreciate your patience. In the meantime, we've seen success with increasing the no_output_timeout to 15 - 20 min to allow more time for Github fetches. This can be done by following this community guide: https://support.circleci.com/hc/en-us/articles/360007188574-Build-has-Hit-Timeout-Limit
- `2026-07-14 15:35` **identified** — Our Mac infrastructure provider is testing changes to their network configuration in an effort to isolate the source of the issue. We're actively validating whether this resolves the impact and will follow up with another update soon. In the meantime, increasing the no_output_timeout to 15-20 min continues to help: https://support.circleci.com/hc/en-us/articles/360007188574-Build-has-Hit-Timeout-Limit
- `2026-07-14 16:31` **identified** — We've identified the issue as some inbound routes from the affected VCS provider causing significantly lower available bandwidth which is negatively impacting repo fetch times. We're working with our Mac infrastructure provider to attempt to shift inbound traffic to a different route, however this process may take an extended period of time to complete. We will continue to provide updates as they become available. As we've shared in previous updates, increasing the no_output_timeout to 15-20 min continues to help: https://support.circleci.com/hc/en-us/articles/360007188574-Build-has-Hit-Timeout-Limit
- `2026-07-14 16:58` **monitoring** — Our Mac infrastructure provider has shifted the inbound routes and we are seeing recovery in VCS fetch times. We will be continuing to monitor this issue over the next few hours. We'll update again soon.
- `2026-07-14 18:28` **monitoring** — We're continuing to see slow, progressive recovery in VCS fetch times. We're going to continuing monitoring and update again soon. We appreciate your patience.
- `2026-07-14 20:18` **investigating** — We're starting to see VCS fetch times rising again. Our engineering team is actively investigating and working with our Mac infrastructure provider to resolve this. Again, we appreciate your patience.
- `2026-07-14 22:15` **monitoring** — Our Mac infrastructure provider has shifted more traffic to working network providers and we're starting to see significant recovery again in VCS fetch times. As always, our engineering team will continue to monitor this closely and we'll provide another update soon.
- `2026-07-14 23:18` **resolved** — We continue to see significant VCS fetch time recovery over the last couple of hours. We're going to resolve this incident, and our engineering team will continue to monitor closely. We greatly appreciate your patience as we worked through this with our Mac infrastructure provider. Please reach out to our Support team if you have any issues.
- `2026-07-16 01:56` **postmortem** — ## Summary From 18:00 UTC on July 13, 2026 to 23:18 UTC on July 14, 2026, some customers running jobs on CircleCI's macOS fleet experienced jobs that hung or failed to progress during GitHub fetch steps. For these customers, GitHub fetch steps which would ordinarily take 1-2 minutes took 20-30 minutes, and timed out. The issue was caused by a capacity and routing problem on the network path between our Mac infrastructure provider and GitHub, and the associated intermittent networking issues slowed down fetches from GitHub. During the entire incident, our team worked directly with our Mac infrastructure provider to identify and reroute traffic away from the affected network paths. We resolved an initial occurrence at 19:40 UTC on July 13. The issue recurred at 21:52 UTC that same day, we reopened the incident, and it was fully resolved by 23:18 UTC on July 14. Our macOS fleet had a similar issue several weeks ago. From 20:18 UTC on June 24, 2026 to 03:09 UTC on June 25, 2026, customers experienced a [similar incident](https://status.circleci.com/incidents/gvysjmkf4ct9) affecting the ability of macOS jobs to reach GitHub. That incident was also caused by a network routing problem in our Mac infrastructure provider. We thank our customers for their patience while we worked through this incident. Please see below for specific actions which CircleCI and our Mac infrastructure provider will be taking. Please reach out to our support team with any questions or concerns. The status pages for this incident can be found [here](https://status.circleci.com/incidents/n1tc2lw9q7l0) and [here](https://status.circleci.com/incidents/7cnl777wp9qz). ## Background CircleCI's macOS jobs are hosted on a third-party Mac infrastructure provider. That provider connects to the broader internet, including services like GitHub, over multiple redundant upstream network paths. When one of those paths is experiencing reduced capacity or a routing problem, jobs that are fetching code or dependencies from a destination along that network path can slow down or hang intermittently, even when CircleCI's own platform, the third-party infrastructure provider, and the destination service are all operating normally. ## What Happened _\(All times UTC\)_ At 18:00 on July 13, some customers running macOS jobs began experiencing intermittent failures and delays fetching code and dependencies from GitHub. We opened an investigation and alerted our customers via our status page at 18:55. By 19:04, our Mac infrastructure provider identified increased latency on one of its network paths and rerouted traffic around it. Fetch times recovered, and we moved the incident to `Monitoring` at 19:29, and to `Resolved` at 19:40. At 20:44 UTC, customers reported that the issue had returned. We reopened the incident and updated our status page to `Investigating` at 21:52. Over the following two hours, we worked closely with our infrastructure provider to gather diagnostic data in an effort to isolate the affected network path. In the meantime, we published guidance recommending that customers increase their `no_output_timeout` setting to 15–20 minutes, to allow more time for GitHub fetches to complete during the intermittent slowdowns, and moved the status page to `Monitoring` at 23:35. After publishing the workaround guidance at 23:35 UTC on July 13, we expected the underlying network issue to improve overnight. It did not. At 13:59 UTC on July 14, we confirmed that customers were still experiencing intermittent failures and moved the status page back to `Investigating`. Our engineering team reproduced the failure directly using our own testing tools, which helped confirm this was a general `git-fetch` issue rather than something specific to any particular build tool, container image, or software update. At 15:35, we updated the status page to share that our infrastructure provider was testing changes to its network configuration to isolate the source of the issue. At 16:31, we updated the status page again to share that we'd identified the cause as reduced inbound bandwidth on a specific route, and that we were working with our provider to shift traffic to a different route. At 16:36, our infrastructure provider identified the affected network paths and shifted traffic away from them, fetch times recovered, and at 16:58, we updated the status page to `Monitoring`. By early afternoon, some customers were again seeing slow fetch times, and we moved the status page back to `Investigating` at 20:18. Our infrastructure provider identified two additional network paths with degraded performance and shifted traffic away from them by 22:07. Fetch times stabilized over the following hour, and we marked the incident `Resolved` at 23:18. In total, customers running Mac jobs may have experienced intermittent job slowdowns or failures for portions of the window between 18:00 UTC on July 13 and 23:18 UTC on July 14. Our infrastructure provider has since reported clean test results across its network and believes the underlying capacity issue originated further upstream, closer to GitHub, rather than within its own network. We are continuing to monitor this closely. ## Future Prevention and Process Improvement We are taking the following steps to prevent a recurrence and improve our response time: **We are building additional automated monitoring for macOS fleet GitHub-fetch performance.** Our extensive automated monitoring did not catch this incident. We are now adding end-to-end monitoring for GitHub fetches in particular, in order to detect similar issues proactively. **We are working directly with our Mac infrastructure provider on faster, more proactive detection.** We are asking our provider to build monitoring that can detect a degraded network path and reroute around it automatically, rather than relying on CircleCI to identify and request a reroute during an active incident. **We are turning the diagnostic tooling we built during this incident into a standing capability.** This will let us detect and reproduce this class of network failure on demand going forward, rather than assembling test infrastructure during an active incident. **We are reviewing our incident response process for a sufficient confirmation window before marking a network-related incident Resolved.** This incident briefly recurred after an early, premature resolution; we are formalizing a minimum period of confirmed-clean monitoring before closing incidents of this type. We will continue to provide regular updates along the way as investigation and remediation continues. **We are actively moving from a single provider for Mac infrastructure** **to multiple providers.** This will allow CircleCI to route incoming customer workloads to the most available and best-performing provider. Customer experience is our top priority, and we commit to continually improving the reliability of our systems to match the trust that our customers place in us. Please reach out to our support team with any questions or concerns.

### 9. digitalocean — Block Storage Volume NYC1, NYC3, SGP1, SYD1 and BLR1

- URL: https://stspg.io/wv68wmnz55ln
- Started: 2026-07-19 09:51 UTC
- Resolved: 2026-07-19 15:32 UTC
- Duration: 5h 41m
- Impact: minor, status: resolved
- Components: BLR1, NYC1, NYC3, SGP1, SYD1
- Canonical incident: [ci-302c651a19c62b6728ed0538](#ci-302c651a19c62b6728ed0538)

- `2026-07-19 09:51` **investigating** — Our Engineering team is investigating an issue with attaching volume to the droplets in the NYC1, NYC3, SGP1, SYD1 and BLR1 regions. During this time you may experience an issue with attaching volume to the Droplet(s). We apologize for the inconvenience and will share an update once we have more information.
- `2026-07-19 13:34` **investigating** — We are continuing to investigate this issue.
- `2026-07-19 13:36` **investigating** — We are continuing to investigate this issue.
- `2026-07-19 14:52` **monitoring** — Our Engineering team has identified the root cause of the issue preventing volume attachments to Droplets in NYC1, NYC3, SGP1, SYD1, and BLR1 and has successfully implemented a fix. You should now be able to attach volumes to your Droplet(s) without encountering any further errors or issues. We are actively monitoring the situation to ensure the fix remains effective and will provide another update once the issue has been fully resolved.
- `2026-07-19 15:32` **resolved** — Our Engineering team has confirmed that the issue regarding volume attachment to Droplets in the NYC1, NYC3, SGP1, SYD1, and BLR1 regions has been fully resolved, and all services are now operating normally. If you continue to experience any issues, please contact our Support team by opening a ticket. We apologize for any inconvenience caused.

### 10. grafana — Partial Outage in prod-eu-west-2

- URL: https://stspg.io/r9ysx33l8xbw
- Started: 2026-07-15 22:24 UTC
- Resolved: 2026-07-15 23:36 UTC
- Duration: 1h 12m
- Impact: major, status: resolved
- Components: AWS Germany - prod-eu-west-2, AWS Germany - prod-eu-west-2, AWS Germany - prod-eu-west-2, AWS Germany - prod-eu-west-2, AWS Germany - prod-eu-west-2, AWS Germany - prod-eu-west-2, AWS Germany - prod-eu-west-2, AWS Germany - prod-eu-west-2, AWS Germany - prod-eu-west-2: API, AWS Germany - prod-eu-west-2: Alertmanager, AWS Germany - prod-eu-west-2: Alertmanager and Rules Configuration API, AWS Germany - prod-eu-west-2: Ingestion, AWS Germany - prod-eu-west-2: Ingestion, AWS Germany - prod-eu-west-2: Metrics Generator, AWS Germany - prod-eu-west-2: Public Probes, AWS Germany - prod-eu-west-2: Querying, AWS Germany - prod-eu-west-2: Querying, AWS Germany - prod-eu-west-2: Rule Evaluation
- Canonical incident: [ci-bc0450b96f0e92e21e6cb2f7](#ci-bc0450b96f0e92e21e6cb2f7)

- `2026-07-15 22:24` **investigating** — We are investigating an issue affecting Loki queries in prod-eu-west-2. We first observed this behavior at approximately 21:57 UTC. Affected users may see elevated query errors, timeouts, or intermittent failures when running Loki queries in this region. The issue appears to be improving, but it is not fully resolved yet. We are continuing to investigate and will share another update as soon as we have more information.
- `2026-07-15 22:35` **investigating** — We are investigating a broader issue affecting multiple Grafana Cloud products in prod-eu-west-2. This appears to be caused by a third-party provider issue rather than a Loki-specific problem. Impact is currently inconsistent: some components are affected while others continue to function normally. We have also seen some impact to Tempo write paths. We are continuing to investigate and will share another update as soon as we have more information.
- `2026-07-15 22:38` **investigating** — We are continuing to investigate this issue.
- `2026-07-15 22:53` **monitoring** — We’ve implemented a fix and are monitoring the results to confirm the issue is fully resolved. Services may start to recover during this time.
- `2026-07-15 23:36` **resolved** — This incident has been resolved. Thank you for your patience.

### 11. gcp — Google Cloud VMware Engine (GCVE) Stretched Cluster customers are experiencing zonal outages impacting network connectivity across multiple regions.

- URL: https://status.cloud.google.com/incidents/T8gmtofFSTGT5tbhyciF
- Started: 2026-07-14 17:00 UTC
- Resolved: 2026-07-15 03:40 UTC
- Duration: 10h 40m
- Impact: major, status: resolved
- Components: VMWare engine
- Canonical incident: [ci-0511caaf3d55692014d0d66a](#ci-0511caaf3d55692014d0d66a)

- `2026-07-14 20:24` **investigating** — **Summary** Google Cloud VMware Engine (GCVE) Stretched Cluster customers are experiencing zonal outages impacting network connectivity across multiple regions. **Description** We are experiencing an issue with Google Cloud VMware Engine (GCVE) beginning Tuesday, 2026-07-14 at 10:00 PDT. We are continuing to investigate and mitigate the issue. We believe this disruption is isolated to network connectivity issues for stretch clusters, while Storage and Compute services appear to be unaffected. GCVE VMs are running as expected but customers may experience connectivity issues to the VMs. We will provide another update by Tuesday, 2026-07-14 14:00 PDT with current details. **Customer Symptoms** Some GCVE customers may experience inter-site communication failures to their GCVE environments within the affected zones. **Workaround** None at this time.
- `2026-07-14 21:31` **investigating** — **Summary** Google Cloud VMware Engine (GCVE) Stretched Cluster customers are experiencing zonal outages impacting network connectivity across multiple regions. **Description** We are experiencing an inter-site communication issue with Google Cloud VMware Engine (GCVE) stretched cluster customers beginning Tuesday, 2026-07-14 at 10:00 PDT. We have identified the customers that are affected by this issue and we are working with them on mitigation. Our preliminary investigation indicates this is stemming from an underlying network connectivity issue affecting the infrastructure that links the zones within a stretch cluster. This disruption is causing synchronization issues between the affected zones. We believe Storage and Compute services remain unaffected, and VMs are running as expected, though connectivity to them may be degraded. We will provide another update by Tuesday, 2026-07-14 15:00 PDT with current details. **Customer Symptoms** Some GCVE customers using Stretched Cluster may experience inter-site communication failures to their GCVE environments within the affected zones. **Workaround** While we work on restoring full connectivity, we recommend the following workarounds to restore access to your workloads: - VM Migration (recommended): Where possible, migrate your affected VMs to the healthy and unaffected side of the stretch cluster. We strongly recommend consulting with Google Support before proceeding.
- `2026-07-14 22:16` **investigating** — **Summary** Google Cloud VMware Engine (GCVE) Stretched Cluster customers are experiencing zonal outages impacting network connectivity across multiple regions. **Description** We are experiencing an inter-site communication issue with Google Cloud VMware Engine (GCVE) stretched cluster customers beginning Tuesday, 2026-07-14 at 10:00 PDT. We have identified the customers that are affected by this issue and we are working with them on mitigation. Our investigation has identified underlying inter-zone communication failures and Border Gateway Protocol (BGP) session flapping between cluster zones. Specifically, network connectivity has been lost between the affected zones and the witness appliance. Because the witness appliance is currently unreachable, the cluster zones are unable to safely synchronize state. As a result, VMs on the affected sites are becoming isolated and may be left without writable data. **Failover / VM Migration:** Migrating affected VMs to the healthy, secondary zone of your stretch cluster remains the primary mitigation strategy. Because of the complexities surrounding failover risks and secondary zone health, we highly encourage you to open a ticket with Google Cloud Support if you are severely impacted. We do not currently have an ETA for resolution. We will provide another update by Tuesday, 2026-07-14 16:00 PDT with current details. **Customer Symptoms** Some GCVE customers using Stretched Cluster may experience inter-site communication failures to their GCVE environments within the affected zones. **Workaround** While we work on restoring full connectivity, we recommend the following workarounds to restore access to your workloads: - VM Migration (recommended): Where possible, migrate your affected VMs to the healthy and unaffected side of the stretch cluster. We strongly recommend consulting with Google Support before proceeding.
- `2026-07-14 23:05` **investigating** — **Summary** Google Cloud VMware Engine (GCVE) Stretched Cluster customers are experiencing zonal outages impacting network connectivity across multiple regions. **Description** We are experiencing an inter-site communication issue with Google Cloud VMware Engine (GCVE) stretched cluster customers beginning Tuesday, 2026-07-14 at 10:00 PDT. We have identified the customers that are affected by this issue and we are working with them on mitigation. Our investigation has identified a recent configuration update that is the likely cause of the inter-zone network disruption. Teams are working on remediation. **Failover / VM Migration:** Migrating affected VMs to the healthy, secondary zone of your stretch cluster remains the primary mitigation strategy. Because of the complexities surrounding failover risks and secondary zone health, we highly encourage you to open a ticket with Google Cloud Support if you are severely impacted. We do not currently have an ETA for resolution. We will provide another update by Tuesday, 2026-07-14 17:00 PDT with current details. **Customer Symptoms** Some GCVE customers using Stretched Cluster may experience inter-site communication failures to their GCVE environments within the affected zones. **Workaround** While we work on restoring full connectivity, we recommend the following workarounds to restore access to your workloads: - VM Migration (recommended): Where possible, migrate your affected VMs to the healthy and unaffected side of the stretch cluster. We strongly recommend consulting with Google Support before proceeding.
- `2026-07-15 00:13` **investigating** — **Summary** Google Cloud VMware Engine (GCVE) Stretched Cluster customers are experiencing zonal outages impacting network connectivity across multiple regions. **Description** We are experiencing an inter-site communication issue with Google Cloud VMware Engine (GCVE) stretched cluster customers beginning Tuesday, 2026-07-14 at 10:00 PDT. We have identified the customers that are affected by this issue and we are working with them on mitigation. A remediation rollout is currently in progress to address the underlying network issue. **Failover / VM Migration:** Migrating affected VMs to the healthy, secondary zone of your stretch cluster remains the primary mitigation strategy. Because of the complexities surrounding failover risks and secondary zone health, we highly encourage you to open a ticket with Google Cloud Support if you are severely impacted. We do not currently have an ETA for resolution. We will provide another update by Tuesday, 2026-07-14 18:00 PDT with current details. **Customer Symptoms** Some GCVE customers using Stretched Cluster may experience inter-site communication failures to their GCVE environments within the affected zones. Upon further investigation we identified that the northamerica-northeast2 region was not impacted. **Workaround** While we work on restoring full connectivity, we recommend the following workarounds to restore access to your workloads: - VM Migration (recommended): Where possible, migrate your affected VMs to the healthy and unaffected side of the stretch cluster. We strongly recommend consulting with Google Support before proceeding.
- `2026-07-15 01:15` **investigating** — **Summary** Google Cloud VMware Engine (GCVE) Stretched Cluster customers are experiencing zonal outages impacting network connectivity across multiple regions. **Description** We are experiencing an inter-site communication issue with Google Cloud VMware Engine (GCVE) stretched cluster customers beginning Tuesday, 2026-07-14 at 10:00 PDT. We have identified the customers that are affected by this issue and we are working with them on mitigation. A remediation rollout is currently in progress to address the underlying network issue. **Failover / VM Migration:** Migrating affected VMs to the healthy, secondary zone of your stretch cluster remains the primary mitigation strategy. Because of the complexities surrounding failover risks and secondary zone health, we highly encourage you to open a ticket with Google Cloud Support if you are severely impacted. We do not currently have an ETA for resolution. We will provide another update by Tuesday, 2026-07-14 19:00 PDT with current details. **Customer Symptoms** Some GCVE customers using Stretched Cluster may experience inter-site communication failures to their GCVE environments within the affected zones. **Workaround** While we work on restoring full connectivity, we recommend the following workarounds to restore access to your workloads: - VM Migration (recommended): Where possible, migrate your affected VMs to the healthy and unaffected side of the stretch cluster. We strongly recommend consulting with Google Support before proceeding.
- `2026-07-15 01:55` **investigating** — **Summary** Google Cloud VMware Engine (GCVE) Stretched Cluster customers are experiencing zonal outages impacting network connectivity across multiple regions. **Description** We are experiencing an inter-site communication issue with Google Cloud VMware Engine (GCVE) stretched cluster customers beginning Tuesday, 2026-07-14 at 10:00 PDT. We have identified the customers that are affected by this issue and we are working with them on mitigation. A remediation rollout is currently in progress to address the underlying network issue. **Failover / VM Migration:** Migrating affected VMs to the healthy, secondary zone of your stretch cluster remains the primary mitigation strategy. Because of the complexities surrounding failover risks and secondary zone health, we highly encourage you to open a ticket with Google Cloud Support if you are severely impacted. We do not currently have an ETA for resolution. We will provide another update by Tuesday, 2026-07-14 20:30 PDT with current details. **Customer Symptoms** Some GCVE customers using Stretched Cluster may experience inter-site communication failures to their GCVE environments within the affected zones. **Workaround** While we work on restoring full connectivity, we recommend the following workarounds to restore access to your workloads: - VM Migration (recommended): Where possible, migrate your affected VMs to the healthy and unaffected side of the stretch cluster. We strongly recommend consulting with Google Support before proceeding.
- `2026-07-15 03:27` **investigating** — **Summary** Google Cloud VMware Engine (GCVE) Stretched Cluster customers are experiencing zonal outages impacting network connectivity across multiple regions. **Description** We are experiencing an inter-site communication issue with Google Cloud VMware Engine (GCVE) stretched cluster customers beginning Tuesday, 2026-07-14 at 10:00 PDT. We have identified the customers that are affected by this issue and we are working with them on mitigation. A remediation rollout is currently in progress to address the underlying network issue. **Failover / VM Migration:** Migrating affected VMs to the healthy, secondary zone of your stretch cluster remains the primary mitigation strategy. Because of the complexities surrounding failover risks and secondary zone health, we highly encourage you to open a ticket with Google Cloud Support if you are severely impacted. We do not currently have an ETA for resolution. We will provide another update by Tuesday, 2026-07-14 22:00 PDT with current details. **Customer Symptoms** Some GCVE customers using Stretched Cluster may experience inter-site communication failures to their GCVE environments within the affected zones. **Workaround** While we work on restoring full connectivity, we recommend the following to restore access to your workloads: * VM Migration (recommended): Where possible, migrate your affected VMs to the healthy and unaffected side of the stretch cluster. We strongly recommend consulting with Google Cloud Support before proceeding.
- `2026-07-15 05:34` **resolved** — **Summary** Google Cloud VMware Engine (GCVE) Stretched Cluster customers may have experienced zonal outages impacting network connectivity across multiple regions. **Description** We experienced an inter-site communication issue with Google Cloud VMware Engine (GCVE) Stretched Cluster customers beginning Tuesday, 2026-07-14 at 10:00 PDT. We identified the affected customers, and have worked with them to fully mitigate the issue by Tuesday, 2026-07-14 at 21:46 PDT. Preliminary analysis indicates that a network configuration update was the cause of the inter-zone network disruption. Our engineering team mitigated the issue by rolling back the faulty configuration to its last-known good value. If customers are still experiencing impact from this issue, please contact Support, and we will work with you to resolve any residual impact. We thank you for your patience while we worked to resolve this issue. **Customer Symptoms** Some GCVE customers using Stretched Cluster may have experienced inter-site communication failures to their GCVE environments within the affected zones. **Workaround** The issue is now mitigated.
- `2026-07-20 18:31` **resolved** — ## Preliminary Incident Report We sincerely apologize for the disruption this incident caused to your business. We know how much you rely on Google Cloud, and we regret the impact on your productivity. We are working to address the root cause and prevent this from occurring in the future. Please note, this information is based on our best knowledge at the time of posting and is subject to change as our investigation continues. A final Incident Report with preventative actions will be posted once our investigation is complete. If you have experienced impact outside of what is listed below, please reach out to Google Cloud Support using [**https://cloud.google.com/support**](https://cloud.google.com/support). ## Date/Time of the Issue (All time US/Pacific) Incident Start: 14 July 2026 10:00 Incident End: 14 July 2026 20:40 Duration: 10 hours, 40 minutes ## Summary On Tuesday, 14 July 2026, Google Cloud VMware Engine (GCVE) Stretched Cluster customers in the australia-southeast2 and europe-west3 zones experienced inter-site communication failures. The disruption was traced to a network configuration update that introduced a conflict, causing inter-site communication failures, which triggered VMware Stretch Cluster failover events. Google engineers successfully mitigated the disruption by rolling back the configuration change. ## Preliminary Root Cause A configuration change deployed within the Google Cloud network affected traffic between the VMware Engine stretched cluster zones. This resulted in a loss of connectivity between zones in the stretched cluster deployment. Google engineers are doing a full root cause analysis and will provide additional information once it is available. ## Remediation To stabilize the environment, engineers identified the working network paths and deployed configuration changes to reroute traffic and restore connectivity. GCVE stretch cluster inter-zonal connectivity was completely restored for all supported locations on Tuesday, 14 July 2026 at 20:40 US/Pacific. ## Description of Impact On Tuesday, 14 July 2026 from 10:00 to 20:40 US/Pacific, customers utilizing GCVE stretch clusters in the affected zones (australia-southeast2, and europe-west3) experienced inter-site communication failures. For some customers, depending on their architecture, this disruption led to VMWare HA event causing VM restarts/movement across zones as designed, Host disconnects, VSAN alarms and Intermittent access to VMware Management (vCenter/NSX Manager) components. ---
- `2026-07-24 07:10` **resolved** — ## Incident Report ## Summary On Tuesday, 14 July 2026 10:00 PT, Google Cloud VMware Engine (GCVE) Stretched Cluster customers in the australia-southeast2 and europe-west3 zones experienced inter-site communication failures. The disruption was traced to a network configuration update that introduced a conflict, causing inter-site communication failures, which triggered VMware Stretched Cluster failover events. Google engineers successfully mitigated the impact by rolling back the issue-causing configuration change and resetting the network state. We sincerely apologize for the disruption this incident caused to your business. We know how much you rely on Google Cloud, and we regret the impact on your productivity. We are working to address the root cause and prevent this from occurring in the future. ## Root Cause A network configuration update intended to prepare the cloud network infrastructure for new capabilities was deployed to the foundational network control plane. While the configuration payload itself was structurally valid, it exposed an implementation gap within the control plane's routing logic. This logical gap caused the underlying network hosts to misconfigure program routing tables. As a result, traffic destined for the private IP address space used by GCVE Stretched Clusters was dropped. Because Stretched Clusters rely on this private address space to establish routing sessions for inter-zonal connectivity, the traffic drops severed communications between the active zones of the clusters, triggering VMware High Availability (HA) failovers. Standard routing health-checking protocols, such as Border Gateway Protocol (BGP) and Bidirectional Forwarding Detection (BFD), remained fully functional because their control plane sessions run on separate, unaffected address spaces. Because the control plane remained healthy, standard failover mechanisms failed to detect that the selective private IP range used for inter-zonal data tunneling was being dropped. Automated safeguards did not block the deployment because the configuration passed initial payload validations. The impact only manifested once the update began routing data traffic through the specific affected IP range. ## Remediation and Prevention To stabilize the environment, engineers identified the working network paths and deployed configuration changes to reroute traffic and restore connectivity. GCVE Stretched Cluster inter-zonal connectivity was completely restored for all supported locations on Tuesday, 14 July 2026 at 20:40 US/Pacific. Google is committed preventing a repeat of this issue in the future and is completing the following actions: * **Expanded Testing:** We are adding more detailed GCVE network setups to our existing testing environments. This allows us to automatically test future network updates against these configurations before they go live. * **Service-Level Data Path Failover:** We are implementing additional service-level data path failover mechanisms that actively probe the specific data-tunneling traffic space. This will ensure a path failover is triggered if the data plane itself is degraded even when the BGP control plane remains functional. * **Detailed Alerting**: We are adding faster, more specific alerts for connection issues between zones. This builds on our current platform monitoring to catch minor disruptions early and speed up our response. * **Improved Cluster Resilience:** We are fine-tuning the cluster's high-availability and storage settings. This makes virtual machines more resilient to short network drops, preventing them from restarting unnecessarily if the main site is still healthy. * **Workload Resilience Alignment (Shared Responsibility):** We are proactively reaching out to customers utilizing non-vSAN replicated virtual machine configurations within Stretched Clusters. Because these workloads are pinned to a single zone without active cross-site replication, they cannot survive inter-site network disruptions. We are ready to assist customers in auditing their storage policies, adjusting Stretched Cluster configurations, and planning the secondary zone capacity required to enable robust high-availability failovers. ## Detailed Description of Impact On Tuesday, 14 July 2026 from 10:00 to 20:40 US/Pacific, customers utilizing GCVE Stretched Clusters in the affected zones (australia-southeast2, and europe-west3) experienced inter-site communication failures. For some customers, depending on their architecture, this disruption led to a VMWare HA event causing VM restarts/movement across zones as designed, host disconnects, VSAN alarms and intermittent access to VMware Management (vCenter/NSX Manager) components.

### 12. grafana — Some Reports of Grafana Not Loading.

- URL: https://stspg.io/ddf067fl4z8h
- Started: 2026-07-15 18:07 UTC
- Resolved: 2026-07-15 21:00 UTC
- Duration: 2h 53m
- Impact: major, status: resolved
- Components: AWS US East - prod-us-east-0, GCP US Central - prod-us-central-0, GCP US Central - prod-us-central-3, GCP US Central - prod-us-central-4
- Canonical incident: [ci-c5d98cc59bbfe51ec56959a7](#ci-c5d98cc59bbfe51ec56959a7)

- `2026-07-15 18:07` **investigating** — We are currently investigating an issue impacting a small subset of stacks in the impacted regions. We will provide more details as they become available.
- `2026-07-15 18:47` **identified** — We believe we have found the cause and are working on remediation. It is also worth mentioning that only stacks on the "slow" release channel are impacted.
- `2026-07-15 20:50` **identified** — The fix is in the process of being rolled out.
- `2026-07-15 21:00` **resolved** — This incident has been resolved. Thank you for your patience.

### 13. grafana — Delayed Aggregated Metrics (prod-us-central-0)

- URL: https://stspg.io/56wgtyyjvnt7
- Started: 2026-07-15 19:19 UTC
- Resolved: 2026-07-16 03:54 UTC
- Duration: 8h 35m
- Impact: minor, status: resolved
- Components: GCP US Central - prod-us-central-0: Ingestion
- Canonical incident: [ci-c5e649e4bd1dc7c56258efc8](#ci-c5e649e4bd1dc7c56258efc8)

- `2026-07-15 19:19` **identified** — We are currently investigating a delay in aggregated metric results for tenants using Adaptive Metrics in the prod-us-central-0 region. Our engineering team has identified the issue and is actively working on a mitigation. We will provide further updates as the investigation progresses.
- `2026-07-15 21:23` **monitoring** — We have applied the mitigation and are monitoring as the backlog catches up. We will post another update once that is complete. Thank you for your patience.
- `2026-07-16 03:54` **resolved** — This incident has been resolved.

### 14. anthropic — Elevated errors across Fable 5

- URL: https://stspg.io/jv85qb4xwc4x
- Started: 2026-07-17 18:32 UTC
- Resolved: 2026-07-17 19:43 UTC
- Duration: 1h 11m
- Impact: major, status: resolved
- Components: Claude API (api.anthropic.com), Claude Code, Claude Cowork, claude.ai
- Canonical incident: [ci-a662c5334eb363b42393d526](#ci-a662c5334eb363b42393d526)

- `2026-07-17 18:32` **investigating** — We are currently investigating this issue.
- `2026-07-17 18:36` **investigating** — We are aware of an issue preventing users from selecting Claude Fable 5 within Claude.ai, Claude Code, and other surfaces, and are working to resolve this issue.
- `2026-07-17 18:48` **monitoring** — We have applied a fix to resolve the erroneous requirement for usage credits on Fable 5, and are monitoring to ensure no further issues. You may need to relaunch Claude.ai or Claude Code to restore access as expected.
- `2026-07-17 19:43` **resolved** — This issue has been resolved, and Claude Fable 5 is available once again without usage credits, as expected.

### 15. anthropic — Elevated errors for multiple models

- URL: https://stspg.io/bdw3y9kd205m
- Started: 2026-07-16 18:36 UTC
- Resolved: 2026-07-16 22:53 UTC
- Duration: 4h 16m
- Impact: major, status: resolved
- Components: —
- Canonical incident: [ci-102b10e3b36d7cac610ec126](#ci-102b10e3b36d7cac610ec126)

- `2026-07-16 18:36` **investigating** — We are investigating elevated errors on requests to multiple models. We will provide an update as soon as possible.
- `2026-07-16 19:14` **identified** — We have identified the cause of elevated errors on requests to multiple models, and are working to resolve this issue. We will provide an update as soon as possible.
- `2026-07-16 20:07` **identified** — We are continuing to work to resolve issues resulting in elevated errors on requests to some Anthropic models. We will provide additional updates as soon as possible.
- `2026-07-16 20:58` **identified** — At this time, we are seeing recovery on requests to most models. We are continuing to see elevated errors on requests to Claude Fable 5 and Mythos 5, and are working to resolve these remaining issues. We will provide an additional update as soon as possible.
- `2026-07-16 22:01` **monitoring** — We have seen success rates recover and are continuing to monitor closely to ensure no further issues. We will provide an additional update shortly.
- `2026-07-16 22:53` **resolved** — This issue has been resolved. From 11:30am PT / 18:30 UTC through to 3:15pm PT / 22:15 UTC some users experienced elevated errors on requests to Claude Models. Most errors were experienced within the first hour of the impact window.

### 16. digitalocean — Reserved IP routing in TOR1

- URL: https://stspg.io/x0s0pghvdcgv
- Started: 2026-07-16 13:39 UTC
- Resolved: 2026-07-16 17:24 UTC
- Duration: 3h 44m
- Impact: minor, status: resolved
- Components: Reserved IP, TOR1
- Canonical incident: [ci-a68b9f479904aabec11c9264](#ci-a68b9f479904aabec11c9264)

- `2026-07-16 13:39` **investigating** — Our Engineering team is investigating an issue affecting Reserved IP routing in the TOR1 region. The issue is causing inbound and outbound traffic to and from Reserved IPs in TOR1 to fail. In the meantime, detaching and reassigning the Reserved IP to the Droplet may restore connectivity. We recommend trying this workaround if you are experiencing issues with your Reserved IP in TOR1. Our engineering team is actively investigating the root cause and working towards a resolution. We will provide updates as more information becomes available.
- `2026-07-16 15:30` **identified** — Our Engineering team has identified the cause of the issue affecting Reserved IP routing in the TOR1 region, which was causing inbound and outbound traffic to and from Reserved IPs in TOR1 to fail. Detaching and reassigning the Reserved IP to the Droplet may restore connectivity in the meantime. We recommend trying this workaround if you are still experiencing issues with your Reserved IP in TOR1. Our engineering team is now working on implementing a fix. We will provide updates as more information becomes available.
- `2026-07-16 17:01` **monitoring** — Our Engineering team has implemented a fix for the issue affecting Reserved IP routing in the TOR1 region, and we are currently monitoring the results. Inbound and outbound traffic to and from Reserved IPs in TOR1 should now be restored. If you reassigned your Reserved IP as a workaround and are still experiencing issues, please reach out to our support team. We will continue to monitor the situation closely and provide a further update once we have confirmed the issue is fully resolved.
- `2026-07-16 17:24` **resolved** — Our Engineering team has confirmed that the underlying issue affecting Reserved IP routing in the TOR1 region has been fully resolved, and services are now operating normally. If you continue to experience any issues, please contact our Support team by opening a ticket. We apologize for any inconvenience caused.

### 17. yandexcloud — The problem is with the UI for the Managed Service for YDB

- URL: https://status.yandex.cloud/en/incidents/1883
- Started: 2026-07-13 12:40 UTC
- Resolved: 2026-07-13 16:10 UTC
- Duration: 3h 30m
- Impact: minor, status: resolved
- Components: Managed Service for YDB
- Canonical incident: [ci-4c390f37329ec200ad9039b7](#ci-4c390f37329ec200ad9039b7)

- `2026-07-13 15:41` **investigating** — We noticed a partial unavailability of the UI when working with the Managed Service for YDB. The developers are already aware of it and are fixing it.
- `2026-07-13 16:26` **resolved** — There are no more problems with the service.

### 18. github — Disruption with some GitHub services

- URL: https://stspg.io/5c43kg01skdw
- Started: 2026-07-16 09:13 UTC
- Resolved: 2026-07-16 12:20 UTC
- Duration: 3h 7m
- Impact: critical, status: resolved
- Components: —
- Canonical incident: [ci-5b251004294104d745470177](#ci-5b251004294104d745470177)

- `2026-07-16 09:13` **investigating** — We are investigating reports of impacted performance for some GitHub services.
- `2026-07-16 09:30` **monitoring** — The degradation has been mitigated. We are monitoring to ensure stability.
- `2026-07-16 10:17` **monitoring** — The degradation has been mitigated. We are monitoring to ensure stability.
- `2026-07-16 12:20` **resolved** — On July 16, 2026, between 08:50 UTC and 09:50 UTC, the GitHub MCP Server’s web_search tool experienced elevated failures. The average error rate was 42% and peaked at 82% of requests to the tool. Other GitHub MCP Server tools were unaffected. This was caused by degradation at a downstream web search provider.<br /><br />The incident was mitigated when the downstream provider recovered, after which we confirmed that the tool’s success rate had returned to normal.<br /><br />We are improving the tool’s resilience and failure handling to reduce the customer impact and duration of similar incidents.

### 19. grafana — Mimir Partial Write Outage

- URL: https://stspg.io/j81jwzycv215
- Started: 2026-07-14 17:09 UTC
- Resolved: 2026-07-14 19:03 UTC
- Duration: 1h 53m
- Impact: major, status: resolved
- Components: GCP US Central - prod-us-central-0: Ingestion, GCP US Central - prod-us-central-0: Querying
- Canonical incident: [ci-60a58a8f99f5d2d08be56375](#ci-60a58a8f99f5d2d08be56375)

- `2026-07-14 17:09` **identified** — We're investigating backend degradation which has resulted in a partial write outage beginning around 16:40 UTC. This degradation has also impacted reads and rule evaluation. Issue has been identified and we are working on mitigation.
- `2026-07-14 17:23` **monitoring** — The outage is recovered as of 16:56 UTC, and we're continuing to monitor.
- `2026-07-14 19:03` **resolved** — This incident has been resolved.

### 20. anthropic — Enterprise SSO sign-in failures

- URL: https://stspg.io/2kbn7kkr22fx
- Started: 2026-07-16 09:23 UTC
- Resolved: 2026-07-16 10:27 UTC
- Duration: 1h 3m
- Impact: major, status: resolved
- Components: Claude Console (platform.claude.com), claude.ai
- Canonical incident: [ci-b5532b5b313b5952aef8dc39](#ci-b5532b5b313b5952aef8dc39)

- `2026-07-16 09:23` **identified** — The issue has been identified and a fix is being implemented.
- `2026-07-16 10:27` **resolved** — This incident has been resolved.

### 21. circleci — Login issues for some Bitbucket users

- URL: https://stspg.io/lzl0fmm6xmts
- Started: 2026-07-14 02:49 UTC
- Resolved: 2026-07-14 05:15 UTC
- Duration: 2h 25m
- Impact: minor, status: postmortem
- Components: Pipelines & Workflows
- Canonical incident: [ci-5c281e7a150c707d7209b592](#ci-5c281e7a150c707d7209b592)

- `2026-07-14 02:49` **investigating** — We are currently investigating login errors affecting some customers using Bitbucket as their identity provider.
- `2026-07-14 03:29` **investigating** — Our team continues to investigate login errors affecting some customers using Bitbucket as their identity provider.
- `2026-07-14 03:51` **investigating** — What's impacted Some customers logging in with a Bitbucket-linked identity are affected. This may also include customers logging in with GitHub but have a linked Bitbucket identity on their account. Additionally, some customers are experiencing workflows that are not starting or updating. What can you expect Affected customers may see login errors. Some customers may also notice workflows that fail to start or do not update as expected. Thank you for your patience while our engineers investigate. Next update We will provide another update as soon as we have more information to share
- `2026-07-14 04:21` **monitoring** — What's impacted Some customers logging in with a Bitbucket-linked identity were affected. This may also have included customers logging in with GitHub who have a linked Bitbucket identity on their account. Some customers also experienced workflows that were not starting or updating. What can you expect A fix has been deployed. We are monitoring the situation to confirm full recovery. Next update We will provide an update if anything changes
- `2026-07-14 05:15` **resolved** — The issue where customers logging in with a Bitbucket-linked identity experienced login errors, and some customers experienced workflows that failed to start or did not update, has been resolved. What you may still experience and need to do - If you use Bitbucket as your login method, please log out and reauthenticate the CircleCI Bitbucket integration on your account. - If any of your workflows are stuck or showing missing status checks, please rerun them to pick up the fix. - If you continue to experience issues after taking these steps, please reach out to CircleCI Support. We thank you for your patience while our team worked on implementing a fix.
- `2026-07-14 22:31` **postmortem** — ## Summary From 01:33 UTC to 05:18 UTC on July 14, 2026, customers with a Bitbucket-linked identity were unable to log in to CircleCI, and some customers experienced workflows that failed to start or update, due to a change in how Bitbucket's OAuth service reports account permissions. At 05:07 UTC, we deployed a fix that corrected how our systems read the updated permissions field. Some customers needed to log out and re-authenticate their Bitbucket integration after we deployed the fix. Our systems continued processing the backlog of affected jobs until 08:39 UTC. We thank our customers for their patience as we resolved this incident. Please reach out to our support team with any questions or concerns. The status page for this incident can be found [here](https://status.circleci.com/incidents/gsyjwybg477g). ## Background CircleCI supports logging in with a GitHub account, a Bitbucket account, or an email and password. When a customer logs in with a Bitbucket-linked identity, or when CircleCI needs to refresh the customer's access on their behalf, we exchange an authorization token with Bitbucket's OAuth service. This exchange includes a list of the permissions, or "scopes," the customer has granted us, which Bitbucket and CircleCI use to confirm what CircleCI is authorized to do on the customer’s behalf. ## What Happened \(All times UTC\) On April 8, 2026, [Bitbucket announced a change to its OAuth service](https://developer.atlassian.com/cloud/bitbucket/changelog/#CHANGE-3139): it would be renaming the field used to report a customer's granted permissions. Bitbucket ran a transition period during which both the old and new field names were available, then fully removed the old field name on May 4, 2026. Our systems had not been updated to recognize the new field name, so once Bitbucket fully phased out the old one, requests that depended on it began to fail. At 01:33 on July 14, 2026, our systems began failing to process permissions information returned for Bitbucket-linked accounts, because our systems still expected the old permissions structure. This caused login attempts for all Bitbucket-linked accounts to fail, including for customers who log in with GitHub but have a Bitbucket identity linked to their account. At 01:59, automated monitoring alerted our engineering team to a spike in errors on the affected systems. The team began investigating immediately, confirmed customer impact at 02:38, and alerted customers via our status page at 02:51. By 03:00, the team had isolated the failures to the renamed Bitbucket field. Beginning at 03:17, workflow status updates for Bitbucket pipelines belonging to customers with an expired access token began to be dropped. The Bitbucket permissions check is used broadly across our platform, and these permissions failures also affected some of our internal job-processing systems. This caused a subset of workflows to become stuck without a final status, and caused some pull requests to show missing or stuck status checks. At 04:21, the team deployed an initial fix that resolved the underlying issue, restoring the login flow and returning our internal permissions system to normal operation. Some affected customers needed to log out and log back in to pick up the fix. By 05:07, the team deployed two additional fixes so our systems would begin accepting Bitbucket's new permissions field format going forward. We resolved the incident at 05:18. Some customers did need to log out and re-authenticate their Bitbucket integration before their account fully recovered. Our systems continued processing a backlog of affected jobs, returning to normal levels by approximately 08:39. ## Future Prevention and Process Improvement We are taking the following steps to prevent a recurrence and improve our response time: **We are hardening our authorization code against upstream API changes.** This incident happened because our system did not gracefully handle a renamed field in a response from a third-party identity provider. We will be immediately updating our authorization code so that unexpected or missing fields from GitHub, Bitbucket, and GitLab are handled safely. **We are improving how we track upstream provider changes.** We currently already monitor changelogs from our identity providers for exactly this kind of breaking change, but we didn't add this particular provider change to our monitoring in time to catch it before it shipped. We are auditing and expanding this monitoring so provider announcements reach our team before they affect customers. **We are improving how we classify and communicate incidents in their earliest minutes.** This incident's initial classification did not immediately reflect its customer-facing severity. We are refining our incident tooling and guidance to help engineers identify and communicate customer impact more quickly. **We are improving the resilience of our workflow-processing pipeline.** A downstream failure in this incident caused some workflow status updates to be dropped rather than retried or clearly surfaced. We are reviewing this system's retry and error-handling behavior so similar downstream failures are more visible and easier to recover from. Customer experience is our top priority, and we commit to continually improving the reliability of our systems to match the trust that our customers place in us. Please reach out to our support team with any questions or concerns.

### 22. digitalocean — Networking in BLR1

- URL: https://stspg.io/8yc90q76srsl
- Started: 2026-07-19 06:57 UTC
- Resolved: 2026-07-19 09:52 UTC
- Duration: 2h 54m
- Impact: minor, status: resolved
- Components: BLR1
- Canonical incident: [ci-68dec95a2a2dfd05de0d5ec1](#ci-68dec95a2a2dfd05de0d5ec1)

- `2026-07-19 06:57` **investigating** — We are currently investigating an issue affecting outbound network traffic for Droplets in the BLR1 region. Customers may experience timeout errors or connectivity issues when attempting to reach external services. Our Engineering team is actively investigating the issue and working to identify the root cause. We will provide additional updates as soon as more information becomes available.
- `2026-07-19 08:27` **identified** — Our Engineering team has identified an issue affecting upstream network connectivity for Droplet-based services in the BLR1 region and is actively working with the relevant upstream providers to implement a resolution. Customers may continue to experience intermittent timeouts or connectivity issues when attempting to reach external endpoints from Droplets in BLR1. We will provide another update once the issue has been fully resolved or when additional information becomes available.
- `2026-07-19 09:31` **monitoring** — Our engineering team has implemented the necessary fixes affecting upstream network connectivity for Droplet-based services in the BLR1 region. Users should now be able to connect to the external endpoints from Droplets in BLR1. We are currently monitoring the situation to ensure that the service has returned to normal operation and remain stable. We appreciate your patience and will provide an update once the issue is fully confirmed as resolved.
- `2026-07-19 09:52` **resolved** — Our engineering team has resolved the outbound network connectivity issue in the BLR1 region. If you continue to experience problems, please open a ticket with our support team. We apologize for any inconvenience.

### 23. github — Degraded REST API Availability

- URL: https://stspg.io/pzvjh0wyg401
- Started: 2026-07-16 22:51 UTC
- Resolved: 2026-07-17 00:14 UTC
- Duration: 1h 22m
- Impact: major, status: resolved
- Components: API Requests
- Canonical incident: [ci-5a0048c0d502177e0fe6f1ba](#ci-5a0048c0d502177e0fe6f1ba)

- `2026-07-16 22:51` **investigating** — We are investigating reports of impacted performance for some GitHub services.
- `2026-07-16 22:58` **investigating** — We are aware of degraded REST API availability and are investigating
- `2026-07-16 22:58` **investigating** — API Requests is experiencing degraded performance. We are continuing to investigate.
- `2026-07-16 23:29` **investigating** — We are continuing to investigate an issue causing approximately 35% of REST API requests to fail. Based on our current understanding, requests are not consistently reaching the application layer, resulting in failed requests returning HTML responses instead of the expected API response format. We are actively investigating the issue and will provide another update as soon as more information is available.
- `2026-07-17 00:00` **monitoring** — The degradation affecting API Requests has been mitigated. We are monitoring to ensure stability.
- `2026-07-17 00:14` **monitoring** — As of 23:46 UTC, the REST API service is reachable and responding to requests normally.
- `2026-07-17 00:14` **resolved** — From 22:21 UTC - 23:50 UTC on July 16, 2026, the REST API experienced significant degradation. During this period, about 39% of REST API requests failed with HTTP 500 level responses, with the errors peaking at 44.3%. <br /><br />We identified the issue as an infrastructure change that wrongly marked the majority of API backends in a single region as unhealthy. As a result, requests routed to those backends failed before reaching the application layer. <br /><br />To prevent this from happening again, we're improving our systems to catch this kind of invalid configuration before it reaches production. We'll also audit the related systems to make them more resilient to future changes, and we're increasing our monitoring sensitivity so we're alerted to problems like this sooner.

### 24. grafana — Adaptive Metrics aggregation delay in eu-west-0 region.

- URL: https://stspg.io/t147yktp9h8k
- Started: 2026-07-16 12:56 UTC
- Resolved: 2026-07-16 14:58 UTC
- Duration: 2h 2m
- Impact: minor, status: resolved
- Components: GCP Belgium - prod-eu-west-0: Ingestion
- Canonical incident: [ci-681523e1556a6b1adcae2396](#ci-681523e1556a6b1adcae2396)

- `2026-07-16 12:56` **identified** — We're currently facing an issue with Adaptive Metrics aggregation delay in the eu-west-0 region (GCP Belgium). The issue started at around 11:50 UTC, but we were able to identify the issue and the appropriate fix is already deployed. We're seeing services recovering. More updated to come soon.
- `2026-07-16 13:45` **monitoring** — Services are fully recovered now. We're monitoring to be sure the issue won't re-occur.
- `2026-07-16 14:58` **resolved** — The incident is now fully resolved.

### 25. anthropic — Elevated error rates for Sonnet 5

- URL: https://stspg.io/zfp1zk8gg7p3
- Started: 2026-07-17 18:12 UTC
- Resolved: 2026-07-17 18:27 UTC
- Duration: 14m
- Impact: major, status: resolved
- Components: Claude API (api.anthropic.com), Claude Code, Claude Console (platform.claude.com), Claude Cowork, claude.ai
- Canonical incident: [ci-a0a6ef7d7ad417dca2fbc7ce](#ci-a0a6ef7d7ad417dca2fbc7ce)

- `2026-07-17 18:12` **investigating** — We are currently investigating this issue.
- `2026-07-17 18:13` **investigating** — We are continuing to investigate this issue.
- `2026-07-17 18:22` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-07-17 18:27` **resolved** — This incident has been resolved.

### 26. github — Actions runs are experiencing failures to start

- URL: https://stspg.io/t0tfsmbt70w7
- Started: 2026-07-13 13:32 UTC
- Resolved: 2026-07-13 13:53 UTC
- Duration: 21m
- Impact: major, status: resolved
- Components: Actions, Pages
- Canonical incident: [ci-1f477a5e8571aadb8872587f](#ci-1f477a5e8571aadb8872587f)

- `2026-07-13 13:32` **investigating** — We are investigating reports of degraded availability for Actions
- `2026-07-13 13:32` **investigating** — Pages is experiencing degraded performance. We are continuing to investigate.
- `2026-07-13 13:39` **monitoring** — The degradation affecting Actions and Pages has been mitigated. We are monitoring to ensure stability.
- `2026-07-13 13:53` **resolved** — On July 13, 2026, between 13:11 and 13:53 UTC, some customers experienced failures starting and running GitHub Actions workflows, which also affected Copilot cloud agent sessions and GitHub Pages builds since they depend on Actions. During the peak of the incident, 30% of Actions jobs failed to start and 2% were delayed more than 5 minutes. <br /><br />The incident was triggered by a configuration change in an internal autoscaling component that contained outdated capacity threshold values. This caused a critical Actions service to scale below its required baseline, reducing capacity for workflow processing. We identified the regression, rolled back the change, and restored service capacity. New workflow executions recovered by 13:39 UTC. Full recovery was reached by 13:53 UTC after the queued backlog was drained. <br /><br />To prevent recurrence, we have added deployment guardrails to validate that autoscaling inputs are current and to detect drift between planned and live scaling state before autoscaling changes are applied.

## Below threshold (12)

- `anthropic` [Elevated errors for Claude Opus 4.8](https://stspg.io/31tdmk4s33mj) — 35m
- `yandexcloud` [The problem is with the API Gateway service](https://status.yandex.cloud/en/incidents/1888) — 1h 54m
- `github` [Disruption with some GitHub services](https://stspg.io/qd0t2882ckl9) — 1h 35m
- `stripe` [Elevated 3DS errors on Mastercard transactions](https://stspg.io/l0pf76xd55ny) — 1h 36m
- `circleci` [Delays starting jobs on Docker (Gen 2) resource classes](https://stspg.io/0t40l14n15yl) — 1h 38m
- `grafana` [Mimir Write Performance Degradation](https://stspg.io/m4g84bc8h8l6) — 1h 32m
- `github` [Claude Fable 5 experiencing degraded performance](https://stspg.io/vgv1m8xfnz3z) — 59m
- `circleci` [Errors accessing CircleCI and delays starting pipelines and scheduled workflows](https://stspg.io/jftpxjjgvqxy) — 1h 22m
- `circleci` [Issues with network routing for Mac jobs](https://stspg.io/vc3sdl5vsqyr) — 44m
- `github` [Incident with Webhooks](https://stspg.io/zgnhmrp08tdy) — 23m
- `stripe` [Elevated Issuing API Errors](https://stspg.io/16s7hszs8vj7) — 28m
- `grafana` [Fleet Managment Interface 404's](https://stspg.io/vxw38mcq61q4) — 0m

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-fa2aba4d84844a02ce489e4a"></a>

- Canonical ID: `ci-fa2aba4d84844a02ce489e4a` (1 source records).
  - [grafana/4pzq4gp7n92g](../../data/2026/W29/grafana.json); vendor URL: https://stspg.io/0gjtnfjby0jy
- Start time basis: `vendor_reported`.

<a id="ci-6e0762625d6fd8ea5d1979c9"></a>

- Canonical ID: `ci-6e0762625d6fd8ea5d1979c9` (1 source records).
  - [grafana/l67ftf1hw18n](../../data/2026/W29/grafana.json); vendor URL: https://stspg.io/9z649qh8csv8
- Start time basis: `vendor_reported`.

<a id="ci-668272b64c22371acd86e65c"></a>

- Canonical ID: `ci-668272b64c22371acd86e65c` (1 source records).
  - [gcp/3BvH3LVGcupoYqV6F4Nw](../../data/2026/W29/gcp.json); vendor URL: https://status.cloud.google.com/incidents/3BvH3LVGcupoYqV6F4Nw
- Start time basis: `vendor_reported`.

<a id="ci-6dcd181a465685cc632e7a19"></a>

- Canonical ID: `ci-6dcd181a465685cc632e7a19` (1 source records).
  - [digitalocean/8sjb7kmqnvk9](../../data/2026/W29/digitalocean.json); vendor URL: https://stspg.io/rzs04wclgpfq
- Start time basis: `vendor_reported`.

<a id="ci-686212956fe2a481e1d2b0d3"></a>

- Canonical ID: `ci-686212956fe2a481e1d2b0d3` (1 source records).
  - [anthropic/vjfp60ngq2zj](../../data/2026/W29/anthropic.json); vendor URL: https://stspg.io/8fyh12n7xnnh
- Start time basis: `vendor_reported`.

<a id="ci-f287bf87d3757080ff23133f"></a>

- Canonical ID: `ci-f287bf87d3757080ff23133f` (1 source records).
  - [anthropic/7gpjd8n56rlq](../../data/2026/W29/anthropic.json); vendor URL: https://stspg.io/bfnl00fvzpzp
- Start time basis: `vendor_reported`.

<a id="ci-bd1ab603180908b002793ec5"></a>

- Canonical ID: `ci-bd1ab603180908b002793ec5` (1 source records).
  - [github/8vfyvq16hzh9](../../data/2026/W29/github.json); vendor URL: https://stspg.io/w8d77c7t94zf
- Start time basis: `vendor_reported`.

<a id="ci-d9d5a770feb641c6d6579b62"></a>

- Canonical ID: `ci-d9d5a770feb641c6d6579b62` (1 source records).
  - [circleci/7cnl777wp9qz](../../data/2026/W29/circleci.json); vendor URL: https://stspg.io/vg35mbl8lxpg
- Start time basis: `vendor_reported`.

<a id="ci-302c651a19c62b6728ed0538"></a>

- Canonical ID: `ci-302c651a19c62b6728ed0538` (1 source records).
  - [digitalocean/84v7d9gkkzzf](../../data/2026/W29/digitalocean.json); vendor URL: https://stspg.io/wv68wmnz55ln
- Start time basis: `vendor_reported`.

<a id="ci-bc0450b96f0e92e21e6cb2f7"></a>

- Canonical ID: `ci-bc0450b96f0e92e21e6cb2f7` (1 source records).
  - [grafana/hcdk41fbltl8](../../data/2026/W29/grafana.json); vendor URL: https://stspg.io/r9ysx33l8xbw
- Start time basis: `vendor_reported`.

<a id="ci-0511caaf3d55692014d0d66a"></a>

- Canonical ID: `ci-0511caaf3d55692014d0d66a` (1 source records).
  - [gcp/T8gmtofFSTGT5tbhyciF](../../data/2026/W29/gcp.json); vendor URL: https://status.cloud.google.com/incidents/T8gmtofFSTGT5tbhyciF
- Start time basis: `vendor_reported`.

<a id="ci-c5d98cc59bbfe51ec56959a7"></a>

- Canonical ID: `ci-c5d98cc59bbfe51ec56959a7` (1 source records).
  - [grafana/v0r3mk7cb4x3](../../data/2026/W29/grafana.json); vendor URL: https://stspg.io/ddf067fl4z8h
- Start time basis: `vendor_reported`.

<a id="ci-c5e649e4bd1dc7c56258efc8"></a>

- Canonical ID: `ci-c5e649e4bd1dc7c56258efc8` (1 source records).
  - [grafana/v44gbjl0610f](../../data/2026/W29/grafana.json); vendor URL: https://stspg.io/56wgtyyjvnt7
- Start time basis: `vendor_reported`.

<a id="ci-a662c5334eb363b42393d526"></a>

- Canonical ID: `ci-a662c5334eb363b42393d526` (1 source records).
  - [anthropic/g613ntyj2pwf](../../data/2026/W29/anthropic.json); vendor URL: https://stspg.io/jv85qb4xwc4x
- Start time basis: `vendor_reported`.

<a id="ci-102b10e3b36d7cac610ec126"></a>

- Canonical ID: `ci-102b10e3b36d7cac610ec126` (1 source records).
  - [anthropic/kqbd7wm6hnnr](../../data/2026/W29/anthropic.json); vendor URL: https://stspg.io/bdw3y9kd205m
- Start time basis: `vendor_reported`.

<a id="ci-a68b9f479904aabec11c9264"></a>

- Canonical ID: `ci-a68b9f479904aabec11c9264` (1 source records).
  - [digitalocean/nfhvp4yhyrfl](../../data/2026/W29/digitalocean.json); vendor URL: https://stspg.io/x0s0pghvdcgv
- Start time basis: `vendor_reported`.

<a id="ci-4c390f37329ec200ad9039b7"></a>

- Canonical ID: `ci-4c390f37329ec200ad9039b7` (1 source records).
  - [yandexcloud/1883](../../data/2026/W29/yandexcloud.json); vendor URL: https://status.yandex.cloud/en/incidents/1883
- Start time basis: `vendor_reported`.

<a id="ci-5b251004294104d745470177"></a>

- Canonical ID: `ci-5b251004294104d745470177` (1 source records).
  - [github/d35pngcy6sb9](../../data/2026/W29/github.json); vendor URL: https://stspg.io/5c43kg01skdw
- Start time basis: `vendor_reported`.

<a id="ci-551a4ba5cb9849708aed550f"></a>

- Canonical ID: `ci-551a4ba5cb9849708aed550f` (1 source records).
  - [anthropic/bwz7ld34gcxt](../../data/2026/W29/anthropic.json); vendor URL: https://stspg.io/31tdmk4s33mj
- Start time basis: `vendor_reported`.

<a id="ci-60a58a8f99f5d2d08be56375"></a>

- Canonical ID: `ci-60a58a8f99f5d2d08be56375` (1 source records).
  - [grafana/1n8sgkngf2kg](../../data/2026/W29/grafana.json); vendor URL: https://stspg.io/j81jwzycv215
- Start time basis: `vendor_reported`.

<a id="ci-b5532b5b313b5952aef8dc39"></a>

- Canonical ID: `ci-b5532b5b313b5952aef8dc39` (1 source records).
  - [anthropic/0hzwq1ylg80g](../../data/2026/W29/anthropic.json); vendor URL: https://stspg.io/2kbn7kkr22fx
- Start time basis: `vendor_reported`.

<a id="ci-c047f38aeb02067a8f2a9dad"></a>

- Canonical ID: `ci-c047f38aeb02067a8f2a9dad` (1 source records).
  - [yandexcloud/1888](../../data/2026/W29/yandexcloud.json); vendor URL: https://status.yandex.cloud/en/incidents/1888
- Start time basis: `vendor_reported`.

<a id="ci-5c281e7a150c707d7209b592"></a>

- Canonical ID: `ci-5c281e7a150c707d7209b592` (1 source records).
  - [circleci/gsyjwybg477g](../../data/2026/W29/circleci.json); vendor URL: https://stspg.io/lzl0fmm6xmts
- Start time basis: `vendor_reported`.

<a id="ci-ce0abbc17704bc5d5d454f3f"></a>

- Canonical ID: `ci-ce0abbc17704bc5d5d454f3f` (1 source records).
  - [github/dfpfsngcwywf](../../data/2026/W29/github.json); vendor URL: https://stspg.io/qd0t2882ckl9
- Start time basis: `vendor_reported`.

<a id="ci-68dec95a2a2dfd05de0d5ec1"></a>

- Canonical ID: `ci-68dec95a2a2dfd05de0d5ec1` (1 source records).
  - [digitalocean/qzbsmdsw9lq8](../../data/2026/W29/digitalocean.json); vendor URL: https://stspg.io/8yc90q76srsl
- Start time basis: `vendor_reported`.

<a id="ci-5a0048c0d502177e0fe6f1ba"></a>

- Canonical ID: `ci-5a0048c0d502177e0fe6f1ba` (1 source records).
  - [github/gxycch3076xk](../../data/2026/W29/github.json); vendor URL: https://stspg.io/pzvjh0wyg401
- Start time basis: `vendor_reported`.

<a id="ci-681523e1556a6b1adcae2396"></a>

- Canonical ID: `ci-681523e1556a6b1adcae2396` (1 source records).
  - [grafana/0fk5b4z41xmw](../../data/2026/W29/grafana.json); vendor URL: https://stspg.io/t147yktp9h8k
- Start time basis: `vendor_reported`.

<a id="ci-cfa4fadf0a40e078b613b933"></a>

- Canonical ID: `ci-cfa4fadf0a40e078b613b933` (1 source records).
  - [stripe/s6hpm263c04x](../../data/2026/W29/stripe.json); vendor URL: https://stspg.io/l0pf76xd55ny
- Start time basis: `vendor_reported`.

<a id="ci-a0a6ef7d7ad417dca2fbc7ce"></a>

- Canonical ID: `ci-a0a6ef7d7ad417dca2fbc7ce` (1 source records).
  - [anthropic/lj0khk5vd7v7](../../data/2026/W29/anthropic.json); vendor URL: https://stspg.io/zfp1zk8gg7p3
- Start time basis: `vendor_reported`.

<a id="ci-016680466610f88c0a3bacdc"></a>

- Canonical ID: `ci-016680466610f88c0a3bacdc` (1 source records).
  - [circleci/rfz5wdpg2nl0](../../data/2026/W29/circleci.json); vendor URL: https://stspg.io/0t40l14n15yl
- Start time basis: `vendor_reported`.

<a id="ci-ab18b3352c70041c5891a27f"></a>

- Canonical ID: `ci-ab18b3352c70041c5891a27f` (1 source records).
  - [grafana/s96rdt3ph2m9](../../data/2026/W29/grafana.json); vendor URL: https://stspg.io/m4g84bc8h8l6
- Start time basis: `vendor_reported`.

<a id="ci-6bb17efa5faf4bf3b6b6c259"></a>

- Canonical ID: `ci-6bb17efa5faf4bf3b6b6c259` (1 source records).
  - [github/0f6hndclsk21](../../data/2026/W29/github.json); vendor URL: https://stspg.io/vgv1m8xfnz3z
- Start time basis: `vendor_reported`.

<a id="ci-ee8f665e6a97134ef7ef0722"></a>

- Canonical ID: `ci-ee8f665e6a97134ef7ef0722` (1 source records).
  - [circleci/2b55fgxy67rb](../../data/2026/W29/circleci.json); vendor URL: https://stspg.io/jftpxjjgvqxy
- Start time basis: `vendor_reported`.

<a id="ci-1f477a5e8571aadb8872587f"></a>

- Canonical ID: `ci-1f477a5e8571aadb8872587f` (1 source records).
  - [github/q27ttsnp0x4g](../../data/2026/W29/github.json); vendor URL: https://stspg.io/t0tfsmbt70w7
- Start time basis: `vendor_reported`.

<a id="ci-d08bac7d2025fad32f865624"></a>

- Canonical ID: `ci-d08bac7d2025fad32f865624` (1 source records).
  - [circleci/n1tc2lw9q7l0](../../data/2026/W29/circleci.json); vendor URL: https://stspg.io/vc3sdl5vsqyr
- Start time basis: `vendor_reported`.

<a id="ci-e52077046be7369d26af8134"></a>

- Canonical ID: `ci-e52077046be7369d26af8134` (1 source records).
  - [github/ydpk76bj34z8](../../data/2026/W29/github.json); vendor URL: https://stspg.io/zgnhmrp08tdy
- Start time basis: `vendor_reported`.

<a id="ci-219ba0527838f68f9ca04c3b"></a>

- Canonical ID: `ci-219ba0527838f68f9ca04c3b` (1 source records).
  - [stripe/v6sllx288pwy](../../data/2026/W29/stripe.json); vendor URL: https://stspg.io/16s7hszs8vj7
- Start time basis: `vendor_reported`.

<a id="ci-6111b525562a8a724b24daab"></a>

- Canonical ID: `ci-6111b525562a8a724b24daab` (1 source records).
  - [grafana/hmxvc7wjj9sg](../../data/2026/W29/grafana.json); vendor URL: https://stspg.io/vxw38mcq61q4
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
