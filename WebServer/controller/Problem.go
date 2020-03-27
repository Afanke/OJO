package ctrl

import (
	"github.com/afanke/OJO/WebServer/dto"
	"github.com/kataras/iris"
)

type Problem struct {
}

func (Problem) GetAllTags(c iris.Context) {
	tags, err := pbdb.GetAllTags()
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: tags})
}

func (Problem) AddProblem(c iris.Context) {
	var p dto.Problem
	err := c.ReadJSON(&p)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	err = pbdb.InsertProblem(&p)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: "save successfully"})
}
