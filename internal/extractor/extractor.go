package extractor

import (
	"log"
	"strings"

	"github.com/gocolly/colly"
)

func Extract(URL string) string {
	var extractedText []string

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (compatible; TextAnalyzer/1.0)"),
	)

	c.OnHTML("script, style, nav, footer, header, noscript", func(e *colly.HTMLElement) {
		e.DOM.Remove()
	})

	c.OnHTML("body", func(e *colly.HTMLElement) {
		text := e.Text
		extractedText = append(extractedText, text)
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Fatal("Request failed:", err)
	})

	err := c.Visit(URL)
	if err != nil {
		log.Fatal(err)
	}

	finalText := strings.Join(extractedText, "\n")
	return finalText
}
