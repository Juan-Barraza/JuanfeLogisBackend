package routes

import (
	"juanfeLogis/config"
	"juanfeLogis/handlers"
	"juanfeLogis/repositories"
	"juanfeLogis/services"

	"github.com/gofiber/fiber/v3"
)

func SetTransactionRoutes(api fiber.Router) {
	txRepo := repositories.NewTransactionRepository(config.DB)
	paginationRepo := repositories.NewPaginationRepository(config.DB)
	txService := services.NewTransactionService(txRepo, paginationRepo)
	txHandler := handlers.NewTransactionHandler(txService)

	group := api.Group("/transactions")
	group.Get("/", txHandler.GetAll)
	group.Get("/:id", txHandler.GetByID)
}
