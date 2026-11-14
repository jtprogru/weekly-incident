// Command dedup audits the local archive without fetching or merging records.
package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/jtprogru/weekly-incident/internal/archive"
	"github.com/jtprogru/weekly-incident/internal/clock"
	"github.com/jtprogru/weekly-incident/internal/config"
	"github.com/jtprogru/weekly-incident/internal/dedup"
)

func main() {
	log.SetFlags(0)
	if err := run(os.Args[1:], os.Stderr); err != nil {
		log.Fatalf("dedup: %v", err)
	}
}

func run(args []string, stderr io.Writer) error {
	if len(args) > 0 && (args[0] == "validate" || args[0] == "plan" || args[0] == "apply" || args[0] == "rebuild") {
		return maintenance(args, stderr)
	}
	if len(args) == 0 || args[0] != "audit" {
		return fmt.Errorf("usage: dedup audit -out <report-directory> [-data data] [-config config.yaml] [-now RFC3339] [-revision label]")
	}
	flags := flag.NewFlagSet("dedup audit", flag.ContinueOnError)
	flags.SetOutput(stderr)
	data := flags.String("data", "data", "read-only archive root")
	configPath := flags.String("config", "config.yaml", "source configuration")
	out := flags.String("out", "", "required report directory, outside the archive and weeks")
	nowFlag := flags.String("now", "", "fixed report time, RFC3339")
	revision := flags.String("revision", "", "optional caller-supplied revision label (file hashes remain authoritative)")
	if err := flags.Parse(args[1:]); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return nil
		}
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("unexpected positional arguments: %s", strings.Join(flags.Args(), " "))
	}
	if *out == "" {
		return fmt.Errorf("-out is required; audit never writes to data or weeks")
	}
	now, err := clock.Now(*nowFlag)
	if err != nil {
		return err
	}
	output, err := checkOutput(*data, *out)
	if err != nil {
		return err
	}
	configInput, err := resolvedPath(*configPath)
	if err != nil {
		return err
	}
	if within(configInput, output) {
		return fmt.Errorf("report directory contains the input configuration")
	}
	// Hash and parse the same config bytes, even if a concurrent edit happens.
	configBytes, err := os.ReadFile(*configPath)
	if err != nil {
		return fmt.Errorf("read config: %w", err)
	}
	cfg, err := config.Parse(configBytes)
	if err != nil {
		return err
	}
	inv, err := archive.New(*data).ReadAll()
	if err != nil {
		return err
	}
	r := dedup.Audit(inv, cfg.Sources, now)
	r.RevisionLabel = *revision
	sum := sha256.Sum256(configBytes)
	r.ConfigSHA256 = hex.EncodeToString(sum[:])
	b, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return fmt.Errorf("encode audit: %w", err)
	}
	b = append(b, '\n')
	// No output is created until all input files have been validated.
	if err := os.MkdirAll(output, 0o755); err != nil {
		return err
	}
	files := []struct {
		name string
		body []byte
	}{{"audit.json", b}, {"audit.md", []byte(dedup.Markdown(r))}}
	for _, file := range files {
		info, err := os.Lstat(filepath.Join(output, file.name))
		if err == nil && !info.Mode().IsRegular() {
			return fmt.Errorf("report %s is not a regular file", file.name)
		}
		if err != nil && !errors.Is(err, os.ErrNotExist) {
			return err
		}
	}
	for _, file := range files {
		if err := writeReport(filepath.Join(output, file.name), file.body); err != nil {
			return err
		}
	}
	_, err = fmt.Fprintf(stderr, "audit: %d records, %d vendors, %d candidate pairs (%d conflicts) -> %s\n", r.Statistics.Records, r.Statistics.Vendors, r.Statistics.CandidatePairs, r.Statistics.ConflictedPairs, output)
	return err
}

// Resolve existing parents, so an output symlink cannot bypass archive/week
// protection. The sibling weeks tree and the cwd's weeks tree are protected
// even when auditing a custom data directory.
func checkOutput(data, out string) (string, error) {
	output, err := resolvedPath(out)
	if err != nil {
		return "", err
	}
	root, err := resolvedPath(data)
	if err != nil {
		return "", err
	}
	for _, protected := range []string{root, filepath.Join(filepath.Dir(root), "weeks"), "weeks"} {
		p, err := resolvedPath(protected)
		if err != nil {
			return "", err
		}
		if within(output, p) || within(p, output) {
			return "", fmt.Errorf("report directory %s overlaps protected input/output tree %s", output, p)
		}
	}
	return output, nil
}

func resolvedPath(path string) (string, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	resolved, err := filepath.EvalSymlinks(abs)
	if err == nil {
		return resolved, nil
	}
	if !errors.Is(err, os.ErrNotExist) {
		return "", err
	}
	if parent := filepath.Dir(abs); parent != abs {
		// A dangling symlink must not be treated as a new ordinary directory.
		if info, statErr := os.Lstat(abs); statErr == nil && info.Mode()&os.ModeSymlink != 0 {
			return "", fmt.Errorf("dangling symlink: %s", abs)
		}
		resolved, err := resolvedPath(parent)
		if err != nil {
			return "", err
		}
		return filepath.Join(resolved, filepath.Base(abs)), nil
	}
	return "", err
}

func within(path, root string) bool {
	rel, err := filepath.Rel(root, path)
	return err == nil && rel != ".." && !strings.HasPrefix(rel, ".."+string(filepath.Separator))
}

func writeReport(path string, body []byte) error {
	old, err := os.ReadFile(path)
	if err == nil && bytes.Equal(old, body) {
		return nil
	}
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}
	f, err := os.CreateTemp(filepath.Dir(path), ".audit-*")
	if err != nil {
		return err
	}
	defer func() { _ = os.Remove(f.Name()) }()
	if _, err := f.Write(body); err != nil {
		_ = f.Close()
		return err
	}
	if err := f.Chmod(0o644); err != nil {
		_ = f.Close()
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return os.Rename(f.Name(), path)
}
