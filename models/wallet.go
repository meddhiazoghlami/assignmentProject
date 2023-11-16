package models

import "github.com/shopspring/decimal"

type Wallet struct {
	Wallet_id    string          `json:"wallet_id"`
	Created_date string          `json:"created_date"`
	Balance      decimal.Decimal `json:"balance"`
	Currency     string          `json:"currency" binding:"required"`
	User_id      string          `json:"user_id"`
}
