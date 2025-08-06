package main

import (
	"AuthInGo/app"
)

func main() {

	// creating a config object
	cfg := app.NewConfig(":3005")

	// created an object of Application struct
	app := app.NewApplication(cfg)

	app.Run()
}
