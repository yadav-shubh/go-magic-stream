package googleai

import (
	"context"
	"sync"

	"github.com/yadav-shubh/go-magic-stream/config"
	"github.com/yadav-shubh/go-magic-stream/utils"
	"go.uber.org/zap"
	"google.golang.org/genai"
)

var (
	googleConfig = config.Get().Google
	once         sync.Once
	googleClient *genai.Client

	// Templates
	Movie_JSON_Template = `
							You are a JSON-only response agent.
							Every response must be valid JSON.
							Never use markdown. Never include explanations or extra text.
							
							Your task:
							Generate exactly 10 movie records for the best movie of 2025.
							
							Each record must match the following JSON format:
							
							{
							  "title": "",
							  "release_year": 2025,
							  "genre": "",
							  "rating": 0
							}
							
							Return an array of 10 objects in valid JSON.
							`
)

func init() {
	once.Do(func() {
		newClient, err := genai.NewClient(context.Background(), &genai.ClientConfig{APIKey: googleConfig.ApiKey})

		if err != nil {
			utils.Log.Error("GoogleAIClient", zap.Error(err))
			panic("Failed to create Google AI client")
		}
		googleClient = newClient
	})
}

func GenerateResponse(ctx context.Context, prompt string) (*genai.GenerateContentResponse, error) {
	utils.Log.Info("GenerateResponse called")

	generateContent, err := googleClient.Models.GenerateContent(
		ctx,
		googleConfig.Model,
		genai.Text(prompt),
		nil,
	)

	if err != nil {
		utils.Log.Error("GenerateContent", zap.Error(err))
		return nil, err
	}
	utils.Log.Info("GenerateContent", zap.Any("generateContent", generateContent))
	return generateContent, nil
}
