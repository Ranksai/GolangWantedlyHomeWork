package main

import (
	"GolangWantedlyHomeWork/src/Handler"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	route.GET("/", Handler.HalloHandler)
	Handler.InitUserHandler(route)
	route.Run(":8080")
}
