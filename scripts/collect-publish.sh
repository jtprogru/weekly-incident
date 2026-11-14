#!/usr/bin/env bash
# Called by collect.yml. The result marker is written by Go only after source
# persistence; a canonical failure may publish that source subset, never a
# partially generated identity/index set. No old plan survives a resync.
set -euo pipefail

git config user.name 'github-actions[bot]'
git config user.email 'github-actions[bot]@users.noreply.github.com'
trap 'rm -f .collect-result.json .collect-summary.txt .collect-warning.txt' EXIT

published=''
collect_status=0
for attempt in 1 2 3; do
  rm -f .collect-result.json .collect-summary.txt .collect-warning.txt
  collect_status=0
  make collect SUMMARY=.collect-summary.txt WARNING=.collect-warning.txt RESULT=.collect-result.json || collect_status=$?

  if ! python3 -c 'import json; d=json.load(open(".collect-result.json")); raise SystemExit(0 if d.get("source_saved") is True else 1)'; then
    echo 'source persistence did not finish; nothing will be committed' >&2
    exit 1
  fi

  if python3 -c 'import json; d=json.load(open(".collect-result.json")); raise SystemExit(0 if d.get("canonical_saved") is True else 1)'; then
    git add -- data
  else
    git add -- data/observations data/state.json \
      ':(glob)data/[0-9][0-9][0-9][0-9]/W[0-9][0-9]/*.json' \
      ':(exclude,glob)data/[0-9][0-9][0-9][0-9]/W[0-9][0-9]/index.json'
  fi

  if git diff --cached --quiet; then
    published='yes'
    break
  fi
  summary='collect: source data saved; see run log'
  if [ -s .collect-summary.txt ]; then summary="$(cat .collect-summary.txt)"; fi
  git commit -q -m "chore(data): ${summary}"
  if git push; then
    published='yes'
    break
  fi
  echo "push rejected on attempt ${attempt}; resyncing and collecting again"
  git fetch origin main
  git reset --hard origin/main
done

if [ -z "${published}" ]; then
  echo 'could not publish after 3 attempts' >&2
  exit 1
fi

if [ -s .collect-warning.txt ]; then
  echo "::warning::$(cat .collect-warning.txt)"
  if [ -n "${GITHUB_OUTPUT:-}" ]; then
    {
      echo 'warning<<EOF'
      cat .collect-warning.txt
      echo 'EOF'
    } >> "$GITHUB_OUTPUT"
  fi
fi

# A source-only commit preserves evidence, but it must not turn the job green.
exit "$collect_status"
