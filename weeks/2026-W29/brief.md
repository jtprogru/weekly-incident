# Week 2026-W29 — top 5 of 26 above threshold

2026-07-13 to 2026-07-19 (UTC).

38 source records (40 saved revisions) → 38 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## 1. grafana — Issue with Dashboard Views Being Registered

- URL: https://stspg.io/0gjtnfjby0jy
- Window: 2026-07-14 08:10 to 2026-07-20 10:15 UTC (6d 2h 5m)
- Impact: minor, status: resolved
- Components: AWS Australia - prod-ap-southeast-2, AWS Australia - prod-au-southeast-1, AWS Brazil - prod-sa-east-1, AWS Canada - prod-ca-east-0, AWS Germany - prod-eu-west-2, AWS Germany - prod-eu-west-4, AWS India - prod-ap-south-1, AWS Ireland - prod-eu-west-6, AWS Japan - prod-ap-northeast-0, AWS Singapore - prod-ap-southeast-1, AWS Sweden - prod-eu-north-0, AWS Switzerland - prod-eu-central-0, AWS UAE - prod-me-central-1, AWS UK - prod-gb-south-1, AWS US East - prod-us-east-0, AWS US East - prod-us-east-2, AWS US West - prod-us-west-0, Azure Netherlands - prod-eu-west-3, Azure US Central - us-central2, Federal Cloud - AWS US Gov West, GCP Australia - prod-au-southeast-0, GCP Belgium - prod-eu-west-0, GCP Brazil - prod-sa-east-0, GCP India - prod-ap-south-0, GCP Singapore - prod-ap-southeast-0, GCP UK - prod-gb-south-0, GCP US Central - prod-us-central-0, GCP US Central - prod-us-central-3, GCP US Central - prod-us-central-4, GCP US East - prod-us-east-1, play.grafana.org

<a id="ci-fa2aba4d84844a02ce489e4a"></a>

- Canonical ID: `ci-fa2aba4d84844a02ce489e4a` (1 source records).
  - [grafana/4pzq4gp7n92g](../../data/2026/W29/grafana.json); vendor URL: https://stspg.io/0gjtnfjby0jy
- Start time basis: `vendor_reported`.

  - `2026-07-14 08:10` investigating — This is related to https://status.grafana.com/incidents/rhrk2ck6ly0y which was resolved by mistake. For some dashboards, the Views/Error counts in the Dashboards/Folder list renders as `-` and never updates, even after dashboards are viewed repeatedly. This is ultimately causing inaccurate or missing view counts. We are continuing to investigate the underlying cause of this issue. At this time, the scope and impact remain unchanged, and we have no new information to share. We will provide another update as soon as more information becomes available.
  - `2026-07-14 14:27` identified — We've identified the cause of the issue impacting dashboard view and error counts in the Dashboards and Folder list. Our team is currently working on implementing a fix and validating the solution. We will provide another update within the next 24 hours, or sooner if we have additional information to share.
  - `2026-07-15 12:05` identified — We continue to work on resolving this issue. We'll provide another update within the next 24 hours, or sooner if additional information becomes available.
  - `2026-07-15 20:52` identified — We continue to work on resolving this issue. We'll provide another update within the next 24 hours, or sooner if additional information becomes available.
  - `2026-07-16 13:34` identified — We continue to work on resolving the issue impacting dashboard view and error counts. There are no new updates to share at this time. Our next update will be provided within 24 hours.
  - `2026-07-16 21:27` identified — We continue to work on resolving the issue impacting dashboard view and error counts. There are no new updates to share at this time. Our next update will be provided within 24 hours.
  - `2026-07-20 08:17` identified — We are currently working on a fix related to this incident. There are no new updates to share at this time.
  - `2026-07-20 10:15` resolved — A fix has been implemented and dashboard view and error counts in the Dashboards and Folder list are updating as expected. Thank you for your patience while we worked to address this issue.

## 2. grafana — Grafana Cloud login issues for users with role of None

