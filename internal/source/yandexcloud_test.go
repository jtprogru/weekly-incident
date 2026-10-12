package source

import (
	"bytes"
	"reflect"
	"strings"
	"testing"

	"github.com/jtprogru/weekly-incident/internal/model"
)

func yandexCloudGolden(t *testing.T) []model.Incident {
	t.Helper()
	y := &yandexCloud{base: base{vendor: "yandexcloud", url: "https://status.cloud.yandex.ru/api/incidents?lang=en&page=4"}}
	res, err := y.parse(golden(t, "yandexcloud_incidents.json"))
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if res.ParseErrors != 0 {
		t.Errorf("ParseErrors = %d, want 0", res.ParseErrors)
	}
	if len(res.Incidents) != 6 {
		t.Fatalf("got %d incidents, want 6", len(res.Incidents))
	}
	return res.Incidents
}

func TestYandexCloudParse(t *testing.T) {
	incidents := yandexCloudGolden(t)

	in := incidents[0]
	if in.Key != "yandexcloud/2136" {
		t.Errorf("Key = %q", in.Key)
	}
	if in.Title != "Unavailability of Yandex Cloud zones" {
		t.Errorf("Title = %q", in.Title)
	}
	if in.URL != "https://status.yandex.cloud/en/incidents/2136" {
		t.Errorf("URL = %q", in.URL)
	}
	// Level 2, Unavailable: the top level the board actually uses.
	if in.Impact != model.ImpactMajor {
		t.Errorf("Impact = %q, want major", in.Impact)
	}
	if in.Status != model.StatusInvestigating {
		t.Errorf("Status = %q, want investigating (the incident is open)", in.Status)
	}
	if want := mustTime(t, "2026-10-11T01:00:00Z"); !in.StartedAt.Equal(want) {
		t.Errorf("StartedAt = %s, want %s", in.StartedAt, want)
	}
	if in.ResolvedAt != nil || in.DurationMinutes != nil {
		t.Errorf("ResolvedAt = %v, DurationMinutes = %v, want both nil on an open incident", in.ResolvedAt, in.DurationMinutes)
	}
	// The capture cuts every service list to three.
	want := []string{"Compute Cloud", "Virtual Private Cloud", "Object Storage"}
	if !reflect.DeepEqual(in.Components, want) {
		t.Errorf("Components = %v, want %v", in.Components, want)
	}

	if len(in.Updates) != 9 {
		t.Fatalf("got %d updates, want 9", len(in.Updates))
	}
	// The feed lists comments newest first; the archive stores them in order.
	first := in.Updates[0]
	if want := mustTime(t, "2026-10-11T01:11:22.672Z"); !first.CreatedAt.Equal(want) {
		t.Errorf("first update at %s, want the oldest (%s)", first.CreatedAt, want)
	}
	if first.NativeID != "2847" || first.Status != model.StatusInvestigating {
		t.Errorf("first update = %q %q, want 2847 investigating", first.NativeID, first.Status)
	}
	if strings.ContainsAny(first.Body, "<>") {
		t.Errorf("first update still carries markup: %q", first.Body)
	}
	assertAscending(t, in.Updates)

	// The icons are in the feed and must not reach the archive.
	if !bytes.Contains(golden(t, "yandexcloud_incidents.json"), []byte(`"icon":`)) {
		t.Fatal("golden carries no service icons, so the check below proves nothing")
	}
	if bytes.Contains(in.Raw, []byte(`"icon":`)) {
		t.Error("Raw still carries the inline SVG icons")
	}
	if !bytes.Contains(in.Raw, []byte(`"iconName":"compute"`)) {
		t.Error("Raw lost more than the icons")
	}

	// A resolved incident takes its end from endDate, not from the closing
	// comment, which the vendor posted four seconds later.
	in = incidents[1]
	if in.Impact != model.ImpactMinor || in.Status != model.StatusResolved {
		t.Errorf("Impact, Status = %q, %q, want minor, resolved", in.Impact, in.Status)
	}
	if in.DurationMinutes == nil || *in.DurationMinutes != 49 {
		t.Errorf("DurationMinutes = %v, want 49", in.DurationMinutes)
	}
}

// Incident 2092 was closed without an end date when it was folded into 2136.
// Without a fallback it would read as resolved and still running.
func TestYandexCloudClosedWithoutEndDate(t *testing.T) {
	in := yandexCloudGolden(t)[4]
	if in.Key != "yandexcloud/2092" {
		t.Fatalf("Key = %q, want yandexcloud/2092", in.Key)
	}
	if in.Status != model.StatusResolved {
		t.Errorf("Status = %q, want resolved", in.Status)
	}
	if in.ResolvedAt == nil {
		t.Fatal("ResolvedAt is nil on a resolved incident")
	}
	if want := mustTime(t, "2026-10-10T17:13:52.935Z"); !in.ResolvedAt.Equal(want) {
		t.Errorf("ResolvedAt = %s, want the last comment (%s)", in.ResolvedAt, want)
	}
	if in.DurationMinutes == nil || *in.DurationMinutes != 4002 {
		t.Errorf("DurationMinutes = %v, want 4002", in.DurationMinutes)
	}
}

