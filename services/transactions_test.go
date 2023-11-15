package services

import (
	"assignmentProject/db"
	"context"
	"testing"
	"time"

	"github.com/shopspring/decimal"
)

func TestMakeDeposit(t *testing.T) {
	inputs := []float64{100.89, 200, 300, 78.091, -50}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	db := db.BuildDBConfig()
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
		err := MakeDeposit(ctx, db, wallet_id, decimalValue)
		if err != nil {
			t.Errorf("Test Failed with error %q", err)
			return
		}

		newWallet, nbErr := GetBalance(db, wallet_id, user_id)
		if nbErr != nil {
			t.Errorf("Test Failed with error %q, caused by getting the 2nd balance", err)
			return
		}
		// Checking if newBalance = balance - amount
		want := balance.Add(decimalValue)
		got := newWallet.Balance
		if !got.Equal(want) {
			t.Errorf("got %q, wanted %q", got, want)
		} else {
			t.Log("TEST PASSED", got, want)
		}
	}
}

func TestMakeWithdraw(t *testing.T) {
	inputs := []float64{100.89, 200, 300, 78.091, -50}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	db := db.BuildDBConfig()
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
			t.Errorf("Test Failed with error %q", err)
			return
		}

		newWallet, nbErr := GetBalance(db, wallet_id, user_id)
		if nbErr != nil {
			t.Errorf("Test Failed with error %q, caused by getting the 2nd balance", err)
			return
		}
		// Checking if newBalance = balance - amount
		want := balance.Sub(decimalValue)
		got := newWallet.Balance
		if !got.Equal(want) {
			t.Errorf("got %q, wanted %q", got, want)
		} else {
			t.Log("TEST PASSED", got, want)
		}
	}
}
