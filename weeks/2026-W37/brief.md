# Week 2026-W37 — top 5 of 23 above threshold

2026-09-07 to 2026-09-13 (UTC).

## 1. grafana — Issues with Geomap Tiles

- URL: https://stspg.io/lr44wpz3m0tz
- Window: 2026-09-08 18:05 to 2026-09-08 22:41 UTC (4h 36m)
- Impact: minor, status: resolved
- Components: AWS Australia - prod-ap-southeast-2, AWS Brazil - prod-sa-east-1, AWS Canada - prod-ca-east-0, AWS Germany - prod-eu-west-2, AWS Germany - prod-eu-west-4, AWS India - prod-ap-south-1, AWS Japan - prod-ap-northeast-0, AWS UAE - prod-me-central-1, AWS Singapore - prod-ap-southeast-1, AWS Sweden - prod-eu-north-0, AWS US East - prod-us-east-0, AWS US East - prod-us-east-2, AWS US West - prod-us-west-0, AWS Australia - prod-au-southeast-1, AWS UK - prod-gb-south-1, AWS Ireland - prod-eu-west-6, Azure US Central - us-central2, AWS Switzerland - prod-eu-central-0, Azure Netherlands - prod-eu-west-3, GCP Australia - prod-au-southeast-0, GCP Belgium - prod-eu-west-0, GCP Brazil - prod-sa-east-0, GCP India - prod-ap-south-0, GCP Singapore - prod-ap-southeast-0, GCP UK - prod-gb-south-0, GCP US Central - prod-us-central-0, GCP US Central - prod-us-central-3, GCP US Central - prod-us-central-4, GCP US East - prod-us-east-1, play.grafana.org, Federal Cloud - AWS US Gov West, GCP Saudi Arabia - prod-me-central-0, AWS US East - prod-us-east-3

  - `2026-09-08 18:05` identified — Some geomap tiles are displaying the “api key required". The issue has been identified, and our team is working on a fix. As a workaround, you can open the Geomap panel, navigate to the Map layers section, and change the Basemap layer from CARTO, or "Default base layer," to OpenStreetMap. The watermark should disappear immediately. Alternatively, if you would prefer to retain a dark-styled map similar to the current CARTO appearance, you can select MapLibre as the basemap and configure the URL as: https://tiles.openfreemap.org/styles/dark This option also renders without the watermark, does not require an API key, and provides a visual style similar to the dark CARTO basemap.
  - `2026-09-08 22:41` resolved — This incident has been resolved.

## 2. anthropic — Degraded functionality for Claude Cowork on Windows

- URL: https://stspg.io/5yp3rhhztm30
- Window: 2026-09-10 15:54 to not yet (3d 21h 34m (ongoing))
- Impact: major, status: identified
- Components: Claude Cowork

  - `2026-09-10 15:54` identified — A Windows update released September 8 has left Claude Cowork on Windows unable to run local commands, because its workspace can no longer reach your computer's drive. For most users, chat and file reading and editing still work. There's no in-app workaround yet, and restarting or reinstalling won't help. Microsoft has developed a fix and is working to release it. We'll post an update when we have more to share.

## 3. cloudflare — Workers Cron Triggers degraded

- URL: https://www.cloudflarestatus.com/incidents/sjs8s0q2x4hw
- Window: 2026-09-09 19:17 to 2026-09-12 03:34 UTC (2d 8h 16m)
- Impact: minor, status: resolved
- Components: Workers

  - `2026-09-09 19:17` investigating — Workers Cron Triggers may not execute or may be delayed in executing. Updates to Workers Cron Triggers may take some time to take effect.
  - `2026-09-09 19:18` identified — The issue has been identified and a fix is being implemented.
  - `2026-09-12 02:25` monitoring — A fix has been implemented and we are monitoring the results.
  - `2026-09-12 03:34` resolved — This incident has been resolved.

## 4. cloudflare — Cloudflare Tunnel Availability Issues

