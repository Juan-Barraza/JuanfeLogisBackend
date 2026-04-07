package routes

import (
	"juanfeLogis/config"
	"juanfeLogis/handlers"
	"juanfeLogis/repositories"
	"juanfeLogis/services"

	"github.com/gofiber/fiber/v3"
)

func SetBoxStockRoutes(apiv1 fiber.Router) {
	stockRepo := repositories.NewBoxStockRepository(config.DB)
	transactionRepo := repositories.NewTransactionRepository(config.DB)
	productRepo := repositories.NewProductRepository(config.DB)
	boxStockService := services.NewBoxStockService(config.DB, stockRepo, transactionRepo, productRepo)
	handler := handlers.NewBoxStockHandler(boxStockService)

	stock := apiv1.Group("/boxes/:boxId/stock")

	stock.Post("/add", handler.AddStock)
	stock.Post("/remove", handler.RemoveStock)
	stock.Post("/return", handler.ReturnStock)
}
