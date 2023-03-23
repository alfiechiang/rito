package server

import (
	"fmt"
	"net/http"
	"rito/internal/app/room"
	"rito/internal/common"

	"rito/internal/router"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Server struct {
	Router router.Router
	Addr   string
	Users  []*User
}

type User struct {
	WsConn *websocket.Conn
	Seq    int64
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func NewUser(seq int64, conn *websocket.Conn) *User {
	return &User{Seq: seq, WsConn: conn}
}

func NewServer(addr string) *Server {
	return &Server{Addr: addr, Users: make([]*User, 0)}
}

func (s *Server) Run() {

	r := gin.Default()
	s.Router.AddHandler()
	r.GET("/ws", func(c *gin.Context) {
		//upgrade get request to websocket protocol
		ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)

		seq := time.Now().UnixNano()
		user := NewUser(seq, ws)
		s.Users = append(s.Users, user)

		if err != nil {
			fmt.Println(err)
			return
		}

		defer ws.Close()
		for {
			//Read Message from client
			mt, message, err := ws.ReadMessage()
			if err != nil {
				fmt.Println(err)
				break
			}
			//If client message is ping will return pong
			if string(message) == "ping" {
				message = []byte("pong")
			}

			req := common.WReq{Conn: ws, Seq: seq}
			resp := s.Router.Exec("room.create", req, seq)
			fmt.Println(resp)
			rooms := room.Rooms
			fmt.Println(rooms)
			//Response message to client
			err = ws.WriteMessage(mt, message)
			if err != nil {
				fmt.Println(err)
				break
			}
		}
	})
	r.Run(s.Addr) // listen and serve on 0.0.0.0:8080
}
