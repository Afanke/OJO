package ctrl

import (
	"errors"
	"github.com/afanke/OJO/WebServer/db"
	"github.com/afanke/OJO/WebServer/dto"
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
	admin, err := isAdmin(c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	form.Cid = admin.Id
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
	admin, err := isAdmin(c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	form.Cid = admin.Id
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

func (Tag) GetAllCommunal(c iris.Context) {
	_, err := isAdmin(c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	tags, err := tagdb.GetAllCommunal()
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
	user, err := isAdmin(c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	t.Cid = user.Id
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
	token, err := getUserToken(c)
	if err != nil {
		return errors.New("please log in")
	}
	if token.Type < 2 {
		return errors.New("not allowed")
	}
	if token.Type == 3 {
		return nil
	}
	id, err := tagdb.GetCreatorId(tid)
	if err != nil {
		return err
	}
	if id == token.Id {
		return nil
	}
	return tagdb.IsShared(tid)
}

func (Tag) isCreator(c iris.Context, tid int64) error {
	token, err := getUserToken(c)
	if err != nil {
		return errors.New("please log in")
	}
	if token.Type == 3 {
		return nil
	}
	id, err := tagdb.GetCreatorId(tid)
	if err != nil {
		return err
	}
	if id != token.Id {
		return errors.New("not allowed")
	}
	return nil
}
