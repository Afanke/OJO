package ctrl

import (
	"fmt"
	"github.com/afanke/OJO/WebServer/db"
	"github.com/afanke/OJO/WebServer/dto"
	"github.com/kataras/iris"
	"io/ioutil"
)

type File struct {
}

func (File) Index(c iris.Context) {
	file, err := ioutil.ReadFile("./dist/index.html")
	if err != nil {
		c.NotFound()
	}
	c.WriteGzip(file)
}

func (File) Favicon(c iris.Context) {
	file, err := ioutil.ReadFile("./dist/favicon.ico")
	if err != nil {
		c.NotFound()
	}
	c.WriteGzip(file)
}

func (File) File(c iris.Context) {
	path := c.Path()
	fmt.Println(path)
	file, err := ioutil.ReadFile("./dist" + path)
	if err != nil {
		c.NotFound()
	}
	c.WriteGzip(file)
}

func (File) GetProgress(c iris.Context) {
	progress := db.GetProgress()
	_, _ = c.JSON(&dto.Res{Error: "", Data: progress})
}
