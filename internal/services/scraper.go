package services

import (
	"credilens-backend/internal/chunker"
	"credilens-backend/internal/models"
	"credilens-backend/internal/processor"
	"credilens-backend/internal/scraper"
)

func Scrape(URL string) (models.IngestResponse, error) {

	rawText, images, err := scraper.ScrapeText(URL)
	if err != nil {
		return models.IngestResponse{}, err
	}

	cleanText := processor.Process(rawText)

	chunks := chunker.Split(
		cleanText,
		50, // chunk size
		5,  // overlap
	)

	return models.IngestResponse{
		SourceURL: URL,
		Chunks:    chunks,
		Images:    images,
	}, nil
}
