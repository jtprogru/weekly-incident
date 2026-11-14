# Week 2023-W27 — top 1 of 1 above threshold

2023-07-03 to 2023-07-09 (UTC).

1 source records (1 saved revisions) → 1 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## 1. atlassian — Performance issues and outages with Cloud products

- URL: https://stspg.io/y73s93jsl9d7
- Window: 2023-07-06 11:18 to 2023-07-06 15:37 UTC (4h 18m)
- Impact: none, status: postmortem
- Components: —

<a id="ci-bfd188d0626de2182d0a7428"></a>

- Canonical ID: `ci-bfd188d0626de2182d0a7428` (1 source records).
  - [atlassian/pzh5b8cvfk2g](../../data/2023/W27/atlassian.json); vendor URL: https://stspg.io/y73s93jsl9d7
- Start time basis: `vendor_reported`.

  - `2023-07-06 11:18` identified — We are investigating an issue with an internal infrastructure component that is impacting multiple Cloud products, including Jira Software, Bitbucket, Jira Service Management and Confluence, and customers. These issues include performance impact and, in some cases, outages. Users may experience slow loading and uploading of attachments, login issues or inability for new customers to sign up. We have identified the root cause and are actively working on the service recovery.
  - `2023-07-06 13:17` monitoring — We have identified the root cause of an issue with an internal infrastructure component that has been impacting multiple Cloud products, including Jira Software, Jira Service Management and Confluence, and customers. This issue had lead to a performance impact and, in some cases, outages. We have implemented a fix to resolve the issue and recovery is in progress.
  - `2023-07-06 15:37` resolved — We experienced performance issues and outages for several Atlassian Cloud Products. The issue has been resolved and the service is operating normally.
  - `2023-07-14 04:56` postmortem — ### **SUMMARY** We understand the importance of providing reliable and consistent service to our valued customers. On July 6, 2023, from 03:52 to 15:11 UTC, we experienced an issue with an upgraded version of a third-party tool that functions as our internal artifact management system. Despite our monitoring system identifying the incident within two minutes, this issue led to the degradation of the scaling capabilities of our internal hosting platform, resulting in service degradation or outages for customers of Atlassian cloud. In response to this situation, we are taking immediate measures to enhance the stability of our system and prevent similar issues from re-occurring. ### **IMPACT** This incident affected multiple regions and products due to the diminished scaling capabilities…
