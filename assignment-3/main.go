package main

import (
	"assignment-3/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("template/*")

	router.GET("/", controller.ReadAndUpdateWeather)

	PORT := ":4040"
	router.Run(PORT)
}
