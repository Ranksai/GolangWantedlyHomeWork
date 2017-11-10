package Dao

import (
	"github.com/go-xorm/xorm"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
)

func NewXormHandler() *xorm.Engine {
	url := "user=root host=localhost port=5432 dbname=wantedlyHomework sslmode=disable"
	engine, err := xorm.NewEngine("postgres", url)
	if err != nil {
		log.Error(err)
		return nil
	}
	//defer engine.Close()

	return engine
}
