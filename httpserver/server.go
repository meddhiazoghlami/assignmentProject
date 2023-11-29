package server

import (
	"database/sql"

	controllers "github.com/meddhiazoghlami/assignmentProject/controllers"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Db *sql.DB
}

func (s *Server) SetupRouter() *gin.Engine {

	r := gin.Default()
	//middleware to intercept token
	// r.Use(middleware.InterceptBearerToken)
	//add User
	r.POST("/users", controllers.AddUser(s.Db))
	//add wallet to a specific user
	r.POST("/users/:id/wallets", controllers.AddWallet(s.Db))
	//get all wallets for a specific user
	r.GET("/users/:id/wallets", controllers.GetUserWallets(s.Db))
	//middleware to get a user id (sub) from context
	// r.GET("/users/:id/wallets", middleware.GetUserFromContext, controllers.GetUserWallets(s.Db))
	//get a specific wallet
	r.GET("/users/:id/wallets/:wallet_id", controllers.GetWallet(s.Db))
	//make a deposit
	r.POST("/users/:id/wallets/:wallet_id/deposit", controllers.MakeDeposit(s.Db))
	//make a withdraw
	r.POST("/users/:id/wallets/:wallet_id/withdraw", controllers.MakeWithdraw(s.Db))
	//get balance of a specific wallet
	r.GET("/users/:id/wallets/:wallet_id/balance", controllers.GetBalance(s.Db))
	//get all transactions
	r.GET("/transactions", controllers.GetAllTransactions(s.Db))

	return r
}
