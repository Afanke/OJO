package main

import (
	"context"
	"github.com/afanke/OJO/utils/log"
	"github.com/afanke/OJO/utils/session"
	"github.com/kataras/iris/v12"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	app := iris.New()
	stop := make(chan int, 1)
	go func() {
		ch := make(chan os.Signal, 5)
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
