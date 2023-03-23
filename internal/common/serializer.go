package common

import (
	"github.com/gorilla/websocket"
)

type WReq struct {
	Conn *websocket.Conn
	Seq  int64
	Body interface{}
}

type WResp struct {
	Code int
	Msg  string
	Data interface{}
}
