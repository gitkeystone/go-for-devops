package model

import "fmt"

type GameRoom struct {
	RoomID   int      `json:"RoomID"`
	Players  []string `json:"Players"`
	IsActive bool     `json:"IsActive "`
}

var defaultGameRoomManager = make(map[int]GameRoom)

func CreateRoom(roomID int, player string) {
	newRoom := GameRoom{
		RoomID:   roomID,
		Players:  []string{player},
		IsActive: true,
	}

	defaultGameRoomManager[roomID] = newRoom

	fmt.Printf("创建房间 %d，玩家 %s 已加入\n", roomID, player)
}

func LeaveRoom(roomID int, player string) {
	if room, ok := defaultGameRoomManager[roomID]; ok {
		lastIndex := len(room.Players) - 1
		for i, v := range room.Players {
			if v == player {
				room.Players[i] = room.Players[lastIndex]
				room.Players = room.Players[:lastIndex]
				break
			}
		}
		if len(room.Players) != 0 {
			defaultGameRoomManager[roomID] = room
		} else {
			delete(defaultGameRoomManager, roomID)
		}
		fmt.Printf("玩家 %s 已离开房间 %d\n", player, roomID)
	} else {
		fmt.Printf("房间 %d 不存在\\n", roomID)
	}
}

func JoinRoom(roomID int, player string) {
	if room, ok := defaultGameRoomManager[roomID]; ok {
		room.Players = append(room.Players, player)
		defaultGameRoomManager[roomID] = room
		fmt.Printf("玩家 %s 已加入房间 %d\n", player, roomID)
	} else {
		fmt.Printf("房间 %d 不存在\\n", roomID)
	}
}

func ShowRoomInfo(roomID int) {
	if room, ok := defaultGameRoomManager[roomID]; ok {
		fmt.Printf("房间 %d 的信息：\n", roomID)
		fmt.Printf("玩家: %v\n", room.Players)
		fmt.Printf("房间状态: %v\n", room.IsActive)
	} else {
		fmt.Printf("房间 %d 不存在\n", roomID)
	}
}
