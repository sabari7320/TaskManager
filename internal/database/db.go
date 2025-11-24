package database

import (
	"fmt"
	"log"
	"os"

	"example.com/task_manager/internal/models"
	// "github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Println("Error loading .env file:", err)
	// }
	// Read database URL from environment variable
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	// Connect to PostgreSQL
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto Migrations
	err = db.AutoMigrate(&models.User{}, &models.Task{})
	if err != nil {
		log.Fatal("Migration error:", err)
	}

	DB = db
	fmt.Println("Connected to PostgreSQL successfully!")
}
