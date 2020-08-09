package db

import (
	"fmt"
	"github.com/afanke/OJO/WebServer/dto"
	"github.com/afanke/OJO/utils/log"
	"github.com/ilibs/gosql/v2"
	"time"
)

type Contest struct{}

var cts Contest

var ContestPageSize = 10

func (Contest) GetAll(form *dto.ContestForm) ([]dto.Contest, error) {
	if form.Page < 1 {
		form.Page = 1
	}
	form.Page -= 1
	form.Limit = ContestPageSize
	form.Offset = form.Page * ContestPageSize
	s := `select id, title, rule, start_time, end_time,create_time,last_update_time,cid,rule,visible,punish,submit_limit  from ojo.contest `
	s += `where 1=1 `
	if form.Keywords != "" {
		s += "and title like concat('%',:keywords,'%') "
	}
	if form.Rule != "" {
		s += "and rule=:rule "
	}
	switch form.Status {
	case 1:
		s += ` and now()<start_time `
	case 2:
		s += " and start_time<now() and now()<end_time "
	case 3:
		s += " and end_time<now() "
	}
	if form.Mine {
		s += " and cid=:cid "
	}
	s += " order by create_time desc limit :offset, :limit"
	rows, err := gosql.Sqlx().NamedQuery(s, &form)
	if err != nil {
		log.Warn("error:%v", err)
		return nil, nil
	}
	t := time.Now().Format("2006-01-02 15:04:05")
	var rest = make([]dto.Contest, 0, form.Limit)
	for rows.Next() {
		var res dto.Contest
		err := rows.StructScan(&res)
		if err != nil {
			log.Warn("error:%v", err)
			return nil, nil
		}
		res.Now = t
		rest = append(rest, res)
	}
	err = cts.SelectCreatorName(len(rest), func(i int) (target int64) {
		return rest[i].Cid
	}, func(i int, res string) {
		rest[i].CreatorName = res
	})
	return rest, err
}

func (Contest) GetAllVisible(form *dto.ContestForm) ([]dto.ContestBrief, error) {
	if form.Page < 1 {
		form.Page = 1
	}
	form.Page -= 1
	form.Limit = ContestPageSize
	form.Offset = form.Page * ContestPageSize
	sql := `select id, title, rule, start_time, end_time from ojo.contest `
	sql += `where 1=1 `
	if form.Keywords != "" {
		sql += "and title like concat('%',:keywords,'%') "
	}
	if form.Rule != "" {
		sql += "and rule=:rule "
	}
	switch form.Status {
	case 1:
		sql += ` and now()<start_time `
	case 2:
		sql += " and start_time<now() and now()<end_time "
	case 3:
		sql += " and end_time<now() "
	}
	sql += " and visible=1 order by create_time desc limit :offset, :limit"
	rows, err := gosql.Sqlx().NamedQuery(sql, &form)
	if err != nil {
		log.Warn("error:%v", err)
		return []dto.ContestBrief{}, nil
	}
	var rest = make([]dto.ContestBrief, 0, form.Limit)
	t := time.Now().Format("2006-01-02 15:04:05")
	for rows.Next() {
		var res dto.ContestBrief
		err := rows.StructScan(&res)
		if err != nil {
			log.Warn("error:%v", err)
			return []dto.ContestBrief{}, nil
		}
		res.Now = t
		rest = append(rest, res)
	}
	return rest, err
}

func (Contest) GetCount(form *dto.ContestForm) (int, error) {
	sql := `select count(*) from ojo.contest `
	sql += `where 1=1 `
	if form.Keywords != "" {
		sql += "and title like concat('%',:keywords,'%') "
	}
	if form.Rule != "" {
		sql += "and rule=:rule "
	}
	switch form.Status {
	case 1:
		sql += ` and now()<start_time `
	case 2:
		sql += " and start_time<now() and now()<end_time "
	case 3:
		sql += " and end_time<now() "
	}
	if form.Mine {
		sql += " and cid=:cid "
	}
	var count int
	rows, err := gosql.Sqlx().NamedQuery(sql, &form)
	if err != nil {
		log.Warn("error:%v", err)
		return 0, err
	}
	_ = rows.Next()
	err = rows.Scan(&count)
	return count, err
}

func (Contest) GetVisibleCount(form *dto.ContestForm) (int, error) {
	sql := `select count(*) from ojo.contest `
	sql += `where 1=1 `
	if form.Keywords != "" {
		sql += "and title like concat('%',:keywords,'%') "
	}
	if form.Rule != "" {
		sql += "and rule=:rule "
	}
	switch form.Status {
	case 1:
		sql += ` and now()<start_time `
	case 2:
		sql += " and start_time<now() and now()<end_time "
	case 3:
		sql += " and end_time<now() "
	}
	sql += " and visible=1 "
	var count int
	rows, err := gosql.Sqlx().NamedQuery(sql, &form)
	if err != nil {
		log.Warn("error:%v", err)
		return 0, err
	}
	_ = rows.Next()
	err = rows.Scan(&count)
	return count, err
}

