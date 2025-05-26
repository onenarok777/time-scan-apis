package gorm

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// check env for use
	requiredEnv := []string{"DATABASE_HOST", "DATABASE_USER", "DATABASE_PASSWORD", "DATABASE_NAME", "DATABASE_PORT", "DATABASE_SSL_MODE"}
	for _, env := range requiredEnv {
		if os.Getenv(env) == "" {
			log.Fatalf("❌ Missing required environment variable: %s", env)
		}
	}
	// set dsn url connect db
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_NAME"),
		os.Getenv("DATABASE_PORT"),
	)
	// open gorm
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}
	DB = db
	log.Println("✅ Connected to database successfully!")
}
