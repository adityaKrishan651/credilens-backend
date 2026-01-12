package extractor

import (
	"credilens-backend/internal/models"
	"credilens-backend/internal/utils"

	"github.com/gocolly/colly"
)

func ExtractImages(e *colly.HTMLElement, baseURL string) []models.ImageMetadata {
	var images []models.ImageMetadata
	position := 0

	e.ForEach("img", func(_ int, img *colly.HTMLElement) {
		position++

		src := img.Attr("src")
		if src == "" {
			return
		}

		imageURL := utils.ResolveURL(src, baseURL)

		if utils.IsDecorative(img) {
			return
		}

		metadata := models.ImageMetadata{
			URL:             imageURL,
			AltText:         img.Attr("alt"),
			Position:        position,
			SurroundingText: utils.GetNearbyText(img),
		}

		images = append(images, metadata)
	})

	return images
}
