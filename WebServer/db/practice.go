package db

import (
	"errors"
	"github.com/afanke/OJO/WebServer/dto"
	"github.com/afanke/OJO/utils/log"
	"github.com/ilibs/gosql/v2"
	"time"
)

type QueryForm struct {
	P1     string `gosql:"p1"`
	P2     string `gosql:"p2"`
	P3     string `gosql:"p3"`
	P4     string `gosql:"p4"`
	Offset int    `gosql:"offset"`
	Limit  int    `gosql:"limit"`
}

type Practice struct{}

var pt = Practice{}

var PracticePageSize = 20

func (Practice) GetAll(form *dto.PracticeForm) ([]dto.PracticeBrief, error) {
	if form.Page < 1 {
		form.Offset = 0
	} else {
		form.Offset = (form.Page - 1) * PracticePageSize
	}
	form.Limit = PracticePageSize
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
	rows, err := gosql.Sqlx().NamedQuery(s, &form)
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
		if form.Uid != 0 {
			flag, err := pt.GetFlagByUidPid(form.Uid, res.Id)
			if err != nil {
				log.Warn("error:%v", err)
				return nil, err
			}
			res.Flag = flag
		}
		res.Tags = tag
		res.Statistic = stat
		rest = append(rest, res)
	}
	return rest, nil

}

func (Practice) GetAllTags() ([]dto.Tag, error) {
	var tags []dto.Tag
	err := gosql.Select(&tags, "select * from ojo.tag")
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
	rows, err := gosql.Sqlx().NamedQuery(s, &form)
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
	err := gosql.Get(&stat, "select * from practice_statistic where pbid=? limit 1", pbid)
	return &stat, err
}

func (Practice) GetDetail(pbid int64) (*dto.Practice, error) {
	var detail dto.Practice
	err := gosql.Get(&detail, `select * from ojo.problem p where p.id=? and p.visible=1 limit 1`, pbid)
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
	template, err := pb.GetTemplate(pbid)
	if err != nil {
		log.Warn("error:%v", err)
		return nil, err
	}
	limit, err := pb.GetLimit(pbid)
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

func (Practice) GetSubmission(uid, pid int64) (*dto.PracticeSubmission, error) {
	var s dto.PracticeSubmission
	err := gosql.Get(&s, "select * from practice_submission ps where ps.uid=? and ps.pid=? order by ps.submit_time desc limit 1", uid, pid)
	return &s, err
}

func (Practice) GetAllStatus(uid int64, offset, limit int) ([]dto.PracticeSubStat, error) {
	var res []dto.PracticeSubStat
	err := gosql.Select(&res, "select ps.id,ps.uid,ps.pid,ps.total_score,ps.lid,ps.flag,ps.submit_time from practice_submission ps where ps.uid=? order by ps.submit_time desc limit ?,?", uid, offset, limit)
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

func (Practice) GetAllStatusCount(uid int64) (int, error) {
	var count int
	err := gosql.Get(&count, "select count(*) from practice_submission ps where ps.uid=?", uid)
	if err != nil {
		log.Warn("error:%v", err)
		return 0, err
	}
	return count, nil
}

func (Practice) GetStatus(psmid int64) (*dto.PracticeSubStat, error) {
	var s dto.PracticeSubStat
	err := gosql.Get(&s, "select * from practice_submission ps where ps.id=? limit 1", psmid)
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
	err := gosql.Select(&res, "select * from ojo.practice_case_result where psmid=?", psmid)
	return res, err
}

func (Practice) GetFlagByUidPid(uid, pid int64) (string, error) {
	var res string
	err := gosql.Get(&res,
		`select flag
				from ojo.practice_submission
				where uid=? and pid=?
				order by submit_time desc limit 1`,
		uid, pid)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return "", nil
		}
	}
	return res, err
}

// -------------------Practice Submit-------------------------

func (Practice) Submit(form *dto.SubmitForm) (*dto.PracticeSubmission, error) {
	var sql = `insert into ojo.practice_submission
			(uid,pid,lid,code,submit_time,total_score,flag,error_msg)
		values(?,?,?,?,now(),0,'JUG','')`
	exec, err := gosql.Exec(sql, form.Uid, form.Pid, form.Lid, form.Code)
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
	err = gosql.Get(&res, "select * from practice_submission where id=? limit 1", id)
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
	_, err := gosql.Exec(sql, total, ac, wa, re, tle, mle, ce, ole, pbid)
	return err
}

func (Practice) SetISE(psmid int64) error {
	_, err := gosql.Exec("update ojo.practice_submission set flag='ISE' where id=? limit 1", psmid)
	if err != nil {
		log.Warn("error:%v", err)
	}
	return err
}

func (Practice) UpdateFlagScoreMsg(psmid int64, score int, flag, errorMsg string) error {
	var sql = ` update ojo.practice_submission set 
                flag =?,
                total_score = ?,
                error_msg=?
        where id = ?`
	_, err := gosql.Exec(sql, flag, score, errorMsg, psmid)
	return err
}

