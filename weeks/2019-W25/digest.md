# Week 2019-W25

2019-06-17 to 2019-06-23 (UTC). Generated 2026-11-14 00:00.

1 canonical incidents, 1 above threshold. All 13 sources responded.

1 source records (1 saved revisions) → 1 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | atlassian | [Atlassian Cloud email notifications issue](http://stspg.io/d5b25c71b) | 2019-06-19 12:00 | 1d 0h 38m | major | — | 1478 = 1478 × 1 × 1.0 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. atlassian — Atlassian Cloud email notifications issue

- URL: http://stspg.io/d5b25c71b
- Started: 2019-06-19 12:00 UTC
- Resolved: 2019-06-20 12:39 UTC
- Duration: 1d 0h 38m
- Impact: major, status: resolved
- Components: —
- Canonical incident: [ci-bb9645b89c2bf851a0ba98c9](#ci-bb9645b89c2bf851a0ba98c9)

- `2019-06-19 12:00` **investigating** — We are investigating an issue with outgoing email that is impacting some Atlassian Cloud customers. We will provide more details within the next hour.
- `2019-06-19 13:27` **identified** — We continue to work on resolving the outgoing emails for Atlassian Cloud. We have identified the root cause.
- `2019-06-20 12:39` **resolved** — This incident has been resolved.

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-bb9645b89c2bf851a0ba98c9"></a>

- Canonical ID: `ci-bb9645b89c2bf851a0ba98c9` (1 source records).
  - [atlassian/sqq15l5bd34s](../../data/2019/W25/atlassian.json); vendor URL: http://stspg.io/d5b25c71b
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
