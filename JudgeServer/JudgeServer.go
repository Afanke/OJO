package main

import (
	"github.com/afanke/OJO/JudgeServer/operator"
	"github.com/kataras/iris/v12"
)

var py3 operator.PythonOperator

func main() {
	app := iris.New()
	BindRoute(app)
	_ = app.Run(iris.Addr(":2333"))
}
