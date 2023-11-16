package services

import (
	"assignmentProject/db"
	"assignmentProject/models"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestGetBalance(t *testing.T) {
	db := db.BuildDBConfig()
	defer db.Close()

	user_id := "41582010-aefd-4a2b-a452-141f5688ff36"
	wallet_id := "4a40cb9b-fe20-470c-96b5-ec57f12970e2"

	got, err := GetBalance(db, wallet_id, user_id)
	if err != nil {
		t.Errorf("Test Failed with error %q", err)
	} else {
		t.Log("got:", got)
	}
	// want := 5400
	// decimalValue := decimal.NewFromInt(int64(want))
	// if !got.Balance.Equal(decimalValue) {
	// 	t.Errorf("wanted: %q, got: %q", got.Balance, decimalValue)
	// }
}

func TestAddWallet(t *testing.T) {
	db := db.BuildDBConfig()
	defer db.Close()
	user_id := "41582010-aefd-4a2b-a452-141f5688ff36"
	wallet1 := models.Wallet{
		Currency: "tnd",
	}
	wallet2, err := AddWallet(db, user_id, wallet1)
	balanceCompare := wallet2.Balance.Equal(decimal.NewFromInt(0))
	require.NoError(t, err)
	require.NotEmpty(t, wallet2)
	require.Equal(t, wallet1.Currency, wallet2.Currency)
	require.Equal(t, user_id, wallet2.User_id)
	require.Equal(t, balanceCompare, true)
}

func TestGetWallet(t *testing.T) {
	db := db.BuildDBConfig()
	defer db.Close()
	user_id := "41582010-aefd-4a2b-a452-141f5688ff36"
	currency := models.Wallet{
		Currency: "eur",
	}
	wallet1, err := AddWallet(db, user_id, currency)
	require.NoError(t, err)
	wallet2, err2 := GetWallet(db, wallet1.Wallet_id)
	require.NoError(t, err2)
	require.NotEmpty(t, wallet2)
	require.Equal(t, wallet1.Wallet_id, wallet2.Wallet_id)
	require.Equal(t, wallet1.User_id, wallet2.User_id)
	require.Equal(t, wallet1.Currency, wallet2.Currency)
	require.Equal(t, wallet1.Balance, wallet2.Balance)
	require.Equal(t, wallet1.Created_date, wallet2.Created_date)
}
