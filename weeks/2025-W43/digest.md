# Week 2025-W43

2025-10-20 to 2025-10-26 (UTC). Generated 2026-11-14 00:00.

2 canonical incidents, 2 above threshold. All 13 sources responded.

2 source records (2 saved revisions) → 2 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | datadog | [Multiple products impacted with data delays](https://stspg.io/gzvrd592kg4t) | 2025-10-20 12:45 | 2d 1h 54m | major | APM, App Builder, Cloud Cost Management, Continuous Profiler, Log Management, RUM, Workflow Automation | 20958 = 2994 × 7 × 1.0 |
| 2 | datadog | [Multiple products impacted with data delays](https://stspg.io/29j9dsc8llxg) | 2025-10-20 10:14 | 1h 16m | major | Cloud Security Management, Incident Response | 152 = 76 × 2 × 1.0 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. datadog — Multiple products impacted with data delays

- URL: https://stspg.io/gzvrd592kg4t
- Started: 2025-10-20 12:45 UTC
- Resolved: 2025-10-22 14:40 UTC
- Duration: 2d 1h 54m
- Impact: major, status: resolved
- Components: APM, App Builder, Cloud Cost Management, Continuous Profiler, Log Management, RUM, Workflow Automation
- Canonical incident: [ci-f140dc983588253f32d7e0a7](#ci-f140dc983588253f32d7e0a7)

- `2025-10-20 12:45` **identified** — We are investigating increased latency processing APM, RUM, Log Management and Profiling. As a result of this issue, some users may see only a subset of their data when querying those different products, other product pages using the same underlying product data will be impacted as well. We are working on bringing new capacity online and the data will be backfilled once the service is fully operational again.
- `2025-10-20 14:07` **identified** — We are investigating increased latency processing APM, RUM, Log Management and Profiling. As a result of this issue, some users may see only a subset of their data when querying those different products, other product pages using the same underlying product data will be impacted as well. Monitors using the impacted data are delayed. We are working on bringing new capacity online and will provide an update once the service is fully operational again.
- `2025-10-20 15:18` **identified** — We are still seeing increased latency processing for those products and the associated monitors are delayed. We are continuing to work on bringing new capacity online and will continue to provide updates on this issue.
- `2025-10-20 17:05` **identified** — APM, RUM, Log Management, Profiling, CCM and Product Analytics data is still delayed. As a result of this issue, some users may see only a subset of their data when querying those different products, other product pages using the same underlying product data will be impacted as well. We are working on bringing new capacity online and for all products except RUM we expect the data will be backfilled once the service is fully operational again. App Builder and Workflow Automation are also experiencing elevated errors, as a result customers might not be to query applications and workflows might take longer to execute. Due to upstream provider issues, we are also continuing to see unavailability of telemetry data coming from AWS into Datadog.
- `2025-10-20 18:04` **identified** — APM, RUM, Log Management, Profiling, CCM and Product Analytics data is still delayed. As a result of this issue, some users may see only a subset of their data when querying those different products, other product pages using the same underlying product data will be impacted as well. We are working on bringing new capacity online and for all products except RUM we expect the data will be backfilled once the service is fully operational again. App Builder and Workflow Automation are also experiencing elevated errors, as a result customers might not be to query applications and workflows might take longer to execute. Due to upstream provider issues, we are also continuing to see unavailability of telemetry data coming from AWS into Datadog.
- `2025-10-20 19:01` **identified** — We are seeing progress in telemetry data coming from AWS into Datadog. Also, we are starting to see our capacity requests being fulfilled. Our processing is still delayed impacting multiple products - Distribution Metrics, APM, RUM, Log Management, Profiling, CCM and Product Analytics data is still delayed. As a result of this issue, some users may see only a subset of their data when querying those different products, other product pages using the same underlying product data will be impacted as well. App Builder and Workflow Automation are also experiencing elevated errors, as a result customers might not be to query applications and workflows might take longer to execute.
- `2025-10-20 20:14` **identified** — We are seeing progress in telemetry data coming from AWS into Datadog. We are starting to see our capacity requests being fulfilled more slowly than usual. App Builder and Workflow Automation are seeing recovery. Our processing is still delayed impacting multiple products - Distribution Metrics, APM, RUM, Log Management, Profiling, CCM and Product Analytics data is still delayed. As a result of this issue, some users may see only a subset of their data when querying those different products, other product pages using the same underlying product data will be impacted as well.
- `2025-10-20 21:47` **identified** — We are seeing recovery in AWS Metrics. Logs data submitted after 21:30 UTC should be processed normally. Users may see gaps in historical logs prior to 21:30 UTC while our backfill is in progress. In addition to Log Management we continue to see delays in processing that impacts the following products: Distribution Metrics, APM, RUM, Profiling, CCM and Product Analytics. As a result of this issue, some users may see only a subset of their data when querying those products or viewing pages that rely on telemetry from those products.
- `2025-10-20 22:40` **identified** — We are seeing recovery in Profiling. Logs data submitted after 21:30 UTC should be processed normally. Users may see gaps in historical logs prior to 21:30 UTC while our backfill is in progress. In addition to Log Management we continue to see delays in processing that impacts the following products: Distribution Metrics, APM, RUM, CCM and Product Analytics. As a result of this issue, some users may see only a subset of their data when querying those products or viewing pages that rely on telemetry from those products.
- `2025-10-20 22:41` **identified** — Logs data have been backfilled, and users should no longer see gaps in their historical logs. Log Archives and Log Forwarding were paused between 15:00 and 18:30 UTC, and we are working to re-forward any logs from that time period. We continue to see delays in processing that impact the following products: Distribution Metrics, APM, RUM, CCM, and Product Analytics. As a result of this issue, some users may see only a subset of their data when querying those products or viewing pages that rely on telemetry from those products.
- `2025-10-21 00:25` **identified** — We are seeing recovery for APM. We continue to see delays in processing that impact the following products: Distribution Metrics, RUM, CCM, and Product Analytics. As a result of this issue, some users may see only a subset of their data when querying those products or viewing pages that rely on telemetry from those products.
- `2025-10-21 01:32` **monitoring** — We are seeing recovery across all of our products, and live data and monitor evaluations have resumed for all affected products. Most historical data in Logs has been backfilled and we have a small number of ongoing backfills in Metrics and other products. We will continue to monitor the situation overnight, and our next update will be 09:00 UTC.
- `2025-10-21 10:20` **monitoring** — All products have been stable since the last update. We are continuing the work on outstanding backfills, during this process queries that include data from the backfilled windows may appear incomplete for the affected subset of customers and products. We will resolve the incident when the backfills are complete or before Oct 21, 16:00 UTC.
- `2025-10-21 18:04` **monitoring** — We are continuing the work on outstanding backfills which are not yet fully complete, during this process queries that include data from the backfilled windows may appear incomplete for the affected subset of customers and products. We will resolve the incident when the backfills are complete or before Oct 21, 22:00 UTC.
- `2025-10-21 21:35` **monitoring** — We are making progress on outstanding backfills. Cloud Cost Monitoring backfill is complete. Metrics and Logs backfills are still in progress. For products still undergoing backfilling, queries that include data from the backfilled windows may appear incomplete for the affected subset of customers. We will provide next update no later than Oct 22, 10:00 UTC.
- `2025-10-22 09:10` **monitoring** — We are making progress on outstanding backfills. Metrics and Logs backfills are still in progress. For products still undergoing backfilling, queries that include data from the backfilled windows may appear incomplete for the affected subset of customers. We will provide next update no later than Oct 22, 16:00 UTC.
- `2025-10-22 14:40` **resolved** — Backfills for Metrics and Log Management data have completed. All systems are back to normal.

### 2. datadog — Multiple products impacted with data delays

- URL: https://stspg.io/29j9dsc8llxg
- Started: 2025-10-20 10:14 UTC
- Resolved: 2025-10-20 11:31 UTC
- Duration: 1h 16m
- Impact: major, status: resolved
- Components: Cloud Security Management, Incident Response
- Canonical incident: [ci-cd239c5828c932cee2002596](#ci-cd239c5828c932cee2002596)

- `2025-10-20 10:14` **identified** — Note: this is a delayed update because this incident impaired our ability to update the status page, we posted banners earlier in the product to let customers know about the ongoing impact. We are still seeing some delays as we are fully recovering from the underlying incident: agentless vulnerability scanning for hosts in AWS us-east-1 is still delayed, On-Call notifications are not fully recovered. This incident started at 07:10 UTC on October 20. So far we have recovered fully from the impact on Synthetics, collection of data from AWS, Bits AI, Codegen, Dashboards (edition features were impaired).
- `2025-10-20 10:53` **monitoring** — We are monitoring and seeing recovery for all products, some customers might still experience for a limited subset of data delays for logs or host vulnerability scanning specific to AWS us-east-1. We will post specific information on the affected product pages for those customers. On-Call notifications are fully operational.
- `2025-10-20 11:31` **resolved** — This incident has been resolved.

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-f140dc983588253f32d7e0a7"></a>

- Canonical ID: `ci-f140dc983588253f32d7e0a7` (1 source records).
  - [datadog/vjgmxjdf0ps5](../../data/2025/W43/datadog.json); vendor URL: https://stspg.io/gzvrd592kg4t
- Start time basis: `vendor_reported`.

<a id="ci-cd239c5828c932cee2002596"></a>

- Canonical ID: `ci-cd239c5828c932cee2002596` (1 source records).
  - [datadog/ww8rhx2j8mfz](../../data/2025/W43/datadog.json); vendor URL: https://stspg.io/29j9dsc8llxg
- Start time basis: `vendor_reported`.

## Sources

| Vendor | Status | Source records seen | Parse errors |
|--------|--------|----------------|--------------|
| anthropic | ok | 50 | 0 |
| atlassian | ok | 34 | 0 |
| aws | ok | 5 | 0 |
| circleci | ok | 50 | 0 |
| cloudflare | ok | 50 | 0 |
| datadog | ok | 50 | 0 |
| digitalocean | ok | 50 | 0 |
| gcp | ok | 4 | 0 |
| github | ok | 50 | 0 |
| grafana | ok | 50 | 0 |
| openai | ok | 25 | 0 |
| slack | ok | 43 | 0 |
| stripe | ok | 50 | 0 |
