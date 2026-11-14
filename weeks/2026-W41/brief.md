# Week 2026-W41 — top 5 of 39 above threshold

2026-10-05 to 2026-10-11 (UTC).

63 source records (68 saved revisions) → 63 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## 1. yandexcloud — Unavailability of Yandex Cloud zones

- URL: https://status.yandex.cloud/en/incidents/2136
- Window: 2026-10-11 01:00 to not yet (34d 11h 36m (ongoing))
- Impact: major, status: investigating
- Components: API Gateway, Application Load Balancer, Audit Trails, Certificate Manager, Cloud Apps, Cloud Backup, Cloud DNS, Cloud Desktop, Cloud Functions, Cloud Interconnect, Cloud Logging, Cloud Notification Service, Cloud Registry, Compute Cloud, Container Registry, Data Processing, Data Streams, Data Transfer, DataSphere, Forms, GOST Gateway, Identity and Access Management, IoT Core, Key Management Service, Managed Service for Apache Spark™, Managed Service for OpenSearch, Managed Service for Prometheus®, Managed Service for Sharded PostgreSQL, Managed Service for YDB, Managed Service for YTsaurus, Managed Service for Apache Kafka®, Managed Service for ClickHouse®, Managed Service for GitLab, Managed Service for Kubernetes®, Managed Service for MySQL®, Managed Service for PostgreSQL, Managed Service for Valkey™, Marketplace, Message Queue, Monitoring, Monium, Network Load Balancer, Object Storage, Resource Manager, Security Deck, Serverless Containers, Smart Web Security, SmartCaptcha, SourceCraft, SourceCraft Code Assistant, SpeechKit, SpeechSense, Stackland, Support Center, Tracker, Translate, VibeCraft, Virtual Private Cloud, Vision OCR, Wiki, Yandex AI Studio, Yandex BareMetal, Yandex Cloud Billing, Yandex Cloud Console, Yandex Cloud Detection and Response, Yandex Cloud Postbox, Yandex Cloud Quota Manager, Yandex Cloud Router, Yandex Cloud Video, Yandex Identity Hub, Yandex Lockbox, Yandex MPP Analytics Engine for PostgreSQL, Yandex Managed Service for Apache Airflow®, Yandex Managed Service for Trino, Yandex MetaData Hub, Yandex Neurosupport, Yandex Query, Yandex SIEM, Yandex Search API, Yandex StoreDoc, Yandex WebSQL, YandexGPT API

<a id="ci-d5585cf81abf4edcfde48cfd"></a>

- Canonical ID: `ci-d5585cf81abf4edcfde48cfd` (1 source records).
  - [yandexcloud/2136](../../data/2026/W41/yandexcloud.json); vendor URL: https://status.yandex.cloud/en/incidents/2136
