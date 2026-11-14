# Week 2026-W14

2026-03-30 to 2026-04-05 (UTC). Generated 2026-11-14 00:00.

6 canonical incidents, 3 above threshold. All 14 sources responded.

6 source records (6 saved revisions) → 6 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | circleci | [Usage API data for 4/1 is incomplete with multiple columns showing null values. Data prior to 4/1 is unaffected. We are investigating.](https://stspg.io/78v0qz8pybdn) | 2026-04-02 19:19 | 6h 58m | minor | CircleCI API | 418 = 418 × 1 × 1.0 |
| 2 | yandexcloud | [Delays in making payments and opening a payment form](https://status.yandex.cloud/en/incidents/1661) | 2026-04-01 12:50 | 2h 40m | minor | Yandex Cloud Billing | 240 = 160 × 1 × 1.5 |
| 3 | circleci | [Some pipelines not triggering due to GitHub API failures](https://stspg.io/3q4dxfr93jhd) | 2026-04-01 11:55 | 28m | major | GitHub API Requests, Pipelines & Workflows | 56 = 28 × 2 × 1.0 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. circleci — Usage API data for 4/1 is incomplete with multiple columns showing null values. Data prior to 4/1 is unaffected. We are investigating.

- URL: https://stspg.io/78v0qz8pybdn
- Started: 2026-04-02 19:19 UTC
- Resolved: 2026-04-03 02:18 UTC
- Duration: 6h 58m
- Impact: minor, status: resolved
- Components: CircleCI API
- Canonical incident: [ci-0ac7ce081ad14a1e659aa129](#ci-0ac7ce081ad14a1e659aa129)

- `2026-04-02 19:19` **investigating** — We are currently investigating this issue.
- `2026-04-02 21:23` **identified** — The Usage API issue has been identified and the fix is being implemented.
- `2026-04-03 02:18` **resolved** — The Usage API issue affecting 4/1 data has been resolved. Updated data will be available on 4/3. Thank you for your patience.

### 2. yandexcloud — Delays in making payments and opening a payment form

- URL: https://status.yandex.cloud/en/incidents/1661
- Started: 2026-04-01 12:50 UTC
- Resolved: 2026-04-01 15:30 UTC
- Duration: 2h 40m
- Impact: minor, status: resolved
- Components: Yandex Cloud Billing
- Canonical incident: [ci-7047e3225e495fcf2d74fe71](#ci-7047e3225e495fcf2d74fe71)

- `2026-04-01 14:32` **investigating** — We noticed partial unavailability and delays when opening the card payment form. The developers are already aware of it and are fixing it.
- `2026-04-01 15:51` **unknown** — Since 18:30 Moscow time, there have been no delays in making payments. We are continuing to investigate the causes of the problem.
- `2026-04-01 15:54` **resolved** — Since 18:30 Moscow time, there have been no delays in making payments. We are continuing to investigate the causes of the problem.

### 3. circleci — Some pipelines not triggering due to GitHub API failures

- URL: https://stspg.io/3q4dxfr93jhd
- Started: 2026-04-01 11:55 UTC
- Resolved: 2026-04-01 12:23 UTC
- Duration: 28m
- Impact: major, status: resolved
- Components: GitHub API Requests, Pipelines & Workflows
- Canonical incident: [ci-46866dcee81c338e2e2d60fd](#ci-46866dcee81c338e2e2d60fd)

- `2026-04-01 11:55` **investigating** — Some pipelines may not trigger due to a high rate of API failures on one of GitHub's API endpoints. This includes some manually-triggered pipelines, some OAuth hooks and some scheduled pipelines.
- `2026-04-01 11:58` **identified** — We are continuing to see very high error rates from GitHub's `GET /repos/{owner}/{repo}/commits` API.
- `2026-04-01 12:12` **monitoring** — GitHub's `GET /repos/{owner}/{repo}/commits` API has appeared to recover. We will continue to monitor this.
- `2026-04-01 12:23` **resolved** — It appears that GitHub's `GET /repos/{owner}/{repo}/commits` API is now stable and responsive.

## Below threshold (3)

- `slack` [Slackbot button in the top navigation is disappearing](https://slack-status.com/2026-04/3e4a3dd1224c4b7d) — 1h 40m
- `stripe` [Elevated Dashboard, Express Dashboard and Mobile surfaces errors](https://stspg.io/p8x22czdkx15) — 1h 11m
- `stripe` [Elevated Bancontact errors](https://stspg.io/r0q1hjmhz3q6) — 44m

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-0ac7ce081ad14a1e659aa129"></a>

- Canonical ID: `ci-0ac7ce081ad14a1e659aa129` (1 source records).
  - [circleci/lxkrhs9rmrxv](../../data/2026/W14/circleci.json); vendor URL: https://stspg.io/78v0qz8pybdn
- Start time basis: `vendor_reported`.

<a id="ci-7047e3225e495fcf2d74fe71"></a>

- Canonical ID: `ci-7047e3225e495fcf2d74fe71` (1 source records).
  - [yandexcloud/1661](../../data/2026/W14/yandexcloud.json); vendor URL: https://status.yandex.cloud/en/incidents/1661
- Start time basis: `vendor_reported`.

<a id="ci-dfd9d23ef519c631f22b184f"></a>

- Canonical ID: `ci-dfd9d23ef519c631f22b184f` (1 source records).
  - [slack/1541](../../data/2026/W14/slack.json); vendor URL: https://slack-status.com/2026-04/3e4a3dd1224c4b7d
- Start time basis: `vendor_reported`.

<a id="ci-a54be0dbd8e1323f054f636c"></a>

- Canonical ID: `ci-a54be0dbd8e1323f054f636c` (1 source records).
  - [stripe/2tcpcfl8rp6j](../../data/2026/W14/stripe.json); vendor URL: https://stspg.io/p8x22czdkx15
- Start time basis: `vendor_reported`.

<a id="ci-46866dcee81c338e2e2d60fd"></a>

- Canonical ID: `ci-46866dcee81c338e2e2d60fd` (1 source records).
  - [circleci/hth87g8tfm2y](../../data/2026/W14/circleci.json); vendor URL: https://stspg.io/3q4dxfr93jhd
- Start time basis: `vendor_reported`.

<a id="ci-406dd78ae55871de613686f9"></a>

- Canonical ID: `ci-406dd78ae55871de613686f9` (1 source records).
  - [stripe/2jz281x5q65q](../../data/2026/W14/stripe.json); vendor URL: https://stspg.io/r0q1hjmhz3q6
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
| gcp | ok | 7 | 0 |
| github | ok | 50 | 0 |
| grafana | ok | 50 | 0 |
| openai | ok | 25 | 0 |
| slack | ok | 35 | 0 |
| stripe | ok | 50 | 0 |
| yandexcloud | ok | 50 | 0 |
