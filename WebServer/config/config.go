package config

import (
	"encoding/json"
	"github.com/afanke/OJO/utils/log"
	"io/ioutil"
	"os"
)

var Config ServerConfig

func init() {
	file, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		log.Fatal("error:%v\n", err)
		os.Exit(-1)
	}
	err = json.Unmarshal(file, &Config)
	if err != nil {
		log.Fatal("error:%v\n", err)
		os.Exit(-1)
	}
}

type ServerConfig struct {
	Port        int
	JudgeServer string
	DataBase    DataBaseConfig
}

type DataBaseConfig struct {
	DriverName     string
	DataSourceName string
}
