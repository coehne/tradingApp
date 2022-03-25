package controllers

import (
	"strconv"
	"time"

	"github.com/dakicka/tradingApp/api/database"
	"github.com/dakicka/tradingApp/api/models"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
)

func GetUserFromToken(c *fiber.Ctx) (models.User, error) {

	cookie := c.Cookies("accessToken")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(verySecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return models.User{}, err
	}

	claims := token.Claims.(*jwt.StandardClaims)
	var user models.User

	database.DB.Where("ID = ?", claims.Issuer).First(&user)
	return user, nil
}

func SetCookieForUser(c *fiber.Ctx, id uint) error {
	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: jwt.At(time.Now().Add(time.Hour * 24)),
		Issuer:    strconv.Itoa(int(id)),
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
	return nil
}
