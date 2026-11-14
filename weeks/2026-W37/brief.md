# Week 2026-W37 — top 5 of 25 above threshold

2026-09-07 to 2026-09-13 (UTC).

46 source records (49 saved revisions) → 46 canonical incidents; 0 repeated records grouped; 0 candidate pairs pending review. Rules: `identity-v1`. Counts refer to collected vendor reports, not all real-world outages.

## 1. cloudflare — Hyperdrive Degraded Performance

- URL: https://www.cloudflarestatus.com/incidents/adw7e1heupa2
- Window: 2026-09-11 00:00 to 2026-10-05 00:00 UTC (24d 0h 0m)
- Impact: none, status: resolved
- Components: Hyperdrive

<a id="ci-74bd65ede6e8fc278492a918"></a>

- Canonical ID: `ci-74bd65ede6e8fc278492a918` (1 source records).
  - [cloudflare/adw7e1heupa2](../../data/2026/W37/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/adw7e1heupa2
- Start time basis: `vendor_reported`.

  - `2026-10-07 23:56` resolved — Cloudflare identified an issue that affected Hyperdrive performance between 11th September and 5th October 2026. During this time, customers may have experienced elevated latency. This incident has been resolved.

## 2. grafana — Issues with Geomap Tiles

- URL: https://stspg.io/lr44wpz3m0tz
- Window: 2026-09-08 18:05 to 2026-09-08 22:41 UTC (4h 36m)
- Impact: minor, status: resolved
- Components: AWS Australia - prod-au-southeast-1, AWS Australia - prod-au-southeast-2, AWS Brazil - prod-sa-east-1, AWS Canada - prod-ca-east-0, AWS Germany - prod-eu-west-2, AWS Germany - prod-eu-west-4, AWS India - prod-ap-south-1, AWS Ireland - prod-eu-west-6, AWS Japan - prod-ap-northeast-0, AWS Singapore - prod-ap-southeast-1, AWS Sweden - prod-eu-north-0, AWS Switzerland - prod-eu-central-0, AWS UAE - prod-me-central-1, AWS UK - prod-gb-south-1, AWS US East (OH) - prod-us-east-0, AWS US East (OH) - prod-us-east-2, AWS US East (VA) - prod-us-east-3, AWS US West - prod-us-west-0, Azure Netherlands - prod-eu-west-3, Azure US Central - us-central2, Federal Cloud - AWS US Gov West, GCP Australia - prod-au-southeast-0, GCP Belgium - prod-eu-west-0, GCP Brazil - prod-sa-east-0, GCP India - prod-ap-south-0, GCP Saudi Arabia - prod-me-central-0, GCP Singapore - prod-ap-southeast-0, GCP UK - prod-gb-south-0, GCP US Central - prod-us-central-0, GCP US Central - prod-us-central-3, GCP US Central - prod-us-central-4, GCP US East - prod-us-east-1, play.grafana.org

<a id="ci-8302dd689d389b0179980548"></a>

- Canonical ID: `ci-8302dd689d389b0179980548` (1 source records).
  - [grafana/nm6vck9btxjb](../../data/2026/W37/grafana.json); vendor URL: https://stspg.io/lr44wpz3m0tz
- Start time basis: `vendor_reported`.

  - `2026-09-08 18:05` identified — Some geomap tiles are displaying the “api key required". The issue has been identified, and our team is working on a fix. As a workaround, you can open the Geomap panel, navigate to the Map layers section, and change the Basemap layer from CARTO, or "Default base layer," to OpenStreetMap. The watermark should disappear immediately. Alternatively, if you would prefer to retain a dark-styled map similar to the current CARTO appearance, you can select MapLibre as the basemap and configure the URL as: https://tiles.openfreemap.org/styles/dark This option also renders without the watermark, does not require an API key, and provides a visual style similar to the dark CARTO basemap.
  - `2026-09-08 22:41` resolved — This incident has been resolved.

## 3. anthropic — Degraded functionality for Claude Cowork on Windows

- URL: https://stspg.io/5yp3rhhztm30
- Window: 2026-09-10 15:54 to 2026-09-14 19:14 UTC (4d 3h 20m)
- Impact: major, status: resolved
- Components: Claude Cowork

<a id="ci-43a9a49695dad394c58e93af"></a>

- Canonical ID: `ci-43a9a49695dad394c58e93af` (1 source records).
  - [anthropic/r1pqn1kb4hvk](../../data/2026/W37/anthropic.json); vendor URL: https://stspg.io/5yp3rhhztm30
- Start time basis: `vendor_reported`.

  - `2026-09-10 15:54` identified — A Windows update released September 8 has left Claude Cowork on Windows unable to run local commands, because its workspace can no longer reach your computer's drive. For most users, chat and file reading and editing still work. There's no in-app workaround yet, and restarting or reinstalling won't help. Microsoft has developed a fix and is working to release it. We'll post an update when we have more to share.
  - `2026-09-14 19:14` resolved — This issue has been resolved. Microsoft has released a September 14 Windows update that fixes it: https://support.microsoft.com/en-us/servicing/os/windows-11/2026/09/kb5129195-windows-11-24h2-25h2-security-update

## 4. cloudflare — Workers Cron Triggers degraded

- URL: https://www.cloudflarestatus.com/incidents/sjs8s0q2x4hw
- Window: 2026-09-09 19:17 to 2026-09-12 03:34 UTC (2d 8h 16m)
- Impact: minor, status: resolved
- Components: Workers

<a id="ci-2ed8c7d7199fd963f73eaa14"></a>

- Canonical ID: `ci-2ed8c7d7199fd963f73eaa14` (1 source records).
  - [cloudflare/sjs8s0q2x4hw](../../data/2026/W37/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/sjs8s0q2x4hw
- Start time basis: `vendor_reported`.

  - `2026-09-09 19:17` investigating — Workers Cron Triggers may not execute or may be delayed in executing. Updates to Workers Cron Triggers may take some time to take effect.
  - `2026-09-09 19:18` identified — The issue has been identified and a fix is being implemented.
  - `2026-09-12 02:25` monitoring — A fix has been implemented and we are monitoring the results.
  - `2026-09-12 03:34` resolved — This incident has been resolved.

## 5. cloudflare — Cloudflare Tunnel Availability Issues

- URL: https://www.cloudflarestatus.com/incidents/k7l4hzmthl3k
- Window: 2026-09-11 21:55 to 2026-09-12 08:57 UTC (11h 1m)
- Impact: minor, status: resolved
- Components: Tunnel

<a id="ci-04ef0ea10c5ad910cb8e2ee1"></a>

- Canonical ID: `ci-04ef0ea10c5ad910cb8e2ee1` (1 source records).
  - [cloudflare/k7l4hzmthl3k](../../data/2026/W37/cloudflare.json); vendor URL: https://www.cloudflarestatus.com/incidents/k7l4hzmthl3k
- Start time basis: `vendor_reported`.

  - `2026-09-11 21:55` investigating — Cloudflare is investigating issues with Cloudflare Tunnel. More updates to follow shortly.
  - `2026-09-11 23:15` identified — The issue has been identified and a fix is being implemented.
  - `2026-09-12 03:49` identified — Mitigation is reducing the affected traffic, and recovery is progressing. Some existing Cloudflare Tunnel connections may continue to experience intermittent issues while this work completes. We continue to monitor progress.
  - `2026-09-12 08:42` monitoring — A fix has been implemented and we are monitoring the results.
  - `2026-09-12 08:57` resolved — This incident has been resolved.

## Also above threshold (20)

- `openai` Unable to open shared ChatGPT Project using direct link (12h 6m) — https://status.openai.com/incidents/01M24HYC9SX2Q4471EVMY6EXQE
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
- `yandexcloud` Slowing down operations with access bindings (2h 22m) — https://status.yandex.cloud/en/incidents/2005
- `openai` Elevated errors for GPT-5.6 Sol on the API (2h 13m) — https://status.openai.com/incidents/01M27VK1VN54RQJGJEHB5JTQJJ
- `slack` Trouble loading Workspace-level Apps & Workflows settings page is returning blank pages (2h 39m) — https://slack-status.com/2026-09/c8ec111fea759b78
- `grafana` Investigating elevated database load in AWS Germany (2h 24m) — https://stspg.io/j1r7zw96mgh5
- `anthropic` Elevated errors for Claude Mythos 5.1 and Claude Fable 5.1 (17m) — https://stspg.io/n9n8kfcxtkfm
