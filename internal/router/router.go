package router

import (
	"rito/internal/app/room"
	"rito/internal/common"
)

type HandlerFunc func(req common.WReq) common.WResp

type Router struct {
	Handler     map[string]HandlerFunc
	RoomHandler map[string]HandlerFunc
}

func (r *Router) AddHandler() {
	r.Handler = make(map[string]HandlerFunc, 0)
	r.RoomHandler = make(map[string]HandlerFunc, 0)
	r.basicRouter()
}

func (r *Router) Exec(name string, req common.WReq, userID int64) common.WResp {

	var resp common.WResp
	for k, fn := range r.Handler {
		if k == name {
			resp = fn(req)
		}
	}

	for _, room := range room.Rooms {
		for _, player := range room.Players {
			if player.Seq == userID {
				resp = room.Exec(name, req)
			}

		}
	}

	return resp

}
