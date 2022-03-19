package routes

import (
	"github.com/dakicka/tradingApp/api/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// Identity endpoint
	app.Post("/api/identity/register", controllers.Register)
	app.Post("/api/identity/login", controllers.Login)
	app.Get("/api/identity/user", controllers.User)
	app.Post("/api/identity/logout", controllers.Logout)

	// Trade endpoint

	// Transaction endpoint
	app.Post("/api/transaction", controllers.DepositCash)

	// Quote endpoint
	app.Get("/api/quote", controllers.GetQuote)
}
