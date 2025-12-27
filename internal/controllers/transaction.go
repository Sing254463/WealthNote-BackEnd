package controllers

import (
	"WealthNoteBackend/internal/services"
	"WealthNoteBackend/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

// GetTransactionAll
func GetTransactionAll(c *fiber.Ctx) error {
	Transaction, err := services.GetTransactionAll()
	if err != nil {
		return utils.ErrorResponse(c, err.Error(), fiber.StatusInternalServerError)
	}
	return utils.SuccessResponse(c, Transaction, "Transaction retrieved successfully")
}
