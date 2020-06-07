package main

import (
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	BindRoute(app)
	_ = app.Run(iris.Addr(":2333"))
}
