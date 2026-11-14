# Week 2024-W23

2024-06-03 to 2024-06-09 (UTC). Generated 2026-11-14 00:00.

1 canonical incidents, 1 above threshold. All 13 sources responded.

1 source records (1 saved revisions) → 1 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | atlassian | [Some of the www.atlassian.com pages are inaccessible](https://stspg.io/qn6bhjbczv28) | 2024-06-06 01:30 | 2h 59m | none | — | 179 = 179 × 1 × 1.0 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. atlassian — Some of the www.atlassian.com pages are inaccessible

- URL: https://stspg.io/qn6bhjbczv28
- Started: 2024-06-06 01:30 UTC
- Resolved: 2024-06-06 04:30 UTC
- Duration: 2h 59m
- Impact: none, status: resolved
- Components: —
- Canonical incident: [ci-db11aedaba6782abfebec120](#ci-db11aedaba6782abfebec120)

- `2024-06-06 01:30` **investigating** — We are investigating reports of intermittent errors on some Atlassian, Opsgenie, Atlassian Access, and Compass-related www.atlassian.com web pages. Once we identify the root cause, we will provide more details.
- `2024-06-06 03:12` **identified** — We continue to work on resolving the issue with www.atlassian.com. We have identified the root cause and expect recovery shortly.
- `2024-06-06 04:30` **resolved** — Between 05 June 23:30 UTC and 06 June 04:34 UTC, we experienced issues with some of the pages on www.atlassian.com. The issue has been resolved, and the pages are operating normally.

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-db11aedaba6782abfebec120"></a>

- Canonical ID: `ci-db11aedaba6782abfebec120` (1 source records).
  - [atlassian/4dh7m57xqhfq](../../data/2024/W23/atlassian.json); vendor URL: https://stspg.io/qn6bhjbczv28
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
