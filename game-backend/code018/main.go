package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

// PlayerRank 定义玩家数据结构，这里只关注影响排行榜的分数
type PlayerRank struct {
	Name  string
	Score int
}

// 获取排行榜数据
func getRankings() ([]PlayerRank, error) {
	// 连接到SQLite数据库
	db, err := sql.Open("sqlite3", "game.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, score FROM players ORDER BY score DESC LIMIT 10")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rankings []PlayerRank
	for rows.Next() {
		var player PlayerRank
		err := rows.Scan(&player.Name, &player.Score)
		if err != nil {
			return nil, err
		}
		rankings = append(rankings, player)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return rankings, nil
}

func createTestData() {
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
        score INTEGER
    )`)
	if err != nil {
		panic(err)
	}

	// 模拟插入一些玩家数据
	_, err = db.Exec("INSERT INTO players (name, score) VALUES ('玩家A', 500)")
	if err != nil {
		fmt.Println("插入数据失败:", err)
	}
	_, err = db.Exec("INSERT INTO players (name, score) VALUES ('玩家B', 800)")
	if err != nil {
		fmt.Println("插入数据失败:", err)
	}
	_, err = db.Exec("INSERT INTO players (name, score) VALUES ('玩家C', 300)")
	if err != nil {
		fmt.Println("插入数据失败:", err)
	}
}

func init() {
	createTestData()
}

func main() {
	// 获取排行榜数据
	rankings, err := getRankings()
	if err != nil {
		fmt.Println("获取排行榜失败:", err)
		return
	}

	fmt.Println("游戏分数排行榜:")
	for i, rank := range rankings {
		fmt.Printf("%d. %s - %d分\n", i+1, rank.Name, rank.Score)
	}
}
