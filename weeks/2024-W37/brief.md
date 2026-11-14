# Week 2024-W37 — top 2 of 2 above threshold

2024-09-09 to 2024-09-15 (UTC).

2 source records (2 saved revisions) → 2 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## 1. atlassian — High RDS CPU on multiple environments

- URL: https://stspg.io/6w6ly5cgydvz
- Window: 2024-09-09 08:56 to 2024-09-09 18:12 UTC (9h 16m)
- Impact: none, status: resolved
- Components: —

<a id="ci-fbd541b325a770d81652f278"></a>

- Canonical ID: `ci-fbd541b325a770d81652f278` (1 source records).
  - [atlassian/d8jljnw353f3](../../data/2024/W37/atlassian.json); vendor URL: https://stspg.io/6w6ly5cgydvz
- Start time basis: `vendor_reported`.

  - `2024-09-09 08:56` investigating — We are investigating issues with Latency and 5xx errors for Jira, Jira Service Management, Jira Work Management due to High RDS CPU on several EU-Central environment and will provide updates here soon.
  - `2024-09-09 10:56` monitoring — We’ve taken steps to address the issue by scaling up the RDS instances, and we will continue to monitor the systems to ensure everything remains stable.
  - `2024-09-09 16:01` identified — We are observing different RDS problems in different environments for Jira, Jira Service Management, and Jira Work Management. We are actively working on this and will provide more updates here soon.
  - `2024-09-09 16:19` investigating — We are observing different RDS problems in different environments for Jira, Jira Service Management, and Jira Work Management. We are actively working on this and will provide more updates here soon.
  - `2024-09-09 17:13` monitoring — We identified the root cause of the RDS problems causing issues for Jira, Jira Service Management, and Jira Work Management. We mitigated the issue and are now monitoring the fix.
  - `2024-09-09 18:12` resolved — Earlier today, we experienced RDS problems for Jira, Jira Service Management, and Jira Work Management. The issue has been resolved, and the services are operating normally.

## 2. atlassian — Slow loading times and extended response times for Cloud products in the EU-West region.

- URL: https://stspg.io/7kv4zlyf4fmn
- Window: 2024-09-10 10:26 to 2024-09-10 12:43 UTC (2h 17m)
- Impact: none, status: resolved
- Components: —

<a id="ci-8661a2eff693fd637888a479"></a>

- Canonical ID: `ci-8661a2eff693fd637888a479` (1 source records).
  - [atlassian/m0nz4gbbjmz2](../../data/2024/W37/atlassian.json); vendor URL: https://stspg.io/7kv4zlyf4fmn
- Start time basis: `vendor_reported`.

  - `2024-09-10 10:26` investigating — We are investigating issues with slow loading and long responsive time for all products (Jira, JSM, JSW and Confluence) and will provide updates here soon.
  - `2024-09-10 12:43` resolved — Systems have recovered, and the issue has been mitigated. All affected products are fully operational, with no further impact observed.
