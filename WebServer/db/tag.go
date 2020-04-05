package db

import (
	"errors"
	"github.com/afanke/OJO/WebServer/dto"
	"github.com/afanke/OJO/utils/log"
	"github.com/ilibs/gosql/v2"
)

type Tag struct {
}

var tag Tag

var tagPageSize = 10

func (Tag) GetAll(form *dto.TagForm) ([]dto.Tag, error) {
	if form.Page < 1 {
		form.Page = 1
	}
	form.Page -= 1
	form.Limit = tagPageSize
	form.Offset = form.Page * tagPageSize
	s := `select id, name, cid, visible, shared, create_time, last_update_time from tag `
	s += " where 1=1 "
	if form.Keywords != "" {
		s += "and name like concat('%',:keywords,'%') "
	}
	if form.Mine {
		s += "and cid=:cid"
	}
	s += " order by id desc limit :offset, :limit"
	rows, err := gosql.Sqlx().NamedQuery(s, &form)
	if err != nil {
		log.Warn("error:%v", err)
		return nil, err
	}
	var rest = make([]dto.Tag, 0, form.Limit)
	for rows.Next() {
		var res dto.Tag
		err := rows.StructScan(&res)
		if err != nil {
			log.Warn("error:%v", err)
			return nil, err
		}
		rest = append(rest, res)
	}
	err = pb.SelectCreatorName(len(rest), func(i int) (target int64) {
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

func (Tag) GetCount(form *dto.TagForm) (int64, error) {
	if form.Page < 1 {
		form.Page = 1
	}
	form.Page -= 1
	form.Limit = tagPageSize
	form.Offset = form.Page * tagPageSize
	s := `select count(*) from tag `
	var count int64
	if form.Mine {
		s += " where cid=? "
		err := gosql.Get(&count, s, form.Cid)
		return count, err
	}
	err := gosql.Get(&count, s)
	return count, err
}

func (Tag) GetBriefByPid(pbid int64) ([]dto.TagBrief, error) {
	var s = `select t.id,t.name from tag t,problem_tag pt 
				where pt.pid=? and pt.tid=t.id`
	var tags []dto.TagBrief
	err := gosql.Select(&tags, s, pbid)
	return tags, err
}

func (Tag) GetAllVisible() ([]dto.TagBrief, error) {
	var s = `select id, name,cid from tag where visible=1`
	var tags []dto.TagBrief
	err := gosql.Select(&tags, s)
	return tags, err
}

func (Tag) GetAllCommunal() ([]dto.TagBrief, error) {
	var s = `select id, name,cid from tag where communal=1`
	var tags []dto.TagBrief
	err := gosql.Select(&tags, s)
	return tags, err
}

func (Tag) GetCreatorId(tid int64) (int64, error) {
	var cid int64
	err := gosql.Get(&cid, "select cid from ojo.tag where id=?", tid)
	return cid, err
}

func (Tag) SetVisibleTrue(id int64) error {
	s := "update ojo.tag set visible=true,last_update_time=NOW() where id=? limit 1"
	_, err := gosql.Exec(s, id)
	return err
}

func (Tag) SetVisibleFalse(id int64) error {
	s := "update ojo.tag set visible=false,last_update_time=NOW() where id=? limit 1"
	_, err := gosql.Exec(s, id)
	return err
}

func (Tag) SetSharedTrue(id int64) error {
	s := "update ojo.tag set shared=true,last_update_time=NOW() where id=? limit 1"
	_, err := gosql.Exec(s, id)
	return err
}

func (Tag) SetSharedFalse(id int64) error {
	s := "update ojo.tag set shared=false,last_update_time=NOW() where id=? limit 1"
	_, err := gosql.Exec(s, id)
	return err
}

func (Tag) InsertTag(t *dto.Tag) error {
	var s = `insert into ojo.tag(name, cid, visible, shared, create_time, last_update_time) VALUES(?,?,?,?,now(),now()) `
	_, err := gosql.Exec(s, t.Name, t.Cid, t.Visible, t.Shared)
	return err
}

func (Tag) UpdateTag(t *dto.TagBrief) error {
	var s = `update ojo.tag set name=?,last_update_time=now() where id=? limit 1`
	_, err := gosql.Exec(s, t.Name, t.Id)
	return err
}

func (Tag) IsShared(tid int64) error {
	var res int
	err := gosql.Get(&res, "select shared from ojo.tag where id=?", tid)
	if err != nil {
		return err
	}
	if res != 1 {
		return errors.New("not allowed")
	}
	return nil
}
