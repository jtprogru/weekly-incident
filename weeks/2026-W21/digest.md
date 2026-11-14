# Week 2026-W21

2026-05-18 to 2026-05-24 (UTC). Generated 2026-11-14 00:00.

12 canonical incidents, 4 above threshold. All 14 sources responded.

12 source records (12 saved revisions) → 12 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | yandexcloud | [The problem is with Virtual Private Cloud services](https://status.yandex.cloud/en/incidents/1772) | 2026-05-22 09:21 | 5h 4m | minor | Compute Cloud, Managed Service for Apache Spark™, Managed Service for OpenSearch, Managed Service for Prometheus®, Managed Service for Sharded PostgreSQL, Managed Service for YTsaurus, Managed Service for Apache Kafka®, Managed Service for ClickHouse®, Managed Service for GitLab, Managed Service for Kubernetes®, Managed Service for MySQL®, Managed Service for PostgreSQL, Managed Service for Valkey™, Yandex StoreDoc | 6384 = 304 × 14 × 1.5 |
| 2 | stripe | [SEPA Instant Partial Outage](https://stspg.io/b8k77546cr1x) | 2026-05-19 11:58 | 18h 48m | none | Acquirers and payment methods | 1354 = 1128 × 1 × 1.2 |
| 3 | digitalocean | [Block Storage Volume Performance](https://stspg.io/1cr62wpw64kx) | 2026-05-20 01:05 | 4h 8m | minor | NYC3 | 198 = 248 × 1 × 0.8 |
| 4 | digitalocean | [Cloud Firewall](https://stspg.io/6y9xmwmgkyy4) | 2026-05-20 05:35 | 2h 12m | minor | Cloud Firewall | 106 = 132 × 1 × 0.8 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. yandexcloud — The problem is with Virtual Private Cloud services

- URL: https://status.yandex.cloud/en/incidents/1772
- Started: 2026-05-22 09:21 UTC
- Resolved: 2026-05-22 14:26 UTC
- Duration: 5h 4m
- Impact: minor, status: resolved
- Components: Compute Cloud, Managed Service for Apache Spark™, Managed Service for OpenSearch, Managed Service for Prometheus®, Managed Service for Sharded PostgreSQL, Managed Service for YTsaurus, Managed Service for Apache Kafka®, Managed Service for ClickHouse®, Managed Service for GitLab, Managed Service for Kubernetes®, Managed Service for MySQL®, Managed Service for PostgreSQL, Managed Service for Valkey™, Yandex StoreDoc
- Canonical incident: [ci-e36d5ddf780b5c37400538ff](#ci-e36d5ddf780b5c37400538ff)

- `2026-05-22 09:22` **investigating** — We are experiencing delays in processing operations for creating and managing virtual machines. This may affect the creation time of resources in Compute Cloud, Managed Service for Kubernetes, as well as managed databases (MySQL, PostgreSQL, ClickHouse, and others).
- `2026-05-22 11:06` **unknown** — Due to ongoing restoration work, operations to create and manage resources in the Compute Cloud service are temporarily unavailable. This makes it impossible to create and modify resources in dependent services: Managed Service for Kubernetes, Managed Service for MySQL, Managed Service for PostgreSQL and other managed databases. Existing previously created resources continue to work.
- `2026-05-22 11:59` **unknown** — We have restored access to operations. Currently, there may be slight delays in processing.
- `2026-05-22 14:37` **resolved** — The services have been fully restored.

### 2. stripe — SEPA Instant Partial Outage

- URL: https://stspg.io/b8k77546cr1x
- Started: 2026-05-19 11:58 UTC
- Resolved: 2026-05-20 06:47 UTC
- Duration: 18h 48m
- Impact: none, status: resolved
- Components: Acquirers and payment methods
- Canonical incident: [ci-f2018f3bb14a96bcbf9e6a4d](#ci-f2018f3bb14a96bcbf9e6a4d)

- `2026-05-19 11:58` **investigating** — Instant payouts in EUR to bank accounts and Bank Transfers payments are experiencing delays and some might fail due to a broader industry issue affecting SEPA Instant. All other payouts, including instant payouts to debit cards, are unaffected. We're investigating and will provide an update once resolved.
- `2026-05-19 15:14` **identified** — Instant payouts in EUR to bank accounts and Bank Transfers payments continue to experience delays due broader industry issues affecting SEPA Instant. All other payouts, including instant payouts to debit cards, remain unaffected. We will provide further updates as soon as possible.
- `2026-05-19 21:15` **identified** — Instant payouts to bank accounts and Bank Transfers payments in Europe continue to experience delays due to broader issues affecting SEPA Instant. This issue is also resulting in elevated errors on Pay by Bank payments beginning at 08:40 UTC. All other payouts, including instant payouts to debit cards, remain unaffected. We are continuing to engage with our upstream partners, and will provide further updates as soon as they are available.
- `2026-05-20 06:47` **resolved** — Instant payouts in EUR to bank accounts, Pay by Bank payments, and Bank Transfers payments in Europe are now working successfully. There was no impact to other payouts, including instant payouts to debit cards during this time.

### 3. digitalocean — Block Storage Volume Performance

- URL: https://stspg.io/1cr62wpw64kx
- Started: 2026-05-20 01:05 UTC
- Resolved: 2026-05-20 05:13 UTC
- Duration: 4h 8m
- Impact: minor, status: resolved
- Components: NYC3
- Canonical incident: [ci-b7783513356ae1b1a5d398cd](#ci-b7783513356ae1b1a5d398cd)

- `2026-05-20 01:05` **investigating** — Our Engineering team is currently investigating an ongoing infrastructure issue in the NYC3 region affecting Block Storage Volumes at the storage layer. Users may experience degraded write performance and intermittent impact to services dependent on the affected storage infrastructure. We apologize for the inconvenience and will continue to provide updates as more information becomes available.
- `2026-05-20 02:19` **monitoring** — Our Engineering team has implemented mitigation measures for the infrastructure issue affecting Block Storage Volumes in the NYC3 region. The team is monitoring the situation, and we will share another update once the issue is fully resolved
- `2026-05-20 05:13` **resolved** — From 00:00 UTC to 01:40 UTC, users may have experienced degraded write performance and intermittent impacts to workloads dependent on Block Storage Volumes in the NYC3 region. Our Engineering team has confirmed that the underlying infrastructure issue affecting Block Storage Volumes at the storage layer has been fully resolved, and services are now operating normally. If you continue to experience any issues, please contact our Support team by opening a ticket. We apologize for any inconvenience caused.

### 4. digitalocean — Cloud Firewall

- URL: https://stspg.io/6y9xmwmgkyy4
- Started: 2026-05-20 05:35 UTC
- Resolved: 2026-05-20 07:48 UTC
- Duration: 2h 12m
- Impact: minor, status: resolved
- Components: Cloud Firewall
- Canonical incident: [ci-1cd9e93d912d021099281323](#ci-1cd9e93d912d021099281323)

- `2026-05-20 05:35` **investigating** — We are currently investigating an issue with incorrect Cloud Firewall rules appearing on the cloud panel. Our engineering team is aware of the situation and actively working to identify the root cause. Since this is a user interface issue, users should not experience any problems with other services or networking. We apologize for the inconvenience and appreciate your patience. We will continue to provide updates as we learn more.
- `2026-05-20 07:15` **monitoring** — Our Engineering team has implemented a fix to resolve the issue causing incorrect Cloud Firewall rules to display on the cloud panel. At this time, the user interface is functioning as expected, and we are actively monitoring the situation to ensure continued stability. We will provide a final update once we have verified that the issue is fully resolved.
- `2026-05-20 07:48` **resolved** — Our Engineering team has verified that the issue causing incorrect Cloud Firewall rules to display on the cloud panel is fully resolved. The user interface continues to function as expected, and all systems have returned to normal operating conditions. If you continue to experience problems, please open a ticket with our support team. Thank you for your patience throughout this incident.

## Below threshold (8)

- `yandexcloud` [The problem of creating resources in ru-central1-A.](https://status.yandex.cloud/en/incidents/1763) — 33m
- `yandexcloud` [The problem is with the services](https://status.yandex.cloud/en/incidents/1775) — 1h 38m
- `slack` [Trouble uploading files and emojis in Slack](https://slack-status.com/2026-05/60ba1e616cf353be) — 1h 36m
- `digitalocean` [Cloud Firewall](https://stspg.io/5dzt41b39qvv) — 54m
- `yandexcloud` [The problem is with the Yandex Managed Service for Trino](https://status.yandex.cloud/en/incidents/1770) — 27m
- `yandexcloud` [The problem is with Compute Cloud services](https://status.yandex.cloud/en/incidents/1756) — 21m
- `stripe` [Elevated Bancontact errors](https://stspg.io/4pv6q73763bm) — 15m
- `slack` [Huddle AI summaries not being generated](https://slack-status.com/2026-05/8ebfa96c50e29550) — 0m

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-e36d5ddf780b5c37400538ff"></a>

- Canonical ID: `ci-e36d5ddf780b5c37400538ff` (1 source records).
  - [yandexcloud/1772](../../data/2026/W21/yandexcloud.json); vendor URL: https://status.yandex.cloud/en/incidents/1772
- Start time basis: `vendor_reported`.

<a id="ci-f2018f3bb14a96bcbf9e6a4d"></a>

- Canonical ID: `ci-f2018f3bb14a96bcbf9e6a4d` (1 source records).
  - [stripe/0plg9bpr660h](../../data/2026/W21/stripe.json); vendor URL: https://stspg.io/b8k77546cr1x
- Start time basis: `vendor_reported`.

<a id="ci-f89a6a62b00db8c42e565771"></a>

- Canonical ID: `ci-f89a6a62b00db8c42e565771` (1 source records).
  - [yandexcloud/1763](../../data/2026/W21/yandexcloud.json); vendor URL: https://status.yandex.cloud/en/incidents/1763
- Start time basis: `vendor_reported`.

<a id="ci-b7783513356ae1b1a5d398cd"></a>

- Canonical ID: `ci-b7783513356ae1b1a5d398cd` (1 source records).
  - [digitalocean/spgd2scvttth](../../data/2026/W21/digitalocean.json); vendor URL: https://stspg.io/1cr62wpw64kx
- Start time basis: `vendor_reported`.

<a id="ci-4d5e59853c2faedc083716c8"></a>

- Canonical ID: `ci-4d5e59853c2faedc083716c8` (1 source records).
  - [yandexcloud/1775](../../data/2026/W21/yandexcloud.json); vendor URL: https://status.yandex.cloud/en/incidents/1775
- Start time basis: `vendor_reported`.

<a id="ci-0e4eb593a9dd8036d626137a"></a>

- Canonical ID: `ci-0e4eb593a9dd8036d626137a` (1 source records).
  - [slack/1555](../../data/2026/W21/slack.json); vendor URL: https://slack-status.com/2026-05/60ba1e616cf353be
- Start time basis: `vendor_reported`.

<a id="ci-1cd9e93d912d021099281323"></a>

- Canonical ID: `ci-1cd9e93d912d021099281323` (1 source records).
  - [digitalocean/sv0tfb6jhqz8](../../data/2026/W21/digitalocean.json); vendor URL: https://stspg.io/6y9xmwmgkyy4
- Start time basis: `vendor_reported`.

<a id="ci-490495e3d5f970ac7c592049"></a>

- Canonical ID: `ci-490495e3d5f970ac7c592049` (1 source records).
  - [digitalocean/55wh1mbv7vbs](../../data/2026/W21/digitalocean.json); vendor URL: https://stspg.io/5dzt41b39qvv
- Start time basis: `vendor_reported`.

<a id="ci-4130191691ffbf9a23f307d7"></a>

- Canonical ID: `ci-4130191691ffbf9a23f307d7` (1 source records).
  - [yandexcloud/1770](../../data/2026/W21/yandexcloud.json); vendor URL: https://status.yandex.cloud/en/incidents/1770
- Start time basis: `vendor_reported`.

<a id="ci-5488cef26581870f9e45b710"></a>

- Canonical ID: `ci-5488cef26581870f9e45b710` (1 source records).
  - [yandexcloud/1756](../../data/2026/W21/yandexcloud.json); vendor URL: https://status.yandex.cloud/en/incidents/1756
- Start time basis: `vendor_reported`.

<a id="ci-47ad279eb28aa27ed2c8f06c"></a>

- Canonical ID: `ci-47ad279eb28aa27ed2c8f06c` (1 source records).
  - [stripe/wjg95ch7scvp](../../data/2026/W21/stripe.json); vendor URL: https://stspg.io/4pv6q73763bm
- Start time basis: `vendor_reported`.

<a id="ci-6deca55495b33c5df6e374c5"></a>

- Canonical ID: `ci-6deca55495b33c5df6e374c5` (1 source records).
  - [slack/1554](../../data/2026/W21/slack.json); vendor URL: https://slack-status.com/2026-05/8ebfa96c50e29550
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
| slack | ok | 32 | 0 |
| stripe | ok | 50 | 0 |
| yandexcloud | ok | 50 | 0 |
