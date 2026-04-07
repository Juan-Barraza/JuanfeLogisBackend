package handlers

import (
	"juanfeLogis/dtos/request"
	"juanfeLogis/services"
	"juanfeLogis/utils"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

type DonorHandler struct {
	donorService *services.DonorService
}

func NewDonorHandler(donorService *services.DonorService) *DonorHandler {
	return &DonorHandler{donorService: donorService}
}

func (h *DonorHandler) Create(c fiber.Ctx) error {
	var req request.DonorRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "JSON inválido"})
	}
	res, err := h.donorService.CreateDonor(req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(res)
}

func (h *DonorHandler) GetAll(c fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("page_size", "10"))
	pagination := &utils.Pagination{
		Page:     page,
		PageSize: pageSize,
	}
	res, err := h.donorService.GetAllDonors(pagination)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(res)
}

func (h *DonorHandler) Update(c fiber.Ctx) error {
	id := c.Params("id")
	var req request.DonorRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "JSON inválido"})
	}
	res, err := h.donorService.UpdateDonor(id, req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(res)
}

func (h *DonorHandler) GetByName(c fiber.Ctx) error {
	name := c.Params("name")
	res, err := h.donorService.GetByName(name)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(res)
}
