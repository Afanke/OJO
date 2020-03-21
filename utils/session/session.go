package session

import (
	"errors"
	"github.com/kataras/iris"
)

type Session struct {
	data map[string]interface{}
}

var pool = map[string]Session{}

func GetSession(c iris.Context) (s Session, err error) {
	cookie := c.GetCookie("GOGONEWWORLD")
	if cookie == "" {
		return Session{}, errors.New("miss cookie or cookie not correct")
	}
	addr := c.RemoteAddr()
	if s, ok := pool[cookie+addr]; ok {
		return s, nil
	} else {
		s = Session{data: map[string]interface{}{}}
		pool[cookie+addr] = s
		return s, nil
	}
}

func (s Session) Get(str string) interface{} {
	return s.data[str]
}

func (s Session) Set(str string, i interface{}) {
	s.data[str] = i
}

func (s Session) Remove(str string) {
	delete(s.data, str)
}
