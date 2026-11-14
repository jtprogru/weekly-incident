# Week 2026-W13 — top 3 of 3 above threshold

2026-03-23 to 2026-03-29 (UTC).

7 source records (7 saved revisions) → 7 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## 1. datadog — Delayed Monitors Notifications

- URL: https://stspg.io/vl9tsrk9z2ld
- Window: 2026-03-26 14:24 to 2026-03-26 19:56 UTC (5h 32m)
- Impact: major, status: resolved
- Components: Monitors

<a id="ci-75741144df53ae276cbd94c6"></a>

- Canonical ID: `ci-75741144df53ae276cbd94c6` (1 source records).
  - [datadog/1qdv7gx89vwv](../../data/2026/W13/datadog.json); vendor URL: https://stspg.io/vl9tsrk9z2ld
- Start time basis: `vendor_reported`.

  - `2026-03-26 14:24` investigating — We are investigating delays in monitor notifications, which began at 12:00 UTC.
  - `2026-03-26 14:55` investigating — We are continuing investigating delays in monitor notifications that started at 12:00 UTC.
  - `2026-03-26 15:32` investigating — We are still continuing investigating delays in monitor notifications that started at 12:00 UTC. We've identified the root cause and are working on a fix.
  - `2026-03-26 16:11` investigating — We are continuing to investigate delays in monitor notifications.
  - `2026-03-26 16:48` identified — The issue has been identified and a fix is being implemented.
  - `2026-03-26 17:28` identified — We are continuing to work on a fix for this issue.
  - `2026-03-26 18:38` identified — We are continuing to work on a fix for this issue.
  - `2026-03-26 19:07` identified — We are continuing to work on a fix for this issue.
  - `2026-03-26 19:42` monitoring — A fix has been implemented and we are monitoring the results.
  - `2026-03-26 19:56` resolved — This incident has been resolved.

## 2. slack — Issue with Sending and Receiving Messages

- URL: https://slack-status.com/2026-03/6da2f3fe8e16319d
- Window: 2026-03-29 22:38 to 2026-03-30 01:17 UTC (2h 38m)
- Impact: unknown, status: resolved
- Components: Messaging

<a id="ci-5927d3838a2eea28a6dbacdd"></a>

- Canonical ID: `ci-5927d3838a2eea28a6dbacdd` (1 source records).
  - [slack/1540](../../data/2026/W13/slack.json); vendor URL: https://slack-status.com/2026-03/6da2f3fe8e16319d
- Start time basis: `vendor_reported`.

  - `2026-03-29 22:38` unknown — We're investigating a new incident. Current impact to end users: Some users may experience elevated error rates and message-related degradations when sending and receiving messages across the Slack platform Teams are actively engaged to identify the root cause and remediate the service impact and will provide updates as information is available. We apologize for the inconvenience and will continue sharing updates as they become available.
  - `2026-03-29 23:30` unknown — We’re continuing to investigate the issue. We have identified two affected shards that are becoming overloaded, causing elevated query failures and message-related degradations. We are in the process of testing a potential fix. We’ll provide the next update once new information becomes available.
  - `2026-03-30 00:29` unknown — A fix has been deployed to the affected shards. ProcessMessage jobs had been paused, and these have now been unpaused on the affected shards. We are monitoring performance to assess whether the shards have returned to expected performance levels.
  - `2026-03-30 01:17` unknown — Following a period of monitoring, our engineering team called this incident resolved as of 01:10 UTC, March 30. We apologize for any inconvenience this issue has caused to users.

## 3. datadog — Delayed Traces in APM Trace Search

- URL: https://stspg.io/4qqxn4k2w7n5
- Window: 2026-03-27 18:57 to 2026-03-27 19:30 UTC (32m)
- Impact: major, status: resolved
- Components: APM

<a id="ci-37f0c49ab8e006bcf69ed82f"></a>

- Canonical ID: `ci-37f0c49ab8e006bcf69ed82f` (1 source records).
  - [datadog/y1trpqc54210](../../data/2026/W13/datadog.json); vendor URL: https://stspg.io/4qqxn4k2w7n5
- Start time basis: `vendor_reported`.

  - `2026-03-27 18:57` investigating — We are investigating increased latency in processing and storing Traces in APM. As a result of this issue, some users may see missing or delayed traces in APM Trace Search since 5pm UTC. They may also experience delay in APM trace-based monitors.
  - `2026-03-27 19:17` identified — The issue has been identified and a fix is being implemented.
  - `2026-03-27 19:21` monitoring — A fix has been implemented and we are monitoring the results.
  - `2026-03-27 19:30` resolved — This incident has been resolved.
