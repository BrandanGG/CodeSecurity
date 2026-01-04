package main

// ingest the OSV data into a database this is a one time operation to build the format for the OSV data to be used by the CLI.

type Findings struct {
	SchemaVersion    string           `json:"schema_version"`
	ID               string           `json:"id"`
	Published        string           `json:"published"`
	Modified         string           `json:"modified"`
	Aliases          []string         `json:"aliases"`
	Related          []string         `json:"related"`
	Summary          string           `json:"summary"`
	Details          string           `json:"details"`
	Affected         []Affected       `json:"affected"`
	References       []References     `json:"references"`
	DatabaseSpecific DatabaseSpecific `json:"database_specific"`
	Severity         []Severity       `json:"severity"`
}

type Affected struct {
	Package          Package          `json:"package"`
	Ranges           []Range          `json:"ranges"`
	DatabaseSpecific DatabaseSpecific `json:"database_specific"`
}

type Package struct {
	Name      string `json:"name"`
	Ecosystem string `json:"ecosystem"`
	PURL      string `json:"purl"`
}

type Range struct {
	Type             string           `json:"type"`
	Events           []Event          `json:"events"`
	DatabaseSpecific DatabaseSpecific `json:"database_specific"`
}

type Event struct {
	Introduced string `json:"introduced,omitempty"`
	Fixed      string `json:"fixed,omitempty"`
}

type References struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

type DatabaseSpecific struct {
	CWEIDs           []string `json:"cwe_ids,omitempty"`
	Severity         string   `json:"severity,omitempty"`
	Source           string   `json:"source,omitempty"`
	GitHubReviewed   bool     `json:"github_reviewed,omitempty"`
	GitHubReviewedAt string   `json:"github_reviewed_at,omitempty"`
	NvdPublishedAt   string   `json:"nvd_published_at,omitempty"`
}

type Severity struct {
	Type  string `json:"type"`
	Score string `json:"score"`
}

// parsed structure for the vuln data
type NormalizedVuln struct {
	VulnID     string
	Alias      string
	Ecosystem  string
	Package    string
	Introduced string
	Fixed      string
	Severity   float64
	Summary    string
}

type NormalizedSeverity struct {
	Score  float64
	Type   string
	vector string
}
