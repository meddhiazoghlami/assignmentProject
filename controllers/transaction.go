package controllers

import (
	"database/sql"

	models "github.com/meddhiazoghlami/assignmentProject/models"
	services "github.com/meddhiazoghlami/assignmentProject/services"

	"github.com/gin-gonic/gin"
)

func MakeDeposit(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		wallet_id := ctx.Param("wallet_id")
		dpReq := models.DepositRequest{}
		err := ctx.BindJSON(&dpReq)
		if err != nil {
			ctx.AbortWithStatusJSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		errs := services.MakeDeposit(ctx.Request.Context(), db, wallet_id, dpReq.Amount)
		if errs != nil {
			ctx.AbortWithStatusJSON(500, gin.H{
				"error": errs.Error(),
			})
			return
		}
		ctx.JSON(200, gin.H{
			"status":  200,
			"message": "Your deposit is done successfully"})
	}
}

func MakeWithdraw(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		wallet_id := ctx.Param("wallet_id")
		dpReq := models.DepositRequest{}
		err := ctx.BindJSON(&dpReq)
		if err != nil {
			ctx.AbortWithStatusJSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		errs := services.MakeWithdraw(ctx.Request.Context(), db, wallet_id, dpReq.Amount)
		if errs != nil {
			ctx.AbortWithStatusJSON(500, gin.H{
				"error": errs.Error(),
			})
			return
		}
		ctx.JSON(200, gin.H{
			"status":  200,
			"message": "Your withdraw is done successfully"})
	}

}

func GetAllTransactions(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		trans, err := services.GetAllTransactions(ctx.Request.Context(), db)
		if err != nil {
			ctx.AbortWithStatusJSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, gin.H{
			"status":       200,
			"transactions": trans})
	}

}
