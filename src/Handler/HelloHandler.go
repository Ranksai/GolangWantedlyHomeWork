package Handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HalloHandler(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
}
