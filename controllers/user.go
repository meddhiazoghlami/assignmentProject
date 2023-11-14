package controllers

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	"assignmentProject/models"
	"assignmentProject/services"
)

func AddUser(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := models.User{}
		errs := ctx.BindJSON(&user)
		if errs != nil {
			ctx.AbortWithStatusJSON(400, gin.H{
				"error": errs.Error(),
			})
			return
		}
		user, err := services.AddUser(db, user)
		if err != nil {
			ctx.AbortWithStatusJSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(201, gin.H{
			"status": 201,
			"user":   user,
		})
	}
}
