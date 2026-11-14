# Week 2026-W43

2026-10-19 to 2026-10-25 (UTC). Generated 2026-11-14 12:36.

63 canonical incidents, 41 above threshold. 2 of 14 sources unavailable: datadog, yandexcloud.

63 source records (73 saved revisions) → 63 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | gcp | [Multiple services were experiencing issues in the europe-west2 region.](https://status.cloud.google.com/incidents/Xs2XTHzAVC9YnafnPyzj) | 2026-10-20 06:00 | 6h 8m | major | AlloyDB for PostgreSQL, Apigee, BigQuery Data Transfer Service, Cloud Data Fusion, Cloud Filestore, Cloud Firestore, Cloud Load Balancing, Cloud Memorystore, Cloud NAT, Cloud Run, Cloud Spanner, Cloud Workstations, Datastream, Google App Engine, Google BigQuery, Google Cloud Bigtable, Google Cloud Composer, Google Cloud Dataflow, Google Cloud Dataproc, Google Cloud NetApp Volumes, Google Cloud Pub/Sub, Google Cloud SQL, Google Cloud Storage, Google Compute Engine, Google Kubernetes Engine, Looker (Google Cloud core), Managed Service for Apache Kafka, Memorystore for Memcached, Memorystore for Redis, Memorystore for Redis Cluster, Memorystore for Valkey, Network Connectivity Center, Persistent Disk, Private Service Connect, Speech-to-Text, Storage Transfer Service, VMWare engine, Virtual Private Cloud (VPC) | 27968 = 368 × 38 × 2.0 |
| 2 | grafana | [Some Grafana Instances Experiencing Issues Coming Online after Restart](https://stspg.io/v6j9g9gqpwsw) | 2026-10-23 13:33 | 4h 42m | minor | AWS Australia - prod-au-southeast-1, AWS Australia - prod-au-southeast-2, AWS Brazil - prod-sa-east-1, AWS Canada - prod-ca-east-0, AWS Canada prod-ca-central-1, AWS East (VA) prod-us-east-1, AWS Germany - prod-eu-west-2, AWS Germany - prod-eu-west-4, AWS Germany prod-eu-central-1, AWS India - prod-ap-south-1, AWS Indonesia - prod-ap-southeast-2, AWS Indonesia - prod-ap-southeast-3, AWS Ireland - prod-eu-west-6, AWS Ireland eu-west-1, AWS Japan - prod-ap-northeast-0, AWS Singapore - prod-ap-southeast-1, AWS Sweden - prod-eu-north-0, AWS Sweden eu-north-1, AWS Switzerland - prod-eu-central-0, AWS Switzerland eu-central-2, AWS UAE - prod-me-central-1, AWS UK - prod-gb-south-1, AWS US East (OH) - prod-us-east-0, AWS US East (OH) - prod-us-east-2, AWS US East (VA) - prod-us-east-3, AWS US West - prod-us-west-0, AWS West prod-us-west-2, Azure Netherlands - prod-eu-west-3, Azure UAE - prod-me-central-2, Azure US Central - us-central2, Federal Cloud - AWS US Gov West, GCP Australia - prod-au-southeast-0, GCP Belgium - prod-eu-west-0, GCP Brazil - prod-sa-east-0, GCP India - prod-ap-south-0, GCP Saudi Arabia - prod-me-central-0, GCP Singapore - prod-ap-southeast-0, GCP UK - prod-gb-south-0, GCP US Central - prod-us-central-0, GCP US Central - prod-us-central-3, GCP US Central - prod-us-central-4, GCP US Central - prod-us-central-7, GCP US East - prod-us-east-1, play.grafana.org | 12408 = 282 × 44 × 1.0 |
| 3 | grafana | [Grafana Cloud unavailibity in the prod-gb-south-0 region](https://stspg.io/blcsc38zxjxw) | 2026-10-20 12:43 | 3h 49m | minor | Cost Management and Billing, Fleet Management, GCP UK - prod-gb-south-0, GCP UK - prod-gb-south-0, GCP UK - prod-gb-south-0, GCP UK - prod-gb-south-0, GCP UK - prod-gb-south-0, GCP UK - prod-gb-south-0, GCP UK - prod-gb-south-0, GCP UK - prod-gb-south-0, GCP UK - prod-gb-south-0, GCP UK - prod-gb-south-0: API, GCP UK - prod-gb-south-0: Alertmanager, GCP UK - prod-gb-south-0: Alertmanager and Rules Configuration API, GCP UK - prod-gb-south-0: Ingestion, GCP UK - prod-gb-south-0: Ingestion, GCP UK - prod-gb-south-0: Metrics Generator, GCP UK - prod-gb-south-0: Public Probes, GCP UK - prod-gb-south-0: Querying, GCP UK - prod-gb-south-0: Querying, GCP UK - prod-gb-south-0: Rule Evaluation, Grafana.com, Incident Management and Response (IRM), Integrations, Public Snapshots, Support Tickets | 5954 = 229 × 26 × 1.0 |
| 4 | cloudflare | [Cloudflare Secure Web Gateway and DLP enforcement issues for Google Gemini](https://www.cloudflarestatus.com/incidents/3z8y23zn7tr2) | 2026-10-19 20:30 | 1d 2h 16m | minor | Data Loss Prevention (DLP), Gateway | 5674 = 1576 × 2 × 1.8 |
| 5 | github | [Project updates delayed](https://stspg.io/sx51v1tlpqjc) | 2026-10-22 15:28 | 11h 2m | minor | API Requests, Actions, Issues, Packages, Pull Requests | 4965 = 662 × 5 × 1.5 |
| 6 | cloudflare | [Network Performance Issues in Ashburn](https://www.cloudflarestatus.com/incidents/39fz3wkl42tr) | 2026-10-24 01:01 | 2h 8m | minor | Artifacts, Ashburn, United States - (IAD), D1, Durable Objects, Network, Queues, R2, Workflows | 1843 = 128 × 8 × 1.8 |
| 7 | cloudflare | [Delayed audit log delivery for WARP and Zero Trust services](https://www.cloudflarestatus.com/incidents/qwrtvb4jfvk7) | 2026-10-22 17:32 | 6h 38m | major | Audit Logs, Zero Trust | 1433 = 398 × 2 × 1.8 |
| 8 | github | [Delays for Pull Requests and Webhook deliveries](https://stspg.io/gnnt0tx9nxmk) | 2026-10-21 14:52 | 4h 30m | critical | Actions, Pull Requests, Webhooks | 1215 = 270 × 3 × 1.5 |
| 9 | grafana | [IRM outage in the region prod-gb-south-0 and prod-eu-west-0](https://stspg.io/nh9hp52f8qzh) | 2026-10-20 07:28 | 3h 5m | none | GCP Belgium - prod-eu-west-0: Alertmanager, GCP Belgium - prod-eu-west-0: Alertmanager and Rules Configuration API, GCP Belgium - prod-eu-west-0: Rule Evaluation, GCP UK - prod-gb-south-0: Alertmanager, GCP UK - prod-gb-south-0: Alertmanager and Rules Configuration API, GCP UK - prod-gb-south-0: Rule Evaluation | 1110 = 185 × 6 × 1.0 |
| 10 | anthropic | [Elevated errors affecting Connectors](https://stspg.io/4dcr8c005yxc) | 2026-10-22 03:50 | 4h 16m | minor | Claude Cowork, claude.ai | 768 = 256 × 2 × 1.5 |
| 11 | cloudflare | [Support Helpdesk Availability Issues](https://www.cloudflarestatus.com/incidents/8m9a1fjjwxp5) | 2026-10-23 03:03 | 6h 39m | minor | Support Site | 718 = 399 × 1 × 1.8 |
| 12 | cloudflare | [Cloudflare Zero Trust connectivity](https://www.cloudflarestatus.com/incidents/0eqtn0khkz87) | 2026-10-20 03:20 | 3h 11m | critical | Dashboard, Zero Trust | 688 = 191 × 2 × 1.8 |
| 13 | stripe | [Elevated Revolut Pay errors](https://stspg.io/80s4dn3gj9ql) | 2026-10-20 07:09 | 9h 32m | none | Acquirers and payment methods | 686 = 572 × 1 × 1.2 |
| 14 | cloudflare | [Elevated error rates for Durable Objects in Western Europe](https://www.cloudflarestatus.com/incidents/v967mz6rrcxw) | 2026-10-22 17:51 | 6h 17m | minor | Durable Objects | 679 = 377 × 1 × 1.8 |
| 15 | cloudflare | [Network Error Log Processing Unavailable](https://www.cloudflarestatus.com/incidents/l3d16334tgjp) | 2026-10-22 16:30 | 5h 30m | major | Logs | 594 = 330 × 1 × 1.8 |
| 16 | anthropic | [Elevated errors for Claude Sonnet 5.5 and other models](https://stspg.io/s3wd6gmdsdhn) | 2026-10-25 06:14 | 1h 39m | major | Claude API (api.anthropic.com), Claude Code, Claude Cowork, claude.ai | 594 = 99 × 4 × 1.5 |
| 17 | openai | [Elevated ChatGPT conversation and Responses API errors](https://status.openai.com/incidents/01M5GPE9C4E19HE461P8JAMYXD) | 2026-10-22 02:18 | 6h 32m | minor | — | 588 = 392 × 1 × 1.5 |
| 18 | cloudflare | [Delayed, missed, and duplicated Workers Cron Triggers](https://www.cloudflarestatus.com/incidents/3ts1mfzs0khz) | 2026-10-23 15:27 | 5h 9m | minor | Workers | 556 = 309 × 1 × 1.8 |
| 19 | openai | [Elevated FedRAMP API error rates for Astra and Sol](https://status.openai.com/incidents/01M5FD63HHDRPQPM5YMQWJ5HTN) | 2026-10-21 14:17 | 5h 17m | minor | — | 476 = 317 × 1 × 1.5 |
| 20 | grafana | [DB issues in prod-sa-east-1](https://stspg.io/28py1pvhnpjf) | 2026-10-20 21:36 | 2h 33m | none | AWS Brazil - prod-sa-east-1, AWS Brazil - prod-sa-east-1: API, AWS Brazil - prod-sa-east-1: Public Probes | 459 = 153 × 3 × 1.0 |
| 21 | cloudflare | [Elevated Errors for R2 and Durable Objects in ENAM](https://www.cloudflarestatus.com/incidents/p1vurmibvmag) | 2026-10-21 19:15 | 2h 6m | major | Durable Objects, R2 | 454 = 126 × 2 × 1.8 |
| 22 | openai | [New Ad Accounts failing to provision admins](https://status.openai.com/incidents/01M5JHVJVC8WAXB4WRYFMKG6HP) | 2026-10-22 19:36 | 4h 31m | minor | — | 406 = 271 × 1 × 1.5 |
| 23 | cloudflare | [Increased Latency in India](https://www.cloudflarestatus.com/incidents/t77gwnpr6oe5) | 2026-10-21 16:42 | 3h 14m | minor | CDN/Cache | 349 = 194 × 1 × 1.8 |
| 24 | github | [Disruption with Dependabot Alerts](https://stspg.io/zdk3zqt69fnz) | 2026-10-23 14:58 | 3h 52m | minor | — | 348 = 232 × 1 × 1.5 |
| 25 | cloudflare | [Network connectivity issues affecting services in North America](https://www.cloudflarestatus.com/incidents/sufgcbe8rrwt) | 2026-10-23 15:47 | 2h 56m | major | Network | 317 = 176 × 1 × 1.8 |
| 26 | github | [Degraded availability for GitHub Actions, GitHub Pages and Webhooks](https://stspg.io/nrv2k8s68bwn) | 2026-10-19 15:49 | 1h 7m | critical | Actions, Pages, Webhooks | 302 = 67 × 3 × 1.5 |
| 27 | stripe | [Elevated Bancontact errors](https://stspg.io/k2jfv0fg9ff6) | 2026-10-22 22:20 | 3h 35m | none | Acquirers and payment methods | 258 = 215 × 1 × 1.2 |
| 28 | anthropic | [Elevated errors for Claude Fable 5 and Claude Fable 5.1](https://stspg.io/7n1wgb6qsfly) | 2026-10-24 05:16 | 36m | major | Claude API (api.anthropic.com), Claude Code, Claude Cowork, claude.ai | 216 = 36 × 4 × 1.5 |
| 29 | openai | [Delays in customer support responses](https://status.openai.com/incidents/01M5KDJ4079XQTZDSRX40NMN6H) | 2026-10-23 03:41 | 2h 17m | none | — | 206 = 137 × 1 × 1.5 |
| 30 | anthropic | [Elevated errors for Claude Fable 5.1, Sonnet 5.5, and Haiku 5.5](https://stspg.io/gqp288t3qzpt) | 2026-10-23 23:37 | 33m | critical | Claude API (api.anthropic.com), Claude Code, Claude Cowork, claude.ai | 198 = 33 × 4 × 1.5 |
| 31 | anthropic | [Elevated errors for Claude Fable 5.1](https://stspg.io/y0gkl12ztjst) | 2026-10-25 08:17 | 32m | major | Claude API (api.anthropic.com), Claude Code, Claude Cowork, claude.ai | 192 = 32 × 4 × 1.5 |
| 32 | openai | [Degraded Performance of Dots on Enterprise](https://status.openai.com/incidents/01M5JF8DR3EGQJNTXPZ24AKX2G) | 2026-10-22 18:51 | 2h 4m | minor | — | 186 = 124 × 1 × 1.5 |
| 33 | stripe | [Elevated Bancontact errors](https://stspg.io/ncmhd5df6f3k) | 2026-10-22 14:18 | 2h 25m | none | Acquirers and payment methods | 174 = 145 × 1 × 1.2 |
| 34 | grafana | [Metrics: Evelvated Read and Ruler Errors.](https://stspg.io/t7cdk3q6bcz9) | 2026-10-19 15:06 | 2h 49m | major | GCP US Central - prod-us-central-0: Querying | 169 = 169 × 1 × 1.0 |
| 35 | stripe | [Spike in cancellations and expired errors with PayByBank payments](https://stspg.io/gmpxzbw40g3d) | 2026-10-20 10:46 | 2h 17m | none | Acquirers and payment methods | 164 = 137 × 1 × 1.2 |
| 36 | circleci | [Increased wait times for macOS jobs on m4pro.large](https://stspg.io/pp6wqy3vlms3) | 2026-10-20 15:58 | 2h 15m | minor | macOS Jobs | 135 = 135 × 1 × 1.0 |
| 37 | grafana | [Some GCK6 App Project Pages are Failing to Load](https://stspg.io/wh6fr7yf9j5h) | 2026-10-21 13:22 | 2h 7m | major | GCk6 App | 127 = 127 × 1 × 1.0 |
| 38 | grafana | [Some GCK6 App Project Pages are Failing to Load](https://stspg.io/y9w2szsh123d) | 2026-10-20 14:36 | 1h 31m | major | GCk6 App | 91 = 91 × 1 × 1.0 |
| 39 | grafana | [Elevated Trace Ingestion Latency and Errors for Grafana Cloud Traces in AWS Germany](https://stspg.io/p5hhk4h6k6dm) | 2026-10-23 09:31 | 44m | major | AWS Germany - prod-eu-west-2, AWS Germany - prod-eu-west-2: Ingestion | 88 = 44 × 2 × 1.0 |
| 40 | anthropic | [Elevated errors for Claude Opus 5.5](https://stspg.io/dmn3bc440bnb) | 2026-10-24 00:31 | 5m | major | Claude API (api.anthropic.com), Claude Code, Claude Cowork, claude.ai | 30 = 5 × 4 × 1.5 |
| 41 | openai | [Authentication errors for some users](https://status.openai.com/incidents/01M5G9QQ9V4CYDWFRYF9BWZ5GG) | 2026-10-21 22:36 | unknown | major | — | 0 = 0 × 1 × 1.5 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. gcp — Multiple services were experiencing issues in the europe-west2 region.

- URL: https://status.cloud.google.com/incidents/Xs2XTHzAVC9YnafnPyzj
- Started: 2026-10-20 06:00 UTC
- Resolved: 2026-10-20 12:08 UTC
- Duration: 6h 8m
- Impact: major, status: resolved
- Components: AlloyDB for PostgreSQL, Apigee, BigQuery Data Transfer Service, Cloud Data Fusion, Cloud Filestore, Cloud Firestore, Cloud Load Balancing, Cloud Memorystore, Cloud NAT, Cloud Run, Cloud Spanner, Cloud Workstations, Datastream, Google App Engine, Google BigQuery, Google Cloud Bigtable, Google Cloud Composer, Google Cloud Dataflow, Google Cloud Dataproc, Google Cloud NetApp Volumes, Google Cloud Pub/Sub, Google Cloud SQL, Google Cloud Storage, Google Compute Engine, Google Kubernetes Engine, Looker (Google Cloud core), Managed Service for Apache Kafka, Memorystore for Memcached, Memorystore for Redis, Memorystore for Redis Cluster, Memorystore for Valkey, Network Connectivity Center, Persistent Disk, Private Service Connect, Speech-to-Text, Storage Transfer Service, VMWare engine, Virtual Private Cloud (VPC)
- Canonical incident: [ci-b2dbd754a165046e8a4caebe](#ci-b2dbd754a165046e8a4caebe)

- `2026-10-20 08:35` **investigating** — **Summary** Multiple services are experiencing issues in the europe-west2 region. **Description** We are experiencing an issue with multiple services in europe-west2 beginning on Monday, 2026-10-19 23:08 PDT. Our engineering team continues to investigate the issue. We will provide an update by Tuesday, 2026-10-20 02:30 PDT with current details. We apologize to all who are affected by the disruption. **Symptoms** Multiple Google Cloud services are experiencing elevated errors and intermittent network failures in europe-west2 region. **Workaround** None for now
- `2026-10-20 09:24` **investigating** — **Summary** Multiple services are experiencing issues in the europe-west2 region. **Description** We are experiencing an issue with multiple services in europe-west2 beginning on Monday, 2026-10-19 23:08 PDT. Our engineering team continues to investigate the issue. We will provide an update by Tuesday, 2026-10-20 04:00 PDT with current details. We apologize to all who are affected by the disruption. **Symptoms** Multiple Google Cloud services are experiencing elevated errors and intermittent network failures in europe-west2 region. Note: Looker Studio is not impacted by this outage. **Workaround** If feasible, the impacted customer may try to migrate their workload impacted to another available location (zone or region).
- `2026-10-20 11:00` **investigating** — **Summary** Multiple services are experiencing issues in the europe-west2 region. **Description** We are experiencing an issue with multiple services in europe-west2 beginning on Monday, 2026-10-19 23:08 PDT. Our engineering team continues to investigate the issue. We will provide an update by Tuesday, 2026-10-20 05:00 PDT with current details. We apologize to all who are affected by the disruption. **Symptoms** Multiple Google Cloud services are experiencing elevated errors and intermittent network failures in europe-west2 region. Note: Looker Studio is not impacted by this outage. **Workaround** If feasible, the impacted customer may try to migrate their workload impacted to another available location (zone or region).
- `2026-10-20 11:12` **investigating** — **Summary** Multiple services are experiencing issues in the europe-west2 region. **Description** We are experiencing an issue with multiple services in europe-west2 beginning on Monday, 2026-10-19 23:08 PDT. Due to a power event we had suffered an outage within europe-west2-b zone. The power event is mitigated. Partial mitigation to infrastructure services will be observed within the next 30 minutes and the remaining services will be mitigated there after. We will provide an update by Tuesday, 2026-10-20 05:00 PDT with current details. We apologize to all who are affected by the disruption. **Symptoms** Google Cloud services are unavailable in europe-west2-b zone; some services in europe-west2 region may experience elevated errors and intermittent network issues. **Workaround** If feasible, the impacted customer may try to migrate their workload impacted to another available location (zone or region).
- `2026-10-20 11:58` **investigating** — **Summary** Multiple services are experiencing issues in the europe-west2 region. **Description** We are experiencing an issue with multiple services in europe-west2 beginning on Monday, 2026-10-19 23:08 PDT. A power event caused an outage within europe-west2-b zone. The power event is mitigated. We are seeing the recovery of services and are prioritizing the recovery of existing workloads.Customers may not be able to yet create new workloads. Full recovery of existing workloads should be completed, tentatively by Tuesday, 2026-10-20 06:00. We will provide an update by Tuesday, 2026-10-20 06:00 PDT with current details. We apologize to all who are affected by the disruption. **Symptoms** Google Cloud services are unavailable in europe-west2-b zone; some services in europe-west2 region may experience elevated errors and intermittent network issues. **Workaround** If feasible, the impacted customer may try to migrate their workload impacted to another available location (zone or region).
- `2026-10-20 12:44` **investigating** — **Summary** Multiple services are experiencing issues in the europe-west2 region. **Description** We are seeing the recovery of services for most workloads in europe-west2 as of Tuesday, 2026-10-20 05:10 PDT. We will provide an update by Tuesday, 2026-10-20 07:00 PDT with current details. We apologize to all who are affected by the disruption. **Symptoms** Google Cloud services are unavailable in europe-west2-b zone; some services in europe-west2 region may experience elevated errors and intermittent network issues. **Workaround** If feasible, the impacted customer may try to migrate their workload impacted to another available location (zone or region).
- `2026-10-20 14:03` **investigating** — **Summary** Multiple services were experiencing issues in the europe-west2 region. **Description** The issue with impacted GCP products has been mitigated for most of the customers and it is believed to be affecting currently a very small number of customers and our Engineering Team is working on it. If you have questions or are impacted, please open a case with the Support Team and we will work with you until this issue is resolved. No further updates will be provided here. We will publish an analysis of this incident once we have completed our internal investigation. We thank you for your patience while we're working on resolving the issue. **Symptoms** Google Cloud services were unavailable in europe-west2-b zone; some services in europe-west2 region might experienced elevated errors and intermittent network issues. **Workaround** None as the issue has been mitigated.
- `2026-10-22 08:42` **resolved** — ## Preliminary Incident Report We recognize that Google Cloud services are critical to your organization, and we sincerely apologize for the disruption this incident may have caused to your business. We are working to address the root cause and prevent this from occurring in the future. Please note, this information is based on our best knowledge at the time of posting and is subject to change as our investigation continues. A final Incident Report with preventative actions will be posted once our investigation is complete. If you have experienced impact outside of what is listed below, please reach out to Google Cloud Support using [**https\://cloud.google.com/support**](https://cloud.google.com/support). ## Date/Time of the Issue Incident Start:19 Oct 2026 23:00 US/Pacific / 20 Oct 2026 07:00 BST Incident End: 20 Oct 2026 05:10 US/Pacific / 20 Oct 2026 13:10 BST Duration: 6 hours, 10 minutes ## Summary On Monday, 19 October 2026 (Tuesday, 20 October BST), multiple Google Cloud Platform services in the europe-west2 region experienced service disruptions, elevated error rates, Application Programming Interface (API) timeouts, and processing latency for a duration of 6 hours and 10 minutes. The disruption was triggered by an electrical power interruption at the datacenter facility hosting availability zone europe-west2-b, which temporarily took physical host servers, network routing hardware, and storage systems offline. This subsequently caused secondary failover and queuing delays across unimpacted regional infrastructure. ## Preliminary Root Cause The incident was triggered when a power grid disruption occurred at the datacenter facility hosting availability zone europe-west2-b. To maintain continuous power during the disruption, automatic transfer to facility backup power systems was triggered immediately. The facility load was transferred to the Uninterruptible Power Supply (UPS) battery banks during the start up sequence of on-site emergency diesel generators. The Programmable Logic Controller (PLC) automation system (both primary and backup PLCs) failed to initiate the logical automatic start/transfer operations of the generators, preventing backup generators from assuming the load. When UPS battery reserves were depleted, physical host servers, network switching fabric, and local storage arrays dedicated to europe-west2-b lost power and went offline. The sequence of events leading to customer impact was as follows: * At 22:53 US/Pacific (20 Oct, 06:53 BST), an electrical power disruption occurred at the datacenter facility hosting availability zone europe-west2-b, and a PLC (Programmable Logic Controller) startup fault prevented facility emergency generators from assuming the electrical load. Facility load continued to be supported by UPS battery banks until the batteries were depleted, starting at 23:00 US/Pacific (20 Oct, 07:00 BST). * When the physical host servers, local storage infrastructure, and network aggregation hardware in europe-west2-b lost power, it caused instances, database replicas, and worker nodes in that zone to become immediately unreachable. * Regional traffic-routing mechanisms and client applications initiated automated failover toward available infrastructure in adjacent availability zones within the region. * Cross-zone remote procedure calls from regional service control planes directed to metadata and catalog components in the offline zone encountered timeouts, causing initial job submission stalls and API latency spikes. * The sudden influx of redirected query executions, high-availability database failovers, and compute workloads saturated available compute slots and processing capacity in other availability zones, creating elevated admission queues, delayed database failover completions, and intermittent 500 and 503 API errors across the region. Google engineers have begun a full root cause analysis (RCA) and will provide additional information once it is available. ## Remediation Google engineers were alerted to the disruption via automated monitoring systems and probers at 23:08 US/Pacific on 19 Oct (20 Oct, 07:08 BST), which detected elevated API error ratios, job submission timeouts, and loss of infrastructure reachability in europe-west2-b. Engineering teams immediately mobilized to isolate the affected infrastructure and commence recovery: * **Traffic Isolation:** Google engineers isolated the offline availability zone (europe-west2-b) by adjusting regional service routing away from the affected infrastructure, preventing cross-zone client requests from timing out against offline endpoints. * **Workload Reprioritization and Resource Safeguards:** Engineers adjusted scheduler allocations and capacity limits across unimpacted infrastructure to prioritize active customer reservations, reduce resource contention on execution slots, and relieve queue backlogs for regional workloads. * **Power Grid Restoration:** On-site facility engineers conducted graceful restoration of a stable utility power feed once utility supply was back, within 1 hour. Following this, they focused and completed a reset and tested the emergency backup generator automation controllers to restore facility power. * **Progressive Infrastructure Re-activation:** Google engineers systematically restarted core Google Cloud infrastructure and services. * **Controlled Traffic Re-admission:** Once system health and storage integrity were verified, engineers executed gradual, controlled traffic re-admission back to zone europe-west2-b allowing accumulated job backlogs to clear smoothly. By 05:10 US/Pacific on Tuesday, 20 Oct(13:10 BST), service availability had recovered for the vast majority of customer workloads, and resolution for most of the customers across Google Cloud Platform services in the region was confirmed. Following the recovery of a small tail of remaining isolated workloads, the closure notification for the incident was posted by 07:03 US/Pacific (15:03 BST). ## Description of Impact Between Monday, 19 Oct 2026, 23:00 US/Pacific (Tuesday, 20 Oct 2026, 07:00 BST) and Tuesday, 20 October 2026, 05:10 US/Pacific (13:10 BST), customers utilizing Google Cloud Platform services in the europe-west2 region experienced service disruptions across two consecutive phases, lasting a total duration of 6 hours, 10 minutes. During the initial phase from 23:00 to 04:29 US/Pacific (07:00 to 12:29 BST) (duration 5 hours, 29 minutes), customers experienced elevated error rates, API timeouts, and processing delays until most Google Cloud Platform services recovered. Following platform service restoration, some customer workloads running on Compute Engine instances experienced a longer recovery tail. By 04:46 US/Pacific (12:46 BST) new Virtual Machine creation resumed and by 05:10 US/Pacific (13:10 BST) over 99% of Compute Engine Virtual Machines were in a healthy state. Any remaining customers that were experiencing Virtual Machine problems were directed to file a customer support issue as per normal procedure. While the physical power outage was isolated strictly to datacenter facilities in europe-west2-b, some customers running workloads in unimpacted zones (europe-west2-a and europe-west2-c) or across regional architectures also experienced unexpected impact. The primary causes were 1\) instantaneous demand surges on unimpacted compute capacity and 2\) synchronous control-plane validation dependencies across zonal boundaries. Affected products included: BigQuery, Looker, Google Compute Engine, Google Kubernetes Engine, Regional Persistent Disk, and Cloud SQL. These services experienced unexpected degradation and failover issues driven by sudden compute capacity exhaustion from surging demand, and control-plane API timeouts waiting on validations from the offline zone. Workloads and resources hosted entirely within europe-west2-b became unreachable. Data storage services with failover capability such as Cloud SQL, AlloyDB, Memorystore for Redis, and Cloud Bigtable failed over into nearby unimpacted zones during the europe-west2-b outage. Multi-zonal and regional services that spanned across the europe-west2 region experienced elevated latency, transient API errors (such as HTTP 500 and 503 errors), and task admission queuing as failover traffic concentrated onto unimpacted regional resources and cross-zone calls to offline components timed out. Storage write operations in the region experienced a transient error rate during the initial failover period, and regional query admission queues temporarily peaked before rapidly subsiding as capacity re-balanced. The following Google Cloud Platform services experienced impact: * AlloyDB for PostgreSQL * Apigee * BigQuery Data Transfer Service * Cloud Data Fusion * Cloud Filestore * Cloud Firestore * Cloud Load Balancing * Cloud Memorystore * Cloud NAT * Cloud Run * Cloud Spanner * Cloud Workstations * Datastream * Dialogflow CX * Google App Engine * Google BigQuery * Google Cloud Bigtable * Google Cloud Dataflow * Google Cloud Dataproc * Google Cloud NetApp Volumes * Google Cloud Pub/Sub * Google Cloud SQL * Google Cloud Storage (GCS) * Google Compute Engine (GCE) * Google Kubernetes Engine (GKE) * Google SecOps (Security Information and Event Management \[SIEM\] and Security Orchestration, Automation, and Response \[SOAR\]) * Looker (Google Cloud core) * Looker Studio * Managed Airflow (Google Cloud Composer) * Managed Service for Apache Kafka * Memorystore for Memcached * Memorystore for Redis * Memorystore for Redis Cluster * Memorystore for Valkey * Network Connectivity Center * Persistent Disk * Private Service Connect * Speech-to-Text * Virtual Private Cloud (VPC) * VMWare engine
- `2026-10-27 15:13` **resolved** — ## Incident Report ### **Summary** On Monday, 19 Oct 2026 (Tuesday, 20 Oct BST), a data center facility that hosts europe-west2-b, experienced an external electrical power disruption. This resulted in service unavailability for workloads located in this zone. As traffic was redirected to alternative cloud regions or unaffected zones within europe-west2 cloud region, these workloads experienced transient latency and elevated errors affecting services across europe-west2. The primary impact was mitigated after 6 hours 8 minutes. ### **Root Cause** **Facility architecture and background:** The data center facility in the UK region hosts Cloud clusters serving the europe-west2-b zone, which form part of the multi-zonal europe-west2 Cloud region. The site relies on a 2N utility power architecture—meaning the facility has complete duplicate power paths for every critical load—backed by a central high-voltage generator bank managed by a dual-redundant Programmable Logic Controller (PLC). Primary utility power is supplied to the facility via two independent feeders (A and B), supplemented by a standby utility connection configured for manual campus-wide transfer during maintenance operations or emergency events. The two high-voltage primary feeders terminate at separate primary intake switchboards capable of closed-loop operation if either board is taken offline. Switching between the primary and standby utility providers utilizes open transitions and key interlocks to prevent parallel connections across sources. Static Uninterruptible Power Supply (UPS) battery systems provide 5 minutes of autonomous runtime. This bridge power ensures uninterrupted operation while the PLC detects utility undervoltage events, initiates diesel generator startup, synchronizes the generators, and transfers the site load. Prior to the utility disruption, the facility had been placed in a temporary single-feed (1N) bridged maintenance posture on the primary utility. Circuit breaker feeder B was physically opened and key-interlocked (Locked Out Tagged Out) to allow for corrective repairs following issues identified during a quarterly generator load test on 14 Oct 2026. Consequently, primary utility feed A was supporting the total critical load while repairs were conducted on the B feeder pathway. **Root Cause:** On Monday, 19 Oct 2026 at 22:53 US/Pacific (Tuesday, 20 Oct 06:53 BST), the primary utility provider experienced an outage on primary utility feed A, caused by an external underground cable fault on the utility network outside the data center facility. When the primary utility feed A was lost, the facility’s generator plant failed to receive a signal to automatically start and take critical load due to a design voltage-sensing and misconfiguration in the PLCs. The facility's static UPS battery banks carried the critical load as designed until their battery reserves depleted resulting in a complete power loss across data halls. The investigation of the PLC configuration identified the root cause as a misconfiguration of the PLC logic that did not verify both an undervoltage condition and physical breaker position for both utility feeders. Instead, the PLC undervoltage sensing inputs were wired upstream (line side) of the isolated utility feed B breaker. The PLC was programmed to verify whether voltage was present on the feeder before commanding a generator start. As the utility feed B remained energized upstream, both PLCs sensed live upstream voltage on the disconnected path. As a result, the PLC logic indicated that a healthy utility source was available, even though the primary and secondary distribution feeds were de-energized, and did not issue a start command to the generator plant.; The onsite facilities team attempted to execute a remote load transfer via the PLC; however, the system inhibited the action. Recognizing that the UPS battery reserves would be exhausted before further remote operations could be resolved, the team transferred the facility load to the standby utility source, restoring the critical load that was lost once batteries depleted during this process. Once utility power was restored via the standby utility source, and UPS batteries started recharging, the onsite facility teams and OEM vendors reconfigured and validated PLC functionality on both undervoltage and breaker status. After the power loss was initially detected, Google engineers began work to mitigate failover failures and prepare service recovery. Upon restoration of the utility source, UPS and generator capabilities, they began the process to re-enable services in europe-west2-b. ### Remediation and Prevention The sequence of events leading to customer impact was as follows: * **19 Oct 22:53 US/Pacific (20 Oct BST 06:53 BST) -** An electrical power disruption occurred at the data center facility hosting availability zone europe-west2-b, and a PLC (Programmable Logic Controller) startup fault prevented facility emergency generators from assuming the electrical load. Facility load continued to be supported by UPS battery banks until the batteries were depleted. * **19 Oct 23:00 US/Pacific (20 Oct 07:00 BST)** **Impact start -** As battery backups were exhausted, physical host servers, local storage infrastructure, and network aggregation hardware in availability zone europe-west2-b lost power, causing instances, database replicas, and worker nodes in that zone to become immediately unreachable. * **19 Oct 23:04 US/Pacific (20 Oct 07:04 BST) -** Google engineers were alerted to the disruption via automated monitoring systems and probers, which detected elevated API error ratios, job submission timeouts, and loss of infrastructure reachability in zone europe-west2-b. Engineering teams immediately mobilized to isolate the affected infrastructure and commence recovery. * **19 Oct 23:09 US/Pacific (20 Oct 07:09 BST) -** Regional traffic-routing mechanisms and client applications initiated automated failover toward available infrastructure in adjacent availability zones within the region. * Cross-zone remote procedure calls from regional service control planes directed to metadata and catalog components in the offline zone encountered timeouts, causing initial job submission stalls and API latency spikes. * The sudden influx of redirected query executions, high-availability database failovers, and compute workloads saturated available compute slots and processing capacity in other availability zones, creating elevated admission queues, delayed database failover completions, and intermittent 500 and 503 API errors across the region. * **Traffic Isolation:** Google engineers isolated the offline availability zone (europe-west2-b) by adjusting regional service routing away from the affected infrastructure, preventing cross-zone client requests from timing out against offline endpoints. * **Workload Reprioritization and Resource Safeguards:** Engineers adjusted scheduler allocations and capacity limits across unimpacted infrastructure to prioritize active customer reservations, reduce resource contention on execution slots, and relieve queue backlogs for regional workloads. * **19 Oct 23:54 US/Pacific (20 Oct 07:54 BST) Utility Power Restoration** **\-** Power supply from the standby utility source was restored and power restoration to the data hall began. * **20 Oct 00:20 US/Pacific (08:20 BST) Datahall Power Restoration** **\-** Utility power was re-established to the datahalls. Mains utility power normalized; power feed to data hall verified. Following this, the site operations and engineering teams focused on and completed a reset and tested the emergency backup generator automation controllers to restore facility power. * **20 Oct 00:40 US/Pacific (08:40 BST)** **\-** Clusters began to recover, cluster manager systems started rescheduling compute jobs. * **20 Oct 00:45 US/Pacific (08:45 BST) \-** Network devices were back online. * **20 Oct 00:50 US/Pacific (08:50 BST) \-** Due to the immediacy of power loss, storage devices with unclean shutdowns began recovery. Responders recognized this and chose not to restore services that would otherwise make this worse. Because of the large volume of storage devices involved, this process required a substantial period to finish prior to restoring service to customers. * **20 Oct 03:13 US/Pacific (11:06 BST)** **\-** Power redundancy was verified and the storage infrastructure was sufficiently restored to prepare the initiation of controlled traffic re-admission. The remaining storage systems continued recovery operations. * **20 Oct 03:45 US/Pacific (11:45 BST) Controlled Traffic Re-admission** **\-** Once system health and storage integrity were verified, engineers executed gradual, controlled traffic re-admission back to zone europe-west2-b allowing accumulated job backlogs to clear smoothly. * **20 Oct 04:11 US/Pacific (12:11 BST)** **\-** Google Cloud Platform services and APIs were made fully available to customers. Some Google Compute Engine virtual machines were still not yet available. * **20 Oct 04:29 US/Pacific (12:29 BST) \-** Regional services and backend compute management systems were fully undrained, marking the end of core platform-level disruption. * **20 Oct 05:08 US/Pacific (13:08 BST) Majority Workload Recovery** **\-** GCE virtual machines were 99.5% operational and healthy, and creating new VMs was fully supported. Service availability was restored for most customer workloads, with resolution confirmed for the majority of users across Google Cloud Platform services in the region. Following the recovery of a small tail of remaining isolated workloads, the closure notification for the incident was posted by 07:03 US/Pacific (15:03 BST). Google is committed to preventing a repeat of this issue in the future and is completing the following actions: **Data Center Action Items:** * Restoration of normal utility configuration (both A & B feeds) * Utility provider to locate the damaged cable fault, using low-frequency radar penetration testing procedures. **\[Completed 26 Oct\]** * Utility provider to send on-site engineers to repair the cable fault and restore the impacted primary utility feed A source. **\[Estimated Completion Date 30 Oct\]** * Onsite Operations teams with a controlled method of procedure (MOP) to transfer the facility to its primary utility source (feeds A and B) from standby utility feed through coordinated switching with the utility provider **\[Estimated Completion Date 31 Oct\]** * Generator confidence testing * Full generator load transfer tests with utility power cutover of the active utility feed. Test automatic startup of generators through PLC signals within the backup battery timer and detailed review of backup procedures **\[Completed on 20 Oct and 26 Oct (full critical load transfer test including updated EOPs)\]** * Full generator load transfer tests in coordination with the primary utility source feeds (A and B) after restoration work completed by the primary utility provider. **\[Estimated Completion Date 31 Oct\]** * PLC configuration fix * Reconfigured and validated the PLC after the outage on 20 Oct and then again on 26 Oct **\[Completed 26 Oct\]** * Google engineers to work in collaboration with the facility vendor teams to perform a comprehensive PLC code review and system resiliency to multiple failure scenarios. **\[Estimated Completion Date (prior to generator transfer test) estimated 31 Oct\]** * Validate PLC telemetry once primary utility feed A and B lineup is restored. **\[Estimated Completion Date 31 Oct\]** * Validate global PLC configurations at all global sites with similar topologies and PLC manufacturers, to ensure confidence in automatic generator starts. **\[Estimated Completion Date 15 Nov\]** * Emergency Operating Procedures (EOP) documentation update: * Comprehensive update to the EOP for site’s response to unexpected loss of active utility source. Update includes \- manual steps to attempt to restore active utility feed, manual restart of generators and manual transfer to other utility sources within the UPS battery run time. **\[Completed 26 Oct\]** * Ensure updated EOP is practiced twice per shift by on-site facility team. **\[Completed 24 Oct and checked with Google DC Operations daily\]** * Strengthen change management: * Freeze planned maintenance activities at other sites with similar topologies, and implement elevated approvals and reviews for data center configuration changes which are outside their standard designs. **\[Estimated Completion Date 30 Oct\]** **Cloud Service Action Items:** * **Communications Enhancement:** Bolster notification speed by developing an automated alert-to-communication mechanism for location-impacting incidents. * **Platform Testing:** Expand platform-wide testing outcomes for infrastructure-level events. ### **Detailed Description of Impact** Between Monday, 19 Oct 2026, 23:00 US/Pacific (Tuesday, 20 Oct 2026, 07:00 BST) and Tuesday, 20 Oct 2026, 05:08 US/Pacific (13:08 BST), customers utilizing Google Cloud Platform services in the europe-west2 region experienced service disruptions across two consecutive phases, lasting a total duration of 6 hours, 08 minutes. During the initial phase from 23:00 to 04:29 US/Pacific (07:00 to 12:29 BST) (duration 5 hours, 29 minutes), customers experienced unavailability of zone-scoped resources in products, including: * **Compute Engine** \- VM instances, zonal MIGs, unmanaged instance groups, zonal NEGs, sole-tenant node groups, zonal reservations, target instances, zonal API operations. * **Persistent Disk / Hyperdisk** \- Zonal PD, zonal Hyperdisk, Hyperdisk storage pools, instant snapshots. * **GKE \-** Zonal clusters (control plane), node pool nodes in \-b, PVs on zonal disks. * **Cloud SQL \-** Non-HA instances, read replicas in \-b. * **AlloyDB \-** Basic (non-HA) primaries in \-b. * **Bigtable \-** Single-cluster instances in \-b. * **Other products with affected zonal resources:** Cloud Batch, Cloud Build, Cloud Storage, Cloud Workstations, Composer, Dataproc, Filestore, Managed Lustre, Migrate to VMs, NetApp Volumes, Pub/Sub Lite, Vertex AI Workbench, VMware Engine. The unavailability of those resources scoped to a zone that undergoes disruption is working within the design of a zonally-scoped system. Requests made to these resulted either in a timeout or error code during the disruption. Any workloads and resources hosted entirely within europe-west2-b became unreachable due to this event. Google Compute Engine's zonal resources such as VMs require customers to provision failover or cross-zone recovery resources if needed, and to set up a recovery mechanism that matches their needs. Some failovers were unsuccessful due to capacity constraints and difficulty in replicating the underlying storage mechanism. Following platform service restoration, some customer workloads running on Compute Engine instances as well as products that rely on Compute Engine experienced a longer recovery. By 04:46 US/Pacific (12:46 BST) new Virtual Machine creation resumed and by 05:08 US/Pacific (13:08 BST) over 99% of Compute Engine Virtual Machines were in a healthy state. Any remaining customers that were experiencing Virtual Machine problems were directed to file a customer support case as per normal procedure. While the physical power outage was isolated strictly to data center facilities in zone europe-west2-b, some customers running workloads in unimpacted zones (europe-west2-a and europe-west2-c) or across regional architectures also experienced unexpected impact. The primary causes of impact outside of zone europe-west2-b were: 1\) Instantaneous demand surges on neighboring zones (a and c) and 2\) Control-plane dependencies across zonal boundaries Affected products included: AlloyDB, Apigee Infra and Apigee Edge, BigQuery, Cloud Dataflow, Cloud Pub/Sub, Cloud SQL, Google Compute Engine, Google Kubernetes Engine, Looker, Regional Persistent Disk, Vertex AI Prediction, and Cloud Speech-to-Text. These services experienced unexpected degradation and failover issues driven by sudden compute capacity exhaustion from surging demand, and control-plane API timeouts. Customers would have experienced the inability to create, update, or delete some regional resources during this period. This was not a full regional control-plane outage, but certainly a degraded state. Multi-zonal and regional services that spanned across the europe-west2 region experienced elevated latency, transient API errors (such as HTTP 500 and 503 errors), and task admission queuing as failover traffic concentrated onto unimpacted regional resources and cross-zone calls to offline components timed out. Storage write operations in the region experienced a transient error rate during the initial failover period, and regional query admission queues temporarily peaked before rapidly subsiding as capacity re-balanced. We did not identify any impact of the outage outside of the europe-west2 region. The following Google Cloud Platform services experienced impact within europe-west2: * AlloyDB for PostgreSQL * Apigee * BigQuery Data Transfer Service * Cloud Data Fusion * Cloud Filestore * Cloud Firestore * Cloud Load Balancing * Cloud Memorystore * Cloud NAT * Cloud Run * Cloud Spanner * Cloud Workstations * Datastream * Dialogflow CX * Google App Engine * Google BigQuery * Google Cloud Bigtable * Google Cloud Dataflow * Google Cloud Dataproc * Google Cloud NetApp Volumes * Google Cloud Pub/Sub * Google Cloud SQL * Google Cloud Storage (GCS) * Google Compute Engine (GCE) * Google Kubernetes Engine (GKE) * Google SecOps (Security Information and Event Management \[SIEM\] and Security Orchestration, Automation, and Response \[SOAR\]) * Looker (Google Cloud core) * Looker Studio * Managed Airflow (Google Cloud Composer) * Managed Service for Apache Kafka * Memorystore for Memcached * Memorystore for Redis * Memorystore for Redis Cluster * Memorystore for Valkey * Network Connectivity Center * Persistent Disk * Private Service Connect * Speech-to-Text * Virtual Private Cloud (VPC) * VMWare engine ----------
- `2026-11-04 21:15` **resolved** — # Incident Report ## Summary On Monday, 19 Oct 2026 at 22:53 US/Pacific (Tuesday, 20 Oct 2026, 06:53 BST), a data center facility that hosts europe-west2-b, a zone within the europe-west2 cloud region, experienced an external electrical power disruption. This resulted in service unavailability for workloads located in this zone. The subsequent failover of these workloads to unaffected zones and alternate regions caused transient latency and elevated error rates across the broader europe-west2 region. The resolution of the main impact was completed within 6 hours and 8 minutes. ## Root Cause **Facility architecture and background:** The data center facility in the UK region hosts Cloud clusters serving the europe-west2-b zone, which form part of the multi-zonal europe-west2 Cloud region. The site relies on a 2N utility power architecture—where the facility has complete duplicate power paths for every critical load—backed by a central high-voltage generator bank managed by a dual-redundant Programmable Logic Controller (PLC). Primary utility power is supplied to the facility via two independent feeders (A and B), supplemented by a standby utility connection configured for manual campus-wide transfer during maintenance operations or emergency events. The two high-voltage primary feeders (A and B) terminate at separate primary intake switchboards capable of closed-loop operation if either board is taken offline. Switching between the primary and standby utility providers utilizes open transitions and key interlocks to prevent parallel connections across sources. Uninterruptible Power Supply (UPS) battery systems provide approximately five minutes of autonomous runtime. This bridge power ensures uninterrupted operation while the PLC detects utility undervoltage events, initiates diesel generator startup, synchronizes the generators, and transfers the site load. Prior to the utility disruption, the facility had been placed in a temporary single-feed (1N) bridged maintenance posture on the primary utility. Circuit breaker feeder B was physically opened to allow for corrective repairs following issues identified during a quarterly generator load test on 14 Oct 2026. Consequently, primary utility feed A was supporting the total critical load while the feed B breaker remained in an open position to allow repairs to be conducted on the B feeder pathway. **Root Cause:** On Monday, 19 Oct 2026 at 22:53 US/Pacific (Tuesday, 20 Oct 2026, 06:53 BST), the primary utility provider experienced an outage on primary utility feed A, caused by an external underground cable fault on the utility network outside the data center facility. Following the loss of primary utility feed A, a voltage-sensing design flaw and PLC misconfiguration prevented the facility’s generator plant from receiving an automatic start signal to assume the critical load. Consequently, static UPS battery banks sustained the critical load as intended until their reserves were used up, leading to a loss of power across the data halls. The investigation of the PLC configuration identified the root cause as a misconfiguration of the PLC logic that did not verify both an undervoltage condition and physical breaker position for both utility feeders. Instead, the PLC undervoltage sensing inputs were wired upstream (line side) of the isolated utility feed breakers. The PLC was programmed to verify whether voltage was present on the line side of the feeder before commanding a generator start. At the time of the power event on 19 Oct 2026, the Utility feed B remained energized upstream while its incomer feed was in open position. Consequently, the active and standby master PLCs, which were running identical control logic, sensed live voltage from feed B through shared upstream sensors on the disconnected path. As a result, the PLC logic indicated that a healthy utility source was available, even though the primary and secondary distribution feeds were de-energized, and did not issue a start command to the generator plant. The onsite facilities team attempted to execute a remote load transfer via the PLC; however, the system inhibited the action due to the live upstream voltage reading on feed B. Recognizing that the UPS battery reserves would be used up before further remote operations could be resolved, the team transferred the facility load to the standby utility source, restoring the critical load that was lost once batteries depleted during this process. Once utility power was restored via the standby utility source, and UPS batteries started recharging, the onsite facility teams and OEM vendors reconfigured and validated PLC functionality on both undervoltage and breaker status. Following the initial detection of the power outage, Google engineering teams initiated efforts to address failover issues and lay the groundwork for service restoration. Once the utility feed, UPS systems, and generator capabilities were successfully restored, engineers initiated the recovery procedure to bring europe-west2-b services back online. ## Remediation and Prevention The sequence of events leading to customer impact was as follows: * **19 Oct 22:53 US/Pacific (20 Oct 2026, 06:53 BST) -** An electrical power disruption occurred at the data center facility hosting availability zone europe-west2-b, and a PLC (Programmable Logic Controller) startup fault prevented facility emergency generators from assuming the electrical load. Facility load continued to be supported by UPS battery banks until the batteries were depleted. * **19 Oct 23:00 US/Pacific (20 Oct 2026, 07:00 BST) - Impact start -** As battery backups were used up, physical host servers, local storage infrastructure, and network aggregation hardware in availability zone europe-west2-b lost power, causing instances, database replicas, and worker nodes in that zone to become immediately unreachable. * **19 Oct 23:04 US/Pacific (20 Oct 2026, 07:04 BST) -** Google engineers were alerted to the disruption via automated monitoring systems and probers, which detected elevated API error ratios, job submission timeouts, and loss of infrastructure reachability in zone europe-west2-b. Engineering teams immediately mobilized to isolate the affected infrastructure and commence recovery. * **19 Oct 23:09 US/Pacific (20 Oct 2026, 07:09 BST) -** Regional traffic-routing mechanisms and client applications initiated automated failover toward available infrastructure in adjacent availability zones within the region. * Cross-zone remote procedure calls from regional service control planes directed to metadata and catalog components in the offline zone encountered timeouts, causing initial job submission stalls and API latency spikes. * The sudden influx of redirected query executions, high-availability database failovers, and compute workloads saturated available compute slots and processing capacity in other availability zones, creating elevated admission queues, delayed database failover completions, and intermittent 500 and 503 API errors across the region. * **Traffic Isolation:** Google engineers isolated the offline availability zone (europe-west2-b) by adjusting regional service routing away from the affected infrastructure, preventing cross-zone client requests from timing out against offline endpoints. * **Workload Reprioritization and Resource Safeguards:** Engineers adjusted scheduler allocations and capacity limits across unimpacted infrastructure to prioritize active customer reservations, reduce resource contention on execution slots, and relieve queue backlogs for regional workloads. * **19 Oct 23:54 US/Pacific (20 Oct 2026, 07:54 BST) - Utility Power Restoration-** Power supply from the standby utility source was restored and power restoration to the data hall began. * **20 Oct 00:20 US/Pacific (20 Oct 2026, 08:20 BST) - Datahall Power Restoration-** Utility power was re-established to the datahalls. Mains utility power normalized; power feed to data hall verified. Following this, the site operations and engineering teams focused on and completed a reset and tested the emergency backup generator automation controllers to restore facility power. * **20 Oct 00:40 US/Pacific (20 Oct 2026, 08:40 BST)-** Clusters began to recover, cluster manager systems started rescheduling compute jobs. * **20 Oct 00:45 US/Pacific (20 Oct 2026, 08:45 BST) -** Network devices were back online. * **20 Oct 00:50 US/Pacific (20 Oct 2026, 08:50 BST) -** Due to the immediacy of power loss, storage devices with unclean shutdowns began recovery. Responders recognized this and chose not to restore services that would otherwise make this worse. Because of the large volume of storage devices involved, this process required a substantial period to finish prior to restoring service to customers. * **20 Oct 03:13 US/Pacific (20 Oct 2026, 11:13 BST)-** Power redundancy was verified and the storage infrastructure was sufficiently restored to prepare the initiation of controlled traffic re-admission. The remaining storage systems continued recovery operations. * **20 Oct 03:45 US/Pacific (20 Oct 2026, 11:45 BST)-** Once system health and storage integrity were verified, engineers executed gradual, controlled traffic re-admission back to zone europe-west2-b allowing accumulated job backlogs to clear smoothly. * **20 Oct 04:11 US/Pacific (20 Oct 2026, 12:11 BST)-** Google Cloud Platform services and APIs were made fully available to customers. Some Google Compute Engine virtual machines were still not yet available. * **20 Oct 04:29 US/Pacific (20 Oct 2026, 12:29 BST) -** Regional services and backend compute management systems were fully undrained, marking the end of core platform-level disruption. * **20 Oct 05:08 US/Pacific (20 Oct 2026, 13:08 BST) - Majority Workload Recovery-** GCE virtual machines were 99.5% operational and healthy, and creating new VMs was fully supported. Service availability was restored for most customer workloads, with resolution confirmed for the majority of users across Google Cloud Platform services in the region. Following the recovery of a small tail of remaining isolated workloads, the closure notification for the incident was posted by 07:03 US/Pacific (20 Oct 2026, 15:03 BST). Google is committed to preventing a repeat of this issue in the future and is completing the following actions: **Data Center Action Items:** * Restoration of normal utility configuration (both A & B feeds) * Utility provider to locate the damaged cable fault, using low-frequency radar penetration testing procedures. **[Completed 26 Oct]** * Utility provider to repair the cable, test and restore the impacted primary utility feed A source. **[Estimated Completion Date 13 Nov]** * Onsite Operations teams with a controlled method of procedure (MOP) to transfer the facility to its primary utility source (feeds A and B) from standby utility feed through coordinated switching with the utility provider **[Estimated Completion Date 17 Nov]** * Generator confidence testing * Full generator load transfer tests with utility power cutover of the active utility feed. Test automatic startup of generators through PLC signals within the backup battery timer and detailed review of backup procedures **[Completed on 20 Oct and 26 Oct (full critical load transfer test including updated EOPs)]** * Full generator load transfer tests in coordination with the primary utility source feeds (A and B) after restoration work completed by the primary utility provider. **[Estimated Completion Date 17 Nov]** * PLC configuration fix * Reconfigured and validated the PLC after the outage on 20 Oct and then again on 26 Oct **[Completed 26 Oct]** * Google engineers to work in collaboration with the facility vendor teams to perform a comprehensive PLC code review and system resiliency to multiple failure scenarios. **[Estimated Completion Date (prior to generator transfer test) 06 Nov]** * Validate PLC telemetry once primary utility feed A and B lineup is restored. **[Estimated Completion Date 17 Nov]** * Validate global PLC configurations at sites with similar topologies and PLC manufacturers **[Estimated Completion Date 15 Nov]** * Emergency Operating Procedures (EOP) documentation update: * Comprehensive update to the EOP for site’s response to unexpected loss of active utility source. Update includes - manual steps to attempt to restore active utility feed, manual restart of generators and manual transfer to other utility sources within the UPS battery run time. **[Completed 26 Oct]** * Ensure updated EOP is practiced twice per shift by on-site facility team. **[Completed 24 Oct and checked with Google DC Operations daily]** * Detect and prevent future PLC misconfigurations * Google to ensure PLC code reviews are fully conducted and validated during facility commissioning and Integrated System Tests (ISTs) prior to Facility Ready (FR). A second Google team will then validate this within a quarter post-FR. **[Estimated Completion Date 18 Dec]** * Strengthen change management: * Freeze planned maintenance activities at other sites with similar topologies, and implement elevated approvals and reviews for data center configuration changes which are outside their standard designs. **[Completed Date 30 Oct]** **Cloud Service Action Items:** * **Communications Enhancement:** Bolster notification speed by developing an automated alert-to-communication mechanism for location-impacting incidents. * **Platform Testing:** Expand platform-wide testing outcomes for infrastructure-level events. ## Detailed Description of Impact Between Monday, 19 Oct 2026, 23:00 US/Pacific (Tuesday, 20 Oct 2026, 07:00 BST) and Tuesday, 20 Oct 2026, 05:08 US/Pacific (Tuesday, 20 Oct 2026, 13:08 BST), customers utilizing Google Cloud Platform services in the europe-west2 region experienced service disruptions across two consecutive phases, lasting a total duration of 6 hours, 08 minutes. During the initial phase from 23:00 US/Pacific (07:00 BST) to 04:29 US/Pacific (12:29 BST) (duration 5 hours, 29 minutes), customers experienced unavailability of zone-scoped resources in products, including: * **Compute Engine** - VM instances, zonal MIGs, unmanaged instance groups, zonal NEGs, sole-tenant node groups, zonal reservations, target instances, zonal API operations. * **Persistent Disk / Hyperdisk** - Zonal PD, zonal Hyperdisk, Hyperdisk storage pools, instant snapshots. * **GKE -** Zonal clusters (control plane), node pool nodes in -b, PVs on zonal disks. * **Cloud SQL -** Non-HA instances, read replicas in -b. * **AlloyDB -** Basic (non-HA) primaries in -b. * **Bigtable -** Single-cluster instances in -b. * **Other products with affected zonal resources:** Cloud Batch, Cloud Build, Cloud Storage, Cloud Workstations, Composer, Dataproc, Filestore, Managed Lustre, Migrate to VMs, NetApp Volumes, Pub/Sub Lite, Vertex AI Workbench, VMware Engine. The unavailability of those resources scoped to a zone that undergoes disruption is working within the design of a zonally-scoped system. Requests made to these resulted either in a timeout or error code during the disruption. Workloads and resources hosted entirely within europe-west2-b became unreachable due to this event. While the physical power outage was isolated strictly to data center facilities in zone europe-west2-b, some customers running workloads in unimpacted zones (europe-west2-a and europe-west2-c) or across regional architectures also experienced unexpected impact. Google Cloud designs regional services to be robust to zone failures, but in this case there were a set of unexpected failures which caused impact outside of the zonal failure domain. We are committed to learning from this incident and preventing zonal events from impacting broader regional operations. The primary causes of impact outside of zone europe-west2-b were: 1) Instantaneous demand surges on neighboring zones (a and c) and 2) Control-plane dependencies across zonal boundaries See below for more detailed descriptions of impact and next steps for each impacted product. ### Google Compute Engine and Persistent Disk **Impact:** Upon power loss in europe-west2-b at 22:53 US/Pacific (06:53 BST), all Compute Engine VMs and zonal Persistent Disk and Hyperdisk volumes in the zone became unavailable. Google Compute Engine's zonal resources, such as VMs, require customers to provision cross-zone failover and recovery mechanisms tailored to their needs. During this incident, however, some of these failover attempts were unsuccessful due to capacity constraints and difficulty in replicating the underlying storage mechanism. Once data hall power was restored at 00:20 US/Pacific (08:20 BST), engineers worked to confirm storage integrity before restoring compute services. Following platform service restoration, some customer workloads running on Compute Engine instances as well as products that rely on Compute Engine experienced a longer recovery due to action taken by responders in order to prevent data loss. Once storage stabilized, engineers brought VMs and zonal disks in the zone back online between 04:01 US/Pacific (12:01 BST) and 04:29 US/Pacific (12:29 BST). By 04:46 US/Pacific (12:46 BST) new VM creation resumed, and by 05:08 US/Pacific (13:08 BST) over 99% of Compute Engine VM were in a healthy state. (The remaining VMs were restarted after their disks recovered). Any remaining customers that were experiencing VM problems were directed to file a customer support case as per normal procedure. **Next Steps:** Google Compute Engine and Persistent Disk engineering teams are conducting a comprehensive postmortem review of this event along with Google data center and infrastructure teams. The follow-up actions focus on preventing a recurrence and on making our clusters recover from an unplanned power loss more quickly and safely. We are also expanding our regular disaster-recovery testing for full cluster power-loss scenarios to validate end-to-end Compute Engine recovery. The zonal failure in a facility within europe-west2-b exhibited sequencing atypical of previous zonal outages. We are investigating additional signals to incorporate into our failure detection capabilities for Regional Persistent Disk, so that should this sequence recur, Regional Persistent Disk can differentiate this case from the more typical zonal outage scenarios. We exercise caution with all changes to detection mechanisms to ensure we do not regress our capabilities to detect and react to more typical outages. ### Cloud SQL **Impact:** Upon power loss, Cloud SQL detected failures and immediately initiated automated failovers for all regional HA instances. The failover was complete for regional instances that did not have Regional Persistent Disk failures (see above), mitigating their impact right away. A majority of the remaining instances recovered by 20 Oct 2026, 01:00 US/Pacific (09:00 BST) once the Regional Persistent Disk recovered, shortly after europe-west2-b recovery. A small number of remaining instances were fully recovered by 20 Oct 2026, 11:00 US/Pacific (19:00 BST). **Next Steps:** Our engineering team is implementing improvements across our database and storage infrastructure to ensure that High Availability instances using Regional Persistent Disks recover more rapidly during unplanned outages. ### Google Kubernetes Engine **Impact:** The power loss in europe-west2-b caused total zonal failure, autoscaling stalls, API error spikes, and regional pod unavailability across GKE clusters. Total incident duration for GKE was 7 hours and 34 minutes, starting at 23:02 US/Pacific (07:02 BST), reaching initial mitigation at 05:22 US/Pacific (13:22 BST), and achieving final resolution at 06:37 US/Pacific (14:37 BST). Zonal clusters in europe-west2-b experienced 100% unavailability between 23:05 US/Pacific (07:05 BST) and 04:30 US/Pacific (12:30 BST), with full restoration by 06:00 US/Pacific (14:00 BST). Regional clusters lost worker nodes in europe-west2-b, causing a spike in regional unavailable pods. Worker nodes scaled up in neighboring zones europe-west2-a and europe-west2-c. The scale ups in surviving zones were impacted by available capacity. Autoscaling was slowed down due to increased load and failures of underlying APIs, such as Google Compute Engine (see above). Regional API operations suffered increased failure rates for cluster creation, deletion, updates and node pool operations due to reduced capacity and increased load on underlying APIs (see above). ### Google Cloud Dataflow **Impact:** Dataflow mostly behaved as expected for a zonal failure: new batch jobs were automatically routed away from europe-west2-b within about one minute, and all workloads (streaming and batch) recovered with no data loss once power returned. Jobs running in the affected zone also resumed responding to user cancellation requests once the zone came back up. Additionally, we identified and have now fixed an issue that caused increased streaming job start latency for about two hours. Google engineering is actively working to resolve a zone-selection issue that incorrectly routed a small number of new jobs to the unavailable zone. ### Looker **Impact:** The power event interrupted Looker's connectivity to GKE and CloudSQL, which Looker uses for its backend compute and metadata database resources, respectively. About 45% of instances became unresponsive as Looker's messaging middle tier was unable to recover and reestablish database connectivity due to degraded GKE and Cloud SQL availability. A manual rolling restart of messaging infrastructure across all instances brought unavailable Looker instances down to about 15% within about two hours of onset. Subsequently, regional control-plane issues in GKE pushed instance unavailability back to about 35% before settling near 20% for several hours. Cloud SQL failover degradation extended the delay as well. Availability returned to pre-incident baseline by 05:20 US/Pacific (13:20 BST) as the GKE outage was mitigated; the incident was declared mitigated at 07:44 US/Pacific (15:44 BST). **Next Steps**: Engineering teams are working to expand zonal outage telemetry and improve Looker's messaging middleware failovers to shorten recovery times during zonal unavailability. ### AlloyDB **Impact:** Customer impact for AlloyDB began at 23:07 US/Pacific (07:07 BST). Regional primary AlloyDB instances experienced up to a 95-minute delay in starting automated failovers due to delayed failover notifications. ~4% of these instances experienced a stalled failover due to Persistent Disk issues that required manual recovery. Additionally, zonal primary instances, zonal secondary instances, and single-node read pool instances encountered downtime directly aligned with the zone's power loss. Over 99% of affected AlloyDB instances in the region were restored to a healthy state by 06:16 US/Pacific (14:16 BST), with the final zonal primary AlloyDB instances recovering by 08:28 US/Pacific (16:28 BST). **Next Steps:** AlloyDB engineering team will be updating their monitoring system to detect zonal outage reliably and send the failover notification. The AlloyDB team is also partnering with the Persistent Disk team to ensure that high availability instances which leverage regional disk will recover faster during unplanned downtime events. ### BigQuery **Impact:** BigQuery experienced a service disruption in the europe-west2 region between 23:03 US/Pacific (07:03 BST) and 04:12 US/Pacific (12:12 BST) (approximately 5 hours and 9 minutes). During this incident, affected customers experienced significant increases in job latencies and high query queuing, resulting in delayed job execution. Customers also experienced intermittent 503 UNAVAILABLE errors on Job API calls (InsertJob, GetQueryResult). Similarly, customers received errors when calling AppendRows on buffered-type streams also after the incident onset until the streams were re-created. Finally, customers experienced reduced processing efficiency and elevated job failure rates across affected projects. The majority of the impact was confined to customers assigned to the impacted zone and was caused by an autoscaling dependency in a critical data plane job, which was unable to scale up during the failover. Some customers not assigned to the impacted zone saw elevated (1%) error rates on JobAPI calls due to overly broad concurrency throttling in the Job API component. **Next Steps:** Engineering teams have updated job provisioning to eliminate the dependency on autoscaling and are enhancing isolation mechanisms for the job server to ensure that future zonal disruptions are contained. ### Memorystore Redis **Impact:** During the power disruption, Basic Tier (non-HA) instances located in europe-west2-b experienced full unavailability, while Standard Tier (HA) instances initiated automated failovers or operated in a degraded state with missing replicas. Following the restoration of the zone, a subset of instances remained stuck in a REPAIRING state. For these affected instances, the data plane remained fully available and online for customer use, but control plane operations such as scaling and RDB exports were temporarily blocked. Google engineers manually mitigated this issue and fully repaired all affected instances within 24 hours. **Next Steps:** The team is working on mechanisms to detect cases when the control plane and data plane get out of sync, then reconcile them so that the system can automatically recover. ### Apigee X **Impact:** During the power disruption, the Apigee X runtime in the region returned elevated HTTP 5xx errors for 7 hours and 19 minutes. Apigee X is **designed to withstand** a single-zone loss as its datastore keeps replicas across the region's three zones, but a scheduling misconfiguration had placed two of three replicas in the zone that lost power, so the datastore lost quorum and the runtime failed requests instead of failing over. The errors were concentrated in a small number of affected clusters rather than spread across the region. As remediation, engineers restored the affected replicas into nearby zones. Apigee Hybrid clusters were not severely impacted. **Next Steps:** The engineering team has already audited the Apigee fleet to rebalance unevenly spread instances without downtime, and added monitoring to alert if datastore replicas drift across zones. Moving forward, the team will modify Apigee X to strictly enforce that every instance's three datastore replicas always run in three different zones. Additionally, the team is improving alert handling practices to help engineers better focus on affected services during similar outages. ### Apigee Edge **Impact:** During the power disruption, Apigee Edge clusters that had a single node in the affected zone successfully auto-recovered once power was restored. However, clusters with multiple nodes in the affected zone experienced extended recovery times. The automated node replacement system didn't handle multiple nodes going down simultaneously. This required manual intervention by Google engineers to restore the affected nodes. Furthermore, during this manual recovery, some virtual machines experienced a timeout when attempting to reach external OS repositories, causing them to boot with an older binary version and further delaying full restoration. **Next Steps:** The engineering team will implement retries and scope repository updates to prevent pre-hook timeouts when reaching external OS repositories during the boot process. The team is also investigating further multiple-node failure scenario resilience improvements. ### Cloud Bigtable **Impact:** Cloud Bigtable behaved as expected during the europe-west2-b power outage. All single-cluster instances located in europe-west2-b experienced expected zonal unavailability while the zone lost power and recovered. Multi-cluster instances configured with High Availability (HA) routing automatically failed over to healthy replica clusters in other zones where replica clusters and the control plane remained available with sufficient capacity to absorb the failover load. ### List of impacted services We did not identify any impact of the outage outside of the europe-west2 region. The following Google Cloud Platform services experienced impact within europe-west2: * AlloyDB for PostgreSQL * Apigee * BigQuery Data Transfer Service * Cloud Data Fusion * Cloud Filestore * Cloud Firestore * Cloud Load Balancing * Cloud Memorystore * Cloud NAT * Cloud Run * Cloud Spanner * Cloud Workstations * Datastream * Dialogflow CX * Google App Engine * Google BigQuery * Google Cloud Bigtable * Google Cloud Dataflow * Google Cloud Dataproc * Google Cloud NetApp Volumes * Google Cloud Pub/Sub * Google Cloud SQL * Google Cloud Storage (GCS) * Google Compute Engine (GCE) * Google Kubernetes Engine (GKE) * Google SecOps (Security Information and Event Management [SIEM] and Security Orchestration, Automation, and Response [SOAR]) * Looker (Google Cloud core) * Looker Studio * Managed Airflow (Google Cloud Composer) * Managed Service for Apache Kafka * Memorystore for Memcached * Memorystore for Redis * Memorystore for Redis Cluster * Memorystore for Valkey * Network Connectivity Center * Persistent Disk * Private Service Connect * Speech-to-Text * Virtual Private Cloud (VPC) * VMware Engine

### 2. grafana — Some Grafana Instances Experiencing Issues Coming Online after Restart

- URL: https://stspg.io/v6j9g9gqpwsw
- Started: 2026-10-23 13:33 UTC
- Resolved: 2026-10-23 18:16 UTC
- Duration: 4h 42m
- Impact: minor, status: resolved
- Components: AWS Australia - prod-au-southeast-1, AWS Australia - prod-au-southeast-2, AWS Brazil - prod-sa-east-1, AWS Canada - prod-ca-east-0, AWS Canada prod-ca-central-1, AWS East (VA) prod-us-east-1, AWS Germany - prod-eu-west-2, AWS Germany - prod-eu-west-4, AWS Germany prod-eu-central-1, AWS India - prod-ap-south-1, AWS Indonesia - prod-ap-southeast-2, AWS Indonesia - prod-ap-southeast-3, AWS Ireland - prod-eu-west-6, AWS Ireland eu-west-1, AWS Japan - prod-ap-northeast-0, AWS Singapore - prod-ap-southeast-1, AWS Sweden - prod-eu-north-0, AWS Sweden eu-north-1, AWS Switzerland - prod-eu-central-0, AWS Switzerland eu-central-2, AWS UAE - prod-me-central-1, AWS UK - prod-gb-south-1, AWS US East (OH) - prod-us-east-0, AWS US East (OH) - prod-us-east-2, AWS US East (VA) - prod-us-east-3, AWS US West - prod-us-west-0, AWS West prod-us-west-2, Azure Netherlands - prod-eu-west-3, Azure UAE - prod-me-central-2, Azure US Central - us-central2, Federal Cloud - AWS US Gov West, GCP Australia - prod-au-southeast-0, GCP Belgium - prod-eu-west-0, GCP Brazil - prod-sa-east-0, GCP India - prod-ap-south-0, GCP Saudi Arabia - prod-me-central-0, GCP Singapore - prod-ap-southeast-0, GCP UK - prod-gb-south-0, GCP US Central - prod-us-central-0, GCP US Central - prod-us-central-3, GCP US Central - prod-us-central-4, GCP US Central - prod-us-central-7, GCP US East - prod-us-east-1, play.grafana.org
- Canonical incident: [ci-a154ea27c29c7b2ae08dc805](#ci-a154ea27c29c7b2ae08dc805)

- `2026-10-23 13:33` **identified** — Starting at approximately 12:37 UTC, a some Hosted Grafana instances across multiple regions have been experiencing availability issues. Our team has identified the root cause and a fix is currently in progress. Affected instances will recover automatically once the fix has been applied. No action is required on your end. We will provide further updates as they become available and apologize for any disruption this may cause.
- `2026-10-23 15:04` **identified** — We are continuing to work on a fix for this issue.
- `2026-10-23 16:16` **monitoring** — A fix as been implemented and Hosted Grafana instances are back to normal levels. We will continue to monitor the issue.
- `2026-10-23 18:16` **resolved** — This incident has been resolved.

### 3. grafana — Grafana Cloud unavailibity in the prod-gb-south-0 region

- URL: https://stspg.io/blcsc38zxjxw
- Started: 2026-10-20 12:43 UTC
- Resolved: 2026-10-20 16:33 UTC
- Duration: 3h 49m
- Impact: minor, status: resolved
- Components: Cost Management and Billing, Fleet Management, GCP UK - prod-gb-south-0, GCP UK - prod-gb-south-0, GCP UK - prod-gb-south-0, GCP UK - prod-gb-south-0, GCP UK - prod-gb-south-0, GCP UK - prod-gb-south-0, GCP UK - prod-gb-south-0, GCP UK - prod-gb-south-0, GCP UK - prod-gb-south-0, GCP UK - prod-gb-south-0: API, GCP UK - prod-gb-south-0: Alertmanager, GCP UK - prod-gb-south-0: Alertmanager and Rules Configuration API, GCP UK - prod-gb-south-0: Ingestion, GCP UK - prod-gb-south-0: Ingestion, GCP UK - prod-gb-south-0: Metrics Generator, GCP UK - prod-gb-south-0: Public Probes, GCP UK - prod-gb-south-0: Querying, GCP UK - prod-gb-south-0: Querying, GCP UK - prod-gb-south-0: Rule Evaluation, Grafana.com, Incident Management and Response (IRM), Integrations, Public Snapshots, Support Tickets
- Canonical incident: [ci-261da6711cd4b1fbba5acb7e](#ci-261da6711cd4b1fbba5acb7e)

- `2026-10-20 06:49` **investigating** — As of 06:30 UTC, we became aware of to an issue with Grafana Cloud unavailibity in the prod-gb-south-0 region. Engineering is actively engaged and assessing the issue. We will provide updates accordingly.
- `2026-10-20 06:54` **investigating** — We are continuing to investigate this issue.
- `2026-10-20 11:37` **investigating** — We are continuing to investigate this issue.
- `2026-10-20 12:43` **identified** — Between 06:15 and 12:15 UTC, Grafana Cloud Frontend Observability in GCP UK (prod-gb-south-0) experienced errors accepting incoming telemetry data. This has been resolved and Frontend Observability is operating normally in this region. We are continuing to investigate and monitor the remaining impacted services.
- `2026-10-20 14:29` **identified** — We are seeing errors in Frontend o11y again. We are continuing to investigate and monitor the impacted services.
- `2026-10-20 16:33` **resolved** — This incident has been resolved.

### 4. cloudflare — Cloudflare Secure Web Gateway and DLP enforcement issues for Google Gemini

- URL: https://www.cloudflarestatus.com/incidents/3z8y23zn7tr2
- Started: 2026-10-19 20:30 UTC
- Resolved: 2026-10-20 22:46 UTC
- Duration: 1d 2h 16m
- Impact: minor, status: resolved
- Components: Data Loss Prevention (DLP), Gateway
- Canonical incident: [ci-a17ba2cdc8c4ab17a0eee508](#ci-a17ba2cdc8c4ab17a0eee508)

- `2026-10-19 20:30` **identified** — Cloudflare has identified an issue affecting Gateway application controls and Data Loss Prevention (DLP) inspection for certain Google Gemini traffic. We are working on a fix and will provide further updates as more information becomes available.
- `2026-10-19 23:46` **identified** — We are rolling out a fix to restore Secure Web Gateway application controls for Google Gemini. Work continues to restore DLP prompt logging and enforcement.
- `2026-10-20 22:46` **resolved** — This incident has been resolved.

### 5. github — Project updates delayed

- URL: https://stspg.io/sx51v1tlpqjc
- Started: 2026-10-22 15:28 UTC
- Resolved: 2026-10-23 02:30 UTC
- Duration: 11h 2m
- Impact: minor, status: resolved
- Components: API Requests, Actions, Issues, Packages, Pull Requests
- Canonical incident: [ci-01510bd2a4cd1e53c09f13f9](#ci-01510bd2a4cd1e53c09f13f9)

- `2026-10-22 15:28` **investigating** — We are investigating reports of degraded performance for Actions, Issues and Pull Requests
- `2026-10-22 15:33` **investigating** — Some users may experience intermittent errors loading Issues, Pull Requests and search results on github.com. <br /><br />We're investigating and working to restore normal service.
- `2026-10-22 16:08` **investigating** — Some users continue to experience intermittent errors loading Issues, Pull Requests and search results on github.com. <br /><br />We're investigating and evaluating mitigation options to restore normal service.
- `2026-10-22 16:26` **investigating** — API Requests is experiencing degraded performance. We are continuing to investigate.
- `2026-10-22 16:28` **investigating** — Packages is experiencing degraded performance. We are continuing to investigate.
- `2026-10-22 16:35` **investigating** — API Requests is operating normally.
- `2026-10-22 16:41` **investigating** — Some users continue to experience errors loading Issues, Pull Requests, package listings and search results on github.com. <br /><br />We're working through multiple mitigation efforts and continuing to investigate.
- `2026-10-22 17:16` **investigating** — We're seeing recovery after adding capacity, with errors decreasing. Some users may still experience errors loading Issues, Pull Requests and search results on github.com. <br /><br />We're continuing to work towards full recovery and monitoring for stability.
- `2026-10-22 17:41` **investigating** — Actions, Issues, Packages and Pull Requests are operating normally.
- `2026-10-22 17:43` **investigating** — We're seeing recovery across the affected services. GitHub Projects is still recovering as queued updates are processed, and we're continuing to monitor for stability.
- `2026-10-22 18:26` **investigating** — The other affected services have recovered. <br /><br />We're continuing to investigate GitHub Projects and verify that normal operation has been restored.
- `2026-10-22 19:04` **investigating** — Some updates in GitHub Projects are still delayed. We're continuing mitigation work and monitoring.
- `2026-10-22 19:56` **investigating** — Some GitHub Projects updates related to issue changes and merged pull requests are still delayed. We're investigating the remaining delays and working to restore normal processing.
- `2026-10-22 20:46` **investigating** — Some GitHub Projects updates, including newly added items and changes to issues and pull requests, remain delayed.<br /><br />We've identified interruptions in update processing and are evaluating mitigations to restore normal operation.
- `2026-10-22 22:07` **investigating** — Some GitHub Projects updates, including newly added items and changes to issues and pull requests, are still delayed.<br /><br />We have seen some recovery, and are working to add more capacity to accelerate the recovery.
- `2026-10-22 23:48` **investigating** — GitHub Projects update processing has improved, but newly added items may still take longer than expected to appear. <br /><br />We're continuing to work on the remaining delays and monitor recovery.
- `2026-10-23 00:55` **monitoring** — The degradation has been mitigated. We are monitoring to ensure stability.
- `2026-10-23 02:30` **resolved** — This incident has been resolved. Thank you for your patience and understanding as we addressed this issue. A detailed root cause analysis will be shared as soon as it is available.

### 6. cloudflare — Network Performance Issues in Ashburn

- URL: https://www.cloudflarestatus.com/incidents/39fz3wkl42tr
- Started: 2026-10-24 01:01 UTC
- Resolved: 2026-10-24 03:09 UTC
- Duration: 2h 8m
- Impact: minor, status: resolved
- Components: Artifacts, Ashburn, United States - (IAD), D1, Durable Objects, Network, Queues, R2, Workflows
- Canonical incident: [ci-e51e85e2fc4bfb8709bdbcbb](#ci-e51e85e2fc4bfb8709bdbcbb)

- `2026-10-24 01:01` **investigating** — Cloudflare is investigating issues with Network Performance in Ashburn, VA. We are working to analyze and mitigate this problem. More updates to follow shortly.
- `2026-10-24 01:05` **investigating** — Affected services: Durable Objects, R2, Workflows, Queues, D1 and Artifacts.
- `2026-10-24 01:28` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-10-24 03:09` **resolved** — This incident has been resolved.

### 7. cloudflare — Delayed audit log delivery for WARP and Zero Trust services

- URL: https://www.cloudflarestatus.com/incidents/qwrtvb4jfvk7
- Started: 2026-10-22 17:32 UTC
- Resolved: 2026-10-23 00:10 UTC
- Duration: 6h 38m
- Impact: major, status: resolved
- Components: Audit Logs, Zero Trust
- Canonical incident: [ci-175bac6dfd66135967c655d3](#ci-175bac6dfd66135967c655d3)

- `2026-10-22 17:32` **investigating** — We are investigating an issue affecting audit log delivery for WARP API. Audit logs for some customers may be delayed or not delivered, and some dependent services may also be affected. We have identified the cause as a capacity problem in a shared internal data pipeline in our PDX region. Mitigations are in place and the pipeline is partially recovered. We are continuing to monitor and work toward full recovery.
- `2026-10-22 17:50` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-10-23 00:10` **resolved** — This incident has been resolved.

### 8. github — Delays for Pull Requests and Webhook deliveries

- URL: https://stspg.io/gnnt0tx9nxmk
- Started: 2026-10-21 14:52 UTC
- Resolved: 2026-10-21 19:22 UTC
- Duration: 4h 30m
- Impact: critical, status: resolved
- Components: Actions, Pull Requests, Webhooks
- Canonical incident: [ci-be8326facd6f44634b5ad3b8](#ci-be8326facd6f44634b5ad3b8)

- `2026-10-21 14:52` **investigating** — We are investigating reports of degraded performance for Actions, Pull Requests and Webhooks
- `2026-10-21 15:15` **investigating** — Actions is experiencing degraded performance. We are continuing to investigate.
- `2026-10-21 15:36` **investigating** — We are investigating delays affecting pull request updates and webhook deliveries. Increased load on our event streaming platform is contributing to hotspots. We are shifting traffic away from affected areas and adding processing capacity; impact is ongoing while these mitigations take effect.
- `2026-10-21 15:58` **investigating** — We continue to see delays affecting pull request updates, webhook deliveries, and audit log events. The event streaming platform is showing signs of recovery, and we are monitoring as the impact continues.
- `2026-10-21 16:17` **investigating** — Audit log events and Actions workflow status updates continue to be delayed. Pull request updates and webhook deliveries are recovering as we restore event processing and work through backlogs. We’re monitoring the recovery and will share further updates as we learn more.
- `2026-10-21 16:25` **investigating** — Audit log events are delayed by up to 100 minutes, and Actions workflow status updates continue to experience extended delays. Pull request updates and webhook deliveries are recovering as we restore event processing and work through backlogs. We’re monitoring recovery and will share further updates as we learn more.
- `2026-10-21 16:33` **investigating** — Pull Requests is operating normally.
- `2026-10-21 16:34` **investigating** — Audit log events continue to be delayed by up to 110 minutes. Pull request updates have recovered. Actions is recovering, and users should no longer experience impact in Actions. Webhook deliveries continue to recover as we work through backlogs. We are monitoring recovery and will share further updates as we learn more.
- `2026-10-21 16:57` **investigating** — Actions and Webhooks are operating normally.
- `2026-10-21 16:59` **investigating** — Actions and webhook deliveries have recovered. Audit log events are still delayed by about 100 minutes. We are working through the backlog, and full recovery will take some time. We will share more updates as recovery continues.
- `2026-10-21 17:09` **investigating** — Pull request updates, Actions, and webhook deliveries have recovered. Audit log events are recovering and delays are decreasing, but we expect the remaining backlog to take approximately two to three hours to clear. We’re monitoring recovery and will share further updates.
- `2026-10-21 18:18` **investigating** — Audit log events are recovering, with delays now below 90 minutes as the remaining backlog is processed. We’ve reduced load on the event-streaming platform and are monitoring recovery. Further updates will follow.
- `2026-10-21 18:52` **investigating** — Audit log events continue to recover, with delays now below 70 minutes. We’ve applied mitigations to speed recovery and expect audit logs to fully recover in about 40 minutes. Pull request updates, Actions, and webhook deliveries remain recovered. We are monitoring progress and will share further updates.
- `2026-10-21 19:22` **investigating** — Pull request updates, Actions, and webhook deliveries have recovered. Audit log delays are now below 40 minutes, within the healthy threshold, and continue to decrease as the remaining backlog is processed.
- `2026-10-21 19:22` **resolved** — This incident has been resolved. Thank you for your patience and understanding as we addressed this issue. A detailed root cause analysis will be shared as soon as it is available.

### 9. grafana — IRM outage in the region prod-gb-south-0 and prod-eu-west-0

- URL: https://stspg.io/nh9hp52f8qzh
- Started: 2026-10-20 07:28 UTC
- Resolved: 2026-10-20 10:33 UTC
- Duration: 3h 5m
- Impact: none, status: resolved
- Components: GCP Belgium - prod-eu-west-0: Alertmanager, GCP Belgium - prod-eu-west-0: Alertmanager and Rules Configuration API, GCP Belgium - prod-eu-west-0: Rule Evaluation, GCP UK - prod-gb-south-0: Alertmanager, GCP UK - prod-gb-south-0: Alertmanager and Rules Configuration API, GCP UK - prod-gb-south-0: Rule Evaluation
- Canonical incident: [ci-f37f0e7288f311b5d6cf90d0](#ci-f37f0e7288f311b5d6cf90d0)

- `2026-10-20 07:28` **investigating** — We are currently investigating this issue.
- `2026-10-20 07:29` **investigating** — We are continuing to investigate this issue.
- `2026-10-20 07:36` **investigating** — Only the public API is affected for prod-eu-west-0
- `2026-10-20 08:04` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-10-20 10:33` **resolved** — This incident has been resolved.

### 10. anthropic — Elevated errors affecting Connectors

- URL: https://stspg.io/4dcr8c005yxc
- Started: 2026-10-22 03:50 UTC
- Resolved: 2026-10-22 08:07 UTC
- Duration: 4h 16m
- Impact: minor, status: resolved
- Components: Claude Cowork, claude.ai
- Canonical incident: [ci-e65647aba2f7e2d8ba0f7393](#ci-e65647aba2f7e2d8ba0f7393)

- `2026-10-22 03:50` **investigating** — We are investigating elevated errors affecting connectors. We will provide an update as soon as possible.
- `2026-10-22 07:28` **monitoring** — A fix has been deployed for the issue affecting claude.ai and Claude Cowork and we are monitoring for recovery.
- `2026-10-22 08:07` **resolved** — The issue affecting claude.ai and Claude Cowork has been resolved.

### 11. cloudflare — Support Helpdesk Availability Issues

- URL: https://www.cloudflarestatus.com/incidents/8m9a1fjjwxp5
- Started: 2026-10-23 03:03 UTC
- Resolved: 2026-10-23 09:42 UTC
- Duration: 6h 39m
- Impact: minor, status: resolved
- Components: Support Site
- Canonical incident: [ci-821ba164c0dd29abf8bab9b9](#ci-821ba164c0dd29abf8bab9b9)

- `2026-10-23 03:03` **investigating** — Our helpdesk provider is currently experiencing issues, and as such customers might encounter errors while submitting or viewing tickets. Responses on customer inquiries may be delayed during this time.
- `2026-10-23 03:37` **identified** — Support portal, live chat support and emergency telephone line are still unavailable due to the ongoing helpdesk provider outage. Customers may experience issues submitting or viewing tickets/chats, and overall response times may be delayed. Customers with urgent issues should revert to using the support email address (backupsupport@cloudflare.com).
- `2026-10-23 04:43` **identified** — We are continuing to work on a fix for the issue.
- `2026-10-23 05:56` **identified** — We are continuing to work on a fix for the issue.
- `2026-10-23 06:37` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-10-23 09:42` **resolved** — This incident has been resolved.

### 12. cloudflare — Cloudflare Zero Trust connectivity

- URL: https://www.cloudflarestatus.com/incidents/0eqtn0khkz87
- Started: 2026-10-20 03:20 UTC
- Resolved: 2026-10-20 06:32 UTC
- Duration: 3h 11m
- Impact: critical, status: resolved
- Components: Dashboard, Zero Trust
- Canonical incident: [ci-bca8bceb3275d7d6111a778a](#ci-bca8bceb3275d7d6111a778a)

- `2026-10-20 03:20` **investigating** — Cloudflare is investigating an issue with traffic that goes through Cloudflare One Gateway when an identity policy is present in the policy decision. This impacts all onramps to Cloudflare One Gateway. Cloudflare One Gateway policies with no identity policy are unaffected. We are also seeing impact to the Cloudflare Dashboard and are investigating further.
- `2026-10-20 03:47` **investigating** — We are continuing to investigate the issue.
- `2026-10-20 04:09` **investigating** — We are continuing to investigate the issue.
- `2026-10-20 04:22` **investigating** — We are continuing to investigate the issue.
- `2026-10-20 04:40` **identified** — The issue has been identified and a fix is being implemented.
- `2026-10-20 04:44` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-10-20 05:46` **monitoring** — The Cloudflare Access API continues to experience higher-than-normal error rates and increased latency. We are working to resolve the remaining API issues. Edge functionality, including Access logins and Cloudflare Secure Web Gateway connectivity, has been fully restored. We continue to monitor these services.
- `2026-10-20 06:14` **monitoring** — The Cloudflare Access API has been fully restored, with error rates and latency returning to normal levels. We continue to monitor these services.
- `2026-10-20 06:32` **resolved** — This incident has been resolved.

### 13. stripe — Elevated Revolut Pay errors

- URL: https://stspg.io/80s4dn3gj9ql
- Started: 2026-10-20 07:09 UTC
- Resolved: 2026-10-20 16:42 UTC
- Duration: 9h 32m
- Impact: none, status: resolved
- Components: Acquirers and payment methods
- Canonical incident: [ci-761acda494454d2fdde259e7](#ci-761acda494454d2fdde259e7)

- `2026-10-20 07:09` **investigating** — We’re investigating elevated errors affecting Revolut Pay, as well as 3D Secure authentication and authorization for some Revolut-issued card transactions. This issue is external to Stripe.
- `2026-10-20 10:32` **monitoring** — We continue to observe elevated authorization errors and payment declines affecting Revolut-issued card transactions and Revolut Pay. We've identified that this issue is caused by an outage at an upstream partner and will continue monitoring for recovery. We will provide further updates once we have observed recovery.
- `2026-10-20 16:42` **resolved** — From 06:00-13:24 UTC, we observed elevated authorization errors and payment declines affecting Revolut-issued card transactions and Revolut Pay. The issue is resolved.

### 14. cloudflare — Elevated error rates for Durable Objects in Western Europe

- URL: https://www.cloudflarestatus.com/incidents/v967mz6rrcxw
- Started: 2026-10-22 17:51 UTC
- Resolved: 2026-10-23 00:09 UTC
- Duration: 6h 17m
- Impact: minor, status: resolved
- Components: Durable Objects
- Canonical incident: [ci-658f2c6102b3854c4f915b42](#ci-658f2c6102b3854c4f915b42)

- `2026-10-22 17:51` **investigating** — We are investigating elevated error rates affecting Durable Objects in Western Europe. Customers may see errors when using Durable Objects and related services. We will post updates as we learn more.
- `2026-10-22 18:19` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-10-23 00:09` **resolved** — This incident has been resolved.

### 15. cloudflare — Network Error Log Processing Unavailable

- URL: https://www.cloudflarestatus.com/incidents/l3d16334tgjp
- Started: 2026-10-22 16:30 UTC
- Resolved: 2026-10-22 22:00 UTC
- Duration: 5h 30m
- Impact: major, status: resolved
- Components: Logs
- Canonical incident: [ci-4c1f2c8a897686ab5707cedd](#ci-4c1f2c8a897686ab5707cedd)

- `2026-10-22 16:30` **investigating** — Since 20:19 UTC, new Network Error Logging (NEL) reports are not being processed. Reports received during this period will not appear in Network Error Logging analytics. Website traffic and other services are not affected. We are investigating and will provide an update as soon as possible.
- `2026-10-22 20:08` **investigating** — Partial service restored. Root cause investigation is currently underway.
- `2026-10-22 21:36` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-10-22 22:00` **resolved** — This incident has been resolved.

### 16. anthropic — Elevated errors for Claude Sonnet 5.5 and other models

- URL: https://stspg.io/s3wd6gmdsdhn
- Started: 2026-10-25 06:14 UTC
- Resolved: 2026-10-25 07:53 UTC
- Duration: 1h 39m
- Impact: major, status: resolved
- Components: Claude API (api.anthropic.com), Claude Code, Claude Cowork, claude.ai
- Canonical incident: [ci-a238d292b9595d18ea969667](#ci-a238d292b9595d18ea969667)

- `2026-10-25 06:14` **investigating** — We are investigating elevated errors on requests to Claude Mythos 5.1, Claude Fable 5.1, Claude Opus 5.5, Claude Sonnet 5.5, and Claude Haiku 5.5. We will provide an update as soon as possible.
- `2026-10-25 06:27` **identified** — We have identified the cause of elevated errors on requests to Claude Sonnet 5.5 and are working on a fix. We will provide an update as soon as possible.
- `2026-10-25 06:35` **monitoring** — A fix has been deployed for the issue affecting Claude services and we are monitoring for recovery.
- `2026-10-25 07:53` **resolved** — The issue affecting Claude services has been resolved. Impact occurred from 23:05 PT on October 24 / 06:05 UTC on October 25 to 23:35 PT on October 24 / 06:35 UTC on October 25.

### 17. openai — Elevated ChatGPT conversation and Responses API errors

- URL: https://status.openai.com/incidents/01M5GPE9C4E19HE461P8JAMYXD
- Started: 2026-10-22 02:18 UTC
- Resolved: 2026-10-22 08:50 UTC
- Duration: 6h 32m
- Impact: minor, status: resolved
- Components: —
- Canonical incident: [ci-38209ec97e89bf26924b7ba9](#ci-38209ec97e89bf26924b7ba9)

- `2026-10-22 02:18` **investigating** — This is impacting users in India, Japan, and Korea. We are investigating the issue for the listed services.
- `2026-10-22 02:42` **monitoring** — This is impacting users in India, Japan, and Korea. We have applied the mitigation and are monitoring the recovery.
- `2026-10-22 08:50` **resolved** — All impacted services have now fully recovered.

### 18. cloudflare — Delayed, missed, and duplicated Workers Cron Triggers

- URL: https://www.cloudflarestatus.com/incidents/3ts1mfzs0khz
- Started: 2026-10-23 15:27 UTC
- Resolved: 2026-10-23 20:37 UTC
- Duration: 5h 9m
- Impact: minor, status: resolved
- Components: Workers
- Canonical incident: [ci-fe367c9ee656ab72e17d18c3](#ci-fe367c9ee656ab72e17d18c3)

- `2026-10-23 15:27` **investigating** — We are investigating reports that some Workers Cron Triggers are running late, being skipped, or running more than once. Affected customers may see scheduled Workers invocations that do not run on time or run more than once. We will post an update as we learn more.
- `2026-10-23 20:29` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-10-23 20:37` **resolved** — This incident has been resolved.

### 19. openai — Elevated FedRAMP API error rates for Astra and Sol

- URL: https://status.openai.com/incidents/01M5FD63HHDRPQPM5YMQWJ5HTN
- Started: 2026-10-21 14:17 UTC
- Resolved: 2026-10-21 19:34 UTC
- Duration: 5h 17m
- Impact: minor, status: resolved
- Components: —
- Canonical incident: [ci-09225cc4401250e1b3349958](#ci-09225cc4401250e1b3349958)

- `2026-10-21 14:17` **identified** — We have identified that users are experiencing elevated errors for the impacted services. We are working on implementing a mitigation.
- `2026-10-21 17:03` **investigating** — We are investigating the issue for the listed services.
- `2026-10-21 17:04` **monitoring** — We have applied the mitigation and are monitoring the recovery.
- `2026-10-21 19:34` **resolved** — All impacted services have now fully recovered.

### 20. grafana — DB issues in prod-sa-east-1

- URL: https://stspg.io/28py1pvhnpjf
- Started: 2026-10-20 21:36 UTC
- Resolved: 2026-10-21 00:09 UTC
- Duration: 2h 33m
- Impact: none, status: resolved
- Components: AWS Brazil - prod-sa-east-1, AWS Brazil - prod-sa-east-1: API, AWS Brazil - prod-sa-east-1: Public Probes
- Canonical incident: [ci-5e97f10fd3ad8b465e170dcc](#ci-5e97f10fd3ad8b465e170dcc)

- `2026-10-20 21:36` **investigating** — We are currently investigating an issue in the prod-sa-east-1 region causing slow requests, or some stacks struggling to start. We will provide an update as soon as we have more information to share.
- `2026-10-20 22:14` **investigating** — Synthetic Monitoring alerting is also impacted. We are continuing to investigate this issue.
- `2026-10-20 23:39` **investigating** — Synthetic Monitoring alerting has recovered. We are continuing to monitor.
- `2026-10-21 00:09` **resolved** — This incident has been resolved.

### 21. cloudflare — Elevated Errors for R2 and Durable Objects in ENAM

- URL: https://www.cloudflarestatus.com/incidents/p1vurmibvmag
- Started: 2026-10-21 19:15 UTC
- Resolved: 2026-10-21 21:21 UTC
- Duration: 2h 6m
- Impact: major, status: resolved
- Components: Durable Objects, R2
- Canonical incident: [ci-fe8781ad2a5d93d98393abcd](#ci-fe8781ad2a5d93d98393abcd)

- `2026-10-21 20:43` **resolved** — Cloudflare has resolved an issue affecting R2 and Durable Objects in the Eastern North America region. Some customers may have seen elevated errors between approximately 19:02 and 19:05 UTC on October 21, 2026.
- `2026-10-21 21:21` **resolved** — This incident has been resolved.

### 22. openai — New Ad Accounts failing to provision admins

- URL: https://status.openai.com/incidents/01M5JHVJVC8WAXB4WRYFMKG6HP
- Started: 2026-10-22 19:36 UTC
- Resolved: 2026-10-23 00:07 UTC
- Duration: 4h 31m
- Impact: minor, status: resolved
- Components: —
- Canonical incident: [ci-a31b0ca45af63b2641680c14](#ci-a31b0ca45af63b2641680c14)

- `2026-10-22 19:36` **identified** — We have identified that new ad accounts are failing to assign their creator admin. We are investigating the issue.
- `2026-10-22 21:30` **identified** — We have identified the issue and are working to implement the mitigation.
- `2026-10-22 23:45` **monitoring** — We have applied the mitigation and are monitoring the recovery.
- `2026-10-23 00:07` **resolved** — All impacted services have now fully recovered.

### 23. cloudflare — Increased Latency in India

- URL: https://www.cloudflarestatus.com/incidents/t77gwnpr6oe5
- Started: 2026-10-21 16:42 UTC
- Resolved: 2026-10-21 19:57 UTC
- Duration: 3h 14m
- Impact: minor, status: resolved
- Components: CDN/Cache
- Canonical incident: [ci-06167ecad8107e08f15af6ce](#ci-06167ecad8107e08f15af6ce)

- `2026-10-21 16:42` **investigating** — We are investigating reports of increased latency affecting some traffic in India. Affected users may experience slower loading times. Our engineering teams are investigating, and we will provide an update as more information becomes available.
- `2026-10-21 18:51` **monitoring** — We have identified and implemented a fix for this issue and are currently monitoring the results.
- `2026-10-21 19:57` **resolved** — This incident has been resolved.

### 24. github — Disruption with Dependabot Alerts

- URL: https://stspg.io/zdk3zqt69fnz
- Started: 2026-10-23 14:58 UTC
- Resolved: 2026-10-23 18:51 UTC
- Duration: 3h 52m
- Impact: minor, status: resolved
- Components: —
- Canonical incident: [ci-a58e3babedc502dcbbb1e2d8](#ci-a58e3babedc502dcbbb1e2d8)

- `2026-10-23 14:58` **investigating** — We are investigating reports of impacted performance for some GitHub services.
- `2026-10-23 15:00` **investigating** — We’re investigating delays in some Dependabot alerts closing after affected repositories have been updated to address a vulnerability. Alert processing is backed up. We’ll share more information as it becomes available.
- `2026-10-23 15:39` **investigating** — We’re scaling alert processing capacity to reduce the backlog causing delays in some Dependabot alerts closing after affected repositories have been updated. We’re continuing to monitor and will share another update as more information becomes available.
- `2026-10-23 16:11` **investigating** — We’re continuing to investigate delays in some Dependabot alerts closing after affected repositories are updated. Alert processing remains backed up, and additional processing capacity is rolling out to help reduce the backlog. We’re monitoring the rollout and will share another update as we learn more.
- `2026-10-23 16:43` **investigating** — We’re continuing to investigate delays in some Dependabot alerts closing after affected repositories are updated. Alert processing remains backed up, and additional processing capacity is rolling out to help reduce the backlog. We’re monitoring progress and will share another update as we learn more.
- `2026-10-23 17:27` **investigating** — Some Dependabot alerts continue to be delayed in closing after repository updates address a vulnerability. <br /><br />We're continuing to investigate and work to restore normal processing.
- `2026-10-23 18:36` **monitoring** — The degradation has been mitigated. We are monitoring to ensure stability.
- `2026-10-23 18:51` **resolved** — We identified that an advisory-publication driven fan-out was overwhelming Dependabot alert processing on repositories by saturating the queue with its own events, delaying processing of events related to repository pushes.<br /><br />We have now isolated this fan-out and applied a throttle so it runs at a fixed rate and cannot delay other events. This remediation has successfully prevented a new fan-out process impacting users.

### 25. cloudflare — Network connectivity issues affecting services in North America

- URL: https://www.cloudflarestatus.com/incidents/sufgcbe8rrwt
- Started: 2026-10-23 15:47 UTC
- Resolved: 2026-10-23 18:43 UTC
- Duration: 2h 56m
- Impact: major, status: resolved
- Components: Network
- Canonical incident: [ci-6e1b833564feb6974890f8ef](#ci-6e1b833564feb6974890f8ef)

- `2026-10-23 15:47` **investigating** — We are investigating degraded network performance affecting some traffic between our Ashburn and Atlanta data centers. Some customers may see increased latency, packet loss, or intermittent errors on Workers and other services in the Eastern US. Our network team is actively working on mitigation.
- `2026-10-23 16:07` **identified** — We have identified a connectivity provider outage causing network congestion and degraded performance for some customers in North America. Affected customers may experience increased latency, packet loss, and intermittent errors across multiple services. Our network team is rerouting traffic away from affected paths, and we are working directly with our provider on restoration.
- `2026-10-23 16:27` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-10-23 16:59` **identified** — We are continuing to work on a fix for the issue.
- `2026-10-23 17:03` **identified** — The infrastructure provider has experienced a recurrence of the failure. Cloudflare is continuing to apply traffic management actions to further remediate the issue.
- `2026-10-23 17:33` **identified** — The team is continuing to perform traffic engineering actions to route around impacted paths. The team will continue performing traffic engineering actions until all packet drops are mitigated.
- `2026-10-23 18:18` **monitoring** — Our network team has rerouted traffic around the impacted paths and we are no longer seeing packet drops. Our connectivity provider reports that its service has been restored, but some links are still being brought back, so we are keeping our traffic management changes in place.
- `2026-10-23 18:43` **resolved** — This incident has been resolved.

### 26. github — Degraded availability for GitHub Actions, GitHub Pages and Webhooks

- URL: https://stspg.io/nrv2k8s68bwn
- Started: 2026-10-19 15:49 UTC
- Resolved: 2026-10-19 16:56 UTC
- Duration: 1h 7m
- Impact: critical, status: resolved
- Components: Actions, Pages, Webhooks
- Canonical incident: [ci-8abff99f6c7e7ecd05f25023](#ci-8abff99f6c7e7ecd05f25023)

- `2026-10-19 15:49` **investigating** — We are investigating reports of degraded performance for Actions and Pages
- `2026-10-19 15:50` **investigating** — Webhooks is experiencing degraded performance. We are continuing to investigate.
- `2026-10-19 15:56` **investigating** — We’re are currently investigating degraded availability affecting GitHub Actions, GitHub Pages and Webhooks. <br /><br />Customers may experience delays in workflow processing, site builds and webhook delivery.<br />
- `2026-10-19 16:30` **investigating** — We’re seeing recovery following mitigation and are continuing to monitor. <br /><br />Customers may still experience delays in GitHub Actions workflow status updates, GitHub Pages deployments, and check and workflow webhook events generated by Actions runs.
- `2026-10-19 16:56` **monitoring** — Actions, Pages and Webhooks are operating normally.
- `2026-10-19 16:56` **resolved** — On October 19, 2026, between 15:24 and 16:52 UTC, GitHub Actions, GitHub Pages, and webhooks generated by Actions were degraded. Live updates on GitHub.com and code search were also degraded for part of this time. No data was lost. <br /><br />Impact <br /><br /> - Workflow status updates, logs, and checks were delayed. <br /> - 5.13% of jobs on GitHub-hosted runners were delayed in starting, and 0.02% of workflow runs failed. <br /> - During the first 30 minutes, 8.47% of requests to the Actions jobs REST API returned errors. <br /><br />Pages deployments and webhooks generated by Actions runs were delayed. <br /><br />An internal Actions service that stores logs and run status lost connections to some of its cache servers in one cloud region and immediately retried them from every server at once. The resulting surge of new connections overloaded a shared network link between our data center and that region, which also disrupted other services that use the same link. We mitigated the incident by moving the Actions service off the affected cache servers and reducing connections from other services on the link. We then processed the delayed Actions work. We are continuing to investigate the cause of the initial connection timeouts. <br /><br />We were already in the process of rolling out a limit on how quickly this service can open new connections, but it was still in shadow mode (logging only, not enforcing) at the time of the incident. We have since fully enforced it and capped the service's concurrent connection attempts. We are extending these limits to other Actions services, adding backoff to retries, limiting how much shared network capacity a single service can use, and improving alerting so we can find the source of connection surges faster.

### 27. stripe — Elevated Bancontact errors

- URL: https://stspg.io/k2jfv0fg9ff6
- Started: 2026-10-22 22:20 UTC
- Resolved: 2026-10-23 01:56 UTC
- Duration: 3h 35m
- Impact: none, status: resolved
- Components: Acquirers and payment methods
- Canonical incident: [ci-9bd5114735766713a9c58d1f](#ci-9bd5114735766713a9c58d1f)

- `2026-10-22 22:20` **identified** — We're currently investigating a partial service disruption of the Bancontact payment method that began at 21:03 UTC. This issue is external to Stripe and is causing increased error rates for transactions processed through the Bancontact payment method. We are actively monitoring the situation and will provide updates as more information becomes available.
- `2026-10-23 01:56` **resolved** — From 21:03 to 23:46 UTC on October 22, there was a partial service disruption of the Bancontact payment method. Some users may have experienced increased error rates for transactions processed through the Bancontact payment method. This issue was external to Stripe and is now resolved.

### 28. anthropic — Elevated errors for Claude Fable 5 and Claude Fable 5.1

- URL: https://stspg.io/7n1wgb6qsfly
- Started: 2026-10-24 05:16 UTC
- Resolved: 2026-10-24 05:53 UTC
- Duration: 36m
- Impact: major, status: resolved
- Components: Claude API (api.anthropic.com), Claude Code, Claude Cowork, claude.ai
- Canonical incident: [ci-e77ad4a5540c2ad86a0c1309](#ci-e77ad4a5540c2ad86a0c1309)

- `2026-10-24 05:16` **investigating** — We are investigating elevated errors on requests to other Claude models. We will provide an update as soon as possible.
- `2026-10-24 05:24` **identified** — We have identified the cause of elevated errors on requests to other Claude models and are working on a fix. We will provide an update as soon as possible.
- `2026-10-24 05:40` **identified** — We have identified the cause of elevated errors on requests to other Claude models and are working on a fix. We will provide an update as soon as possible.
- `2026-10-24 05:41` **monitoring** — A fix has been deployed for the issue affecting Claude services and we are monitoring for recovery.
- `2026-10-24 05:53` **resolved** — The issue affecting Claude services has been resolved. Impact occurred from 22:00 PT on October 23 / 05:00 UTC on October 24 to 22:35 PT on October 23 / 05:35 UTC on October 24.

### 29. openai — Delays in customer support responses

- URL: https://status.openai.com/incidents/01M5KDJ4079XQTZDSRX40NMN6H
- Started: 2026-10-23 03:41 UTC
- Resolved: 2026-10-23 05:58 UTC
- Duration: 2h 17m
- Impact: none, status: resolved
- Components: —
- Canonical incident: [ci-81b99475136688f5534f1fed](#ci-81b99475136688f5534f1fed)

- `2026-10-23 03:41` **investigating** — We are investigating technical issues affecting our customer support systems. Customers may experience longer than usual response times. We are working to restore normal operations.
- `2026-10-23 03:57` **investigating** — As a temporary solution, customers looking to reach out to OpenAI Support can email [support@openai.com](mailto:support@openai.com "support@openai.com").
- `2026-10-23 04:43` **identified** — A fix is being implemented for the issue affecting our support systems. Customers may continue to experience delays in receiving support responses. We will provide further updates as recovery progresses.
- `2026-10-23 05:47` **monitoring** — We have applied the mitigation and are monitoring the recovery.
- `2026-10-23 05:58` **resolved** — All impacted services have now fully recovered.

### 30. anthropic — Elevated errors for Claude Fable 5.1, Sonnet 5.5, and Haiku 5.5

- URL: https://stspg.io/gqp288t3qzpt
- Started: 2026-10-23 23:37 UTC
- Resolved: 2026-10-24 00:11 UTC
- Duration: 33m
- Impact: critical, status: resolved
- Components: Claude API (api.anthropic.com), Claude Code, Claude Cowork, claude.ai
- Canonical incident: [ci-a2cd10e989fc6f692657b7b2](#ci-a2cd10e989fc6f692657b7b2)

- `2026-10-23 23:37` **investigating** — We are investigating elevated errors on requests to Claude Fable 5.1. We will provide an update as soon as possible.
- `2026-10-23 23:39` **identified** — We have identified the cause of elevated errors on requests to Claude Fable 5.1, Sonnet 5.5, and Haiku 5.5, as well as fast mode requests to Opus 5.5 and are working on a fix. We will provide an update as soon as possible.
- `2026-10-23 23:54` **monitoring** — We are seeing success rates return to normal for all affected models, and are continuing to monitor to ensure no further issues.
- `2026-10-24 00:11` **resolved** — The issue affecting Claude models has been resolved. Impact occurred from 16:28 PT / 23:28 UTC on October 23 to 16:59 PT / 23:59 UTC on October 23.

### 31. anthropic — Elevated errors for Claude Fable 5.1

- URL: https://stspg.io/y0gkl12ztjst
- Started: 2026-10-25 08:17 UTC
- Resolved: 2026-10-25 08:49 UTC
- Duration: 32m
- Impact: major, status: resolved
- Components: Claude API (api.anthropic.com), Claude Code, Claude Cowork, claude.ai
- Canonical incident: [ci-2b70e77251e9dbd7c301bf32](#ci-2b70e77251e9dbd7c301bf32)

- `2026-10-25 08:17` **investigating** — We are investigating elevated errors on requests to Claude Fable 5.1. We will provide an update as soon as possible.
- `2026-10-25 08:17` **identified** — We have identified the cause of elevated errors on requests to Claude Fable 5.1 and are working on a fix. We will provide an update as soon as possible.
- `2026-10-25 08:30` **monitoring** — A fix has been deployed for the issue affecting Claude services and we are monitoring for recovery.
- `2026-10-25 08:49` **resolved** — The issue affecting Claude services has been resolved. Impact occurred from 01:08 PT / 08:08 UTC to 01:29 PT / 08:29 UTC.

### 32. openai — Degraded Performance of Dots on Enterprise

- URL: https://status.openai.com/incidents/01M5JF8DR3EGQJNTXPZ24AKX2G
- Started: 2026-10-22 18:51 UTC
- Resolved: 2026-10-22 20:55 UTC
- Duration: 2h 4m
- Impact: minor, status: resolved
- Components: —
- Canonical incident: [ci-4c4b0b5bccc5651f9fa80b23](#ci-4c4b0b5bccc5651f9fa80b23)

- `2026-10-22 18:51` **investigating** — Users are experiencing failures sending messages to dots.
- `2026-10-22 18:59` **identified** — Our teams have identified the issue and are actively implementing a mitigation. At this time the majority of impact is affecting enterprise customers.
- `2026-10-22 20:55` **resolved** — All impacted services have now fully recovered.

### 33. stripe — Elevated Bancontact errors

- URL: https://stspg.io/ncmhd5df6f3k
- Started: 2026-10-22 14:18 UTC
- Resolved: 2026-10-22 16:44 UTC
- Duration: 2h 25m
- Impact: none, status: resolved
- Components: Acquirers and payment methods
- Canonical incident: [ci-f77218a758067ac55f5709af](#ci-f77218a758067ac55f5709af)

- `2026-10-22 14:18` **identified** — We're currently investigating a Partial Service disruption of the Bancontact payment method that began at 14:08 UTC. This issue is external to Stripe and is causing increased error rates for transactions processed through the Bancontact payment method. We are actively monitoring the situation and will provide updates as more information becomes available.
- `2026-10-22 16:44` **resolved** — From 14:08 - 16:26 UTC there was a partial service disruption of the Bancontact payment method. Users may have experienced increased error rates for transactions processed through the Bancontact payment method. This issue was external to Stripe and is now resolved.

### 34. grafana — Metrics: Evelvated Read and Ruler Errors.

- URL: https://stspg.io/t7cdk3q6bcz9
- Started: 2026-10-19 15:06 UTC
- Resolved: 2026-10-19 17:56 UTC
- Duration: 2h 49m
- Impact: major, status: resolved
- Components: GCP US Central - prod-us-central-0: Querying
- Canonical incident: [ci-55d6f7311d52643fa1267d42](#ci-55d6f7311d52643fa1267d42)

- `2026-10-19 15:06` **investigating** — We are currently investigating an issue causing elevated read and ruler errors. We will provide another update as soon as we have more to share.
- `2026-10-19 15:48` **monitoring** — Errors have stopped since 15:15 UTC. We are continuing to monitor.
- `2026-10-19 17:56` **resolved** — This incident has been resolved.

### 35. stripe — Spike in cancellations and expired errors with PayByBank payments

- URL: https://stspg.io/gmpxzbw40g3d
- Started: 2026-10-20 10:46 UTC
- Resolved: 2026-10-20 13:04 UTC
- Duration: 2h 17m
- Impact: none, status: resolved
- Components: Acquirers and payment methods
- Canonical incident: [ci-02eab138d0f5a589a622c3d4](#ci-02eab138d0f5a589a622c3d4)

- `2026-10-20 10:46` **identified** — We're currently seeing elevated cancellations and expired errors with PayByBank payments. We've identified that this is caused by an outage at an upstream partner and our engineers will monitor for recovery. We will provide further updates once recovery is observed.
- `2026-10-20 13:04` **resolved** — From 2026-10-20 10:47 UTC to approximately 11:18 UTC, we saw elevated cancellation and expired payment errors affecting PayByBank payments through a third-party payment partner. The issue was limited to payments processed through a specific issuing bank and has now recovered.

### 36. circleci — Increased wait times for macOS jobs on m4pro.large

- URL: https://stspg.io/pp6wqy3vlms3
- Started: 2026-10-20 15:58 UTC
- Resolved: 2026-10-20 18:14 UTC
- Duration: 2h 15m
- Impact: minor, status: resolved
- Components: macOS Jobs
- Canonical incident: [ci-0cdfa8a494248e3c6140c18c](#ci-0cdfa8a494248e3c6140c18c)

- `2026-10-20 15:58` **investigating** — We're seeing increased wait times for some macOS jobs using the m4pro.large resource class. We're investigating and will provide another update soon. Thank you for your patience.
- `2026-10-20 16:32` **monitoring** — Customers running macOS jobs on the m4pro.large resource class are still seeing increased wait times, with some jobs waiting up to 30 minutes to start. We're continuing to monitor this and will share another update shortly.
- `2026-10-20 17:02` **monitoring** — Wait times for m4pro.large are still elevated due to heavy large demand. All m4pro.medium are normal. We're actively monitoring, we appreciate patience.
- `2026-10-20 17:31` **monitoring** — Wait times for macOS jobs on the m4pro.large resource class are recovering. Jobs using Xcode 27.0 are now waiting about ~3.5 minutes on average, down from about ~40 minutes at the peak. Jobs using Xcode 26.6 and 26.4.1 are still working through a backlog, which is clearing. Wait times on m4pro.medium are normal. We will share another update soon as we continue to monitor.
- `2026-10-20 18:14` **resolved** — Wait times for macOS jobs on the m4pro.large resource class have returned to normal. We thank you for your patience while the backlogs cleared.

### 37. grafana — Some GCK6 App Project Pages are Failing to Load

- URL: https://stspg.io/wh6fr7yf9j5h
- Started: 2026-10-21 13:22 UTC
- Resolved: 2026-10-21 15:29 UTC
- Duration: 2h 7m
- Impact: major, status: resolved
- Components: GCk6 App
- Canonical incident: [ci-ff88dcdc38bab1e99967a161](#ci-ff88dcdc38bab1e99967a161)

- `2026-10-21 13:22` **identified** — We are currently investigating this issue impacting some GCK6 app project pages. We are in the process of providing a fix for this.
- `2026-10-21 14:22` **identified** — We are continuing to work on a fix for the issue.
- `2026-10-21 15:29` **resolved** — This incident has been resolved.

### 38. grafana — Some GCK6 App Project Pages are Failing to Load

- URL: https://stspg.io/y9w2szsh123d
- Started: 2026-10-20 14:36 UTC
- Resolved: 2026-10-20 16:08 UTC
- Duration: 1h 31m
- Impact: major, status: resolved
- Components: GCk6 App
- Canonical incident: [ci-d9cc4bc829d37ed93096a4a9](#ci-d9cc4bc829d37ed93096a4a9)

- `2026-10-20 14:36` **investigating** — We are currently investigating this issue impacting some GCK6 app project pages. We will provide an update once we have more information.
- `2026-10-20 16:08` **resolved** — This incident has been resolved.

### 39. grafana — Elevated Trace Ingestion Latency and Errors for Grafana Cloud Traces in AWS Germany

- URL: https://stspg.io/p5hhk4h6k6dm
- Started: 2026-10-23 09:31 UTC
- Resolved: 2026-10-23 10:16 UTC
- Duration: 44m
- Impact: major, status: resolved
- Components: AWS Germany - prod-eu-west-2, AWS Germany - prod-eu-west-2: Ingestion
- Canonical incident: [ci-2deb47fcafe3e70c16fd7f36](#ci-2deb47fcafe3e70c16fd7f36)

- `2026-10-23 09:31` **investigating** — We are investigating elevated latency and intermittent errors when ingesting traces for Grafana Cloud Traces in AWS Germany. Customers may see delayed trace ingestion and failed trace writes. Next update by 10:30 UTC.
- `2026-10-23 10:16` **resolved** — Between 09:31 and 10:13 UTC on 23 October 2026 (42 minutes), Grafana Cloud Traces in AWS Germany experienced elevated ingestion latency and intermittent trace write failures, affecting both direct trace ingestion and ingestion via the OpenTelemetry endpoint. All services are now operating normally.

### 40. anthropic — Elevated errors for Claude Opus 5.5

- URL: https://stspg.io/dmn3bc440bnb
- Started: 2026-10-24 00:31 UTC
- Resolved: 2026-10-24 00:37 UTC
- Duration: 5m
- Impact: major, status: resolved
- Components: Claude API (api.anthropic.com), Claude Code, Claude Cowork, claude.ai
- Canonical incident: [ci-bf4547b0669eacc6e8a6631a](#ci-bf4547b0669eacc6e8a6631a)

- `2026-10-24 00:31` **investigating** — We are investigating elevated errors on requests to Claude Opus 5.5. We will provide an update as soon as possible.
- `2026-10-24 00:37` **resolved** — The issue affecting Claude Opus 5.5 has been resolved. Impact occurred from 17:23 PT on October 23 / 00:23 UTC on October 24 to 17:34 PT on October 23 / 00:34 UTC on October 24.

### 41. openai — Authentication errors for some users

- URL: https://status.openai.com/incidents/01M5G9QQ9V4CYDWFRYF9BWZ5GG
- Started: 2026-10-21 22:36 UTC
- Resolved: 2026-10-20 17:35 UTC
- Duration: unknown
- Impact: major, status: resolved
- Components: —
- Canonical incident: [ci-2c32255340a752854a0604f5](#ci-2c32255340a752854a0604f5)

- `2026-10-21 22:36` **resolved** — On October 20, from 9:35 to 10:35 a.m. PDT, some users near Illinois experienced authentication errors with the API, ChatGPT, and Codex. We rolled back a configuration change, and service was restored.
- `2026-10-21 22:36` **investigating** — Some users near Illinois experienced authentication errors with the API, ChatGPT, and Codex.

## Below threshold (22)

- `cloudflare` [Network congestion in Portland and Seattle, US](https://www.cloudflarestatus.com/incidents/nwgwgbrn9a71) — 1h 46m
- `anthropic` [Degraded performance for Claude Fable 5](https://stspg.io/cgklr4v2sgkh) — 40m
- `cloudflare` [Elevated Error Rate for Durable Objects in Lisbon](https://www.cloudflarestatus.com/incidents/33k6g8j7mhah) — 44m
- `cloudflare` [Elevated errors for some requests in Ashburn, Virginia](https://www.cloudflarestatus.com/incidents/l0gsc0sccx2v) — 29m
- `openai` [Elevated error rates for gpt-5.6-luna API requests](https://status.openai.com/incidents/01M5FHKHPQJMSHCE2B2WDPZW2N) — 1h 9m
- `grafana` [Increased timeout errors in alert rule evaluations (AWS Germany, prod-eu-west-2)](https://stspg.io/xk032xv99zbr) — 1h 38m
- `openai` [Degraded performance in ChatGPT, Codex, and APIs](https://status.openai.com/incidents/01M5E4G39TD1BW1QWHZG6RFA5H) — 1h 1m
- `github` [Incident with Actions](https://stspg.io/44w94j036yc8) — 1h 0m
- `circleci` [Increased Docker gen2 wait times](https://stspg.io/8g54j07z69lb) — 1h 23m
- `cloudflare` [Increased error rate in Workflow Operations](https://www.cloudflarestatus.com/incidents/pij2josz3rvz) — 44m
- `cloudflare` [Cloudflare Workers API Latency](https://www.cloudflarestatus.com/incidents/lhdg2tq46faq) — 42m
- `openai` [Degraded experience with ChatGPT app connections using CIMD](https://status.openai.com/incidents/01M5NFF94XCJ6CGTD4S78A9BEB) — 50m
- `cloudflare` [Issue with Bot management traffic categorization for verified bots](https://www.cloudflarestatus.com/incidents/qlgdup1n0ge7) — 41m
- `stripe` [Elevated Blik errors](https://stspg.io/w2qqtf6tbrsy) — 1h 0m
- `openai` [File Download errors](https://status.openai.com/incidents/01M5A5S1ATN3QB2WCHQ3VDZC81) — 42m
- `github` [Grok 4.6 and Grok 4.7 Service Issues](https://stspg.io/9hysn9g59lpg) — 41m
- `github` [Delayed Actions status check results](https://stspg.io/kmk42jg3v1sm) — 36m
- `openai` [Help Center Chat Support Degraded](https://status.openai.com/incidents/01M5NAR5GZC5M7D0S7AQ6CGEEN) — 32m
- `cloudflare` [DNS Update Delays](https://www.cloudflarestatus.com/incidents/ce61s9m7vz48) — 22m
- `cloudflare` [Cloudflare Access Availability Issues](https://www.cloudflarestatus.com/incidents/t0p23w62n13j) — 20m
- `grafana` [Possible incorrect evaluation of Synthetic Monitoring alert rules in cluster prod-us-east-3](https://stspg.io/3xcm1c985jpl) — 0m
- `slack` [Slack Real-Time Connection Disruption](https://slack-status.com/2026-10/58a30c7fc8aa3e80) — 0m

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-b2dbd754a165046e8a4caebe"></a>

- Canonical ID: `ci-b2dbd754a165046e8a4caebe` (1 source records).
  - [gcp/Xs2XTHzAVC9YnafnPyzj](../../data/2026/W43/gcp.json); vendor URL: https://status.cloud.google.com/incidents/Xs2XTHzAVC9YnafnPyzj
- Start time basis: `vendor_reported`.

<a id="ci-a154ea27c29c7b2ae08dc805"></a>

- Canonical ID: `ci-a154ea27c29c7b2ae08dc805` (1 source records).
  - [grafana/mthmxgzkn610](../../data/2026/W43/grafana.json); vendor URL: https://stspg.io/v6j9g9gqpwsw
- Start time basis: `vendor_reported`.

<a id="ci-261da6711cd4b1fbba5acb7e"></a>

- Canonical ID: `ci-261da6711cd4b1fbba5acb7e` (1 source records).
  - [grafana/4z62z17qh88l](../../data/2026/W43/grafana.json); vendor URL: https://stspg.io/blcsc38zxjxw
- Start time basis: `vendor_reported`.

<a id="ci-a17ba2cdc8c4ab17a0eee508"></a>

- Canonical ID: `ci-a17ba2cdc8c4ab17a0eee508` (1 source records).
  - [cloudflare/3z8y23zn7tr2](../../data/2026/W43/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/3z8y23zn7tr2
- Start time basis: `vendor_reported`.

<a id="ci-01510bd2a4cd1e53c09f13f9"></a>

- Canonical ID: `ci-01510bd2a4cd1e53c09f13f9` (1 source records).
  - [github/wq8305lppdc2](../../data/2026/W43/github.json); vendor URL: https://stspg.io/sx51v1tlpqjc
- Start time basis: `vendor_reported`.

<a id="ci-e51e85e2fc4bfb8709bdbcbb"></a>

- Canonical ID: `ci-e51e85e2fc4bfb8709bdbcbb` (1 source records).
  - [cloudflare/39fz3wkl42tr](../../data/2026/W43/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/39fz3wkl42tr
- Start time basis: `vendor_reported`.

<a id="ci-175bac6dfd66135967c655d3"></a>

- Canonical ID: `ci-175bac6dfd66135967c655d3` (1 source records).
  - [cloudflare/qwrtvb4jfvk7](../../data/2026/W43/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/qwrtvb4jfvk7
- Start time basis: `vendor_reported`.

<a id="ci-be8326facd6f44634b5ad3b8"></a>

- Canonical ID: `ci-be8326facd6f44634b5ad3b8` (1 source records).
  - [github/nms1dpzmh9d9](../../data/2026/W43/github.json); vendor URL: https://stspg.io/gnnt0tx9nxmk
- Start time basis: `vendor_reported`.

<a id="ci-f37f0e7288f311b5d6cf90d0"></a>

- Canonical ID: `ci-f37f0e7288f311b5d6cf90d0` (1 source records).
  - [grafana/cwcjkcghg3g0](../../data/2026/W43/grafana.json); vendor URL: https://stspg.io/nh9hp52f8qzh
- Start time basis: `vendor_reported`.

<a id="ci-e65647aba2f7e2d8ba0f7393"></a>

- Canonical ID: `ci-e65647aba2f7e2d8ba0f7393` (1 source records).
  - [anthropic/wj6r4t6mxq1r](../../data/2026/W43/anthropic.json); vendor URL: https://stspg.io/4dcr8c005yxc
- Start time basis: `vendor_reported`.

<a id="ci-821ba164c0dd29abf8bab9b9"></a>

- Canonical ID: `ci-821ba164c0dd29abf8bab9b9` (1 source records).
  - [cloudflare/8m9a1fjjwxp5](../../data/2026/W43/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/8m9a1fjjwxp5
- Start time basis: `vendor_reported`.

<a id="ci-bca8bceb3275d7d6111a778a"></a>

- Canonical ID: `ci-bca8bceb3275d7d6111a778a` (1 source records).
  - [cloudflare/0eqtn0khkz87](../../data/2026/W43/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/0eqtn0khkz87
- Start time basis: `vendor_reported`.

<a id="ci-761acda494454d2fdde259e7"></a>

- Canonical ID: `ci-761acda494454d2fdde259e7` (1 source records).
  - [stripe/kjn6f9xs515n](../../data/2026/W43/stripe.json); vendor URL: https://stspg.io/80s4dn3gj9ql
- Start time basis: `vendor_reported`.

<a id="ci-658f2c6102b3854c4f915b42"></a>

- Canonical ID: `ci-658f2c6102b3854c4f915b42` (1 source records).
  - [cloudflare/v967mz6rrcxw](../../data/2026/W43/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/v967mz6rrcxw
- Start time basis: `vendor_reported`.

<a id="ci-4c1f2c8a897686ab5707cedd"></a>

- Canonical ID: `ci-4c1f2c8a897686ab5707cedd` (1 source records).
  - [cloudflare/l3d16334tgjp](../../data/2026/W43/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/l3d16334tgjp
- Start time basis: `vendor_reported`.

<a id="ci-a238d292b9595d18ea969667"></a>

- Canonical ID: `ci-a238d292b9595d18ea969667` (1 source records).
  - [anthropic/334jvgbrc4ff](../../data/2026/W43/anthropic.json); vendor URL: https://stspg.io/s3wd6gmdsdhn
- Start time basis: `vendor_reported`.

<a id="ci-38209ec97e89bf26924b7ba9"></a>

- Canonical ID: `ci-38209ec97e89bf26924b7ba9` (1 source records).
  - [openai/01M5GPE9C4E19HE461P8JAMYXD](../../data/2026/W43/openai.json); vendor URL: https://status.openai.com/incidents/01M5GPE9C4E19HE461P8JAMYXD
- Start time basis: `vendor_reported`.

<a id="ci-fe367c9ee656ab72e17d18c3"></a>

- Canonical ID: `ci-fe367c9ee656ab72e17d18c3` (1 source records).
  - [cloudflare/3ts1mfzs0khz](../../data/2026/W43/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/3ts1mfzs0khz
- Start time basis: `vendor_reported`.

<a id="ci-09225cc4401250e1b3349958"></a>

- Canonical ID: `ci-09225cc4401250e1b3349958` (1 source records).
  - [openai/01M5FD63HHDRPQPM5YMQWJ5HTN](../../data/2026/W43/openai.json); vendor URL: https://status.openai.com/incidents/01M5FD63HHDRPQPM5YMQWJ5HTN
- Start time basis: `vendor_reported`.

<a id="ci-5e97f10fd3ad8b465e170dcc"></a>

- Canonical ID: `ci-5e97f10fd3ad8b465e170dcc` (1 source records).
  - [grafana/3ft2hqqwnl2m](../../data/2026/W43/grafana.json); vendor URL: https://stspg.io/28py1pvhnpjf
- Start time basis: `vendor_reported`.

<a id="ci-fe8781ad2a5d93d98393abcd"></a>

- Canonical ID: `ci-fe8781ad2a5d93d98393abcd` (1 source records).
  - [cloudflare/p1vurmibvmag](../../data/2026/W43/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/p1vurmibvmag
- Start time basis: `vendor_reported`.

<a id="ci-a31b0ca45af63b2641680c14"></a>

- Canonical ID: `ci-a31b0ca45af63b2641680c14` (1 source records).
  - [openai/01M5JHVJVC8WAXB4WRYFMKG6HP](../../data/2026/W43/openai.json); vendor URL: https://status.openai.com/incidents/01M5JHVJVC8WAXB4WRYFMKG6HP
- Start time basis: `vendor_reported`.

<a id="ci-453fc38764c99e73a7be226a"></a>

- Canonical ID: `ci-453fc38764c99e73a7be226a` (1 source records).
  - [cloudflare/nwgwgbrn9a71](../../data/2026/W43/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/nwgwgbrn9a71
- Start time basis: `vendor_reported`.

<a id="ci-06167ecad8107e08f15af6ce"></a>

- Canonical ID: `ci-06167ecad8107e08f15af6ce` (1 source records).
  - [cloudflare/t77gwnpr6oe5](../../data/2026/W43/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/t77gwnpr6oe5
- Start time basis: `vendor_reported`.

<a id="ci-a58e3babedc502dcbbb1e2d8"></a>

- Canonical ID: `ci-a58e3babedc502dcbbb1e2d8` (1 source records).
  - [github/71y513kjvbzh](../../data/2026/W43/github.json); vendor URL: https://stspg.io/zdk3zqt69fnz
- Start time basis: `vendor_reported`.

<a id="ci-6e1b833564feb6974890f8ef"></a>

- Canonical ID: `ci-6e1b833564feb6974890f8ef` (1 source records).
  - [cloudflare/sufgcbe8rrwt](../../data/2026/W43/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/sufgcbe8rrwt
- Start time basis: `vendor_reported`.

<a id="ci-8abff99f6c7e7ecd05f25023"></a>

- Canonical ID: `ci-8abff99f6c7e7ecd05f25023` (1 source records).
  - [github/8tzbpcnl75jp](../../data/2026/W43/github.json); vendor URL: https://stspg.io/nrv2k8s68bwn
- Start time basis: `vendor_reported`.

<a id="ci-9bd5114735766713a9c58d1f"></a>

- Canonical ID: `ci-9bd5114735766713a9c58d1f` (1 source records).
  - [stripe/34c66nvjvk2g](../../data/2026/W43/stripe.json); vendor URL: https://stspg.io/k2jfv0fg9ff6
- Start time basis: `vendor_reported`.

<a id="ci-4aee7b2692cef9d264269623"></a>

- Canonical ID: `ci-4aee7b2692cef9d264269623` (1 source records).
  - [anthropic/njk29r81v3wg](../../data/2026/W43/anthropic.json); vendor URL: https://stspg.io/cgklr4v2sgkh
- Start time basis: `vendor_reported`.

<a id="ci-e77ad4a5540c2ad86a0c1309"></a>

- Canonical ID: `ci-e77ad4a5540c2ad86a0c1309` (1 source records).
  - [anthropic/0ns5clsk8p2x](../../data/2026/W43/anthropic.json); vendor URL: https://stspg.io/7n1wgb6qsfly
- Start time basis: `vendor_reported`.

<a id="ci-81b99475136688f5534f1fed"></a>

- Canonical ID: `ci-81b99475136688f5534f1fed` (1 source records).
  - [openai/01M5KDJ4079XQTZDSRX40NMN6H](../../data/2026/W43/openai.json); vendor URL: https://status.openai.com/incidents/01M5KDJ4079XQTZDSRX40NMN6H
- Start time basis: `vendor_reported`.

<a id="ci-a2cd10e989fc6f692657b7b2"></a>

- Canonical ID: `ci-a2cd10e989fc6f692657b7b2` (1 source records).
  - [anthropic/h101vcmhnfky](../../data/2026/W43/anthropic.json); vendor URL: https://stspg.io/gqp288t3qzpt
- Start time basis: `vendor_reported`.

<a id="ci-2b70e77251e9dbd7c301bf32"></a>

- Canonical ID: `ci-2b70e77251e9dbd7c301bf32` (1 source records).
  - [anthropic/vx06kp8wfpq2](../../data/2026/W43/anthropic.json); vendor URL: https://stspg.io/y0gkl12ztjst
- Start time basis: `vendor_reported`.

<a id="ci-4c4b0b5bccc5651f9fa80b23"></a>

- Canonical ID: `ci-4c4b0b5bccc5651f9fa80b23` (1 source records).
  - [openai/01M5JF8DR3EGQJNTXPZ24AKX2G](../../data/2026/W43/openai.json); vendor URL: https://status.openai.com/incidents/01M5JF8DR3EGQJNTXPZ24AKX2G
- Start time basis: `vendor_reported`.

<a id="ci-f77218a758067ac55f5709af"></a>

- Canonical ID: `ci-f77218a758067ac55f5709af` (1 source records).
  - [stripe/9jlv78n17jhj](../../data/2026/W43/stripe.json); vendor URL: https://stspg.io/ncmhd5df6f3k
- Start time basis: `vendor_reported`.

<a id="ci-55d6f7311d52643fa1267d42"></a>

- Canonical ID: `ci-55d6f7311d52643fa1267d42` (1 source records).
  - [grafana/hznszb0tghz6](../../data/2026/W43/grafana.json); vendor URL: https://stspg.io/t7cdk3q6bcz9
- Start time basis: `vendor_reported`.

<a id="ci-02eab138d0f5a589a622c3d4"></a>

- Canonical ID: `ci-02eab138d0f5a589a622c3d4` (1 source records).
  - [stripe/wklqnr97wbqs](../../data/2026/W43/stripe.json); vendor URL: https://stspg.io/gmpxzbw40g3d
- Start time basis: `vendor_reported`.

<a id="ci-d10fb3679f8349287e4de6bc"></a>

- Canonical ID: `ci-d10fb3679f8349287e4de6bc` (1 source records).
  - [cloudflare/33k6g8j7mhah](../../data/2026/W43/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/33k6g8j7mhah
- Start time basis: `vendor_reported`.

<a id="ci-50c27fe370d083e5257808e9"></a>

- Canonical ID: `ci-50c27fe370d083e5257808e9` (1 source records).
  - [cloudflare/l0gsc0sccx2v](../../data/2026/W43/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/l0gsc0sccx2v
- Start time basis: `vendor_reported`.

<a id="ci-0cdfa8a494248e3c6140c18c"></a>

- Canonical ID: `ci-0cdfa8a494248e3c6140c18c` (1 source records).
  - [circleci/rrrvc4lqplb5](../../data/2026/W43/circleci.json); vendor URL: https://stspg.io/pp6wqy3vlms3
- Start time basis: `vendor_reported`.

<a id="ci-ff88dcdc38bab1e99967a161"></a>

- Canonical ID: `ci-ff88dcdc38bab1e99967a161` (1 source records).
  - [grafana/zdwn7zxghn9q](../../data/2026/W43/grafana.json); vendor URL: https://stspg.io/wh6fr7yf9j5h
- Start time basis: `vendor_reported`.

<a id="ci-cf380e3999894535a76095e4"></a>

- Canonical ID: `ci-cf380e3999894535a76095e4` (1 source records).
  - [openai/01M5FHKHPQJMSHCE2B2WDPZW2N](../../data/2026/W43/openai.json); vendor URL: https://status.openai.com/incidents/01M5FHKHPQJMSHCE2B2WDPZW2N
- Start time basis: `vendor_reported`.

<a id="ci-458f7145445816ce851f9158"></a>

- Canonical ID: `ci-458f7145445816ce851f9158` (1 source records).
  - [grafana/bpld6z1v27qx](../../data/2026/W43/grafana.json); vendor URL: https://stspg.io/xk032xv99zbr
- Start time basis: `vendor_reported`.

<a id="ci-921164f487b5331e59ba59df"></a>

- Canonical ID: `ci-921164f487b5331e59ba59df` (1 source records).
  - [openai/01M5E4G39TD1BW1QWHZG6RFA5H](../../data/2026/W43/openai.json); vendor URL: https://status.openai.com/incidents/01M5E4G39TD1BW1QWHZG6RFA5H
- Start time basis: `vendor_reported`.

<a id="ci-d9cc4bc829d37ed93096a4a9"></a>

- Canonical ID: `ci-d9cc4bc829d37ed93096a4a9` (1 source records).
  - [grafana/gt3xltfs1m60](../../data/2026/W43/grafana.json); vendor URL: https://stspg.io/y9w2szsh123d
- Start time basis: `vendor_reported`.

<a id="ci-f6e55546037ba2c221fc1850"></a>

- Canonical ID: `ci-f6e55546037ba2c221fc1850` (1 source records).
  - [github/q29p76qtnkyk](../../data/2026/W43/github.json); vendor URL: https://stspg.io/44w94j036yc8
- Start time basis: `vendor_reported`.

<a id="ci-2deb47fcafe3e70c16fd7f36"></a>

- Canonical ID: `ci-2deb47fcafe3e70c16fd7f36` (1 source records).
  - [grafana/x264qsyj6cgd](../../data/2026/W43/grafana.json); vendor URL: https://stspg.io/p5hhk4h6k6dm
- Start time basis: `vendor_reported`.

<a id="ci-2d735d5c8139579e4b57c138"></a>

- Canonical ID: `ci-2d735d5c8139579e4b57c138` (1 source records).
  - [circleci/5jb50gh55gtj](../../data/2026/W43/circleci.json); vendor URL: https://stspg.io/8g54j07z69lb
- Start time basis: `vendor_reported`.

<a id="ci-60a91fc8ae52377941ed8cad"></a>

- Canonical ID: `ci-60a91fc8ae52377941ed8cad` (1 source records).
  - [cloudflare/pij2josz3rvz](../../data/2026/W43/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/pij2josz3rvz
- Start time basis: `vendor_reported`.

<a id="ci-584456f2aeadcf33508cead5"></a>

- Canonical ID: `ci-584456f2aeadcf33508cead5` (1 source records).
  - [cloudflare/lhdg2tq46faq](../../data/2026/W43/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/lhdg2tq46faq
- Start time basis: `vendor_reported`.

<a id="ci-1b87e39f9e5d5648fe7f1f41"></a>

- Canonical ID: `ci-1b87e39f9e5d5648fe7f1f41` (1 source records).
  - [openai/01M5NFF94XCJ6CGTD4S78A9BEB](../../data/2026/W43/openai.json); vendor URL: https://status.openai.com/incidents/01M5NFF94XCJ6CGTD4S78A9BEB
- Start time basis: `vendor_reported`.

<a id="ci-f1c67b705ad4d71463330c57"></a>

- Canonical ID: `ci-f1c67b705ad4d71463330c57` (1 source records).
  - [cloudflare/qlgdup1n0ge7](../../data/2026/W43/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/qlgdup1n0ge7
- Start time basis: `vendor_reported`.

<a id="ci-37397510c3db2a1e6b641bbd"></a>

- Canonical ID: `ci-37397510c3db2a1e6b641bbd` (1 source records).
  - [stripe/vjz0mvsck768](../../data/2026/W43/stripe.json); vendor URL: https://stspg.io/w2qqtf6tbrsy
- Start time basis: `vendor_reported`.

<a id="ci-eaa614fa75868b7f98852e07"></a>

- Canonical ID: `ci-eaa614fa75868b7f98852e07` (1 source records).
  - [openai/01M5A5S1ATN3QB2WCHQ3VDZC81](../../data/2026/W43/openai.json); vendor URL: https://status.openai.com/incidents/01M5A5S1ATN3QB2WCHQ3VDZC81
- Start time basis: `vendor_reported`.

<a id="ci-b8aefa5735613ca94827f433"></a>

- Canonical ID: `ci-b8aefa5735613ca94827f433` (1 source records).
  - [github/khv5nj9qy2kk](../../data/2026/W43/github.json); vendor URL: https://stspg.io/9hysn9g59lpg
- Start time basis: `vendor_reported`.

<a id="ci-6069ac409566a4f21c7c92e0"></a>

- Canonical ID: `ci-6069ac409566a4f21c7c92e0` (1 source records).
  - [github/ntrfd2jqldn6](../../data/2026/W43/github.json); vendor URL: https://stspg.io/kmk42jg3v1sm
- Start time basis: `vendor_reported`.

<a id="ci-706ecc4c0e4b7a170f96ba0c"></a>

- Canonical ID: `ci-706ecc4c0e4b7a170f96ba0c` (1 source records).
  - [openai/01M5NAR5GZC5M7D0S7AQ6CGEEN](../../data/2026/W43/openai.json); vendor URL: https://status.openai.com/incidents/01M5NAR5GZC5M7D0S7AQ6CGEEN
- Start time basis: `vendor_reported`.

<a id="ci-8446ebf9444aecee24c74fbe"></a>

- Canonical ID: `ci-8446ebf9444aecee24c74fbe` (1 source records).
  - [cloudflare/ce61s9m7vz48](../../data/2026/W43/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/ce61s9m7vz48
- Start time basis: `vendor_reported`.

<a id="ci-6f62729cd086ee0352b91c28"></a>

- Canonical ID: `ci-6f62729cd086ee0352b91c28` (1 source records).
  - [cloudflare/t0p23w62n13j](../../data/2026/W43/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/t0p23w62n13j
- Start time basis: `vendor_reported`.

<a id="ci-bf4547b0669eacc6e8a6631a"></a>

- Canonical ID: `ci-bf4547b0669eacc6e8a6631a` (1 source records).
  - [anthropic/3cqf3pmlnr96](../../data/2026/W43/anthropic.json); vendor URL: https://stspg.io/dmn3bc440bnb
- Start time basis: `vendor_reported`.

<a id="ci-5febe131b295252ccfbc706a"></a>

- Canonical ID: `ci-5febe131b295252ccfbc706a` (1 source records).
  - [grafana/lvwsbpw8pr3n](../../data/2026/W43/grafana.json); vendor URL: https://stspg.io/3xcm1c985jpl
- Start time basis: `vendor_reported`.

<a id="ci-2c32255340a752854a0604f5"></a>

- Canonical ID: `ci-2c32255340a752854a0604f5` (1 source records).
  - [openai/01M5G9QQ9V4CYDWFRYF9BWZ5GG](../../data/2026/W43/openai.json); vendor URL: https://status.openai.com/incidents/01M5G9QQ9V4CYDWFRYF9BWZ5GG
- Start time basis: `vendor_reported`.
- Warning: Vendor resolved_at precedes started_at; duration is unknown, score uses zero duration; decision:legacy-time-openai-01M5G9QQ9V4CYDWFRYF9BWZ5GG

<a id="ci-2c165761bc4d805b2e9112b7"></a>

- Canonical ID: `ci-2c165761bc4d805b2e9112b7` (1 source records).
  - [slack/1582](../../data/2026/W43/slack.json); vendor URL: https://slack-status.com/2026-10/58a30c7fc8aa3e80
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
