# Week 2026-W24 — top 5 of 6 above threshold

2026-06-08 to 2026-06-14 (UTC).

15 source records (15 saved revisions) → 15 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## 1. slack — Trouble adding multiple users to channels

- URL: https://slack-status.com/2026-06/767f9891b3838e3f
- Window: 2026-06-09 15:24 to 2026-06-10 21:43 UTC (1d 6h 18m)
- Impact: unknown, status: resolved
- Components: Workspace/Org Administration

<a id="ci-0c0a528ec7cd9fc0ec327c5b"></a>

- Canonical ID: `ci-0c0a528ec7cd9fc0ec327c5b` (1 source records).
  - [slack/1563](../../data/2026/W24/slack.json); vendor URL: https://slack-status.com/2026-06/767f9891b3838e3f
- Start time basis: `vendor_reported`.

  - `2026-06-09 15:24` unknown — Some users may experience issues when attempting to add multiple users to channels. Multiple matches error may appear along with a failure to identify team members. This may cause subsequent workflows not to trigger. We're investigating and will let you know as soon as we know more.
  - `2026-06-09 17:39` unknown — Our work on this issue is still ongoing. Current status and actions being taken: We have identified the trigger for this issue and a fix is in the process of deploying. The progress is being monitored and we’ll provide an update as soon as the fix is completely deployed and verified. We'll be back with another update once we have news to share.
  - `2026-06-09 20:18` unknown — This issue is now resolved for all users. Current status and recent actions taken: We identified the trigger for this issue and a fix was deployed. After a period of monitoring, we have concluded that there is no further impact to users. Previous impact to end users: Some users may have experienced issues when attempting to add multiple users to channels and received an error indicating a failure to identify team members. This could have caused subsequent workflows not to trigger. Required resolution steps: No further action is necessary. Thanks for sticking with us as we resolved this. We apologize for any disruptions to your day.
  - `2026-06-10 09:47` unknown — We've received reports that a subset of users may continue to experience issues when attempting to add multiple users to channels using email addresses. These users may encounter a "Multiple matches" error, preventing the intended user from being identified and added to the channel. We’re actively investigating these reports to determine whether additional remediation is required. We'll provide further updates as more information becomes available.
  - `2026-06-10 10:35` unknown — We’re working to reproduce the issue. This will help us ‌identify the underlying cause and provide a resolution. We'll provide further update as soon as more information becomes available.
  - `2026-06-10 11:44` unknown — We’ve identified the cause of the issue and we’re validating a fix to ensure it fully resolves the issue before we deploy it. We'll provide further update as soon as more information becomes available.
  - `2026-06-10 14:35` unknown — We've deployed a fix for the issue and are monitoring to confirm that the issue has been resolved. We'll provide another update as soon as more information becomes available.
  - `2026-06-10 21:43` unknown — We've resolved an issue that could cause incorrect email address tokens to appear when users pasted multiple email addresses into the member invite flow. In some cases, pasted email addresses were incorrectly matched to existing members with similar email addresses, requiring users to manually correct the entries before proceeding. We identified the cause of the issue and deployed a fix to ensure email addresses are now matched correctly when pasted into the invite flow. The issue is resolved for all affected users. If you're still experiencing any trouble, please reload Slack using Command + Shift + R (Mac) or Ctrl + Shift + R (Windows/Linux). We apologize for the disruption and appreciate your patience.

## 2. slack — Trouble with Gmail add-on Integration

- URL: https://slack-status.com/2026-06/8e2c9f5e6f18c289
- Window: 2026-06-10 20:36 to 2026-06-11 06:39 UTC (10h 3m)
- Impact: unknown, status: resolved
- Components: Apps/Integrations/APIs

<a id="ci-f0bf528214a9c092f80216a0"></a>

- Canonical ID: `ci-f0bf528214a9c092f80216a0` (1 source records).
  - [slack/1564](../../data/2026/W24/slack.json); vendor URL: https://slack-status.com/2026-06/8e2c9f5e6f18c289
