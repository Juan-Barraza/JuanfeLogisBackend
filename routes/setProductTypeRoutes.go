package routes

import (
	"juanfeLogis/config"
	"juanfeLogis/handlers"
	"juanfeLogis/repositories"
	"juanfeLogis/services"

	"github.com/gofiber/fiber/v3"
)

func SetProductTypeRouter(api fiber.Router) {
	ptRepo := repositories.NewProductTypeRepository(config.DB)
	ptService := services.NewProductTypeService(ptRepo)
	ptHandler := handlers.NewProductTypeHandler(ptService)

	api.Post("/product-types", ptHandler.Create)
	api.Get("/product-types", ptHandler.GetAll)
	api.Put("/product-types/:id", ptHandler.Update)
}
