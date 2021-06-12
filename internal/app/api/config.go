package api

import (
	"os"

	"github.com/KirillNikoda/api/api/storage"
)

//General instance for API server of REST application

type Config struct {
	//Port
	BindArr string
	//Logger Level
	LoggerLevel string
	//Storage Level
	Storage *storage.Config
}

func NewConfig() *Config {
	return &Config{
		BindArr:     os.Getenv("port"),
		LoggerLevel: "debug",
		Storage:     storage.NewConfig(),
	}

}


