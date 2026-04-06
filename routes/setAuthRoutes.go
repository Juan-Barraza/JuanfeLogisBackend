package routes

import (
	"juanfeLogis/config"
	"juanfeLogis/handlers"
	"juanfeLogis/repositories"
	"juanfeLogis/services"

	"github.com/gofiber/fiber/v3"
)

func SetAuthRouter(apiv1 fiber.Router) {
	userRepo := repositories.NewUserRepository(config.DB)
	authServ := services.NewAuthService(userRepo)
	authHandler := handlers.NewAuthHandler(authServ)

	apiv1.Post("/auth/login", authHandler.Login)
}
