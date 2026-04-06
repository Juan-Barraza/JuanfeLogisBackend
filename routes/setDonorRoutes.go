package routes

import (
	"juanfeLogis/config"
	"juanfeLogis/handlers"
	"juanfeLogis/repositories"
	"juanfeLogis/services"

	"github.com/gofiber/fiber/v3"
)

func SetDonorRouter(api fiber.Router) {
	repo := repositories.NewDonorRepository(config.DB)
	paginationRepo := repositories.NewPaginationRepository(config.DB)
	svc := services.NewDonorService(repo, paginationRepo)
	hdl := handlers.NewDonorHandler(svc)

	api.Post("/donors", hdl.Create)
	api.Get("/donors", hdl.GetAll)
	api.Get("/donors/:name", hdl.GetByName)
	api.Put("/donors/:id", hdl.Update)
}
