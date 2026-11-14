# Week 2026-W08

2026-02-16 to 2026-02-22 (UTC). Generated 2026-11-14 00:00.

4 canonical incidents, 1 above threshold. All 14 sources responded.

4 source records (4 saved revisions) → 4 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | circleci | [Sign-In Issue Affecting Email/Password Logins](https://stspg.io/8hdkcgmsfw21) | 2026-02-22 17:25 | 1h 40m | major | — | 100 = 100 × 1 × 1.0 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. circleci — Sign-In Issue Affecting Email/Password Logins

- URL: https://stspg.io/8hdkcgmsfw21
- Started: 2026-02-22 17:25 UTC
- Resolved: 2026-02-22 19:05 UTC
- Duration: 1h 40m
- Impact: major, status: resolved
- Components: —
- Canonical incident: [ci-22d65fb9b2b12ef254dd38af](#ci-22d65fb9b2b12ef254dd38af)

- `2026-02-22 17:25` **investigating** — We are currently investigating an issue affecting users attempting to log in with email and password credentials. Sign-in via GitHub or Bitbucket OAuth is not affected and remains fully operational. This issue appears to be related to an outage with a downstream authentication provider: https://status.auth0.com/
- `2026-02-22 17:37` **identified** — We’ve identified that the issue affecting email/password sign-ins is caused by a degradation in service from our downstream authentication provider, Auth0: https://status.auth0.com/incidents/kknl0nbdzbvx Users should still be able to sign in using GitHub or Bitbucket OAuth, which remain operational. We continue to monitor and will update as we learn more.
- `2026-02-22 18:53` **monitoring** — Our downstream authentication provider has implemented mitigation actions and is reporting a reduction in errors. We are seeing corresponding improvement in email/password login success rates. We are continuing to monitor closely to ensure stability before marking this incident resolved.
- `2026-02-22 19:05` **resolved** — Our authentication provider has confirmed the earlier service disruption has been fully mitigated: https://status.auth0.com/incidents/kknl0nbdzbvx Email/password login functionality has been restored and is operating normally.

## Below threshold (3)

- `datadog` [Monitors - Delayed Evaluation for Multiple Products](https://stspg.io/c6d05hsplhbv) — 1h 6m
- `circleci` [Pipeline UI Degradation and GitHub Checks Delivery](https://stspg.io/hczpgwq7t5cq) — 1h 18m
- `stripe` [Service disruption with a French issuing bank](https://stspg.io/j8ynh35wfv35) — 1h 16m

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-fcdc74572df17113f2efd997"></a>

- Canonical ID: `ci-fcdc74572df17113f2efd997` (1 source records).
  - [datadog/l3xqsgb18fvy](../../data/2026/W08/datadog.json); vendor URL: https://stspg.io/c6d05hsplhbv
- Start time basis: `vendor_reported`.

<a id="ci-4586ee56fed36e8b79e962d0"></a>

- Canonical ID: `ci-4586ee56fed36e8b79e962d0` (1 source records).
  - [circleci/6rpnwxd8zbt6](../../data/2026/W08/circleci.json); vendor URL: https://stspg.io/hczpgwq7t5cq
- Start time basis: `vendor_reported`.

<a id="ci-22d65fb9b2b12ef254dd38af"></a>

- Canonical ID: `ci-22d65fb9b2b12ef254dd38af` (1 source records).
  - [circleci/5xrywvlmqd9b](../../data/2026/W08/circleci.json); vendor URL: https://stspg.io/8hdkcgmsfw21
- Start time basis: `vendor_reported`.

<a id="ci-2f31de7f0713c2c2f9bfa6a9"></a>

- Canonical ID: `ci-2f31de7f0713c2c2f9bfa6a9` (1 source records).
  - [stripe/ktlly27q0x0h](../../data/2026/W08/stripe.json); vendor URL: https://stspg.io/j8ynh35wfv35
- Start time basis: `vendor_reported`.

## Sources

| Vendor | Status | Source records seen | Parse errors |
|--------|--------|----------------|--------------|
| anthropic | ok | 50 | 0 |
| atlassian | ok | 34 | 0 |
| aws | ok | 6 | 0 |
| circleci | ok | 50 | 0 |
| cloudflare | ok | 50 | 0 |
| datadog | ok | 50 | 0 |
| digitalocean | ok | 50 | 0 |
| gcp | ok | 8 | 0 |
| github | ok | 50 | 0 |
| grafana | ok | 50 | 0 |
| openai | ok | 25 | 0 |
| slack | ok | 31 | 0 |
| stripe | ok | 50 | 0 |
| yandexcloud | ok | 50 | 0 |
