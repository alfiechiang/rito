package room

import (
	"rito/internal/common"
)

func RoomList(req common.WReq) common.WResp {

	return common.WResp{
		Code: 200,
		Data: Rooms,
		Msg:  "創建房間成功",
	}

}
