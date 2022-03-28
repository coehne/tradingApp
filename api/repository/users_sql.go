package repository

import (
	"github.com/dakicka/tradingApp/api/database"
	"github.com/dakicka/tradingApp/api/entity"
	"github.com/gofiber/fiber/v2"
)

// UsersSQL wraps the SQL DB and implements the required operations.
type UsersSQL struct {
	*database.GormDB
}

// NewUsersSQL instanciates and returns a new users repository.
func NewUsersSQL(db *database.GormDB) Users {
	return &UsersSQL{db}
}

func (r *UsersSQL) Create(ctx fiber.Ctx, user entity.User) (entity.User, error) {

	// Insert into DB
	result := r.DB.Create(&user)

	// Check for errors during insertion
	if result.Error != nil {
		return entity.User{}, result.Error
	}

	return user, nil
}
func (r *UsersSQL) Get(ctx fiber.Ctx, user entity.User) (entity.User, error) {

	// Query from DB
	result := r.DB.Find(&user)

	// Check for errors during query
	if result.Error != nil {
		return entity.User{}, result.Error
	}

	return user, nil
}
func (r *UsersSQL) GetByEmail(ctx fiber.Ctx, user entity.User) (entity.User, error) {

	// Query from DB
	result := r.DB.First(&user, "email = ?", user.Email)

	// Check for errors during Query
	if result.Error != nil {
		return entity.User{}, result.Error
	}

	return user, nil
}
