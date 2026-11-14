package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/jtprogru/weekly-incident/internal/storage"
)

func TestPublishWorkflowSourceOnlyAndResync(t *testing.T) {
	git, err := exec.LookPath("git")
	if err != nil {
		t.Fatal(err)
	}
	script, err := filepath.Abs("../../scripts/collect-publish.sh")
	if err != nil {
		t.Fatal(err)
	}
	if b, err := exec.Command("bash", "-n", script).CombinedOutput(); err != nil {
		t.Fatalf("shell syntax: %s %v", b, err)
	}
	for _, scenario := range []string{"partial", "retry", "source-failure"} {
		t.Run(scenario, func(t *testing.T) {
			root := t.TempDir()
			repo := filepath.Join(root, "repo")
			bin := filepath.Join(root, "bin")
			for _, dir := range []string{repo, bin} {
				if err := os.Mkdir(dir, 0o755); err != nil {
					t.Fatal(err)
				}
			}
			env := append(os.Environ(), "GIT_CONFIG_GLOBAL=/dev/null", "GIT_CONFIG_SYSTEM=/dev/null", "SCENARIO="+scenario, "FIXTURE_ROOT="+root, "REAL_GIT="+git)
			runGit := func(args ...string) string {
				t.Helper()
				c := exec.Command(git, args...)
				c.Dir = repo
				c.Env = env
				b, err := c.CombinedOutput()
				if err != nil {
					t.Fatalf("git %v: %s %v", args, b, err)
				}
				return string(b)
			}
			runGit("init", "-q")
			runGit("config", "user.name", "fixture")
			runGit("config", "user.email", "fixture@example.invalid")
			for _, path := range []string{"data/2026/W09/aws.json", "data/2026/W09/index.json", "data/state.json", "data/observations/manifest.json", "data/canonical/incidents.json", "data/identity/registry.json"} {
				if err := storage.Write(filepath.Join(repo, path), []byte("baseline\n")); err != nil {
					t.Fatal(err)
				}
			}
			runGit("add", "data")
			runGit("commit", "-qm", "baseline")
			gitMock := `#!/usr/bin/env bash
set -euo pipefail
case "$1" in
 push)
  n=0; if [ -f "$FIXTURE_ROOT/pushes" ]; then n=$(cat "$FIXTURE_ROOT/pushes"); fi
  n=$((n+1)); echo "$n" > "$FIXTURE_ROOT/pushes"
  if [ "$SCENARIO" = retry ] && [ "$n" = 1 ]; then exit 1; fi
  exit 0 ;;
 fetch) exit 0 ;;
 reset) exec "$REAL_GIT" reset --hard HEAD~1 ;;
 *) exec "$REAL_GIT" "$@" ;;
esac
`
			makeMock := `#!/usr/bin/env bash
set -euo pipefail
n=0; if [ -f "$FIXTURE_ROOT/calls" ]; then n=$(cat "$FIXTURE_ROOT/calls"); fi
n=$((n+1)); echo "$n" > "$FIXTURE_ROOT/calls"
if [ "$SCENARIO" = source-failure ]; then
 echo '{"source_saved":false,"canonical_saved":false}' > .collect-result.json
 exit 2
fi
printf 'source-%s\n' "$n" > data/2026/W09/aws.json
printf 'observations-%s\n' "$n" > data/observations/manifest.json
printf 'state-%s\n' "$n" > data/state.json
printf 'derived-%s\n' "$n" > data/2026/W09/index.json
printf 'derived-%s\n' "$n" > data/canonical/incidents.json
printf 'derived-%s\n' "$n" > data/identity/registry.json
echo 'fixture collect' > .collect-summary.txt
if [ "$SCENARIO" = partial ]; then
 echo '{"source_saved":true,"canonical_saved":false}' > .collect-result.json
 exit 2
fi
echo '{"source_saved":true,"canonical_saved":true}' > .collect-result.json
`
			for name, body := range map[string]string{"git": gitMock, "make": makeMock} {
				if err := os.WriteFile(filepath.Join(bin, name), []byte(body), 0o755); err != nil {
					t.Fatal(err)
				}
			}
			cmd := exec.Command("bash", script)
			cmd.Dir = repo
			cmd.Env = append(env, "PATH="+bin+string(os.PathListSeparator)+os.Getenv("PATH"))
			output, err := cmd.CombinedOutput()
			if scenario == "retry" && err != nil {
				t.Fatalf("retry: %s %v", output, err)
			}
			if scenario != "retry" && err == nil {
				t.Fatalf("failure turned green: %s", output)
			}
			if scenario == "source-failure" {
				if strings.TrimSpace(runGit("rev-list", "--count", "HEAD")) != "1" {
					t.Fatal("committed unsuccessful persistence")
				}
				return
			}
			n := 1
			if scenario == "retry" {
				n = 2
			}
			if strings.TrimSpace(runGit("show", "HEAD:data/2026/W09/aws.json")) != "source-"+strconv.Itoa(n) {
				t.Fatal("did not publish the fresh source run")
			}
			for _, path := range []string{"data/2026/W09/index.json", "data/canonical/incidents.json", "data/identity/registry.json"} {
				want := "baseline"
				if scenario == "retry" {
					want = "derived-2"
				}
				if strings.TrimSpace(runGit("show", "HEAD:"+path)) != want {
					t.Fatalf("wrong commit scope for %s", path)
				}
			}
			if _, err := os.Stat(filepath.Join(repo, ".collect-result.json")); !os.IsNotExist(err) {
				t.Fatal("result marker leaked")
			}
		})
	}
}
