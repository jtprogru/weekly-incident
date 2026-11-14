package model

// SourceLink points back to the source record and its physical archive file.
type SourceLink struct {
	Key  string `json:"key"`
	Path string `json:"path"`
	URL  string `json:"url"`
}

// CanonicalInfo accompanies a report projection, never a vendor payload.
type CanonicalInfo struct {
	DurationUnknown bool         `json:"duration_unknown,omitempty"`
	ID              string       `json:"id"`
	Sources         []SourceLink `json:"sources"`
	Reasons         []string     `json:"reasons"`
	StartQuality    string       `json:"start_quality"`
	Warnings        []string     `json:"warnings"`
}

type DedupStatistics struct {
	RulesVersion     string `json:"rules_version"`
	SourceRecords    int    `json:"source_records_total"`
	SourceVersions   int    `json:"source_versions_total"`
	Incidents        int    `json:"canonical_incidents_total"`
	MergedRecords    int    `json:"merged_source_records_total"`
	PendingPairs     int    `json:"pending_candidate_pairs"`
	ConflictedGroups int    `json:"conflicted_groups_total"`
}
