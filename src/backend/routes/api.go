package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"hppku-nextgen/backend/handlers"
	"hppku-nextgen/backend/repositories"
	"hppku-nextgen/backend/services"
)

// SetupRoutes configures all the application routes.
func SetupRoutes(app *fiber.App, db *gorm.DB) {
	api := app.Group("/api")

	// --- Customer Routes ---
	customerRepo := repositories.NewCustomerRepository(db)
	customerService := services.NewCustomerService(customerRepo)
	customerHandler := handlers.NewCustomerHandler(customerService)
	customerHandler.SetupCustomerRoutes(api)

	// Other routes can be added here...
}
