package clock

import (
	"testing"
	"time"
)

func TestNowDefaultsToWallClock(t *testing.T) {
	got, err := Now("")
	if err != nil {
		t.Fatalf("Now: %v", err)
	}
	if got.Location() != time.UTC {
		t.Errorf("location = %s, want UTC", got.Location())
	}
	if time.Since(got) > time.Minute {
		t.Errorf("Now() = %s, which is not now", got)
	}
}

func TestNowHonoursTheOverride(t *testing.T) {
	// Pinning the clock is what makes a run reproducible; the archive stamps
	// FirstSeenAt with it and ongoing durations are measured against it.
	got, err := Now("2026-08-18T06:00:00Z")
	if err != nil {
		t.Fatalf("Now: %v", err)
	}
	if want := time.Date(2026, 8, 18, 6, 0, 0, 0, time.UTC); !got.Equal(want) {
		t.Errorf("Now = %s, want %s", got, want)
	}

	// A non-UTC offset is normalized rather than carried around.
	got, err = Now("2026-08-18T09:00:00+03:00")
	if err != nil {
		t.Fatalf("Now: %v", err)
	}
	if want := time.Date(2026, 8, 18, 6, 0, 0, 0, time.UTC); !got.Equal(want) || got.Location() != time.UTC {
		t.Errorf("Now = %s (%s), want %s in UTC", got, got.Location(), want)
	}
}

func TestNowRejectsGarbage(t *testing.T) {
	for _, bad := range []string{"yesterday", "2026-08-18", "18/08/2026"} {
		if _, err := Now(bad); err == nil {
			t.Errorf("Now(%q) succeeded, want an error", bad)
		}
	}
}
