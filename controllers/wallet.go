package controllers

import (
	"assignmentProject/models"
	"assignmentProject/services"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func AddWallet(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user_id := ctx.Param("id")
		wallet := models.Wallet{}
		errs := ctx.BindJSON(&wallet)
		if errs != nil {
			ctx.AbortWithStatusJSON(400, gin.H{
				"status":  400,
				"message": errs.Error(),
			})
		}
		wallet, err := services.AddWallet(db, user_id, wallet)
		if err != nil {
			ctx.AbortWithStatusJSON(500, gin.H{
				"status":  500,
				"message": err.Error(),
			})
		}
		ctx.JSON(201, gin.H{
			"status": 201,
			"wallet": wallet})
	}
}

func GetWallet(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		wallet_id := ctx.Param("wallet_id")

		wallet, err := services.GetWallet(db, wallet_id)
		if err != nil {
			ctx.AbortWithStatusJSON(500, gin.H{
				"status":  500,
				"message": err.Error(),
			})
		}
		ctx.JSON(200, gin.H{"status": "ok",
			"data": wallet,
		})
	}
}

func GetBalance(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user_id := ctx.Param("id")
		wallet_id := ctx.Param("wallet_id")
		wallet, err := services.GetBalance(db, wallet_id, user_id)
		if err != nil {
			ctx.AbortWithStatusJSON(500, gin.H{
				"status":  500,
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(200, gin.H{"balance": wallet.Balance})

	}
}
