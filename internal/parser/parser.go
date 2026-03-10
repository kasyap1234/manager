// Package parser
package parser

import (
	"manager/internal/model"
	"manager/pkg/llm"
	"time"
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
	expense := model.Transaction{
		ID:          response.ID,
		Amount:      response.Amount,
		Date:        response.Date,
		Merchant:    response.Merchant,
		Credit:      response.Credit,
		Category:    response.Category,
		Description: response.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	return expense, nil
}
