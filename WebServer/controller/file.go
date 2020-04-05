package ctrl

import (
	"github.com/afanke/OJO/WebServer/db"
	"github.com/afanke/OJO/WebServer/dto"
	"github.com/afanke/OJO/utils/log"
	"github.com/kataras/iris/v12"
	"io"
	"io/ioutil"
	"os"
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

func (File) UploadImg(c iris.Context) {
	c.SetMaxRequestBodySize(2 * 1024 * 1024)
	file, header, err := c.FormFile("img")
	if err != nil {
		c.StatusCode(iris.StatusInternalServerError)
		log.Warn("error:%v\n", err)
		return
	}
	defer file.Close()
	// 创建一个具有相同名称的文件
	// 假设你有一个名为'uploads'的文件夹
	out, err := os.OpenFile("./dist/img/"+header.Filename,
		os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		c.StatusCode(iris.StatusInternalServerError)
		log.Warn("error:%v\n", err)
		return
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		c.StatusCode(iris.StatusInternalServerError)
		log.Warn("error:%v\n", err)
		return
	}
	c.JSON(struct {
		Url string `json:"url"`
	}{
		Url: "/img/" + header.Filename,
	})
}
