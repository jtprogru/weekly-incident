# Week 2023-W51

2023-12-18 to 2023-12-24 (UTC). Generated 2026-11-14 00:00.

1 canonical incidents, 1 above threshold. All 13 sources responded.

1 source records (1 saved revisions) → 1 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | atlassian | [Atlassian's cross product user search service is currently degraded.](https://stspg.io/3yxq4k25m6sr) | 2023-12-18 15:04 | 12h 23m | none | — | 743 = 743 × 1 × 1.0 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. atlassian — Atlassian's cross product user search service is currently degraded.

- URL: https://stspg.io/3yxq4k25m6sr
- Started: 2023-12-18 15:04 UTC
- Resolved: 2023-12-19 03:28 UTC
- Duration: 12h 23m
- Impact: none, status: postmortem
- Components: —
- Canonical incident: [ci-62c7ef5a72a10bb1128da827](#ci-62c7ef5a72a10bb1128da827)

- `2023-12-18 15:04` **investigating** — We are investigating reports of intermittent errors for <SOME/ALL> Atlassian, Confluence, Jira Work Management, Jira Service Management, Jira Software, Atlassian Bitbucket, Jira Align, Jira Product Discovery, Atlas, and Compass Cloud customers. We will provide more details once we identify the root cause.
- `2023-12-18 16:01` **investigating** — Atlassian's cross product user search service is recovering. Searches for users within Atlassian products are returning to normal.
- `2023-12-18 16:02` **investigating** — Atlassian's cross product user search service is recovering. Searches for users within Atlassian products are returning to normal.
- `2023-12-18 16:06` **investigating** — Atlassian's cross product user search service is recovering. Searches for users within Atlassian products are returning to normal.
- `2023-12-18 16:23` **investigating** — Atlassian's cross product user search service is currently healthy. Searches for users within Atlassian products are working as expected. We are in the process of investigating the root cause of this incident.
- `2023-12-19 03:28` **resolved** — It has been resolved. Atlassian's cross product user search is working.
- `2023-12-28 02:43` **postmortem** — ### **SUMMARY** On Dec 18, 2023, between 12:29 p.m. and 3:35 p.m. UTC, Dec 18, 2023, Atlassian's cloud customers using Atlas, Bitbucket Cloud, Compass, Confluence Cloud, Jira Service Management, Jira Software, Jira Work Management, Jira Product Discovery products were unable to search for users or use the "@mention" functionality. Customers' user search results failed or were delayed as Atlassian's service returning user search results was degraded in several regions. The incident originated from a computationally intensive operation that was triggered multiple times in rapid succession, resulting in degraded performance of Atlassian's user search service across several regions. Notably, customers in the EU west region were most affected. The incident was detected within 2 minutes by automated monitoring, and our team promptly took action by recovering unhealthy systems and scaling up the service's infrastructure temporarily. The resolution process concluded in 3 hours and 06 minutes. ### **IMPACT** The overall impact was between Dec 18, 2023, between 12:29 p.m. UTC and Dec 18, 2023, 3:35 p.m. UTC. The Incident caused service disruption to cloud customers worldwide. Customers experienced delayed or failed user searches when using the following Atlassian cloud products: * Atlas * Bitbucket Cloud * Compass * Confluence Cloud * Jira Service Management * Jira Software * Jira Work Management * Jira Product Discovery ### **ROOT CAUSE** The incident stemmed from Atlassian's user search service receiving commands to process multiple computationally intensive operations in rapid succession. These operations were directed at the same customer data set, and therefore overloaded resources within a clustered database system, leading to memory exhaustion and subsequent unresponsiveness to user search requests. ### **REMEDIAL ACTIONS PLAN & NEXT STEPS** To prevent a recurrence of such incidents, we are implementing the following measures: * Implement a mechanism to queue computationally intensive operations in order to avoid overloading the resources within the systems and process them without impact on customer experience. * Fine-tune our clustered database settings to mitigate the impact of resource exhaustion on the overall system. We apologize to customers whose services were affected during this incident; we are taking immediate steps to improve the service’s resiliency. Thanks, Atlassian Customer Support

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-62c7ef5a72a10bb1128da827"></a>

- Canonical ID: `ci-62c7ef5a72a10bb1128da827` (1 source records).
  - [atlassian/lvq7pq0zbgvr](../../data/2023/W51/atlassian.json); vendor URL: https://stspg.io/3yxq4k25m6sr
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
