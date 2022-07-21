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

func EnvSetting() string {
	var (
		user     = os.Getenv("MYSQL_USER")
		pass     = os.Getenv("MYSQL_PASSWORD")
		host     = os.Getenv("MYSQL_HOST")
		port     = os.Getenv("MYSQL_HOST_PORT")
		database = os.Getenv("MYSQL_DATABASE")
	)
	conf := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8mb4"
	return conf
}

func TestDbConnector(t *testing.T) {
	conf := EnvSetting()

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

func TestDbInsert(t *testing.T) {
	conf := EnvSetting()
	db, err := sql.Open("mysql", conf)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer db.Close()

	query := "INSERT INTO user(user_id, user_name, email_address, tel_number) VALUES(?, ?, ?, ?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		fmt.Println("Failure of Query issuing process.\n", err)
		return
	}

	var (
		user_id       = 10001
		user_name     = "sample_user"
		email_address = "sample-mail@example.co.jp"
		tel_number    = "050-1234-5678"
	)

	result, err := stmt.Exec(user_id, user_name, email_address, tel_number)
	if err != nil {
		fmt.Println("Failure of query execution process.\n", err)
		return
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		fmt.Println("Insert ID: ", -1)
		return
	}
	fmt.Println("Insert ID: ", insertId)
}
