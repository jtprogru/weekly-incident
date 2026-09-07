# Week 2026-W36 — top 5 of 30 above threshold

2026-08-31 to 2026-09-06 (UTC).

## 1. cloudflare — HTTP/3 issue affecting R2 custom domains

- URL: https://www.cloudflarestatus.com/incidents/8x66bpk6p9kk
- Window: 2026-08-31 18:47 to 2026-09-03 19:32 UTC (3d 0h 45m)
- Impact: minor, status: resolved
- Components: R2

  - `2026-08-31 18:47` investigating — Some Firefox users may see delayed or failed asset loads on R2 custom domains due to an incorrect HTTP/3 advertisement. We are working on a fix.
  - `2026-09-01 14:46` identified — The issue has been identified and a fix is being implemented.
  - `2026-09-03 15:14` monitoring — A fix has been implemented and we are monitoring the results.
  - `2026-09-03 19:32` resolved — This incident has been resolved.

## 2. gcp — Multiple products in us-central1-b are experiencing network service degradation.

- URL: https://status.cloud.google.com/incidents/J5ia5t9p3g9Q5Wi7r8Ev
- Window: 2026-09-01 14:44 to 2026-09-01 18:52 UTC (4h 8m)
- Impact: major, status: resolved
- Components: AlloyDB for PostgreSQL, Apigee, Cloud Filestore, Cloud Run, Cloud Spanner, Google App Engine, Google BigQuery, Google Cloud Bigtable, Google Cloud Dataflow, Google Cloud SQL, Google Compute Engine, Google Kubernetes Engine, Hybrid Connectivity, Looker (Google Cloud core), Virtual Private Cloud (VPC)

  - `2026-09-01 16:03` investigating — **Summary** Multiple products in us-central1-b are experiencing network service degradation. **Description** We are experiencing a networking issue starting Tuesday, 2026-09-01 07:41 US/Pacific. The issue is affecting a portion of the us-central1-b zone. The remaining zones in the region are unaffected. Our network engineering team has put in mitigations and we are seeing services starting to recover. We will provide more information by Tuesday, 2026-09-01 09:30 US/Pacific with current details. We apologize to all who are affected by the disruption. **Customer Symptoms** Customers are experiencing elevated packet loss and errors. * BigQuery * Cloud Dataflow * Cloud SQL * App Engine * Cloud Run * AlloyDB * Cloud Bigtable * Cloud Spanner * Compute Engine * Cloud VPN * Cloud Interconnect * C…
  - `2026-09-01 16:19` investigating — **Summary** Multiple products in us-central1-b are experiencing network service degradation. **Description** The issue is affecting a portion of the us-central1-b zone. The remaining zones in the region are unaffected. Following network-level mitigation efforts, the underlying network infrastructure has fully recovered. Product engineering teams across Cloud are currently validating the recovery status of individual services. We will provide more information by Tuesday, 2026-09-01 10:00 US/Pacific with current details. **Customer Symptoms** Customers are experiencing elevated packet loss and errors. * BigQuery * Cloud Dataflow * Cloud SQL * App Engine * Cloud Run * AlloyDB * Cloud Bigtable * Cloud Spanner * Compute Engine * Cloud VPN * Cloud Interconnect * Cloud NAT * Cloud Filestore * Clo…
  - `2026-09-01 17:09` investigating — **Summary** Multiple products in us-central1-b are experiencing network service degradation. **Description** The issue is affecting a portion of the us-central1-b zone. The remaining zones in the region are unaffected. Most of the products have recovered and the engineering team is still validating a few remaining product impacts. We will provide more information by Tuesday, 2026-09-01 11:00 US/Pacific with current details. **Customer Symptoms** Customers are experiencing elevated packet loss and errors. * Cloud Run * App Engine Following products have recovered * BigQuery * Cloud Dataflow * Cloud Bigtable * Compute Engine * Cloud Hybrid Connectivity * Cloud Filestore * Virtual Private Cloud * Cloud Spanner * Google Kubernetes Engine * Cloud SQL * AlloyDBCloud VPN * Apigee * Cloud Intercon…
  - `2026-09-01 18:01` investigating — **Summary** Multiple products in us-central1-b are experiencing network service degradation. **Description** The issue is affecting a portion of the us-central1-b zone. The remaining zones in the region are unaffected. Most of the products have recovered and the engineering team is still validating a few remaining product impacts. Routine network maintenance triggered issues in us-central1-b. We have recovered the network capacity that was impacted and restored all zonal and regional services. Maintenance in the region has been halted while proactive audits are carried out. We will provide more information by Tuesday, 2026-09-01 12:30 US/Pacific with current details. **Customer Symptoms** Customers are experiencing elevated packet loss and errors. * Cloud Run * App Engine Following product…
  - `2026-09-01 18:36` investigating — **Summary** Multiple products in us-central1-b are experiencing network service degradation. **Description** The issue is affecting a portion of the us-central1-b zone. The remaining zones in the region are unaffected. Most of the products have recovered. The impacted products Cloud Run and App Engine are also seeing recovery and the errors have lowered pre-incident levels. Engineering continues to monitor and validate. Routine network fabric path maintenance triggered unexpected issues in two clusters in us-central1-b. We have recovered the network capacity that was impacted and restored all zonal and regional services. Maintenance in the region has been halted while proactive audits are carried out. We will provide more information by Tuesday, 2026-09-01 12:30 US/Pacific with current det…
  - `2026-09-01 19:12` investigating — **Summary** Multiple products in us-central1-b experienced network service degradation. **Description** The issue was affecting a portion of the us-central1-b zone. The remaining zones in the region were unaffected. All the impacted products have recovered. From preliminary analysis, routine network fabric path maintenance triggered unexpected issues in one cluster in us-central1-b. We have recovered the network capacity that was impacted and restored all zonal and regional services. Maintenance in the region has been halted while proactive audits are carried out. We thank you for your patience while we worked on resolving the issue. **Customer Symptoms** Customers experienced elevated packet loss and errors. **Workaround** The issue has been mitigated.
  - `2026-09-03 17:47` resolved — ## Preliminary Incident Report We sincerely apologize for the disruption this incident caused to your business. We know how much you rely on Google Cloud, and we regret the impact on your environment. We are working to address the root cause and prevent this from occurring in the future. Please note, this information is based on our best knowledge at the time of posting and is subject to change as our investigation continues. A final Incident Report with preventative actions will be posted once our investigation is complete. If you have experienced impact outside of what is listed below, please reach out to Google Cloud Support using https://cloud.google.com/support ## Date/Time of the Issue (All time US/Pacific) **Incident Start**: 1 September 2026 07:41 **Incident End**: 1 September 2026…

