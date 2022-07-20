package main

import (
	"fmt"
	"log"

	"developer/database"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

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
