package models

import "github.com/shopspring/decimal"

type Transaction struct {
	Tran_id   string          `json:"tran_id"`
	Tran_type string          `json:"tran_type" binding:"required"`
	Amount    decimal.Decimal `json:"amount" binding:"required"`
	Tran_date string          `json:"tran_date" binding:"required"`
	Wallet_id string          `json:"wallet_id" binding:"required"`
}

type DepositRequest struct {
	Amount decimal.Decimal `json:"amount" binding:"required"`
}
