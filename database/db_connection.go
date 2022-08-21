package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// Userの構造体
type User struct {
	UserId       uint32
	UserName     string
	EmailAddress string
	TelNumber    string
}

// 構造体Userの配列版の構造体を新たに定義
type UserList []User

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

func GetAllUsers() (UserList, error) {
	// 構造体User(配列版)の変数宣言
	var ul UserList

	db, err := DbConnector()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	query := "SELECT user_id, user_name, email_address, tel_number FROM user"
	stmt, err := db.Prepare(query)
	if err != nil {
		return ul, fmt.Errorf("Prepare Error. %v\n", err)
	}

	rows, err := stmt.Query() // 複数レコード取得
	if err != nil {
		return ul, fmt.Errorf("Query Error. %v\n", err)
	}

	for rows.Next() {
		var user User
		err = rows.Scan(&user.UserId, &user.UserName, &user.EmailAddress, &user.TelNumber)
		if err != nil {
			return ul, fmt.Errorf("Scan Error. %v\n", err)
		}
		ul = append(ul, user)
	}

	err = rows.Err()
	if err != nil {
		return ul, fmt.Errorf("Rows Error. %v\n", err)
	}

	return ul, nil
}
