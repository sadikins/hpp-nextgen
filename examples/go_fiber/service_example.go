// examples/go_fiber/service_example.go
package services

import (
	"errors"

	"github.com/yourusername/hppku-nextgen/backend/models"       // Sesuaikan path modul
	"github.com/yourusername/hppku-nextgen/backend/repositories" // Sesuaikan path modul
)

// RawMaterialService defines the interface for raw material business logic.
type RawMaterialService interface {
	CreateRawMaterial(rm *models.RawMaterial) (*models.RawMaterial, error)
	GetRawMaterialByID(id uint) (*models.RawMaterial, error)
	GetAllRawMaterials() ([]models.RawMaterial, error)
	UpdateRawMaterial(rm *models.RawMaterial) (*models.RawMaterial, error)
	DeleteRawMaterial(id uint) error
}

// rawMaterialServiceImpl implements RawMaterialService.
type rawMaterialServiceImpl struct {
	repo repositories.RawMaterialRepository
}

// NewRawMaterialService creates a new instance of RawMaterialService.
func NewRawMaterialService(repo repositories.RawMaterialRepository) RawMaterialService {
	return &rawMaterialServiceImpl{repo: repo}
}

// CreateRawMaterial handles the creation of a new raw material.
func (s *rawMaterialServiceImpl) CreateRawMaterial(rm *models.RawMaterial) (*models.RawMaterial, error) {
	// Add any business logic/validation here before saving
	if rm.UnitPrice <= 0 {
		return nil, errors.New("unit price must be positive")
	}
	err := s.repo.Create(rm)
	if err != nil {
		return nil, err
	}
	return rm, nil
}

// GetRawMaterialByID retrieves a raw material by its ID.
func (s *rawMaterialServiceImpl) GetRawMaterialByID(id uint) (*models.RawMaterial, error) {
	return s.repo.FindByID(id)
}

// GetAllRawMaterials retrieves all raw materials.
func (s *rawMaterialServiceImpl) GetAllRawMaterials() ([]models.RawMaterial, error) {
	return s.repo.FindAll()
}

// UpdateRawMaterial handles updating an existing raw material.
func (s *rawMaterialServiceImpl) UpdateRawMaterial(rm *models.RawMaterial) (*models.RawMaterial, error) {
	existing, err := s.repo.FindByID(rm.ID)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, errors.New("raw material not found")
	}
	// Copy fields to update (or use specific update methods in repo)
	existing.Code = rm.Code
	existing.Name = rm.Name
	existing.Unit = rm.Unit
	existing.UnitPrice = rm.UnitPrice
	existing.Description = rm.Description

	err = s.repo.Update(existing)
	if err != nil {
		return nil, err
	}
	return existing, nil
}

// DeleteRawMaterial handles deleting a raw material.
func (s *rawMaterialServiceImpl) DeleteRawMaterial(id uint) error {
	return s.repo.Delete(id)
}