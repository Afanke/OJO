package main

import (
	"fmt"
	"github.com/afanke/OJO/utils/log"
	"github.com/kataras/iris/v12"
)

var password string = ""

func main() {
	fmt.Printf("Please set a password:")
	_, err := fmt.Scanf("%s", &password)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.InitDefaultLog()
	app := iris.New()
	BindRoute(app)
	_ = app.Run(iris.Addr(":2333"))
}
