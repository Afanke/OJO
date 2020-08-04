package main

import (
	"fmt"
	"github.com/afanke/OJO/JudgeServer/dto"
	"github.com/afanke/OJO/JudgeServer/judge"
	"github.com/afanke/OJO/utils/log"
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

func GetRtJug(lid int64) judge.RtJug {
	switch lid {
	case 1:
		return judge.C{}
	case 2:
		return judge.Cpp{}
	case 3:
		return judge.Java{}
	case 4:
		return judge.Python{}
	case 5:
		return judge.Go{}
	default:
		return judge.C{}
	}
}

func GetSpJug(lid int64) judge.SpJug {
	switch lid {
	case 1:
		return judge.C{}
	case 2:
		return judge.Cpp{}
	case 3:
		return judge.Java{}
	case 4:
		return judge.Python{}
	case 5:
		return judge.Go{}
	default:
		return judge.C{}
	}
}

func GetJug(lid1, lid2 int64) judge.Base {
	return judge.NewJudge(GetRtJug(lid1), GetSpJug(lid2))
}

func checkForm(form *dto.JudgeForm) {
	if form.MaxRealTime < 0 {
		form.MaxRealTime = 0
	}
	if form.MaxCpuTime < 0 {
		form.MaxCpuTime = 0
	}
	if form.MaxMemory < 0 {
		form.MaxMemory = 0
	}
	if form.SPJMp <= 0 {
		form.SPJMp = 1
	}
	if form.CompMp <= 0 {
		form.CompMp = 1
	}
	form.TotalScore = 0
	form.Flag = "JUG"
	form.ErrorMsg = ""
}

func BindRoute(app *iris.Application) {
	app.Use(PanicMidWare)
	{
		app.Post("/touch", func(c iris.Context) {
			p := &struct {
				Password string `json:"password"`
			}{}
			err := c.ReadJSON(p)
			if err != nil {
				c.JSON(&dto.TouchResult{
					Message:   "invalid input",
					Connected: false,
				})
				return
			}
			if password != p.Password {
				c.JSON(&dto.TouchResult{
					Message:   "Wrong Password",
					Connected: false,
				})
				return
			}
			c.JSON(&dto.TouchResult{
				Message:   "Connected",
				Connected: true,
			})
		})
		app.Post("/judge", func(c iris.Context) {
			var form dto.JudgeForm
			err := c.ReadJSON(&form)
			if err != nil {
				for i, j := 0, len(form.TestCase); i < j; i++ {
					form.TestCase[i].Flag = "ISE"
				}
				form.Flag = "ISE"
				log.Error("%v", err)
				c.JSON(&form)
				return
			}
			log.Debug("%+v", &form)
			checkForm(&form)
			jug := GetJug(form.Lid, form.SPJLid)
			jug.Judge(&form)
			c.JSON(&form)
		})
	}
}
