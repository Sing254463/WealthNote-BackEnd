package utils

import (
	"github.com/gofiber/fiber/v2"
)

func SuccessResponse(c *fiber.Ctx, data interface{}, message string) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": message,
		"data":    data,
	})
}

func ErrorResponse(c *fiber.Ctx, message string, code int) error {
	return c.Status(code).JSON(fiber.Map{
		"status":  "error",
		"message": message,
	})
}