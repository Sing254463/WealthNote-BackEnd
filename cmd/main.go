package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"

	_ "WealthNoteBackend/docs" // Import the generated docs
	"WealthNoteBackend/internal/config"
	"WealthNoteBackend/internal/database"
	"WealthNoteBackend/internal/middleware"
	"WealthNoteBackend/internal/routes"
)

// @title           Go Fiber API
// @version         1.0
// @description     This is a Go Fiber API with JWT authentication and multiple database support
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3000
// @BasePath  /api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	// Load configuration
	config.Load()

	// Initialize databases based on configuration
	initDatabases()

	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// Middleware
	app.Use(logger.New())
	app.Use(middleware.Cors())

	// Swagger
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Setup routes
	routes.SetupRoutes(app)

	// Start server
	port := config.AppConfig.Port
	log.Printf("Server starting on port %s", port)
	log.Printf("Swagger UI available at: http://localhost:%s/swagger/index.html", port)
	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}

func initDatabases() {
	if config.AppConfig.EnableMongoDB {
		database.ConnectMongoDB()
		log.Println("MongoDB enabled and connected")
	}

	if config.AppConfig.EnableMySQL && config.AppConfig.DatabaseType == 2 {
		database.Connect()
		log.Println("MySQL enabled and connected")
	}

	if config.AppConfig.EnablePostgreSQL && config.AppConfig.DatabaseType == 1 {
		database.ConnectPostgres()
		log.Println("PostgreSQL enabled and connected")

		// Run migrations
		database.RunMigrations()
		log.Println("Database migrations completed")
	}

	if config.AppConfig.EnableSQLServer && config.AppConfig.DatabaseType == 1 {
		database.ConnectSQLServer()
		log.Println("SQL Server enabled and connected")
	}
}
