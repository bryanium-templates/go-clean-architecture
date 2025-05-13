package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	err := godotenv.Load()
	if err != nil{
		log.Fatal("Error loading the .env file")
	}

	dbstring := os.Getenv("DATABASE_URL")
	if dbstring == "" {
		log.Fatal("DATABASE_URL is undefined in the .env file")
	}

	database, err := gorm.Open(postgres.Open(dbstring), &gorm.Config{})
	if err != nil{
		log.Fatal("Failed to connect to the database", err)
	}

	DB = database
	log.Println("Database has been connected successfully")
}