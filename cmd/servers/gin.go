package servers

import (
	router "go-template-api/api/routers"
	"go-template-api/config"
	"log"

	"github.com/gin-gonic/gin"
)

func InitServer() {
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
