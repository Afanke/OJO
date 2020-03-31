package db

import (
	"database/sql"
	"github.com/afanke/OJO/WebServer/dto"
	"github.com/afanke/OJO/utils/log"
)

type Problem struct {
}

var pb Problem

var ProblemPageSize = 10

func (Problem) GetPbCase(pid int64) ([]dto.ProblemCase, error) {
	var res []dto.ProblemCase
	err := db.Select(&res, "select * from problem_case where pid=?", pid)
	return res, err
}

func (Problem) GetAll(form *dto.ProblemForm) ([]dto.ProblemBrief, error) {
	if form.Page < 1 {
		form.Page = 1
	}
	form.Page -= 1
	form.Limit = ProblemPageSize
	form.Offset = form.Page * ProblemPageSize
	var s = `select id,ref,cid,title, difficulty,create_time,last_update_time,visible
			from ojo.problem  where 1=1 `
	if form.Keywords != "" {
		s += "and title like concat('%',:keywords,'%') "
	}
	if form.Difficulty != "" {
		s += "and difficulty=:difficulty "
	}
	s += " order by id desc limit :offset, :limit"
	rows, err := db.NamedQuery(s, &form)
	if err != nil {
		log.Warn("error:%v", err)
		return nil, err
	}
	var rest = make([]dto.ProblemBrief, 0, form.Limit)
	for rows.Next() {
		var res dto.ProblemBrief
		err := rows.StructScan(&res)
		if err != nil {
			log.Warn("error:%v", err)
			return []dto.ProblemBrief{}, err
		}
		tag, err := pb.GetProblemTag(res.Id)
		if err != nil {
			log.Warn("error:%v", err)
			return []dto.ProblemBrief{}, err
		}
		res.Tags = tag
		rest = append(rest, res)
	}
	return rest, nil
}

