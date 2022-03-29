package api

import (
	"github.com/dakicka/tradingApp/api/usecase"
	"github.com/gofiber/fiber/v2"
)

type transactionController struct {
	usecase.UseCases
}

// NewTransaction sets up a new transactions service with the given repositories, helpers and
// registers the corresponding routes.

func NewTransaction(app *fiber.App, service usecase.UseCases) {
	ctr := transactionController{service}

	apiEndpoint := app.Group("/api/")
	apiEndpoint.Post("transaction", ctr.postTransaction)
	apiEndpoint.Get("transaction", ctr.getTransactions)
}

func (ctr *transactionController) postTransaction(ctx *fiber.Ctx) error {
	var req createTransactionReq

	// TODO: add validation here
	err := ctx.BodyParser(&req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	// Send transaction amount to transaction service
	_, err = ctr.CreateTransaction(ctx, req.Amount)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})
}
func (ctr *transactionController) getTransactions(ctx *fiber.Ctx) error {

	transactions, err := ctr.GetAllForUserId(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return ctx.Status(fiber.StatusOK).JSON(transactions)
}
