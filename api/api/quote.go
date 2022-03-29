package api

import (
	"github.com/dakicka/tradingApp/api/integration"
	"github.com/dakicka/tradingApp/api/usecase"
	"github.com/gofiber/fiber/v2"
)

type quoteController struct {
	usecase.UseCases
}

// NewUser sets up a new user service with the given repositories, helpers and
// registers the corresponding routes.

func NewQuote(app *fiber.App, service usecase.UseCases) {
	ctr := quoteController{service}

	apiEndpoint := app.Group("/api/")
	apiEndpoint.Get("quote/:symbol", ctr.getQuote)

}

func (ctr *quoteController) getQuote(ctx *fiber.Ctx) error {

	// Get symbol from params
	symbol := ctx.Params("symbol")

	// Get stock info from iex cloud API
	quote, err := integration.GetStockInfo(symbol)
	if err != nil {
		ctx.Status(fiber.StatusOK)
		return ctx.JSON(fiber.Map{
			"message": "invalid symbol or stock data",
		})
	}
	// Send response
	return ctx.Status(fiber.StatusOK).JSON(quote)
}
