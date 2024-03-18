package main

import (
	"flag"
	"os"
	"transaction-parser/internal/app"
)

func main() {
	cfgPath := loadConfigPath()

	application := app.MustLoad(cfgPath)

	application.MustRun()
}

func loadConfigPath() string {
	var path string

	flag.StringVar(&path, "config", "./config/local.toml", "path to config file")
	flag.Parse()

	if path == "" {
		path = os.Getenv("CONFIG_PATH")
	}

	return path
}
