# Week 2026-W27 — top 5 of 22 above threshold

2026-06-29 to 2026-07-05 (UTC).

28 source records (28 saved revisions) → 28 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## 1. grafana — Some Queries Failing

- URL: https://stspg.io/hmkbllwz66w2
- Window: 2026-07-02 18:10 to 2026-07-02 21:35 UTC (3h 25m)
- Impact: minor, status: resolved
- Components: AWS Australia - prod-ap-southeast-2, AWS Brazil - prod-sa-east-1, AWS Canada - prod-ca-east-0, AWS Germany - prod-eu-west-2, AWS India - prod-ap-south-1, AWS Ireland - prod-eu-west-6, AWS Japan - prod-ap-northeast-0, AWS Singapore - prod-ap-southeast-1, AWS Sweden - prod-eu-north-0, AWS Switzerland - prod-eu-central-0, AWS UAE - prod-me-central-1, AWS US East - prod-us-east-0, AWS US West - prod-us-west-0, Azure Netherlands - prod-eu-west-3, Azure US Central - us-central2, Azure US Central - us-central7, Federal Cloud - AWS US Gov West, GCP Australia - prod-au-southeast-0, GCP Belgium - prod-eu-west-0, GCP Brazil - prod-sa-east-0, GCP India - prod-ap-south-0, GCP Singapore - prod-ap-southeast-0, GCP UK - prod-gb-south-0, GCP US Central - prod-us-central-0, GCP US Central - prod-us-central-5

<a id="ci-b2091b3ec81242da991f3243"></a>

- Canonical ID: `ci-b2091b3ec81242da991f3243` (1 source records).
  - [grafana/x0h2rsqpqk3w](../../data/2026/W27/grafana.json); vendor URL: https://stspg.io/hmkbllwz66w2
- Start time basis: `vendor_reported`.

  - `2026-07-02 18:10` identified — Queries with drop __error__, including log volume histogram queries (the queries that generate the histogram visualization in Grafana), are failing due to a bug with series limit checks. The Root cause has been identified, and we are working on a fix.
  - `2026-07-02 20:16` identified — We are continuing to work on rolling back a PR responsible for this behavior. Once we have more information, we will share it here. Thank you for your patience.
  - `2026-07-02 21:35` resolved — This incident has been resolved. Thank you for your patience.

## 2. grafana — Elevated Loki Query Bytes Reporting

- URL: https://stspg.io/wqbptt5wfqtg
- Window: 2026-07-01 19:43 to 2026-07-01 22:00 UTC (2h 17m)
- Impact: minor, status: resolved
- Components: AWS Australia - prod-ap-southeast-2, AWS Brazil - prod-sa-east-1, AWS Canada - prod-ca-east-0, AWS Germany - prod-eu-west-2, AWS India - prod-ap-south-1, AWS Ireland - prod-eu-west-6, AWS Japan - prod-ap-northeast-0, AWS Singapore - prod-ap-southeast-1, AWS Sweden - prod-eu-north-0, AWS Switzerland - prod-eu-central-0, AWS UAE - prod-me-central-1, AWS US East - prod-us-east-0, AWS US West - prod-us-west-0, Azure Netherlands - prod-eu-west-3, Azure US Central - us-central2, Azure US Central - us-central7, Federal Cloud - AWS US Gov West, GCP Australia - prod-au-southeast-0, GCP Belgium - prod-eu-west-0, GCP Brazil - prod-sa-east-0, GCP India - prod-ap-south-0, GCP Singapore - prod-ap-southeast-0, GCP UK - prod-gb-south-0, GCP US Central - prod-us-central-0, GCP US Central - prod-us-central-5

<a id="ci-403c0d3369f9b580168e1f9e"></a>

- Canonical ID: `ci-403c0d3369f9b580168e1f9e` (1 source records).
  - [grafana/j5vgks57cphf](../../data/2026/W27/grafana.json); vendor URL: https://stspg.io/wqbptt5wfqtg
