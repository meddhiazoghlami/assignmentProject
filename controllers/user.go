package controllers

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"

	"assignmentProject/models"
)

func AddUser(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := addUser(ctx, db)
		if err != nil {
			fmt.Println("err", err)
			return
		}
		ctx.JSON(201, gin.H{
			"status": 201,
			"user":   user,
		})
	}
}

// Service
func addUser(ctx *gin.Context, db *sql.DB) (models.User, error) {
	user := models.User{}
	ctx.BindJSON(&user)
	sqlStatement := `INSERT INTO users (username) VALUES ($1) RETURNING user_id`
	err := db.QueryRow(sqlStatement, user.Username).Scan(&user.User_id)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
	}

	return user, err
}
