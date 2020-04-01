package db

import (
	"database/sql"
	"errors"
	"github.com/afanke/OJO/WebServer/dto"
	"github.com/afanke/OJO/utils/log"
)

type QueryForm struct {
	P1     string `db:"p1"`
	P2     string `db:"p2"`
	P3     string `db:"p3"`
	P4     string `db:"p4"`
	Offset int    `db:"offset"`
	Limit  int    `db:"limit"`
}

type Practice struct{}

var pt = Practice{}

func (Practice) GetAll(form *dto.PracticeForm) ([]dto.PracticeBrief, error) {
	var s = `select p.id, p.ref, p.title, p.difficulty
			from ojo.problem p `
	if form.Tid != 0 {
		s += ", ojo.problem_tag pt "
	}
	s += "where 1=1 "
	if form.Tid != 0 {
		s += "and pt.tid = :tid and pt.pid=p.id "
	}
	if form.Keywords != "" {
		s += "and title like concat('%',:keywords,'%') "
	}
	if form.Difficulty != "" {
		s += "and difficulty=:difficulty "
	}
	s += "and p.visible=1 order by p.ref limit :offset, :limit"
	rows, err := db.NamedQuery(s, &form)
	if err != nil {
		log.Warn("error:%v", err)
		return nil, err
	}
	var rest = make([]dto.PracticeBrief, 0, form.Limit)
	for rows.Next() {
		var res dto.PracticeBrief
		err := rows.StructScan(&res)
		if err != nil {
			log.Warn("error:%v", err)
			return nil, err
		}
		stat, err := pt.GetStatistic(res.Id)
		if err != nil {
			log.Warn("error:%v", err)
			return nil, err
		}
		tag, err := tag.GetBriefByPid(res.Id)
		if err != nil {
			log.Warn("error:%v", err)
			return nil, err
		}
		res.Tags = tag
		res.Statistic = stat
		rest = append(rest, res)
	}
	return rest, nil

}

func (Practice) GetAllTags() ([]dto.Tag, error) {
	var tags []dto.Tag
	err := db.Select(&tags, "select * from ojo.tag")
	return tags, err
}

func (Practice) GetCount(form *dto.PracticeForm) (int, error) {
	var s = `select count(*)
			from ojo.problem p `
	if form.Tid != 0 {
		s += ", ojo.problem_tag pt "
	}
	s += "where 1=1 "
	if form.Tid != 0 {
		s += "and pt.tid =:tid and pt.pid=p.id "
	}
	if form.Keywords != "" {
		s += "and title like concat('%',:keywords,'%') "
	}
	if form.Difficulty != "" {
		s += "and difficulty=:difficulty "
	}
	s += "and p.visible=1 "
	var count int
	rows, err := db.NamedQuery(s, &form)
	if err != nil {
		log.Warn("error:%v", err)
		return 0, err
	}
	_ = rows.Next()
	err = rows.Scan(&count)
	return count, err
}

func (Practice) GetStatistic(pbid int64) (*dto.PracticeStatistic, error) {
	var stat dto.PracticeStatistic
	err := db.Get(&stat, "select * from practice_statistic where pbid=? limit 1", pbid)
	return &stat, err
}

