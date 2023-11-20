package models

type User struct {
	User_id  string `json:"user_id"`
	Username string `json:"username" binding:"required"`
}

type UserWallets struct {
	User
	Wallets []Wallet
}
