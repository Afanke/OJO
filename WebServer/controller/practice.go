package ctrl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/afanke/OJO/WebServer/db"
	"github.com/afanke/OJO/WebServer/dto"
	jsp "github.com/afanke/OJO/WebServer/judge"
	"github.com/afanke/OJO/utils/log"
	"github.com/kataras/iris/v12"
	"io/ioutil"
	"net/http"
	"time"
)

// Practice 为show=true 的Problem
type Practice struct{}

var pctdb db.Practice
var pbdb db.Problem
var pt Practice

// 获得所有Practice
func (Practice) GetAll(c iris.Context) {
	var form dto.PracticeForm
	err := c.ReadJSON(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	res, err := pctdb.GetAll(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: res})
}

// 获得所有Problem的Tags
func (Practice) GetTags(c iris.Context) {
	tags, err := pctdb.GetAllTags()
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: tags})
}

// 获得所有Problem的Tags
func (Practice) GetCount(c iris.Context) {
	var form dto.PracticeForm
	err := c.ReadJSON(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	tags, err := pctdb.GetCount(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: tags})
}

// 获得对应Problem的具体信息
func (Practice) GetDetail(c iris.Context) {
	var ptid dto.Id
	err := c.ReadJSON(&ptid)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	detail, err := pctdb.GetDetail(ptid.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: detail})
}

// 获得当前Practice最新一次的提交状态
func (Practice) GetCurrentStatus(c iris.Context) {
	var ptid dto.Id
	err := c.ReadJSON(&ptid)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	userId, err := getUserId(c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	detail, err := pctdb.GetSubmission(userId, ptid.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: detail})
}

// 根据pcmid获得对应Practice提交的总体信息
func (Practice) GetStatus(c iris.Context) {
	var psmid dto.Id
	err := c.ReadJSON(&psmid)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	detail, err := pctdb.GetStat(psmid.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: detail})
}

// 根据pcmid获得对应Practice提交的各个判题点信息
func (Practice) GetStatusDetail(c iris.Context) {
	var psmid dto.Id
	err := c.ReadJSON(&psmid)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	data, err := pctdb.GetCaseRes(psmid.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: data})
}

// 获得当前用户的所有Practice提交的记录
func (Practice) GetAllStatus(c iris.Context) {
	var form dto.PracticeForm
	err := c.ReadJSON(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	if form.Page < 1 {
		form.Offset = 0
	} else {
		form.Offset = (form.Page - 1) * 10
	}
	form.Limit = 10
	userId, err := getUserId(c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	data, err := pctdb.GetAllStat(userId, form.Offset, form.Limit)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: data})
}

// 获得当前用户的所有Practice提交的记录之和
func (Practice) GetAllStatusCount(c iris.Context) {
	userId, err := getUserId(c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	data, err := pctdb.GetAllStatCount(userId)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: data})
}

func (Practice) GetTodayCount(c iris.Context) {
	res, err := pctdb.GetTodayCount()
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: res})
}

func (Practice) GetWeekCount(c iris.Context) {
	res, err := pctdb.GetWeekCount()
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: res})
}

func (Practice) GetMonthCount(c iris.Context) {
	res, err := pctdb.GetMonthCount()
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: res})
}

// 提交代码
func (Practice) Submit(c iris.Context) {
	var form dto.SubmitForm
	err := c.ReadJSON(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	userId, err := getUserId(c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	form.Uid = userId
	data, err := pctdb.Submit(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	form.Sid = data.Id
	go pt.handleSubmit(&form)
	c.JSON(&dto.Res{Error: "", Data: data})
}

func (Practice) handleSubmit(submitForm *dto.SubmitForm) {
	form, err := pt.prepareForms(submitForm)
	if err != nil {
		log.Warn("error:%v", err)
		_ = pctdb.SetISE(submitForm.Sid)
		return
	}
	form, err = pt.sendToJudge(form)
	if err != nil {
		log.Warn("error:%v", err)
		_ = pctdb.SetISE(submitForm.Sid)
		return
	}
	err = pt.updateStatistic(form)
	if err != nil {
		log.Warn("error:%v", err)
		_ = pctdb.SetISE(submitForm.Sid)
		return
	}
	err = pt.InsertCaseRes(form)
	if err != nil {
		log.Warn("error:%v", err)
		_ = pctdb.SetISE(submitForm.Sid)
		return
	}
	err = pctdb.UpdateFlagAndScore(form.Sid, form.TotalScore, form.Flag)
	if err != nil {
		log.Warn("error:%v", err)
		_ = pctdb.SetISE(submitForm.Sid)
		return
	}
}

func (Practice) sendToJudge(form *dto.JudgeForm) (*dto.JudgeForm, error) {
	fmt.Println(form)
	addr, err := jsp.GetAddr()
	if err != nil {
		log.Error("error:%v", err)
		return nil, err
	}
	client := &http.Client{
		Timeout: time.Duration(4 * form.MaxRealTime * form.SPJMp * form.CompMp),
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
	fmt.Println(string(body))
	err = json.Unmarshal(body, &form)
	if err != nil {
		log.Error("error:%v", err)
		return nil, err
	}
	fmt.Println(form)
	return form, err
}

func (Practice) prepareForms(subForm *dto.SubmitForm) (*dto.JudgeForm, error) {
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
		Code:        subForm.Code,
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

func (Practice) countTotalScore(forms []dto.OperationForm) int {
	var count int
	for i := 0; i < len(forms); i++ {
		count += forms[i].Score
	}
	return count
}

func (Practice) concludeFlag(forms []dto.OperationForm) string {
	var flag = false
	var res = "NULL"
	for i := 0; i < len(forms); i++ {
		if forms[i].Flag != "AC" {
			if forms[i].Flag == "ISE" {
				return "ISE"
			} else if forms[i].Flag == "CE" {
				return "CE"
			} else if res == "NULL" {
				res = forms[i].Flag
			}
		} else {
			flag = true
		}
	}
	if flag && res == "NULL" {
		return "AC"
	} else if flag {
		return "PA"
	} else {
		return res
	}
}

func (Practice) updateStatistic(form *dto.JudgeForm) error {
	var total = 1
	var ac = 0
	var wa = 0
	var ce = 0
	var re = 0
	var tle = 0
	var mle = 0
	var ole = 0
	switch form.Flag {
	case "ISE":
	case "AC":
		ac++
	case "RE":
		re++
	case "CE":
		ce++
	case "TLE":
		tle++
	case "WA":
		wa++
	case "MLE":
		mle++
	case "OLE":
		ole++
	}
	err := pctdb.UpdateStat(form.Pid, total, ac, wa, ce, mle, re, tle, ole)
	return err
}

func (Practice) InsertCaseRes(form *dto.JudgeForm) error {
	for i, j := 0, len(form.TestCase); i < j; i++ {
		err := pctdb.InsertCaseRes(form.Sid, form.Uid, &form.TestCase[i])
		if err != nil {
			log.Warn("error:%v", err)
			return err
		}
	}
}
