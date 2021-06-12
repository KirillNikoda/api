package main

import (
	"log"

	"github.com/KirillNikoda/api/api/internal/app/api"
	"github.com/joho/godotenv"
)


func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("could not find .env file:", err)
	}
}


func main() {
	log.Println("It works")
	//Server instance initialization
	config := api.NewConfig()

	server := api.New(config)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
