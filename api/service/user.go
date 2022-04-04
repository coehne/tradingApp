package service

import (
	"strings"

	"github.com/dakicka/tradingApp/api/entity"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func (s Service) RegisterUser(firstName, email, password string) (entity.User, error) {

	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 14)

	u := entity.User{
		Email:     strings.ToLower(email),
		FirstName: firstName,
		Hash:      hash,
	}

	existingUser, err := s.users.GetByEmail(u)

	// If email is already in db, return 400
	if existingUser.ID != 0 {
		return entity.User{}, fiber.NewError(fiber.StatusBadRequest, "User already registered")
	}

	user, err := s.users.Create(u)
	if err != nil {
		return entity.User{}, errors.Wrapf(err, "could not create user with email %s", u.Email)
	}

	return user, nil
}

func (s Service) GetUserFromId(id uint) (entity.User, error) {

	// Build empty user object and assign id
	user := entity.User{
		ID: id,
	}

	// Pass down the architecture to get the user from user repository
	user, err := s.users.Get(user)
	if err != nil {
		errors.Wrapf(err, "could get user data from repo for user with id %f", id)
		return entity.User{}, err
	}

	return user, nil
}
func (s Service) Login(email, password string) (entity.User, error) {

	// Build empty user object and assign email
	user := entity.User{
		Email: email,
	}

	// Pass down the architecture to get the user from user repository
	user, err := s.users.GetByEmail(user)
	if err != nil {
		errors.Wrapf(err, "could not get user data from repo for user with email %s", email)
		return entity.User{}, err
	}

	// Validate Password with hash
	if err := bcrypt.CompareHashAndPassword(user.Hash, []byte(password)); err != nil {
		errors.Wrapf(err, "password incorrect for user with email %s", email)
		return entity.User{}, err
	}

	return user, nil
}
