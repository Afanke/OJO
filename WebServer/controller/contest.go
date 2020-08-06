package ctrl

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/afanke/OJO/WebServer/db"
	"github.com/afanke/OJO/WebServer/dto"
	jsp "github.com/afanke/OJO/WebServer/judge"
	"github.com/afanke/OJO/utils/log"
	"github.com/afanke/OJO/utils/session"
	"github.com/kataras/iris/v12"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Contest struct{}

var ctsdb db.Contest
var cts Contest

// 获得所有的Contest
func (Contest) GetAllVisible(c iris.Context) {
	var form dto.ContestForm
	err := c.ReadJSON(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	res, err := ctsdb.GetAllVisible(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: res})
}

func (Contest) GetAll(c iris.Context) {
	var form dto.ContestForm
	err := c.ReadJSON(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	userId, err := isAdmin(c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	form.Cid = userId
	data, err := ctsdb.GetAll(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: data})
}

func (Contest) GetCtsProblem(c iris.Context) {
	var id dto.Id3
	err := c.ReadJSON(&id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	err = cts.isPermitted(c, id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	data, err := ctsdb.GetCtsProblem(id.Id)
	c.JSON(&dto.Res{Error: "", Data: data})
}

func (Contest) SetVisibleTrue(c iris.Context) {
	var id dto.Id
	err := c.ReadJSON(&id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	err = cts.isPermitted(c, id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	err = ctsdb.SetVisibleTrue(id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: "update successfully"})
}

func (Contest) SetVisibleFalse(c iris.Context) {
	var id dto.Id
	err := c.ReadJSON(&id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	err = cts.isPermitted(c, id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	err = ctsdb.SetVisibleFalse(id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: "update successfully"})
}

// 获得所有的Contest的数量
func (Contest) GetCount(c iris.Context) {
	var form dto.ContestForm
	err := c.ReadJSON(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	userId, err := isAdmin(c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	form.Cid = userId
	res, err := ctsdb.GetCount(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: res})
}

func (Contest) GetVisibleCount(c iris.Context) {
	var form dto.ContestForm
	err := c.ReadJSON(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	res, err := ctsdb.GetVisibleCount(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: res})
}

// 获得对应id的Contest的详细信息
func (Contest) GetVisibleDetail(c iris.Context) {
	var id dto.Id
	err := c.ReadJSON(&id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	res, err := ctsdb.GetVisibleDetail(id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: res})
}

// 获得对应id的Contest的时间信息
func (Contest) GetTime(c iris.Context) {
	var id dto.Id
	err := c.ReadJSON(&id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	res, err := ctsdb.GetTime(id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: res})
}

// 获得用户对Contest的访问权限
func (Contest) GetQualification(c iris.Context) {
	var id dto.Id
	err := c.ReadJSON(&id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	qualified, _, err := cts.isQualified(id.Id, c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: qualified})
}

// 根据password让用户尝试获得Contest权限
func (Contest) Qualify(c iris.Context) {
	var form dto.ContestQualifyForm
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
	startTime, err := ctsdb.GetStartTime(form.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	if time.Now().Before(startTime) {
		c.JSON(&dto.Res{Error: errors.New("the contest isn't started").Error(), Data: nil})
		return
	}
	password, err := ctsdb.GetPassword(form.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	if form.Password != password {
		c.JSON(&dto.Res{Error: errors.New("the password is not correct").Error(), Data: nil})
		return
	}
	err = ctsdb.AddQualification(userId, form.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: "success"})
}

// 判断Contest是否已经结束
func (Contest) isOver(cid int64) (bool, error) {
	res, err := ctsdb.GetTime(cid)
	if err != nil {
		log.Warn("error:%v\n", err)
		return true, err
	}
	end, err := time.Parse("2006-01-02 15:04:05", res.EndTime)
	if err != nil {
		log.Warn("error:%v\n", err)
		return true, err
	}
	now, err := time.Parse("2006-01-02 15:04:05", res.Now)
	if err != nil {
		log.Warn("error:%v\n", err)
		return true, err
	}
	return now.After(end), nil
}

// 根据Session和cid比对用户是否具有Contest的访问权限
func (Contest) isQualified(cid int64, c iris.Context) (bool, *dto.UserToken, error) {
	s, err := session.GetSession(c)
	if err != nil {
		log.Warn("error:%v\n", err)
		return false, nil, err
	}
	user, ok := s.Get("user").(dto.UserToken)
	if !ok {
		return false, nil, errors.New("please login")
	}
	qualified, err := ctsdb.GetQualification(user.Id, cid)
	return qualified, &user, err
}

// 获得用户在cid对应Contest的所有提交记录
func (Contest) GetAllStatus(c iris.Context) {
	var form dto.ContestForm
	err := c.ReadJSON(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	qualified, user, err := cts.isQualified(form.Cid, c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	if !qualified {
		c.JSON(&dto.Res{Error: errors.New("you are not qualified").Error(), Data: nil})
		return
	}
	if form.Page < 1 {
		form.Offset = 0
	} else {
		form.Offset = (form.Page - 1) * 10
	}
	form.Limit = 10
	data, err := ctsdb.GetAllStat(form.Cid, user.Id, form.Offset, form.Limit)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: data})
}

// 获得用户在cid对应Contest的所有提交记录数目
func (Contest) GetAllStatusCount(c iris.Context) {
	var form dto.ContestForm
	err := c.ReadJSON(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	qualified, user, err := cts.isQualified(form.Cid, c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	if !qualified {
		c.JSON(&dto.Res{Error: errors.New("you are not qualified").Error(), Data: nil})
		return
	}
	data, err := ctsdb.GetAllStatCount(form.Cid, user.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: data})
}

// 获得Contest下的所有Problem
func (Contest) GetAllProblem(c iris.Context) {
	var id dto.Id
	err := c.ReadJSON(&id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	qualified, _, err := cts.isQualified(id.Id, c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	if !qualified {
		c.JSON(&dto.Res{Error: errors.New("you are not qualified").Error(), Data: nil})
		return
	}
	data, err := ctsdb.GetAllProblem(id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: data})
}

// 获得Contest下的所有Problem的名字
func (Contest) GetAllProblemName(c iris.Context) {
	var id dto.Id
	err := c.ReadJSON(&id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	qualified, _, err := cts.isQualified(id.Id, c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	if !qualified {
		c.JSON(&dto.Res{Error: errors.New("you are not qualified").Error(), Data: nil})
		return
	}
	data, err := ctsdb.GetAllProblemName(id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: data})
}

// 根据pid和cid得到对应Problem的具体信息
func (Contest) GetProblemDetail(c iris.Context) {
	var id dto.Id2
	err := c.ReadJSON(&id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	qualified, _, err := cts.isQualified(id.Cid, c)
	if !qualified {
		c.JSON(&dto.Res{Error: errors.New("you are not qualified").Error(), Data: nil})
		return
	}
	matched, err := ctsdb.IsMatched(id.Cid, id.Pid)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	if !matched {
		c.JSON(&dto.Res{Error: errors.New("the contest and the problem is not matched").Error(), Data: nil})
		return
	}
	data, err := ctsdb.GetProblemDetail(id.Cid, id.Pid)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: data})
}

// 根据pid和cid得到对应Problem的最新一次的提交记录
func (Contest) GetCurrentStatus(c iris.Context) {
	var form dto.SubmitForm
	err := c.ReadJSON(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	qualified, user, err := cts.isQualified(form.Cid, c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	if !qualified {
		c.JSON(&dto.Res{Error: errors.New("you are not qualified").Error(), Data: nil})
		return
	}
	detail, err := ctsdb.GetSubmission(user.Id, form.Pid, form.Cid)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: detail})
}

// 根据csmid获得Contest提交记录的总体信息
func (Contest) GetStatus(c iris.Context) {
	var csmid dto.Id
	err := c.ReadJSON(&csmid)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	detail, err := ctsdb.GetStat(csmid.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: detail})
}

// 根据csmid获得Contest提交记录的具体各个判题点信息
func (Contest) GetStatusDetail(c iris.Context) {
	var csmid dto.Id
	err := c.ReadJSON(&csmid)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	data, err := ctsdb.GetCaseRes(csmid.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: data})
}

// 根据cid获得对应Contest的OI排名
func (Contest) GetOIRank(c iris.Context) {
	var form dto.ContestForm
	err := c.ReadJSON(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	qualified, _, err := cts.isQualified(form.Cid, c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	if !qualified {
		c.JSON(&dto.Res{Error: errors.New("you are not qualified").Error(), Data: nil})
		return
	}
	detail, err := ctsdb.GetOIRank(form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: detail})
}

// 根据cid获得对应Contest的OI排名总人数
func (Contest) GetOIRankCount(c iris.Context) {
	var id dto.Id
	err := c.ReadJSON(&id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	qualified, _, err := cts.isQualified(id.Id, c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	if !qualified {
		c.JSON(&dto.Res{Error: errors.New("you are not qualified").Error(), Data: nil})
		return
	}
	detail, err := ctsdb.GetOIRankCount(id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: detail})
}

// 根据cid获得对应Contest的OI排名前十位
func (Contest) GetOITop10(c iris.Context) {
	var id dto.Id
	err := c.ReadJSON(&id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	qualified, _, err := cts.isQualified(id.Id, c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	if !qualified {
		c.JSON(&dto.Res{Error: errors.New("you are not qualified").Error(), Data: nil})
		return
	}
	detail, err := ctsdb.GetOITop10(id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: detail})
}

func (Contest) GetACMTop10(c iris.Context) {
	var id dto.Id
	err := c.ReadJSON(&id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	qualified, _, err := cts.isQualified(id.Id, c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	if !qualified {
		c.JSON(&dto.Res{Error: errors.New("you are not qualified").Error(), Data: nil})
		return
	}
	detail, err := ctsdb.GetACMTop10(id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: detail})
}

func (Contest) GetACMRank(c iris.Context) {
	var form dto.ContestForm
	err := c.ReadJSON(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	qualified, _, err := cts.isQualified(form.Cid, c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	if !qualified {
		c.JSON(&dto.Res{Error: errors.New("you are not qualified").Error(), Data: nil})
		return
	}
	detail, err := ctsdb.GetACMRank(form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: detail})
}

func (Contest) GetACMRankCount(c iris.Context) {
	var id dto.Id
	err := c.ReadJSON(&id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	qualified, _, err := cts.isQualified(id.Id, c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	if !qualified {
		c.JSON(&dto.Res{Error: errors.New("you are not qualified").Error(), Data: nil})
		return
	}
	detail, err := ctsdb.GetACMRankCount(id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: detail})
}

// -------------------------------------------------------------
// 提交代码
func (Contest) Submit(c iris.Context) {
	var form dto.SubmitForm
	err := c.ReadJSON(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	over, err := cts.isOver(form.Cid)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	if over {
		c.JSON(&dto.Res{Error: errors.New("the contest is over").Error(), Data: nil})
		return
	}
	qualified, user, err := cts.isQualified(form.Cid, c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	if !qualified {
		c.JSON(&dto.Res{Error: errors.New("you are not qualified").Error(), Data: nil})
		return
	}
	form.Uid = user.Id
	data, err := ctsdb.Submit(form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	form.Sid = data.Id
	go cts.handleSubmit(&form)
	c.JSON(&dto.Res{Error: "", Data: data})
}

func (Contest) handleSubmit(form *dto.SubmitForm) {
	contest, err := ctsdb.GetVisibleDetail(form.Cid)
	if err != nil {
		log.Warn("error:%v", err)
		_ = ctsdb.SetISE(form.Sid)
		return
	}
	if contest.Rule == "OI" {
		handleOI(contest, form)
	} else {
		handleACM(contest, form)
	}
}

func handleOI(contest *dto.ContestDetail, form *dto.SubmitForm) {
	forms, err := cts.prepareForms(form)
	if err != nil {
		log.Warn("error:%v", err)
		_ = ctsdb.SetISE(form.Sid)
		return
	}
	forms, err = cts.sendToJudge(forms)
	if err != nil {
		log.Warn("error:%v", err)
		_ = ctsdb.SetISE(form.Sid)
		return
	}
	err = cts.updateStatistic(form.Cid, form.Pid, form.Sid, form.Uid, forms)
	if err != nil {
		log.Warn("error:%v", err)
		_ = ctsdb.SetISE(form.Sid)
		return
	}
	flag := cts.concludeFlag(forms)
	score := cts.countTotalScore(forms)
	err = ctsdb.UpdateFlagAndScore(form.Sid, score, flag)
	if err != nil {
		log.Warn("error:%v", err)
		_ = ctsdb.SetISE(form.Sid)
		return
	}
}

func handleACM(contest *dto.ContestDetail, form *dto.SubmitForm) {
	forms, err := cts.prepareForms(form)
	if err != nil {
		log.Warn("error:%v", err)
		_ = ctsdb.SetISE(form.Sid)
		return
	}
	forms, err = cts.sendToJudge(forms)
	if err != nil {
		log.Warn("error:%v", err)
		_ = ctsdb.SetISE(form.Sid)
		return
	}
	err = cts.updateStatistic(form.Cid, form.Pid, form.Sid, form.Uid, forms)
	if err != nil {
		log.Warn("error:%v", err)
		_ = ctsdb.SetISE(form.Sid)
		return
	}
	flag := cts.concludeFlag(forms)
	stat, err := ctsdb.GetStat(form.Sid)
	if err != nil {
		log.Warn("error:%v", err)
		_ = ctsdb.SetISE(form.Sid)
		return
	}
	startTime, err := time.Parse("2006-01-02 15:04:05", contest.StartTime)
	if err != nil {
		log.Warn("error:%v", err)
		_ = ctsdb.SetISE(form.Sid)
		return
	}
	subTime, err := time.Parse("2006-01-02 15:04:05", stat.SubmitTime)
	if err != nil {
		log.Warn("error:%v", err)
		_ = ctsdb.SetISE(form.Sid)
		return
	}
	duration := int(subTime.Unix() - startTime.Unix())
	yes, err := ctsdb.HasACMOverAll(form)
	if err != nil {
		log.Warn("error:%v", err)
		_ = ctsdb.SetISE(form.Sid)
		return
	}
	if yes {
		wrong, err := ctsdb.GetACMWrong(form)
		if err != nil {
			log.Warn("error:%v", err)
			_ = ctsdb.SetISE(form.Sid)
			return
		}
		du := duration
		if flag != "AC" {
			du += (wrong + 1) * contest.Punish
		} else {
			du += wrong * contest.Punish
		}
		err = ctsdb.UpdateACMOverAll(form, du, flag == "AC")
		if err != nil {
			log.Warn("error:%v", err)
			_ = ctsdb.SetISE(form.Sid)
			return
		}
	} else {
		du := duration
		if flag != "AC" {
			du += contest.Punish
		}
		err = ctsdb.InsertACMOverAll(form, du, flag == "AC")
		if err != nil {
			log.Warn("error:%v", err)
			_ = ctsdb.SetISE(form.Sid)
			return
		}
	}
	yes, err = ctsdb.HasACMDetail(form)
	if err != nil {
		log.Warn("error:%v", err)
		_ = ctsdb.SetISE(form.Sid)
		return
	}
	first, err := ctsdb.HasACMFirstDetail(form)
	if yes {
		err = ctsdb.UpdateACMDetail(form, duration, flag == "AC", first && flag == "AC")
		if err != nil {
			log.Warn("error:%v", err)
			_ = ctsdb.SetISE(form.Sid)
			return
		}
	} else {
		if err != nil {
			log.Warn("error:%v", err)
			_ = ctsdb.SetISE(form.Sid)
			return
		}
		err = ctsdb.InsertACMDetail(form, duration, flag == "AC", first && flag == "AC")
		if err != nil {
			log.Warn("error:%v", err)
			_ = ctsdb.SetISE(form.Sid)
			return
		}
	}
	score := 0
	if flag == "AC" {
		score = cts.countTotalScore(forms)
	}
	err = ctsdb.UpdateFlagAndScore(form.Sid, score, flag)
	if err != nil {
		log.Warn("error:%v", err)
		_ = ctsdb.SetISE(form.Sid)
		return
	}
}

func (Contest) sendToJudge(forms []dto.OperationForm) ([]dto.OperationForm, error) {
	addr, err := jsp.GetAddr()
	if err != nil {
		log.Error("error:%v", err)
		return nil, err
	}
	client := &http.Client{
		Timeout: 1 * time.Second,
	}
	buff, err := json.Marshal(forms)
	if err != nil {
		log.Error("error:%v", err)
		return nil, err
	}
	res, err := client.Post("http://"+addr+"/"+forms[0].Language, "application/json", bytes.NewBuffer(buff))
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
	err = json.Unmarshal(body, &forms)
	if err != nil {
		log.Error("error:%v", err)
		return nil, err
	}
	return forms, err
}

func (Contest) prepareForms(subForm *dto.SubmitForm) ([]dto.OperationForm, error) {
	cases, err := pbdb.GetCase(subForm.Pid)
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
		forms[i].Code = subForm.Code
	}
	return forms, nil
}

func (Contest) countTotalScore(forms []dto.OperationForm) int {
	var count int
	for i := 0; i < len(forms); i++ {
		count += forms[i].Score
	}
	return count
}

func (Contest) concludeFlag(forms []dto.OperationForm) string {
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

func (Contest) updateStatistic(cid, pid, csmid, uid int64, forms []dto.OperationForm) error {
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
		err := ctsdb.InsertCaseRes(csmid, uid, forms[i])
		if err != nil {
			log.Warn("error:%v", err)
			return err
		}
	}
	err := ctsdb.UpdateStat(cid, pid, total, ac, wa, ce, mle, re, tle, ole)
	return err
}

// -------------------------------------------------------------

func (Contest) AddContest(c iris.Context) {
	var contest dto.Contest
	err := c.ReadJSON(&contest)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	log.Debug("%v", contest)
	userId, err := isAdmin(c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	contest.Cid = userId
	err = ctsdb.InsertContest(&contest)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: "add contest successfully"})
}

func (Contest) AddProblem(c iris.Context) {
	var id dto.Id4
	err := c.ReadJSON(&id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	err = cts.isPermitted(c, id.Cid)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	err = pb.isPermitted(c, id.Pid)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	detail, err := ctsdb.GetDetail(id.Cid)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	endTime, err := time.Parse("2006-01-02 15:04:05", detail.EndTime)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	before := endTime.Before(time.Now())
	if before {
		c.JSON(&dto.Res{Error: "the contest is ended", Data: nil})
		return
	}
	err = ctsdb.InsertCtsPb(id.Cid, id.Pid)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: "add problem successfully"})
}

func (Contest) DeleteProblem(c iris.Context) {
	var id dto.Id4
	err := c.ReadJSON(&id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	err = cts.isPermitted(c, id.Cid)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	err = pb.isPermitted(c, id.Pid)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	detail, err := ctsdb.GetDetail(id.Cid)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	startTime, err := time.Parse("2006-01-02 15:04:05", detail.StartTime)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	before := startTime.Before(time.Now())
	if before {
		c.JSON(&dto.Res{Error: "can't delete problem once the contest begun", Data: nil})
		return
	}
	err = ctsdb.DeleteCtsPb(id.Cid, id.Pid)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: "delete problem successfully"})
}

func (Contest) DeleteContest(c iris.Context) {
	var id dto.Id
	err := c.ReadJSON(&id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	err = cts.isPermitted(c, id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	detail, err := ctsdb.GetDetail(id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	startTime, err := time.ParseInLocation("2006-01-02 15:04:05", detail.StartTime, time.Local)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	endTime, err := time.ParseInLocation("2006-01-02 15:04:05", detail.EndTime, time.Local)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	now := time.Now()
	before := startTime.Before(now)
	if before {
		after := endTime.After(time.Now())
		log.Debug("after:%v", after)
		if after {
			c.JSON(&dto.Res{Error: "can't delete contest underway", Data: nil})
			return
		}
	}
	err = ctsdb.DeleteContest(id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: "delete contest successfully"})
}

func (Contest) TryEdit(c iris.Context) {
	var id dto.Id
	err := c.ReadJSON(&id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	err = cts.isPermitted(c, id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: "ok"})
}

func (Contest) GetDetail(c iris.Context) {
	var id dto.Id
	err := c.ReadJSON(&id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	err = cts.isPermitted(c, id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	data, err := ctsdb.GetDetail(id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: data})
}

func (Contest) UpdateContest(c iris.Context) {
	var contest dto.Contest
	err := c.ReadJSON(&contest)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	log.Debug("%#v", contest)
	err = cts.isPermitted(c, contest.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	err = ctsdb.UpdateContest(&contest)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: "save successfully"})
}

func (Contest) GetTodayCount(c iris.Context) {
	res, err := ctsdb.GetTodayCount()
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: res})
}

func (Contest) GetWeekCount(c iris.Context) {
	res, err := ctsdb.GetWeekCount()
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: res})
}

func (Contest) GetMonthCount(c iris.Context) {
	res, err := ctsdb.GetMonthCount()
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: res})
}

func (Contest) GetRecentCount(c iris.Context) {
	res, err := ctsdb.GetRecentCount()
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: res})
}

func (Contest) isCreator(c iris.Context, id int64) error {
	i, err := session.GetInt64(c, "userId")
	if err != nil {
		return err
	}
	creatorId, err := ctsdb.GetCreatorId(id)
	if err != nil {
		return err
	}
	if i != creatorId {
		return errors.New("not allowed")
	}
	return nil
}

// to see whether he is super admin or the creator of the contest
func (Contest) isPermitted(c iris.Context, id int64) error {
	_, err := isSuperAdmin(c)
	if err != nil {
		err := cts.isCreator(c, id)
		if err != nil {
			return err
		}
	}
	return nil
}
