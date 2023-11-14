package services

import (
	"assignmentProject/models"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/shopspring/decimal"
)

// Services

func MakeDeposit(ctx context.Context, db *sql.DB, wallet_id string, amount decimal.Decimal) error {
	if amount.IsNegative() {
		ClientErr := errors.New("Amount you provided is negative")
		fmt.Println("err", ClientErr)
		return ClientErr
	}
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		ServerErr := errors.New("Something Went wrong.(Start transaction)")
		fmt.Println("err", ServerErr)
		return ServerErr
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	sqlStatement := `UPDATE wallets SET balance = balance + $1 WHERE wallet_id = $2`

	_, err = tx.ExecContext(ctx, sqlStatement, amount, wallet_id)
	if err != nil {
		ServerErr := errors.New("Something Went wrong.(DB:Update wallet)")
		fmt.Println("err", ServerErr)
		return ServerErr
	}

	sqlStatement2 := `INSERT INTO transactions (tran_type,amount,wallet_id) VALUES ($1,$2,$3)`
	t_type := "deposit"
	_, err = tx.ExecContext(ctx, sqlStatement2, t_type, amount, wallet_id)
	if err != nil {
		ServerErr := errors.New("Something Went wrong.(DB:Create transaction)")
		fmt.Println("err", ServerErr)
		return ServerErr
	}

	if err = tx.Commit(); err != nil {
		ServerErr := errors.New("Something Went wrong.(Commit Tx)")
		fmt.Println("err", ServerErr)
		return ServerErr
	}
	return err
}

func MakeWithdraw(ctx context.Context, db *sql.DB, wallet_id string, amount decimal.Decimal) error {
	if amount.IsNegative() {
		ClientErr := errors.New("Amount you provided is negative")
		fmt.Println("err", ClientErr)
		return ClientErr
	}
	wallet := models.Wallet{}
	sqlSt := `SELECT * FROM wallets WHERE wallet_id = $1`
	err := db.QueryRow(sqlSt, wallet_id).Scan(&wallet.Wallet_id, &wallet.Created_date, &wallet.Balance, &wallet.Currency, &wallet.User_id)
	if err != nil {
		ServerErr := errors.New("Something Went wrong.(DB)")
		fmt.Println("err", ServerErr)
		return ServerErr
	}
	if amount.GreaterThan(wallet.Balance) {
		ClientErr := errors.New("Not enough balance to perform your action.")
		fmt.Println("err", ClientErr)
		return ClientErr
	}

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		ServerErr := errors.New("Something Went wrong.(Start transaction)")
		fmt.Println("err", ServerErr)
		return ServerErr
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	sqlStatement := `UPDATE wallets SET balance = balance - $2 WHERE wallet_id = $1`

	_, err = tx.ExecContext(ctx, sqlStatement, wallet_id, amount)
	if err != nil {
		ServerErr := errors.New("Something Went wrong.(DB:Update wallet)")
		fmt.Println("err", ServerErr)
		return ServerErr
	}

	sqlStatement2 := `INSERT INTO transactions (tran_type,amount,wallet_id) VALUES ($1,$2,$3)`
	t_type := "withdraw"
	_, err = tx.ExecContext(ctx, sqlStatement2, t_type, amount, wallet_id)
	if err != nil {
		ServerErr := errors.New("Something Went wrong.(DB:Create transaction)")
		fmt.Println("err", ServerErr)
		return ServerErr
	}

	if err = tx.Commit(); err != nil {
		ServerErr := errors.New("Something Went wrong.(Commit Tx)")
		fmt.Println("err", ServerErr)
		return ServerErr
	}

	return errors.New("In")
}

func GetAllTransactions(ctx context.Context, db *sql.DB) ([]models.Transaction, error) {

	var trans []models.Transaction
	var allTrans []models.Transaction
	transaction := models.Transaction{}
	sqlStatement := `SELECT * FROM transactions`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		ServerErr := errors.New("Something Went wrong.(DB)")
		fmt.Println("err", ServerErr)
		return trans, ServerErr
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&transaction.Tran_id, &transaction.Tran_type, &transaction.Amount, &transaction.Tran_date, &transaction.Wallet_id)
		if err != nil {
			ServerErr := errors.New("Something Went wrong.(DB)")
			fmt.Println("err", ServerErr)
			return allTrans, ServerErr
		}

		trans = append(trans, transaction)
	}
	if err := rows.Err(); err != nil {
		ServerErr := errors.New("Something Went wrong.(DB)")
		fmt.Println("err", ServerErr)
		return allTrans, ServerErr
	}
	return trans, err
}
