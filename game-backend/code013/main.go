package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"percxh/code013/utils"
)

func main() {
	createTable()
	insertTestData()

	username := "zhangsan"
	password := "zs123456"
	if checkLogin(username, password) {
		fmt.Println("登录成功")
	} else {
		fmt.Println("登录失败")
	}
}

func createTable() {
	db, err := sql.Open("sqlite3", "./game.db")
	if err != nil {
		log.Fatalf("无法打开数据库：%v", err)
	}
	defer db.Close()

	// create `user` table
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
        username VARCHAR(20) NOT NULL PRIMARY KEY,
        password VARCHAR(20) NOT NULL,
        nickname VARCHAR(20) NOT NULL
    );`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("创建表失败: %v", err)
	}
	fmt.Println("users表创建成功（或已存在）")
}

func insertTestData() {
	db, err := sql.Open("sqlite3", "./game.db")
	if err != nil {
		log.Fatalf("无法打开数据库: %v", err)
	}
	defer db.Close()

	// insert entries
	insertSQL := `
 	INSERT INTO users (username, password, nickname)
    VALUES (?, ?, ?);`

	users := []struct {
		username string
		password string
		nickname string
	}{
		{"zhangsan", utils.HashPassword("zs123456"), "张三"},
		{"lisi", utils.HashPassword("ls123456"), "李四"},
		{"wangwu", utils.HashPassword("ww123456"), "王五"},
	}

	for _, user := range users {
		db.Exec(insertSQL, user.username, user.password, user.nickname)
	}
}

func checkLogin(username, password string) bool {
	db, err := sql.Open("sqlite3", "./game.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var hashedPassword string
	querySQL := `SELECT password FROM users WHERE username =?`
	err = db.QueryRow(querySQL, username).Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		panic(err)
	}
	return utils.CheckPasswordEquivalent(hashedPassword, password)
}
