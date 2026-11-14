# Week 2024-W02

2024-01-08 to 2024-01-14 (UTC). Generated 2026-11-14 00:00.

1 canonical incidents, 1 above threshold. All 13 sources responded.

1 source records (1 saved revisions) → 1 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | atlassian | [Delayed SCIM provisioning syncs of users and groups from Identity Providers](https://stspg.io/trcvys1bxjkn) | 2024-01-09 22:03 | 2d 19h 9m | none | — | 4029 = 4029 × 1 × 1.0 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. atlassian — Delayed SCIM provisioning syncs of users and groups from Identity Providers

- URL: https://stspg.io/trcvys1bxjkn
- Started: 2024-01-09 22:03 UTC
- Resolved: 2024-01-12 17:13 UTC
- Duration: 2d 19h 9m
- Impact: none, status: resolved
- Components: —
- Canonical incident: [ci-018a773ce5674989761759d9](#ci-018a773ce5674989761759d9)

- `2024-01-09 22:03` **investigating** — We are investigating cases of degraded performance when SCIM provisioning users/groups for Confluence, Jira Work Management, Jira Service Management, Jira Software, and Atlassian Access Cloud customers. We will provide more details shortly.
- `2024-01-10 03:17` **monitoring** — We have identified a bottleneck in the Group synchronisation process for SCIM Provisioning. We have increased the resources allocated to the process in order to mitigate the issue. We are now monitoring the service.
- `2024-01-10 19:56` **monitoring** — The bottleneck in the Group synchronization process for SCIM Provisioning is still under work. We are closely monitoring the service.
- `2024-01-11 18:54` **monitoring** — The bottleneck in the Group synchronization process for SCIM Provisioning is still under work. The team is all hands on deck to improve the situation.
- `2024-01-11 22:36` **monitoring** — The bottleneck in the Group synchronization process for SCIM Provisioning has seen considerable improvement due to recent changes, but it is still under work. The team is all hands on deck to keep improving the situation.
- `2024-01-12 06:55` **monitoring** — A fix for the bottleneck identified in the Group synchronization process for SCIM Provisioning has been made. We are seeing processing start to return to normal levels and will be monitoring over the next few hours. The team is all hands on deck to keep improving the situation.
- `2024-01-12 17:13` **resolved** — We experienced degraded SCIM provisioning from external Identity Providers for Confluence, Jira Work Management, Jira Service Management, Jira Software, and Atlassian Access. The issue has been resolved and the service is operating normally.

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-018a773ce5674989761759d9"></a>

- Canonical ID: `ci-018a773ce5674989761759d9` (1 source records).
  - [atlassian/71sdcbz9gy53](../../data/2024/W02/atlassian.json); vendor URL: https://stspg.io/trcvys1bxjkn
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
