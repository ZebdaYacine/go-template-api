package main

import (
	//server "go-template-api/servers"
	ginServer "go-template-api/cmd/servers"
	"go-template-api/internal/db"
	"log"
)

func luanch() {
	log.Println("Initializing Server")
	go db.ConnectDb()
	ginServer.InitServer()
}

func main() {
	luanch()
}
