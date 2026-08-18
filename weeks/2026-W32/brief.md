# Week 2026-W32 — top 5 of 24 above threshold

2026-08-03 to 2026-08-09 (UTC).

## 1. cloudflare — R2 Availability Issues

- URL: https://www.cloudflarestatus.com/incidents/9fzfjzc5ms0h
- Window: 2026-08-07 18:42 to 2026-08-09 02:35 UTC (1d 7h 53m)
- Impact: minor, status: resolved
- Components: R2

  - `2026-08-07 18:42` investigating — Cloudflare is aware, and investigating an issue affecting availability in ENAM for a small number of R2 buckets.
  - `2026-08-07 22:54` identified — We have identified the cause and mitigated further impact. Writes to a small set of buckets in ENAM were impacted between 14:52-17:02 UTC. We are continuing to work to restore normal availability for these buckets.
  - `2026-08-07 23:05` identified — We are continuing to work on a fix for this issue.
  - `2026-08-08 10:52` identified — We are continuing to restore availability for impacted R2 buckets in ENAM. Recovery is progressing well, and we will provide status updates as they becomes available
  - `2026-08-08 23:28` identified — Availability has been restored for nearly all impacted R2 buckets. A small number of objects uploaded through multipart uploads in ENAM between 14:52 and 17:02 UTC on August 7 remain unavailable. We are continuing to restore access and expect to be able to provide another update soon.
  - `2026-08-09 02:35` resolved — Availability has now been fully restored for all impacted R2 buckets, including the remaining objects uploaded through multipart uploads in ENAM.

## 2. anthropic — Degraded performance of multiple models

- URL: https://stspg.io/v7xpxd5xg5pc
- Window: 2026-08-05 07:05 to 2026-08-05 14:14 UTC (7h 8m)
- Impact: minor, status: resolved
- Components: claude.ai, Claude API (api.anthropic.com), Claude Code, Claude Cowork

  - `2026-08-05 07:05` identified — We have identified the cause of elevated errors on requests to Claude Mythos 5, Claude Fable 5, and Claude Opus 5, Claude Sonnet 5 and are working on a fix. We will provide an update as soon as possible.
  - `2026-08-05 09:13` identified — We are continuing to work on a fix for this issue.
  - `2026-08-05 13:08` monitoring — A fix has been deployed for the issue affecting Claude services and we are monitoring for recovery.
  - `2026-08-05 14:14` resolved — This incident has been resolved.

## 3. github — Incident with Actions