func (Problem) GetCount(form *dto.ProblemForm) (int, error) {
	var s = `select count(*)
			from ojo.problem  where 1=1 `
	if form.Keywords != "" {
		s += "and title like concat('%',:keywords,'%') "
	}
	if form.Difficulty != "" {
		s += "and difficulty=:difficulty "
	}
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

func (Problem) GetProblem(id int64) (*dto.Problem, error) {
	var res dto.Problem
	err := db.Get(&res, `select * from ojo.problem p where p.id=? limit 1`, id)
	return &res, err
}

func (Problem) GetProblemTag(pbid int64) ([]dto.Tag, error) {
	var s = `select t.id,t.name from tag t,problem_tag pt 
			where pt.pid=? and pt.tid=t.id`
	var tags []dto.Tag
	err := db.Select(&tags, s, pbid)
	return tags, err
}

func (Problem) GetLanguage(pbid int64) ([]dto.Language, error) {
	var s = `select l.id,l.name from language l,problem_language pl 
			where pl.pid=? and pl.lid=l.id`
	var languages []dto.Language
	err := db.Select(&languages, s, pbid)
	return languages, err
}

func (Problem) GetSample(pbid int64) ([]dto.ProblemSample, error) {
	var s = `select id,pid, input, output from problem_sample where pid=?`
	var samples []dto.ProblemSample
	err := db.Select(&samples, s, pbid)
	return samples, err
}

func (Problem) GetCreatorName(creatorId int64) (string, error) {
	var s string
	err := db.Get(&s, "select username from ojo.user  where id=? limit 1", creatorId)
	return s, err
}

func (Problem) GetName(pbid int64) (string, error) {
	var s string
	err := db.Get(&s, "select title from ojo.problem where id=? limit 1", pbid)
	return s, err
}

func (Problem) GetAllTags() ([]dto.Tag, error) {
	var s = `select id, name from tag `
	var tags []dto.Tag
	err := db.Select(&tags, s)
	return tags, err
}

func (Problem) InsertProblem(p *dto.Problem) error {
	var s = `insert into ojo.problem(cid,
                        ref,
                        title,
                        description,
                        input_description,
                        output_description,
                        hint, create_time,
                        last_update_time,
                        cpu_time_limit,
                        memory_limit,
                        difficulty,
                        real_time_limit,
                        source,
                        visible) VALUES(?,?,?,?,?,?,?,now(),now(),?,?,?,?,?,?) `
	tx, err := db.Begin()
	if err != nil {
		log.Warn("%v", err)
		return err
	}
	res, err := tx.Exec(s, p.Cid, p.Ref, p.Title, p.Description, p.InputDescription,
		p.OutputDescription, p.Hint, p.CpuTimeLimit, p.MemoryLimit, p.Difficulty, p.RealTimeLimit, p.Source, p.Visible)
	if err != nil {
		log.Warn("%v", err)
		err2 := tx.Rollback()
		if err2 != nil {
			log.Warn("%v", err2)
		}
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Warn("%v", err)
		err2 := tx.Rollback()
		if err2 != nil {
			log.Warn("%v", err2)
		}
		return err
	}
	for i, j := 0, len(p.ProblemCase); i < j; i++ {
		p.ProblemCase[i].Pid = id
		err := pb.InsertProblemCase(tx, &p.ProblemCase[i])
		if err != nil {
			log.Warn("%v", err)
			err2 := tx.Rollback()
			if err2 != nil {
				log.Warn("%v", err2)
			}
			return err
		}
	}
	for i, j := 0, len(p.Language); i < j; i++ {
		err := pb.InsertProblemLanguage(tx, id, p.Language[i].Id)
		if err != nil {
			log.Warn("%v", err)
			err2 := tx.Rollback()
			if err2 != nil {
				log.Warn("%v", err2)
			}
			return err
		}
	}
	for i, j := 0, len(p.Sample); i < j; i++ {
		p.Sample[i].Pid = id
		err := pb.InsertProblemSample(tx, &p.Sample[i])
		if err != nil {
			log.Warn("%v", err)
			err2 := tx.Rollback()
			if err2 != nil {
				log.Warn("%v", err2)
			}
			return err
		}
	}
	for i, j := 0, len(p.Tag); i < j; i++ {
		err := pb.InsertProblemTag(tx, id, p.Tag[i].Id)
		if err != nil {
			log.Warn("%v", err)
			err2 := tx.Rollback()
			if err2 != nil {
				log.Warn("%v", err2)
			}
			return err
		}
	}
	err = pt.InsertStatistic(tx, id)
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
	return nil
}

func (Problem) InsertProblemCase(tx *sql.Tx, pc *dto.ProblemCase) error {
	var s = "insert into ojo.problem_case(pid, input, output,score) VALUES (?,?,?,?)"
	_, err := tx.Exec(s, pc.Pid, pc.Input, pc.Output, pc.Score)
	return err
}

func (Problem) InsertProblemLanguage(tx *sql.Tx, pid, lid int64) error {
	var s = "insert into ojo.problem_language(pid, lid) VALUES (?,?)"
	_, err := tx.Exec(s, pid, lid)
	return err
}

func (Problem) InsertProblemSample(tx *sql.Tx, ps *dto.ProblemSample) error {
	var s = "insert into ojo.problem_sample(pid, input, output) VALUES (?,?,?)"
	_, err := tx.Exec(s, ps.Pid, ps.Input, ps.Output)
	return err
}

func (Problem) InsertProblemTag(tx *sql.Tx, pid, tid int64) error {
	var s = "insert into ojo.problem_tag(tid, pid) VALUES (?,?)"
	_, err := tx.Exec(s, tid, pid)
	return err
}

func (Problem) SetVisibleTrue(id int64) error {
	s := "update ojo.problem set visible=true where id=? limit 1"
	_, err := db.Exec(s, id)
	return err
}

func (Problem) SetVisibleFalse(id int64) error {
	s := "update ojo.problem set visible=false where id=? limit 1"
	_, err := db.Exec(s, id)
	return err
}

func (Problem) GetCreatorId(pid int64) (int64, error) {
	var cid int64
	err := db.Get(&cid, "select cid from ojo.problem where id=?", pid)
	return cid, err
}
