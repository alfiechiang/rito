package player

import (
	"rito/internal/common"
)

func PlayerInfo(req common.WReq) common.WResp {

	return common.WResp{
		Code: 200,
		Msg:  "遊戲成功",
	}

}
