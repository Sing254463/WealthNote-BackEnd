package routes

import (
	"WealthNoteBackend/internal/controllers"
	"WealthNoteBackend/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Health check
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "OK",
			"message": "Server is running",
		})
	})

	// Auth routes
	auth := api.Group("/auth")
	auth.Post("/login", controllers.Login)
	auth.Post("/register", controllers.Register)

	// User routes
	users := api.Group("/users")
	users.Use(middleware.AuthMiddleware()) // ⚠️ ต้องมีบรรทัดนี้
	users.Get("/", controllers.GetAllUsers)
	users.Get("/:id", controllers.GetUserByID)
	users.Put("/:id", controllers.UpdateUser)
	users.Delete("/:id", controllers.DeleteUser)

	transaction := api.Group("/transactions")
	transaction.Use(middleware.AuthMiddleware()) // ⚠️ ต้องมีบรรทัดนี้
	transaction.Get("/", controllers.GetTransactionAll)
	transaction.Post("/", controllers.CreateTransaction)
	transaction.Put("/:id", controllers.UpdateTransaction)

}
