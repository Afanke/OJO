package ctrl

import (
	"github.com/afanke/OJO/WebServer/db"
	"github.com/afanke/OJO/WebServer/dto"
	"github.com/kataras/iris/v12"
)

type Rank struct{}

var rankdb db.Rank

func (Rank) GetACMTop10(c iris.Context) {
	detail, err := rankdb.GetACMTop10()
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: detail})
}

func (Rank) GetACMRank(c iris.Context) {
	var form dto.RankForm
	err := c.ReadJSON(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	detail, err := rankdb.GetACMRank(form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: detail})
}

func (Rank) GetACMRankCount(c iris.Context) {
	detail, err := rankdb.GetACMRankCount()
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: detail})
}
