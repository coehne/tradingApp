package auth

import (
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
)

var verySecretKey = os.Getenv("VERY_SECRET_KEY")

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

func GetUserIdFromToken(c *fiber.Ctx) (uint, error) {

	cookie := c.Cookies("accessToken")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(verySecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return 0, err
	}
	claims := token.Claims.(*jwt.StandardClaims)

	// Convert string to uint64
	ID, err := strconv.ParseUint(claims.Issuer, 10, 64)
	return uint(ID), nil
}
