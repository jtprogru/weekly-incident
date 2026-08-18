package source

import (
	"html"
	"strings"
)

// blockBreaks are the tags that carry a line break's worth of meaning. They are
// turned into newlines before the rest of the markup is dropped, so a Slack
// note made of several <p> blocks does not collapse into one run-on line.
var blockBreaks = strings.NewReplacer(
	"</p>", "\n",
	"<br>", "\n",
	"<br/>", "\n",
	"<br />", "\n",
	"</div>", "\n",
	"</li>", "\n",
)

// stripHTML renders a vendor's HTML update body as plain text. The original
// markup survives in the incident's Raw field, so nothing is lost here.
func stripHTML(s string) string {
	s = blockBreaks.Replace(s)

	var b strings.Builder
	b.Grow(len(s))
	depth := 0
	for _, r := range s {
		switch {
		case r == '<':
			depth++
		case r == '>':
			if depth > 0 {
				depth--
			}
		case depth == 0:
			b.WriteRune(r)
		}
	}

	out := html.UnescapeString(b.String())

	// Collapse the whitespace the tags left behind, but keep paragraph breaks.
	lines := strings.Split(out, "\n")
	kept := lines[:0]
	for _, l := range lines {
		if l = strings.TrimSpace(l); l != "" {
			kept = append(kept, l)
		}
	}
	return strings.Join(kept, "\n")
}

// firstSentence returns the leading sentence of s, or the first limit runes if
// no sentence boundary shows up first. Used for vendors that ship a paragraph
// where a title belongs; the full text stays in Raw.
func firstSentence(s string, limit int) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return ""
	}
	runes := []rune(s)
	for i, r := range runes {
		if r != '.' && r != '!' && r != '?' {
			continue
		}
		// Skip abbreviations and decimals: a boundary needs a space after it.
		if i+1 < len(runes) && runes[i+1] != ' ' && runes[i+1] != '\n' {
			continue
		}
		if i+1 <= limit {
			return strings.TrimSpace(string(runes[:i+1]))
		}
		break
	}
	if len(runes) <= limit {
		return s
	}
	return strings.TrimSpace(string(runes[:limit])) + "…"
}
