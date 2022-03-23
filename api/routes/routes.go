package routes

import (
	"github.com/dakicka/tradingApp/api/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// Identity endpoint
	app.Post("/api/identity/signup", controllers.Register)
	app.Post("/api/identity/login", controllers.Login)
	app.Get("/api/identity/me", controllers.User)
	app.Post("/api/identity/logout", controllers.Logout)

	// Trade endpoint
	app.Post("api/trade", controllers.CreateTrade)
	app.Get("api/trade", controllers.GetTrades)
	app.Get("api/trade/:id", controllers.GetTrade)

	// Transaction endpoint
	app.Post("/api/transaction", controllers.CreateTransaction)
	app.Get("/api/transaction", controllers.GetTransactions)

	// Quote endpoint
	app.Get("/api/quote", controllers.GetQuote)
}