- URL: https://stspg.io/9z649qh8csv8
- Window: 2026-07-17 18:22 to 2026-07-18 13:15 UTC (18h 52m)
- Impact: minor, status: resolved
- Components: AWS Australia - prod-ap-southeast-2, AWS Australia - prod-au-southeast-1, AWS Brazil - prod-sa-east-1, AWS Canada - prod-ca-east-0, AWS Germany - prod-eu-west-2, AWS Germany - prod-eu-west-4, AWS India - prod-ap-south-1, AWS Ireland - prod-eu-west-6, AWS Japan - prod-ap-northeast-0, AWS Singapore - prod-ap-southeast-1, AWS Sweden - prod-eu-north-0, AWS Switzerland - prod-eu-central-0, AWS UAE - prod-me-central-1, AWS UK - prod-gb-south-1, AWS US East - prod-us-east-0, AWS US East - prod-us-east-2, AWS US West - prod-us-west-0, Azure Netherlands - prod-eu-west-3, Azure US Central - us-central2, Federal Cloud - AWS US Gov West, GCP Australia - prod-au-southeast-0, GCP Belgium - prod-eu-west-0, GCP Brazil - prod-sa-east-0, GCP India - prod-ap-south-0, GCP Singapore - prod-ap-southeast-0, GCP UK - prod-gb-south-0, GCP US Central - prod-us-central-0, GCP US Central - prod-us-central-3, GCP US Central - prod-us-central-4, GCP US East - prod-us-east-1, play.grafana.org

<a id="ci-6e0762625d6fd8ea5d1979c9"></a>

- Canonical ID: `ci-6e0762625d6fd8ea5d1979c9` (1 source records).
  - [grafana/l67ftf1hw18n](../../data/2026/W29/grafana.json); vendor URL: https://stspg.io/9z649qh8csv8
- Start time basis: `vendor_reported`.

  - `2026-07-17 18:22` investigating — We’re currently investigating an issue with user logins with the None role in Grafana Cloud. Our team is actively working to identify the cause. Thank you for your patience.
  - `2026-07-17 19:28` identified — We’ve identified the cause of the issue impacting users with the role of None from logging in. Our team is currently implementing a fix.
  - `2026-07-17 21:33` identified — We are in the process of implementing a fix and are monitoring the rollout. Thank you for your patience.
  - `2026-07-18 13:15` resolved — Users with the None role authenticating via Grafana.com should be able to access their stacks again

## 3. gcp — Google Cloud VMware Engine (GCVE), Google Cloud NetApp Volumes, and Bare Metal Solutions (BMS) services are experiencing a service outage in europe-west4-a due to a cooling failure.

- URL: https://status.cloud.google.com/incidents/3BvH3LVGcupoYqV6F4Nw
- Window: 2026-07-15 23:57 to 2026-07-16 12:25 UTC (12h 28m)
- Impact: major, status: resolved
- Components: Bare Metal Solution, Google Cloud NetApp Volumes, VMWare engine

<a id="ci-668272b64c22371acd86e65c"></a>

- Canonical ID: `ci-668272b64c22371acd86e65c` (1 source records).
  - [gcp/3BvH3LVGcupoYqV6F4Nw](../../data/2026/W29/gcp.json); vendor URL: https://status.cloud.google.com/incidents/3BvH3LVGcupoYqV6F4Nw
