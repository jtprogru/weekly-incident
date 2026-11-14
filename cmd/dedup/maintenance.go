package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"path/filepath"
	"strings"

	"github.com/jtprogru/weekly-incident/internal/archive"
	"github.com/jtprogru/weekly-incident/internal/canonical"
	"github.com/jtprogru/weekly-incident/internal/clock"
	"github.com/jtprogru/weekly-incident/internal/config"
	"github.com/jtprogru/weekly-incident/internal/storage"
)

func maintenance(args []string, stderr io.Writer) error {
	f := flag.NewFlagSet("dedup "+args[0], flag.ContinueOnError)
	f.SetOutput(stderr)
	data := f.String("data", "data", "archive root")
	cp := f.String("config", "config.yaml", "configuration")
	dp := f.String("decisions", "", "override decisions YAML path")
	out := f.String("out", "", "plan report directory (required for plan)")
	planPath := f.String("plan", "", "reviewed plan.json to apply")
	weeks := f.String("weeks", "", "render every historical week here after rebuild, without notifications")
	nowFlag := f.String("now", "", "fixed time, RFC3339")
	if err := f.Parse(args[1:]); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return nil
		}
		return err
	}
	if f.NArg() != 0 {
		return fmt.Errorf("unexpected arguments: %s", strings.Join(f.Args(), " "))
	}
	if args[0] == "plan" && *out == "" {
		return fmt.Errorf("plan requires -out")
	}
	if args[0] == "apply" && *planPath == "" {
		return fmt.Errorf("apply requires -plan")
	}
	if args[0] != "rebuild" && *weeks != "" {
		return fmt.Errorf("-weeks is only supported by rebuild")
	}
	now, err := clock.Now(*nowFlag)
	if err != nil {
		return err
	}
	cfg, err := config.Load(*cp)
	if err != nil {
		return err
	}
	decisions := *dp
	if decisions == "" {
		decisions = cfg.DecisionsPath(*cp)
	}
	if *weeks != "" {
		root, err := resolvedPath(*data)
		if err != nil {
			return err
		}
		dest, err := resolvedPath(*weeks)
		if err != nil {
			return err
		}
		if within(dest, root) || within(root, dest) {
			return fmt.Errorf("weeks directory overlaps archive")
		}
	}
	if args[0] == "apply" || args[0] == "rebuild" {
		release, err := storage.Lock(*data)
		if err != nil {
			return err
		}
		defer release()
	}
	switch args[0] {
	case "plan":
		dir, err := checkOutput(*data, *out)
		if err != nil {
			return err
		}
		for _, input := range []string{*cp, decisions} {
			path, err := resolvedPath(input)
			if err != nil {
				return err
			}
			if within(path, dir) {
				return fmt.Errorf("report directory contains input %s", input)
			}
		}
		p, err := canonical.Prepare(*data, cfg, decisions, now, true, true)
		if err != nil {
			return err
		}
		if err := storage.WriteJSON(filepath.Join(dir, "plan.json"), p.Plan); err != nil {
			return err
		}
		if err := storage.Write(filepath.Join(dir, "plan.md"), []byte(planMarkdown(p.Plan))); err != nil {
			return err
		}
		_, err = fmt.Fprintf(stderr, "plan: %d source records -> %d canonical incidents; %d changed groups, %d conflicts -> %s\n", p.Plan.Statistics.SourceRecords, p.Plan.Statistics.Incidents, len(p.Plan.Changes), len(p.Plan.Conflicts), dir)
		return err
	case "apply":
		var p canonical.Plan
		if err := storage.Read(*planPath, &p); err != nil {
			return err
		}
		if err := canonical.Apply(*data, cfg, decisions, p); err != nil {
			return err
		}
		_, err = fmt.Fprintf(stderr, "applied: %d canonical incidents, %d weeks; source records retained\n", p.Statistics.Incidents, len(p.Weeks))
		return err
	case "validate":
		p, err := canonical.Load(*data, cfg, decisions)
		if err != nil {
			return err
		}
		built, err := canonical.Build(p.Observations, p.Registry, p.Decisions, cfg.Sources)
		if err != nil {
			return err
		}
		for _, in := range built {
			if len(in.Conflicts) > 0 {
				return fmt.Errorf("canonical conflicts in %s: %v", in.ID, in.Conflicts)
			}
		}
		actual, err := storage.Digest(built)
		if err != nil {
			return err
		}
		expected, err := storage.Digest(p.Incidents)
		if err != nil {
			return err
		}
		if actual != expected {
			return fmt.Errorf("canonical content differs from current builder; rebuild required")
		}
		_, err = fmt.Fprintf(stderr, "valid: %d canonical incidents, %d source records\n", len(p.Incidents), len(p.Observations.Manifest.Records))
		return err
	case "rebuild":
		p, err := canonical.Prepare(*data, cfg, decisions, now, false, false)
		if err != nil {
			return err
		}
		if err := canonical.Commit(*data, p, cfg, nil, nil); err != nil {
			return err
		}
		if *weeks != "" {
			for _, id := range p.Plan.Weeks {
				w, err := archive.ParseWeek(id)
				if err != nil {
					return err
				}
				if err := canonical.RenderWeek(*data, *weeks, p, cfg, w, now); err != nil {
					return err
				}
			}
		}
		_, err = fmt.Fprintf(stderr, "rebuilt: %d canonical incidents, %d weeks; no notifications sent\n", len(p.Incidents), len(p.Plan.Weeks))
		return err
	}
	return fmt.Errorf("unknown command")
}

