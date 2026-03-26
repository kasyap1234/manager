package parser

import (
	"errors"
	"strings"
	"testing"
	"time"

	"manager/internal/model"
)

type fakeLLM struct {
	sms      string
	response model.Transaction
	err      error
	called   bool
}

func (f *fakeLLM) Call(sms string) (model.Transaction, error) {
	f.called = true
	f.sms = sms
	if f.err != nil {
		return model.Transaction{}, f.err
	}
	return f.response, nil
}

func TestParseRejectsNonTransactionSMS(t *testing.T) {
	llm := &fakeLLM{}
	parser := NewSMSParser(llm)

	_, err := parser.Parse("Your OTP is 123456 for login")
	if !errors.Is(err, ErrNonTransactionSMS) {
		t.Fatalf("expected ErrNonTransactionSMS, got %v", err)
	}

	if llm.called {
		t.Fatal("LLM should not be called for non-transaction SMS")
	}
}

func TestParseRedactsBeforeCallingLLM(t *testing.T) {
	llm := &fakeLLM{
		response: model.Transaction{
			ID:          "tx-1",
			Amount:      250.00,
			Date:        time.Now(),
			Merchant:    "Swiggy",
			Credit:      false,
			Medium:      "upi",
			Category:    "food",
			Description: "UPI payment to Swiggy",
		},
	}
	parser := NewSMSParser(llm)

	input := "Your A/c no 123456789012 is debited by Rs. 250.00 via UPI to john.doe@okicici. Avl Bal Rs. 1,249.00"
	_, err := parser.Parse(input)
	if err != nil {
		t.Fatalf("Parse returned error: %v", err)
	}

	if !llm.called {
		t.Fatal("LLM should have been called for transaction SMS")
	}

	for _, secret := range []string{
		"123456789012",
		"john.doe@okicici",
	} {
		if strings.Contains(llm.sms, secret) {
			t.Fatalf("LLM input still contains %q: %q", secret, llm.sms)
		}
	}

	for _, want := range []string{"debited", "Rs. 250.00", "UPI"} {
		if !strings.Contains(strings.ToLower(llm.sms), strings.ToLower(want)) {
			t.Fatalf("LLM input lost %q: %q", want, llm.sms)
		}
	}
}
