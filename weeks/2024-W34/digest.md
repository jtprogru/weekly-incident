# Week 2024-W34

2024-08-19 to 2024-08-25 (UTC). Generated 2026-11-14 00:00.

1 canonical incidents, 1 above threshold. All 13 sources responded.

1 source records (1 saved revisions) → 1 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | atlassian | [Data residency migrations halted.](https://stspg.io/pd9n7p9p3vb4) | 2024-08-20 11:46 | 2d 23h 28m | minor | — | 4288 = 4288 × 1 × 1.0 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. atlassian — Data residency migrations halted.

- URL: https://stspg.io/pd9n7p9p3vb4
- Started: 2024-08-20 11:46 UTC
- Resolved: 2024-08-23 11:15 UTC
- Duration: 2d 23h 28m
- Impact: minor, status: resolved
- Components: —
- Canonical incident: [ci-332dd047fcb01d885f9aca36](#ci-332dd047fcb01d885f9aca36)

- `2024-08-20 11:46` **investigating** — Data residency migrations are currently not possible for Cloud customers. If you attempt a Data residency migration, you will encounter an error message, “We are unable to schedule your move” when you get to the “Review and submit your request” page. An update will be shared here as soon as this service is resumed.
- `2024-08-23 11:15` **resolved** — The issue with scheduling App data residency migrations has been resolved and the service is operating normally.

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-332dd047fcb01d885f9aca36"></a>

- Canonical ID: `ci-332dd047fcb01d885f9aca36` (1 source records).
  - [atlassian/2rsqj1hmdg82](../../data/2024/W34/atlassian.json); vendor URL: https://stspg.io/pd9n7p9p3vb4
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
