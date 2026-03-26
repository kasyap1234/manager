package filter

import "testing"

func TestIsTransactionSMS(t *testing.T) {
	tests := []struct {
		name string
		sms  string
		want bool
	}{
		{
			name: "upi debit",
			sms:  "Your A/c XX1234 is debited by INR 250.00 via UPI to Swiggy. Avl Bal INR 1,249.00",
			want: true,
		},
		{
			name: "imps credit",
			sms:  "INR 15,000 credited to your account 123456789012 via IMPS from John",
			want: true,
		},
		{
			name: "card purchase",
			sms:  "Spent Rs. 899.00 on your HDFC Bank Credit Card at Zomato",
			want: true,
		},
		{
			name: "atm cash withdrawal",
			sms:  "Cash withdrawn from ATM Rs. 2000 from your account ending 1234",
			want: true,
		},
		{
			name: "otp message",
			sms:  "Your OTP is 123456 for SBI login",
			want: false,
		},
		{
			name: "statement message",
			sms:  "Your monthly e-statement is ready for download",
			want: false,
		},
		{
			name: "kyc reminder",
			sms:  "KYC update pending for your account",
			want: false,
		},
		{
			name: "marketing offer",
			sms:  "Great offers on personal loans with instant approval",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsTransactionSMS(tt.sms)
			if got != tt.want {
				t.Fatalf("IsTransactionSMS(%q) = %v, want %v", tt.sms, got, tt.want)
			}
		})
	}
}

func TestTransAndCreditChecks(t *testing.T) {
	if !TransCheck("debited Rs. 100 from your account") {
		t.Fatal("expected TransCheck to match debit message")
	}

	if !CreditCheck("credited Rs. 100 to your account") {
		t.Fatal("expected CreditCheck to match credit message")
	}
}
