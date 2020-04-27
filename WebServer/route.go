package main

import (
	"fmt"
	ctrl "github.com/afanke/OJO/WebServer/controller"
	"github.com/afanke/OJO/WebServer/dto"
	jsp "github.com/afanke/OJO/WebServer/judge"
	"github.com/afanke/OJO/utils/log"
	"github.com/afanke/OJO/utils/randstr"
	"github.com/afanke/OJO/utils/session"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
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
		c.Header("Access-Control-Allow-Origin", c.Request().Header.Get("Origin"))
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.StatusCode(204)
		return
	}
	c.Header("Access-Control-Allow-Origin", c.Request().Header.Get("Origin"))
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
		var sys ctrl.System
		app.Get("/", file.Index)
		app.Get("/admin", file.Admin)
		app.Get("/vds", file.VDS)
		app.Get("/favicon.ico", file.Favicon)
		app.Get("/img/*", file.File)
		app.Get("/fonts/*", file.File)
		app.Get("/css/*", file.File)
		app.Get("/js/*", file.File)
		app.Get("/sys/getWebConfig", sys.GetWebConfig)
		app.Post("/getProgress", file.GetProgress)
		app.Post("/uploadImg", file.UploadImg)
		app.Options("*", func(c iris.Context) {
			c.Next()
		})
	}
	user := app.Party("/user")
	{
		var u ctrl.User
		user.Post("/getInfo", u.GetInfo)
		user.Post("/login", u.Login)
		user.Post("/login1", u.Login1)
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
		contest.Post("/getAll", c.GetAllVisible)
		contest.Post("/getCount", c.GetVisibleCount)
		contest.Post("/getDetail", c.GetVisibleDetail)
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
	admin := app.Party("/admin")
	{
		var pb ctrl.Problem
		{
			admin.Post("/problem/addProblem", pb.AddProblem)
			admin.Post("/problem/getAll", pb.GetAll)
			admin.Post("/problem/getCount", pb.GetCount)
			admin.Post("/problem/setVisibleTrue", pb.SetVisibleTrue)
			admin.Post("/problem/setVisibleFalse", pb.SetVisibleFalse)
			admin.Post("/problem/tryEdit", pb.TryEdit)
			admin.Post("/problem/getDetail", pb.GetDetail)
			admin.Post("/problem/updateProblem", pb.UpdateProblem)
		}
		var user ctrl.User
		{
			admin.Post("/user/getAll", user.GetAll)
			admin.Post("/user/getCount", user.GetCount)
			admin.Post("/user/getDetail", user.GetDetail)
			admin.Post("/user/updateDetail", user.UpdateDetail)
			admin.Post("/user/enable", user.Enable)
			admin.Post("/user/disable", user.Disable)
		}
		var tag ctrl.Tag
		{
			admin.Post("/tag/getAll", tag.GetAll)
			admin.Get("/tag/getAllShared", tag.GetAllShared)
			admin.Post("/tag/getCount", tag.GetCount)
			admin.Post("/tag/getAllVisible", tag.GetAllVisible)
			admin.Post("/tag/setVisibleTrue", tag.SetVisibleTrue)
			admin.Post("/tag/setVisibleFalse", tag.SetVisibleFalse)
			admin.Post("/tag/setSharedTrue", tag.SetSharedTrue)
			admin.Post("/tag/setSharedFalse", tag.SetSharedFalse)
			admin.Post("/tag/addTag", tag.AddTag)
			admin.Post("/tag/updateTag", tag.UpdateTag)
			admin.Post("/tag/deleteTag", tag.DeleteTag)
		}
		var cts ctrl.Contest
		{
			admin.Post("/contest/getAll", cts.GetAll)
			admin.Post("/contest/getCount", cts.GetCount)
			admin.Post("/contest/setVisibleTrue", cts.SetVisibleTrue)
			admin.Post("/contest/setVisibleFalse", cts.SetVisibleFalse)
			admin.Post("/contest/addContest", cts.AddContest)
			admin.Post("/contest/tryEdit", cts.TryEdit)
			admin.Post("/contest/getDetail", cts.GetDetail)
			admin.Post("/contest/updateContest", cts.UpdateContest)
		}
		var sys ctrl.System
		{
			admin.Get("/sys/getAll", sys.GetAll)
			admin.Post("/sys/updateSMTP", sys.UpdateSMTP)
			admin.Post("/sys/updateWeb", sys.UpdateWeb)
		}
		var anno ctrl.Announcement
		{
			admin.Post("/announcement/getAll", anno.GetAll)
			admin.Post("/announcement/getCount", anno.GetCount)
			admin.Post("/announcement/setVisibleTrue", anno.SetVisibleTrue)
			admin.Post("/announcement/setVisibleFalse", anno.SetVisibleFalse)
			admin.Post("/announcement/addAnnouncement", anno.AddAnnouncement)
			admin.Post("/announcement/getDetail", anno.GetDetail)
			admin.Post("/announcement/updateAnnouncement", anno.UpdateAnnouncement)
		}
		{
			admin.Get("/jsp/getAllInfo", func(c context.Context) {
				_, _ = c.JSON(&dto.Res{
					Error: "",
					Data:  jsp.GetAllInfo(),
				})
			})
		}

	}
}
