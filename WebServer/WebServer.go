package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.Default()
	// yaag.Init(&yaag.Config{
	// 	On:       true,                 //是否开启自动生成API文档功能
	// 	DocTitle: "Iris",
	// 	DocPath:  "apidoc.html",        //生成API文档名称存放路径
	// 	BaseUrls: map[string]string{"Production": "", "Staging": ""},
	// })
	//注册中间件
	// app.Use(irisyaag.New())
	BindRoute(app)
	app.Run(iris.Addr(":80"))
}

// type myXML struct {
// 	Result string `xml:"result"`
// }
