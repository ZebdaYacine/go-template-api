package router

import (
	"go-template-api/handlers"
	"go-template-api/middlewares"
	"log"

	"github.com/gin-gonic/gin"
)

func GetRouting(server *gin.Engine) {
	log.Println("SETUP ROUTING OF API")
	userRoute := server.Group("/")
	{
		userRoute.
			POST(
				"login",
				middlewares.EnsureNotLoggedIn(),
				handlers.Loging,
			)
		userRoute.
			GET(
				"logout",
				middlewares.EnsureLoggedIn(),
				handlers.LogOut,
			)
		userRoute.
			POST(
				"register",
				middlewares.EnsureNotLoggedIn(),
				middlewares.IsAuthorizedToNewUser(),
				handlers.Register,
			)
	}
	productRouter := server.Group("/product")
	{
		productRouter.
			GET(
				"productList",
				middlewares.EnsureLoggedIn(),
				middlewares.IsAuthorizedToGetProductsList(),
				handlers.ProductList,
			)
		productRouter.
			POST(
				"addProduct",
				middlewares.EnsureLoggedIn(),
				middlewares.IsAuthorizedToNewProduct(),
				handlers.NewProduct,
			)
	}

}
