# Week 2019-W20

2019-05-13 to 2019-05-19 (UTC). Generated 2026-11-14 00:00.

1 canonical incidents, 1 above threshold. All 13 sources responded.

1 source records (1 saved revisions) → 1 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | atlassian | [Atlassian Cloud users experiencing outages](http://stspg.io/9c0c8c681) | 2019-05-14 08:57 | 2h 6m | critical | — | 126 = 126 × 1 × 1.0 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. atlassian — Atlassian Cloud users experiencing outages

- URL: http://stspg.io/9c0c8c681
- Started: 2019-05-14 08:57 UTC
- Resolved: 2019-05-14 11:04 UTC
- Duration: 2h 6m
- Impact: critical, status: resolved
- Components: —
- Canonical incident: [ci-7c8af48c1b5ff734fc2d2dbb](#ci-7c8af48c1b5ff734fc2d2dbb)

- `2019-05-14 08:57` **investigating** — We are investigating an outage that is impacting some Atlassian Cloud customers. We will provide more details within the next hour.
- `2019-05-14 09:28` **investigating** — We are continuing to investigate this issue.
- `2019-05-14 11:04` **resolved** — We have identified the service that caused a high load on the server. We have deployed a couple of changes to mitigate the situation. The services have been recovered and operating normally.

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-7c8af48c1b5ff734fc2d2dbb"></a>

- Canonical ID: `ci-7c8af48c1b5ff734fc2d2dbb` (1 source records).
  - [atlassian/3p4zfxggg0vv](../../data/2019/W20/atlassian.json); vendor URL: http://stspg.io/9c0c8c681
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
