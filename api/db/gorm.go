package db

import (
	"fmt"
	"log"
	"time"

	"github.com/dakicka/tradingApp/api/config"
	"github.com/dakicka/tradingApp/api/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type GormDB struct {
	DB     *gorm.DB
	Config *config.Config
}

func NewGorm(config *config.Config) *GormDB {
	var err error
	db := &GormDB{}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", config.DB.Host, config.DB.User, config.DB.Password, config.DB.DBName, config.DB.Port)

	db.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		NowFunc: func() time.Time {
			return time.Now().UTC()

		},
	})
	if err != nil {
		panic(err)
	}
	log.Println("Successfully connected to database")

	// How verbose the terminal log should be
	db.DB.Logger = logger.Default.LogMode(logger.Info)

	return db
}

func (d *GormDB) Migrate() {
	err := d.DB.AutoMigrate(&entity.User{})
	if err != nil {
		panic(err)
	}
}
