package server

import (
	"go-template-api/config"
	router "go-template-api/routers"
	"log"

	"github.com/gin-gonic/gin"
)

func InitGinServer() {
	//gin.SetMode(gin.DebugMode)

	server := gin.New()
	//server.Use(gin.Logger())
	//server.Use(gin.Recovery())

	router.GetRouting(server)
	err := server.Run(config.GetURL())
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
		return
	}
}
