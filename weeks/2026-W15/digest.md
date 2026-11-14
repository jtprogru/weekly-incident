# Week 2026-W15

2026-04-06 to 2026-04-12 (UTC). Generated 2026-11-14 00:00.

6 canonical incidents, 5 above threshold. All 14 sources responded.

6 source records (6 saved revisions) → 6 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | slack | [Trouble with Slackbot and AI-powered features](https://slack-status.com/2026-04/2e11633c5821fcd0) | 2026-04-09 07:58 | 11h 30m | unknown | Connectivity, Workflows | 1656 = 690 × 2 × 1.2 |
| 2 | circleci | [Delays in Workflow and Job data in the Usage API](https://stspg.io/b6482c23v6nr) | 2026-04-08 13:20 | 7h 13m | minor | CircleCI API | 433 = 433 × 1 × 1.0 |
| 3 | datadog | [Notifications to Pagerduty, VictorOps, Slack, Webhook, Microsoft Teams and OpsGenie are not delivered](https://stspg.io/18986pkc2dgr) | 2026-04-10 14:54 | 6h 5m | major | Monitors | 365 = 365 × 1 × 1.0 |
| 4 | slack | [Trouble Sending Messages](https://slack-status.com/2026-04/a982d547f16ab676) | 2026-04-10 19:09 | 2h 35m | unknown | Messaging | 186 = 155 × 1 × 1.2 |
| 5 | circleci | [Degradation in CircleCI UI elements](https://stspg.io/3qm63nb96zjy) | 2026-04-06 20:33 | 39m | major | CircleCI UI | 39 = 39 × 1 × 1.0 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. slack — Trouble with Slackbot and AI-powered features

- URL: https://slack-status.com/2026-04/2e11633c5821fcd0
- Started: 2026-04-09 07:58 UTC
- Resolved: 2026-04-09 19:29 UTC
- Duration: 11h 30m
- Impact: unknown, status: resolved
- Components: Connectivity, Workflows
- Canonical incident: [ci-329c303051e04f868417a8fb](#ci-329c303051e04f868417a8fb)

- `2026-04-09 07:58` **unknown** — We're investigating a new incident. Current impact to end users: Some users may be experiencing trouble with Slackbot responses. Other Slack functionality that leverages AI services may also be experiencing issues. We're on the case and we'll provide an update within 30 minutes including further details on current actions being taken.
- `2026-04-09 08:28` **unknown** — Our work on this issue is still ongoing. Current status and actions being taken: Users may continue to experience trouble with Slackbot responses and other Slack AI-powered features. We've identified the trigger as a third-party infrastructure issue and are working with the provider to resolve this issue. We'll provide another update in 30 minutes until the impact is resolved for all users.
- `2026-04-09 08:59` **unknown** — Our work on this issue is still ongoing. Users may continue to experience intermittent trouble with Slackbot responses and other Slack AI-powered features such as Search Answers, Meeting Notes, and Translation. Current status and actions being taken: We remain engaged with our third-party provider. We'll provide another update in 60 minutes until the impact is resolved for all users.
- `2026-04-09 09:51` **unknown** — Current status and actions being taken: We're continuing to engage with our third-party provider to find a solution. We're also exploring multiple avenues on our side to try to alleviate impact on Slackbot and AI-powered features until the upstream issue is resolved. We'll provide another update in 60 minutes until the impact is resolved for all users.
- `2026-04-09 10:54` **unknown** — Current status and actions being taken: We've taken steps to try to reduce the impact on Slackbot and its response capabilities. Customers may notice signs of improvement, however, the impact may still be intermittent until full resolution. Other AI-powered features, including Search Answers, Meeting Notes, and Translation, are still experiencing issues due to the ongoing upstream problem. We're investigating similar alleviative actions for these features. We'll provide another update in 60 minutes until the impact is resolved for all users.
- `2026-04-09 11:56` **unknown** — Current status and actions being taken: We're continuing to engage with our third-party provider to find a solution. No further alleviative actions have been undertaken since our last update. We'll provide another update in 60 minutes until the impact is resolved for all users.
- `2026-04-09 12:58` **unknown** — Current status and actions being taken: We believe we've identified an issue that's still causing impact to occur even after our upstream provider has resolved impact from their end. We expect a code fix to take at least an hour to complete and pass through various test environments. We'll provide another update in 60 minutes.
- `2026-04-09 13:59` **unknown** — Current status and actions being taken: We've begun deploying the code fix to production. Once complete we'll begin monitoring the performance of Slackbot, as well as our other AI-powered features. We'll provide another update in 60 minutes.
- `2026-04-09 14:52` **unknown** — Current status and actions being taken: The code fix hasn't successfully addressed the impact and we're resuming our investigations. We'll provide another update in 60 minutes.
- `2026-04-09 15:55` **unknown** — Current status and actions being taken: Upon further review, we've identified issues with secondary caching and are currently working on a configuration fix. In parallel, we're continuing to investigate and assess several solutions to restore ‌services. We'll provide another update in 60 minutes.
- `2026-04-09 17:11` **unknown** — Current status and actions being taken: We’ve deployed a configuration fix, resulting in significant improvements to Slackbot success rates and health metrics. We're now monitoring the situation as we investigate a secondary cache issue that may be contributing to intermittent errors or degraded workflow performance. We'll provide another update in 60 minutes.
- `2026-04-09 18:13` **unknown** — Current status and actions being taken: We've deployed an emergency fix to address the issues contributing to intermittent errors with Slackbot and Workflows. We're now seeing significant recovery across all affected services and are continuing to monitor the situation closely to ensure performance remains stable. We'll provide another update in 60 minutes.
- `2026-04-09 19:29` **unknown** — Current status and recent actions taken: After monitoring system performance and infrastructure stability, we have confirmed that all Slack features are operating within expected levels. This incident is now resolved and users should no longer experience any issues with Slackbot, AI Features, or workflows. We apologize for any disruptions to your day.

### 2. circleci — Delays in Workflow and Job data in the Usage API

- URL: https://stspg.io/b6482c23v6nr
- Started: 2026-04-08 13:20 UTC
- Resolved: 2026-04-08 20:33 UTC
- Duration: 7h 13m
- Impact: minor, status: resolved
- Components: CircleCI API
- Canonical incident: [ci-d7db9f771b9eb1d60ce023c4](#ci-d7db9f771b9eb1d60ce023c4)

- `2026-04-08 13:20` **identified** — We have identified an issue causing Insights workflow and job data to be delayed when requested through the Usage API. Some data for the last 24 hours may not be fully reflected while we work to process the backlog of events. Insights UI is not impacted by this issue. Thank you for your patience.
- `2026-04-08 20:33` **resolved** — The Usage API incident is resolved on an ongoing basis. There was a 4-hour period yesterday 4/7 for which we did not capture RAM and CPU utilization in our logs. We will not be backfilling the data for that period. Appreciate your patience.

### 3. datadog — Notifications to Pagerduty, VictorOps, Slack, Webhook, Microsoft Teams and OpsGenie are not delivered

- URL: https://stspg.io/18986pkc2dgr
- Started: 2026-04-10 14:54 UTC
- Resolved: 2026-04-10 20:59 UTC
- Duration: 6h 5m
- Impact: major, status: resolved
- Components: Monitors
- Canonical incident: [ci-c878375696ccda4284faea49](#ci-c878375696ccda4284faea49)

- `2026-04-10 14:54` **investigating** — We are investigating an issue to send notifications to Pagerduty, VictorOps, Slack, Webhook, Microsoft Teams and OpsGenie.
- `2026-04-10 15:00` **identified** — The issue has been identified and a fix is being implemented.
- `2026-04-10 15:23` **identified** — We are continuing to work on a fix for this issue.
- `2026-04-10 15:42` **identified** — New notifications are now successfully delivered, and we are catching-up on past notifications that may have been delayed.
- `2026-04-10 15:42` **monitoring** — New notifications are now successfully delivered, and we are catching-up on past notifications that may have been delayed
- `2026-04-10 15:59` **monitoring** — We are still catching-up on past notifications that may have been delayed.
- `2026-04-10 16:15` **monitoring** — We are still catching-up on past notifications that may have been delayed.
- `2026-04-10 17:12` **monitoring** — Notifications for Opsgenie, VictorOps, BigPanda, Zendesk, and Sumologic are caught up. Webhooks, Microsoft Teams, Slack, JIRA, and ServiceNow are still catching up.
- `2026-04-10 17:37` **monitoring** — Notifications for JIRA are now caught up.
- `2026-04-10 18:33` **monitoring** — Notifications for ServiceNow and Microsoft Teams are now caught up. Webhooks and Slack notifications are still catching up.
- `2026-04-10 19:02` **monitoring** — Webhooks and Slack notifications are still catching up.
- `2026-04-10 19:44` **monitoring** — Webhooks and Slack notifications are still catching up.
- `2026-04-10 20:13` **monitoring** — Webhooks and Slack notifications are still catching up.
- `2026-04-10 20:59` **resolved** — Between approximately 9:42 AM and 10:38 AM ET on April 10, we observed delivery failures for customers in the US1, AP1, and AP2 regions. During this window, notifications to the following integrations were delayed or temporarily undelivered: • PagerDuty • Slack • Microsoft Teams • Webhooks • Jira • ServiceNow • OpsGenie • VictorOps • BigPanda • Zendesk • Sumologic All delayed notifications have been replayed with the following exceptions: PagerDuty pages and Slack notifications queued during the outage window were not replayed, as they were no longer actionable at the time of recovery. Webhook notifications will be reprocessed and delivered at a later date. No monitoring data was lost during this incident. Monitor evaluations continued normally — only the delivery of triggered notifications was affected. Thank you for your patience.

### 4. slack — Trouble Sending Messages

- URL: https://slack-status.com/2026-04/a982d547f16ab676
- Started: 2026-04-10 19:09 UTC
- Resolved: 2026-04-10 21:44 UTC
- Duration: 2h 35m
- Impact: unknown, status: resolved
- Components: Messaging
- Canonical incident: [ci-8696a1908f924950eaa662ef](#ci-8696a1908f924950eaa662ef)

- `2026-04-10 19:09` **unknown** — We're investigating a new incident. Current impact to end users: Some users may experience trouble sending messages. We're investigating the issue and we regret any inconvenience this is causing. We'll provide an update within 30 minutes including further details on current actions being taken.
- `2026-04-10 19:44` **unknown** — We are actively working to resolve the current platform functionality issues. Current status and actions being taken: We've identified a fix and are actively deploying it across the platform. Once the rollout is complete, we'll monitor system performance to ensure full stability. Estimated time to resolution: No ETA at this time. Scope: This issue isn't affecting any specific region, but a percentage of users may be affected globally. Current impact to end users: Some users may experience trouble sending messages. Known workarounds: No known workarounds at this time. We appreciate your patience as we sort this out.
- `2026-04-10 20:21` **unknown** — We are actively working to resolve the current platform functionality issues. Current status and actions being taken: We have processed the affected service PODs to expedite recovery, resulting in significant recovery across the platform. We're now monitoring all systems to ensure performance remains stable. Estimated time to resolution: No ETA at this time. Scope: This issue isn't affecting any specific region, but a percentage of users may be affected globally. Current impact to end users: Some users may experience trouble sending messages. Known workarounds: No known workarounds at this time. We'll continue to provide updates every 60 minutes until the impact is resolved for all users.
- `2026-04-10 21:22` **unknown** — We are actively working to resolve the current platform functionality issues. Current status and actions being taken: We're still working toward resolving the issue and to reduce the likelihood of symptoms recurring. We'll continue to monitor internal metrics as we work through this. Estimated time to resolution: No ETA at this time. Scope: This issue isn't affecting any specific region, but a percentage of users may be affected globally. Current impact to end users: Some users may experience trouble sending messages. Known workarounds: No known workarounds at this time. We'll provide the next update upon a significant change in status, or sooner if additional information becomes available.
- `2026-04-10 21:44` **unknown** — This issue is now resolved for all users. Current status and recent actions taken: After monitoring system performance and infrastructure stability, we have confirmed that all Slack features are operating within expected levels. Users should no longer experience any issues when sending messages. Scope: This issue wasn't affecting any specific region, but a percentage of users may have been affected globally. Previous impact to end users: Some users may have experienced trouble sending messages. Required resolution steps: No additional steps are required. We apologize for any disruptions to your day.

### 5. circleci — Degradation in CircleCI UI elements

- URL: https://stspg.io/3qm63nb96zjy
- Started: 2026-04-06 20:33 UTC
- Resolved: 2026-04-06 21:12 UTC
- Duration: 39m
- Impact: major, status: postmortem
- Components: CircleCI UI
- Canonical incident: [ci-a597a6d7ca8da256269c0483](#ci-a597a6d7ca8da256269c0483)

- `2026-04-06 20:33` **investigating** — We are currently investigating an issue in the CircleCI UI impacting approval jobs and SSH key navigation. This is an interface-only issue and builds continue to flow as usual.
- `2026-04-06 20:43` **monitoring** — A fix has been implemented and we are monitoring the results. Thank you for your patience.
- `2026-04-06 21:12` **resolved** — This incident has been resolved.
- `2026-04-07 00:17` **postmortem** — **What happened** All modals and dialogs across the CircleCI application became invisible. They were being rendered correctly in the background, but were not visible to users. This affected approval jobs, SSH key management, and other settings pages across the app. **Root cause** The issue was introduced by an upgrade to our internal design system library. The new version added animation support to modal components using a library called Framer Motion. These animations are designed to fade modals in from invisible \(`opacity: 0`\) to visible \(`opacity: 1`\) when opened. However, the animation logic reads open/close state from a specific React context that is only available when modals are opened via a particular component \(`DialogTrigger`\). Our application manages modal state differently, passing open/close state directly via props, so that context was never populated. As a result, every modal in the app was permanently stuck in the "hidden" animation state and never transitioned to visible. **Resolution** We reverted the design system upgrade to restore the previous working version, which resolved the issue immediately. We are now working on a follow-up fix that will upgrade Framer Motion properly, by ensuring all modal components receive the required context. **Impact** The issue affected all modal dialogs across project settings, org settings, user settings, plans & payments, and other areas of the app for the duration of the incident. No data was lost or corrupted. The underlying functionality was intact, just inaccessible via the UI.

## Below threshold (1)

- `digitalocean` [Droplet Availability in All Regions](https://stspg.io/nm9g257py11j) — 33m

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-329c303051e04f868417a8fb"></a>

- Canonical ID: `ci-329c303051e04f868417a8fb` (1 source records).
  - [slack/1542](../../data/2026/W15/slack.json); vendor URL: https://slack-status.com/2026-04/2e11633c5821fcd0
- Start time basis: `vendor_reported`.

<a id="ci-d7db9f771b9eb1d60ce023c4"></a>

- Canonical ID: `ci-d7db9f771b9eb1d60ce023c4` (1 source records).
  - [circleci/l7vz5yr3wrp0](../../data/2026/W15/circleci.json); vendor URL: https://stspg.io/b6482c23v6nr
- Start time basis: `vendor_reported`.

<a id="ci-c878375696ccda4284faea49"></a>

- Canonical ID: `ci-c878375696ccda4284faea49` (1 source records).
  - [datadog/q2njm9l08tjb](../../data/2026/W15/datadog.json); vendor URL: https://stspg.io/18986pkc2dgr
- Start time basis: `vendor_reported`.

<a id="ci-8696a1908f924950eaa662ef"></a>

- Canonical ID: `ci-8696a1908f924950eaa662ef` (1 source records).
  - [slack/1543](../../data/2026/W15/slack.json); vendor URL: https://slack-status.com/2026-04/a982d547f16ab676
- Start time basis: `vendor_reported`.

<a id="ci-314310cc83952d9969e43ba2"></a>

- Canonical ID: `ci-314310cc83952d9969e43ba2` (1 source records).
  - [digitalocean/wjyr2cz6y3xb](../../data/2026/W15/digitalocean.json); vendor URL: https://stspg.io/nm9g257py11j
- Start time basis: `vendor_reported`.

<a id="ci-a597a6d7ca8da256269c0483"></a>

- Canonical ID: `ci-a597a6d7ca8da256269c0483` (1 source records).
  - [circleci/2l1y8dhrvtzd](../../data/2026/W15/circleci.json); vendor URL: https://stspg.io/3qm63nb96zjy
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
