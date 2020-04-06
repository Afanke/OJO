package main

import (
	"fmt"
	"github.com/afanke/OJO/JudgeServer/dto"
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
		app.Post("/Python3", func(c iris.Context) {
			var forms []dto.OperationForm
			err := c.ReadJSON(&forms)
			fmt.Println(1, forms)
			if err != nil {
				for i, j := 0, len(forms); i < j; i++ {
					forms[i].Flag = "ISE"
				}
				c.JSON(&forms)
				return
			}

			for i, j := 0, len(forms); i < j; i++ {
				py3.Operate(&forms[i])
			}
			fmt.Println(2, forms)
			c.JSON(&forms)
		})
	}
}
