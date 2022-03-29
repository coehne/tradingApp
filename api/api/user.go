package api

import (
	"github.com/dakicka/tradingApp/api/auth"
	"github.com/dakicka/tradingApp/api/entity"
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
	endpointIdentity.Post("/signup", ctr.create)
	endpointIdentity.Get("/me", ctr.get)
	endpointIdentity.Post("/login", ctr.login)
	endpointIdentity.Post("/logout", ctr.logout)

}

func (ctr *userController) create(ctx *fiber.Ctx) error {

	var req registerReq

	// TODO: add validation here
	err := ctx.BodyParser(&req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}
	// Pass down the user object through the clean architecture shells
	user, err := ctr.RegisterUser(req.FirstName, req.Email, req.Password)

	// Check if everything went well down the line
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	// Remove clear text pw for security
	req.Password = ""

	// Set cookie with accessToken
	auth.SetCookieForUser(ctx, user.ID)

	// Build response
	res := ctr.meResponse(user)

	// Send response
	return ctx.Status(fiber.StatusOK).JSON(res)
}

func (ctr *userController) get(ctx *fiber.Ctx) error {
	user := entity.User{}

	// Get the user Id from the accessToken inside the cookie
	userId, err := auth.GetUserIdFromToken(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	// Pass down the user object through the clean architecture shells
	user, err = ctr.GetUserFromId(userId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	// Build response
	res := ctr.meResponse(user)

	// Send response
	return ctx.Status(fiber.StatusOK).JSON(res)
}
func (ctr *userController) login(ctx *fiber.Ctx) error {
	var req loginReq

	// TODO: add validation here
	if err := ctx.BodyParser(&req); err != nil {

		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	// Pass down the request data through the clean architecture shells to get user object
	user, err := ctr.Login(req.Email, req.Password)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}
	// Set cookie with accessToken
	auth.SetCookieForUser(ctx, user.ID)

	// Build response
	res := ctr.meResponse(user)

	// Send response
	return ctx.Status(fiber.StatusOK).JSON(res)
}
func (ctr *userController) logout(ctx *fiber.Ctx) error {

	// Call logout function from user service
	ctr.logout(ctx)

	return ctx.JSON(fiber.Map{
		"message": "success",
	})
}

// MeResponse takes in the user entity and only reponse with the necessary fields
func (ctr *userController) meResponse(u entity.User) meRes {
	return meRes{
		ID:        u.ID,
		FirstName: u.FirstName,
	}
}