func (Practice) GetDetail(pbid int64) (*dto.Practice, error) {
	var detail dto.Practice
	err := db.Get(&detail, `select * from ojo.problem p where p.id=? and p.visible=1 limit 1`, pbid)
	if err != nil {
		log.Warn("error:%v", err)
		return nil, err
	}
	if !detail.Visible {
		return nil, errors.New("failed to access")
	}
	statistic, err := pt.GetStatistic(pbid)
	if err != nil {
		log.Warn("error:%v", err)
		return nil, err
	}
	tags, err := tag.GetBriefByPid(pbid)
	if err != nil {
		log.Warn("error:%v", err)
		return nil, err
	}
	languages, err := pb.GetLanguage(pbid)
	if err != nil {
		log.Warn("error:%v", err)
		return nil, err
	}
	samples, err := pb.GetSample(pbid)
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

func (Practice) GetSubmission(uid, pid int64) (*dto.PracticeSubmission, error) {
	var s dto.PracticeSubmission
	err := db.Get(&s, "select * from practice_submission ps where ps.uid=? and ps.pid=? order by ps.submit_time desc limit 1", uid, pid)
	return &s, err
}

func (Practice) GetAllStat(uid int64, offset, limit int) ([]dto.PracticeSubStat, error) {
	var res []dto.PracticeSubStat
	err := db.Select(&res, "select ps.id,ps.uid,ps.pid,ps.total_score,ps.language,ps.status,ps.submit_time from practice_submission ps where ps.uid=? order by ps.submit_time desc limit ?,?", uid, offset, limit)
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

func (Practice) GetAllStatCount(uid int64) (int, error) {
	var count int
	err := db.Get(&count, "select count(*) from practice_submission ps where ps.uid=?", uid)
	if err != nil {
		log.Warn("error:%v", err)
		return 0, err
	}
	return count, nil
}

func (Practice) GetStat(psmid int64) (*dto.PracticeSubStat, error) {
	var s dto.PracticeSubStat
	err := db.Get(&s, "select * from practice_submission ps where ps.id=? limit 1", psmid)
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

func (Practice) GetCaseRes(psmid int64) ([]dto.PracticeCaseResult, error) {
	var res []dto.PracticeCaseResult
	err := db.Select(&res, "select * from ojo.practice_case_result where psmid=?", psmid)
	return res, err
}

func (Practice) Submit(form dto.SubmitForm) (*dto.PracticeSubmission, error) {
	var sql = `insert into ojo.practice_submission
			(uid,pid,language,status,total_score,submit_time,code)
		values(?,?,?,'Judging',0,now(),?)`
	exec, err := db.Exec(sql, form.Uid, form.Pid, form.Language, form.Code)
	if err != nil {
		log.Warn("error:%v", err)
		return nil, err
	}
	id, err := exec.LastInsertId()
	if err != nil {
		log.Warn("error:%v", err)
		return nil, err
	}
	var res dto.PracticeSubmission
	err = db.Get(&res, "select * from practice_submission where id=? limit 1", id)
	return &res, err
}

func (Practice) UpdateStat(pbid int64, total, ac, wa, ce, mle, re, tle, ole int) error {
	var sql = `  update ojo.practice_statistic set 
                total =total+ ?,
                ac =ac+ ?,
                wa =wa+ ?,
                re =re+ ?,
                tle =tle+ ?,
                mle =mle+ ?,
                ce =ce+ ?,
                ole =ole+ ?
        where pbid = ?`
	_, err := db.Exec(sql, total, ac, wa, re, tle, mle, ce, ole, pbid)
	return err
}

func (Practice) SetISE(psmid int64) error {
	_, err := db.Exec("update ojo.practice_submission set status='ISE' where id=?", psmid)
	if err != nil {
		log.Warn("error:%v", err)
	}
	return err
}

func (Practice) UpdateFlagAndScore(psmid int64, score int, flag string) error {
	var sql = `  update ojo.practice_submission set 
                status =?,
                total_score = ?
        where id = ?`
	_, err := db.Exec(sql, flag, score, psmid)
	return err
}

func (Practice) InsertCaseRes(psmid, uid int64, form dto.OperationForm) error {
	var sql = `  insert into ojo.practice_case_result
  (psmid,pcaseid,uid,flag,cpu_time,real_time,real_memory,real_output,error_output,score)
  				values(?,?,?,?,?,?,?,?,?,?)`
	_, err := db.Exec(sql, psmid, form.PcId, uid, form.Flag, form.ActualCpuTime,
		form.ActualRealTime, form.RealMemory, form.RealOutput, form.ErrorOutput, form.Score)
	return err
}

func (Practice) InsertStatistic(tx *sql.Tx, pbid int64) error {
	s := "insert into ojo.practice_statistic(pbid) values (?)"
	_, err := tx.Exec(s, pbid)
	return err
}
