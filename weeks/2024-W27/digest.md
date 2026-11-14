# Week 2024-W27

2024-07-01 to 2024-07-07 (UTC). Generated 2026-11-14 00:00.

1 canonical incidents, 1 above threshold. All 13 sources responded.

1 source records (1 saved revisions) → 1 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | atlassian | [Some products are hard down](https://stspg.io/7h7xnnb9dcq5) | 2024-07-03 20:51 | 7h 3m | none | — | 423 = 423 × 1 × 1.0 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. atlassian — Some products are hard down

- URL: https://stspg.io/7h7xnnb9dcq5
- Started: 2024-07-03 20:51 UTC
- Resolved: 2024-07-04 03:54 UTC
- Duration: 7h 3m
- Impact: none, status: resolved
- Components: —
- Canonical incident: [ci-4be32123f53b880b66778710](#ci-4be32123f53b880b66778710)

- `2024-07-03 20:51` **investigating** — We are investigating an issue with <FUNCTIONALITY IMPACTED> that is impacting <SOME/ALL> Atlassian, Atlassian Partners, Atlassian Support, Confluence, Jira Work Management, Jira Service Management, Jira, Opsgenie, Atlassian Developer, Atlassian (deprecated), Trello, Atlassian Bitbucket, Guard, Jira Align, Jira Product Discovery, Atlas, Atlassian Analytics, and Rovo Cloud customers. We will provide more details within the next hour.
- `2024-07-03 21:45` **monitoring** — We have mitigated the problem and continue looking into the root cause. The outage was between 8:08pm 03/07 UTC - 08:31pm 03/07 UTC We are now monitoring closely.
- `2024-07-04 03:54` **resolved** — Between 03-07-2024 20:08 UTC to 03-07-2024 20:31 UTC, we experienced downtime for some of the products . JIRA product impacted with issue creation and project creation. While Confluence was impacted with public link access by anonymous user and the mission control page. Confluence issue was fixed at 04-07-2024 02:30 UTC. The issue has been resolved and the service is operating normally.

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-4be32123f53b880b66778710"></a>

- Canonical ID: `ci-4be32123f53b880b66778710` (1 source records).
  - [atlassian/2spys43jxslj](../../data/2024/W27/atlassian.json); vendor URL: https://stspg.io/7h7xnnb9dcq5
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
