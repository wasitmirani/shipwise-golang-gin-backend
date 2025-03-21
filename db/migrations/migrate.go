package migrations

import (
	"log"
	"shipwise/internal/app/models"
	"gorm.io/gorm"
)

// RunMigrations runs the database migrations for the given GORM DB instance
func RunMigrations(db *gorm.DB) {
	// Auto migrate the User model
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Error running user migrations: %v", err)
	}

	log.Println("Migrations completed successfully!")
}