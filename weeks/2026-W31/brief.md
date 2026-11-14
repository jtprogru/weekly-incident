# Week 2026-W31 — top 5 of 19 above threshold

2026-07-27 to 2026-08-02 (UTC).

51 source records (53 saved revisions) → 51 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## 1. grafana — Issues with Billing/Usage Dashboard Metrics and Panels.

- URL: https://stspg.io/fz2tz9xsqx46
- Window: 2026-07-30 13:15 to 2026-07-30 15:17 UTC (2h 1m)
- Impact: major, status: resolved
- Components: AWS Australia - prod-ap-southeast-2: Ingestion, AWS Australia - prod-ap-southeast-2: Querying, AWS Brazil - prod-sa-east-1: Ingestion, AWS Brazil - prod-sa-east-1: Querying, AWS Canada - prod-ca-east-0: Ingestion, AWS Canada - prod-ca-east-0: Querying, AWS Germany - prod-eu-west-2: Ingestion, AWS Germany - prod-eu-west-2: Querying, AWS India - prod-ap-south-1: Ingestion, AWS India - prod-ap-south-1: Querying, AWS Ireland - prod-eu-west-6, AWS Japan - prod-ap-northeast-0: Ingestion, AWS Japan - prod-ap-northeast-0: Querying, AWS Singapore - prod-ap-southeast-1: Ingestion, AWS Singapore - prod-ap-southeast-1: Querying, AWS Sweden - prod-eu-north-0: Ingestion, AWS Sweden - prod-eu-north-0: Querying, AWS Switzerland - prod-eu-central-0, AWS UAE - prod-me-central-1: Ingestion, AWS UAE - prod-me-central-1: Querying, AWS US East - prod-us-east-0: Ingestion, AWS US East - prod-us-east-0: Querying, AWS US West - prod-us-west-0: Ingestion, AWS US West - prod-us-west-0: Querying, Azure Netherlands - prod-eu-west-3: Ingestion, Azure Netherlands - prod-eu-west-3: Querying, Azure US Central - prod-us-central-7: Ingestion, Azure US Central - prod-us-central-7: Querying, Azure US Central - us-central2: Ingestion, Azure US Central - us-central2: Querying, Azure US Central - us-central7: Ingestion, Federal Cloud - AWS US Gov West, GCP Australia - prod-au-southeast-0: Ingestion, GCP Australia - prod-au-southeast-0: Querying, GCP Belgium - prod-eu-west-0: Ingestion, GCP Belgium - prod-eu-west-0: Querying, GCP Brazil - prod-sa-east-0: Ingestion, GCP Brazil - prod-sa-east-0: Querying, GCP India - prod-ap-south-0: Ingestion, GCP India - prod-ap-south-0: Querying, GCP Singapore - prod-ap-southeast-0: Ingestion, GCP Singapore - prod-ap-southeast-0: Querying, GCP UK - prod-gb-south-0: Ingestion, GCP UK - prod-gb-south-0: Querying, GCP US Central - prod-us-central-0: Ingestion, GCP US Central - prod-us-central-0: Querying, GCP US Central - prod-us-central-5: Ingestion, GCP US Central - prod-us-central-5: Querying, GCS US - cortex-prod-04: Ingestion, GCS US - cortex-prod-04: Querying

<a id="ci-aa15d823569b32f896053338"></a>

- Canonical ID: `ci-aa15d823569b32f896053338` (1 source records).
  - [grafana/79nwp08zs9vj](../../data/2026/W31/grafana.json); vendor URL: https://stspg.io/fz2tz9xsqx46
- Start time basis: `vendor_reported`.

  - `2026-07-30 13:15` investigating — We are investigating an outage for the Billing / Usage dashboard metrics and panels. This appears to be partially affecting organizations using Grafana Cloud. We are working on identifying and resolving the issue
  - `2026-07-30 13:50` identified — We have identified the issue, and are working on deploying a fix.
  - `2026-07-30 14:34` monitoring — A fix has been implemented and we are monitoring the results.
  - `2026-07-30 15:17` resolved — This incident has been resolved.

## 2. cloudflare — Network Performance Issues in Hamburg, Germany

- URL: https://www.cloudflarestatus.com/incidents/3ywn8wy3kqh8
- Window: 2026-07-31 15:20 to 2026-08-01 15:18 UTC (23h 58m)
- Impact: minor, status: resolved
- Components: Hamburg, Germany - (HAM)

<a id="ci-cc131d6ee61cf325905371da"></a>

