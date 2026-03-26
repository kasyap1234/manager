package redact

import (
	"strings"
	"testing"
)

func TestRedactMasksSensitiveData(t *testing.T) {
	input := "A/c no 123456789012 debited Rs. 1,234.50 via UPI to john.doe@okicici. Call +91 9876543210 or email support@example.com. PAN ABCDE1234F IFSC HDFC0001234 OTP 123456"

	got := Redact(input)

	for _, secret := range []string{
		"123456789012",
		"john.doe@okicici",
		"9876543210",
		"support@example.com",
		"ABCDE1234F",
		"HDFC0001234",
		"123456",
	} {
		if strings.Contains(got, secret) {
			t.Fatalf("redacted output still contains %q: %q", secret, got)
		}
	}

	for _, want := range []string{"debited", "Rs. 1,234.50", "UPI"} {
		if !strings.Contains(strings.ToLower(got), strings.ToLower(want)) {
			t.Fatalf("redacted output lost %q: %q", want, got)
		}
	}
}

func TestRedactIsIdempotent(t *testing.T) {
	input := "Your account 123456789012 was debited via UPI to merchant@example.com"

	first := Redact(input)
	second := Redact(first)

	if first != second {
		t.Fatalf("redaction should be idempotent: first=%q second=%q", first, second)
	}
}
