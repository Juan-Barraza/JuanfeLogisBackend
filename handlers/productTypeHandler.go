package handlers

import (
	"strconv"

	"juanfeLogis/dtos/request"
	"juanfeLogis/services"

	"github.com/gofiber/fiber/v3"
)

type ProductTypeHandler struct {
	ptService *services.ProductTypeService
}

func NewProductTypeHandler(ptService *services.ProductTypeService) *ProductTypeHandler {
	return &ProductTypeHandler{ptService: ptService}
}

func (h *ProductTypeHandler) Create(c fiber.Ctx) error {
	var req request.ProductTypeRequest

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Datos de entrada inválidos"})
	}

	if req.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "el nombre es requerido"})
	}

	res, err := h.ptService.CreateProductType(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(res)
}

func (h *ProductTypeHandler) GetAll(c fiber.Ctx) error {
	res, err := h.ptService.GetAllProductTypes()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *ProductTypeHandler) Update(c fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID inválido"})
	}

	var req request.ProductTypeRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Datos de entrada inválidos"})
	}

	if req.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "el nombre es requerido"})
	}

	res, err := h.ptService.UpdateProductType(uint(id), req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
