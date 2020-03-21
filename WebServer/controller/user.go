package ctrl

import (
	"errors"
	"fmt"
	"github.com/afanke/OJO/WebServer/db"
	"github.com/afanke/OJO/WebServer/dto"
	captcha "github.com/afanke/OJO/utils/chapcha"
	"github.com/afanke/OJO/utils/session"
	"github.com/kataras/iris"
	"image/png"
)

type User struct{}

var userdb = db.User{}

func (User) Login(c iris.Context) {
	var loginForm dto.LoginForm
	err := c.ReadJSON(&loginForm)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	s, err := session.GetSession(c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	cp, ok := s.Get("captcha").(string)
	if !ok {
		c.JSON(&dto.Res{Error: errors.New("please refresh your captcha").Error(), Data: nil})
		return
	}
	if cp != loginForm.Captcha {
		c.JSON(&dto.Res{Error: errors.New("captcha is not correct").Error(), Data: nil})
		return
	}
	res, err := userdb.Query(loginForm.Username, loginForm.Password)
	if err != nil {
		c.JSON(&dto.Res{Error: errors.New("username or password not correct").Error(), Data: nil})
		return
	}
	s.Set("user", res)
	c.JSON(&dto.Res{Error: "", Data: res})
}

func (User) GetInfo(c iris.Context) {
	s, err := session.GetSession(c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	res := s.Get("user")
	if res == nil {
		c.JSON(&dto.Res{Error: errors.New("not log in").Error(), Data: nil})
		return
	}
	user := res.(dto.User)
	c.JSON(&dto.Res{Error: "", Data: user})
}

func (User) Captcha(c iris.Context) {
	//服务器通知浏览器不要缓存
	c.Header("pragma", "no-cache")
	c.Header("cache-control", "no-cache")
	c.Header("expires", "0")
	cp := captcha.NewCaptcha(120, 40, 4)
	cp.SetFontPath("./dist/fonts/Wesley.ttf")
	cp.SetMode(1) // 设置为数学公式
	code, img := cp.OutPut()
	s, err := session.GetSession(c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	s.Set("captcha", code)
	//备注：code 可以根据情况存储到session，并在使用时取出验证
	fmt.Println(code)
	_ = png.Encode(c.ResponseWriter(), img)

}

func (User) Logout(c iris.Context) {
	s, err := session.GetSession(c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	s.Remove("user")
	c.JSON(&dto.Res{Error: "", Data: "success to log out"})

}

func (User) Register(c iris.Context) {
	var regForm dto.RegisterForm
	err := c.ReadJSON(&regForm)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	s, err := session.GetSession(c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	cp, ok := s.Get("captcha").(string)
	if !ok {
		c.JSON(&dto.Res{Error: errors.New("please refresh your captcha").Error(), Data: nil})
		return
	}
	if cp != regForm.Captcha {
		c.JSON(&dto.Res{Error: errors.New("captcha is not correct").Error(), Data: nil})
		return
	}
	err = userdb.Insert(&regForm)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: "success to register"})
}
