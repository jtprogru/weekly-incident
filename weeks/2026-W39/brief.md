# Week 2026-W39 — top 5 of 18 above threshold

2026-09-21 to 2026-09-27 (UTC).

49 source records (53 saved revisions) → 49 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## 1. aws — Increased Error Rates

- URL: https://health.aws.amazon.com/health/status
- Window: 2026-09-21 23:26 to not yet (53d 13h 9m (ongoing))
- Impact: unknown, status: investigating
- Components: ec2-us-east-1

<a id="ci-1f6486ca51776216be85a1cf"></a>

- Canonical ID: `ci-1f6486ca51776216be85a1cf` (1 source records).
  - [aws/f5f654e733c72f13](../../data/2026/W39/aws.json); vendor URL: https://health.aws.amazon.com/health/status
- Start time basis: `first_observed_update`.
- Warning: AWS onset is the first observed message, not a verified impact start

  - `2026-09-21 23:26` investigating — We are investigating increased API error rates in the US-EAST-1 Region.
  - `2026-09-21 23:44` investigating — We can confirm increased error rates affecting new instance launches, including elevated API error rates for RunInstances, CreateFleet, and related EC2 launch workflows in the US-EAST-1 Region. Customers attempting to launch new EC2 instances may experience launch failures or elevated API response latency. Our engineers are actively engaged and working towards recovering the issue. We are seeing initial signs of recovery. Existing instances remain unaffected. We recommend retrying failed RunInstances and CreateFleet API calls. We will provide an update by 5:30 PM PDT or sooner.
  - `2026-09-22 00:24` investigating — Between 3:49 PM and 5:06 PM PDT, we experienced increased error rates affecting new instance launches in the US-EAST-1 Region. During this time, customers may have experienced elevated error rates when attempting to launch new EC2 instances through RunInstances, CreateFleet, and related EC2 launch workflows. AWS services that launch EC2 instances on behalf of customers, including Amazon ECS, AWS Fargate, Amazon EMR, amongst others were also impacted during this period. We identified the root cause as a recent deployment to a subsystem within the EC2 service responsible for processing instance launch requests. We initiated a rollback of the problematic deployment and began observing initial signs of recovery at 4:44 PM. Existing instances remain unaffected. The issue has been resolved and t…

## 2. cloudflare — Cloudflare One Clients are incorrectly challenged on some sites

- URL: https://www.cloudflarestatus.com/incidents/tl0vrgjy99r0
- Window: 2026-09-22 19:23 to 2026-10-14 16:54 UTC (21d 21h 31m)
- Impact: minor, status: resolved
- Components: Cloudflare One Client

<a id="ci-39234b54d251a8c74d44a8ac"></a>

