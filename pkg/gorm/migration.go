package gorm

import (
	"log"
	"time-scan/internal/users/models/tables"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&tables.User{},
	)

	if err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	log.Println("✅ Migration completed.")
}
