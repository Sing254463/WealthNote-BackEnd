package controllers

import (
	"WealthNoteBackend/internal/models"
	"WealthNoteBackend/internal/services"
	"WealthNoteBackend/pkg/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// GetAllUsers godoc
// @Summary      Get all users
// @Description  Retrieve a list of all users
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /users [get]
// @Security     BearerAuth
func GetAllUsers(c *fiber.Ctx) error {
	// ✅ เชื่อมกับ database จริง
	users, err := services.GetAllUsers()
	if err != nil {
		return utils.ErrorResponse(c, err.Error(), fiber.StatusInternalServerError)
	}

	return utils.SuccessResponse(c, users, "Users retrieved successfully")
}

// GetUserByID godoc
// @Summary      Get user by ID
// @Description  Retrieve a specific user by their ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      404  {object}  map[string]interface{}
// @Router       /users/{id} [get]
// @Security     BearerAuth
func GetUserByID(c *fiber.Ctx) error {
	// ✅ แปลง string เป็น int
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.ErrorResponse(c, "Invalid user ID", fiber.StatusBadRequest)
	}

	// ✅ เชื่อมกับ database จริง
	user, err := services.GetUserByID(id)
	if err != nil {
		return utils.ErrorResponse(c, err.Error(), fiber.StatusNotFound)
	}

	return utils.SuccessResponse(c, user, "User retrieved successfully")
}

type UpdateUserRequest struct {
	UserCode *string `json:"usercode,omitempty"`
	Email    *string `json:"email,omitempty" validate:"omitempty,email"`
	FNameT   *string `json:"fnamet,omitempty"`
	LNameT   *string `json:"lnamet,omitempty"`
	FNameE   *string `json:"fnamee,omitempty"`
	LNameE   *string `json:"lnamee,omitempty"`
}

// UpdateUser godoc
// @Summary      Update an existing user
// @Description  Update user information by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id      path      string            true  "User ID"
// @Param        request body      UpdateUserRequest true  "User update data"
// @Success      200     {object}  map[string]interface{}
// @Failure      400     {object}  map[string]interface{}
// @Failure      404     {object}  map[string]interface{}
// @Router       /users/{id} [put]
// @Security     BearerAuth
func UpdateUser(c *fiber.Ctx) error {
	// ✅ แปลง string เป็น int
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.ErrorResponse(c, "Invalid user ID", fiber.StatusBadRequest)
	}

	var userData UpdateUserRequest
	if err := c.BodyParser(&userData); err != nil {
		return utils.ErrorResponse(c, "Invalid input", fiber.StatusBadRequest)
	}

	// ✅ แปลง request เป็น UpdateUserInput
	input := models.UpdateUserInput{
		UserCode: userData.UserCode,
		Email:    userData.Email,
		FNameT:   userData.FNameT,
		LNameT:   userData.LNameT,
		FNameE:   userData.FNameE,
		LNameE:   userData.LNameE,
	}

	// ✅ เชื่อมกับ database จริง
	user, err := services.UpdateUser(id, input)
	if err != nil {
		return utils.ErrorResponse(c, err.Error(), fiber.StatusInternalServerError)
	}

	return utils.SuccessResponse(c, user, "User updated successfully")
}

// DeleteUser godoc
// @Summary      Delete a user
// @Description  Delete a user account by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      404  {object}  map[string]interface{}
// @Router       /users/{id} [delete]
// @Security     BearerAuth
func DeleteUser(c *fiber.Ctx) error {
	// ✅ แปลง string เป็น int
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.ErrorResponse(c, "Invalid user ID", fiber.StatusBadRequest)
	}

	// ✅ เชื่อมกับ database จริง
	if err := services.DeleteUser(id); err != nil {
		return utils.ErrorResponse(c, err.Error(), fiber.StatusInternalServerError)
	}

	return utils.SuccessResponse(c, fiber.Map{
		"id": id,
	}, "User deleted successfully")
}
