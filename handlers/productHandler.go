package handlers

import (
	"strconv"

	"juanfeLogis/dtos/request"
	"juanfeLogis/services"
	"juanfeLogis/utils"

	"github.com/gofiber/fiber/v3"
)

type ProductHandler struct {
	productService *services.ProductService
}

func NewProductHandler(productService *services.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

func (h *ProductHandler) Create(c fiber.Ctx) error {
	var req request.ProductRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Datos inválidos"})
	}

	if err := utils.ValidateProductRequest(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	res, err := h.productService.CreateProduct(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(res)
}

func (h *ProductHandler) GetAll(c fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("page_size", "10"))
	itemType := c.Query("type", "")
	donor := c.Query("donor", "")
	size := c.Query("size", "")
	disposition := c.Query("disposition", "")

	pagination := &utils.Pagination{
		Page:     page,
		PageSize: pageSize,
	}

	res, err := h.productService.GetAllProducts(pagination, itemType, donor, size, disposition)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(res)
}

func (h *ProductHandler) GetByID(c fiber.Ctx) error {
	id := c.Params("id")
	res, err := h.productService.GetByID(id)
	if err != nil && err.Error() == "producto no encontrado" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(res)
}

func (h *ProductHandler) Update(c fiber.Ctx) error {
	id := c.Params("id")
	var req request.ProductRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Datos inválidos"})
	}

	res, err := h.productService.UpdateProduct(id, req)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(res)
}

func (h *ProductHandler) Delete(c fiber.Ctx) error {
	id := c.Params("id")
	if err := h.productService.DeleteProduct(id); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
