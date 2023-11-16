package services

import (
	"assignmentProject/models"
	"database/sql"
	"errors"
)

func AddWallet(db *sql.DB, user_id string, wallet models.Wallet) (models.Wallet, error) {

	sqlStatement := `INSERT INTO wallets (currency,user_id) VALUES ($1,$2) RETURNING *`
	err := db.QueryRow(sqlStatement, wallet.Currency, user_id).Scan(&wallet.Wallet_id, &wallet.Created_date, &wallet.Balance, &wallet.Currency, &wallet.User_id)
	if err != nil {
		return wallet, err
	}
	return wallet, nil
}

func GetWallet(db *sql.DB, wallet_id string) (models.Wallet, error) {
	wallet := models.Wallet{}
	sqlStatement := `SELECT * FROM wallets WHERE wallet_id = $1`
	err := db.QueryRow(sqlStatement, wallet_id).Scan(&wallet.Wallet_id, &wallet.Created_date, &wallet.Balance, &wallet.Currency, &wallet.User_id)
	if err != nil {
		return wallet, err
	}
	return wallet, err
}

func GetBalance(db *sql.DB, wallet_id string, user_id string) (models.Wallet, error) {

	wallet := models.Wallet{}
	sqlStatement := `SELECT * FROM wallets WHERE wallet_id = $1`
	err := db.QueryRow(sqlStatement, wallet_id).Scan(&wallet.Wallet_id, &wallet.Created_date, &wallet.Balance, &wallet.Currency, &wallet.User_id)
	if err != nil {
		return wallet, err
	}
	if user_id != wallet.User_id {
		return wallet, errors.New("The wallet of this user doesn't exist")
	}
	return wallet, err
}
