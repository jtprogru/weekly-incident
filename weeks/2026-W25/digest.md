# Week 2026-W25

2026-06-15 to 2026-06-21 (UTC). Generated 2026-11-14 00:00.

12 canonical incidents, 4 above threshold. All 14 sources responded.

12 source records (12 saved revisions) → 12 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | slack | [Trouble submitting workflow forms with file attachments on Android.](https://slack-status.com/2026-06/83210c13677ffff1) | 2026-06-16 15:50 | 6d 4h 49m | unknown | Workflows | 10715 = 8929 × 1 × 1.2 |
| 2 | slack | [Trouble loading SSO configuration settings for some workspace admins](https://slack-status.com/2026-06/66c001cc76d073f8) | 2026-06-17 22:33 | 5h 41m | unknown | Login/SSO | 409 = 341 × 1 × 1.2 |
| 3 | digitalocean | [Monitoring Graphs in the Cloud Control Panel](https://stspg.io/8jy6pdvsszwm) | 2026-06-19 14:00 | 2h 37m | minor | Global | 126 = 157 × 1 × 0.8 |
| 4 | digitalocean | [Intermittent 500 Errors on Serverless/GenAI Inference API](https://stspg.io/40ln0gl1x56j) | 2026-06-16 22:00 | 2h 28m | none | Inference | 118 = 148 × 1 × 0.8 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. slack — Trouble submitting workflow forms with file attachments on Android.

- URL: https://slack-status.com/2026-06/83210c13677ffff1
- Started: 2026-06-16 15:50 UTC
- Resolved: 2026-06-22 20:39 UTC
- Duration: 6d 4h 49m
- Impact: unknown, status: resolved
- Components: Workflows
- Canonical incident: [ci-71e689789f34b24cb1970646](#ci-71e689789f34b24cb1970646)

- `2026-06-16 15:50` **unknown** — We are aware that some users on the Slack Android app are unable to submit workflow forms that contain a required file upload field. After a file is selected, the Submit button becomes disabled and the form cannot be completed. Users on Slack for iOS and Slack for web are not affected. We are actively investigating the issue and working to identify the root cause. We'll provide an update in 60 minutes or sooner if additional information becomes available.
- `2026-06-16 16:48` **unknown** — We have identified the cause of this issue and confirmed it affects a small subset of users on a single version of the Slack Android app. Users on other Android versions, iOS, and web are not affected. A fix is already available and will begin rolling out to Android users tomorrow. We will provide an update as soon as additional information becomes available.
- `2026-06-17 19:38` **unknown** — A fix for this issue has been delivered in a new version of the Slack Android app, and rollout has now begun. Android users will start seeing the updated app version become available in the Google Play Store as the rollout progresses. Once updated, users who were experiencing issues submitting workflow forms containing a required file upload field will be able to do so successfully. We will continue to monitor the rollout and provide another update as additional information becomes available.
- `2026-06-18 22:30` **unknown** — The rollout of the new version of the Slack Android app is ongoing and may take up to several days to fully complete. We continue to monitor the rollout and will provide another update when we have more information to share. Thank you for your patience with us while we work through this.
- `2026-06-21 17:33` **unknown** — We continue to monitor the rollout of the new version of Slack for Android. We'll provide another update when we have further information to share. Thank you for your continued patience.
- `2026-06-22 20:39` **unknown** — This issue has been resolved. Android users who were experiencing issues submitting workflow forms containing a required file upload field can now do so successfully. A fix was delivered via a new update to the Slack Android app. Users who are still experiencing this issue should update the Slack app from the Google Play Store to resolve it. No other action is required. We apologize for the inconvenience and thank you for your patience.

### 2. slack — Trouble loading SSO configuration settings for some workspace admins

- URL: https://slack-status.com/2026-06/66c001cc76d073f8
- Started: 2026-06-17 22:33 UTC
- Resolved: 2026-06-18 04:14 UTC
- Duration: 5h 41m
- Impact: unknown, status: resolved
- Components: Login/SSO
- Canonical incident: [ci-cf621911c5ab80e448160010](#ci-cf621911c5ab80e448160010)

- `2026-06-17 22:33` **unknown** — Some workspace admins on standalone (non-Enterprise Grid) workspaces may be experiencing an issue where the SSO and authentication configuration page does not display current settings as expected. Admins may see a blank value for the "Require SSO authentication for" field on the admin SSO settings page, which does not accurately reflect the actual backend configuration — SSO settings remain intact and in effect. We are actively investigating this issue and will provide an update as soon as we have more information to share. We apologize for the inconvenience.
- `2026-06-18 01:49` **unknown** — We have identified the trigger and developed a fix. We are currently deploying this fix and expect to resolve the issue shortly. We will provide another update once the deployment is complete. We apologize for the inconvenience.
- `2026-06-18 04:14` **unknown** — We have successfully deployed the fix to production and the issue is resolved. The Admin SSO and Authentication Configuration page for standalone workspaces should now display the correct settings. If you continue to see outdated information, refresh the SSO Admin page to load the latest changes.

### 3. digitalocean — Monitoring Graphs in the Cloud Control Panel

- URL: https://stspg.io/8jy6pdvsszwm
- Started: 2026-06-19 14:00 UTC
- Resolved: 2026-06-19 16:37 UTC
- Duration: 2h 37m
- Impact: minor, status: resolved
- Components: Global
- Canonical incident: [ci-ab19514f13cb1137863242a4](#ci-ab19514f13cb1137863242a4)

- `2026-06-19 14:00` **investigating** — Our Engineering team is currently investigating an issue affecting the visibility of monitoring graphs within the Cloud Panel. During this period, users may notice missing or unavailable monitoring graphs for services such as Droplets(with DO agent installed), Load Balancers, Databases, etc. We apologize for the inconvenience caused. We'll update once we have more information
- `2026-06-19 14:50` **monitoring** — Our engineering team has implemented the necessary fixes to address the issue affecting the visibility of monitoring graphs within the Cloud Panel. Users should now be able to view monitoring graphs for their services, including Droplets with DO Agent installed, Load Balancers, Databases, etc. We are currently monitoring the situation to ensure that the service has returned to normal operation and remain stable. We appreciate your patience and will provide an update once the issue is fully confirmed as resolved.
- `2026-06-19 16:37` **resolved** — Between 13:22 & 14:26 UTC today, users have experienced missing monitoring graphs for Droplets (with DO agent installed), Load Balancers, Databases, and other services within the Cloud Control Panel. Our engineering team has identified the root cause of the issue and has taken appropriate steps to restore functionality. We can confirm that services have been restored and are functioning as expected. We apologize for any inconvenience this may have caused. If you continue to experience issues viewing monitoring graphs, please create a support ticket for further analysis. Thank you for your patience and understanding

### 4. digitalocean — Intermittent 500 Errors on Serverless/GenAI Inference API

- URL: https://stspg.io/40ln0gl1x56j
- Started: 2026-06-16 22:00 UTC
- Resolved: 2026-06-17 00:29 UTC
- Duration: 2h 28m
- Impact: none, status: resolved
- Components: Inference
- Canonical incident: [ci-b9b1f1438d415a36a37cbc47](#ci-b9b1f1438d415a36a37cbc47)

- `2026-06-16 22:52` **investigating** — We are actively investigating an issue causing elevated HTTP 500 error rates for customers utilizing our Serverless/GenAI Inference API. Customer Impact: Customers making calls to the inference API—specifically targeting /v1/* endpoints—will experience intermittent HTTP 500 errors and failed requests.
- `2026-06-16 23:29` **monitoring** — Our engineering teams have successfully begun implementing mitigation steps to resolve the connectivity issues affecting the inference API. We will provide another update once the API has fully recovered and error rates return to normal.
- `2026-06-17 00:29` **resolved** — The connectivity issues affecting our Serverless/GenAI Inference API have been fully resolved. Our engineering teams successfully completed the connectivity restoration. All systems should be functioning normally, and the endpoint should be fully operational.

## Below threshold (8)

- `github` [Disruption with Copilot next edit suggestions](https://stspg.io/7tn8k5v6kzjj) — 1h 30m
- `yandexcloud` [Routine maintenance at the Cloud Backup provider](https://status.yandex.cloud/en/incidents/1842) — 55m
- `yandexcloud` [Planned infrastructure improvements.](https://status.yandex.cloud/en/incidents/1845) — 34m
- `slack` [Trouble performing some actions](https://slack-status.com/2026-06/2af2b628658941f3) — 28m
- `datadog` [Delayed Events](https://stspg.io/tdw3spf8slj1) — 20m
- `yandexcloud` [Network connectivity issue in the ru-central1-d zone](https://status.yandex.cloud/en/incidents/1843) — 5m
- `slack` [Slack Messaging issue](https://slack-status.com/2026-06/7504b505289fa2a9) — 0m
- `github` [Incident With Webhooks](https://stspg.io/t0qkhqpm5vlm) — unknown

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-71e689789f34b24cb1970646"></a>

- Canonical ID: `ci-71e689789f34b24cb1970646` (1 source records).
  - [slack/1568](../../data/2026/W25/slack.json); vendor URL: https://slack-status.com/2026-06/83210c13677ffff1
- Start time basis: `vendor_reported`.

<a id="ci-cf621911c5ab80e448160010"></a>

- Canonical ID: `ci-cf621911c5ab80e448160010` (1 source records).
  - [slack/1570](../../data/2026/W25/slack.json); vendor URL: https://slack-status.com/2026-06/66c001cc76d073f8
- Start time basis: `vendor_reported`.

<a id="ci-d64da6fdbb9935c42415cbaf"></a>

- Canonical ID: `ci-d64da6fdbb9935c42415cbaf` (1 source records).
  - [github/5t81zk0vrk3z](../../data/2026/W25/github.json); vendor URL: https://stspg.io/7tn8k5v6kzjj
- Start time basis: `vendor_reported`.

<a id="ci-ab19514f13cb1137863242a4"></a>

- Canonical ID: `ci-ab19514f13cb1137863242a4` (1 source records).
  - [digitalocean/nqkv7bnzr4l3](../../data/2026/W25/digitalocean.json); vendor URL: https://stspg.io/8jy6pdvsszwm
- Start time basis: `vendor_reported`.

<a id="ci-b9b1f1438d415a36a37cbc47"></a>

- Canonical ID: `ci-b9b1f1438d415a36a37cbc47` (1 source records).
  - [digitalocean/bfsh34myw3bc](../../data/2026/W25/digitalocean.json); vendor URL: https://stspg.io/40ln0gl1x56j
- Start time basis: `vendor_reported`.

<a id="ci-414516cb60cdec485303268d"></a>

- Canonical ID: `ci-414516cb60cdec485303268d` (1 source records).
  - [yandexcloud/1842](../../data/2026/W25/yandexcloud.json); vendor URL: https://status.yandex.cloud/en/incidents/1842
- Start time basis: `vendor_reported`.

<a id="ci-b5c7f7f344b95fbfd8b30dde"></a>

- Canonical ID: `ci-b5c7f7f344b95fbfd8b30dde` (1 source records).
  - [yandexcloud/1845](../../data/2026/W25/yandexcloud.json); vendor URL: https://status.yandex.cloud/en/incidents/1845
- Start time basis: `vendor_reported`.

<a id="ci-4af55f43fa9d27b79ec5ad8a"></a>

- Canonical ID: `ci-4af55f43fa9d27b79ec5ad8a` (1 source records).
  - [slack/1569](../../data/2026/W25/slack.json); vendor URL: https://slack-status.com/2026-06/2af2b628658941f3
- Start time basis: `vendor_reported`.

<a id="ci-9cc360844d6ad6a917f66d66"></a>

- Canonical ID: `ci-9cc360844d6ad6a917f66d66` (1 source records).
  - [datadog/7l0kwdnv43j8](../../data/2026/W25/datadog.json); vendor URL: https://stspg.io/tdw3spf8slj1
- Start time basis: `vendor_reported`.

<a id="ci-c133a5f11f30b7bb82113d7e"></a>

- Canonical ID: `ci-c133a5f11f30b7bb82113d7e` (1 source records).
  - [yandexcloud/1843](../../data/2026/W25/yandexcloud.json); vendor URL: https://status.yandex.cloud/en/incidents/1843
- Start time basis: `vendor_reported`.

<a id="ci-27c33ec0f1ea3b990881912c"></a>

- Canonical ID: `ci-27c33ec0f1ea3b990881912c` (1 source records).
  - [slack/1567](../../data/2026/W25/slack.json); vendor URL: https://slack-status.com/2026-06/7504b505289fa2a9
- Start time basis: `vendor_reported`.

<a id="ci-ea834c7c3ca6730ff3ed2b50"></a>

- Canonical ID: `ci-ea834c7c3ca6730ff3ed2b50` (1 source records).
  - [github/kn7gv3tlfc54](../../data/2026/W25/github.json); vendor URL: https://stspg.io/t0qkhqpm5vlm
- Start time basis: `vendor_reported`.
- Warning: Vendor resolved_at precedes started_at; duration is unknown, score uses zero duration; decision:legacy-time-github-kn7gv3tlfc54

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
