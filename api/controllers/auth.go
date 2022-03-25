package controllers

import (
	"os"
	"strconv"
	"time"

	"github.com/dakicka/tradingApp/api/database"
	"github.com/dakicka/tradingApp/api/models"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

var verySecretKey = os.Getenv("VERY_SECRET_KEY")

type SignupRequest struct {
	FirstName string `json:"firstName"`
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password"`
}

func Register(c *fiber.Ctx) error {
	var request SignupRequest

	if err := c.BodyParser(&request); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 14)
	user := models.User{
		FirstName: request.FirstName,
		Email:     request.Email,
		Hash:      password,
	}
	database.DB.Where("email = ?", request.Email).First(&user)

	// If email is already in db, return 400
	if user.ID != 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "User already registered",
		})
	}

	//
	database.DB.Create(&user)
	c.Status(fiber.StatusOK)
	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)

	// If no user is found, return 404
	if user.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "User not found",
		})
	}

	//TODO: Check if password was provided or not

	// Compare the probided hashed password with the hash from the db
	if err := bcrypt.CompareHashAndPassword(user.Hash, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: jwt.At(time.Now().Add(time.Hour * 24)),
		Issuer:    strconv.Itoa(int(user.ID)),
	}

	// Create token with claims
	tokenData := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenData.SignedString([]byte(verySecretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not login",
		})

	}

	cookie := fiber.Cookie{
		Name:     "accessToken",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})

}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("accessToken")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(verySecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "no valid cookie found",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	database.DB.Where("ID = ?", claims.Issuer).First(&user)

	return c.JSON(user)

}

func Logout(c *fiber.Ctx) error {
	// Set expired new cookie to invalidate the old one
	cookie := fiber.Cookie{
		Name:     "accessToken",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})

}
