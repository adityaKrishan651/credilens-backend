package utils

import "strings"

func IsBoilerplate(line string) bool {
	lower := strings.ToLower(line)

	boilerplate := []string{
		"cookie",
		"privacy policy",
		"terms of service",
		"sign in",
		"log in",
		"subscribe",
		"©",
	}

	for _, phrase := range boilerplate {
		if strings.Contains(lower, phrase) {
			return true
		}
	}

	// discard very short UI fragments
	if len(line) < 40 {
		return true
	}

	return false
}
