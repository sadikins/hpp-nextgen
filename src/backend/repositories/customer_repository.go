package repositories

import (
	"hppku-nextgen/backend/models"
	"gorm.io/gorm"
)

// CustomerRepository defines the interface for customer data operations.
type CustomerRepository interface {
	FindAll() ([]models.Customer, error)
}

type customerRepository struct {
	db *gorm.DB
}

// NewCustomerRepository creates a new instance of CustomerRepository.
func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{db: db}
}

// FindAll retrieves all customers from the database.
func (r *customerRepository) FindAll() ([]models.Customer, error) {
	var customers []models.Customer
	err := r.db.Find(&customers).Error
	return customers, err
}