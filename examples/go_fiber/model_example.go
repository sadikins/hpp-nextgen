// examples/go_fiber/model_example.go
package models

import (
	"time"

	"gorm.io/gorm"
)

// RawMaterial represents a raw material used in food production.
type RawMaterial struct {
	gorm.Model
	Code        string `gorm:"type:varchar(50);uniqueIndex;not null" json:"code" validate:"required,min=3,max=50"`
	Name        string `gorm:"type:varchar(255);not null" json:"name" validate:"required,min=3,max=255"`
	Unit        string `gorm:"type:varchar(50);not null" json:"unit" validate:"required,min=1,max=50"` // e.g., "gram", "kg", "pcs"
	UnitPrice   float64 `gorm:"type:numeric;not null" json:"unit_price" validate:"required,gt=0"`
	Description string `gorm:"type:text" json:"description"`
}

// TableName overrides the table name for RawMaterial.
func (RawMaterial) TableName() string {
	return "raw_materials"
}

// BeforeCreate hook to set CreatedAt and UpdatedAt
func (rm *RawMaterial) BeforeCreate(tx *gorm.DB) (err error) {
    rm.CreatedAt = time.Now()
    rm.UpdatedAt = time.Now()
    return
}

// BeforeUpdate hook to set UpdatedAt
func (rm *RawMaterial) BeforeUpdate(tx *gorm.DB) (err error) {
    rm.UpdatedAt = time.Now()
    return
}