func (Contest) GetVisibleDetail(id int64) (*dto.ContestDetail, error) {
	var data dto.ContestDetail
	err := gosql.Get(&data, "select id, title, description, rule, start_time, end_time, cid, punish from contest where id=?", id)
	if err != nil {
		log.Warn("error:%v\n", err)
		return nil, err
	}
	name, err := pb.GetCreatorName(data.Cid)
	if err != nil {
		log.Warn("error:%v\n", err)
		return nil, err
	}
	data.CreatorName = name
	data.Now = time.Now().Format("2006-01-02 15:04:05")
	return &data, nil
}

func (Contest) HasPassword(id int64) (bool, error) {
	var res bool
	err := gosql.Get(&res, "select Length(password)>0 from contest where id=? limit 1;", id)
	if err != nil {
		log.Warn("error:%v\n", err)
		return false, err
	}
	return res, nil
}

func (Contest) GetQualification(uid, cid int64) (bool, error) {
	var ok int
	err := gosql.Get(&ok, "select count(*) from ojo.contest_user where uid=? and cid=? limit 1", uid, cid)
	return ok >= 1, err
}

func (Contest) AddQualification(uid, cid int64) error {
	_, err := gosql.Exec("insert into ojo.contest_user(cid, uid) values(?,?)", cid, uid)
	return err
}

func (Contest) GetPassword(cid int64) (string, error) {
	var res string
	err := gosql.Get(&res, "select password from ojo.contest where id=?", cid)
	return res, err
}

func (Contest) GetStartTime(cid int64) (time.Time, error) {
	var res string
	err := gosql.Get(&res, "select start_time from ojo.contest where id=?", cid)
	if err != nil {
		log.Warn("error:%v\n", err)
		return time.Time{}, err
	}
	data, err := time.ParseInLocation("2006-01-02 15:04:05", res, time.Local)
	if err != nil {
		log.Warn("error:%v\n", err)
		return time.Time{}, err
	}
	return data, err
}

func (Contest) GetCtsProblem(cid int64) ([]dto.ProblemBrief, error) {
	var data []dto.ProblemBrief
	s := `select p.id,p.ref,p.cid,p.title, p.difficulty,p.create_time,p.last_update_time,p.visible
			from ojo.contest_problem cp,ojo.problem p  where cp.pid=p.id and cp.cid=?`
	err := gosql.Select(&data, s, cid)
	if err != nil {
		log.Warn("error:%v", err)
		return nil, err
	}
	err = pb.SelectCreatorName(len(data), func(i int) int64 {
		return data[i].Cid
	}, func(i int, res string) {
		data[i].CreatorName = res
	})
	if err != nil {
		log.Warn("%v", err)
		return nil, err
	}
	return data, nil
}

func (Contest) GetAllProblem(cid int64) ([]dto.CtsPbBrief, error) {
	var sql = `select p.id,p.title,p.ref from contest_problem cp,problem p where cp.pid=p.id and cp.cid=?`
	var data []dto.CtsPbBrief
	err := gosql.Select(&data, sql, cid)
	if err != nil {
		log.Warn("error:%v\n", err)
		return nil, err
	}
	for i, l := 0, len(data); i < l; i++ {
		stat, err := cts.GetStatistic(cid, data[i].Id)
		if err != nil {
			log.Warn("error:%v\n", err)
			return nil, err
		}
		data[i].Statistic = stat
	}
	return data, nil
}

func (Contest) GetAllProblemName(cid int64) ([]dto.CtsPbBrief, error) {
	var sql = `select p.id,p.title,p.ref from contest_problem cp,problem p where cp.pid=p.id and cp.cid=? order by p.ref`
	var data []dto.CtsPbBrief
	err := gosql.Select(&data, sql, cid)
	return data, err
}

func (Contest) GetStatistic(cid, pid int64) (*dto.ContestStatistic, error) {
	var stat dto.ContestStatistic
	err := gosql.Get(&stat, "select * from ojo.contest_statistic where cid=? and pid=? limit 1", cid, pid)
	return &stat, err
}

