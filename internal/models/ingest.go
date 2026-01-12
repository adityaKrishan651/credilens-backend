package models

type IngestRequest struct {
	URL string `json:"url"`
}

type IngestResponse struct {
	SourceURL string          `json:"source_url"`
	Chunks    []string        `json:"chunks"`
	Images    []ImageMetadata `json:"images"`
}

type ImageMetadata struct {
	URL             string
	AltText         string
	Position        int
	SurroundingText string
}
