package main

import (
	"GolangWantedlyHomeWork/src/Handler"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	route.GET("/", Handler.HalloHandler)
	route.Run("localhost:8080")
}
