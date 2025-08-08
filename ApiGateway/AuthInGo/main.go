package main

import (
	"AuthInGo/app"
	config "AuthInGo/config/env"
)

func main() {

	config.Load() // Load environment variables from .env file

	// creating a config object
	cfg := app.NewConfig(":3005")

	// created an object of Application struct
	app := app.NewApplication(cfg)

	app.Run()
}
