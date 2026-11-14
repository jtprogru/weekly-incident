# Week 2026-W26 — top 5 of 7 above threshold

2026-06-22 to 2026-06-28 (UTC).

16 source records (16 saved revisions) → 16 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## 1. github — Disruption with some GitHub services

- URL: https://stspg.io/h85w6wz46mvk
- Window: 2026-06-27 14:02 to 2026-06-27 20:33 UTC (6h 30m)
- Impact: minor, status: resolved
- Components: —

<a id="ci-1b5eee852daa125aba07a3cc"></a>

- Canonical ID: `ci-1b5eee852daa125aba07a3cc` (1 source records).
  - [github/v0b3bpsyvqtk](../../data/2026/W26/github.json); vendor URL: https://stspg.io/h85w6wz46mvk
- Start time basis: `vendor_reported`.

  - `2026-06-27 14:02` investigating — We are investigating reports of impacted performance for some GitHub services.
  - `2026-06-27 14:02` investigating — We are conducting routine maintenance on our network infrastructure in the EU. This will not impact production traffic, but may result in slightly increased latency for the remainder of our work. We expect this to last until 19:00 UTC.
  - `2026-06-27 14:04` monitoring — The degradation has been mitigated. We are monitoring to ensure stability.
  - `2026-06-27 18:22` monitoring — This is progressing as expected, but we're going to extend the maintenance window by a couple of hours. Our new expected completion time is 21:00 UTC.
  - `2026-06-27 20:19` monitoring — Maintenance has completed and we've begun normalizing traffic in EU. We will monitor for a while longer before confirming resolution.
  - `2026-06-27 20:33` resolved — This incident was used to notify for a maintenance event. There is no specific root cause analysis. Work progressed as planned without any issues to report.

## 2. github — Disruption with some GitHub services

- URL: https://stspg.io/v6kqsytzxg2j
- Window: 2026-06-28 17:50 to 2026-06-28 20:55 UTC (3h 5m)
- Impact: minor, status: resolved
- Components: —

<a id="ci-68a50d6d4a276ff97efe9b57"></a>

- Canonical ID: `ci-68a50d6d4a276ff97efe9b57` (1 source records).
  - [github/0rkjjs2ssp7z](../../data/2026/W26/github.json); vendor URL: https://stspg.io/v6kqsytzxg2j
- Start time basis: `vendor_reported`.

  - `2026-06-28 17:50` investigating — We are investigating reports of impacted performance for some GitHub services.
  - `2026-06-28 17:59` investigating — Copilot cloud agent is experiencing intermittent problems with opening pull requests, pushing changes and replying to comments. We have identified the issue and are validating a fix.
  - `2026-06-28 18:31` investigating — Copilot cloud agent is experiencing intermittent problems with opening pull requests, pushing changes and replying to comments. We have validated a fix and are deploying that fix now.
  - `2026-06-28 20:02` investigating — Copilot cloud agent had been experiencing intermittent problems with opening pull requests, pushing changes and replying to comments. A fix has been deployed and we are validating that fix.
  - `2026-06-28 20:55` resolved — From June 26, 2026 at 23:40 UTC through June 28, 2026 at 20:55 UTC, Copilot Cloud Agent was degraded. The agent could fail when reporting progress, replying to pull request comments, or opening pull requests. For affected built-in tool calls, the average error rate was approximately 8%, with hourly error rates peaking around 26%.<br /><br />This was due to a regression introduced during a Copilot Cloud Agent runtime deployment that caused several built-in agent tools to become unavailable. In many cases, the affected tool calls failed silently so agent jobs appeared to succeed. This monitoring gap meant it took longer than expected to identify the failure. We mitigated the incident by reverting the runtime deployment to the previously stable version.<br /><br />We've added monitoring and a…

## 3. circleci — Issues with network routing for Mac jobs

- URL: https://stspg.io/mrq29l3fw2cm
- Window: 2026-06-24 23:13 to 2026-06-25 03:06 UTC (3h 52m)
- Impact: minor, status: resolved
- Components: macOS Jobs

<a id="ci-91c9280c704e772cecddf7c4"></a>

- Canonical ID: `ci-91c9280c704e772cecddf7c4` (1 source records).
  - [circleci/gvysjmkf4ct9](../../data/2026/W26/circleci.json); vendor URL: https://stspg.io/mrq29l3fw2cm