func (Contest) GetProblemDetail(cid, pid int64) (*dto.ContestProblem, error) {
	var detail dto.ContestProblem
	err := gosql.Get(&detail, `select * from ojo.problem p where p.id=? and p.visible=1 limit 1`, pid)
	if err != nil {
		log.Warn("error:%v", err)
		return nil, err
	}
	statistic, err := cts.GetStatistic(cid, pid)
	if err != nil {
		log.Warn("error:%v", err)
		return nil, err
	}
	tags, err := tag.GetBriefByPid(pid)
	if err != nil {
		log.Warn("error:%v", err)
		return nil, err
	}
	languages, err := pb.GetLanguage(pid)
	if err != nil {
		log.Warn("error:%v", err)
		return nil, err
	}
	samples, err := pb.GetSample(pid)
	if err != nil {
		log.Warn("error:%v", err)
		return nil, err
	}
	name, err := pb.GetCreatorName(detail.Cid)
	if err != nil {
		log.Warn("error:%v", err)
		return nil, err
	}
	template, err := pb.GetTemplate(pid)
	if err != nil {
		log.Warn("error:%v", err)
		return nil, err
	}
	limit, err := pb.GetLimit(pid)
	if err != nil {
		log.Warn("error:%v", err)
		return nil, err
	}
	for i, j := 0, len(template); i < j; i++ {
		template[i].Append = ""
		template[i].Prepend = ""
	}
	for i, j := 0, len(limit); i < j; i++ {
		limit[i].SPJMp = 0
		limit[i].CompMp = 0
	}
	detail.CreatorName = name
	detail.Tag = tags
	detail.Language = languages
	detail.Statistic = statistic
	detail.Sample = samples
	detail.Template = template
	detail.Limit = limit
	return &detail, err
}

func (Contest) IsMatched(cid, pid int64) (bool, error) {
	var ok int
	err := gosql.Get(&ok, "select count(*) from ojo.contest_problem where cid=? and pid=? limit 1", cid, pid)
	return ok == 1, err
}

func (Contest) GetStat(psmid int64) (*dto.ContestSubStat, error) {
	var s dto.ContestSubStat
	err := gosql.Get(&s, "select * from contest_submission ps where ps.id=? limit 1", psmid)
	if err != nil {
		log.Warn("error:%v", err)
		return nil, err
	}
	ptName, err := pb.GetName(s.Pid)
	if err != nil {
		log.Warn("error:%v", err)
		return nil, err
	}
	userName, err := user.GetName(s.Uid)
	if err != nil {
		log.Warn("error:%v", err)
		return nil, err
	}
	s.Username = userName
	s.ProblemName = ptName
	return &s, nil
}

func (Contest) GetSubmission(uid, pid, cid int64) (*dto.ContestSubmission, error) {
	var s dto.ContestSubmission
	err := gosql.Get(&s, "select * from contest_submission cs where cs.uid=? and cs.pid=? and cs.cid=? order by cs.submit_time desc limit 1", uid, pid, cid)
	return &s, err
}

func (Contest) GetOIRank(form dto.ContestForm) ([]dto.OIRank, error) {
	if form.Page < 1 {
		form.Page = 1
	}
	form.Page -= 1
	form.Limit = 10
	form.Offset = form.Page * 10
	// When I wrote this code, only God and I understand what it did
	// Now only God knows
	sql := `select a.uid,sum(a.max_score) total_score,max(a.submit_time) last_submit_time
			from
				(select uid,max(total_score) max_score ,max(submit_time) submit_time
				 from ojo.contest_submission
				 where cid=?
				 group by uid,pid) a
			group by uid
			order by total_score desc ,last_submit_time
			limit ?,?`
	var res []dto.OIRank
	err := gosql.Select(&res, sql, form.Cid, form.Offset, form.Limit)
	if err != nil {
		log.Warn("error:%v\n", err)
		return nil, err
	}
	for i, j := 0, len(res); i < j; i++ {
		detail, err := cts.GetOIDetail(form.Cid, res[i].Uid)
		if err != nil {
			log.Warn("error:%v\n", err)
			return nil, err
		}
		name, err := user.GetName(res[i].Uid)
		if err != nil {
			log.Warn("error:%v\n", err)
			return nil, err
		}
		res[i].OIDetail = detail
		res[i].Username = name
		res[i].Cid = form.Cid
	}
	return res, err
}

func (Contest) GetOIRankCount(cid int64) (int, error) {
	// When I wrote this code, only God and I understand what it did
	// Now only God knows
	sql := `select count(*)
			from
				(select sum(max_score)
				 from
					 (select uid,max(total_score) max_score
					  from ojo.contest_submission
					  where cid=?
					  group by uid,pid) a
				 group by uid) b`
	var count int
	err := gosql.Get(&count, sql, cid)
	fmt.Println(count)
	return count, err
}

func (Contest) GetOIDetail(cid, uid int64) ([]dto.OIDetail, error) {
	sql := `select cs.pid,max(total_score) max_score
		    from ojo.contest_submission cs,ojo.problem p
		    where cs.cid=? and cs.uid=? and cs.pid=p.id
		    group by cs.uid,cs.pid ,p.ref
		    order by p.ref`
	var res []dto.OIDetail
	err := gosql.Select(&res, sql, cid, uid)
	return res, err
}

