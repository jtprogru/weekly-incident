# Week 2026-W08 — top 1 of 1 above threshold

2026-02-16 to 2026-02-22 (UTC).

4 source records (4 saved revisions) → 4 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## 1. circleci — Sign-In Issue Affecting Email/Password Logins

- URL: https://stspg.io/8hdkcgmsfw21
- Window: 2026-02-22 17:25 to 2026-02-22 19:05 UTC (1h 40m)
- Impact: major, status: resolved
- Components: —

<a id="ci-22d65fb9b2b12ef254dd38af"></a>

- Canonical ID: `ci-22d65fb9b2b12ef254dd38af` (1 source records).
  - [circleci/5xrywvlmqd9b](../../data/2026/W08/circleci.json); vendor URL: https://stspg.io/8hdkcgmsfw21
- Start time basis: `vendor_reported`.

  - `2026-02-22 17:25` investigating — We are currently investigating an issue affecting users attempting to log in with email and password credentials. Sign-in via GitHub or Bitbucket OAuth is not affected and remains fully operational. This issue appears to be related to an outage with a downstream authentication provider: https://status.auth0.com/
  - `2026-02-22 17:37` identified — We’ve identified that the issue affecting email/password sign-ins is caused by a degradation in service from our downstream authentication provider, Auth0: https://status.auth0.com/incidents/kknl0nbdzbvx Users should still be able to sign in using GitHub or Bitbucket OAuth, which remain operational. We continue to monitor and will update as we learn more.
  - `2026-02-22 18:53` monitoring — Our downstream authentication provider has implemented mitigation actions and is reporting a reduction in errors. We are seeing corresponding improvement in email/password login success rates. We are continuing to monitor closely to ensure stability before marking this incident resolved.
  - `2026-02-22 19:05` resolved — Our authentication provider has confirmed the earlier service disruption has been fully mitigated: https://status.auth0.com/incidents/kknl0nbdzbvx Email/password login functionality has been restored and is operating normally.
