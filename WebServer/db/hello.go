package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gogotime/OJO/WebServer/config"
	"github.com/gogotime/OJO/utils/log"
	"github.com/ilibs/gosql/v2"
	"time"
)

func init() {
	cfg := config.Config.DataBase
	configs := make(map[string]*gosql.Config)
	configs["default"] = &gosql.Config{
		Enable:       true,
		Driver:       cfg.DriverName,
		Dsn:          cfg.DataSourceName,
		ShowSql:      true,
		MaxIdleConns: 10,
		MaxLifetime:  int(100 * time.Second),
		MaxOpenConns: 10,
	}
	gosql.SetLogging(true)
	gosql.SetLogger(log.GetLogger())
	err := gosql.Connect(configs)
	if err != nil {
		log.Fatal("error:%v", err)
	}
}
