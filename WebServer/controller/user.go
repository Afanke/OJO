package ctrl

import (
	"errors"
	"fmt"
	"github.com/afanke/OJO/WebServer/db"
	"github.com/afanke/OJO/WebServer/dto"
	captcha "github.com/afanke/OJO/utils/chapcha"
	"github.com/afanke/OJO/utils/session"
	"github.com/kataras/iris/v12"
	"image/png"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type User struct{}

var userdb = db.User{}

var ImgMaxSize int64 = 2 << 20 // 2MB

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
	if !res.Enabled {
		c.JSON(&dto.Res{Error: errors.New("you are not allowed to login").Error(), Data: nil})
		return
	}
	err = userdb.UpdateLoginTime(res.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	s.Set("userId", res.Id)
	c.JSON(&dto.Res{Error: "", Data: res})
}

func (User) AdminLogin(c iris.Context) {
	var loginForm dto.LoginForm
	err := c.ReadJSON(&loginForm)
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
	fmt.Println(loginForm)
	res, err := userdb.Query(loginForm.Username, loginForm.Password)
	if err != nil {
		c.JSON(&dto.Res{Error: errors.New("username or password not correct").Error(), Data: nil})
		return
	}
	if !res.Enabled {
		c.JSON(&dto.Res{Error: errors.New("you are not allowed to login").Error(), Data: nil})
		return
	}
	if res.Type < 2 {
		c.JSON(&dto.Res{Error: errors.New("username or password not correct").Error(), Data: nil})
		return
	}
	err = userdb.UpdateLoginTime(res.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	err = session.SetInt64(c, "userId", res.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: res})
}

func (User) Login1(c iris.Context) {
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
	s.Set("userId", res.Id)
	c.JSON(&dto.Res{Error: "", Data: res})
}

func (User) GetInfo(c iris.Context) {
	s, err := session.GetSession(c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	res := s.Get("userId")
	if res == nil {
		c.JSON(&dto.Res{Error: errors.New("not log in").Error(), Data: nil})
		return
	}
	user := res.(dto.UserToken)
	c.JSON(&dto.Res{Error: "", Data: user})
}

func (User) UploadImg(c iris.Context) {
	userId, err := getUserId(c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	file, info, err := c.FormFile("file")
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	defer file.Close()
	if info.Size > ImgMaxSize {
		c.JSON(&dto.Res{Error: errors.New("file to large").Error(), Data: nil})
		return
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	path := "/img/data/" + strconv.Itoa(int(userId)) + "_" +
		time.Now().Format("2006_01_02_15_04_05") + "_" + info.Filename
	err = ioutil.WriteFile("./dist"+path, bytes, 0666)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	err = userdb.UpdateIcon(userId, path)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: "upload icon successfully"})
}

func (User) GetAdminInfo(c iris.Context) {
	s, err := session.GetSession(c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	res := s.Get("userId")
	if res == nil {
		c.JSON(&dto.Res{Error: errors.New("not log in").Error(), Data: nil})
		return
	}
	user := res.(dto.UserToken)
	if user.Type < 2 {
		c.JSON(&dto.Res{Error: errors.New("you are not admin").Error(), Data: nil})
	}
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
	s.Set("captcha", strings.ToLower(code))
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
	s.Remove("userId")
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
	if id.Id == 0 {
		userId, err := getUserId(c)
		if err != nil {
			c.JSON(&dto.Res{Error: err.Error(), Data: nil})
			return
		}
		id.Id = userId
	}
	data, err := userdb.GetDetail(id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: data})
}

func (User) GetStatistic(c iris.Context) {
	var id dto.Id
	err := c.ReadJSON(&id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	var res dto.UserStatistic
	data, err := pctdb.GetUserACCount(id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	res.AC = data
	data, err = pctdb.GetUserSubmissionCount(id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	res.Submission = data
	data, err = pctdb.GetUserScore(id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	res.Score = data
	sl, err := pctdb.GetUserSolvedList(id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	res.SolvedList = sl
	c.JSON(&dto.Res{Error: "", Data: res})
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
	s, err := session.GetSessionByInt64("userId", form.Id)
	if err == nil {
		if token, ok := s.Get("data").(dto.UserToken); ok {
			token.Username = form.Username
			s.Set("data", token)
		}
	}
	c.JSON(&dto.Res{Error: "", Data: "update data successfully"})
}

func (User) UpdateProfile(c iris.Context) {
	var form dto.UserDetail
	err := c.ReadJSON(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	id, err := getUserId(c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	form.Id = id
	err = userdb.UpdateProfile(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: "update profile successfully"})
}

func (User) UpdatePassword(c iris.Context) {
	var form dto.UpdateForm
	err := c.ReadJSON(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	id, err := getUserId(c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	form.Id = id
	err = userdb.CheckPassword(form.Id, form.Password)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	err = userdb.UpdatePassword(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: "update password successfully"})
}

func (User) UpdateEmail(c iris.Context) {
	var form dto.UpdateForm
	err := c.ReadJSON(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	id, err := getUserId(c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	form.Id = id
	err = userdb.CheckPassword(form.Id, form.Password)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	err = userdb.UpdateEmail(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: "update email successfully"})
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
	c.JSON(&dto.Res{Error: "", Data: "update data successfully"})
}

func (User) Disable(c iris.Context) {
	var id dto.Id
	err := c.ReadJSON(&id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	userId, err := isSuperAdmin(c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	if userId == id.Id {
		c.JSON(&dto.Res{Error: errors.New("can't disable yourself").Error(), Data: nil})
		return
	}
	err = userdb.Disable(id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	session.DelByInt64("userId", id.Id)
	c.JSON(&dto.Res{Error: "", Data: "update data successfully"})
}

func getUserId(c iris.Context) (int64, error) {
	id, err := session.GetInt64(c, "userId")
	if err != nil {
		return 0, errors.New("please login ")
	}
	return id, nil
}

func isAdmin(c iris.Context) (int64, error) {
	userId, err := session.GetInt64(c, "userId")
	if err != nil {
		return 0, errors.New("please login")
	}
	userType, err := userdb.GetUserType(userId)
	if userType < 2 {
		return 0, errors.New("not allowed")
	}
	return userId, err
}

func isSuperAdmin(c iris.Context) (int64, error) {
	userId, err := session.GetInt64(c, "userId")
	if err != nil {
		return 0, errors.New("not login in or not permitted")
	}
	userType, err := userdb.GetUserType(userId)
	if userType < 3 {
		return 0, errors.New("not allowed")
	}
	return userId, err
}
