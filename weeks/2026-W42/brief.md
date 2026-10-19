# Week 2026-W42 — top 5 of 31 above threshold

2026-10-12 to 2026-10-18 (UTC).

## 1. grafana — Phone call notifications to numbers in Russia and Kazakhstan are blocked

- URL: https://stspg.io/3kqkpqlqm6f5
- Window: 2026-10-13 09:23 to not yet (6d 4h 45m (ongoing))
- Impact: minor, status: investigating
- Components: Incident Management and Response (IRM)

  - `2026-10-13 09:23` investigating — Phone call notifications to numbers in Russia and Kazakhstan are currently blocked. We are working to resolve the issue. As SMS support is limited in these countries, we recommend using mobile push notifications as an alternative. We do not currently have an estimated resolution time. We will provide further updates as more information becomes available.
  - `2026-10-13 21:35` investigating — Phone call notifications to numbers in Russia and Kazakhstan are currently blocked. We are working to resolve the issue. As SMS support is limited in these countries, we recommend using mobile push notifications as an alternative. We do not currently have an estimated resolution time. We will provide further updates as more information becomes available.
  - `2026-10-16 19:44` investigating — Phone call notifications to numbers in Russia and Kazakhstan are still blocked. We are working to resolve the issue. As SMS support is limited in these countries, we recommend using mobile push notifications as an alternative. We do not currently have an estimated resolution time. We will provide further updates as more information becomes available.

## 2. cloudflare — Audit Logs v2 ingestion is missing multiple Cloudflare products

- URL: https://www.cloudflarestatus.com/incidents/hpmrg8d5vwr6
- Window: 2026-10-13 10:14 to 2026-10-16 00:36 UTC (2d 14h 21m)
- Impact: minor, status: resolved
- Components: Audit Logs

  - `2026-10-13 10:14` identified — Since 2026-10-12 at 16:57:24 UTC, audit logs from multiple Cloudflare products are not ingested into the audit logs v2 database anymore. V1 audit logs are not affected
  - `2026-10-13 11:45` monitoring — We are continuing to monitor for any further issues.
  - `2026-10-16 00:36` resolved — This incident has been resolved.

## 3. openai — Delayed Costs data in the Compliance API

- URL: https://status.openai.com/incidents/01M4XNQX7C6RCBNGX9H711WFSG
- Window: 2026-10-14 17:00 to 2026-10-16 10:44 UTC (1d 17h 44m)
- Impact: minor, status: resolved
- Components: —

  - `2026-10-14 17:00` identified — We have identified an issue causing delays in publishing organization cost data through the Costs API. API and ChatGPT usage are not affected, and other reporting channels remain available. We are working to restore normal data freshness.
  - `2026-10-14 17:02` identified — We have identified an issue causing delays in publishing organization cost data through the Costs API. API and ChatGPT usage are not affected, and other reporting channels remain available. We are working to restore normal data freshness.
  - `2026-10-14 20:00` identified — We have identified an issue causing delays in publishing organization cost data through the Costs API. A fix is in progress and currently undergoing testing. API and ChatGPT usage are not affected, and other reporting channels remain available. We are continuing work to restore normal data freshness.
  - `2026-10-15 22:36` monitoring — We have deployed a fix and are monitoring recovery. Costs data in the Compliance API remains delayed while we process the backlog. We will provide further updates as processing catches up.
  - `2026-10-16 10:44` resolved — All impacted services have now fully recovered.

## 4. cloudflare — Elevated errors in Los Angeles (LAX)

- URL: https://www.cloudflarestatus.com/incidents/vyy1a1dspusg
- Window: 2026-10-12 08:52 to 2026-10-13 17:40 UTC (1d 8h 48m)
- Impact: minor, status: resolved
- Components: Los Angeles, CA, United States - (LAX)

  - `2026-10-12 08:52` investigating — Customers reaching Los Angeles (LAX) colo may see an elevate number of errors. Our team is investigating and working on mitigating the issue.
  - `2026-10-12 21:41` identified — We are continuing to work on a fix for the issue.
  - `2026-10-13 17:40` resolved — This incident has been resolved.

## 5. cloudflare — Cloudflare Workers - Cron Trigger Activation Issues

