package middleware

import (
	"strings"

	"WealthNoteBackend/pkg/jwt"
	"WealthNoteBackend/pkg/utils"
	"log"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return utils.ErrorResponse(c, "Authorization header is missing", fiber.StatusUnauthorized)
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		claims, err := jwt.ValidateToken(tokenString)
		if err != nil {
			return utils.ErrorResponse(c, "Invalid or expired token", fiber.StatusUnauthorized)
		}

		// ✅ ตอนนี้ใช้งานได้แล้ว เพราะ claims เป็น *CustomClaims
		c.Locals("user", claims)
		c.Locals("user_id", claims.UserID)
		log.Println("Authenticated user ID:", claims.UserID)
		return c.Next()
	}
}
