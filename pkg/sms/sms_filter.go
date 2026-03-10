package sms

import "strings"


positiveWords :=map[string]int{
	"credited": 3 ,
	"debited": 3,
	"spent": 3,
	"paid": 2,
	"received": 2,
	"withdrawn": 2,
	"deposited": 2,
	"upi": 3,
	"transferred": 2,
	"sent": 1,
	"emi": 1 ,
	"withdrawal": 2 ,

}


var negativeKeywords = map[string]int{
	"otp":              -5,
	"one time password": -5,
	"due":              -4,
	"statement":        -4,
	"minimum due":      -5,
	"credit card due":  -5,
	"offer":            -5,
	"sale":             -5,
	"discount":         -5,
	"promo":            -5,
	"reminder":         -4,
	"failed":           -4,
	"declined":         -4,
	"login":            -3,
	"kyc":              -3,
	"reward points":    -3,
	"welcome":          -2,
}


func isTransaction(sms string) bool {
	count :=0
	text :=strings.ToLower(sms)
	for _,word := range strings.Split(text," "){
		if _,ok := positiveWords[word]; ok{
			count+=positiveWords[word]
		}
		if _,ok := negativeKeywords[word]; ok{
			count+=negativeKeywords[word]
		}
	}
	return count > 0
}
