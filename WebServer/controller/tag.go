package ctrl

import (
	"errors"
	"github.com/gogotime/OJO/WebServer/db"
	"github.com/gogotime/OJO/WebServer/dto"
	"github.com/gogotime/OJO/utils/session"
	"github.com/kataras/iris/v12"
	"strconv"
	"time"
)

type Tag struct {
}

var tag Tag

var tagdb db.Tag

func (Tag) GetAll(c iris.Context) {
	var form dto.TagForm
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
	tags, err := tagdb.GetAll(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: tags})
}

func (Tag) GetCount(c iris.Context) {
	var form dto.TagForm
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
	tags, err := tagdb.GetCount(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: tags})
}

func (Tag) GetAllVisible(c iris.Context) {
	tags, err := tagdb.GetAllVisible()
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: tags})
}

func (Tag) GetAllShared(c iris.Context) {
	_, err := isAdmin(c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	tags, err := tagdb.GetAllShared()
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: tags})
}

func (Tag) SetVisibleTrue(c iris.Context) {
	var id dto.Id
	err := c.ReadJSON(&id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	err = tag.isPermitted(c, id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	err = tagdb.SetVisibleTrue(id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: time.Now().Format("2006-01-02 15:04:05")})
}

func (Tag) SetVisibleFalse(c iris.Context) {
	var id dto.Id
	err := c.ReadJSON(&id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	err = tag.isPermitted(c, id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	err = tagdb.SetVisibleFalse(id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: time.Now().Format("2006-01-02 15:04:05")})
}

func (Tag) SetSharedTrue(c iris.Context) {
	var id dto.Id
	err := c.ReadJSON(&id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	err = tag.isCreator(c, id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	err = tagdb.SetSharedTrue(id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: time.Now().Format("2006-01-02 15:04:05")})
}

func (Tag) SetSharedFalse(c iris.Context) {
	var id dto.Id
	err := c.ReadJSON(&id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	err = tag.isCreator(c, id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	err = tagdb.SetSharedFalse(id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: time.Now().Format("2006-01-02 15:04:05")})
}

func (Tag) AddTag(c iris.Context) {
	var t dto.Tag
	err := c.ReadJSON(&t)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	userId, err := isAdmin(c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	t.Cid = userId
	err = tagdb.InsertTag(&t)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: "save successfully"})
}

func (Tag) UpdateTag(c iris.Context) {
	var t dto.TagBrief
	err := c.ReadJSON(&t)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	err = tag.isPermitted(c, t.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	err = tagdb.UpdateTag(&t)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: "save successfully"})
}

func (Tag) DeleteTag(c iris.Context) {
	var id dto.Id
	err := c.ReadJSON(&id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	err = tag.isCreator(c, id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	count, err := tagdb.IsDepended(id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	if count != 0 {
		c.JSON(&dto.Res{Error: errors.New("can't delete tag: the tag is depended by " + strconv.Itoa(count) + " problems").Error(), Data: nil})
		return
	}
	err = tagdb.DeleteTag(id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: "delete tag successfully"})
}

func (Tag) isPermitted(c iris.Context, tid int64) error {
	id, err := session.GetInt64(c, "userId")
	if err != nil {
		return errors.New("please log in")
	}
	userType, err := userdb.GetUserType(id)
	if userType < 2 {
		return errors.New("not allowed")
	}
	if userType == 3 {
		return nil
	}
	creatorId, err := tagdb.GetCreatorId(tid)
	if err != nil {
		return err
	}
	if id == creatorId {
		return nil
	}
	return tagdb.IsShared(tid)
}

func (Tag) isCreator(c iris.Context, tid int64) error {
	id, err := session.GetInt64(c, "userId")
	if err != nil {
		return errors.New("please log in")
	}
	userType, err := userdb.GetUserType(id)
	if userType < 2 {
		return errors.New("not allowed")
	}
	if userType == 3 {
		return nil
	}
	creatorId, err := tagdb.GetCreatorId(tid)
	if err != nil {
		return err
	}
	if id != creatorId {
		return errors.New("not allowed")
	}
	return nil
}
