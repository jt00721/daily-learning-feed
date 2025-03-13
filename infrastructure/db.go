package infrastructure

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/jt00721/daily-learning-feed/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using Railway environment variables")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = DB.AutoMigrate(&domain.Resource{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
	log.Println("Database initialized and migrated successfully")
}
