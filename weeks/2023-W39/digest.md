# Week 2023-W39

2023-09-25 to 2023-10-01 (UTC). Generated 2026-11-14 00:00.

1 canonical incidents, 1 above threshold. All 13 sources responded.

1 source records (1 saved revisions) → 1 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | atlassian | [We are experiencing intermittent issues on Atlassian.com and Bitbucket Marketing pages](https://stspg.io/7nk14plf0xcd) | 2023-09-29 16:26 | 4h 0m | none | — | 240 = 240 × 1 × 1.0 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. atlassian — We are experiencing intermittent issues on Atlassian.com and Bitbucket Marketing pages

- URL: https://stspg.io/7nk14plf0xcd
- Started: 2023-09-29 16:26 UTC
- Resolved: 2023-09-29 20:27 UTC
- Duration: 4h 0m
- Impact: none, status: resolved
- Components: —
- Canonical incident: [ci-9715c5cb1816e6544c716b7a](#ci-9715c5cb1816e6544c716b7a)

- `2023-09-29 16:26` **investigating** — We are investigating cases of degraded performance for Atlassian Cloud customers. We will provide more details within the next hour.
- `2023-09-29 17:44` **investigating** — We are still investigating intermittent issues with these two domains. We will provide more details within the next hour.
- `2023-09-29 18:42` **investigating** — We are still investigating intermittent issues with these two domains. We will provide more details within the next hour.
- `2023-09-29 20:27` **resolved** — Between 04:36 UTC to 17:50 UTC, we experienced intermittent issues on Atlassian.com and Bitbucket Marketing pages. The issue has been resolved and the service is operating normally.

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-9715c5cb1816e6544c716b7a"></a>

- Canonical ID: `ci-9715c5cb1816e6544c716b7a` (1 source records).
  - [atlassian/lgl84rzhxd0d](../../data/2023/W39/atlassian.json); vendor URL: https://stspg.io/7nk14plf0xcd
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
