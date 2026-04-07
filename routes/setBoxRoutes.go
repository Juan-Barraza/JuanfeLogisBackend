package routes

import (
	"juanfeLogis/config"
	"juanfeLogis/handlers"
	"juanfeLogis/repositories"
	"juanfeLogis/services"

	"github.com/gofiber/fiber/v3"
)

func SetBoxRouter(api fiber.Router) {
	boxRepo := repositories.NewBoxRepository(config.DB)
	paginationRepo := repositories.NewPaginationRepository(config.DB)
	boxService := services.NewBoxService(boxRepo, paginationRepo)
	boxHandler := handlers.NewBoxHandler(boxService)

	api.Get("/boxes", boxHandler.GetAll)
	api.Get("/boxes/:id", boxHandler.GetByID)
	api.Post("/boxes", boxHandler.Create)
	api.Put("/boxes/:id", boxHandler.Update)
	api.Delete("/boxes/:id", boxHandler.Delete)
	api.Get("/boxes/:id/qr", boxHandler.GetQR)
}
