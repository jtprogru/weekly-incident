# Week 2021-W16

2021-04-19 to 2021-04-25 (UTC). Generated 2026-11-14 00:00.

1 canonical incidents, 1 above threshold. All 13 sources responded.

1 source records (1 saved revisions) → 1 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | atlassian | [Atlassian account session invalidations](https://stspg.io/ql9jsw73ndfh) | 2021-04-19 04:32 | 10d 21h 1m | none | — | 15661 = 15661 × 1 × 1.0 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. atlassian — Atlassian account session invalidations

- URL: https://stspg.io/ql9jsw73ndfh
- Started: 2021-04-19 04:32 UTC
- Resolved: 2021-04-30 01:34 UTC
- Duration: 10d 21h 1m
- Impact: none, status: resolved
- Components: —
- Canonical incident: [ci-097af74616e841c66ac3232c](#ci-097af74616e841c66ac3232c)

- `2021-04-19 04:32` **investigating** — We are investigating an incident impacting Jira Cloud, Confluence Cloud, Bitbucket Cloud, and Statuspage. During our investigation, users may be logged out of their accounts as we work towards a resolution. We are continuing to investigate and will update this incident with more details as they are available.
- `2021-04-30 01:34` **resolved** — Between 15/Apr/21 17:20 AM PDT to 27/Apr/21 10:00 PM PDT, we experienced some cloud customers of Atlassian Support, Confluence, Jira Work Management, Jira Service Management, Jira Software, Opsgenie, Atlassian Developer, Trello, Atlassian Bitbucket, Atlassian Access, and Jira Align were logged out of there account. The issue has been resolved and the service is operating normally.

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-097af74616e841c66ac3232c"></a>

- Canonical ID: `ci-097af74616e841c66ac3232c` (1 source records).
  - [atlassian/yz4mkc6j501q](../../data/2021/W16/atlassian.json); vendor URL: https://stspg.io/ql9jsw73ndfh
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
