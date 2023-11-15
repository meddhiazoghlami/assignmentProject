package services

import (
	"assignmentProject/db"
	"context"
	"testing"
	"time"

	"github.com/shopspring/decimal"
)

func TestMakeDeposit(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	db := db.BuildDBConfig()
	defer db.Close()
	user_id := "41582010-aefd-4a2b-a452-141f5688ff36"
	wallet_id := "4a40cb9b-fe20-470c-96b5-ec57f12970e2"
	wallet, bErr := GetBalance(db, wallet_id, user_id)
	if bErr != nil {
		t.Errorf("Test Failed with error %q, caused by getting balance", bErr)
		return
	}
	balance := wallet.Balance
	amount := 4200
	decimalValue := decimal.NewFromInt(int64(amount))
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
	// Checking if newBalance = balance + amount
	want := balance.Add(decimalValue)
	got := newWallet.Balance
	if !got.Equal(want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestMakeWithdraw(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	db := db.BuildDBConfig()
	defer db.Close()
	user_id := "41582010-aefd-4a2b-a452-141f5688ff36"
	wallet_id := "4a40cb9b-fe20-470c-96b5-ec57f12970e2"
	wallet, bErr := GetBalance(db, wallet_id, user_id)
	if bErr != nil {
		t.Errorf("Test Failed with error %q, caused by getting balance", bErr)
		return
	}
	balance := wallet.Balance
	amount := 4200
	decimalValue := decimal.NewFromInt(int64(amount))
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
	}
}
