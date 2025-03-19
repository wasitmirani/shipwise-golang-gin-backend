package main

import (
	"shipwise/config"
	database "shipwise/db"
	"shipwise/internal/app/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	db, err := database.InitPostgres(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}
	defer db.Close()

	r := gin.Default()
	routes.SetupRoutes(r)

	port := os.Getenv("PORT")
	log.Println("Starting server on port", port)
	if port == "" {
		port = "8080"
	}
	log.Fatal(r.Run(":" + port))
}
