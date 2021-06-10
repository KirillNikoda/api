package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/KirillNikoda/api/api/internal/app/api"
	"log"
)

var (
	configPath = "configs/api.toml"
)

func init() {
	//В этот момент происходит инициализация переменной configPath значением
	flag.StringVar(&configPath, "path", "configs/api.toml", "path to config file in .toml format")
}

func main() {
	log.Println("It works")
	//Server instance initialization
	config := api.NewConfig()
	_, err := toml.DecodeFile(configPath, config)

	if err != nil {
		log.Println("can't find configs file. using default values:", err)
	}

	server := api.New(config)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
