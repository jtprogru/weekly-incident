# Week 2026-W10

2026-03-02 to 2026-03-08 (UTC). Generated 2026-11-14 12:36.

9 canonical incidents, 4 above threshold. All 13 sources responded.

9 source records (9 saved revisions) → 9 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | aws | [Increased Connectivity Issues and API Error Rates](https://health.aws.amazon.com/health/status) | 2026-03-02 05:56 | 257d 6h 40m (ongoing) | unknown | multipleservices-me-south-1 | 740960 = 370480 × 1 × 2.0 |
| 2 | slack | [Issues with App Mentions, Grid Migrations, and Notifications](https://slack-status.com/2026-03/e63ca9cc0d76539c) | 2026-03-05 14:56 | 1d 1h 48m | unknown | Apps/Integrations/APIs, Notifications, Workspace/Org Administration | 5573 = 1548 × 3 × 1.2 |
| 3 | slack | [Degraded Search Performance for Dubai IDR Users – Middle East (Dubai) Region](https://slack-status.com/2026-03/5d92c691d977da03) | 2026-03-02 08:50 | 2d 9h 51m | unknown | Search | 4165 = 3471 × 1 × 1.2 |
| 4 | stripe | [Federal Reserve ACH system disruption](https://stspg.io/tc7z8ct6pyqd) | 2026-03-03 15:06 | 5h 1m | minor | Stripe core components | 361 = 301 × 1 × 1.2 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. aws — Increased Connectivity Issues and API Error Rates

- URL: https://health.aws.amazon.com/health/status
- Started: 2026-03-02 05:56 UTC
- Resolved: not yet
- Duration: 257d 6h 40m (ongoing)
- Impact: unknown, status: investigating
- Components: multipleservices-me-south-1
- Canonical incident: [ci-ca1adfedc3b0d11bb596da39](#ci-ca1adfedc3b0d11bb596da39)

- `2026-03-02 05:56` **investigating** — We are investigating increased API error rates in a single Availability Zone (mes1-az2) in the ME-SOUTH-1 Region.
- `2026-03-02 07:09` **investigating** — We are investigating connectivity and power issues affecting APIs and instances in a single Availability Zone (mes1-az2) in the ME-SOUTH-1 Region due to a localized power issue. Existing instances in this zone will also be affected. Other AWS Services may also be experiencing increased errors and latencies for their workflows, and we are working to route requests away from this affected Availability Zone. We recommend customers make use of other Availability Zones at this time. During this time, we are also experiencing delays in propagating DNS changes for Route53 to pops (Points of Presence) in ME-SOUTH-1. Targeting new launches using RunInstances in the remaining AZs should succeed. Existing instances in the other AZs are not affected.
- `2026-03-02 09:03` **investigating** — We continue to work on a localized power issue affecting a single Availability Zone (mes1-az2) in the ME-SOUTH-1 Region. In the impacted Availability Zone, EC2 Instances, DB Instances, EBS Volumes, and other AWS Services are also experiencing elevated error rates and latencies for some workflows. As part of our recovery effort, we have shifted traffic away from the impacted Availability Zone for most services. We recommend customers utilize one of the other Availability Zones in the ME-SOUTH-1 Region, as existing instances in other AZs remain unaffected by this issue. We are actively working to restore power and connectivity, at which time we will begin recovering affected resources. Currently, we expect recovery to take many hours. We will provide an update by 2:30 AM PST, or sooner if we have additional information to share.
- `2026-03-02 10:41` **investigating** — We continue to work toward restoring power in the affected Availability Zone (mes1-az2) in the ME-SOUTH-1 Region. At this time, some AWS services have shifted traffic away from the affected Availability Zone and are seeing recovery for their affected operations and workflows. EC2 Instances, EBS Volumes, and other resources impacted in the affected Availability Zone will require a longer recovery timeline. Power has not yet been restored to the affected Availability Zone. If immediate recovery is required, we recommend customers restore from EBS Snapshots and/or launch replacement resources in one of the unaffected Availability Zones or an alternate Region. In parallel, we are actively working on reducing the error rates and latencies that some customers are experiencing with EC2 APIs. For now, we recommend continuing to retry any failed API requests. We will provide an update by 6:00 AM PST on March 2, or sooner if we have additional information to share.
- `2026-03-02 14:23` **investigating** — We continue to work toward restoring power in the impacted Availability Zone (mes1-az2) in the ME-SOUTH-1 Region. Meanwhile, EC2 instance and networking APIs have been restored for the other Availability Zones. Additionally, we have made improvements to the availability of RDS multi-AZ databases while operating with the impaired Availability Zone. These improvements will help customers create database exports to preserve data, and we recommend customers with databases in the affected Availability Zone consider creating exports as a precautionary measure. EC2 Instances, EBS Volumes, and other resources impacted in the affected Availability Zone will require a longer recovery timeline, as power has not yet been restored. We are expecting recovery to take at least a day, as it requires repair of facilities, cooling and power systems, coordination with local authorities, and careful assessment to ensure the safety of our operators. If immediate recovery is required, we recommend customers restore from EBS Snapshots and/or launch replacement resources in one of the unaffected Availability Zones or an alternate AWS Region. We will provide an update by 11:00 AM PST on March 2, or sooner if we have additional information to share.
- `2026-03-02 18:52` **investigating** — We continue to work towards restoring power in the affected Availability Zone (mes1-az2) in the ME-SOUTH-1 Region. We currently expect our recovery efforts to take at least a day. Our current guidance regarding immediate recovery remains unchanged from our previous update. Customers are able to disassociate Elastic IP addresses from resources in the affected Availability Zone and associate those with resources in the unaffected Availability Zones. This can be done by specifying --allow-reassociation when attempting to associate the Elastic IP to the new resource. We will provide you with further updates by 2:00 PM PST or sooner if new information becomes available.
- `2026-03-02 22:29` **investigating** — We continue to work towards restoring power in the affected Availability Zone (mes1-az2) in the ME-SOUTH-1 Region. We have no updated guidance on expected recovery times, and still expect this to take at least a day to fully restore power and connectivity. We continue to advise customers to launch replacement resources in one of the unaffected Availability Zones or an alternate AWS Region. At this time we recommend that customers that are capable of backing up data outside of the region consider doing so. You can view the current status of affected AWS services below. We will provide you with another update by 7:00 PM PST, or sooner if we have additional information to share.
- `2026-03-03 00:22` **investigating** — We are providing an update on the ongoing service disruptions affecting the AWS Middle East (UAE) Region (ME-CENTRAL-1) and the AWS Middle East (Bahrain) Region (ME-SOUTH-1). Due to the ongoing conflict in the Middle East, both affected regions have experienced physical impacts to infrastructure as a result of drone strikes. In the UAE, two of our facilities were directly struck, while in Bahrain, a drone strike in close proximity to one of our facilities caused physical impacts to our infrastructure. These strikes have caused structural damage, disrupted power delivery to our infrastructure, and in some cases required fire suppression activities that resulted in additional water damage. We are working closely with local authorities and prioritizing the safety of our personnel throughout our recovery efforts. In the ME-CENTRAL-1 (UAE) Region, two of our three Availability Zones (mec1-az2 and mec1-az3) remain significantly impaired. The third Availability Zone (mec1-az1) continues to operate normally, though some services have experienced indirect impact due to dependencies on the affected zones. In the ME-SOUTH-1 (Bahrain) Region, one facility has been impacted. Across both regions, customers are experiencing elevated error rates and degraded availability for services including Amazon EC2, Amazon S3, Amazon DynamoDB, AWS Lambda, Amazon Kinesis, Amazon CloudWatch, Amazon RDS, and the AWS Management Console and CLI. We are working to restore full service availability as quickly as possible, though we expect recovery to be prolonged given the nature of the physical damage involved. In parallel with efforts to restore the physical infrastructure at the affected sites, we are pursuing multiple software-based recovery paths that do not depend on the underlying facilities being fully brought back online. For Amazon S3 and Amazon DynamoDB, we are actively working to restore data access and service availability through software mitigations, including deploying updates to enable S3 to operate within the current infrastructure constraints and remediating impaired DynamoDB tables to restore read and write availability for dependent services. Our focus on restoring these foundational services is deliberate, as recovery of Amazon S3 and Amazon DynamoDB will in turn enable a broad range of dependent AWS services to recover. For other affected service APIs, we are deploying targeted software updates to reduce error rates and restore functionality where possible, independent of the physical recovery timeline. We are also working to restore access to the AWS Management Console and CLI through network-level changes that route traffic away from the affected infrastructure. While these software-based mitigations can address many of the service-level impacts, some recovery actions are constrained by the physical state of the affected facilities — meaning that full restoration of certain services will require the underlying infrastructure to be repaired and brought back online. Across all services, our teams are working in parallel on both the physical restoration of the affected facilities and these software-based mitigations, with the goal of restoring as much customer access as possible as quickly as possible, even ahead of full infrastructure recovery. In addition, we are prioritizing the restoration of services and tools that enable customers to back up and migrate their data and applications out of the affected regions. Finally, even as we work to restore these facilities, the ongoing conflict in the region means that the broader operating environment in the Middle East remains unpredictable. We recommend that customers with workloads running in the Middle East consider taking action now to backup data and potentially migrate your workloads to alternate AWS Regions. We recommend customers exercise their disaster recovery plans, recover from remote backups stored in other regions, and update their applications to direct traffic away from the affected regions. For customers requiring guidance on alternate regions, we recommend considering AWS Regions in the United States, Europe, or Asia Pacific, as appropriate for your latency and data residency requirements. We will continue to provide updates as recovery progresses and as the situation evolves. Our next update will be provided by 9:00 PM PST on March 2, 2026, or sooner if new information becomes available.
- `2026-03-03 06:27` **investigating** — We continue to work towards restoring power in the affected Availability Zone (mes1-az2) in the ME-SOUTH-1 Region. We have no updated guidance on expected recovery times, and still expect this to take at least a day to fully restore power and connectivity. AWS infrastructure is designed to be highly resilient, but given the uncertainty of the current situation, we encourage our customers to replicate Amazon S3 and critical data from the ME-SOUTH-1 Region to another AWS Region. For customers requiring guidance on alternate regions, we recommend considering AWS Regions in the United States, Europe, or Asia Pacific, as appropriate for your latency and data residency requirements. We will provide another update by March 3 at 3:00 AM PST, or sooner if new information becomes available. For more information on Cross-Region Replication, refer [1]. For more information on S3 Batch Replication, see [2]. For a simple script to quickly set up and start S3 Replication, see [3]. If you have questions or concerns, please contact AWS Support [4]. [1] <a href="https://docs.aws.amazon.com/AmazonS3/latest/userguide/replication.html">https://docs.aws.amazon.com/AmazonS3/latest/userguide/replication.html</a> [2] <a href="https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-batch-replication-batch.html">https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-batch-replication-batch.html</a> [3] <a href="https://github.com/awslabs/aws-support-tools/blob/master/S3/Setup_Replication/setup_replication.py">https://github.com/awslabs/aws-support-tools/blob/master/S3/Setup_Replication/setup_replication.py</a> [4] <a href="https://aws.amazon.com/support">https://aws.amazon.com/support</a>
- `2026-03-03 11:10` **investigating** — We continue to work toward restoring power in the affected Availability Zone (mes1-az2) in the ME-SOUTH-1 Region. The overall state of the region remains largely unchanged from our previous update. At this time, we have no updated guidance on expected timelines for fully restoring power and connectivity. We are taking all necessary steps to support the recovery process. While progress is being made, significant work remains before full restoration is complete. Given the ongoing uncertainty, we encourage customers to replicate their Amazon S3 data and other critical data from the ME-SOUTH-1 Region to another AWS Region, using the guidance provided in our previous update. We will continue to provide updates as recovery progresses and as the situation evolves. Our next update will be provided by 6:00 AM PST on March 3, or sooner if new information becomes available.
- `2026-03-03 14:02` **investigating** — Recovery efforts in the affected Availability Zone (mes1-az2) in the ME-SOUTH-1 Region are ongoing, with the situation remaining consistent with our last update. We have no change to expected timelines for fully restoring power and connectivity. While progress is being made, significant work remains before full restoration is complete. We continue to recommend customers launch replacement resources in one of the unaffected Availability Zones or an alternate AWS Region. Given the extended nature of this event, we continue to encourage customers to replicate Amazon S3 data and other critical workloads from ME-SOUTH-1 to another AWS Region using the guidance shared previously. We will provide our next update by 12:00 PM PST on March 3, or sooner if conditions change.
- `2026-03-03 16:40` **investigating** — We are providing an update on the ongoing service disruptions affecting the AWS Middle East (Bahrain) Region (ME-SOUTH-1). We continue to make progress on recovery efforts across multiple workstreams. With the immediate phase of this event now better understood, we are moving to a more targeted communication model. Going forward, updates will be delivered directly to affected customers through the AWS Personal Health Dashboard. Customers who require assistance with this event are encouraged to contact AWS Support through the AWS Management Console or the AWS Support Center. We continue to strongly recommend that customers with workloads running in the Middle East take action now to migrate those workloads to alternate AWS Regions. Customers should enact their disaster recovery plans, recover from remote backups stored in other Regions, and update their applications to direct traffic away from the affected Regions. For customers requiring guidance on alternate regions, we recommend considering AWS Regions in the United States, Europe, or Asia Pacific, as appropriate for your latency and data residency requirements.

### 2. slack — Issues with App Mentions, Grid Migrations, and Notifications

- URL: https://slack-status.com/2026-03/e63ca9cc0d76539c
- Started: 2026-03-05 14:56 UTC
- Resolved: 2026-03-06 16:44 UTC
- Duration: 1d 1h 48m
- Impact: unknown, status: resolved
- Components: Apps/Integrations/APIs, Notifications, Workspace/Org Administration
- Canonical incident: [ci-28b809aa3e401e96c0708eba](#ci-28b809aa3e401e96c0708eba)

- `2026-03-05 14:56` **unknown** — We're aware of an issue impacting the completion of ongoing enterprise migrations, resulting in longer than expected downtime. In efforts to resolve the issue we're increasing processing capacity and have deployed code changes to address the delayed job processing. We'll provide an update when we next have additional information.
- `2026-03-05 19:36` **unknown** — We've identified additional impact where developers may observe no responses when @mentioning an application. This incident extends impact to Grid migrations and iOS push notifications may not deliver. Current status and actions being taken: We continue to investigate potential triggers in recent deployments for this issue. Scope: This will impact any developers @mentioning an app in a Slack channel, upcoming/planned migrations, and iOS notifications. Current impact to end users: Developers may not receive responses when @mentioning an app to execute an action in a Slack channel, org admins will observe delays with upcoming/planned migrations, and iOS push notifications may not deliver. Known workarounds: No known workarounds. We appreciate your patience as we sort this out. We'll continue to provide updates every 30 minutes until the impact is resolved for all users.
- `2026-03-05 20:07` **unknown** — We’re still actively investigating this issue, but we don’t have any new information to share at this time. Thanks for sticking with us as we continue to work towards a fix. We'll provide another update in 30 minutes until impact is resolved for all users.
- `2026-03-05 20:40` **unknown** — We're actively working to resolve the current platform functionality issues. Current status and actions being taken: We’re continuing to investigate the underlying trigger and have noticed an improvement in health metrics with Push notifications. Estimated time to resolution: No ETA at this time. Scope: This issue isn't affecting any specific region, but a percentage of users may be affected. Current impact to end users: Developers may not receive responses when @mentioning an app to execute an action in a Slack channel, org admins will observe delays with upcoming/planned migrations, and mobile push notifications may not deliver. Known workarounds: No known workarounds at this time. We'll continue to provide updates every 30 minutes until the impact is resolved for all users.
- `2026-03-05 21:14` **unknown** — We're actively working to resolve the current platform functionality issues. Current status and actions being taken: We're currently working on a fix to block dispatch jobs while increasing resources to expedite processing. Estimated time to resolution: No ETA at this time. Scope: This issue isn't affecting any specific region, but a percentage of users may be affected. Current impact to end users: Developers may not receive responses when @mentioning an app to execute an action in a Slack channel, org admins will observe delays with upcoming/planned migrations, and mobile push notifications may not deliver. Known workarounds: No known workarounds at this time. We'll continue to provide updates every 30 minutes until the impact is resolved for all users.
- `2026-03-05 21:46` **unknown** — We're actively working to resolve the current platform functionality issues. Current status and actions being taken: Our investigation continues, and we're progressing toward a fix. Estimated time to resolution: No ETA at this time. Scope: This issue isn't affecting any specific region, but a percentage of users may be affected. Current impact to end users: Developers may not receive responses when @mentioning an app to execute an action in a Slack channel, org admins will observe delays with upcoming/planned migrations, and mobile push notifications may not deliver. Known workarounds: No known workarounds at this time. We'll continue to provide updates every 30 minutes until the impact is resolved for all users.
- `2026-03-05 22:20` **unknown** — We're actively working to resolve the current platform functionality issues. Current status and actions being taken: We've increased resources to expedite processing and are continuing to monitor the situation. We're currently observing improved health metrics. Estimated time to resolution: No ETA at this time. Scope: This issue isn't affecting any specific region, but a percentage of users may be affected. Current impact to end users: Developers may not receive responses when @mentioning an app to execute an action in a Slack channel, org admins will observe delays with upcoming/planned migrations, and mobile push notifications may not deliver. Known workarounds: No known workarounds at this time. We'll continue to provide updates every 60 minutes until the impact is resolved for all users.
- `2026-03-05 23:27` **unknown** — We're actively working to resolve the current platform functionality issues. Current status and actions being taken: Our platform is currently processing the message backlog, and we're observing improvements in our health metrics. Estimated time to resolution: No ETA at this time. Scope: This issue isn't affecting any specific region, but a percentage of users may be affected. Current impact to end users: Developers may not receive responses when @mentioning an app to execute an action in a Slack channel, org admins will observe delays with upcoming/planned migrations, and mobile push notifications may not deliver. Known workarounds: No known workarounds at this time. We'll continue to provide updates every 60 minutes until the impact is resolved for all users.
- `2026-03-06 00:35` **unknown** — We are seeing improvements and working toward full resolution. Current status and actions being taken: We have cleared the message backlog and restored normal processing. Most activity has now completed, and the remaining work is progressing as expected. We will continue monitoring as the remaining scheduled activity runs. Estimated time to resolution: No ETA at this time. Scope: This issue isn't affecting any specific region, but a percentage of users may be affected. Current impact to end users: Developers may not receive responses when @mentioning an app to execute an action in a Slack channel, org admins will observe delays with upcoming/planned migrations, and mobile push notifications may not deliver. Known workarounds: No known workarounds at this time. We will provide the next update as we make further progress or if significant developments occur.
- `2026-03-06 16:44` **unknown** — This issue is now resolved for all users. Current status and recent actions taken: The issue is now resolved, and users should no longer be experiencing any trouble with when using app mentions, migrations, and push notifications. Enterprise Migrations should also be completing as expected. Scope: This issue wasn't affecting any specific region, but a percentage of users may have been affected globally. Previous impact to end users: Developers may have experienced issues in receiving responses when @mentioning an app to execute an action in a Slack channel, org admins may have observed delays with upcoming/planned migrations, and mobile push notifications may have failed to deliver. Required resolution steps: No steps required. We apologize for any disruptions to your day.

### 3. slack — Degraded Search Performance for Dubai IDR Users – Middle East (Dubai) Region

- URL: https://slack-status.com/2026-03/5d92c691d977da03
- Started: 2026-03-02 08:50 UTC
- Resolved: 2026-03-04 18:41 UTC
- Duration: 2d 9h 51m
- Impact: unknown, status: resolved
- Components: Search
- Canonical incident: [ci-8c6e7b740e7c2c59d04e884f](#ci-8c6e7b740e7c2c59d04e884f)

- `2026-03-02 08:50` **unknown** — We’ve noted that International Data Residency (IDR) users in the Middle East (Dubai) region may experience degraded search performance or search failures. Services for other users in the region are operating normally. We're aware of a vendor issue affecting multiple availability zones in the region. We are actively engaged in restoring the services as quickly as possible.
- `2026-03-02 14:47` **unknown** — Our work on this issue is still ongoing. Current status and actions being taken: International Data Residency (IDR) users continue to experience search issues due to an ongoing regional incident. We're working with a third-party vendor to restore service... We don’t have an ETA but it’s likely to take at least a day for service to be restored. We continue to monitor the situation closely and will provide updates as more information becomes available.
- `2026-03-02 22:47` **unknown** — Our work on this issue is still ongoing. Current status and actions being taken: International Data Residency (IDR) customers continue to experience search disruptions due to a regional incident affecting our third-party provider, Amazon Web Services. This issue is related to the broader Salesforce incident currently underway. We are actively working to restore service, but there is no estimated time of recovery. Conditions at the impacted facilities, including loss of power and cooling, continue to affect progress. We are closely monitoring the situation, and will provide further updates as new information becomes available. Current impact to end users: International Data Residency (IDR) users in the Middle East (Dubai) region may experience degraded search performance or search failures. Services for other users in the region are operating normally. We continue to monitor the situation closely and will provide updates as more information becomes available.
- `2026-03-03 20:30` **unknown** — Our work on this issue is still ongoing. Current status and actions being taken: We are currently managing a severe physical infrastructure incident in the Middle East following significant damage to regional data centers. While one facility remains operational, we are continuing to work closely with AWS to resolve current data access issues and restore full service, with a major recovery milestone expected within the next 24 hours. This recovery will resolve current data access issues where some older files and database records remain temporarily unreachable. Our teams are prioritizing the restoration of missing data shards and database records as quickly as safety protocols allow. Current impact to end users: International Data Residency (IDR) users in the Middle East (Dubai) region may experience degraded search performance or search failures. Services for other users in the region are operating normally. We continue to monitor the situation closely and will provide updates as more information becomes available.
- `2026-03-04 18:41` **unknown** — This issue with a third party vendor impacts only a subset of IDR customers. Current status and recent actions taken: We are staying engaged with our International Data Residency (IDR) customers on a path to recovery. Scope: Isolated to a subset of IDR customers. Previous impact to end users: Severe degradation to services only impacting IDR customers. Required resolution steps: Remain engaged with our IDR customers until they are fully recovered.

### 4. stripe — Federal Reserve ACH system disruption

- URL: https://stspg.io/tc7z8ct6pyqd
- Started: 2026-03-03 15:06 UTC
- Resolved: 2026-03-03 20:07 UTC
- Duration: 5h 1m
- Impact: minor, status: resolved
- Components: Stripe core components
- Canonical incident: [ci-694efca0d56424e23e262173](#ci-694efca0d56424e23e262173)

- `2026-03-03 15:06` **investigating** — We are aware of a Federal Reserve ACH system disruption that may affect ACH transaction processing, including payouts, top ups bank debits and bank transfers. We are actively working with our banking partners to assess any impact to Stripe users. We will continue to monitor the situation and provide updates as more information becomes available.
- `2026-03-03 18:22` **investigating** — We’re continuing to monitor the US Federal Reserve ACH system disruption. This may cause delays to ACH transaction processing, including US payouts, top ups, bank debits, and bank transfers. We anticipate that all the ACH transactions will still settle later today and will provide updates as more information becomes available.
- `2026-03-03 20:07` **resolved** — The US Federal Reserve's ACH system (FedACH) experienced an outage which delayed ACH processing, including US payouts, top ups, bank debits, and bank transfers. The issue is now resolved. ACH processing expected to be received March 3rd should settle later today.

## Below threshold (5)

- `stripe` [Dashboard pages export errors](https://stspg.io/f411dzxr6x7b) — 1h 19m
- `stripe` [Elevated errors for multiple Payment Methods](https://stspg.io/29xyj5xknp7p) — 1h 18m
- `stripe` [Elevated API errors](https://stspg.io/5r6k5ht4vvg6) — 21m
- `circleci` [docker not working in cimg:*](https://stspg.io/wqmcz2chv2f6) — 21m
- `stripe` [Elevated API errors](https://stspg.io/sjvl135mk2yv) — 11m

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-ca1adfedc3b0d11bb596da39"></a>

- Canonical ID: `ci-ca1adfedc3b0d11bb596da39` (1 source records).
  - [aws/93d92e456f200868](../../data/2026/W10/aws.json); vendor URL: https://health.aws.amazon.com/health/status
- Start time basis: `first_observed_update`.
- Warning: AWS onset is the first observed message, not a verified impact start

<a id="ci-28b809aa3e401e96c0708eba"></a>

- Canonical ID: `ci-28b809aa3e401e96c0708eba` (1 source records).
  - [slack/1535](../../data/2026/W10/slack.json); vendor URL: https://slack-status.com/2026-03/e63ca9cc0d76539c
- Start time basis: `vendor_reported`.

<a id="ci-8c6e7b740e7c2c59d04e884f"></a>

- Canonical ID: `ci-8c6e7b740e7c2c59d04e884f` (1 source records).
  - [slack/1534](../../data/2026/W10/slack.json); vendor URL: https://slack-status.com/2026-03/5d92c691d977da03
- Start time basis: `vendor_reported`.

<a id="ci-694efca0d56424e23e262173"></a>

- Canonical ID: `ci-694efca0d56424e23e262173` (1 source records).
  - [stripe/vbrq9l2x48y6](../../data/2026/W10/stripe.json); vendor URL: https://stspg.io/tc7z8ct6pyqd
- Start time basis: `vendor_reported`.

<a id="ci-adc7ba4bc92103665e4cfa40"></a>

- Canonical ID: `ci-adc7ba4bc92103665e4cfa40` (1 source records).
  - [stripe/rv8m0mtn024h](../../data/2026/W10/stripe.json); vendor URL: https://stspg.io/f411dzxr6x7b
- Start time basis: `vendor_reported`.

<a id="ci-38b483378a53b7b39e3139dd"></a>

- Canonical ID: `ci-38b483378a53b7b39e3139dd` (1 source records).
  - [stripe/cqvlkm3fcr7x](../../data/2026/W10/stripe.json); vendor URL: https://stspg.io/29xyj5xknp7p
- Start time basis: `vendor_reported`.

<a id="ci-07122084487a66276aa043ad"></a>

- Canonical ID: `ci-07122084487a66276aa043ad` (1 source records).
  - [stripe/5jywybwt5qcz](../../data/2026/W10/stripe.json); vendor URL: https://stspg.io/5r6k5ht4vvg6
- Start time basis: `vendor_reported`.

<a id="ci-8da5b1de9672a87ce89fcf29"></a>

- Canonical ID: `ci-8da5b1de9672a87ce89fcf29` (1 source records).
  - [circleci/lg8x09p2d2n3](../../data/2026/W10/circleci.json); vendor URL: https://stspg.io/wqmcz2chv2f6
- Start time basis: `vendor_reported`.

<a id="ci-50aa9eb6d358e17cafc4524a"></a>

- Canonical ID: `ci-50aa9eb6d358e17cafc4524a` (1 source records).
  - [stripe/xw65m0409rp0](../../data/2026/W10/stripe.json); vendor URL: https://stspg.io/sjvl135mk2yv
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
| gcp | ok | 6 | 0 |
| github | ok | 50 | 0 |
| grafana | ok | 50 | 0 |
| openai | ok | 25 | 0 |
| slack | ok | 36 | 0 |
| stripe | ok | 50 | 0 |
