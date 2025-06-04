package infrastructure

import (
	"log"

	"github.com/bryanium-templates/go-clean-architecture/models"
)

func SyncDatabase() {
	log.Println("Syncing declared database schema...")

	err := DB.AutoMigrate(
		&models.User{},
		&models.Guitar{},
	)

	if err != nil {
		log.Fatalf("Failed to auto-migrate the provided schema: %v", err)
	}
	log.Println("Database migration is successful")
}