- URL: https://stspg.io/rcz3fcm83sff
- Window: 2026-08-06 15:22 to 2026-08-07 02:04 UTC (10h 41m)
- Impact: critical, status: resolved
- Components: Actions, Pages

  - `2026-08-06 15:22` investigating — We are investigating reports of degraded performance for Actions
  - `2026-08-06 15:41` investigating — Actions is experiencing degraded availability. We are continuing to investigate.
  - `2026-08-06 15:45` investigating — We are investigating errors affecting GitHub Actions. Some workflow runs are failing to start or failing partway through, and some requests to the Actions REST API are returning errors. <br /><br />Some customers may also see unexpected rate limiting in their workflows. <br /><br />Engineers have identified the source of the disruption and are actively working on a mitigation
  - `2026-08-06 15:53` investigating — Pages is experiencing degraded performance. We are continuing to investigate.
  - `2026-08-06 16:19` investigating — Pages is operating normally.
  - `2026-08-06 16:27` investigating — Pages is experiencing degraded performance. We are continuing to investigate.
  - `2026-08-06 16:27` investigating — We are continuing to work on the issue affecting GitHub Actions. <br /><br />Some workflow runs are still delayed or failing to complete, and some requests to the Actions API are returning errors. <br /><br />Customers running migrations with GitHub Enterprise Importer may also see failures. <br /><br />Engineers are actively working towards full recovery.
  - `2026-08-06 16:33` investigating — Actions and Pages are experiencing degraded availability. We are continuing to investigate.
  - `2026-08-06 17:02` investigating — We are continuing to work on the issue affecting GitHub Actions. <br /><br />Workflow runs are still failing or delayed in starting, and some queued jobs may time out. <br /><br />Some requests to the Actions API are returning errors. Customers running migrations with GitHub Enterprise Importer may see failures. <br /><br />Our engineers have applied several mitigations and are rolling out a further fix now.
  - `2026-08-06 17:40` investigating — We are continuing to work on an issue affecting multiple GitHub services.<br /><br />Workflow runs are failing or delayed in starting, and some queued jobs may time out. <br /><br />Copilot code review, Copilot coding agent, hosted runners, and migrations using GitHub Enterprise Importer might also affected. <br /><br />Webhook deliveries may be delayed. <br /><br />Engineers have applied a number of mitigations and are rolling out a further fix across all affected systems now.
  - `2026-08-06 18:11` investigating — We are continuing to work on an issue affecting multiple GitHub services. <br /><br />Workflow runs are still failing or delayed in starting, and some queued jobs may time out. <br /><br />Customers using self-hosted runners may see errors or rate limiting when runners register. <br /><br />Copilot code review, Copilot coding agent, hosted runners, and migrations using GitHub Enterprise Importer may also be affected. <br /><br />Webhook deliveries may be delayed.<br /><br />Engineers have applied further mitigations and are continuing to work towards full recovery.
  - `2026-08-06 18:46` investigating — We are continuing to work on an issue affecting multiple GitHub services. <br /><br />Workflow runs are still failing, and jobs may remain queued for an extended period before starting or may time out. Jobs using GitHub-hosted runners are particularly affected while capacity is constrained. <br /><br />Customers using self-hosted runners may see errors or rate limiting when runners register. <br /><br />Copilot code review, Copilot coding agent, and migrations using GitHub Enterprise Importer may also be affected. Webhook deliveries may be delayed. <br /><br />Recovery is taking longer than we expected, and engineers remain actively engaged.
  - `2026-08-06 19:43` investigating — We are continuing to work on an issue affecting GitHub Actions. <br /><br />Capacity remains constrained and jobs may still be delayed or fail while it recovers gradually. Customers using self-hosted runners may see errors or rate limiting when runners register. <br /><br />Copilot code review, Copilot coding agent, and migrations using GitHub Enterprise Importer may also be affected. Webhook deliveries may be delayed. <br /><br />Our engineers remain actively engaged.
  - `2026-08-06 20:34` investigating — We are continuing to work on an issue affecting GitHub Actions. Webhook triggers are currently throttled to help with recovery and and we are processing approximately 15% of webhooks, so many events such as pushes and pull requests are not triggering workflow runs. Of jobs queued, approximately 65% are succeeding, improved from a low of 30 to 40% earlier in this incident.<br /><br />We have narrowed the remaining impact to runners that are stuck retrying jobs that are no longer available. Both GitHub-hosted and self-hosted runners are affected, and we are working to recover them.<br /><br />Copilot code review, Copilot coding agent, and migrations using GitHub Enterprise Importer may also be affected.
  - `2026-08-06 21:30` investigating — We are continuing to work on an issue affecting GitHub Actions. Webhook triggers remain throttled to aid recovery, so many push and pull request events are not triggering new workflow runs.<br /><br />We identified runners being assigned jobs that are no longer valid and are deploying a change to address this issue. Both GitHub-hosted and self-hosted runners are affected.<br /><br />Copilot code review, Copilot coding agent, and GitHub Pages may experience failures or delays. Migrations using GitHub Enterprise Importer have been paused to support mitigation efforts.
  - `2026-08-06 22:18` investigating — We continue to make progress on the issue affecting GitHub Actions. We have deployed a fix that addresses runners being assigned jobs that are no longer valid, and are seeing improvement in job completion rates. For workflow runs that are starting, success rates have increased significantly and are now at 97%. Standard and larger runners are now draining queued work. A change is also in progress to mitigate issues with existing self-hosted runners that are not picking up jobs.<br /><br />Webhook triggers remain throttled to support recovery. Many push and pull request events are not yet triggering new workflow runs, and we are working to safely restore full throughput.<br /><br />GitHub Pages, Copilot code review, and Copilot coding agent may still experience failures or delays. Migrations…
  - `2026-08-06 23:13` investigating — We have deployed fixes that address runners being assigned invalid jobs and are taking additional steps to clear the backlog of affected jobs. Job completion rates for running workflows have improved significantly, with success rates now at 99%. Global queues for hosted runner assignment are nearly burned down and concurrency queues for customers are being processed. Another change was deployed to accelerate processing the backlog of job requests.<br /><br />We are gradually restoring throughput for webhook-triggered Actions workflows and monitoring system stability. We have deployed a fix for self-hosted runners that were not picking up jobs and are enabling it incrementally.<br /><br />GitHub Pages, Copilot code review, and Copilot coding agent may still experience intermittent failures…
  - `2026-08-07 00:01` investigating — <br />System-wide queues have been drained, and new jobs are being processed as expected. The fix for self-hosted runners not picking up jobs has been fully rolled out.<br /><br />Webhook-triggered Actions workflows have been restored to full throughput. GitHub Pages, Copilot code review, and Copilot coding agent are showing recovery. Migrations using GitHub Enterprise Importer remain paused as a precaution.<br /><br />We are monitoring all affected services for sustained recovery and will provide another update shortly.
  - `2026-08-07 00:01` investigating — System-wide queues have been drained, and new jobs are being processed as expected. The fix for self-hosted runners not picking up jobs has been fully rolled out.<br /><br />Webhook-triggered Actions workflows have been restored to full throughput. GitHub Pages, Copilot code review, and Copilot coding agent are showing recovery. Migrations using GitHub Enterprise Importer remain paused as a precaution.<br /><br />We are monitoring all affected services for sustained recovery and will provide another update shortly.
  - `2026-08-07 00:05` investigating — The degradation affecting Actions and Pages has been mitigated. We are monitoring to ensure stability.
  - `2026-08-07 00:06` monitoring — The degradation has been mitigated. We are monitoring to ensure stability.
  - `2026-08-07 00:59` monitoring — We’re investigating reports that some Actions Runner Controller runners are taking longer than expected to recover. We’ll provide an update as our investigation progresses.
  - `2026-08-07 02:03` monitoring — During the incident, some Actions Runner Controller (ARC) runner pods became stuck in an idle state. Affected users can delete those pods using kubectl or redeploy their Actions Runner Controller application. ARC will automatically create replacement runners.<br /><br />The next releases of Actions Runner and Actions Runner Controller will include an automatic recovery mechanism, preventing the need for these manual steps in the future.<br /><br />Some workflow-triggering events, including push and pull request events, were not processed during the incident and cannot be replayed automatically. Customers may need to repeat the triggering action by pushing a new commit, updating the pull request, or manually re-running the workflow where applicable.
  - `2026-08-07 02:04` resolved — On August 6, 2026, between 15:05 UTC and 00:14 UTC on August 7, GitHub Actions experienced degraded availability. During the incident, workflow runs failed or remained queued for an extended period of time. Customers using both GitHub-hosted and self-hosted runners were affected. At peak, 71% of workflow runs experienced infrastructure failures and 75% of the remaining workflow runs were delayed by more than 5 minutes. <br /><br />The incident was triggered by a routine deployment to an internal Actions service responsible for processing events and generating Actions jobs. The deployment exposed an existing capacity and concurrency weakness. As pods were replaced during the deployment, remaining capacity became saturated, causing services to crash and triggering a cascading impact across m…

