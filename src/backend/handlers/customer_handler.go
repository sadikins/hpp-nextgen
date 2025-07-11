package handlers

import (
	"hppku-nextgen/backend/services"
	"github.com/gofiber/fiber/v2"
)

// CustomerHandler handles HTTP requests for customers.
type CustomerHandler struct {
	service services.CustomerService
}

// NewCustomerHandler creates a new instance of CustomerHandler.
func NewCustomerHandler(service services.CustomerService) *CustomerHandler {
	return &CustomerHandler{service: service}
}

// GetCustomers handles the request to get all customers.
func (h *CustomerHandler) GetCustomers(c *fiber.Ctx) error {
	customers, err := h.service.GetAllCustomers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not retrieve customers",
		})
	}
	return c.JSON(customers)
}

// SetupCustomerRoutes configures the routes for the customer feature.
func (h *CustomerHandler) SetupCustomerRoutes(api fiber.Router) {
	api.Get("/customers", h.GetCustomers)
}