// examples/go_fiber/repository_example.go
package repositories

import (
	"github.com/yourusername/hppku-nextgen/backend/models" // Sesuaikan path modul
	"gorm.io/gorm"
)

// RawMaterialRepository defines the interface for raw material data operations.
type RawMaterialRepository interface {
	Create(rm *models.RawMaterial) error
	FindByID(id uint) (*models.RawMaterial, error)
	FindAll() ([]models.RawMaterial, error)
	Update(rm *models.RawMaterial) error
	Delete(id uint) error
}

// rawMaterialRepositoryImpl implements RawMaterialRepository using GORM.
type rawMaterialRepositoryImpl struct {
	db *gorm.DB
}

// NewRawMaterialRepository creates a new instance of RawMaterialRepository.
func NewRawMaterialRepository(db *gorm.DB) RawMaterialRepository {
	return &rawMaterialRepositoryImpl{db: db}
}

// Create creates a new raw material record in the database.
func (r *rawMaterialRepositoryImpl) Create(rm *models.RawMaterial) error {
	return r.db.Create(rm).Error
}

// FindByID retrieves a raw material by its ID.
func (r *rawMaterialRepositoryImpl) FindByID(id uint) (*models.RawMaterial, error) {
	var rm models.RawMaterial
	err := r.db.First(&rm, id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil // Or return a custom "not found" error
	}
	return &rm, err
}

// FindAll retrieves all raw material records.
func (r *rawMaterialRepositoryImpl) FindAll() ([]models.RawMaterial, error) {
	var rms []models.RawMaterial
	err := r.db.Find(&rms).Error
	return rms, err
}

// Update updates an existing raw material record.
func (r *rawMaterialRepositoryImpl) Update(rm *models.RawMaterial) error {
	return r.db.Save(rm).Error
}

// Delete deletes a raw material record by its ID.
func (r *rawMaterialRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&models.RawMaterial{}, id).Error
}