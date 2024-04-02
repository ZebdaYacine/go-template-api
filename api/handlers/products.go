package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProductList(c *gin.Context) {
	// log.Println("ACCESS TO PRODUCTS LIST PAGE")
	// list := make([]model.Goods, 0)
	// for i := 0; i < 20; i++ {
	// 	list = append(
	// 		list,
	// 		model.Goods{Name: "Goods" + strconv.Itoa(i), Price: float64(i * 2500)},
	// 	)
	// }
	// c.JSON(http.StatusOK, gin.H{
	// 	"productsList": list,
	// })
}

func NewProduct(c *gin.Context) {
	log.Println("ACCESS TO NEW PRODUCT PAGE")
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, this is new product page",
	})
}
