package Dao

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-xorm/xorm"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
)

func NewXormHandler() *xorm.Engine {
	host := os.Getenv("POSTGRES_HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("POSTGRES_PORT")
	if port == "" {
		port = "5432"
	}
	_, err := strconv.Atoi(port)
	if err != nil {
		log.Error(err)
	}
	user := os.Getenv("POSTGRES_USER")
	if user == "" {
		user = "postgres"
	}
	dbname := "wantedlyhomework"

	dsn := fmt.Sprintf(
		"user=%s host=%s port=%s dbname=%s sslmode=disable",
		user,
		host,
		port,
		dbname,
	)
	engine, err := xorm.NewEngine("postgres", dsn)
	if err != nil {
		log.Error(err)
		return nil
	}
	//defer engine.Close()

	return engine
}
