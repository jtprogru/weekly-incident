package source

import "testing"

func TestStripHTML(t *testing.T) {
	cases := map[string]string{
		"<p>Users might be unable to see older messages.</p>": "Users might be unable to see older messages.",
		"<p>one</p><p>two</p>":                                "one\ntwo",
		"a &amp; b &lt;c&gt;":                                 "a & b <c>",
		"line<br/>break":                                      "line\nbreak",
		"<ul><li>first</li><li>second</li></ul>":              "first\nsecond",
		"plain text":                                          "plain text",
		"":                                                    "",
	}
	for in, want := range cases {
		if got := stripHTML(in); got != want {
			t.Errorf("stripHTML(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestFirstSentence(t *testing.T) {
	cases := []struct {
		in    string
		limit int
		want  string
	}{
		{"Short sentence. And another one.", 200, "Short sentence."},
		{"No terminator here", 200, "No terminator here"},
		// A decimal point is not a sentence boundary: no space follows it.
		{"Latency rose to 1.5s and stayed there.", 200, "Latency rose to 1.5s and stayed there."},
		{"aaaaaaaaaa bbbbbbbbbb cccccccccc", 10, "aaaaaaaaaa…"},
		{"   ", 200, ""},
	}
	for _, c := range cases {
		if got := firstSentence(c.in, c.limit); got != c.want {
			t.Errorf("firstSentence(%q, %d) = %q, want %q", c.in, c.limit, got, c.want)
		}
	}
}
