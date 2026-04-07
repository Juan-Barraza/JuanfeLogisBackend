package utils

import "github.com/gofiber/fiber/v3"

func GetUserID(c fiber.Ctx) string {
	val, _ := c.Locals("userID").(string)
	return val
}
