# Week 2026-W18 — top 5 of 7 above threshold

2026-04-27 to 2026-05-03 (UTC).

10 source records (10 saved revisions) → 10 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## 1. aws — Increased Connectivity Issues and API Error Rates

- URL: https://health.aws.amazon.com/health/status
- Window: 2026-04-30 07:07 to not yet (198d 5h 29m (ongoing))
- Impact: unknown, status: investigating
- Components: multipleservices-me-south-1

<a id="ci-4f155127c97b96a474760521"></a>

- Canonical ID: `ci-4f155127c97b96a474760521` (1 source records).
  - [aws/dd4f7c1f4ca3fe41](../../data/2026/W18/aws.json); vendor URL: https://health.aws.amazon.com/health/status
- Start time basis: `first_observed_update`.
- Warning: AWS onset is the first observed message, not a verified impact start

  - `2026-04-30 07:07` investigating — We are providing an update on the ongoing service disruption. The Middle East (Bahrain) Region (ME-SOUTH-1) has suffered damage due to the conflict in the Middle East and is currently unavailable. Customers should recover their resources in other Regions from remote backups. Relevant billing operations are currently suspended while we restore normal operations in this AWS Region. This process is expected to take several months.

## 2. aws — Increased Error Rates

- URL: https://health.aws.amazon.com/health/status
- Window: 2026-04-30 07:25 to not yet (198d 5h 10m (ongoing))
- Impact: unknown, status: investigating
- Components: multipleservices-me-central-1

<a id="ci-0e9fcfad9f58dcd7e204aafb"></a>

- Canonical ID: `ci-0e9fcfad9f58dcd7e204aafb` (1 source records).
  - [aws/ca5f1cc2b6034dd6](../../data/2026/W18/aws.json); vendor URL: https://health.aws.amazon.com/health/status
- Start time basis: `first_observed_update`.
- Warning: AWS onset is the first observed message, not a verified impact start

  - `2026-04-30 07:25` investigating — We are providing an update on the ongoing service disruption. The Middle East (UAE) Region (ME-CENTRAL-1) has suffered damage as a result of the conflict in the Middle East and is currently unable to reliably support customer applications. While some workloads continue to function normally, we strongly recommend customers migrate all accessible resources to other Regions and restore inaccessible resources from remote backups as soon as possible. Relevant billing operations are currently suspended while we restore normal operations in this AWS Region. This process is expected to take several months.

## 3. digitalocean — Elevated 5xx “context canceled” errors impacting serverless inference

- URL: https://stspg.io/qf13crcjzkmn
- Window: 2026-04-28 13:45 to 2026-04-28 19:33 UTC (5h 48m)
- Impact: minor, status: resolved
- Components: Agent Runtime

<a id="ci-35367f529f55a4569db74813"></a>

- Canonical ID: `ci-35367f529f55a4569db74813` (1 source records).
  - [digitalocean/z5fjzx3kp16z](../../data/2026/W18/digitalocean.json); vendor URL: https://stspg.io/qf13crcjzkmn
- Start time basis: `vendor_reported`.

  - `2026-04-28 13:45` investigating — Serverless inference customers are experiencing elevated 5xx errors, including “context canceled” responses. This may result in intermittent request failures. Our team is actively investigating and will provide updates as more information becomes available.
  - `2026-04-28 15:59` identified — We have identified an issue affecting our service and are currently working to implement a fix. Our team is actively investigating and taking the necessary steps to restore normal operations as quickly as possible. We appreciate your patience and will provide updates as soon as more information becomes available.
  - `2026-04-28 19:00` monitoring — Service for Serverless Inference has been restored. We’ve implemented tighter rate limits to help prevent recurrence and are closely monitoring system performance. Some users may still experience intermittent latency as we complete final stabilization efforts. Our team remains actively engaged to ensure full recovery. We appreciate your patience and will provide further updates as needed.
  - `2026-04-28 19:33` resolved — All services are operating normally. We will continue to monitor the system to ensure ongoing reliability. Thank you for your patience while we worked to resolve this issue.

## 4. slack — Users May Not Receive Notifications on Slack Thread Replies

- URL: https://slack-status.com/2026-04/22195f9f2ca47237
- Window: 2026-04-27 17:11 to 2026-04-27 20:36 UTC (3h 24m)
- Impact: unknown, status: resolved
- Components: Notifications

<a id="ci-21a7f32d8d03f521c103d75b"></a>

- Canonical ID: `ci-21a7f32d8d03f521c103d75b` (1 source records).
  - [slack/1547](../../data/2026/W18/slack.json); vendor URL: https://slack-status.com/2026-04/22195f9f2ca47237
- Start time basis: `vendor_reported`.

  - `2026-04-27 17:11` unknown — Some users have reported issues with thread replies not generating a notification banner. We’re investigating the issue, and have identified a recent deployment as a possible trigger. We apologize for the inconvenience and will report back as more information becomes available.
  - `2026-04-27 18:07` unknown — We’re observing positive signals after rolling back the deployment in test environments. We will complete some final validations before rolling back the deployment in the production environment. Thank you for your patience as we work through this one. We’ll report back once production rollback is completed or if there are substantial updates to share.
  - `2026-04-27 20:36` unknown — We’ve successfully performed a rollback in the production environment, and all systems are now functioning as expected. Users should no longer experience any issues with thread replies not generating a notification banner. We apologize for any inconvenience and appreciate your patience.

## 5. circleci — Usage API data is delayed for data from yesterday, 4/30. All prior data is available. We are investigating.

- URL: https://stspg.io/ld03v9l2cm5c
- Window: 2026-05-01 10:04 to 2026-05-01 14:02 UTC (3h 57m)
- Impact: none, status: resolved
- Components: CircleCI API

<a id="ci-df5d09dc1620e3242436ac57"></a>

- Canonical ID: `ci-df5d09dc1620e3242436ac57` (1 source records).
  - [circleci/mspf5pqp7d42](../../data/2026/W18/circleci.json); vendor URL: https://stspg.io/ld03v9l2cm5c
- Start time basis: `vendor_reported`.

  - `2026-05-01 10:04` investigating — We are currently investigating
  - `2026-05-01 14:02` resolved — This issue has been resolved. All data is up to date through yesterday.

## Also above threshold (2)

- `digitalocean` Serverless Inference - Intermittent Rate Limiting Affecting Some Customers Using Anthropic Models (59m) — https://stspg.io/y8cgtdqrkct1
- `datadog` Elevated Error Rates for Metrics Queries (34m) — https://stspg.io/17n4grhvbmjt
