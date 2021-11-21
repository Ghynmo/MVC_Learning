package main

import (
	"github.com/ghynmo/MVC_Learning/config"
	"github.com/ghynmo/MVC_Learning/middlewares"
	"github.com/ghynmo/MVC_Learning/routes"
)

func main() {
	config.InitDB()
	// config.MigrateDB()
	echoApp := routes.NewRoutes()

	middlewares.LogMiddlewareInit(echoApp)
	//Penamaan function if huruf awal gede artinya Global function

	echoApp.Start(":8080") //start server
}
