package services

import (
	"assignmentProject/models"
	"database/sql"
	"fmt"
)

func AddUser(db *sql.DB, user models.User) (models.User, error) {

	sqlStatement := `INSERT INTO users (username) VALUES ($1) RETURNING user_id`
	err := db.QueryRow(sqlStatement, user.Username).Scan(&user.User_id)
	if err != nil {
		fmt.Println("err", err)
	}

	return user, err
}
