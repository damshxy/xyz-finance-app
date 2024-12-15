package database

import (
	"fmt"
	"log"
	"time"

	"github.com/damshxy/xyz-finance-app/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// PostgresInit initializes the PostgreSQL database connection.
func PostgresInit(c *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
		c.PGHost,
		c.PGPort,
		c.PGUser,
		c.PGPassword,
		c.PGDatabase,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return nil, err
	}

	// Test the database connection
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
		return nil, err
	}

	// Ping the database to ensure the connection is valid
	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
		return nil, err
	}

	// Configure connection pool
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	log.Println("Database connection established successfully.")
	return db, nil
}
