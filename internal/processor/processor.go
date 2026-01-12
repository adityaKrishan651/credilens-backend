package processor

import (
	"credilens-backend/internal/helpers"
	"strings"
)

func Process(rawText string) string {
	lines := strings.Split(rawText, "\n")
	var cleaned []string

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		if helpers.IsBoilerplate(line) {
			continue
		}

		cleaned = append(cleaned, line)
	}

	return strings.Join(cleaned, "\n")
}
