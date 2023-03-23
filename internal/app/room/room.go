package room

import (
	"rito/internal/common"
)

type HandlerFunc func(req common.WReq) common.WResp

var Rooms = make([]*Room, 0)

type Room struct {
	RoomID  uint
	Players []*common.Player
	Handler map[string]HandlerFunc
}

func (r *Room) Router() {
	r.Handler["room.enter"] = r.EnterRoom

}

func (r *Room) Exec(name string, req common.WReq) common.WResp {

	var resp common.WResp
	for k, fn := range r.Handler {
		if k == name {
			resp = fn(req)
		}
	}

	return resp

}
