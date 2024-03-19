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
	"transaction-parser/internal/usecases/parser"
)

type App struct {
	port int
	log  *slog.Logger
}

func MustLoad(cfgPath string) *App {
	cfg := config.MustLoad(cfgPath)
	log := logger.MustLoad(cfg.Env)

	app := &App{
		port: cfg.HTTP.Port,
		log:  log,
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
				factory, err := parser.GetParsersFactory(
					config.ParserConfig{
						ID:          "ETH:mainnet",
						ProviderUrl: "https://eth-mainnet.g.alchemy.com/v2/9Iwzq4BnnPoiF8i0X_QHKtczKVCy2bkS",
						Blockchain:  "ethereum",
					},
				)
				if err != nil {
					panic("cant create factory")
				}

				controller := prscontroller.New(a.log, factory.GetHistoryParser(
					0,
					1000,
					4,
					a.log,
				))

				prs.Get("/run", controller.RunParsing)
			}
		}
	}

	// Запуск HTTP сервера
	if err := app.Listen(fmt.Sprintf(":%d", a.port)); err != nil {
		panic(fmt.Sprintf("%s: %e", op, err))
	}
}
