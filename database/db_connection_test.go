package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func TestDbConnector(t *testing.T) {
	var (
		user     = os.Getenv("MYSQL_USER")
		pass     = os.Getenv("MYSQL_PASSWORD")
		host     = os.Getenv("MYSQL_HOST")
		port     = os.Getenv("MYSQL_HOST_PORT")
		database = os.Getenv("MYSQL_DATABASE")
	)

	conf := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8mb4"

	db, err := sql.Open("mysql", conf)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Database connection failed.")
		t.Logf("%+v", err)
		return
	} else {
		fmt.Println("Database connection succeeded.")
		return
	}
}
