# Week 2026-W38 — top 5 of 21 above threshold

2026-09-14 to 2026-09-20 (UTC).

## 1. aws — Increased Error Rates

- URL: https://health.aws.amazon.com/health/status
- Window: 2026-09-15 10:22 to not yet (6d 3h 5m (ongoing))
- Impact: unknown, status: investigating
- Components: multipleservices-me-south-1

  - `2026-09-15 10:22` investigating — We are providing an update on the disruption affecting the Middle East (Bahrain) (me-south-1) Region. The damage to our infrastructure spanned multiple Availability Zones and exceeded what our regional and multi-AZ services are designed to withstand. After a thorough assessment, we have determined that we are unable to restore access to the resources and data hosted exclusively in this Region. When the first Availability Zone was damaged in March, we began recommending that customers migrate their workloads to other Regions, and most did so before the Region became unavailable following the disruption of a second Availability Zone in April. Since then, we have supported the remaining customers in re-establishing their operations in alternate Regions, using backups where available or imple…

## 2. aws — Increased Error Rates

- URL: https://health.aws.amazon.com/health/status
- Window: 2026-09-15 10:27 to not yet (6d 3h 0m (ongoing))
- Impact: unknown, status: investigating
- Components: multipleservices-me-central-1

  - `2026-09-15 10:27` investigating — We are providing an update on the disruption affecting the Middle East (UAE) (me-central-1) Region. After a thorough assessment, we have determined that we are unable to restore access to the resources and data hosted exclusively in the mec1-az2 Availability Zone. We continue to work on recovering regional resources, as well as zonal resources hosted in the other affected Availability Zones (mec1-az1 and mec1-az3). Since the disruption began in March, most customers have been able to re-establish their operations in other Regions by restoring backups or copying data that remained accessible. AWS Support remains available to help customers who need assistance moving their applications to alternate Regions. We remain committed to supporting our customers in the UAE. We are working on repla…

## 3. cloudflare — Cloudflare Queues issues

- URL: https://www.cloudflarestatus.com/incidents/s4x7n3nfqnyy
- Window: 2026-09-14 20:22 to 2026-09-15 22:31 UTC (1d 2h 8m)
- Impact: minor, status: resolved
- Components: Queues

  - `2026-09-14 20:22` investigating — Cloudflare is aware of investigating issues with our Cloudflare Queues product.
  - `2026-09-15 02:01` identified — We are investigating an issue where messages published to some queues are not being delivered to their consumers, resulting in delivery delays. Publishing to queues is unaffected, and messages are being received and safely retained — no messages are being lost. We have identified the root cause and are actively working on a fix. Affected messages will be delivered automatically once the issue is resolved.
  - `2026-09-15 15:14` monitoring — A fix has been implemented and we are monitoring the results.
  - `2026-09-15 22:31` resolved — This incident has been resolved.

## 4. cloudflare — Replicate Pruna issue

- URL: https://www.cloudflarestatus.com/incidents/823rbdm2k25x
- Window: 2026-09-17 01:09 to 2026-09-17 21:51 UTC (20h 42m)
- Impact: minor, status: resolved
- Components: Replicate

  - `2026-09-17 01:09` identified — Some Pruna models may not be able to scale up and are therefore either experiencing long queue times or entirely unavailable. We are working with Pruna to resolve the issue.
  - `2026-09-17 21:51` resolved — This incident has been resolved.

## 5. cloudflare — Some Replicate models are failing to start

- URL: https://www.cloudflarestatus.com/incidents/19g1m7tsvncw
- Window: 2026-09-14 22:25 to 2026-09-15 14:06 UTC (15h 40m)
- Impact: minor, status: resolved
- Components: Replicate

  - `2026-09-14 22:25` investigating — We are investigating an issue which is preventing Replicate models from failing to start new replicas.
  - `2026-09-15 00:03` identified — Some third party models as well as https://replicate.com/qwen/qwen-edit-multiangle are affected and are unable to scale out. Please use alternate models while we work with providers to publish fixed versions.
  - `2026-09-15 14:06` resolved — This incident has been resolved.

## Also above threshold (16)

- `cloudflare` Network Performance Issues in Los Angeles (14h 20m) — https://www.cloudflarestatus.com/incidents/qc2v4gxy3gfw
- `stripe` Elevated errors on Google Pay payments (7h 46m) — https://stspg.io/j7b7wsrfq7hm
- `openai` Elevated errors affecting Work Mode in ChatGPT (11h 12m) — https://status.openai.com/incidents/01M2GA8XTS6VB3QCDEGZ0HNAQ5
- `github` Degradation with Gemini 3.8 Flash (10h 27m) — https://stspg.io/hsjy12f57b24
- `openai` Overbilling for OpenAI-hosted containers in the Agent API (9h 22m) — https://status.openai.com/incidents/01M2VA7X37P1ASADSNZ1CG4N4D
- `cloudflare` Support Helpdesk Availability Issues (6h 26m) — https://www.cloudflarestatus.com/incidents/zlyf7lvpp05s
- `circleci` Customers may experience delays in UI updates (4h 34m) — https://stspg.io/40g5hyq2p6j4
- `digitalocean` Support Portal Service Disruption (11h 21m) — https://stspg.io/n6hr8jd64qjl
- `anthropic` Intermittent error spikes for Claude Mythos 5.1 and Claude Fable 5.1 (1h 26m) — https://stspg.io/x1qpycpg6ln7
- `stripe` Stripe Support Systems issue (7h 4m) — https://stspg.io/1pqrkshz4vdd
- `cloudflare` Load Balancing Analytics and Notification Delays (3h 23m) — https://www.cloudflarestatus.com/incidents/rnmq02lhv5w1
- `cloudflare` Issues with 1.1.1.1 public resolver on .tn queries (3h 11m) — https://www.cloudflarestatus.com/incidents/9tgkg3ww89kf
- `openai` Delayed support responses (3h 33m) — https://status.openai.com/incidents/01M2V76GEPJB0HQGRA0QERRZ3T
- `circleci` Customers may experience delays in UI updates (2h 20m) — https://stspg.io/0djh5p3ll96t
- `openai` Elevated errors affecting ChatGPT Work mode (1h 21m) — https://status.openai.com/incidents/01M2RC0WHGTX8EHDH7D60JJ7FC
- `openai` Elevated errors with gpt-image-2.5-flare (40m) — https://status.openai.com/incidents/01M2KQNE5C42NEZPX6V01NHH5W

