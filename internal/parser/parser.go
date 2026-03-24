// Package parser
package parser

import (
	"time"

	"manager/internal/model"
	"manager/pkg/llm"
)

type Parser interface {
	Parse(sms string) (model.Transaction, error)
}

type SMSParser struct {
	llmClient llm.CallLLM
}

func NewSMSParser(llmClient llm.CallLLM) Parser {
	return &SMSParser{llmClient: llmClient}
}

func (p *SMSParser) Parse(sms string) (model.Transaction, error) {
	response, err := p.llmClient.Call(sms)
	if err != nil {
		return model.Transaction{}, err
	}

	now := time.Now()
	transaction := model.Transaction{
		ID:          response.ID,
		Amount:      response.Amount,
		Date:        response.Date,
		Merchant:    response.Merchant,
		Credit:      response.Credit,
		Medium:      response.Medium,
		Category:    response.Category,
		Description: response.Description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	return transaction, nil
}
