package controllers

import (
	"assignmentProject/models"
	"database/sql"
	"errors"

	"github.com/gin-gonic/gin"
)

func MakeDeposit(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := makeDeposit(ctx, db)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		ctx.JSON(200, gin.H{"message": "Your deposit is done successfully"})
	}

}

func MakeWithdraw(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := makeWithdraw(ctx, db)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}
		ctx.JSON(200, gin.H{"message": "Your withdraw is done successfully"})
	}

}

func GetAllTransactions(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		trans, err := getAllTransactions(ctx, db)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}
		ctx.JSON(200, gin.H{"transactions": trans})
	}

}

// Services

func makeDeposit(ctx *gin.Context, db *sql.DB) error {
	wallet_id := ctx.Param("wallet_id")
	depositReq := models.DepositRequest{}
	err := ctx.BindJSON(&depositReq)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return err
	}
	if depositReq.Amount.IsNegative() {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": "Amount should be higher than 0",
		})
		return err
	}
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": "Something went wrong",
		})
		return err
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	sqlStatement := `UPDATE wallets SET balance = balance + $1 WHERE wallet_id = $2`

	_, err = tx.ExecContext(ctx, sqlStatement, depositReq.Amount, wallet_id)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"error": "Something went wrong",
		})
		return err
	}

	sqlStatement2 := `INSERT INTO transactions (tran_type,amount,wallet_id) VALUES ($1,$2,$3)`
	t_type := "deposit"
	_, err = tx.ExecContext(ctx, sqlStatement2, t_type, depositReq.Amount, wallet_id)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"error": "Something went wrong",
		})
		return err
	}

	if err = tx.Commit(); err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"error": "Something went wrong",
		})
		return err
	}

	return err
}

func makeWithdraw(ctx *gin.Context, db *sql.DB) error {
	wallet_id := ctx.Param("wallet_id")
	depositReq := models.DepositRequest{}
	if depositReq.Amount.IsNegative() {
		return errors.New("Amount should be higher than 0")
	}
	wallet := models.Wallet{}
	err := ctx.BindJSON(&depositReq)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return err
	}
	sqlSt := `SELECT * FROM wallets WHERE wallet_id = $1`
	err = db.QueryRow(sqlSt, wallet_id).Scan(&wallet.Wallet_id, &wallet.Created_date, &wallet.Balance, &wallet.Currency, &wallet.User_id)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
		return err
	}
	if depositReq.Amount.GreaterThan(wallet.Balance) {
		return errors.New("You Dont have enough money")
	}

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
		return err
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	sqlStatement := `UPDATE wallets SET balance = balance - $2 WHERE wallet_id = $1`

	_, err = tx.ExecContext(ctx, sqlStatement, wallet_id, depositReq.Amount)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
		return err
	}

	sqlStatement2 := `INSERT INTO transactions (tran_type,amount,wallet_id) VALUES ($1,$2,$3)`
	t_type := "withdraw"
	_, err = tx.ExecContext(ctx, sqlStatement2, t_type, depositReq.Amount, wallet_id)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
		return err
	}

	if err = tx.Commit(); err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
		return err
	}

	return err
}

func getAllTransactions(ctx *gin.Context, db *sql.DB) ([]models.Transaction, error) {

	var trans []models.Transaction
	var allTrans []models.Transaction
	transaction := models.Transaction{}
	sqlStatement := `SELECT * FROM transactions`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		ctx.JSON(500, err.Error())
		return trans, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&transaction.Tran_id, &transaction.Tran_type, &transaction.Amount, &transaction.Tran_date, &transaction.Wallet_id)
		if err != nil {
			ctx.JSON(500, err.Error())
			return allTrans, err
		}

		trans = append(trans, transaction)
	}
	if err := rows.Err(); err != nil {
		ctx.JSON(500, err.Error())
		return allTrans, err
	}
	return trans, err

}
