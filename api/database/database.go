package database

import "github.com/dakicka/tradingApp/api/config"

type DB interface {
	Migrate()
}

func New(config *config.Config) *GormDB {
	return NewGorm(config)
}
