package models

type IngestRequest struct {
	URL string `json:"url"`
}

type IngestResponse struct {
	SourceURL string   `json:"source_url"`
	Chunks    []string `json:"chunks"`
}
