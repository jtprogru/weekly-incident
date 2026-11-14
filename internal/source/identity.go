package source

import (
	"encoding/json"
	"slices"
	"strings"

	"github.com/jtprogru/weekly-incident/internal/config"
	"github.com/jtprogru/weekly-incident/internal/model"
)

// IdentityUpdates backfills RSS message GUIDs from stored raw payloads without
// mutating the legacy vendor archive. A GUID identifies a message, not an event.
func IdentityUpdates(in model.Incident, kind config.Kind) []model.Update {
	out := slices.Clone(in.Updates)
	if kind != config.KindAWSRSS {
		return out
	}
	var items []awsItem
	if json.Unmarshal(in.Raw, &items) != nil {
		return out
	}
	for i := range out {
		if out[i].NativeID != "" {
			continue
		}
		for _, item := range items {
			at, err := parseAWSTime(item.PubDate)
			service, serviceErr := parseAWSService(item.GUID)
			if err != nil || serviceErr != nil || !slices.Contains(in.Components, service) {
				continue
			}
			if at.Equal(out[i].CreatedAt) && strings.TrimSpace(item.Description) == strings.TrimSpace(out[i].Body) {
				out[i].NativeID = item.GUID
				break
			}
		}
	}
	return out
}
