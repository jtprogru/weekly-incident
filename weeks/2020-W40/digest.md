# Week 2020-W40

2020-09-28 to 2020-10-04 (UTC). Generated 2026-11-14 00:00.

1 canonical incidents, 1 above threshold. All 13 sources responded.

1 source records (1 saved revisions) → 1 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | atlassian | [Login / Signup with Microsoft not working](http://stspg.io/n55q9mvx3w62) | 2020-09-28 23:39 | 1h 23m | major | — | 83 = 83 × 1 × 1.0 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. atlassian — Login / Signup with Microsoft not working

- URL: http://stspg.io/n55q9mvx3w62
- Started: 2020-09-28 23:39 UTC
- Resolved: 2020-09-29 01:03 UTC
- Duration: 1h 23m
- Impact: major, status: resolved
- Components: —
- Canonical incident: [ci-18b291c935b96c9193e296fb](#ci-18b291c935b96c9193e296fb)

- `2020-09-28 23:39` **identified** — We are investigating reports of intermittent errors for some Atlassian, Confluence, Jira Core, Jira Service Desk, Jira Software, Opsgenie, Atlassian Developer, Trello, and Atlassian Bitbucket Cloud customers. We will provide more details once we identify the root cause.
- `2020-09-29 01:03` **resolved** — Between 2020-09-28 21:32 UTC to 2020-09-29 00:09 UTC, we experienced a degradation in functionality for users attempting to authenticate using Microsoft. This affected Confluence, Jira Core, Jira Service Desk, Jira Software, Opsgenie, Atlassian Developer, Trello, and Atlassian Bitbucket. The issue has been resolved and the service is operating normally.

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-18b291c935b96c9193e296fb"></a>

- Canonical ID: `ci-18b291c935b96c9193e296fb` (1 source records).
  - [atlassian/wkg98r6wc1b5](../../data/2020/W40/atlassian.json); vendor URL: http://stspg.io/n55q9mvx3w62
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
