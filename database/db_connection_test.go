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
	conf := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + database + "?parseTime=true"
	return conf
}

// DB接続
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

// User情報登録
func TestInsertUser(t *testing.T) {
	conf := EnvSetting()
	db, err := sql.Open("mysql", conf)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	var uid uint32 = 10001
	var uName, addr, tNum string = "proto_user", "proto@example.co.jp", "050-1234-5678"
	lastUid, err := InsertUser(uid, uName, addr, tNum)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("[SUCCESS] uid:", lastUid)
	// t.Logf("%+v", lastUid)
}

// User検索
func TestGetUser(t *testing.T) {
	conf := EnvSetting()
	db, err := sql.Open("mysql", conf)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	var uid uint32 = 100020
	u, err := GetUser(uid)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", u)
}

// 全User検索
func TestGetAllUsers(t *testing.T) {
	conf := EnvSetting()
	db, err := sql.Open("mysql", conf)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	userList, err := GetAllUsers()
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", userList)
}

// User検索 -> UserName変更
// User検索 -> 消去
