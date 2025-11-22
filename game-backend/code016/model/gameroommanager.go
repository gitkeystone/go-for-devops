package model

import "fmt"

type GameRoom struct {
	RoomID   int
	IsActive bool
	*Hub
}

func CreateRoom(roomID int) map[int]GameRoom {
	var gameRoomManager = make(map[int]GameRoom)
	hub := NewHub()
	go hub.Run()
	newRoom := GameRoom{roomID, true, hub}
	gameRoomManager[roomID] = newRoom

	fmt.Printf("创建房间 %d\n", roomID)
	return gameRoomManager
}
