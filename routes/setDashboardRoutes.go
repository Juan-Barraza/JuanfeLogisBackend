package routes

import (
	"juanfeLogis/config"
	"juanfeLogis/handlers"
	"juanfeLogis/repositories"
	"juanfeLogis/services"

	"github.com/gofiber/fiber/v3"
)

func SetupDashboardRoutes(apiv1 fiber.Router) {

	dashboardRepo := repositories.NewDashboardRepository(config.DB)
	dashboardService := services.NewDashboardService(dashboardRepo)
	dashboardHandler := handlers.NewDashboardHandler(dashboardService)

	dashboardGroup := apiv1.Group("/dashboard")

	dashboardGroup.Get("/logistics/kpis", dashboardHandler.GetLogisticsKPIs)
	dashboardGroup.Get("/logistics/distribution", dashboardHandler.GetLogisticsDistribution)
	dashboardGroup.Get("/logistics/locations", dashboardHandler.GetLocationDistribution)
	dashboardGroup.Get("/logistics/donors/top", dashboardHandler.GetTopDonorsLogistics)
	dashboardGroup.Get("/financial/kpis", dashboardHandler.GetFinancialKPIs)
	dashboardGroup.Get("/financial/trends", dashboardHandler.GetFinancialTrends)
	dashboardGroup.Get("/financial/donors/top", dashboardHandler.GetTopDonorsFinancial)
	dashboardGroup.Get("/financial/profitability", dashboardHandler.GetProfitability)
}
