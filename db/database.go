package database

import (
	"fmt"
	"shipwise/config"

	"gorm.io/driver/postgres"

	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitPostgres initializes the PostgreSQL database connection using GORM
func InitPostgres(cfg config.Config) (*gorm.DB, error) {
	// Validate the configuration values
	if cfg.DBHost == "" || cfg.DBPort == "" || cfg.DBUser == "" || cfg.DBPassword == "" || cfg.DBName == "" {
		return nil, fmt.Errorf("database configuration is incomplete: host=%s port=%s user=%s dbname=%s",
			cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBName)
	}

	// Create the DSN (Data Source Name)
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	// Open a connection to the database using GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Assign the GORM DB instance to the global DB variable
	DB = db

	return db, nil
}
