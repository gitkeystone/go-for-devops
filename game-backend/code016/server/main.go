package main

import (
	"fmt"
	"log"
	"net/http"
	"percxh/code016/model"
	"percxh/code016/utils"
	"time"
)

func main() {
	//创建新房间
	grm := model.CreateRoom(1)
	fmt.Printf("%s\n", "欢迎大家来玩！")

	//接待新玩家
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		utils.ServeWs(grm[1].Hub, w, r)
	})

	go func() {
		for {
			battleInfo := model.BattleInfo{
				PlayerID: "玩家1",
				Health:   80,
				Mana:     60,
				Action:   "攻击",
				Skills: []model.Skill{
					{Name: "魔法水晶箭", Damage: 200, ManaCost: 80, CoolDown: 60, IsOnCooldown: false},
				},
				Position: model.Position{
					X: 10,
					Y: 20,
				},
				StatusEffects:  []string{"眩晕"},
				TargetPlayerID: "玩家2",
			}
			grm[1].Hub.Broadcast <- battleInfo
			time.Sleep(5e9)
		}
	}()

	log.Fatal(http.ListenAndServe(":8081", nil))

}
