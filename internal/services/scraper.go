package services

import (
	"credilens-backend/internal/chunker"
	"credilens-backend/internal/extractor"
	"credilens-backend/internal/models"
	"credilens-backend/internal/processor"
)

func Scrape(URL string) (models.IngestResponse, error) {

	rawText := extractor.Extract(URL)
	cleanText := processor.Process(rawText)

	chunks := chunker.Split(
		cleanText,
		0, //chunk size
		0, //overlap,
	)

	return models.IngestResponse{
		SourceURL: URL,
		Chunks:    chunks,
	}, nil

}
