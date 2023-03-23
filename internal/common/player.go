package common

import (
	"github.com/gorilla/websocket"
)

type Player struct {
	WsConn   *websocket.Conn
	UserName string
	Seq int64
}
