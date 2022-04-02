package repository

import (
	"github.com/dakicka/tradingApp/api/entity"
	"github.com/dakicka/tradingApp/api/package/db"
)

// UsersSQL wraps the SQL DB and implements the required operations.
type UsersSQL struct {
	db.GormDB
}

// NewUsersSQLRepo instanciates and returns a new users repository.
func NewUsersSQLRepo(db *db.GormDB) Users {
	return UsersSQL{*db}
}

// Create inserts a new user into the db
func (r UsersSQL) Create(user entity.User) (entity.User, error) {

	// Insert into DB
	result := r.DB.Create(&user)

	// Check for errors during insertion
	if result.Error != nil {
		return entity.User{}, result.Error
	}

	user = entity.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		Email:     user.Email,
	}

	return user, nil
}
func (r UsersSQL) Get(user entity.User) (entity.User, error) {

	// Query from DB
	result := r.DB.Find(&user)

	// Check for errors during query
	if result.Error != nil {
		return entity.User{}, result.Error
	}

	return user, nil
}
func (r UsersSQL) GetByEmail(user entity.User) (entity.User, error) {

	// Query from DB
	result := r.DB.First(&user, "email = ?", user.Email)

	// Check for errors during Query
	if result.Error != nil {
		return entity.User{}, result.Error
	}

	return user, nil
}
