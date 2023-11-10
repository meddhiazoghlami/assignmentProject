package controllers

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	"assignmentProject/models"
)

func AddUser(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := addUser(ctx, db)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}
		ctx.JSON(200, gin.H{
			"data": user,
		})
	}
}

// Service
func addUser(ctx *gin.Context, db *sql.DB) (models.User, error) {
	user := models.User{}
	ctx.BindJSON(&user)
	sqlStatement := `INSERT INTO users (username) VALUES ($1) RETURNING user_id`
	err := db.QueryRow(sqlStatement, user.Username).Scan(&user.User_id)

	return user, err
}
