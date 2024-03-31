package middlewares

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func EnsureLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		log.Printf("Authorization : %s", authHeader)
		if authHeader != "Basic emVkOjU0NTY0Zw==" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Wrong Token"})
			return
		}
		c.Next()
	}
}

func EnsureNotLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.GetHeader("Authorization")) == 0 {
			c.Next()
			return
		}
	}
}
