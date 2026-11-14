# Week 2026-W22 — top 5 of 7 above threshold

2026-05-25 to 2026-05-31 (UTC).

8 source records (8 saved revisions) → 8 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## 1. slack — Trouble with Slack - General Slack functionality issues

- URL: https://slack-status.com/2026-05/46d9d5d41fbedd9f
- Window: 2026-05-27 23:00 to 2026-05-28 15:22 UTC (16h 21m)
- Impact: unknown, status: resolved
- Components: Apps/Integrations/APIs, Canvases, Messaging

<a id="ci-400a96f7032fa078718ac923"></a>

- Canonical ID: `ci-400a96f7032fa078718ac923` (1 source records).
  - [slack/1556](../../data/2026/W22/slack.json); vendor URL: https://slack-status.com/2026-05/46d9d5d41fbedd9f
- Start time basis: `vendor_reported`.

  - `2026-05-27 23:00` unknown — We're investigating a new incident. Current impact to end users: We’re aware of an issue affecting workflows, message edits, channel renames, apps, and channel creations. We'll provide an update within 30 minutes including further details on current actions being taken.
  - `2026-05-27 23:31` unknown — We continue to investigate the issue, and are‌ implementing fixes to restore services. We're starting to see some improvements from the fixes implemented so far. We’ll provide another update as soon as we have more information, and again we are sorry for the inconvenience this is causing.
  - `2026-05-28 01:03` unknown — Teams continue to implement configuration adjustments and deployment fixes to resolve the issue. All features except for Workflows and Messaging should be recovered. Some lingering impact remains. Users may notice: Empty pills instead of @-mention Delayed reacji Mismatch in Channel member count and channel membership Slow loading of messages, user avatars, and unfurls We are working with urgency to improve platform stability and reduce impact. We’ll provide an update in the next 30 minutes or sooner, if additional information becomes available.
  - `2026-05-28 01:41` unknown — We are continuing to work on implementing fixes to fully resolve the issue. Messaging should now be recovered. Users may continue to experience issues with workflow steps completing. Empty pills instead of @-mention; delayed reacji; mismatch in Channel member count and channel membership; and slow loading of messages, user avatars, and unfurls should now all be working properly. We’ll provide an update in the next 30 minutes or sooner, if additional information becomes available.
  - `2026-05-28 02:17` unknown — Messaging latency has now cleared. Newly triggered workflows should start and complete as normal. Workflows that executed during the impact period ending at approximately 7pm PDT may not have completed. Customers may notice delayed workflow execution for jobs that triggered during the disruption period. We’ll provide an update in the next 30 minutes or sooner, if additional information becomes available.
  - `2026-05-28 02:46` unknown — Workflow builder jobs that were triggered during the disruption period have been replayed and will complete over the next several hours. We’ll provide an update in the next 30 minutes or sooner, if additional information becomes available.
  - `2026-05-28 03:22` unknown — Cleanup and recovery efforts continue. Workflow builder job backlogs are being drained. Event subscriptions have been re-enabled. No further impact is expected. We’ll provide an update in the next 30 minutes or sooner, if additional information becomes available.
  - `2026-05-28 03:55` unknown — We’re glad to inform you that the Workflow Builder job backlogs were cleared and workflow execution has returned to normal. We continue progressing through the final stages of recovery activities across affected service areas as overall platform stability improves. We’ll provide an update in the next 30 minutes or sooner, if additional information becomes available.
  - `2026-05-28 04:30` unknown — Recovery activities continue, and restoration efforts are progressing well. Platform stability is improving, and no further customer impact is expected. We’ll provide an update in the next 60 minutes or sooner, if additional information becomes available.
  - `2026-05-28 05:13` unknown — Slack services continue to progress through the final stages of ‌recovery activities to restore functionality across all services. We’ll provide an update in the next 30 minutes or sooner, if additional information becomes available.
  - `2026-05-28 05:47` unknown — Slack services have largely returned to normal operational levels. Messaging, Workflows, Apps, Integrations, APIs, and Connections have all been restored. We'll continue with the final restoration activities and we'll provide updates as significant developments occur.
  - `2026-05-28 10:04` unknown — We're investigating a new impact to Slack. Current impact to end users: Slack users are experiencing issues with managing channels, statuses aren't clearing or updating correctly, and edits on Canvas aren't being saved. We'll provide an update within 30 minutes including further details on current actions being taken.
  - `2026-05-28 10:45` unknown — We continue to investigate the issue and are working with urgency to improve stability and reduce impact. We’ll provide another in 30 minutes or sooner, including further details on current actions being taken. We're sorry for the inconvenience this is causing.
  - `2026-05-28 11:17` unknown — Users continue to experience issues with their ability to manage channels, Canvas edits not saving, and user statuses not clearing or updating correctly. Investigations are underway. We're working with urgency to improve platform stability and reduce impact. We’ll provide an update in the next 30 minutes or sooner, if additional information becomes available.
  - `2026-05-28 11:46` unknown — We continue to investigate the issue, and are‌ gathering the resources necessary to restore services. We’ll provide another update in 30 minutes or as soon as we have more information, and again we're sorry for the inconvenience this is causing.
  - `2026-05-28 12:01` unknown — Our work on this issue is still ongoing. Current status and actions being taken: Remediation efforts are underway via configuration adjustments and deployment. We're still working to resolve the issues with Canvas. We'll be back with another update once we have news to share.
  - `2026-05-28 12:25` unknown — Recovery activities continue, and restoration efforts are progressing well. Slack users should now be able to edit Canvases as the errors have dropped. We’ll provide an update in the next 30 minutes or sooner, if additional information becomes available.
  - `2026-05-28 12:55` unknown — Remediation efforts via configuration adjustments and deployment are in progress. Users may still run into issues updating their status and managing channels. We’ll provide another update in 30 minutes or as soon as we have more information, and again we are sorry for the inconvenience this is causing.
  - `2026-05-28 13:35` unknown — Remediation efforts via configuration adjustments and deployment will take several hours to complete to fully resolve channel management and status issues. Slack users may continue to experience intermittent impact. We'll be back with another update in 30 minutes when we have news to share.
  - `2026-05-28 14:09` unknown — Slack users continue to experience permissions-related issues such as: unable to manage channels, add Canvas to tabs, update statuses, convert public to private channels while we work on a complete restoration. We’ll provide an update in the next 30 minutes or sooner, if additional information becomes available.
  - `2026-05-28 15:22` unknown — This issue is now resolved for all users. Current status and recent actions taken: We deployed a series of configuration adjustments and infrastructure fixes to restore services. Messaging, Workflows, Apps, and Connections were progressively restored over the last several hours. Some users continued to experience permissions-related issues, including inability to manage channels and Canvas edits not saving, as a result of residual stale cache data. Cache deployment activities completed at 14:13 UTC on May 28, 2026, fully resolving the issue. Scope: This issue affected Slack users globally. Previous impact to end users: Slack users experienced difficulties accessing Slack features, including Messaging, Workflows, Apps, Integrations, APIs, Canvas, and Connections, beginning at 22:21 UTC on M…

