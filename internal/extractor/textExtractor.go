package extractor

import "strings"

func ExtractText(bodyText string) string {
	lines := strings.Split(bodyText, "\n")
	var cleaned []string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		cleaned = append(cleaned, line)
	}

	return strings.Join(cleaned, "\n")
}
