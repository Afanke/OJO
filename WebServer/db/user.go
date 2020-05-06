package db

import (
	"fmt"
	"github.com/afanke/OJO/WebServer/dto"
	"github.com/afanke/OJO/utils/log"
	"github.com/ilibs/gosql/v2"
)

type User struct{}

var user User
var UserPageSize = 10

func (User) Query(username string, password string) (dto.UserToken, error) {
	var user dto.UserToken
	err := gosql.Get(&user, `select id,username,type,icon_path,real_name,enabled from ojo.user  
		where username=? and password=?  limit 1`, username, password)
	return user, err
}

func (User) UpdateLoginTime(id int64) error {
	_, err := gosql.Exec("update ojo.user set last_login_time=now() where id=?", id)
	return err
}

func (User) GetName(id int64) (string, error) {
	var s string
	err := gosql.Get(&s, "select username from user where id=?", id)
	return s, err
}

func (User) Insert(form *dto.RegisterForm) error {
	_, err := gosql.Exec("insert into `user`(username,password,email,create_time,last_login_time) values (?,?,?,now(),now())",
		form.Username, form.Password, form.Email)
	return err
}

func (User) GetAll(form *dto.UserForm) ([]dto.UserBrief, error) {
	if form.Page < 1 {
		form.Page = 1
	}
	form.Page -= 1
	form.Limit = UserPageSize
	form.Offset = form.Page * UserPageSize
	var s = `select id, username, email, create_time, last_login_time, type, enabled, real_name
			from ojo.user where 1=1 `
	if form.Keywords != "" {
		s += "and (username like concat('%',:keywords,'%') or real_name like concat('%',:keywords,'%'))"
	}
	if form.Type != 0 {
		s += "and type=:type "
	}
	s += "  order by id desc limit :offset, :limit"
	rows, err := gosql.Sqlx().NamedQuery(s, &form)
	if err != nil {
		log.Warn("error:%v", err)
		return nil, err
	}
	var rest = make([]dto.UserBrief, 0, form.Limit)
	for rows.Next() {
		var res dto.UserBrief
		err := rows.StructScan(&res)
		if err != nil {
			log.Warn("error:%v", err)
			return nil, err
		}
		rest = append(rest, res)
	}
	return rest, nil
}

func (User) GetCount(form *dto.UserForm) (int, error) {
	var s = `select count(*)
			from ojo.user where 1=1 `
	if form.Keywords != "" {
		s += "and (username like concat('%',:keywords,'%') or real_name like concat('%',:keywords,'%'))"
	}
	if form.Type != 0 {
		s += "and type=:type "
	}
	rows, err := gosql.Sqlx().NamedQuery(s, &form)
	if err != nil {
		log.Warn("error:%v", err)
		return 0, err
	}
	var count int
	rows.Next()
	err = rows.Scan(&count)
	return count, err
}

func (User) GetDetail(id int64) (*dto.UserDetail, error) {
	var s dto.UserDetail
	err := gosql.Get(&s, "select id, username, email, type, icon_path, real_name, school, signature, blog, major, github from user where id=?", id)
	return &s, err
}

func (User) UpdateDetail(form *dto.UserDetail2) error {
	s := `update ojo.user set username=:username,email=:email,
 			type=:type,real_name=:real_name,signature=:signature,school=:school,
 			icon_path=:icon_path,blog=:blog,major=:major,github=:github`
	if form.Password != "" {
		s += ",password=:password"
	}
	s += " where id=:id"
	_, err := gosql.Sqlx().NamedExec(s, form)
	return err
}

func (User) Enable(id int64) error {
	s := `update ojo.user set enabled=1 where id=?`
	_, err := gosql.Exec(s, id)
	return err
}

func (User) Disable(id int64) error {
	s := `update ojo.user set enabled=0 where id=?`
	_, err := gosql.Exec(s, id)
	return err
}

func (User) SelectUserName(lens int, getId func(i int) (target int64), setName func(i int, res string)) error {
	if lens <= 0 {
		return nil
	}
	ids := make([]int64, 0, lens)
	for i := 0; i < lens; i++ {
		ids = append(ids, getId(i))
	}
	var s []dto.Username
	err := gosql.Select(&s, "select id,username from ojo.user where id in (?) ", ids)
	if err != nil {
		return err
	}
	fmt.Println(s)
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

func (User) SelectUserNameAndSig(lens int, getId func(i int) (target int64), setData func(i int, res *dto.UsernameAndSig)) error {
	if lens <= 0 {
		return nil
	}
	ids := make([]int64, 0, lens)
	for i := 0; i < lens; i++ {
		ids = append(ids, getId(i))
	}
	var s []dto.UsernameAndSig
	err := gosql.Select(&s, "select id,username,signature from ojo.user where id in (?) ", ids)
	if err != nil {
		return err
	}
	fmt.Println(s)
	for i := 0; i < lens; i++ {
		for j, k := 0, len(s); j < k; j++ {
			if getId(i) == s[j].Id {
				setData(i, &s[j])
				break
			}
		}
	}
	return nil
}
