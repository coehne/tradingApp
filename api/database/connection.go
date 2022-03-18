package database

import (
	"log"

	"github.com/dakicka/tradingApp/api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(sqlite.Open("database/db.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database\n", err.Error())
	}
	DB = connection

	log.Println("Successfully connected to database")

	// How verbose the terminal log should be
	DB.Logger = logger.Default.LogMode(logger.Info)

	// Add auto migrations
	DB.AutoMigrate(&models.User{}, &models.Trade{}, &models.Transaction{})

}