func (Contest) Submit(form dto.SubmitForm) (*dto.ContestSubmission, error) {
	var sql = `insert into ojo.contest_submission
			(cid,uid,pid,language,status,total_score,submit_time,code)
		values(?,?,?,?,'Judging',0,now(),?)`
	exec, err := gosql.Exec(sql, form.Cid, form.Uid, form.Pid, form.Lid, form.Code)
	if err != nil {
		log.Warn("error:%v", err)
		return nil, err
	}
	id, err := exec.LastInsertId()
	if err != nil {
		log.Warn("error:%v", err)
		return nil, err
	}
	var res dto.ContestSubmission
	err = gosql.Get(&res, "select * from contest_submission where id=? limit 1", id)
	return &res, err
}

func (Contest) UpdateStat(cid, pid int64, total, ac, wa, ce, mle, re, tle, ole int) error {
	var sql = `  update ojo.contest_statistic set 
                total =total+ ?,
                ac =ac+ ?,
                wa =wa+ ?,
                re =re+ ?,
                tle =tle+ ?,
                mle =mle+ ?,
                ce =ce+ ?,
                ole =ole+ ?
        where cid = ? and pid=?`
	_, err := gosql.Exec(sql, total, ac, wa, re, tle, mle, ce, ole, cid, pid)
	return err
}

func (Contest) SetISE(csmid int64) error {
	_, err := gosql.Exec("update ojo.contest_submission set status='ISE' where id=?", csmid)
	if err != nil {
		log.Warn("error:%v", err)
	}
	return err
}

func (Contest) UpdateFlagAndScore(csmid int64, score int, flag string) error {
	var sql = `  update ojo.contest_submission set 
                status =?,
                total_score = ?
        where id = ?`
	_, err := gosql.Exec(sql, flag, score, csmid)
	return err
}

func (Contest) InsertCaseRes(csmid, uid int64, form dto.OperationForm) error {
	var sql = `  insert into ojo.contest_case_result
  (csmid,pcaseid,uid,flag,cpu_time,real_time,real_memory,real_output,error_output,score)
  				values(?,?,?,?,?,?,?,?,?,?)`
	_, err := gosql.Exec(sql, csmid, form.PcId, uid, form.Flag, form.ActualCpuTime,
		form.ActualRealTime, form.RealMemory, form.RealOutput, form.ErrorOutput, form.Score)
	return err
}
func (Contest) GetCaseRes(csmid int64) ([]dto.ContestCaseResult, error) {
	var res []dto.ContestCaseResult
	err := gosql.Select(&res, "select * from ojo.contest_case_result where csmid=?", csmid)
	return res, err
}

func (Contest) GetAllStat(cid, uid int64, offset, limit int) ([]dto.ContestSubStat, error) {
	var res []dto.ContestSubStat
	err := gosql.Select(&res, "select cs.id,cs.uid,cs.cid,cs.pid,cs.total_score,cs.language,cs.status,cs.submit_time from contest_submission cs where cs.cid=? and cs.uid=? order by cs.submit_time desc limit ?,?", cid, uid, offset, limit)
	if err != nil {
		log.Warn("error:%v", err)
		return nil, err
	}
	for i := 0; i < len(res); i++ {
		ptName, err := pb.GetName(res[i].Pid)
		if err != nil {
			log.Warn("error:%v", err)
			return nil, err
		}
		res[i].ProblemName = ptName
	}
	return res, nil
}

func (Contest) GetAllStatCount(cid, uid int64) (int, error) {
	var count int
	err := gosql.Get(&count, "select count(*) from contest_submission cs where cs.cid=? and cs.uid=?", cid, uid)
	if err != nil {
		log.Warn("error:%v", err)
		return 0, err
	}
	return count, nil
}

func (Contest) GetTime(id int64) (*dto.ContestDetail, error) {
	var data dto.ContestDetail
	err := gosql.Get(&data, "select id, start_time, end_time from contest where id=?", id)
	if err != nil {
		log.Warn("error:%v\n", err)
		return nil, err
	}
	data.Now = time.Now().Format("2006-01-02 15:04:05")
	return &data, nil
}

func (Contest) GetOITop10(cid int64) ([]dto.OIRank, error) {
	// When I wrote this code, only God and I understand what it did
	// Now only God knows
	sql := `select a.cid,a.uid,sum(a.max_score) total_score,max(a.submit_time) last_submit_time 
            from
	       (select cid,uid,max(total_score) max_score ,max(submit_time) submit_time 
	        from ojo.contest_submission 
	        where cid=? 
	        group by uid,pid,cid) a
	        group by uid,cid 
            order by total_score desc ,last_submit_time 
            limit 10 `
	var res []dto.OIRank
	err := gosql.Select(&res, sql, cid)
	if err != nil {
		log.Warn("error:%v\n", err)
		return nil, err
	}
	for i, j := 0, len(res); i < j; i++ {
		name, err := user.GetName(res[i].Uid)
		if err != nil {
			log.Warn("error:%v\n", err)
			return nil, err
		}
		res[i].Username = name
	}
	return res, err
}