- Canonical ID: `ci-39234b54d251a8c74d44a8ac` (1 source records).
  - [cloudflare/tl0vrgjy99r0](../../data/2026/W39/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/tl0vrgjy99r0
- Start time basis: `vendor_reported`.

  - `2026-09-22 19:23` investigating — Cloudflare is investigating an issue where Cloudflare One Client customers are incorrectly challenged while visiting some websites.
  - `2026-09-23 11:02` identified — The issue has been identified and a fix is being implemented.
  - `2026-10-14 16:43` monitoring — Cloudflare has implemented a fix and is currently monitoring for exceptions. We will update the status once the issue is resolved.
  - `2026-10-14 16:54` resolved — This incident has been resolved.

## 3. cloudflare — Network Performance Degradation — Asia-Pacific

- URL: https://www.cloudflarestatus.com/incidents/8wmvkv5jkf15
- Window: 2026-09-23 18:42 to 2026-09-29 23:01 UTC (6d 4h 18m)
- Impact: minor, status: resolved
- Components: Network

<a id="ci-cea3259e86f3208c2a946dae"></a>

- Canonical ID: `ci-cea3259e86f3208c2a946dae` (1 source records).
  - [cloudflare/8wmvkv5jkf15](../../data/2026/W39/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/8wmvkv5jkf15
- Start time basis: `vendor_reported`.

  - `2026-09-23 18:42` investigating — Since September 21, 2026 at 02:20 UTC, multiple subsea cable outages have caused congestion between Tokyo and Singapore datacenters. We have rerouted traffic to reduce impact and are working with third-party vendors to restore capacity.
  - `2026-09-23 18:44` identified — The issue has been identified and a fix is being implemented.
  - `2026-09-25 08:50` identified — Cloudflare continues to mitigate intermittent connectivity degradation for a small subset of customers caused by significant degradation of capacity in the region.
  - `2026-09-29 23:01` resolved — This incident has been resolved.

## 4. github — Incident across several services

- URL: https://stspg.io/m1ps7yrhp4n8
- Window: 2026-09-23 10:11 to 2026-09-24 04:55 UTC (18h 44m)
- Impact: minor, status: resolved
- Components: API Requests

<a id="ci-04770eb84a28e9bd42bc0367"></a>

- Canonical ID: `ci-04770eb84a28e9bd42bc0367` (1 source records).
  - [github/8zc63m64hy36](../../data/2026/W39/github.json); vendor URL: https://stspg.io/m1ps7yrhp4n8
- Start time basis: `vendor_reported`.

  - `2026-09-23 10:11` investigating — We are investigating reports of degraded performance for API Requests
  - `2026-09-23 10:20` investigating — Database replicas have detached. We're working to restore the database replicas. Users may experience issues beyond creating organizations and a degraded experience with the GitHub API and Projects.
  - `2026-09-23 10:57` investigating — Database replicas have been restored. Org creation and the GitHub API are no longer degraded.
  - `2026-09-23 10:58` investigating — The degradation affecting API Requests has been mitigated. We are monitoring to ensure stability.
  - `2026-09-23 11:38` investigating — We are seeing recovery for Projects. Users may experience stale search results for Projects while indexing catches up.
  - `2026-09-23 13:35` investigating — Users may experience stale Project search results. We are working to increase indexing speed. All other services are available.
  - `2026-09-23 17:01` investigating — Updates to issue labels may be delayed in being reflected in Projects by about ~10 minutes. We have added some capacity to work through the backlog more quickly, but it'll likely be a few hours to complete processing the full backlog of messages. All other services are available.
  - `2026-09-23 17:30` investigating — We will post another update in approximately one hour to share our progress.
  - `2026-09-23 18:42` investigating — Continuing to investigate the lag that may be experienced in issue labels being accurately reflected in Projects. We are working on alternate solutions to process the backlog of label updates.
  - `2026-09-23 20:26` investigating — We are preparing to deploy a change that will mitigate the impact.
  - `2026-09-23 21:39` investigating — Updates to issue labels may be delayed in being reflected in Projects. We are continuing to deploy a change that will accelerate processing of the backlog of label updates. All other services are available. We will provide another update within the next hour.
  - `2026-09-24 00:16` investigating — We've deployed a change intended to accelerate processing of the backlog of issue label updates in Projects. A sizable backlog still remains and we continue working through it. All other services are operating normally. We will provide another update within the next hour.
  - `2026-09-24 02:00` investigating — We are continuing to process the backlog of issue label updates for Projects. Users may still see delays before label changes are reflected in Projects. All other services are operating normally.
  - `2026-09-24 03:14` investigating — We are continuing to process the backlog of issue label updates for Projects. Label changes may still be delayed. All other services are operating normally.
  - `2026-09-24 04:55` monitoring — The degradation has been mitigated. We are monitoring to ensure stability.
  - `2026-09-24 04:55` resolved — Starting at 07:57 UTC on September 23, GitHub experienced elevated 500 and 404 responses across several application pages. This caused failures when installing GitHub Apps, creating organizations, and making some organization membership changes. Customers also experienced delayed label updates and stale search results in Projects. <br /><br />The infrastructure failure was isolated to our Azure Central US region. The API errors were mitigated by 10:58 UTC on September 23. Projects' processing continued to recover while an accumulated backlog was drained, and full service was restored at 04:55 UTC on September 24. <br /><br />The incident was caused by a failed planned maintenance operation on a primary database. Automated recovery initiated an emergency database failover, after which sever…

## 5. cloudflare — Issues with Durable Objects

- URL: https://www.cloudflarestatus.com/incidents/gv29fd076q2t
- Window: 2026-09-25 00:03 to 2026-09-25 14:49 UTC (14h 45m)
- Impact: minor, status: resolved
- Components: Durable Objects

<a id="ci-68070bea63cd77f508eda034"></a>

- Canonical ID: `ci-68070bea63cd77f508eda034` (1 source records).
  - [cloudflare/gv29fd076q2t](../../data/2026/W39/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/gv29fd076q2t
- Start time basis: `vendor_reported`.

  - `2026-09-25 00:03` investigating — Cloudflare is aware of and investigating an issue impacting some customers where new Workflows instances may be stuck in a Queued state. More updates to follow shortly.
  - `2026-09-25 01:31` monitoring — A fix has been implemented and we are monitoring the results.
  - `2026-09-25 02:32` investigating — Cloudflare is investigating issues causing an increase of errors for Durable Objects .We are working to analyse and mitigate this problem. More updates to follow shortly.
  - `2026-09-25 08:48` monitoring — A fix has been implemented and we are monitoring the results.
  - `2026-09-25 14:49` resolved — This incident has been resolved.

## Also above threshold (13)

- `cloudflare` Intermittent authentication errors for API and R2 (11h 50m) — https://www.cloudflarestatus.com/incidents/jg6ys8sbd6ny
- `grafana` IRM Access Issues for a Small Group of Users (20h 10m) — https://stspg.io/cqfnrz8bdkvz
- `grafana` Kubernetes Observability Billing & Usage Incorrect (16h 38m) — https://stspg.io/4ntjl2yr75n4
- `anthropic` Elevated errors for multiple models (1h 37m) — https://stspg.io/90s4l9yywzbk
- `cloudflare` Browser isolation session data sync issues (5h 11m) — https://www.cloudflarestatus.com/incidents/mltthh0jxxpz
- `cloudflare` Delays Starting Cloudflare Workers Builds (3h 18m) — https://www.cloudflarestatus.com/incidents/dwfzz0bwhmsq
- `cloudflare` Partial visibility of Workers and Durable Objects analytics for Fedramp High customers (3h 15m) — https://www.cloudflarestatus.com/incidents/dcmwvjwfyhjg
- `github` Disruption with billing information updates (3h 49m) — https://stspg.io/g2q4g5kv9frp
- `cloudflare` Cloudflare Access updates delayed (3h 4m) — https://www.cloudflarestatus.com/incidents/xbgbb80yw0jh
- `stripe` API call success rate drop (2h 14m) — https://stspg.io/vfb8m896t1fk
- `cloudflare` Elevated Errors with any / all in http_response_cache_settings (2h 13m) — https://www.cloudflarestatus.com/incidents/f8wg5htd3zz4
- `openai` Issues with Codex (55m) — https://status.openai.com/incidents/01M3DCNWMW57HYK8FJ5FBFPA39
- `stripe` Elevated API Errors (18m) — https://stspg.io/4ntbqnhzg7bb
