package ctrl

import (
	"github.com/gogotime/OJO/utils/log"
	"github.com/kataras/iris/v12"
	"io"
	"io/ioutil"
	"os"
)

type File struct {
}

func (File) Index(c iris.Context) {
	file, err := os.Open("./dist/index.html")
	if err != nil {
		c.NotFound()
		return
	}
	defer func() {
		err2 := file.Close()
		if err2 != nil {
			log.Error("%v\n", err)
			return
		}
	}()
	stat, err := file.Stat()
	if err != nil {
		c.NotFound()
		return
	}
	if modified, err := c.CheckIfModifiedSince(stat.ModTime()); !modified && err == nil {
		c.Header("Cache-Control", "max-age=64200")
		c.WriteNotModified()
		return
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		c.NotFound()
		return
	}
	c.SetLastModified(stat.ModTime())
	c.Header("Cache-Control", "max-age=64200")
	_, err = c.TryWriteGzip(bytes)
	if err != nil {
		log.Error("%v", err)
		return
	}
}

func (File) Admin(c iris.Context) {
	file, err := ioutil.ReadFile("./dist/admin.html")
	if err != nil {
		c.NotFound()
	}
	_, err = c.TryWriteGzip(file)
	if err != nil {
		log.Error("%v", err)
		return
	}
}

func (File) VDS(c iris.Context) {
	file, err := ioutil.ReadFile("./dist/vds.html")
	if err != nil {
		c.NotFound()
	}
	_, err = c.TryWriteGzip(file)
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
	_, _ = c.TryWriteGzip(file)
}

func (File) File(c iris.Context) {
	path := c.Path()
	stat, err := os.Stat("./dist" + path)
	if err != nil {
		c.NotFound()
		return
	}
	if modified, err := c.CheckIfModifiedSince(stat.ModTime()); !modified && err == nil {
		c.Header("Cache-Control", "public,max-age=64200")
		c.WriteNotModified()
		return
	}
	file, err := os.Open("./dist" + path)
	if err != nil {
		c.NotFound()
		return
	}
	defer func() {
		err2 := file.Close()
		if err2 != nil {
			log.Error("%v\n", err)
			return
		}
	}()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		c.NotFound()
		return
	}
	c.SetLastModified(stat.ModTime())
	c.Header("Cache-Control", "public,max-age=64200")
	_, err = c.TryWriteGzip(bytes)
	if err != nil {
		log.Error("%v", err)
		return
	}
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