func TestYandexCloudReportBecomesPostmortem(t *testing.T) {
	in := yandexCloudGolden(t)[5]
	if in.Key != "yandexcloud/1009" {
		t.Fatalf("Key = %q, want yandexcloud/1009", in.Key)
	}
	if in.Status != model.StatusPostmortem {
		t.Errorf("Status = %q, want postmortem", in.Status)
	}
	// Seven comments plus the report.
	if len(in.Updates) != 8 {
		t.Fatalf("got %d updates, want 8", len(in.Updates))
	}
	last := in.Updates[len(in.Updates)-1]
	if last.Status != model.StatusPostmortem {
		t.Errorf("last update status = %q, want postmortem", last.Status)
	}
	if want := mustTime(t, "2024-12-12T06:33:47.822Z"); !last.CreatedAt.Equal(want) {
		t.Errorf("report at %s, want %s", last.CreatedAt, want)
	}
	if !strings.HasPrefix(last.Body, "On November 29 an incident occurred") {
		t.Errorf("report body = %.60q", last.Body)
	}
	// The report lands two weeks later and must not stretch the incident.
	if in.DurationMinutes == nil || *in.DurationMinutes != 284 {
		t.Errorf("DurationMinutes = %v, want 284", in.DurationMinutes)
	}
	assertAscending(t, in.Updates)
}

func TestYandexCloudImpact(t *testing.T) {
	cases := map[int]model.Impact{
		1: model.ImpactMinor,
		2: model.ImpactMajor,
		3: model.ImpactMajor,
		4: model.ImpactCritical,
		0: model.ImpactUnknown,
		5: model.ImpactUnknown,
	}
	for level, want := range cases {
		if got := ycImpact(level); got != want {
			t.Errorf("ycImpact(%d) = %q, want %q", level, got, want)
		}
	}
}

func TestYandexCloudUsesPeakLevel(t *testing.T) {
	// Seen on incident 1904: unavailable at its worst, minor by the time it
	// closed. The digest ranks the outage, not its last hour.
	body := []byte(`{"items":[
		{"id":1,"title":"wound down","status":"resolved","levelId":1,"maxLevelId":2,
		 "startDate":"2026-07-01T10:00:00.000Z","endDate":"2026-07-01T11:00:00.000Z"}
	]}`)
	y := &yandexCloud{base: base{vendor: "test"}}
	res, err := y.parse(body)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if len(res.Incidents) != 1 {
		t.Fatalf("got %d incidents, want 1", len(res.Incidents))
	}
	if got := res.Incidents[0].Impact; got != model.ImpactMajor {
		t.Errorf("Impact = %q, want major", got)
	}
}

func TestYandexCloudStatuses(t *testing.T) {
	for in, want := range map[string]model.Status{
		"open":     model.StatusInvestigating,
		"resolved": model.StatusResolved,
		"draft":    model.StatusUnknown,
	} {
		if got := ycStatus(in); got != want {
			t.Errorf("ycStatus(%q) = %q, want %q", in, got, want)
		}
	}
	for in, want := range map[string]model.Status{
		"investigation": model.StatusInvestigating,
		"resolved":      model.StatusResolved,
		"update":        model.StatusUnknown,
	} {
		if got := ycCommentStatus(in); got != want {
			t.Errorf("ycCommentStatus(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestYandexCloudSkipsUnparseableIncident(t *testing.T) {
	body := []byte(`{"items":[
		{"id":1,"title":"fine","status":"resolved","levelId":1,"maxLevelId":1,
		 "startDate":"2026-07-01T10:00:00.000Z","endDate":"2026-07-01T11:00:00.000Z"},
		{"id":2,"title":"no start"}
	]}`)
	y := &yandexCloud{base: base{vendor: "test"}}
	res, err := y.parse(body)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if len(res.Incidents) != 1 {
		t.Errorf("got %d incidents, want 1", len(res.Incidents))
	}
	if res.ParseErrors != 1 {
		t.Errorf("ParseErrors = %d, want 1", res.ParseErrors)
	}
}

func TestYandexCloudRejectsPayloadWithoutItems(t *testing.T) {
	// What the antirobot in front of the board sends instead of the feed. It
	// must fail the source, not pass for a week without incidents.
	y := &yandexCloud{base: base{vendor: "test"}}
	if _, err := y.parse([]byte(`{"type": "blocked"}`)); err == nil {
		t.Error("parse accepted a payload without items")
	}
}
