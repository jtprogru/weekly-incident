# Week 2026-W04

2026-01-19 to 2026-01-25 (UTC). Generated 2026-11-14 00:00.

1 canonical incidents, 1 above threshold. All 14 sources responded.

1 source records (1 saved revisions) → 1 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | datadog | [Web Application Not Loading](https://stspg.io/qzwqj368qbvk) | 2026-01-22 18:51 | 36m | critical | APM, App Builder, Application Security Management, Application Vulnerability Management, Audit Trail, CI Visibility, Cloud Cost Management, Cloud Network Monitoring, Cloud SIEM, Cloud Security Management, Code Security, Continuous Profiler, Data Jobs Monitoring, Data Streams Monitoring, Database Monitoring, Incident Response, LLM Observability, Log Management, Metrics and Infra Monitoring, Mobile Application, Monitors, NDM, Observability Pipelines, Package Repositories, RUM, Sensitive Data Scanner, Serverless, Synthetics, Universal Service Monitoring, Web Application | 1080 = 36 × 30 × 1.0 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. datadog — Web Application Not Loading

- URL: https://stspg.io/qzwqj368qbvk
- Started: 2026-01-22 18:51 UTC
- Resolved: 2026-01-22 19:27 UTC
- Duration: 36m
- Impact: critical, status: resolved
- Components: APM, App Builder, Application Security Management, Application Vulnerability Management, Audit Trail, CI Visibility, Cloud Cost Management, Cloud Network Monitoring, Cloud SIEM, Cloud Security Management, Code Security, Continuous Profiler, Data Jobs Monitoring, Data Streams Monitoring, Database Monitoring, Incident Response, LLM Observability, Log Management, Metrics and Infra Monitoring, Mobile Application, Monitors, NDM, Observability Pipelines, Package Repositories, RUM, Sensitive Data Scanner, Serverless, Synthetics, Universal Service Monitoring, Web Application
- Canonical incident: [ci-874fd359f9be51b93c9954f1](#ci-874fd359f9be51b93c9954f1)

- `2026-01-22 18:51` **investigating** — We are investigating loading issues on our web application. As a result, some users might be getting errors when loading the web application. Please note that data processing and alerts are not affected by this incident.
- `2026-01-22 19:01` **investigating** — We are continuing to investigate this issue.
- `2026-01-22 19:13` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-01-22 19:27` **resolved** — This incident has been resolved.

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-874fd359f9be51b93c9954f1"></a>

- Canonical ID: `ci-874fd359f9be51b93c9954f1` (1 source records).
  - [datadog/0vy6kj9tgmt6](../../data/2026/W04/datadog.json); vendor URL: https://stspg.io/qzwqj368qbvk
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
