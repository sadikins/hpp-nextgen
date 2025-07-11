package database

import (
	
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"hppku-nextgen/backend/models"
)

var DB *gorm.DB

// Connect connects to the database and runs migrations.
func Connect() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL not set")
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connection successful.")

	// Run migrations
	log.Println("Running migrations...")
	if err := DB.AutoMigrate(&models.Customer{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Migrations completed.")
}
