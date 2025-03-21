package main

import (
	"log"
	"os"
	"shipwise/config"
	database "shipwise/db"
	"shipwise/db/migrations"
	"shipwise/internal/app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	// Initialize the database connection
	db, err := database.InitPostgres(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Run migrations by passing the db instance
	migrations.RunMigrations(db)

	r := gin.Default()
	routes.SetupRoutes(r)

	port := os.Getenv("PORT")
	log.Println("Starting server on port", port)
	if port == "" {
		port = "8080"
	}
	log.Fatal(r.Run(":" + port))
}
