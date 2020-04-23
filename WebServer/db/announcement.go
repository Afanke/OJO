package db

import (
	_ "database/sql"
	"github.com/afanke/OJO/WebServer/dto"
	"github.com/afanke/OJO/utils/log"
	"github.com/ilibs/gosql/v2"
)

type Announcement struct {
}

var AnnouncementPageSize = 10

var anno Announcement

func (Announcement) GetAll(form *dto.AnnouncementForm) ([]dto.Announcement, error) {
	if form.Page < 1 {
		form.Page = 1
	}
	form.Page -= 1
	form.Limit = AnnouncementPageSize
	form.Offset = form.Page * AnnouncementPageSize
	var s = `select id, title, visible, cid, create_time, last_update_time
			from ojo.announcement  where 1=1 `
	if form.Keywords != "" {
		s += "and title like concat('%',:keywords,'%') "
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
	var rest = make([]dto.Announcement, 0, form.Limit)
	for rows.Next() {
		var res dto.Announcement
		err := rows.StructScan(&res)
		if err != nil {
			log.Warn("error:%v", err)
			return nil, err
		}
		rest = append(rest, res)
	}
	err = anno.SelectCreatorName(len(rest), func(i int) int64 {
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

func (Announcement) GetCount(form *dto.AnnouncementForm) (int, error) {
	var s = `select count(*)
			from ojo.announcement  where 1=1 `
	if form.Keywords != "" {
		s += "and title like concat('%',:keywords,'%') "
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

func (Announcement) GetDetail(id int64) (*dto.Announcement, error) {
	var detail dto.Announcement
	err := gosql.Get(&detail, `select title,content from ojo.announcement a where a.id=? limit 1`, id)
	if err != nil {
		log.Warn("error:%v", err)
		return nil, err
	}
	return &detail, err
}

func (Announcement) SelectCreatorName(lens int, getId func(i int) (target int64), setName func(i int, res string)) error {
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

func (Announcement) InsertAnnouncement(a *dto.Announcement) error {
	var s = `insert into ojo.announcement(title,
                             content,
                             visible,
                             cid,
                             create_time,
                             last_update_time) VALUES(?,?,?,?,now(),now()) `
	_, err := gosql.Exec(s, a.Title, a.Content, a.Visible, a.Cid, a.CreateTime, a.LastUpdateTime)
	if err != nil {
		log.Warn("%v", err)
		return err
	}
	return nil
}

func (Announcement) UpdateAnnouncement(a *dto.Announcement) error {
	var s = `update ojo.announcement set title=?,
                        content=?,
                        visible=?,                  
                        last_update_time=now()
                   where id=? `
	_, err := gosql.Exec(s, a.Title, a.Content, a.Visible, a.Id)
	if err != nil {
		log.Warn("%v", err)
		return err
	}
	return nil
}

func (Announcement) DeleteAnnouncement(id int64) error {
	var s = "delete from ojo.announcement where id=? limit 1"
	_, err := gosql.Exec(s, id)
	return err
}

func (Announcement) SetVisibleTrue(id int64) error {
	s := "update ojo.announcement set visible=true where id=? limit 1"
	_, err := gosql.Exec(s, id)
	return err
}

func (Announcement) SetVisibleFalse(id int64) error {
	s := "update ojo.announcement set visible=false where id=? limit 1"
	_, err := gosql.Exec(s, id)
	return err
}

func (Announcement) GetCreatorId(id int64) (int64, error) {
	var cid int64
	err := gosql.Get(&cid, "select cid from ojo.announcement where id=?", id)
	return cid, err
}
