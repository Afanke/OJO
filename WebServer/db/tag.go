package db

import (
	"github.com/afanke/OJO/WebServer/dto"
	"github.com/afanke/OJO/utils/log"
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
	s := `select id, name, cid, visible, communal, create_time, last_update_time from tag `
	s += " where 1=1 "
	if form.Keywords != "" {
		s += "and name like concat('%',:keywords,'%') "
	}
	if form.Mine {
		s += "and cid=:cid"
	}
	s += " order by id desc limit :offset, :limit"
	rows, err := db.NamedQuery(s, &form)
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
		name, err := pb.GetCreatorName(res.Cid)
		if err != nil {
			log.Warn("%v", err)
			return nil, err
		}
		res.CreatorName = name
		rest = append(rest, res)
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
		err := db.Get(&count, s, form.Cid)
		return count, err
	}
	err := db.Get(&count, s)
	return count, err
}

func (Tag) GetBriefByPid(pbid int64) ([]dto.TagBrief, error) {
	var s = `select t.id,t.name from tag t,problem_tag pt 
				where pt.pid=? and pt.tid=t.id`
	var tags []dto.TagBrief
	err := db.Select(&tags, s, pbid)
	return tags, err
}

func (Tag) GetAllVisible() ([]dto.TagBrief, error) {
	var s = `select id, name,cid from tag where visible=1`
	var tags []dto.TagBrief
	err := db.Select(&tags, s)
	return tags, err
}

func (Tag) GetAllCommunal() ([]dto.TagBrief, error) {
	var s = `select id, name,cid from tag where communal=1`
	var tags []dto.TagBrief
	err := db.Select(&tags, s)
	return tags, err
}

func (Tag) GetTagCreatorId(tid int64) (int64, error) {
	var cid int64
	err := db.Get(&cid, "select cid from ojo.tag where id=?", tid)
	return cid, err
}
