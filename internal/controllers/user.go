package controllers

import (
	"WealthNoteBackend/pkg/utils"

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
	// TODO: Implement with actual database
	users := []fiber.Map{
		{
			"id":    "1",
			"name":  "John Doe",
			"email": "john@example.com",
		},
		{
			"id":    "2",
			"name":  "Jane Smith",
			"email": "jane@example.com",
		},
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
	userID := c.Params("id")

	// TODO: Implement with actual database
	user := fiber.Map{
		"id":    userID,
		"name":  "John Doe",
		"email": "john@example.com",
	}

	return utils.SuccessResponse(c, user, "User retrieved successfully")
}

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required" example:"John Doe"`
	Email    string `json:"email" validate:"required,email" example:"john@example.com"`
	Password string `json:"password" validate:"required,min=6" example:"password123"`
}

// CreateUser godoc
// @Summary      Create a new user
// @Description  Create a new user account
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request body CreateUserRequest true "User creation data"
// @Success      201  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Router       /users [post]
// @Security     BearerAuth
func CreateUser(c *fiber.Ctx) error {
	var userData CreateUserRequest

	if err := c.BodyParser(&userData); err != nil {
		return utils.ErrorResponse(c, "Invalid input", fiber.StatusBadRequest)
	}

	// TODO: Implement with actual database
	user := fiber.Map{
		"id":    "new-user-id",
		"name":  userData.Name,
		"email": userData.Email,
	}

	return utils.SuccessResponse(c, user, "User created successfully")
}

type UpdateUserRequest struct {
	Name  string `json:"name" validate:"required" example:"John Doe"`
	Email string `json:"email" validate:"required,email" example:"john@example.com"`
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
	userID := c.Params("id")

	var userData UpdateUserRequest

	if err := c.BodyParser(&userData); err != nil {
		return utils.ErrorResponse(c, "Invalid input", fiber.StatusBadRequest)
	}

	// TODO: Implement with actual database
	user := fiber.Map{
		"id":    userID,
		"name":  userData.Name,
		"email": userData.Email,
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
	userID := c.Params("id")

	// TODO: Implement with actual database

	return utils.SuccessResponse(c, fiber.Map{
		"id": userID,
	}, "User deleted successfully")
}
