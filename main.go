package main

import (
	"rito/internal/server"
)



func main(){
	server := server.NewServer(":3000")
	server.Run()
}