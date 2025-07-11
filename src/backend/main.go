// src/backend/main.go
package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"hppku-nextgen/backend/config"
	"hppku-nextgen/backend/database"
	"hppku-nextgen/backend/routes"
)

func main() {
	// Load configuration from .env file
	config.LoadConfig()

	// Connect to the database
	database.Connect()

	// Create Fiber app
	app := fiber.New()

	// Setup Middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173", // Default Vite dev port
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Setup routes
	routes.SetupRoutes(app, database.DB)

	// Welcome route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to HPPKu NextGen Backend! Refactored!")
	})

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(app.Listen(":" + port))
}