## 4. cloudflare — Cloudflare Workers Runtime API issues

- URL: https://www.cloudflarestatus.com/incidents/238b69fw6l55
- Window: 2026-08-03 15:13 to 2026-08-04 07:12 UTC (15h 59m)
- Impact: minor, status: resolved
- Components: Workers

  - `2026-08-03 15:13` investigating — Cloudflare is aware of an issue with the latest Workers Runtime API release that exposes native Temporal with clock stuck at epoch 0. Updates to follow.
  - `2026-08-03 15:21` identified — Workers runtime: Temporal global exposed with an incorrect clock Since 2026-07-30, the Workers runtime has exposed a global Temporal object that was not intended to be available. Its clock is wrong: Temporal.Now reports 1970-01-01 rather than the current time. Date and Date.now() are unaffected and report the correct time. This can affect Workers that install a Temporal polyfill only when no native Temporal is detected. On those Workers, the polyfill stopped being installed and time-dependent logic began computing against 1970 — for example, tokens minted with an issue time of 0 and an expiry in 1970, incorrect TTLs, or incorrect date arithmetic. No error is raised when this happens.
  - `2026-08-04 07:00` monitoring — A fix has been implemented and we are monitoring the results.
  - `2026-08-04 07:12` resolved — This incident has been resolved.

