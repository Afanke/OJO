package jsp

// judge server pool
import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/afanke/OJO/WebServer/db"
	"github.com/afanke/OJO/WebServer/dto"
	"github.com/afanke/OJO/utils/log"
	"github.com/afanke/OJO/utils/session"
	"github.com/ilibs/gosql/v2"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var (
	jsp     []dto.JudgeServer
	lens    int
	current int
	count   int
	lock    sync.RWMutex
)

var userdb = db.User{}
var pbdb = db.Problem{}

func init() {
	initJSP()
	TouchJSP()
	go func() {
		for {
			<-time.After(5 * time.Second)
			TouchJSP()
		}
	}()
}

func GetAddr() (string, error) {
	lock.RLock()
	defer lock.RUnlock()
	en := 0
	for i, j := 0, len(jsp); i < j; i++ {
		if jsp[i].Enabled {
			en++
		}
	}
	if en == 0 {
		return "", errors.New("now judge server available, please wait a minute or contact with the admin")
	}
	for {
		if count > jsp[current].Weight {
			current++
			current = current % lens
			count = 0
		} else {
			count++
		}
		if jsp[current].Connected {
			break
		} else {
			current++
			current = current % lens
			count = 0
		}
	}
	return jsp[current].Address + ":" + strconv.Itoa(jsp[current].Port), nil
}

func initJSP() {
	var js []dto.JudgeServer
	err := gosql.Select(&js, "select id, name, address, port, weight, enabled,password from ojo.judge_server")
	if err != nil {
		log.Fatal("init jsp server failed:%v", err)
		return
	}
	lock.Lock()
	defer lock.Unlock()
	jsp = js
	lens = len(js)
}

func UpdateJSP() {
	var js []dto.JudgeServer
	err := gosql.Select(&js, "select id, name, address, port, weight, enabled,password from ojo.judge_server")
	if err != nil {
		log.Error("update jsp server failed:%v", err)
	}
	lock.Lock()
	defer lock.Unlock()
	jsp = js
	lens = len(js)
	current = 0
	count = 0
}

func TouchJSP() {
	lock.RLock()
	defer lock.RUnlock()
	// log.Debug("%v", len(jsp))
	for i, j := 0, len(jsp); i < j; i++ {
		// log.Debug("%v", i)
		k := i
		if jsp[k].Enabled {
			go func() {
				client := &http.Client{
					Timeout: 1 * time.Second,
				}
				p := &struct {
					Password string `json:"password"`
				}{}
				p.Password = jsp[k].Password
				buff, err := json.Marshal(p)
				if err != nil {
					log.Error("error:%v", err)
					jsp[k].Connected = false
					jsp[k].Message = "JSON failed"
					return
				}
				res, err := client.Post("http://"+jsp[k].Address+":"+strconv.Itoa(jsp[k].Port)+"/touch", "application/json", bytes.NewBuffer(buff))
				if err != nil {
					log.Error("error:%v", err)
					jsp[k].Connected = false
					jsp[k].Message = "Connection Failed"
					return
				}
				defer res.Body.Close()
				body, err := ioutil.ReadAll(res.Body)
				if err != nil {
					log.Error("error:%v", err)
					jsp[k].Connected = false
					jsp[k].Message = "Connection Failed"
					return
				}
				var rest dto.TouchResult
				err = json.Unmarshal(body, &rest)
				if err != nil {
					log.Error("error:%v", err)
					jsp[k].Connected = false
					jsp[k].Message = "Connection Failed"
					return
				}
				jsp[k].Connected = rest.Connected
				jsp[k].Message = rest.Message
			}()
		}
	}
}

func SendToJudge(form *dto.JudgeForm) (*dto.JudgeForm, error) {
	addr, err := GetAddr()
	if err != nil {
		log.Error("error:%v", err)
		return nil, err
	}
	timeOut := time.Millisecond * time.Duration(4*form.MaxRealTime*form.SPJMp*form.CompMp)
	client := &http.Client{
		Timeout: timeOut,
	}
	buff, err := json.Marshal(form)
	if err != nil {
		log.Error("error:%v", err)
		return nil, err
	}
	res, err := client.Post("http://"+addr+"/judge", "application/json", bytes.NewBuffer(buff))
	if err != nil {
		log.Error("error:%v", err)
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error("error:%v", err)
		return nil, err
	}
	err = json.Unmarshal(body, &form)
	if err != nil {
		log.Error("error:%v", err)
		return nil, err
	}
	return form, err
}

