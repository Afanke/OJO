package session

import (
	"encoding/gob"
	"encoding/json"
	"errors"
	"github.com/gogotime/OJO/utils/log"
	"github.com/kataras/iris/v12"
	"io/ioutil"
	"os"
	"sync"
	"time"
)

type Session struct {
	Data map[string]interface{}
	Time int64
}

type Config struct {
	CleanCycle int64
	SaveCycle  int64
	MaxAge     int64
}

type SConfig struct {
	Config Config `json:"Session"`
}

var cfg Config
var Pool = map[string]Session{}
var PoolLock sync.RWMutex

func init() {
	gob.Register(sync.RWMutex{})
	LoadSession()
	err := LoadConfig()
	if err != nil {
		log.Error("use default config")
		cfg.SaveCycle = 15000
		cfg.CleanCycle = 15000
		cfg.MaxAge = 15000
		return
	}
	go regularTask()
}

func LoadConfig() error {
	file, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		log.Error("%v", err)
		return err
	}
	var scfg SConfig
	err = json.Unmarshal(file, &scfg)
	if err != nil {
		log.Error("%v", err)
		return err
	}
	cfg = scfg.Config
	if cfg.CleanCycle < 60 || cfg.SaveCycle < 60 {
		log.Error("session config params not permitted")
		return errors.New("")
	}
	log.Info("success to load session config")
	return nil
}

func GetSession(c iris.Context) (s *Session, err error) {
	cookie := c.GetCookie("GOGONEWWORLD")
	if cookie == "" {
		return nil, errors.New("miss cookie or cookie not correct")
	}
	addr := c.RemoteAddr()
	PoolLock.RLock()
	defer PoolLock.RUnlock()
	if s, ok := Pool[cookie+addr]; ok {
		return &s, nil
	} else {
		s = Session{Data: map[string]interface{}{}, Time: time.Now().Unix()}
		Pool[cookie+addr] = s
		return &s, nil
	}
}

func GetSessionByInt64(str string, id int64) (Session, error) {
	for k := range Pool {
		i, ok := Pool[k].Get(str).(int64)
		if ok {
			if i == id {
				s := Pool[k]
				return s, nil
			}
		}
	}
	return Session{}, errors.New("now such session")
}

func Get(c iris.Context, key string) (interface{}, error) {
	s, err := GetSession(c)
	if err != nil {
		return nil, err
	}
	return s.Get(key), nil
}

func Set(c iris.Context, key string, value interface{}) error {
	s, err := GetSession(c)
	if err != nil {
		return err
	}
	s.Set(key, value)
	return nil
}

func SetInt(c iris.Context, key string, value int) error {
	s, err := GetSession(c)
	if err != nil {
		return err
	}
	s.Set(key, value)
	return nil
}

func GetInt(c iris.Context, key string) (int, error) {
	s, err := GetSession(c)
	if err != nil {
		return 0, err
	}
	i, ok := s.Get(key).(int)
	if !ok {
		return 0, errors.New("can't convert interface to int")
	}
	return i, err
}

func SetInt64(c iris.Context, key string, value int64) error {
	s, err := GetSession(c)
	if err != nil {
		return err
	}
	s.Set(key, value)
	return nil
}

func GetInt64(c iris.Context, key string) (int64, error) {
	s, err := GetSession(c)
	if err != nil {
		return 0, err
	}
	i, ok := s.Get(key).(int64)
	if !ok {
		return 0, errors.New("can't convert interface to int64")
	}
	return i, err
}

func (s Session) Get(str string) interface{} {
	return s.Data[str]
}

func (s Session) Set(str string, i interface{}) {
	s.Time = time.Now().Unix()
	s.Data[str] = i
}

func (s Session) Remove(str string) {
	delete(s.Data, str)
}

func DelByInt64(str string, id int64) {
	PoolLock.Lock()
	defer PoolLock.Unlock()
	for k := range Pool {
		i, ok := Pool[k].Get(str).(int64)
		if ok {
			if i == id {
				delete(Pool, k)
			}
		}
	}
}

func regularTask() {
	for {
		select {
		case <-time.Tick(time.Duration(cfg.CleanCycle) * time.Second):
			CleanPool()
		case <-time.Tick(time.Duration(cfg.SaveCycle) * time.Second):
			SaveSession()
		}
	}
}

func CleanPool() {
	PoolLock.Lock()
	defer PoolLock.Unlock()
	now := time.Now().Unix()
	t := 0
	d := 0
	for k := range Pool {
		t++
		if now-Pool[k].Time > cfg.MaxAge {
			delete(Pool, k)
			d++
		}
	}
	log.Info("clean session pool: all:%d save:%d del:%d", t, t-d, d)
}

func SaveSession() {
	PoolLock.Lock()
	defer PoolLock.Unlock()
	file, err := os.OpenFile("./config/session.gob", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
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
	PoolLock.Lock()
	defer PoolLock.Unlock()
	file, err := os.OpenFile("./config/session.gob", os.O_RDONLY, 0666)
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
