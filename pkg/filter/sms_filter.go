package filter

import (
	"regexp"
	"strings"
)

var (
	strictNonTransactionPatterns = []*regexp.Regexp{
		regexp.MustCompile(`(?i)\b(?:otp|passcode|pin|cvv|cvc|verification\s*code|one\s*time\s*password)\b`),
		regexp.MustCompile(`(?i)\b(?:login|logged\s*in|sign\s*in|sign-in|new\s*device|device\s*change)\b`),
		regexp.MustCompile(`(?i)\b(?:offer|offers|sale|discount|cashback|reward|rewards|win|lucky\s*draw|contest|promotion|promo|coupon)\b`),
		regexp.MustCompile(`(?i)\b(?:kyc|re-?kyc|pan\s*update|aadhaar\s*update|profile\s*update)\b`),
		regexp.MustCompile(`(?i)\b(?:monthly\s*statement|mini\s*statement|e-?statement|statement\s*is\s*ready|account\s*summary)\b`),
		regexp.MustCompile(`(?i)\b(?:loan\s*due|emi\s*due|premium\s*due|subscription\s*renewal|bill\s*due|payment\s*reminder|credit\s*score)\b`),
		regexp.MustCompile(`(?i)\b(?:survey|feedback|download\s*the\s*app|install\s*the\s*app|refer\s*and\s*earn)\b`),
	}
	strongTransactionPatterns = []*regexp.Regexp{
		regexp.MustCompile(`(?i)\b(?:debited|credited|deducted|withdrawn|withdrawal|transferred|transfer|paid|spent|refunded|reversal|reversed|received|deposited|charged|purchase|payment|sent)\b`),
		regexp.MustCompile(`(?i)\b(?:upi|imps|neft|rtgs|atm|card|wallet|net\s*banking|netbanking|ach|ecs|nach|cash)\b`),
	}
	transactionContextPatterns = []*regexp.Regexp{
		regexp.MustCompile(`(?i)\b(?:rs\.?|inr|₹)\s*[\d,]+(?:\.\d{2})?\b`),
		regexp.MustCompile(`(?i)\b(?:a\/c|acct?|account|acct\.|card)\b`),
		regexp.MustCompile(`(?i)\b(?:available\s*balance|avl\.?\s*bal|ledger\s*balance|balance\s*(?:after|is|left))\b`),
		regexp.MustCompile(`(?i)\b(?:txn|transaction|ref|rrn|utr)\b`),
		regexp.MustCompile(`(?i)\b(?:merchant|payee|beneficiary|recipient|sender|from\s+your|to\s+your|at\s+[a-z])\b`),
	}
)

func IsTransactionSMS(sms string) bool {
	text := normalizeSMS(sms)
	if text == "" {
		return false
	}

	txScore := 0
	txScore += scoreMatches(text, strongTransactionPatterns, 2)
	txScore += scoreMatches(text, transactionContextPatterns, 1)

	if txScore == 0 {
		return false
	}

	nonTxScore := scoreMatches(text, strictNonTransactionPatterns, 2)
	if nonTxScore >= 4 && txScore < 5 {
		return false
	}
	if nonTxScore > txScore && txScore < 4 {
		return false
	}

	return txScore >= 3
}

func TransCheck(text string) bool {
	normalized := normalizeSMS(text)
	if normalized == "" {
		return false
	}

	verbPattern := regexp.MustCompile(`(?i)\b(?:debited|credited|paid|spent|sent|deducted|transferred|received)\b`)
	return verbPattern.MatchString(normalized)
}

func CreditCheck(text string) bool {
	normalized := normalizeSMS(text)
	if normalized == "" {
		return false
	}

	creditPattern := regexp.MustCompile(`(?i)\b(?:credited|received|deposited|refund|reversal|reversed)\b`)
	return creditPattern.MatchString(normalized)
}

func normalizeSMS(s string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(strings.ToLower(s))), " ")
}

func scoreMatches(text string, patterns []*regexp.Regexp, weight int) int {
	score := 0
	for _, pattern := range patterns {
		if pattern.MatchString(text) {
			score += weight
		}
	}
	return score
}
