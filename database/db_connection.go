package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	UserId       uint32
	UserName     string
	EmailAddress string
	TelNumber    string
}

// type User struct {
// 	UserId       uint32 `json:"user_id"`
// 	UserName     string `json:"user_name"`
// 	EmailAddress string `json:"email_address"`
// 	TelNumber    string `json:"tel_number"`
// }

func DbConnector() (*sql.DB, error) {
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
	return db, nil
}

func InsertUser(userId uint32, userName string, emailAddress string, telNumber string) (uint32, error) {
	db, err := DbConnector()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	query := "INSERT INTO user(user_id, user_name, email_address, tel_number) VALUES(?, ?, ?, ?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return 0, fmt.Errorf("Failure of Query issuing process.%v\n", err)
	}

	result, err := stmt.Exec(userId, userName, emailAddress, telNumber)
	if err != nil {
		return 0, fmt.Errorf("Failure of query execution process. %v\n", err)
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("Insert ID: %d, %v ", -1, err)
	}

	return uint32(insertId), nil
}

func GetAllUsers() {
	db, err := DbConnector()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	query := "SELECT user_id, user_name, email_address, tel_number FROM user"
	stmt, err := db.Prepare(query)
	if err != nil {
		fmt.Println("Prepare Error.\n", err)
		return
	}

	rows, err := stmt.Query() // 複数レコード取得
	if err != nil {
		fmt.Println("Query Error.\n", err)
		return
	}

	for rows.Next() {
		u := &User{}
		err = rows.Scan(&u.UserId, &u.UserName, &u.EmailAddress, &u.TelNumber)
		if err != nil {
			fmt.Println("Scan Error\n", err)
			return
		}
		fmt.Println(u)
	}

	err = rows.Err()
	if err != nil {
		fmt.Println("Rows Error\n", err)
		return
	}
}
