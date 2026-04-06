package main

import (
	"juanfeLogis/config"
	"juanfeLogis/routes"
	"juanfeLogis/utils"
	"log"
)

func main() {

	config.ConnectDB()
	config.MigrateDB()

	app, err := utils.InitFiber()
	if err != nil {
		log.Fatalf("Error to initializing fiber: %v", err)
	}

	routes.SetRoutes(app)

	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}
