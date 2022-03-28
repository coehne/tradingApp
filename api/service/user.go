package service

import (
	"strings"
	"time"

	"github.com/dakicka/tradingApp/api/entity"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func (s Service) RegisterUser(ctx fiber.Ctx, firstName, email, password string) (entity.User, error) {

	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user := entity.User{
		Email:     strings.ToLower(email),
		FirstName: firstName,
		Hash:      hash,
	}

	s.users.Create(ctx, user)

	return user, nil
}

func (s Service) GetUserFromId(ctx fiber.Ctx, id uint) (entity.User, error) {
	// Build empty user object and assign id
	user := entity.User{
		ID: id,
	}

	// Pass down the architecture to get the user from user repository
	s.users.Get(ctx, user)

	return user, nil
}
func (s Service) Login(ctx fiber.Ctx, email, password string) (entity.User, error) {

	// Build empty user object and assign email
	user := entity.User{
		Email: email,
	}

	// Pass down the architecture to get the user from user repository
	s.users.GetByEmail(ctx, user)

	// Validate Password with hash
	if err := bcrypt.CompareHashAndPassword(user.Hash, []byte(password)); err != nil {
		return entity.User{}, err
	}

	return user, nil
}
func (s Service) Logout(ctx fiber.Ctx) error {

	// Set expired new cookie to invalidate the old one
	cookie := fiber.Cookie{
		Name:     "accessToken",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	ctx.Cookie(&cookie)

	return ctx.JSON(fiber.Map{
		"message": "success",
	})

}
