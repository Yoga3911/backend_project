package utils

import "github.com/gofiber/fiber/v2"

func Response(c *fiber.Ctx, code int, data interface{}, message string, status bool) error {
	return c.Status(code).JSON(fiber.Map{
		"data":    data,
		"message": message,
		"status":  status,
	})
}
