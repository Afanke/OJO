package ctrl

import (
	"github.com/afanke/OJO/WebServer/db"
	"github.com/afanke/OJO/WebServer/dto"
	"github.com/afanke/OJO/utils/log"
	"github.com/kataras/iris"
	"io/ioutil"
)

type File struct {
}

func (File) Index(c iris.Context) {
	file, err := ioutil.ReadFile("./dist/index2.html")
	if err != nil {
		c.NotFound()
	}
	_, err = c.WriteGzip(file)
	if err != nil {
		log.Error("%v", err)
		return
	}
}

func (File) Favicon(c iris.Context) {
	file, err := ioutil.ReadFile("./dist/favicon.ico")
	if err != nil {
		c.NotFound()
	}
	_, _ = c.WriteGzip(file)
}

func (File) File(c iris.Context) {
	path := c.Path()
	file, err := ioutil.ReadFile("./dist" + path)
	if err != nil {
		c.NotFound()
	}
	_, _ = c.WriteGzip(file)
}

func (File) GetProgress(c iris.Context) {
	progress := db.GetProgress()
	_, _ = c.JSON(&dto.Res{Error: "", Data: progress})
}
