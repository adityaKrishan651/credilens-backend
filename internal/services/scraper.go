package services

import (
	"github.com/gocolly/colly"
)

func Scrape() {

	c := colly.NewCollector(
		colly.AllowedDomains("www.scrapingcourse.com"),
	)

	c.OnRequest(func(r *colly.Request) {
		// before http request
	})

	c.OnResponse(func(r *colly.Response) {
		// once site is reached
	})

}
