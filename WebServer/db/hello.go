package db

import (
	"fmt"
	"github.com/afanke/OJO/WebServer/config"
	"github.com/afanke/OJO/WebServer/dto"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

var db *sqlx.DB

func init() {
	var err error
	cfg := config.Config.DataBase
	db, err = sqlx.Open(cfg.DriverName, cfg.DataSourceName)
	if err != nil {
		fmt.Printf("error:%v", err)
		os.Exit(-1)
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("error:%v", err)
		os.Exit(-1)
	}
}

func GetProgress() []dto.Progress {
	var res []dto.Progress
	err := db.Select(&res, "select * from ojo.progress")
	if err != nil {
		log.Printf("error:%v", err)
		return nil
	}
	return res
}
