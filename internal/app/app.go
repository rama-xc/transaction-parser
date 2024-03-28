package app

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/kataras/iris/v12"
	interceptor "github.com/kataras/iris/v12/middleware/logger"
	"log/slog"
	prscontroller "transaction-parser/internal/adapters/controllers/parsers"
	"transaction-parser/internal/app/common/config"
	"transaction-parser/internal/app/common/logger"
	"transaction-parser/internal/app/drivers/redis"
	"transaction-parser/internal/usecases/parser"
)

type App struct {
	port  int
	log   *slog.Logger
	prsrs map[string]parser.IHistory
}

func MustLoad(cfgPath string) *App {
	cfg := config.MustLoad(cfgPath)
	log := logger.MustLoad(cfg.Env)
	rds := rdsdrv.MustLoad()
	prs := parser.MustLoad(cfg.ParsersFactories, log, rds)

	app := &App{
		port:  cfg.HTTP.Port,
		log:   log,
		prsrs: prs,
	}

	log.Info("Welcome to BlockChain Transaction Parser ;) Application created successfully.")

	return app
}

func (a *App) MustRun() {
	op := "App.MustRun"
	app := iris.New()
	app.Validator = validator.New()

	// Загрузка голбальных Middlewares
	app.Use(interceptor.New())

	// Загрузка Роутера
	core := app.Party("/")
	{
		core.Get("/", func(ctx iris.Context) { _ = ctx.JSON(map[string]string{"ping": "pong"}) })
		v1 := core.Party("/v1")
		{
			prs := v1.Party("/prs")
			{
				controller := prscontroller.New(a.prsrs)

				prs.Get("/", controller.ShowAll)
				prs.Post("/run", controller.Run)
				prs.Get("/profile/{id}", controller.Profiling)
			}
		}
	}

	// Запуск HTTP сервера
	if err := app.Listen(fmt.Sprintf(":%d", a.port)); err != nil {
		panic(fmt.Sprintf("%s: %e", op, err))
	}
}
