# Week 2026-W23 — top 5 of 8 above threshold

2026-06-01 to 2026-06-07 (UTC).

14 source records (14 saved revisions) → 14 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## 1. gcp — Network traffic to Google Cloud originating from Delhi, Chennai, Mumbai and surrounding areas is experiencing intermittent periods of elevated latency and possible packet loss.

- URL: https://status.cloud.google.com/incidents/5fGQt4VbkDnr3Yp8PXPr
- Window: 2026-06-05 07:00 to 2026-06-26 19:00 UTC (21d 12h 0m)
- Impact: major, status: resolved
- Components: Hybrid Connectivity, Media CDN, Virtual Private Cloud (VPC)

<a id="ci-cf47cdc71eaf54c1e652582b"></a>

- Canonical ID: `ci-cf47cdc71eaf54c1e652582b` (1 source records).
  - [gcp/5fGQt4VbkDnr3Yp8PXPr](../../data/2026/W23/gcp.json); vendor URL: https://status.cloud.google.com/incidents/5fGQt4VbkDnr3Yp8PXPr
- Start time basis: `vendor_reported`.

  - `2026-06-10 01:13` investigating — **Summary** Network traffic to Google Cloud originating from Delhi, Chennai, Mumbai and surrounding areas is experiencing intermittent periods of elevated latency and possible packet loss. **Description** A fire at a third-party data center facility required an emergency power shutdown of networking equipment, isolating a non-compute local Point of Presence (POP) in Delhi and reducing available network capacity in the metro area. We rerouted significant traffic from the impacted facility in Delhi to address reduced local serving capabilities. As a result, a subset of Hybrid Connectivity and Virtual Private Cloud (VPC) customers may be impacted by the routing changes made to address reduced local, latency-optimized serving capabilities in Delhi. Affected customers may experience intermitten…
  - `2026-06-10 22:25` investigating — **Summary** Network traffic to Google Cloud originating from Delhi, Chennai, Mumbai and surrounding areas is experiencing intermittent periods of elevated latency and possible packet loss. **Description** A fire at a third-party data center facility required an emergency power shutdown of networking equipment, isolating a non-compute local Point of Presence (POP) in Delhi and reducing available network capacity in the metro area. We rerouted significant traffic from the impacted facility in Delhi to address reduced local serving capabilities. As a result, a subset of Hybrid Connectivity and Virtual Private Cloud (VPC) customers may be impacted by the routing changes made to address reduced local, latency-optimized serving capabilities in Delhi. Affected customers may experience intermitten…
  - `2026-06-11 23:25` investigating — **Summary** Network traffic to Google Cloud originating from Delhi, Chennai, Mumbai and surrounding areas is experiencing intermittent periods of elevated latency and possible packet loss. **Description** A fire at a third-party data center facility required an emergency power shutdown of networking equipment, isolating a non-compute local Point of Presence (POP) in Delhi and reducing available network capacity in the metro area. We rerouted significant traffic from the impacted facility in Delhi to address reduced local serving capabilities. As a result, a subset of Hybrid Connectivity and Virtual Private Cloud (VPC) customers may be impacted by the routing changes made to address reduced local, latency-optimized serving capabilities in Delhi. Affected customers may experience intermitten…
  - `2026-06-12 22:50` investigating — **Summary** Network traffic to Google Cloud originating from Delhi, Chennai, Mumbai and surrounding areas is experiencing intermittent periods of elevated latency and possible packet loss. **Description** A fire at a third-party data center facility required an emergency power shutdown of networking equipment, isolating a non-compute local Point of Presence (POP) in Delhi and reducing available network capacity in the metro area. We rerouted significant traffic from the impacted facility in Delhi to address reduced local serving capabilities. As a result, a subset of Hybrid Connectivity and Virtual Private Cloud (VPC) customers may be impacted by the routing changes made to address reduced local, latency-optimized serving capabilities in Delhi. Affected customers may experience intermitten…
  - `2026-06-15 23:11` investigating — **Summary** Network traffic to Google Cloud originating from Delhi, Chennai, Mumbai and surrounding areas is experiencing intermittent periods of elevated latency and possible packet loss. **Description** A fire at a third-party data center facility required an emergency power shutdown of networking equipment, isolating a non-compute local Point of Presence (POP) in Delhi and reducing available network capacity in the metro area. We rerouted significant traffic from the impacted facility in Delhi to address reduced local serving capabilities. As a result, a subset of Hybrid Connectivity and Virtual Private Cloud (VPC) customers may be impacted by the routing changes made to address reduced local, latency-optimized serving capabilities in Delhi. Affected customers may experience intermitten…
  - `2026-06-17 22:36` investigating — **Summary** Network traffic to Google Cloud originating from Delhi, Chennai, Mumbai and surrounding areas is experiencing intermittent periods of elevated latency and possible packet loss. **Description** Traffic rerouting from the impacted Delhi facility has caused a subset of Hybrid Connectivity and Virtual Private Cloud (VPC) customers to experience intermittent latency spikes as demand exceeds regional capacity. Media CDN customers may also notice increased latencies. Initial traffic mitigations have yielded positive results for some Cloud customers. We have restored a portion of Internet Edge peering capacity to reduce latency in the local Delhi metropolitan area. Further, we are nearly finished with the augmentation of out-of-region Internet Edge regional peering capacity in Chennai…
  - `2026-06-22 23:51` investigating — **Summary** Network traffic to Google Cloud originating from Delhi, Chennai, Mumbai and surrounding areas is experiencing intermittent periods of elevated latency and possible packet loss. **Description** Traffic rerouting from the impacted Delhi facility has caused a subset of Hybrid Connectivity and Virtual Private Cloud (VPC) customers to experience intermittent latency spikes as demand exceeds regional capacity. Media CDN customers may also notice increased latencies. Initial traffic mitigations have yielded positive results for some Cloud customers. We have restored a large portion of Internet Edge peering capacity to reduce latency in the local Delhi metropolitan area. Since the last update, we have completed the augmentation of out-of-region Internet Edge regional peering capacity i…
  - `2026-06-23 22:52` investigating — **Summary** Network traffic to Google Cloud originating from Delhi, Chennai, Mumbai and surrounding areas is experiencing intermittent periods of elevated latency and possible packet loss. **Description** Traffic rerouting from the impacted Delhi facility has caused a subset of Hybrid Connectivity and Virtual Private Cloud (VPC) customers to experience intermittent latency spikes as demand exceeds regional capacity. Media CDN customers may also notice increased latencies. Initial traffic mitigations have yielded positive results for some Cloud customers. We have restored a large portion of Internet Edge peering capacity to reduce latency in the local Delhi metropolitan area. Since the last update, we have completed the augmentation of out-of-region Internet Edge regional peering capacity i…
  - `2026-06-29 23:06` resolved — **Summary** Network traffic to Google Cloud originating from Delhi, Chennai, Mumbai and surrounding areas experienced intermittent periods of elevated latency and possible packet loss. **Description** Traffic rerouting from the impacted Delhi facility caused a subset of Hybrid Connectivity, Virtual Private Cloud (VPC) and Media CDN customers to experience intermittent latency spikes as demand exceeded regional capacity. We completed the augmentation of out-of-region Internet Edge regional peering capacity in Chennai to provide additional load-balancing and redundancy to large ISPs in India. Service to a large portion of Internet Edge peering capacity has been restored to reduce latency in the local Delhi metropolitan area. We are now recovered and returned to normal service as of Friday, 2…

