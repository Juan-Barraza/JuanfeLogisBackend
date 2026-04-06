package routes

import (
	"juanfeLogis/config"
	"juanfeLogis/handlers"
	"juanfeLogis/repositories"
	"juanfeLogis/services"

	"github.com/gofiber/fiber/v3"
)

func SetLocationRouter(api fiber.Router) {
	locationRepo := repositories.NewLocationRepository(config.DB)
	locationService := services.NewLocationService(locationRepo)
	locationHandler := handlers.NewLocationHandler(locationService)

	api.Post("/locations", locationHandler.Create)
	api.Get("/locations", locationHandler.GetAll)
	api.Put("/locations/:id", locationHandler.Update)
}