func (Practice) InsertCaseRes(psmid, uid int64, tc *dto.TestCase) error {
	var sql = `  insert into ojo.practice_case_result
  (psmid,pcaseid,uid,flag,cpu_time,real_time,real_memory,real_output,error_output,spj_output,spj_error_output,score)
  				values(?,?,?,?,?,?,?,?,?,?,?,?)`
	_, err := gosql.Exec(sql, psmid, tc.Id, uid, tc.Flag, tc.ActualCpuTime,
		tc.ActualRealTime, tc.RealMemory, tc.RealOutput, tc.ErrorOutput,
		tc.SPJOutput, tc.SPJErrorOutput, tc.Score)
	return err
}

func (Practice) InsertStatistic(tx *gosql.DB, pbid int64) error {
	s := "insert into ojo.practice_statistic(pbid) values (?)"
	_, err := tx.Exec(s, pbid)
	return err
}

// -------------------Practice Submit-------------------------

// -------------------Practice Sub Count-------------------------

func (Practice) GetTodayCount() ([]dto.TodayCount, error) {
	var data []dto.TodayCount
	err := gosql.Select(&data, `SELECT
    DATE_FORMAT(submit_time, '%H') hour,
       count( * ) AS count
	FROM
    ojo.practice_submission
	WHERE
    submit_time between curdate()  and date_sub(curdate(),interval -1 day )
	GROUP BY
    hour
	ORDER BY
    hour;`)
	return data, err
}

func (Practice) GetWeekCount() (dto.WeekCount, error) {
	var data []dto.DayCount
	err := gosql.Select(&data, `SELECT
    DATE_FORMAT(submit_time, '%Y-%m-%d') day,
    count( * ) AS count
FROM
    ojo.practice_submission
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

func (Practice) GetMonthCount() (dto.MonthCount, error) {
	var data []dto.DayCount
	err := gosql.Select(&data, `SELECT
    DATE_FORMAT(submit_time, '%Y-%m-%d') day,
    count( * ) AS count
FROM
    ojo.practice_submission
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

// -------------------Practice Sub Count-------------------------

// -------------------Practice User Info-------------------------

func (Practice) GetUserSubmissionCount(uid int64) (int, error) {
	var count int
	s := `select count(*) 
			from ojo.practice_submission ps
			where ps.uid=?`
	err := gosql.Get(&count, s, uid)
	return count, err
}

func (Practice) GetUserACCount(uid int64) (int, error) {
	var count int
	s := `select count(*) as ac
			from (
			 select pid
			 from ojo.practice_submission
			 where flag='AC' and uid=?
			 group by pid
     		) as p`
	err := gosql.Get(&count, s, uid)
	return count, err
}

func (Practice) GetUserScore(uid int64) (int, error) {
	var count int
	s := `select IFNULL(sum(a.score),0) from (select max(ps.total_score) score
                                    from ojo.practice_submission ps
                                    where ps.uid=?
                                    group by ps.pid) a;`
	err := gosql.Get(&count, s, uid)
	return count, err
}

func (Practice) GetUserSolvedList(uid int64) ([]int, error) {
	var data []int
	s := `select distinct ps.pid
		from ojo.practice_submission ps
		where ps.uid=? and ps.flag='AC'`
	err := gosql.Select(&data, s, uid)
	return data, err
}

// -------------------Practice User Info-------------------------

// ----------------------Practice Rank---------------------------

var PctRankPageSize = 30

func (Practice) GetPctTop10() ([]dto.PctRank, error) {
	sql := `select count(*) as ac,uid
			from (
				select pid,uid
				from ojo.practice_submission
				where flag='AC'
				group by pid,uid
				) as p group by p.uid
			order by ac desc
			limit 10`
	var data []dto.PctRank
	err := gosql.Select(&data, sql)
	if err != nil {
		log.Warn("error:%v\n", err)
		return nil, err
	}
	err = user.SelectUserName(len(data), func(i int) (target int64) {
		return data[i].Uid
	}, func(i int, res string) {
		data[i].Username = res
	})
	return data, err
}

func (Practice) GetPctRank(form dto.RankForm) ([]dto.PctRank, error) {
	if form.Page < 1 {
		form.Page = 1
	}
	form.Page -= 1
	form.Limit = PctRankPageSize
	form.Offset = form.Page * PctRankPageSize
	sql := `select count(*) as ac,uid
			from (
				select pid,uid
				from ojo.practice_submission
				where flag='AC'
				group by pid,uid
				) as p group by p.uid
			order by ac desc 
			limit ?,?`
	var data []dto.PctRank
	err := gosql.Select(&data, sql, form.Offset, form.Limit)
	if err != nil {
		log.Warn("error:%v\n", err)
		return nil, err
	}
	err = user.SelectUserNameAndSig(len(data), func(i int) (target int64) {
		return data[i].Uid
	}, func(i int, res *dto.UsernameAndSig) {
		data[i].Username = res.Username
		data[i].Signature = res.Signature
	})
	return data, err
}

func (Practice) GetPctRankCount() (int, error) {
	sql := `select count(distinct uid) as total
			from ojo.practice_submission`
	var count int
	err := gosql.Get(&count, sql)
	return count, err
}

// ----------------------Practice Rank---------------------------
