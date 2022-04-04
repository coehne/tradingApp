package api

import (
	"github.com/dakicka/tradingApp/api/usecase"
	"github.com/gofiber/fiber/v2"
)

type tradeController struct {
	usecase.UseCases
}

// NewUser sets up a new user service with the given repositories, helpers and
// registers the corresponding routes.

func NewTrade(app *fiber.App, service usecase.UseCases) {
	ctr := tradeController{service}

	apiEndpoint := app.Group("/api/")
	apiEndpoint.Post("trades", ctr.create)
	apiEndpoint.Get("trades/depot", ctr.getDepot)
	apiEndpoint.Get("trades", ctr.getAllByUserId)
	apiEndpoint.Get("trades/:id", ctr.getById)

}

func (ctr *tradeController) create(ctx *fiber.Ctx) error {
	var req createTradeReq

	err := ctx.BodyParser(&req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	// Validate not 0
	if req.Qty == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON("the qty can not be 0")
	}

	_, err = ctr.CreateTrade(ctx, req.Qty, req.Symbol)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	// Send response
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success"})
}

func (ctr *tradeController) getAllByUserId(ctx *fiber.Ctx) error {

	trades, err := ctr.GetTradesByUserId(ctx)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	return ctx.Status(fiber.StatusOK).JSON(trades)

}
func (ctr *tradeController) getDepot(ctx *fiber.Ctx) error {

	trades, err := ctr.GetTradesForDepot(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	return ctx.Status(fiber.StatusOK).JSON(trades)
}
func (ctr *tradeController) getById(ctx *fiber.Ctx) error {

	trade, err := ctr.GetTradeById(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	return ctx.Status(fiber.StatusOK).JSON(trade)
}
