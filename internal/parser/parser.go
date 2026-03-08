package parser

import (
	"manager/internal/model"
	"manager/pkg/llm"
	"time"
)

type Parser interface {
	Parse(sms string) (model.Expense, error)
}

type SMSParser struct {
	llmClient llm.CallLLM
}

func NewSMSParser(llmClient llm.CallLLM) Parser {
	return &SMSParser{llmClient: llmClient}
}

func (p *SMSParser) Parse(sms string) (model.Expense, error) {
	response, err := p.llmClient.Call(sms)
	if err != nil {
		return model.Expense{}, err
	}
	expense := model.Expense{
		Amount:      response.Amount,
		Date:        response.Date,
		Category:    response.Category,
		Description: response.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	return expense, nil
}