- Start time basis: `vendor_reported`.

  - `2026-10-11 01:11` investigating — The Yandex Cloud platform is still operating in a limited mode. Our specialists continue restoration work around the clock. The technical support team receives an increased number of requests, so the response time can be increased. Our recommendations continue to be relevant. We recommend using alternative sites for hosting and reserving resources. The next update will be published on October 12 at 19:00 Moscow time.
  - `2026-10-11 03:08` unknown — As a result of the UAV attack, the infrastructure of Yandex’s data center in Vladimir was damaged. The data center’s operation has been completely suspended. There are no casualties; the relevant services are already on site and dealing with the consequences. We consider the remaining resource configuration to be unstable. The platform’s operating mode is emergency — we recommend using any other alternative disaster recovery plans. Communication channels with technical support may be unstable. The Yandex Cloud team is working to restore service operations.
  - `2026-10-11 03:11` unknown — We are sharing the status of individual components and services as of 6:10: • Yandex Cloud Console and API are available; • Compute: partially available; • MDB: most HA database clusters have switched over; some may operate in read‑only mode or be unavailable. • S3: available for reading, potential issues with writing; • Managed Service for Kubernetes: we are assessing the status of regional clusters; • Yandex Identity and Access Management (IAM) is available; • Other services may experience disruptions – we will provide updates on their status later. Communication channels with technical support may be unstable. Information on the status board appears with a delay. We recommend using alternative resource recovery plans on external platforms.
  - `2026-10-11 04:13` unknown — Updates on the operation of services and components as of 07:13: • Yandex Cloud Console and API are available; • Compute: partially available; • MDB: most HA database clusters have switched over; some may operate in read‑only mode or be unavailable. • S3: available for reading, write issues may be observed. • Managed Service for Kubernetes: partially available; • Yandex Identity and Access Management (IAM) is available; • Yandex Monium is unavailable. The team is working on restoring technical support communication channels. We recommend using alternative sites to restore resources.
  - `2026-10-11 05:55` unknown — The team continues to work on restoring and stabilizing key services. We’re sharing an update on the service operations: Yandex Cloud Console and API are available; Compute: partially available; MDB: most HA database clusters have switched over; some may operate in read‑only mode or be unavailable. S3: available for reading, but there may be issues with writing; Managed Service for Kubernetes: available, but issues with individual clusters may persist; Yandex Monium: unavailable; Yandex Container Registry: operational; Yandex AI Studio: services are partially available, more details about the status. Security services are partially available, more details about the status. The Yandex Cloud team is working to resume the operation of technical support communication channels. The main recomme…
  - `2026-10-11 08:14` unknown — The current status of individual platform services: Yandex Cloud Console and API are available; Compute: partially available; Managed Databases: the service is available, but difficulties persist with Managed Service for Apache Kafka®. S3: available for reading and writing, with possible slowdown in operation speed. Managed Service for Kubernetes: available, but issues with individual clusters may persist; Yandex Monium is unavailable; Yandex Container Registry is operational; Yandex AI Studio: services are partially available, more details about the status. Security services are partially available, more details about the status. The operation of the technical support channels has been partially restored; we are working to speed up the processing of inquiries, but there may still be signi…
  - `2026-10-11 10:12` unknown — Service availability remains unchanged. We continue working around the clock to restore services and platform components. Technical support channels have been partially restored. The Yandex Cloud platform team is working to process support requests faster, but responses may still be significantly delayed. As previously announced, quota increases for additional cloud resources remain suspended. We will provide the next update on October 11 at 18:00 Moscow time (UTC+3).
  - `2026-10-11 15:01` unknown — The Yandex Cloud platform continues to operate with limited functionality. Our technical teams are restoring services, while our support team is handling requests around the clock. We continue to recommend considering alternative platforms to deploy your resources more quickly and provide redundancy. We will provide the next update on October 12 at 11:00 Moscow time (UTC+3).
  - `2026-10-12 08:17` unknown — The Yandex Cloud platform is still operating in a limited mode. Our specialists continue restoration work around the clock. The technical support team receives an increased number of requests, so the response time can be increased. Our recommendations continue to be relevant. We recommend using alternative sites for hosting and reserving resources. The next update will be published on October 12 at 19:00 Moscow time.
  - `2026-10-12 16:16` unknown — The Yandex Cloud platform operates in a limited mode. Technical support receives an increased number of requests, so the response time can be increased. We collected answers to the most common questions and recommendations for restoring work, which will help you navigate the available options. The next update will be published on October 13 at 11:00 Moscow time.
  - `2026-10-13 09:05` unknown — The Yandex Cloud platform continues to operate in a limited capacity. Our engineers are continuing recovery efforts. We understand that your teams are currently focused on restoring their services. To ensure you do not have to worry about additional cloud costs during this time, we have decided to make the use of all Yandex Cloud services across all availability zones in the Russian Federation free of charge for existing customers from October 8 through October 31, inclusive. Your October invoice will reflect only the costs for resources used up to and including October 7. Charges for the period starting October 8 may still appear in your billing account, but we will adjust them automatically. No action is required on your part. We will publish the next update on October 13 at 19:00 Moscow…
  - `2026-10-13 16:44` unknown — The status of the cloud platform remains unchanged; the Yandex Cloud team continues to restore services. We have updated the statuses and recommendations for specific services—we do this daily on a dedicated page.
  - `2026-10-14 16:23` unknown — The Yandex Cloud platform is still operating in a limited mode. Our specialists continue the restoration work around the clock. Recommendations for restoring services are updated daily on a separate page: https://yandex.cloud/en/recommendations
  - `2026-10-15 17:04` unknown — Yandex’s cloud platform continues to operate with limitations. Engineers are working around the clock to restore service functionality. We update all currently available product‑related recommendations on a daily basis.
  - `2026-10-16 17:19` unknown — The Yandex Cloud platform operates in a limited mode. Our specialists are carrying out restoration work. Users can keep up-to-date information about the availability of individual services in the management console. We also regularly update recommendations for restoring services.

## 2. yandexcloud — Power failure in the ru-central-b availability zone

