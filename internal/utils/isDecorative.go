package utils

import (
	"strings"

	"github.com/gocolly/colly"
)

func IsDecorative(e *colly.HTMLElement) bool {
	alt := strings.TrimSpace(strings.ToLower(e.Attr("alt")))
	class := strings.ToLower(e.Attr("class"))

	if alt == "" {
		return true
	}

	decorativeKeywords := []string{
		"icon", "logo", "sprite", "button", "arrow",
	}

	for _, k := range decorativeKeywords {
		if strings.Contains(class, k) || strings.Contains(alt, k) {
			return true
		}
	}

	return false
}
