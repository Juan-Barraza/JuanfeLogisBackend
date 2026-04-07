package handlers

import (
	"juanfeLogis/dtos/request"
	"juanfeLogis/services"

	"github.com/gofiber/fiber/v3"
)

type DashboardHandler struct {
	dashboardService *services.DashboardService
}

func NewDashboardHandler(service *services.DashboardService) *DashboardHandler {
	return &DashboardHandler{dashboardService: service}
}

func (h *DashboardHandler) GetLogisticsKPIs(c fiber.Ctx) error {
	var req request.DashboardFilterRequest

	// Fiber v3 atrapa los parámetros de la URL ?size=M&condition=nuevo
	if err := c.Bind().Query(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Parámetros de filtro inválidos",
		})
	}

	kpis, err := h.dashboardService.GetLogisticsKPIs(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al calcular KPIs logísticos",
		})
	}

	return c.Status(fiber.StatusOK).JSON(kpis)
}

func (h *DashboardHandler) GetLogisticsDistribution(c fiber.Ctx) error {
	var req request.DashboardFilterRequest

	if err := c.Bind().Query(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Parámetros de filtro inválidos",
		})
	}

	distribution, err := h.dashboardService.GetLogisticsDistribution(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al calcular distribuciones logísticas",
		})
	}

	return c.Status(fiber.StatusOK).JSON(distribution)
}

func (h *DashboardHandler) GetLocationDistribution(c fiber.Ctx) error {
	var req request.DashboardFilterRequest
	if err := c.Bind().Query(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Filtros inválidos"})
	}
	distribution, err := h.dashboardService.GetLocationDistribution(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error al calcular ubicaciones"})
	}
	return c.Status(fiber.StatusOK).JSON(distribution)
}

func (h *DashboardHandler) GetTopDonorsLogistics(c fiber.Ctx) error {
	var req request.DashboardFilterRequest

	if err := c.Bind().Query(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Parámetros de filtro inválidos",
		})
	}

	donors, err := h.dashboardService.GetTopDonorsLogistics(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al calcular el ranking de donantes",
		})
	}

	return c.Status(fiber.StatusOK).JSON(donors)
}

func (h *DashboardHandler) GetFinancialKPIs(c fiber.Ctx) error {
	var req request.DashboardFilterRequest

	if err := c.Bind().Query(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Parámetros de filtro inválidos",
		})
	}

	kpis, err := h.dashboardService.GetFinancialKPIs(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al calcular KPIs financieros",
		})
	}

	return c.Status(fiber.StatusOK).JSON(kpis)
}

func (h *DashboardHandler) GetFinancialTrends(c fiber.Ctx) error {
	var req request.DashboardFilterRequest

	if err := c.Bind().Query(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Parámetros de filtro inválidos",
		})
	}

	trends, err := h.dashboardService.GetFinancialTrends(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al calcular las tendencias financieras",
		})
	}

	return c.Status(fiber.StatusOK).JSON(trends)
}

func (h *DashboardHandler) GetTopDonorsFinancial(c fiber.Ctx) error {
	var req request.DashboardFilterRequest
	if err := c.Bind().Query(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Parámetros inválidos"})
	}

	donors, err := h.dashboardService.GetTopDonorsFinancial(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error al calcular donantes financieros"})
	}
	return c.Status(fiber.StatusOK).JSON(donors)
}

func (h *DashboardHandler) GetProfitability(c fiber.Ctx) error {
	var req request.DashboardFilterRequest
	if err := c.Bind().Query(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Parámetros inválidos"})
	}

	profit, err := h.dashboardService.GetProfitability(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error al calcular rentabilidad"})
	}
	return c.Status(fiber.StatusOK).JSON(profit)
}
