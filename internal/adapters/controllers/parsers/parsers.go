package prscontroller

import (
	"errors"
	"github.com/kataras/iris/v12"
	"log/slog"
	prsrcmps "transaction-parser/internal/usecases/parser-composer"
	"transaction-parser/internal/usecases/parser-composer/parser"
)

type Controller struct {
	log  *slog.Logger
	cprs map[string]prsrcmps.IParserComposer
}

func New(log *slog.Logger, cprs map[string]prsrcmps.IParserComposer) *Controller {
	return &Controller{log: log, cprs: cprs}
}

func (c *Controller) ParsersIDs(ctx iris.Context) {
	res := map[string][]string{"ids": []string{}}

	for s, _ := range c.cprs {
		res["ids"] = append(res["ids"], s)
	}

	_ = ctx.JSON(res)
}

func (c *Controller) ParserExist(ctx iris.Context) {
	var p ParserExistDto
	if err := ctx.ReadParams(&p); err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	_, ok := c.cprs[p.ID]

	_ = ctx.JSON(
		map[string]bool{"is_exist": ok},
	)
}

func (c *Controller) Run(ctx iris.Context) {
	var dto RunDto
	if err := ctx.ReadJSON(&dto); err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	composer, ok := c.cprs[dto.ID]
	if ok == false {
		ctx.StopWithError(iris.StatusBadRequest, errors.New("composer doesn't exist"))
		return
	}

	var opts map[parser.Type]int

	for _, prs := range dto.Parsers {
		opts[prs.Type] = prs.Workers
	}

	command := prsrcmps.NewRunCommand(opts, composer)
	cc := prsrcmps.NewComposerControl(
		command,
	)

	cc.Run()

	_ = ctx.JSON(
		map[string]bool{"is_exist": ok},
	)
}
