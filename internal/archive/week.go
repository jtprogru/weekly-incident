package archive

import (
	"fmt"
	"time"
)

// Week identifies an ISO-8601 week, the unit the whole archive is bucketed by.
// An incident belongs to the week its impact started in, so one that begins on
// Sunday night and ends on Monday stays whole in the earlier week.
type Week struct {
	Year int
	Num  int
}

// WeekOf returns the ISO week containing t. The ISO year is not always the
// calendar year: 2027-01-01 falls in 2026-W53.
func WeekOf(t time.Time) Week {
	y, n := t.UTC().ISOWeek()
	return Week{Year: y, Num: n}
}

// String renders the week as "2026-W34".
func (w Week) String() string { return fmt.Sprintf("%d-W%02d", w.Year, w.Num) }

// Dir is the archive path segment for the week: "2026/W34".
func (w Week) Dir() string { return fmt.Sprintf("%d/W%02d", w.Year, w.Num) }

// Start returns the Monday 00:00:00 UTC that opens the week.
//
// January 4th is by definition always in ISO week 1, which makes it the anchor
// for finding any week's Monday without a lookup table.
func (w Week) Start() time.Time {
	jan4 := time.Date(w.Year, 1, 4, 0, 0, 0, 0, time.UTC)
	weekday := int(jan4.Weekday())
	if weekday == 0 {
		weekday = 7 // Go counts Sunday as 0; ISO counts it as 7.
	}
	firstMonday := jan4.AddDate(0, 0, -(weekday - 1))
	return firstMonday.AddDate(0, 0, (w.Num-1)*7)
}

// End returns the exclusive upper bound of the week: the next Monday.
func (w Week) End() time.Time { return w.Start().AddDate(0, 0, 7) }

// Contains reports whether t falls inside the week.
func (w Week) Contains(t time.Time) bool {
	u := t.UTC()
	return !u.Before(w.Start()) && u.Before(w.End())
}

// Prev returns the week before w.
func (w Week) Prev() Week { return WeekOf(w.Start().AddDate(0, 0, -1)) }

// ParseWeek reads the "2026-W34" form.
func ParseWeek(s string) (Week, error) {
	var w Week
	if _, err := fmt.Sscanf(s, "%d-W%02d", &w.Year, &w.Num); err != nil {
		return Week{}, fmt.Errorf("parse week %q: %w", s, err)
	}
	if w.Num < 1 || w.Num > 53 {
		return Week{}, fmt.Errorf("parse week %q: week number out of range", s)
	}
	return w, nil
}
