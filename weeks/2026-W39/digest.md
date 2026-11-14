# Week 2026-W39

2026-09-21 to 2026-09-27 (UTC). Generated 2026-11-14 12:36.

49 canonical incidents, 18 above threshold. 2 of 14 sources unavailable: datadog, yandexcloud.

49 source records (53 saved revisions) → 49 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## Above threshold

| # | Vendor | Incident | Started (UTC) | Duration | Impact | Components | Score |
|---|--------|----------|---------------|----------|--------|------------|-------|
| 1 | aws | [Increased Error Rates](https://health.aws.amazon.com/health/status) | 2026-09-21 23:26 | 53d 13h 9m (ongoing) | unknown | ec2-us-east-1 | 154218 = 77109 × 1 × 2.0 |
| 2 | cloudflare | [Cloudflare One Clients are incorrectly challenged on some sites](https://www.cloudflarestatus.com/incidents/tl0vrgjy99r0) | 2026-09-22 19:23 | 21d 21h 31m | minor | Cloudflare One Client | 56756 = 31531 × 1 × 1.8 |
| 3 | cloudflare | [Network Performance Degradation — Asia-Pacific](https://www.cloudflarestatus.com/incidents/8wmvkv5jkf15) | 2026-09-23 18:42 | 6d 4h 18m | minor | Network | 16016 = 8898 × 1 × 1.8 |
| 4 | github | [Incident across several services](https://stspg.io/m1ps7yrhp4n8) | 2026-09-23 10:11 | 18h 44m | minor | API Requests | 1686 = 1124 × 1 × 1.5 |
| 5 | cloudflare | [Issues with Durable Objects](https://www.cloudflarestatus.com/incidents/gv29fd076q2t) | 2026-09-25 00:03 | 14h 45m | minor | Durable Objects | 1593 = 885 × 1 × 1.8 |
| 6 | cloudflare | [Intermittent authentication errors for API and R2](https://www.cloudflarestatus.com/incidents/jg6ys8sbd6ny) | 2026-09-23 09:08 | 11h 50m | minor | R2 | 1278 = 710 × 1 × 1.8 |
| 7 | grafana | [IRM Access Issues for a Small Group of Users](https://stspg.io/cqfnrz8bdkvz) | 2026-09-22 14:57 | 20h 10m | major | Incident Management and Response (IRM) | 1210 = 1210 × 1 × 1.0 |
| 8 | grafana | [Kubernetes Observability Billing & Usage Incorrect](https://stspg.io/4ntjl2yr75n4) | 2026-09-22 21:27 | 16h 38m | minor | Cost Management and Billing | 998 = 998 × 1 × 1.0 |
| 9 | anthropic | [Elevated errors for multiple models](https://stspg.io/90s4l9yywzbk) | 2026-09-22 00:57 | 1h 37m | major | Claude API (api.anthropic.com), Claude Code, Claude Cowork, claude.ai | 582 = 97 × 4 × 1.5 |
| 10 | cloudflare | [Browser isolation session data sync issues](https://www.cloudflarestatus.com/incidents/mltthh0jxxpz) | 2026-09-24 17:26 | 5h 11m | minor | Browser Isolation | 560 = 311 × 1 × 1.8 |
| 11 | cloudflare | [Delays Starting Cloudflare Workers Builds](https://www.cloudflarestatus.com/incidents/dwfzz0bwhmsq) | 2026-09-24 17:33 | 3h 18m | minor | Workers Builds | 356 = 198 × 1 × 1.8 |
| 12 | cloudflare | [Partial visibility of Workers and Durable Objects analytics for Fedramp High customers](https://www.cloudflarestatus.com/incidents/dcmwvjwfyhjg) | 2026-09-22 09:36 | 3h 15m | minor | Analytics | 351 = 195 × 1 × 1.8 |
| 13 | github | [Disruption with billing information updates](https://stspg.io/g2q4g5kv9frp) | 2026-09-24 16:51 | 3h 49m | minor | — | 344 = 229 × 1 × 1.5 |
| 14 | cloudflare | [Cloudflare Access updates delayed](https://www.cloudflarestatus.com/incidents/xbgbb80yw0jh) | 2026-09-24 15:20 | 3h 4m | minor | Access | 331 = 184 × 1 × 1.8 |
| 15 | stripe | [API call success rate drop](https://stspg.io/vfb8m896t1fk) | 2026-09-22 20:32 | 2h 14m | none | Stripe API, Stripe core components | 322 = 134 × 2 × 1.2 |
| 16 | cloudflare | [Elevated Errors with any / all in http_response_cache_settings](https://www.cloudflarestatus.com/incidents/f8wg5htd3zz4) | 2026-09-23 15:54 | 2h 13m | minor | Rules | 239 = 133 × 1 × 1.8 |
| 17 | openai | [Issues with Codex](https://status.openai.com/incidents/01M3DCNWMW57HYK8FJ5FBFPA39) | 2026-09-25 22:58 | 55m | critical | — | 82 = 55 × 1 × 1.5 |
| 18 | stripe | [Elevated API Errors](https://stspg.io/4ntbqnhzg7bb) | 2026-09-27 01:03 | 18m | major | Stripe API, Stripe core components | 43 = 18 × 2 × 1.2 |

Score is duration in minutes, times component count, times vendor weight.

## Timelines

### 1. aws — Increased Error Rates

- URL: https://health.aws.amazon.com/health/status
- Started: 2026-09-21 23:26 UTC
- Resolved: not yet
- Duration: 53d 13h 9m (ongoing)
- Impact: unknown, status: investigating
- Components: ec2-us-east-1
- Canonical incident: [ci-1f6486ca51776216be85a1cf](#ci-1f6486ca51776216be85a1cf)

- `2026-09-21 23:26` **investigating** — We are investigating increased API error rates in the US-EAST-1 Region.
- `2026-09-21 23:44` **investigating** — We can confirm increased error rates affecting new instance launches, including elevated API error rates for RunInstances, CreateFleet, and related EC2 launch workflows in the US-EAST-1 Region. Customers attempting to launch new EC2 instances may experience launch failures or elevated API response latency. Our engineers are actively engaged and working towards recovering the issue. We are seeing initial signs of recovery. Existing instances remain unaffected. We recommend retrying failed RunInstances and CreateFleet API calls. We will provide an update by 5:30 PM PDT or sooner.
- `2026-09-22 00:24` **investigating** — Between 3:49 PM and 5:06 PM PDT, we experienced increased error rates affecting new instance launches in the US-EAST-1 Region. During this time, customers may have experienced elevated error rates when attempting to launch new EC2 instances through RunInstances, CreateFleet, and related EC2 launch workflows. AWS services that launch EC2 instances on behalf of customers, including Amazon ECS, AWS Fargate, Amazon EMR, amongst others were also impacted during this period. We identified the root cause as a recent deployment to a subsystem within the EC2 service responsible for processing instance launch requests. We initiated a rollback of the problematic deployment and began observing initial signs of recovery at 4:44 PM. Existing instances remain unaffected. The issue has been resolved and the service is operating normally.

### 2. cloudflare — Cloudflare One Clients are incorrectly challenged on some sites

- URL: https://www.cloudflarestatus.com/incidents/tl0vrgjy99r0
- Started: 2026-09-22 19:23 UTC
- Resolved: 2026-10-14 16:54 UTC
- Duration: 21d 21h 31m
- Impact: minor, status: resolved
- Components: Cloudflare One Client
- Canonical incident: [ci-39234b54d251a8c74d44a8ac](#ci-39234b54d251a8c74d44a8ac)

- `2026-09-22 19:23` **investigating** — Cloudflare is investigating an issue where Cloudflare One Client customers are incorrectly challenged while visiting some websites.
- `2026-09-23 11:02` **identified** — The issue has been identified and a fix is being implemented.
- `2026-10-14 16:43` **monitoring** — Cloudflare has implemented a fix and is currently monitoring for exceptions. We will update the status once the issue is resolved.
- `2026-10-14 16:54` **resolved** — This incident has been resolved.

### 3. cloudflare — Network Performance Degradation — Asia-Pacific

- URL: https://www.cloudflarestatus.com/incidents/8wmvkv5jkf15
- Started: 2026-09-23 18:42 UTC
- Resolved: 2026-09-29 23:01 UTC
- Duration: 6d 4h 18m
- Impact: minor, status: resolved
- Components: Network
- Canonical incident: [ci-cea3259e86f3208c2a946dae](#ci-cea3259e86f3208c2a946dae)

- `2026-09-23 18:42` **investigating** — Since September 21, 2026 at 02:20 UTC, multiple subsea cable outages have caused congestion between Tokyo and Singapore datacenters. We have rerouted traffic to reduce impact and are working with third-party vendors to restore capacity.
- `2026-09-23 18:44` **identified** — The issue has been identified and a fix is being implemented.
- `2026-09-25 08:50` **identified** — Cloudflare continues to mitigate intermittent connectivity degradation for a small subset of customers caused by significant degradation of capacity in the region.
- `2026-09-29 23:01` **resolved** — This incident has been resolved.

### 4. github — Incident across several services

- URL: https://stspg.io/m1ps7yrhp4n8
- Started: 2026-09-23 10:11 UTC
- Resolved: 2026-09-24 04:55 UTC
- Duration: 18h 44m
- Impact: minor, status: resolved
- Components: API Requests
- Canonical incident: [ci-04770eb84a28e9bd42bc0367](#ci-04770eb84a28e9bd42bc0367)

- `2026-09-23 10:11` **investigating** — We are investigating reports of degraded performance for API Requests
- `2026-09-23 10:20` **investigating** — Database replicas have detached. We're working to restore the database replicas. Users may experience issues beyond creating organizations and a degraded experience with the GitHub API and Projects.
- `2026-09-23 10:57` **investigating** — Database replicas have been restored. Org creation and the GitHub API are no longer degraded.
- `2026-09-23 10:58` **investigating** — The degradation affecting API Requests has been mitigated. We are monitoring to ensure stability.
- `2026-09-23 11:38` **investigating** — We are seeing recovery for Projects. Users may experience stale search results for Projects while indexing catches up.
- `2026-09-23 13:35` **investigating** — Users may experience stale Project search results. We are working to increase indexing speed. All other services are available.
- `2026-09-23 17:01` **investigating** — Updates to issue labels may be delayed in being reflected in Projects by about ~10 minutes. We have added some capacity to work through the backlog more quickly, but it'll likely be a few hours to complete processing the full backlog of messages. All other services are available.
- `2026-09-23 17:30` **investigating** — We will post another update in approximately one hour to share our progress.
- `2026-09-23 18:42` **investigating** — Continuing to investigate the lag that may be experienced in issue labels being accurately reflected in Projects. We are working on alternate solutions to process the backlog of label updates.
- `2026-09-23 20:26` **investigating** — We are preparing to deploy a change that will mitigate the impact.
- `2026-09-23 21:39` **investigating** — Updates to issue labels may be delayed in being reflected in Projects. We are continuing to deploy a change that will accelerate processing of the backlog of label updates. All other services are available. We will provide another update within the next hour.
- `2026-09-24 00:16` **investigating** — We've deployed a change intended to accelerate processing of the backlog of issue label updates in Projects. A sizable backlog still remains and we continue working through it. All other services are operating normally. We will provide another update within the next hour.
- `2026-09-24 02:00` **investigating** — We are continuing to process the backlog of issue label updates for Projects. Users may still see delays before label changes are reflected in Projects. All other services are operating normally.
- `2026-09-24 03:14` **investigating** — We are continuing to process the backlog of issue label updates for Projects. Label changes may still be delayed. All other services are operating normally.
- `2026-09-24 04:55` **monitoring** — The degradation has been mitigated. We are monitoring to ensure stability.
- `2026-09-24 04:55` **resolved** — Starting at 07:57 UTC on September 23, GitHub experienced elevated 500 and 404 responses across several application pages. This caused failures when installing GitHub Apps, creating organizations, and making some organization membership changes. Customers also experienced delayed label updates and stale search results in Projects. <br /><br />The infrastructure failure was isolated to our Azure Central US region. The API errors were mitigated by 10:58 UTC on September 23. Projects' processing continued to recover while an accumulated backlog was drained, and full service was restored at 04:55 UTC on September 24. <br /><br />The incident was caused by a failed planned maintenance operation on a primary database. Automated recovery initiated an emergency database failover, after which several replicas in the affected region were unable to resume replication correctly. This reduced available database capacity and caused the API errors and downstream Projects processing delays. <br /><br />We have mitigated the immediate failure mode. We are also improving maintenance safety checks, database failover handling, post-failover replica validation, and downstream processing resilience to reduce the likelihood and impact of similar incidents.

### 5. cloudflare — Issues with Durable Objects

- URL: https://www.cloudflarestatus.com/incidents/gv29fd076q2t
- Started: 2026-09-25 00:03 UTC
- Resolved: 2026-09-25 14:49 UTC
- Duration: 14h 45m
- Impact: minor, status: resolved
- Components: Durable Objects
- Canonical incident: [ci-68070bea63cd77f508eda034](#ci-68070bea63cd77f508eda034)

- `2026-09-25 00:03` **investigating** — Cloudflare is aware of and investigating an issue impacting some customers where new Workflows instances may be stuck in a Queued state. More updates to follow shortly.
- `2026-09-25 01:31` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-09-25 02:32` **investigating** — Cloudflare is investigating issues causing an increase of errors for Durable Objects .We are working to analyse and mitigate this problem. More updates to follow shortly.
- `2026-09-25 08:48` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-09-25 14:49` **resolved** — This incident has been resolved.

### 6. cloudflare — Intermittent authentication errors for API and R2

- URL: https://www.cloudflarestatus.com/incidents/jg6ys8sbd6ny
- Started: 2026-09-23 09:08 UTC
- Resolved: 2026-09-23 20:59 UTC
- Duration: 11h 50m
- Impact: minor, status: resolved
- Components: R2
- Canonical incident: [ci-434bb27d04b018913a1fff67](#ci-434bb27d04b018913a1fff67)

- `2026-09-23 09:08` **identified** — A known issue affected authentication for a small percentage of requests to the API and R2. The issue was identified on Sep 22 at 13:30 UTC, and major impact was mitigated at 19:00 UTC. Our team is actively working to resolve the residual impact.
- `2026-09-23 20:09` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-09-23 20:59` **resolved** — This incident has been resolved.

### 7. grafana — IRM Access Issues for a Small Group of Users

- URL: https://stspg.io/cqfnrz8bdkvz
- Started: 2026-09-22 14:57 UTC
- Resolved: 2026-09-23 11:07 UTC
- Duration: 20h 10m
- Impact: major, status: resolved
- Components: Incident Management and Response (IRM)
- Canonical incident: [ci-7a823295b5dda7b73729aa90](#ci-7a823295b5dda7b73729aa90)

- `2026-09-22 14:57` **investigating** — We are currently investigating an issue that may cause a small group of users to experience unexpected IRM access issues, or undelivered pages. Most users should not be affected.
- `2026-09-22 15:55` **identified** — The issue has been identified, and we are working on a fix.
- `2026-09-22 17:19` **monitoring** — A fix is currently being rolled out and will be deployed progressively over the next few hours. We appreciate your patience as this reaches all affected customers.
- `2026-09-22 18:57` **monitoring** — The fix has been deployed to the majority of deployments. We continue to monitor progress.
- `2026-09-23 11:07` **resolved** — Between 14:57 UTC on 22 September and 09:00 UTC on 23 September, a small group of users experienced unexpected access issues and undelivered pages in Grafana IRM. This has been fully resolved and all affected users should now have normal access. Thank you for your patience.

### 8. grafana — Kubernetes Observability Billing & Usage Incorrect

- URL: https://stspg.io/4ntjl2yr75n4
- Started: 2026-09-22 21:27 UTC
- Resolved: 2026-09-23 14:06 UTC
- Duration: 16h 38m
- Impact: minor, status: resolved
- Components: Cost Management and Billing
- Canonical incident: [ci-b91802663877f3ca7681de43](#ci-b91802663877f3ca7681de43)

- `2026-09-22 21:27` **investigating** — Some customers may see a sudden spike in reported Kubernetes Observability usage even if they haven't activated the product. We are investigating the root cause, and will ensure the amount is corrected before any usage is invoiced.
- `2026-09-22 23:10` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-09-23 14:06` **resolved** — This incident has been resolved.

### 9. anthropic — Elevated errors for multiple models

- URL: https://stspg.io/90s4l9yywzbk
- Started: 2026-09-22 00:57 UTC
- Resolved: 2026-09-22 02:35 UTC
- Duration: 1h 37m
- Impact: major, status: resolved
- Components: Claude API (api.anthropic.com), Claude Code, Claude Cowork, claude.ai
- Canonical incident: [ci-90caeb635c9c77d87523f942](#ci-90caeb635c9c77d87523f942)

- `2026-09-22 00:57` **investigating** — We are investigating elevated errors on requests to Claude Mythos 5.1, Claude Fable 5.1, and Claude Opus 5. We will provide an update as soon as possible.
- `2026-09-22 01:17` **identified** — We have identified the cause of elevated errors on requests to Claude Mythos 5.1, Claude Fable 5.1, and Claude Opus 5 and are working on a fix. We will provide an update as soon as possible.
- `2026-09-22 01:35` **identified** — We are continuing to work to resolve errors affecting some models. At this time, requests to Claude Fable 5 and 5.1 and Mythos 5 and 5.1 have returned to normal success rates. We are working to resolve remaining errors affecting Claude Opus 5, and will provide an additional update shortly.
- `2026-09-22 02:11` **monitoring** — We have seen success rates return to normal across affected models, and are monitoring closely to ensure no further issues.
- `2026-09-22 02:35` **resolved** — This issue has been resolved. Impact occurred from 5:50pm PT / 00:50 UTC to 7:10pm PT / 02:10 UTC.

### 10. cloudflare — Browser isolation session data sync issues

- URL: https://www.cloudflarestatus.com/incidents/mltthh0jxxpz
- Started: 2026-09-24 17:26 UTC
- Resolved: 2026-09-24 22:37 UTC
- Duration: 5h 11m
- Impact: minor, status: resolved
- Components: Browser Isolation
- Canonical incident: [ci-92a8bc8535ec4a3a450f1af8](#ci-92a8bc8535ec4a3a450f1af8)

- `2026-09-24 17:26` **identified** — Cloudflare is investigating issues with Browser Isolation's session data sync. Customers that use Browser Isolation may observe errors on acquiring a browser stating that browser storage is unavailable, and they will need to sign into web apps again. They will continue to be able to use browsers during this time.
- `2026-09-24 17:28` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-09-24 22:37` **resolved** — This incident has been resolved.

### 11. cloudflare — Delays Starting Cloudflare Workers Builds

- URL: https://www.cloudflarestatus.com/incidents/dwfzz0bwhmsq
- Started: 2026-09-24 17:33 UTC
- Resolved: 2026-09-24 20:52 UTC
- Duration: 3h 18m
- Impact: minor, status: resolved
- Components: Workers Builds
- Canonical incident: [ci-cf225c9cb7eabbc58dc7d19b](#ci-cf225c9cb7eabbc58dc7d19b)

- `2026-09-24 17:33` **monitoring** — Customers may experience delays starting Cloudflare Workers builds.
- `2026-09-24 19:54` **monitoring** — Queue levels & latency are returning to normal
- `2026-09-24 20:52` **resolved** — This incident has been resolved.

### 12. cloudflare — Partial visibility of Workers and Durable Objects analytics for Fedramp High customers

- URL: https://www.cloudflarestatus.com/incidents/dcmwvjwfyhjg
- Started: 2026-09-22 09:36 UTC
- Resolved: 2026-09-22 12:51 UTC
- Duration: 3h 15m
- Impact: minor, status: resolved
- Components: Analytics
- Canonical incident: [ci-3f3eb5bd4ea507cb1460c2e6](#ci-3f3eb5bd4ea507cb1460c2e6)

- `2026-09-22 09:36` **identified** — Cloudflare is monitoring the restoration of Workers and Durable Objects analytics visibility for all Fedramp High customers. While the underlying data remains intact, affected datasets were not being correctly returned by our analytics dashboard and GraphQL API since June 11th. Services are currently recovering as we propagate the fix.
- `2026-09-22 10:08` **monitoring** — A fix has been implemented and we are monitoring the results.
- `2026-09-22 12:51` **resolved** — This incident has been resolved.

### 13. github — Disruption with billing information updates

- URL: https://stspg.io/g2q4g5kv9frp
- Started: 2026-09-24 16:51 UTC
- Resolved: 2026-09-24 20:41 UTC
- Duration: 3h 49m
- Impact: minor, status: resolved
- Components: —
- Canonical incident: [ci-02e842160015b8bcf9b815d1](#ci-02e842160015b8bcf9b815d1)

- `2026-09-24 16:51` **investigating** — We are investigating reports of impacted performance for some GitHub services.
- `2026-09-24 16:51` **investigating** — We are investigating failures on billing information updates. Customers may be unable to create or update their billing information in the meantime.
- `2026-09-24 17:40` **investigating** — We have identified the problem and are actively working on mitigation.
- `2026-09-24 18:22` **investigating** — We are continuing to work on mitigation. We will post another update in approximately one hour.
- `2026-09-24 20:10` **investigating** — We have applied the mitigation and are seeing signs of recovery. We are continuing to monitor the system closely.
- `2026-09-24 20:24` **monitoring** — The degradation has been mitigated. We are monitoring to ensure stability.
- `2026-09-24 20:41` **resolved** — Beginning on September 22, 2026 at 11:51 UTC, approximately 1,700 users experienced delays and errors when adding or updating billing address information because requests to the address validation service timed out or failed. Within the affected address validation flow, requests failed at an average rate of 69.3%, reaching 100% at peak. <br /><br />An operating system upgrade exposed a compatibility issue in an adapter used by our HTTP client library, causing requests to stall. A subsequent change to shorten request timeouts caused stalled requests to return connection errors and increased the failure rate. <br /><br />We mitigated the incident on September 24, 2026 at 20:24 UTC by changing the HTTP client used for address validation. We have added alerting to detect similar issues sooner and are extending the fix to other integrations.

### 14. cloudflare — Cloudflare Access updates delayed

- URL: https://www.cloudflarestatus.com/incidents/xbgbb80yw0jh
- Started: 2026-09-24 15:20 UTC
- Resolved: 2026-09-24 18:25 UTC
- Duration: 3h 4m
- Impact: minor, status: resolved
- Components: Access
- Canonical incident: [ci-ae494d7db7bd8e14189afad3](#ci-ae494d7db7bd8e14189afad3)

- `2026-09-24 15:20` **investigating** — Cloudflare engineering is investigating an issue causing a delay in configuration updates. This includes applications and policy updates. Authentication and policy enforcement are not impacted.
- `2026-09-24 15:31` **monitoring** — We have applied a fix and are actively monitoring.
- `2026-09-24 18:25` **resolved** — This incident has been resolved.

### 15. stripe — API call success rate drop

- URL: https://stspg.io/vfb8m896t1fk
- Started: 2026-09-22 20:32 UTC
- Resolved: 2026-09-22 22:47 UTC
- Duration: 2h 14m
- Impact: none, status: resolved
- Components: Stripe API, Stripe core components
- Canonical incident: [ci-48c078e62ca2aae8f3df281c](#ci-48c078e62ca2aae8f3df281c)

- `2026-09-22 20:32` **identified** — We're currently investigating an issue affecting API accounts and capability call success rates. We've identified the root cause and are working on a mitigation. We will post further updates as we have more information.
- `2026-09-22 21:27` **monitoring** — API error rates are recovering as part of our mitigations. We're continuing to monitor until the issue is fully resolved. We will provide further updates once we have observed sustained recovery.
- `2026-09-22 22:47` **resolved** — From 19:13 to 22:19 UTC, we saw elevated error rates and response times with the API. This is now resolved.

### 16. cloudflare — Elevated Errors with any / all in http_response_cache_settings

- URL: https://www.cloudflarestatus.com/incidents/f8wg5htd3zz4
- Started: 2026-09-23 15:54 UTC
- Resolved: 2026-09-23 18:08 UTC
- Duration: 2h 13m
- Impact: minor, status: resolved
- Components: Rules
- Canonical incident: [ci-82858f1ad2f887f239826b5b](#ci-82858f1ad2f887f239826b5b)

- `2026-09-23 15:54` **identified** — We have identified an issue causing unexpected errors when creating or updating configurations using any or all expressions in the http_response_cache_settings. This issue strictly affects API operations; existing active configurations and live traffic are not impacted. A fix is currently being deployed, and we will provide an update once the rollout is complete.
- `2026-09-23 18:08` **resolved** — This incident has been resolved.

### 17. openai — Issues with Codex

- URL: https://status.openai.com/incidents/01M3DCNWMW57HYK8FJ5FBFPA39
- Started: 2026-09-25 22:58 UTC
- Resolved: 2026-09-25 23:54 UTC
- Duration: 55m
- Impact: critical, status: resolved
- Components: —
- Canonical incident: [ci-3109e9afbc523c7d259c894a](#ci-3109e9afbc523c7d259c894a)

- `2026-09-25 22:58` **identified** — We have identified that users are experiencing a codex outage. We have identified the internal issue and are heading towards a mitigation.
- `2026-09-25 23:03` **identified** — We have identified that users are experiencing elevated errors for the impacted services. We are working on implementing a mitigation.
- `2026-09-25 23:19` **identified** — Login via API key will unblock access at this time.
- `2026-09-25 23:34` **identified** — We have identified root cause and are moving towards mitigation.
- `2026-09-25 23:45` **monitoring** — We have applied the mitigation and are monitoring the recovery.
- `2026-09-25 23:54` **resolved** — All impacted services have now fully recovered.

### 18. stripe — Elevated API Errors

- URL: https://stspg.io/4ntbqnhzg7bb
- Started: 2026-09-27 01:03 UTC
- Resolved: 2026-09-27 01:21 UTC
- Duration: 18m
- Impact: major, status: resolved
- Components: Stripe API, Stripe core components
- Canonical incident: [ci-c25a69d9c1311c50db50b4d0](#ci-c25a69d9c1311c50db50b4d0)

- `2026-09-27 01:03` **investigating** — Starting 00:32 UTC, we’ve seen elevated failures with the API. We’re investigating and will post an update as soon as possible.
- `2026-09-27 01:21` **resolved** — From 00:32 UTC to 01:01 UTC, we saw elevated API errors. This issue is now resolved.

## Below threshold (31)

- `cloudflare` [Intermittent 500 errors for backend services on Replicate](https://www.cloudflarestatus.com/incidents/2x4qt2q2ds9x) — 1h 54m
- `cloudflare` [Increased Errors for Durable Objects](https://www.cloudflarestatus.com/incidents/78c19w4xw4z2) — 1h 50m
- `stripe` [Elevated Klarna errors](https://stspg.io/4kvs6zwpn12h) — 1h 19m
- `cloudflare` [Edit Compression Rule Issue](https://www.cloudflarestatus.com/incidents/drq3pdx1f11y) — 1h 32m
- `cloudflare` [Network connectivity issues in LHR, London](https://www.cloudflarestatus.com/incidents/nhf2kr8q3nfw) — 1h 11m
- `cloudflare` [Bot Management Configuration Propagation Issues](https://www.cloudflarestatus.com/incidents/c2f7zj7fs8l8) — 1h 3m
- `circleci` [Delays in processing new signups and plan changes](https://stspg.io/qdgh114h7kwz) — 1h 52m
- `cloudflare` [Elevated error rates for Stream Live LL-HLS](https://www.cloudflarestatus.com/incidents/pvvq3811tpyq) — 1h 1m
- `cloudflare` [Network Performance Issues in San Jose (SJC-A)](https://www.cloudflarestatus.com/incidents/g4vvd0xrc2m8) — 1h 0m
- `openai` [Elevated Error Rates for ChatGPT Work, across Plus, Pro, Business, Enterprise and Education plans.](https://status.openai.com/incidents/01M34TAFM9WBGEBE3G1GBH4SE8) — 54m
- `stripe` [Elevated API errors](https://stspg.io/zxkgb0lnrxc8) — 33m
- `cloudflare` [Issues with 1.1.1.1 for Families](https://www.cloudflarestatus.com/incidents/rcjbfb0lyn65) — 43m
- `cloudflare` [Unable to start containers in Asia-Pacific](https://www.cloudflarestatus.com/incidents/rxmmx07ys7qg) — 41m
- `openai` [Elevated Error Rates on GPT-6 Astra Pro](https://status.openai.com/incidents/01M3ACKDE4GYY67FXXNRRSC0CE) — 46m
- `cloudflare` [Cloudflare Cloud Access Security Broker (CASB) Issues](https://www.cloudflarestatus.com/incidents/39dzwkhzcvv5) — 37m
- `openai` [Elevated Error Rates for ChatGPT across Plus and Pro plans.](https://status.openai.com/incidents/01M36RMC01ZFWQ861WYJKC4XE1) — 41m
- `cloudflare` [Elevated number of R2 503 errors in Australian Eastern Coast region](https://www.cloudflarestatus.com/incidents/x86v1qk8yp5v) — 17m
- `openai` [Increased error rate for Plus and Pro users.](https://status.openai.com/incidents/01M348TX2DX7KPM9X857N4APNV) — 39m
- `grafana` [Cloud Provider Observability - Hosted Azure metrics issues](https://stspg.io/y9pvmm8jfz0c) — 48m
- `circleci` [Delays starting Linux Machine jobs](https://stspg.io/8gpc503nlf95) — 47m
- `datadog` [Delayed Monitors Notifications](https://stspg.io/k1zlzpm9dfm2) — 46m
- `openai` [Mobile users unable to see Work Mode and the Model Picker](https://status.openai.com/incidents/01M389CDQ97B11QPMSRAAYSE5S) — 24m
- `datadog` [Users are unable to acknowledge, escalate, or resolve On-Call Pages](https://stspg.io/03h92q1l96z2) — 27m
- `cloudflare` [Unable to register certain domains](https://www.cloudflarestatus.com/incidents/lb5514cvngd7) — 11m
- `cloudflare` [Customers using BYOIP can have issues updating their BGP prefixes, including advertising or withdrawing prefixes.](https://www.cloudflarestatus.com/incidents/gwrzznc9f642) — 3m
- `cloudflare` [Increased Cache Failures](https://www.cloudflarestatus.com/incidents/mrqb4d6q0y26) — 0m
- `grafana` [Mimir write request errors](https://stspg.io/j6zm954jcg7m) — unknown
- `cloudflare` [Connectivity issues in Los Angeles (LAX)](https://www.cloudflarestatus.com/incidents/z84lp04f9sht) — 0m
- `cloudflare` [Network Performance Issues in Taipei](https://www.cloudflarestatus.com/incidents/zpkwd9fxsby5) — 0m
- `cloudflare` [Secondary DNS Update Delays](https://www.cloudflarestatus.com/incidents/7dt6f0chnclh) — unknown
- `cloudflare` [R2 Service Issues in Western North America (WNAM)](https://www.cloudflarestatus.com/incidents/6m2w2ls1gv75) — unknown

## Identity and evidence

Source record paths refer to the preserved vendor archive. Unconfirmed candidates remain separate.

<a id="ci-1f6486ca51776216be85a1cf"></a>

- Canonical ID: `ci-1f6486ca51776216be85a1cf` (1 source records).
  - [aws/f5f654e733c72f13](../../data/2026/W39/aws.json); vendor URL: https://health.aws.amazon.com/health/status
- Start time basis: `first_observed_update`.
- Warning: AWS onset is the first observed message, not a verified impact start

<a id="ci-39234b54d251a8c74d44a8ac"></a>

- Canonical ID: `ci-39234b54d251a8c74d44a8ac` (1 source records).
  - [cloudflare/tl0vrgjy99r0](../../data/2026/W39/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/tl0vrgjy99r0
- Start time basis: `vendor_reported`.

<a id="ci-cea3259e86f3208c2a946dae"></a>

- Canonical ID: `ci-cea3259e86f3208c2a946dae` (1 source records).
  - [cloudflare/8wmvkv5jkf15](../../data/2026/W39/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/8wmvkv5jkf15
- Start time basis: `vendor_reported`.

<a id="ci-04770eb84a28e9bd42bc0367"></a>

- Canonical ID: `ci-04770eb84a28e9bd42bc0367` (1 source records).
  - [github/8zc63m64hy36](../../data/2026/W39/github.json); vendor URL: https://stspg.io/m1ps7yrhp4n8
- Start time basis: `vendor_reported`.

<a id="ci-68070bea63cd77f508eda034"></a>

- Canonical ID: `ci-68070bea63cd77f508eda034` (1 source records).
  - [cloudflare/gv29fd076q2t](../../data/2026/W39/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/gv29fd076q2t
- Start time basis: `vendor_reported`.

<a id="ci-434bb27d04b018913a1fff67"></a>

- Canonical ID: `ci-434bb27d04b018913a1fff67` (1 source records).
  - [cloudflare/jg6ys8sbd6ny](../../data/2026/W39/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/jg6ys8sbd6ny
- Start time basis: `vendor_reported`.

<a id="ci-7a823295b5dda7b73729aa90"></a>

- Canonical ID: `ci-7a823295b5dda7b73729aa90` (1 source records).
  - [grafana/vc8dbf0jf1ds](../../data/2026/W39/grafana.json); vendor URL: https://stspg.io/cqfnrz8bdkvz
- Start time basis: `vendor_reported`.

<a id="ci-b91802663877f3ca7681de43"></a>

- Canonical ID: `ci-b91802663877f3ca7681de43` (1 source records).
  - [grafana/gjqvmydq1d56](../../data/2026/W39/grafana.json); vendor URL: https://stspg.io/4ntjl2yr75n4
- Start time basis: `vendor_reported`.

<a id="ci-90caeb635c9c77d87523f942"></a>

- Canonical ID: `ci-90caeb635c9c77d87523f942` (1 source records).
  - [anthropic/7g1qpkyz5gxh](../../data/2026/W39/anthropic.json); vendor URL: https://stspg.io/90s4l9yywzbk
- Start time basis: `vendor_reported`.

<a id="ci-92a8bc8535ec4a3a450f1af8"></a>

- Canonical ID: `ci-92a8bc8535ec4a3a450f1af8` (1 source records).
  - [cloudflare/mltthh0jxxpz](../../data/2026/W39/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/mltthh0jxxpz
- Start time basis: `vendor_reported`.

<a id="ci-cf225c9cb7eabbc58dc7d19b"></a>

- Canonical ID: `ci-cf225c9cb7eabbc58dc7d19b` (1 source records).
  - [cloudflare/dwfzz0bwhmsq](../../data/2026/W39/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/dwfzz0bwhmsq
- Start time basis: `vendor_reported`.

<a id="ci-3f3eb5bd4ea507cb1460c2e6"></a>

- Canonical ID: `ci-3f3eb5bd4ea507cb1460c2e6` (1 source records).
  - [cloudflare/dcmwvjwfyhjg](../../data/2026/W39/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/dcmwvjwfyhjg
- Start time basis: `vendor_reported`.

<a id="ci-02e842160015b8bcf9b815d1"></a>

- Canonical ID: `ci-02e842160015b8bcf9b815d1` (1 source records).
  - [github/1dk955gg3bvz](../../data/2026/W39/github.json); vendor URL: https://stspg.io/g2q4g5kv9frp
- Start time basis: `vendor_reported`.

<a id="ci-ae494d7db7bd8e14189afad3"></a>

- Canonical ID: `ci-ae494d7db7bd8e14189afad3` (1 source records).
  - [cloudflare/xbgbb80yw0jh](../../data/2026/W39/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/xbgbb80yw0jh
- Start time basis: `vendor_reported`.

<a id="ci-48c078e62ca2aae8f3df281c"></a>

- Canonical ID: `ci-48c078e62ca2aae8f3df281c` (1 source records).
  - [stripe/tk5f1plzpwnm](../../data/2026/W39/stripe.json); vendor URL: https://stspg.io/vfb8m896t1fk
- Start time basis: `vendor_reported`.

<a id="ci-82858f1ad2f887f239826b5b"></a>

- Canonical ID: `ci-82858f1ad2f887f239826b5b` (1 source records).
  - [cloudflare/f8wg5htd3zz4](../../data/2026/W39/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/f8wg5htd3zz4
- Start time basis: `vendor_reported`.

<a id="ci-d9ce48beb56921d67c2bd118"></a>

- Canonical ID: `ci-d9ce48beb56921d67c2bd118` (1 source records).
  - [cloudflare/2x4qt2q2ds9x](../../data/2026/W39/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/2x4qt2q2ds9x
- Start time basis: `vendor_reported`.

<a id="ci-cc4a7ca8054a05537b396db1"></a>

- Canonical ID: `ci-cc4a7ca8054a05537b396db1` (1 source records).
  - [cloudflare/78c19w4xw4z2](../../data/2026/W39/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/78c19w4xw4z2
- Start time basis: `vendor_reported`.

<a id="ci-d6fc354e0f53373eb8affd77"></a>

- Canonical ID: `ci-d6fc354e0f53373eb8affd77` (1 source records).
  - [stripe/3g746nm5qf15](../../data/2026/W39/stripe.json); vendor URL: https://stspg.io/4kvs6zwpn12h
- Start time basis: `vendor_reported`.

<a id="ci-beecf9d0e839cc13effe8e3a"></a>

- Canonical ID: `ci-beecf9d0e839cc13effe8e3a` (1 source records).
  - [cloudflare/drq3pdx1f11y](../../data/2026/W39/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/drq3pdx1f11y
- Start time basis: `vendor_reported`.

<a id="ci-b20044d1367cd141d31f1f5e"></a>

- Canonical ID: `ci-b20044d1367cd141d31f1f5e` (1 source records).
  - [cloudflare/nhf2kr8q3nfw](../../data/2026/W39/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/nhf2kr8q3nfw
- Start time basis: `vendor_reported`.

<a id="ci-9e3cc31c69f331b3a57b45bb"></a>

- Canonical ID: `ci-9e3cc31c69f331b3a57b45bb` (1 source records).
  - [cloudflare/c2f7zj7fs8l8](../../data/2026/W39/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/c2f7zj7fs8l8
- Start time basis: `vendor_reported`.

<a id="ci-eb5220d57590228e3ea3c5bc"></a>

- Canonical ID: `ci-eb5220d57590228e3ea3c5bc` (1 source records).
  - [circleci/zxqlp0926vcz](../../data/2026/W39/circleci.json); vendor URL: https://stspg.io/qdgh114h7kwz
- Start time basis: `vendor_reported`.

<a id="ci-f6e6cbb183b7a3152f8f44a2"></a>

- Canonical ID: `ci-f6e6cbb183b7a3152f8f44a2` (1 source records).
  - [cloudflare/pvvq3811tpyq](../../data/2026/W39/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/pvvq3811tpyq
- Start time basis: `vendor_reported`.

<a id="ci-1b4455036694b9bb7cab91fa"></a>

- Canonical ID: `ci-1b4455036694b9bb7cab91fa` (1 source records).
  - [cloudflare/g4vvd0xrc2m8](../../data/2026/W39/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/g4vvd0xrc2m8
- Start time basis: `vendor_reported`.

<a id="ci-3109e9afbc523c7d259c894a"></a>

- Canonical ID: `ci-3109e9afbc523c7d259c894a` (1 source records).
  - [openai/01M3DCNWMW57HYK8FJ5FBFPA39](../../data/2026/W39/openai.json); vendor URL: https://status.openai.com/incidents/01M3DCNWMW57HYK8FJ5FBFPA39
- Start time basis: `vendor_reported`.

<a id="ci-7c716591639daa3cda27a4a5"></a>

- Canonical ID: `ci-7c716591639daa3cda27a4a5` (1 source records).
  - [openai/01M34TAFM9WBGEBE3G1GBH4SE8](../../data/2026/W39/openai.json); vendor URL: https://status.openai.com/incidents/01M34TAFM9WBGEBE3G1GBH4SE8
- Start time basis: `vendor_reported`.

<a id="ci-86f3451565d9315d21a3d183"></a>

- Canonical ID: `ci-86f3451565d9315d21a3d183` (1 source records).
  - [stripe/45t05tqxy5w7](../../data/2026/W39/stripe.json); vendor URL: https://stspg.io/zxkgb0lnrxc8
- Start time basis: `vendor_reported`.

<a id="ci-2b7d25f33ad35705f6922e4f"></a>

- Canonical ID: `ci-2b7d25f33ad35705f6922e4f` (1 source records).
  - [cloudflare/rcjbfb0lyn65](../../data/2026/W39/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/rcjbfb0lyn65
- Start time basis: `vendor_reported`.

<a id="ci-256be7c2987090e82a9dff18"></a>

- Canonical ID: `ci-256be7c2987090e82a9dff18` (1 source records).
  - [cloudflare/rxmmx07ys7qg](../../data/2026/W39/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/rxmmx07ys7qg
- Start time basis: `vendor_reported`.

<a id="ci-4be69ec19a2b8dfb27cba7b4"></a>

- Canonical ID: `ci-4be69ec19a2b8dfb27cba7b4` (1 source records).
  - [openai/01M3ACKDE4GYY67FXXNRRSC0CE](../../data/2026/W39/openai.json); vendor URL: https://status.openai.com/incidents/01M3ACKDE4GYY67FXXNRRSC0CE
- Start time basis: `vendor_reported`.

<a id="ci-7415fc0f8ee5b1eee3b4bed8"></a>

- Canonical ID: `ci-7415fc0f8ee5b1eee3b4bed8` (1 source records).
  - [cloudflare/39dzwkhzcvv5](../../data/2026/W39/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/39dzwkhzcvv5
- Start time basis: `vendor_reported`.

<a id="ci-4c683c57754556fd042015ab"></a>

- Canonical ID: `ci-4c683c57754556fd042015ab` (1 source records).
  - [openai/01M36RMC01ZFWQ861WYJKC4XE1](../../data/2026/W39/openai.json); vendor URL: https://status.openai.com/incidents/01M36RMC01ZFWQ861WYJKC4XE1
- Start time basis: `vendor_reported`.

<a id="ci-ab29801c581871ef1a3d0b1d"></a>

- Canonical ID: `ci-ab29801c581871ef1a3d0b1d` (1 source records).
  - [cloudflare/x86v1qk8yp5v](../../data/2026/W39/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/x86v1qk8yp5v
- Start time basis: `vendor_reported`.

<a id="ci-17e2262c2ba09557db77afda"></a>

- Canonical ID: `ci-17e2262c2ba09557db77afda` (1 source records).
  - [openai/01M348TX2DX7KPM9X857N4APNV](../../data/2026/W39/openai.json); vendor URL: https://status.openai.com/incidents/01M348TX2DX7KPM9X857N4APNV
- Start time basis: `vendor_reported`.

<a id="ci-6c7dfb26d647dc24bd933baa"></a>

- Canonical ID: `ci-6c7dfb26d647dc24bd933baa` (1 source records).
  - [grafana/bdf4jlb79sp7](../../data/2026/W39/grafana.json); vendor URL: https://stspg.io/y9pvmm8jfz0c
- Start time basis: `vendor_reported`.

<a id="ci-4b56232a2369f423b27b5193"></a>

- Canonical ID: `ci-4b56232a2369f423b27b5193` (1 source records).
  - [circleci/3kx4rv23998z](../../data/2026/W39/circleci.json); vendor URL: https://stspg.io/8gpc503nlf95
- Start time basis: `vendor_reported`.

<a id="ci-d86f2ac1fc8227f2b16d93a6"></a>

- Canonical ID: `ci-d86f2ac1fc8227f2b16d93a6` (1 source records).
  - [datadog/g67lc6tmjnlt](../../data/2026/W39/datadog.json); vendor URL: https://stspg.io/k1zlzpm9dfm2
- Start time basis: `vendor_reported`.

<a id="ci-c25a69d9c1311c50db50b4d0"></a>

- Canonical ID: `ci-c25a69d9c1311c50db50b4d0` (1 source records).
  - [stripe/7lr9lwxqj3c0](../../data/2026/W39/stripe.json); vendor URL: https://stspg.io/4ntbqnhzg7bb
- Start time basis: `vendor_reported`.

<a id="ci-b56cbf1ec0b991fc55e9f755"></a>

- Canonical ID: `ci-b56cbf1ec0b991fc55e9f755` (1 source records).
  - [openai/01M389CDQ97B11QPMSRAAYSE5S](../../data/2026/W39/openai.json); vendor URL: https://status.openai.com/incidents/01M389CDQ97B11QPMSRAAYSE5S
- Start time basis: `vendor_reported`.

<a id="ci-b2c7a365178c52db4bfae3db"></a>

- Canonical ID: `ci-b2c7a365178c52db4bfae3db` (1 source records).
  - [datadog/cqxdntdk4c7d](../../data/2026/W39/datadog.json); vendor URL: https://stspg.io/03h92q1l96z2
- Start time basis: `vendor_reported`.

<a id="ci-0094db6a09b88f589e573e08"></a>

- Canonical ID: `ci-0094db6a09b88f589e573e08` (1 source records).
  - [cloudflare/lb5514cvngd7](../../data/2026/W39/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/lb5514cvngd7
- Start time basis: `vendor_reported`.

<a id="ci-ff19559d1d5f024384ea97e1"></a>

- Canonical ID: `ci-ff19559d1d5f024384ea97e1` (1 source records).
  - [cloudflare/gwrzznc9f642](../../data/2026/W39/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/gwrzznc9f642
- Start time basis: `vendor_reported`.

<a id="ci-443ede6b2aded893b384448a"></a>

- Canonical ID: `ci-443ede6b2aded893b384448a` (1 source records).
  - [cloudflare/mrqb4d6q0y26](../../data/2026/W39/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/mrqb4d6q0y26
- Start time basis: `vendor_reported`.

<a id="ci-df9f47e1f86195d8fea46c1b"></a>

- Canonical ID: `ci-df9f47e1f86195d8fea46c1b` (1 source records).
  - [grafana/gtdjr6zl1jp0](../../data/2026/W39/grafana.json); vendor URL: https://stspg.io/j6zm954jcg7m
- Start time basis: `vendor_reported`.
- Warning: Vendor resolved_at precedes started_at; duration is unknown, score uses zero duration; decision:legacy-time-grafana-gtdjr6zl1jp0

<a id="ci-99e437da00a2a4106b6872d1"></a>

- Canonical ID: `ci-99e437da00a2a4106b6872d1` (1 source records).
  - [cloudflare/z84lp04f9sht](../../data/2026/W39/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/z84lp04f9sht
- Start time basis: `vendor_reported`.

<a id="ci-0247d85f6a2b03021f713b07"></a>

- Canonical ID: `ci-0247d85f6a2b03021f713b07` (1 source records).
  - [cloudflare/zpkwd9fxsby5](../../data/2026/W39/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/zpkwd9fxsby5
- Start time basis: `vendor_reported`.

<a id="ci-e1bdfa64817e8102c8f6d655"></a>

- Canonical ID: `ci-e1bdfa64817e8102c8f6d655` (1 source records).
  - [cloudflare/7dt6f0chnclh](../../data/2026/W39/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/7dt6f0chnclh
- Start time basis: `vendor_reported`.
- Warning: Vendor resolved_at precedes started_at; duration is unknown, score uses zero duration; decision:legacy-time-cloudflare-7dt6f0chnclh

<a id="ci-ede9178a6bde81677f5a36ab"></a>

- Canonical ID: `ci-ede9178a6bde81677f5a36ab` (1 source records).
  - [cloudflare/6m2w2ls1gv75](../../data/2026/W39/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/6m2w2ls1gv75
- Start time basis: `vendor_reported`.
- Warning: Vendor resolved_at precedes started_at; duration is unknown, score uses zero duration; decision:legacy-time-cloudflare-6m2w2ls1gv75

## Sources

| Vendor | Status | Source records seen | Parse errors |
|--------|--------|----------------|--------------|
| anthropic | ok | 50 | 0 |
| atlassian | ok | 34 | 0 |
| aws | ok | 6 | 0 |
| circleci | ok | 50 | 0 |
| cloudflare | ok | 50 | 0 |
| datadog | FETCH FAILED: GET https://status.datadoghq.com/api/v2/incidents.json: status 403 | 0 | 0 |
| digitalocean | ok | 50 | 0 |
| gcp | ok | 8 | 0 |
| github | ok | 50 | 0 |
| grafana | ok | 50 | 0 |
| openai | ok | 25 | 0 |
| slack | ok | 28 | 0 |
| stripe | ok | 50 | 0 |
| yandexcloud | FETCH FAILED: GET https://status.cloud.yandex.ru/api/incidents?lang=en&page=4: status 403 | 0 | 0 |
