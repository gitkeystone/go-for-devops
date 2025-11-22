package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

// Player 定义玩家数据结构
type Player struct {
	ID    int
	Name  string
	Level int
	Coins int
}

// 插入玩家数据到数据库
func insertPlayer(db *sql.DB, player Player) (int, error) {
	stmt, err := db.Prepare("INSERT INTO players (name, level, coins) VALUES (?,?,?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(player.Name, player.Level, player.Coins)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(lastInsertID), nil
}

// 从数据库查询玩家数据
func getPlayer(db *sql.DB, playerID int) (Player, error) {
	var player Player
	err := db.QueryRow("SELECT id, name, level, coins FROM players WHERE id =?", playerID).
		Scan(&player.ID, &player.Name, &player.Level, &player.Coins)
	if err != nil {
		return player, err
	}
	return player, nil
}

func main() {
	// 连接到SQLite数据库
	db, err := sql.Open("sqlite3", "game.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 创建表（如果不存在）
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS players (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        level INTEGER,
        coins INTEGER
    )`)
	if err != nil {
		panic(err)
	}

	// 插入新玩家数据
	newPlayer := Player{
		Name:  "小明",
		Level: 1,
		Coins: 100,
	}
	playerID, err := insertPlayer(db, newPlayer)
	if err != nil {
		fmt.Println("插入数据失败:", err)
		return
	}
	fmt.Printf("插入玩家成功，ID: %d\n", playerID)

	// 查询玩家数据
	player, err := getPlayer(db, playerID)
	if err != nil {
		fmt.Println("查询数据失败:", err)
		return
	}
	fmt.Printf("查询到玩家信息: ID: %d, 名字: %s, 等级: %d, 金币: %d\n", player.ID, player.Name, player.Level, player.Coins)
}
