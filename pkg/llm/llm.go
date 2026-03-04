package llm

import (
	"context"

	"manager/internal/model"

	"google.golang.org/genai"
)

type CallLLM interface {
	Call(sms string) (model.Transaction, error)
}

type GeminiClient struct {
	client *genai.Client
}

func NewGeminiClient(apiKey string) *GeminiClient {
	client, err := genai.NewClient(context.Background(), &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return nil
	}
	return &GeminiClient{
		client: client,
	}
}

func (c *GeminiClient) Call(sms string) (model.Transaction, error) {
}
