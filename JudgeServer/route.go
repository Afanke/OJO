package main

import (
	"fmt"
	"github.com/afanke/OJO/JudgeServer/dto"
	"github.com/afanke/OJO/JudgeServer/operator"
	"github.com/kataras/iris/v12"
	"runtime"
)

func PanicMidWare(ctx iris.Context) {
	defer func() {
		if err := recover(); err != nil {
			if ctx.IsStopped() {
				return
			}
			var stacktrace string
			for i := 1; ; i++ {
				_, f, l, got := runtime.Caller(i)
				if !got {
					break
				}
				stacktrace += fmt.Sprintf("%s:%d\n", f, l)
			}
			// when stack finishes
			println("Recovered from a route's Handler('%s')", ctx.HandlerName())
			println("At Request: %s", ctx.Path())
			println("Trace: %s", err)
			println("%s", stacktrace)
			ctx.StatusCode(500)
			ctx.StopExecution()
		}
	}()

	ctx.Next()
}

func BindRoute(app *iris.Application) {
	app.Use(PanicMidWare)
	{
		app.Get("/touch", func(c iris.Context) {
			c.JSON(dto.Res{Error: "", Data: "success"})
		})
		app.Get("/judgePython3", func(c iris.Context) {
			var form dto.OperationForm
			err := c.ReadJSON(form)
			if err != nil {
				c.JSON(&dto.Res{
					Error: err.Error(),
					Data:  nil,
				})
				return
			}
			operator.PythonOperator{}.Operate(&form)
			c.JSON(&dto.Res{Error: "", Data: form})
		})
	}
}
