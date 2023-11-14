package services

import (
	"assignmentProject/db"
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/shopspring/decimal"
)

func TestMakeDeposit(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	db := db.BuildDBConfig()
	defer db.Close()
	wallet_id := "4a40cb9b-fe20-470c-96b5-ec57f12970e2"
	amount := 4200
	decimalValue := decimal.NewFromInt(int64(amount))
	err := MakeDeposit(ctx, db, wallet_id, decimalValue)
	if err != nil {
		t.Errorf("Test Failed with error %q", err)
		return
	} else {
		//in case of success - withdraw the amount
		errorW := MakeWithdraw(ctx, db, wallet_id, decimalValue)
		if errorW != nil {
			fmt.Println("Something went wrong in withdrawing")
			return
		}
	}
}

func TestMakeWithdraw(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	db := db.BuildDBConfig()
	defer db.Close()
	wallet_id := "4a40cb9b-fe20-470c-96b5-ec57f12970e2"
	amount := 4200
	decimalValue := decimal.NewFromInt(int64(amount))
	err := MakeWithdraw(ctx, db, wallet_id, decimalValue)
	if err != nil {
		t.Errorf("Test Failed with error %q", err)
		return
	} else {
		//in case of success - withdraw the amount
		errorW := MakeDeposit(ctx, db, wallet_id, decimalValue)
		if errorW != nil {
			fmt.Println("Something went wrong in depositing")
			return
		}
	}
}