- Start time basis: `vendor_reported`.

  - `2026-07-16 03:30` investigating — **Summary** Google Cloud VMware Engine (GCVE), Google Cloud NetApp Volumes, and Bare Metal Solution (BMS) services are experiencing temperature alerts in europe-west4-a. **Description** **GCVE:** Our engineering teams are currently working with data center facility teams to address the rising temperatures. A shutdown has already occurred for some private clouds, and we expect that others will shut down as well. **We recommend shutting down your VMware workload VMs for precaution so that recovery is safer. A shutdown of all systems started at 19:00 PDT \[4:00 Local Time (GMT+2)\]** We will notify you as soon as your Private Cloud has been powered back up and it is safe to restart the VMs. **BMS:** Some Bare Metal Solution servers are experiencing high temperatures and we are in the process…
  - `2026-07-16 04:17` investigating — **Summary** Google Cloud VMware Engine (GCVE), Google Cloud NetApp Volumes, and Bare Metal Solution (BMS) services are experiencing temperature alerts in europe-west4-a **Description** The datacenter serving europe-west4-a for GCVE, BMS, and NetApp has experienced a power failure, which subsequently caused a cooling failure. Google has proactively turned down workloads in order to protect customer data from any risks posed by running infrastructure in a high temperature environment. The datacenter operator team is working diligently to bring cooling and power back online so we can safely restart affected workloads. **GCVE:** Our engineering teams have shut down private clouds, and we expect that others will shut down as well. We will notify customers as soon as their private cloud has been…
  - `2026-07-16 04:41` investigating — **Summary** Google Cloud VMware Engine (GCVE), Google Cloud NetApp Volumes, and Bare Metal Solution (BMS) services are experiencing a service outage in europe-west4-a due to a cooling failure **Description** The datacenter serving europe-west4-a for GCVE, BMS, and NetApp has experienced a power failure, which subsequently caused a cooling failure. Google has proactively turned down workloads in order to protect customer data from any risks posed by running infrastructure in a high temperature environment. The datacenter operator team is working diligently to bring cooling and power back online so we can safely restart affected workloads. Some cooling has been restored, and we are working to restore affected workloads as available capacity allows. **GCVE:** Cooling has been restored at the…
  - `2026-07-16 06:29` investigating — **Summary** Google Cloud VMware Engine (GCVE), Google Cloud NetApp Volumes, and Bare Metal Solution (BMS) services are experiencing a service outage in europe-west4-a due to a cooling failure. **Description** The datacenter serving europe-west4-a for GCVE, BMS, and NetApp has experienced a power failure, which subsequently caused a cooling failure. Google proactively turned down workloads in order to protect customer data from any risks posed by running infrastructure in a high temperature environment. Cooling and power have been restored at the site. The datacenter operator team is now working to safely restart affected workloads. **GCVE:** As a precautionary measure, our engineering teams had shut down private clouds to protect these clouds from any damage. Service restoration is still i…
  - `2026-07-16 07:39` investigating — **Summary** Google Cloud VMware Engine (GCVE), Google Cloud NetApp Volumes, and Bare Metal Solution (BMS) services are experiencing a service outage in europe-west4-a due to a cooling failure. **Description** The datacenter serving europe-west4-a for GCVE, BMS, and NetApp has experienced a power failure, which subsequently caused a cooling failure. Google proactively turned down workloads in order to protect customer data from any risks posed by running infrastructure in a high temperature environment. Cooling and power have been restored at the site. The datacenter operator team is now working to safely restart affected workloads. **GCVE:** As a precautionary measure, our engineering teams had shut down private clouds to protect these clouds from any damage. Network fabric has been restor…
  - `2026-07-16 08:32` investigating — **Summary** Google Cloud VMware Engine (GCVE), Google Cloud NetApp Volumes, and Bare Metal Solution (BMS) services experienced a service outage in europe-west4-a due to a cooling failure. **Description** The datacenter serving europe-west4-a for GCVE, BMS, and NetApp has experienced a power failure, which subsequently caused a cooling failure. Google proactively turned down workloads in order to protect customer data from any risks posed by running infrastructure in a high temperature environment. Cooling and power have been restored at the site. The datacenter operator team is now validating site stability. **GCVE:** Network fabric has been restored, and Private Cloud restoration remains underway. We are notifying the affected customers as their Private Clouds are restored to allow for VM…
  - `2026-07-16 09:21` investigating — **Summary** Google Cloud VMware Engine (GCVE), Google Cloud NetApp Volumes, and Bare Metal Solution (BMS) services experienced a service outage in europe-west4-a due to a cooling failure. **Description** The datacenter serving europe-west4-a for GCVE, BMS, and NetApp has experienced a power failure, which subsequently caused a cooling failure. Google proactively turned down workloads in order to protect customer data from any risks posed by running infrastructure in a high temperature environment. Cooling and power have been restored at the site. The datacenter operator team is now validating site stability. **GCVE:** Network fabric has been restored, and Private Cloud restoration remains underway. We are notifying the affected customers as their Private Clouds are restored to allow for VM…
  - `2026-07-16 10:14` investigating — **Summary** Google Cloud VMware Engine (GCVE), Google Cloud NetApp Volumes, and Bare Metal Solution (BMS) services experienced a service outage in europe-west4-a due to a cooling failure. **Description** The datacenter serving europe-west4-a for GCVE, BMS, and NetApp has experienced a power failure, which subsequently caused a cooling failure. Google proactively turned down workloads in order to protect customer data from any risks posed by running infrastructure in a high temperature environment. Cooling and power have been restored at the site. The issue with GCVE and Google Cloud NetApp Volumes has been resolved for all affected users. We continue to recover the impact for BMS at the moment. **GCVE [Mitigated]:** Service restoration is complete across all the Private Clouds. Customers h…
  - `2026-07-16 11:39` investigating — **Summary** Google Cloud VMware Engine (GCVE), Google Cloud NetApp Volumes, and Bare Metal Solution (BMS) services experienced a service outage in europe-west4-a due to a cooling failure. **Description** The datacenter serving europe-west4-a for GCVE, BMS, and NetApp has experienced a power failure, which subsequently caused a cooling failure. Google proactively turned down workloads in order to protect customer data from any risks posed by running infrastructure in a high temperature environment. Cooling and power have been restored at the site. The issue with GCVE and Google Cloud NetApp Volumes has been resolved for all affected users. We continue to recover the impact for BMS at the moment. **GCVE [Mitigated]:** Service restoration is complete across all the Private Clouds. Customers h…
  - `2026-07-16 12:25` investigating — **Summary** Google Cloud VMware Engine (GCVE), Google Cloud NetApp Volumes, and Bare Metal Solution (BMS) services experienced a service outage in europe-west4-a due to a cooling failure. **Description** The datacenter serving europe-west4-a for GCVE, BMS, and NetApp has experienced a power failure, which subsequently caused a cooling failure. Google proactively turned down workloads in order to protect customer data from any risks posed by running infrastructure in a high temperature environment. Cooling and power have been restored at the site. The issue with GCVE and Google Cloud NetApp Volumes has been resolved for all affected users as of Thursday, 2026-07-16 02:31 PDT and 01:00 PDT respectively. The issue with BMS is largely resolved and is believed to be affecting a very small numbe…
  - `2026-07-18 00:22` resolved — # Preliminary Incident Report We sincerely apologize for the disruption this incident caused to your business. We know how much you rely on Google Cloud, and we regret the impact on your productivity. Please note, this information is based on our best knowledge at the time of posting and is subject to change as our investigation continues. A final Incident Report with preventative actions will be posted once our investigation is complete. If you have experienced impact outside of what is listed below, please reach out to Google Cloud Support using https://cloud.google.com/support. ## Date/Time of the Issue (All time US/Pacific) **Google Cloud VMware Engine Impact:** - Start: 15 July 2026 17:09 - End: 16 July 2026 2:33 - Duration: 9 Hours, 24 Minutes **Google Cloud NetApp Volumes Impact:**…
  - `2026-07-25 13:16` resolved — ## \# Incident Report ## \#\# Summary On Wednesday, 15 July 2026, Google Cloud VMware Engine (GCVE), Bare Metal Solution (BMS), and Google Cloud NetApp Volumes (GCNV) experienced service interruptions for a total duration of 14 hours, 55 minutes. The root cause for this outage is a 3ms voltage drop in the power feed from the utility provider and subsequent utility breaker protective action. This incident was mitigated by the data center provider rectifying the failed systems followed by the Google engineering team restoring the services. ## \#\# Root Cause A regional data center hosting services in europe-west4-a experienced an upstream voltage transient affecting both utility power feeds A and B. During this event, the utility breakers on both A and B feeds tripped and initiated the trans…

