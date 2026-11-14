package dedup

import (
	"fmt"
	"html"
	"strings"
	"time"
)

// Markdown is a review view of the JSON report. All evidence is retained in
// JSON; the prose view limits examples to keep long vendor updates readable.
func Markdown(r Report) string {
	var b strings.Builder
	fmt.Fprintf(&b, "# Аудит дубликатов\n\nПравила: %s. Время: %s.\n\n", r.RulesVersion, r.GeneratedAt.Format(time.RFC3339))
	fmt.Fprintf(&b, "Fingerprint архива: `%s`.\n\nSHA-256 конфигурации: `%s`.\n\n", r.InputFingerprint, r.ConfigSHA256)
	if r.RevisionLabel != "" {
		fmt.Fprintf(&b, "Метка ревизии, переданная при запуске: %s. Fingerprints описывают фактически прочитанные файлы.\n\n", escape(r.RevisionLabel))
	}
	s := r.Statistics
	fmt.Fprintf(&b, "Проверено записей: %d; различных source keys: %d; вендоров: %d; файлов: %d.\n\n", s.Records, s.SourceKeys, s.Vendors, s.Files)
	fmt.Fprintf(&b, "Групп повторяющихся ключей: %d. Пар кандидатов: %d, из них с конфликтующими признаками: %d.\n\n", s.RepeatedKeyGroups, s.CandidatePairs, s.ConflictedPairs)
	b.WriteString("Все пары требуют проверки. Объединения не выполнялись; число уникальных аварий не рассчитано. Отсутствие кандидатов не доказывает отсутствие дублей: похожесть заголовков, пересказ текста и связи между вендорами этот этап не проверяет.\n\n")
	for _, w := range r.Warnings {
		fmt.Fprintf(&b, "- Предупреждение: %s\n", escape(w))
	}
	if len(r.Warnings) > 0 {
		b.WriteString("\n")
	}
	b.WriteString("## Возможности источников\n\n| Вендор | Парсер | Стабильный ID записи | ID обновлений в архиве | Ссылки на события |\n| --- | --- | --- | --- | --- |\n")
	for _, p := range r.Policies {
		fmt.Fprintf(&b, "| %s | %s | %t | %t | %t |\n", escape(p.Vendor), escape(string(p.Kind)), p.StableIncidentID, p.UpdateIDs, p.EventURLs)
	}
	b.WriteString("\nВозможности описывают текущие парсеры. Совпадение ID обновления у разных записей — кандидат для проверки области уникальности, а не доказательство общей аварии.\n\n")
	b.WriteString("## Кандидаты\n\n| # | Левая запись | Правая запись | Уровень | Общие updates | Основания |\n| --- | --- | --- | --- | --- | --- |\n")
	for i, c := range r.Candidates {
		fmt.Fprintf(&b, "| %d | %s | %s | %s | %d | %s |\n", i+1, escape(c.Left.Key), escape(c.Right.Key), c.Level, len(c.SharedUpdates), escape(strings.Join(c.Reasons, ", ")))
	}
	if len(r.Candidates) == 0 {
		b.WriteString("\nПо текущим правилам кандидаты не найдены.\n")
	}
	for i, c := range r.Candidates {
		fmt.Fprintf(&b, "\n## Пара %d\n\nID кандидата: `%s`. Уровень: %s.\n\n", i+1, c.ID, c.Level)
		writeRef(&b, "A", c.Left)
		writeRef(&b, "B", c.Right)
		fmt.Fprintf(&b, "Основания: %s.\n\n", escape(strings.Join(c.Reasons, ", ")))
		if len(c.Cautions) > 0 {
			fmt.Fprintf(&b, "Оговорки / противоречия: %s.\n\n", escape(strings.Join(c.Cautions, ", ")))
		}
		if len(c.SharedUpdateIDs) > 0 {
			fmt.Fprintf(&b, "Совпавшие ID обновлений: %s.\n\n", escape(strings.Join(c.SharedUpdateIDs, ", ")))
		}
		n := min(3, len(c.SharedUpdates))
		if n > 0 {
			fmt.Fprintf(&b, "Примеры совпавших сообщений: %d из %d. Полные тексты всех совпадений — в audit.json.\n\n", n, len(c.SharedUpdates))
		}
		for _, u := range c.SharedUpdates[:n] {
			text := []rune(u.Body)
			if len(text) > 600 {
				text = append(text[:600], []rune("… [сокращено]")...)
			}
			fmt.Fprintf(&b, "- %s: %s\n", u.CreatedAt.Format(time.RFC3339Nano), escape(string(text)))
		}
	}
	return b.String()
}

func writeRef(b *strings.Builder, side string, r RecordRef) {
	end := "неизвестно"
	if r.ResolvedAt != nil {
		end = r.ResolvedAt.UTC().Format(time.RFC3339Nano)
	}
	fmt.Fprintf(b, "- %s: %s — %s\n", side, escape(r.Key), escape(r.Title))
	fmt.Fprintf(b, "  - Файл относительно корня архива: %s; позиция в JSON-массиве: %d.\n", escape(r.Path), r.Position)
	fmt.Fprintf(b, "  - Начало: %s; завершение: %s; сообщений: %d.\n", r.StartedAt.Format(time.RFC3339Nano), end, r.Updates)
	fmt.Fprintf(b, "  - Компоненты: %s. URL источника: %s.\n\n", escape(strings.Join(r.Components, ", ")), escape(r.URL))
}

// Vendor strings are data, never Markdown/HTML instructions in the report.
func escape(s string) string {
	s = html.EscapeString(s)
	r := strings.NewReplacer("\\", "\\\\", "|", "\\|", "`", "\\`", "[", "\\[", "]", "\\]", "*", "\\*", "_", "\\_", "\r", " ", "\n", " ")
	return r.Replace(s)
}
