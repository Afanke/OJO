package jsp

// judge server pool
import (
	"encoding/json"
	"errors"
	"github.com/afanke/OJO/WebServer/dto"
	"github.com/afanke/OJO/utils/log"
	"github.com/ilibs/gosql/v2"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var (
	jsp     []dto.JudgeServer
	lens    int
	current int
	count   int
	lock    sync.RWMutex
)

func init() {
	initJSP()
	TouchJSP()
	go func() {
		for {
			<-time.After(5 * time.Second)
			TouchJSP()
		}
	}()
}

func GetAddr() (string, error) {
	lock.RLock()
	defer lock.RUnlock()
	en := 0
	for i, j := 0, len(jsp); i < j; i++ {
		if jsp[i].Enabled {
			en++
		}
	}
	if en == 0 {
		return "", errors.New("now jsp server available, please wait a minute or contact with the admin")
	}
	if count > jsp[current].Weight {
		current++
		current = current % lens
		count = 0
	} else {
		count++
	}
	return jsp[current].Address + ":" + strconv.Itoa(jsp[current].Port), nil
}

func initJSP() {
	var js []dto.JudgeServer
	err := gosql.Select(&js, "select id, name, address, port, weight, enabled from ojo.judge_server")
	if err != nil {
		log.Fatal("init jsp server failed:%v", err)
		return
	}
	lock.Lock()
	defer lock.Unlock()
	jsp = js
	lens = len(js)
}

func UpdateJSP() error {
	var js []dto.JudgeServer
	err := gosql.Select(&js, "select id, name, address, port, weight, enabled from ojo.judge_server")
	if err != nil {
		log.Error("update jsp server failed:%v", err)
		return err
	}
	lock.Lock()
	defer lock.Unlock()
	jsp = js
	lens = len(js)
	current = 0
	count = 0
	return nil
}

func TouchJSP() {
	lock.RLock()
	defer lock.RUnlock()
	log.Debug("%v", len(jsp))
	for i, j := 0, len(jsp); i < j; i++ {
		log.Debug("%v", i)
		k := i
		if jsp[k].Enabled {
			go func() {
				client := &http.Client{
					Timeout: 1 * time.Second,
				}
				res, err := client.Get("http://" + jsp[k].Address + ":" + strconv.Itoa(jsp[k].Port) + "/touch")
				if err != nil {
					log.Error("error:%v", err)
					jsp[k].Status = false
					return
				}
				defer res.Body.Close()
				body, err := ioutil.ReadAll(res.Body)
				if err != nil {
					log.Error("error:%v", err)
					jsp[k].Status = false
					return
				}
				var rest dto.Res
				err = json.Unmarshal(body, &res)
				if err != nil {
					log.Error("error:%v", err)
					jsp[k].Status = false
					return
				}
				if rest.Error != "" {
					log.Error("error:%v", rest.Error)
					jsp[k].Status = false
					return
				}
				jsp[k].Status = true
			}()
		}
	}
}

func GetAllInfo() []dto.JudgeServer {
	lock.RLock()
	defer lock.RUnlock()
	return jsp
}
