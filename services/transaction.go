package services

import (
	"assignmentProject/models"
	"context"
	"database/sql"
	"errors"

	"github.com/shopspring/decimal"
)

func MakeDeposit(ctx context.Context, db *sql.DB, wallet_id string, amount decimal.Decimal) error {
	if amount.IsNegative() {
		return errors.New("Amount you provided is negative")
	}
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return errors.New("Something Went wrong.(Start transaction)")
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	sqlStatement := `UPDATE wallets SET balance = balance + $1 WHERE wallet_id = $2`

	_, err = tx.ExecContext(ctx, sqlStatement, amount, wallet_id)
	if err != nil {
		return errors.New("Something Went wrong.(DB:Update wallet)")
	}

	sqlStatement2 := `INSERT INTO transactions (tran_type,amount,wallet_id) VALUES ($1,$2,$3)`
	t_type := "deposit"
	_, err = tx.ExecContext(ctx, sqlStatement2, t_type, amount, wallet_id)
	if err != nil {
		return errors.New("Something Went wrong.(DB:Create transaction)")
	}

	if err = tx.Commit(); err != nil {
		return errors.New("Something Went wrong.(Commit Tx)")
	}
	return err
}

func MakeWithdraw(ctx context.Context, db *sql.DB, wallet_id string, amount decimal.Decimal) error {
	if amount.IsNegative() {
		return errors.New("Amount you provided is negative")
	}
	wallet := models.Wallet{}
	sqlSt := `SELECT * FROM wallets WHERE wallet_id = $1`
	err := db.QueryRow(sqlSt, wallet_id).Scan(&wallet.Wallet_id, &wallet.Created_date, &wallet.Balance, &wallet.Currency, &wallet.User_id)
	if err != nil {
		return errors.New("Something Went wrong.(DB)")
	}
	if amount.GreaterThan(wallet.Balance) {
		return errors.New("Not enough balance to perform your action.")
	}

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return errors.New("Something Went wrong.(Start transaction)")
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	sqlStatement := `UPDATE wallets SET balance = balance - $2 WHERE wallet_id = $1`

	_, err = tx.ExecContext(ctx, sqlStatement, wallet_id, amount)
	if err != nil {
		return errors.New("Something Went wrong.(DB:Update wallet)")
	}

	sqlStatement2 := `INSERT INTO transactions (tran_type,amount,wallet_id) VALUES ($1,$2,$3)`
	t_type := "withdraw"
	_, err = tx.ExecContext(ctx, sqlStatement2, t_type, amount, wallet_id)
	if err != nil {
		return errors.New("Something Went wrong.(DB:Create transaction)")
	}

	if err = tx.Commit(); err != nil {
		return errors.New("Something Went wrong.(Commit Tx)")
	}

	return err
}

func GetAllTransactions(ctx context.Context, db *sql.DB) ([]models.Transaction, error) {

	var trans []models.Transaction
	var allTrans []models.Transaction
	transaction := models.Transaction{}
	sqlStatement := `SELECT * FROM transactions`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return allTrans, errors.New("Something Went wrong.(DB)")
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&transaction.Tran_id, &transaction.Tran_type, &transaction.Amount, &transaction.Tran_date, &transaction.Wallet_id)
		if err != nil {
			return allTrans, errors.New("Something Went wrong.(DB)")
		}

		trans = append(trans, transaction)
	}
	if err := rows.Err(); err != nil {
		return allTrans, errors.New("Something Went wrong.(DB)")
	}
	return trans, err
}
