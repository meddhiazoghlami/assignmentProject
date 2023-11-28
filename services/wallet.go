package services

import (
	"database/sql"
	"errors"

	models "github.com/meddhiazoghlami/assignmentProject/models"
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

func GetUserWallets(db *sql.DB, user_id string) (models.UserWallets, error) {
	userWallets := models.UserWallets{}
	wallet := models.Wallet{}
	sqlGetUser := `SELECT * FROM users WHERE user_id = $1`
	err := db.QueryRow(sqlGetUser, user_id).Scan(&userWallets.User_id, &userWallets.Username)
	if err != nil {
		return userWallets, err
	}
	sqlGetWallets := `SELECT * FROM wallets where user_id = $1`
	rows, err := db.Query(sqlGetWallets, user_id)
	if err != nil {
		return userWallets, errors.New("Cannot fetch wallets")
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&wallet.Wallet_id, &wallet.Currency, &wallet.Balance, &wallet.Created_date, &wallet.User_id)
		if err != nil {
			return userWallets, errors.New("Can't scan wallet")
		}
		userWallets.Wallets = append(userWallets.Wallets, wallet)
	}
	if err := rows.Err(); err != nil {
		return userWallets, errors.New("Something Went wrong.(DB ROWS)")
	}
	return userWallets, err
}
