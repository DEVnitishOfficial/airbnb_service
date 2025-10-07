package main

import (
	"ReviewService/app"
	config "ReviewService/config/env"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	log.Println("Loading environment variables from .env file")
	config.Load()
	cfg := app.NewConfig()
	application := app.NewApplication(cfg)

	application.Run()
}
