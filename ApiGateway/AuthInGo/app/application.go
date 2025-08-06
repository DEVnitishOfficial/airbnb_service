package app

import (
	"fmt"
	"net/http"
	"time"
)

// config holds configuration for the server(application)
type Config struct {
	Addr string // port
}

// Application struct holds ther server configuration
// and this configuration will be provided by the config struct
type Application struct {
	Config Config
}

// defining member function inside application struct
// (app *Application) --->> it is a receiver function
// it means that this function is associated with the Application struct
// func (receiver_name receiver_type) MethodName() return_type ---->>> its a syntax to define a method in Go
func (app *Application) Run() error {
	//returning reference of the created server
	server := &http.Server{
		// below is the configuration of the server
		Addr:         app.Config.Addr,
		Handler:      nil, // TODO : setup a chi router and put it here
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("Starting server on http://localhost", app.Config.Addr)
	// server provide us ListenAndServe method which starts the server
	// and if any error occurs it will return that error
	// and that error we can return from this Run method
	return server.ListenAndServe() // it will listen on the port specified in the config
}
