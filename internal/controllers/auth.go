package controllers

import (
	"WealthNoteBackend/internal/models"
	"WealthNoteBackend/internal/services"
	"WealthNoteBackend/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

type LoginRequest struct {
	UserCode string `json:"usercode" validate:"required" example:"user001"`
	Password string `json:"password" validate:"required" example:"password123"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	RefreshToken string `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	ExpiresIn    int    `json:"expires_in" example:"900"`
}

type RegisterRequest struct {
	UserCode string  `json:"usercode" validate:"required" example:"user001"`
	Email    string  `json:"email" validate:"required,email" example:"user@example.com"`
	FNameT   *string `json:"fnamet,omitempty" example:"สมชาย"`
	LNameT   *string `json:"lnamet,omitempty" example:"ใจดี"`
	FNameE   *string `json:"fnamee,omitempty" example:"Somchai"`
	LNameE   *string `json:"lnamee,omitempty" example:"Jaidee"`
	Password string  `json:"password" validate:"required,min=6" example:"password123"`
}

// Login godoc
// @Summary      User login
// @Description  Authenticate user with usercode and password, return JWT tokens
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body LoginRequest true "Login credentials"
// @Success      200  {object}  LoginResponse
// @Failure      400  {object}  map[string]interface{}
// @Failure      401  {object}  map[string]interface{}
// @Router       /auth/login [post]
func Login(c *fiber.Ctx) error {
	var input LoginRequest

	if err := c.BodyParser(&input); err != nil {
		return utils.ErrorResponse(c, "Invalid input", fiber.StatusBadRequest)
	}

	// เรียกใช้ service เพื่อ login
	accessToken, refreshToken, err := services.Login(input.UserCode, input.Password)
	if err != nil {
		return utils.ErrorResponse(c, err.Error(), fiber.StatusUnauthorized)
	}

	return utils.SuccessResponse(c, LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    15 * 60, // 15 minutes
	}, "Login successful")
}

// Register godoc
// @Summary      User registration
// @Description  Register a new user with usercode, email and password
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body RegisterRequest true "Registration data"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Router       /auth/register [post]
func Register(c *fiber.Ctx) error {
	var input models.RegisterInput

	if err := c.BodyParser(&input); err != nil {
		return utils.ErrorResponse(c, "Invalid input", fiber.StatusBadRequest)
	}

	// เรียกใช้ service เพื่อสมัครสมาชิก
	user, err := services.Register(input)
	if err != nil {
		return utils.ErrorResponse(c, err.Error(), fiber.StatusBadRequest)
	}

	return utils.SuccessResponse(c, user, "User registered successfully")
}
