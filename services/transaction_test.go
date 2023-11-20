package services

import (
	"assignmentProject/db"
	"context"
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestMakeDeposit(t *testing.T) {
	inputs := []float64{78.091, -67}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	db := db.BuildDBConfig("test")
	defer db.Close()
	user_id := "41582010-aefd-4a2b-a452-141f5688ff36"
	wallet_id := "4a40cb9b-fe20-470c-96b5-ec57f12970e2"

	for _, amount := range inputs {
		wallet, bErr := GetBalance(db, wallet_id, user_id)
		if bErr != nil {
			t.Errorf("Test Failed with error %q, caused by getting balance", bErr)
			return
		}
		balance := wallet.Balance
		decimalValue := decimal.NewFromFloat(float64(amount))
		t.Log("DZOVI=====", balance)
		err := MakeDeposit(ctx, db, wallet_id, decimalValue)
		if err != nil {
			if decimalValue.IsNegative() {
				assert.Error(t, err)
				assert.EqualError(t, err, "Amount you provided is negative")
			}
			return
		}

		newWallet, nbErr := GetBalance(db, wallet_id, user_id)
		if nbErr != nil {
			t.Errorf("Test Failed with error %q, caused by getting the 2nd balance", err)
			return
		}
		t.Log("DZOVI=====", newWallet.Balance)
		// Checking if newBalance = balance + amount

		var want, got decimal.Decimal
		if amount <= 0 {
			want = decimal.NewFromFloat(float64(0))
			got = decimal.NewFromFloat(float64(0))
		} else {
			want = balance.Add(decimalValue)
			got = newWallet.Balance
		}
		t.Log("DZOVI ===", got, want)
		if !got.Equal(want) {
			t.Errorf("got %q, wanted %q", got, want)
		} else {
			t.Log("TEST PASSED :", got, want)
		}
	}
}

func TestMakeWithdraw(t *testing.T) {
	inputs := []float64{78.091, -67}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	db := db.BuildDBConfig("test")
	defer db.Close()
	user_id := "41582010-aefd-4a2b-a452-141f5688ff36"
	wallet_id := "4a40cb9b-fe20-470c-96b5-ec57f12970e2"

	for _, amount := range inputs {
		wallet, bErr := GetBalance(db, wallet_id, user_id)
		if bErr != nil {
			t.Errorf("Test Failed with error %q, caused by getting balance", bErr)
			return
		}
		balance := wallet.Balance
		decimalValue := decimal.NewFromFloat(float64(amount))
		err := MakeWithdraw(ctx, db, wallet_id, decimalValue)
		if err != nil {
			if decimalValue.IsNegative() {
				assert.Error(t, err)
				assert.EqualError(t, err, "Amount you provided is negative")
			}
			return
		}
		newWallet, nbErr := GetBalance(db, wallet_id, user_id)
		if nbErr != nil {
			t.Errorf("Test Failed with error %q, caused by getting the 2nd balance", err)
			return
		}
		// Checking if newBalance = balance - amount
		var want, got decimal.Decimal
		if amount <= 0 {
			want = decimal.NewFromFloat(float64(0))
			got = decimal.NewFromFloat(float64(0))
		} else {
			want = balance.Sub(decimalValue)
			got = newWallet.Balance
		}

		if !got.Equal(want) {
			t.Errorf("got %q, wanted %q", got, want)
		} else {
			t.Log("TEST PASSED", got, want)
		}
	}

}

func TestGetTransactions(t *testing.T) {
	db := db.BuildDBConfig("test")
	defer db.Close()
	transactions, err := GetAllTransactions(context.Background(), db)
	assert.NoError(t, err)
	assert.NotEmpty(t, transactions)
}
