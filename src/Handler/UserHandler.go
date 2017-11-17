package Handler

import (
	"GolangWantedlyHomeWork/src/Dao"

	"net/http"

	"GolangWantedlyHomeWork/src/model/entity"

	"strconv"

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
	var wantedlyUsers entity.WantedlyUsers
	wantedlyUsers = make(entity.WantedlyUsers, 0, 1)
	err := session.Asc("id").Find(&wantedlyUsers)
	if err != nil {
		log.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, wantedlyUsers)
}

func GetUser(c *gin.Context) {
	n := c.Param("id")
	session := Dao.NewXormHandler()
	wantedlyUser := new(entity.WantedlyUser)
	id, err := strconv.Atoi(n)
	if err != nil {
		log.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	wantedlyUser.Id = id
	has, err := session.ID(wantedlyUser.Id).Get(wantedlyUser)
	if !has {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}
	if err != nil {
		log.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, wantedlyUser)
}

type registerUserParam struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func PostUser(c *gin.Context) {
	param := new(registerUserParam)
	if err := c.Bind(param); err != nil {
		log.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	session := Dao.NewXormHandler()
	wantedlyUser := new(entity.WantedlyUser)
	wantedlyUser.Name = param.Name
	wantedlyUser.Email = param.Email

	insert, err := session.Insert(wantedlyUser)
	if insert == int64(0) {
		log.Error(err)
		c.AbortWithStatus(http.StatusNoContent)
		return
	}
	if err != nil {
		log.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	getByEmailWantedlyUser := new(entity.WantedlyUser)
	has, err := session.Where("email=?", wantedlyUser.Email).Get(getByEmailWantedlyUser)
	if !has {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}
	if err != nil {
		log.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusCreated, getByEmailWantedlyUser)
}

type updateUserParam struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func PutUser(c *gin.Context) {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		log.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	param := new(updateUserParam)
	if err := c.Bind(param); err != nil {
		log.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	session := Dao.NewXormHandler()
	wantedlyUser := new(entity.WantedlyUser)
	wantedlyUser.Id = id
	has, err := session.Id(wantedlyUser.Id).Get(wantedlyUser)
	if !has {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if err != nil {
		log.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if param.Name != "" && wantedlyUser.Name != param.Name {
		wantedlyUser.Name = param.Name
	}
	if param.Email != "" && wantedlyUser.Email != param.Email {
		getByEmailParamWantedlyUser := new(entity.WantedlyUser)
		has, err = session.Where("email=?", param.Email).Get(getByEmailParamWantedlyUser)
		if err != nil {
			log.Error(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if !has {
			wantedlyUser.Email = param.Email
		}
	}
	_, err = session.ID(wantedlyUser.Id).Update(wantedlyUser)
	if err != nil {
		log.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	getByEmailWantedlyUser := new(entity.WantedlyUser)
	has, err = session.Where("email=?", wantedlyUser.Email).Get(getByEmailWantedlyUser)
	if !has {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}
	if err != nil {
		log.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusCreated, getByEmailWantedlyUser)
}

func DeleteUser(c *gin.Context) {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		log.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	session := Dao.NewXormHandler()
	wantedlyUser := new(entity.WantedlyUser)
	wantedlyUser.Id = id
	_, err = session.ID(wantedlyUser.Id).Delete(wantedlyUser)
	if err != nil {
		log.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
