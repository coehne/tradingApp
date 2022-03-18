package controllers

import (
	"time"

	"github.com/dakicka/tradingApp/api/database"
	"github.com/dakicka/tradingApp/api/models"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

const VERY_SECRET_KEY = "verySecretKey"

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	user := models.User{
		FirstName: data["firstName"],
		Email:     data["email"],
		Hash:      password,
	}

	database.DB.Create(&user)
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

	// Compare the probided hashed password with the hash from the db
	if err := bcrypt.CompareHashAndPassword(user.Hash, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: jwt.NewTime(15000),
		Issuer:    user.FirstName,
	}

	// Create token with claims
	tokenData := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenData.SignedString([]byte(VERY_SECRET_KEY))

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
