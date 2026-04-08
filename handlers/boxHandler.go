package handlers

import (
	"fmt"
	"strconv"

	"juanfeLogis/dtos/request"
	"juanfeLogis/services"
	"juanfeLogis/utils"

	"github.com/gofiber/fiber/v3"
	"github.com/skip2/go-qrcode"
)

type BoxHandler struct {
	boxService *services.BoxService
}

func NewBoxHandler(boxService *services.BoxService) *BoxHandler {
	return &BoxHandler{boxService: boxService}
}

func (h *BoxHandler) Create(c fiber.Ctx) error {
	var req request.BoxRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Datos de entrada inválidos"})
	}

	if req.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "El nombre de la caja es obligatorio"})
	}

	if req.LocationID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "La ubicación de la caja es obligatoria"})
	}

	if len(req.LabelIDs) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Las etiquetas de la caja son obligatorias"})
	}

	res, err := h.boxService.CreateBox(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(res)
}

func (h *BoxHandler) GetAll(c fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("page_size", "10"))
	name := c.Query("name", "")
	location := c.Query("location", "")

	pagination := &utils.Pagination{
		Page:     page,
		PageSize: pageSize,
	}

	res, err := h.boxService.GetAllBoxes(pagination, name, location)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(res)
}

func (h *BoxHandler) GetByID(c fiber.Ctx) error {
	id := c.Params("id")
	res, err := h.boxService.GetBoxDetail(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(res)
}

func (h *BoxHandler) Update(c fiber.Ctx) error {
	id := c.Params("id")

	var req request.BoxRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Datos inválidos"})
	}

	res, err := h.boxService.UpdateBox(id, req)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(res)
}

func (h *BoxHandler) Delete(c fiber.Ctx) error {
	id := c.Params("id")

	if err := h.boxService.DeleteBox(id); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *BoxHandler) GetQR(c fiber.Ctx) error {
	id := c.Params("id")

	box, err := h.boxService.GetByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "caja no encontrada"})
	}

	qrData := box.QRCodeURL
	if qrData == "" {
		// Respaldamos generando el QR si el campo está vacío en DB
		qrData = utils.GenerateBoxQR(id)
	}

	pngBytes, err := qrcode.Encode(qrData, qrcode.Medium, 256)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "error generando QR"})
	}

	c.Set("Content-Type", "image/png")
	c.Set("Content-Disposition", fmt.Sprintf(`attachment; filename="qr-caja-%s.png"`, id))
	return c.Send(pngBytes)
}
