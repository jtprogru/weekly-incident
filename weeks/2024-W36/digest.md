# Week 2024-W36

2024-09-02 to 2024-09-08 (UTC). Generated 2026-11-14 00:00.

1 canonical incidents, 1 above threshold. All 13 sources responded.

1 source records (1 saved revisions) → 1 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | atlassian | [Slowness in Jira.](https://stspg.io/gnd70qwf3n98) | 2024-09-02 09:23 | 8d 3h 51m | none | — | 11751 = 11751 × 1 × 1.0 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. atlassian — Slowness in Jira.

- URL: https://stspg.io/gnd70qwf3n98
- Started: 2024-09-02 09:23 UTC
- Resolved: 2024-09-10 13:15 UTC
- Duration: 8d 3h 51m
- Impact: none, status: resolved
- Components: —
- Canonical incident: [ci-255a889fc6d75ef90b67d149](#ci-255a889fc6d75ef90b67d149)

- `2024-09-02 09:23` **investigating** — Our team is currently investigating the issue. We are working to identify the cause and will provide updates as soon as possible.
- `2024-09-02 10:56` **monitoring** — The slowness issue has been resolved, and we are actively monitoring the systems to ensure stability.
- `2024-09-02 18:33` **monitoring** — We have identified the root cause of an issue with an internal infrastructure component that has been impacting the provisioning process for the Mumbai Region. This issue has led to a performance impact or slowness in product provisioning. We are working on a fix to resolve the issue and recovery is in progress.
- `2024-09-10 13:15` **resolved** — We have identified the root cause of the issue with an internal infrastructure component that has been impacting the provisioning process for the Mumbai Region with several Atlassian Cloud products. All affected products are now back online and no further impact has been observed.

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-255a889fc6d75ef90b67d149"></a>

- Canonical ID: `ci-255a889fc6d75ef90b67d149` (1 source records).
  - [atlassian/kl4d5mfn26kc](../../data/2024/W36/atlassian.json); vendor URL: https://stspg.io/gnd70qwf3n98
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
