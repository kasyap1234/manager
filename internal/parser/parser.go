// Package parser
package parser

import (
	"errors"
	"time"

	"manager/internal/model"
	"manager/pkg/filter"
	"manager/pkg/redact"
)

type Parser interface {
	Parse(sms string) (model.Transaction, error)
}

type CallLLM interface {
	Call(sms string) (model.Transaction, error)
}

var ErrNonTransactionSMS = errors.New("sms does not look like a transaction")

type SMSParser struct {
	llmClient CallLLM
}

func NewSMSParser(llmClient CallLLM) Parser {
	return &SMSParser{llmClient: llmClient}
}

func (p *SMSParser) Parse(sms string) (model.Transaction, error) {
	if !filter.IsTransactionSMS(sms) {
		return model.Transaction{}, ErrNonTransactionSMS
	}

	sanitized := redact.Redact(sms)
	response, err := p.llmClient.Call(sanitized)
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
