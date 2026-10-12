#!/usr/bin/env bash
# Re-capture the golden feed responses the parser tests run against.
#
# Run this when a vendor changes its schema and a parser test starts failing:
# the point of the goldens is that the tests describe the feeds as they really
# are, so they have to be refreshed from the real endpoints, not hand-edited.
#
# The captures are trimmed to keep the repository small. Statuspage, Slack, GCP
# and Yandex Cloud are cut to a handful of incidents; the AWS feed is kept whole
# because its grouping logic needs several chains to exercise.
set -euo pipefail

cd "$(dirname "$0")/.."
UA='weekly-incident/0.1 (+https://github.com/jtprogru/weekly-incident)'

python3 - "$UA" <<'PY'
import json, re, sys, urllib.request

ua = sys.argv[1]

def get(url):
    return urllib.request.urlopen(
        urllib.request.Request(url, headers={'User-Agent': ua}), timeout=30
    ).read()

def dump(path, obj):
    with open(path, 'w', encoding='utf-8') as f:
        f.write(json.dumps(obj, indent=2, ensure_ascii=False) + '\n')
    print(f'wrote {path}')

# Statuspage: GitHub stands in for all ten vendors on this schema.
sp = json.loads(get('https://www.githubstatus.com/api/v2/incidents.json'))
sp['incidents'] = sp['incidents'][:5]
dump('testdata/statuspage_github.json', sp)

dump('testdata/slack_history.json',
     json.loads(get('https://slack-status.com/api/v2.0.0/history'))[:5])

# Two GCP incidents is enough: one of them carries a full postmortem.
dump('testdata/gcp_incidents.json',
     json.loads(get('https://status.cloud.google.com/incidents.json'))[:2])

# Yandex Cloud: the newest five, plus the newest incident with a published
# report, because reports land weeks late and the recent window rarely holds
# one. `page` is cumulative, so one deep page is the whole history. A zone-wide
# outage names eighty-odd services and each drags an inline SVG icon along, so
# every service list is cut to three. The old host is deliberate: see the note
# on this source in config.yaml.
yc = json.loads(get('https://status.cloud.yandex.ru/api/incidents?lang=en&page=99'))
yc['items'] = yc['items'][:5] + [i for i in yc['items'][5:] if i['isReportPublished']][:1]
for item in yc['items']:
    item['services'] = item['services'][:3]
dump('testdata/yandexcloud_incidents.json', yc)

# AWS is kept in full: the grouping needs several service chains, one closed by
# a [RESOLVED] headline and the others only by the gap rule.
rss = get('https://status.aws.amazon.com/rss/all.rss').decode('utf-8', 'replace')
head = rss[:rss.index('<item>')]
items = re.findall(r'[ \t]*<item>.*?</item>\n?', rss, re.S)
with open('testdata/aws_all.rss', 'w', encoding='utf-8') as f:
    f.write(head + ''.join(items) + '  </channel>\n</rss>\n')
print(f'wrote testdata/aws_all.rss ({len(items)} items)')
PY

echo
echo 'Goldens refreshed. Run "make test" — failures now describe real schema changes.'
