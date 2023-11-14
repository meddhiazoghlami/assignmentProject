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

	users := []string{"41582010-aefd-4a2b-a452-141f5688ff36", "41582010-aefd-4a2b-a452-141f5688ff36"}
	wallets := []string{"4a40cb9b-fe20-470c-96b5-ec57f12970e2", "4a40cb9b-fe20-470c-96b5-ec57f12970e2"}

	for _, uValue := range users {
		for _, wValue := range wallets {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/users/"+uValue+"/wallets/"+wValue+"/balance", nil)
			router.ServeHTTP(w, req)
			assert.Equal(t, 200, w.Code)
			// assert.Equal(t, string("{\"balance\":\"1800\"}"), w.Body.String())
		}
	}
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
		Amount: 1300,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	user_id := "41582010-aefd-4a2b-a452-141f5688ff36"
	wallet_id := "4a40cb9b-fe20-470c-96b5-ec57f12970e2"
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users/"+user_id+"/wallets/"+wallet_id+"/deposit", bytes.NewBuffer(jsonData))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string("{\"message\":\"Your deposit is done successfully\",\"status\":200}"), w.Body.String())

	// If success withdraw the sum amount added (to not affect DB changes)
	reqBodys := &MockAmount{
		Amount: 1300,
	}
	jsonDatas, err := json.Marshal(reqBodys)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	reqs, _ := http.NewRequest("POST", "/users/"+user_id+"/wallets/"+wallet_id+"/withdraw", bytes.NewBuffer(jsonDatas))
	router.ServeHTTP(w, reqs)
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
		Amount: 500,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	user_id := "41582010-aefd-4a2b-a452-141f5688ff36"
	wallet_id := "4a40cb9b-fe20-470c-96b5-ec57f12970e2"
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users/"+user_id+"/wallets/"+wallet_id+"/withdraw", bytes.NewBuffer(jsonData))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string("{\"message\":\"Your withdraw is done successfully\",\"status\":200}"), w.Body.String())

	// If success deposit the same amount removed (to not affect DB changes)
	reqBodys := &MockAmount{
		Amount: 500,
	}
	jsonDatas, err := json.Marshal(reqBodys)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	reqs, _ := http.NewRequest("POST", "/users/"+user_id+"/wallets/"+wallet_id+"/deposit", bytes.NewBuffer(jsonDatas))
	router.ServeHTTP(w, reqs)
}