- Start time basis: `vendor_reported`.

  - `2026-06-10 20:36` unknown — We're aware of an issue impacting the Slack Gmail add-on integration. As a result, users are unable to interact with the add-on and may encounter a “Something went wrong when executing the add-on. For help, contact the add-on developer” error message when attempting to use it. We are investigating the cause of this issue. We'll provide an update in 30 minutes or sooner if additional information becomes available.
  - `2026-06-10 21:05` unknown — We’re still actively investigating the issue impacting the Slack Gmail add-on. We are working to identify the cause, and no workaround is currently available. We’ll provide an update in 30 minutes or sooner if additional information becomes available.
  - `2026-06-10 21:37` unknown — We’re still investigating the issue impacting Slack Gmail add-ons. No additional news to share just yet, but we’re focused on getting things back to normal as quickly as we can. We apologize for the continued trouble. We’ll provide an update in 60 minutes or sooner if additional information becomes available.
  - `2026-06-10 22:49` unknown — We’re still investigating the issue impacting Slack Gmail add-ons. No additional news to share just yet, but we’re focused on getting things back to normal as quickly as we can. We apologize for the continued trouble. We’ll provide an update in 60 minutes or sooner if additional information becomes available.
  - `2026-06-10 23:39` unknown — We continue to investigate the issue impacting Slack Gmail add-ons. We've identified the cause and our teams are working on a fix to resolve the issue as soon as possible. We’ll provide an update once additional information becomes available.
  - `2026-06-11 00:18` unknown — We’re actively working on the fix to resolve the issue. We estimate this will take approximately four hours or more to complete, and validate the fix, before it can be deployed. As such, the next status update will be once the fix has been completed and validated, or we have new information to share. We apologize for any inconvenience this issue is causing for users.
  - `2026-06-11 06:04` unknown — We’re continuing to investigate the issues with the Slack Gmail add-on and are actively working on a fix to resolve the issue. There’s no new information to share at this time, but we’ll provide an update as soon as we have more details on implementing the fix. We apologize for any inconvenience this issue is causing for users.
  - `2026-06-11 06:39` unknown — We've resolved the issue that prevented customers from using the Slack Gmail add-on. The fix has been successfully deployed; health metrics have stabilized, and the Gmail add-on is now fully operational for all affected users. We apologize for any inconvenience this may have caused.

## 3. datadog — Delayed Evaluation

- URL: https://stspg.io/j0xmm8v17tmx
- Window: 2026-06-12 17:12 to 2026-06-12 19:39 UTC (2h 27m)
- Impact: minor, status: resolved
- Components: APM, Metrics and Infra Monitoring, Monitors

<a id="ci-4b867524298f47b09826146b"></a>

- Canonical ID: `ci-4b867524298f47b09826146b` (1 source records).
  - [datadog/3fnvmf6kckds](../../data/2026/W24/datadog.json); vendor URL: https://stspg.io/j0xmm8v17tmx
- Start time basis: `vendor_reported`.

  - `2026-06-12 17:12` identified — We are investigating delays in monitor evaluation, which began at 16:10 UTC.
  - `2026-06-12 17:59` identified — We are continuing to work on a fix for this issue.
  - `2026-06-12 18:30` identified — We are continuing to work on a fix for this issue.
  - `2026-06-12 18:39` identified — We are continuing to work on a fix for this issue.
  - `2026-06-12 19:01` identified — We are continuing to work on a fix for this issue.
  - `2026-06-12 19:08` identified — We are continuing to work on a fix for this issue. During this time, the following components are affected: Monitor evaluation Metrics and Infra Monitoring APM metrics Estimated usage metrics
  - `2026-06-12 19:23` monitoring — A fix has been implemented and we are monitoring the results.
  - `2026-06-12 19:39` resolved — This incident has been resolved.

