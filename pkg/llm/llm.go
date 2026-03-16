// Package llm
package llm

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

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

// Call sends the SMS text to the Gemini 3.1 Flash Lite model and
// parses the response into a Transaction struct.
func (c *GeminiClient) Call(sms string) (model.Transaction, error) {
	var tx model.Transaction

	if c == nil || c.client == nil {
		return tx, fmt.Errorf("gemini client is not initialized")
	}

	ctx := context.Background()

	// Prompt Gemini to return ONLY a JSON object matching model.Transaction.
	prompt := `
You are a financial SMS transaction parser.
if the sms is not a transaction, return an empty JSON object.

Given the following SMS describing a financial transaction, produce a JSON object
that matches EXACTLY the following Go struct (field names and types):

type Transaction struct {
  ID          string    // a unique identifier, or empty string if unknown
  Amount      float64   // transaction amount
  Date        time.Time // transaction date
  Merchant    string    // name of the merchant
  Credit      bool      // whether the transaction is a credit or debit
  Merchant    string    // name of the merchant
  Credit      bool      // whether the transaction is a credit or debit
  Category    string    // short category like "groceries", "restaurant"
  Description string    // concise description of the transaction
  CreatedAt   time.Time // creation timestamp
  UpdatedAt   time.Time // last update timestamp
}
This is additional context for relevant transactions : 
	raise securities is dhan app category : investment . 

Rules:
- All timestamps (date, created_at, updated_at) MUST be RFC3339 strings, e.g. "2006-01-02T15:04:05Z".
- If you don't know a value, make a reasonable guess rather than leaving it null.
- Use only these 7 fields and no others.

Respond with ONLY a single JSON object using these exact JSON keys:
id, amount, date, merchant, credit, category, description, created_at, updated_at.
Do not include any additional text before or after the JSON.

SMS:
` + sms

	parts := []*genai.Part{{Text: prompt}}

	resp, err := c.client.Models.GenerateContent(ctx, "gemini-3.1-flash-lite", []*genai.Content{
		{Parts: parts},
	}, nil)
	if err != nil {
		return tx, fmt.Errorf("calling gemini: %w", err)
	}

	if len(resp.Candidates) == 0 || resp.Candidates[0].Content == nil {
		return tx, fmt.Errorf("no candidates returned from gemini")
	}

	var b strings.Builder
	for _, part := range resp.Candidates[0].Content.Parts {
		if part.Text != "" {
			b.WriteString(part.Text)
		}
	}

	raw := strings.TrimSpace(b.String())
	if raw == "" {
		return tx, fmt.Errorf("empty response from gemini")
	}

	if err := json.Unmarshal([]byte(raw), &tx); err != nil {
		return tx, fmt.Errorf("unmarshal gemini JSON into Transaction: %w; raw=%s", err, raw)
	}

	return tx, nil
}
