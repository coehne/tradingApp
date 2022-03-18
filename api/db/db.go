package db

import (
	"log"

	"github.com/dakicka/tradingApp/api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("db.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database\n", err.Error())
	}

	log.Println("Successfully connected to database")

	// How verbose the terminal log should be
	db.Logger = logger.Default.LogMode(logger.Info)

	// Add auto migrations
	db.AutoMigrate(&models.User{}, &models.Trade{}, &models.Transaction{})

	Database = DbInstance{Db: db}
}