- URL: https://status.yandex.cloud/en/incidents/2092
- Window: 2026-10-07 22:31 to 2026-10-10 17:13 UTC (2d 18h 42m)
- Impact: major, status: resolved
- Components: API Gateway, Application Load Balancer, Audit Trails, Certificate Manager, Cloud Apps, Cloud Backup, Cloud CDN, Cloud DNS, Cloud Desktop, Cloud Functions, Cloud Interconnect, Cloud Logging, Cloud Notification Service, Cloud Registry, Compute Cloud, Container Registry, Data Processing, Data Streams, Data Transfer, DataLens, DataSphere, Forms, GOST Gateway, Identity and Access Management, IoT Core, Key Management Service, Managed Service for Apache Spark™, Managed Service for OpenSearch, Managed Service for Prometheus®, Managed Service for Sharded PostgreSQL, Managed Service for YDB, Managed Service for YTsaurus, Managed Service for Apache Kafka®, Managed Service for ClickHouse®, Managed Service for GitLab, Managed Service for Kubernetes®, Managed Service for MySQL®, Managed Service for PostgreSQL, Managed Service for Valkey™, Marketplace, Message Queue, Monitoring, Monium, Network Load Balancer, Object Storage, Resource Manager, Security Deck, Serverless Containers, Smart Web Security, SmartCaptcha, SourceCraft, SourceCraft Code Assistant, SpeechKit, SpeechSense, Stackland, Support Center, Tracker, Translate, VibeCraft, Virtual Private Cloud, Vision OCR, Wiki, Yandex AI Studio, Yandex BareMetal, Yandex Cloud Billing, Yandex Cloud Console, Yandex Cloud Detection and Response, Yandex Cloud Postbox, Yandex Cloud Quota Manager, Yandex Cloud Router, Yandex Cloud Video, Yandex Identity Hub, Yandex Lockbox, Yandex MPP Analytics Engine for PostgreSQL, Yandex Managed Service for Apache Airflow®, Yandex Managed Service for Trino, Yandex MetaData Hub, Yandex Neurosupport, Yandex Query, Yandex SIEM, Yandex Search API, Yandex StoreDoc, Yandex WebSQL, YandexGPT API

<a id="ci-8e0bde59a25eb3bfe75e60a3"></a>

- Canonical ID: `ci-8e0bde59a25eb3bfe75e60a3` (1 source records).
  - [yandexcloud/2092](../../data/2026/W41/yandexcloud.json); vendor URL: https://status.yandex.cloud/en/incidents/2092
