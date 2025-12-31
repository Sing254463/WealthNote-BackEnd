package controllers

import (
	"WealthNoteBackend/internal/models"
	"WealthNoteBackend/internal/services"
	"WealthNoteBackend/pkg/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// GetTransactionAll godoc
// @Summary Get User Transactions
// @Description ดึงข้อมูล transactions ทั้งหมดของ user ที่ login
// @Tags Transactions
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "Transactions retrieved successfully"
// @Failure 401 {object} map[string]interface{} "User not authenticated"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /transactions [get]
func GetTransactionAll(c *fiber.Ctx) error {
	// ✅ Debug: ดูว่า user_id มีค่าอะไร
	userID := c.Locals("user_id")
	if userID == nil {
		return utils.ErrorResponse(c, "User not authenticated", fiber.StatusUnauthorized)
	}

	// ✅ แปลง user_id เป็น int
	id, err := strconv.Atoi(userID.(string))
	if err != nil {
		return utils.ErrorResponse(c, "Invalid user ID", fiber.StatusBadRequest)
	}

	// ✅ ส่ง user_id ไปยัง service
	transactions, err := services.GetTransactionByUserID(id)
	if err != nil {
		return utils.ErrorResponse(c, err.Error(), fiber.StatusInternalServerError)
	}

	return utils.SuccessResponse(c, transactions, "Transactions retrieved successfully")
}

// CreateTransaction - สร้าง transaction ใหม่

// CreateTransaction godoc
// @Summary Create Transaction
// @Description สร้าง transaction ใหม่สำหรับ user ที่ login
// @Tags Transactions
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param transaction body models.CreateTransactionInput true "Transaction data"
// @Success 201 {object} map[string]interface{} "Transaction created successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request body"
// @Failure 401 {object} map[string]interface{} "User not authenticated"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /transactions [post]
func CreateTransaction(c *fiber.Ctx) error {
	// ✅ ดึง user_id จาก JWT Token
	userID := c.Locals("user_id")
	if userID == nil {
		return utils.ErrorResponse(c, "User not authenticated", fiber.StatusUnauthorized)
	}

	id, err := strconv.Atoi(userID.(string))
	if err != nil {
		return utils.ErrorResponse(c, "Invalid user ID", fiber.StatusBadRequest)
	}

	// ✅ Parse request body เป็น CreateTransactionInput
	var input models.CreateTransactionInput
	if err := c.BodyParser(&input); err != nil {
		return utils.ErrorResponse(c, "Invalid request body", fiber.StatusBadRequest)
	}

	// ✅ Set user_id จาก JWT (ไม่ให้ client ส่งมา)
	input.IDUser = id

	// ✅ Validation
	if input.Title == "" {
		return utils.ErrorResponse(c, "Title is required", fiber.StatusBadRequest)
	}
	if input.Amount <= 0 {
		return utils.ErrorResponse(c, "Amount must be greater than 0", fiber.StatusBadRequest)
	}
	if input.IDType == 0 {
		return utils.ErrorResponse(c, "ID Type is required", fiber.StatusBadRequest)
	}
	if input.IDCategory == 0 {
		return utils.ErrorResponse(c, "ID Category is required", fiber.StatusBadRequest)
	}

	// ✅ เรียก service
	transaction, err := services.CreateTransaction(input)
	if err != nil {
		return utils.ErrorResponse(c, err.Error(), fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "Transaction created successfully",
		"data":    transaction,
	})
}
