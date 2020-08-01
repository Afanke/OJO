package ctrl

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/afanke/OJO/WebServer/dto"
	jsp "github.com/afanke/OJO/WebServer/judge"
	"github.com/kataras/iris/v12"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type Problem struct {
}

var pb Problem

func (Problem) LocalTest(c iris.Context) {
	var form dto.JudgeForm
	err := c.ReadJSON(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	_, err = isAdmin(c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	addr, err := jsp.GetAddr()
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	client := &http.Client{
		Timeout: time.Duration(form.MaxRealTime*form.SPJMp*form.CompMp) * time.Second * 5,
	}
	buff, err := json.Marshal(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	res, err := client.Post("http://"+addr+"/judge", "application/json", bytes.NewBuffer(buff))
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: form})
}

func (Problem) AddProblem(c iris.Context) {
	var p dto.Problem
	err := c.ReadJSON(&p)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	userId, err := isAdmin(c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	p.Cid = userId
	err = pbdb.InsertProblem(&p)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: "save successfully"})
}

func (Problem) UpdateProblem(c iris.Context) {
	var p dto.Problem
	err := c.ReadJSON(&p)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	userId, err := isAdmin(c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	p.Cid = userId
	err = pbdb.UpdateProblem(&p)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: "save successfully"})
}

func (Problem) DeleteProblem(c iris.Context) {
	var id dto.Id
	err := c.ReadJSON(&id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	err = pb.isCreator(c, id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	count, err := pbdb.IsDepended(id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	if count != 0 {
		c.JSON(&dto.Res{Error: errors.New("can't delete problem: the problem is depended by " + strconv.Itoa(count) + " contests").Error(), Data: nil})
		return
	}
	err = pbdb.DeleteProblem(id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: "delete successfully"})
}

func (Problem) GetAll(c iris.Context) {
	var form dto.ProblemForm
	err := c.ReadJSON(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	userId, err := isAdmin(c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	form.Cid = userId
	data, err := pbdb.GetAll(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: data})
}

func (Problem) GetAllShared(c iris.Context) {
	var form dto.ProblemForm
	err := c.ReadJSON(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	userId, err := isAdmin(c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	form.Cid = userId
	data, err := pbdb.GetAll(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: data})
}

func (Problem) GetSharedCount(c iris.Context) {
	userId, err := isAdmin(c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	data, err := pbdb.GetSharedCount(userId)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: data})
}

func (Problem) GetDetail(c iris.Context) {
	var id dto.Id
	err := c.ReadJSON(&id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	err = pb.isPermitted(c, id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	data, err := pbdb.GetDetail(id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: data})
}

func (Problem) GetCount(c iris.Context) {
	var form dto.ProblemForm
	err := c.ReadJSON(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	userId, err := isAdmin(c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	form.Cid = userId
	tags, err := pbdb.GetCount(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: tags})
}

func (Problem) SetVisibleTrue(c iris.Context) {
	var id dto.Id
	err := c.ReadJSON(&id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	err = pb.isPermitted(c, id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	err = pbdb.SetVisibleTrue(id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: "set up successfully"})
}

func (Problem) SetVisibleFalse(c iris.Context) {
	var id dto.Id
	err := c.ReadJSON(&id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	err = pb.isPermitted(c, id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	err = pbdb.SetVisibleFalse(id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: "set up successfully"})
}

func (Problem) TryEdit(c iris.Context) {
	var id dto.Id
	err := c.ReadJSON(&id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	err = pb.isPermitted(c, id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: "ok"})
}

func (Problem) isCreator(c iris.Context, id int64) error {
	adminId, err := isSuperAdmin(c)
	if err == nil {
		return nil
	}
	creatorId, err := pbdb.GetCreatorId(id)
	if err != nil {
		return err
	}
	if adminId != creatorId {
		return errors.New("not allowed")
	}
	return nil
}

func (Problem) isPermitted(c iris.Context, id int64) error {
	err := pb.isCreator(c, id)
	if err == nil {
		return nil
	}
	err = pbdb.IsShared(id)
	return err
}