## 2. yandexcloud — Incorrect charge of funds.

- URL: https://status.yandex.cloud/en/incidents/1824
- Window: 2026-06-05 12:25 to 2026-06-05 20:55 UTC (8h 30m)
- Impact: major, status: resolved
- Components: Yandex Cloud Billing

<a id="ci-f4f2bb971e0fb64afd429e2b"></a>

- Canonical ID: `ci-f4f2bb971e0fb64afd429e2b` (1 source records).
  - [yandexcloud/1824](../../data/2026/W23/yandexcloud.json); vendor URL: https://status.yandex.cloud/en/incidents/1824
- Start time basis: `vendor_reported`.

  - `2026-06-05 13:13` investigating — An error was found in calculating charges for some of the resources. Billing processing has been temporarily stopped to prevent incorrect charges. After the cause is eliminated, the necessary accrual adjustments will be made. No additional user actions are required.
  - `2026-06-05 13:42` unknown — The team is engaged in unblocking previously stopped resources. As soon as availability is restored, users will need to manually launch these resources. We will inform you about the completion of the process additionally.
  - `2026-06-05 14:56` unknown — Most of the clouds that were mistakenly blocked earlier have already been restored, and users will need to manually start the stopped resources themselves. The team is also working on unblocking mistakenly blocked billing accounts and displaying the details correctly; we expect the details to start displaying correctly by the end of the day.
  - `2026-06-05 15:12` unknown — All cloud accounts and billing accounts that were mistakenly blocked have now been unblocked. We ask users to manually restart any resources that should be active. We are also continuing to work on restoring the correct display of billing details and expect it to be displayed correctly by the end of the day. Separately, the team is handling erroneous charges: lists of affected cases are currently being compiled, and refunds will be issued shortly. We will provide an additional update on this matter.
  - `2026-06-05 15:34` unknown — We have already started processing refunds; the actual time it takes for the funds to be credited may depend on the card-issuing bank.
  - `2026-06-05 16:02` unknown — The detail has been fixed. Erroneous charges have been canceled and no longer affect your personal account balance.
  - `2026-06-05 18:01` unknown — We have refunded the funds that were mistakenly debited. The actual transfer of the refund may depend on the bank that issued the card.
  - `2026-06-05 20:56` resolved — The billing systems have been restored.

