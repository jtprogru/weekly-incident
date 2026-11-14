# Week 2026-W16

2026-04-13 to 2026-04-19 (UTC). Generated 2026-11-14 00:00.

12 canonical incidents, 6 above threshold. All 14 sources responded.

12 source records (12 saved revisions) → 12 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | stripe | [Canadian payouts delay](https://stspg.io/bccfqr59yk3d) | 2026-04-14 17:20 | 22h 23m | none | Stripe core components | 1612 = 1343 × 1 × 1.2 |
| 2 | digitalocean | [DNS Resolution for .co TLD](https://stspg.io/2s124dt3v3gl) | 2026-04-17 21:22 | 16h 47m | minor | DNS | 806 = 1007 × 1 × 0.8 |
| 3 | digitalocean | [Managed Database Resizes](https://stspg.io/v90w741klgft) | 2026-04-14 12:53 | 2h 26m | minor | Global, Global | 234 = 146 × 2 × 0.8 |
| 4 | stripe | [SMS delivery issues in the US](https://stspg.io/hcll9q8xngg8) | 2026-04-18 03:48 | 2h 54m | none | Stripe core components | 209 = 174 × 1 × 1.2 |
| 5 | circleci | [Usage API data delayed this morning](https://stspg.io/h7yfy5vyf3dw) | 2026-04-18 12:07 | 2h 57m | none | CircleCI API | 177 = 177 × 1 × 1.0 |
| 6 | datadog | [Delayed Metrics](https://stspg.io/32s2j42tn8xz) | 2026-04-15 22:38 | 1h 59m | major | Metrics and Infra Monitoring | 119 = 119 × 1 × 1.0 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. stripe — Canadian payouts delay

- URL: https://stspg.io/bccfqr59yk3d
- Started: 2026-04-14 17:20 UTC
- Resolved: 2026-04-15 15:43 UTC
- Duration: 22h 23m
- Impact: none, status: resolved
- Components: Stripe core components
- Canonical incident: [ci-1933dcc4af3a82746edc49aa](#ci-1933dcc4af3a82746edc49aa)

- `2026-04-14 17:20` **investigating** — Due to a processing issue with our banking partner, some Canadian Payouts processed on April 14, 2026 are experiencing a delay of a few hours. We are actively working with our banking partner to resolve this issue to ensure all affected payouts land today.
- `2026-04-15 00:47` **monitoring** — As of 22:00 UTC, we have identified the root cause with our banking partner and expect affected payouts to land by the end of the day. We will continue to monitor closely and post further updates as needed.
- `2026-04-15 15:43` **resolved** — We have confirmed with our banking partner that the issue has been resolved. All impacted Canadian payouts scheduled for April 14th have successfully posted on the same day. Thank you for your patience.

### 2. digitalocean — DNS Resolution for .co TLD

- URL: https://stspg.io/2s124dt3v3gl
- Started: 2026-04-17 21:22 UTC
- Resolved: 2026-04-18 14:10 UTC
- Duration: 16h 47m
- Impact: minor, status: resolved
- Components: DNS
- Canonical incident: [ci-33a48a145e4e716654d644b6](#ci-33a48a145e4e716654d644b6)

- `2026-04-17 21:22` **identified** — Our Engineering team is aware of a widespread, external issue affecting .co top-level domain (TLD). While this incident originates outside of DigitalOcean's infrastructure, you may experience errors when querying a .co domain, regardless of the DNS resolver being used. Our Engineers are actively deploying temporary backend mitigations to help minimize the impact on our customers. We will continue to monitor the situation closely and post updates as more information becomes available.
- `2026-04-17 22:53` **monitoring** — Our Engineering team has implemented a fix to resolve the issue that was affecting the DNS resolution of .co top-level domain (TLD). During this time, users who are using DigitalOcean DNS resolvers in their resources should no longer experience issues related to the DNS resolution. Our Engineers are currently monitoring the situation. We will post an update as soon as the issue is fully resolved.
- `2026-04-18 14:10` **resolved** — Our Engineering team has confirmed that the issue affecting the DNS resolution of .co top-level domain (TLD) has been resolved. The DNS resolution for .co domains is now working as expected. We apologize for the inconvenience. If you continue to face any issues, please open a support ticket from within your account.

### 3. digitalocean — Managed Database Resizes

- URL: https://stspg.io/v90w741klgft
- Started: 2026-04-14 12:53 UTC
- Resolved: 2026-04-14 15:19 UTC
- Duration: 2h 26m
- Impact: minor, status: resolved
- Components: Global, Global
- Canonical incident: [ci-3b79f9fbd8faadbe2b225ef9](#ci-3b79f9fbd8faadbe2b225ef9)

- `2026-04-14 12:53` **investigating** — Our engineering team is investigating an issue impacting resize operations for Managed Databases. During this time, users may experience error when attempting to resize Managed Database via Cloud Control Panel and API in all regions. We apologize for the inconvenience and will share an update once we have more information.
- `2026-04-14 14:30` **monitoring** — Our Engineering team has taken action to mitigate the issue with resize operations for Managed Databases and implemented a fix . We are monitoring the situation and will post an update as soon as we confirm that the issue is fully resolved.
- `2026-04-14 15:19` **resolved** — Our Engineering team has resolved the issue with resize operations for Managed Databases and should now be operating normally. If you continue to experience problems, please open a ticket with our support team from within your Cloud Control Panel. We apologize for any inconvenience.

### 4. stripe — SMS delivery issues in the US

- URL: https://stspg.io/hcll9q8xngg8
- Started: 2026-04-18 03:48 UTC
- Resolved: 2026-04-18 06:42 UTC
- Duration: 2h 54m
- Impact: none, status: resolved
- Components: Stripe core components
- Canonical incident: [ci-7a8d762b562a7bbd99dc05a6](#ci-7a8d762b562a7bbd99dc05a6)

- `2026-04-18 03:48` **investigating** — We're currently experiencing SMS delivery issues in the US, affecting 2-Factor authentication for dashboard users as well as the OTP code for Link and Express users. We'll post an update as soon as we have it.
- `2026-04-18 04:40` **identified** — Since 01:35 UTC, we’ve been experiencing US SMS delivery issues affecting 2FA and OTP codes for Dashboard, Link, Issuing, and Express users. Non-SMS 2FA methods (security keys, authenticator apps) still work. We’re actively monitoring and will update when resolved.
- `2026-04-18 06:42` **resolved** — From 01:35 - 05:20 UTC, we saw elevated errors with SMS delivery in the US. This is now resolved.

### 5. circleci — Usage API data delayed this morning

- URL: https://stspg.io/h7yfy5vyf3dw
- Started: 2026-04-18 12:07 UTC
- Resolved: 2026-04-18 15:05 UTC
- Duration: 2h 57m
- Impact: none, status: resolved
- Components: CircleCI API
- Canonical incident: [ci-93b342eade73e708e9403d34](#ci-93b342eade73e708e9403d34)

- `2026-04-18 12:07` **investigating** — Usage API data is delayed this morning. Team is investigating.
- `2026-04-18 13:17` **identified** — The Usage API issue has been identified and a fix is being implemented. We expect to the issue be resolved within a few hours.
- `2026-04-18 15:05` **resolved** — The Usage API issue has been resolved. All data is now current. Appreciate your patience.

### 6. datadog — Delayed Metrics

- URL: https://stspg.io/32s2j42tn8xz
- Started: 2026-04-15 22:38 UTC
- Resolved: 2026-04-16 00:37 UTC
- Duration: 1h 59m
- Impact: major, status: resolved
- Components: Metrics and Infra Monitoring
- Canonical incident: [ci-355fcf3adbdaf9fd8c2f1f0a](#ci-355fcf3adbdaf9fd8c2f1f0a)

- `2026-04-15 22:38` **investigating** — We are investigating increased latency processing Metrics. As a result of this issue, some users may see delays or gaps for metrics on graphs. To prevent false monitor alerts due to delayed data, monitors affected by the delay will not notify and will automatically resume once current data is available. All other monitors will operate normally.
- `2026-04-15 23:16` **investigating** — We are investigating increased latency processing Distribution Metrics. As a result of this issue, some users may see delays or gaps for distribution metrics on graphs. To prevent false monitor alerts due to delayed data, monitors affected by the delay will not notify and will automatically resume once current data is available. All other monitors will operate normally.
- `2026-04-15 23:54` **investigating** — We are still investigating increased latency processing Distribution Metrics, delays have gone down but have not fully recovered.
- `2026-04-16 00:32` **monitoring** — We have resolved the lag in distribution metrics and we are monitoring the results. We will provide another update once the issue is fully resolved.
- `2026-04-16 00:37` **resolved** — This incident has been resolved.

## Below threshold (6)

- `digitalocean` [Serverless Inference](https://stspg.io/yhy97jfdzszn) — 1h 48m
- `stripe` [Elevated BLIK errors](https://stspg.io/wgrdkk5szv89) — 1h 3m
- `stripe` [Elevated Swish errors](https://stspg.io/b83lnsn0ls35) — 53m
- `stripe` [Elevated Swish errors](https://stspg.io/bs93ly27kw4x) — 38m
- `digitalocean` [App Platform Deployments](https://stspg.io/qk6vnqw22fph) — 0m
- `digitalocean` [Spaces availability in NYC3](https://stspg.io/nftmlzbn2kd4) — 0m

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-1933dcc4af3a82746edc49aa"></a>

- Canonical ID: `ci-1933dcc4af3a82746edc49aa` (1 source records).
  - [stripe/ldd0qtrsq6c0](../../data/2026/W16/stripe.json); vendor URL: https://stspg.io/bccfqr59yk3d
- Start time basis: `vendor_reported`.

<a id="ci-33a48a145e4e716654d644b6"></a>

- Canonical ID: `ci-33a48a145e4e716654d644b6` (1 source records).
  - [digitalocean/v6qdbpt1dcd3](../../data/2026/W16/digitalocean.json); vendor URL: https://stspg.io/2s124dt3v3gl
- Start time basis: `vendor_reported`.

<a id="ci-3b79f9fbd8faadbe2b225ef9"></a>

- Canonical ID: `ci-3b79f9fbd8faadbe2b225ef9` (1 source records).
  - [digitalocean/8hj57rv32g6n](../../data/2026/W16/digitalocean.json); vendor URL: https://stspg.io/v90w741klgft
- Start time basis: `vendor_reported`.

<a id="ci-7a8d762b562a7bbd99dc05a6"></a>

- Canonical ID: `ci-7a8d762b562a7bbd99dc05a6` (1 source records).
  - [stripe/3y9q97y8cl7r](../../data/2026/W16/stripe.json); vendor URL: https://stspg.io/hcll9q8xngg8
- Start time basis: `vendor_reported`.

<a id="ci-93b342eade73e708e9403d34"></a>

- Canonical ID: `ci-93b342eade73e708e9403d34` (1 source records).
  - [circleci/2v2q3ht88fh1](../../data/2026/W16/circleci.json); vendor URL: https://stspg.io/h7yfy5vyf3dw
- Start time basis: `vendor_reported`.

<a id="ci-355fcf3adbdaf9fd8c2f1f0a"></a>

- Canonical ID: `ci-355fcf3adbdaf9fd8c2f1f0a` (1 source records).
  - [datadog/7ht7d7xsnbhy](../../data/2026/W16/datadog.json); vendor URL: https://stspg.io/32s2j42tn8xz
- Start time basis: `vendor_reported`.

<a id="ci-85897c7312bdf9279b0173e8"></a>

- Canonical ID: `ci-85897c7312bdf9279b0173e8` (1 source records).
  - [digitalocean/cxh9f0hgw934](../../data/2026/W16/digitalocean.json); vendor URL: https://stspg.io/yhy97jfdzszn
- Start time basis: `vendor_reported`.

<a id="ci-d2c57b47b82f481422deb1c6"></a>

- Canonical ID: `ci-d2c57b47b82f481422deb1c6` (1 source records).
  - [stripe/7kmbgqhlvykm](../../data/2026/W16/stripe.json); vendor URL: https://stspg.io/wgrdkk5szv89
- Start time basis: `vendor_reported`.

<a id="ci-d08f090bcb5e045906680ce6"></a>

- Canonical ID: `ci-d08f090bcb5e045906680ce6` (1 source records).
  - [stripe/dv3xdy81j9hh](../../data/2026/W16/stripe.json); vendor URL: https://stspg.io/b83lnsn0ls35
- Start time basis: `vendor_reported`.

<a id="ci-5ba6970dfe23f89fc20be1d2"></a>

- Canonical ID: `ci-5ba6970dfe23f89fc20be1d2` (1 source records).
  - [stripe/hzlysrxhzmh9](../../data/2026/W16/stripe.json); vendor URL: https://stspg.io/bs93ly27kw4x
- Start time basis: `vendor_reported`.

<a id="ci-b095b1e779b22d692136234e"></a>

- Canonical ID: `ci-b095b1e779b22d692136234e` (1 source records).
  - [digitalocean/4cc9krk5gy9p](../../data/2026/W16/digitalocean.json); vendor URL: https://stspg.io/qk6vnqw22fph
- Start time basis: `vendor_reported`.

<a id="ci-2bc0712ab4a9892f9ea69b05"></a>

- Canonical ID: `ci-2bc0712ab4a9892f9ea69b05` (1 source records).
  - [digitalocean/8whvfslclskl](../../data/2026/W16/digitalocean.json); vendor URL: https://stspg.io/nftmlzbn2kd4
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
