package controllers

import (
	"assignmentProject/models"
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
)

func AddWallet(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		wallet, err := addWallet(ctx, db)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}
		ctx.JSON(200, gin.H{"data": wallet})

	}
}

func GetWallet(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		wallet, err := getWallet(ctx, db)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}
		ctx.JSON(200, gin.H{"status": "ok",
			"data": wallet,
		})

	}
}

func GetBalance(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		wallet, err := getBalance(ctx, db)

		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		ctx.JSON(200, gin.H{"balance": wallet.Balance})

	}
}

//Services

func addWallet(ctx *gin.Context, db *sql.DB) (models.Wallet, error) {
	user_id := ctx.Param("id")
	wallet := models.Wallet{}
	ctx.BindJSON(&wallet)
	sqlStatement := `INSERT INTO wallets (currency,user_id) VALUES ($1,$2) RETURNING wallet_id`
	err := db.QueryRow(sqlStatement, wallet.Currency, user_id).Scan(&wallet.Wallet_id)

	return wallet, err
}

func getWallet(ctx *gin.Context, db *sql.DB) (models.Wallet, error) {
	user_id := ctx.Param("id")
	fmt.Println(user_id)
	wallet_id := ctx.Param("wallet_id")
	wallet := models.Wallet{}
	sqlStatement := `SELECT * FROM wallets WHERE wallet_id = $1`
	err := db.QueryRow(sqlStatement, wallet_id).Scan(&wallet.Wallet_id, &wallet.Created_date, &wallet.Balance, &wallet.Currency, &wallet.User_id)

	return wallet, err
}

func getBalance(ctx *gin.Context, db *sql.DB) (models.Wallet, error) {
	wallet_id := ctx.Param("wallet_id")
	wallet := models.Wallet{}
	sqlStatement := `SELECT * FROM wallets WHERE wallet_id = $1`
	err := db.QueryRow(sqlStatement, wallet_id).Scan(&wallet.Wallet_id, &wallet.Created_date, &wallet.Balance, &wallet.Currency, &wallet.User_id)

	return wallet, err
}
