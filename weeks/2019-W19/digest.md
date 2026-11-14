# Week 2019-W19

2019-05-06 to 2019-05-12 (UTC). Generated 2026-11-14 00:00.

2 canonical incidents, 2 above threshold. All 13 sources responded.

2 source records (2 saved revisions) → 2 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | atlassian | [Jira Users Receiving HTTP 500 Errors](http://stspg.io/cb7cd57ba) | 2019-05-10 14:10 | 26m | major | — | 26 = 26 × 1 × 1.0 |
| 2 | atlassian | [Jira Imports are Failing](http://stspg.io/06e0ed12a) | 2019-05-07 09:45 | 0m | major | — | 0 = 0 × 1 × 1.0 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. atlassian — Jira Users Receiving HTTP 500 Errors

- URL: http://stspg.io/cb7cd57ba
- Started: 2019-05-10 14:10 UTC
- Resolved: 2019-05-10 14:37 UTC
- Duration: 26m
- Impact: major, status: resolved
- Components: —
- Canonical incident: [ci-e659433cc9489712bacbf039](#ci-e659433cc9489712bacbf039)

- `2019-05-10 14:10` **investigating** — We are investigating an issue with Jira Cloud that is impacting some of our Atlassian Cloud customers. We will provide more details within the next hour.
- `2019-05-10 14:37` **resolved** — Between 13:30 UTC to 14:05 UTC, we experienced a partial service outage in the US region for Jira Cloud products. The issue has been resolved and the service is operating normally.

### 2. atlassian — Jira Imports are Failing

- URL: http://stspg.io/06e0ed12a
- Started: 2019-05-07 09:45 UTC
- Resolved: 2019-05-07 09:45 UTC
- Duration: 0m
- Impact: major, status: resolved
- Components: —
- Canonical incident: [ci-9a8de6eabdeb3ec639bfe106](#ci-9a8de6eabdeb3ec639bfe106)

- `2019-05-07 12:56` **resolved** — Between 06/May/2019 06:00 AM UTC to 07/May/2019 09:10 AM UTC, we experienced issues with Jira imports. The issue has been resolved and the service is operating normally.

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-e659433cc9489712bacbf039"></a>

- Canonical ID: `ci-e659433cc9489712bacbf039` (1 source records).
  - [atlassian/dgsyrnqfr2r0](../../data/2019/W19/atlassian.json); vendor URL: http://stspg.io/cb7cd57ba
- Start time basis: `vendor_reported`.

<a id="ci-9a8de6eabdeb3ec639bfe106"></a>

- Canonical ID: `ci-9a8de6eabdeb3ec639bfe106` (1 source records).
  - [atlassian/mx1wvlvbw23y](../../data/2019/W19/atlassian.json); vendor URL: http://stspg.io/06e0ed12a
- Start time basis: `vendor_reported`.

## Sources

| Vendor | Status | Source records seen | Parse errors |
|--------|--------|----------------|--------------|
| anthropic | ok | 50 | 0 |
| atlassian | ok | 34 | 0 |
| aws | ok | 5 | 0 |
| circleci | ok | 50 | 0 |
| cloudflare | ok | 50 | 0 |
| datadog | ok | 50 | 0 |
| digitalocean | ok | 50 | 0 |
| gcp | ok | 4 | 0 |
| github | ok | 50 | 0 |
| grafana | ok | 50 | 0 |
| openai | ok | 25 | 0 |
| slack | ok | 43 | 0 |
| stripe | ok | 50 | 0 |
