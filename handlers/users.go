package handlers

import (
	"encoding/json"
	"go-template-api/common"
	"go-template-api/controller"
	"go-template-api/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func accessTo(service string, c *gin.Context) (int, model.User_Table) {
	var userReq model.User_Table
	json.Unmarshal([]byte(c.PostForm("user")), &userReq)
	var result common.SqlQueryStatus
	switch service {
	case "login":
		result = controller.CheckCredentials(userReq)
	case "register":
		result = controller.CreateNewUser(userReq)
	default:
	}
	if code := result.Code; code != 1 {
		c.JSON(http.StatusOK, gin.H{
			"message": result.Message,
		})
	}
	return result.Code, userReq
}
func Loging(c *gin.Context) {
	log.Println("LOGIN USER")
	code, userReq := accessTo("login", c)
	if code == 1 {
		token, err := common.GenetatJwt(userReq)
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, common.LoginResponse{Message: "Jwt Generated", Token: token})
	}
}

func Register(c *gin.Context) {
	log.Println("REGISTER  USER")
	code, userReq := accessTo("register", c)
	if code == 1 {
		c.JSON(http.StatusOK, common.ReisterResponse{Message: "Register successfully", User: userReq})
	}
}

func LogOut(c *gin.Context) {
	log.Println("LOGOUT USER")
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, this is Setting page",
	})
}
