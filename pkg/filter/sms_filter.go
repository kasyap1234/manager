// Package filter
package filter

import (
	"regexp"
	"strings"
)

var (
	excludePatterns = []*regexp.Regexp{
		regexp.MustCompile(`(?i)\b(otp|otp\s*\d+)\b`),
		regexp.MustCompile(`(?i)\b(cdsl|nsdl)\b`),
		regexp.MustCompile(`(?i)\b(refund|reversal)\b.*?(pending|initiated|processing)`),
		regexp.MustCompile(`(?i)^.*?(loan|credit\s*card|payment)\s*due.*$`),
		regexp.MustCompile(`(?i)\b(failed|declined|unsuccessful)\b`),
		regexp.MustCompile(`(?i)\b(mpin|mpin\s*changed)\b`),
	}
	transactionIndicators = []*regexp.Regexp{
		regexp.MustCompile(`(?i)\b(debited|credited)\b`),
		regexp.MustCompile(`(?i)\b(spent|paid)\b`),
		regexp.MustCompile(`(?i)\b(upi\s*(payment|transfer|transaction|sent|received))\b`),
		regexp.MustCompile(`(?i)\b(imps|neft|rtgs|upi)\b`),
		regexp.MustCompile(`(?i)\b(card\s*(payment|purchase|transaction|swipe))\b`),
		regexp.MustCompile(`(?i)\b(wallet\s*(credit|debit|load|add))\b`),
		regexp.MustCompile(`(?i)\b(acct?|a/?c)\s*(no?|number)?\s*[:\-]?\s*[\dx]{8,}`),
		regexp.MustCompile(`(?i)\b(rs\.?|inr|₹)\s*[\d,]+(\.\d{2})?\b`),
		regexp.MustCompile(`(?i)\bagainst\s+your\s+(acct?|a/?c)\b`),
		regexp.MustCompile(`(?i)\b(transfer\s*(of|to|from))\b`),
		regexp.MustCompile(`(?i)\breceived\s+(?:rs\.?|inr|₹)\b`),
		regexp.MustCompile(`(?i)\bsent\s+to\b`),
		regexp.MustCompile(`(?i)\breceived\s+from\b`),
	}
)

func isTransactionSMS(sms string) bool {
	text := strings.TrimSpace(sms)
	if text == "" {
		return false
	}

	lowerText := strings.ToLower(text)

	for _, pattern := range excludePatterns {
		if pattern.MatchString(lowerText) {
			return false
		}
	}

	hasTransactionIndicator := false
	for _, pattern := range transactionIndicators {
		if pattern.MatchString(lowerText) {
			hasTransactionIndicator = true
			break
		}
	}

	return hasTransactionIndicator
}

func TransCheck(text string) bool {
	lowerText := strings.ToLower(text)
	transPatterns := []string{"debited", "spent", "paid", "sent", "deducted"}
	for _, pattern := range transPatterns {
		if strings.Contains(lowerText, pattern) {
			return true
		}
	}
	return false
}

func CreditCheck(text string) bool {
	lowerText := strings.ToLower(text)
	creditPatterns := []string{"credited", "received", "deposited", "refund", "credited to"}
	for _, pattern := range creditPatterns {
		if strings.Contains(lowerText, pattern) {
			return true
		}
	}
	return false
}

func isCDSL(text string) bool {
	lowerText := strings.ToLower(text)
	return strings.Contains(lowerText, "cdsl") || strings.Contains(lowerText, "nsdl")
}

func isOTP(text string) bool {
	lowerText := strings.ToLower(text)
	otpPattern := regexp.MustCompile(`(?i)\botp\b|\botp\s*\d{4,6}\b`)
	return otpPattern.MatchString(lowerText)
}
