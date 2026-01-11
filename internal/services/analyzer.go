package services

import (
	"credilens-backend/internal/models"

	"context"
	"log"

	"google.golang.org/genai"
)

func AnalyzeContent(content string) models.AnalyzeResponse {
	ctx := context.Background()

	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	config := &genai.GenerateContentConfig{
		ResponseMIMEType: `application/json`,
	}

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.5-flash",
		genai.Text(content),
		config,
	)
	if err != nil {
		log.Fatal(err)
	}

	return models.AnalyzeResponse{
		RiskLevel: "Medium",
		Summary:   "successfull response",
		Signals: []string{
			result.Text(),
		},
	}

}
