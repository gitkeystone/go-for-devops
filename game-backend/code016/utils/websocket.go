package utils

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"percxh/code016/model"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// ServeWs ws连接处理函数
func ServeWs(hub *model.Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil) // 将http协议转化成websocket协议
	if err != nil {
		log.Println(err)
		return
	}
	client := &model.Client{
		Conn: conn,
		Send: make(chan []byte, 256),
	}
	hub.Register <- client

	go client.Write()
	go client.Read(hub)

}
