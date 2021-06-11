package api

import "github.com/KirillNikoda/api/api/storage"

//General instance for API server of REST application

type Config struct {
	//Port
	BindArr string `toml:"bind_addr"`
	//Logger Level
	LoggerLevel string `toml:"logger_level"`
	//Storage Level
	Storage *storage.Config `toml:"database_uri"`
}

func NewConfig() *Config {
	return &Config{
		BindArr:     "8080",
		LoggerLevel: "debug",
		Storage:     storage.NewConfig(),
	}
}
