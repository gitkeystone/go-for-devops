package main

import "percxh/code015/model"

func main() {
	model.CreateRoom(1, "Player1")
	model.JoinRoom(1, "Player2")
	model.ShowRoomInfo(1)

	model.LeaveRoom(1, "Player1")
	model.LeaveRoom(1, "Player2")
	model.ShowRoomInfo(1)
}
