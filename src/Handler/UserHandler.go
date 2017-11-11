package Handler

import (
	"GolangWantedlyHomeWork/src/Dao"

	"net/http"

	"GolangWantedlyHomeWork/src/model/entity"

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
	wantedlyUsers := new(entity.WantedlyUser)
	_, err := session.Get(wantedlyUsers)
	if err != nil {
		log.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, wantedlyUsers)
}

func GetUser(c *gin.Context) {
	session := Dao.NewXormHandler()
	wantedlyUser := new(entity.WantedlyUser)
	_, err := session.ID(wantedlyUser.Id).Get(wantedlyUser)
	if err != nil {
		log.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, wantedlyUser)
}

func PostUser(c *gin.Context) {

}

func PutUser(c *gin.Context) {

}

func DeleteUser(c *gin.Context) {

}
