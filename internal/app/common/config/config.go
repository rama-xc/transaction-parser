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
	Env     string         `toml:"env" env-default:"local"`
	HTTP    HTTPConfig     `toml:"http-server"`
	Parsers []ParserConfig `toml:"parsers"`
}

type HTTPConfig struct {
	Port int `toml:"port" env-required:"true"`
}

type ParserConfig struct {
	ID          string         `toml:"id" env-required:"true"`
	ProviderUrl string         `toml:"provider_url" env-required:"true"`
	Blockchain  BlockchainName `toml:"blockchain" env-required:"true"`
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
