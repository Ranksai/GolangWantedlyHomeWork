package main

import (
	"./Handler"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	route.GET("/", Handler.HalloHandler)
	route.Run("localhost:3000")
}
