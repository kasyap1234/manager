// Package model
package model

import "time"

type Transaction struct {
	ID          string    `json:"id"`
	Amount      float64   `json:"amount"`
	Date        time.Time `json:"date"`
	Merchant    string    `json:"merchant"`
	Credit      bool      `json:"credit"`
	Medium      string    `json:"medium"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