func (Contest) GetACMTop10(cid int64) ([]dto.ACMRank, error) {
	// When I wrote this code, only God and I understand what it did
	// Now only God knows
	sql := `select * from ojo.contest_acm_overall where cid=? order by ac desc ,total_time limit 10 `
	var res []dto.ACMRank
	err := gosql.Select(&res, sql, cid)
	if err != nil {
		log.Warn("error:%v\n", err)
		return nil, err
	}
	for i, j := 0, len(res); i < j; i++ {
		detail, err := cts.GetACMDetail(cid, res[i].Uid)
		if err != nil {
			log.Warn("error:%v\n", err)
			return nil, err
		}
		name, err := user.GetName(res[i].Uid)
		if err != nil {
			log.Warn("error:%v\n", err)
			return nil, err
		}
		res[i].ACMDetail = detail
		res[i].Username = name
	}
	return res, err
}

func (Contest) GetACMDetail(cid, uid int64) ([]dto.ACMDetail, error) {
	sql := `select cad.*
from ojo.contest_acm_detail cad,ojo.problem p
where cad.cid=? and cad.uid=? and cad.pid=p.id
order by p.ref`
	var res []dto.ACMDetail
	err := gosql.Select(&res, sql, cid, uid)
	return res, err
}

func (Contest) GetACMRank(form dto.ContestForm) ([]dto.ACMRank, error) {
	if form.Page < 1 {
		form.Page = 1
	}
	form.Page -= 1
	form.Limit = 10
	form.Offset = form.Page * 10

	sql := `select * from ojo.contest_acm_overall where cid=? order by ac desc ,total_time limit ?,?`
	var res []dto.ACMRank
	err := gosql.Select(&res, sql, form.Cid, form.Offset, form.Limit)
	if err != nil {
		log.Warn("error:%v\n", err)
		return nil, err
	}
	for i, j := 0, len(res); i < j; i++ {
		detail, err := cts.GetACMDetail(form.Cid, res[i].Uid)
		if err != nil {
			log.Warn("error:%v\n", err)
			return nil, err
		}
		name, err := user.GetName(res[i].Uid)
		if err != nil {
			log.Warn("error:%v\n", err)
			return nil, err
		}
		res[i].ACMDetail = detail
		res[i].Username = name
	}
	return res, err
}

func (Contest) GetACMRankCount(cid int64) (int, error) {
	// When I wrote this code, only God and I understand what it did
	// Now only God knows
	sql := `select count(*) from ojo.contest_acm_overall where cid=? `
	var count int
	err := gosql.Get(&count, sql, cid)
	return count, err
}

func (Contest) HasACMOverAll(form *dto.SubmitForm) (bool, error) {
	sql := "select count(*) from ojo.contest_acm_overall where cid=? and uid=?"
	var count int
	err := gosql.Get(&count, sql, form.Cid, form.Uid)
	return count == 1, err
}

func (Contest) HasACMDetail(form *dto.SubmitForm) (bool, error) {
	sql := "select count(*) from ojo.contest_acm_detail where cid=? and uid=? and pid=?"
	var count int
	err := gosql.Get(&count, sql, form.Cid, form.Uid, form.Pid)
	return count == 1, err
}

func (Contest) InsertACMOverAll(form *dto.SubmitForm, time int, ac bool) error {
	aa := 0
	if ac {
		aa = 1
	}
	sql := "insert into ojo.contest_acm_overall(cid, uid, total, ac, total_time) VALUES (?,?,?,?,?)"
	_, err := gosql.Exec(sql, form.Cid, form.Uid, 1, aa, time)
	return err
}

func (Contest) InsertACMDetail(form *dto.SubmitForm, time int, ac, firstAc bool) error {
	sql := "insert into ojo.contest_acm_detail(cid, uid, pid, last_submit_time, total, ac, first_ac) VALUES (?,?,?,?,?,?,?)"
	_, err := gosql.Exec(sql, form.Cid, form.Uid, form.Pid, time, 1, ac, firstAc)
	return err
}

func (Contest) UpdateACMOverAll(form *dto.SubmitForm, time int, ac bool) error {
	sql := `update ojo.contest_acm_overall 
			set total=total+? , ac=ac+? , total_time=? 
			where cid=? and uid=?`
	aa := 0
	if ac {
		aa = 1
	}
	_, err := gosql.Exec(sql, 1, aa, time, form.Cid, form.Uid)
	return err
}

