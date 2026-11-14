# Week 2026-W23

2026-06-01 to 2026-06-07 (UTC). Generated 2026-11-14 00:00.

14 canonical incidents, 8 above threshold. All 14 sources responded.

14 source records (14 saved revisions) → 14 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | gcp | [Network traffic to Google Cloud originating from Delhi, Chennai, Mumbai and surrounding areas is experiencing intermittent periods of elevated latency and possible packet loss.](https://status.cloud.google.com/incidents/5fGQt4VbkDnr3Yp8PXPr) | 2026-06-05 07:00 | 21d 12h 0m | major | Hybrid Connectivity, Media CDN, Virtual Private Cloud (VPC) | 185760 = 30960 × 3 × 2.0 |
| 2 | yandexcloud | [Incorrect charge of funds.](https://status.yandex.cloud/en/incidents/1824) | 2026-06-05 12:25 | 8h 30m | major | Yandex Cloud Billing | 765 = 510 × 1 × 1.5 |
| 3 | stripe | [Elevated declines on Cartes Bancaires](https://stspg.io/ydwywzw7gnp3) | 2026-06-05 05:33 | 9h 8m | none | Acquirers and payment methods | 658 = 548 × 1 × 1.2 |
| 4 | yandexcloud | [The problem is with the AI Studio service](https://status.yandex.cloud/en/incidents/1807) | 2026-06-01 02:37 | 4h 14m | minor | Yandex AI Studio | 381 = 254 × 1 × 1.5 |
| 5 | yandexcloud | [Partial unavailability of services](https://status.yandex.cloud/en/incidents/1822) | 2026-06-04 08:30 | 3h 0m | major | Yandex BareMetal | 270 = 180 × 1 × 1.5 |
| 6 | yandexcloud | [Problems getting a list of models in the AI Studio service](https://status.yandex.cloud/en/incidents/1820) | 2026-06-03 13:08 | 2h 13m | minor | Yandex AI Studio | 200 = 133 × 1 × 1.5 |
| 7 | circleci | [Mac jobs failing to start](https://stspg.io/n3brb41nn16r) | 2026-06-01 01:04 | 2h 21m | minor | macOS Jobs | 141 = 141 × 1 × 1.0 |
| 8 | circleci | [macOS jobs failing to start](https://stspg.io/fcqx1c8m4dqx) | 2026-06-07 04:20 | 24m | major | macOS Jobs | 24 = 24 × 1 × 1.0 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. gcp — Network traffic to Google Cloud originating from Delhi, Chennai, Mumbai and surrounding areas is experiencing intermittent periods of elevated latency and possible packet loss.

- URL: https://status.cloud.google.com/incidents/5fGQt4VbkDnr3Yp8PXPr
- Started: 2026-06-05 07:00 UTC
- Resolved: 2026-06-26 19:00 UTC
- Duration: 21d 12h 0m
- Impact: major, status: resolved
- Components: Hybrid Connectivity, Media CDN, Virtual Private Cloud (VPC)
- Canonical incident: [ci-cf47cdc71eaf54c1e652582b](#ci-cf47cdc71eaf54c1e652582b)

- `2026-06-10 01:13` **investigating** — **Summary** Network traffic to Google Cloud originating from Delhi, Chennai, Mumbai and surrounding areas is experiencing intermittent periods of elevated latency and possible packet loss. **Description** A fire at a third-party data center facility required an emergency power shutdown of networking equipment, isolating a non-compute local Point of Presence (POP) in Delhi and reducing available network capacity in the metro area. We rerouted significant traffic from the impacted facility in Delhi to address reduced local serving capabilities. As a result, a subset of Hybrid Connectivity and Virtual Private Cloud (VPC) customers may be impacted by the routing changes made to address reduced local, latency-optimized serving capabilities in Delhi. Affected customers may experience intermittent latency spikes due to demand exceeding capacity across Indian metros and regional ISPs. We are investigating additional traffic mitigations and Internet Edge peering augmentation to alleviate the latency issues affecting our customers. We will provide our next update by Wednesday, 2026-06-10 17:00 PDT. **Symptoms** Customers may experience slightly elevated latency and non-optimal network routing into Google Cloud until the affected facility is fully restored. **Workaround** There is no workaround at this time.
- `2026-06-10 22:25` **investigating** — **Summary** Network traffic to Google Cloud originating from Delhi, Chennai, Mumbai and surrounding areas is experiencing intermittent periods of elevated latency and possible packet loss. **Description** A fire at a third-party data center facility required an emergency power shutdown of networking equipment, isolating a non-compute local Point of Presence (POP) in Delhi and reducing available network capacity in the metro area. We rerouted significant traffic from the impacted facility in Delhi to address reduced local serving capabilities. As a result, a subset of Hybrid Connectivity and Virtual Private Cloud (VPC) customers may be impacted by the routing changes made to address reduced local, latency-optimized serving capabilities in Delhi. Affected customers may experience intermittent latency spikes due to demand exceeding capacity across Indian metros and regional ISPs. We are investigating additional traffic mitigations and Internet Edge peering augmentation to alleviate the latency issues affecting our customers. We are continuing to work with the ISP partners in the region to mitigate any additional impact from unplanned failures. We will provide our next update by Thursday, 2026-06-11 17:00 PDT. **Symptoms** Customers may experience slightly elevated latency and non-optimal network routing into Google Cloud until the affected facility is fully restored. **Workaround** There is no workaround at this time.
- `2026-06-11 23:25` **investigating** — **Summary** Network traffic to Google Cloud originating from Delhi, Chennai, Mumbai and surrounding areas is experiencing intermittent periods of elevated latency and possible packet loss. **Description** A fire at a third-party data center facility required an emergency power shutdown of networking equipment, isolating a non-compute local Point of Presence (POP) in Delhi and reducing available network capacity in the metro area. We rerouted significant traffic from the impacted facility in Delhi to address reduced local serving capabilities. As a result, a subset of Hybrid Connectivity and Virtual Private Cloud (VPC) customers may be impacted by the routing changes made to address reduced local, latency-optimized serving capabilities in Delhi. Affected customers may experience intermittent latency spikes due to demand exceeding capacity across Indian metros and regional ISPs. Initial traffic mitigations have yielded positive results for some Cloud customers. In parallel, we are pursuing additional Internet Edge peering capacity to reduce existing fragility in the metro. We are also optimizing capacity across network backbones to increase additional headroom for VPC customers in the affected region. Additionally, we are planning to augment our local Delhi POP and migrating select peering partners to further increase regional capacity. We will continue to closely monitor latency deviations and packet drops. We will provide our next update by Friday, 2026-06-12 at 17:00 PDT. **Symptoms** Customers may experience slightly elevated latency and non-optimal network routing into Google Cloud until the affected facility is fully restored. **Workaround** There is no workaround at this time.
- `2026-06-12 22:50` **investigating** — **Summary** Network traffic to Google Cloud originating from Delhi, Chennai, Mumbai and surrounding areas is experiencing intermittent periods of elevated latency and possible packet loss. **Description** A fire at a third-party data center facility required an emergency power shutdown of networking equipment, isolating a non-compute local Point of Presence (POP) in Delhi and reducing available network capacity in the metro area. We rerouted significant traffic from the impacted facility in Delhi to address reduced local serving capabilities. As a result, a subset of Hybrid Connectivity and Virtual Private Cloud (VPC) customers may be impacted by the routing changes made to address reduced local, latency-optimized serving capabilities in Delhi. Affected customers may experience intermittent latency spikes due to demand exceeding capacity across Indian metros and regional ISPs. Initial traffic mitigations have yielded positive results for some Cloud customers. In parallel, we are pursuing additional Internet Edge peering capacity to reduce latency in the local Delhi metropolitan area. Further, we are augmenting out-of-region Internet Edge regional peering capacity in Chennai to provide additional load-balancing and redundancy to large ISPs in India (expected to be done by Wednesday, 2026-06-17 PDT). We have optimized capacity across network backbones to increase available headroom. Additionally, we are further augmenting our Delhi backbone capacity (expected to be complete by Monday, 2026-06-15 PDT). We will continue to closely monitor latency deviations and packet drops. We will provide our next update by Monday, 2026-06-15 17:00 PDT. **Symptoms** Customers may experience slightly elevated latency and non-optimal network routing into Google Cloud until the affected facility is fully restored. **Workaround** There is no workaround at this time.
- `2026-06-15 23:11` **investigating** — **Summary** Network traffic to Google Cloud originating from Delhi, Chennai, Mumbai and surrounding areas is experiencing intermittent periods of elevated latency and possible packet loss. **Description** A fire at a third-party data center facility required an emergency power shutdown of networking equipment, isolating a non-compute local Point of Presence (POP) in Delhi and reducing available network capacity in the metro area. We rerouted significant traffic from the impacted facility in Delhi to address reduced local serving capabilities. As a result, a subset of Hybrid Connectivity and Virtual Private Cloud (VPC) customers may be impacted by the routing changes made to address reduced local, latency-optimized serving capabilities in Delhi. Affected customers may experience intermittent latency spikes due to demand exceeding capacity across Indian metros and regional ISPs. Media CDN customers may have experienced higher latencies than normal. Initial traffic mitigations have yielded positive results for some Cloud customers. We have restored a portion of Internet Edge peering capacity to reduce latency in the local Delhi metropolitan area. Further, we are augmenting out-of-region Internet Edge regional peering capacity in Chennai to provide additional load-balancing and redundancy to large ISPs in India (expected to be done by Wednesday, 2026-06-17 PDT). We have optimized capacity across network backbones to increase available headroom. Additionally, we have augmented our Delhi backbone capacity over the weekend. We will continue to closely monitor latency deviations and packet drops. We will provide our next update by Wednesday, 2026-06-17 at 17:00 PDT. **Symptoms** Customers may experience slightly elevated latency and non-optimal network routing into Google Cloud until the affected facility is fully restored. **Workaround** There is no workaround at this time.
- `2026-06-17 22:36` **investigating** — **Summary** Network traffic to Google Cloud originating from Delhi, Chennai, Mumbai and surrounding areas is experiencing intermittent periods of elevated latency and possible packet loss. **Description** Traffic rerouting from the impacted Delhi facility has caused a subset of Hybrid Connectivity and Virtual Private Cloud (VPC) customers to experience intermittent latency spikes as demand exceeds regional capacity. Media CDN customers may also notice increased latencies. Initial traffic mitigations have yielded positive results for some Cloud customers. We have restored a portion of Internet Edge peering capacity to reduce latency in the local Delhi metropolitan area. Further, we are nearly finished with the augmentation of out-of-region Internet Edge regional peering capacity in Chennai to provide additional load-balancing and redundancy to large ISPs in India. We have optimized capacity across network backbones to increase available headroom. Additionally, we have augmented our Delhi backbone capacity. We have also restored additional capacity between Delhi-Chennai and Delhi-Mumbai. We will continue to closely monitor latency deviations and packet drops. We will provide our next update by Monday, 2026-06-22 17:00 PDT. **Symptoms** Customers may experience slightly elevated latency and non-optimal network routing into Google Cloud until the affected facility is fully restored. **Workaround** There is no workaround at this time.
- `2026-06-22 23:51` **investigating** — **Summary** Network traffic to Google Cloud originating from Delhi, Chennai, Mumbai and surrounding areas is experiencing intermittent periods of elevated latency and possible packet loss. **Description** Traffic rerouting from the impacted Delhi facility has caused a subset of Hybrid Connectivity and Virtual Private Cloud (VPC) customers to experience intermittent latency spikes as demand exceeds regional capacity. Media CDN customers may also notice increased latencies. Initial traffic mitigations have yielded positive results for some Cloud customers. We have restored a large portion of Internet Edge peering capacity to reduce latency in the local Delhi metropolitan area. Since the last update, we have completed the augmentation of out-of-region Internet Edge regional peering capacity in Chennai to provide additional load-balancing and redundancy to large ISPs in India. Following safety clearance, teams obtained access to the damaged site and will further restore additional incremental user-facing backbone capacity on Tuesday, 2026-06-23. We have optimized capacity across network backbones to increase available headroom, and augmented our Delhi user-facing backbone capacity. We will continue to closely monitor latency deviations and packet drops. We will provide our next update by Tuesday, 2026-06-23 at 17:00 PDT. **Symptoms** Customers may experience slightly elevated latency and non-optimal network routing into Google Cloud until the affected facility is fully restored. **Workaround** There is no workaround at this time.
- `2026-06-23 22:52` **investigating** — **Summary** Network traffic to Google Cloud originating from Delhi, Chennai, Mumbai and surrounding areas is experiencing intermittent periods of elevated latency and possible packet loss. **Description** Traffic rerouting from the impacted Delhi facility has caused a subset of Hybrid Connectivity and Virtual Private Cloud (VPC) customers to experience intermittent latency spikes as demand exceeds regional capacity. Media CDN customers may also notice increased latencies. Initial traffic mitigations have yielded positive results for some Cloud customers. We have restored a large portion of Internet Edge peering capacity to reduce latency in the local Delhi metropolitan area. Since the last update, we have completed the augmentation of out-of-region Internet Edge regional peering capacity in Chennai to provide additional load-balancing and redundancy to large ISPs in India. Following safety clearance, our team has accessed the damaged site and is continuing to restore additional capacity throughout this week. We have optimized capacity across network backbones to increase available headroom and augmented our Delhi user-facing backbone capacity. We will continue to closely monitor latency deviations and packet drops. We will provide our next update by Monday, 2026-06-29 at 17:00 PDT. **Symptoms** Customers may experience slightly elevated latency and non-optimal network routing into Google Cloud until the affected facility is fully restored. **Workaround** There is no workaround at this time.
- `2026-06-29 23:06` **resolved** — **Summary** Network traffic to Google Cloud originating from Delhi, Chennai, Mumbai and surrounding areas experienced intermittent periods of elevated latency and possible packet loss. **Description** Traffic rerouting from the impacted Delhi facility caused a subset of Hybrid Connectivity, Virtual Private Cloud (VPC) and Media CDN customers to experience intermittent latency spikes as demand exceeded regional capacity. We completed the augmentation of out-of-region Internet Edge regional peering capacity in Chennai to provide additional load-balancing and redundancy to large ISPs in India. Service to a large portion of Internet Edge peering capacity has been restored to reduce latency in the local Delhi metropolitan area. We are now recovered and returned to normal service as of Friday, 2026-06-26 PDT. Following safety clearance, our team restored all lost capacity. We have restored capacity between Delhi-Chennai and Delhi-Mumbai and will continue to closely monitor latency deviations and packet drops. We thank you for your patience during the resolution of this issue. **Symptoms** The impacted customers may have experienced slightly elevated latency and non-optimal network routing into Google Cloud. **Workaround** None

### 2. yandexcloud — Incorrect charge of funds.

- URL: https://status.yandex.cloud/en/incidents/1824
- Started: 2026-06-05 12:25 UTC
- Resolved: 2026-06-05 20:55 UTC
- Duration: 8h 30m
- Impact: major, status: resolved
- Components: Yandex Cloud Billing
- Canonical incident: [ci-f4f2bb971e0fb64afd429e2b](#ci-f4f2bb971e0fb64afd429e2b)

- `2026-06-05 13:13` **investigating** — An error was found in calculating charges for some of the resources. Billing processing has been temporarily stopped to prevent incorrect charges. After the cause is eliminated, the necessary accrual adjustments will be made. No additional user actions are required.
- `2026-06-05 13:42` **unknown** — The team is engaged in unblocking previously stopped resources. As soon as availability is restored, users will need to manually launch these resources. We will inform you about the completion of the process additionally.
- `2026-06-05 14:56` **unknown** — Most of the clouds that were mistakenly blocked earlier have already been restored, and users will need to manually start the stopped resources themselves. The team is also working on unblocking mistakenly blocked billing accounts and displaying the details correctly; we expect the details to start displaying correctly by the end of the day.
- `2026-06-05 15:12` **unknown** — All cloud accounts and billing accounts that were mistakenly blocked have now been unblocked. We ask users to manually restart any resources that should be active. We are also continuing to work on restoring the correct display of billing details and expect it to be displayed correctly by the end of the day. Separately, the team is handling erroneous charges: lists of affected cases are currently being compiled, and refunds will be issued shortly. We will provide an additional update on this matter.
- `2026-06-05 15:34` **unknown** — We have already started processing refunds; the actual time it takes for the funds to be credited may depend on the card-issuing bank.
- `2026-06-05 16:02` **unknown** — The detail has been fixed. Erroneous charges have been canceled and no longer affect your personal account balance.
- `2026-06-05 18:01` **unknown** — We have refunded the funds that were mistakenly debited. The actual transfer of the refund may depend on the bank that issued the card.
- `2026-06-05 20:56` **resolved** — The billing systems have been restored.

### 3. stripe — Elevated declines on Cartes Bancaires

- URL: https://stspg.io/ydwywzw7gnp3
- Started: 2026-06-05 05:33 UTC
- Resolved: 2026-06-05 14:41 UTC
- Duration: 9h 8m
- Impact: none, status: resolved
- Components: Acquirers and payment methods
- Canonical incident: [ci-3f5cddd160fc320f88711ef2](#ci-3f5cddd160fc320f88711ef2)

- `2026-06-05 05:33` **identified** — We’re currently seeing elevated error rates and response times with some Cartes Bancaires payments due to an issuer-side degradation starting 07:23 UTC June 4th. We believe this outage is outside of Stripe's infrastructure. This is likely impacting a number of payment processors. We’re continuing to monitor and will post further updates.
- `2026-06-05 12:22` **identified** — We are continuing to see elevated error rates and response times with some Cartes Bancaires payments due to an issuer-side degradation starting 07:23 UTC June 4th. We are actively monitoring and will provide a further update as more information becomes available.
- `2026-06-05 14:41` **resolved** — Between 07:23 UTC on June 4th and 13:51 UTC on June 5th, we saw elevated error rates and response times with some Cartes Bancaires payments. This is now resolved.

### 4. yandexcloud — The problem is with the AI Studio service

- URL: https://status.yandex.cloud/en/incidents/1807
- Started: 2026-06-01 02:37 UTC
- Resolved: 2026-06-01 06:51 UTC
- Duration: 4h 14m
- Impact: minor, status: resolved
- Components: Yandex AI Studio
- Canonical incident: [ci-ab551a950d372dbbf26f0544](#ci-ab551a950d372dbbf26f0544)

- `2026-06-01 06:50` **investigating** — We noticed a partial unavailability of the AI Studio service in the Cloud. There may be 503 errors when working through the web console or API. The developers are already aware of it and are fixing it.
- `2026-06-01 08:08` **resolved** — At 09:51 Moscow time, the problem was fixed, and now the Cloud services are working correctly. We are investigating the causes of its occurrence.

### 5. yandexcloud — Partial unavailability of services

- URL: https://status.yandex.cloud/en/incidents/1822
- Started: 2026-06-04 08:30 UTC
- Resolved: 2026-06-04 11:30 UTC
- Duration: 3h 0m
- Impact: major, status: resolved
- Components: Yandex BareMetal
- Canonical incident: [ci-4b61b24fda7dd190cb06d6dc](#ci-4b61b24fda7dd190cb06d6dc)

- `2026-06-04 10:42` **investigating** — We have noticed a degradation of the VM's public network connectivity in the Baremetal service. The developers are already aware of it and are fixing it.
- `2026-06-04 13:33` **resolved** — At 15:30 Moscow time, network connectivity in the Baremetal service was restored. If you still have questions about cloud resources, please contact technical support.

### 6. yandexcloud — Problems getting a list of models in the AI Studio service

- URL: https://status.yandex.cloud/en/incidents/1820
- Started: 2026-06-03 13:08 UTC
- Resolved: 2026-06-03 15:21 UTC
- Duration: 2h 13m
- Impact: minor, status: resolved
- Components: Yandex AI Studio
- Canonical incident: [ci-bef7a344911503009bc9d64f](#ci-bef7a344911503009bc9d64f)

- `2026-06-03 15:02` **investigating** — We noticed that the API does not return the correct list of available models. The developers are already aware and are fixing it.
- `2026-06-03 16:10` **resolved** — At 6:21 p.m., the problem was resolved. If you still have questions about cloud resources, please contact technical support.

### 7. circleci — Mac jobs failing to start

- URL: https://stspg.io/n3brb41nn16r
- Started: 2026-06-01 01:04 UTC
- Resolved: 2026-06-01 03:26 UTC
- Duration: 2h 21m
- Impact: minor, status: resolved
- Components: macOS Jobs
- Canonical incident: [ci-694a4b0137b3cc4ada38f7d5](#ci-694a4b0137b3cc4ada38f7d5)

- `2026-06-01 01:04` **investigating** — Mac jobs are continuing to experience failures. Our infrastructure provider is continuing to report that they are experiencing a DDoS attack causing high network latency and packet loss for our Mac fleet. We are working to bring up additional capacity to restore service.
- `2026-06-01 01:05` **investigating** — We are continuing to investigate this issue.
- `2026-06-01 01:36` **monitoring** — We are seeing Mac jobs starting successfully and improved network connectivity to the Mac fleet, however our infrastructure provider is continuing to experience a DDoS attack. We will continue to monitor the recovery while bringing additional capacity to our platform.
- `2026-06-01 02:05` **monitoring** — Mac jobs are running successfully, and we are continuing to monitor the performance of our Mac fleet. Thank you for your patience.
- `2026-06-01 03:26` **resolved** — All Mac jobs are running successfully. Our infrastructure provider has confirmed that the DDoS attack has been mitigated and that services have been restored. Thank you for your patience during this incident.

### 8. circleci — macOS jobs failing to start

- URL: https://stspg.io/fcqx1c8m4dqx
- Started: 2026-06-07 04:20 UTC
- Resolved: 2026-06-07 04:44 UTC
- Duration: 24m
- Impact: major, status: resolved
- Components: macOS Jobs
- Canonical incident: [ci-774d17c7b13e6cbcd3bb0e45](#ci-774d17c7b13e6cbcd3bb0e45)

- `2026-06-07 04:20` **identified** — We have identified an issue causing macOS jobs to fail to start and are working to restore service. This impacts customers running macOS jobs across all resource classes. Thank you for your patience while our engineers work to resolve this. We will provide another update within 30 minutes.
- `2026-06-07 04:44` **resolved** — This issue has been resolved and macOS jobs across all resource classes are starting normally again. Customers whose jobs failed to start during this window may rerun them. We thank you for your patience while our team worked on implementing a fix.

## Below threshold (6)

- `digitalocean` [DNS API Service](https://stspg.io/c9tgc3txntpg) — 1h 34m
- `slack` [Trouble connecting or loading Slack](https://slack-status.com/2026-06/79438d4a833ca70d) — 48m
- `slack` [Administrators may experience errors in member management.](https://slack-status.com/2026-06/898527c2a0f5c9be) — 1h 1m
- `slack` [Partial Degradation Impacting Usergroups](https://slack-status.com/2026-05/8f583746d60de2fc) — 57m
- `circleci` [Errors on Cloud Docker Pulls](https://stspg.io/4yrvgyvkfqpy) — 19m
- `slack` [Connectivity Issues Affecting Some Users in Europe](https://slack-status.com/2026-06/2e2e3eb0560b8566) — 0m

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-cf47cdc71eaf54c1e652582b"></a>

- Canonical ID: `ci-cf47cdc71eaf54c1e652582b` (1 source records).
  - [gcp/5fGQt4VbkDnr3Yp8PXPr](../../data/2026/W23/gcp.json); vendor URL: https://status.cloud.google.com/incidents/5fGQt4VbkDnr3Yp8PXPr
- Start time basis: `vendor_reported`.

<a id="ci-f4f2bb971e0fb64afd429e2b"></a>

- Canonical ID: `ci-f4f2bb971e0fb64afd429e2b` (1 source records).
  - [yandexcloud/1824](../../data/2026/W23/yandexcloud.json); vendor URL: https://status.yandex.cloud/en/incidents/1824
- Start time basis: `vendor_reported`.

<a id="ci-3f5cddd160fc320f88711ef2"></a>

- Canonical ID: `ci-3f5cddd160fc320f88711ef2` (1 source records).
  - [stripe/jkxl171k6zq7](../../data/2026/W23/stripe.json); vendor URL: https://stspg.io/ydwywzw7gnp3
- Start time basis: `vendor_reported`.

<a id="ci-ab551a950d372dbbf26f0544"></a>

- Canonical ID: `ci-ab551a950d372dbbf26f0544` (1 source records).
  - [yandexcloud/1807](../../data/2026/W23/yandexcloud.json); vendor URL: https://status.yandex.cloud/en/incidents/1807
- Start time basis: `vendor_reported`.

<a id="ci-b1154dc554c1f8842a028a06"></a>

- Canonical ID: `ci-b1154dc554c1f8842a028a06` (1 source records).
  - [digitalocean/nnt4rfbgn1y6](../../data/2026/W23/digitalocean.json); vendor URL: https://stspg.io/c9tgc3txntpg
- Start time basis: `vendor_reported`.

<a id="ci-4b61b24fda7dd190cb06d6dc"></a>

- Canonical ID: `ci-4b61b24fda7dd190cb06d6dc` (1 source records).
  - [yandexcloud/1822](../../data/2026/W23/yandexcloud.json); vendor URL: https://status.yandex.cloud/en/incidents/1822
- Start time basis: `vendor_reported`.

<a id="ci-bef7a344911503009bc9d64f"></a>

- Canonical ID: `ci-bef7a344911503009bc9d64f` (1 source records).
  - [yandexcloud/1820](../../data/2026/W23/yandexcloud.json); vendor URL: https://status.yandex.cloud/en/incidents/1820
- Start time basis: `vendor_reported`.

<a id="ci-694a4b0137b3cc4ada38f7d5"></a>

- Canonical ID: `ci-694a4b0137b3cc4ada38f7d5` (1 source records).
  - [circleci/wcsh4hxwkrh7](../../data/2026/W23/circleci.json); vendor URL: https://stspg.io/n3brb41nn16r
- Start time basis: `vendor_reported`.

<a id="ci-5dae324bc852953ef2ac90df"></a>

- Canonical ID: `ci-5dae324bc852953ef2ac90df` (1 source records).
  - [slack/1560](../../data/2026/W23/slack.json); vendor URL: https://slack-status.com/2026-06/79438d4a833ca70d
- Start time basis: `vendor_reported`.

<a id="ci-4a0c6f1aa889d2599505d2fb"></a>

- Canonical ID: `ci-4a0c6f1aa889d2599505d2fb` (1 source records).
  - [slack/1561](../../data/2026/W23/slack.json); vendor URL: https://slack-status.com/2026-06/898527c2a0f5c9be
- Start time basis: `vendor_reported`.

<a id="ci-aada16fe86e9def51339da88"></a>

- Canonical ID: `ci-aada16fe86e9def51339da88` (1 source records).
  - [slack/1558](../../data/2026/W23/slack.json); vendor URL: https://slack-status.com/2026-05/8f583746d60de2fc
- Start time basis: `vendor_reported`.

<a id="ci-774d17c7b13e6cbcd3bb0e45"></a>

- Canonical ID: `ci-774d17c7b13e6cbcd3bb0e45` (1 source records).
  - [circleci/2rhbvp9382px](../../data/2026/W23/circleci.json); vendor URL: https://stspg.io/fcqx1c8m4dqx
- Start time basis: `vendor_reported`.

<a id="ci-e936f360e883ec67d4f556f7"></a>

- Canonical ID: `ci-e936f360e883ec67d4f556f7` (1 source records).
  - [circleci/b4jcpwt024vc](../../data/2026/W23/circleci.json); vendor URL: https://stspg.io/4yrvgyvkfqpy
- Start time basis: `vendor_reported`.

<a id="ci-3759171fb14e8538c3ca1c93"></a>

- Canonical ID: `ci-3759171fb14e8538c3ca1c93` (1 source records).
  - [slack/1559](../../data/2026/W23/slack.json); vendor URL: https://slack-status.com/2026-06/2e2e3eb0560b8566
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
