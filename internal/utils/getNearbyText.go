package utils

import (
	"strings"

	"github.com/gocolly/colly"
)

func GetNearbyText(e *colly.HTMLElement) string {
	parent := e.DOM.Parent()
	if parent == nil {
		return ""
	}

	text := strings.TrimSpace(parent.Text())
	if len(text) > 300 {
		text = text[:300]
	}

	return text
}
