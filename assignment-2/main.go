package main

import (
	"assignment-2/controllers"
	"assignment-2/database"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	db := database.ConnectDB()

	orderController := controllers.NewOrderController(db)

	router.GET("/order", orderController.GetAllOrder)
	router.POST("/order", orderController.CreateOrder)
	router.GET("/order/:orderID", orderController.GetOneOrder)
	router.PUT("/order/:orderID", orderController.UpdateOrder)
	router.DELETE("/order/:orderID", orderController.DeleteOrder)

	PORT := ":4000"
	router.Run(PORT)
}
