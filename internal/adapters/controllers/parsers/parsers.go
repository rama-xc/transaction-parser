package prscontroller

import (
	"github.com/kataras/iris/v12"
	"log/slog"
	"transaction-parser/internal/usecases/parser"
)

type Controller struct {
	log       *slog.Logger
	htrParser parser.IHistory
}

func New(log *slog.Logger, htrParser parser.IHistory) *Controller {
	return &Controller{log: log, htrParser: htrParser}
}

func (c *Controller) RunParsing(ctx iris.Context) {
	c.htrParser.Start()
	_ = ctx.JSON("success")
}
