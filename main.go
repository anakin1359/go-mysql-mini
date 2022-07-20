package main

import (
	"fmt"

	"developer/database"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := database.DbConnector()

	defer db.Close()

	err := db.Ping()
	if err != nil {
		fmt.Println("Database connection failed.")
		return
	} else {
		fmt.Println("Database connection succeeded.")
	}
}
