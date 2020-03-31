package ctrl

import (
	"errors"
	"fmt"
	"github.com/afanke/OJO/WebServer/db"
	"github.com/afanke/OJO/WebServer/dto"
	captcha "github.com/afanke/OJO/utils/chapcha"
	"github.com/afanke/OJO/utils/log"
	"github.com/afanke/OJO/utils/session"
	"github.com/kataras/iris"
	"image/png"
	"math/rand"
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
	// cp, ok := s.Get("captcha").(string)
	// if !ok {
	// 	c.JSON(&dto.Res{Error: errors.New("please refresh your captcha").Error(), Data: nil})
	// 	return
	// }
	// if cp != loginForm.Captcha {
	// 	c.JSON(&dto.Res{Error: errors.New("captcha is not correct").Error(), Data: nil})
	// 	return
	// }
	res, err := userdb.Query(loginForm.Username, loginForm.Password)
	if err != nil {
		c.JSON(&dto.Res{Error: errors.New("username or password not correct").Error(), Data: nil})
		return
	}
	if !res.Enabled {
		c.JSON(&dto.Res{Error: errors.New("you are not allowed to login").Error(), Data: nil})
		return
	}
	err = userdb.UpdateLoginTime(res.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	s.Set("user", res)
	s.Set("userid", res.Id)
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
	user := res.(dto.UserToken)
	c.JSON(&dto.Res{Error: "", Data: user})
}

func (User) Captcha(c iris.Context) {
	// 服务器通知浏览器不要缓存
	c.Header("pragma", "no-cache")
	c.Header("cache-control", "no-cache")
	c.Header("expires", "0")
	cp := captcha.NewCaptcha(120, 40, 4)
	cp.SetFontPath("./config/xindexingcao57.ttf")
	cp.SetMode(rand.Int() & 1) // 设置为数学公式
	code, img := cp.OutPut()
	s, err := session.GetSession(c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	s.Set("captcha", code)
	// 备注：code 可以根据情况存储到session，并在使用时取出验证
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

func (User) GetAll(c iris.Context) {
	var form dto.UserForm
	err := c.ReadJSON(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	_, err = isAdmin(c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	data, err := userdb.GetAll(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: data})
}

func (User) GetCount(c iris.Context) {
	var form dto.UserForm
	err := c.ReadJSON(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	_, err = isAdmin(c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	data, err := userdb.GetCount(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: data})
}

func (User) GetDetail(c iris.Context) {
	var id dto.Id
	err := c.ReadJSON(&id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	_, err = isSuperAdmin(c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	data, err := userdb.GetDetail(id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: data})
}

func (User) UpdateDetail(c iris.Context) {
	var form dto.UserDetail2
	err := c.ReadJSON(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	_, err = isSuperAdmin(c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	err = userdb.UpdateDetail(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	s, err := session.GetSessionByInt64("userid", form.Id)
	if err == nil {
		if token, ok := s.Get("user").(dto.UserToken); ok {
			token.Type = form.Type
			token.Username = form.Username
			token.IconPath = form.IconPath
			token.RealName = form.RealName
			s.Set("user", token)
		}
	}
	c.JSON(&dto.Res{Error: "", Data: "update user successfully"})
}

func (User) Enable(c iris.Context) {
	var id dto.Id
	err := c.ReadJSON(&id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	_, err = isSuperAdmin(c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	err = userdb.Enable(id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: "update user successfully"})
}

func (User) Disable(c iris.Context) {
	var id dto.Id
	err := c.ReadJSON(&id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	spAdmin, err := isSuperAdmin(c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	if spAdmin.Id == id.Id {
		c.JSON(&dto.Res{Error: errors.New("can't disable yourself").Error(), Data: nil})
		return
	}
	err = userdb.Disable(id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	session.DelByInt64("userid", id.Id)
	c.JSON(&dto.Res{Error: "", Data: "update user successfully"})
}

func getUserToken(c iris.Context) (*dto.UserToken, error) {
	get, err := session.Get(c, "user")
	if err != nil {
		return nil, errors.New("please login")
	}
	if user, ok := get.(dto.UserToken); !ok {
		return nil, errors.New("please login")
	} else {
		return &user, nil
	}
}

func getUserId(c iris.Context) (int64, error) {
	id, err := session.GetInt64(c, "userid")
	if err != nil {
		return 0, errors.New("please login")
	}
	return id, nil
}

func isAdmin(c iris.Context) (*dto.UserToken, error) {
	get, err := session.Get(c, "user")
	if err != nil {
		return nil, errors.New("not login in or not permitted")
	}
	if user, ok := get.(dto.UserToken); !ok {
		return nil, errors.New("not login in or not permitted")
	} else {
		log.Debug("%#v", user)
		if user.Type < 2 {
			return nil, errors.New("not permitted")
		}
		return &user, nil
	}
}

func isSuperAdmin(c iris.Context) (*dto.UserToken, error) {
	get, err := session.Get(c, "user")
	if err != nil {
		return nil, errors.New("not login in or not permitted")
	}
	if user, ok := get.(dto.UserToken); !ok {
		return nil, errors.New("not login in or not permitted")
	} else {
		if user.Type < 3 {
			return nil, errors.New("not permitted")
		}
		return &user, nil
	}
}
