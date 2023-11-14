package server

import (
	"assignmentProject/db"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBalanceRoute(t *testing.T) {
	db := db.BuildDBConfig()
	defer db.Close()
	server := &Server{
		Db: db,
	}
	router := server.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users/872a9d57-4a4c-49be-a447-5f44d3176f72/wallets/4a40cb9b-fe20-470c-96b5-ec57f12970e2/balance", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	// assert.Equal(t, string("{\"balance\":\"1800\"}"), w.Body.String())
}

func TestMakeDepositRoute(t *testing.T) {
	db := db.BuildDBConfig()
	defer db.Close()
	server := &Server{
		Db: db,
	}
	router := server.SetupRouter()

	type MockAmount struct {
		Amount int `json:"amount"`
	}

	reqBody := &MockAmount{
		Amount: 1200,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users/872a9d57-4a4c-49be-a447-5f44d3176f72/wallets/4a40cb9b-fe20-470c-96b5-ec57f12970e2/deposit", bytes.NewBuffer(jsonData))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string("{\"message\":\"Your deposit is done successfully\",\"status\":200}"), w.Body.String())
}

func TestMakeWithdrawRoute(t *testing.T) {
	db := db.BuildDBConfig()
	defer db.Close()
	server := &Server{
		Db: db,
	}
	router := server.SetupRouter()

	type MockAmount struct {
		Amount int `json:"amount"`
	}

	reqBody := &MockAmount{
		Amount: 1200,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users/872a9d57-4a4c-49be-a447-5f44d3176f72/wallets/4a40cb9b-fe20-470c-96b5-ec57f12970e2/withdraw", bytes.NewBuffer(jsonData))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string("{\"message\":\"Your withdraw is done successfully\",\"status\":200}"), w.Body.String())
}
