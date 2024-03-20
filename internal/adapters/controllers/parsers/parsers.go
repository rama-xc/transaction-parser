package prscontroller

import (
	"github.com/kataras/iris/v12"
	"transaction-parser/internal/usecases/parser"
)

type Controller struct {
	prsrs map[string]parser.IHistory
}

func New(prsrs map[string]parser.IHistory) *Controller {
	return &Controller{prsrs: prsrs}
}

func (c *Controller) ShowAll(ctx iris.Context) {
	ids := map[string][]string{"ids": {}}

	for id, _ := range c.prsrs {
		ids["ids"] = append(ids["ids"], id)
	}

	_ = ctx.JSON(ids)
}
