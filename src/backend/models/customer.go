package models

import "gorm.io/gorm"

// Customer represents the model for a customer
// gorm.Model provides fields like ID, CreatedAt, UpdatedAt, DeletedAt
type Customer struct {
	gorm.Model
	Name    string `json:"name" gorm:"not null"`
	Email   string `json:"email" gorm:"unique;not null"`
	Phone   string `json:"phone"`
}