## 2. yandexcloud — Temporary unavailability of services

- URL: https://status.yandex.cloud/en/incidents/1782
- Window: 2026-05-27 12:22 to 2026-05-27 12:36 UTC (14m)
- Impact: major, status: resolved
- Components: API Gateway, Application Load Balancer, Audit Trails, Certificate Manager, Cloud Apps, Cloud Backup, Cloud CDN, Cloud DNS, Cloud Desktop, Cloud Functions, Cloud Interconnect, Cloud Logging, Cloud Notification Service, Cloud Registry, Compute Cloud, Container Registry, Data Processing, Data Streams, Data Transfer, DataLens, DataSphere, Forms, GOST Gateway, Identity and Access Management, IoT Core, Key Management Service, Load Testing, Managed Service for Apache Spark™, Managed Service for OpenSearch, Managed Service for Prometheus®, Managed Service for Sharded PostgreSQL, Managed Service for YDB, Managed Service for YTsaurus, Managed Service for Apache Kafka®, Managed Service for ClickHouse®, Managed Service for GitLab, Managed Service for Kubernetes®, Managed Service for MySQL®, Managed Service for PostgreSQL, Managed Service for Valkey™, Marketplace, Message Queue, Monitoring, Monium, Monium Logs, Monium Metrics, Monium Traces, Network Load Balancer, Object Storage, Resource Manager, Security Deck, Serverless Containers, Serverless Integrations, Smart Web Security, SmartCaptcha, SourceCraft, SourceCraft Code Assistant, SpeechKit, SpeechSense, Stackland, Support Center, Tracker, Translate, Virtual Private Cloud, Vision OCR, Wiki, Yandex AI Studio, Yandex BareMetal, Yandex Cloud Billing, Yandex Cloud Console, Yandex Cloud Detection and Response, Yandex Cloud Postbox, Yandex Cloud Quota Manager, Yandex Cloud Video, Yandex Identity Hub, Yandex Lockbox, Yandex MPP Analytics Engine for PostgreSQL, Yandex Managed Service for Apache Airflow®, Yandex Managed Service for Trino, Yandex MetaData Hub, Yandex Neurosupport, Yandex Query, Yandex Search API, Yandex StoreDoc, Yandex WebSQL, YandexGPT API

<a id="ci-744de9288fd6dd0996a466c6"></a>

- Canonical ID: `ci-744de9288fd6dd0996a466c6` (1 source records).
  - [yandexcloud/1782](../../data/2026/W22/yandexcloud.json); vendor URL: https://status.yandex.cloud/en/incidents/1782
- Start time basis: `vendor_reported`.

  - `2026-05-27 12:29` investigating — There is a short-term loss of connectivity between external services and Yandex Cloud. The problem is localized at the external contour level: the operation of internal services and the data-plane is not disrupted. The development team is already working on the restoration.
  - `2026-05-27 12:43` unknown — At 15:33 Moscow time (UTC+3), we are seeing the restoration of most services. The developers continue to investigate the situation.
  - `2026-05-27 12:58` resolved — The availability of external services has been fully restored. The cause of the unavailability has been fixed. The operation of the data-plane and internal systems remained stable throughout the incident. The problem is not expected to happen again. If you have any questions or problems, please contact technical support.