func (Contest) UpdateACMDetail(form *dto.SubmitForm, time int, ac, first bool) error {
	aa := 0
	if ac {
		aa = 1
	}
	sql := `update ojo.contest_acm_detail 
			set total=total+? ,ac=? ,last_submit_time=? ,first_ac=?
			where cid=? and uid=? and pid=?`

	_, err := gosql.Exec(sql, 1, aa, time, first, form.Cid, form.Uid, form.Pid)
	return err
}

func (Contest) HasACMFirstDetail(form *dto.SubmitForm) (bool, error) {
	sql := "select count(*) from ojo.contest_acm_detail where cid=? and pid=? and ac=1"
	var count int
	err := gosql.Get(&count, sql, form.Cid, form.Pid)
	return count == 0, err
}

// 根据Cid获得ACM提交错误的次数
func (Contest) GetACMWrong(form *dto.SubmitForm) (int, error) {
	sql := "select a.total-a.ac from (select total,ac from ojo.contest_acm_overall where cid=? and uid=?) a"
	var count int
	err := gosql.Get(&count, sql, form.Cid, form.Uid)
	return count, err
}

func (Contest) SelectCreatorName(lens int, getId func(i int) (target int64), setName func(i int, res string)) error {
	if lens == 0 {
		return nil
	}
	ids := make([]int64, 0, lens)
	for i := 0; i < lens; i++ {
		ids = append(ids, getId(i))
	}
	var s []dto.Username
	err := gosql.Select(&s, "select id,username from ojo.user  where id in (?) ", ids)
	if err != nil {
		return err
	}
	for i := 0; i < lens; i++ {
		for j, k := 0, len(s); j < k; j++ {
			if getId(i) == s[j].Id {
				setName(i, s[j].Username)
				break
			}
		}
	}
	return nil
}

func (Contest) GetCreatorId(id int64) (int64, error) {
	var cid int64
	err := gosql.Get(&cid, "select cid from ojo.contest where id=?", id)
	return cid, err
}

func (Contest) SetVisibleTrue(id int64) error {
	s := "update ojo.contest set visible=true where id=? limit 1"
	_, err := gosql.Exec(s, id)
	return err
}

func (Contest) SetVisibleFalse(id int64) error {
	s := "update ojo.contest set visible=false where id=? limit 1"
	_, err := gosql.Exec(s, id)
	return err
}

func (Contest) InsertContest(c *dto.Contest) error {
	s := `insert into ojo.contest(
                title, description, rule,
                cid, password,punish,visible,
				submit_limit,show_output,show_rank,                        
                start_time, end_time, create_time,
                last_update_time ) 
                VALUES (?,?,?,?,?,?,?,?,?,?,?,?,now(),now())`
	db, err := gosql.Begin()
	if err != nil {
		log.Warn("%v", err)
		return err
	}
	res, err := db.Exec(s, c.Title, c.Description, c.Rule,
		c.Cid, c.Password, c.Punish, c.Visible, c.SubmitLimit,
		c.ShowOutput, c.ShowRank, c.StartTime, c.EndTime)
	if err != nil {
		log.Warn("%v", err)
		err2 := db.Rollback()
		if err2 != nil {
			log.Warn("%v", err2)
		}
		return err
	}
	id, err := res.LastInsertId()
	if len(c.IPLimit) == 0 {
		err := cts.InsertIPRange(db, &dto.ContestIPLimit{
			Id:      0,
			Cid:     id,
			Address: "0.0.0.0",
			Mask:    0,
		})
		if err != nil {
			log.Warn("%v", err)
			err2 := db.Rollback()
			if err2 != nil {
				log.Warn("%v", err2)
			}
			return err
		}
	} else {
		for i, j := 0, len(c.IPLimit); i < j; i++ {
			c.IPLimit[i].Cid = id
			err := cts.InsertIPRange(db, &c.IPLimit[i])
			if err != nil {
				log.Warn("%v", err)
				err2 := db.Rollback()
				if err2 != nil {
					log.Warn("%v", err2)
				}
				return err
			}
		}
	}
	err = db.Commit()
	if err != nil {
		log.Warn("%v", err)
		err2 := db.Rollback()
		if err2 != nil {
			log.Warn("%v", err2)
		}
		return err
	}
	return err
}

func (Contest) InsertCtsPb(cid, pid int64) error {
	s1 := `insert into ojo.contest_problem(cid, pid) 
                VALUES (?,?)`
	s2 := `insert into ojo.contest_statistic(cid, pid)
 				VALUES (?,?)`
	tx, err := gosql.Begin()
	if err != nil {
		log.Warn("%v", err)
		return err
	}
	_, err = tx.Exec(s1, cid, pid)
	if err != nil {
		log.Warn("%v", err)
		err2 := tx.Rollback()
		if err2 != nil {
			log.Warn("%v", err2)
		}
		return err
	}
	_, err = tx.Exec(s2, cid, pid)
	if err != nil {
		log.Warn("%v", err)
		err2 := tx.Rollback()
		if err2 != nil {
			log.Warn("%v", err2)
		}
		return err
	}
	err = tx.Commit()
	if err != nil {
		log.Warn("%v", err)
		err2 := tx.Rollback()
		if err2 != nil {
			log.Warn("%v", err2)
		}
		return err
	}
	return err
}