- URL: https://www.cloudflarestatus.com/incidents/npjny9agana8
- Window: 2026-10-14 05:36 to 2026-10-15 01:29 UTC (19h 53m)
- Impact: minor, status: resolved
- Components: Workers

  - `2026-10-14 05:36` investigating — Cloudflare is investigating issues with Workers Cron Triggers. Customers who have recently created or modified a Cron Trigger schedule may find that the schedule is not executing as expected. Pre-existing Cron Trigger schedules that have not been recently modified are not affected. We are working to identify the cause and mitigate this problem. More updates to follow shortly.
  - `2026-10-14 07:23` identified — A fix has been identified and is underway
  - `2026-10-14 08:14` monitoring — We are continuing to monitor for any further issues.
  - `2026-10-15 01:29` resolved — Workers Cron Trigger backlogs have cleared and appear healthy. Cron Trigger schedules should be executing as expected. Some schedule redeploys may be required for a small fraction of Workers who were impacted during the incident.

## Also above threshold (26)

- `github` Incident with Actions, Pull Requests and Issues (7h 54m) — https://stspg.io/flldq0dn2wdz
- `cloudflare` Increased HTTP 502 Errors (9h 51m) — https://www.cloudflarestatus.com/incidents/mxthd6onaplz
- `cloudflare` Increased Errors for Durable Objects (2h 23m) — https://www.cloudflarestatus.com/incidents/9yl8m5h16g1z
- `cloudflare` Network performance issues in Ashburn, US (IAD) (10h 56m) — https://www.cloudflarestatus.com/incidents/gqm2xrm9wk8t
- `cloudflare` Cloudflare Workers - Cron Trigger Activation Issues (8h 51m) — https://www.cloudflarestatus.com/incidents/cbvocc87nsne
- `cloudflare` Network Performance Issues in Seattle (8h 12m) — https://www.cloudflarestatus.com/incidents/4kaixilcb426
- `cloudflare` Cloudflare Network Performance Issues (7h 27m) — https://www.cloudflarestatus.com/incidents/ja8aolszxtrg
- `openai` Degraded performance in GPT-6.1 Sol Codex and Work mode (8h 26m) — https://status.openai.com/incidents/01M4T4ZCM54D3ZPTQXRSWFG0XS
- `cloudflare` Pages.dev sites experiencing availability issues in Norway (4h 43m) — https://www.cloudflarestatus.com/incidents/5py8l9ees6di
- `grafana` Elevated Errors Ingesting Metrics in AWS US East (VA) (6h 56m) — https://stspg.io/q5r4k8bnw6l6
- `cloudflare` Cloudflare Workers - Cron Trigger Activation Issues (3h 10m) — https://www.cloudflarestatus.com/incidents/mbv1dfp471bh
- `grafana` Tempo Write Failures in prod-us-central-0 (5h 42m) — https://stspg.io/0w7m43wy213v
- `cloudflare` DNS Update Delays (2h 45m) — https://www.cloudflarestatus.com/incidents/wpq9v3pzhcwl
- `grafana` Metric reads and rule evaluations are degraded in mimir-prod-66 (4h 49m) — https://stspg.io/489nw4cdzhsw
- `cloudflare` Network Performance Issues in Eastern North America (49m) — https://www.cloudflarestatus.com/incidents/fiv2bu3zf8kx
- `cloudflare` Cloudflare Workers AI Issues (2h 26m) — https://www.cloudflarestatus.com/incidents/alzmexlqangd
- `cloudflare` Cloudflare Fedramp Dashboard and Access (1h 12m) — https://www.cloudflarestatus.com/incidents/qwx5xtl8mv7d
- `openai` Degraded performance in Codex and ChatGPT Work (2h 47m) — https://status.openai.com/incidents/01M4WKPE1EY48ZB4RZJ80YH7ES
- `cloudflare` Network Performance Issues in LAX (2h 6m) — https://www.cloudflarestatus.com/incidents/fdgz8zlmfhw6
- `cloudflare` Elevated number of 5XX errors in PDX (2h 5m) — https://www.cloudflarestatus.com/incidents/7hicc5bfkvx3
- `openai` Widespread feature degradation (2h 26m) — https://status.openai.com/incidents/01M506BA4XRSWH7THR53HQGN3Z
- `grafana` Authentication errors in prod-us-west-0 (2h 48m) — https://stspg.io/wrxbbbsyx0f7
- `stripe` Elevated declines on Mastercard (2h 18m) — https://stspg.io/csnmnfc5cgjv
- `grafana` Tempo: High Latency in prod-eu-west-2 (2h 23m) — https://stspg.io/qth8ndlqhdj1
- `cloudflare` Network Performance Issues in London (13m) — https://www.cloudflarestatus.com/incidents/ybyy2whvd8ef
- `cloudflare` Cloudflare Service Issues (28m) — https://www.cloudflarestatus.com/incidents/z7vongh8rqc1