- Start time basis: `vendor_reported`.

  - `2026-06-24 23:13` investigating — What's happening Our engineers are investigating network routing issues for some Mac jobs. What's impacted Mac jobs that make outbound requests, including checking out from GitHub.
  - `2026-06-24 23:33` identified — We have implemented a change to address the routing issues and are monitoring the results.
  - `2026-06-24 23:56` identified — Mac jobs are still seeing intermittent failures with network routing. We are working with our infrastructure provider to address the networking issues.
  - `2026-06-25 00:39` identified — We are seeing improvement in Mac job network routing. We will continue to monitor for signs of recovery.
  - `2026-06-25 01:13` monitoring — Our team is moving this incident to Monitoring status based on significant improvement seen in our tests and low error rate. We are seeing improvement in network routing for Mac jobs and are actively engaged with our infrastructure provider. If your Mac jobs experience network-related failures or degraded service please reach out to CircleCI support.
  - `2026-06-25 01:22` monitoring — We are continuing to monitor for any further issues.
  - `2026-06-25 02:20` monitoring — We are continuing to monitor. We still observe a low percentage of intermittent failures. The failure rate is trending down as we continue to monitor. We advise customers to contact CircleCI Support if they see a change in degraded services and increased failure rate.
  - `2026-06-25 02:54` investigating — Due to continued reports of intermittent failures and a lack of sustained improvement while in monitoring status, this incident is elevated back to investigating to escalate with our infrastructure provider.
  - `2026-06-25 02:58` monitoring — A fix has been implemented and we are monitoring the results.
  - `2026-06-25 03:06` resolved — We are concluding this incident based on a confirmed fix with our infrastructure provider and successful tests. For any customers seeing failed runs, please run again for a successful connection.

## 4. github — Degradation with Webhooks, Pull Requests and Actions

- URL: https://stspg.io/t882b0gccf1h
- Window: 2026-06-25 17:50 to 2026-06-25 18:27 UTC (37m)
- Impact: major, status: resolved
- Components: Actions, Issues, Pull Requests, Webhooks

<a id="ci-fa4036707296a3abe6a9a819"></a>

- Canonical ID: `ci-fa4036707296a3abe6a9a819` (1 source records).
  - [github/9ndxtnrwjf37](../../data/2026/W26/github.json); vendor URL: https://stspg.io/t882b0gccf1h
- Start time basis: `vendor_reported`.

  - `2026-06-25 17:50` investigating — We are investigating reports of degraded performance for Actions, Pull Requests and Webhooks
  - `2026-06-25 17:53` investigating — Actions is experiencing degraded availability. We are continuing to investigate.
  - `2026-06-25 17:58` investigating — Issues is experiencing degraded performance. We are continuing to investigate.
  - `2026-06-25 18:07` monitoring — The degradation affecting Actions, Issues, Pull Requests and Webhooks has been mitigated. We are monitoring to ensure stability.
  - `2026-06-25 18:27` monitoring — We identified an issue that caused degradation across multiple services including Webhooks, Pull Requests, Actions, and Issues. Customers may have experienced delays or failures with these services. We have applied mitigations and affected services have recovered.
  - `2026-06-25 18:27` resolved — On June 25, 2026, between 17:33 UTC and 17:55 UTC, our background job service experienced degradation which increased delays to pull requests, repository pushes, Actions workflows, and Webhooks, with delays peaking at 7m. The issue was caused by underlying hypervisor issues and an incoming traffic spike, causing service timeouts which led to a connection storm and continual rebalances. <br /><br />The issue was mitigated by replacing the problem node at 17:49, after which all services saw recovery by 18:07.

## 5. yandexcloud — Errors when accessing Object Storage

- URL: https://status.yandex.cloud/en/incidents/1851
- Window: 2026-06-23 17:00 to 2026-06-23 18:34 UTC (1h 34m)
- Impact: major, status: resolved
- Components: Object Storage

<a id="ci-3a4214d2d0687c84cbb209cb"></a>

- Canonical ID: `ci-3a4214d2d0687c84cbb209cb` (1 source records).
  - [yandexcloud/1851](../../data/2026/W26/yandexcloud.json); vendor URL: https://status.yandex.cloud/en/incidents/1851
- Start time basis: `vendor_reported`.

  - `2026-06-23 18:31` investigating — We noticed errors when accessing Object Storage in Yandex Cloud. The developers are already aware and are fixing them.
  - `2026-06-23 18:48` unknown — The Object Storage API worked correctly, but there were errors when accessing the website.yandexcloud.net. The developers continue to investigate the situation.
  - `2026-06-23 19:12` resolved — At 21:34, errors occurred when accessing the website.yandexcloud.net they were resolved. If you still have questions about cloud resources, please contact technical support.

## Also above threshold (2)

- `circleci` Delays starting Docker jobs (1h 23m) — https://stspg.io/dw1r4sr0mrtl
- `datadog` Metrics Queries (1h 13m) — https://stspg.io/z9s5hk7jc45c
