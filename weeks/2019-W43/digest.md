# Week 2019-W43

2019-10-21 to 2019-10-27 (UTC). Generated 2026-11-14 00:00.

1 canonical incidents, 1 above threshold. All 13 sources responded.

1 source records (1 saved revisions) → 1 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | atlassian | [Intermittent DNS Resolution Errors](http://stspg.io/e1be96263) | 2019-10-22 22:32 | 5h 29m | none | — | 329 = 329 × 1 × 1.0 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. atlassian — Intermittent DNS Resolution Errors

- URL: http://stspg.io/e1be96263
- Started: 2019-10-22 22:32 UTC
- Resolved: 2019-10-23 04:02 UTC
- Duration: 5h 29m
- Impact: none, status: resolved
- Components: —
- Canonical incident: [ci-09947484c0995c6b58b2eed9](#ci-09947484c0995c6b58b2eed9)

- `2019-10-22 22:32` **investigating** — We are aware that some customers may be experiencing intermittent failure connecting to Atlassian Cloud. We are working with our cloud hosting provider (AWS) to resolve the DNS related errors. We will provide hourly updates.
- `2019-10-23 00:12` **identified** — We continue to work on resolving intermittent failure connecting to Atlassian Cloud for some cloud customers. We will provide hourly updates.
- `2019-10-23 01:08` **identified** — We continue to work on resolving intermittent failure connecting to Atlassian Cloud for some cloud customers. We will provide hourly updates.
- `2019-10-23 01:32` **identified** — We have identified the root cause of the intermittent failure connecting to Atlassian Cloud for some cloud customers and have mitigated the problem. We are now monitoring closely.
- `2019-10-23 01:32` **monitoring** — We have identified the root cause of the intermittent failure connecting to Atlassian Cloud for some cloud customers and have mitigated the problem. We are now monitoring closely.
- `2019-10-23 04:02` **resolved** — Between 23:45 UTC to 02:19 UTC, some customers experienced intermittent failure connecting to Atlassian Cloud. The root cause was an increased DNS error rate from our infrastructure supplier. The supplier fixed the upstream issue and we have verified that the services have recovered. The conditions that caused the issue have been addressed and we are actively working on a permanent fix. The issue has been resolved and the service is operating normally.

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-09947484c0995c6b58b2eed9"></a>

- Canonical ID: `ci-09947484c0995c6b58b2eed9` (1 source records).
  - [atlassian/kw28xr3wyxmq](../../data/2019/W43/atlassian.json); vendor URL: http://stspg.io/e1be96263
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
