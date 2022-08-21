package database

import (
	"fmt"
	"log"
	"testing"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

// DB接続
func TestDbConnector(t *testing.T) {
	db, err := DbConnector()
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		t.Logf("[FAILED] Database connection. %+v", err)
	} else {
		fmt.Println("[SUCCESS] Database connection succeeded.")
	}
}

// User情報登録
func TestInsertUser(t *testing.T) {
	db, err := DbConnector()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var uid uint32 = 10012
	var uName, addr, tNum string = "proto_user", "proto@example.co.jp", "050-1234-5678"
	lastUid, err := InsertUser(db, uid, uName, addr, tNum)
	if err != nil {
		t.Error(err)
	}
	t.Logf("[SUCCESS] user_id: %+v", lastUid)
}

// User検索
func TestGetUser(t *testing.T) {
	db, err := DbConnector()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var uid uint32 = 10012
	u, err := GetUser(db, uid)
	if err != nil {
		t.Error(err)
	}
	t.Logf("[SUCCESS] User Info: %+v", u)
}

// 全User検索
func TestGetAllUsers(t *testing.T) {
	db, err := DbConnector()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	userList, err := GetAllUsers(db)
	if err != nil {
		t.Error(err)
	}
	t.Logf("[SUCCESS] All User Info %+v", userList)
}

// User検索 -> UserName変更
func TestUpdateUserName(t *testing.T) {
	db, err := DbConnector()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var uid uint32 = 10012
	var userName string = "proto_user_aaa"
	_, err = UpdateUserName(db, uid, userName)
	if err != nil {
		t.Error(err)
	}
	t.Log("[SUCCESS]")
}

// User検索 -> 消去