## 3. stripe — Elevated declines on Cartes Bancaires

- URL: https://stspg.io/ydwywzw7gnp3
- Window: 2026-06-05 05:33 to 2026-06-05 14:41 UTC (9h 8m)
- Impact: none, status: resolved
- Components: Acquirers and payment methods

<a id="ci-3f5cddd160fc320f88711ef2"></a>

- Canonical ID: `ci-3f5cddd160fc320f88711ef2` (1 source records).
  - [stripe/jkxl171k6zq7](../../data/2026/W23/stripe.json); vendor URL: https://stspg.io/ydwywzw7gnp3
- Start time basis: `vendor_reported`.

  - `2026-06-05 05:33` identified — We’re currently seeing elevated error rates and response times with some Cartes Bancaires payments due to an issuer-side degradation starting 07:23 UTC June 4th. We believe this outage is outside of Stripe's infrastructure. This is likely impacting a number of payment processors. We’re continuing to monitor and will post further updates.
  - `2026-06-05 12:22` identified — We are continuing to see elevated error rates and response times with some Cartes Bancaires payments due to an issuer-side degradation starting 07:23 UTC June 4th. We are actively monitoring and will provide a further update as more information becomes available.
  - `2026-06-05 14:41` resolved — Between 07:23 UTC on June 4th and 13:51 UTC on June 5th, we saw elevated error rates and response times with some Cartes Bancaires payments. This is now resolved.

## 4. yandexcloud — The problem is with the AI Studio service

- URL: https://status.yandex.cloud/en/incidents/1807
- Window: 2026-06-01 02:37 to 2026-06-01 06:51 UTC (4h 14m)
- Impact: minor, status: resolved
- Components: Yandex AI Studio

<a id="ci-ab551a950d372dbbf26f0544"></a>

- Canonical ID: `ci-ab551a950d372dbbf26f0544` (1 source records).
  - [yandexcloud/1807](../../data/2026/W23/yandexcloud.json); vendor URL: https://status.yandex.cloud/en/incidents/1807
- Start time basis: `vendor_reported`.

  - `2026-06-01 06:50` investigating — We noticed a partial unavailability of the AI Studio service in the Cloud. There may be 503 errors when working through the web console or API. The developers are already aware of it and are fixing it.
  - `2026-06-01 08:08` resolved — At 09:51 Moscow time, the problem was fixed, and now the Cloud services are working correctly. We are investigating the causes of its occurrence.

## 5. yandexcloud — Partial unavailability of services

- URL: https://status.yandex.cloud/en/incidents/1822
- Window: 2026-06-04 08:30 to 2026-06-04 11:30 UTC (3h 0m)
- Impact: major, status: resolved
- Components: Yandex BareMetal

<a id="ci-4b61b24fda7dd190cb06d6dc"></a>

- Canonical ID: `ci-4b61b24fda7dd190cb06d6dc` (1 source records).
  - [yandexcloud/1822](../../data/2026/W23/yandexcloud.json); vendor URL: https://status.yandex.cloud/en/incidents/1822
- Start time basis: `vendor_reported`.

  - `2026-06-04 10:42` investigating — We have noticed a degradation of the VM's public network connectivity in the Baremetal service. The developers are already aware of it and are fixing it.
  - `2026-06-04 13:33` resolved — At 15:30 Moscow time, network connectivity in the Baremetal service was restored. If you still have questions about cloud resources, please contact technical support.

## Also above threshold (3)

- `yandexcloud` Problems getting a list of models in the AI Studio service (2h 13m) — https://status.yandex.cloud/en/incidents/1820
- `circleci` Mac jobs failing to start (2h 21m) — https://stspg.io/n3brb41nn16r
- `circleci` macOS jobs failing to start (24m) — https://stspg.io/fcqx1c8m4dqx