- Canonical ID: `ci-cc131d6ee61cf325905371da` (1 source records).
  - [cloudflare/3ywn8wy3kqh8](../../data/2026/W31/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/3ywn8wy3kqh8
- Start time basis: `vendor_reported`.

  - `2026-07-31 15:20` identified — Cloudflare is investigating issues with network performance in Hamburg, Germany (HAM). Customers routing through this location may experience request errors or failures. We have identified the problem and are working on a fix.
  - `2026-07-31 16:58` monitoring — A fix has been implemented and we are monitoring the results.
  - `2026-08-01 15:18` resolved — This incident has been resolved.

## 3. anthropic — Elevated errors across many models

- URL: https://stspg.io/cpzpstxn0z6k
- Window: 2026-07-30 05:57 to 2026-07-30 10:48 UTC (4h 51m)
- Impact: major, status: resolved
- Components: Claude API (api.anthropic.com), Claude Code, Claude Cowork, claude.ai

<a id="ci-d053ed79dc135551f56c34ab"></a>

- Canonical ID: `ci-d053ed79dc135551f56c34ab` (1 source records).
  - [anthropic/fsh2zzzl2c4l](../../data/2026/W31/anthropic.json); vendor URL: https://stspg.io/cpzpstxn0z6k
- Start time basis: `vendor_reported`.

  - `2026-07-30 05:57` investigating — We are currently investigating this issue.
  - `2026-07-30 06:26` identified — All models except Opus 5 have recovered. We're working to recover Opus 5 as well.
  - `2026-07-30 07:33` identified — As of 7:32 UTC / 0:32 PT, Opus 5 errors are back to baseline but we're seeing again Sonnet 5's errors up.
  - `2026-07-30 08:06` identified — As of 8:05 UTC / 1:05 PT, Sonnet 5 errors are back to baseline, but we're seeing an increased error rate on Fable 5.
  - `2026-07-30 08:33` identified — As of 01:33 PT / 08:33 UTC we are back to errors across all models.
  - `2026-07-30 10:03` monitoring — As of 2:58 PT / 9:58 UTC the errors are back to baseline.
  - `2026-07-30 10:23` monitoring — We are also investigating issues with Opus 4.7 and Opus 4.5.
  - `2026-07-30 10:48` resolved — This incident has been resolved.

## 4. slack — Feature Degradation Affecting the Receipt of Emails in Slack

- URL: https://slack-status.com/2026-07/6b1014d8bfb3afc0
- Window: 2026-07-28 02:13 to 2026-07-28 23:09 UTC (20h 55m)
- Impact: unknown, status: resolved
- Components: Apps/Integrations/APIs

<a id="ci-0626e37ed376c79ddab893c3"></a>

- Canonical ID: `ci-0626e37ed376c79ddab893c3` (1 source records).
  - [slack/1574](../../data/2026/W31/slack.json); vendor URL: https://slack-status.com/2026-07/6b1014d8bfb3afc0
- Start time basis: `vendor_reported`.

  - `2026-07-28 02:13` unknown — We're currently investigating an issue affecting users' ability to receive incoming emails into Slack, caused by some issues with our email ingestion pipeline. We're sorry for any inconvenience this is causing, our team is actively working on a fix and we'll share an update as soon as we have more information.
  - `2026-07-28 03:25` unknown — We identified the cause of the issue preventing users from receiving incoming emails. We are deploying a fix and will share a further update once deployment and verification are complete.
  - `2026-07-28 03:43` unknown — The fix has been successfully deployed, and validation is now complete. The issue preventing users from receiving incoming emails into Slack is fully resolved, and normal service functionality has been restored. We apologize for any disruptions to your day.

## 5. digitalocean — Agent Platform Requests Returning HTTP 500 Errors

- URL: https://stspg.io/54mh0lzk9733
- Window: 2026-07-27 18:31 to 2026-07-28 15:24 UTC (20h 52m)
- Impact: major, status: resolved
- Components: Agent Runtime

<a id="ci-6c1b5d39356b2854e0b85c2c"></a>

- Canonical ID: `ci-6c1b5d39356b2854e0b85c2c` (1 source records).
  - [digitalocean/fcfplv03vbwt](../../data/2026/W31/digitalocean.json); vendor URL: https://stspg.io/54mh0lzk9733
- Start time basis: `vendor_reported`.

  - `2026-07-27 18:31` investigating — We are currently investigating this issue.
  - `2026-07-27 19:44` identified — The issue has been identified and a fix is being implemented.
  - `2026-07-28 15:24` resolved — This incident has been resolved.

## Also above threshold (14)

- `yandexcloud` The problem is with the services (7h 54m) — https://status.yandex.cloud/en/incidents/1901
- `grafana` PDC Authentication Issues (2h 54m) — https://stspg.io/p8c7l7pzd5f5
- `digitalocean` Response Degradation Impacting Kimi-K3 (12h 19m) — https://stspg.io/ynhlvpy6ctcm
- `cloudflare` Increased HTTP 5XX Errors in IAD (4h 30m) — https://www.cloudflarestatus.com/incidents/l63k37vrcd9c
- `cloudflare` Possible Network Congestion in North America region (4h 20m) — https://www.cloudflarestatus.com/incidents/z33kqrzksxqy
- `yandexcloud` The problem is with DataSphere services (4h 29m) — https://status.yandex.cloud/en/incidents/1893
- `cloudflare` Possible Network Congestion between Singapore and Tokyo (3h 16m) — https://www.cloudflarestatus.com/incidents/h1kxms9wcbwj
- `anthropic` Elevated errors across all models (2h 46m) — https://stspg.io/6xr5mmpjs1k3
- `cloudflare` Increased HTTP 5xx Errors for us-east-1-aws (2h 18m) — https://www.cloudflarestatus.com/incidents/s18kw61f2ht5
- `grafana` IRM Performance Degradation in EU Region (2h 52m) — https://stspg.io/35m2lfvv6h87
- `github` Incident with Actions (34m) — https://stspg.io/8nsh6820sff9
- `cloudflare` Cloudflare Workers — Errors Deploying Workers Scripts (27m) — https://www.cloudflarestatus.com/incidents/z325n2dc14tk
- `datadog` Synthetics Test Results Delayed (46m) — https://stspg.io/9trx2m1xsd4h
- `grafana` Partial OTLP Write Outage in prod-us-east-3 (3m) — https://stspg.io/05bt4v67qbyk