func planMarkdown(p canonical.Plan) string {
	var b strings.Builder
	fmt.Fprintf(&b, "# План применения дедупликации\n\nВремя: %s. Правила: %s.\n\n", p.At.Format("2006-01-02T15:04:05Z07:00"), p.RulesVersion)
	fmt.Fprintf(&b, "Записей источников: %d; версий: %d; канонических инцидентов: %d; повторных записей внутри групп: %d; непроверенных пар: %d; конфликтных групп: %d.\n\n", p.Statistics.SourceRecords, p.Statistics.SourceVersions, p.Statistics.Incidents, p.Statistics.MergedRecords, p.Statistics.PendingPairs, p.Statistics.ConflictedGroups)
	fmt.Fprintf(&b, "Fingerprint архива: `%s`.\n\nПолные входы, proposed hashes и изменения полей — в plan.json. Apply повторно вычисляет план и отказывает при изменении входов.\n\n", p.Inputs.Archive)
	b.WriteString("| Canonical ID | Исходные записи | Недели до → после | Score до → после | Основание |\n| --- | --- | --- | --- | --- |\n")
	for _, c := range p.Changes {
		beforeWeeks := []string{}
		beforeScores := []string{}
		for _, row := range c.Before {
			beforeWeeks = append(beforeWeeks, row.Week)
			beforeScores = append(beforeScores, fmt.Sprintf("%.2f", row.Score.Value))
		}
		fmt.Fprintf(&b, "| %s | %s | %s → %s | %s → %.2f | %s |\n", c.CanonicalID, strings.Join(c.Members, ", "), strings.Join(beforeWeeks, ", "), c.After.Week, strings.Join(beforeScores, ", "), c.After.Score.Value, strings.Join(c.Reasons, ", "))
	}
	if len(p.Conflicts) > 0 {
		b.WriteString("\nПрименение заблокировано конфликтами; подробности — в plan.json.\n")
	}
	b.WriteString("\n## Недельные сравнения\n\nОбе стороны рассчитаны с одинаковым -now. Top до содержит source keys, top после — canonical IDs; соответствие дано выше и в реестре.\n\n| Неделя | Записи → инциденты | Выше threshold до → после | Top до | Top после |\n| --- | --- | --- | --- | --- |\n")
	for _, w := range p.WeekComparisons {
		fmt.Fprintf(&b, "| %s | %d → %d | %d → %d | %s | %s |\n", w.Week, w.SourceRecords, w.CanonicalIncidents, w.AboveBefore, w.AboveAfter, strings.Join(w.TopBefore, ", "), strings.Join(w.TopAfter, ", "))
	}
	b.WriteString("\nИсходные записи сохраняются. Число канонических инцидентов относится к собранным сообщениям, а не ко всем реальным авариям.\n")
	return b.String()
}
