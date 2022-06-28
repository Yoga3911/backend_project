package helpers

import "github.com/gofiber/fiber/v2"

func SuccessResponse(c *fiber.Ctx, code int, message string, data interface{}) error {
	return c.Status(code).JSON(fiber.Map{
		"data":    data,
		"message": message,
		"status":  true,
	})
}

func FailedResponse(c *fiber.Ctx, code int, message string) error {
	return c.Status(code).JSON(fiber.Map{
		"data":    nil,
		"message": message,
		"status":  false,
	})
}
