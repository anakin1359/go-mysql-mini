package main

import (
	"developer/database"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
}

// func main() {
// 	fmt.Println("test")
// }

func main() {
	db, err := database.DbConnector()
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	query := "SELECT * FROM user"
	stmt, err := db.Prepare(query)
	if err != nil {
		fmt.Println("Failure of Query issuing process.\n", err)
		return
	}
	var (
		user_id       = 1001
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
