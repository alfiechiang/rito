package room

import (
	"math/rand"
	"rito/internal/common"
)
func CreateRoom(req common.WReq) common.WResp {
	
	roomID := rand.Intn(1000000)
	player := &common.Player{
		WsConn:   req.Conn,
		UserName: "",
		Seq: req.Seq,
	}
	players := make([]*common.Player, 0)
	players = append(players, player)

	room := &Room{
		RoomID:  uint(roomID),
		Players: players,
	}


	Rooms = append(Rooms, room)

	return common.WResp{
		Code: 200,
		Msg:  "創建房間成功",
	}

}