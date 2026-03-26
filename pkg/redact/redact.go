package redact

import (
	"regexp"
	"strings"
)

var (
	emailPattern      = regexp.MustCompile(`(?i)\b[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}\b`)
	upiHandlePattern  = regexp.MustCompile(`(?i)\b[a-z0-9._%+\-]{2,}@[a-z0-9.\-]{2,}\b`)
	aadhaarPattern    = regexp.MustCompile(`\b\d{4}[\s-]?\d{4}[\s-]?\d{4}\b`)
	panPattern        = regexp.MustCompile(`\b[A-Z]{5}\d{4}[A-Z]\b`)
	ifscPattern       = regexp.MustCompile(`\b[A-Z]{4}0[A-Z0-9]{6}\b`)
	phonePattern      = regexp.MustCompile(`(?i)\b(?:\+?91[\s-]?)?[6-9]\d{9}\b`)
	longNumberPattern = regexp.MustCompile(`\b(?:\d[ -]?){12,19}\b`)
	labelPattern      = regexp.MustCompile(`(?i)\b(?:a\/c|acct(?:ount)?|acc(?:ount)?|card|card\s*no|account\s*no|loan\s*account|savings\s*account|current\s*account|reference|ref(?:erence)?|txn(?:\s*id)?|transaction\s*id|utr|rrn|auth(?:orization)?\s*code|upi\s*ref(?:erence)?|order\s*id)\b\s*[:#-]?\s*[A-Z0-9xX/\- ]{4,}`)
	otpPattern        = regexp.MustCompile(`(?i)\b(?:otp|pin|passcode|cvv|cvc)\b\s*[:#-]?\s*\d{3,8}\b`)
	maskedPattern     = regexp.MustCompile(`(?i)\b(?:x{2,}|x{4,}|\*{2,})[\s-]*\d{2,4}\b`)
)

func Redact(s string) string {
	if strings.TrimSpace(s) == "" {
		return ""
	}

	redacted := s
	replacements := []struct {
		pattern     *regexp.Regexp
		replacement string
	}{
		{emailPattern, "[redacted-email]"},
		{upiHandlePattern, "[redacted-upi]"},
		{panPattern, "[redacted-pan]"},
		{aadhaarPattern, "[redacted-aadhaar]"},
		{ifscPattern, "[redacted-ifsc]"},
		{otpPattern, "[redacted-secret]"},
		{maskedPattern, "[redacted-number]"},
		{labelPattern, "[redacted-reference]"},
		{phonePattern, "[redacted-phone]"},
		{longNumberPattern, "[redacted-number]"},
	}

	for _, replacement := range replacements {
		redacted = replacement.pattern.ReplaceAllString(redacted, replacement.replacement)
	}

	return strings.Join(strings.Fields(redacted), " ")
}
