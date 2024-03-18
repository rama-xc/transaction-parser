package app

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/kataras/iris/v12"
	interceptor "github.com/kataras/iris/v12/middleware/logger"
	"log/slog"
	corecontroller "transaction-parser/internal/adapters/controllers/core"
	prscontroller "transaction-parser/internal/adapters/controllers/parsers"
	"transaction-parser/internal/app/common/config"
	"transaction-parser/internal/app/common/logger"
	prsrcmps "transaction-parser/internal/usecases/parser-composer"
)

type ComposerMap map[string]prsrcmps.IParserComposer

type App struct {
	port        int
	log         *slog.Logger
	prsComposer ComposerMap
}

func MustLoad(cfgPath string) *App {
	cfg := config.MustLoad(cfgPath)
	log := logger.MustLoad(cfg.Env)
	prsComposer := prsrcmps.MustLoadParsers(cfg.Parsers)

	app := &App{
		port:        cfg.HTTP.Port,
		log:         log,
		prsComposer: prsComposer,
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
		core.Get("/", corecontroller.Ping)
		v1 := core.Party("/v1")
		{
			prs := v1.Party("/parsers")
			{
				ctrl := prscontroller.New(
					a.log,
					a.prsComposer,
				)

				prs.Get("/", ctrl.ParsersIDs)
				prs.Get("/{id}", ctrl.ParserExist)
				prs.Post("/run", ctrl.Run)
			}
		}
	}

	// Запуск HTTP сервера
	if err := app.Listen(fmt.Sprintf(":%d", a.port)); err != nil {
		panic(fmt.Sprintf("%s: %e", op, err))
	}
}
