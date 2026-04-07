package handlers

import (
	"juanfeLogis/dtos/request"
	"juanfeLogis/services"
	"juanfeLogis/utils"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type BoxStockHandler struct {
	service *services.BoxStockService
}

func NewBoxStockHandler(service *services.BoxStockService) *BoxStockHandler {
	return &BoxStockHandler{service: service}
}

func (h *BoxStockHandler) AddStock(c fiber.Ctx) error {
	boxID, err := uuid.Parse(c.Params("boxId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "boxId inválido",
		})
	}

	var req request.BoxStockRequest
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "body inválido",
		})
	}
	if err := utils.ValidateBoxStockRequest(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	userIDStr := utils.GetUserID(c)
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "user_id inválido",
		})
	}

	if err := h.service.AddStock(boxID, req, userID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "stock agregado correctamente",
	})
}

func (h *BoxStockHandler) RemoveStock(c fiber.Ctx) error {
	boxID, err := uuid.Parse(c.Params("boxId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "boxId inválido",
		})
	}

	var req request.BoxStockRequest
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "body inválido",
		})
	}
	if err := utils.ValidateBoxStockRequest(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	userIDStr := utils.GetUserID(c)
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "user_id inválido",
		})
	}

	if err := h.service.RemoveStock(boxID, req, userID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "stock removido correctamente",
	})
}

func (h *BoxStockHandler) ReturnStock(c fiber.Ctx) error {
	boxID, err := uuid.Parse(c.Params("boxId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "boxId inválido",
		})
	}

	var req request.BoxStockRequest
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "body inválido",
		})
	}
	if err := utils.ValidateBoxStockRequest(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	userIDStr := utils.GetUserID(c)
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "user_id inválido",
		})
	}

	if err := h.service.ReturnStock(boxID, req, userID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "devolución registrada correctamente",
	})
}
