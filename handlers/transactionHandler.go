package handlers

import (
	"juanfeLogis/services"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

type TransactionHandler struct {
	txService *services.TransactionService
}

func NewTransactionHandler(txService *services.TransactionService) *TransactionHandler {
	return &TransactionHandler{txService: txService}
}

func (h *TransactionHandler) GetAll(c fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize", "10"))
	txType := c.Query("type", "")
	startDate := c.Query("startDate", "")
	endDate := c.Query("endDate", "")

	results, err := h.txService.GetAllTransactions(page, pageSize, txType, startDate, endDate)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(results)
}

func (h *TransactionHandler) GetByID(c fiber.Ctx) error {
	id := c.Params("id")
	result, err := h.txService.GetTransactionByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(result)
}
