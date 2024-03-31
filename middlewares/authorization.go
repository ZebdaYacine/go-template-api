package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getRoleUser(c *gin.Context) string {
	switch c.Request.Method {
	case "POST":
		return c.PostForm("role")
	case "GET":
		return c.Query("role")
	default:
		return c.PostForm("role")
	}
}
func isAdminOrSuperUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := getRoleUser(c)
		if role != "admin" && role != "superUser" {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{"error": "You are not authorized to access this resource"},
			)
			return
		}
		c.Next()
	}
}
func IsAuthorizedToGetProductsList() gin.HandlerFunc {
	return isAdminOrSuperUser()
}

func IsAuthorizedToNewProduct() gin.HandlerFunc {
	return isAdminOrSuperUser()
}

func IsAuthorizedToDeleteAccount() gin.HandlerFunc {
	return isAdminOrSuperUser()
}
