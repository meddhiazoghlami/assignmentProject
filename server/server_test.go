package server

import (
	"assignmentProject/db"
	"assignmentProject/models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

type ResponseUser struct {
	Status int         `json:"status"`
	User   models.User `json:"user"`
}

type ResponseWallet struct {
	Status int           `json:"status"`
	Wallet models.Wallet `json:"wallet"`
}

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
		Amount decimal.Decimal `json:"amount"`
	}
	amount := 123.567

	reqBody := &MockAmount{
		Amount: decimal.NewFromFloat(float64(amount)),
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
		Amount: decimal.NewFromFloat(float64(amount)),
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
		Amount decimal.Decimal `json:"amount"`
	}
	amount := 123.567
	reqBody := &MockAmount{
		Amount: decimal.NewFromFloat(float64(amount)),
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
		Amount: decimal.NewFromFloat(float64(amount)),
	}
	jsonDatas, err := json.Marshal(reqBodys)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	reqs, _ := http.NewRequest("POST", "/users/"+user_id+"/wallets/"+wallet_id+"/deposit", bytes.NewBuffer(jsonDatas))
	router.ServeHTTP(w, reqs)
}

//Scenario

func TestScenario(t *testing.T) {
	db := db.BuildDBConfig()
	defer db.Close()
	server := &Server{
		Db: db,
	}
	router := server.SetupRouter()

	//Step 1: Create a new User then save its user_id
	reqBodys := models.User{
		Username: "Dhia",
	}
	jsonDatas, err := json.Marshal(reqBodys)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonDatas))
	router.ServeHTTP(w, req)
	var responseUser ResponseUser
	errUser := json.Unmarshal([]byte(w.Body.String()), &responseUser)
	if errUser != nil {
		fmt.Println("errUseror: ", errUser)
	}
	assert.Equal(t, 201, w.Code)
	assert.Equal(t, reqBodys.Username, responseUser.User.Username)
	user_id := responseUser.User.User_id
	//Step 2: Create a new wallet and save its wallet_id
	walletBody := models.Wallet{
		Currency: "EUR",
	}
	jsonWallet, err := json.Marshal(walletBody)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	w2 := httptest.NewRecorder()
	reqW, _ := http.NewRequest("POST", "/users/"+user_id+"/wallets", bytes.NewBuffer(jsonWallet))
	router.ServeHTTP(w2, reqW)
	var responseWallet ResponseWallet
	errWallet := json.Unmarshal([]byte(w2.Body.String()), &responseWallet)
	if errWallet != nil {
		fmt.Println("errWallet", errWallet)
	}
	assert.Equal(t, 201, w2.Code)
	assert.Equal(t, walletBody.Currency, responseWallet.Wallet.Currency)
	assert.Equal(t, user_id, responseWallet.Wallet.User_id)
	balanceCompare := responseWallet.Wallet.Balance.Equal(decimal.NewFromInt(0))
	assert.Equal(t, balanceCompare, true)
	wallet_id := responseWallet.Wallet.Wallet_id
	//Step 3: Make a deposit on that wallet then check its balance
	type MockAmount struct {
		Amount decimal.Decimal `json:"amount"`
	}
	amount := 123.567
	reqBody := &MockAmount{
		Amount: decimal.NewFromFloat(float64(amount)),
	}
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	w3 := httptest.NewRecorder()
	req3, _ := http.NewRequest("POST", "/users/"+user_id+"/wallets/"+wallet_id+"/deposit", bytes.NewBuffer(jsonData))
	router.ServeHTTP(w3, req3)
	assert.Equal(t, 200, w3.Code)
	assert.Equal(t, string("{\"message\":\"Your deposit is done successfully\",\"status\":200}"), w3.Body.String())
	//Step 4: Check if the new balance is equal to the amount
	type Balance struct {
		Balance decimal.Decimal
	}
	var balance Balance
	w4 := httptest.NewRecorder()
	req4, _ := http.NewRequest("GET", "/users/"+user_id+"/wallets/"+wallet_id+"/balance", bytes.NewBuffer(jsonData))
	router.ServeHTTP(w4, req4)
	errBalance := json.Unmarshal([]byte(w4.Body.String()), &balance)
	if errBalance != nil {
		fmt.Println("errBalance", errBalance)
	}
	balanceCom := balance.Balance.Equal(reqBody.Amount)
	assert.Equal(t, 200, w4.Code)
	assert.Equal(t, balanceCom, true)
	//Step 5: Make a withdraw with the same amount of money
	w5 := httptest.NewRecorder()
	req5, _ := http.NewRequest("POST", "/users/"+user_id+"/wallets/"+wallet_id+"/withdraw", bytes.NewBuffer(jsonData))
	router.ServeHTTP(w5, req5)
	assert.Equal(t, 200, w5.Code)
	assert.Equal(t, string("{\"message\":\"Your withdraw is done successfully\",\"status\":200}"), w5.Body.String())
	//Step 6: Compare the balance again and its has to be 0
	var balance2 Balance
	w6 := httptest.NewRecorder()
	req6, _ := http.NewRequest("GET", "/users/"+user_id+"/wallets/"+wallet_id+"/balance", bytes.NewBuffer(jsonData))
	router.ServeHTTP(w6, req6)
	errBalance2 := json.Unmarshal([]byte(w6.Body.String()), &balance2)
	if errBalance2 != nil {
		fmt.Println("errBalance2", errBalance2)
	}
	balanceCom2 := balance2.Balance.Equal(decimal.NewFromInt(0))
	assert.Equal(t, 200, w6.Code)
	assert.Equal(t, balanceCom2, true)
}
