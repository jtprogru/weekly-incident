# Week 2023-W27

2023-07-03 to 2023-07-09 (UTC). Generated 2026-11-14 00:00.

1 canonical incidents, 1 above threshold. All 13 sources responded.

1 source records (1 saved revisions) → 1 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | atlassian | [Performance issues and outages with Cloud products](https://stspg.io/y73s93jsl9d7) | 2023-07-06 11:18 | 4h 18m | none | — | 258 = 258 × 1 × 1.0 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. atlassian — Performance issues and outages with Cloud products

- URL: https://stspg.io/y73s93jsl9d7
- Started: 2023-07-06 11:18 UTC
- Resolved: 2023-07-06 15:37 UTC
- Duration: 4h 18m
- Impact: none, status: postmortem
- Components: —
- Canonical incident: [ci-bfd188d0626de2182d0a7428](#ci-bfd188d0626de2182d0a7428)

- `2023-07-06 11:18` **identified** — We are investigating an issue with an internal infrastructure component that is impacting multiple Cloud products, including Jira Software, Bitbucket, Jira Service Management and Confluence, and customers. These issues include performance impact and, in some cases, outages. Users may experience slow loading and uploading of attachments, login issues or inability for new customers to sign up. We have identified the root cause and are actively working on the service recovery.
- `2023-07-06 13:17` **monitoring** — We have identified the root cause of an issue with an internal infrastructure component that has been impacting multiple Cloud products, including Jira Software, Jira Service Management and Confluence, and customers. This issue had lead to a performance impact and, in some cases, outages. We have implemented a fix to resolve the issue and recovery is in progress.
- `2023-07-06 15:37` **resolved** — We experienced performance issues and outages for several Atlassian Cloud Products. The issue has been resolved and the service is operating normally.
- `2023-07-14 04:56` **postmortem** — ### **SUMMARY** We understand the importance of providing reliable and consistent service to our valued customers. On July 6, 2023, from 03:52 to 15:11 UTC, we experienced an issue with an upgraded version of a third-party tool that functions as our internal artifact management system. Despite our monitoring system identifying the incident within two minutes, this issue led to the degradation of the scaling capabilities of our internal hosting platform, resulting in service degradation or outages for customers of Atlassian cloud. In response to this situation, we are taking immediate measures to enhance the stability of our system and prevent similar issues from re-occurring. ### **IMPACT** This incident affected multiple regions and products due to the diminished scaling capabilities of our internal hosting platform. In most products and offerings, customers faced reduced functionality, slower response times, and limited access to specific features. ### **ROOT CAUSE** The root cause of the incident was the introduction of new functionality in a third-party tool that functions as our internal artifact management system. It led to an unexpected increase in the load on the primary database of the artifact system. Upon identifying and localizing the problem, we promptly adjusted the system configuration to regain stability. ### **REMEDIAL ACTIONS PLAN & NEXT STEPS** Over the next months, we will enact a temporary freeze on non-critical upgrades of the artifact management system, and we will focus our efforts on three high-priority initiatives: 1. **Enhancing system scaling:** We prioritized work ensuring that downtime in a critical infrastructure component does not affect the scaling of other components. We expect to complete this initiative within the next two months. 2. **Reducing interdependencies:** We are working to mitigate the risk of potential cascading failures by ensuring that significant system components are able to operate independently in the case of issues. Initiatives 1 and 2 are already in progress but have been given priority to be completed as soon as possible. 3. **Strengthening testing procedures:** Alongside these initiatives, we are addressing the need for even more stringent testing procedures than we already have in place to prevent potential issues in future updates. We are committed to collaborating closely with our technology partners to ensure the most optimal experience for our customers. We apologize for any inconvenience caused by this incident and appreciate your understanding. Our team is dedicated to continually improving our systems and processes to provide you with the exceptional service you deserve. Thank you for your continued support and trust in us. Sincerely, Atlassian Customer Support

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-bfd188d0626de2182d0a7428"></a>

- Canonical ID: `ci-bfd188d0626de2182d0a7428` (1 source records).
  - [atlassian/pzh5b8cvfk2g](../../data/2023/W27/atlassian.json); vendor URL: https://stspg.io/y73s93jsl9d7
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
