package services

import (
	"assignmentProject/models"
	"database/sql"
	"fmt"
)

func AddWallet(db *sql.DB, user_id string, wallet models.Wallet) (models.Wallet, error) {

	sqlStatement := `INSERT INTO wallets (currency,user_id) VALUES ($1,$2) RETURNING wallet_id`
	err := db.QueryRow(sqlStatement, wallet.Currency, user_id).Scan(&wallet.Wallet_id)
	if err != nil {
		fmt.Println("err", err)
		return wallet, err
	}
	return wallet, nil
}

func GetWallet(db *sql.DB, wallet_id string) (models.Wallet, error) {
	wallet := models.Wallet{}
	sqlStatement := `SELECT * FROM wallets WHERE wallet_id = $1`
	err := db.QueryRow(sqlStatement, wallet_id).Scan(&wallet.Wallet_id, &wallet.Created_date, &wallet.Balance, &wallet.Currency, &wallet.User_id)
	if err != nil {
		fmt.Println("error", err)
		return wallet, err
	}
	return wallet, err
}

func GetBalance(db *sql.DB, wallet_id string) (models.Wallet, error) {

	wallet := models.Wallet{}
	sqlStatement := `SELECT * FROM wallets WHERE wallet_id = $1`
	err := db.QueryRow(sqlStatement, wallet_id).Scan(&wallet.Wallet_id, &wallet.Created_date, &wallet.Balance, &wallet.Currency, &wallet.User_id)
	if err != nil {
		fmt.Println("err", err)
		return wallet, err
	}
	return wallet, err
}
