package main

import (
	"github.com/ghynmo/MVC_Learning/config"
	"github.com/ghynmo/MVC_Learning/routes"
)

func main() {
	config.InitDB()
	echoApp := routes.NewRoutes()
	echoApp.Start(":8080")
}
