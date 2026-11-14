# Week 2026-W17

2026-04-20 to 2026-04-26 (UTC). Generated 2026-11-14 00:00.

12 canonical incidents, 6 above threshold. All 14 sources responded.

12 source records (12 saved revisions) → 12 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | yandexcloud | [Unavailability of Yandex Cloud Backup during maintenance work](https://status.yandex.cloud/en/incidents/1688) | 2026-04-21 08:00 | 7h 10m | major | Cloud Backup | 645 = 430 × 1 × 1.5 |
| 2 | slack | [Slack Threads Not Loading](https://slack-status.com/2026-04/17580e43d51e0bf2) | 2026-04-23 07:27 | 2h 23m | unknown | Connectivity | 172 = 143 × 1 × 1.2 |
| 3 | slack | [Custom status not expiring](https://slack-status.com/2026-04/be1299791e068265) | 2026-04-21 13:46 | 2h 1m | unknown | Messaging | 145 = 121 × 1 × 1.2 |
| 4 | yandexcloud | [Unavailability of dashboards in Monium](https://status.yandex.cloud/en/incidents/1686) | 2026-04-20 14:34 | 1h 21m | major | Monium | 122 = 81 × 1 × 1.5 |
| 5 | digitalocean | [App Platform Deployments](https://stspg.io/0rby3knkk6s1) | 2026-04-23 08:08 | 2h 1m | minor | Global | 97 = 121 × 1 × 0.8 |
| 6 | datadog | [Delays in AWS and Azure cloud integration metrics ingestion](https://stspg.io/xw01xpvggm3n) | 2026-04-24 16:27 | 23m | major | Metrics and Infra Monitoring, Monitors | 46 = 23 × 2 × 1.0 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. yandexcloud — Unavailability of Yandex Cloud Backup during maintenance work

- URL: https://status.yandex.cloud/en/incidents/1688
- Started: 2026-04-21 08:00 UTC
- Resolved: 2026-04-21 15:10 UTC
- Duration: 7h 10m
- Impact: major, status: resolved
- Components: Cloud Backup
- Canonical incident: [ci-6b1ea05d787d7ffece86be09](#ci-6b1ea05d787d7ffece86be09)

- `2026-04-21 07:41` **investigating** — Technical work is underway in the Cloud Backup service: During the period 11:00 - 13:00 Moscow time, service operations may be unavailable for up to 45 minutes in all availability zones, including the kz1-a zone.; During the period 13:00 - 19:00 Moscow time, it is possible to reduce the speed of creating and restoring backups in the RU availability zones.
- `2026-04-21 15:21` **resolved** — The technical work has been completed, and the service is operating normally. If you have any questions or problems, please contact technical support.

### 2. slack — Slack Threads Not Loading

- URL: https://slack-status.com/2026-04/17580e43d51e0bf2
- Started: 2026-04-23 07:27 UTC
- Resolved: 2026-04-23 09:50 UTC
- Duration: 2h 23m
- Impact: unknown, status: resolved
- Components: Connectivity
- Canonical incident: [ci-ce2635506b22350da8a15dc1](#ci-ce2635506b22350da8a15dc1)

- `2026-04-23 07:27` **unknown** — Some users are having trouble loading Threads in Slack, this is affecting Windows and Chrome users. Users may receive an error asking them to ‘Try Again’. We apologize for the inconvenience and will report back as more information becomes available.
- `2026-04-23 08:05` **unknown** — We’re still investigating these issues with Threads not loading. No additional information to share just yet, but we’ll keep you posted as soon as we have an update.
- `2026-04-23 09:50` **unknown** — Further investigation confirms that the issue impacting Thread loading has been self- resolved, and our team will continue to investigate to ensure that this doesn’t reoccur. We apologize for the trouble today and thank you for your patience.

### 3. slack — Custom status not expiring

- URL: https://slack-status.com/2026-04/be1299791e068265
- Started: 2026-04-21 13:46 UTC
- Resolved: 2026-04-21 15:47 UTC
- Duration: 2h 1m
- Impact: unknown, status: resolved
- Components: Messaging
- Canonical incident: [ci-b14ab11f938e3cc921c51700](#ci-b14ab11f938e3cc921c51700)

- `2026-04-21 13:46` **unknown** — We're aware of an issue affecting Slack. Some users who have set a custom status to expire at a specific time are finding that their status is not automatically clearing as expected. We're actively investigating and we'll provide an update as soon as we have more information available.
- `2026-04-21 15:47` **unknown** — Resolved This issue is now resolved for all users. Current status and recent actions taken: Custom statuses set to expire at a specific time are now clearing as expected. The issue self-resolved when the next scheduled expiration job successfully processed the affected users. Scope: Some external users globally were affected. Users on desktop were impacted; the issue was not reproducible internally or across test workspaces. Previous impact to end users: Some users may have experienced custom statuses that did not automatically clear at their scheduled expiration time. Required resolution steps: No steps required. We apologize for any disruptions to your day.

### 4. yandexcloud — Unavailability of dashboards in Monium

- URL: https://status.yandex.cloud/en/incidents/1686
- Started: 2026-04-20 14:34 UTC
- Resolved: 2026-04-20 15:55 UTC
- Duration: 1h 21m
- Impact: major, status: resolved
- Components: Monium
- Canonical incident: [ci-d72c6bc2ce382d2fd92d7425](#ci-d72c6bc2ce382d2fd92d7425)

- `2026-04-20 14:53` **investigating** — We noticed the unavailability of service dashboards in Monium. The service team is aware of the problem and is fixing it.
- `2026-04-20 16:09` **resolved** — At 18.55 (+3), the service dashboards were restored. The service is working correctly.

### 5. digitalocean — App Platform Deployments

- URL: https://stspg.io/0rby3knkk6s1
- Started: 2026-04-23 08:08 UTC
- Resolved: 2026-04-23 10:09 UTC
- Duration: 2h 1m
- Impact: minor, status: resolved
- Components: Global
- Canonical incident: [ci-c5f294960054b9df459f9b11](#ci-c5f294960054b9df459f9b11)

- `2026-04-23 08:08` **investigating** — Our Engineering team is currently investigating reports of build failures on App Platform. Users may experience errors when attempting to build their applications, resulting in failed deployments. Our Engineering team is working to fix the issue and will share an update once we have more information. We apologize for the inconvenience this issue may be causing and appreciate your patience as we work to resolve it.
- `2026-04-23 08:31` **investigating** — Our Engineering team is currently investigating reports of build failures on App Platform. During this time, some users may encounter errors while building their applications, which may result in failed deployments. In addition, we are observing an issue where Kubernetes (DOKS) nodes are being marked as unhealthy by load balancers, which may impact traffic routing for affected services. Our Engineering team is actively working to resolve these issues and will share an update as soon as more information becomes available. We apologize for the inconvenience this may be causing and appreciate your patience.
- `2026-04-23 09:52` **monitoring** — Our Engineering team has implemented a fix to address the issue causing in App Platform deployments and Kubernetes (DOKS) nodes. We are actively monitoring the situation to ensure overall stability. Users may already notice improvements while deploying apps and DOKS nodes. We appreciate your patience throughout the process and will provide a further update once the issue is fully confirmed to be resolved.
- `2026-04-23 10:09` **resolved** — Our Engineering team has confirmed that the issue impacting App Platform deployments and Kubernetes (DOKS) nodes has been fully resolved at 09:22 UTC. Users may already notice improvements while deploying apps and DOKS nodes. All App Platform deployments are now succeeding as expected. Customers who previously encountered build failures should now be able to deploy their applications without further issues. If you continue to experience any problems, please open a ticket with our support team. Thank you for your patience, and we apologize for any inconvenience.

### 6. datadog — Delays in AWS and Azure cloud integration metrics ingestion

- URL: https://stspg.io/xw01xpvggm3n
- Started: 2026-04-24 16:27 UTC
- Resolved: 2026-04-24 16:50 UTC
- Duration: 23m
- Impact: major, status: resolved
- Components: Metrics and Infra Monitoring, Monitors
- Canonical incident: [ci-c5eabf6d3fc4b02ab93b1446](#ci-c5eabf6d3fc4b02ab93b1446)

- `2026-04-24 16:27` **investigating** — We are currently investigating this issue.
- `2026-04-24 16:43` **identified** — The issue has been identified and a fix is being implemented.
- `2026-04-24 16:44` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-04-24 16:50` **resolved** — This incident has been resolved.

## Below threshold (6)

- `digitalocean` [Cloud UI for Managed Kubernetes](https://stspg.io/3tff32qrsv77) — 1h 44m
- `circleci` [Build failures Xcode 26.4.0](https://stspg.io/c4b1vth00hcg) — 1h 11m
- `yandexcloud` [The problem is with the services](https://status.yandex.cloud/en/incidents/1687) — 46m
- `digitalocean` [Intermittent errors impacting some Serverless Inference models in ATL1](https://stspg.io/z4pxt0jh8n97) — 1h 25m
- `digitalocean` [Cloud Control Panel and API](https://stspg.io/2wd37dxsqcrk) — 40m
- `slack` [Users may experience delays due to increased load on bychan keyspace affecting system performance](https://slack-status.com/2026-04/78d9794be1584bae) — 0m

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-6b1ea05d787d7ffece86be09"></a>

- Canonical ID: `ci-6b1ea05d787d7ffece86be09` (1 source records).
  - [yandexcloud/1688](../../data/2026/W17/yandexcloud.json); vendor URL: https://status.yandex.cloud/en/incidents/1688
- Start time basis: `vendor_reported`.

<a id="ci-ce2635506b22350da8a15dc1"></a>

- Canonical ID: `ci-ce2635506b22350da8a15dc1` (1 source records).
  - [slack/1546](../../data/2026/W17/slack.json); vendor URL: https://slack-status.com/2026-04/17580e43d51e0bf2
- Start time basis: `vendor_reported`.

<a id="ci-b14ab11f938e3cc921c51700"></a>

- Canonical ID: `ci-b14ab11f938e3cc921c51700` (1 source records).
  - [slack/1545](../../data/2026/W17/slack.json); vendor URL: https://slack-status.com/2026-04/be1299791e068265
- Start time basis: `vendor_reported`.

<a id="ci-d72c6bc2ce382d2fd92d7425"></a>

- Canonical ID: `ci-d72c6bc2ce382d2fd92d7425` (1 source records).
  - [yandexcloud/1686](../../data/2026/W17/yandexcloud.json); vendor URL: https://status.yandex.cloud/en/incidents/1686
- Start time basis: `vendor_reported`.

<a id="ci-c5f294960054b9df459f9b11"></a>

- Canonical ID: `ci-c5f294960054b9df459f9b11` (1 source records).
  - [digitalocean/2rby4tk1q6mm](../../data/2026/W17/digitalocean.json); vendor URL: https://stspg.io/0rby3knkk6s1
- Start time basis: `vendor_reported`.

<a id="ci-5466c377a9d3e33858402aad"></a>

- Canonical ID: `ci-5466c377a9d3e33858402aad` (1 source records).
  - [digitalocean/p1jbzg2bvhj9](../../data/2026/W17/digitalocean.json); vendor URL: https://stspg.io/3tff32qrsv77
- Start time basis: `vendor_reported`.

<a id="ci-f3ae1aec623f1789a152a0ed"></a>

- Canonical ID: `ci-f3ae1aec623f1789a152a0ed` (1 source records).
  - [circleci/90354ss2tb8t](../../data/2026/W17/circleci.json); vendor URL: https://stspg.io/c4b1vth00hcg
- Start time basis: `vendor_reported`.

<a id="ci-ba73aabc5869223f7a5a323b"></a>

- Canonical ID: `ci-ba73aabc5869223f7a5a323b` (1 source records).
  - [yandexcloud/1687](../../data/2026/W17/yandexcloud.json); vendor URL: https://status.yandex.cloud/en/incidents/1687
- Start time basis: `vendor_reported`.

<a id="ci-26e4cc09a9d3b2758961181f"></a>

- Canonical ID: `ci-26e4cc09a9d3b2758961181f` (1 source records).
  - [digitalocean/1y9gxbdcqb71](../../data/2026/W17/digitalocean.json); vendor URL: https://stspg.io/z4pxt0jh8n97
- Start time basis: `vendor_reported`.

<a id="ci-9ca463de4c007d1645d15bb0"></a>

- Canonical ID: `ci-9ca463de4c007d1645d15bb0` (1 source records).
  - [digitalocean/zjd34dpv45vr](../../data/2026/W17/digitalocean.json); vendor URL: https://stspg.io/2wd37dxsqcrk
- Start time basis: `vendor_reported`.

<a id="ci-c5eabf6d3fc4b02ab93b1446"></a>

- Canonical ID: `ci-c5eabf6d3fc4b02ab93b1446` (1 source records).
  - [datadog/0p99flrd9mln](../../data/2026/W17/datadog.json); vendor URL: https://stspg.io/xw01xpvggm3n
- Start time basis: `vendor_reported`.

<a id="ci-9fce10134f5b307d5c9a1af9"></a>

- Canonical ID: `ci-9fce10134f5b307d5c9a1af9` (1 source records).
  - [slack/1544](../../data/2026/W17/slack.json); vendor URL: https://slack-status.com/2026-04/78d9794be1584bae
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