- Start time basis: `vendor_reported`.

  - `2026-10-07 22:31` investigating — Power outages are observed in the ru-central1-b availability zone. The team is working to fix the problem. If possible, we recommend shifting the load to other availability zones. We will keep you informed of the progress of the work.
  - `2026-10-07 23:19` unknown — We have taken measures to localize the problem in the ru-central1-b availability zone. The situation has not stabilized yet — the team continues to find out the reasons and work on recovery. We will share updated information in the near future.
  - `2026-10-08 01:16` unknown — The Yandex Cloud team continues to work on eliminating the consequences of the incident. The recovery dates are being clarified — we will inform them as soon as confirmed information becomes available.
  - `2026-10-08 01:55` unknown — Due to the incident in the ru-central1-b availability zone, problems may arise in other zones when creating new VMs, k8s clusters, and databases. We are continuing to address the consequences of the incident and will report on the progress of the recovery.
  - `2026-10-08 05:26` unknown — The unavailability of the ru-central1-b zone remains for now, we have completed all the necessary work to switch the load to other availability zones. All services in the remaining zones are operating normally. Currently, the creation of new resources in the cloud infrastructure may be limited. If you have any questions, please contact technical support. Our specialists continue to work on eliminating the consequences of the incident. We will be back with an update before 15:00 Moscow time.
  - `2026-10-08 12:19` unknown — The ru-central1-b availability zone remains unavailable. The cause of the incident was a fire that occurred as a result of a UAV attack on the Yandex data center in Sasovo and affected part of the Yandex Cloud infrastructure. The data center has been completely shut down. Currently, the creation of new resources in the cloud infrastructure may be limited. Our team continues to work on eliminating the consequences of the incident. If you have any questions, please contact technical support. The next update will be published at 19:00 Moscow time.
  - `2026-10-08 15:25` unknown — The situation with the availability of the ru-central1-b zone has not changed yet. Our specialists are continuing the restoration work, however, according to current estimates, it will not be possible to resume the work of the zone in the near future. The recommendation remains the same, if possible, to use alternative plans to restore the services. The next update will be given at 23:00 Moscow time.
  - `2026-10-08 20:05` unknown — The availability of the ru-central1-b zone has not yet been restored. Our team continues to eliminate the consequences of the incident. If possible, we recommend using alternative service recovery options, including dedicated physical servers Bare Metal. The next update will be published on October 9 at 12:00 Moscow time.
  - `2026-10-09 07:38` unknown — The ru-central1-b zone is still unavailable. The other availability zones are working. Our specialists are working around the clock to restore the previous level of infrastructure redundancy for key services in the near future and ensure the availability of additional capacities and backup components. The technical support team helps customers to resume their services. We have also agreed with other cloud providers to speed up the placement and backup of your resources on their sites. If you see a need, you can contact your colleagues directly: Selectel K2Cloud VK Cloud Using the code word "CloudBro”, our colleagues will be ready to process your application as a matter of priority and provide the necessary technical assistance together with our teams. If you need additional advice, please…
  - `2026-10-09 12:42` unknown — The ru-central1-b zone is unavailable. Our specialists continue to work to eliminate the consequences of the incident and restore cloud services. Current status for some services: Managed Service for Kubernetes: it is now possible to transfer the management components of regional clusters from the ru-central1-b zone to other availability zones. To restore the applications, we recommend creating node groups in other zones and transferring the load to them. Managed Databases: in fault-tolerant clusters, it is now possible to remove hosts from the ru-central1-b zone. After deleting them, cluster management operations will become available, including changing settings and adding users. Yandex AI Studio: since yesterday, all models and components of the platform have been operating normally. At…
  - `2026-10-09 18:47` unknown — The situation with the availability of the ru-central1-b zone remains unchanged so far. Our engineers continue to work around the clock to eliminate the consequences of the incident and restore services. The next update will be given on October 10 at 12:00 Moscow time.
  - `2026-10-10 06:00` unknown — At the moment, the ru-central1-b zone remains unavailable. Our engineers continue to work to eliminate the consequences of the incident. To ensure stable operation of the cloud platform for the coming weekend, we are switching to fully manual management of quotas for computing resources and disks. All users have access to a small additional amount of resources beyond those already used. This will allow some clients, if necessary, to restore resources from backups, including for their subsequent transfer to other sites, or perform other operations. The creation of new resources beyond the limits set by us will be limited. Priority in expansion will be given to those whose critical services are still not restored. If you have any questions, please contact technical support or your account ma…
  - `2026-10-10 17:13` unknown — The ru-central1-b zone remains unavailable. In the morning, we switched to fully manual management of quotas for computing resources and disks. Technical support specialists sort out each situation manually, so responses may be delayed. We will try to sort out all quota requests as soon as possible. Our team continues to work around the clock to increase the amount of available computing resources. The next update will be given on October 11 at 12:00 Moscow time.

## 3. cloudflare — Missing Logs and Traces in Cloudflare Observability Platform

- URL: https://www.cloudflarestatus.com/incidents/pfvfug94jq4q
- Window: 2026-10-05 13:00 to 2026-10-23 21:00 UTC (18d 8h 0m)
- Impact: none, status: resolved
- Components: —

<a id="ci-dcb212550be05a0c4eb3a32a"></a>

