package prscontroller

import (
	"github.com/kataras/iris/v12"
	"net/http"
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

func (c *Controller) Run(ctx iris.Context) {
	var dto RunBody

	if err := ctx.ReadBody(&dto); err != nil {
		ctx.StopWithError(http.StatusBadRequest, ValidationErr)
		return
	}

	prsr, ok := c.prsrs[dto.ID]
	if ok != true {
		ctx.StopWithError(http.StatusNotFound, ParserNotFoundErr)
		return
	}

	prsr.SendCommand(
		parser.NewRunCommand(prsr),
	)

	_ = ctx.JSON(map[string]string{"message": "command send"})
}

func (c *Controller) Option(ctx iris.Context) {
	panic("Implement me!")

	//var dto OptionsBody
	//
	//if err := ctx.ReadParams(&dto); err != nil {
	//	ctx.StopWithError(http.StatusBadRequest, ValidationErr)
	//	return
	//}
	//
	//prsr, ok := c.prsrs[dto.ID]
	//if ok != true {
	//	ctx.StopWithError(http.StatusNotFound, ParserNotFoundErr)
	//	return
	//}
	//
	//prsr.SendCommand(
	//	block.NewOptionsCommand(prsr, dto.Wrks),
	//)
	//
	//_ = ctx.JSON(map[string]string{"message": "command send"})
}

func (c *Controller) Profiling(ctx iris.Context) {
	panic("Implement me!")

	//var dto ProfilingParams
	//
	//if err := ctx.ReadParams(&dto); err != nil {
	//	ctx.StopWithError(http.StatusBadRequest, ValidationErr)
	//	return
	//}
	//
	//prsr, ok := c.prsrs[dto.ID]
	//if ok != true {
	//	ctx.StopWithError(http.StatusNotFound, ParserNotFoundErr)
	//	return
	//}
	//
	//_ = ctx.JSON(
	//	prsr.Profile(),
	//)
}
