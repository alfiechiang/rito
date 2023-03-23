package room

import (
	"fmt"
	"math/rand"
	"rito/internal/common"
)

func (r *Room) EnterRoom(req common.WReq) common.WResp {

	roomID := rand.Intn(1000000)
	player := &common.Player{
		WsConn:   req.Conn,
		UserName: "",
		Seq:      req.Seq,
	}
	players := make([]*common.Player, 0)
	players = append(players, player)

	room := &Room{
		RoomID:  uint(roomID),
		Players: players,
	}

	fmt.Println(room)

	return common.WResp{
		Code: 200,
		Msg:  "遊戲成功",
	}

}
