package session

import (
	"encoding/gob"
	"errors"
	"github.com/afanke/OJO/WebServer/dto"
	"github.com/afanke/OJO/utils/log"
	"github.com/kataras/iris"
	"os"
	"time"
)

type Session struct {
	Data map[string]interface{}
}

func init() {
	gob.Register(dto.User{})
	LoadSession()
	go func() {
		for {
			select {
			case <-time.Tick(time.Minute):
				SaveSession()
			}
		}
	}()
}

var Pool = map[string]Session{}

func GetSession(c iris.Context) (s Session, err error) {
	cookie := c.GetCookie("GOGONEWWORLD")
	if cookie == "" {
		return Session{}, errors.New("miss cookie or cookie not correct")
	}
	addr := c.RemoteAddr()
	if s, ok := Pool[cookie+addr]; ok {
		return s, nil
	} else {
		s = Session{Data: map[string]interface{}{}}
		Pool[cookie+addr] = s
		return s, nil
	}
}

func (s Session) Get(str string) interface{} {
	return s.Data[str]
}

func (s Session) Set(str string, i interface{}) {
	s.Data[str] = i
}

func (s Session) Remove(str string) {
	delete(s.Data, str)
}

func SaveSession() {
	log.Info("start to save session")
	file, err := os.OpenFile("./config/session", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		log.Error("failed to save session:%v", err)
		return
	}
	err = gob.NewEncoder(file).Encode(Pool)
	if err != nil {
		log.Error("failed to save session:%v", err)
		return
	}
	log.Info("success to save session")
}

func LoadSession() {
	file, err := os.OpenFile("./config/session", os.O_RDONLY, 0666)
	if err != nil {
		log.Error("failed to load session:%v", err)
		return
	}
	err = gob.NewDecoder(file).Decode(&Pool)
	if err != nil {
		log.Error("failed to load session:%v", err)
		return
	}
	log.Info("success to load session")
}