- Canonical ID: `ci-dcb212550be05a0c4eb3a32a` (1 source records).
  - [cloudflare/pfvfug94jq4q](../../data/2026/W41/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/pfvfug94jq4q
- Start time basis: `vendor_reported`.

  - `2026-11-03 18:40` resolved — Between October 5 at 1:00 PM UTC and October 23 at 9:00 PM UTC, some customers may have experienced missing log events, incomplete traces, or inconsistent issue counts due to ingestion pipeline capacity limits. Traffic and Workers were not affected, and customers will not be billed for dropped events. Capacity has since been increased, and service has returned to normal.

## 4. grafana — IRM Unable to Send SMS to Thailand Numbers

- URL: https://stspg.io/9l8htmqt82s2
- Window: 2026-10-05 16:17 to 2026-10-19 11:40 UTC (13d 19h 22m)
- Impact: minor, status: resolved
- Components: Incident Management and Response (IRM)

<a id="ci-aa6f59211a96cfbcca9feb49"></a>

- Canonical ID: `ci-aa6f59211a96cfbcca9feb49` (1 source records).
  - [grafana/6x77ccl9shpy](../../data/2026/W41/grafana.json); vendor URL: https://stspg.io/9l8htmqt82s2
- Start time basis: `vendor_reported`.

  - `2026-10-05 16:17` investigating — We are currently investigating an issue impacting all SMS attempts to Austria for IRM. All SMS attempts to Austria for IRM are currently blocked. We are working on a resolution, and an update will follow as soon as we have more info. Users should configure mobile push or phone calls as notifications.
  - `2026-10-05 21:47` identified — The issue has been identified. We are working on a resolution, but do not have an ETA at this time. We will provide another update as soon as we have more information to share. Users should configure mobile push or phone calls as notifications.
  - `2026-10-06 13:42` identified — This issue is also impacting numbers from Thailand and Bangladesh. The title of this status page has been updated to reflect this. Users should configure mobile push or phone calls as notifications. We are working on a resolution, and we will provide another update as soon as we have more information to share.
  - `2026-10-09 12:51` identified — We are continuing working on a resolution, and will provide another update as soon as we have more information to share. In the meantime we recommend, affected users to configure mobile push or phone calls as notifications.
  - `2026-10-13 21:35` identified — We are continuing working on a resolution, and will provide another update as soon as we have more information to share. In the meantime we recommend, affected users to configure mobile push or phone calls as notifications.
  - `2026-10-15 09:33` identified — SMS delivery to Austrian numbers has been restored. SMS delivery to Thailand and Bangladesh is still affected, and we are continuing to work on a resolution. We will provide another update as soon as we have more information to share. In the meantime, we recommend that affected users configure mobile push or phone calls as notifications.
  - `2026-10-16 15:20` identified — SMS delivery to Bangladesh has been restored. SMS delivery to Thailand is still affected, and we are continuing to work on a resolution. We will provide another update as soon as we have more information to share. In the meantime, we recommend that affected users configure mobile push or phone calls as notifications.
  - `2026-10-19 11:40` resolved — Following further investigation, we've confirmed that SMS delivery to certain destinations is not currently supported. We've updated our documentation to clarify the supported destinations and available notification options. We recommend using mobile push notifications or phone calls as alternatives where SMS delivery is unavailable.

## 5. cloudflare — Elevated Errors in China Network Colos

- URL: https://www.cloudflarestatus.com/incidents/cyxclg7nnhjy
- Window: 2026-10-10 04:13 to 2026-10-15 09:28 UTC (5d 5h 15m)
- Impact: minor, status: resolved
- Components: China Network

<a id="ci-1d7fe76c931dba8c5cb73366"></a>

- Canonical ID: `ci-1d7fe76c931dba8c5cb73366` (1 source records).
  - [cloudflare/cyxclg7nnhjy](../../data/2026/W41/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/cyxclg7nnhjy
- Start time basis: `vendor_reported`.

  - `2026-10-10 04:13` investigating — We are investigating a spike in errors observed in our China colos.
  - `2026-10-10 05:48` identified — The issue has been identified and a fix is being implemented.
  - `2026-10-12 04:23` identified — We are investigating reports of degraded performance for traffic traversing the public internet between certain cities and provinces in mainland China and origin servers located outside the region. The impact is being reported across multiple ISPs and manifests as slower speeds or intermittent timeouts. While this congestion occurs upstream of Cloudflare's network, our engineering team is actively working to mitigate the effects. Intermittent 524 timeout errors is expected to continue due to ongoing congestion. While there is no fixed timeline for full resolution, as there are dependencies outside our network, we are taking steps to reduce customer impact.
  - `2026-10-12 15:33` identified — We have identified that increased latency and 5xx errors affecting certain China Network colos are caused by cross-border routing instability in the China network. We are actively working to mitigate impact for affected customers by adjusting traffic routing, and are continuing to work with our network partners on a resolution. Customers experiencing issues may contact Cloudflare Support.
  - `2026-10-14 17:50` monitoring — A fix has been implemented and we are monitoring the results.
  - `2026-10-15 09:28` resolved — This incident has been resolved.

## Also above threshold (34)

- `anthropic` Elevated errors on platform.claude.com (6d 4h 35m) — https://stspg.io/sjkw7njwf0mt
- `openai` Delayed Costs data in the Compliance API (2d 22h 39m) — https://status.openai.com/incidents/01M4GV1M2A3K38T73JCVCD9G0R
- `slack` Slack for iOS closing unexpectedly when opening channels (3d 1h 28m) — https://slack-status.com/2026-10/a0cfb0208a1d2032
- `gcp` Google Cloud Storage (GCS) customers experienced elevated latencies and error rates when performing upload or write operations. (1d 5h 46m) — https://status.cloud.google.com/incidents/mSFVt6PiP557bjRyYsut
- `grafana` Degraded Performance with Increased Latency for Grafana Cloud Metrics and Traces in GCP US Central (11h 33m) — https://stspg.io/cjbf6knm4jq4
- `cloudflare` Durable Objects Availability Issue in Eastern North America (5h 21m) — https://www.cloudflarestatus.com/incidents/z28czd65h29l
- `cloudflare` Network Performance Issues in Columbus (9h 2m) — https://www.cloudflarestatus.com/incidents/tt8f4ttcuonf
- `circleci` Delay in data appearing in UI and API (14h 35m) — https://stspg.io/hkfkwz292bz8
- `openai` Elevated Work Mode errors (19h 6m) — https://status.openai.com/incidents/01M45E0JFFJZ6EQ3BWYPW74R2Z
- `cloudflare` Cloudflare DNS Resolver 1.1.1.1 and Gateway DNS are having issues in SIN (4h 49m) — https://www.cloudflarestatus.com/incidents/4knp3pzs3qvx
- `yandexcloud` Increasing timings in the Object Storage service (14h 20m) — https://status.yandex.cloud/en/incidents/2130
- `cloudflare` API Shield JWT Validation Errors (11h 43m) — https://www.cloudflarestatus.com/incidents/ts295zqycqn7
- `cloudflare` Cloudflare API service issues (7h 19m) — https://www.cloudflarestatus.com/incidents/x8n12150xwf6
- `github` Several services are degraded (3h 43m) — https://stspg.io/w3rcssslq22w
- `github` Incident with Actions (3h 37m) — https://stspg.io/c11dc9nb1zdq
- `circleci` Delays in pipelines, workflows, UI data, and notifications (5h 18m) — https://stspg.io/s7tpg47srxdd
- `openai` Elevated errors in ChatGPT conversations (7h 2m) — https://status.openai.com/incidents/01M46G248GGHV9RSAJE82BZY1X
- `openai` Unable to create new ChatGPT Work threads in latest version of desktop app (6h 21m) — https://status.openai.com/incidents/01M4ADHMMBK7K92NJFJ0DXRYJM
- `openai` Elevated error rates in Codex Cloud (6h 2m) — https://status.openai.com/incidents/01M46KED6KQ6PR195KZYYVNY8W
- `openai` Admin Console Intermittently Inaccessible (5h 17m) — https://status.openai.com/incidents/01M46C47WYYD9NMZVGH83X3ZR1
- `stripe` Elevated Swish errors (6h 33m) — https://stspg.io/m7j171ppknbs
- `cloudflare` Network Performance Issues in Kochi (3h 0m) — https://www.cloudflarestatus.com/incidents/xczvsioz1un8
- `openai` Some ChatGPT content may be unavailable on Android (3h 26m) — https://status.openai.com/incidents/01M4H7FKXDA425TDNP2CRGEYA3
- `github` Incident with Git Operations, Issues, Actions and Pull Requests (47m) — https://stspg.io/4p8x0jnnsx7m
- `yandexcloud` The problem is with the Cloud Functions and Serverless Containers services (1h 25m) — https://status.yandex.cloud/en/incidents/2124
- `openai` Computer Use Is Experiencing Issues (2h 45m) — https://status.openai.com/incidents/01M493CZ0CB14KPC0EER83MNVB
- `circleci` Elevated wait times for machine jobs (3h 35m) — https://stspg.io/vcsrgfwdsy4z
- `openai` Errors processing large PDFs in the Responses API (2h 4m) — https://status.openai.com/incidents/01M48Y8E81QAFF9XKVGE118E6X
- `github` Disruption with some GitHub services (1h 44m) — https://stspg.io/pjcfc4c3pm7g
- `openai` Workspace admins may be unable to manage plugins (1h 38m) — https://status.openai.com/incidents/01M4HA7BCCXVHTR5DCYMSBQF2H
- `anthropic` Elevated errors loading usage data on platform.claude.com (1h 13m) — https://stspg.io/scdhhnplc2jn
- `anthropic` Elevated errors for Claude Opus 5.5 (18m) — https://stspg.io/bzc81gbv1b99
- `anthropic` Elevated errors for Claude Mythos 5.1 and Claude Fable 5.1 (15m) — https://stspg.io/wdqhxt8xt574
- `github` Incident with Git Operations, Pull Requests and Actions (0m) — https://stspg.io/96smrcth8bpg
