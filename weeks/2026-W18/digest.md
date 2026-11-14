# Week 2026-W18

2026-04-27 to 2026-05-03 (UTC). Generated 2026-11-14 12:36.

10 canonical incidents, 7 above threshold. All 14 sources responded.

10 source records (10 saved revisions) → 10 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | aws | [Increased Connectivity Issues and API Error Rates](https://health.aws.amazon.com/health/status) | 2026-04-30 07:07 | 198d 5h 29m (ongoing) | unknown | multipleservices-me-south-1 | 570898 = 285449 × 1 × 2.0 |
| 2 | aws | [Increased Error Rates](https://health.aws.amazon.com/health/status) | 2026-04-30 07:25 | 198d 5h 10m (ongoing) | unknown | multipleservices-me-central-1 | 570860 = 285430 × 1 × 2.0 |
| 3 | digitalocean | [Elevated 5xx “context canceled” errors impacting serverless inference](https://stspg.io/qf13crcjzkmn) | 2026-04-28 13:45 | 5h 48m | minor | Agent Runtime | 278 = 348 × 1 × 0.8 |
| 4 | slack | [Users May Not Receive Notifications on Slack Thread Replies](https://slack-status.com/2026-04/22195f9f2ca47237) | 2026-04-27 17:11 | 3h 24m | unknown | Notifications | 245 = 204 × 1 × 1.2 |
| 5 | circleci | [Usage API data is delayed for data from yesterday, 4/30. All prior data is available. We are investigating.](https://stspg.io/ld03v9l2cm5c) | 2026-05-01 10:04 | 3h 57m | none | CircleCI API | 237 = 237 × 1 × 1.0 |
| 6 | digitalocean | [Serverless Inference - Intermittent Rate Limiting Affecting Some Customers Using Anthropic Models](https://stspg.io/y8cgtdqrkct1) | 2026-04-27 10:38 | 59m | major | Agent Runtime | 47 = 59 × 1 × 0.8 |
| 7 | datadog | [Elevated Error Rates for Metrics Queries](https://stspg.io/17n4grhvbmjt) | 2026-05-01 00:45 | 34m | major | Metrics and Infra Monitoring | 34 = 34 × 1 × 1.0 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. aws — Increased Connectivity Issues and API Error Rates

- URL: https://health.aws.amazon.com/health/status
- Started: 2026-04-30 07:07 UTC
- Resolved: not yet
- Duration: 198d 5h 29m (ongoing)
- Impact: unknown, status: investigating
- Components: multipleservices-me-south-1
- Canonical incident: [ci-4f155127c97b96a474760521](#ci-4f155127c97b96a474760521)

- `2026-04-30 07:07` **investigating** — We are providing an update on the ongoing service disruption. The Middle East (Bahrain) Region (ME-SOUTH-1) has suffered damage due to the conflict in the Middle East and is currently unavailable. Customers should recover their resources in other Regions from remote backups. Relevant billing operations are currently suspended while we restore normal operations in this AWS Region. This process is expected to take several months.

### 2. aws — Increased Error Rates

- URL: https://health.aws.amazon.com/health/status
- Started: 2026-04-30 07:25 UTC
- Resolved: not yet
- Duration: 198d 5h 10m (ongoing)
- Impact: unknown, status: investigating
- Components: multipleservices-me-central-1
- Canonical incident: [ci-0e9fcfad9f58dcd7e204aafb](#ci-0e9fcfad9f58dcd7e204aafb)

- `2026-04-30 07:25` **investigating** — We are providing an update on the ongoing service disruption. The Middle East (UAE) Region (ME-CENTRAL-1) has suffered damage as a result of the conflict in the Middle East and is currently unable to reliably support customer applications. While some workloads continue to function normally, we strongly recommend customers migrate all accessible resources to other Regions and restore inaccessible resources from remote backups as soon as possible. Relevant billing operations are currently suspended while we restore normal operations in this AWS Region. This process is expected to take several months.

### 3. digitalocean — Elevated 5xx “context canceled” errors impacting serverless inference

- URL: https://stspg.io/qf13crcjzkmn
- Started: 2026-04-28 13:45 UTC
- Resolved: 2026-04-28 19:33 UTC
- Duration: 5h 48m
- Impact: minor, status: resolved
- Components: Agent Runtime
- Canonical incident: [ci-35367f529f55a4569db74813](#ci-35367f529f55a4569db74813)

- `2026-04-28 13:45` **investigating** — Serverless inference customers are experiencing elevated 5xx errors, including “context canceled” responses. This may result in intermittent request failures. Our team is actively investigating and will provide updates as more information becomes available.
- `2026-04-28 15:59` **identified** — We have identified an issue affecting our service and are currently working to implement a fix. Our team is actively investigating and taking the necessary steps to restore normal operations as quickly as possible. We appreciate your patience and will provide updates as soon as more information becomes available.
- `2026-04-28 19:00` **monitoring** — Service for Serverless Inference has been restored. We’ve implemented tighter rate limits to help prevent recurrence and are closely monitoring system performance. Some users may still experience intermittent latency as we complete final stabilization efforts. Our team remains actively engaged to ensure full recovery. We appreciate your patience and will provide further updates as needed.
- `2026-04-28 19:33` **resolved** — All services are operating normally. We will continue to monitor the system to ensure ongoing reliability. Thank you for your patience while we worked to resolve this issue.

### 4. slack — Users May Not Receive Notifications on Slack Thread Replies

- URL: https://slack-status.com/2026-04/22195f9f2ca47237
- Started: 2026-04-27 17:11 UTC
- Resolved: 2026-04-27 20:36 UTC
- Duration: 3h 24m
- Impact: unknown, status: resolved
- Components: Notifications
- Canonical incident: [ci-21a7f32d8d03f521c103d75b](#ci-21a7f32d8d03f521c103d75b)

- `2026-04-27 17:11` **unknown** — Some users have reported issues with thread replies not generating a notification banner. We’re investigating the issue, and have identified a recent deployment as a possible trigger. We apologize for the inconvenience and will report back as more information becomes available.
- `2026-04-27 18:07` **unknown** — We’re observing positive signals after rolling back the deployment in test environments. We will complete some final validations before rolling back the deployment in the production environment. Thank you for your patience as we work through this one. We’ll report back once production rollback is completed or if there are substantial updates to share.
- `2026-04-27 20:36` **unknown** — We’ve successfully performed a rollback in the production environment, and all systems are now functioning as expected. Users should no longer experience any issues with thread replies not generating a notification banner. We apologize for any inconvenience and appreciate your patience.

### 5. circleci — Usage API data is delayed for data from yesterday, 4/30. All prior data is available. We are investigating.

- URL: https://stspg.io/ld03v9l2cm5c
- Started: 2026-05-01 10:04 UTC
- Resolved: 2026-05-01 14:02 UTC
- Duration: 3h 57m
- Impact: none, status: resolved
- Components: CircleCI API
- Canonical incident: [ci-df5d09dc1620e3242436ac57](#ci-df5d09dc1620e3242436ac57)

- `2026-05-01 10:04` **investigating** — We are currently investigating
- `2026-05-01 14:02` **resolved** — This issue has been resolved. All data is up to date through yesterday.

### 6. digitalocean — Serverless Inference - Intermittent Rate Limiting Affecting Some Customers Using Anthropic Models

- URL: https://stspg.io/y8cgtdqrkct1
- Started: 2026-04-27 10:38 UTC
- Resolved: 2026-04-27 11:37 UTC
- Duration: 59m
- Impact: major, status: resolved
- Components: Agent Runtime
- Canonical incident: [ci-5ed8106e69c1175914306ca2](#ci-5ed8106e69c1175914306ca2)

- `2026-04-27 10:38` **investigating** — We are investigating an issue affecting some customers using DigitalOcean Serverless Inference with Anthropic models. Over the last two hours, impacted customers may have experienced intermittent request failures, including HTTP 429 responses, on some Anthropic model requests. Our engineering team is actively investigating the issue. We apologize for the inconvenience and will share another update as soon as more information is available.
- `2026-04-27 11:07` **monitoring** — We identified the cause of intermittent HTTP 429 responses affecting some customers using Anthropic models on DigitalOcean Serverless Inference and applied a mitigation. Service has recovered, and we are monitoring stability.
- `2026-04-27 11:37` **resolved** — The issue is resolved, and service is operating normally.

### 7. datadog — Elevated Error Rates for Metrics Queries

- URL: https://stspg.io/17n4grhvbmjt
- Started: 2026-05-01 00:45 UTC
- Resolved: 2026-05-01 01:19 UTC
- Duration: 34m
- Impact: major, status: resolved
- Components: Metrics and Infra Monitoring
- Canonical incident: [ci-0807986938bb64d909571aa3](#ci-0807986938bb64d909571aa3)

- `2026-05-01 00:45` **investigating** — We are actively investigating elevated error rates for Metrics Queries. As a result of this issue, some users may see errors with metrics graphs on the web application or API.
- `2026-05-01 01:09` **monitoring** — We have deployed a fix and we are monitoring the results. We will provide another update once the service is fully operational.
- `2026-05-01 01:19` **resolved** — This incident has been resolved.

## Below threshold (3)

- `circleci` [Jobs using Xcode 14.3.1 on m4pro.large are delayed.](https://stspg.io/f6kcq9nb85wt) — 16m
- `stripe` [Certificate warnings on stripe.com and dashboard logins](https://stspg.io/3x0dw69bql3k) — 0m
- `slack` [Slack Search Feature Disruption - Japan region](https://slack-status.com/2026-04/4bd5196b0c972734) — 0m

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-4f155127c97b96a474760521"></a>

- Canonical ID: `ci-4f155127c97b96a474760521` (1 source records).
  - [aws/dd4f7c1f4ca3fe41](../../data/2026/W18/aws.json); vendor URL: https://health.aws.amazon.com/health/status
- Start time basis: `first_observed_update`.
- Warning: AWS onset is the first observed message, not a verified impact start

<a id="ci-0e9fcfad9f58dcd7e204aafb"></a>

- Canonical ID: `ci-0e9fcfad9f58dcd7e204aafb` (1 source records).
  - [aws/ca5f1cc2b6034dd6](../../data/2026/W18/aws.json); vendor URL: https://health.aws.amazon.com/health/status
- Start time basis: `first_observed_update`.
- Warning: AWS onset is the first observed message, not a verified impact start

<a id="ci-35367f529f55a4569db74813"></a>

- Canonical ID: `ci-35367f529f55a4569db74813` (1 source records).
  - [digitalocean/z5fjzx3kp16z](../../data/2026/W18/digitalocean.json); vendor URL: https://stspg.io/qf13crcjzkmn
- Start time basis: `vendor_reported`.

<a id="ci-21a7f32d8d03f521c103d75b"></a>

- Canonical ID: `ci-21a7f32d8d03f521c103d75b` (1 source records).
  - [slack/1547](../../data/2026/W18/slack.json); vendor URL: https://slack-status.com/2026-04/22195f9f2ca47237
- Start time basis: `vendor_reported`.

<a id="ci-df5d09dc1620e3242436ac57"></a>

- Canonical ID: `ci-df5d09dc1620e3242436ac57` (1 source records).
  - [circleci/mspf5pqp7d42](../../data/2026/W18/circleci.json); vendor URL: https://stspg.io/ld03v9l2cm5c
- Start time basis: `vendor_reported`.

<a id="ci-5ed8106e69c1175914306ca2"></a>

- Canonical ID: `ci-5ed8106e69c1175914306ca2` (1 source records).
  - [digitalocean/5gsk0hfctqp5](../../data/2026/W18/digitalocean.json); vendor URL: https://stspg.io/y8cgtdqrkct1
- Start time basis: `vendor_reported`.

<a id="ci-0807986938bb64d909571aa3"></a>

- Canonical ID: `ci-0807986938bb64d909571aa3` (1 source records).
  - [datadog/dw41jpjjghhs](../../data/2026/W18/datadog.json); vendor URL: https://stspg.io/17n4grhvbmjt
- Start time basis: `vendor_reported`.

<a id="ci-a9cf852b04a73ed291ca58e6"></a>

- Canonical ID: `ci-a9cf852b04a73ed291ca58e6` (1 source records).
  - [circleci/t30mkznmqskh](../../data/2026/W18/circleci.json); vendor URL: https://stspg.io/f6kcq9nb85wt
- Start time basis: `vendor_reported`.

<a id="ci-8cd491fa5c718c100f3c9dd7"></a>

- Canonical ID: `ci-8cd491fa5c718c100f3c9dd7` (1 source records).
  - [stripe/jnvb0qp0c7pd](../../data/2026/W18/stripe.json); vendor URL: https://stspg.io/3x0dw69bql3k
- Start time basis: `vendor_reported`.

<a id="ci-6839b46b334262e9c00550d8"></a>

- Canonical ID: `ci-6839b46b334262e9c00550d8` (1 source records).
  - [slack/1548](../../data/2026/W18/slack.json); vendor URL: https://slack-status.com/2026-04/4bd5196b0c972734
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
