package router

import (
	"rito/internal/app/player"

	"rito/internal/app/room"
)

func (r *Router) basicRouter() {
	r.Handler["player.game"] = player.PlayerInfo
	r.Handler["room.create"] = room.CreateRoom
	r.Handler["room.list"] = room.RoomList

}
