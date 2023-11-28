package services

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/meddhiazoghlami/assignmentProject/db"
	"github.com/meddhiazoghlami/assignmentProject/models"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestGetBalance(t *testing.T) {
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

	user_id := "41582010-aefd-4a2b-a452-141f5688ff36"
	wallet_id := "4a40cb9b-fe20-470c-96b5-ec57f12970e2"

	got, err := GetBalance(db, wallet_id, user_id)
	if err != nil {
		t.Errorf("Test Failed with error %q", err)
	} else {
		t.Log("got:", got)
	}
	user_id2 := "aaaaaaaaaaaaaaaaa"
	wallet_id2 := "aaaaaaaaaaaaaaaa"
	_, err2 := GetBalance(db, wallet_id, user_id2)
	if err2 != nil {
		assert.Error(t, err2)
		assert.EqualError(t, err2, "The wallet of this user doesn't exist")

	}
	wallet3, err3 := GetBalance(db, wallet_id2, user_id)
	if err3 != nil {
		assert.Error(t, err3)
		assert.Empty(t, wallet3)
	}
	user_id3 := "41582010-aefd-4a2b-a452-141f5688ff36"
	wallet3, err4 := GetBalance(db, wallet_id, user_id3)
	if err4 != nil {
		assert.Error(t, err4)
		assert.EqualError(t, err2, "The wallet of this user doesn't exist")
		assert.Empty(t, wallet3)
	}
}

func TestAddWallet(t *testing.T) {
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
	user_id := "41582010-aefd-4a2b-a452-141f5688ff36"
	wallet1 := models.Wallet{
		Currency: "tnd",
	}
	wallet2, err := AddWallet(db, user_id, wallet1)
	balanceCompare := wallet2.Balance.Equal(decimal.NewFromInt(0))
	assert.NoError(t, err)
	assert.NotEmpty(t, wallet2)
	assert.Equal(t, wallet1.Currency, wallet2.Currency)
	assert.Equal(t, user_id, wallet2.User_id)
	assert.Equal(t, balanceCompare, true)

	user_id2 := "azert"
	_, err1 := AddWallet(db, user_id2, wallet1)
	if err1 != nil {
		assert.Error(t, err1)
	}
}

func TestGetWallet(t *testing.T) {
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
	user_id := "41582010-aefd-4a2b-a452-141f5688ff36"
	currency := models.Wallet{
		Currency: "eur",
	}
	wallet1, err := AddWallet(db, user_id, currency)
	assert.NoError(t, err)
	wallet2, err2 := GetWallet(db, wallet1.Wallet_id)
	assert.NoError(t, err2)
	assert.NotEmpty(t, wallet2)
	assert.Equal(t, wallet1.Wallet_id, wallet2.Wallet_id)
	assert.Equal(t, wallet1.User_id, wallet2.User_id)
	assert.Equal(t, wallet1.Currency, wallet2.Currency)
	assert.Equal(t, wallet1.Balance, wallet2.Balance)
	assert.Equal(t, wallet1.Created_date, wallet2.Created_date)

	_, wErr := GetWallet(db, "AZERkk")
	if wErr != nil {
		assert.Error(t, wErr)
	}
}

func TestGetUserWallets(t *testing.T) {
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
	user_id := "41582010-aefd-4a2b-a452-141f5688ff36"
	user, err := GetUserWallets(db, user_id)
	assert.NoError(t, err)
	assert.NotEmpty(t, user)
	assert.NotEmpty(t, user.Wallets)
}
