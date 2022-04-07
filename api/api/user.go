package api

import (
	"fmt"

	"github.com/dakicka/tradingApp/api/entity"
	"github.com/dakicka/tradingApp/api/package/auth"
	"github.com/dakicka/tradingApp/api/usecase"
	"github.com/gofiber/fiber/v2"
)

type userController struct {
	usecase.UseCases
}

// NewUser sets up a new user service with the given repositories, helpers and
// registers the corresponding routes.
func NewUser(app *fiber.App, service usecase.UseCases) {
	ctr := userController{service}

	endpointIdentity := app.Group("/api/identity")
	endpointIdentity.Post("/signup", ctr.register)
	endpointIdentity.Get("/me", ctr.getMe)
	endpointIdentity.Post("/login", ctr.login)
	endpointIdentity.Post("/logout", ctr.logout)

}

func (ctr *userController) register(ctx *fiber.Ctx) error {

	var req registerReq

	err := ctx.BodyParser(&req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	user, err := ctr.RegisterUser(req.FirstName, req.Email, req.Password)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	// Remove pw for security
	req.Password = ""

	// Set cookie with accessToken
	auth.SetCookieForUser(ctx, user.ID)

	// Send response
	return ctx.SendStatus(fiber.StatusNoContent)
}

func (ctr *userController) getMe(ctx *fiber.Ctx) error {
	user := entity.User{}

	// Get the user Id from the accessToken inside the cookie
	userId, err := auth.GetUserIdByContext(ctx)
	if err != nil {
		fmt.Println("userId error")
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	if userId == 0 {
		return ctx.Status(fiber.StatusBadRequest).SendString("no user found")
	}

	user, err = ctr.GetUserFromId(userId)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	// Build response
	res := ctr.meResponse(user)

	// Send response
	return ctx.Status(fiber.StatusOK).JSON(res)
}

func (ctr *userController) login(ctx *fiber.Ctx) error {
	var req loginReq

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("bad request")
	}

	// Pass down the request data through the clean architecture shells to get user object
	user, err := ctr.Login(req.Email, req.Password)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("bad request")
	}

	// Set cookie with accessToken
	auth.SetCookieForUser(ctx, user.ID)

	// Build response
	res := ctr.meResponse(user)

	// Send response
	return ctx.Status(fiber.StatusOK).JSON(res)
}
func (ctr *userController) logout(ctx *fiber.Ctx) error {

	// Use an expired cookie to invalidate the users cookie
	auth.SetExpiredToken(ctx)

	return ctx.SendStatus(fiber.StatusNoContent)
}

// MeResponse takes in the user entity and only reponse with the necessary fields
func (ctr *userController) meResponse(u entity.User) meRes {
	return meRes{
		ID:        u.ID,
		FirstName: u.FirstName,
	}
}
