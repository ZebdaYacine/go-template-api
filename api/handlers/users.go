package handlers

import (
	"encoding/json"
	"go-template-api/api/security"
	"go-template-api/common"
	"go-template-api/internal/controller"
	"go-template-api/internal/model"
	"go-template-api/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func accessTo(service string, c *gin.Context) (common.SqlQueryStatus, model.User_Table) {
	var userReq *model.User_Table
	json.Unmarshal([]byte(c.PostForm("user")), &userReq)
	var result common.SqlQueryStatus
	if userReq == nil {
		result.Code = 500
		result.Message = "EMPTY REQUEST"
		return result, model.User_Table{ID: 0}
	}
	switch service {
	case "login":
		result = controller.CheckCredentials(userReq)
	case "register":
		result = controller.CreateNewUser(userReq)
	default:
	}
	if code := result.Code; code != 1 {
		utils.PrintError(result.Message)
		c.JSON(http.StatusOK, gin.H{
			"message": result.Message,
		})
	}
	return result, *userReq
}
func Loging(c *gin.Context) {
	log.Println("LOGIN SERVICE")
	resualt, userReq := accessTo("login", c)
	if resualt.Code == 1 {
		token, err := security.GenetatJwt(userReq)
		if err != nil {
			return
		}
		utils.PrintInfo(resualt.Message)
		c.JSON(http.StatusOK, common.LoginResponse{Message: resualt.Message, Token: token})
	}
	if resualt.Code == 500 {
		utils.PrintError(resualt.Message)
		c.JSON(http.StatusOK, common.LoginResponse{Message: resualt.Message})
	}
}

func Register(c *gin.Context) {
	log.Println("REGISTER SERVICE")
	resualt, userReq := accessTo("register", c)
	if resualt.Code == 1 {
		utils.PrintInfo(resualt.Message)
		c.JSON(http.StatusOK, common.ReisterResponse{Message: resualt.Message, User: userReq})
	}
	if resualt.Code == 500 {
		utils.PrintError(resualt.Message)
		c.JSON(http.StatusOK, common.ReisterResponse{Message: resualt.Message})
	}
}

func LogOut(c *gin.Context) {
	log.Println("LOGOUT SERVICE")
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, this is logout page",
	})
}
