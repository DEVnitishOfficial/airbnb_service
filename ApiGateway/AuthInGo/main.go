package main

import (
	"AuthInGo/app"
)

func main() {

	// creating a config object
	// this config object will be used to configure the server
	cfg := app.Config{
		Addr: ":3005", // specify the port
	}
	// the run() function is available on the Application struct
	// so we have to create an object of Application struct
	app := app.Application{
		Config: cfg,
	}

	app.Run()
}
