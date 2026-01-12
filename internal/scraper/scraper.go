package scraper

import (
	"credilens-backend/internal/extractor"
	"credilens-backend/internal/models"

	"github.com/gocolly/colly"
)

func ScrapeText(url string) (string, []models.ImageMetadata, error) {
	var (
		bodyText string
		images   []models.ImageMetadata
	)

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (compatible; CredilensBot/1.0)"),
	)

	c.OnHTML("script, style, nav, footer, header, noscript", func(e *colly.HTMLElement) {
		e.DOM.Remove()
	})

	c.OnHTML("body", func(e *colly.HTMLElement) {
		bodyText = e.Text
		images = extractor.ExtractImages(e, e.Request.URL.String())
	})

	if err := c.Visit(url); err != nil {
		return "", nil, err
	}

	return extractor.ExtractText(bodyText), images, nil
}
