package main

import (
	"fmt"
	ctrl "github.com/afanke/OJO/WebServer/controller"
	"github.com/afanke/OJO/WebServer/dto"
	"github.com/afanke/OJO/utils/log"
	"github.com/afanke/OJO/utils/randstr"
	"github.com/afanke/OJO/utils/session"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"net/http"
	"runtime"
	"time"
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
			log.Error("Recovered from a route's Handler('%s')", ctx.HandlerName())
			log.Error("At Request: %s", ctx.Path())
			log.Error("Trace: %s", err)
			log.Error("%s", stacktrace)
			ctx.StatusCode(500)
			ctx.StopExecution()
		}
	}()

	ctx.Next()
}
func ReqMidWare(ctx iris.Context) {
	var latency time.Duration
	var startTime, endTime time.Time
	startTime = time.Now()

	ctx.Next()

	endTime = time.Now()
	latency = endTime.Sub(startTime)
	ip := ctx.RemoteAddr()
	method := ctx.Method()
	path := ctx.Path()
	status := ctx.GetStatusCode()
	log.Info("%d %s %s %s %s", status, latency, ip, method, path)
}
func CorsMidWare(c iris.Context) {
	if c.Request().Method == "OPTIONS" {
		c.Header("Access-Control-Allow-Origin", "http://localhost:8080")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.StatusCode(204)
		return
	}
	c.Header("Access-Control-Allow-Origin", "http://localhost:8080")
	c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Next()
}
func SessionMidWare(c iris.Context) {
	cookie := c.GetCookie("GOGONEWWORLD")
	if cookie == "" {
		c.SetCookieKV("GOGONEWWORLD", randstr.BigRandN(16), func(c *http.Cookie) {
			c.MaxAge = 0
		})
	}
	c.Next()
}
func TemUserMidWare(c iris.Context) {
	s, err := session.GetSession(c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	user := s.Get("user")
	_, ok := user.(dto.User)
	if !ok {
		s.Set("user", dto.User{Id: 1, Username: "visitor"})
	}
	c.Next()
}

func BindRoute(app *iris.Application) {
	app.Use(PanicMidWare)
	app.Use(SessionMidWare)
	app.Use(CorsMidWare)
	app.Use(ReqMidWare)
	// app.Use(TemUserMidWare)
	{
		var file ctrl.File
		app.Get("/pp", func(context context.Context) {
			log.Fatal("asdwer")
		})
		app.Get("/", file.Index)
		app.Get("/favicon.ico", file.Favicon)
		app.Get("/img/*", file.File)
		app.Get("/fonts/*", file.File)
		app.Get("/css/*", file.File)
		app.Get("/js/*", file.File)
		app.Post("/getProgress", file.GetProgress)
		app.Options("*", func(c iris.Context) {
			c.Next()
		})
	}
	user := app.Party("/user")
	{
		var u ctrl.User
		user.Post("/getInfo", u.GetInfo)
		user.Post("/login", u.Login)
		user.Get("/captcha", u.Captcha)
		user.Post("/logout", u.Logout)
		user.Post("/register", u.Register)
	}
	practice := app.Party("/practice")
	{
		var p ctrl.Practice
		practice.Post("/getAll", p.GetAll)
		practice.Post("/getTags", p.GetTags)
		practice.Post("/getCount", p.GetCount)
		practice.Post("/getDetail", p.GetDetail)
		practice.Post("/getCurrentStatus", p.GetCurrentStatus)
		practice.Post("/getStatus", p.GetStatus)
		practice.Post("/getStatusDetail", p.GetStatusDetail)
		practice.Post("/getAllStatus", p.GetAllStatus)
		practice.Post("/getAllStatusCount", p.GetAllStatusCount)
		practice.Post("/submit", p.Submit)
	}
	contest := app.Party("/contest")
	{
		var c ctrl.Contest
		contest.Post("/getAll", c.GetAll)
		contest.Post("/getCount", c.GetCount)
		contest.Post("/getDetail", c.GetDetail)
		contest.Post("/getQualification", c.GetQualification)
		contest.Post("/qualify", c.Qualify)
		contest.Post("/getAllProblem", c.GetAllProblem)
		contest.Post("/getAllProblemName", c.GetAllProblemName)
		contest.Post("/getProblemDetail", c.GetProblemDetail)
		contest.Post("/submit", c.Submit)
		contest.Post("/getStatus", c.GetStatus)
		contest.Post("/getCurrentStatus", c.GetCurrentStatus)
		contest.Post("/getStatusDetail", c.GetStatusDetail)
		contest.Post("/getAllStatus", c.GetAllStatus)
		contest.Post("/getAllStatusCount", c.GetAllStatusCount)
		contest.Post("/getTime", c.GetTime)
		contest.Post("/getOIRank", c.GetOIRank)
		contest.Post("/getOITop10", c.GetOITop10)
		contest.Post("/getOIRankCount", c.GetOIRankCount)
		contest.Post("/getACMRank", c.GetACMRank)
		contest.Post("/getACMTop10", c.GetACMTop10)
		contest.Post("/getACMRankCount", c.GetACMRankCount)
	}
}
