package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var Config ServerConfig

func init() {
	file, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		fmt.Printf("error:%v\n", err)
		os.Exit(-1)
	}
	err = json.Unmarshal(file, &Config)
	if err != nil {
		fmt.Printf("error:%v\n", err)
		os.Exit(-1)
	}
	// fmt.Println(Config)
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