## 4. stripe — Elevated card declines for some users in Mexico

- URL: https://stspg.io/rfnl7whdr3cr
- Window: 2026-06-10 05:17 to 2026-06-10 07:41 UTC (2h 23m)
- Impact: none, status: resolved
- Components: Acquirers and payment methods

<a id="ci-ce1951a252715017fccf5dbd"></a>

- Canonical ID: `ci-ce1951a252715017fccf5dbd` (1 source records).
  - [stripe/7mb3cx0lfctr](../../data/2026/W24/stripe.json); vendor URL: https://stspg.io/rfnl7whdr3cr
- Start time basis: `vendor_reported`.

  - `2026-06-10 05:17` investigating — As of 04:16 UTC, some card transactions in Mexico are experiencing elevated declines due to an issuing bank being temporarily unavailable. We're closely monitoring the issue and will post an update once we have observed recovery.
  - `2026-06-10 07:41` resolved — Between 04:16 – 07:06 UTC, we saw elevated declines with some card transactions in Mexico due to an issuing bank being temporarily unavailable. This is now resolved.

## 5. circleci — Issue authenticating only Bitbucket users

- URL: https://stspg.io/qjt275qv8plh
- Window: 2026-06-11 21:31 to 2026-06-12 00:18 UTC (2h 46m)
- Impact: minor, status: resolved
- Components: Pipelines & Workflows

<a id="ci-ddbad23456671965483bae8c"></a>

- Canonical ID: `ci-ddbad23456671965483bae8c` (1 source records).
  - [circleci/brp558x5kfn1](../../data/2026/W24/circleci.json); vendor URL: https://stspg.io/qjt275qv8plh
- Start time basis: `vendor_reported`.

  - `2026-06-11 21:31` investigating — We are currently investigating this issue.
  - `2026-06-11 22:05` identified — What's happening We are investigating an issue we believe is related to an ongoing incident with Atlassian Bitbucket's authentication service. We are actively monitoring Bitbucket's status. What's impacted Customers using Bitbucket as their VCS provider may be experiencing disruptions. What can you expect You may see errors when CircleCI attempts to interact with your Bitbucket repositories. Thank you for your patience while our engineers investigate. We will provide an update as soon as we have more information.
  - `2026-06-11 22:57` identified — What's happening Atlassian Bitbucket has acknowledged an incident affecting their OAuth service, which is the cause of the errors our customers are experiencing. The status of their incident is available at: https://bitbucket.status.atlassian.com/incidents/h51sr145gsmx What's impacted Customers using Bitbucket as their VCS provider may be experiencing disruptions. What can you expect You may see errors when CircleCI attempts to interact with your Bitbucket repositories. Customers using Bitbucket as their login provider may also be unable to authenticate to CircleCI. We will provide an update as soon as we have more information. Thank you for your patience while we monitor Bitbucket's incident and recovery.
  - `2026-06-12 00:10` monitoring — What's happening Atlassian Bitbucket has moved their OAuth service incident to monitoring. You can follow their status at https://bitbucket.status.atlassian.com/incidents/h51sr145gsmx What's impacted Customers using Bitbucket as their VCS provider may be experiencing disruptions. What can you expect You may see errors when CircleCI attempts to interact with your Bitbucket repositories. Customers using Bitbucket as their login provider may also be unable to authenticate to CircleCI. Thank you for your patience while we continue to monitor recovery.
  - `2026-06-12 00:18` resolved — Customers using Bitbucket as their VCS provider or as their login provider experienced errors when interacting with Bitbucket repositories and authenticating to CircleCI. This was caused by an incident with Atlassian Bitbucket's OAuth service, which has now been resolved and all affected functionality has returned to normal.

## Also above threshold (1)

- `digitalocean` DeepSeek V4 Pro is returning HTTP 429 "Rate limit exceeded" (3h 17m) — https://stspg.io/r8mzkb94pf9g
