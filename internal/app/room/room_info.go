package room

import (
	"rito/internal/common"
)

func (room *Room) RoomInfo(req common.WReq) common.WResp {

	return common.WResp{
		Code: 200,
		Msg:  "遊戲成功",
	}

}