## 3. circleci — Mac jobs failing to start

- URL: https://stspg.io/fd7x5v1g2kw9
- Window: 2026-05-31 18:15 to 2026-06-01 00:45 UTC (6h 29m)
- Impact: minor, status: resolved
- Components: macOS Jobs

<a id="ci-b5120b92f48d7074f1cb5cec"></a>

- Canonical ID: `ci-b5120b92f48d7074f1cb5cec` (1 source records).
  - [circleci/t7vt8j14jnbg](../../data/2026/W22/circleci.json); vendor URL: https://stspg.io/fd7x5v1g2kw9
- Start time basis: `vendor_reported`.

  - `2026-05-31 18:15` identified — We have identified a network outage that is preventing Mac jobs from starting. This include m4pro.medium and m4pro.large resource types. We are working with our infrastructure provider to restore service. We will update this incident as we work to resolve the issue.
  - `2026-05-31 19:09` identified — We are continuing to work on resolving the networking issue for Mac jobs.
  - `2026-05-31 20:05` identified — We are still seeing failures with Mac jobs starting due to high network latency and packet loss. We will continue to update as we have more information on a resolution.
  - `2026-05-31 21:24` identified — We are continuing to investigate failures with Mac jobs starting due to high network latency and packet loss. Our engineers remain engaged with our infrastructure provider to restore service. We will provide another update within the hour, or sooner if circumstances change. We thank you for your patience while our engineers work to resolve this.
  - `2026-05-31 22:38` identified — Mac jobs are continuing to experience failures. Our infrastructure provider is reporting they are experiencing a DDoS attack causing high network latency and packet loss for our Mac fleet. We are working to bring up additional capacity to restore service. We will continue to update as more information is available.
  - `2026-06-01 00:24` monitoring — We are seeing Mac jobs starting successfully and improved network connectivity to the Mac fleet. We will continue to monitor the recovery.
  - `2026-06-01 00:42` monitoring — We have processed the backlog of Mac jobs and job start times are recovering to normal parameters.
  - `2026-06-01 00:45` resolved — Mac jobs have now fully recovered. Thank you for your patience during this incident.

## 4. digitalocean — App Platform Deployments

- URL: https://stspg.io/c5cpdwx9rc37
- Window: 2026-05-29 00:14 to 2026-05-29 07:58 UTC (7h 44m)
- Impact: minor, status: resolved
- Components: Global

<a id="ci-2893bf3fd510dc864d0a4c9f"></a>

- Canonical ID: `ci-2893bf3fd510dc864d0a4c9f` (1 source records).
  - [digitalocean/y3p9v8vm4tfn](../../data/2026/W22/digitalocean.json); vendor URL: https://stspg.io/c5cpdwx9rc37
- Start time basis: `vendor_reported`.

  - `2026-05-29 00:14` investigating — Our Engineering team is currently investigating an issue with build failures on App Platform in multiple regions. Users may experience errors when attempting to build their applications, resulting in failed deployments. Our Engineering team is working to fix the issue and will share an update once we have more details. We apologize for the inconvenience this issue may be causing and appreciate your patience as we work to resolve it.
  - `2026-05-29 01:58` monitoring — Our Engineering team has implemented a fix to resolve the issue with build failures on App Platform. Users should see their builds deploy successfully now. We are closely monitoring the situation, and will post an update once we've confirmed this is fully resolved.
  - `2026-05-29 07:58` resolved — Our Engineering team has confirmed that the issue impacting build failures on App Platform has been resolved at 07:00 am UTC. All App Platform builds are now succeeding as expected. Customers who previously encountered build failures should now be able to deploy their applications without further issues. If you continue to experience any problems, please open a ticket with our support team. Thank you for your patience, and we apologize for any inconvenience.

## 5. circleci — Usage API data delayed for 5/28

- URL: https://stspg.io/vc8xbt2r7rmw
- Window: 2026-05-29 14:51 to 2026-05-29 18:10 UTC (3h 18m)
- Impact: none, status: resolved
- Components: CircleCI API

<a id="ci-09e195b42c448ddb971f501f"></a>

- Canonical ID: `ci-09e195b42c448ddb971f501f` (1 source records).
  - [circleci/dqlq3v1q8pvn](../../data/2026/W22/circleci.json); vendor URL: https://stspg.io/vc8xbt2r7rmw
- Start time basis: `vendor_reported`.

  - `2026-05-29 14:51` investigating — We are currently investigating this issue.
  - `2026-05-29 14:57` investigating — We've identified a small number of duplicate records in the Usage API data. Data is available, but may be inaccurate. We continue to investigate.
  - `2026-05-29 18:10` resolved — The Usage API incident has been resolved.

## Also above threshold (2)

- `yandexcloud` Reduced performance when working with deployment-optimized disks. (2h 3m) — https://status.yandex.cloud/en/incidents/1786
- `yandexcloud` Unavailability of the Cloud Backup service (2h 0m) — https://status.yandex.cloud/en/incidents/1797
