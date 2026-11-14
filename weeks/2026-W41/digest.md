# Week 2026-W41

2026-10-05 to 2026-10-11 (UTC). Generated 2026-11-14 12:36.

63 canonical incidents, 39 above threshold. 2 of 14 sources unavailable: datadog, yandexcloud.

63 source records (68 saved revisions) → 63 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | yandexcloud | [Unavailability of Yandex Cloud zones](https://status.yandex.cloud/en/incidents/2136) | 2026-10-11 01:00 | 34d 11h 36m (ongoing) | major | API Gateway, Application Load Balancer, Audit Trails, Certificate Manager, Cloud Apps, Cloud Backup, Cloud DNS, Cloud Desktop, Cloud Functions, Cloud Interconnect, Cloud Logging, Cloud Notification Service, Cloud Registry, Compute Cloud, Container Registry, Data Processing, Data Streams, Data Transfer, DataSphere, Forms, GOST Gateway, Identity and Access Management, IoT Core, Key Management Service, Managed Service for Apache Spark™, Managed Service for OpenSearch, Managed Service for Prometheus®, Managed Service for Sharded PostgreSQL, Managed Service for YDB, Managed Service for YTsaurus, Managed Service for Apache Kafka®, Managed Service for ClickHouse®, Managed Service for GitLab, Managed Service for Kubernetes®, Managed Service for MySQL®, Managed Service for PostgreSQL, Managed Service for Valkey™, Marketplace, Message Queue, Monitoring, Monium, Network Load Balancer, Object Storage, Resource Manager, Security Deck, Serverless Containers, Smart Web Security, SmartCaptcha, SourceCraft, SourceCraft Code Assistant, SpeechKit, SpeechSense, Stackland, Support Center, Tracker, Translate, VibeCraft, Virtual Private Cloud, Vision OCR, Wiki, Yandex AI Studio, Yandex BareMetal, Yandex Cloud Billing, Yandex Cloud Console, Yandex Cloud Detection and Response, Yandex Cloud Postbox, Yandex Cloud Quota Manager, Yandex Cloud Router, Yandex Cloud Video, Yandex Identity Hub, Yandex Lockbox, Yandex MPP Analytics Engine for PostgreSQL, Yandex Managed Service for Apache Airflow®, Yandex Managed Service for Trino, Yandex MetaData Hub, Yandex Neurosupport, Yandex Query, Yandex SIEM, Yandex Search API, Yandex StoreDoc, Yandex WebSQL, YandexGPT API | 6107688 = 49656 × 82 × 1.5 |
| 2 | yandexcloud | [Power failure in the ru-central-b availability zone](https://status.yandex.cloud/en/incidents/2092) | 2026-10-07 22:31 | 2d 18h 42m | major | API Gateway, Application Load Balancer, Audit Trails, Certificate Manager, Cloud Apps, Cloud Backup, Cloud CDN, Cloud DNS, Cloud Desktop, Cloud Functions, Cloud Interconnect, Cloud Logging, Cloud Notification Service, Cloud Registry, Compute Cloud, Container Registry, Data Processing, Data Streams, Data Transfer, DataLens, DataSphere, Forms, GOST Gateway, Identity and Access Management, IoT Core, Key Management Service, Managed Service for Apache Spark™, Managed Service for OpenSearch, Managed Service for Prometheus®, Managed Service for Sharded PostgreSQL, Managed Service for YDB, Managed Service for YTsaurus, Managed Service for Apache Kafka®, Managed Service for ClickHouse®, Managed Service for GitLab, Managed Service for Kubernetes®, Managed Service for MySQL®, Managed Service for PostgreSQL, Managed Service for Valkey™, Marketplace, Message Queue, Monitoring, Monium, Network Load Balancer, Object Storage, Resource Manager, Security Deck, Serverless Containers, Smart Web Security, SmartCaptcha, SourceCraft, SourceCraft Code Assistant, SpeechKit, SpeechSense, Stackland, Support Center, Tracker, Translate, VibeCraft, Virtual Private Cloud, Vision OCR, Wiki, Yandex AI Studio, Yandex BareMetal, Yandex Cloud Billing, Yandex Cloud Console, Yandex Cloud Detection and Response, Yandex Cloud Postbox, Yandex Cloud Quota Manager, Yandex Cloud Router, Yandex Cloud Video, Yandex Identity Hub, Yandex Lockbox, Yandex MPP Analytics Engine for PostgreSQL, Yandex Managed Service for Apache Airflow®, Yandex Managed Service for Trino, Yandex MetaData Hub, Yandex Neurosupport, Yandex Query, Yandex SIEM, Yandex Search API, Yandex StoreDoc, Yandex WebSQL, YandexGPT API | 504252 = 4002 × 84 × 1.5 |
| 3 | cloudflare | [Missing Logs and Traces in Cloudflare Observability Platform](https://www.cloudflarestatus.com/incidents/pfvfug94jq4q) | 2026-10-05 13:00 | 18d 8h 0m | none | — | 47520 = 26400 × 1 × 1.8 |
| 4 | grafana | [IRM Unable to Send SMS to Thailand Numbers](https://stspg.io/9l8htmqt82s2) | 2026-10-05 16:17 | 13d 19h 22m | minor | Incident Management and Response (IRM) | 19882 = 19882 × 1 × 1.0 |
| 5 | cloudflare | [Elevated Errors in China Network Colos](https://www.cloudflarestatus.com/incidents/cyxclg7nnhjy) | 2026-10-10 04:13 | 5d 5h 15m | minor | China Network | 13527 = 7515 × 1 × 1.8 |
| 6 | anthropic | [Elevated errors on platform.claude.com](https://stspg.io/sjkw7njwf0mt) | 2026-10-07 13:25 | 6d 4h 35m | major | Claude Console (platform.claude.com) | 13372 = 8915 × 1 × 1.5 |
| 7 | openai | [Delayed Costs data in the Compliance API](https://status.openai.com/incidents/01M4GV1M2A3K38T73JCVCD9G0R) | 2026-10-09 17:23 | 2d 22h 39m | minor | — | 6358 = 4239 × 1 × 1.5 |
| 8 | slack | [Slack for iOS closing unexpectedly when opening channels](https://slack-status.com/2026-10/a0cfb0208a1d2032) | 2026-10-09 15:35 | 3d 1h 28m | unknown | Connectivity | 5290 = 4408 × 1 × 1.2 |
| 9 | gcp | [Google Cloud Storage (GCS) customers experienced elevated latencies and error rates when performing upload or write operations.](https://status.cloud.google.com/incidents/mSFVt6PiP557bjRyYsut) | 2026-10-08 11:00 | 1d 5h 46m | major | Google Cloud Storage | 3572 = 1786 × 1 × 2.0 |
| 10 | grafana | [Degraded Performance with Increased Latency for Grafana Cloud Metrics and Traces in GCP US Central](https://stspg.io/cjbf6knm4jq4) | 2026-10-08 13:06 | 11h 33m | none | GCP US Central - prod-us-central-0, GCP US Central - prod-us-central-0: Ingestion, GCP US Central - prod-us-central-0: Ingestion, GCP US Central - prod-us-central-5, GCP US Central - prod-us-central-5: Ingestion | 3465 = 693 × 5 × 1.0 |
| 11 | cloudflare | [Durable Objects Availability Issue in Eastern North America](https://www.cloudflarestatus.com/incidents/z28czd65h29l) | 2026-10-05 17:14 | 5h 21m | minor | Containers, Durable Objects, R2, Workers Assets, Workflows | 2889 = 321 × 5 × 1.8 |
| 12 | cloudflare | [Network Performance Issues in Columbus](https://www.cloudflarestatus.com/incidents/tt8f4ttcuonf) | 2026-10-07 02:48 | 9h 2m | major | Columbus, OH, United States - (CMH), Network | 1951 = 542 × 2 × 1.8 |
| 13 | circleci | [Delay in data appearing in UI and API](https://stspg.io/hkfkwz292bz8) | 2026-10-09 11:33 | 14h 35m | major | CircleCI API, CircleCI UI | 1750 = 875 × 2 × 1.0 |
| 14 | openai | [Elevated Work Mode errors](https://status.openai.com/incidents/01M45E0JFFJZ6EQ3BWYPW74R2Z) | 2026-10-05 07:03 | 19h 6m | minor | — | 1719 = 1146 × 1 × 1.5 |
| 15 | cloudflare | [Cloudflare DNS Resolver 1.1.1.1 and Gateway DNS are having issues in SIN](https://www.cloudflarestatus.com/incidents/4knp3pzs3qvx) | 2026-10-09 14:41 | 4h 49m | minor | Gateway, Recursive DNS, Singapore, Singapore - (SIN) | 1561 = 289 × 3 × 1.8 |
| 16 | yandexcloud | [Increasing timings in the Object Storage service](https://status.yandex.cloud/en/incidents/2130) | 2026-10-08 19:30 | 14h 20m | minor | Object Storage | 1290 = 860 × 1 × 1.5 |
| 17 | cloudflare | [API Shield JWT Validation Errors](https://www.cloudflarestatus.com/incidents/ts295zqycqn7) | 2026-10-05 10:55 | 11h 43m | minor | API Shield | 1265 = 703 × 1 × 1.8 |
| 18 | cloudflare | [Cloudflare API service issues](https://www.cloudflarestatus.com/incidents/x8n12150xwf6) | 2026-10-09 18:52 | 7h 19m | minor | API | 790 = 439 × 1 × 1.8 |
| 19 | github | [Several services are degraded](https://stspg.io/w3rcssslq22w) | 2026-10-06 19:57 | 3h 43m | critical | Pull Requests, Webhooks | 669 = 223 × 2 × 1.5 |
| 20 | github | [Incident with Actions](https://stspg.io/c11dc9nb1zdq) | 2026-10-05 19:11 | 3h 37m | critical | Actions, Pages | 651 = 217 × 2 × 1.5 |
| 21 | circleci | [Delays in pipelines, workflows, UI data, and notifications](https://stspg.io/s7tpg47srxdd) | 2026-10-06 16:00 | 5h 18m | major | Notifications & Status Updates, Pipelines & Workflows | 636 = 318 × 2 × 1.0 |
| 22 | openai | [Elevated errors in ChatGPT conversations](https://status.openai.com/incidents/01M46G248GGHV9RSAJE82BZY1X) | 2026-10-05 16:58 | 7h 2m | minor | — | 633 = 422 × 1 × 1.5 |
| 23 | openai | [Unable to create new ChatGPT Work threads in latest version of desktop app](https://status.openai.com/incidents/01M4ADHMMBK7K92NJFJ0DXRYJM) | 2026-10-07 05:31 | 6h 21m | major | — | 572 = 381 × 1 × 1.5 |
| 24 | openai | [Elevated error rates in Codex Cloud](https://status.openai.com/incidents/01M46KED6KQ6PR195KZYYVNY8W) | 2026-10-05 17:58 | 6h 2m | none | — | 543 = 362 × 1 × 1.5 |
| 25 | openai | [Admin Console Intermittently Inaccessible](https://status.openai.com/incidents/01M46C47WYYD9NMZVGH83X3ZR1) | 2026-10-05 15:50 | 5h 17m | none | — | 476 = 317 × 1 × 1.5 |
| 26 | stripe | [Elevated Swish errors](https://stspg.io/m7j171ppknbs) | 2026-10-06 00:30 | 6h 33m | none | Acquirers and payment methods | 472 = 393 × 1 × 1.2 |
| 27 | cloudflare | [Network Performance Issues in Kochi](https://www.cloudflarestatus.com/incidents/xczvsioz1un8) | 2026-10-08 14:00 | 3h 0m | minor | — | 324 = 180 × 1 × 1.8 |
| 28 | openai | [Some ChatGPT content may be unavailable on Android](https://status.openai.com/incidents/01M4H7FKXDA425TDNP2CRGEYA3) | 2026-10-09 21:00 | 3h 26m | major | — | 309 = 206 × 1 × 1.5 |
| 29 | github | [Incident with Git Operations, Issues, Actions and Pull Requests](https://stspg.io/4p8x0jnnsx7m) | 2026-10-07 17:17 | 47m | critical | Actions, Git Operations, Issues, Pull Requests | 282 = 47 × 4 × 1.5 |
| 30 | yandexcloud | [The problem is with the Cloud Functions and Serverless Containers services](https://status.yandex.cloud/en/incidents/2124) | 2026-10-08 11:00 | 1h 25m | major | Cloud Functions, Serverless Containers | 255 = 85 × 2 × 1.5 |
| 31 | openai | [Computer Use Is Experiencing Issues](https://status.openai.com/incidents/01M493CZ0CB14KPC0EER83MNVB) | 2026-10-06 17:15 | 2h 45m | minor | — | 248 = 165 × 1 × 1.5 |
| 32 | circleci | [Elevated wait times for machine jobs](https://stspg.io/vcsrgfwdsy4z) | 2026-10-08 13:00 | 3h 35m | minor | Machine Jobs | 215 = 215 × 1 × 1.0 |
| 33 | openai | [Errors processing large PDFs in the Responses API](https://status.openai.com/incidents/01M48Y8E81QAFF9XKVGE118E6X) | 2026-10-06 15:45 | 2h 4m | minor | — | 186 = 124 × 1 × 1.5 |
| 34 | github | [Disruption with some GitHub services](https://stspg.io/pjcfc4c3pm7g) | 2026-10-05 23:47 | 1h 44m | major | — | 156 = 104 × 1 × 1.5 |
| 35 | openai | [Workspace admins may be unable to manage plugins](https://status.openai.com/incidents/01M4HA7BCCXVHTR5DCYMSBQF2H) | 2026-10-09 21:48 | 1h 38m | major | — | 147 = 98 × 1 × 1.5 |
| 36 | anthropic | [Elevated errors loading usage data on platform.claude.com](https://stspg.io/scdhhnplc2jn) | 2026-10-07 01:20 | 1h 13m | major | Claude Console (platform.claude.com) | 110 = 73 × 1 × 1.5 |
| 37 | anthropic | [Elevated errors for Claude Opus 5.5](https://stspg.io/bzc81gbv1b99) | 2026-10-06 12:24 | 18m | major | Claude API (api.anthropic.com), Claude Code, Claude Cowork, claude.ai | 108 = 18 × 4 × 1.5 |
| 38 | anthropic | [Elevated errors for Claude Mythos 5.1 and Claude Fable 5.1](https://stspg.io/wdqhxt8xt574) | 2026-10-05 13:05 | 15m | major | Claude API (api.anthropic.com), Claude Code, Claude Cowork, claude.ai | 90 = 15 × 4 × 1.5 |
| 39 | github | [Incident with Git Operations, Pull Requests and Actions](https://stspg.io/96smrcth8bpg) | 2026-10-07 16:25 | 0m | critical | Actions, Git Operations, Issues, Pull Requests, Webhooks | 0 = 0 × 5 × 1.5 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. yandexcloud — Unavailability of Yandex Cloud zones

- URL: https://status.yandex.cloud/en/incidents/2136
- Started: 2026-10-11 01:00 UTC
- Resolved: not yet
- Duration: 34d 11h 36m (ongoing)
- Impact: major, status: investigating
- Components: API Gateway, Application Load Balancer, Audit Trails, Certificate Manager, Cloud Apps, Cloud Backup, Cloud DNS, Cloud Desktop, Cloud Functions, Cloud Interconnect, Cloud Logging, Cloud Notification Service, Cloud Registry, Compute Cloud, Container Registry, Data Processing, Data Streams, Data Transfer, DataSphere, Forms, GOST Gateway, Identity and Access Management, IoT Core, Key Management Service, Managed Service for Apache Spark™, Managed Service for OpenSearch, Managed Service for Prometheus®, Managed Service for Sharded PostgreSQL, Managed Service for YDB, Managed Service for YTsaurus, Managed Service for Apache Kafka®, Managed Service for ClickHouse®, Managed Service for GitLab, Managed Service for Kubernetes®, Managed Service for MySQL®, Managed Service for PostgreSQL, Managed Service for Valkey™, Marketplace, Message Queue, Monitoring, Monium, Network Load Balancer, Object Storage, Resource Manager, Security Deck, Serverless Containers, Smart Web Security, SmartCaptcha, SourceCraft, SourceCraft Code Assistant, SpeechKit, SpeechSense, Stackland, Support Center, Tracker, Translate, VibeCraft, Virtual Private Cloud, Vision OCR, Wiki, Yandex AI Studio, Yandex BareMetal, Yandex Cloud Billing, Yandex Cloud Console, Yandex Cloud Detection and Response, Yandex Cloud Postbox, Yandex Cloud Quota Manager, Yandex Cloud Router, Yandex Cloud Video, Yandex Identity Hub, Yandex Lockbox, Yandex MPP Analytics Engine for PostgreSQL, Yandex Managed Service for Apache Airflow®, Yandex Managed Service for Trino, Yandex MetaData Hub, Yandex Neurosupport, Yandex Query, Yandex SIEM, Yandex Search API, Yandex StoreDoc, Yandex WebSQL, YandexGPT API
- Canonical incident: [ci-d5585cf81abf4edcfde48cfd](#ci-d5585cf81abf4edcfde48cfd)

- `2026-10-11 01:11` **investigating** — The Yandex Cloud platform is still operating in a limited mode. Our specialists continue restoration work around the clock. The technical support team receives an increased number of requests, so the response time can be increased. Our recommendations continue to be relevant. We recommend using alternative sites for hosting and reserving resources. The next update will be published on October 12 at 19:00 Moscow time.
- `2026-10-11 03:08` **unknown** — As a result of the UAV attack, the infrastructure of Yandex’s data center in Vladimir was damaged. The data center’s operation has been completely suspended. There are no casualties; the relevant services are already on site and dealing with the consequences. We consider the remaining resource configuration to be unstable. The platform’s operating mode is emergency — we recommend using any other alternative disaster recovery plans. Communication channels with technical support may be unstable. The Yandex Cloud team is working to restore service operations.
- `2026-10-11 03:11` **unknown** — We are sharing the status of individual components and services as of 6:10: • Yandex Cloud Console and API are available; • Compute: partially available; • MDB: most HA database clusters have switched over; some may operate in read‑only mode or be unavailable. • S3: available for reading, potential issues with writing; • Managed Service for Kubernetes: we are assessing the status of regional clusters; • Yandex Identity and Access Management (IAM) is available; • Other services may experience disruptions – we will provide updates on their status later. Communication channels with technical support may be unstable. Information on the status board appears with a delay. We recommend using alternative resource recovery plans on external platforms.
- `2026-10-11 04:13` **unknown** — Updates on the operation of services and components as of 07:13: • Yandex Cloud Console and API are available; • Compute: partially available; • MDB: most HA database clusters have switched over; some may operate in read‑only mode or be unavailable. • S3: available for reading, write issues may be observed. • Managed Service for Kubernetes: partially available; • Yandex Identity and Access Management (IAM) is available; • Yandex Monium is unavailable. The team is working on restoring technical support communication channels. We recommend using alternative sites to restore resources.
- `2026-10-11 05:55` **unknown** — The team continues to work on restoring and stabilizing key services. We’re sharing an update on the service operations: Yandex Cloud Console and API are available; Compute: partially available; MDB: most HA database clusters have switched over; some may operate in read‑only mode or be unavailable. S3: available for reading, but there may be issues with writing; Managed Service for Kubernetes: available, but issues with individual clusters may persist; Yandex Monium: unavailable; Yandex Container Registry: operational; Yandex AI Studio: services are partially available, more details about the status. Security services are partially available, more details about the status. The Yandex Cloud team is working to resume the operation of technical support communication channels. The main recommendation remains the same: please use alternative platforms to restore resources.
- `2026-10-11 08:14` **unknown** — The current status of individual platform services: Yandex Cloud Console and API are available; Compute: partially available; Managed Databases: the service is available, but difficulties persist with Managed Service for Apache Kafka®. S3: available for reading and writing, with possible slowdown in operation speed. Managed Service for Kubernetes: available, but issues with individual clusters may persist; Yandex Monium is unavailable; Yandex Container Registry is operational; Yandex AI Studio: services are partially available, more details about the status. Security services are partially available, more details about the status. The operation of the technical support channels has been partially restored; we are working to speed up the processing of inquiries, but there may still be significant delays in responses. We have also temporarily suspended the issuance of quotas for additional cloud resources. The next update will be released on October 11 at 13:00 Moscow time.
- `2026-10-11 10:12` **unknown** — Service availability remains unchanged. We continue working around the clock to restore services and platform components. Technical support channels have been partially restored. The Yandex Cloud platform team is working to process support requests faster, but responses may still be significantly delayed. As previously announced, quota increases for additional cloud resources remain suspended. We will provide the next update on October 11 at 18:00 Moscow time (UTC+3).
- `2026-10-11 15:01` **unknown** — The Yandex Cloud platform continues to operate with limited functionality. Our technical teams are restoring services, while our support team is handling requests around the clock. We continue to recommend considering alternative platforms to deploy your resources more quickly and provide redundancy. We will provide the next update on October 12 at 11:00 Moscow time (UTC+3).
- `2026-10-12 08:17` **unknown** — The Yandex Cloud platform is still operating in a limited mode. Our specialists continue restoration work around the clock. The technical support team receives an increased number of requests, so the response time can be increased. Our recommendations continue to be relevant. We recommend using alternative sites for hosting and reserving resources. The next update will be published on October 12 at 19:00 Moscow time.
- `2026-10-12 16:16` **unknown** — The Yandex Cloud platform operates in a limited mode. Technical support receives an increased number of requests, so the response time can be increased. We collected answers to the most common questions and recommendations for restoring work, which will help you navigate the available options. The next update will be published on October 13 at 11:00 Moscow time.
- `2026-10-13 09:05` **unknown** — The Yandex Cloud platform continues to operate in a limited capacity. Our engineers are continuing recovery efforts. We understand that your teams are currently focused on restoring their services. To ensure you do not have to worry about additional cloud costs during this time, we have decided to make the use of all Yandex Cloud services across all availability zones in the Russian Federation free of charge for existing customers from October 8 through October 31, inclusive. Your October invoice will reflect only the costs for resources used up to and including October 7. Charges for the period starting October 8 may still appear in your billing account, but we will adjust them automatically. No action is required on your part. We will publish the next update on October 13 at 19:00 Moscow time.
- `2026-10-13 16:44` **unknown** — The status of the cloud platform remains unchanged; the Yandex Cloud team continues to restore services. We have updated the statuses and recommendations for specific services—we do this daily on a dedicated page.
- `2026-10-14 16:23` **unknown** — The Yandex Cloud platform is still operating in a limited mode. Our specialists continue the restoration work around the clock. Recommendations for restoring services are updated daily on a separate page: https://yandex.cloud/en/recommendations
- `2026-10-15 17:04` **unknown** — Yandex’s cloud platform continues to operate with limitations. Engineers are working around the clock to restore service functionality. We update all currently available product‑related recommendations on a daily basis.
- `2026-10-16 17:19` **unknown** — The Yandex Cloud platform operates in a limited mode. Our specialists are carrying out restoration work. Users can keep up-to-date information about the availability of individual services in the management console. We also regularly update recommendations for restoring services.

### 2. yandexcloud — Power failure in the ru-central-b availability zone

- URL: https://status.yandex.cloud/en/incidents/2092
- Started: 2026-10-07 22:31 UTC
- Resolved: 2026-10-10 17:13 UTC
- Duration: 2d 18h 42m
- Impact: major, status: resolved
- Components: API Gateway, Application Load Balancer, Audit Trails, Certificate Manager, Cloud Apps, Cloud Backup, Cloud CDN, Cloud DNS, Cloud Desktop, Cloud Functions, Cloud Interconnect, Cloud Logging, Cloud Notification Service, Cloud Registry, Compute Cloud, Container Registry, Data Processing, Data Streams, Data Transfer, DataLens, DataSphere, Forms, GOST Gateway, Identity and Access Management, IoT Core, Key Management Service, Managed Service for Apache Spark™, Managed Service for OpenSearch, Managed Service for Prometheus®, Managed Service for Sharded PostgreSQL, Managed Service for YDB, Managed Service for YTsaurus, Managed Service for Apache Kafka®, Managed Service for ClickHouse®, Managed Service for GitLab, Managed Service for Kubernetes®, Managed Service for MySQL®, Managed Service for PostgreSQL, Managed Service for Valkey™, Marketplace, Message Queue, Monitoring, Monium, Network Load Balancer, Object Storage, Resource Manager, Security Deck, Serverless Containers, Smart Web Security, SmartCaptcha, SourceCraft, SourceCraft Code Assistant, SpeechKit, SpeechSense, Stackland, Support Center, Tracker, Translate, VibeCraft, Virtual Private Cloud, Vision OCR, Wiki, Yandex AI Studio, Yandex BareMetal, Yandex Cloud Billing, Yandex Cloud Console, Yandex Cloud Detection and Response, Yandex Cloud Postbox, Yandex Cloud Quota Manager, Yandex Cloud Router, Yandex Cloud Video, Yandex Identity Hub, Yandex Lockbox, Yandex MPP Analytics Engine for PostgreSQL, Yandex Managed Service for Apache Airflow®, Yandex Managed Service for Trino, Yandex MetaData Hub, Yandex Neurosupport, Yandex Query, Yandex SIEM, Yandex Search API, Yandex StoreDoc, Yandex WebSQL, YandexGPT API
- Canonical incident: [ci-8e0bde59a25eb3bfe75e60a3](#ci-8e0bde59a25eb3bfe75e60a3)

- `2026-10-07 22:31` **investigating** — Power outages are observed in the ru-central1-b availability zone. The team is working to fix the problem. If possible, we recommend shifting the load to other availability zones. We will keep you informed of the progress of the work.
- `2026-10-07 23:19` **unknown** — We have taken measures to localize the problem in the ru-central1-b availability zone. The situation has not stabilized yet — the team continues to find out the reasons and work on recovery. We will share updated information in the near future.
- `2026-10-08 01:16` **unknown** — The Yandex Cloud team continues to work on eliminating the consequences of the incident. The recovery dates are being clarified — we will inform them as soon as confirmed information becomes available.
- `2026-10-08 01:55` **unknown** — Due to the incident in the ru-central1-b availability zone, problems may arise in other zones when creating new VMs, k8s clusters, and databases. We are continuing to address the consequences of the incident and will report on the progress of the recovery.
- `2026-10-08 05:26` **unknown** — The unavailability of the ru-central1-b zone remains for now, we have completed all the necessary work to switch the load to other availability zones. All services in the remaining zones are operating normally. Currently, the creation of new resources in the cloud infrastructure may be limited. If you have any questions, please contact technical support. Our specialists continue to work on eliminating the consequences of the incident. We will be back with an update before 15:00 Moscow time.
- `2026-10-08 12:19` **unknown** — The ru-central1-b availability zone remains unavailable. The cause of the incident was a fire that occurred as a result of a UAV attack on the Yandex data center in Sasovo and affected part of the Yandex Cloud infrastructure. The data center has been completely shut down. Currently, the creation of new resources in the cloud infrastructure may be limited. Our team continues to work on eliminating the consequences of the incident. If you have any questions, please contact technical support. The next update will be published at 19:00 Moscow time.
- `2026-10-08 15:25` **unknown** — The situation with the availability of the ru-central1-b zone has not changed yet. Our specialists are continuing the restoration work, however, according to current estimates, it will not be possible to resume the work of the zone in the near future. The recommendation remains the same, if possible, to use alternative plans to restore the services. The next update will be given at 23:00 Moscow time.
- `2026-10-08 20:05` **unknown** — The availability of the ru-central1-b zone has not yet been restored. Our team continues to eliminate the consequences of the incident. If possible, we recommend using alternative service recovery options, including dedicated physical servers Bare Metal. The next update will be published on October 9 at 12:00 Moscow time.
- `2026-10-09 07:38` **unknown** — The ru-central1-b zone is still unavailable. The other availability zones are working. Our specialists are working around the clock to restore the previous level of infrastructure redundancy for key services in the near future and ensure the availability of additional capacities and backup components. The technical support team helps customers to resume their services. We have also agreed with other cloud providers to speed up the placement and backup of your resources on their sites. If you see a need, you can contact your colleagues directly: Selectel K2Cloud VK Cloud Using the code word "CloudBro”, our colleagues will be ready to process your application as a matter of priority and provide the necessary technical assistance together with our teams. If you need additional advice, please contact technical support or your account manager. The next update will be published at 15:00 Moscow time.
- `2026-10-09 12:42` **unknown** — The ru-central1-b zone is unavailable. Our specialists continue to work to eliminate the consequences of the incident and restore cloud services. Current status for some services: Managed Service for Kubernetes: it is now possible to transfer the management components of regional clusters from the ru-central1-b zone to other availability zones. To restore the applications, we recommend creating node groups in other zones and transferring the load to them. Managed Databases: in fault-tolerant clusters, it is now possible to remove hosts from the ru-central1-b zone. After deleting them, cluster management operations will become available, including changing settings and adding users. Yandex AI Studio: since yesterday, all models and components of the platform have been operating normally. At the same time, some features of DeepSeek models related to the size of the context window are preserved. Other AI services are also operating normally. If you have any questions or need help restoring resources, please contact technical support. The next update will be published at 22:00 Moscow time.
- `2026-10-09 18:47` **unknown** — The situation with the availability of the ru-central1-b zone remains unchanged so far. Our engineers continue to work around the clock to eliminate the consequences of the incident and restore services. The next update will be given on October 10 at 12:00 Moscow time.
- `2026-10-10 06:00` **unknown** — At the moment, the ru-central1-b zone remains unavailable. Our engineers continue to work to eliminate the consequences of the incident. To ensure stable operation of the cloud platform for the coming weekend, we are switching to fully manual management of quotas for computing resources and disks. All users have access to a small additional amount of resources beyond those already used. This will allow some clients, if necessary, to restore resources from backups, including for their subsequent transfer to other sites, or perform other operations. The creation of new resources beyond the limits set by us will be limited. Priority in expansion will be given to those whose critical services are still not restored. If you have any questions, please contact technical support or your account manager. The next update will be given on October 10 at 20:00 Moscow time.
- `2026-10-10 17:13` **unknown** — The ru-central1-b zone remains unavailable. In the morning, we switched to fully manual management of quotas for computing resources and disks. Technical support specialists sort out each situation manually, so responses may be delayed. We will try to sort out all quota requests as soon as possible. Our team continues to work around the clock to increase the amount of available computing resources. The next update will be given on October 11 at 12:00 Moscow time.

### 3. cloudflare — Missing Logs and Traces in Cloudflare Observability Platform

- URL: https://www.cloudflarestatus.com/incidents/pfvfug94jq4q
- Started: 2026-10-05 13:00 UTC
- Resolved: 2026-10-23 21:00 UTC
- Duration: 18d 8h 0m
- Impact: none, status: resolved
- Components: —
- Canonical incident: [ci-dcb212550be05a0c4eb3a32a](#ci-dcb212550be05a0c4eb3a32a)

- `2026-11-03 18:40` **resolved** — Between October 5 at 1:00 PM UTC and October 23 at 9:00 PM UTC, some customers may have experienced missing log events, incomplete traces, or inconsistent issue counts due to ingestion pipeline capacity limits. Traffic and Workers were not affected, and customers will not be billed for dropped events. Capacity has since been increased, and service has returned to normal.

### 4. grafana — IRM Unable to Send SMS to Thailand Numbers

- URL: https://stspg.io/9l8htmqt82s2
- Started: 2026-10-05 16:17 UTC
- Resolved: 2026-10-19 11:40 UTC
- Duration: 13d 19h 22m
- Impact: minor, status: resolved
- Components: Incident Management and Response (IRM)
- Canonical incident: [ci-aa6f59211a96cfbcca9feb49](#ci-aa6f59211a96cfbcca9feb49)

- `2026-10-05 16:17` **investigating** — We are currently investigating an issue impacting all SMS attempts to Austria for IRM. All SMS attempts to Austria for IRM are currently blocked. We are working on a resolution, and an update will follow as soon as we have more info. Users should configure mobile push or phone calls as notifications.
- `2026-10-05 21:47` **identified** — The issue has been identified. We are working on a resolution, but do not have an ETA at this time. We will provide another update as soon as we have more information to share. Users should configure mobile push or phone calls as notifications.
- `2026-10-06 13:42` **identified** — This issue is also impacting numbers from Thailand and Bangladesh. The title of this status page has been updated to reflect this. Users should configure mobile push or phone calls as notifications. We are working on a resolution, and we will provide another update as soon as we have more information to share.
- `2026-10-09 12:51` **identified** — We are continuing working on a resolution, and will provide another update as soon as we have more information to share. In the meantime we recommend, affected users to configure mobile push or phone calls as notifications.
- `2026-10-13 21:35` **identified** — We are continuing working on a resolution, and will provide another update as soon as we have more information to share. In the meantime we recommend, affected users to configure mobile push or phone calls as notifications.
- `2026-10-15 09:33` **identified** — SMS delivery to Austrian numbers has been restored. SMS delivery to Thailand and Bangladesh is still affected, and we are continuing to work on a resolution. We will provide another update as soon as we have more information to share. In the meantime, we recommend that affected users configure mobile push or phone calls as notifications.
- `2026-10-16 15:20` **identified** — SMS delivery to Bangladesh has been restored. SMS delivery to Thailand is still affected, and we are continuing to work on a resolution. We will provide another update as soon as we have more information to share. In the meantime, we recommend that affected users configure mobile push or phone calls as notifications.
- `2026-10-19 11:40` **resolved** — Following further investigation, we've confirmed that SMS delivery to certain destinations is not currently supported. We've updated our documentation to clarify the supported destinations and available notification options. We recommend using mobile push notifications or phone calls as alternatives where SMS delivery is unavailable.

### 5. cloudflare — Elevated Errors in China Network Colos

- URL: https://www.cloudflarestatus.com/incidents/cyxclg7nnhjy
- Started: 2026-10-10 04:13 UTC
- Resolved: 2026-10-15 09:28 UTC
- Duration: 5d 5h 15m
- Impact: minor, status: resolved
- Components: China Network
- Canonical incident: [ci-1d7fe76c931dba8c5cb73366](#ci-1d7fe76c931dba8c5cb73366)

- `2026-10-10 04:13` **investigating** — We are investigating a spike in errors observed in our China colos.
- `2026-10-10 05:48` **identified** — The issue has been identified and a fix is being implemented.
- `2026-10-12 04:23` **identified** — We are investigating reports of degraded performance for traffic traversing the public internet between certain cities and provinces in mainland China and origin servers located outside the region. The impact is being reported across multiple ISPs and manifests as slower speeds or intermittent timeouts. While this congestion occurs upstream of Cloudflare's network, our engineering team is actively working to mitigate the effects. Intermittent 524 timeout errors is expected to continue due to ongoing congestion. While there is no fixed timeline for full resolution, as there are dependencies outside our network, we are taking steps to reduce customer impact.
- `2026-10-12 15:33` **identified** — We have identified that increased latency and 5xx errors affecting certain China Network colos are caused by cross-border routing instability in the China network. We are actively working to mitigate impact for affected customers by adjusting traffic routing, and are continuing to work with our network partners on a resolution. Customers experiencing issues may contact Cloudflare Support.
- `2026-10-14 17:50` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-10-15 09:28` **resolved** — This incident has been resolved.

### 6. anthropic — Elevated errors on platform.claude.com

- URL: https://stspg.io/sjkw7njwf0mt
- Started: 2026-10-07 13:25 UTC
- Resolved: 2026-10-13 18:00 UTC
- Duration: 6d 4h 35m
- Impact: major, status: resolved
- Components: Claude Console (platform.claude.com)
- Canonical incident: [ci-f64b047b06bad8235f9ab606](#ci-f64b047b06bad8235f9ab606)

- `2026-10-07 13:25` **investigating** — We are investigating elevated errors affecting usage data on the platform.claude.com Usage page and the Admin API usage report. Core functionality is not affected. We will provide an update as soon as possible.
- `2026-10-07 16:36` **investigating** — We are continuing to work on elevated errors affecting usage data on the platform.claude.com Usage and Rate limits pages and the Admin API usage report. A mitigation for the Usage page and the Admin API usage report is rolling out; once it is in place, the most recent usage may take about 30 minutes to appear there. Core functionality is not affected.
- `2026-10-07 17:28` **identified** — We've identified the cause of errors loading usage data. The Admin API usage report has recovered, and fixes for the platform.claude.com Usage and Rate limits pages are in progress. The most recent 30 minutes of usage may be delayed. Core functionality is not affected.
- `2026-10-07 21:44` **monitoring** — We've mitigated the errors loading usage data on platform.claude.com and in the Admin API usage report. While we work on a full fix, the most recent ~30 minutes of usage may take longer to appear on the platform.claude.com. API requests are not affected, and we will provide additional updates here.
- `2026-10-13 18:00` **resolved** — This issue has been resolved.

### 7. openai — Delayed Costs data in the Compliance API

- URL: https://status.openai.com/incidents/01M4GV1M2A3K38T73JCVCD9G0R
- Started: 2026-10-09 17:23 UTC
- Resolved: 2026-10-12 16:02 UTC
- Duration: 2d 22h 39m
- Impact: minor, status: resolved
- Components: —
- Canonical incident: [ci-25c27c306bfe03faf64a026b](#ci-25c27c306bfe03faf64a026b)

- `2026-10-09 17:23` **identified** — We have identified an issue causing delays in publishing organization cost data through the Costs API. API and ChatGPT usage are not affected, and other reporting channels remain available. We are working to restore normal data freshness.
- `2026-10-09 17:27` **identified** — We have identified an issue causing delays in publishing Costs data through the Compliance API. The API Platform organization costs endpoint and ChatGPT usage are not affected, and other reporting channels remain available. We are working to restore normal data freshness.
- `2026-10-09 18:46` **identified** — Some customers may continue to experience delayed Costs data through the Compliance API. ChatGPT usage and API Platform requests are not affected, and other reporting channels remain available. We have identified the cause and are continuing work to restore normal data freshness. We will provide another update as work progresses.
- `2026-10-09 22:03` **monitoring** — We have implemented a mitigation, and delayed Costs data is now processing through the Compliance API. Some customers may continue to see incomplete Costs data while the backlog catches up. We expect this catch-up process to continue through Sunday, October 11. ChatGPT usage and API Platform requests remain unaffected. We are monitoring recovery.
- `2026-10-12 16:02` **resolved** — All impacted services have now fully recovered.

### 8. slack — Slack for iOS closing unexpectedly when opening channels

- URL: https://slack-status.com/2026-10/a0cfb0208a1d2032
- Started: 2026-10-09 15:35 UTC
- Resolved: 2026-10-12 17:03 UTC
- Duration: 3d 1h 28m
- Impact: unknown, status: resolved
- Components: Connectivity
- Canonical incident: [ci-057564c342d6384d89c544ba](#ci-057564c342d6384d89c544ba)

- `2026-10-09 15:35` **unknown** — Users may experience iOS closing unexpectedly when opening channels. We're actively looking in to this, and we'll share another update as soon as we have more information. We appreciate your patience in the meantime.
- `2026-10-09 16:20` **unknown** — We've developed a fix for the iOS app issue and are working to validate it before releasing it to customers. We'll share another update as soon as we have more information.
- `2026-10-09 23:58` **unknown** — An updated version of the Slack iOS app that fixes the crash was approved and began rolling out gradually. Customers who download the updated version from the app store can now open channels without the app closing unexpectedly. We'll provide an update once additional information becomes available.
- `2026-10-12 17:03` **unknown** — We've rolled out an updated version of the Slack iOS app to all impacted customers. Channels now open without the app closing unexpectedly. Users who are still experiencing this issue should update the Slack IOS app from the app store Store to resolve it. We apologize for the inconvenience and thank you for your patience.

### 9. gcp — Google Cloud Storage (GCS) customers experienced elevated latencies and error rates when performing upload or write operations.

- URL: https://status.cloud.google.com/incidents/mSFVt6PiP557bjRyYsut
- Started: 2026-10-08 11:00 UTC
- Resolved: 2026-10-09 16:46 UTC
- Duration: 1d 5h 46m
- Impact: major, status: resolved
- Components: Google Cloud Storage
- Canonical incident: [ci-74a2f2cec65976aed2992881](#ci-74a2f2cec65976aed2992881)

- `2026-10-08 18:24` **investigating** — **Summary** Google Cloud Storage customers are experiencing elevated latencies and error rates when performing upload or write operations **Description** We are experiencing an issue with Google Cloud Storage beginning on Thursday, 2026-10-08 05:19 PDT. Our engineering team continues to investigate the issue. We will provide an update by Thursday, 2026-10-08 11:50 PDT with current details. **Customer Symptoms** * Customers are experiencing elevated latencies and error rates when performing upload or write operations. * A smaller subset of read operations may also experience increased latency and errors **Workaround** None at this time.
- `2026-10-08 18:34` **investigating** — **Summary** Google Cloud Storage customers are experiencing elevated latencies and error rates when performing upload or write operations **Description** We are experiencing an issue with Google Cloud Storage beginning on Thursday, 2026-10-08 05:19 PDT. Initial mitigations applied by our engineering team leading to observable improvements in error rates in our internal monitoring. Work remains underway to roll out further mitigations. We will provide an update by Thursday, 2026-10-08 12:00 PDT with current details. **Customer Symptoms** * Customers are experiencing elevated latencies and error rates when performing upload or write operations. * A smaller subset of read operations may also experience increased latency and errors **Workaround** None at this time.
- `2026-10-08 18:55` **investigating** — **Summary** Google Cloud Storage customers are experiencing elevated latencies and error rates when performing upload or write operations **Description** We are experiencing an issue with Google Cloud Storage beginning on Thursday, 2026-10-08 05:19 PDT. Initial mitigations applied by our engineering team leading to observable improvements in error rates in our internal monitoring beginning 10:57 PDT. Work remains underway to roll out further mitigations. We will provide an update by Thursday, 2026-10-08 12:45 PDT with current details. **Customer Symptoms** * Customers are experiencing elevated latencies and error rates when performing upload or write operations. * A smaller subset of read operations may also experience increased latency and errors **Workaround** None at this time.
- `2026-10-08 19:38` **investigating** — **Summary** Google Cloud Storage customers are experiencing elevated latencies and error rates when performing upload or write operations **Description** We are continuing to mitigate the issue with Google Cloud Storage in the us-central1 region that began on Thursday, 2026-10-08 05:19 PDT. The initial mitigations applied by our engineering team continue to show observable improvements in error rates. Our engineering team is actively monitoring the environment and rolling out further mitigations to fully stabilize the service. We will provide an update by Thursday, 2026-10-08 14:00 PDT with current details. **Customer Symptoms** * Customers are experiencing elevated latencies and error rates when performing upload or write operations. * A smaller subset of read operations may also experience increased latency and errors **Workaround** None at this time.
- `2026-10-08 20:48` **investigating** — **Summary** Google Cloud Storage customers are experiencing elevated latencies and error rates when performing upload or write operations **Description** We are continuing to mitigate the issue with Google Cloud Storage in the us-central1 region that began on Thursday, 2026-10-08 05:19 PDT. The initial mitigations applied by our engineering team continue to show observable improvements in error rates. Our engineering team is actively monitoring the environment and rolling out further mitigations to fully stabilize the service. We will provide an update by Thursday, 2026-10-08 16:00 PDT with current details. **Customer Symptoms** * Customers are experiencing elevated latencies and error rates when performing upload or write operations. * A smaller subset of read operations may also experience increased latency and errors **Workaround** None at this time.
- `2026-10-09 16:44` **resolved** — ## Preliminary Incident Report We sincerely apologize for the disruption this incident caused to your business. We know how much you rely on Google Cloud, and we regret the impact on your productivity. We are working to address the root cause and prevent this from occurring in the future. Please note, this information is based on our best knowledge at the time of posting and is subject to change as our investigation continues. A final Incident Report with preventative actions will be posted once our investigation is complete. If you have experienced impact outside of what is listed below, please reach out to Google Cloud Support using [https\://cloud.google.com/support](https://cloud.google.com/support) ## Date/Time of the Issue (All time US/Pacific) Incident Start: 08 October 2026 04:00 Incident End: 08 October 2026 16:30 Duration: 12 hours, 30 minutes ## Summary On Thursday, 08 October 2026, Google Cloud Storage experienced elevated write latency and write error rates in the us-central1 region for a duration of 12 hours, 30 minutes. We are taking immediate steps to ensure this doesn’t happen again. ## Preliminary Root Cause Google Cloud Storage relies on a highly distributed backend architecture to process and store incoming data. During the incident window, specific storage backend nodes within the us-central1 region experienced a localized capacity shortfall. The sequence of events leading to customer impact was as follows: * A traffic sorting optimization intended to address localized capacity constraints caused severe resource exhaustion on specific backend storage nodes in the affected region. * This resource exhaustion caused backend services handling object data ingestion to begin timing out. * As a result, incoming upload and write requests routed to these nodes experienced elevated latency and failed to complete, resulting in errors surfaced to customers. * Concurrently, a separate issue with the underlying database service contributed to additional metadata processing errors during a portion of the incident window. Google engineers have begun a full root cause analysis (RCA) and will provide additional information once it is available. ## Remediation The issue was initially detected by Google engineering teams actively monitoring system telemetry early in the incident window. To mitigate the issue, Google engineers executed the following step-by-step strategy: * Engineers initially attempted to dynamically allocate additional emergency capacity to the affected region, but this was blocked by systemic regional capacity limits. * To relieve immediate pressure, engineers selectively drained traffic from the most overloaded backend storage nodes. * A configuration change was deployed to disable the traffic sorting optimization, significantly reducing resource overhead on the storage backend. * For the concurrent metadata issue, engineers rolled back a recent update to the underlying metadata database service. Once traffic levels stabilized and resource utilization returned to normal baseline levels, the previously drained nodes were restored to active service, successfully restoring normal operations. ## Description of Impact On 08 October, 2026, from 04:00 to 16:30 US/Pacific, customers utilizing the us-central1 region experienced elevated latency and error rates during Cloud Storage upload and write operations. * Customers primarily observed HTTP 503 UNAVAILABLE errors on WriteObject operations and client-side timeouts. * A minor impact was also observed on read operations during this timeframe. * Cloud Storage operations in all other regions, as well as all other Google Cloud services, remained unaffected by this specific event.
- `2026-10-14 07:15` **resolved** — ## Incident Report ## Date/Time of the Issue (All time US/Pacific) **Incident First Occurrence** Incident Start: 08 October 2026 04:00 Incident End: 08 October 2026 16:30 Duration: 12 hours, 30 minutes **Incident Recurrence** Recurrence Start: 09 October 2026 04:05 Recurrence End: 09 October 2026 09:46 Duration: 5 hours, 41 minutes **Total Duration:** 18 hours, 11 minutes ## Summary On Thursday, 08 October 2026, Google Cloud Storage experienced elevated write latency and write error rates while performing upload or write operations in the us-central1 region for a duration of 12 hours, 30 minutes. This was followed by a recurrence on Friday, 09 October 2026 in the us-central1-b zone for a duration of 5 hours, 41 minutes. We are taking immediate steps to ensure this doesn’t happen again. ## Root Cause Google Cloud Storage relies on a highly distributed backend architecture to process and store incoming data. During the incident window, specific storage backend nodes within the us-central1 region experienced a localized capacity shortfall due to increased traffic. The sequence of events leading to customer impact was as follows: * A traffic sorting optimization configuration update to the Google Cloud Storage load balancing layer intended to address localized capacity constraints caused severe resource exhaustion on specific storage backend nodes in the affected region by inadvertently bypassing the standard capacity-aware routing logic, resulting in a severe load imbalance. * Consequently, a small fraction of storage backend nodes in the us-central1 region received an overwhelming proportion of incoming write traffic despite having low available serving capacity. This concentration of traffic led to resource exhaustion on the affected backend servers causing backend services handling object data ingestion to begin timing out. * Because the underlying instances could not process the high volume of requests quickly enough, the incoming upload and write requests routed to these nodes experienced elevated latency and failed to complete, resulting in errors surfaced to customers. * Concurrently, a separate issue with the underlying database service contributed to additional metadata processing errors during a portion of the incident window. * The recurrence of the issue on October 9 was caused by an incomplete rollback of an upload planner feature, which broke the load balancing mechanism and overloaded a few smaller storage backend nodes. ## Remediation and Prevention The issue was initially detected on 08 October 2026 at 04:50 by Google engineering teams actively monitoring system telemetry early in the incident window. To mitigate the issue, Google engineers executed the following step-by-step strategy: * Engineers initially attempted to dynamically allocate additional emergency capacity to the affected region, but this was blocked by systemic regional capacity limits. * To relieve immediate pressure, engineers selectively drained traffic from the most overloaded storage backend nodes. * A configuration change was deployed to disable the traffic sorting optimization, significantly reducing resource overhead on the storage backend. * For the concurrent metadata issue, engineers rolled back a recent update to the underlying metadata database service. * While these steps provided temporary relief, a recurrence of the load imbalance required further traffic redirection and draining of affected storage components. The incident was permanently resolved when engineering fully reverted the problematic load balancing configuration change, restoring the previous capacity-aware routing behavior and allowing traffic to distribute evenly across the region. Once traffic levels stabilized and resource utilization returned to normal baseline levels, the previously drained nodes were restored to active service, successfully restoring normal operations. **Preventative Actions:** To ensure this issue does not recur, we are implementing the following improvements: - **Safe Configuration Rollouts:** We are updating our deployment protocols to ensure that load balancing logic changes are rigorously validated in alternative, smaller regions before being rolled out to larger, highly utilized regions. - **Intelligent Operational Controls:** We are enhancing the storage load balancing architecture to strictly enforce capacity constraints, ensuring that traffic distribution logic cannot bypass pre-spilling safeguards and route traffic to clusters lacking sufficient serving resources. - **Resilient Scaling Logic:** We are implementing targeted improvements to accelerate internal data movement within highly utilized storage cells, enabling the system to process data more efficiently and handle temporary load spikes with greater resilience. ## Description of Impact On 08 October 2026, from 04:00 to 16:30 in us-central1 region and 09 October 2026, from 04:05 to 09:46 in us-central1-b zone, customers experienced elevated latency and error rates during Cloud Storage upload and write operations. * **Upload & Write Failures:** Customers primarily observed HTTP 503 UNAVAILABLE errors on WriteObject operations and client-side timeouts. * **Increased Write Latency:** Write operations that did succeed experienced significantly higher latencies than normal. * Cloud Storage operations in all other regions, as well as all other Google Cloud services, remained unaffected by this specific event.

### 10. grafana — Degraded Performance with Increased Latency for Grafana Cloud Metrics and Traces in GCP US Central

- URL: https://stspg.io/cjbf6knm4jq4
- Started: 2026-10-08 13:06 UTC
- Resolved: 2026-10-09 00:39 UTC
- Duration: 11h 33m
- Impact: none, status: resolved
- Components: GCP US Central - prod-us-central-0, GCP US Central - prod-us-central-0: Ingestion, GCP US Central - prod-us-central-0: Ingestion, GCP US Central - prod-us-central-5, GCP US Central - prod-us-central-5: Ingestion
- Canonical incident: [ci-c309a05b6bfa84541633c192](#ci-c309a05b6bfa84541633c192)

- `2026-10-08 13:06` **investigating** — Beginning around 10:20 UTC, we have been seeing intermittent elevated error rates and latency affecting Grafana Cloud Metrics and Logs in GCP US Central (prod-us-central-0 and prod-us-central-5), and periodic increased latency and occasional errors affecting Grafana Cloud Traces in prod-us-central-0. Customers may see errors or delays when sending metrics, logs and traces. We are actively investigating. Next update by 13:40 UTC.
- `2026-10-08 13:25` **investigating** — We continue to investigate intermittent elevated error rates and latency affecting Grafana Cloud Metrics and Logs in GCP US Central (prod-us-central-0 and prod-us-central-5), and periodic increased latency affecting Grafana Cloud Traces in prod-us-central-0. Error rates have decreased, but intermittent errors and latency may still occur. We are working with our cloud provider as part of the investigation. We do not have an estimated time of resolution at this time. Next update by 14:30 UTC.
- `2026-10-08 13:39` **investigating** — Update to scope: we have not confirmed any impact to Grafana Cloud Logs, and Logs has been removed from the affected components. We continue to investigate intermittent elevated error rates and latency affecting Grafana Cloud Metrics in GCP US Central (prod-us-central-0 and prod-us-central-5), and periodic increased latency affecting Grafana Cloud Traces in prod-us-central-0. Error rates have decreased, but intermittent errors and latency may still occur. We are working with our cloud provider as part of the investigation. We do not have an estimated time of resolution at this time. Next update by 14:30 UTC.
- `2026-10-08 13:52` **investigating** — We continue to investigate intermittent elevated error rates and latency affecting Grafana Cloud Metrics in GCP US Central (prod-us-central-0 and prod-us-central-5), and periodic increased latency affecting Grafana Cloud Traces in prod-us-central-0. Intermittent errors and latency may still occur. We are working with our cloud provider as part of the investigation. We do not have an estimated time of resolution at this time. Next update by 14:30 UTC.
- `2026-10-08 14:40` **investigating** — We are investigating degraded performance with increased latency in the ingestion path for Grafana Cloud Metrics in GCP US Central (prod-us-central-0 and prod-us-central-5) and for Grafana Cloud Traces in prod-us-central-0. Customers may experience intermittent errors or delays when sending metrics and traces. We are investigating this with our cloud provider. We do not have an estimated time of resolution at this time. Next update by 15:30 UTC.
- `2026-10-08 15:19` **investigating** — We are investigating degraded performance with increased latency in the ingestion path for Grafana Cloud Metrics in GCP US Central (prod-us-central-0 and prod-us-central-5) and for Grafana Cloud Traces in prod-us-central-0. Customers may experience intermittent errors or delays when sending metrics and traces. We are investigating this with our cloud provider. We do not have an estimated time of resolution at this time. Next update by 16:30 UTC.
- `2026-10-08 16:21` **investigating** — We continue seeing elevated error rates and latency, and keep working with our cloud provider on this issue. Next update by 17:30 UTC.
- `2026-10-08 17:30` **investigating** — We are investigating degraded performance with increased latency in the ingestion path for Grafana Cloud Metrics in GCP US Central (prod-us-central-0 and prod-us-central-5) and for Grafana Cloud Traces in prod-us-central-0. Customers may experience intermittent errors or delays when sending metrics and traces. We continue to investigate this with our cloud provider. We do not have an estimated time of resolution at this time. Next update by 18:30 UTC
- `2026-10-08 18:15` **investigating** — We are investigating degraded performance and increased latency in the ingestion path for Grafana Cloud Metrics in GCP US Central (prod-us-central-0 and prod-us-central-5) and Grafana Cloud Traces in prod-us-central-0. Customers may experience intermittent errors or delays when sending metrics and traces. We have also received reports that some customers using Google Managed Prometheus in this region are experiencing issues when querying their GMP instances. We continue to investigate the issue with our cloud provider. We do not have an estimated time to resolution at this time.
- `2026-10-08 19:13` **investigating** — We are investigating degraded performance and increased latency in the ingestion path for Grafana Cloud Metrics in GCP US Central (prod-us-central-0 and prod-us-central-5) and Grafana Cloud Traces in prod-us-central-0. Customers may experience intermittent errors or delays when sending metrics and traces. We have also received reports that some customers using Google Managed Prometheus in this region are experiencing issues when querying their GMP instances. We continue to investigate the issue with our cloud provider. We do not have an estimated time to resolution at this time.
- `2026-10-08 20:41` **investigating** — We are investigating degraded performance and increased latency in the ingestion path for Grafana Cloud Metrics in GCP US Central (prod-us-central-0 and prod-us-central-5) and Grafana Cloud Traces in prod-us-central-0. Customers may experience intermittent errors or delays when sending metrics and traces. We continue to investigate the issue with our cloud provider. We do not have an estimated time to resolution at this time.
- `2026-10-08 21:42` **investigating** — We are investigating degraded performance and increased latency in the ingestion path for Grafana Cloud Metrics in GCP US Central (prod-us-central-0 and prod-us-central-5) and Grafana Cloud Traces in prod-us-central-0. Customers may experience intermittent errors or delays when sending metrics and traces. We continue to investigate the issue with our cloud provider. We do not have an estimated time to resolution at this time.
- `2026-10-09 00:39` **resolved** — This incident has been resolved

### 11. cloudflare — Durable Objects Availability Issue in Eastern North America

- URL: https://www.cloudflarestatus.com/incidents/z28czd65h29l
- Started: 2026-10-05 17:14 UTC
- Resolved: 2026-10-05 22:35 UTC
- Duration: 5h 21m
- Impact: minor, status: resolved
- Components: Containers, Durable Objects, R2, Workers Assets, Workflows
- Canonical incident: [ci-52f863d412c65a5c1bbf1f11](#ci-52f863d412c65a5c1bbf1f11)

- `2026-10-05 17:14` **investigating** — We are investigating an issue leading to decreased availability for Durable Objects and dependent services in the Eastern North America region. Our engineering team is actively working to mitigate the impact, and further updates will be provided as soon as more information becomes available.
- `2026-10-05 17:53` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-10-05 22:35` **resolved** — This incident has been resolved.

### 12. cloudflare — Network Performance Issues in Columbus

- URL: https://www.cloudflarestatus.com/incidents/tt8f4ttcuonf
- Started: 2026-10-07 02:48 UTC
- Resolved: 2026-10-07 11:51 UTC
- Duration: 9h 2m
- Impact: major, status: resolved
- Components: Columbus, OH, United States - (CMH), Network
- Canonical incident: [ci-7e70d415c6e2be105272e1c1](#ci-7e70d415c6e2be105272e1c1)

- `2026-10-07 02:48` **investigating** — Cloudflare is investigating issues with Network Performance in Columbus, OH (CMH). We are working to analyze and mitigate this problem. More updates to follow shortly.
- `2026-10-07 03:06` **identified** — The issue has been identified and a fix is being implemented.
- `2026-10-07 05:11` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-10-07 11:51` **resolved** — This incident has been resolved.

### 13. circleci — Delay in data appearing in UI and API

- URL: https://stspg.io/hkfkwz292bz8
- Started: 2026-10-09 11:33 UTC
- Resolved: 2026-10-10 02:08 UTC
- Duration: 14h 35m
- Impact: major, status: postmortem
- Components: CircleCI API, CircleCI UI
- Canonical incident: [ci-fb8b1117166bd38c3550921f](#ci-fb8b1117166bd38c3550921f)

- `2026-10-09 11:33` **identified** — There is a delay in our data infrastructure causing delays before build data and some notifications become visible. We are working on a mitigation.
- `2026-10-09 13:00` **identified** — What's impacted: Customers triggering pipelines, and customers relying on outbound webhooks and email and Slack notifications. What can you expect Since 11:15 UTC: - Pipelines are still being created while this is ongoing but not yet processed, so retriggering may result in duplicate pipelines. - Outbound webhooks and transactional notifications (Slack and email) are not being sent. - Pipeline statuses may not be visible. - Usage charge data on the Project Usage page will be stale. Thank you for your patience while our engineers work to resolve this. Next update: We will provide another update by 13:30 UTC.
- `2026-10-09 13:30` **identified** — Between 11:15 UTC and about 13:00 UTC, pipelines were created but not processed, and outbound webhooks and transactional notifications (Slack and email) were not sent. What's impacted: Customers triggering pipelines, and customers relying on outbound webhooks and email and Slack notifications. What can you expect: - Since about 13:00 UTC, pipelines are being processed again and jobs are starting. - Most jobs are currently starting within about 4 minutes, and some are waiting up to about 10 minutes. Pipelines triggered during the incident are still being processed and may take longer to start. - Outbound webhooks and notifications have resumed. - Pipelines are still being processed from the backlog, so retriggering may result in duplicate pipelines. Thank you for your patience while our engineers work to resolve this. Next update: We will provide another update by 14:00 UTC.
- `2026-10-09 14:00` **identified** — Since 11:15 UTC, customers have experienced disruption to pipelines and notifications. Between 11:15 UTC and about 13:00 UTC, pipelines were created but not processed, and outbound webhooks and transactional notifications (Slack and email) were not sent. What's impacted Customers running workflows and jobs, and customers viewing usage data on the Project Usage page. What can you expect - Pipelines are being processed again, including those triggered during the incident. - Since about 13:25 UTC, workflows and jobs have been starting with increasing delays as we work through the backlog. Most jobs are currently starting within about 10 minutes, and some are waiting up to about 25 minutes. - We expect to work through this backlog by about 14:35 UTC. - Workflows from pipelines triggered during the incident are still waiting to start, so retriggering may result in duplicate workflows. - Usage charge data on the Project Usage page is currently stale. Thank you for your patience while our engineers work to resolve this. Next update We will provide another update by 14:30 UTC.
- `2026-10-09 14:29` **identified** — Since 11:15 UTC, customers have experienced disruption to pipelines and notifications. Between 11:15 UTC and about 13:00 UTC, pipelines were created but not processed, and outbound webhooks and transactional notifications (Slack and email) were not sent. What's impacted Customers running workflows and jobs, and customers viewing usage data on the Project Usage page. What can you expect - Delays starting workflows and jobs are improving. Most jobs are now starting within about 3 minutes, down from about 10 minutes at 13:50 UTC. - Some workflows from pipelines triggered during the incident are still waiting up to about 40 minutes to start. - Working through the backlog is taking slightly longer than expected. We now expect to clear it by about 14:45 UTC. - Workflows from pipelines triggered during the incident are still waiting to start, so retriggering may result in duplicate workflows. - Usage charge data on the Project Usage page is currently stale. Thank you for your patience while our engineers work to resolve this. Next update We will provide another update by 15:00 UTC.
- `2026-10-09 14:43` **identified** — Since 11:15 UTC, customers have experienced disruption to pipelines and notifications. Between 11:15 UTC and about 13:00 UTC, pipelines were created but not processed, and outbound webhooks and transactional notifications (Slack and email) were not sent. Between about 13:00 UTC and 14:40 UTC, workflows and jobs started with delays while we worked through the backlog. What’s impacted Customers running workflows and jobs, and customers viewing usage data on the Project Usage page. What can you expect - The backlog of workflows from pipelines triggered during the incident cleared at about 14:40 UTC. - Most jobs are now starting within about 30 seconds, and some are waiting up to about 3 minutes. - Usage charge data on the Project Usage page is currently stale. Thank you for your patience while our engineers continue to work on the underlying issue. Next update We will provide another update by 15:10 UTC.
- `2026-10-09 14:54` **monitoring** — Since 11:15 UTC, customers have experienced disruption to pipelines and notifications. Between 11:15 UTC and about 13:00 UTC, pipelines were created but not processed, and outbound webhooks and transactional notifications (Slack and email) were not sent. Between about 13:00 UTC and 14:40 UTC, workflows and jobs started with delays while we worked through the backlog. What’s impacted Customers running workflows and jobs, and customers viewing usage data on the Project Usage page. What can you expect - The backlog of pipelines and workflows from the incident has cleared. - Since about 14:45 UTC, most workflows and jobs are starting within about 15 seconds, and start times are continuing to return to normal. - Usage charge data on the Project Usage page is currently stale. Thank you for your patience while our engineers monitor the recovery. Next update We will provide another update by 15:20 UTC.
- `2026-10-09 15:32` **monitoring** — Since 11:15 UTC, customers have experienced disruption to pipelines and notifications. Between 11:15 UTC and about 13:00 UTC, pipelines were created but not processed, and outbound webhooks and transactional notifications (Slack and email) were not sent. Between about 13:00 UTC and 14:40 UTC, workflows and jobs started with delays while we worked through the backlog. What's impacted Customers viewing pipelines, workflows and jobs in the CircleCI UI, and customers viewing usage data on the Project Usage page. What can you expect - Workflows and jobs are running on time. - The CircleCI UI is about 30 minutes behind. New workflows and status changes may take that long to appear. - Our engineers are actively working to speed up processing so the UI can catch up. - Please avoid retriggering workflows that don't yet appear in the UI, as this may create duplicates. - Usage charge data on the Project Usage page is currently stale. Thank you for your patience while our engineers work to reduce this delay. Next update We will provide another update by 16:00 UTC.
- `2026-10-09 16:02` **monitoring** — Since 11:15 UTC, customers have experienced disruption to pipelines and notifications. Between 11:15 UTC and about 13:00 UTC, pipelines were created but not processed, and outbound webhooks and transactional notifications (Slack and email) were not sent. Between about 13:00 UTC and 14:40 UTC, workflows and jobs started with delays while we worked through the backlog. What's impacted Customers running workflows and jobs, customers viewing pipelines, workflows and jobs in the CircleCI UI, and customers viewing usage data on the Project Usage page. What can you expect - Since about 15:15 UTC, workflows and jobs have been starting with increasing delays. Most jobs are currently starting within about 4 minutes, and some are waiting up to about 20 minutes. - The CircleCI UI remains about 30 minutes behind on average, and some updates are taking longer. New workflows and status changes may take that long to appear. - Please avoid retriggering workflows that don't yet appear in the UI, as this may create duplicates. - Usage charge data on the Project Usage page is currently stale. Thank you for your patience while our engineers work to reduce these delays. Next update We will provide another update by 16:30 UTC.
- `2026-10-09 16:32` **monitoring** — Since 11:15 UTC, customers have experienced disruption to pipelines and notifications. Between 11:15 UTC and about 13:00 UTC, pipelines were created but not processed, and outbound webhooks and transactional notifications (Slack and email) were not sent. Between about 13:00 UTC and 14:40 UTC, workflows and jobs started with delays while we worked through the backlog. Since about 15:15 UTC, workflows and jobs have again been starting with delays. What's impacted Customers running workflows and jobs, customers viewing pipelines, workflows and jobs in the CircleCI UI, and customers viewing usage data on the Project Usage page. What can you expect - Delays starting workflows and jobs have stopped increasing and are beginning to ease. Most jobs are starting within about 4 minutes, and some are waiting up to about 25 minutes. - The CircleCI UI is now about 20 minutes behind, down from about 30 minutes. New workflows and status changes may take that long to appear. - Please avoid retriggering workflows that don't yet appear in the UI, as this may create duplicates. - Usage charge data on the Project Usage page is currently stale. Thank you for your patience while our engineers work to reduce these delays. Next update We will provide another update by 17:00 UTC.
- `2026-10-09 17:06` **monitoring** — Since 11:15 UTC, customers have experienced disruption to pipelines and notifications. Between 11:15 UTC and about 13:00 UTC, pipelines were created but not processed, and outbound webhooks and transactional notifications (Slack and email) were not sent. Between about 13:00 UTC and 14:40 UTC, workflows and jobs started with delays while we worked through the backlog. Since about 15:15 UTC, workflows and jobs have again been starting with delays. What’s impacted Customers running workflows and jobs, customers viewing pipelines, workflows and jobs in the CircleCI UI, and customers viewing usage data on the Project Usage page. What can you expect - Delays starting workflows and jobs are holding steady. Most jobs are starting within about 4 minutes, and some are waiting up to about 8 minutes, down from about 25 minutes. - The CircleCI UI remains about 20 minutes behind on average, and some updates are taking longer. New workflows and status changes may take that long to appear. - Please avoid retriggering workflows that don’t yet appear in the UI, as this may create duplicates. - Usage charge data on the Project Usage page is currently stale. Thank you for your patience while our engineers work to reduce these delays. Next update We will provide another update by 17:30 UTC.
- `2026-10-09 17:34` **identified** — Since 11:15 UTC, customers have experienced disruption to pipelines and notifications. Between 11:15 UTC and about 13:00 UTC, pipelines were created but not processed, and outbound webhooks and transactional notifications (Slack and email) were not sent. Between about 13:00 UTC and 14:40 UTC, workflows and jobs started with delays while we worked through the backlog. Since about 15:15 UTC, workflows and jobs have again been starting with delays. What’s impacted Customers triggering pipelines and running workflows and jobs, customers viewing pipelines, workflows and jobs in the CircleCI UI, and customers viewing usage data on the Project Usage page. What can you expect - Since about 17:00 UTC, delays have increased significantly. Most jobs are now taking about 12 minutes to start, and some are waiting up to about 30 minutes. - Since about 17:15 UTC, some jobs are failing to start. - Pipelines are being processed more slowly, so some may take longer to start. - The CircleCI UI is more than an hour behind. New workflows and status changes may take that long to appear. - Please avoid retriggering workflows that don’t yet appear in the UI, as this may create duplicates. - Usage charge data on the Project Usage page is currently stale. Thank you for your patience while our engineers work to resolve this. Next update We will provide another update by 18:00 UTC.
- `2026-10-09 18:08` **identified** — Since 11:15 UTC, customers have experienced disruption to pipelines and notifications. Between 11:15 UTC and about 13:00 UTC, pipelines were created but not processed, and outbound webhooks and transactional notifications (Slack and email) were not sent. Between about 13:00 UTC and 14:40 UTC, workflows and jobs started with delays while we worked through the backlog. Since about 15:15 UTC, workflows and jobs have again been starting with delays. What’s impacted Customers triggering pipelines and running workflows and jobs, customers viewing pipelines, workflows and jobs in the CircleCI UI, and customers viewing usage data on the Project Usage page. What to expect - Since about 17:00 UTC, we are seeing an increase in delays. Most jobs are now taking about 25 minutes to start, and some are waiting up to about an hour. - Since about 17:15 UTC, some in-progress workflows have failed and a small number of jobs are failing to start. - New pipelines are being queued and are taking longer to start. - The CircleCI UI is more than an hour behind. New workflows and status changes may take that long to appear. - Please avoid retriggering workflows that don't yet appear in the UI, as this may create duplicates. - Usage charge data on the Project Usage page is currently stale. Our engineers continue to work to resolve the issue at hand. Next update We will provide another update by 18:30 UTC.
- `2026-10-09 18:32` **identified** — What's changed since our last update * The CircleCI UI is about an hour behind, and up to about 90 minutes for some updates. Workflows and jobs may already have run, or may be running, even if the UI shows them as not started or does not show them at all. * Jobs that are running are starting faster. Most are now starting within about 15 minutes, down from about 28 minutes, though some are still waiting over an hour. * New pipelines are increasingly delayed. * Some workflows are still failing. Workflows that could not be processed in time are being marked as failed. * Please avoid retriggering workflows that don't yet appear in the UI, as this may create duplicates. Still ongoing: Usage charge data on the Project Usage page is stale. Next update We will provide another update by 19:00 UTC.
- `2026-10-09 18:57` **identified** — What's changed since our last update * Jobs are starting faster. Most are now starting within about 8 minutes, down from about 15 minutes, and the longest waits have dropped from over an hour to about 45 minutes. * The CircleCI UI is beginning to catch up, but it is still behind. Workflows and jobs may already have run, or may be running, even if the UI shows them as not started or does not show them at all. * Pipelines are being processed faster, but new pipelines are still delayed. Still ongoing * Some workflows and jobs are still failing. * Please avoid retriggering workflows that don't yet appear in the UI, as this may create duplicates. * Usage charge data on the Project Usage page is stale. Next update We will provide another update by 19:30 UTC.
- `2026-10-09 19:29` **identified** — Since our last update, jobs are starting faster: most are now starting within about 6 minutes, down from about 8 minutes. The CircleCI UI is catching up but still about an hour behind. New pipelines are being processed as they arrive. Jobs may already have run even if the UI does not show them yet, so please avoid retriggering workflows that don't appear in the UI. We will provide another update by 20:00 UTC.
- `2026-10-09 20:06` **identified** — Jobs are starting faster and fewer are failing: most jobs are now starting within about 4 minutes, down from about 6 minutes, and the longest waits have dropped from about 30 minutes to about 20 minutes. Pipelines from earlier in the incident are now being worked through. The CircleCI UI is still behind, and some older updates may take longer to appear. Jobs may already have run even if the UI does not show them yet, so please avoid retriggering workflows that don't appear in the UI. We will provide another update by 20:30 UTC.
- `2026-10-09 21:22` **identified** — Since about 20:48 UTC, new workflows and jobs are not starting. We have temporarily paused starting new work so that our systems can catch up on work already in progress, including delayed updates to the CircleCI UI and API. Older updates are now appearing in the UI and API as this backlog clears. New pipelines are still being accepted and queued Please avoid retriggering pipelines or workflows during this time, as this may create duplicates once work resumes. We will provide another update by 22:00 UTC.
- `2026-10-09 22:09` **identified** — Since our last update, most of the pipelines that were delayed earlier in the incident have now been processed, and their workflows are queued to start. New workflows and jobs remain paused while we prepare to resume work safely. Delayed updates continue to appear in the CircleCI UI and API. Please continue to avoid retriggering pipelines or workflows, as this may create duplicates once work resumes. We will provide another update by 22:30 UTC.
- `2026-10-09 22:28` **identified** — Since our last update, we have begun gradually resuming work, and some queued workflows and jobs are starting again. Most queued work is still waiting, and jobs that do start may have waited over an hour. We are increasing the rate carefully to avoid overloading our systems. Delayed updates continue to appear in the CircleCI UI and API. Please continue to avoid retriggering pipelines or workflows, as this may create duplicates as queued work resumes. We will provide another update by 23:00 UTC.
- `2026-10-09 23:00` **identified** — A large number of queued workflows and jobs have started, and the CircleCI UI and API have caught up and are showing updates normally with few delays. We are continuing to work through the remaining queued work. Jobs from the queue may have waited about an hour or more to start, and some Linux jobs are waiting up to about 10 minutes for capacity. Please continue to avoid retriggering while queued work continues to process. We will provide another update by 23:30 UTC.
- `2026-10-09 23:38` **identified** — We are still working through a significant backlog of queued workflows and jobs. Wait times are improving but they remain longer than normal and some jobs have waited over an hour. The CircleCI UI and API remain up to date. Please avoid retriggering pipelines or workflows while the backlog clears, as this may create duplicates. We will provide another update by 00:00 UTC.
- `2026-10-09 23:59` **identified** — Since our last update, the backlog of queued workflows and jobs has cleared, as of about 23:47 UTC. Workflows and jobs are now starting within seconds, and the CircleCI UI and API are up to date. We are monitoring our systems closely to make sure they remain stable. We will provide another update by 00:30 UTC.
- `2026-10-10 00:45` **identified** — Workflows and jobs have continued to start normally, within seconds, and the CircleCI UI and API remain up to date. Slack and email notifications delayed earlier in the incident are still being delivered. Usage charge data on the Project Usage page is stale. We are continuing to monitor our systems closely and will provide another update by 01:15 UTC.
- `2026-10-10 01:23` **monitoring** — Workflows and jobs are starting and running normally, and the CircleCI UI and API are up to date. We are continuing to monitor our systems closely to make sure they remain stable. We will provide another update by 02:00 UTC.
- `2026-10-10 02:08` **resolved** — Between 11:15 UTC and 23:47 UTC on October 9, customers experienced delays and failures across pipelines, workflows and jobs, along with delays in the CircleCI UI, API and notifications. From 11:15 to about 13:00 UTC, pipelines were created but not processed, and outbound webhooks and Slack and email notifications were not sent. From about 13:00 to 14:40 UTC, workflows and jobs started with delays while we worked through the backlog. From about 15:15 to 20:48 UTC, workflows and jobs started with increasing delays, of up to about an hour, and the CircleCI UI and API fell behind by up to about 2 hours. During this time, some jobs failed to start and some in-progress workflows failed. From 20:48 to about 22:10 UTC, we temporarily paused starting new workflows and jobs so our systems could recover. New pipelines were accepted and queued during this time. From about 22:10 to 23:47 UTC, queued workflows and jobs resumed and were processed. Some of them had waited up to about 2 hours. The issue has been resolved and all affected functionality has returned to normal. Usage charge data on the Project Usage page is not yet up to date and will be updated once our data backfills complete. Customers whose jobs failed during this window may rerun affected jobs. We thank you for your patience while our engineers worked on implementing a fix.
- `2026-10-12 22:32` **postmortem** — ## Summary Between 11:15 UTC and 23:47 UTC on October 9, CircleCI customers experienced hours-long delays and failures across pipelines, workflows and jobs, along with delays in the CircleCI UI, API and notifications. * From 11:15 to 13:00 UTC, pipelines were created but not processed, and outbound webhooks and Slack and email notifications were not sent. * From 13:00 to 14:40 UTC, workflows and jobs started with delays while we worked through the backlog. * From 15:15 to 20:48 UTC, workflows and jobs started with increasing delays, of up to about an hour, and the CircleCI UI and API fell behind by up to about 2 hours. During this time, some jobs failed to start and some in-progress workflows failed. * From 20:48 to 22:10 UTC, we temporarily paused starting new workflows and jobs so our systems could recover. New pipelines were accepted and queued during this time. * From 22:10 to 23:47 UTC, queued workflows and jobs resumed and were successfully processed. Some of them had waited up to about 2 hours. * By 23:47, the incident was mitigated. The status page for this incident can be found at [October 9 - Delay in data appearing in UI and API](https://status.circleci.com/incidents/6cfhhlr4rmm0). Customers experienced a similar incident earlier in the week at [October 6 - Delays in pipelines, workflows, UI data, and notifications](https://status.circleci.com/incidents/z4hbq8h49j8c). Multiple CircleCI services use Temporal backed by Postgres as a reliable queueing mechanism, and both incidents had the same cause: the rate of writes to Temporal's Postgres instance were faster than Postgres’s transaction log could archive those writes. In both cases, this caused multi-hour delays in processing both pipelines and status notifications. We evaluated several mitigations, then paused new workflow starts so the system could recover safely, and ramped back up gradually. We then throttled work back up until our systems returned to normal. To permanently resolve the bottleneck, we have split work across multiple Temporal instances, each with its own independent database and worker clusters, and we load-balance across them. We had already scheduled infrastructure work starting the week of October 12 that would have removed this exact bottleneck and prevented these incidents. So starting 19:11 on October 9 our engineering team implemented, tested, and rolled out a simpler version of these changes within a few hours. We further added new alerting and tuned the underlying databases for better performance under high load. Customers whose jobs failed during this window may rerun them. ## Background To run customer jobs, CircleCI executes a series of steps: we accept customer pipelines; we assign jobs for those pipelines to workers; we deliver status updates for those pipelines and jobs; and we charge customers for the credits those jobs consume. Each of these steps is implemented by an independent service, with reliable queues in between. Historically we have implemented these reliable queues with RabbitMQ on top of durable storage, with exponential backoff and dead letter queues for reliability. Over the last year, we have increasingly reimplemented our reliable queues on [Temporal](https://temporal.io/), which provides far better durability and reliability guarantees over a simpler and more powerful programming model. Our build service \(which turns a pipeline config into an executable plan\), our workflow orchestrator service \(which distributes tasks onto workers\), our notification service \(which delivers status updates\), and our charging service \(which charges customers for credits\) all now use Temporal. Temporal stores the state for its workflows and activities \(called “history”\) in a database, and this is the only component in a Temporal instance that is synchronous and stateful. As a consequence, a Temporal instance’s ability to accept and execute workflows is dependent on the write rate of its backing database. We have also configured Temporal to delete completed workflows and activities from its history 3 days after they complete. At the time of this incident, we used a single clustered Temporal instance backed by a Multi-AZ Postgres RDS instance for all of these reliable queues. While a Multi-AZ RDS instance is resilient to zone failures, it maintains database consistency via a single write-ahead-log \(WAL\). At the time we had configured our Postgres instance to store its WAL on a Dedicated Log Volume \(DLV\) so log writes did not compete with data writes. The database archives the log continuously to backup storage, and it can only reuse space on the volume after the log has been archived. If the database writes logs faster than it archives, the volume fills and Postgres has no space to the write its WAL log. Then the database stops accepting incoming writes and goes entirely read-only. An RDS DLV is limited to 1TB of storage and 3000 IOPS and per AWS these limits are hardcoded and cannot be increased. ## What Happened _\(All times UTC\)_ One of the contributing factors to the October 6 incident was that we were processing a high rate of Temporal history deletions at the same time as incoming work, which in turn put pressure on the backing database. To resolve the October 6 incident, we changed cleanup settings on Temporal’s largest tables and moved the transaction log back to its dedicated volume. These changes eliminated the major source of transaction log writes, but did not consider the additional log writes from scheduled deletions. Because we have configured Temporal to delete completed workflows and activities from its history 3 days after they complete, the October 6 history was scheduled to be deleted on October 9. Starting at 08:00 on October 9, Temporal began processing those scheduled deletions, and as a consequence free space on the Postgres DLV began dropping precipitously, from 0.8TB at 08:00 to 0.1TB by 10:10. Our monitoring alerted at 10:28 and we declared an incident at 10:31. While our infrastructure team was investigating, at 11:15 the DLV ran out of space and the Postgres instance went read-only. This stopped all Temporal processing in its tracks. At 11:42, we restarted Postgres to move the WAL log from the DLV onto the data volume and use the spare disk space and IOPS there. This maneuver had mitigated the incident on October 6, and we knew it would take roughly 60 minutes. While the database was restarting, we began implementing several load-shedding techniques upstream of Temporal: * We began investigating whether we should shed incoming workloads altogether to keep the rest moving, but ultimately decided that this would provide no value. * At 12:29, we flagged off credit charging updates, which represented 12% of Temporal activities. This left the Pipeline Usage page with stale charging data, until we would backfill it, but would not lose data or prevent pipelines from running. At 12:59, the database was back up and began processing work that had queued since 11:15 and by 14:40, this backlog had cleared and jobs were starting on time. However, by 15:15, delays returned as database’s routing cleanup had fallen far behind on the tables Temporal uses to track queued and scheduled work. Those tables normally hold tens of thousands of deleted rows awaiting cleanup, but by late afternoon, they held tens of millions of rows. Every Temporal operation had to read past those rows, so the database spent much more processing time on each one and Temporal ran slower. This resulted in customers continuing to see hour-long delays in starting pipelines and in status updates. For the next several hours we attempted various mitigations, including: * Manipulating Temporal configuration values to reduce database pressure. This had minimal effect. * Running manual vacuum jobs to clean up dead rows in Postgres. This ended up competing for CPU with queries themselves, and was ultimately counterproductive. We cancelled them at 17:45 to relieve the pressure on the database load. * We began trying to move all of the notification usecases to a separate Temporal instance. This represented 18% of Temporal activities. We successfully finished moving notifications by 17:07. * At 17:16, we successfully implemented an optimization in our representation of pipeline state, which reduced it by 5%. That in turn reduced the size of database writes and therefore data pressure on the Postgres WAL log. None of these fully mitigated the issue. By 20:48, we realized that Postgres was now thrashing on database CPU and was no longer bottlenecked on log writes. We further realized that our upstream systems were both submitting new work and also retrying work that had earlier been submitted and failed. Recognizing that we were in a classic retry storm, we paused all incoming and retry work and let it queue up, in order to allow the database to fully complete in-process work. We then slowly throttled up queued work from 22:07 to keep the database at a sustainable load. This fully succeeded by 23:47 when all queued work had started, and we closed the incident at 02:08 on October 10. ## Future Prevention and Process Improvement ### Multiple Temporal Instances In parallel with incident mitigation and resolution, starting at 19:11 on October 9, we also began standing up multiple sharded Temporal instances and implementing the ability to load-balance incoming work between them. Each of the multiple Temporal instances maintains its own independent database and worker clusters, and so is isolated from the others. As a postmortem action item from the October 6 incident, we had already scheduled infrastructure work that would have removed the bottleneck of a single Temporal instance and Postgres database, starting October 12 and expecting to take 2 weeks. On October 9, we realized that this work simply could not wait, so our engineering team implemented, tested, and rolled out a simpler version of the planned changes within a few hours. These changes completed successfully at 01:45 on October 10. **In addition, we are changing how we store and monitor transaction logs.** Our Temporal database no longer uses the fixed-size log volume. We are adding alerts on how far log archiving falls behind and how fast the database writes logs, so we can act hours earlier than we did on October 9. **We are also revisiting the Temporal's deletion schedule.** Temporal deletes finished work a few days after it runs, so a busy day returns as database load later in the week. We are changing when and how quickly those deletions run. **We are changing how our systems retry when Temporal is slow.** On October 9, our systems retried failed work against Temporal and added to its load. As we committed after the October 6 incident, we are changing our systems to wait for Temporal to recover so work is not dropped. ## Conclusion Customer experience is our top priority, and we commit to continually improving the reliability of our systems to match the trust that our customers place in us. We thank our customers for their patience while our team worked to resolve these incidents. Please reach out to our support team with any questions or concerns.

### 14. openai — Elevated Work Mode errors

- URL: https://status.openai.com/incidents/01M45E0JFFJZ6EQ3BWYPW74R2Z
- Started: 2026-10-05 07:03 UTC
- Resolved: 2026-10-06 02:10 UTC
- Duration: 19h 6m
- Impact: minor, status: resolved
- Components: —
- Canonical incident: [ci-16c02e0e6cbe0a1b216975a4](#ci-16c02e0e6cbe0a1b216975a4)

- `2026-10-05 07:03` **investigating** — We are observing elevated Work Mode errors. Scheduled tasks may also be impacted.
- `2026-10-05 07:14` **investigating** — We are continuing to observe elevated Work Mode errors. Scheduled tasks may also be impacted.
- `2026-10-05 09:50` **monitoring** — We have mitigated the issue, and are monitoring recovery.
- `2026-10-06 08:29` **resolved** — All impacted services have now fully recovered.

### 15. cloudflare — Cloudflare DNS Resolver 1.1.1.1 and Gateway DNS are having issues in SIN

- URL: https://www.cloudflarestatus.com/incidents/4knp3pzs3qvx
- Started: 2026-10-09 14:41 UTC
- Resolved: 2026-10-09 19:30 UTC
- Duration: 4h 49m
- Impact: minor, status: resolved
- Components: Gateway, Recursive DNS, Singapore, Singapore - (SIN)
- Canonical incident: [ci-2c1e26dd8f2e3ec55cffa8a8](#ci-2c1e26dd8f2e3ec55cffa8a8)

- `2026-10-09 14:41` **investigating** — Cloudflare is aware of, and investigating, an issue which potentially impacts multiple users in Singapore that use the 1.1.1.1 public resolver or Gateway DNS. Further detail will be provided as more information becomes available.
- `2026-10-09 14:42` **investigating** — We are continuing to investigate this issue.
- `2026-10-09 16:03` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-10-09 19:30` **resolved** — This incident has been resolved.

### 16. yandexcloud — Increasing timings in the Object Storage service

- URL: https://status.yandex.cloud/en/incidents/2130
- Started: 2026-10-08 19:30 UTC
- Resolved: 2026-10-09 09:50 UTC
- Duration: 14h 20m
- Impact: minor, status: resolved
- Components: Object Storage
- Canonical incident: [ci-b9d6e1019a1436b9ecb42027](#ci-b9d6e1019a1436b9ecb42027)

- `2026-10-08 21:35` **investigating** — There is a partial unavailability of individual components of Object Storage. For some requests to the S3 API, an increase in response time was recorded, primarily for object recording operations. The read operations are running normally. The developers are already aware of it and are fixing it.
- `2026-10-09 10:15` **resolved** — At 12:50 Moscow time, the problem was fixed, the Object Storage service is working correctly.

### 17. cloudflare — API Shield JWT Validation Errors

- URL: https://www.cloudflarestatus.com/incidents/ts295zqycqn7
- Started: 2026-10-05 10:55 UTC
- Resolved: 2026-10-05 22:39 UTC
- Duration: 11h 43m
- Impact: minor, status: resolved
- Components: API Shield
- Canonical incident: [ci-f49247b63ccb439cf646ba98](#ci-f49247b63ccb439cf646ba98)

- `2026-10-05 10:55` **investigating** — Cloudflare is investigating an issue affecting API Shield JWT validation for customers using symmetric keying material. Our team is actively working on restoring normal service.
- `2026-10-05 11:24` **identified** — The issue has been identified and a fix is being implemented.
- `2026-10-05 20:09` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-10-05 22:39` **resolved** — This incident has been resolved.

### 18. cloudflare — Cloudflare API service issues

- URL: https://www.cloudflarestatus.com/incidents/x8n12150xwf6
- Started: 2026-10-09 18:52 UTC
- Resolved: 2026-10-10 02:12 UTC
- Duration: 7h 19m
- Impact: minor, status: resolved
- Components: API
- Canonical incident: [ci-06d5bce0f029e3c9c82e3c57](#ci-06d5bce0f029e3c9c82e3c57)

- `2026-10-09 18:52` **investigating** — Cloudflare is investigating issues with Cloudflare Dashboard APIs. These issues do not affect the serving of cached files via the Cloudflare CDN or other security features at the Cloudflare Edge. Customers using the Cloudflare APIs are impacted as requests might fail and/or errors may be displayed.
- `2026-10-09 20:58` **identified** — The issue has been identified and a fix is being implemented.
- `2026-10-09 21:10` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-10-10 02:12` **resolved** — This incident has been resolved.

### 19. github — Several services are degraded

- URL: https://stspg.io/w3rcssslq22w
- Started: 2026-10-06 19:57 UTC
- Resolved: 2026-10-06 23:40 UTC
- Duration: 3h 43m
- Impact: critical, status: resolved
- Components: Pull Requests, Webhooks
- Canonical incident: [ci-81b00cc38adda810f181cd6d](#ci-81b00cc38adda810f181cd6d)

- `2026-10-06 19:57` **investigating** — We are investigating reports of degraded performance for Pull Requests and Webhooks
- `2026-10-06 20:04` **investigating** — Webhooks is operating normally.
- `2026-10-06 20:12` **investigating** — Pull Requests is operating normally.
- `2026-10-06 20:22` **investigating** — A number of services including issues, pull requests, and webhooks, were briefly degraded, but have now recovered. Enterprise migrations are paused while we validate causes and mitigations.
- `2026-10-06 20:26` **monitoring** — The degradation has been mitigated. We are monitoring to ensure stability.
- `2026-10-06 20:57` **investigating** — All services are functioning normally except for enterprise migrations, which are still paused. They will resume soon.
- `2026-10-06 21:29` **investigating** — We are working towards a mitigation that will allow us to resume processing enterprise migrations.
- `2026-10-06 22:40` **investigating** — We are deploying a mitigation and plan to resume migrations shortly thereafter
- `2026-10-06 23:40` **resolved** — On October 6, 2026, from 19:44 to 19:50 UTC, GitHub experienced a service outage with elevated error rates peaking at 1.30% of all requests, totaling approximately 1.16 million failed requests. This impacted 12 services across GitHub, including Pull Requests, Webhooks, Repos, OAuth, Git over SSH, Codespaces, and Gists. The incident was triggered by GitHub Enterprise Importer (GEI) making multiple, long duration queries which saturated a database connection pool. Recovery was seen at approximately 19:50 UTC once the migration's calls stopped; GEI migrations were reenabled at 23:32 UTC after a code fix was merged.<br /><br />We have completed a detailed root-cause review and are implementing improved database-health safeguards. We will continue to review and add additional safeguards to prevent future recurrence.

### 20. github — Incident with Actions

- URL: https://stspg.io/c11dc9nb1zdq
- Started: 2026-10-05 19:11 UTC
- Resolved: 2026-10-05 22:49 UTC
- Duration: 3h 37m
- Impact: critical, status: resolved
- Components: Actions, Pages
- Canonical incident: [ci-b531d0fe442ab4c7a92d0e1d](#ci-b531d0fe442ab4c7a92d0e1d)

- `2026-10-05 19:11` **investigating** — We are investigating reports of degraded performance for Actions
- `2026-10-05 19:15` **investigating** — We’re investigating an issue causing delays when assigning GitHub-hosted runners to Actions jobs. Some workflows may take longer to start across runner configurations. Our teams are working to mitigate the issue, and we’ll provide an update as we learn more.
- `2026-10-05 19:50` **investigating** — We’re still investigating delays in assigning GitHub-hosted runners, affecting workflow start times across multiple runner configurations. Our teams are working to mitigate the impact; we’ll share another update as we learn more.
- `2026-10-05 20:39` **investigating** — We’re continuing to investigate job failures and delays affecting GitHub-hosted runner assignment and workflow start times. Our teams are actively working to mitigate the impact and will provide another update as we learn more.
- `2026-10-05 20:47` **investigating** — Actions is experiencing degraded availability. We are continuing to investigate.
- `2026-10-05 21:09` **investigating** — In addition to ongoing issues with Actions and Hosted Runners, some customers are also unable to access repository list, licensing and billing pages within their GitHub accounts. We continue to investigate and are working on mitigation.
- `2026-10-05 21:22` **investigating** — Pages is experiencing degraded performance. We are continuing to investigate.
- `2026-10-05 21:31` **investigating** — Actions is experiencing degraded performance. We are continuing to investigate.
- `2026-10-05 21:32` **investigating** — We have applied mitigations to address GitHub Actions failures. Queued jobs are clearing, new jobs are not delayed, and we expect full runner recovery shortly. We continue to monitor for full recovery and are working to mitigate ongoing issues affecting other services, including access to repository lists, licensing, and billing pages.
- `2026-10-05 21:54` **investigating** — Actions is operating normally.
- `2026-10-05 22:40` **investigating** — Pages is operating normally.
- `2026-10-05 22:49` **resolved** — On October 5, 2026, between 18:48 and 22:49 UTC, GitHub Actions and Hosted Runners experienced degraded performance. Over the full incident, 14.3% of workflow runs using GitHub-hosted runners and 26.5% of individual GitHub-hosted jobs did not start within five minutes. <br /><br /> <br />During the roughly 90-minute period of highest impact, 47.0% of workflow runs using GitHub-hosted runners failed and 71.9% of GitHub-hosted jobs did not start within five minutes, while existing Pages sites continued to serve normally. A portion of customers also encountered intermittent errors in repository lists, licensing and billing pages, and some Copilot features. Self-hosted runners were not affected. <br /><br /><br />The incident was caused by a partial networking failure that disrupted connectivity between one of our on-premises datacenter regions and a subset of cloud-hosted regional database and storage services. <br /><br /><br />We mitigated the incident by shifting database and storage traffic to healthy regions and endpoints and increasing capacity for affected compute services. GitHub Actions and Hosted Runners recovered by 21:54 UTC, and all affected services recovered by 22:49 UTC. <br /><br /> <br />We are partnering with our cloud provider to understand the network failure and identify improvements to network-path resilience, detection, and recovery for similar incidents.

### 21. circleci — Delays in pipelines, workflows, UI data, and notifications

- URL: https://stspg.io/s7tpg47srxdd
- Started: 2026-10-06 16:00 UTC
- Resolved: 2026-10-06 21:18 UTC
- Duration: 5h 18m
- Impact: major, status: postmortem
- Components: Notifications & Status Updates, Pipelines & Workflows
- Canonical incident: [ci-266fbf9dfabc9fed11b692e4](#ci-266fbf9dfabc9fed11b692e4)

- `2026-10-06 16:00` **investigating** — What's impacted Customers using CircleCI pipelines and workflows, including UI data display, outbound notifications, and outbound webhooks. What can you expect Delays in pipelines being created and in workflows and jobs starting. Job, workflow, and pipeline data may take up to 15 minutes to appear in the UI. Outbound notifications and webhooks are delayed by up to 1 minute. Next update We will provide another update within 30 minutes or as we have more information to share.
- `2026-10-06 16:10` **investigating** — Since 15:47 UTC, customers may also see failed jobs, in addition to the delays previously reported. Our engineers are continuing to investigate and mitigate the issue.
- `2026-10-06 16:31` **identified** — As of 16:02 UTC, pipelines are failing entirely, so no new workflows or jobs are running at this point. This is in addition to the delays and failed jobs reported earlier. Our engineers are continuing to work on mitigating the issue, and we will share more information as soon as we have it.
- `2026-10-06 16:42` **identified** — As of 16:39 UTC, we are starting to see recovery for new pipelines and workflows being submitted, and jobs are beginning to run again. Our engineers are continuing to work on the recovery and will share more information as soon as we have it.
- `2026-10-06 17:21` **identified** — While Pipelines, workflows, and jobs are running again, customers may continue to see longer than usual delays between jobs in a workflow, and outbound notifications are delayed by approximately 1 minute. Notifications and status updates to your version control system are beginning to go out again. Our engineers are continuing to work toward a full recovery.
- `2026-10-06 18:04` **identified** — Customers continue to experience longer than usual delays between jobs in a workflow, as well as delays in outbound notifications. We thank you for your patience while our engineers continue to investigate and work on mitigating the issue. We will share another update within the next 30 minutes, or sooner if more information becomes available.
- `2026-10-06 18:55` **identified** — Customers continue to experience longer than usual delays before workflows and jobs start, in some cases up to approximately 30 minutes. This is improved from earlier in the incident, but delays remain. Delays in outbound notifications have improved. We thank you for your patience while our engineers continue to work toward restoring normal job start times. We will share another update within the next 30 minutes, or sooner if more information becomes available.
- `2026-10-06 19:16` **identified** — Customers continue to experience delays. Most are now starting within approximately 20 minutes, improved from approximately 40 minutes earlier in the incident. A very small number of workflows may have waited more than 1 hour. Our engineers have made changes that have improved service stability, and the remaining delay is due to a backlog of work that built up during the incident. The backlog is decreasing, and we are focused on clearing it faster. We will share more information as progress continues. We thank you for your patience while our engineers work through the backlog. We will share another update within the next 30 minutes, or sooner if more information becomes available.
- `2026-10-06 19:51` **identified** — Delays for workflows and jobs to start are continuing to decrease. Most are now starting within approximately 10 minutes, improved from approximately 40 minutes earlier in the incident. A very small number of workflows may have waited more than 1 hour. If your job failed to start, please rerun the job. Delays in outbound notifications have improved. The remaining delay is due to a backlog of work that built up during the incident, and that backlog has reduced by more than half. We expect the backlog to be cleared in approximately 15 minutes. We will share another update once it has cleared, or within the next 30 minutes at the latest.
- `2026-10-06 20:13` **identified** — The backlog of delayed work has been cleared, and workflows are starting again with no delays in notifications. However, customers using Docker and Linux jobs across multiple resource classes, including medium, large, and 2xlarge gen2, may now see jobs waiting in the queue for few minutes before they start. Our engineers are continuing to work on mitigating this issue. If your job failed to start, please rerun the job. We thank you for your patience while our engineers work on this. We will share another update within the next 30 minutes, or sooner if more information becomes available.
- `2026-10-06 20:31` **monitoring** — Workflows and jobs are starting normally, and wait times for Docker and Linux jobs have returned to normal. Our engineers are continuing to monitor this issue. If your job failed to start, please rerun the job. We thank you for your patience while our engineers work on this. We will share another update within the next 30 minutes, or sooner if more information becomes available.
- `2026-10-06 21:18` **resolved** — Between approximately 15:30 UTC and 20:30 UTC on October 6, 2026, customers experienced delays and failures affecting pipelines, workflows, and jobs. From approximately 16:02 UTC to 16:39 UTC, pipelines were failing and no new workflows or jobs could start. After service resumed, workflows and jobs continued to start with delays until approximately 20:00 UTC. Start times averaged up to approximately 40 minutes at their peak, and a very small number of workflows waited more than 1 hour. Outbound notifications and webhooks were also delayed until approximately 19:00 UTC, typically by about a minute. As the backlog of delayed work cleared, the resulting increase in demand placed pressure on our execution capacity, causing further delays for Docker and Linux jobs. From approximately 19:45 UTC to 20:30 UTC, these jobs waited in the queue longer than usual before starting. At the peak, around 20:05 UTC, waits reached up to approximately 8 minutes for 2xlarge gen2, 6 minutes for xlarge gen2, and less than 5 minutes for a few other resource classes. The issue has been resolved, and all affected functionality has returned to normal. Customers whose jobs failed during this window may rerun affected jobs. We thank you for your patience while our team worked on implementing a fix.
- `2026-10-07 22:08` **postmortem** — ## Summary On October 6, 2026, from 15:30 to 20:30 UTC, CircleCI customers experienced delays creating pipelines and starting workflows and jobs. From 15:48 to 16:38 UTC, no workflows or jobs started, and pipelines triggered from 16:02 UTC failed. Shorter, intermittent delays began earlier, at 13:30 UTC. Data in the CircleCI UI, notifications, and commit status updates to version control providers were also delayed. The incident started in the database behind our workflow orchestration service, which coordinates every pipeline, workflow and job on CircleCI. Routine maintenance on the database's largest tables wrote transaction logs faster than the database could archive them, and the logs filled the storage volume set aside for them. To keep the database from becoming read-only, our engineers moved the logs to the database's main storage volume. That move took the database offline for 50 minutes. When the database came back, it processed work at about half its normal rate until our engineers changed a database setting at 18:19 UTC. We then worked through the backlog, and job start times returned to normal by 20:30 UTC. Customers whose jobs failed during this window may rerun them. The original status page can be found [here](https://status.circleci.com/incidents/z4hbq8h49j8c). ## Background When you trigger a pipeline, CircleCI hands it to a workflow orchestration service. That service tracks the state of every workflow and job, decides when each job is ready to run, and passes ready jobs to our execution fleet. It also drives the job data you see in the UI and the notifications and commit statuses we send when work finishes. The orchestration service stores its state in a database. The database writes every change to a transaction log before applying it, and it keeps those logs on a dedicated storage volume. It archives the logs continuously so it can reuse that space. If archiving falls behind and the volume fills, the database stops accepting writes, and no new work can move forward. The database also runs a periodic maintenance process \(autovacuums\) on each table to keep its internal records valid over time. On a very large table, this process reads every page of the table and writes a large volume of transaction logs. ## What Happened _\(All times UTC\)_ From 10:02 on October 6, our monitoring raised short-lived alerts about errors in the orchestration service. Each alert cleared on its own within about 15 minutes. Our engineers had seen similar brief alerts before, and the system showed no other signs of stress, so they followed the runbook each time, and each alert cleared on its own partway through and also started work to make the alerts less sensitive. Around 12:40, along with the regular workload which generates its own transaction logs, the maintenance process on several of the database's largest tables began writing additional transaction logs faster than the database could archive them, and the log volume started to fill. At 13:30, the orchestration service began to slow down. Through 15:30, some pipelines and jobs took longer to start, in short bursts. At 13:35, our monitoring alerted again, this time alongside related alerts from several other services. Infrastructure Engineers followed the alert runbooks, clearing up some of the alerts. The alerts returned again at around 14:30 and after investigating further, we declared an incident at 14:54. We should have recognized the combination of alerts sooner, and we have covered how we are fixing that below. After declaring the incident, our engineers traced the slowdown to the maintenance process \(anti-wraparound autovacuum\). The database restarts this process automatically if it is cancelled, so our engineers changed its settings to help it finish faster. By 15:30, the delays were continuous, and some jobs began to fail. At 15:43, the transaction log volume was close to full. If it filled, the database would stop accepting writes and no work could run. Our engineers decided to move the transaction logs onto the database's main storage volume, which had plenty of free space. The move required the database to go offline, and we could not predict in advance how long that would take. The move started at 15:47. From 15:48 to 16:38, the orchestration service could not process any work. Customers could still trigger pipelines, but no workflows or jobs started, and from 16:02 newly triggered pipelines failed. We posted to our status page at 15:58 and raised it to a major outage at 16:20. While the database was offline, our engineers prepared a replacement database as a fallback. The original database came back first, so we kept it in service. At 16:38, the database came back online and work started flowing again. With transaction logs and regular data now sharing the same storage, the database ran more slowly than before. From 16:38 to 18:19, the orchestration service processed work at about half its normal rate, and jobs waited tens of minutes to start, up to about 50 minutes at the longest. During that time, the database had to finish archiving its backlog of transaction logs before we could move them back to a dedicated volume. We also investigated and implemented several database parameter changes to stop further maintenance runs from starting and prepared ways to reduce the work reaching the service. Along with the replacement database, we also prepared a complete stack of the service during that time so that we could move to it at the risk of data loss and opted against that move. At 18:19, our engineers changed a database setting to cut the time each write spent waiting on storage. Processing speed recovered, and the backlog began to clear. To clear it faster, we added database and server capacity to the systems that hand jobs to our execution fleet. By about 19:55, new jobs were again starting on time. The backlog then reached our execution fleet. From 19:49 to 20:30, some Docker jobs on larger resource classes waited up to about 10 minutes to start while capacity scaled up. By 20:30, job start times were back to normal. We moved the status page to monitoring at 20:32 and resolved the incident at 21:19. About 3,500 jobs failed to start during the incident. Customers may rerun these jobs. We found no evidence that CircleCI ran any job more than once. Customers who reran a workflow while the original run was still delayed may have seen both runs complete. Annual plan customers can work with their account team to review usage. ## Future Prevention and Process Improvement We are taking the following steps to prevent a recurrence and improve our response time: **We are adding alerts on the conditions that caused this incident.** Our monitoring caught the slowdown, but we had no alert on how full the transaction log volume was or on how far archiving had fallen behind. We are adding alerts on both, along with alerts on the database maintenance that drives them, so we can act earlier. We are also revisiting all our alerts to ensure noisy or sensitive alerts don’t hide real issues. **We are tuning how maintenance process \(autovacuum\) runs on our database.** We are looking at the frequency and the aggressiveness with which the autovacuums run on our database. We are also tracking the size of these tables over time so they stay within safe limits. **We are removing the single transaction log bottleneck.** All of the orchestration service's writes currently go through one database and one transaction log, so when that log fell behind, every pipeline, workflow and job slowed down with it. We are evaluating two approaches: splitting the service across multiple database instances, each with its own transaction log, and moving it to a distributed database built to spread writes across many nodes. Either approach would limit the effect of a backlog to part of the workload. **We are making our systems wait for the orchestration service to recover.** During the incident, some jobs failed because our systems gave up after retrying. We are working on changing that behavior. **We are improving our status page updates during long incidents.** Customers told us our early updates did not give enough detail. We are updating our guidance so updates include specific impact and timing sooner. ‌ Customer experience is our top priority, and we commit to continually improving the reliability of our systems to match the trust that our customers place in us. We thank our customers for their patience while our team worked to resolve this incident. Please reach out to our support team with any questions or concerns.

### 22. openai — Elevated errors in ChatGPT conversations

- URL: https://status.openai.com/incidents/01M46G248GGHV9RSAJE82BZY1X
- Started: 2026-10-05 16:58 UTC
- Resolved: 2026-10-06 00:01 UTC
- Duration: 7h 2m
- Impact: minor, status: resolved
- Components: —
- Canonical incident: [ci-f6f4424427f4046d5e2cc46a](#ci-f6f4424427f4046d5e2cc46a)

- `2026-10-05 16:58` **investigating** — We are investigating elevated error rates affecting some ChatGPT conversations. Error rates are improving, and our team is continuing to assess recovery. We will share further updates as they become available.
- `2026-10-05 17:55` **monitoring** — We have mitigated the issue affecting ChatGPT conversations and are monitoring recovery.
- `2026-10-06 00:01` **resolved** — All impacted services have now fully recovered.

### 23. openai — Unable to create new ChatGPT Work threads in latest version of desktop app

- URL: https://status.openai.com/incidents/01M4ADHMMBK7K92NJFJ0DXRYJM
- Started: 2026-10-07 05:31 UTC
- Resolved: 2026-10-07 11:53 UTC
- Duration: 6h 21m
- Impact: major, status: resolved
- Components: —
- Canonical incident: [ci-58fe92f1c3a4fa125c4de798](#ci-58fe92f1c3a4fa125c4de798)

- `2026-10-07 05:31` **investigating** — We are investigating the issue. Version 26.1002.51308 is affected
- `2026-10-07 05:54` **identified** — we have identified the issue, working on rolling back to the previous version.
- `2026-10-07 09:28` **identified** — We have identified the issue. We have released an updated app version for Linux and macOS. We are actively working to roll out the fix for Windows.
- `2026-10-07 11:06` **identified** — We are actively working on a fixed app version for Windows.
- `2026-10-07 11:53` **resolved** — We have released an updated app version with the fix for Linux, macOS and Windows. Affected users should update to the latest version.

### 24. openai — Elevated error rates in Codex Cloud

- URL: https://status.openai.com/incidents/01M46KED6KQ6PR195KZYYVNY8W
- Started: 2026-10-05 17:58 UTC
- Resolved: 2026-10-06 00:01 UTC
- Duration: 6h 2m
- Impact: none, status: resolved
- Components: —
- Canonical incident: [ci-9712f4ec943d180f692b6acb](#ci-9712f4ec943d180f692b6acb)

- `2026-10-05 17:58` **identified** — We have identified that users are experiencing elevated errors within Codex Cloud. We are working on implementing a mitigation.
- `2026-10-05 18:07` **monitoring** — We have applied the mitigation and are monitoring the recovery.
- `2026-10-06 08:22` **resolved** — All impacted services have now fully recovered.

### 25. openai — Admin Console Intermittently Inaccessible

- URL: https://status.openai.com/incidents/01M46C47WYYD9NMZVGH83X3ZR1
- Started: 2026-10-05 15:50 UTC
- Resolved: 2026-10-05 21:07 UTC
- Duration: 5h 17m
- Impact: none, status: resolved
- Components: —
- Canonical incident: [ci-a1a70a6819b7b9547baf2129](#ci-a1a70a6819b7b9547baf2129)

- `2026-10-05 15:50` **investigating** — We are investigating the issue for the listed services.
- `2026-10-05 19:33` **monitoring** — We have applied the mitigation and are monitoring the recovery.
- `2026-10-05 21:07` **resolved** — All impacted services have now fully recovered.

### 26. stripe — Elevated Swish errors

- URL: https://stspg.io/m7j171ppknbs
- Started: 2026-10-06 00:30 UTC
- Resolved: 2026-10-06 07:04 UTC
- Duration: 6h 33m
- Impact: none, status: resolved
- Components: Acquirers and payment methods
- Canonical incident: [ci-a3dd9b08db53546229d15e29](#ci-a3dd9b08db53546229d15e29)

- `2026-10-06 00:30` **investigating** — As of 00:28 UTC, we've seen elevated errors on Swish payments. We're continuing to investigate and will post further updates.
- `2026-10-06 07:04` **resolved** — From 00:00–03:00 UTC October 6th, we saw elevated errors on Swish payments as a result of scheduled maintenance. This is now resolved.

### 27. cloudflare — Network Performance Issues in Kochi

- URL: https://www.cloudflarestatus.com/incidents/xczvsioz1un8
- Started: 2026-10-08 14:00 UTC
- Resolved: 2026-10-08 17:00 UTC
- Duration: 3h 0m
- Impact: minor, status: resolved
- Components: —
- Canonical incident: [ci-3a323695409125036a7a6913](#ci-3a323695409125036a7a6913)

- `2026-10-08 17:14` **resolved** — Cloudflare resolved an issue with network performance for our Kochi (COK) facility which occurred today, 2026-10-08, between approximately 14:00 and 17:00 UTC. Customers may have observed elevated latency and/or error rates from this facility during this period.

### 28. openai — Some ChatGPT content may be unavailable on Android

- URL: https://status.openai.com/incidents/01M4H7FKXDA425TDNP2CRGEYA3
- Started: 2026-10-09 21:00 UTC
- Resolved: 2026-10-10 00:26 UTC
- Duration: 3h 26m
- Impact: major, status: resolved
- Components: —
- Canonical incident: [ci-1504dfe18c7a0d6c7d36211f](#ci-1504dfe18c7a0d6c7d36211f)

- `2026-10-09 21:00` **identified** — We have identified an issue affecting the ChatGPT Android app where users may see “This content is unavailable in this version of the app” instead of certain content. A fix is rolling out gradually in the latest Android version. Users can use ChatGPT on the web while we work to restore the experience for older app versions.
- `2026-10-10 00:26` **resolved** — All impacted services have now fully recovered.

### 29. github — Incident with Git Operations, Issues, Actions and Pull Requests

- URL: https://stspg.io/4p8x0jnnsx7m
- Started: 2026-10-07 17:17 UTC
- Resolved: 2026-10-07 18:04 UTC
- Duration: 47m
- Impact: critical, status: resolved
- Components: Actions, Git Operations, Issues, Pull Requests
- Canonical incident: [ci-8a0991181da00ff1f8976d77](#ci-8a0991181da00ff1f8976d77)

- `2026-10-07 17:17` **investigating** — We are investigating reports of degraded performance for Actions, Git Operations, Issues and Pull Requests
- `2026-10-07 17:25` **investigating** — We observed widespread impact across GitHub services from 16:52 UTC to 17:01 UTC, and are investigating the root cause. All systems have recovered. This incident is a reoccurrence of the incident from earlier in the day https://www.githubstatus.com/incidents/djlmxz2zd0j7. Durable mitigation is in progress. We will provide another update in 30 minutes.
- `2026-10-07 17:27` **monitoring** — The degradation has been mitigated. We are monitoring to ensure stability.
- `2026-10-07 18:00` **monitoring** — All systems have recovered. We have applied mitigations and do not anticipate another recurrence.
- `2026-10-07 18:04` **resolved** — On October 7, 2026, from 16:52 to 17:01 UTC, GitHub experienced a recurrence of the write-operation outage from earlier that day. Both outages shared the same root cause. Please refer to the combined summary covering both incidents, their impact, and our prevention measures: https://www.githubstatus.com/incidents/djlmxz2zd0j7

### 30. yandexcloud — The problem is with the Cloud Functions and Serverless Containers services

- URL: https://status.yandex.cloud/en/incidents/2124
- Started: 2026-10-08 11:00 UTC
- Resolved: 2026-10-08 12:25 UTC
- Duration: 1h 25m
- Impact: major, status: resolved
- Components: Cloud Functions, Serverless Containers
- Canonical incident: [ci-fd0e36afed5adea81c871030](#ci-fd0e36afed5adea81c871030)

- `2026-10-08 11:24` **investigating** — We noticed the complete unavailability of Cloud Functions and Serverless Containers services in the Cloud. The developers are already aware of it and are fixing it.
- `2026-10-08 18:30` **resolved** — At 3:25 p.m., the problem was fixed, and the Cloud Functions and Serverless Containers services are working correctly. We will tell you about the reasons for the incident later.

### 31. openai — Computer Use Is Experiencing Issues

- URL: https://status.openai.com/incidents/01M493CZ0CB14KPC0EER83MNVB
- Started: 2026-10-06 17:15 UTC
- Resolved: 2026-10-06 20:00 UTC
- Duration: 2h 45m
- Impact: minor, status: resolved
- Components: —
- Canonical incident: [ci-ef5d020ca1dcb8d3aca305e0](#ci-ef5d020ca1dcb8d3aca305e0)

- `2026-10-06 17:15` **investigating** — We are investigating issues for agents api and computer use.
- `2026-10-06 20:00` **resolved** — All impacted services have now fully recovered.

### 32. circleci — Elevated wait times for machine jobs

- URL: https://stspg.io/vcsrgfwdsy4z
- Started: 2026-10-08 13:00 UTC
- Resolved: 2026-10-08 16:35 UTC
- Duration: 3h 35m
- Impact: minor, status: postmortem
- Components: Machine Jobs
- Canonical incident: [ci-363019f9a151bb752e2ec0ea](#ci-363019f9a151bb752e2ec0ea)

- `2026-10-08 13:00` **identified** — Customers may be experiencing elevated wait times for machine jobs. We are working to resolve this.
- `2026-10-08 13:34` **identified** — Customers using Linux machine jobs are experiencing elevated wait times, averaging about 11 minutes, with the longest waits exceeding 30 minutes on the medium, arm.medium and arm.large resource classes. Our engineers have identified the issue and are working on a fix. We will provide another update by 14:00 UTC.
- `2026-10-08 14:04` **identified** — Customers using Linux machine jobs are experiencing elevated wait times, averaging about 40 minutes, with the longest waits exceeding 50 minutes. Most Linux machine resource classes are affected, including medium, large, xlarge, 2xlarge and their Arm equivalents. Our engineers have identified the issue and are working on a fix. We will provide another update by 14:30 UTC.
- `2026-10-08 14:37` **identified** — Customers using Linux machine jobs and remote Docker are experiencing elevated wait times. Wait times have started to decrease and now average about 20 minutes, with the longest waits exceeding 40 minutes on some resource classes. Our engineers have identified the issue and are working on a fix. We will provide another update by 15:00 UTC.
- `2026-10-08 15:01` **identified** — A fix has been deployed and wait times are decreasing, but customers using Linux machine jobs and Remote Docker are still experiencing delays. Wait times currently average about 11 minutes, with the longest waits exceeding 35 minutes on some resource classes. We are working to add capacity as quickly as possible. We will provide another update by 15:30 UTC.
- `2026-10-08 15:31` **identified** — Wait times continue to decrease, but customers using Linux machine jobs and remote Docker are still experiencing delays. Wait times currently average about 6 minutes. The longest waits, up to about 20 minutes, are on the 2xlarge, arm.2xlarge and gpu.nvidia.small resource classes. We are working to add capacity as quickly as possible. We will provide another update by 16:00 UTC.
- `2026-10-08 15:59` **identified** — Wait times continue to decrease, but customers using Linux machine jobs and remote Docker are still experiencing delays. Wait times currently average about 90 seconds, with the longest waits up to about 6 minutes on some resource classes. We are working to add capacity as quickly as possible. We will provide another update by 16:30 UTC.
- `2026-10-08 16:09` **monitoring** — Wait times for customers using Linux machine jobs and remote Docker have returned to normal. We are monitoring to confirm wait times remain stable while we continue to add capacity. We will provide another update by 16:30 UTC.
- `2026-10-08 16:35` **resolved** — Between 12:40 UTC and 16:00 UTC on October 8, customers using Linux machine jobs and remote Docker experienced elevated wait times. The issue has been resolved and wait times have returned to normal. We thank you for your patience while our team worked on implementing a fix.
- `2026-10-08 17:54` **postmortem** — ## **Summary** Between October 1 and October 8, 2026, CircleCI customers experienced five separate incidents in which jobs waited longer than normal to start, or failed to start. Three of the incidents were driven by limited instance availability from our cloud provider. One was caused by a configuration issue in our scheduling infrastructure, and one by memory limits on an internal service, which we raised during the incident. * **October 1, 07:57 to 10:24 UTC:** Docker Gen 2 jobs waited up to 30 minutes to start. A configuration issue prevented the Gen 2 scheduling servers from using their high-performance local disks. During a burst in traffic, the fallback disks saturated and job placement slowed. * **October 7, 13:27 to 14:06 UTC:** Machine jobs waited longer than normal because we could not acquire enough virtual machines from our cloud provider. * **October 7, 14:24 to 15:21 UTC:** Some jobs failed with an infrastructure error, and workflows were delayed. The service that gives each job its configuration at startup ran out of memory and restarted. * **October 7, 18:49 to 20:05 UTC:** Docker Gen 2 jobs waited longer than normal because we could not acquire enough instances from our cloud provider. * **October 8, 12:48 to 16:00 UTC:** Linux machine jobs and remote Docker jobs waited up to 50 minutes to start. Our cloud provider did not have enough of the instance type we request, and a defect in the logic that picks a different instance type kept us from using alternatives. Jobs that waited during these incidents were delayed, not lost. Customers whose jobs failed with an infrastructure error during the October 7 incident from 14:24 to 15:21 UTC can rerun them. The original status pages can be found below: * [October 1 - Delays starting Gen 2 Docker Jobs](https://status.circleci.com/incidents/22mrv46g6n7g) * [October 7 - Delay on starting Machine Job Tasks](https://status.circleci.com/incidents/tj6wdnfjwr64) * [October 7 - Elevated level of infra fails on customer jobs](https://status.circleci.com/incidents/0z31ldjd8m5v) * [October 7 - Increased task wait times for Docker Gen2](https://status.circleci.com/incidents/39rfrjmk5pvd) * [October 8 - Elevated wait times for machine jobs](https://status.circleci.com/incidents/jn756xtsy4xg) ## **Background** Every job that runs on CircleCI needs compute. Where that compute comes from depends on the executor. Docker jobs run on clusters of servers that schedule work onto a fleet of instances. The scheduler is the software on those servers that decides which instance runs each job. Gen 2 Docker jobs ran on one such cluster at the time of the October 1 incident. Gen 1 Docker jobs are spread across several clusters. Machine jobs, including remote Docker jobs, get a dedicated virtual machine. A remote Docker job is handled as a Linux machine job from submission through provisioning. Our machine provisioning service requests each machine from our cloud provider in one of the two regions we normally use for machine jobs. When the provider cannot supply the instance type we request, the service falls back: it requests a different instance type, or a different region, from a list we maintain. A job waits until a machine is available. When any job starts, it asks a job configuration service for its configuration. That service runs as a set of pods, which are small copies of the same program. If too many pods stop, the remaining pods carry more load, and jobs cannot start until the service recovers. These paths depend on three kinds of capacity: the throughput of our scheduling systems, the supply of instances from our cloud provider, and the memory and number of copies of the services that start jobs. A shortfall in any of them makes jobs wait or fail. ## **What Happened** _\(All times UTC\)_ ### **October 1: Docker Gen 2 jobs delayed up to 30 minutes** At 07:44, disk activity on the servers that schedule Gen 2 Docker jobs rose to 90% of capacity and stayed there. At 07:57, the number of Docker jobs waiting to start began to grow, and wait times rose with it. An automated alert paged the on-call engineer at 08:08. We declared an incident at 08:42. At 08:43, engineers identified high disk latency on the scheduling servers. They first believed the disks were too slow for the load, so they prepared a change to a faster disk type. At 09:01, the team agreed on two workstreams: move the servers to faster disks, and add a second Gen 2 cluster so that jobs would split across two clusters. At 09:41, engineers found the root cause. Each scheduling server has fast disks attached directly to it, and the scheduler is meant to keep its working data there. These servers run on a security-hardened operating system image, and on that image the fast disks were not being used. The scheduler wrote its data to slower network-attached disks instead. The script that sets up the fast disks failed without an error, and a gap in logging on these servers hid the failure. Gen 2 traffic had also become increasingly bursty. The slower disks saturated during those bursts, and the scheduler fell behind. Wait times peaked at 09:54, when jobs on the Gen 2 medium resource class waited up to 30 minutes. At 09:55, engineers applied a fix that set up the fast disks correctly. At 10:11, the scheduler began placing jobs faster, which confirmed the servers could keep up with new jobs. Wait times for Gen 2 medium fell to 13 minutes at 10:14, 7 minutes at 10:17, and 1 minute 30 seconds at 10:20. We moved the status page to monitoring at 10:24 and resolved the incident at 10:45. At 11:20, the second Gen 2 cluster began running jobs. ### **October 7, 13:27 to 14:06: Elevated wait times for machine jobs** We declared an incident at 13:27 after seeing elevated queueing and wait times. We posted to the status page at 13:31. At 13:31, engineers saw that we were struggling to acquire virtual machines for machine jobs, and that the number of jobs waiting to start was high mainly for machine jobs. At 13:34, wait times were elevated across executors. At 13:48, wait times began to recover as more machines became available, and we resolved the incident at 14:06. ### **October 7, 14:24 to 15:21: Job failures and workflow delays** Shortly after the previous incident was marked as resolved, we declared a new incident at 14:24 due to a sharp drop in the number of jobs submitted to our execution systems. At 14:26, engineers saw that pods of the job configuration service were restarting. Pods that restart leave fewer copies of the service to answer requests, and jobs that could not retrieve their configuration after repeated attempts failed with an infrastructure error. Workflows that were already running waited for updates from jobs that had not started. At 14:35, engineers increased the amount of available memory to the pods and increased the number of pods, which gave the service more copies to share the load. We also increased the amount of memory available Job submissions began recovering, and the number of messages flowing through our workflow system returned to normal levels. At 14:44, we moved the status page to monitoring. Pods in the job configuration service reached their memory limits and restarted. We increased memory allocations and pod counts, and the service recovered. At 15:13, we increased the number of instances in our Gen 2 Docker clusters so we could work through the backlog of waiting jobs quickly. We resolved the incident at 15:21. ### **October 7, 18:49 to 20:05: Elevated wait times for Docker Gen 2 jobs** We declared a new incident at 18:49 after Docker Gen 2 wait times rose. At 18:51, engineers confirmed that a shortage of available instances kept us from acquiring enough of them, which raised wait times. Wait times differed by resource class size. At 19:52, average wait times were lower than an hour earlier as more capacity became available, though brief spikes continued, and we moved the status page to monitoring. We resolved the incident at 20:05. ### **October 8: Linux machine and remote Docker jobs wait up to 50 minutes** We declared an incident at 12:48 after seeing a large number of jobs waiting to start and long wait times for machine and remote Docker jobs. At 12:49, engineers saw that our cloud provider could not supply enough instances. Delays continued to climb as we investigated. At 13:13, we began examining the logic that chooses a different instance type or region when our first choice is unavailable. At 13:24, we suspected that a defect in that logic was stopping us from using other instance types in our primary region, so we disabled our fallback region and confirmed the cause of the defect. Engineers began working on a fix and deployed it at 14:49. The service began requesting other instance types and started machines in large numbers. At 14:58, the queue of jobs waiting for a machine dropped steadily. Average wait times across resource classes continued to drop over the next hour as we worked through the backlog of waiting jobs. At 16:00, the queue of jobs waiting for a machine cleared and the number of waiting jobs returned to normal. We moved the status page to monitoring at 16:09 and resolved the incident at 16:35. ## **Future Prevention and Process Improvement** We are taking the following steps to strengthen the resilience of job start times, including supply diversification, fallback logic, and earlier detection. **We are expanding our compute regions.** In the October 7 and October 8 incidents, our cloud provider could not supply enough instances in the regions we used. We are adding new regions to our fallback logic in order handle spikes for demand that outpace our preferred regions. **We are fixing the logic that picks a different instance type.** On October 8, a defect kept the machine provisioning service from using other instance types in the same region. We deployed a fix during the incident and are now reviewing and refactoring this logic as a whole, with tests, so that a shortage of one instance type moves work to another. **We are expanding the instance types we offer and securing additional capacity.** We are working with our cloud provider to reserve additional capacity in our existing regions, and we are adding instance types that jobs can run on. Industry demand for high-performance compute is growing. We are diversifying where and how we source capacity so our customers are insulated from supply variation. **We added a second Gen 2 Docker cluster so traffic bursts are distributed across two clusters.** We are also adding alerts on scheduling server disk saturation and on how fast the scheduler places jobs, and an alert that fires when a server starts without its fast disks in use. We made the disk setup script fail with an error, and we restored log shipping from these servers so that a similar problem is visible sooner. Customer experience is our top priority, and we commit to continually improving the reliability of our systems to match the trust that our customers place in us. We thank our customers for their patience while our team worked to resolve these incidents. Please reach out to our support team with any questions or concerns.

### 33. openai — Errors processing large PDFs in the Responses API

- URL: https://status.openai.com/incidents/01M48Y8E81QAFF9XKVGE118E6X
- Started: 2026-10-06 15:45 UTC
- Resolved: 2026-10-06 17:50 UTC
- Duration: 2h 4m
- Impact: minor, status: resolved
- Components: —
- Canonical incident: [ci-899fdc48100f185284250152](#ci-899fdc48100f185284250152)

- `2026-10-06 15:45` **investigating** — We are investigating the issue for the listed services.
- `2026-10-06 16:39` **identified** — We are implementing the mitigation.
- `2026-10-06 20:11` **resolved** — All impacted services have now fully recovered.

### 34. github — Disruption with some GitHub services

- URL: https://stspg.io/pjcfc4c3pm7g
- Started: 2026-10-05 23:47 UTC
- Resolved: 2026-10-06 01:32 UTC
- Duration: 1h 44m
- Impact: major, status: resolved
- Components: —
- Canonical incident: [ci-14bbdb741d4488be842edc30](#ci-14bbdb741d4488be842edc30)

- `2026-10-05 23:47` **investigating** — We are investigating reports of impacted performance for some GitHub services.
- `2026-10-06 00:05` **investigating** — Users may experience issues accessing organization and enterprise billing and licensing pages. We are deploying a mitigation and monitoring for recovery.
- `2026-10-06 01:32` **resolved** — On October 5, 2026, between 23:05 UTC and 00:11 UTC on October 6, some users experienced slow page loads, timeouts, and errors when accessing organization and enterprise billing, licensing, and settings pages. On average, the error rate was 1.9% and peaked at 13.8% of requests to the service that powers these pages.<br /><br />This was caused by a networking failure, which disrupted connectivity between one of our on-premises datacenter regions and a subset of cloud-hosted regional database services. We mitigated the incident by shifting the database traffic to a healthy region.<br /><br />We are working to make regional database failover faster and more complete, reduce timeouts so these pages fail fast instead of hanging, and add latency-based alerting to reduce our time to detection and mitigation of issues like this one in the future. We continue to partner with our cloud provider to understand the underlying network failure.

### 35. openai — Workspace admins may be unable to manage plugins

- URL: https://status.openai.com/incidents/01M4HA7BCCXVHTR5DCYMSBQF2H
- Started: 2026-10-09 21:48 UTC
- Resolved: 2026-10-09 23:26 UTC
- Duration: 1h 38m
- Impact: major, status: resolved
- Components: —
- Canonical incident: [ci-34efd4facf7251208fbe1394](#ci-34efd4facf7251208fbe1394)

- `2026-10-09 21:48` **investigating** — We are investigating an issue affecting some ChatGPT workspace administrators’ ability to view and manage plugin availability and settings. Some users may also see plugins as unavailable. We will provide another update as soon as we have more information.
- `2026-10-09 22:52` **monitoring** — We are seeing recovery in workspace administrators’ ability to view and manage plugin availability and settings. We are monitoring to ensure service remains stable.
- `2026-10-09 23:26` **resolved** — All impacted services have now fully recovered.

### 36. anthropic — Elevated errors loading usage data on platform.claude.com

- URL: https://stspg.io/scdhhnplc2jn
- Started: 2026-10-07 01:20 UTC
- Resolved: 2026-10-07 02:33 UTC
- Duration: 1h 13m
- Impact: major, status: resolved
- Components: Claude Console (platform.claude.com)
- Canonical incident: [ci-adbaab36de2e68a0ca1a88ff](#ci-adbaab36de2e68a0ca1a88ff)

- `2026-10-07 01:20` **investigating** — We are investigating elevated errors loading usage data on the platform.claude.com Usage page and the Admin API usage report endpoint since 5 PM PT. Messages API requests are not affected. We will provide an update as soon as possible.
- `2026-10-07 02:33` **resolved** — This incident has been resolved. From 17:00 PT on October 6 / 00:00 UTC on October 7 to 18:40 PT on October 6 / 01:40 UTC on October 7, many requests to load usage data on the platform.claude.com Usage page and the Admin API usage report returned errors.

### 37. anthropic — Elevated errors for Claude Opus 5.5

- URL: https://stspg.io/bzc81gbv1b99
- Started: 2026-10-06 12:24 UTC
- Resolved: 2026-10-06 12:43 UTC
- Duration: 18m
- Impact: major, status: resolved
- Components: Claude API (api.anthropic.com), Claude Code, Claude Cowork, claude.ai
- Canonical incident: [ci-2c0e3710114f15eb32a62531](#ci-2c0e3710114f15eb32a62531)

- `2026-10-06 12:24` **investigating** — We are investigating elevated errors on requests to Claude Opus 5.5. We will provide an update as soon as possible.
- `2026-10-06 12:37` **identified** — We have identified the cause of elevated errors on requests to Claude Opus 5.5 and are working on a fix. We will provide an update as soon as possible.
- `2026-10-06 12:43` **resolved** — The issue affecting Claude Opus 5.5 has been resolved.

### 38. anthropic — Elevated errors for Claude Mythos 5.1 and Claude Fable 5.1

- URL: https://stspg.io/wdqhxt8xt574
- Started: 2026-10-05 13:05 UTC
- Resolved: 2026-10-05 13:21 UTC
- Duration: 15m
- Impact: major, status: resolved
- Components: Claude API (api.anthropic.com), Claude Code, Claude Cowork, claude.ai
- Canonical incident: [ci-b2a29683e064e606703b3672](#ci-b2a29683e064e606703b3672)

- `2026-10-05 13:05` **investigating** — We are investigating elevated errors on requests to Claude Mythos 5.1 and Claude Fable 5.1. We will provide an update as soon as possible.
- `2026-10-05 13:21` **resolved** — The issue affecting Claude Mythos 5.1 and Claude Fable 5.1 has been resolved. Impact occurred from 05:40 PT / 12:40 UTC to 06:10 PT / 13:10 UTC.

### 39. github — Incident with Git Operations, Pull Requests and Actions

- URL: https://stspg.io/96smrcth8bpg
- Started: 2026-10-07 16:25 UTC
- Resolved: 2026-10-07 16:25 UTC
- Duration: 0m
- Impact: critical, status: resolved
- Components: Actions, Git Operations, Issues, Pull Requests, Webhooks
- Canonical incident: [ci-20cff0acc195e836836e1e49](#ci-20cff0acc195e836836e1e49)

- `2026-10-07 15:14` **investigating** — We are investigating reports of degraded availability for Actions, Git Operations and Pull Requests
- `2026-10-07 15:14` **investigating** — Webhooks is experiencing degraded performance. We are continuing to investigate.
- `2026-10-07 15:15` **investigating** — Issues is experiencing degraded performance. We are continuing to investigate.
- `2026-10-07 15:16` **investigating** — Webhooks is experiencing degraded availability. We are continuing to investigate.
- `2026-10-07 15:21` **investigating** — Git Operations is experiencing degraded performance. We are continuing to investigate.
- `2026-10-07 15:24` **investigating** — We observed widespread impact across GitHub services between 15:06 UTC and 15:16 UTC, and are investigating the root cause. We are seeing all systems recover, and will provide another update shortly.
- `2026-10-07 15:25` **investigating** — Pull Requests is experiencing degraded performance. We are continuing to investigate.
- `2026-10-07 15:27` **investigating** — Git Operations is operating normally.
- `2026-10-07 15:27` **investigating** — Webhooks is experiencing degraded performance. We are continuing to investigate.
- `2026-10-07 15:31` **investigating** — The degradation affecting Pull Requests has been mitigated. We are monitoring to ensure stability.
- `2026-10-07 15:32` **investigating** — Actions is operating normally.
- `2026-10-07 15:35` **investigating** — Git Operations is experiencing degraded performance. We are continuing to investigate.
- `2026-10-07 15:41` **investigating** — We are observing partial recovery across all systems. We are continuing to monitor and investigate the root cause of the disruption.
- `2026-10-07 15:47` **investigating** — The degradation affecting Issues has been mitigated. We are monitoring to ensure stability.
- `2026-10-07 15:49` **investigating** — The degradation affecting Git Operations and Webhooks has been mitigated. We are monitoring to ensure stability.
- `2026-10-07 15:49` **monitoring** — The degradation has been mitigated. We are monitoring to ensure stability.
- `2026-10-07 15:56` **monitoring** — The degradation has been mitigated. We are monitoring to ensure stability.
- `2026-10-07 15:58` **monitoring** — We are seeing full recovery across all systems including Git Operations, Pull Requests and Actions. We are continuing to investigate the root cause of the disruption between 15:06 UTC and 15:16 UTC, and will share another update shortly.
- `2026-10-07 16:23` **monitoring** — Git Operations, Pull Requests and Actions are fully recovered and stable. The root cause has been identified and mitigations are being applied to prevent a reoccurrence.
- `2026-10-07 16:25` **resolved** — On October 7, 2026, GitHub had two outages of write operations with the same cause, totaling about 19 minutes. The first ran from 15:06 to 15:16 UTC (about 10 minutes) and the second from 16:52 to 17:01 UTC (about 9 minutes). In both windows, actions that save or change data failed across Git operations, pull requests, issues, GitHub Actions, the API, webhooks, sign-in and authentication, GitHub Apps and OAuth, code scanning, Codespaces, and Copilot. Reading data was largely unaffected. In the first outage, 5.6% of all web and API requests failed, peaking at 8.9%, and more than 12,000 organizations saw failed requests. Up to 26% of Git authentication requests failed, and 17% of Actions workflow runs failed with infrastructure errors. In the second outage, 6.2% of web and API requests failed, peaking at 9.3%, and more than 11,000 organizations were affected. Up to 25% of Git authentication requests failed, and 19% of Actions workflow runs failed. No databases experienced data loss, and operations retried after recovery succeeded. GitHub relies on an internal service that keeps track of which database servers should receive new and updated data. While replacing machines that run this service as part of regular maintenance, one of the new machines started with an outdated copy of its records, which still listed database servers we had already retired. When this node came online, it incorrectly modified routing configuration with the outdated records, routing write requests to non-existent database servers. Each outage ended when one of our established machines with correct configuration became the leader of the quorum and sent the correct configuration. After the first outage, we identified the new machine as the source and paused automated maintenance to stop it from being leader again. At 16:52 UTC, an automated maintenance task that did not honor the pause put it back in charge, and the outage happened again. We shut that machine down at 17:02 UTC and have since removed it, and the automation now honors the pause, which was a new regression due to an update of our puppet infrastructure. To prevent recurrence, we are clearing retired database servers from the service's records before any new machine is added, and we are keeping new machines from being elected leader until their records have been checked. We are also changing our systems to reject configuration that point to servers that don't exist and to keep using the last known-good settings instead. Finally, we are giving responders a faster way to take a misbehaving machine out of service, and adding alerts so changes like this are spotted within seconds. We will share a more detailed write-up of these incidents in our monthly availability report, along with the status of the repairs we are making to prevent issues like these in the future.

## Below threshold (24)

- `cloudflare` [Network routing issue](https://www.cloudflarestatus.com/incidents/ryvhi2suec0d) — 1h 59m
- `openai` [Workspace Agents Degraded and Responses Impacted](https://status.openai.com/incidents/01M49YQE6WZT02WYF0KBXVMQ5G) — 1h 47m
- `openai` [Dot in Codex may fail to create new threads](https://status.openai.com/incidents/01M4BK9J2M83R4SSC71D9FXW45) — 1h 35m
- `openai` [Some APAC users are seeing errors in ChatGPT Work, conversations, and GPTs.](https://status.openai.com/incidents/01M4CR8GCAFC1RYQXZDXAZP2VT) — 1h 33m
- `cloudflare` [Network Performance Issues in Sofia](https://www.cloudflarestatus.com/incidents/9o6q1l6wwtmr) — 57m
- `openai` [Elevated error rates in ChatGPT Go conversations](https://status.openai.com/incidents/01M48PT6B8SDTW023NC9XZTBC0) — 1h 5m
- `openai` [Degraded performance in Codex and Work mode: Dot turn failures, delayed responses, and reduced proactivity](https://status.openai.com/incidents/01M4BP5JE7DKJSP3VG2S2DCWF2) — 1h 1m
- `openai` [Elevated errors in API Platform login and administration APIs.](https://status.openai.com/incidents/01M47AFBNRNXKJWN549T3PMYV8) — 55m
- `cloudflare` [RSA Keyless SSL handshake issues](https://www.cloudflarestatus.com/incidents/6mlprt9v5si5) — 44m
- `cloudflare` [Cloudflare API service issues](https://www.cloudflarestatus.com/incidents/p86jkvmrc3pc) — 43m
- `digitalocean` [VPC Networking Availability](https://stspg.io/1n5v2k9sxq46) — 1h 35m
- `yandexcloud` [The problem is with the Yandex Cloud Console services](https://status.yandex.cloud/en/incidents/2134) — 49m
- `circleci` [Increased task wait times for Docker Gen2](https://stspg.io/fpv4m66zyd77) — 1h 10m
- `digitalocean` [DNS Resolution for Managed Kubernetes](https://stspg.io/61n77v08qp0j) — 1h 27m
- `cloudflare` [Network Performance Issues in Queretaro](https://www.cloudflarestatus.com/incidents/60w284knplry) — 19m
- `circleci` [Elevated level of infra fails on customer jobs](https://stspg.io/n2l0044kmkdr) — 53m
- `openai` [Elevated errors affecting multiple products including ChatGPT (conversations), Image generations, Spaces and Pages, API rehydration, and Dots.](https://status.openai.com/incidents/01M47HC9SGHGFPBN44KGC6SKAQ) — 35m
- `openai` [Elevated errors with GPT-5.6 Instant in FedRAMP workspaces](https://status.openai.com/incidents/01M4CBXZ62A9SVH8D74KVYQC8M) — 21m
- `openai` [Issues saving and editing Pages](https://status.openai.com/incidents/01M461GYDMCVJJK8330PVCH42C) — 19m
- `circleci` [Delay on starting Machine Job Tasks](https://stspg.io/7q8hkm5sgt8h) — 26m
- `cloudflare` [Increase in Latency Across Multiple Locations](https://www.cloudflarestatus.com/incidents/y3hstd1lvhc9) — 0m
- `grafana` [Grafana Cloud k6 – Brief metrics unavailability](https://stspg.io/lw22c35nvf3l) — 0m
- `anthropic` [Issue with spend limits](https://stspg.io/frlrjsqqn605) — 0m
- `cloudflare` [Network Performance Issues in Kochi](https://www.cloudflarestatus.com/incidents/76m2gfyfx1xc) — 0m

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-d5585cf81abf4edcfde48cfd"></a>

- Canonical ID: `ci-d5585cf81abf4edcfde48cfd` (1 source records).
  - [yandexcloud/2136](../../data/2026/W41/yandexcloud.json); vendor URL: https://status.yandex.cloud/en/incidents/2136
- Start time basis: `vendor_reported`.

<a id="ci-8e0bde59a25eb3bfe75e60a3"></a>

- Canonical ID: `ci-8e0bde59a25eb3bfe75e60a3` (1 source records).
  - [yandexcloud/2092](../../data/2026/W41/yandexcloud.json); vendor URL: https://status.yandex.cloud/en/incidents/2092
- Start time basis: `vendor_reported`.

<a id="ci-dcb212550be05a0c4eb3a32a"></a>

- Canonical ID: `ci-dcb212550be05a0c4eb3a32a` (1 source records).
  - [cloudflare/pfvfug94jq4q](../../data/2026/W41/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/pfvfug94jq4q
- Start time basis: `vendor_reported`.

<a id="ci-aa6f59211a96cfbcca9feb49"></a>

- Canonical ID: `ci-aa6f59211a96cfbcca9feb49` (1 source records).
  - [grafana/6x77ccl9shpy](../../data/2026/W41/grafana.json); vendor URL: https://stspg.io/9l8htmqt82s2
- Start time basis: `vendor_reported`.

<a id="ci-1d7fe76c931dba8c5cb73366"></a>

- Canonical ID: `ci-1d7fe76c931dba8c5cb73366` (1 source records).
  - [cloudflare/cyxclg7nnhjy](../../data/2026/W41/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/cyxclg7nnhjy
- Start time basis: `vendor_reported`.

<a id="ci-f64b047b06bad8235f9ab606"></a>

- Canonical ID: `ci-f64b047b06bad8235f9ab606` (1 source records).
  - [anthropic/978mjkgw6mh7](../../data/2026/W41/anthropic.json); vendor URL: https://stspg.io/sjkw7njwf0mt
- Start time basis: `vendor_reported`.

<a id="ci-25c27c306bfe03faf64a026b"></a>

- Canonical ID: `ci-25c27c306bfe03faf64a026b` (1 source records).
  - [openai/01M4GV1M2A3K38T73JCVCD9G0R](../../data/2026/W41/openai.json); vendor URL: https://status.openai.com/incidents/01M4GV1M2A3K38T73JCVCD9G0R
- Start time basis: `vendor_reported`.

<a id="ci-057564c342d6384d89c544ba"></a>

- Canonical ID: `ci-057564c342d6384d89c544ba` (1 source records).
  - [slack/1580](../../data/2026/W41/slack.json); vendor URL: https://slack-status.com/2026-10/a0cfb0208a1d2032
- Start time basis: `vendor_reported`.

<a id="ci-74a2f2cec65976aed2992881"></a>

- Canonical ID: `ci-74a2f2cec65976aed2992881` (1 source records).
  - [gcp/mSFVt6PiP557bjRyYsut](../../data/2026/W41/gcp.json); vendor URL: https://status.cloud.google.com/incidents/mSFVt6PiP557bjRyYsut
- Start time basis: `vendor_reported`.

<a id="ci-c309a05b6bfa84541633c192"></a>

- Canonical ID: `ci-c309a05b6bfa84541633c192` (1 source records).
  - [grafana/91ct2lc9xfmn](../../data/2026/W41/grafana.json); vendor URL: https://stspg.io/cjbf6knm4jq4
- Start time basis: `vendor_reported`.

<a id="ci-52f863d412c65a5c1bbf1f11"></a>

- Canonical ID: `ci-52f863d412c65a5c1bbf1f11` (1 source records).
  - [cloudflare/z28czd65h29l](../../data/2026/W41/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/z28czd65h29l
- Start time basis: `vendor_reported`.

<a id="ci-7e70d415c6e2be105272e1c1"></a>

- Canonical ID: `ci-7e70d415c6e2be105272e1c1` (1 source records).
  - [cloudflare/tt8f4ttcuonf](../../data/2026/W41/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/tt8f4ttcuonf
- Start time basis: `vendor_reported`.

<a id="ci-fb8b1117166bd38c3550921f"></a>

- Canonical ID: `ci-fb8b1117166bd38c3550921f` (1 source records).
  - [circleci/6cfhhlr4rmm0](../../data/2026/W41/circleci.json); vendor URL: https://stspg.io/hkfkwz292bz8
- Start time basis: `vendor_reported`.

<a id="ci-16c02e0e6cbe0a1b216975a4"></a>

- Canonical ID: `ci-16c02e0e6cbe0a1b216975a4` (1 source records).
  - [openai/01M45E0JFFJZ6EQ3BWYPW74R2Z](../../data/2026/W41/openai.json); vendor URL: https://status.openai.com/incidents/01M45E0JFFJZ6EQ3BWYPW74R2Z
- Start time basis: `vendor_reported`.

<a id="ci-2c1e26dd8f2e3ec55cffa8a8"></a>

- Canonical ID: `ci-2c1e26dd8f2e3ec55cffa8a8` (1 source records).
  - [cloudflare/4knp3pzs3qvx](../../data/2026/W41/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/4knp3pzs3qvx
- Start time basis: `vendor_reported`.

<a id="ci-b9d6e1019a1436b9ecb42027"></a>

- Canonical ID: `ci-b9d6e1019a1436b9ecb42027` (1 source records).
  - [yandexcloud/2130](../../data/2026/W41/yandexcloud.json); vendor URL: https://status.yandex.cloud/en/incidents/2130
- Start time basis: `vendor_reported`.

<a id="ci-f49247b63ccb439cf646ba98"></a>

- Canonical ID: `ci-f49247b63ccb439cf646ba98` (1 source records).
  - [cloudflare/ts295zqycqn7](../../data/2026/W41/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/ts295zqycqn7
- Start time basis: `vendor_reported`.

<a id="ci-06d5bce0f029e3c9c82e3c57"></a>

- Canonical ID: `ci-06d5bce0f029e3c9c82e3c57` (1 source records).
  - [cloudflare/x8n12150xwf6](../../data/2026/W41/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/x8n12150xwf6
- Start time basis: `vendor_reported`.

<a id="ci-81b00cc38adda810f181cd6d"></a>

- Canonical ID: `ci-81b00cc38adda810f181cd6d` (1 source records).
  - [github/cjwszcf57bff](../../data/2026/W41/github.json); vendor URL: https://stspg.io/w3rcssslq22w
- Start time basis: `vendor_reported`.

<a id="ci-b531d0fe442ab4c7a92d0e1d"></a>

- Canonical ID: `ci-b531d0fe442ab4c7a92d0e1d` (1 source records).
  - [github/3q1yb5m7ltvb](../../data/2026/W41/github.json); vendor URL: https://stspg.io/c11dc9nb1zdq
- Start time basis: `vendor_reported`.

<a id="ci-266fbf9dfabc9fed11b692e4"></a>

- Canonical ID: `ci-266fbf9dfabc9fed11b692e4` (1 source records).
  - [circleci/z4hbq8h49j8c](../../data/2026/W41/circleci.json); vendor URL: https://stspg.io/s7tpg47srxdd
- Start time basis: `vendor_reported`.

<a id="ci-f6f4424427f4046d5e2cc46a"></a>

- Canonical ID: `ci-f6f4424427f4046d5e2cc46a` (1 source records).
  - [openai/01M46G248GGHV9RSAJE82BZY1X](../../data/2026/W41/openai.json); vendor URL: https://status.openai.com/incidents/01M46G248GGHV9RSAJE82BZY1X
- Start time basis: `vendor_reported`.

<a id="ci-58fe92f1c3a4fa125c4de798"></a>

- Canonical ID: `ci-58fe92f1c3a4fa125c4de798` (1 source records).
  - [openai/01M4ADHMMBK7K92NJFJ0DXRYJM](../../data/2026/W41/openai.json); vendor URL: https://status.openai.com/incidents/01M4ADHMMBK7K92NJFJ0DXRYJM
- Start time basis: `vendor_reported`.

<a id="ci-9712f4ec943d180f692b6acb"></a>

- Canonical ID: `ci-9712f4ec943d180f692b6acb` (1 source records).
  - [openai/01M46KED6KQ6PR195KZYYVNY8W](../../data/2026/W41/openai.json); vendor URL: https://status.openai.com/incidents/01M46KED6KQ6PR195KZYYVNY8W
- Start time basis: `vendor_reported`.

<a id="ci-a1a70a6819b7b9547baf2129"></a>

- Canonical ID: `ci-a1a70a6819b7b9547baf2129` (1 source records).
  - [openai/01M46C47WYYD9NMZVGH83X3ZR1](../../data/2026/W41/openai.json); vendor URL: https://status.openai.com/incidents/01M46C47WYYD9NMZVGH83X3ZR1
- Start time basis: `vendor_reported`.

<a id="ci-a3dd9b08db53546229d15e29"></a>

- Canonical ID: `ci-a3dd9b08db53546229d15e29` (1 source records).
  - [stripe/0zs9ncsr6t0l](../../data/2026/W41/stripe.json); vendor URL: https://stspg.io/m7j171ppknbs
- Start time basis: `vendor_reported`.

<a id="ci-3a323695409125036a7a6913"></a>

- Canonical ID: `ci-3a323695409125036a7a6913` (1 source records).
  - [cloudflare/xczvsioz1un8](../../data/2026/W41/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/xczvsioz1un8
- Start time basis: `vendor_reported`.

<a id="ci-1504dfe18c7a0d6c7d36211f"></a>

- Canonical ID: `ci-1504dfe18c7a0d6c7d36211f` (1 source records).
  - [openai/01M4H7FKXDA425TDNP2CRGEYA3](../../data/2026/W41/openai.json); vendor URL: https://status.openai.com/incidents/01M4H7FKXDA425TDNP2CRGEYA3
- Start time basis: `vendor_reported`.

<a id="ci-8a0991181da00ff1f8976d77"></a>

- Canonical ID: `ci-8a0991181da00ff1f8976d77` (1 source records).
  - [github/qpfv5p86dmrl](../../data/2026/W41/github.json); vendor URL: https://stspg.io/4p8x0jnnsx7m
- Start time basis: `vendor_reported`.

<a id="ci-fd0e36afed5adea81c871030"></a>

- Canonical ID: `ci-fd0e36afed5adea81c871030` (1 source records).
  - [yandexcloud/2124](../../data/2026/W41/yandexcloud.json); vendor URL: https://status.yandex.cloud/en/incidents/2124
- Start time basis: `vendor_reported`.

<a id="ci-ef5d020ca1dcb8d3aca305e0"></a>

- Canonical ID: `ci-ef5d020ca1dcb8d3aca305e0` (1 source records).
  - [openai/01M493CZ0CB14KPC0EER83MNVB](../../data/2026/W41/openai.json); vendor URL: https://status.openai.com/incidents/01M493CZ0CB14KPC0EER83MNVB
- Start time basis: `vendor_reported`.

<a id="ci-363019f9a151bb752e2ec0ea"></a>

- Canonical ID: `ci-363019f9a151bb752e2ec0ea` (1 source records).
  - [circleci/jn756xtsy4xg](../../data/2026/W41/circleci.json); vendor URL: https://stspg.io/vcsrgfwdsy4z
- Start time basis: `vendor_reported`.

<a id="ci-02341b4e8a9e7f91bb4d84d1"></a>

- Canonical ID: `ci-02341b4e8a9e7f91bb4d84d1` (1 source records).
  - [cloudflare/ryvhi2suec0d](../../data/2026/W41/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/ryvhi2suec0d
- Start time basis: `vendor_reported`.

<a id="ci-899fdc48100f185284250152"></a>

- Canonical ID: `ci-899fdc48100f185284250152` (1 source records).
  - [openai/01M48Y8E81QAFF9XKVGE118E6X](../../data/2026/W41/openai.json); vendor URL: https://status.openai.com/incidents/01M48Y8E81QAFF9XKVGE118E6X
- Start time basis: `vendor_reported`.

<a id="ci-6ec39652fd41bab024b58cae"></a>

- Canonical ID: `ci-6ec39652fd41bab024b58cae` (1 source records).
  - [openai/01M49YQE6WZT02WYF0KBXVMQ5G](../../data/2026/W41/openai.json); vendor URL: https://status.openai.com/incidents/01M49YQE6WZT02WYF0KBXVMQ5G
- Start time basis: `vendor_reported`.

<a id="ci-14bbdb741d4488be842edc30"></a>

- Canonical ID: `ci-14bbdb741d4488be842edc30` (1 source records).
  - [github/2kxnqcxcrr7f](../../data/2026/W41/github.json); vendor URL: https://stspg.io/pjcfc4c3pm7g
- Start time basis: `vendor_reported`.

<a id="ci-34efd4facf7251208fbe1394"></a>

- Canonical ID: `ci-34efd4facf7251208fbe1394` (1 source records).
  - [openai/01M4HA7BCCXVHTR5DCYMSBQF2H](../../data/2026/W41/openai.json); vendor URL: https://status.openai.com/incidents/01M4HA7BCCXVHTR5DCYMSBQF2H
- Start time basis: `vendor_reported`.

<a id="ci-4a3bf32092b0619fc46879ba"></a>

- Canonical ID: `ci-4a3bf32092b0619fc46879ba` (1 source records).
  - [openai/01M4BK9J2M83R4SSC71D9FXW45](../../data/2026/W41/openai.json); vendor URL: https://status.openai.com/incidents/01M4BK9J2M83R4SSC71D9FXW45
- Start time basis: `vendor_reported`.

<a id="ci-f4f73ae0c6db3f2be69e4fdd"></a>

- Canonical ID: `ci-f4f73ae0c6db3f2be69e4fdd` (1 source records).
  - [openai/01M4CR8GCAFC1RYQXZDXAZP2VT](../../data/2026/W41/openai.json); vendor URL: https://status.openai.com/incidents/01M4CR8GCAFC1RYQXZDXAZP2VT
- Start time basis: `vendor_reported`.

<a id="ci-adbaab36de2e68a0ca1a88ff"></a>

- Canonical ID: `ci-adbaab36de2e68a0ca1a88ff` (1 source records).
  - [anthropic/wf2v8ms031sx](../../data/2026/W41/anthropic.json); vendor URL: https://stspg.io/scdhhnplc2jn
- Start time basis: `vendor_reported`.

<a id="ci-2c0e3710114f15eb32a62531"></a>

- Canonical ID: `ci-2c0e3710114f15eb32a62531` (1 source records).
  - [anthropic/ch27pb90bn85](../../data/2026/W41/anthropic.json); vendor URL: https://stspg.io/bzc81gbv1b99
- Start time basis: `vendor_reported`.

<a id="ci-bd9b218ed67aa993339cb8d8"></a>

- Canonical ID: `ci-bd9b218ed67aa993339cb8d8` (1 source records).
  - [cloudflare/9o6q1l6wwtmr](../../data/2026/W41/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/9o6q1l6wwtmr
- Start time basis: `vendor_reported`.

<a id="ci-796ed797e406491fac8add80"></a>

- Canonical ID: `ci-796ed797e406491fac8add80` (1 source records).
  - [openai/01M48PT6B8SDTW023NC9XZTBC0](../../data/2026/W41/openai.json); vendor URL: https://status.openai.com/incidents/01M48PT6B8SDTW023NC9XZTBC0
- Start time basis: `vendor_reported`.

<a id="ci-4daeb9178f091494d45f763a"></a>

- Canonical ID: `ci-4daeb9178f091494d45f763a` (1 source records).
  - [openai/01M4BP5JE7DKJSP3VG2S2DCWF2](../../data/2026/W41/openai.json); vendor URL: https://status.openai.com/incidents/01M4BP5JE7DKJSP3VG2S2DCWF2
- Start time basis: `vendor_reported`.

<a id="ci-b2a29683e064e606703b3672"></a>

- Canonical ID: `ci-b2a29683e064e606703b3672` (1 source records).
  - [anthropic/bhphxz3vr58g](../../data/2026/W41/anthropic.json); vendor URL: https://stspg.io/wdqhxt8xt574
- Start time basis: `vendor_reported`.

<a id="ci-c3101d05c5977ddc4955d693"></a>

- Canonical ID: `ci-c3101d05c5977ddc4955d693` (1 source records).
  - [openai/01M47AFBNRNXKJWN549T3PMYV8](../../data/2026/W41/openai.json); vendor URL: https://status.openai.com/incidents/01M47AFBNRNXKJWN549T3PMYV8
- Start time basis: `vendor_reported`.

<a id="ci-18a1092c7e252f7b1e5cbb70"></a>

- Canonical ID: `ci-18a1092c7e252f7b1e5cbb70` (1 source records).
  - [cloudflare/6mlprt9v5si5](../../data/2026/W41/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/6mlprt9v5si5
- Start time basis: `vendor_reported`.

<a id="ci-10ddaa6b07418af4b2f23265"></a>

- Canonical ID: `ci-10ddaa6b07418af4b2f23265` (1 source records).
  - [cloudflare/p86jkvmrc3pc](../../data/2026/W41/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/p86jkvmrc3pc
- Start time basis: `vendor_reported`.

<a id="ci-1c33748f558f404f7e9229a4"></a>

- Canonical ID: `ci-1c33748f558f404f7e9229a4` (1 source records).
  - [digitalocean/5fxykytg48dx](../../data/2026/W41/digitalocean.json); vendor URL: https://stspg.io/1n5v2k9sxq46
- Start time basis: `vendor_reported`.

<a id="ci-aa1f60a734a82fc0ce00fa85"></a>

- Canonical ID: `ci-aa1f60a734a82fc0ce00fa85` (1 source records).
  - [yandexcloud/2134](../../data/2026/W41/yandexcloud.json); vendor URL: https://status.yandex.cloud/en/incidents/2134
- Start time basis: `vendor_reported`.

<a id="ci-d96807d0b516c785be4572ad"></a>

- Canonical ID: `ci-d96807d0b516c785be4572ad` (1 source records).
  - [circleci/39rfrjmk5pvd](../../data/2026/W41/circleci.json); vendor URL: https://stspg.io/fpv4m66zyd77
- Start time basis: `vendor_reported`.

<a id="ci-63e835b12afc098814d7ccd5"></a>

- Canonical ID: `ci-63e835b12afc098814d7ccd5` (1 source records).
  - [digitalocean/qfg5hmfhxw5d](../../data/2026/W41/digitalocean.json); vendor URL: https://stspg.io/61n77v08qp0j
- Start time basis: `vendor_reported`.

<a id="ci-adea61d0eff07c44cbd48e80"></a>

- Canonical ID: `ci-adea61d0eff07c44cbd48e80` (1 source records).
  - [cloudflare/60w284knplry](../../data/2026/W41/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/60w284knplry
- Start time basis: `vendor_reported`.

<a id="ci-069d54eb78d7605c97ec4b97"></a>

- Canonical ID: `ci-069d54eb78d7605c97ec4b97` (1 source records).
  - [circleci/0z31ldjd8m5v](../../data/2026/W41/circleci.json); vendor URL: https://stspg.io/n2l0044kmkdr
- Start time basis: `vendor_reported`.

<a id="ci-641698e61525ab845ba33b13"></a>

- Canonical ID: `ci-641698e61525ab845ba33b13` (1 source records).
  - [openai/01M47HC9SGHGFPBN44KGC6SKAQ](../../data/2026/W41/openai.json); vendor URL: https://status.openai.com/incidents/01M47HC9SGHGFPBN44KGC6SKAQ
- Start time basis: `vendor_reported`.

<a id="ci-549fd707e25d355d14ccce79"></a>

- Canonical ID: `ci-549fd707e25d355d14ccce79` (1 source records).
  - [openai/01M4CBXZ62A9SVH8D74KVYQC8M](../../data/2026/W41/openai.json); vendor URL: https://status.openai.com/incidents/01M4CBXZ62A9SVH8D74KVYQC8M
- Start time basis: `vendor_reported`.

<a id="ci-c9b20e7d3a18b281a171aab9"></a>

- Canonical ID: `ci-c9b20e7d3a18b281a171aab9` (1 source records).
  - [openai/01M461GYDMCVJJK8330PVCH42C](../../data/2026/W41/openai.json); vendor URL: https://status.openai.com/incidents/01M461GYDMCVJJK8330PVCH42C
- Start time basis: `vendor_reported`.

<a id="ci-936940e1fd341f29c694c0fb"></a>

- Canonical ID: `ci-936940e1fd341f29c694c0fb` (1 source records).
  - [circleci/tj6wdnfjwr64](../../data/2026/W41/circleci.json); vendor URL: https://stspg.io/7q8hkm5sgt8h
- Start time basis: `vendor_reported`.

<a id="ci-da786a534281572ae546b136"></a>

- Canonical ID: `ci-da786a534281572ae546b136` (1 source records).
  - [cloudflare/y3hstd1lvhc9](../../data/2026/W41/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/y3hstd1lvhc9
- Start time basis: `vendor_reported`.

<a id="ci-5b8bf8a7409d5f4e80f3db41"></a>

- Canonical ID: `ci-5b8bf8a7409d5f4e80f3db41` (1 source records).
  - [grafana/1t50vxwq2q11](../../data/2026/W41/grafana.json); vendor URL: https://stspg.io/lw22c35nvf3l
- Start time basis: `vendor_reported`.

<a id="ci-20cff0acc195e836836e1e49"></a>

- Canonical ID: `ci-20cff0acc195e836836e1e49` (1 source records).
  - [github/djlmxz2zd0j7](../../data/2026/W41/github.json); vendor URL: https://stspg.io/96smrcth8bpg
- Start time basis: `vendor_reported`.

<a id="ci-710572d3eba4735846b6b393"></a>

- Canonical ID: `ci-710572d3eba4735846b6b393` (1 source records).
  - [anthropic/vmys9qn874h4](../../data/2026/W41/anthropic.json); vendor URL: https://stspg.io/frlrjsqqn605
- Start time basis: `vendor_reported`.

<a id="ci-6a739d8a179dffec5a1944da"></a>

- Canonical ID: `ci-6a739d8a179dffec5a1944da` (1 source records).
  - [cloudflare/76m2gfyfx1xc](../../data/2026/W41/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/76m2gfyfx1xc
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