## 3. grafana — US Central Region Instability

- URL: https://stspg.io/b3b4ppnkx23p
- Window: 2026-09-04 16:45 to 2026-09-04 23:00 UTC (6h 14m)
- Impact: minor, status: resolved
- Components: GCP US Central - prod-us-central-0: Querying, GCP US Central - prod-us-central-0: Ingestion, GCP US Central - prod-us-central-5: Querying, GCP US Central - prod-us-central-5: Ingestion, GCP US Central - prod-us-central-0: Querying, GCP US Central - prod-us-central-0: Ingestion, GCP US Central - prod-us-central-0: Metrics Generator, GCP US Central - prod-us-central-5: Querying, GCP US Central - prod-us-central-5: Ingestion

  - `2026-09-04 16:45` identified — We’re currently investigating an issue affecting the us central region due to ongoing cloud provider incident. Multiple components impacted. Our team is actively working on this. Thank you for your patience.
  - `2026-09-04 18:58` monitoring — The underlying cloud provider issue is not yet fully resolved, but our systems have recovered and affected Grafana Cloud stacks are loading normally again. Some customers may still see brief intermittent slowness. We are continuing to monitor until our provider confirms resolution.
  - `2026-09-04 23:00` resolved — The US Central region is fully recovered now and the issue is now resolved.

## 4. openai — Elevated latency in the Responses API

- URL: https://status.openai.com/incidents/01M1CYZDD84XD0H56XNEV2PT98
- Window: 2026-08-31 22:27 to 2026-09-01 19:05 UTC (20h 37m)
- Impact: minor, status: resolved
- Components: —

  - `2026-08-31 22:27` identified — We have identified an issue causing elevated latency for some Responses API requests. We are working on implementing a mitigation.
  - `2026-08-31 23:13` monitoring — We have applied a mitigation and are monitoring recovery.
  - `2026-09-01 01:46` monitoring — We have applied the mitigation and are monitoring the recovery.
  - `2026-09-01 15:15` identified — We have identified elevated latencies impacting Responses API on the US endpoint.
  - `2026-09-01 17:36` identified — We are implementing the mitigation.
  - `2026-09-01 18:09` monitoring — We have applied the mitigation and are monitoring the recovery.
  - `2026-09-01 19:05` resolved — All impacted services have now fully recovered.

