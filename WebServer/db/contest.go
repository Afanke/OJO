package db

import (
	"fmt"
	"github.com/afanke/OJO/WebServer/dto"
	"github.com/afanke/OJO/utils/log"
	"time"
)

type Contest struct{}

var cts Contest

func (Contest) GetAll(form *dto.ContestForm) ([]dto.ContestBrief, error) {
	if form.Page < 1 {
		form.Page = 1
	}
	form.Page -= 1
	form.Limit = 5
	form.Offset = form.Page * 5
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
	sql += " order by create_time desc limit :offset, :limit"
	rows, err := db.NamedQuery(sql, &form)
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
	var count int
	rows, err := db.NamedQuery(sql, &form)
	if err != nil {
		log.Warn("error:%v", err)
		return 0, err
	}
	_ = rows.Next()
	err = rows.Scan(&count)
	return count, err
}

func (Contest) GetDetail(id int) (*dto.ContestDetail, error) {
	var data dto.ContestDetail
	err := db.Get(&data, "select id, title, description, rule, start_time, end_time, cid,punish_time from contest where id=?", id)
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

func (Contest) GetQualification(uid, cid int) (bool, error) {
	var ok int
	err := db.Get(&ok, "select count(*) from ojo.contest_user where uid=? and cid=? limit 1", uid, cid)
	return ok == 1, err
}

func (Contest) AddQualification(uid, cid int) error {
	_, err := db.Exec("insert into ojo.contest_user(cid, uid) values(?,?)", cid, uid)
	return err
}

func (Contest) GetPassword(cid int) (string, error) {
	var res string
	err := db.Get(&res, "select password from ojo.contest where id=?", cid)
	return res, err
}

func (Contest) GetStartTime(cid int) (time.Time, error) {
	var res string
	err := db.Get(&res, "select start_time from ojo.contest where id=?", cid)
	if err != nil {
		log.Warn("error:%v\n", err)
		return time.Time{}, err
	}
	data, err := time.Parse("2006-01-02 15:04:05", res)
	if err != nil {
		log.Warn("error:%v\n", err)
		return time.Time{}, err
	}
	return data, err
}

func (Contest) GetAllProblem(cid int) ([]dto.CtsPbBrief, error) {
	var sql = `select p.id,p.title,p.ref from contest_problem cp,problem p where cp.pid=p.id and cp.cid=?`
	var data []dto.CtsPbBrief
	err := db.Select(&data, sql, cid)
	if err != nil {
		log.Warn("error:%v\n", err)
		return nil, err
	}
	for i, l := 0, len(data); i < l; i++ {
		stat, err := cts.GetStatistic(cid, int64(data[i].Id))
		if err != nil {
			log.Warn("error:%v\n", err)
			return nil, err
		}
		data[i].Statistic = stat
	}
	return data, nil
}

func (Contest) GetAllProblemName(cid int) ([]dto.CtsPbBrief, error) {
	var sql = `select p.id,p.title,p.ref from contest_problem cp,problem p where cp.pid=p.id and cp.cid=? order by p.ref`
	var data []dto.CtsPbBrief
	err := db.Select(&data, sql, cid)
	return data, err
}

func (Contest) GetStatistic(cid int, pid int64) (*dto.ContestStatistic, error) {
	var stat dto.ContestStatistic
	err := db.Get(&stat, "select * from ojo.contest_statistic where cid=? and pid=? limit 1", cid, pid)
	return &stat, err
}

func (Contest) GetProblemDetail(cid int, pid int64) (*dto.ContestProblem, error) {
	var detail dto.ContestProblem
	err := db.Get(&detail, `select id, cid, ref, title, description, input_description, output_description, hint, create_time, last_update_time, cpu_time_limit, memory_limit, io_mode, difficulty, real_time_limit, source from ojo.problem p where p.id=? limit 1`, pid)
	if err != nil {
		log.Warn("error:%v", err)
		return nil, err
	}
	statistic, err := cts.GetStatistic(cid, pid)
	if err != nil {
		log.Warn("error:%v", err)
		return nil, err
	}
	tags, err := pb.GetProblemTag(pid)
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
	detail.CreatorName = name
	detail.Tags = tags
	detail.Languages = languages
	detail.Statistic = statistic
	detail.Samples = samples
	return &detail, err
}

func (Contest) IsMatched(cid int, pid int64) (bool, error) {
	var ok int
	err := db.Get(&ok, "select count(*) from ojo.contest_problem where cid=? and pid=? limit 1", cid, pid)
	return ok == 1, err
}

func (Contest) GetStat(psmid int) (*dto.ContestSubStat, error) {
	var s dto.ContestSubStat
	err := db.Get(&s, "select * from contest_submission ps where ps.id=? limit 1", psmid)
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

func (Contest) GetSubmission(uid int, pid int64, cid int) (*dto.ContestSubmission, error) {
	var s dto.ContestSubmission
	err := db.Get(&s, "select * from contest_submission cs where cs.uid=? and cs.pid=? and cs.cid=? order by cs.submit_time desc limit 1", uid, pid, cid)
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
	err := db.Select(&res, sql, form.Cid, form.Offset, form.Limit)
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

func (Contest) GetOIRankCount(cid int) (int, error) {
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
	err := db.Get(&count, sql, cid)
	fmt.Println(count)
	return count, err
}

func (Contest) GetOIDetail(cid, uid int) ([]dto.OIDetail, error) {
	sql := `select cs.pid,max(total_score) max_score
		    from ojo.contest_submission cs,ojo.problem p
		    where cs.cid=? and cs.uid=? and cs.pid=p.id
		    group by cs.uid,cs.pid ,p.ref
		    order by p.ref`
	var res []dto.OIDetail
	err := db.Select(&res, sql, cid, uid)
	return res, err
}

func (Contest) Submit(form dto.SubmitForm) (*dto.ContestSubmission, error) {
	var sql = `insert into ojo.contest_submission
			(cid,uid,pid,language,status,total_score,submit_time,code)
		values(?,?,?,?,'Judging',0,now(),?)`
	exec, err := db.Exec(sql, form.Cid, form.Uid, form.Pid, form.Language, form.Code)
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
	err = db.Get(&res, "select * from contest_submission where id=? limit 1", id)
	return &res, err
}

func (Contest) UpdateStat(cid int, pid int64, total, ac, wa, ce, mle, re, tle, ole int) error {
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
	_, err := db.Exec(sql, total, ac, wa, re, tle, mle, ce, ole, cid, pid)
	return err
}

func (Contest) SetISE(csmid int) error {
	_, err := db.Exec("update ojo.contest_submission set status='ISE' where id=?", csmid)
	if err != nil {
		log.Warn("error:%v", err)
	}
	return err
}

func (Contest) UpdateFlagAndScore(csmid, score int, flag string) error {
	var sql = `  update ojo.contest_submission set 
                status =?,
                total_score = ?
        where id = ?`
	_, err := db.Exec(sql, flag, score, csmid)
	return err
}

func (Contest) InsertCaseRes(csmid, uid int, form dto.OperationForm) error {
	var sql = `  insert into ojo.contest_case_result
  (csmid,pcaseid,uid,flag,cpu_time,real_time,real_memory,real_output,error_output,score)
  				values(?,?,?,?,?,?,?,?,?,?)`
	_, err := db.Exec(sql, csmid, form.PcId, uid, form.Flag, form.ActualCpuTime,
		form.ActualRealTime, form.RealMemory, form.RealOutput, form.ErrorOutput, form.Score)
	return err
}
func (Contest) GetCaseRes(csmid int) ([]dto.ContestCaseResult, error) {
	var res []dto.ContestCaseResult
	err := db.Select(&res, "select * from ojo.contest_case_result where csmid=?", csmid)
	return res, err
}

func (Contest) GetAllStat(cid, uid, offset, limit int) ([]dto.ContestSubStat, error) {
	var res []dto.ContestSubStat
	err := db.Select(&res, "select cs.id,cs.uid,cs.cid,cs.pid,cs.total_score,cs.language,cs.status,cs.submit_time from contest_submission cs where cs.cid=? and cs.uid=? order by cs.submit_time desc limit ?,?", cid, uid, offset, limit)
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

func (Contest) GetAllStatCount(cid, uid int) (int, error) {
	var count int
	err := db.Get(&count, "select count(*) from contest_submission cs where cs.cid=? and cs.uid=?", cid, uid)
	if err != nil {
		log.Warn("error:%v", err)
		return 0, err
	}
	return count, nil
}

func (Contest) GetTime(id int) (*dto.ContestDetail, error) {
	var data dto.ContestDetail
	err := db.Get(&data, "select id, start_time, end_time from contest where id=?", id)
	if err != nil {
		log.Warn("error:%v\n", err)
		return nil, err
	}
	data.Now = time.Now().Format("2006-01-02 15:04:05")
	return &data, nil
}

func (Contest) GetOITop10(cid int) ([]dto.OIRank, error) {
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
	err := db.Select(&res, sql, cid)
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

func (Contest) GetACMTop10(cid int) ([]dto.ACMRank, error) {
	// When I wrote this code, only God and I understand what it did
	// Now only God knows
	sql := `select * from ojo.contest_acm_overall where cid=? order by ac desc ,total_time limit 10 `
	var res []dto.ACMRank
	err := db.Select(&res, sql, cid)
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

func (Contest) GetACMDetail(cid, uid int) ([]dto.ACMDetail, error) {
	sql := `select cad.*
from ojo.contest_acm_detail cad,ojo.problem p
where cad.cid=? and cad.uid=? and cad.pid=p.id
order by p.ref`
	var res []dto.ACMDetail
	err := db.Select(&res, sql, cid, uid)
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
	err := db.Select(&res, sql, form.Cid, form.Offset, form.Limit)
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

func (Contest) GetACMRankCount(cid int) (int, error) {
	// When I wrote this code, only God and I understand what it did
	// Now only God knows
	sql := `select count(*) from ojo.contest_acm_overall where cid=? `
	var count int
	err := db.Get(&count, sql, cid)
	return count, err
}

func (Contest) HasACMOverAll(form *dto.SubmitForm) (bool, error) {
	sql := "select count(*) from ojo.contest_acm_overall where cid=? and uid=?"
	var count int
	err := db.Get(&count, sql, form.Cid, form.Uid)
	return count == 1, err
}

func (Contest) HasACMDetail(form *dto.SubmitForm) (bool, error) {
	sql := "select count(*) from ojo.contest_acm_detail where cid=? and uid=? and pid=?"
	var count int
	err := db.Get(&count, sql, form.Cid, form.Uid, form.Pid)
	return count == 1, err
}

func (Contest) InsertACMOverAll(form *dto.SubmitForm, time int, ac bool) error {
	aa := 0
	if ac {
		aa = 1
	}
	sql := "insert into ojo.contest_acm_overall(cid, uid, total, ac, total_time) VALUES (?,?,?,?,?)"
	_, err := db.Exec(sql, form.Cid, form.Uid, 1, aa, time)
	return err
}

func (Contest) InsertACMDetail(form *dto.SubmitForm, time int, ac, firstAc bool) error {
	sql := "insert into ojo.contest_acm_detail(cid, uid, pid, last_submit_time, total, ac, first_ac) VALUES (?,?,?,?,?,?,?)"
	_, err := db.Exec(sql, form.Cid, form.Uid, form.Pid, time, 1, ac, firstAc)
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
	_, err := db.Exec(sql, 1, aa, time, form.Cid, form.Uid)
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

	_, err := db.Exec(sql, 1, aa, time, first, form.Cid, form.Uid, form.Pid)
	return err
}

func (Contest) HasACMFirstDetail(form *dto.SubmitForm) (bool, error) {
	sql := "select count(*) from ojo.contest_acm_detail where cid=? and pid=? and ac=1"
	var count int
	err := db.Get(&count, sql, form.Cid, form.Pid)
	return count == 0, err
}

func (Contest) GetACMWrong(form *dto.SubmitForm) (int, error) {
	sql := "select a.total-a.ac from (select total,ac from ojo.contest_acm_overall where cid=? and uid=?) a"
	var count int
	err := db.Get(&count, sql, form.Cid, form.Uid)
	return count, err
}
