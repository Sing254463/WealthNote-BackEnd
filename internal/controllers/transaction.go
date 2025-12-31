package controllers

import (
	"WealthNoteBackend/internal/services"
	"WealthNoteBackend/pkg/utils"
	"fmt" // ⚠️ เพิ่ม
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
	fmt.Printf("🔍 user_id from context: %v (type: %T)\n", userID, userID) // ⚠️ เพิ่ม log

	if userID == nil {
		return utils.ErrorResponse(c, "User not authenticated", fiber.StatusUnauthorized)
	}

	// ✅ แปลง user_id เป็น int
	id, err := strconv.Atoi(userID.(string))
	if err != nil {
		fmt.Printf("❌ Error converting user_id: %v\n", err) // ⚠️ เพิ่ม log
		return utils.ErrorResponse(c, "Invalid user ID", fiber.StatusBadRequest)
	}

	fmt.Printf("✅ User ID: %d\n", id) // ⚠️ เพิ่ม log

	// ✅ ส่ง user_id ไปยัง service
	transactions, err := services.GetTransactionByUserID(id)
	if err != nil {
		return utils.ErrorResponse(c, err.Error(), fiber.StatusInternalServerError)
	}

	return utils.SuccessResponse(c, transactions, "Transactions retrieved successfully")
}
