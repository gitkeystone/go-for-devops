package model

import (
	"github.com/gorilla/websocket"
	"log"
)

type Client struct {
	Conn *websocket.Conn
	Send chan []byte
}

// 服务端读取客户端消息的方法
func (c *Client) Read(hub *Hub) {
	defer func() {
		hub.Unregister <- c
		_ = c.Conn.Close()
	}()
	for {
		_, _, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
	}
}

// 服务端写入客户端消息的方法
func (c *Client) Write() {
	defer func() {
		_ = c.Conn.Close()
	}()

	for message := range c.Send {
		err := c.Conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Printf("error: %v", err)
			break
		}
	}
}