- URL: https://www.cloudflarestatus.com/incidents/k7l4hzmthl3k
- Window: 2026-09-11 21:55 to 2026-09-12 08:57 UTC (11h 1m)
- Impact: minor, status: resolved
- Components: Tunnel

  - `2026-09-11 21:55` investigating — Cloudflare is investigating issues with Cloudflare Tunnel. More updates to follow shortly.
  - `2026-09-11 23:15` identified — The issue has been identified and a fix is being implemented.
  - `2026-09-12 03:49` identified — Mitigation is reducing the affected traffic, and recovery is progressing. Some existing Cloudflare Tunnel connections may continue to experience intermittent issues while this work completes. We continue to monitor progress.
  - `2026-09-12 08:42` monitoring — A fix has been implemented and we are monitoring the results.
  - `2026-09-12 08:57` resolved — This incident has been resolved.

## 5. openai — Unable to open shared ChatGPT Project using direct link

- URL: https://status.openai.com/incidents/01M24HYC9SX2Q4471EVMY6EXQE
- Window: 2026-09-10 02:21 to 2026-09-10 14:28 UTC (12h 6m)
- Impact: minor, status: resolved
- Components: —

  - `2026-09-10 02:21` investigating — We are investigating the issue for the listed services.
  - `2026-09-10 06:01` identified — We are still implementing the mitigation.
  - `2026-09-10 09:37` identified — We are still implementing the mitigation.
  - `2026-09-10 12:06` monitoring — We have applied the mitigation and are monitoring the recovery.
  - `2026-09-10 14:28` resolved — All impacted services have now fully recovered.

## Also above threshold (18)

- `cloudflare` Browser Isolation Session Initialization Failures (4h 49m) — https://www.cloudflarestatus.com/incidents/z16209cfb1xv
- `cloudflare` DNS Update Delays (8h 24m) — https://www.cloudflarestatus.com/incidents/nhfcmzd1pv4p
- `openai` 1% of ChatGPT Work (mobile/web) turns are failing for existing threads (9h 34m) — https://status.openai.com/incidents/01M28MEQWTQJDRCRPFD9FWQ0H3
- `stripe` Elevated Interac errors (11h 4m) — https://stspg.io/j5s6qs75yrjz
- `openai` Elevated errors for image generation (7h 26m) — https://status.openai.com/incidents/01M20PYYYGRT9303VHAPA7YNT2
- `github` Incident with several GitHub Services (1h 28m) — https://stspg.io/8f3xch4y0v2k
- `slack` Some customers experiencing blurry image previews and download issues (8h 55m) — https://slack-status.com/2026-09/f171324a097abb9a
- `cloudflare` Cloudflare Workers AI increased errors (5h 44m) — https://www.cloudflarestatus.com/incidents/jsvvrtk5hvym
- `cloudflare` Increased Workflow Instance Creation Errors (2h 23m) — https://www.cloudflarestatus.com/incidents/np4n61ckjdw7
- `openai` File uploads are delayed or failing (4h 36m) — https://status.openai.com/incidents/01M20QBACYZMK2PRCJQCBSWTYJ
- `cloudflare` Issues creating and updating Cloud Connector Rules (2h 36m) — https://www.cloudflarestatus.com/incidents/t4gby3dqzbg1
- `cloudflare` Network Performance Issues with Taiwan datacenter (TPE) (2h 28m) — https://www.cloudflarestatus.com/incidents/2njywrxj8qjf
- `grafana` High Latency in prod-us-east-2 (4h 19m) — https://stspg.io/f5m752d4l87f
- `openai` Elevated errors for ChatGPT users in Europe (2h 23m) — https://status.openai.com/incidents/01M27Q9GH7WVKNA61ZR3RDMSY5
- `openai` Elevated errors for GPT-5.6 Sol on the API (2h 13m) — https://status.openai.com/incidents/01M27VK1VN54RQJGJEHB5JTQJJ
- `slack` Trouble loading Workspace-level Apps & Workflows settings page is returning blank pages (2h 39m) — https://slack-status.com/2026-09/c8ec111fea759b78
- `grafana` Investigating elevated database load in AWS Germany (2h 24m) — https://stspg.io/j1r7zw96mgh5
- `anthropic` Elevated errors for Claude Mythos 5.1 and Claude Fable 5.1 (17m) — https://stspg.io/n9n8kfcxtkfm

