package config

import (
	"log"

	"github.com/harsh6373/chat-app-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := GetEnv("DATABASE_URL", "")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	DB = db

	DB.AutoMigrate(&models.User{}, &models.Message{})
}
