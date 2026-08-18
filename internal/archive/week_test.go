package archive

import (
	"testing"
	"time"
)

func TestWeekString(t *testing.T) {
	w := WeekOf(time.Date(2026, 8, 18, 12, 0, 0, 0, time.UTC))
	if got := w.String(); got != "2026-W34" {
		t.Errorf("String() = %q, want 2026-W34", got)
	}
	if got := w.Dir(); got != "2026/W34" {
		t.Errorf("Dir() = %q, want 2026/W34", got)
	}
}

// TestWeekBoundsHoldEverywhere walks four years a day at a time. Both ends of a
// calendar year are where ISO weeks stop matching intuition, and a bucket that
// silently drifts would misfile incidents without ever failing loudly.
func TestWeekBoundsHoldEverywhere(t *testing.T) {
	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	for d := range 4 * 366 {
		day := start.AddDate(0, 0, d)
		w := WeekOf(day)

		if !w.Contains(day) {
			t.Fatalf("%s: WeekOf gave %s, which does not contain it (start %s, end %s)",
				day.Format(time.RFC3339), w, w.Start(), w.End())
		}
		if w.Start().Weekday() != time.Monday {
			t.Fatalf("%s: week %s starts on %s, want Monday", day.Format(time.RFC3339), w, w.Start().Weekday())
		}
		if got := w.End().Sub(w.Start()); got != 7*24*time.Hour {
			t.Fatalf("week %s spans %s, want 168h", w, got)
		}
		if WeekOf(w.Start()) != w || WeekOf(w.End().Add(-time.Nanosecond)) != w {
			t.Fatalf("week %s does not round-trip at its own edges", w)
		}
	}
}

func TestWeekAcrossYearBoundary(t *testing.T) {
	// 2027-01-01 is a Friday, so ISO puts it in the last week of 2026. The
	// archive path has to follow the ISO year, not the calendar year, or the
	// incident lands in a directory of its own.
	w := WeekOf(time.Date(2027, 1, 1, 12, 0, 0, 0, time.UTC))
	if w.Year != 2026 {
		t.Errorf("ISO year = %d, want 2026", w.Year)
	}
	if w.Start().Year() != 2026 || w.Start().Month() != time.December {
		t.Errorf("week starts %s, want a December 2026 Monday", w.Start().Format(time.RFC3339))
	}
}

func TestPrev(t *testing.T) {
	w := WeekOf(time.Date(2026, 8, 18, 0, 0, 0, 0, time.UTC))
	prev := w.Prev()
	if got := prev.String(); got != "2026-W33" {
		t.Errorf("Prev() = %q, want 2026-W33", got)
	}
	if got := w.Start().Sub(prev.Start()); got != 7*24*time.Hour {
		t.Errorf("Prev() is %s earlier, want 168h", got)
	}
	// The first week of a year steps back into the previous ISO year.
	first := Week{Year: 2026, Num: 1}
	if p := first.Prev(); p.Year != 2025 {
		t.Errorf("Prev() of the first week = %s, want an ISO week of 2025", p)
	}
}

func TestParseWeek(t *testing.T) {
	w, err := ParseWeek("2026-W34")
	if err != nil {
		t.Fatalf("ParseWeek: %v", err)
	}
	if w.Year != 2026 || w.Num != 34 {
		t.Errorf("ParseWeek = %+v", w)
	}
	if w != WeekOf(time.Date(2026, 8, 18, 0, 0, 0, 0, time.UTC)) {
		t.Error("ParseWeek does not round-trip with WeekOf")
	}
	for _, bad := range []string{"", "nonsense", "2026-W00", "2026-W54"} {
		if _, err := ParseWeek(bad); err == nil {
			t.Errorf("ParseWeek(%q) succeeded, want an error", bad)
		}
	}
}
