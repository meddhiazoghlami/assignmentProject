package services

import (
	"assignmentProject/db"
	"assignmentProject/models"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddUser(t *testing.T) {
	db := db.BuildDBConfig()
	defer db.Close()

	username := models.User{
		Username: "Test User",
	}

	user, err := AddUser(db, username)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, user.Username, user.Username)
}
