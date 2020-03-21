package db

import (
	"fmt"
	"github.com/afanke/OJO/WebServer/dto"
)

type User struct{}

var user User

func (User) Query(username string, password string) (dto.User, error) {
	var user dto.User
	err := db.Get(&user, "select id,username,email,create_time,last_login_time from `user`  "+
		"where username=? and password=? limit 1", username, password)
	fmt.Println(err)
	return user, err
}

func (User) GetName(id int) (string, error) {
	var s string
	err := db.Get(&s, "select username from user where id=?", id)
	return s, err
}

func (User) Insert(form *dto.RegisterForm) error {
	_, err := db.Exec("insert into `user`(username,password,email,create_time,last_login_time) values (?,?,?,now(),now())",
		form.Username, form.Password, form.Email)
	return err
}