func PrepareForm(subForm *dto.SubmitForm) (*dto.JudgeForm, error) {
	useSPJ, err := pbdb.UseSPJ(subForm.Pid)
	if err != nil {
		log.Warn("error:%v", err)
		return nil, err
	}
	limit, err := pbdb.GetLimitByLid(subForm.Pid, subForm.Lid)
	if err != nil {
		log.Warn("error:%v", err)
		return nil, err
	}
	t, err := pbdb.GetTemplateByLid(subForm.Pid, subForm.Lid)
	if err != nil {
		log.Warn("error:%v", err)
		return nil, err
	}
	code := t.Prepend + subForm.Code + t.Append
	form := &dto.JudgeForm{
		UseSPJ:      useSPJ,
		MaxCpuTime:  limit.MaxCpuTime,
		MaxRealTime: limit.MaxRealTime,
		MaxMemory:   limit.MaxMemory,
		TotalScore:  0,
		CompMp:      limit.CompMp,
		SPJMp:       limit.SPJMp,
		Id:          subForm.Sid,
		Lid:         subForm.Lid,
		Sid:         subForm.Sid,
		Pid:         subForm.Pid,
		Cid:         subForm.Cid,
		Uid:         subForm.Uid,
		SPJLid:      0,
		SPJCode:     "",
		Code:        code,
		Flag:        "",
		ErrorMsg:    "",
		TestCase:    nil,
	}
	if useSPJ {
		spj, err := pbdb.GetSPJ(subForm.Pid)
		if err != nil {
			log.Warn("error:%v", err)
			return nil, err
		}
		form.SPJCode = spj.Code
		form.SPJLid = spj.Lid
	}
	cases, err := pbdb.GetCase(subForm.Pid)
	if err != nil {
		log.Warn("error:%v", err)
		return nil, err
	}
	testCase := make([]dto.TestCase, len(cases))
	for i := 0; i < len(cases); i++ {
		testCase[i].Input = cases[i].Input
		testCase[i].ExpectedOutput = cases[i].Output
		testCase[i].Score = cases[i].Score
		testCase[i].Id = cases[i].Id
	}
	form.TestCase = testCase
	return form, nil
}

func GetAllInfo(c context.Context) {
	_, err := isSuperAdmin(c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	lock.RLock()
	newJsp := make([]dto.JudgeServer, len(jsp))
	copy(newJsp, jsp)
	lock.RUnlock()
	for i, j := 0, len(newJsp); i < j; i++ {
		newJsp[i].Password = ""
	}
	_, _ = c.JSON(dto.Res{
		Error: "",
		Data:  &newJsp,
	})
}

func AddJudgeServer(c context.Context) {
	var js dto.JudgeServer
	err := c.ReadJSON(&js)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	_, err = isSuperAdmin(c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	_, err = gosql.Exec(`insert into ojo.judge_server(name, address, port, weight, enabled, password)
			values(?,?,?,?,?,?)`, js.Name, js.Address, js.Port, js.Weight, js.Enabled, js.Password)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	go UpdateJSP()
	c.JSON(dto.Res{
		Error: "",
		Data:  "success",
	})
}

func UpdateJudgeServer(c context.Context) {
	var js dto.JudgeServer
	err := c.ReadJSON(&js)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	_, err = isSuperAdmin(c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	if js.Password == "" {
		_, err = gosql.Exec(`update ojo.judge_server set 
                            name=?,
                            address=?,
                            port=?,
                            weight=?,
                            enabled=?
                            where id=?`,
			js.Name, js.Address, js.Port,
			js.Weight, js.Enabled, js.Id)
		if err != nil {
			c.JSON(&dto.Res{Error: err.Error(), Data: nil})
			return
		}
	} else {
		_, err = gosql.Exec(`update ojo.judge_server set 
                            name=?,
                            address=?,
                            port=?,
                            weight=?,
                            enabled=?,
                            password=? 
                            where id=?`,
			js.Name, js.Address, js.Port,
			js.Weight, js.Enabled, js.Password, js.Id)
		if err != nil {
			c.JSON(&dto.Res{Error: err.Error(), Data: nil})
			return
		}
	}
	go UpdateJSP()
	c.JSON(dto.Res{
		Error: "",
		Data:  "success",
	})
}

func DeleteJudgeServer(c context.Context) {
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
	_, err = gosql.Exec(`delete from ojo.judge_server 
                            where id=?`,
		id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	go UpdateJSP()
	c.JSON(dto.Res{
		Error: "",
		Data:  "success",
	})
}

func SetEnabledTrue(c context.Context) {
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
	_, err = gosql.Exec("update ojo.judge_server set enabled=1 where id=?", id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	go UpdateJSP()
	c.JSON(dto.Res{
		Error: "",
		Data:  "success",
	})
}

func SetEnabledFalse(c context.Context) {
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
	_, err = gosql.Exec("update ojo.judge_server set enabled=0 where id=?", id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	go UpdateJSP()
	c.JSON(dto.Res{
		Error: "",
		Data:  "success",
	})
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
