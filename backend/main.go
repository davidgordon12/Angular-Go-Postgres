package main

import (
	"api/data"
	"api/services/orders"

	"github.com/gin-gonic/gin"
)

func main() {
	ctx := data.CreateConn()
	defer ctx.Close()

	orders := orders.GetAllOrders(ctx)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"orders": orders,
		})
	})
	r.Run()
}
