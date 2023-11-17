package services

import (
	"assignmentProject/db"
	"assignmentProject/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddUser(t *testing.T) {
	db := db.BuildDBConfig("test")
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
