package main

import (
	"go-template-api/db"
	server "go-template-api/servers"

	//server "go-template-api/servers"
	"log"
	//server "go-template-api/servers"
)

func luanch() {
	log.Println("Initializing Server")
	db.ConnectDb()
	server.InitGinServer()
}

func main() {
	luanch()
}
