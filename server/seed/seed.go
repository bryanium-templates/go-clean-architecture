package main

import (
	"encoding/json"
	"log"
	"os"

	guitar_core "server/guitar/core"
	"server/infrastructure/database"

	"github.com/google/uuid"
)

func main() {
	// Database connection & migration
	database.ConnectDb()

	if err := database.DB.AutoMigrate(&guitar_core.Guitar{}); err != nil {
		log.Fatalf("DB Migration Failed: %v", err)
	}

	// JSON seed loading
	file, err := os.ReadFile("seed/guitar.json")
	if err != nil {
		log.Fatalf("Failed to read JSON file: %v", err)
	}


	// Seeding Starts
	var guitars []guitar_core.Guitar
	if err := json.Unmarshal(file, &guitars); err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	insertedCount := 0
	skippedCount := 0
	failedCount := 0

	for _, g := range guitars {
		var existing guitar_core.Guitar
		if err := database.DB.Where("name = ?", g.Name).First(&existing).Error; err == nil {
			log.Printf("Skipped existing guitar: %s\n", g.Name)
			skippedCount++
			continue
		}

		g.ID = uuid.New()
		if err := database.DB.Create(&g).Error; err != nil {
			log.Printf("Failed to insert %s: %v\n", g.Name, err)
			failedCount++
		} else {
			log.Printf("Inserted guitar: %s\n", g.Name)
			insertedCount++
		}
	}

	log.Println("🎉 Guitar seed process completed.")
	log.Printf("Summary: Inserted = %d | Skipped = %d | Failed = %d\n", insertedCount, skippedCount, failedCount)
}

