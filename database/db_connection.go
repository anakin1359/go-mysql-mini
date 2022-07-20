package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	user     = "developer"
	pass     = "password"
	host     = "127.0.0.1"
	port     = "3333"
	database = "proto"
)

func DbConnector() *sql.DB {
	conf := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8mb4"

	db, err := sql.Open("mysql", conf)
	if err != nil {
		fmt.Println(err.Error())
	}
	return db
}
