package chunker

import (
	"strings"
)

func Split(cleanText string, chunkSize int, overlap int) []string {
	if chunkSize <= 0 || overlap >= chunkSize {
		return nil
	}

	tokens := strings.Fields(cleanText)
	var chunks []string

	index := 0
	for index < len(tokens) {
		end := index + chunkSize
		if end > len(tokens) {
			end = len(tokens)
		}

		chunkTokens := tokens[index:end]
		chunks = append(chunks, strings.Join(chunkTokens, " "))

		index += chunkSize - overlap
	}

	return chunks
}
