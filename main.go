package main

import (
	"github.com/ghynmo/MVC_Learning/config"
	"github.com/ghynmo/MVC_Learning/routes"
)

func main() {
	config.InitDB()
	config.MigrateDB()
	echoApp := routes.NewRoutes()
	echoApp.Start(":8080")
}
