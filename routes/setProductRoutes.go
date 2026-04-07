package routes

import (
	"juanfeLogis/config"
	"juanfeLogis/handlers"
	"juanfeLogis/repositories"
	"juanfeLogis/services"

	"github.com/gofiber/fiber/v3"
)

func SetProductRouter(api fiber.Router) {
	productRepo := repositories.NewProductRepository(config.DB)
	paginationRepo := repositories.NewPaginationRepository(config.DB)
	productService := services.NewProductService(productRepo, paginationRepo)
	productHandler := handlers.NewProductHandler(productService)

	api.Get("/products", productHandler.GetAll)
	api.Get("/products/:id", productHandler.GetByID)
	api.Post("/products", productHandler.Create)
	api.Put("/products/:id", productHandler.Update)
	api.Delete("/products/:id", productHandler.Delete)
}
