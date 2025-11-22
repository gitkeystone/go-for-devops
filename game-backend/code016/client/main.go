package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"percxh/code016/model"
)

func main() {
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8081/ws", nil)
	if err != nil {
		log.Fatal(err)
	}

	defer func(conn *websocket.Conn) {
		_ = conn.Close()
	}(conn)

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}

		var battleInfo model.BattleInfo
		err = json.Unmarshal(message, &battleInfo)
		if err != nil {
			log.Println("Unmarshal error:", err)
			continue
		}
		log.Printf(`收到战斗数据: 
            玩家ID: %s
            生命值: %d
            法力值: %d
            当前行为: %s
            技能: %+v
            位置: (%d, %d)
            附加效果: %s
            被攻击玩家ID: %v`,
			battleInfo.PlayerID, battleInfo.Health,
			battleInfo.Mana, battleInfo.Action,
			battleInfo.Skills,
			battleInfo.Position.X, battleInfo.Position.Y,
			battleInfo.StatusEffects, battleInfo.TargetPlayerID)
	}
}
