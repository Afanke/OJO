package db

import (
	_ "database/sql"
	"errors"
	"github.com/afanke/OJO/WebServer/dto"
	"github.com/afanke/OJO/utils/log"
	"github.com/ilibs/gosql/v2"
)

type Problem struct {
}

var pb Problem

var ProblemPageSize = 10

func (Problem) GetCase(pid int64) ([]dto.ProblemCase, error) {
	var res []dto.ProblemCase
	err := gosql.Select(&res, "select * from problem_case where pid=?", pid)
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
	if form.Mine {
		s += " and cid=:cid "
	}
	s += " order by id desc limit :offset, :limit"
	rows, err := gosql.Sqlx().NamedQuery(s, &form)
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
			return nil, err
		}
		rest = append(rest, res)
	}
	err = pb.SelectCreatorName(len(rest), func(i int) int64 {
		return rest[i].Cid
	}, func(i int, res string) {
		rest[i].CreatorName = res
	})
	if err != nil {
		log.Warn("%v", err)
		return nil, err
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
	if form.Mine {
		s += " and cid=:cid "
	}
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

func (Problem) GetAllShared(form *dto.ProblemForm) ([]dto.ProblemBrief, error) {
	if form.Page < 1 {
		form.Page = 1
	}
	form.Page -= 1
	form.Limit = ProblemPageSize
	form.Offset = form.Page * ProblemPageSize
	var data []dto.ProblemBrief
	var s = `select id,ref,cid,title,difficulty,create_time,last_update_time,visible
			from ojo.problem  where shared=1 or cid=? order by id desc limit ?,? `
	err := gosql.Select(&data, s, form.Cid, form.Limit, form.Offset)
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

func (Problem) GetSharedCount(cid int64) (int, error) {
	var s = `select count(*)
			from ojo.problem  where shared=1 or cid=?  `
	var count int
	err := gosql.Get(&count, s, cid)
	return count, err
}

func (Problem) GetProblem(id int64) (*dto.Problem, error) {
	var res dto.Problem
	err := gosql.Get(&res, `select * from ojo.problem p where p.id=? limit 1`, id)
	return &res, err
}

func (Problem) GetDetail(id int64) (*dto.Problem, error) {
	var detail dto.Problem
	err := gosql.Get(&detail, `select id, cid,
       ref, title, description, input_description,
       output_description, hint, create_time,
       last_update_time, cpu_time_limit, memory_limit,
       difficulty, real_time_limit, source,
       visible from ojo.problem p where p.id=? limit 1`, id)
	if err != nil {
		log.Warn("error:%v", err)
		return nil, err
	}
	tags, err := tag.GetBriefByPid(id)
	if err != nil {
		log.Warn("error:%v", err)
		return nil, err
	}
	languages, err := pb.GetLanguage(id)
	if err != nil {
		log.Warn("error:%v", err)
		return nil, err
	}
	samples, err := pb.GetSample(id)
	if err != nil {
		log.Warn("error:%v", err)
		return nil, err
	}
	cases, err := pb.GetCase(id)
	if err != nil {
		log.Warn("error:%v", err)
		return nil, err
	}
	detail.Tag = tags
	detail.Language = languages
	detail.Sample = samples
	detail.ProblemCase = cases
	return &detail, err
}

func (Problem) GetLanguage(pbid int64) ([]dto.Language, error) {
	var s = `select l.id,l.name from language l,problem_language pl 
			where pl.pid=? and pl.lid=l.id`
	var languages []dto.Language
	err := gosql.Select(&languages, s, pbid)
	return languages, err
}

func (Problem) GetSample(pbid int64) ([]dto.ProblemSample, error) {
	var s = `select id,pid, input, output from problem_sample where pid=?`
	var samples []dto.ProblemSample
	err := gosql.Select(&samples, s, pbid)
	return samples, err
}

func (Problem) GetCreatorName(creatorId int64) (string, error) {
	var s string
	err := gosql.Get(&s, "select username from ojo.user where id=? ", creatorId)
	return s, err
}

func (Problem) SelectCreatorName(lens int, getId func(i int) (target int64), setName func(i int, res string)) error {
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

func (Problem) GetName(pbid int64) (string, error) {
	var s string
	err := gosql.Get(&s, "select title from ojo.problem where id=? limit 1", pbid)
	return s, err
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
	tx, err := gosql.Begin()
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

func (Problem) UpdateProblem(p *dto.Problem) error {
	var s = `update ojo.problem set cid=?,
                        ref=?,
                        title=?,
                        description=?,
                        input_description=?,
                        output_description=?,
                        hint=?,
                        last_update_time=now(),
                        cpu_time_limit=?,
                        memory_limit=?,
                        difficulty=?,
                        real_time_limit=?,
                        source=?,
                        visible=? where id=? `
	tx, err := gosql.Begin()
	if err != nil {
		log.Warn("%v", err)
		return err
	}
	_, err = tx.Exec(s, p.Cid, p.Ref, p.Title, p.Description, p.InputDescription,
		p.OutputDescription, p.Hint, p.CpuTimeLimit, p.MemoryLimit, p.Difficulty, p.RealTimeLimit, p.Source, p.Visible, p.Id)
	if err != nil {
		log.Warn("%v", err)
		err2 := tx.Rollback()
		if err2 != nil {
			log.Warn("%v", err2)
		}
		return err
	}
	err = pb.DeleteProblemCase(tx, p.Id)
	if err != nil {
		log.Warn("%v", err)
		err2 := tx.Rollback()
		if err2 != nil {
			log.Warn("%v", err2)
		}
		return err
	}
	err = pb.DeleteProblemLanguage(tx, p.Id)
	if err != nil {
		log.Warn("%v", err)
		err2 := tx.Rollback()
		if err2 != nil {
			log.Warn("%v", err2)
		}
		return err
	}
	err = pb.DeleteProblemSample(tx, p.Id)
	if err != nil {
		log.Warn("%v", err)
		err2 := tx.Rollback()
		if err2 != nil {
			log.Warn("%v", err2)
		}
		return err
	}
	err = pb.DeleteProblemTag(tx, p.Id)
	if err != nil {
		log.Warn("%v", err)
		err2 := tx.Rollback()
		if err2 != nil {
			log.Warn("%v", err2)
		}
		return err
	}
	for i, j := 0, len(p.ProblemCase); i < j; i++ {
		p.ProblemCase[i].Pid = p.Id
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
		err := pb.InsertProblemLanguage(tx, p.Id, p.Language[i].Id)
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
		p.Sample[i].Pid = p.Id
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
		err := pb.InsertProblemTag(tx, p.Id, p.Tag[i].Id)
		if err != nil {
			log.Warn("%v", err)
			err2 := tx.Rollback()
			if err2 != nil {
				log.Warn("%v", err2)
			}
			return err
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

func (Problem) DeleteProblem(pid int64) error {
	var s = `delete from ojo.problem where id=? limit 1`
	tx, err := gosql.Begin()
	if err != nil {
		log.Warn("%v", err)
		return err
	}
	_, err = tx.Exec(s, pid)
	if err != nil {
		log.Warn("%v", err)
		err2 := tx.Rollback()
		if err2 != nil {
			log.Warn("%v", err2)
		}
		return err
	}
	err = pb.DeleteProblemCase(tx, pid)
	if err != nil {
		log.Warn("%v", err)
		err2 := tx.Rollback()
		if err2 != nil {
			log.Warn("%v", err2)
		}
		return err
	}
	err = pb.DeleteProblemLanguage(tx, pid)
	if err != nil {
		log.Warn("%v", err)
		err2 := tx.Rollback()
		if err2 != nil {
			log.Warn("%v", err2)
		}
		return err
	}
	err = pb.DeleteProblemSample(tx, pid)
	if err != nil {
		log.Warn("%v", err)
		err2 := tx.Rollback()
		if err2 != nil {
			log.Warn("%v", err2)
		}
		return err
	}
	err = pb.DeleteProblemTag(tx, pid)
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

func (Problem) IsDepended(pid int64) (int, error) {
	var count int
	err := gosql.Get(&count, "select count(*) from ojo.contest_problem where pid=?", pid)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (Problem) DeleteProblemCase(tx *gosql.DB, pid int64) error {
	var s = "delete from ojo.problem_case where pid=?"
	_, err := tx.Exec(s, pid)
	return err
}

func (Problem) DeleteProblemLanguage(tx *gosql.DB, pid int64) error {
	var s = "delete from ojo.problem_language where pid=?"
	_, err := tx.Exec(s, pid)
	return err
}

func (Problem) DeleteProblemSample(tx *gosql.DB, pid int64) error {
	var s = "delete from ojo.problem_sample where pid=?"
	_, err := tx.Exec(s, pid)
	return err
}

func (Problem) DeleteProblemTag(tx *gosql.DB, pid int64) error {
	var s = "delete from ojo.problem_tag where pid=?"
	_, err := tx.Exec(s, pid)
	return err
}

func (Problem) InsertProblemCase(tx *gosql.DB, pc *dto.ProblemCase) error {
	var s = "insert into ojo.problem_case(pid, input, output,score) VALUES (?,?,?,?)"
	_, err := tx.Exec(s, pc.Pid, pc.Input, pc.Output, pc.Score)
	return err
}

func (Problem) InsertProblemLanguage(tx *gosql.DB, pid, lid int64) error {
	var s = "insert into ojo.problem_language(pid, lid) VALUES (?,?)"
	_, err := tx.Exec(s, pid, lid)
	return err
}

func (Problem) InsertProblemSample(tx *gosql.DB, ps *dto.ProblemSample) error {
	var s = "insert into ojo.problem_sample(pid, input, output) VALUES (?,?,?)"
	_, err := tx.Exec(s, ps.Pid, ps.Input, ps.Output)
	return err
}

func (Problem) InsertProblemTag(tx *gosql.DB, pid, tid int64) error {
	var s = "insert into ojo.problem_tag(tid, pid) VALUES (?,?)"
	_, err := tx.Exec(s, tid, pid)
	return err
}

func (Problem) SetVisibleTrue(id int64) error {
	s := "update ojo.problem set visible=true where id=? limit 1"
	_, err := gosql.Exec(s, id)
	return err
}

func (Problem) SetVisibleFalse(id int64) error {
	s := "update ojo.problem set visible=false where id=? limit 1"
	_, err := gosql.Exec(s, id)
	return err
}

func (Problem) IsShared(pid int64) error {
	var res bool
	err := gosql.Get(&res, "select shared from ojo.problem where id=?", pid)
	if err != nil {
		return err
	}
	if !res {
		return errors.New("problem isn't shared")
	}
	return nil
}

func (Problem) GetCreatorId(pid int64) (int64, error) {
	var cid int64
	err := gosql.Get(&cid, "select cid from ojo.problem where id=?", pid)
	return cid, err
}
