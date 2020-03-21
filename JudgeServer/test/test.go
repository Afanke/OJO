package main

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
)

var db *sqlx.DB

func init() {
	var err error
	db, err = sqlx.Open("mysql", "root:123123@tcp(127.0.0.1:3306)/ojo?charset=utf8mb4")
	if err != nil {
		fmt.Printf("error:%v", err)
		os.Exit(-1)
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("error:%v", err)
		os.Exit(-1)
	}
}

type S struct {
	Username string `db:"username"`
	Password string `db:"password"`
}
type User struct {
	Id            int    `json:"id" db:"id"`
	Username      string `json:"username" db:"username"`
	Password      string `json:"password" db:"password"`
	Email         string `json:"email" db:"email"`
	CreateTime    string `json:"createTime" db:"create_time"`
	LastLoginTime string `json:"lastLoginTime" db:"last_login_time"`
}

func hello() (string, error) {
	return "sadasd", errors.New("sadasd")
}

func nohello() (string, error) {
	return "sdsa", nil
}

func main() {
	_, err := hello()
	s, err := nohello()
	fmt.Println(err)
	fmt.Println(s)
}
