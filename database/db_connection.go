package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// User構造体
type User struct {
	UserId       uint32
	UserName     string
	EmailAddress string
	TelNumber    string
}

// 構造体User (配列版)
type UserList []User

type Notification struct {
	At   int64  `json:"at"`
	Item string `json:"item"`
}

// DB接続
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

// User情報登録
func InsertUser(db *sql.DB, userId uint32, userName string, emailAddress string, telNumber string) (uint32, error) {
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

// User検索
func GetUser(db *sql.DB, userId uint32) (*User, error) {
	var u User

	query := "SELECT user_id, user_name, email_address, tel_number FROM user WHERE user_id = ? LIMIT 1"
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("Prepare Error. %v\n", err)
	}

	row := stmt.QueryRow(userId)

	err = row.Scan(&u.UserId, &u.UserName, &u.EmailAddress, &u.TelNumber)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("指定したユーザは存在しません。 %v\n", err)
	}
	if err != nil {
		return nil, err
	}

	return &u, nil
}

// UserId検索 -> UserName変更
func UpdateUserName(db *sql.DB, userId uint32, userName string) (bool, error) {
	u, err := GetUser(db, userId)
	if err != nil {
		return false, err
	}
	if u.UserName == userName {
		return false, fmt.Errorf("前回と同じユーザ名です。\n")
	}

	query := "UPDATE user SET user_name = ? WHERE user_id = ?;"

	stmt, err := db.Prepare(query)
	if err != nil {
		return false, fmt.Errorf("Prepare Error. %v\n", err)
	}

	u.UserName = userName
	_, err = stmt.Exec(u.UserName, u.UserId)
	if err != nil {
		return false, fmt.Errorf("Failure of query execution process. %v\n", err)
	}
	return true, nil
}

// UserId検索 -> 消去
// func DeleteUserInfo() bool {}

// 全User検索
func GetAllUsers(db *sql.DB) (*UserList, error) {
	// 構造体User(配列版)の変数宣言
	var ul UserList

	query := "SELECT user_id, user_name, email_address, tel_number FROM user"
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("Prepare Error. %v\n", err)
	}

	rows, err := stmt.Query() // 複数レコード取得
	if err != nil {
		return nil, fmt.Errorf("Query Error. %v\n", err)
	}

	for rows.Next() {
		var user User
		err = rows.Scan(&user.UserId, &user.UserName, &user.EmailAddress, &user.TelNumber)
		if err != nil {
			return nil, fmt.Errorf("Scan Error. %v\n", err)
		}
		ul = append(ul, user)
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("Rows Error. %v\n", err)
	}

	return &ul, nil
}
