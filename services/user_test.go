package services

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/meddhiazoghlami/assignmentProject/db"
	"github.com/meddhiazoghlami/assignmentProject/models"

	"github.com/stretchr/testify/assert"
)

func TestAddUser(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	dbconfig := db.DBConfig{
		Host:     os.Getenv("DB_HOST_TEST"),
		Port:     os.Getenv("DB_PORT_TEST"),
		User:     os.Getenv("DB_USER_TEST"),
		Password: os.Getenv("DB_PASSWORD_TEST"),
		Dbname:   os.Getenv("DB_NAME_TEST"),
	}
	db := db.BuildDBConfig(dbconfig)
	defer db.Close()

	username := models.User{
		Username: "Test User",
	}

	user, err := AddUser(db, username)
	assert.NoError(t, err)
	assert.NotEmpty(t, user)
	assert.Equal(t, user.Username, user.Username)

	username2 := models.User{
		Username: "",
	}

	user2, err2 := AddUser(db, username2)

	assert.Error(t, err2)
	assert.Empty(t, user2)

}
