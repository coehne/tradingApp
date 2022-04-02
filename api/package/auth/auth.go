package auth

import (
	"strconv"
	"time"

	"github.com/pkg/errors"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

var JWT_KEY = viper.GetString("JWT_KEY")

func SetCookieForUser(ctx *fiber.Ctx, id uint) error {
	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: jwt.At(time.Now().Add(time.Hour * 24)),
		Issuer:    strconv.Itoa(int(id)),
	}

	// Create token with claims
	tokenData := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenData.SignedString([]byte(JWT_KEY))

	if err != nil {
		errors.Wrapf(err, "could  not sign token with claims")
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	cookie := fiber.Cookie{
		Name:     "accessToken",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	ctx.Cookie(&cookie)
	return nil
}

func GetUserIdFromToken(ctx *fiber.Ctx) (uint, error) {

	cookie := ctx.Cookies("accessToken")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWT_KEY), nil
	})

	if err != nil {
		errors.Wrapf(err, "could not get token from cookie")
		return 0, ctx.SendStatus(fiber.StatusInternalServerError)
	}
	claims := token.Claims.(*jwt.StandardClaims)

	// Convert string to uint64
	ID, err := strconv.ParseUint(claims.Issuer, 10, 64)
	return uint(ID), nil
}

func SetExpiredToken(ctx *fiber.Ctx) {
	// Set expired new cookie to invalidate the old one
	cookie := fiber.Cookie{
		Name:     "accessToken",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	ctx.Cookie(&cookie)
}
