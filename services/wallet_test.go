package services

import (
	"assignmentProject/db"
	"testing"

	"github.com/shopspring/decimal"
)

func TestGetBalance(t *testing.T) {
	db := db.BuildDBConfig()
	defer db.Close()

	user_id := "41582010-aefd-4a2b-a452-141f5688ff36"
	wallet_id := "4a40cb9b-fe20-470c-96b5-ec57f12970e2"

	got, err := GetBalance(db, wallet_id, user_id)
	if err != nil {
		t.Errorf("Test Failed with error %q", err)
	}
	want := 400
	decimalValue := decimal.NewFromInt(int64(want))
	if !got.Balance.Equal(decimalValue) {
		t.Errorf("wanted: %q, got: %q", got.Balance, decimalValue)
	}
}
