package Handler

import "github.com/gin-gonic/gin"

func HalloHandler(c *gin.Context)  {
	c.JSON(200, gin.H{"message": "Hello, World!",})
}
