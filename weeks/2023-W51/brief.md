# Week 2023-W51 — top 1 of 1 above threshold

2023-12-18 to 2023-12-24 (UTC).

1 source records (1 saved revisions) → 1 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## 1. atlassian — Atlassian's cross product user search service is currently degraded.

- URL: https://stspg.io/3yxq4k25m6sr
- Window: 2023-12-18 15:04 to 2023-12-19 03:28 UTC (12h 23m)
- Impact: none, status: postmortem
- Components: —

<a id="ci-62c7ef5a72a10bb1128da827"></a>

- Canonical ID: `ci-62c7ef5a72a10bb1128da827` (1 source records).
  - [atlassian/lvq7pq0zbgvr](../../data/2023/W51/atlassian.json); vendor URL: https://stspg.io/3yxq4k25m6sr
- Start time basis: `vendor_reported`.

  - `2023-12-18 15:04` investigating — We are investigating reports of intermittent errors for <SOME/ALL> Atlassian, Confluence, Jira Work Management, Jira Service Management, Jira Software, Atlassian Bitbucket, Jira Align, Jira Product Discovery, Atlas, and Compass Cloud customers. We will provide more details once we identify the root cause.
  - `2023-12-18 16:01` investigating — Atlassian's cross product user search service is recovering. Searches for users within Atlassian products are returning to normal.
  - `2023-12-18 16:02` investigating — Atlassian's cross product user search service is recovering. Searches for users within Atlassian products are returning to normal.
  - `2023-12-18 16:06` investigating — Atlassian's cross product user search service is recovering. Searches for users within Atlassian products are returning to normal.
  - `2023-12-18 16:23` investigating — Atlassian's cross product user search service is currently healthy. Searches for users within Atlassian products are working as expected. We are in the process of investigating the root cause of this incident.
  - `2023-12-19 03:28` resolved — It has been resolved. Atlassian's cross product user search is working.
  - `2023-12-28 02:43` postmortem — ### **SUMMARY** On Dec 18, 2023, between 12:29 p.m. and 3:35 p.m. UTC, Dec 18, 2023, Atlassian's cloud customers using Atlas, Bitbucket Cloud, Compass, Confluence Cloud, Jira Service Management, Jira Software, Jira Work Management, Jira Product Discovery products were unable to search for users or use the "@mention" functionality. Customers' user search results failed or were delayed as Atlassian's service returning user search results was degraded in several regions. The incident originated from a computationally intensive operation that was triggered multiple times in rapid succession, resulting in degraded performance of Atlassian's user search service across several regions. Notably, customers in the EU west region were most affected. The incident was detected within 2 minutes by aut…
