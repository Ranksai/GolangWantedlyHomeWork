package main

import (
	"GolangWantedlyHomeWork/src/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	route.GET("/", handler.HalloHandler)
	route.Run("localhost:8080")
}