## 5. cloudflare — Cloudflare Access one-time PIN emails blocked by certain email security gateways, including Proofpoint

- URL: https://www.cloudflarestatus.com/incidents/cczd80189gvc
- Window: 2026-09-01 20:06 to 2026-09-02 07:16 UTC (11h 9m)
- Impact: minor, status: resolved
- Components: Access

  - `2026-09-01 20:06` investigating — Some customers using email security gateway like Proofpoint may experience delayed Cloudflare notifications and login codes. Proofpoint is temporarily deferring messages with an SMTP 421 4.7.0 response. Cloudflare will continue retrying delivery.
  - `2026-09-02 07:16` resolved — This incident has been resolved.

## Also above threshold (25)

- `cloudflare` Pipelines Creation Issue (10h 7m) — https://www.cloudflarestatus.com/incidents/w6ng4jcxv7m1
- `anthropic` Elevated errors for multiple models (2h 57m) — https://stspg.io/9xz4hhmd1jzn
- `grafana` High Latency in prod-ap-south-1 (3h 17m) — https://stspg.io/f6h7k3lq3v62
- `cloudflare` Requests from wrangler or other OAuth clients may experience elevated errors (3h 27m) — https://www.cloudflarestatus.com/incidents/3mjmdcmm86nx
- `grafana` Degradation of Hosted Grafana in US Central Region (3h 39m) — https://stspg.io/7n6x5qw9x9d5
- `grafana` Tempo and Mimir Read and Write Failures (2h 42m) — https://stspg.io/kw7xrpjrg134
- `cloudflare` Purchased domains from registrar not appearing in Cloudflare Dashboard (4h 43m) — https://www.cloudflarestatus.com/incidents/1ghtnh94pby6
- `cloudflare` Workers Builds elevated queue times (4h 33m) — https://www.cloudflarestatus.com/incidents/48wc4y9p1vym
- `openai` ChatGPT Work seeing elevated errors and latency (5h 23m) — https://status.openai.com/incidents/01M1C5M4K0WC8PPT0Z175RJA1E
- `cloudflare` Workers KV experiencing elevated error rates in Western Europe Region (4h 2m) — https://www.cloudflarestatus.com/incidents/kby271yt9vwx
- `cloudflare` Elevated number of R2 503 errors in Eastern North America region (3h 31m) — https://www.cloudflarestatus.com/incidents/ftvf8c3m4mv5
- `grafana` Alert rule creation, deletion, and update degradation in prod-us-east-2 (2h 56m) — https://stspg.io/8gw4nt4mmp09
- `cloudflare` Cloudflare Access bypass IP service experiencing availability issues (3h 10m) — https://www.cloudflarestatus.com/incidents/fcdl2yb3q7qw
- `datadog` Delayed CICD Optimization, Code Coverage, Code Security, and DORA data (5h 40m) — https://stspg.io/zstf1zx7hfy3
- `openai` Users in APAC region may face increased error in ChatGPT, Work, image generation, file upload, Voice, and Codex Cloud (3h 46m) — https://status.openai.com/incidents/01M1NKFZH5EEYEREC54HNAHY35
- `grafana` Investigating issues in US Central (prod-us-central-0, prod-us-central-5) (4h 54m) — https://stspg.io/jvcf7d0j5x37
- `aws` Increased API Error Rates (2h 17m) — https://health.aws.amazon.com/health/status
- `github` Incident with Grok Copilot AI Model Provider (2h 54m) — https://stspg.io/p217h6l54208
- `cloudflare` Durable Objects increased errors in Western North America (2h 7m) — https://www.cloudflarestatus.com/incidents/3qs7m80p13v2
- `cloudflare` Cache Purging Errors (2h 7m) — https://www.cloudflarestatus.com/incidents/ff0yvx1d6tby
- `github` Disruption with Copilot Code Review (1h 47m) — https://stspg.io/dddfhld6fj0z
- `anthropic` Elevated errors for Claude Sonnet 5 (26m) — https://stspg.io/9dgzwz3ymz42
- `anthropic` Elevated errors for Claude Sonnet 5 (18m) — https://stspg.io/jt14jk41pjsf
- `cloudflare` Network Route Leak in Palmas, Brazil (18m) — https://www.cloudflarestatus.com/incidents/1w42pbthkqjf
- `openai` ChatGPT Work Mode High Error Rates (6m) — https://status.openai.com/incidents/01M1J99EDCYED5GP2W8GZ3N2NA

