package ctrl

import (
	"encoding/json"
	"errors"
	"github.com/afanke/OJO/WebServer/config"
	"github.com/afanke/OJO/WebServer/db"
	"github.com/afanke/OJO/WebServer/dto"
	"github.com/afanke/OJO/utils/log"
	"github.com/afanke/OJO/utils/session"
	"github.com/afanke/OJO/utils/tcp"
	"github.com/kataras/iris"
	"strings"
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
	if form.Page < 1 {
		form.Offset = 0
	} else {
		form.Offset = (form.Page - 1) * 5
	}
	form.Limit = 5
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
	detail, err := pctdb.GetDetail(int64(ptid.Id))
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
	s, err := session.GetSession(c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	user, ok := s.Get("user").(dto.User)
	if !ok {
		c.JSON(&dto.Res{Error: errors.New("please login").Error(), Data: nil})
		return
	}
	detail, err := pctdb.GetSubmission(user.Id, ptid.Id)
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
	//s, err := session.GetSession(c)
	//if err != nil {c.JSON(&dto.Res{Error:err.Error(),Data:nil});return}
	//user,ok:= s.Get("user").(dto.User)
	//if !ok{user=dto.User{Id:1,}}
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
	s, err := session.GetSession(c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	user, ok := s.Get("user").(dto.User)
	if !ok {
		c.JSON(&dto.Res{Error: errors.New("please login").Error(), Data: nil})
		return
	}
	data, err := pctdb.GetAllStat(user.Id, form.Offset, form.Limit)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: data})
}

// 获得当前用户的所有Practice提交的记录之和
func (Practice) GetAllStatusCount(c iris.Context) {
	s, err := session.GetSession(c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	res := s.Get("user")
	var user dto.User
	user, ok := res.(dto.User)
	if !ok {
		c.JSON(&dto.Res{Error: errors.New("please login").Error(), Data: nil})
		return
	}
	data, err := pctdb.GetAllStatCount(user.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: data})
}

// 提交代码
func (Practice) Submit(c iris.Context) {
	var form dto.SubmitForm
	err := c.ReadJSON(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	s, err := session.GetSession(c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	res := s.Get("user")
	var user dto.User
	user, ok := res.(dto.User)
	if !ok {
		c.JSON(&dto.Res{Error: errors.New("please login").Error(), Data: nil})
		return
	}
	form.Uid = user.Id
	data, err := pctdb.Submit(form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	form.Sid = data.Id
	go pt.handleSubmit(form)
	c.JSON(&dto.Res{Error: "", Data: data})
}

func (Practice) handleSubmit(form dto.SubmitForm) {
	forms, err := pt.prepareForms(&form)
	if err != nil {
		log.Warn("error:%v", err)
		_ = pctdb.SetISE(form.Sid)
		return
	}
	forms, err = pt.sendToJudge(forms)
	if err != nil {
		log.Warn("error:%v", err)
		_ = pctdb.SetISE(form.Sid)
		return
	}
	err = pt.updateStatistic(form.Pid, form.Sid, form.Uid, forms)
	if err != nil {
		log.Warn("error:%v", err)
		_ = pctdb.SetISE(form.Sid)
		return
	}
	flag := pt.concludeFlag(forms)
	score := pt.countTotalScore(forms)
	err = pctdb.UpdateFlagAndScore(form.Sid, score, flag)
	if err != nil {
		log.Warn("error:%v", err)
		_ = pctdb.SetISE(form.Sid)
		return
	}
}

func (Practice) sendToJudge(forms []dto.OperationForm) ([]dto.OperationForm, error) {
	conn, err := tcp.Dial(config.Config.JudgeServer)
	if err != nil {
		log.Warn("error:%v", err)
		return nil, err
	}
	bytes, err := json.Marshal(&forms)
	_, err = conn.Send(bytes)
	if err != nil {
		log.Warn("error:%v", err)
		return nil, err
	}
	_, recv, err := conn.Recv()
	if err != nil {
		log.Warn("error:%v", err)
		return nil, err
	}
	err = json.Unmarshal(recv, &forms)
	return forms, err
}

func (Practice) prepareForms(subForm *dto.SubmitForm) ([]dto.OperationForm, error) {
	cases, err := pbdb.GetPbCase(subForm.Pid)
	if err != nil {
		log.Warn("error:%v", err)
		return nil, err
	}
	problem, err := pbdb.GetProblem(subForm.Pid)
	if err != nil {
		log.Warn("error:%v", err)
		return nil, err
	}
	forms := make([]dto.OperationForm, len(cases))
	for i := 0; i < len(cases); i++ {
		forms[i].Input = cases[i].Input
		forms[i].ExpectOutput = strings.ReplaceAll(cases[i].Output, "\r\n", "\n")
		forms[i].Score = cases[i].Score
		forms[i].PcId = cases[i].Id
		forms[i].Language = subForm.Language
		forms[i].Code = subForm.Code
		forms[i].MaxCpuTime = problem.CpuTimeLimit
		forms[i].MaxMemory = problem.MemoryLimit
		forms[i].MaxRealTime = problem.RealTimeLimit
	}
	return forms, nil
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

func (Practice) updateStatistic(pbid int64, psmid, uid int, forms []dto.OperationForm) error {
	var total = 0
	var ac = 0
	var wa = 0
	var ce = 0
	var re = 0
	var tle = 0
	var mle = 0
	var ole = 0
	for i := 0; i < len(forms); i++ {
		switch forms[i].Flag {
		case "ISE":
		case "AC":
			total++
			ac++
		case "RE":
			total++
			re++
		case "CE":
			total++
			ce++
		case "TLE":
			total++
			tle++
		case "WA":
			total++
			wa++
		case "MLE":
			total++
			mle++
		case "OLE":
			total++
			ole++
		}
		err := pctdb.InsertCaseRes(psmid, uid, forms[i])
		if err != nil {
			log.Warn("error:%v", err)
			return err
		}
	}
	err := pctdb.UpdateStat(pbid, total, ac, wa, ce, mle, re, tle, ole)
	return err
}