- Start time basis: `vendor_reported`.

  - `2026-07-01 19:43` investigating — We are investigating an issue where some customers may see higher Loki query byte usage reported than was actually consumed. This affects usage reporting only; there is no impact to query execution or service availability. The issue began at approximately 13:20 UTC and is ongoing. We expect the issue to be resolved soon and will provide another update as more information becomes available.
  - `2026-07-01 22:00` resolved — We’ve implemented a fix and can confirm the issue is fully resolved as of 20:25 UTC. Thank you for your patience.

## 3. grafana — Mimir read errors and high latency in prod-eu-west-0

- URL: https://stspg.io/h6ydsghh5vvs
- Window: 2026-06-29 11:37 to 2026-06-30 08:20 UTC (20h 43m)
- Impact: minor, status: resolved
- Components: GCP Belgium - prod-eu-west-0: Ingestion, GCP Belgium - prod-eu-west-0: Querying

<a id="ci-2643670698537376487b6cc4"></a>

- Canonical ID: `ci-2643670698537376487b6cc4` (1 source records).
  - [grafana/grtt5vk2dd3v](../../data/2026/W27/grafana.json); vendor URL: https://stspg.io/h6ydsghh5vvs
- Start time basis: `vendor_reported`.

  - `2026-06-29 11:37` investigating — We are currently investigating an issue with Mimir in prod-eu-west-0 we are seeing read errors and high latency. This incident is currently ongoing. The errors are recovering but we are currently looking into the route cause of this.
  - `2026-06-29 13:52` investigating — The errors are recovering, and we are still looking into the root cause of this.
  - `2026-06-29 15:37` investigating — The errors and latency have now recovered, we continue investigating the root cause.
  - `2026-06-29 17:47` identified — We've identified a possible cause, and a mitigation is in place to prevent further occurrences.
  - `2026-06-30 08:20` resolved — Since the mitigation has been applied, we have not seen the errors return. At this point, we are considering the incident resolved.

## 4. digitalocean — Agent Timeouts While Retrieving Data from Knowledge Bases

- URL: https://stspg.io/sp8b09xxqvjt
- Window: 2026-07-01 06:25 to 2026-07-02 03:38 UTC (21h 13m)
- Impact: none, status: resolved
- Components: Agent Runtime, Knowledge Bases

<a id="ci-c016a22dbd3c939ab94c7e55"></a>

- Canonical ID: `ci-c016a22dbd3c939ab94c7e55` (1 source records).
  - [digitalocean/76l7xym7jd5m](../../data/2026/W27/digitalocean.json); vendor URL: https://stspg.io/sp8b09xxqvjt
- Start time basis: `vendor_reported`.

  - `2026-07-01 06:25` investigating — Our Engineering team is currently investigating an issue where Agents are experiencing timeouts while retrieving data from Knowledge Bases. As a result, affected Agents may fail to retrieve data and return the following error: "Failed to retrieve data from Knowledge base(s) - timeout" Please be assured that we are treating this as a high-priority issue and are actively working to mitigate it. We sincerely apologise for any inconvenience this may have caused and appreciate your patience and understanding. If you have any further questions, please create a support ticket so that we can investigate your specific case further.
  - `2026-07-01 08:39` identified — Our Engineering team has identified the issue causing Agents to experience timeouts while retrieving data from Knowledge Bases. Please be assured that our Engineering team is actively working on a fix and is treating this issue with high priority. We sincerely apologise for any inconvenience this may have caused and appreciate your patience and understanding. If you have any further questions, please create a support ticket so that we can investigate your specific case further.
  - `2026-07-02 03:38` resolved — This incident has been resolved, and our teams continue to monitor the results. If you experience any further issues, please contact support.

## 5. datadog — Delayed data processing and errors in multiple products

