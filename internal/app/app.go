package app

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"log/slog"
	"transaction-parser/internal/app/common/config"
	"transaction-parser/internal/app/common/logger"

	interceptor "github.com/kataras/iris/v12/middleware/logger"
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

	// Загрузка голбальных Middlewares
	app.Use(interceptor.New())

	// Загрузка Роутера
	core := app.Party("/")
	{
		core.Get("/", func(ctx iris.Context) {
			_ = ctx.JSON(
				map[string]string{
					"ping": "pong",
				},
			)
		})
	}

	// Запуск HTTP сервера
	if err := app.Listen(fmt.Sprintf(":%d", a.port)); err != nil {
		panic(fmt.Sprintf("%s: %e", op, err))
	}
}
