// Package clock resolves the run's notion of "now".
//
// Both commands accept -now so a run can be pinned and reproduced: the archive
// stamps FirstSeenAt with it, ongoing incidents are measured against it, and
// the default week is derived from it. Leaving that parsing duplicated in two
// mains meant two places to get the format wrong.
package clock

import (
	"fmt"
	"time"
)

// Now returns the current UTC time, or the parsed override when one is given.
func Now(override string) (time.Time, error) {
	if override == "" {
		return time.Now().UTC(), nil
	}
	t, err := time.Parse(time.RFC3339, override)
	if err != nil {
		return time.Time{}, fmt.Errorf("parse -now: %w", err)
	}
	return t.UTC(), nil
}
