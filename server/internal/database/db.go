package database

import (
	"accountantapp/go-service/internal/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"accountantapp/go-service/internal/models"
)

var DB *gorm.DB

func Connect() {
	cfg := config.Load()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	// Авто-миграция моделей
	err = db.AutoMigrate(
		&models.User{},
		&models.Account{},
		&models.Transaction{},
		&models.Category{},
		&models.Template{},
		&models.AppSettings{},
		&models.Transfer{},
	)
	if err != nil {
		panic("Failed to auto-migrate models!")
	}

	DB = db
	fmt.Println("Database connected and migrated")
}
