// examples/go_fiber/handler_example.go
package handlers

import (
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/yourusername/hppku-nextgen/backend/models"   // Sesuaikan path modul
	"github.com/yourusername/hppku-nextgen/backend/services" // Sesuaikan path modul
)

// RawMaterialHandler handles HTTP requests for raw materials.
type RawMaterialHandler struct {
	service  services.RawMaterialService
	validate *validator.Validate
}

// NewRawMaterialHandler creates a new instance of RawMaterialHandler.
func NewRawMaterialHandler(s services.RawMaterialService) *RawMaterialHandler {
	return &RawMaterialHandler{
		service:  s,
		validate: validator.New(),
	}
}

// CreateRawMaterial handles POST /api/v1/raw-materials
func (h *RawMaterialHandler) CreateRawMaterial(c *fiber.Ctx) error {
	rm := new(models.RawMaterial)
	if err := c.BodyParser(rm); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Validate input
	if err := h.validate.Struct(rm); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	createdRM, err := h.service.CreateRawMaterial(rm)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(createdRM)
}

// GetRawMaterialByID handles GET /api/v1/raw-materials/:id
func (h *RawMaterialHandler) GetRawMaterialByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	rm, err := h.service.GetRawMaterialByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if rm == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Raw material not found"})
	}

	return c.JSON(rm)
}

// GetAllRawMaterials handles GET /api/v1/raw-materials
func (h *RawMaterialHandler) GetAllRawMaterials(c *fiber.Ctx) error {
	rms, err := h.service.GetAllRawMaterials()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(rms)
}

// UpdateRawMaterial handles PUT /api/v1/raw-materials/:id
func (h *RawMaterialHandler) UpdateRawMaterial(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	rm := new(models.RawMaterial)
	if err := c.BodyParser(rm); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
    rm.ID = uint(id) // Ensure ID from URL is used

	if err := h.validate.Struct(rm); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	updatedRM, err := h.service.UpdateRawMaterial(rm)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(updatedRM)
}

// DeleteRawMaterial handles DELETE /api/v1/raw-materials/:id
func (h *RawMaterialHandler) DeleteRawMaterial(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	err = h.service.DeleteRawMaterial(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusNoContent).Send(nil) // 204 No Content
}