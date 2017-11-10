package Handler

import (
	"GolangWantedlyHomeWork/src/Dao"

	"GolangWantedlyHomeWork/src/model/row"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

func InitUserHandler(e *gin.Engine) {
	e.GET("/users", FindUsers)
	e.GET("/users/:id", GetUser)
	e.POST("/users", PostUser)
	e.PUT("/users/:id", PutUser)
	e.DELETE("/users/:id", DeleteUser)
}

func FindUsers(c *gin.Context) {
	session := Dao.NewXormHandler()
	wantedlyUsers := new(row.WantedlyUsers)
	_, err := session.Get(wantedlyUsers)
	if err != nil {
		log.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return

	}
	c.JSON(http.StatusOK, wantedlyUsers)
}

func GetUser(c *gin.Context) {

}

func PostUser(c *gin.Context) {

}

func PutUser(c *gin.Context) {

}

func DeleteUser(c *gin.Context) {

}