func (Contest) DeleteContest(cid int64) error {
	s1 := `delete from contest where id=? limit 1`
	s2 := `delete from contest_ip_limit where cid=?`
	s3 := `delete from contest_problem where cid=?`
	tx, err := gosql.Begin()
	if err != nil {
		log.Warn("%v", err)
		return err
	}
	_, err = tx.Exec(s1, cid)
	if err != nil {
		log.Warn("%v", err)
		err2 := tx.Rollback()
		if err2 != nil {
			log.Warn("%v", err2)
		}
		return err
	}
	_, err = tx.Exec(s2, cid)
	if err != nil {
		log.Warn("%v", err)
		err2 := tx.Rollback()
		if err2 != nil {
			log.Warn("%v", err2)
		}
		return err
	}
	_, err = tx.Exec(s3, cid)
	if err != nil {
		log.Warn("%v", err)
		err2 := tx.Rollback()
		if err2 != nil {
			log.Warn("%v", err2)
		}
		return err
	}
	err = tx.Commit()
	if err != nil {
		log.Warn("%v", err)
		err2 := tx.Rollback()
		if err2 != nil {
			log.Warn("%v", err2)
		}
		return err
	}
	return err
}

func (Contest) DeleteCtsPb(cid, pid int64) error {
	s1 := `delete from ojo.contest_problem
			where cid=? and pid=? limit 1`
	s2 := `delete from ojo.contest_statistic
			where cid=? and pid=? limit 1`
	tx, err := gosql.Begin()
	if err != nil {
		log.Warn("%v", err)
		return err
	}
	_, err = tx.Exec(s1, cid, pid)
	if err != nil {
		log.Warn("%v", err)
		err2 := tx.Rollback()
		if err2 != nil {
			log.Warn("%v", err2)
		}
		return err
	}
	_, err = tx.Exec(s2, cid, pid)
	if err != nil {
		log.Warn("%v", err)
		err2 := tx.Rollback()
		if err2 != nil {
			log.Warn("%v", err2)
		}
		return err
	}
	err = tx.Commit()
	if err != nil {
		log.Warn("%v", err)
		err2 := tx.Rollback()
		if err2 != nil {
			log.Warn("%v", err2)
		}
		return err
	}
	return err
}

func (Contest) GetStartEndTime(pid int64) (*dto.ContestTime, error) {
	s := "select start_time,end_time from contest where id=?"
	var res dto.ContestTime
	err := gosql.Get(&res, s, pid)
	return &res, err
}

func (Contest) InsertIPRange(db *gosql.DB, limit *dto.ContestIPLimit) error {
	s := `insert into ojo.contest_ip_limit(cid, address,mask) VALUES (?,?,?)`
	_, err := db.Exec(s, limit.Cid, limit.Address, limit.Mask)
	return err
}

func (Contest) GetDetail(id int64) (*dto.Contest, error) {
	var detail dto.Contest
	err := gosql.Get(&detail, `select id, title,
       description, rule, start_time,
       end_time, create_time, last_update_time,
       cid, punish, visible, submit_limit,show_output,show_rank
       from ojo.contest c where c.id=? limit 1`, id)
	if err != nil {
		log.Warn("error:%v", err)
		return nil, err
	}
	limit, err := cts.GetIPLimit(id)
	if err != nil {
		log.Warn("error:%v", err)
		return nil, err
	}
	detail.IPLimit = limit
	detail.Now = time.Now().Format("2006-01-02 15:04:05")
	return &detail, err
}

func (Contest) GetIPLimit(id int64) ([]dto.ContestIPLimit, error) {
	var res []dto.ContestIPLimit
	err := gosql.Select(&res, `select id, cid,address,mask
       from ojo.contest_ip_limit c where c.cid=?`, id)
	return res, err
}

