# Week 2019-W27

2019-07-01 to 2019-07-07 (UTC). Generated 2026-11-14 00:00.

1 canonical incidents, 1 above threshold. All 13 sources responded.

1 source records (1 saved revisions) → 1 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | atlassian | [App switcher not loading in Jira and Confluence for some customers](http://stspg.io/d2bfd779a) | 2019-07-01 09:03 | 2h 0m | major | — | 120 = 120 × 1 × 1.0 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. atlassian — App switcher not loading in Jira and Confluence for some customers

- URL: http://stspg.io/d2bfd779a
- Started: 2019-07-01 09:03 UTC
- Resolved: 2019-07-01 11:03 UTC
- Duration: 2h 0m
- Impact: major, status: resolved
- Components: —
- Canonical incident: [ci-ff81391fb676709acd02670b](#ci-ff81391fb676709acd02670b)

- `2019-07-01 09:03` **investigating** — We are investigating an issue with app switcher that is impacting some Atlassian Cloud customers. We will provide more details within the next hour.
- `2019-07-01 10:04` **investigating** — We are continuing to investigate this issue.
- `2019-07-01 11:03` **resolved** — The issue has been resolved and the service is operating normally.

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-ff81391fb676709acd02670b"></a>

- Canonical ID: `ci-ff81391fb676709acd02670b` (1 source records).
  - [atlassian/w81fyrpz5qcl](../../data/2019/W27/atlassian.json); vendor URL: http://stspg.io/d2bfd779a
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
