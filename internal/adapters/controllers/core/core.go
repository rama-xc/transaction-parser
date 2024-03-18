package corecontroller

import "github.com/kataras/iris/v12"

func Ping(ctx iris.Context) {
	_ = ctx.JSON(
		map[string]string{
			"ping": "pong",
		},
	)
}
