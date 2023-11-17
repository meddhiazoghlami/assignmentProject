package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Dbname   string
	Password string
}

func BuildDBConfig(typeDB string) *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	var port, host, name, user, password string
	if typeDB == "app" {

		port = os.Getenv("DB_PORT")
		host = os.Getenv("DB_HOST")
		name = os.Getenv("DB_NAME")
		user = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
	} else if typeDB == "test" {
		port = os.Getenv("DB_PORT_TEST")
		host = os.Getenv("DB_HOST_TEST")
		name = os.Getenv("DB_NAME_TEST")
		user = os.Getenv("DB_USER_TEST")
		password = os.Getenv("DB_PASSWORD_TEST")
	}

	a, _ := strconv.Atoi(port)

	dbconfig := DBConfig{
		Host:     host,
		Port:     a,
		User:     user,
		Password: password,
		Dbname:   name,
	}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", dbconfig.Host, dbconfig.Port, dbconfig.User, dbconfig.Password, dbconfig.Dbname)
	fmt.Println(psqlInfo)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	return db
}
