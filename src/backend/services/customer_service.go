package services

import (
	"hppku-nextgen/backend/models"
	"hppku-nextgen/backend/repositories"
)

// CustomerService defines the interface for customer business logic.
type CustomerService interface {
	GetAllCustomers() ([]models.Customer, error)
}

type customerService struct {
	repo repositories.CustomerRepository
}

// NewCustomerService creates a new instance of CustomerService.
func NewCustomerService(repo repositories.CustomerRepository) CustomerService {
	return &customerService{repo: repo}
}

// GetAllCustomers retrieves all customers by calling the repository.
func (s *customerService) GetAllCustomers() ([]models.Customer, error) {
	return s.repo.FindAll()
}