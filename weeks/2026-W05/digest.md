# Week 2026-W05

2026-01-26 to 2026-02-01 (UTC). Generated 2026-11-14 00:00.

5 canonical incidents, 4 above threshold. All 14 sources responded.

5 source records (5 saved revisions) → 5 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | datadog | [Monitors - Delayed Evaluation](https://stspg.io/55l2cqg035fj) | 2026-01-28 20:30 | 3h 5m | minor | Monitors | 185 = 185 × 1 × 1.0 |
| 2 | stripe | [Elevated errors on 3DS transactions for Commonwealth Bank of Australia](https://stspg.io/xhlp9h3kzr5y) | 2026-01-29 10:18 | 2h 14m | minor | Acquirers and payment methods | 161 = 134 × 1 × 1.2 |
| 3 | datadog | [Delayed Distribution Monitors Evaluations](https://stspg.io/kywknq2z0ht1) | 2026-01-29 17:35 | 2h 19m | critical | Monitors | 139 = 139 × 1 × 1.0 |
| 4 | stripe | [Support site not loading](https://stspg.io/2nnjjm8kjhfb) | 2026-01-26 23:25 | 7m | major | Stripe core components | 8 = 7 × 1 × 1.2 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. datadog — Monitors - Delayed Evaluation

- URL: https://stspg.io/55l2cqg035fj
- Started: 2026-01-28 20:30 UTC
- Resolved: 2026-01-28 23:36 UTC
- Duration: 3h 5m
- Impact: minor, status: resolved
- Components: Monitors
- Canonical incident: [ci-5d50d684d5519ed975ba48a5](#ci-5d50d684d5519ed975ba48a5)

- `2026-01-28 21:18` **investigating** — We are investigating delays in service checks monitors evaluation, which began at 20:26 1/28/2026 UTC.
- `2026-01-28 21:21` **investigating** — We are continuing to investigate this issue.
- `2026-01-28 22:08` **identified** — The issue has been identified and a fix is being implemented.
- `2026-01-28 23:13` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-01-28 23:36` **resolved** — This incident has been resolved.

### 2. stripe — Elevated errors on 3DS transactions for Commonwealth Bank of Australia

- URL: https://stspg.io/xhlp9h3kzr5y
- Started: 2026-01-29 10:18 UTC
- Resolved: 2026-01-29 12:32 UTC
- Duration: 2h 14m
- Impact: minor, status: resolved
- Components: Acquirers and payment methods
- Canonical incident: [ci-fbe9c9a978b329b1553ae88b](#ci-fbe9c9a978b329b1553ae88b)

- `2026-01-29 10:18` **investigating** — We are currently investigating elevated failures for 3DS transactions on Commonwealth Bank of Australia issued cards starting at 09:07 UTC. This issue is external to Stripe and we are actively monitoring the situation. We will provide updates as more information becomes available.
- `2026-01-29 12:32` **resolved** — Between 09:07 UTC and 11:54 UTC we observed elevated failures for 3DS transactions on Commonwealth Bank of Australia issued cards. This issue was external to Stripe and is now resolved.

### 3. datadog — Delayed Distribution Monitors Evaluations

- URL: https://stspg.io/kywknq2z0ht1
- Started: 2026-01-29 17:35 UTC
- Resolved: 2026-01-29 19:55 UTC
- Duration: 2h 19m
- Impact: critical, status: resolved
- Components: Monitors
- Canonical incident: [ci-31730b4041adaf2c1cb1289c](#ci-31730b4041adaf2c1cb1289c)

- `2026-01-29 17:35` **investigating** — We are investigating delays in Monitors evaluations, which began at 17:15 UTC.
- `2026-01-29 17:46` **identified** — We have identified the underlying issue and are working on a fix. It is important to note that no data has been lost, and evaluations will be caught up once the service is operational again.
- `2026-01-29 18:06` **identified** — We are continuing to work on a fix for this issue. It is important to note that no data has been lost, and evaluations will be caught up once the service is operational again.
- `2026-01-29 18:36` **monitoring** — We have deployed a fix and we are monitoring the results. We will continue to provide regular updates.
- `2026-01-29 19:09` **monitoring** — We are continuing to monitor the fix and will continue to provide regular updates.
- `2026-01-29 19:55` **resolved** — This incident has been resolved.

### 4. stripe — Support site not loading

- URL: https://stspg.io/2nnjjm8kjhfb
- Started: 2026-01-26 23:25 UTC
- Resolved: 2026-01-26 23:32 UTC
- Duration: 7m
- Impact: major, status: resolved
- Components: Stripe core components
- Canonical incident: [ci-7861671c3a91772806dd6630](#ci-7861671c3a91772806dd6630)

- `2026-01-26 23:25` **investigating** — We're currently investigating an issue starting at 22:58 UTC, wherein support.stripe.com is not loading and users are seeing 'service not found' on the page. You can still contact Stripe support by using the Stripe Dashboard help icon. We'll post an update as soon as the issue is resolved.
- `2026-01-26 23:32` **resolved** — There was an issue between 22:58 to 23:29 UTC, wherein support.stripe.com was not loading and users were seeing 'service not found' on the page. The issue is now resolved.

## Below threshold (1)

- `stripe` [Elevated Bancontact errors](https://stspg.io/yxxcm2f8sltb) — 38m

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-5d50d684d5519ed975ba48a5"></a>

- Canonical ID: `ci-5d50d684d5519ed975ba48a5` (1 source records).
  - [datadog/6vndlbw8vzqp](../../data/2026/W05/datadog.json); vendor URL: https://stspg.io/55l2cqg035fj
- Start time basis: `vendor_reported`.

<a id="ci-fbe9c9a978b329b1553ae88b"></a>

- Canonical ID: `ci-fbe9c9a978b329b1553ae88b` (1 source records).
  - [stripe/s61prt5s7pyy](../../data/2026/W05/stripe.json); vendor URL: https://stspg.io/xhlp9h3kzr5y
- Start time basis: `vendor_reported`.

<a id="ci-31730b4041adaf2c1cb1289c"></a>

- Canonical ID: `ci-31730b4041adaf2c1cb1289c` (1 source records).
  - [datadog/1hmvg62sxmjp](../../data/2026/W05/datadog.json); vendor URL: https://stspg.io/kywknq2z0ht1
- Start time basis: `vendor_reported`.

<a id="ci-bfdc1349fad6218350d0d4b0"></a>

- Canonical ID: `ci-bfdc1349fad6218350d0d4b0` (1 source records).
  - [stripe/2jhh765rxrdd](../../data/2026/W05/stripe.json); vendor URL: https://stspg.io/yxxcm2f8sltb
- Start time basis: `vendor_reported`.

<a id="ci-7861671c3a91772806dd6630"></a>

- Canonical ID: `ci-7861671c3a91772806dd6630` (1 source records).
  - [stripe/pwzz43xcw9f5](../../data/2026/W05/stripe.json); vendor URL: https://stspg.io/2nnjjm8kjhfb
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