func (Contest) UpdateContest(c *dto.Contest) error {
	s := ""
	if c.ChangePassword {
		s = `update ojo.contest set 
                        title=?,
                        description=?,
                        rule=?,
                        password=?,
                        punish=?,
                        visible=?,
						submit_limit=?,    
                       	show_rank=?,
                        show_output=?,
                		start_time=?,
                        end_time=?, 
                        last_update_time=now()
			 where id=?`
	} else {
		s = `update ojo.contest set 
                        title=?,
                        description=?,
                        rule=?,
                        punish=?,
                        visible=?,
						submit_limit=?,    
                       	show_rank=?,
                        show_output=?,
                		start_time=?,
                        end_time=?, 
                        last_update_time=now()
			 where id=?`
	}
	tx, err := gosql.Begin()
	if err != nil {
		log.Warn("%v", err)
		return err
	}
	if c.ChangePassword {
		_, err = tx.Exec(s, c.Title, c.Description, c.Rule,
			c.Password, c.Punish, c.Visible, c.SubmitLimit,
			c.ShowRank, c.ShowOutput, c.StartTime, c.EndTime, c.Id)
		if err != nil {
			log.Warn("%v", err)
			err2 := tx.Rollback()
			if err2 != nil {
				log.Warn("%v", err2)
			}
			return err
		}
	} else {
		_, err = tx.Exec(s, c.Title, c.Description, c.Rule,
			c.Punish, c.Visible, c.SubmitLimit,
			c.ShowRank, c.ShowOutput, c.StartTime, c.EndTime, c.Id)
		if err != nil {
			log.Warn("%v", err)
			err2 := tx.Rollback()
			if err2 != nil {
				log.Warn("%v", err2)
			}
			return err
		}
	}
	err = cts.DeleteIPLimit(tx, c.Id)
	if err != nil {
		log.Warn("%v", err)
		err2 := tx.Rollback()
		if err2 != nil {
			log.Warn("%v", err2)
		}
		return err
	}
	if len(c.IPLimit) == 0 {
		err := cts.InsertIPRange(tx, &dto.ContestIPLimit{
			Id:      0,
			Cid:     c.Id,
			Address: "0.0.0.0",
			Mask:    0,
		})
		if err != nil {
			log.Warn("%v", err)
			err2 := tx.Rollback()
			if err2 != nil {
				log.Warn("%v", err2)
			}
			return err
		}
	} else {
		for i, j := 0, len(c.IPLimit); i < j; i++ {
			c.IPLimit[i].Cid = c.Id
			err := cts.InsertIPRange(tx, &c.IPLimit[i])
			if err != nil {
				log.Warn("%v", err)
				err2 := tx.Rollback()
				if err2 != nil {
					log.Warn("%v", err2)
				}
				return err
			}
		}
	}
	err = tx.Commit()
	if err != nil {
		log.Warn("%v", err)
		err2 := tx.Rollback()
		if err2 != nil {
			log.Warn("%v", err2)
		}
		return err
	}
	return nil
}

func (Contest) DeleteIPLimit(db *gosql.DB, id int64) error {
	_, err := db.Exec("delete from contest_ip_limit where cid=?", id)
	return err
}

func (Contest) GetTodayCount() ([]dto.TodayCount, error) {
	var data []dto.TodayCount
	err := gosql.Select(&data, `SELECT
    DATE_FORMAT(submit_time, '%H') hour,
       count( * ) AS count
	FROM
    ojo.contest_submission
	WHERE
    submit_time between curdate()  and date_sub(curdate(),interval -1 day )
	GROUP BY
    hour
	ORDER BY
    hour;`)
	return data, err
}

func (Contest) GetWeekCount() (dto.WeekCount, error) {
	var data []dto.DayCount
	err := gosql.Select(&data, `SELECT
    DATE_FORMAT(submit_time, '%Y-%m-%d') day,
    count( * ) AS count
FROM
    ojo.contest_submission
WHERE
    submit_time between date_sub(curdate(),interval 7 day) and curdate()
GROUP BY
    day
ORDER BY
    day;`)
	if err != nil {
		log.Warn("%v", err)
		return dto.WeekCount{}, err
	}
	var res dto.WeekCount
	res.DayCount = data
	var now = time.Now()
	res.Today = now.Format("2006-01-02")
	return res, err
}

func (Contest) GetMonthCount() (dto.MonthCount, error) {
	var data []dto.DayCount
	err := gosql.Select(&data, `SELECT
    DATE_FORMAT(submit_time, '%Y-%m-%d') day,
    count( * ) AS count
FROM
    ojo.contest_submission
WHERE
    submit_time between date_sub(curdate(),interval 30 day) and curdate()
GROUP BY
    day
ORDER BY
    day;`)
	if err != nil {
		log.Warn("%v", err)
		return dto.MonthCount{}, err
	}
	var res dto.MonthCount
	res.DayCount = data
	var now = time.Now()
	res.Today = now.Format("2006-01-02")
	return res, err
}

func (Contest) GetRecentCount() (int, error) {
	var count int
	err := gosql.Get(&count, `SELECT
    count( * ) AS count
	FROM
    ojo.contest
	WHERE
    start_time<= date_sub(now(),interval -6 hour ) and end_time>now()`)
	return count, err
}
