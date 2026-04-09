package routes

import (
	"juanfeLogis/middlewares"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func SetRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// Global middleware
	// CORS
	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")
	if allowedOrigins == "" {
		allowedOrigins = "http://localhost:5173"
	}
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{allowedOrigins},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
	}))

	// Health check
	v1.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "Inventory Juanfe API is running ",
		})
	})

	// api publica
	SetAuthRouter(v1)

	protected := v1.Group("/", middlewares.JWTAuth())
	SetLocationRouter(protected)
	SetProductTypeRouter(protected)
	SetDonorRouter(protected)
	SetBoxRouter(protected)
	SetProductRouter(protected)
	SetBoxStockRoutes(protected)
	SetTransactionRoutes(protected)
	SetupDashboardRoutes(protected)
}