- URL: https://stspg.io/5snql687z9c4
- Window: 2026-06-30 17:51 to 2026-06-30 20:32 UTC (2h 40m)
- Impact: major, status: resolved
- Components: APM, CI Visibility, Code Coverage, Database Monitoring, Error Tracking, Observability Pipelines, RUM, Synthetics, Test Optimization

<a id="ci-caba25cbe35e7151a8c9f079"></a>

- Canonical ID: `ci-caba25cbe35e7151a8c9f079` (1 source records).
  - [datadog/71vvxf0rxxnz](../../data/2026/W27/datadog.json); vendor URL: https://stspg.io/5snql687z9c4
- Start time basis: `vendor_reported`.

  - `2026-06-30 17:51` investigating — We are currently investigating an issue affecting CI Visibility, Test Optimization, Code Coverage, Software Delivery, Preemptive Alerting, Session Replay, RUM Explorer, Agent Observability, Database Monitoring, Event Correlation, Error Tracking, and all Security products in our US1 datacenter. As a result, some users may experience delays in data processing and errors when using affected products. Our engineering teams have identified the issue and are actively working on a resolution. Our engineering teams are actively working to resolve the issue, and we'll follow up with more information shortly. Please let us know if you have any questions or need anything else in the meantime.
  - `2026-06-30 17:54` investigating — We are continuing to investigate this issue.
  - `2026-06-30 18:28` investigating — We are continuing to investigate this issue.
  - `2026-06-30 18:36` investigating — We are continuing to investigate this issue.
  - `2026-06-30 18:57` identified — The issue has been identified and a fix is being implemented.
  - `2026-06-30 19:06` identified — We are continuing to work on a fix for this issue.
  - `2026-06-30 19:36` monitoring — A fix has been implemented and we are monitoring the results.
  - `2026-06-30 20:32` resolved — This incident has been resolved.

## Also above threshold (17)

- `grafana` Grafana Cloud IRM alert groups failing to produce alerts in us-east-3 region. (19h 59m) — https://stspg.io/34p7wm05ydt0
- `yandexcloud` The problem is with SpeechKit speech recognition services. (12h 0m) — https://status.yandex.cloud/en/incidents/1869
- `digitalocean` Managed Database Clusters (10h 15m) — https://stspg.io/rpkck6zr7ht5
- `stripe` Elevated WeChat Pay errors (12h 48m) — https://stspg.io/17c6mmm4lg2f
- `digitalocean` DNS Lookup Failures for Managed Database (4h 52m) — https://stspg.io/r6yw4xc720r8
- `circleci` Some customers are unable to access their organizations (6h 34m) — https://stspg.io/ty6ll6xj9c18
- `digitalocean` Limited access to Deepseek V4 Pro model (2h 39m) — https://stspg.io/wwsc9zft1lxd
- `github` Delays in copilot budget limits resets for some users (2h 34m) — https://stspg.io/jy3yhfvwxpjh
- `grafana` Loki and Frontend Observability - Major Outage in prod-us-central-0 region (1h 53m) — https://stspg.io/3v5hcx1v7p32
- `grafana` Rule evaluation error on cluster prod-gb-south-0 (2h 38m) — https://stspg.io/qldkxzbdb7qp
- `circleci` Outage impacting web experience and pipelines starting (1h 18m) — https://stspg.io/n86hwbb1whpc
- `datadog` Elevated Error Rates (2h 5m) — https://stspg.io/h1vdjddg90hd
- `grafana` Confluent API Outage (1h 49m) — https://stspg.io/b3bbl34zgmbw
- `digitalocean` Droplet Resize (2h 14m) — https://stspg.io/m6y2gmlz3jrv
- `yandexcloud` Partial unavailability of Virtual Private Cloud services (1h 6m) — https://status.yandex.cloud/en/incidents/1865
- `datadog` Errors Pulling Container Images (1h 3m) — https://stspg.io/gs85vq0xmxcx
- `datadog` Delayed Monitors Notifications (43m) — https://stspg.io/hsz1r8ww26gn
