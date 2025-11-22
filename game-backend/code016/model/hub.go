package model

import (
	"encoding/json"
	"fmt"
	"log"
)

type Hub struct {
	Clients    map[*Client]bool
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan BattleInfo
}

// NewHub 创建新的Hub实例
func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[*Client]bool),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan BattleInfo),
	}
}

// Run Hub的运行方法，处理客户端的注册、注销和广播消息
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
			fmt.Printf("玩家 %s 已加入房间\n", client.Conn.RemoteAddr())
		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
			}
			fmt.Printf("玩家 %s 已退出房间\n", client.Conn.RemoteAddr())
		case battleInfo := <-h.Broadcast:
			data, err := json.Marshal(battleInfo)
			if err != nil {
				log.Println("Marshal error:", err)
				continue
			}
			for client := range h.Clients {
				select {
				case client.Send <- data:
				default:
					delete(h.Clients, client)
					close(client.Send)
				}
			}
		}
	}
}
