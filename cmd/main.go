package main

import (
	"log"

	"github.com/damshxy/xyz-finance-app/config"
	"github.com/damshxy/xyz-finance-app/database"
	"github.com/damshxy/xyz-finance-app/internal/delivery/routes"
	"github.com/damshxy/xyz-finance-app/internal/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func main() {
	// Load application configuration
	config := config.LoadConfig()

	// Initialize database
	db, err := database.PostgresInit(config)
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
		panic(err)
	}

	// Auto migrate tables
	err = autoMigrateTables(db)
	if err != nil {
		log.Fatalf("Error migrating tables: %v", err)
	}

	// Setup Fiber
	app := fiber.New()

	// Setup routes
	routes.RoutesInit(app)

	// Start Server
	log.Println("Starting server on :4000")
	log.Fatal(app.Listen(":4000"))
}

func autoMigrateTables(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Consumer{},
		&models.Transaction{},
	)
}