## 5. cloudflare — Network Performance Issues in Istanbul

- URL: https://www.cloudflarestatus.com/incidents/x44ny80374v7
- Window: 2026-08-08 08:25 to 2026-08-08 19:45 UTC (11h 19m)
- Impact: minor, status: resolved
- Components: Network

  - `2026-08-08 08:25` investigating — We have lost dark fiber in the Istanbul region, we are working to follow up with our vendors to restore connectivity. More updates to follow shortly.
  - `2026-08-08 19:45` resolved — This incident has been resolved.

## Also above threshold (19)

- `digitalocean` Account Registration, Droplets, and Related Services (6h 3m) — https://stspg.io/grfh9nf5bp3f
- `grafana` Some Grafana Instances Unavailable (6h 20m) — https://stspg.io/4jyq27wvr4kk
- `openai` Elevated errors with image generation (8h 26m) — https://status.openai.com/incidents/01KZ96S5HEWX0CX26W2WFYFKE6
- `openai` Issues with Custom GPT actions (5h 53m) — https://status.openai.com/incidents/01KZ9DMQD2GJ8JJWDN7572RH78
- `digitalocean` Spaces Cold Storage Billing (5h 29m) — https://stspg.io/qb6zv0kdjy50
- `circleci` Increased Job Queue Times: Windows, Android, and GPU (4h 12m) — https://stspg.io/525771jb0gj4
- `circleci` Usage API data delayed for 8/5 (6h 48m) — https://stspg.io/kntq9cf4jgpk
- `cloudflare` Connectivity issues affecting access to some website from certain networks in Egypt (2h 53m) — https://www.cloudflarestatus.com/incidents/ydgyl9v3g02s
- `openai` Elevated errors in ChatGPT conversations with files (3h 13m) — https://status.openai.com/incidents/01KZ9E7QD8WFJD5DMF73QGVYH8
- `cloudflare` Cloudflare Dedicated Egress in London (2h 17m) — https://www.cloudflarestatus.com/incidents/7dk3g4k188ky
- `circleci` Insights service data is lagging (4h 2m) — https://stspg.io/6qdzqdj5ww81
- `openai` Elevated ChatGPT conversation errors affecting Plus, Pro, Business, and Edu users (2h 41m) — https://status.openai.com/incidents/01KZ6CT9K9Q52S3707W75MNBGB
- `cloudflare` Elevated R2 error rates and latency (2h 6m) — https://www.cloudflarestatus.com/incidents/fkpqy4zrm28t
- `cloudflare` Web Analytics Configuration issues (2h 2m) — https://www.cloudflarestatus.com/incidents/3wbkwvf5ypbk
- `slack` Trouble Using Search Bar For Some Admins (2h 45m) — https://slack-status.com/2026-08/1db1f200d420086a
- `digitalocean` Functions on App Platform (2h 3m) — https://stspg.io/7pm9l1jwpkr6
- `grafana` Logs latency increase within prod-eu-west-3 (3h 14m) — https://stspg.io/tf9jp15y4f5k
- `digitalocean` Cloud Control Panel Access (4h 1m) — https://stspg.io/045lx5xqfdns
- `anthropic` Elevated errors across many models (1h 11m) — https://stspg.io/j5npxk25yrp3

