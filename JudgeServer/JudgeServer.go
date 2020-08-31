package main

import (
	"fmt"
	"github.com/gogotime/OJO/utils/log"
	"github.com/kataras/iris/v12"
)

var password string = ""

var port string = ""

func main() {
	fmt.Printf("Please enter the listening port:")
	_, err := fmt.Scanf("%s", &port)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Please set a password:")
	_, err = fmt.Scanf("%s", &password)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.InitDefaultLog(log.INFO)
	app := iris.New()
	BindRoute(app)
	_ = app.Run(iris.Addr(":" + port))
}
