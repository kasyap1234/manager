package parser

import (
	"manager/internal/model"
	"manager/pkg/llm"
)

type Parser interface {
	Parse(sms string) (model.Transaction, error)
}

type SMSParser struct {
	llmClient *llm.GeminiClient
}

func (p SMSParser) Parse(sms string) (model.Transaction, error) {
	// call llm to parse the sms
	//
	response, err := p.llmClient.Call(sms)
	if err != nil {
		return model.Transaction{}, err
	}
	transaction := model.Transaction{
		Amount:      response.Amount,
		Date:        response.Date,
		Description: response.Description,
	}
	return transaction, nil
}
