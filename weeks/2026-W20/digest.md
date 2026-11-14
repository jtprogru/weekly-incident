# Week 2026-W20

2026-05-11 to 2026-05-17 (UTC). Generated 2026-11-14 00:00.

11 canonical incidents, 6 above threshold. All 14 sources responded.

11 source records (11 saved revisions) → 11 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | digitalocean | [DNS service, Certificates and Managed MongoDB](https://stspg.io/r6dtr7gc9pch) | 2026-05-14 13:57 | 6h 49m | minor | DNS, Global | 654 = 409 × 2 × 0.8 |
| 2 | circleci | [Elevated 429 errors pulling from maven central](https://stspg.io/c8h5p945vnkw) | 2026-05-14 16:49 | 2h 49m | minor | Docker Jobs, Machine Jobs, macOS Jobs | 507 = 169 × 3 × 1.0 |
| 3 | stripe | [Elevated MobilePay errors](https://stspg.io/06tvvhhcgy80) | 2026-05-16 16:01 | 2h 51m | major | Acquirers and payment methods | 205 = 171 × 1 × 1.2 |
| 4 | datadog | [Azure Metrics Reporting](https://stspg.io/l0b8n3g0ysky) | 2026-05-16 02:17 | 1h 25m | critical | Metrics and Infra Monitoring | 85 = 85 × 1 × 1.0 |
| 5 | digitalocean | [GradientAI: Agent Platform Playground Interaction Errors](https://stspg.io/ll1m0fw2m4g9) | 2026-05-12 10:50 | 47m | major | Agent Runtime | 38 = 47 × 1 × 0.8 |
| 6 | digitalocean | [Anthropic reported outage that's impacting access to their Serverless Inference models](https://stspg.io/d47rft5t1t68) | 2026-05-15 01:21 | 38m | major | Agent Runtime | 30 = 38 × 1 × 0.8 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. digitalocean — DNS service, Certificates and Managed MongoDB

- URL: https://stspg.io/r6dtr7gc9pch
- Started: 2026-05-14 13:57 UTC
- Resolved: 2026-05-14 20:46 UTC
- Duration: 6h 49m
- Impact: minor, status: resolved
- Components: DNS, Global
- Canonical incident: [ci-8dc9c4e7028b3dcfe8964dba](#ci-8dc9c4e7028b3dcfe8964dba)

- `2026-05-14 13:57` **investigating** — Our Engineering team is investigating an issue affecting our DNS service. At this time, DNS resolution remains functional but any changes to DNS records are not being reflected at the edge. Additionally, customers may be unable to create new Let's Encrypt certificates at this time. Our engineering team is actively working to identify the root cause and restore full functionality. We apologize for any inconvenience, and we'll share more information as it becomes available.
- `2026-05-14 15:52` **investigating** — Our engineering team continues to investigate an issue affecting our DNS service. At this time, DNS resolution remains functional; however, new changes to DNS records are currently not being reflected at the edge. This issue is also impacting related services. Specifically, customers may be unable to create new Let's Encrypt certificates. Regarding Managed MongoDB, customers will be able to submit requests to create new clusters or scale existing ones, but the completion of the provisioning process is currently delayed. This is due to the dependency on Let's Encrypt certificate issuance, which requires functional DNS propagation. Affected clusters will automatically recover and complete their provisioning once the DNS issue is resolved. Our engineering team is actively working to restore full functionality across all affected services. We apologize for the inconvenience and will share more information as it becomes available.
- `2026-05-14 16:42` **monitoring** — Our Engineering team has implemented a fix for the issue affecting our DNS service and are now seeing DNS record updates successfully propagate to the edge. As a result, Let's Encrypt certificate issuance and Managed MongoDB provisioning (including scaling operations) should resume. We are currently monitoring the systems as they process the backlog of delayed requests. Affected MongoDB clusters and certificate requests should complete their provisioning automatically. We will continue to monitor the situation closely to ensure full stability. We appreciate your patience throughout this process and will provide an update once the issue is fully confirmed as resolved.
- `2026-05-14 20:46` **resolved** — As of 19:50 UTC, the issue affecting the DNS service has been fully resolved, and all related services, including Let's Encrypt certificate issuance and Managed MongoDB provisioning, are now operating normally. The backlog of delayed requests has been successfully processed, and all DNS record updates, pending certificate, and MongoDB operations should now be complete. We apologize for the disruption this caused and appreciate your patience while our team worked to restore full functionality. However, if you continue to experience any issues, please don't hesitate to raise a support ticket for further investigation. We'll be happy to assist you.

### 2. circleci — Elevated 429 errors pulling from maven central

- URL: https://stspg.io/c8h5p945vnkw
- Started: 2026-05-14 16:49 UTC
- Resolved: 2026-05-14 19:38 UTC
- Duration: 2h 49m
- Impact: minor, status: resolved
- Components: Docker Jobs, Machine Jobs, macOS Jobs
- Canonical incident: [ci-a76c2e4b19a99a21c55a6b6d](#ci-a76c2e4b19a99a21c55a6b6d)

- `2026-05-14 16:49` **investigating** — Jobs pulling dependencies from maven central are experiencing increased 429 errors due to maven central rate limits. We are investigating the issue and working to identify mitigations for customer jobs.
- `2026-05-14 18:08` **monitoring** — We're monitoring to verify that rate limited jobs are recovering.
- `2026-05-14 19:38` **resolved** — Customers running jobs that pulled dependencies from Maven Central experienced elevated 429 (rate limit) errors. We have worked with Sonatype to allow list CircleCI's dedicated IP ranges. Customers using our IP ranges feature should no longer see Maven Central rate limits affect their jobs while we continue work to fully address the issue for all customers. Customers whose jobs failed during this window may rerun affected jobs. We thank you for your patience while our team worked on implementing a fix.

### 3. stripe — Elevated MobilePay errors

- URL: https://stspg.io/06tvvhhcgy80
- Started: 2026-05-16 16:01 UTC
- Resolved: 2026-05-16 18:53 UTC
- Duration: 2h 51m
- Impact: major, status: resolved
- Components: Acquirers and payment methods
- Canonical incident: [ci-e811d98eb0a3c2f907d6623d](#ci-e811d98eb0a3c2f907d6623d)

- `2026-05-16 16:01` **investigating** — We're currently investigating elevated MobilePay errors. We'll post an update soon.
- `2026-05-16 18:23` **investigating** — Between 14:18–17:20 UTC May 16th, we saw elevated error rates and response times with MobilePay requests. This is now resolved.
- `2026-05-16 18:53` **resolved** — Between 14:18–17:20 UTC May 16th, we saw elevated error rates and response times with MobilePay requests. This is now resolved.

### 4. datadog — Azure Metrics Reporting

- URL: https://stspg.io/l0b8n3g0ysky
- Started: 2026-05-16 02:17 UTC
- Resolved: 2026-05-16 03:43 UTC
- Duration: 1h 25m
- Impact: critical, status: resolved
- Components: Metrics and Infra Monitoring
- Canonical incident: [ci-115f93cb7e4c7ab94b7db4de](#ci-115f93cb7e4c7ab94b7db4de)

- `2026-05-16 02:17` **investigating** — We are investigating an issue submitting Azure metrics.
- `2026-05-16 02:40` **identified** — The issue has been identified and a fix is being implemented.
- `2026-05-16 03:09` **identified** — We are continuing to work on a fix for this issue.
- `2026-05-16 03:27` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-05-16 03:43` **resolved** — This incident has been resolved and Azure metrics are reporting as expected.

### 5. digitalocean — GradientAI: Agent Platform Playground Interaction Errors

- URL: https://stspg.io/ll1m0fw2m4g9
- Started: 2026-05-12 10:50 UTC
- Resolved: 2026-05-12 11:37 UTC
- Duration: 47m
- Impact: major, status: resolved
- Components: Agent Runtime
- Canonical incident: [ci-e8665ca5d4f274e915898b73](#ci-e8665ca5d4f274e915898b73)

- `2026-05-12 10:24` **investigating** — We are currently investigating an issue affecting the GradientAI Agent Platform Playground. Users may see a “Something went wrong” error for all agent interactions in the Playground. Agent functionality through API endpoints remains unaffected. We are actively working to identify the cause and will provide an update as soon as more information is available.
- `2026-05-12 10:50` **identified** — We have identified the cause of the issue affecting the GradientAI Agent Platform Playground. Our engineering team is implementing a fix, and we will provide another update as soon as it is available.
- `2026-05-12 11:14` **identified** — A fix has been applied for the issue affecting the GradientAI Agent Platform Playground, and we are observing recovery in service behavior.
- `2026-05-12 11:36` **monitoring** — We have applied a fix for the issue affecting the GradientAI Agent Platform Playground and are monitoring the service to ensure continued stability. At this time, we are no longer observing new reports related to this issue.
- `2026-05-12 11:37` **resolved** — The remediation for the issue affecting the GradientAI Agent Platform Playground has been fully implemented, and the service is operating normally. We are no longer observing customer impact related to this issue, and the incident has been resolved.

### 6. digitalocean — Anthropic reported outage that's impacting access to their Serverless Inference models

- URL: https://stspg.io/d47rft5t1t68
- Started: 2026-05-15 01:21 UTC
- Resolved: 2026-05-15 02:00 UTC
- Duration: 38m
- Impact: major, status: resolved
- Components: Agent Runtime
- Canonical incident: [ci-e120edc6536dbd164c96001c](#ci-e120edc6536dbd164c96001c)

- `2026-05-15 01:21` **monitoring** — As of the current time, our Engineering team is aware of an ongoing incident with our provider, Anthropic, that is impacting Serverless Inference. The outage is affecting all users attempting to use Sonnet 4.6 and Opus 4.7 models. According to Anthropic's status page (https://status.claude.com/incidents/8z7l5zcy0v3b), they are currently experiencing an outage that is causing this disruption. We apologize for the inconvenience and will provide updates as more information becomes available.
- `2026-05-15 01:23` **identified** — According to Anthropic's status page (https://status.claude.com/incidents/8z7l5zcy0v3b), they have identified the root cause and are actively working on a fix. At this time, users may continue to experience errors when attempting to use Sonnet 4.6 and Opus 4.7 models. We will provide another update once Anthropic has implemented a fix and service is restored.
- `2026-05-15 02:00` **resolved** — As of the current time, our Engineering team has confirmed that the issue with Serverless Inference has been resolved. The root cause of the issue was an outage with our provider, Anthropic, which affected users of Sonnet 4.6 and Opus 4.7 models. According to Anthropic's status page (https://status.claude.com/incidents/8z7l5zcy0v3b), they have resolved the outage. If you continue to experience problems, please open a ticket with our Support team. We apologize for any inconvenience this may have caused.

## Below threshold (5)

- `slack` [Issues affecting file uploads, message edits, channel renames, and channel creations](https://slack-status.com/2026-05/fe557ca05fdb64fa) — 57m
- `datadog` [Delayed Metric Loading](https://stspg.io/sskp2jtfd9sj) — 1h 23m
- `stripe` [Elevated Apple Pay errors](https://stspg.io/m72fwk8k8qfg) — 1h 42m
- `datadog` [Degraded Web Application Performance](https://stspg.io/cwzjd95rnzhd) — 36m
- `stripe` [Elevated Swish errors](https://stspg.io/0zpfx7w4g20f) — 36m

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-8dc9c4e7028b3dcfe8964dba"></a>

- Canonical ID: `ci-8dc9c4e7028b3dcfe8964dba` (1 source records).
  - [digitalocean/gwgh41d5r0nz](../../data/2026/W20/digitalocean.json); vendor URL: https://stspg.io/r6dtr7gc9pch
- Start time basis: `vendor_reported`.

<a id="ci-a76c2e4b19a99a21c55a6b6d"></a>

- Canonical ID: `ci-a76c2e4b19a99a21c55a6b6d` (1 source records).
  - [circleci/qhxl36d29vfl](../../data/2026/W20/circleci.json); vendor URL: https://stspg.io/c8h5p945vnkw
- Start time basis: `vendor_reported`.

<a id="ci-114752388364a22478b40082"></a>

- Canonical ID: `ci-114752388364a22478b40082` (1 source records).
  - [slack/1553](../../data/2026/W20/slack.json); vendor URL: https://slack-status.com/2026-05/fe557ca05fdb64fa
- Start time basis: `vendor_reported`.

<a id="ci-e811d98eb0a3c2f907d6623d"></a>

- Canonical ID: `ci-e811d98eb0a3c2f907d6623d` (1 source records).
  - [stripe/ws0j885npfx1](../../data/2026/W20/stripe.json); vendor URL: https://stspg.io/06tvvhhcgy80
- Start time basis: `vendor_reported`.

<a id="ci-89746412d1f29bb81a71ba15"></a>

- Canonical ID: `ci-89746412d1f29bb81a71ba15` (1 source records).
  - [datadog/shjkvbwy06nl](../../data/2026/W20/datadog.json); vendor URL: https://stspg.io/sskp2jtfd9sj
- Start time basis: `vendor_reported`.

<a id="ci-26ddd961b436cb4073952d0a"></a>

- Canonical ID: `ci-26ddd961b436cb4073952d0a` (1 source records).
  - [stripe/f5l7tqzmc9mm](../../data/2026/W20/stripe.json); vendor URL: https://stspg.io/m72fwk8k8qfg
- Start time basis: `vendor_reported`.

<a id="ci-115f93cb7e4c7ab94b7db4de"></a>

- Canonical ID: `ci-115f93cb7e4c7ab94b7db4de` (1 source records).
  - [datadog/xqyrslhjs303](../../data/2026/W20/datadog.json); vendor URL: https://stspg.io/l0b8n3g0ysky
- Start time basis: `vendor_reported`.

<a id="ci-d110986be9f5076f273c3e9b"></a>

- Canonical ID: `ci-d110986be9f5076f273c3e9b` (1 source records).
  - [datadog/j4v5lc1lvgf1](../../data/2026/W20/datadog.json); vendor URL: https://stspg.io/cwzjd95rnzhd
- Start time basis: `vendor_reported`.

<a id="ci-dd7261e4b346ea17bf7452d6"></a>

- Canonical ID: `ci-dd7261e4b346ea17bf7452d6` (1 source records).
  - [stripe/mkbj4zw0w85w](../../data/2026/W20/stripe.json); vendor URL: https://stspg.io/0zpfx7w4g20f
- Start time basis: `vendor_reported`.

<a id="ci-e8665ca5d4f274e915898b73"></a>

- Canonical ID: `ci-e8665ca5d4f274e915898b73` (1 source records).
  - [digitalocean/bnq93nhgvtm6](../../data/2026/W20/digitalocean.json); vendor URL: https://stspg.io/ll1m0fw2m4g9
- Start time basis: `vendor_reported`.

<a id="ci-e120edc6536dbd164c96001c"></a>

- Canonical ID: `ci-e120edc6536dbd164c96001c` (1 source records).
  - [digitalocean/lrrw084g3whn](../../data/2026/W20/digitalocean.json); vendor URL: https://stspg.io/d47rft5t1t68
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
