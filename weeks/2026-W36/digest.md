# Week 2026-W36

2026-08-31 to 2026-09-06 (UTC). Generated 2026-11-14 12:36.

56 canonical incidents, 32 above threshold. 2 of 14 sources unavailable: datadog, yandexcloud.

56 source records (63 saved revisions) → 56 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | gcp | [Multiple products in us-central1-b are experiencing network service degradation.](https://status.cloud.google.com/incidents/J5ia5t9p3g9Q5Wi7r8Ev) | 2026-09-01 14:42 | 4h 10m | major | AlloyDB for PostgreSQL, Apigee, Cloud Filestore, Cloud Firestore, Cloud Run, Cloud Spanner, Google App Engine, Google BigQuery, Google Cloud Bigtable, Google Cloud Dataflow, Google Cloud SQL, Google Compute Engine, Google Kubernetes Engine, Hybrid Connectivity, Looker (Google Cloud core), Virtual Private Cloud (VPC) | 8000 = 250 × 16 × 2.0 |
| 2 | cloudflare | [HTTP/3 issue affecting R2 custom domains](https://www.cloudflarestatus.com/incidents/8x66bpk6p9kk) | 2026-08-31 18:47 | 3d 0h 45m | minor | R2 | 7857 = 4365 × 1 × 1.8 |
| 3 | grafana | [US Central Region Instability](https://stspg.io/b3b4ppnkx23p) | 2026-09-04 16:45 | 6h 14m | minor | GCP US Central - prod-us-central-0: Ingestion, GCP US Central - prod-us-central-0: Ingestion, GCP US Central - prod-us-central-0: Metrics Generator, GCP US Central - prod-us-central-0: Querying, GCP US Central - prod-us-central-0: Querying, GCP US Central - prod-us-central-5: Ingestion, GCP US Central - prod-us-central-5: Ingestion, GCP US Central - prod-us-central-5: Querying, GCP US Central - prod-us-central-5: Querying | 3366 = 374 × 9 × 1.0 |
| 4 | openai | [Elevated latency in the Responses API](https://status.openai.com/incidents/01M1CYZDD84XD0H56XNEV2PT98) | 2026-08-31 22:27 | 20h 37m | minor | — | 1856 = 1237 × 1 × 1.5 |
| 5 | yandexcloud | [Problems with multiple services in the kz1-a availability zone](https://status.yandex.cloud/en/incidents/1992) | 2026-09-01 16:47 | 1h 13m | major | Cloud DNS, Data Transfer, Identity and Access Management, Managed Service for Apache Spark™, Managed Service for OpenSearch, Managed Service for YDB, Managed Service for Apache Kafka®, Managed Service for ClickHouse®, Managed Service for MySQL®, Managed Service for PostgreSQL, Managed Service for Valkey™, SpeechKit, Yandex Cloud Console, Yandex MPP Analytics Engine for PostgreSQL | 1533 = 73 × 14 × 1.5 |
| 6 | cloudflare | [Cloudflare Access one-time PIN emails blocked by certain email security gateways, including Proofpoint](https://www.cloudflarestatus.com/incidents/cczd80189gvc) | 2026-09-01 20:06 | 11h 9m | minor | Access | 1204 = 669 × 1 × 1.8 |
| 7 | cloudflare | [Pipelines Creation Issue](https://www.cloudflarestatus.com/incidents/w6ng4jcxv7m1) | 2026-09-01 21:51 | 10h 7m | minor | Pipelines | 1093 = 607 × 1 × 1.8 |
| 8 | anthropic | [Elevated errors for multiple models](https://stspg.io/9xz4hhmd1jzn) | 2026-09-03 13:26 | 2h 57m | major | Claude API (api.anthropic.com), Claude Code, Claude Cowork, claude.ai | 1062 = 177 × 4 × 1.5 |
| 9 | grafana | [High Latency in prod-ap-south-1](https://stspg.io/f6h7k3lq3v62) | 2026-09-03 09:32 | 3h 17m | major | AWS India - prod-ap-south-1, AWS India - prod-ap-south-1, AWS India - prod-ap-south-1: Ingestion, AWS India - prod-ap-south-1: Ingestion | 788 = 197 × 4 × 1.0 |
| 10 | cloudflare | [Requests from wrangler or other OAuth clients may experience elevated errors](https://www.cloudflarestatus.com/incidents/3mjmdcmm86nx) | 2026-09-01 08:38 | 3h 27m | minor | D1, Workflows | 745 = 207 × 2 × 1.8 |
| 11 | grafana | [Degradation of Hosted Grafana in US Central Region](https://stspg.io/7n6x5qw9x9d5) | 2026-09-04 17:03 | 3h 39m | major | GCP US Central - prod-us-central-0, GCP US Central - prod-us-central-3, GCP US Central - prod-us-central-4 | 657 = 219 × 3 × 1.0 |
| 12 | grafana | [Tempo and Mimir Read and Write Failures](https://stspg.io/kw7xrpjrg134) | 2026-09-03 22:13 | 2h 42m | minor | AWS US West - prod-us-west-0: Ingestion, AWS US West - prod-us-west-0: Ingestion, AWS US West - prod-us-west-0: Querying, AWS US West - prod-us-west-0: Querying | 648 = 162 × 4 × 1.0 |
| 13 | cloudflare | [Purchased domains from registrar not appearing in Cloudflare Dashboard](https://www.cloudflarestatus.com/incidents/1ghtnh94pby6) | 2026-09-06 17:06 | 4h 43m | minor | Registrar | 509 = 283 × 1 × 1.8 |
| 14 | cloudflare | [Workers Builds elevated queue times](https://www.cloudflarestatus.com/incidents/48wc4y9p1vym) | 2026-09-03 23:35 | 4h 33m | minor | Workers Builds | 491 = 273 × 1 × 1.8 |
| 15 | openai | [ChatGPT Work seeing elevated errors and latency](https://status.openai.com/incidents/01M1C5M4K0WC8PPT0Z175RJA1E) | 2026-08-31 15:04 | 5h 23m | major | — | 484 = 323 × 1 × 1.5 |
| 16 | cloudflare | [Workers KV experiencing elevated error rates in Western Europe Region](https://www.cloudflarestatus.com/incidents/kby271yt9vwx) | 2026-08-31 11:01 | 4h 2m | minor | Workers KV | 436 = 242 × 1 × 1.8 |
| 17 | cloudflare | [Elevated number of R2 503 errors in Eastern North America region](https://www.cloudflarestatus.com/incidents/ftvf8c3m4mv5) | 2026-09-04 21:10 | 3h 31m | minor | R2 | 380 = 211 × 1 × 1.8 |
| 18 | grafana | [Alert rule creation, deletion, and update degradation in prod-us-east-2](https://stspg.io/8gw4nt4mmp09) | 2026-09-04 13:36 | 2h 56m | minor | AWS US East (OH) - prod-us-east-0: Alertmanager, AWS US East (OH) - prod-us-east-0: Alertmanager and Rules Configuration API | 352 = 176 × 2 × 1.0 |
| 19 | cloudflare | [Cloudflare Access bypass IP service experiencing availability issues](https://www.cloudflarestatus.com/incidents/fcdl2yb3q7qw) | 2026-08-31 11:40 | 3h 10m | minor | Access | 342 = 190 × 1 × 1.8 |
| 20 | datadog | [Delayed CICD Optimization, Code Coverage, Code Security, and DORA data](https://stspg.io/zstf1zx7hfy3) | 2026-09-03 17:20 | 5h 40m | minor | CI Visibility | 340 = 340 × 1 × 1.0 |
| 21 | openai | [Users in APAC region may face increased error in ChatGPT, Work, image generation, file upload, Voice, and Codex Cloud](https://status.openai.com/incidents/01M1NKFZH5EEYEREC54HNAHY35) | 2026-09-04 07:00 | 3h 46m | minor | — | 339 = 226 × 1 × 1.5 |
| 22 | grafana | [Investigating issues in US Central (prod-us-central-0, prod-us-central-5)](https://stspg.io/jvcf7d0j5x37) | 2026-09-01 15:25 | 4h 54m | minor | — | 294 = 294 × 1 × 1.0 |
| 23 | aws | [Increased API Error Rates](https://health.aws.amazon.com/health/status) | 2026-09-03 21:49 | 2h 17m | unknown | multipleservices-us-west-2 | 274 = 137 × 1 × 2.0 |
| 24 | github | [Incident with Grok Copilot AI Model Provider](https://stspg.io/p217h6l54208) | 2026-09-03 14:17 | 2h 54m | minor | Copilot AI Model Providers | 261 = 174 × 1 × 1.5 |
| 25 | yandexcloud | [Partial unavailability of the Cloud Backup service](https://status.yandex.cloud/en/incidents/1993) | 2026-09-02 07:00 | 2h 51m | minor | Cloud Backup | 256 = 171 × 1 × 1.5 |
| 26 | cloudflare | [Durable Objects increased errors in Western North America](https://www.cloudflarestatus.com/incidents/3qs7m80p13v2) | 2026-09-02 13:43 | 2h 7m | minor | Durable Objects | 229 = 127 × 1 × 1.8 |
| 27 | cloudflare | [Cache Purging Errors](https://www.cloudflarestatus.com/incidents/ff0yvx1d6tby) | 2026-09-04 14:51 | 2h 7m | minor | CDN Cache Purge | 229 = 127 × 1 × 1.8 |
| 28 | github | [Disruption with Copilot Code Review](https://stspg.io/dddfhld6fj0z) | 2026-09-04 20:39 | 1h 47m | major | — | 160 = 107 × 1 × 1.5 |
| 29 | anthropic | [Elevated errors for Claude Sonnet 5](https://stspg.io/9dgzwz3ymz42) | 2026-09-02 21:17 | 26m | major | Claude API (api.anthropic.com), Claude Code, Claude Cowork, claude.ai | 156 = 26 × 4 × 1.5 |
| 30 | anthropic | [Elevated errors for Claude Sonnet 5](https://stspg.io/jt14jk41pjsf) | 2026-09-03 12:37 | 18m | major | Claude API (api.anthropic.com), Claude Code, Claude Cowork, claude.ai | 108 = 18 × 4 × 1.5 |
| 31 | cloudflare | [Network Route Leak in Palmas, Brazil](https://www.cloudflarestatus.com/incidents/1w42pbthkqjf) | 2026-09-01 21:47 | 18m | major | Palmas, Brazil - (PMW) | 32 = 18 × 1 × 1.8 |
| 32 | openai | [ChatGPT Work Mode High Error Rates](https://status.openai.com/incidents/01M1J99EDCYED5GP2W8GZ3N2NA) | 2026-09-03 00:04 | 6m | major | — | 9 = 6 × 1 × 1.5 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. gcp — Multiple products in us-central1-b are experiencing network service degradation.

- URL: https://status.cloud.google.com/incidents/J5ia5t9p3g9Q5Wi7r8Ev
- Started: 2026-09-01 14:42 UTC
- Resolved: 2026-09-01 18:52 UTC
- Duration: 4h 10m
- Impact: major, status: resolved
- Components: AlloyDB for PostgreSQL, Apigee, Cloud Filestore, Cloud Firestore, Cloud Run, Cloud Spanner, Google App Engine, Google BigQuery, Google Cloud Bigtable, Google Cloud Dataflow, Google Cloud SQL, Google Compute Engine, Google Kubernetes Engine, Hybrid Connectivity, Looker (Google Cloud core), Virtual Private Cloud (VPC)
- Canonical incident: [ci-5650bfa3e91cb26e3635de5e](#ci-5650bfa3e91cb26e3635de5e)

- `2026-09-01 16:03` **investigating** — **Summary** Multiple products in us-central1-b are experiencing network service degradation. **Description** We are experiencing a networking issue starting Tuesday, 2026-09-01 07:41 US/Pacific. The issue is affecting a portion of the us-central1-b zone. The remaining zones in the region are unaffected. Our network engineering team has put in mitigations and we are seeing services starting to recover. We will provide more information by Tuesday, 2026-09-01 09:30 US/Pacific with current details. We apologize to all who are affected by the disruption. **Customer Symptoms** Customers are experiencing elevated packet loss and errors. * BigQuery * Cloud Dataflow * Cloud SQL * App Engine * Cloud Run * AlloyDB * Cloud Bigtable * Cloud Spanner * Compute Engine * Cloud VPN * Cloud Interconnect * Cloud NAT * Cloud Filestore * Cloud Hybrid Connectivity * Virtual Private Cloud * Google Kubernetes Engine * Looker (Google Cloud core) **Workaround** * Customers with service layer retries may see successful completions as only a portion of the zone is affected. * Customers that are severely affected and have services available in alternate zones can failover to alternate zones.
- `2026-09-01 16:19` **investigating** — **Summary** Multiple products in us-central1-b are experiencing network service degradation. **Description** The issue is affecting a portion of the us-central1-b zone. The remaining zones in the region are unaffected. Following network-level mitigation efforts, the underlying network infrastructure has fully recovered. Product engineering teams across Cloud are currently validating the recovery status of individual services. We will provide more information by Tuesday, 2026-09-01 10:00 US/Pacific with current details. **Customer Symptoms** Customers are experiencing elevated packet loss and errors. * BigQuery * Cloud Dataflow * Cloud SQL * App Engine * Cloud Run * AlloyDB * Cloud Bigtable * Cloud Spanner * Compute Engine * Cloud VPN * Cloud Interconnect * Cloud NAT * Cloud Filestore * Cloud Hybrid Connectivity * Virtual Private Cloud * Google Kubernetes Engine * Looker (Google Cloud core) * Apigee * Google SecOps SOAR **Workaround** * Customers with service layer retries may see successful completions as only a portion of the zone is affected. * Customers that are severely affected and have services available in alternate zones can failover to alternate zones.
- `2026-09-01 17:09` **investigating** — **Summary** Multiple products in us-central1-b are experiencing network service degradation. **Description** The issue is affecting a portion of the us-central1-b zone. The remaining zones in the region are unaffected. Most of the products have recovered and the engineering team is still validating a few remaining product impacts. We will provide more information by Tuesday, 2026-09-01 11:00 US/Pacific with current details. **Customer Symptoms** Customers are experiencing elevated packet loss and errors. * Cloud Run * App Engine Following products have recovered * BigQuery * Cloud Dataflow * Cloud Bigtable * Compute Engine * Cloud Hybrid Connectivity * Cloud Filestore * Virtual Private Cloud * Cloud Spanner * Google Kubernetes Engine * Cloud SQL * AlloyDBCloud VPN * Apigee * Cloud Interconnect * Cloud NAT * Looker (Google Cloud core) * Google SecOps SOAR * Cloud Firestore **Workaround** * Customers with service layer retries may see successful completions as only a portion of the zone is affected. * Customers that are severely affected and have services available in alternate zones can failover to alternate zones.
- `2026-09-01 18:01` **investigating** — **Summary** Multiple products in us-central1-b are experiencing network service degradation. **Description** The issue is affecting a portion of the us-central1-b zone. The remaining zones in the region are unaffected. Most of the products have recovered and the engineering team is still validating a few remaining product impacts. Routine network maintenance triggered issues in us-central1-b. We have recovered the network capacity that was impacted and restored all zonal and regional services. Maintenance in the region has been halted while proactive audits are carried out. We will provide more information by Tuesday, 2026-09-01 12:30 US/Pacific with current details. **Customer Symptoms** Customers are experiencing elevated packet loss and errors. * Cloud Run * App Engine Following products have recovered * BigQuery * Cloud Dataflow * Cloud Bigtable * Compute Engine * Cloud Hybrid Connectivity * Cloud Filestore * Virtual Private Cloud * Cloud Spanner * Google Kubernetes Engine * Cloud SQL * AlloyDBCloud VPN * Apigee * Cloud Interconnect * Cloud NAT * Looker (Google Cloud core) * Google SecOps SOAR * Cloud Firestore **Workaround** * Customers with service layer retries may see successful completions as only a portion of the zone is affected.
- `2026-09-01 18:36` **investigating** — **Summary** Multiple products in us-central1-b are experiencing network service degradation. **Description** The issue is affecting a portion of the us-central1-b zone. The remaining zones in the region are unaffected. Most of the products have recovered. The impacted products Cloud Run and App Engine are also seeing recovery and the errors have lowered pre-incident levels. Engineering continues to monitor and validate. Routine network fabric path maintenance triggered unexpected issues in two clusters in us-central1-b. We have recovered the network capacity that was impacted and restored all zonal and regional services. Maintenance in the region has been halted while proactive audits are carried out. We will provide more information by Tuesday, 2026-09-01 12:30 US/Pacific with current details. **Customer Symptoms** Customers are experiencing elevated packet loss and errors. * Cloud Run * App Engine Following products have recovered * BigQuery * Cloud Dataflow * Cloud Bigtable * Compute Engine * Cloud Hybrid Connectivity * Cloud Filestore * Virtual Private Cloud * Cloud Spanner * Google Kubernetes Engine * Cloud SQL * AlloyDBCloud VPN * Apigee * Cloud Interconnect * Cloud NAT * Looker (Google Cloud core) * Google SecOps SOAR * Cloud Firestore **Workaround** * Customers with service layer retries may see successful completions as only a portion of the zone is affected.
- `2026-09-01 19:12` **investigating** — **Summary** Multiple products in us-central1-b experienced network service degradation. **Description** The issue was affecting a portion of the us-central1-b zone. The remaining zones in the region were unaffected. All the impacted products have recovered. From preliminary analysis, routine network fabric path maintenance triggered unexpected issues in one cluster in us-central1-b. We have recovered the network capacity that was impacted and restored all zonal and regional services. Maintenance in the region has been halted while proactive audits are carried out. We thank you for your patience while we worked on resolving the issue. **Customer Symptoms** Customers experienced elevated packet loss and errors. **Workaround** The issue has been mitigated.
- `2026-09-03 17:47` **resolved** — ## Preliminary Incident Report We sincerely apologize for the disruption this incident caused to your business. We know how much you rely on Google Cloud, and we regret the impact on your environment. We are working to address the root cause and prevent this from occurring in the future. Please note, this information is based on our best knowledge at the time of posting and is subject to change as our investigation continues. A final Incident Report with preventative actions will be posted once our investigation is complete. If you have experienced impact outside of what is listed below, please reach out to Google Cloud Support using https://cloud.google.com/support ## Date/Time of the Issue (All time US/Pacific) **Incident Start**: 1 September 2026 07:41 **Incident End**: 1 September 2026 11:52 **Duration**: 4 hours, 11 minutes ## Summary On Tuesday, 1 September 2026, some customers in us-central1-b experienced network service degradation and instance isolation for a duration of 4 hours and 11 minutes. To our customers whose businesses were impacted during this disruption, we sincerely apologize. This is not the level of quality and reliability we strive to offer you, and we are taking immediate steps to improve the platform’s performance and availability. ## Preliminary Root Cause The immediate technical trigger for this event was the inadvertent physical disconnection of network fiber-optic cables during a routine hardware maintenance procedure. An engineer was performing a scheduled capacity upgrade on datacenter routers that support a fraction of capacity in the us-central1-b zone. The network architecture for each datacenter is designed with redundancy across multiple routing devices. The system is designed to be resilient to all single device or fiber-path failures, and most double- or triple-failures do not affect customer traffic. To ensure this, the devices and fiber paths are physically separated in each datacenter, with diverse power sources. A procedural error meant that the physical maintenance action sequentially unplugged 100% of fiber paths across all devices within 13 minutes. The nature of the error, combined with the speed of the action, prevented warnings of incorrect action reaching the engineer before complete disconnection. This resulted in compute capacity in the impacted zone being isolated from the network. Customers were unable to reach their virtual machines, and those virtual machines could not establish connections outside their zone. ## Remediation The issue was detected immediately by automated network loss monitoring systems as well as proactive probes, which rapidly engaged the network engineering and incident response teams. To mitigate the immediate customer impact, engineering teams actively moved traffic away from the impacted infrastructure, automatically shifting compatible workload traffic to healthy capacity elsewhere in the region. Concurrently, hardware operations technicians on site identified the disconnected optical links and physically reseated the fibers. Once the physical links were fully restored, traffic flow rates normalized and the traffic was redirected back to return the capacity to service. ## Description of Impact On Tuesday, 1 September, from 07:41 to 11:52 US/Pacific, a portion of the us-central1-b zone experienced severe network degradation and resource isolation. * **Multiple products affected:** The network disruption to/from a portion of us-central1-b impacted multiple GCP Products for some customers in that zone. * **Error rate:** Traffic flow drop rates for resources hosted in the affected area reached 100% during the peak of the incident, resulting in unreachable virtual machines and elevated packet loss. **Affected Services and Features** * **Google Compute Engine:** Inability to access virtual machines externally, and inability for VMs to reach remote resources. * **Google Kubernetes Engine:** Unreachable regional/zonal clusters and nodes within the affected zone, prompting delayed cluster recovery times. * **Cloud Run / Google App Engine:** Temporary latency spikes and pending queue aborts as backend workloads automatically evacuated and shifted to healthy capacity. A subset of workloads in the specifically impacted area experienced degradation until 11:52 US/Pacific. * **Cloud SQL, AlloyDB for PostgreSQL, Cloud Spanner, Google Cloud Bigtable, Cloud Filestore:** Data access and connectivity severed for instances localized strictly to the impacted infrastructure. * **Virtual Private Cloud (VPC), Cloud NAT, Apigee, Google BigQuery, Google Cloud Dataflow, Looker, Google SecOps SOAR, Hybrid Connectivity, Cloud Interconnect:** Elevated network packet loss, connectivity drops, and API timeouts in the impacted zone.
- `2026-09-10 21:20` **resolved** — ## Incident Report ## Summary On Tuesday, 1 September 2026, some customers in us-central1 experienced network service degradation and instance isolation for a duration of 4 hours and 11 minutes. To our customers whose businesses were impacted during this disruption, we sincerely apologize. This is not the level of quality and reliability we strive to offer you, and we are taking immediate steps to improve the platform’s performance and availability. ## Root Cause The event was triggered by the inadvertent physical disconnection of network fiber-optic cables during routine hardware maintenance. A technician was performing a scheduled capacity upgrade on data center routers that support a fraction of capacity in the us-central1-b zone and a small fraction of capacity in the us-central1-f zone. During the hands-on maintenance, optical fibers connecting the data center routers to the network fabric were unplugged and the optical transceivers were replaced with a transceiver that supports higher-density connectivity. This transceiver was not compatible with the transceivers still in place in the data center fabric, causing a loss of connectivity for each fiber. The network is designed with redundant routers in separate data center rooms within the same building so that failure of one router does not interrupt service. The maintenance was intended to upgrade one router at a time over several days, with traffic diversions at each step and verification between steps that the network has returned to a fully-connected state. However, a procedural error in the manually orchestrated upgrade process caused the complete list of transceiver replacements across all routers to be issued to the technician without instructions to sequence the work one router at a time. The maintenance workflow also did not include the expected human and software verification steps for detecting unintended disruption. As a result, the technician sequentially unplugged all fiber paths across the affected devices within 13 minutes. The speed and nature of the error prevented warnings of incorrect action from reaching the engineer before connectivity was lost. Additionally, a standing procedure to halt if light is detected on any fiber-optic cable after unplug was not followed. This resulted in compute capacity in the impacted zone being isolated from the network. Customers were unable to reach their virtual machines, and those virtual machines could not establish connections outside their zone. ## Remediation and Prevention The issue was detected immediately by automated network loss monitoring systems as well as proactive probes, which rapidly engaged the network engineering and incident response teams. To mitigate the immediate customer impact, engineering teams actively moved traffic away from the impacted infrastructure. Concurrently, hardware operations technicians on site identified the disconnected optical links and physically re-inserted the original optical transceivers. Once the physical links were fully restored, traffic flow rates normalized and the traffic was redirected back to return the capacity to service. Google is committed to preventing a repeat of this issue in the future and is completing the following actions: * Reinforce training in us-central1 region and globally around safe maintenance techniques, including the mandatory requirement to verify whether light is on every unplugged fiber. * Finish the migration of network upgrade workflows to a fully automatically sequenced orchestration system. Most common regional network workflows were completed in 2025, and zonal workflows are in progress. * Complete the deployment of work stop alerting for actions resulting in the unintended disconnection of live fiber cables in Google data centers. * Complete the deployment of the automated system that will move regional traffic away from faulty zones, reducing the time from approximately 19 minutes in this instance to around 5 minutes. ## Detailed Description of Impact On Tuesday, 1 September, from 07:41 to 11:52 US/Pacific, a portion of the us-central1-b and us-central1-f zones experienced severe network degradation and resource isolation. 1. Multiple products affected: The network disruption to/from a portion of us-central1-b and us-central1-f zones impacted all zonal GCP Products for some customers in those zones. 2. Regional impact: Regional products using capacity in the affected data center were affected until traffic diversions were fully in place at 08:00 US/Pacific. Google Kubernetes Engine and Cloud Run / Google App Engine had longer impacts, discussed below. 3. Error rate: Traffic flow drop rates for resources hosted in the affected area reached 100% during the peak of the incident, resulting in unreachable virtual machines and elevated packet loss. **Affected Services and Features** * Google Compute Engine: Inability for users in the affected portion of us-central1-b and us-central1-f zones to access virtual machines externally, and inability for VMs to reach remote resources. * Cloud SQL, AlloyDB for PostgreSQL, Google Cloud Bigtable, Cloud Filestore: Data access and connectivity severed for instances localized strictly to the impacted infrastructure. * Cloud Spanner, Cloud Firestore: Elevated Remote Procedure Call (RPC) error rates, write timeouts, and latency spikes for the nam5 multi-region instances due to severed connectivity to components running in the impacted zone. * Virtual Private Cloud (VPC), Cloud NAT, Apigee, Google BigQuery, Google Cloud Dataflow, Looker, Google SecOps SOAR, Hybrid Connectivity, Cloud Interconnect: Elevated network packet loss, connectivity drops, and API timeouts in the impacted zone. * Google Kubernetes Engine (GKE): Between 07:40 and 09:15 US/Pacific on 2026-09-01, customers experienced errors and timeouts on requests to their GKE cluster control planes (Kubernetes API servers) in us-central1. This could have caused new or updated workloads to fail scheduling and Kubernetes API operations to fail. * Cloud Run / Google App Engine: Temporary latency spikes and pending queue aborts as backend workloads automatically evacuated and shifted to healthy capacity. A subset of workloads in the specifically impacted area experienced degradation until 11:52 US/Pacific. **Customer Impact** Customers with resources hosted in the specific affected clusters within us-central1-b and us-central1-f experienced a complete loss of network connectivity starting at 07:41 US/Pacific. During this time, virtual machines and associated services became unreachable from the internet and from other Google Cloud regions, and those internal resources could not initiate outbound connections. The issue was detected almost immediately and engineers took traffic diversion actions at 07:45 and 08:00 US/Pacific to reduce impact to regional products. By 08:50 US/Pacific, the majority of physical connections were restored and traffic began recovering, and the traffic diversion actions were removed at 09:19 US/Pacific. Full recovery across nearly all affected services and long-running operations was confirmed by 11:52 US/Pacific. Because the network isolation was strictly limited to specific clusters within us-central1-b and us-central1-f, multi-zonal deployments correctly utilizing redundancy across other zones in us-central1 were largely able to bypass the physical hardware failure and continue serving traffic. Customers with projects in both us-central1-b and us-central1-f would not have seen multi-zonal impact - at most one of their zones would have capacity in the affected data center.
- `2026-09-29 16:09` **resolved** — ## Addendum to Incident Report This addendum extends the Detailed Description of Impact from the previous message. Though only two clusters in a datacenter in us-central1-b and us-central1-f experienced network isolation, a few regional products with a dependency on the affected datacenter were also affected. Between 07:41 and 11:52 US/Pacific on Tuesday, September 1, 2026, the following products were affected: - **Compute Engine VMs**: VMs running in the affected data center lost network connectivity to resources located outside of their datacenter. From a Compute Engine perspective, this incident only affected a subset of VMs in us-central1-b and us-central1-f. - **Zonal and Regional Products Based on Compute Engine** Products with dependencies on VMs in the affected datacenter experienced symptoms like connection timeouts or service unavailability because the underlying VMs lost network connectivity. Examples of products that are based on Compute Engine VMs include: Apigee, Cloud SQL, AlloyDB for PostgreSQL, Dataflow, Cloud Filestore, and Looker. - **Google Kubernetes Engine (GKE)**: GKE clusters whose control plane had dependencies on Compute Engine VMs in the affected clusters experienced errors and timeouts when administrators or nodes attempted to connect to the Kubernetes API server. Consequently, deployments might have been delayed, Kubernetes API operations might have failed, and new nodes might not have registered with the control plane (causing node pool operational delays). - **Cloud Run and App Engine:** Workloads in the us-central1 region experienced elevated serving latency, 5xx aborted request errors, and outbound connectivity degradation. To prevent severe request failures, traffic was diverted to other clusters. Diverting requests to other clusters temporarily caused elevated latency or errors in those clusters.Once the impacted cluster was restored, workload demand for these new backends created a 'thundering herd' of concurrent instance initializations. This sudden surge bypassed warm caches and overwhelmed internal file-serving tiers, extending symptoms until emergency capacity was provisioned. - **Other Google APIs and Services**, in us-central1 or a multi-region that includes us-central1, were affected if they had dependencies on software tasks in the affected clusters and weren't able to use software tasks in an unaffected cluster. Symptoms included connection timeouts or service unavailability. Examples of other Google APIs and services are: Google Cloud Bigtable, BigQuery, Google SecOps SOAR, Cloud Spanner, and Cloud Firestore. For some services, symptoms were limited to temporary latency spikes and service unavailability until workloads shifted to unaffected clusters. - **Private Service Connect**: Clients (in any region) experienced connectivity issues involving PSC endpoints with dependencies in us-central1, in specific circumstances: - PSC endpoints for Google APIs were affected in the same way as noted in the Google APIs and services section. - Producer/consumer PSC endpoints were affected if the producer had load balancers or backend VMs in the affected clusters of us-central1. - **Hybrid Connectivity**: Some Cloud Router, Cloud VPN, and Cloud Interconnect VLAN attachments in us-central1 were affected in specific circumstances: - Cloud Router BGP software tasks for some Cloud Routers in us-central1 experienced a Cloud Router maintenance event if their BGP software tasks were running in the affected clusters. - Cloud VPN tunnel software tasks running in the affected clusters lost network connectivity. This caused some Cloud VPN tunnels in us-central1 to go down until a replacement tunnel task in a different cluster was assigned. - Resources (in any region) sending packets through Cloud Interconnect VLAN attachments in us-central1 might have experienced intermittent connectivity issues in certain situations where network flows weren’t already programmed. - **Envoy-based load balancers and other products**: Load balancers and other products that depended on managed Envoy tasks of a proxy-only subnet in us-central1 experienced increased latency or connectivity issues if enough Envoy tasks were in the affected clusters. The following products were affected: - Regional external application load balancers in us-central1 - Regional internal application load balancers in us-central1 - Cross-region internal application load balancers with proxies in us-central1 - Regional external proxy network load balancers in us-central1 - Regional internal proxy network load balancers in us-central1 - Cross-region internal proxy network load balancers with proxies in us-central1 - Secure Web Proxy in us-central1

### 2. cloudflare — HTTP/3 issue affecting R2 custom domains

- URL: https://www.cloudflarestatus.com/incidents/8x66bpk6p9kk
- Started: 2026-08-31 18:47 UTC
- Resolved: 2026-09-03 19:32 UTC
- Duration: 3d 0h 45m
- Impact: minor, status: resolved
- Components: R2
- Canonical incident: [ci-93c7789a5f6839ae3a7974bb](#ci-93c7789a5f6839ae3a7974bb)

- `2026-08-31 18:47` **investigating** — Some Firefox users may see delayed or failed asset loads on R2 custom domains due to an incorrect HTTP/3 advertisement. We are working on a fix.
- `2026-09-01 14:46` **identified** — The issue has been identified and a fix is being implemented.
- `2026-09-03 15:14` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-09-03 19:32` **resolved** — This incident has been resolved.

### 3. grafana — US Central Region Instability

- URL: https://stspg.io/b3b4ppnkx23p
- Started: 2026-09-04 16:45 UTC
- Resolved: 2026-09-04 23:00 UTC
- Duration: 6h 14m
- Impact: minor, status: resolved
- Components: GCP US Central - prod-us-central-0: Ingestion, GCP US Central - prod-us-central-0: Ingestion, GCP US Central - prod-us-central-0: Metrics Generator, GCP US Central - prod-us-central-0: Querying, GCP US Central - prod-us-central-0: Querying, GCP US Central - prod-us-central-5: Ingestion, GCP US Central - prod-us-central-5: Ingestion, GCP US Central - prod-us-central-5: Querying, GCP US Central - prod-us-central-5: Querying
- Canonical incident: [ci-70f9b5947df78d16177cefc9](#ci-70f9b5947df78d16177cefc9)

- `2026-09-04 16:45` **identified** — We’re currently investigating an issue affecting the us central region due to ongoing cloud provider incident. Multiple components impacted. Our team is actively working on this. Thank you for your patience.
- `2026-09-04 18:58` **monitoring** — The underlying cloud provider issue is not yet fully resolved, but our systems have recovered and affected Grafana Cloud stacks are loading normally again. Some customers may still see brief intermittent slowness. We are continuing to monitor until our provider confirms resolution.
- `2026-09-04 23:00` **resolved** — The US Central region is fully recovered now and the issue is now resolved.

### 4. openai — Elevated latency in the Responses API

- URL: https://status.openai.com/incidents/01M1CYZDD84XD0H56XNEV2PT98
- Started: 2026-08-31 22:27 UTC
- Resolved: 2026-09-01 19:05 UTC
- Duration: 20h 37m
- Impact: minor, status: resolved
- Components: —
- Canonical incident: [ci-f78b9503c3eb6ba7e99b3bcd](#ci-f78b9503c3eb6ba7e99b3bcd)

- `2026-08-31 22:27` **identified** — We have identified an issue causing elevated latency for some Responses API requests. We are working on implementing a mitigation.
- `2026-08-31 23:13` **monitoring** — We have applied a mitigation and are monitoring recovery.
- `2026-09-01 01:46` **monitoring** — We have applied the mitigation and are monitoring the recovery.
- `2026-09-01 15:15` **identified** — We have identified elevated latencies impacting Responses API on the US endpoint.
- `2026-09-01 17:36` **identified** — We are implementing the mitigation.
- `2026-09-01 18:09` **monitoring** — We have applied the mitigation and are monitoring the recovery.
- `2026-09-01 19:05` **resolved** — All impacted services have now fully recovered.

### 5. yandexcloud — Problems with multiple services in the kz1-a availability zone

- URL: https://status.yandex.cloud/en/incidents/1992
- Started: 2026-09-01 16:47 UTC
- Resolved: 2026-09-01 18:00 UTC
- Duration: 1h 13m
- Impact: major, status: resolved
- Components: Cloud DNS, Data Transfer, Identity and Access Management, Managed Service for Apache Spark™, Managed Service for OpenSearch, Managed Service for YDB, Managed Service for Apache Kafka®, Managed Service for ClickHouse®, Managed Service for MySQL®, Managed Service for PostgreSQL, Managed Service for Valkey™, SpeechKit, Yandex Cloud Console, Yandex MPP Analytics Engine for PostgreSQL
- Canonical incident: [ci-2b37818ff156857c8838818e](#ci-2b37818ff156857c8838818e)

- `2026-09-01 17:01` **investigating** — We noticed problems with a number of services in the kz1-a availability zone. The developers are already aware and are fixing it.
- `2026-09-01 18:15` **resolved** — At 21:00 Moscow time, the problems with the services were resolved. If you still have questions about cloud resources, please contact technical support.

### 6. cloudflare — Cloudflare Access one-time PIN emails blocked by certain email security gateways, including Proofpoint

- URL: https://www.cloudflarestatus.com/incidents/cczd80189gvc
- Started: 2026-09-01 20:06 UTC
- Resolved: 2026-09-02 07:16 UTC
- Duration: 11h 9m
- Impact: minor, status: resolved
- Components: Access
- Canonical incident: [ci-8c18817154780b777347ab6e](#ci-8c18817154780b777347ab6e)

- `2026-09-01 20:06` **investigating** — Some customers using email security gateway like Proofpoint may experience delayed Cloudflare notifications and login codes. Proofpoint is temporarily deferring messages with an SMTP 421 4.7.0 response. Cloudflare will continue retrying delivery.
- `2026-09-02 07:16` **resolved** — This incident has been resolved.

### 7. cloudflare — Pipelines Creation Issue

- URL: https://www.cloudflarestatus.com/incidents/w6ng4jcxv7m1
- Started: 2026-09-01 21:51 UTC
- Resolved: 2026-09-02 07:58 UTC
- Duration: 10h 7m
- Impact: minor, status: resolved
- Components: Pipelines
- Canonical incident: [ci-4f3dd69d94dd51ba12a21465](#ci-4f3dd69d94dd51ba12a21465)

- `2026-09-01 21:51` **investigating** — Cloudflare is investigating an issue with creating new Pipeline resources, including Pipelines, Streams, and Sinks. Existing Pipelines are unaffected. More updates to follow shortly.
- `2026-09-01 23:00` **identified** — The issue has been identified and a fix is being implemented.
- `2026-09-01 23:46` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-09-02 07:58` **resolved** — This incident has been resolved.

### 8. anthropic — Elevated errors for multiple models

- URL: https://stspg.io/9xz4hhmd1jzn
- Started: 2026-09-03 13:26 UTC
- Resolved: 2026-09-03 16:23 UTC
- Duration: 2h 57m
- Impact: major, status: resolved
- Components: Claude API (api.anthropic.com), Claude Code, Claude Cowork, claude.ai
- Canonical incident: [ci-d842d706e233d509f3083a67](#ci-d842d706e233d509f3083a67)

- `2026-09-03 13:26` **investigating** — We are investigating elevated errors on requests to Claude Mythos 5.1, Claude Fable 5.1, and Claude Opus 5. We will provide an update as soon as possible.
- `2026-09-03 13:41` **identified** — We have identified the cause of elevated errors on requests to Claude Mythos 5.1, Claude Fable 5.1, and Claude Opus 5 and are working on a fix. We will provide an update as soon as possible.
- `2026-09-03 13:50` **identified** — An exhaustive list of affected models: Mythos/Fable 5.1, Mythos/Fable 5, Opus 5, Opus 4.8, Opus 4.6.
- `2026-09-03 14:49` **identified** — We are continuing to work on a fix for this issue.
- `2026-09-03 15:25` **identified** — The only affected models right now are Opus 4.8 and Opus 5. The rest of the models have recovered to baseline error rate.
- `2026-09-03 16:06` **monitoring** — A fix has been deployed for Claude Mythos 5.1, Claude Fable 5.1, and Claude Opus 5 and we are monitoring for recovery.
- `2026-09-03 16:23` **resolved** — The issue affecting Claude Mythos 5.1, Claude Fable 5.1, and Claude Opus 5 has been resolved. Impact has ended as of 9:16 PT / 16:16 UTC.

### 9. grafana — High Latency in prod-ap-south-1

- URL: https://stspg.io/f6h7k3lq3v62
- Started: 2026-09-03 09:32 UTC
- Resolved: 2026-09-03 12:49 UTC
- Duration: 3h 17m
- Impact: major, status: resolved
- Components: AWS India - prod-ap-south-1, AWS India - prod-ap-south-1, AWS India - prod-ap-south-1: Ingestion, AWS India - prod-ap-south-1: Ingestion
- Canonical incident: [ci-769082f75e1a72f2fcc729f6](#ci-769082f75e1a72f2fcc729f6)

- `2026-09-03 09:32` **investigating** — We are currently investigating an issue impacting deployments in prod-ap-south-1. Impacted customers may be experiencing increased latency and issues with ingestion.
- `2026-09-03 11:37` **monitoring** — A fix has been applied and we are seeing signs of recovery. We will continue to monitor the recovery.
- `2026-09-03 12:49` **resolved** — This incident has been resolved.

### 10. cloudflare — Requests from wrangler or other OAuth clients may experience elevated errors

- URL: https://www.cloudflarestatus.com/incidents/3mjmdcmm86nx
- Started: 2026-09-01 08:38 UTC
- Resolved: 2026-09-01 12:05 UTC
- Duration: 3h 27m
- Impact: minor, status: resolved
- Components: D1, Workflows
- Canonical incident: [ci-a96eeb9ccce075e6b5a733b3](#ci-a96eeb9ccce075e6b5a733b3)

- `2026-09-01 08:38` **investigating** — We are currently investigating this issue.
- `2026-09-01 08:42` **investigating** — We are continuing to investigate this issue.
- `2026-09-01 09:41` **identified** — The issue has been identified and a fix is being implemented.
- `2026-09-01 12:05` **resolved** — This incident has been resolved.

### 11. grafana — Degradation of Hosted Grafana in US Central Region

- URL: https://stspg.io/7n6x5qw9x9d5
- Started: 2026-09-04 17:03 UTC
- Resolved: 2026-09-04 20:42 UTC
- Duration: 3h 39m
- Impact: major, status: resolved
- Components: GCP US Central - prod-us-central-0, GCP US Central - prod-us-central-3, GCP US Central - prod-us-central-4
- Canonical incident: [ci-ac11f1a6fd089e9ddb1b326c](#ci-ac11f1a6fd089e9ddb1b326c)

- `2026-09-04 17:03` **identified** — We are aware of an issue causing some Grafana Cloud stacks in the US region to be unavailable or fail to load. The cause is reduced compute capacity in one of our cloud provider's availability zones, which is preventing affected stacks from starting. Other regions are unaffected. We are working with our cloud provider to restore capacity.
- `2026-09-04 18:24` **monitoring** — We are currently observing minimal impact and continuing to monitor the situation.
- `2026-09-04 20:42` **resolved** — The US Central region is now fully operational and this incident is now resolved.

### 12. grafana — Tempo and Mimir Read and Write Failures

- URL: https://stspg.io/kw7xrpjrg134
- Started: 2026-09-03 22:13 UTC
- Resolved: 2026-09-04 00:55 UTC
- Duration: 2h 42m
- Impact: minor, status: resolved
- Components: AWS US West - prod-us-west-0: Ingestion, AWS US West - prod-us-west-0: Ingestion, AWS US West - prod-us-west-0: Querying, AWS US West - prod-us-west-0: Querying
- Canonical incident: [ci-a420a3ad5b4504b9b5c08b7e](#ci-a420a3ad5b4504b9b5c08b7e)

- `2026-09-03 22:13` **identified** — The issue has been identified and a fix is being implemented.
- `2026-09-03 22:31` **identified** — Cluster Metrics are healthy for the last 15 minutes
- `2026-09-03 22:48` **monitoring** — Tempo recovered.
- `2026-09-04 00:55` **resolved** — This incident has been resolved.

### 13. cloudflare — Purchased domains from registrar not appearing in Cloudflare Dashboard

- URL: https://www.cloudflarestatus.com/incidents/1ghtnh94pby6
- Started: 2026-09-06 17:06 UTC
- Resolved: 2026-09-06 21:50 UTC
- Duration: 4h 43m
- Impact: minor, status: resolved
- Components: Registrar
- Canonical incident: [ci-459646994d184708b9977b47](#ci-459646994d184708b9977b47)

- `2026-09-06 17:06` **investigating** — We are investigating reports of delays in the provisioning of newly purchased domains. Our engineering team is currently working to resolve these delays affecting domain registration.
- `2026-09-06 21:50` **resolved** — This incident has been resolved.

### 14. cloudflare — Workers Builds elevated queue times

- URL: https://www.cloudflarestatus.com/incidents/48wc4y9p1vym
- Started: 2026-09-03 23:35 UTC
- Resolved: 2026-09-04 04:09 UTC
- Duration: 4h 33m
- Impact: minor, status: resolved
- Components: Workers Builds
- Canonical incident: [ci-378316fc56173d6c4b79caea](#ci-378316fc56173d6c4b79caea)

- `2026-09-03 23:35` **investigating** — Cloudflare is aware of and investigating an issue with delays running Cloudflare Workers builds which potentially impacts multiple customers.
- `2026-09-03 23:43` **identified** — The issue has been identified and a fix is being implemented.
- `2026-09-04 03:56` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-09-04 04:09` **resolved** — This incident has been resolved.

### 15. openai — ChatGPT Work seeing elevated errors and latency

- URL: https://status.openai.com/incidents/01M1C5M4K0WC8PPT0Z175RJA1E
- Started: 2026-08-31 15:04 UTC
- Resolved: 2026-08-31 20:28 UTC
- Duration: 5h 23m
- Impact: major, status: resolved
- Components: —
- Canonical incident: [ci-46a6010020a7471ec0d2af52](#ci-46a6010020a7471ec0d2af52)

- `2026-08-31 15:04` **identified** — We have identified that users are experiencing elevated latency for the impacted services. We are working on implementing a mitigation.
- `2026-08-31 15:21` **identified** — We have identified that users are experiencing elevated errors for the impacted services. Plus users are particularly impacted, where Work mode is currently unavailable. We are working on implementing a mitigation.
- `2026-08-31 16:02` **identified** — User across multiple subscription plans may be unable to start or continue tasks in ChatGPT Work. We are continuing to work on a mitigation.
- `2026-08-31 17:01` **identified** — User across multiple subscription plans may be unable to start or continue tasks in ChatGPT Work. We We are continuing work on implementing a mitigation.
- `2026-08-31 17:52` **identified** — User across multiple subscription plans may be unable to start or continue tasks in ChatGPT Work. We are continuing work on implementing a mitigation.
- `2026-08-31 19:29` **identified** — User across multiple subscription plans may be unable to start or continue tasks in ChatGPT Work. We are continuing work on implementing a mitigation.
- `2026-08-31 20:01` **monitoring** — We have applied a mitigation and are monitoring recovery.
- `2026-08-31 20:28` **resolved** — All impacted services have now fully recovered.

### 16. cloudflare — Workers KV experiencing elevated error rates in Western Europe Region

- URL: https://www.cloudflarestatus.com/incidents/kby271yt9vwx
- Started: 2026-08-31 11:01 UTC
- Resolved: 2026-08-31 15:04 UTC
- Duration: 4h 2m
- Impact: minor, status: resolved
- Components: Workers KV
- Canonical incident: [ci-adfb5f33512e06d00e6e16d1](#ci-adfb5f33512e06d00e6e16d1)

- `2026-08-31 11:02` **investigating** — We are investigating an issue affecting Workers KV in the Western Europe Region, which may result in elevated error rates for read requests. Writes are not affected. Our engineering team is currently working to identify and resolve the root cause. We will provide updates as more information becomes available.
- `2026-08-31 12:40` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-08-31 12:40` **monitoring** — We are continuing to monitor for any further issues.
- `2026-08-31 15:04` **resolved** — This incident has been resolved.

### 17. cloudflare — Elevated number of R2 503 errors in Eastern North America region

- URL: https://www.cloudflarestatus.com/incidents/ftvf8c3m4mv5
- Started: 2026-09-04 21:10 UTC
- Resolved: 2026-09-05 00:42 UTC
- Duration: 3h 31m
- Impact: minor, status: resolved
- Components: R2
- Canonical incident: [ci-c622743dec6d1dec6f33f398](#ci-c622743dec6d1dec6f33f398)

- `2026-09-04 21:10` **investigating** — Cloudflare is investigating issues with R2 buckets in Eastern North America.
- `2026-09-04 21:27` **identified** — The cause of this issue has been identified and a fix is being implemented.
- `2026-09-05 00:13` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-09-05 00:42` **resolved** — This incident has been resolved.

### 18. grafana — Alert rule creation, deletion, and update degradation in prod-us-east-2

- URL: https://stspg.io/8gw4nt4mmp09
- Started: 2026-09-04 13:36 UTC
- Resolved: 2026-09-04 16:32 UTC
- Duration: 2h 56m
- Impact: minor, status: resolved
- Components: AWS US East (OH) - prod-us-east-0: Alertmanager, AWS US East (OH) - prod-us-east-0: Alertmanager and Rules Configuration API
- Canonical incident: [ci-1320404fc71af98183f97366](#ci-1320404fc71af98183f97366)

- `2026-09-04 13:36` **investigating** — We’re currently investigating an issue affecting alert rule creation, deletion, and updates in prod-us-east-2. Our team is actively working to identify the cause. Thank you for your patience.
- `2026-09-04 16:32` **resolved** — The issue has been mitigated and there's a code fix pending.

### 19. cloudflare — Cloudflare Access bypass IP service experiencing availability issues

- URL: https://www.cloudflarestatus.com/incidents/fcdl2yb3q7qw
- Started: 2026-08-31 11:40 UTC
- Resolved: 2026-08-31 14:50 UTC
- Duration: 3h 10m
- Impact: minor, status: resolved
- Components: Access
- Canonical incident: [ci-42f150177c196c3833cd05c3](#ci-42f150177c196c3833cd05c3)

- `2026-08-31 11:40` **investigating** — We are investigating an issue affecting the Access IP List service, which may result in failures when bypassing users whose IPs are listed in them. Our engineering team is currently working to identify and resolve the root cause. We will provide updates as more information becomes available.
- `2026-08-31 13:08` **identified** — The issue has been identified and a fix is being implemented.
- `2026-08-31 14:50` **resolved** — This incident has been resolved.

### 20. datadog — Delayed CICD Optimization, Code Coverage, Code Security, and DORA data

- URL: https://stspg.io/zstf1zx7hfy3
- Started: 2026-09-03 17:20 UTC
- Resolved: 2026-09-03 23:01 UTC
- Duration: 5h 40m
- Impact: minor, status: resolved
- Components: CI Visibility
- Canonical incident: [ci-5eebb41b9c34c3b2a26ea1be](#ci-5eebb41b9c34c3b2a26ea1be)

- `2026-09-03 17:20` **investigating** — We are investigating increased latency processing CICD Optimization, Code Coverage, Code Security, and DORA data. As a result of this issue, some users may see delays with test runs and pipeline executions.
- `2026-09-03 17:27` **identified** — We identified the issue.
- `2026-09-03 17:45` **monitoring** — We are monitoring the issue.
- `2026-09-03 17:59` **identified** — We identified the issue.
- `2026-09-03 19:01` **monitoring** — We are monitoring the issue.
- `2026-09-03 21:51` **monitoring** — We have applied remediations, and recovery is in progress. We are continuing to monitor the issue.
- `2026-09-03 23:01` **resolved** — The issue is now resolved.

### 21. openai — Users in APAC region may face increased error in ChatGPT, Work, image generation, file upload, Voice, and Codex Cloud

- URL: https://status.openai.com/incidents/01M1NKFZH5EEYEREC54HNAHY35
- Started: 2026-09-04 07:00 UTC
- Resolved: 2026-09-04 10:46 UTC
- Duration: 3h 46m
- Impact: minor, status: resolved
- Components: —
- Canonical incident: [ci-8b787f6cd11e7586b127c5b2](#ci-8b787f6cd11e7586b127c5b2)

- `2026-09-04 07:00` **investigating** — We are investigating the issue for the listed services.
- `2026-09-04 09:47` **monitoring** — We have applied the mitigation and are monitoring the recovery.
- `2026-09-04 10:46` **resolved** — All impacted services have now fully recovered.

### 22. grafana — Investigating issues in US Central (prod-us-central-0, prod-us-central-5)

- URL: https://stspg.io/jvcf7d0j5x37
- Started: 2026-09-01 15:25 UTC
- Resolved: 2026-09-01 20:20 UTC
- Duration: 4h 54m
- Impact: minor, status: resolved
- Components: —
- Canonical incident: [ci-9f2faa3a5009d8624d790189](#ci-9f2faa3a5009d8624d790189)

- `2026-09-01 15:25` **investigating** — We are investigating reports of errors affecting Grafana Cloud in US Central (prod-us-central-0, prod-us-central-5). Customers in this region may see timeouts or failed requests when writing or querying logs and metrics. We are looking into this and will post another update as soon as we have more information.
- `2026-09-01 16:14` **identified** — We have identified the cause as an issue in our cloud provider’s US Central region. A mitigation is in place and recovery is in progress. Customers in prod-us-central-0 and prod-us-central-5 may still see errors or delays with metrics, logs, or Grafana while services come back. We will update this page as recovery continues.
- `2026-09-01 20:20` **resolved** — This incident has been resolved.

### 23. aws — Increased API Error Rates

- URL: https://health.aws.amazon.com/health/status
- Started: 2026-09-03 21:49 UTC
- Resolved: 2026-09-04 00:07 UTC
- Duration: 2h 17m
- Impact: unknown, status: resolved
- Components: multipleservices-us-west-2
- Canonical incident: [ci-e8e8f6a77b376e83f25d7cf4](#ci-e8e8f6a77b376e83f25d7cf4)

- `2026-09-03 21:49` **investigating** — Starting at 2:12 PM PDT, we began experiencing increased API error rates for STS and Sign In when using SAML in the US-WEST-2 Region. Our engineering team was automatically engaged at 2:19 PM to begin investigating the root cause. There is no work around available at this time. We will provide another update by 3:30 PM PDT.
- `2026-09-03 22:26` **investigating** — We are seeing early signs of recovery and continue to monitor for full recovery. We will provide another update at 4:15 PM, or sooner if we have additional information to share.
- `2026-09-03 23:20` **investigating** — We continue to see recovery holding steady for the STS AssumeRoleWithSAML and AssumeRoleWithWebIdentity APIs in the US-WEST-2 Region. Error rates are now back to pre-event levels, and we continue actively monitoring to confirm full recovery. We will provide another update by 5:15 PM or earlier.
- `2026-09-04 00:07` **resolved** — Between 2:12 PM and 3:18 PM PDT, we experienced increased API error rates affecting the STS AssumeRoleWithSAML and AssumeRoleWithWebIdentity APIs in the US-WEST-2 Region. The root cause was determined to be due to an issue with an STS subsystem responsible for communicating with external identity providers. Other AWS Services that rely on these identity federation protocols were also affected. At 3:18 PM, we observed signs of recovery and continued to monitor to ensure stability and full recovery. The issue is resolved and the service is operating normally at this time.

### 24. github — Incident with Grok Copilot AI Model Provider

- URL: https://stspg.io/p217h6l54208
- Started: 2026-09-03 14:17 UTC
- Resolved: 2026-09-03 17:11 UTC
- Duration: 2h 54m
- Impact: minor, status: resolved
- Components: Copilot AI Model Providers
- Canonical incident: [ci-f7215da25777c0cc42667b87](#ci-f7215da25777c0cc42667b87)

- `2026-09-03 14:17` **investigating** — We are investigating reports of degraded performance for Copilot AI Model Providers
- `2026-09-03 14:20` **investigating** — We are experiencing degraded availability for the Grok 4.6 model in Copilot Chat, VS Code and other Copilot products. This is due to an issue with an upstream model provider. We are working with them to resolve the issue.<br />
- `2026-09-03 14:22` **investigating** — The Grok 4.5 model has degraded availability as well. We are working with the upstream provider to resolve the issue.
- `2026-09-03 17:11` **investigating** — The issues with our upstream model provider have been resolved, and Grok models are once again available in Copilot products and IDE surfaces.<br />We will continue monitoring to ensure stability, but mitigation is complete.
- `2026-09-03 17:11` **resolved** — Between 13:22 and 17:11 UTC on September 03, 2026, GitHub Copilot experienced degradation affecting several Grok models, including Grok 4.5 and Grok 4.6. Users encountered elevated error rates, but other models were not affected. The degradation was caused by an issue with an upstream model provider. GitHub engineers detected the issue through automated monitoring, displayed in-product warnings for the affected models, and coordinated with the provider. Service returned to normal after the provider implemented a mitigation.

### 25. yandexcloud — Partial unavailability of the Cloud Backup service

- URL: https://status.yandex.cloud/en/incidents/1993
- Started: 2026-09-02 07:00 UTC
- Resolved: 2026-09-02 09:51 UTC
- Duration: 2h 51m
- Impact: minor, status: resolved
- Components: Cloud Backup
- Canonical incident: [ci-d673566ad5de60511cf14751](#ci-d673566ad5de60511cf14751)

- `2026-09-02 06:51` **unknown** — From 10:00 to 13:00 Moscow time, the service will be operated by the backup provider. At this point, it can be observed: – Unavailability of operations for connecting virtual machines to the service, management of backup and backup policies through the user interface, CLI and API for up to 60 minutes; – There may be problems in starting and executing backup tasks, as well as a drop in the speed of creating and restoring backups.
- `2026-09-02 11:40` **resolved** — The provider's work has been completed, and the service is operating normally.

### 26. cloudflare — Durable Objects increased errors in Western North America

- URL: https://www.cloudflarestatus.com/incidents/3qs7m80p13v2
- Started: 2026-09-02 13:43 UTC
- Resolved: 2026-09-02 15:51 UTC
- Duration: 2h 7m
- Impact: minor, status: resolved
- Components: Durable Objects
- Canonical incident: [ci-6cc8b0ff4e4446877ca31ce6](#ci-6cc8b0ff4e4446877ca31ce6)

- `2026-09-02 14:47` **investigating** — Cloudflare is aware of, and investigating, an increased number of errors for Durable Objects in Western North America. The suspected cause is a reduction in routing capacity due to a fibre cut. We are working to mitigate this issue, and update will follow.
- `2026-09-02 14:48` **identified** — The issue has been identified and a fix is being implemented.
- `2026-09-02 15:35` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-09-02 15:51` **resolved** — This incident has been resolved.

### 27. cloudflare — Cache Purging Errors

- URL: https://www.cloudflarestatus.com/incidents/ff0yvx1d6tby
- Started: 2026-09-04 14:51 UTC
- Resolved: 2026-09-04 16:58 UTC
- Duration: 2h 7m
- Impact: minor, status: resolved
- Components: CDN Cache Purge
- Canonical incident: [ci-5b3a5d907037f65fac7337c7](#ci-5b3a5d907037f65fac7337c7)

- `2026-09-04 14:51` **investigating** — Cloudflare cache purging called via the Dashboard or API is returning errors, with the message "unable to purge. internal error". We are working to resolve this issue as soon as possible. More updates to follow.
- `2026-09-04 15:31` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-09-04 16:58` **resolved** — This incident is now resolved.

### 28. github — Disruption with Copilot Code Review

- URL: https://stspg.io/dddfhld6fj0z
- Started: 2026-09-04 20:39 UTC
- Resolved: 2026-09-04 22:26 UTC
- Duration: 1h 47m
- Impact: major, status: resolved
- Components: —
- Canonical incident: [ci-5bf23adefa31d8cff78965a5](#ci-5bf23adefa31d8cff78965a5)

- `2026-09-04 20:39` **investigating** — We are investigating reports of impacted performance for some GitHub services.
- `2026-09-04 20:57` **investigating** — Some users may be experiencing failures when using Copilot code review. We have identified the root cause and are working on a mitigation.
- `2026-09-04 21:54` **investigating** — We are applying the mitigation and expect recovery within approximately 30 minutes.
- `2026-09-04 22:25` **monitoring** — The degradation has been mitigated. We are monitoring to ensure stability.
- `2026-09-04 22:26` **resolved** — On September 4, 2026, between 20:04 and 22:26 UTC, GitHub Copilot code review experienced an increased failure rate. Affected pull request reviews failed to complete or post review comments.<br /><br />The incident was caused by a change to the service’s authentication permissions that prevented it from submitting affected reviews to the GitHub API. We reverted the change and restored normal operation by 22:26 UTC.<br /><br />We apologize for the disruption.

### 29. anthropic — Elevated errors for Claude Sonnet 5

- URL: https://stspg.io/9dgzwz3ymz42
- Started: 2026-09-02 21:17 UTC
- Resolved: 2026-09-02 21:44 UTC
- Duration: 26m
- Impact: major, status: resolved
- Components: Claude API (api.anthropic.com), Claude Code, Claude Cowork, claude.ai
- Canonical incident: [ci-c3ec83ce59c8bc0a05e2aeb6](#ci-c3ec83ce59c8bc0a05e2aeb6)

- `2026-09-02 21:17` **investigating** — We are investigating elevated errors on requests to Claude Sonnet 5. We will provide an update as soon as possible.
- `2026-09-02 21:44` **resolved** — The issue affecting Claude Sonnet 5 has been resolved. Impact occurred from 2:05pm PT / 21:05 UTC to 2:19pm PT / 21:19 UTC.

### 30. anthropic — Elevated errors for Claude Sonnet 5

- URL: https://stspg.io/jt14jk41pjsf
- Started: 2026-09-03 12:37 UTC
- Resolved: 2026-09-03 12:56 UTC
- Duration: 18m
- Impact: major, status: resolved
- Components: Claude API (api.anthropic.com), Claude Code, Claude Cowork, claude.ai
- Canonical incident: [ci-c908724cf5da881bc4680624](#ci-c908724cf5da881bc4680624)

- `2026-09-03 12:37` **investigating** — We are investigating elevated errors on requests to Claude Sonnet 5. We will provide an update as soon as possible.
- `2026-09-03 12:47` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-09-03 12:56` **resolved** — This incident has been resolved.

### 31. cloudflare — Network Route Leak in Palmas, Brazil

- URL: https://www.cloudflarestatus.com/incidents/1w42pbthkqjf
- Started: 2026-09-01 21:47 UTC
- Resolved: 2026-09-01 22:05 UTC
- Duration: 18m
- Impact: major, status: resolved
- Components: Palmas, Brazil - (PMW)
- Canonical incident: [ci-26627d577d5c35c53f8cb596](#ci-26627d577d5c35c53f8cb596)

- `2026-09-01 21:47` **monitoring** — One of Cloudflare's peers experienced a route leak, misdirecting some traffic to Palmas, Brazil. We are working to mitigate impact to Internet users in this region.
- `2026-09-01 22:05` **resolved** — This incident has been resolved.

### 32. openai — ChatGPT Work Mode High Error Rates

- URL: https://status.openai.com/incidents/01M1J99EDCYED5GP2W8GZ3N2NA
- Started: 2026-09-03 00:04 UTC
- Resolved: 2026-09-03 00:10 UTC
- Duration: 6m
- Impact: major, status: resolved
- Components: —
- Canonical incident: [ci-c2548008c4f3a3f8f8292e73](#ci-c2548008c4f3a3f8f8292e73)

- `2026-09-03 00:04` **investigating** — ChatGPT Work Mode is experiencing high error rates. We are investigating the issue.
- `2026-09-03 00:10` **resolved** — All impacted services have now fully recovered.

## Below threshold (24)

- `cloudflare` [Service Connectivity Issues in Canberra](https://www.cloudflarestatus.com/incidents/htwds1qczkzd) — 1h 50m
- `cloudflare` [Cloudflare Workers Issues in Narita](https://www.cloudflarestatus.com/incidents/lmsfz6pcb65c) — 1h 21m
- `anthropic` [Degraded performance on claude.ai and Claude Code](https://stspg.io/gny7dn3k231b) — 1h 14m
- `cloudflare` [Cloudflare Access one-time PIN emails blocked by certain email security gateways, including Proofpoint](https://www.cloudflarestatus.com/incidents/72lc3n642lkf) — 1h 46m
- `anthropic` [Degraded performance on platform.claude.com and Claude for Microsoft Office 365](https://stspg.io/4rf459q4pjc3) — 1h 1m
- `anthropic` [Delays in credit purchases](https://stspg.io/jc510qy51bfh) — 1h 57m
- `openai` [Elevated errors across ChatGPT and Codex](https://status.openai.com/incidents/01M1KWEDH417T2CF44YYHZDFCR) — 1h 57m
- `cloudflare` [Increased HTTP 522 Errors in Seattle (SEA)](https://www.cloudflarestatus.com/incidents/vp04n6xtv51r) — 1h 0m
- `circleci` [macOS Job Delays](https://stspg.io/l8rtj1qhr625) — 1h 39m
- `github` [Delays in commit processing](https://stspg.io/0xbhhq84v4mt) — 1h 0m
- `github` [Elevated rate of errors for OpenAI models provided by Copilot](https://stspg.io/8trzn5bjrtnn) — 42m
- `cloudflare` [Elevated number of R2 503 errors in Western North America region](https://www.cloudflarestatus.com/incidents/y27rg9kcyryn) — 29m
- `openai` [Elevated errors creating new accounts](https://status.openai.com/incidents/01M1HN7GRR6FJNWFVX2F8QMA3R) — 31m
- `anthropic` [Degraded performance on claude.ai](https://stspg.io/hmpwykr9h567) — 29m
- `openai` [Elevated errors in ChatGPT conversations for Free and Go plans](https://status.openai.com/incidents/01M1D6E14BQSYQA3R5CNQBB7EB) — 21m
- `github` [Degradation in repos contents API](https://stspg.io/vhmsw2rhdcgv) — 21m
- `anthropic` [Degraded performance on Claude for Microsoft Office 365](https://stspg.io/k8bzm26b6sy9) — 19m
- `cloudflare` [Increased HTTP 5xx Errors in Hong Kong.](https://www.cloudflarestatus.com/incidents/hv0y80k98t9b) — 14m
- `cloudflare` [Zero Trust Dashboard - Posture Check Visibility Issue](https://www.cloudflarestatus.com/incidents/6ztvhhp2ll11) — 12m
- `datadog` [Delayed DBM Monitor Evaluation](https://stspg.io/2lq8c212gd6p) — 17m
- `cloudflare` [Increased latency in Taipei](https://www.cloudflarestatus.com/incidents/vthcyk0986fd) — 0m
- `grafana` [Elevated Latency in Grafana Cloud Logs (prod-ap-south-1)](https://stspg.io/39bjpwxypjqn) — 0m
- `cloudflare` [Elevated number of R2 503 errors in Western North America region](https://www.cloudflarestatus.com/incidents/pzw3kysls711) — 0m
- `cloudflare` [Elevated connectivity degradation in NRT](https://www.cloudflarestatus.com/incidents/bgbz898hh98g) — 0m

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-5650bfa3e91cb26e3635de5e"></a>

- Canonical ID: `ci-5650bfa3e91cb26e3635de5e` (1 source records).
  - [gcp/J5ia5t9p3g9Q5Wi7r8Ev](../../data/2026/W36/gcp.json); vendor URL: https://status.cloud.google.com/incidents/J5ia5t9p3g9Q5Wi7r8Ev
- Start time basis: `vendor_reported`.

<a id="ci-93c7789a5f6839ae3a7974bb"></a>

- Canonical ID: `ci-93c7789a5f6839ae3a7974bb` (1 source records).
  - [cloudflare/8x66bpk6p9kk](../../data/2026/W36/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/8x66bpk6p9kk
- Start time basis: `vendor_reported`.

<a id="ci-70f9b5947df78d16177cefc9"></a>

- Canonical ID: `ci-70f9b5947df78d16177cefc9` (1 source records).
  - [grafana/k5c91vq67f8f](../../data/2026/W36/grafana.json); vendor URL: https://stspg.io/b3b4ppnkx23p
- Start time basis: `vendor_reported`.

<a id="ci-f78b9503c3eb6ba7e99b3bcd"></a>

- Canonical ID: `ci-f78b9503c3eb6ba7e99b3bcd` (1 source records).
  - [openai/01M1CYZDD84XD0H56XNEV2PT98](../../data/2026/W36/openai.json); vendor URL: https://status.openai.com/incidents/01M1CYZDD84XD0H56XNEV2PT98
- Start time basis: `vendor_reported`.

<a id="ci-2b37818ff156857c8838818e"></a>

- Canonical ID: `ci-2b37818ff156857c8838818e` (1 source records).
  - [yandexcloud/1992](../../data/2026/W36/yandexcloud.json); vendor URL: https://status.yandex.cloud/en/incidents/1992
- Start time basis: `vendor_reported`.

<a id="ci-8c18817154780b777347ab6e"></a>

- Canonical ID: `ci-8c18817154780b777347ab6e` (1 source records).
  - [cloudflare/cczd80189gvc](../../data/2026/W36/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/cczd80189gvc
- Start time basis: `vendor_reported`.

<a id="ci-4f3dd69d94dd51ba12a21465"></a>

- Canonical ID: `ci-4f3dd69d94dd51ba12a21465` (1 source records).
  - [cloudflare/w6ng4jcxv7m1](../../data/2026/W36/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/w6ng4jcxv7m1
- Start time basis: `vendor_reported`.

<a id="ci-d842d706e233d509f3083a67"></a>

- Canonical ID: `ci-d842d706e233d509f3083a67` (1 source records).
  - [anthropic/461yvfrzpwtt](../../data/2026/W36/anthropic.json); vendor URL: https://stspg.io/9xz4hhmd1jzn
- Start time basis: `vendor_reported`.

<a id="ci-769082f75e1a72f2fcc729f6"></a>

- Canonical ID: `ci-769082f75e1a72f2fcc729f6` (1 source records).
  - [grafana/mdbl0xl76v3q](../../data/2026/W36/grafana.json); vendor URL: https://stspg.io/f6h7k3lq3v62
- Start time basis: `vendor_reported`.

<a id="ci-a96eeb9ccce075e6b5a733b3"></a>

- Canonical ID: `ci-a96eeb9ccce075e6b5a733b3` (1 source records).
  - [cloudflare/3mjmdcmm86nx](../../data/2026/W36/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/3mjmdcmm86nx
- Start time basis: `vendor_reported`.

<a id="ci-ac11f1a6fd089e9ddb1b326c"></a>

- Canonical ID: `ci-ac11f1a6fd089e9ddb1b326c` (1 source records).
  - [grafana/w1c8md3mnthx](../../data/2026/W36/grafana.json); vendor URL: https://stspg.io/7n6x5qw9x9d5
- Start time basis: `vendor_reported`.

<a id="ci-a420a3ad5b4504b9b5c08b7e"></a>

- Canonical ID: `ci-a420a3ad5b4504b9b5c08b7e` (1 source records).
  - [grafana/lv9wsn3ksb79](../../data/2026/W36/grafana.json); vendor URL: https://stspg.io/kw7xrpjrg134
- Start time basis: `vendor_reported`.

<a id="ci-459646994d184708b9977b47"></a>

- Canonical ID: `ci-459646994d184708b9977b47` (1 source records).
  - [cloudflare/1ghtnh94pby6](../../data/2026/W36/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/1ghtnh94pby6
- Start time basis: `vendor_reported`.

<a id="ci-378316fc56173d6c4b79caea"></a>

- Canonical ID: `ci-378316fc56173d6c4b79caea` (1 source records).
  - [cloudflare/48wc4y9p1vym](../../data/2026/W36/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/48wc4y9p1vym
- Start time basis: `vendor_reported`.

<a id="ci-46a6010020a7471ec0d2af52"></a>

- Canonical ID: `ci-46a6010020a7471ec0d2af52` (1 source records).
  - [openai/01M1C5M4K0WC8PPT0Z175RJA1E](../../data/2026/W36/openai.json); vendor URL: https://status.openai.com/incidents/01M1C5M4K0WC8PPT0Z175RJA1E
- Start time basis: `vendor_reported`.

<a id="ci-adfb5f33512e06d00e6e16d1"></a>

- Canonical ID: `ci-adfb5f33512e06d00e6e16d1` (1 source records).
  - [cloudflare/kby271yt9vwx](../../data/2026/W36/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/kby271yt9vwx
- Start time basis: `vendor_reported`.

<a id="ci-7c8c3a4145969d08fdc1b8cd"></a>

- Canonical ID: `ci-7c8c3a4145969d08fdc1b8cd` (1 source records).
  - [cloudflare/htwds1qczkzd](../../data/2026/W36/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/htwds1qczkzd
- Start time basis: `vendor_reported`.

<a id="ci-c622743dec6d1dec6f33f398"></a>

- Canonical ID: `ci-c622743dec6d1dec6f33f398` (1 source records).
  - [cloudflare/ftvf8c3m4mv5](../../data/2026/W36/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/ftvf8c3m4mv5
- Start time basis: `vendor_reported`.

<a id="ci-1320404fc71af98183f97366"></a>

- Canonical ID: `ci-1320404fc71af98183f97366` (1 source records).
  - [grafana/5847d1tvtkdd](../../data/2026/W36/grafana.json); vendor URL: https://stspg.io/8gw4nt4mmp09
- Start time basis: `vendor_reported`.

<a id="ci-42f150177c196c3833cd05c3"></a>

- Canonical ID: `ci-42f150177c196c3833cd05c3` (1 source records).
  - [cloudflare/fcdl2yb3q7qw](../../data/2026/W36/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/fcdl2yb3q7qw
- Start time basis: `vendor_reported`.

<a id="ci-5eebb41b9c34c3b2a26ea1be"></a>

- Canonical ID: `ci-5eebb41b9c34c3b2a26ea1be` (1 source records).
  - [datadog/4zd1rbp4pbt5](../../data/2026/W36/datadog.json); vendor URL: https://stspg.io/zstf1zx7hfy3
- Start time basis: `vendor_reported`.

<a id="ci-8b787f6cd11e7586b127c5b2"></a>

- Canonical ID: `ci-8b787f6cd11e7586b127c5b2` (1 source records).
  - [openai/01M1NKFZH5EEYEREC54HNAHY35](../../data/2026/W36/openai.json); vendor URL: https://status.openai.com/incidents/01M1NKFZH5EEYEREC54HNAHY35
- Start time basis: `vendor_reported`.

<a id="ci-9f2faa3a5009d8624d790189"></a>

- Canonical ID: `ci-9f2faa3a5009d8624d790189` (1 source records).
  - [grafana/328b9jvhlwwf](../../data/2026/W36/grafana.json); vendor URL: https://stspg.io/jvcf7d0j5x37
- Start time basis: `vendor_reported`.

<a id="ci-daee8aab0ca1e3dc7030097c"></a>

- Canonical ID: `ci-daee8aab0ca1e3dc7030097c` (1 source records).
  - [cloudflare/lmsfz6pcb65c](../../data/2026/W36/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/lmsfz6pcb65c
- Start time basis: `vendor_reported`.

<a id="ci-e8e8f6a77b376e83f25d7cf4"></a>

- Canonical ID: `ci-e8e8f6a77b376e83f25d7cf4` (1 source records).
  - [aws/151d07b0bd5e8d08](../../data/2026/W36/aws.json); vendor URL: https://health.aws.amazon.com/health/status
- Start time basis: `first_observed_update`.
- Warning: AWS onset is the first observed message, not a verified impact start

<a id="ci-f7215da25777c0cc42667b87"></a>

- Canonical ID: `ci-f7215da25777c0cc42667b87` (1 source records).
  - [github/ktdr5t0xwnhp](../../data/2026/W36/github.json); vendor URL: https://stspg.io/p217h6l54208
- Start time basis: `vendor_reported`.

<a id="ci-d673566ad5de60511cf14751"></a>

- Canonical ID: `ci-d673566ad5de60511cf14751` (1 source records).
  - [yandexcloud/1993](../../data/2026/W36/yandexcloud.json); vendor URL: https://status.yandex.cloud/en/incidents/1993
- Start time basis: `vendor_reported`.

<a id="ci-6cc8b0ff4e4446877ca31ce6"></a>

- Canonical ID: `ci-6cc8b0ff4e4446877ca31ce6` (1 source records).
  - [cloudflare/3qs7m80p13v2](../../data/2026/W36/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/3qs7m80p13v2
- Start time basis: `vendor_reported`.

<a id="ci-5b3a5d907037f65fac7337c7"></a>

- Canonical ID: `ci-5b3a5d907037f65fac7337c7` (1 source records).
  - [cloudflare/ff0yvx1d6tby](../../data/2026/W36/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/ff0yvx1d6tby
- Start time basis: `vendor_reported`.

<a id="ci-4f0572a67a1300d68a970b8a"></a>

- Canonical ID: `ci-4f0572a67a1300d68a970b8a` (1 source records).
  - [anthropic/r82kdk0m7vqh](../../data/2026/W36/anthropic.json); vendor URL: https://stspg.io/gny7dn3k231b
- Start time basis: `vendor_reported`.

<a id="ci-a99357246e887d2a39d9d656"></a>

- Canonical ID: `ci-a99357246e887d2a39d9d656` (1 source records).
  - [cloudflare/72lc3n642lkf](../../data/2026/W36/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/72lc3n642lkf
- Start time basis: `vendor_reported`.

<a id="ci-c99e5327bdbec2413d3b3878"></a>

- Canonical ID: `ci-c99e5327bdbec2413d3b3878` (1 source records).
  - [anthropic/3g3d55q6vk3h](../../data/2026/W36/anthropic.json); vendor URL: https://stspg.io/4rf459q4pjc3
- Start time basis: `vendor_reported`.

<a id="ci-cdbdf5929644c1df3d2d3847"></a>

- Canonical ID: `ci-cdbdf5929644c1df3d2d3847` (1 source records).
  - [anthropic/620swtqyn24k](../../data/2026/W36/anthropic.json); vendor URL: https://stspg.io/jc510qy51bfh
- Start time basis: `vendor_reported`.

<a id="ci-46ebfc228d1ac2bc83913366"></a>

- Canonical ID: `ci-46ebfc228d1ac2bc83913366` (1 source records).
  - [openai/01M1KWEDH417T2CF44YYHZDFCR](../../data/2026/W36/openai.json); vendor URL: https://status.openai.com/incidents/01M1KWEDH417T2CF44YYHZDFCR
- Start time basis: `vendor_reported`.

<a id="ci-5bf23adefa31d8cff78965a5"></a>

- Canonical ID: `ci-5bf23adefa31d8cff78965a5` (1 source records).
  - [github/zw1hbx2yyhvr](../../data/2026/W36/github.json); vendor URL: https://stspg.io/dddfhld6fj0z
- Start time basis: `vendor_reported`.

<a id="ci-c3ec83ce59c8bc0a05e2aeb6"></a>

- Canonical ID: `ci-c3ec83ce59c8bc0a05e2aeb6` (1 source records).
  - [anthropic/ls6bn1x81m0w](../../data/2026/W36/anthropic.json); vendor URL: https://stspg.io/9dgzwz3ymz42
- Start time basis: `vendor_reported`.

<a id="ci-99f8e82e4af55f1f5b86887c"></a>

- Canonical ID: `ci-99f8e82e4af55f1f5b86887c` (1 source records).
  - [cloudflare/vp04n6xtv51r](../../data/2026/W36/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/vp04n6xtv51r
- Start time basis: `vendor_reported`.

<a id="ci-c908724cf5da881bc4680624"></a>

- Canonical ID: `ci-c908724cf5da881bc4680624` (1 source records).
  - [anthropic/288w7p4hk1l1](../../data/2026/W36/anthropic.json); vendor URL: https://stspg.io/jt14jk41pjsf
- Start time basis: `vendor_reported`.

<a id="ci-dff6abdcb98527cc9b309b38"></a>

- Canonical ID: `ci-dff6abdcb98527cc9b309b38` (1 source records).
  - [circleci/wl4yqd96zl2h](../../data/2026/W36/circleci.json); vendor URL: https://stspg.io/l8rtj1qhr625
- Start time basis: `vendor_reported`.

<a id="ci-77810a0da696ab3c4cdd052f"></a>

- Canonical ID: `ci-77810a0da696ab3c4cdd052f` (1 source records).
  - [github/jk2dy5h9yp3m](../../data/2026/W36/github.json); vendor URL: https://stspg.io/0xbhhq84v4mt
- Start time basis: `vendor_reported`.

<a id="ci-1e92b4d1274d9b2bba2298b0"></a>

- Canonical ID: `ci-1e92b4d1274d9b2bba2298b0` (1 source records).
  - [github/7fxts6gmq5gr](../../data/2026/W36/github.json); vendor URL: https://stspg.io/8trzn5bjrtnn
- Start time basis: `vendor_reported`.

<a id="ci-7c0992345872a319df49e5c7"></a>

- Canonical ID: `ci-7c0992345872a319df49e5c7` (1 source records).
  - [cloudflare/y27rg9kcyryn](../../data/2026/W36/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/y27rg9kcyryn
- Start time basis: `vendor_reported`.

<a id="ci-0511001f7927cdcd843133ca"></a>

- Canonical ID: `ci-0511001f7927cdcd843133ca` (1 source records).
  - [openai/01M1HN7GRR6FJNWFVX2F8QMA3R](../../data/2026/W36/openai.json); vendor URL: https://status.openai.com/incidents/01M1HN7GRR6FJNWFVX2F8QMA3R
- Start time basis: `vendor_reported`.

<a id="ci-7a99d723667f3037a7acb3bd"></a>

- Canonical ID: `ci-7a99d723667f3037a7acb3bd` (1 source records).
  - [anthropic/9jrp5rtyzrf6](../../data/2026/W36/anthropic.json); vendor URL: https://stspg.io/hmpwykr9h567
- Start time basis: `vendor_reported`.

<a id="ci-26627d577d5c35c53f8cb596"></a>

- Canonical ID: `ci-26627d577d5c35c53f8cb596` (1 source records).
  - [cloudflare/1w42pbthkqjf](../../data/2026/W36/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/1w42pbthkqjf
- Start time basis: `vendor_reported`.

<a id="ci-78bfad5a74d6da60c346ec04"></a>

- Canonical ID: `ci-78bfad5a74d6da60c346ec04` (1 source records).
  - [openai/01M1D6E14BQSYQA3R5CNQBB7EB](../../data/2026/W36/openai.json); vendor URL: https://status.openai.com/incidents/01M1D6E14BQSYQA3R5CNQBB7EB
- Start time basis: `vendor_reported`.

<a id="ci-66bd88d3e21f53ceca59cb62"></a>

- Canonical ID: `ci-66bd88d3e21f53ceca59cb62` (1 source records).
  - [github/r05vk75c594j](../../data/2026/W36/github.json); vendor URL: https://stspg.io/vhmsw2rhdcgv
- Start time basis: `vendor_reported`.

<a id="ci-9577738f2f8bf46405559ef5"></a>

- Canonical ID: `ci-9577738f2f8bf46405559ef5` (1 source records).
  - [anthropic/nr3h7bw8b3k3](../../data/2026/W36/anthropic.json); vendor URL: https://stspg.io/k8bzm26b6sy9
- Start time basis: `vendor_reported`.

<a id="ci-d4f28229058a43c5d8279d5a"></a>

- Canonical ID: `ci-d4f28229058a43c5d8279d5a` (1 source records).
  - [cloudflare/hv0y80k98t9b](../../data/2026/W36/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/hv0y80k98t9b
- Start time basis: `vendor_reported`.

<a id="ci-197010993dc8b4faacf7f3fa"></a>

- Canonical ID: `ci-197010993dc8b4faacf7f3fa` (1 source records).
  - [cloudflare/6ztvhhp2ll11](../../data/2026/W36/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/6ztvhhp2ll11
- Start time basis: `vendor_reported`.

<a id="ci-3cdf35ef2b1999edd1d024f5"></a>

- Canonical ID: `ci-3cdf35ef2b1999edd1d024f5` (1 source records).
  - [datadog/6k7fhxz0q7k1](../../data/2026/W36/datadog.json); vendor URL: https://stspg.io/2lq8c212gd6p
- Start time basis: `vendor_reported`.

<a id="ci-c2548008c4f3a3f8f8292e73"></a>

- Canonical ID: `ci-c2548008c4f3a3f8f8292e73` (1 source records).
  - [openai/01M1J99EDCYED5GP2W8GZ3N2NA](../../data/2026/W36/openai.json); vendor URL: https://status.openai.com/incidents/01M1J99EDCYED5GP2W8GZ3N2NA
- Start time basis: `vendor_reported`.

<a id="ci-54f70e687a1f4dce9ee6d909"></a>

- Canonical ID: `ci-54f70e687a1f4dce9ee6d909` (1 source records).
  - [cloudflare/vthcyk0986fd](../../data/2026/W36/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/vthcyk0986fd
- Start time basis: `vendor_reported`.

<a id="ci-7cc1bd65c92da3fc9f72f62f"></a>

- Canonical ID: `ci-7cc1bd65c92da3fc9f72f62f` (1 source records).
  - [grafana/486pg31dzsl1](../../data/2026/W36/grafana.json); vendor URL: https://stspg.io/39bjpwxypjqn
- Start time basis: `vendor_reported`.

<a id="ci-d1f747f9b30a8e861b8c662a"></a>

- Canonical ID: `ci-d1f747f9b30a8e861b8c662a` (1 source records).
  - [cloudflare/pzw3kysls711](../../data/2026/W36/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/pzw3kysls711
- Start time basis: `vendor_reported`.

<a id="ci-f12f9766ce53e33cf7a31b66"></a>

- Canonical ID: `ci-f12f9766ce53e33cf7a31b66` (1 source records).
  - [cloudflare/bgbz898hh98g](../../data/2026/W36/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/bgbz898hh98g
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
