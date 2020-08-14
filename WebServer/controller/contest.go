package ctrl

import (
	"errors"
	"github.com/afanke/OJO/WebServer/db"
	"github.com/afanke/OJO/WebServer/dto"
	jsp "github.com/afanke/OJO/WebServer/judge"
	"github.com/afanke/OJO/utils/log"
	"github.com/afanke/OJO/utils/session"
	"github.com/kataras/iris/v12"
	"net"
	"sort"
	"strings"
	"sync"
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

func (Contest) HasPassword(c iris.Context) {
	var id dto.Id
	err := c.ReadJSON(&id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	res, err := ctsdb.HasPassword(id.Id)
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

// 根据Session和cid比对用户是否具有Contest的访问权限
func (Contest) isQualified(cid int64, c iris.Context) (bool, int64, error) {
	userId, err := session.GetInt64(c, "userId")
	if err != nil {
		return false, 0, errors.New("please login")
	}
	qualified, err := ctsdb.GetQualification(userId, cid)
	return qualified, userId, err
}

var statusPageSize = 15

// 获得用户在cid对应Contest的所有提交记录
func (Contest) GetAllStatus(c iris.Context) {
	var form dto.ContestForm
	err := c.ReadJSON(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	qualified, userId, err := cts.isQualified(form.Cid, c)
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
		form.Offset = (form.Page - 1) * statusPageSize
	}
	form.Limit = statusPageSize
	data, err := ctsdb.GetAllStat(form.Cid, userId, form.Offset, form.Limit)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	err = BatchEncrypt(len(data), func(i int) *int64 {
		return &data[i].Id
	}, func(i int) *string {
		return &data[i].Eid
	})
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
	qualified, userId, err := cts.isQualified(form.Cid, c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	if !qualified {
		c.JSON(&dto.Res{Error: errors.New("you are not qualified").Error(), Data: nil})
		return
	}
	data, err := ctsdb.GetAllStatCount(form.Cid, userId)
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
	qualified, userId, err := cts.isQualified(form.Cid, c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	if !qualified {
		c.JSON(&dto.Res{Error: errors.New("you are not qualified").Error(), Data: nil})
		return
	}
	data, err := ctsdb.GetSubmission(userId, form.Pid, form.Cid)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	err = BatchEncrypt(1, func(i int) *int64 {
		return &data.Id
	}, func(i int) *string {
		return &data.Eid
	})
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: data})
}

// 根据csmid获得Contest提交记录的总体信息
func (Contest) GetStatus(c iris.Context) {
	var csmid dto.Eid
	err := c.ReadJSON(&csmid)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	id, err := DecryptId(csmid.Id)
	data, err := ctsdb.GetStatus(id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	err = BatchEncrypt(1, func(i int) *int64 {
		return &data.Id
	}, func(i int) *string {
		return &data.Eid
	})
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: data})
}

// 根据csmid获得Contest提交记录的具体各个判题点信息
func (Contest) GetStatusDetail(c iris.Context) {
	var csmid dto.Eid
	err := c.ReadJSON(&csmid)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	id, err := DecryptId(csmid.Id)
	data, err := ctsdb.GetCaseRes(id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	showOutput, err := ctsdb.GetShowOutput(id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	if !showOutput {
		for i, j := 0, len(data); i < j; i++ {
			data[i].RealOutput = ""
			data[i].ErrorOutput = ""
			data[i].SPJOutput = ""
			data[i].SPJErrorOutput = ""
		}
	}
	err = BatchEncrypt(len(data), func(i int) *int64 {
		return &data[i].Csmid
	}, func(i int) *string {
		return &data[i].Ecsmid
	})
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
	showRank, err := ctsdb.GetShowRank(form.Cid)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	if !showRank {
		c.JSON(&dto.Res{Error: errors.New("rank closed").Error(), Data: nil})
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
	showRank, err := ctsdb.GetShowRank(id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	if !showRank {
		c.JSON(&dto.Res{Error: errors.New("rank closed").Error(), Data: nil})
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
	showRank, err := ctsdb.GetShowRank(id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	if !showRank {
		c.JSON(&dto.Res{Error: errors.New("rank closed").Error(), Data: nil})
		return
	}
	detail, err := ctsdb.GetOITop10(id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: detail})
}

//-------------------------------------------------------------
// ACM Rank

type ACMRankPool map[int64]*ACMRank

var ACMPoolLock sync.Mutex

type ACMRank struct {
	FirstAC    map[int64]time.Time `json:"firstAC"`
	Rank       RankList            `json:"rank"`
	UpdateTime time.Time           `json:"updateTime"`
	data       map[int64]*ACMData
	lock       sync.RWMutex
}

type ACMRankForm struct {
	FirstAC    map[int64]time.Time `json:"firstAC"`
	Rank       RankList            `json:"rank"`
	UpdateTime time.Time           `json:"updateTime"`
}

type ACMData struct {
	Uid        int64      `json:"uid" db:"uid"`
	Total      int        `json:"total" db:"total"`
	AC         int        `json:"ac" db:"ac"`
	TotalTime  time.Time  `json:"totalTime" db:"total_time"`
	Username   string     `json:"username" db:"username"`
	Detail     DetailList `json:"detail" db:"detail"`
	detailData map[int64]*ACMDetail
}

type ACMDetail struct {
	Pid            int64     `json:"pid" db:"pid"`
	LastSubmitTime time.Time `json:"lastSubmitTime" db:"last_submit_time"`
	Total          int       `json:"total" db:"total"`
	AC             bool      `json:"ac" db:"ac"`
}

type RankList []ACMData

type DetailList []ACMDetail

func (l RankList) Len() int {
	return len(l)
}

func (l RankList) Less(i, j int) bool {
	return l[i].AC > l[j].AC || (l[i].AC == l[j].AC && l[i].TotalTime.Before(l[j].TotalTime))
}

func (l RankList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l DetailList) Len() int {
	return len(l)
}

func (l DetailList) Less(i, j int) bool {
	return l[i].LastSubmitTime.Before(l[j].LastSubmitTime)
}

func (l DetailList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

var acm ACMRankPool = map[int64]*ACMRank{}

func TryUpdateACMRank(cid int64) error {
	rank, ok := acm[cid]
	if !ok {
		ACMPoolLock.Lock()
		defer ACMPoolLock.Unlock()
		if _, ok = acm[cid]; !ok {
			acm[cid] = &ACMRank{
				FirstAC:    map[int64]time.Time{},
				data:       map[int64]*ACMData{},
				Rank:       []ACMData{},
				UpdateTime: time.Time{},
				lock:       sync.RWMutex{},
			}
		}
		rank = acm[cid]
	}
	rank.lock.RLock()
	if time.Now().Before(rank.UpdateTime.Add(time.Second * 60)) {
		rank.lock.RUnlock()
		return nil
	}
	rank.lock.RUnlock()
	rank.lock.Lock()
	defer rank.lock.Unlock()
	if time.Now().Before(rank.UpdateTime.Add(time.Second * 60)) {
		return nil
	}
	next := time.Now()
	s, err := ctsdb.GetACMSubByTime(cid, rank.UpdateTime, next)
	if err != nil {
		return err
	}
	d, err := ctsdb.GetDetail(cid)
	if err != nil {
		return err
	}
	punishTime := d.Punish
	for i, j := 0, len(s); i < j; i++ {
		uid := s[i].Uid
		pid := s[i].Pid
		user, ok := rank.data[uid]
		if !ok {
			rank.data[uid] = &ACMData{detailData: map[int64]*ACMDetail{}}
			user = rank.data[uid]
		}
		user.Uid = uid
		detail, ok := user.detailData[pid]
		if !ok {
			user.detailData[pid] = &ACMDetail{}
			detail = user.detailData[pid]
		}
		detail.Pid = pid
		if s[i].Flag == "AC" {
			detail.AC = true
			user.AC++
		}
		if s[i].Flag != "ISE" {
			detail.Total++
			user.Total++
			submitTime, err := time.Parse("2006-01-02 15:04:05", s[i].SubmitTime)
			if err != nil {
				return err
			}
			if submitTime.After(detail.LastSubmitTime) {
				detail.LastSubmitTime = submitTime
			}
			if submitTime.After(user.TotalTime) {
				user.TotalTime = submitTime
			}
		}
		firstAC := rank.FirstAC[pid]
		if detail.AC && ((detail.LastSubmitTime.Before(firstAC)) || firstAC.Equal(time.Time{})) {
			rank.FirstAC[pid] = detail.LastSubmitTime
		}
	}
	rank.Rank = []ACMData{}
	for k := range rank.data {
		rank.Rank = append(rank.Rank, *rank.data[k])
	}
	for i, j := 0, len(rank.Rank); i < j; i++ {
		ac := rank.Rank[i].AC
		total := rank.Rank[i].Total
		totalTime := rank.Rank[i].TotalTime
		rank.Rank[i].TotalTime = totalTime.Add(time.Second * time.Duration((total-ac)*punishTime))
		name, err := userdb.GetName(rank.Rank[i].Uid)
		if err != nil {
			log.Warn("error:%v", err)
		}
		rank.Rank[i].Username = name
		for k := range rank.Rank[i].detailData {
			rank.Rank[i].Detail = append(rank.Rank[i].Detail, *rank.Rank[i].detailData[k])
		}
		sort.Sort(rank.Rank[i].Detail)
	}
	sort.Sort(rank.Rank)
	rank.UpdateTime = next
	return nil
}

var ACMRankPageSize = 20

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
	showRank, err := ctsdb.GetShowRank(id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	if !showRank {
		c.JSON(&dto.Res{Error: errors.New("rank closed").Error(), Data: nil})
		return
	}
	err = TryUpdateACMRank(id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	acm[id.Id].lock.RLock()
	defer acm[id.Id].lock.RUnlock()
	length := len(acm[id.Id].Rank)
	right := 10
	if right > length {
		right = length
	}
	data := ACMRankForm{
		Rank:       acm[id.Id].Rank[0:right],
		UpdateTime: acm[id.Id].UpdateTime,
	}
	c.JSON(&dto.Res{Error: "", Data: data})
}

func (Contest) GetACMRank(c iris.Context) {
	var form dto.ContestForm
	err := c.ReadJSON(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	cid := form.Cid
	qualified, _, err := cts.isQualified(cid, c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	if !qualified {
		c.JSON(&dto.Res{Error: errors.New("you are not qualified").Error(), Data: nil})
		return
	}
	showRank, err := ctsdb.GetShowRank(cid)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	if !showRank {
		c.JSON(&dto.Res{Error: errors.New("rank closed").Error(), Data: nil})
		return
	}
	page := form.Page
	if page <= 0 {
		page = 1
	}
	err = TryUpdateACMRank(cid)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	acm[cid].lock.RLock()
	defer acm[cid].lock.RUnlock()
	length := len(acm[cid].Rank)
	left := (page - 1) * ACMRankPageSize
	right := page * ACMRankPageSize
	if length < left {
		right = 0
		left = 0
	} else if right > length {
		right = length
	}
	log.Debug("%v %v", left, right)
	data := ACMRankForm{
		FirstAC:    acm[cid].FirstAC,
		Rank:       acm[cid].Rank[left:right],
		UpdateTime: acm[cid].UpdateTime,
	}

	c.JSON(&dto.Res{Error: "", Data: data})
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
	showRank, err := ctsdb.GetShowRank(id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	if !showRank {
		c.JSON(&dto.Res{Error: errors.New("rank closed").Error(), Data: nil})
		return
	}
	err = TryUpdateACMRank(id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	acm[id.Id].lock.RLock()
	defer acm[id.Id].lock.RUnlock()
	data := len(acm[id.Id].Rank)
	c.JSON(&dto.Res{Error: "", Data: data})
}

//-------------------------------------------------------------

// -------------------------------------------------------------
// 提交代码
func (Contest) Submit(c iris.Context) {
	var form dto.SubmitForm
	err := c.ReadJSON(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	over, err := cts.isEnded(form.Cid)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	if over {
		c.JSON(&dto.Res{Error: errors.New("the contest is over").Error(), Data: nil})
		return
	}
	qualified, userId, err := cts.isQualified(form.Cid, c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	if !qualified {
		c.JSON(&dto.Res{Error: errors.New("you are not qualified").Error(), Data: nil})
		return
	}
	form.Uid = userId
	data, err := ctsdb.Submit(form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	form.Sid = data.Id
	go cts.handleSubmit(&form)
	err = BatchEncrypt(1, func(i int) *int64 {
		return &data.Id
	}, func(i int) *string {
		return &data.Eid
	})
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: data})
}

func (Contest) handleSubmit(submitForm *dto.SubmitForm) {
	form, err := jsp.PrepareForm(submitForm)
	if err != nil {
		log.Warn("error:%v", err)
		_ = ctsdb.SetISE(submitForm.Sid)
		return
	}
	form, err = jsp.SendToJudge(form)
	if err != nil {
		log.Warn("error:%v", err)
		_ = pctdb.SetISE(submitForm.Sid)
		return
	}
	err = cts.updateStatistic(form)
	if err != nil {
		log.Warn("error:%v", err)
		_ = pctdb.SetISE(submitForm.Sid)
		return
	}
	err = cts.InsertCaseRes(form)
	if err != nil {
		log.Warn("error:%v", err)
		_ = pctdb.SetISE(submitForm.Sid)
		return
	}
	err = ctsdb.UpdateFlagScoreMsg(form.Sid, form.TotalScore, form.Flag, form.ErrorMsg)
	if err != nil {
		log.Warn("error:%v", err)
		_ = pctdb.SetISE(submitForm.Sid)
		return
	}
}

func (Contest) updateStatistic(form *dto.JudgeForm) error {
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
		total--
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
	err := ctsdb.UpdateStat(form.Cid, form.Pid, total, ac, wa, ce, re, tle, mle, ole)
	return err
}

func (Contest) InsertCaseRes(form *dto.JudgeForm) error {
	for i, j := 0, len(form.TestCase); i < j; i++ {
		err := ctsdb.InsertCaseRes(form.Sid, form.Uid, &form.TestCase[i])
		if err != nil {
			log.Warn("error:%v", err)
			return err
		}
	}
	return nil
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
	if contest.Password != "" {
		contest.Password = Encrypt(contest.Password)
	}
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
	isStarted, err := cts.isStarted(id.Cid)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	if isStarted {
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
	isUnderway, err := cts.isUnderway(id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	if isUnderway {
		c.JSON(&dto.Res{Error: "can't delete contest underway", Data: nil})
		return
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
	isStarted, err := cts.isStarted(contest.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	if isStarted {
		data, err := ctsdb.GetDetail(contest.Id)
		if err != nil {
			c.JSON(&dto.Res{Error: err.Error(), Data: nil})
			return
		}
		contest.StartTime = data.StartTime
		contest.Rule = data.Rule
		contest.Punish = data.Punish
		contest.SubmitLimit = data.SubmitLimit
	}
	if contest.Password != "" {
		contest.Password = Encrypt(contest.Password)
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

func (Contest) isIPMatched(c iris.Context, id int64) error {
	s := c.RemoteAddr()
	split := strings.Split(s, ":")
	ip := net.ParseIP(split[0])
	if ip == nil {
		return errors.New("parse ip error")
	}
	limit, err := ctsdb.GetIPLimit(id)
	if err != nil {
		return err
	}
	matched := false
	for i, j := 0, len(limit); i < j; i++ {
		lip := net.ParseIP(limit[i].Address)
		if lip == nil {
			return errors.New("illegal ip limit, please contact with admin")
		}
		mask := net.CIDRMask(limit[i].Mask, 32)
		rlip := lip.Mask(mask)
		if rlip == nil {
			return errors.New("illegal ip limit, please contact with admin")
		}
		rip := ip.Mask(mask)
		if rip == nil {
			return errors.New("illegal ip limit, please contact with admin")
		}
		equal := rip.Equal(ip)
		if equal {
			matched = true
			break
		}
	}
	if !matched {
		return errors.New("ip range not allowed")
	}
	return nil
}

func (Contest) isStarted(pid int64) (bool, error) {
	ctsTime, err := ctsdb.GetTime(pid)
	if err != nil {
		return false, err
	}
	startTime, err := time.ParseInLocation("2006-01-02 15:04:05", ctsTime.StartTime, time.Local)
	if err != nil {
		return false, err
	}
	now := time.Now()
	return now.After(startTime), nil
}

func (Contest) isEnded(pid int64) (bool, error) {
	ctsTime, err := ctsdb.GetTime(pid)
	if err != nil {
		return false, err
	}
	endTime, err := time.ParseInLocation("2006-01-02 15:04:05", ctsTime.EndTime, time.Local)
	if err != nil {
		return false, err
	}
	now := time.Now()
	return now.After(endTime), nil
}

func (Contest) isUnderway(pid int64) (bool, error) {
	ctsTime, err := ctsdb.GetTime(pid)
	if err != nil {
		return false, err
	}
	startTime, err := time.ParseInLocation("2006-01-02 15:04:05", ctsTime.StartTime, time.Local)
	if err != nil {
		return false, err
	}
	endTime, err := time.ParseInLocation("2006-01-02 15:04:05", ctsTime.EndTime, time.Local)
	if err != nil {
		return false, err
	}
	now := time.Now()
	return now.After(startTime) && now.Before(endTime), nil
}
