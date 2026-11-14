# Week 2023-W43 — top 1 of 1 above threshold

2023-10-23 to 2023-10-29 (UTC).

1 source records (1 saved revisions) → 1 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## 1. atlassian — Email domain unverified

- URL: https://stspg.io/xlcrsntd8s3p
- Window: 2023-10-24 00:22 to 2023-10-24 10:00 UTC (9h 38m)
- Impact: major, status: resolved
- Components: —

<a id="ci-c6e5539135eab374811266ce"></a>

- Canonical ID: `ci-c6e5539135eab374811266ce` (1 source records).
  - [atlassian/3sf6vg9v8nqm](../../data/2023/W43/atlassian.json); vendor URL: https://stspg.io/xlcrsntd8s3p
- Start time basis: `vendor_reported`.

  - `2023-10-24 00:22` investigating — We are investigating an issue with emails that have been sent out to admins stating domain verification is failing. We will provide more details within the next hour.
  - `2023-10-24 01:49` investigating — We are continuing to investigate this issue. We are aware that this incident is impacting custom email configuration and some Jira notifications are being sent with the default email address instead of the configured custom email. We will provide more details in the next hour.
  - `2023-10-24 02:43` investigating — The investigation is still ongoing. We are working on a mitigation.
  - `2023-10-24 05:04` identified — We continue to work on resolving the issue with domain verification. We have identified the root cause and expect recovery shortly.
  - `2023-10-24 06:14` monitoring — We have mitigated the problem. We are now monitoring this closely. Please contact us at https://support.atlassian.com/contact/ if you are still facing issues with email domain verification.
  - `2023-10-24 06:48` identified — We identified that the mitigation method doesn't fix the issue in some affected Cloud instances. We are continuing with the investigation.
  - `2023-10-24 07:49` identified — We have identified the issue and a fix is currently being implemented. We expect this issue to be resolved shortly.
  - `2023-10-24 10:00` resolved — Between 22:00 UTC 23rd October and 07:45 UTC 24th October, we experienced an issue with custom domain validation, this resulted in Jira Software, Jira Work Management, Jira Service management, Atlassian Access notification emails for some customers to be sent from generic Atlassian email domains rather than the custom domain email address. The issue has been resolved and the service is operating normally.
