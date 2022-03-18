package routes

import (
	"github.com/dakicka/tradingApp/api/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// Welcome endpoint
	app.Post("/api/identity/register", controllers.Register)
	app.Post("/api/identity/login", controllers.Login)

	// User endpoint

}
