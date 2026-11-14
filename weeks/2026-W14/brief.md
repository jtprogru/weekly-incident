# Week 2026-W14 — top 3 of 3 above threshold

2026-03-30 to 2026-04-05 (UTC).

6 source records (6 saved revisions) → 6 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## 1. circleci — Usage API data for 4/1 is incomplete with multiple columns showing null values. Data prior to 4/1 is unaffected. We are investigating.

- URL: https://stspg.io/78v0qz8pybdn
- Window: 2026-04-02 19:19 to 2026-04-03 02:18 UTC (6h 58m)
- Impact: minor, status: resolved
- Components: CircleCI API

<a id="ci-0ac7ce081ad14a1e659aa129"></a>

- Canonical ID: `ci-0ac7ce081ad14a1e659aa129` (1 source records).
  - [circleci/lxkrhs9rmrxv](../../data/2026/W14/circleci.json); vendor URL: https://stspg.io/78v0qz8pybdn
- Start time basis: `vendor_reported`.

  - `2026-04-02 19:19` investigating — We are currently investigating this issue.
  - `2026-04-02 21:23` identified — The Usage API issue has been identified and the fix is being implemented.
  - `2026-04-03 02:18` resolved — The Usage API issue affecting 4/1 data has been resolved. Updated data will be available on 4/3. Thank you for your patience.

## 2. yandexcloud — Delays in making payments and opening a payment form

- URL: https://status.yandex.cloud/en/incidents/1661
- Window: 2026-04-01 12:50 to 2026-04-01 15:30 UTC (2h 40m)
- Impact: minor, status: resolved
- Components: Yandex Cloud Billing

<a id="ci-7047e3225e495fcf2d74fe71"></a>

- Canonical ID: `ci-7047e3225e495fcf2d74fe71` (1 source records).
  - [yandexcloud/1661](../../data/2026/W14/yandexcloud.json); vendor URL: https://status.yandex.cloud/en/incidents/1661
- Start time basis: `vendor_reported`.

  - `2026-04-01 14:32` investigating — We noticed partial unavailability and delays when opening the card payment form. The developers are already aware of it and are fixing it.
  - `2026-04-01 15:51` unknown — Since 18:30 Moscow time, there have been no delays in making payments. We are continuing to investigate the causes of the problem.
  - `2026-04-01 15:54` resolved — Since 18:30 Moscow time, there have been no delays in making payments. We are continuing to investigate the causes of the problem.

## 3. circleci — Some pipelines not triggering due to GitHub API failures

- URL: https://stspg.io/3q4dxfr93jhd
- Window: 2026-04-01 11:55 to 2026-04-01 12:23 UTC (28m)
- Impact: major, status: resolved
- Components: GitHub API Requests, Pipelines & Workflows

<a id="ci-46866dcee81c338e2e2d60fd"></a>

- Canonical ID: `ci-46866dcee81c338e2e2d60fd` (1 source records).
  - [circleci/hth87g8tfm2y](../../data/2026/W14/circleci.json); vendor URL: https://stspg.io/3q4dxfr93jhd
- Start time basis: `vendor_reported`.

  - `2026-04-01 11:55` investigating — Some pipelines may not trigger due to a high rate of API failures on one of GitHub's API endpoints. This includes some manually-triggered pipelines, some OAuth hooks and some scheduled pipelines.
  - `2026-04-01 11:58` identified — We are continuing to see very high error rates from GitHub's `GET /repos/{owner}/{repo}/commits` API.
  - `2026-04-01 12:12` monitoring — GitHub's `GET /repos/{owner}/{repo}/commits` API has appeared to recover. We will continue to monitor this.
  - `2026-04-01 12:23` resolved — It appears that GitHub's `GET /repos/{owner}/{repo}/commits` API is now stable and responsive.
