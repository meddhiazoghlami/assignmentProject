package main

import (
	"assignmentProject/db"
	"assignmentProject/server"
)

func main() {
	db := db.BuildDBConfig()
	defer db.Close()
	server := &server.Server{
		Db: db,
	}
	r := server.SetupRouter()

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
