# Week 2026-W11 — top 3 of 3 above threshold

2026-03-09 to 2026-03-15 (UTC).

4 source records (4 saved revisions) → 4 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## 1. slack — Intermittent SSO login errors accessing Slack workspace.

- URL: https://slack-status.com/2026-03/0c433efa2ea295fb
- Window: 2026-03-09 03:02 to 2026-03-09 19:17 UTC (16h 15m)
- Impact: unknown, status: resolved
- Components: Login/SSO

<a id="ci-ecc5a65c7cd10783b328c8dc"></a>

- Canonical ID: `ci-ecc5a65c7cd10783b328c8dc` (1 source records).
  - [slack/1536](../../data/2026/W11/slack.json); vendor URL: https://slack-status.com/2026-03/0c433efa2ea295fb
- Start time basis: `vendor_reported`.

  - `2026-03-09 03:02` unknown — Some users may experience trouble logging into Slack. We’re on the case and will provide another update once we have more details to share. We’re sorry for the disruption.
  - `2026-03-09 03:43` unknown — Our work on this issue is still ongoing. Current status and actions being taken: Some users are experiencing intermittent SSO login failures when trying to access their Slack workspace. Our team is investigating the cause of this issue and working to resolve it as quickly as possible. We'll be back with another update once we have news to share.
  - `2026-03-09 04:29` unknown — Our work on this issue is still ongoing. Current status and actions being taken: the team continues investigating. We'll provide another update in 60 minutes until impact is resolved for all users.
  - `2026-03-09 05:22` unknown — Current status and recent actions taken: Single Sign-On (SSO) login functionality has been fully restored. Following the deployment of a fix by our engineering team, users are now able to log in normally and without any issues. Scope: Previous impact to end users: customers were experiencing intermittent Single Sign-On (SSO) error accessing Slack workspaces. Required resolution steps: NA
  - `2026-03-09 07:19` unknown — This issue has resurfaced and we're actively engaged to resolve the issue as quickly as possible. Current status and actions being taken: We have identified the trigger ‌cause and are working on a comprehensive fix to prevent this from happening again. Estimated time to resolution: We are working urgently to resolve this issue as quickly as possible. Scope: A subset of Enterprise customers using Single Sign-On (SSO) are experiencing intermittent login issues. This is not affecting all users within impacted organizations, and the issue may vary by session. Current impact to end users: Users attempting to log in via SSO may encounter errors requiring them to refresh their browser multiple times before successfully accessing their workspace. Some users may experience a degraded login ex…
  - `2026-03-09 19:17` unknown — This issue is now resolved for all users. Current status and recent actions taken: We identified the cause as a login routing component that had incorrect entries, which prevented users from being directed to their workspace correctly. A fix has been deployed to address this issue. Previous impact to end users: Users attempting to log in via SSO may have encountered errors requiring them to refresh their browser multiple times before successfully accessing their workspace. Some users may have experienced a degraded login experience even after gaining access. We apologize for any disruptions to your day.

## 2. datadog — Web Application Not Loading

- URL: https://stspg.io/mfp6qzq5bd51
- Window: 2026-03-13 09:26 to 2026-03-13 10:30 UTC (1h 3m)
- Impact: critical, status: resolved
- Components: Web Application

<a id="ci-c81e3a8954c6b4ad425b7337"></a>

- Canonical ID: `ci-c81e3a8954c6b4ad425b7337` (1 source records).
  - [datadog/9ksfbdb59mtn](../../data/2026/W11/datadog.json); vendor URL: https://stspg.io/mfp6qzq5bd51
- Start time basis: `vendor_reported`.

  - `2026-03-13 09:26` investigating — We are investigating loading issues on our web application. As a result, some users might be getting errors when loading the web application.
  - `2026-03-13 09:27` investigating — We are continuing to investigate this issue.
  - `2026-03-13 09:32` investigating — We are continuing to investigate this issue.
  - `2026-03-13 09:52` investigating — We are continuing to investigate this issue.
  - `2026-03-13 10:00` identified — The issue has been identified and a fix is being implemented.
  - `2026-03-13 10:05` monitoring — A fix has been implemented and we are monitoring the results.
  - `2026-03-13 10:30` resolved — This incident has been resolved.

## 3. stripe — Errors downloading files

- URL: https://stspg.io/9swt2kr5kqlm
- Window: 2026-03-10 20:43 to 2026-03-10 20:53 UTC (9m)
- Impact: major, status: resolved
- Components: Stripe core components

<a id="ci-7691bf9a7c76e2615fc16765"></a>

- Canonical ID: `ci-7691bf9a7c76e2615fc16765` (1 source records).
  - [stripe/57rkhg708hrf](../../data/2026/W11/stripe.json); vendor URL: https://stspg.io/9swt2kr5kqlm
- Start time basis: `vendor_reported`.

  - `2026-03-10 20:43` investigating — We’re currently investigating an issue causing file downloads to fail with 404 errors. We’ve identified the cause, and are working on a fix. We’ll post an update soon.
  - `2026-03-10 20:53` resolved — Between 19:22 to 20:49 UTC on March 10th, file downloads were returning 404 errors. The issue is now resolved.
