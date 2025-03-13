package infrastructure

import (
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

	dsn := os.Getenv("DATABASE_URL") + "?sslmode=require"
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set")
	}

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
