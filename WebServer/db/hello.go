package db

import (
	"github.com/afanke/OJO/WebServer/config"
	"github.com/afanke/OJO/WebServer/dto"
	"github.com/afanke/OJO/utils/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
)

var db *sqlx.DB

func init() {
	var err error
	cfg := config.Config.DataBase
	db, err = sqlx.Open(cfg.DriverName, cfg.DataSourceName)
	if err != nil {
		log.Warn("error:%v", err)
		os.Exit(-1)
	}
	err = db.Ping()
	if err != nil {
		log.Warn("error:%v", err)
		os.Exit(-1)
	}
}

func GetProgress() []dto.Progress {
	var res []dto.Progress
	err := db.Select(&res, "select * from ojo.progress")
	if err != nil {
		log.Warn("error:%v", err)
		return nil
	}
	return res
}
