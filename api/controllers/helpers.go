package controllers

import (
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
