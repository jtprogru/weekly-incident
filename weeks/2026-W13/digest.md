# Week 2026-W13

2026-03-23 to 2026-03-29 (UTC). Generated 2026-11-14 00:00.

7 canonical incidents, 3 above threshold. All 14 sources responded.

7 source records (7 saved revisions) → 7 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | datadog | [Delayed Monitors Notifications](https://stspg.io/vl9tsrk9z2ld) | 2026-03-26 14:24 | 5h 32m | major | Monitors | 332 = 332 × 1 × 1.0 |
| 2 | slack | [Issue with Sending and Receiving Messages](https://slack-status.com/2026-03/6da2f3fe8e16319d) | 2026-03-29 22:38 | 2h 38m | unknown | Messaging | 190 = 158 × 1 × 1.2 |
| 3 | datadog | [Delayed Traces in APM Trace Search](https://stspg.io/4qqxn4k2w7n5) | 2026-03-27 18:57 | 32m | major | APM | 32 = 32 × 1 × 1.0 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. datadog — Delayed Monitors Notifications

- URL: https://stspg.io/vl9tsrk9z2ld
- Started: 2026-03-26 14:24 UTC
- Resolved: 2026-03-26 19:56 UTC
- Duration: 5h 32m
- Impact: major, status: resolved
- Components: Monitors
- Canonical incident: [ci-75741144df53ae276cbd94c6](#ci-75741144df53ae276cbd94c6)

- `2026-03-26 14:24` **investigating** — We are investigating delays in monitor notifications, which began at 12:00 UTC.
- `2026-03-26 14:55` **investigating** — We are continuing investigating delays in monitor notifications that started at 12:00 UTC.
- `2026-03-26 15:32` **investigating** — We are still continuing investigating delays in monitor notifications that started at 12:00 UTC. We've identified the root cause and are working on a fix.
- `2026-03-26 16:11` **investigating** — We are continuing to investigate delays in monitor notifications.
- `2026-03-26 16:48` **identified** — The issue has been identified and a fix is being implemented.
- `2026-03-26 17:28` **identified** — We are continuing to work on a fix for this issue.
- `2026-03-26 18:38` **identified** — We are continuing to work on a fix for this issue.
- `2026-03-26 19:07` **identified** — We are continuing to work on a fix for this issue.
- `2026-03-26 19:42` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-03-26 19:56` **resolved** — This incident has been resolved.

### 2. slack — Issue with Sending and Receiving Messages

- URL: https://slack-status.com/2026-03/6da2f3fe8e16319d
- Started: 2026-03-29 22:38 UTC
- Resolved: 2026-03-30 01:17 UTC
- Duration: 2h 38m
- Impact: unknown, status: resolved
- Components: Messaging
- Canonical incident: [ci-5927d3838a2eea28a6dbacdd](#ci-5927d3838a2eea28a6dbacdd)

- `2026-03-29 22:38` **unknown** — We're investigating a new incident. Current impact to end users: Some users may experience elevated error rates and message-related degradations when sending and receiving messages across the Slack platform Teams are actively engaged to identify the root cause and remediate the service impact and will provide updates as information is available. We apologize for the inconvenience and will continue sharing updates as they become available.
- `2026-03-29 23:30` **unknown** — We’re continuing to investigate the issue. We have identified two affected shards that are becoming overloaded, causing elevated query failures and message-related degradations. We are in the process of testing a potential fix. We’ll provide the next update once new information becomes available.
- `2026-03-30 00:29` **unknown** — A fix has been deployed to the affected shards. ProcessMessage jobs had been paused, and these have now been unpaused on the affected shards. We are monitoring performance to assess whether the shards have returned to expected performance levels.
- `2026-03-30 01:17` **unknown** — Following a period of monitoring, our engineering team called this incident resolved as of 01:10 UTC, March 30. We apologize for any inconvenience this issue has caused to users.

### 3. datadog — Delayed Traces in APM Trace Search

- URL: https://stspg.io/4qqxn4k2w7n5
- Started: 2026-03-27 18:57 UTC
- Resolved: 2026-03-27 19:30 UTC
- Duration: 32m
- Impact: major, status: resolved
- Components: APM
- Canonical incident: [ci-37f0c49ab8e006bcf69ed82f](#ci-37f0c49ab8e006bcf69ed82f)

- `2026-03-27 18:57` **investigating** — We are investigating increased latency in processing and storing Traces in APM. As a result of this issue, some users may see missing or delayed traces in APM Trace Search since 5pm UTC. They may also experience delay in APM trace-based monitors.
- `2026-03-27 19:17` **identified** — The issue has been identified and a fix is being implemented.
- `2026-03-27 19:21` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-03-27 19:30` **resolved** — This incident has been resolved.

## Below threshold (4)

- `slack` [Trouble connecting to Slack](https://slack-status.com/2026-03/8269f159feb011e3) — 55m
- `slack` [Issues while sending messages and loading message history is impacting some users](https://slack-status.com/2026-03/5e14cc3c13b8ae14) — 1h 27m
- `slack` [Authorization Failure Errors Connecting to the Slack MCP Sever](https://slack-status.com/2026-03/009e069324bea32e) — 1h 5m
- `stripe` [Elevated iDeal errors](https://stspg.io/224dj2tf08n9) — 15m

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-75741144df53ae276cbd94c6"></a>

- Canonical ID: `ci-75741144df53ae276cbd94c6` (1 source records).
  - [datadog/1qdv7gx89vwv](../../data/2026/W13/datadog.json); vendor URL: https://stspg.io/vl9tsrk9z2ld
- Start time basis: `vendor_reported`.

<a id="ci-5927d3838a2eea28a6dbacdd"></a>

- Canonical ID: `ci-5927d3838a2eea28a6dbacdd` (1 source records).
  - [slack/1540](../../data/2026/W13/slack.json); vendor URL: https://slack-status.com/2026-03/6da2f3fe8e16319d
- Start time basis: `vendor_reported`.

<a id="ci-6ff75d75b2e45a6d51f47943"></a>

- Canonical ID: `ci-6ff75d75b2e45a6d51f47943` (1 source records).
  - [slack/1538](../../data/2026/W13/slack.json); vendor URL: https://slack-status.com/2026-03/8269f159feb011e3
- Start time basis: `vendor_reported`.

<a id="ci-de938cd83af6948a92834869"></a>

- Canonical ID: `ci-de938cd83af6948a92834869` (1 source records).
  - [slack/1539](../../data/2026/W13/slack.json); vendor URL: https://slack-status.com/2026-03/5e14cc3c13b8ae14
- Start time basis: `vendor_reported`.

<a id="ci-b2d61fecd8015fa14fd19465"></a>

- Canonical ID: `ci-b2d61fecd8015fa14fd19465` (1 source records).
  - [slack/1537](../../data/2026/W13/slack.json); vendor URL: https://slack-status.com/2026-03/009e069324bea32e
- Start time basis: `vendor_reported`.

<a id="ci-37f0c49ab8e006bcf69ed82f"></a>

- Canonical ID: `ci-37f0c49ab8e006bcf69ed82f` (1 source records).
  - [datadog/y1trpqc54210](../../data/2026/W13/datadog.json); vendor URL: https://stspg.io/4qqxn4k2w7n5
- Start time basis: `vendor_reported`.

<a id="ci-96d3be6254c72e5794b8dc44"></a>

- Canonical ID: `ci-96d3be6254c72e5794b8dc44` (1 source records).
  - [stripe/tvh0drctzb1h](../../data/2026/W13/stripe.json); vendor URL: https://stspg.io/224dj2tf08n9
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