## 4. digitalocean — Serverless Inference - Gemma4 Latency Issues Causing Timeouts & Slow Responses

- URL: https://stspg.io/rzs04wclgpfq
- Window: 2026-07-13 19:04 to 2026-07-15 21:46 UTC (2d 2h 42m)
- Impact: minor, status: resolved
- Components: Inference

<a id="ci-6dcd181a465685cc632e7a19"></a>

- Canonical ID: `ci-6dcd181a465685cc632e7a19` (1 source records).
  - [digitalocean/8sjb7kmqnvk9](../../data/2026/W29/digitalocean.json); vendor URL: https://stspg.io/rzs04wclgpfq
- Start time basis: `vendor_reported`.

  - `2026-07-13 19:04` identified — We are currently experiencing an issue where customers using the Gemma 4 model on our Serverless Inference and Dedicated Inference platforms may experience severe latency or request timeouts. Our engineering team has identified a backend configuration issue as the root cause, which is temporarily degrading performance. We are actively working on a fix to restore normal response times and will provide another update as soon as the mitigation is in place
  - `2026-07-13 22:24` monitoring — A fix has been deployed to resolve the backend configuration issue. We are closely monitoring system performance to ensure full recovery and normal response times for all Gemma 4 inference workflows.
  - `2026-07-14 11:07` identified — We are currently experiencing an issue affecting customers using the Gemma 4 model on our Serverless Inference platform. Customers may experience significantly increased latency or request timeouts. Our Engineering team has identified a backend configuration issue as the root cause, which is temporarily impacting model performance. Please be assured that our Engineering team is actively working on a fix and is treating this issue with high priority. We sincerely apologise for any inconvenience this may have caused and appreciate your patience and understanding. If you have any further questions, please create a support ticket so that we can investigate your specific case further.
  - `2026-07-15 10:06` monitoring — A fix has been deployed to resolve the issue. We are closely monitoring system performance to ensure full recovery and normal response times for all Gemma 4 inference workflows.
  - `2026-07-15 21:46` resolved — The deployed fix has successfully restored full functionality, and our monitoring shows that system performance has completely stabilized. Response times for all Gemma 4 inference workflows have returned to normal baseline levels. We will continue to track platform stability moving forward to ensure long-term reliability. We apologize for any disruption this may have caused to your workflows and appreciate your patience throughout the recovery process.

