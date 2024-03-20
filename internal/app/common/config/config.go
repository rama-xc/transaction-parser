package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type BlockchainName string

const (
	Ethereum BlockchainName = "ethereum"
)

type Config struct {
	Env              string                   `toml:"env" env-default:"local"`
	HTTP             HTTPConfig               `toml:"http-server"`
	ParsersFactories []ParsersFactoriesConfig `toml:"parsers-factories"`
}

type HTTPConfig struct {
	Port int `toml:"port" env-required:"true"`
}

type ParsersFactoriesConfig struct {
	ID          string         `toml:"id" env-required:"true"`
	ProviderUrl string         `toml:"provider_url" env-required:"true"`
	Blockchain  BlockchainName `toml:"blockchain" env-required:"true"`
	Parsers     ParsersConfig  `toml:"parsers" env-required:"true"`
}

type ParsersConfig struct {
	History HistoryParserConfig `toml:"history"`
}

type HistoryParserConfig struct {
	Workers   int   `toml:"workers" env-required:"true"`
	BlockTo   int64 `toml:"blockTo" env-required:"true"`
	BlockFrom int64 `toml:"blockFrom" env-required:"true"`
}

func MustLoad(path string) *Config {
	op := "Config.MustLoad"

	if path == "" {
		panic(op + ": config path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic(op + ": config file does not exist: " + path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}

	return &cfg
}
