package main

import (
	"context"
	ctrl "github.com/gogotime/OJO/WebServer/controller"
	"github.com/gogotime/OJO/utils/log"
	"github.com/gogotime/OJO/utils/session"
	"github.com/kataras/iris/v12"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	app := iris.New()
	log.InitLog()
	session.Register(ctrl.RSTForm{})
	session.InitSession()
	stop := make(chan int, 1)
	go func() {
		http.ListenAndServe("localhost:8082", nil)
		log.Debug("start to pprof")
	}()
	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch,
			syscall.SIGQUIT,
			syscall.SIGINT,
			syscall.SIGKILL,
			syscall.SIGTERM,
			syscall.SIGABRT,
		)
		select {
		case <-ch:
			session.SaveSession()
			timeout := 5 * time.Second
			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()
			ctx.Done()
			err := app.Shutdown(ctx)
			if err != nil {
				log.Error("%v", err)
				return
			}
			stop <- 1
		}
	}()
	BindRoute(app)
	// yaag.Init(&yaag.Config{
	// 	On:       true,                 //是否开启自动生成API文档功能
	// 	DocTitle: "Iris",
	// 	DocPath:  "apidoc.html",        //生成API文档名称存放路径
	// 	BaseUrls: map[string]string{"Production": "", "Staging": ""},
	// })
	// 注册中间件
	// app.Use(irisyaag.New())
	_ = app.Run(iris.Addr(":80"))
	<-stop
}