## 5. anthropic — Elevated errors for Claude Opus 4.7

- URL: https://stspg.io/8fyh12n7xnnh
- Window: 2026-07-16 08:58 to 2026-07-16 13:30 UTC (4h 32m)
- Impact: minor, status: resolved
- Components: Claude API (api.anthropic.com), Claude Code, Claude Console (platform.claude.com), Claude Cowork, claude.ai

<a id="ci-686212956fe2a481e1d2b0d3"></a>

- Canonical ID: `ci-686212956fe2a481e1d2b0d3` (1 source records).
  - [anthropic/vjfp60ngq2zj](../../data/2026/W29/anthropic.json); vendor URL: https://stspg.io/8fyh12n7xnnh
- Start time basis: `vendor_reported`.

  - `2026-07-16 08:58` investigating — We are currently investigating this issue.
  - `2026-07-16 13:30` resolved — This incident has been resolved.

## Also above threshold (21)

- `anthropic` Elevated errors on Sonnet 5 and Haiku 4.5 (5h 33m) — https://stspg.io/bfnl00fvzpzp
- `github` Incident with GitHub Actions (5h 9m) — https://stspg.io/w8d77c7t94zf
- `circleci` Issues with network routing for Mac jobs (1d 1h 26m) — https://stspg.io/vg35mbl8lxpg
- `digitalocean` Block Storage Volume NYC1, NYC3, SGP1, SYD1 and BLR1 (5h 41m) — https://stspg.io/wv68wmnz55ln
- `grafana` Partial Outage in prod-eu-west-2 (1h 12m) — https://stspg.io/r9ysx33l8xbw
- `gcp` Google Cloud VMware Engine (GCVE) Stretched Cluster customers are experiencing zonal outages impacting network connectivity across multiple regions. (10h 40m) — https://status.cloud.google.com/incidents/T8gmtofFSTGT5tbhyciF
- `grafana` Some Reports of Grafana Not Loading. (2h 53m) — https://stspg.io/ddf067fl4z8h
- `grafana` Delayed Aggregated Metrics (prod-us-central-0) (8h 35m) — https://stspg.io/56wgtyyjvnt7
- `anthropic` Elevated errors across Fable 5 (1h 11m) — https://stspg.io/jv85qb4xwc4x
- `anthropic` Elevated errors for multiple models (4h 16m) — https://stspg.io/bdw3y9kd205m
- `digitalocean` Reserved IP routing in TOR1 (3h 44m) — https://stspg.io/x0s0pghvdcgv
- `yandexcloud` The problem is with the UI for the Managed Service for YDB (3h 30m) — https://status.yandex.cloud/en/incidents/1883
- `github` Disruption with some GitHub services (3h 7m) — https://stspg.io/5c43kg01skdw
- `grafana` Mimir Partial Write Outage (1h 53m) — https://stspg.io/j81jwzycv215
- `anthropic` Enterprise SSO sign-in failures (1h 3m) — https://stspg.io/2kbn7kkr22fx
- `circleci` Login issues for some Bitbucket users (2h 25m) — https://stspg.io/lzl0fmm6xmts
- `digitalocean` Networking in BLR1 (2h 54m) — https://stspg.io/8yc90q76srsl
- `github` Degraded REST API Availability (1h 22m) — https://stspg.io/pzvjh0wyg401
- `grafana` Adaptive Metrics aggregation delay in eu-west-0 region. (2h 2m) — https://stspg.io/t147yktp9h8k
- `anthropic` Elevated error rates for Sonnet 5 (14m) — https://stspg.io/zfp1zk8gg7p3
- `github` Actions runs are experiencing failures to start (21m) — https://stspg.io/t0tfsmbt70w7
