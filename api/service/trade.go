package service

import (
	"fmt"

	"github.com/dakicka/tradingApp/api/auth"
	"github.com/dakicka/tradingApp/api/entity"
	"github.com/dakicka/tradingApp/api/integration"
	"github.com/gofiber/fiber/v2"
)

func (s Service) CreateTrade(ctx *fiber.Ctx, qty int, symbol string) (entity.Trade, error) {
	userId, err := auth.GetUserIdFromToken(ctx)
	if err != nil {
		return entity.Trade{}, fiber.NewError(fiber.ErrInternalServerError.Code, "could not get user id from token")
	}

	// Get current stock price
	stock, err := integration.GetStockInfo(symbol)
	if err != nil {
		fmt.Println(err.Error())
		return entity.Trade{}, err
	}

	// build trade
	t := entity.Trade{
		UserID:      userId,
		Qty:         qty,
		Symbol:      symbol,
		Price:       stock.LatestPrice,
		CompanyName: stock.CompanyName,
	}

	// Build transaction
	tx := entity.Transaction{
		UserID: userId,
		Amount: t.Price * float64(t.Qty) * (-1),
	}
	t, err = s.trades.Create(t)
	if err != nil {
		return entity.Trade{}, fiber.NewError(fiber.ErrInternalServerError.Code, "could not send trade to repo")
	}

	tx.TradeID = &t.ID
	tx, err = s.transactions.CreateTransaction(tx)

	return t, nil
}

func (s Service) GetTradesForDepot(ctx *fiber.Ctx) ([]entity.Trade, error) {
	userId, err := auth.GetUserIdFromToken(ctx)
	if err != nil {
		return []entity.Trade{}, fiber.NewError(fiber.ErrInternalServerError.Code, "could not get user id from token")
	}

	trades, err := s.trades.GetAllByUserId(userId)
	if err != nil {
		return []entity.Trade{}, fiber.NewError(fiber.ErrInternalServerError.Code, "could not get trades from repo")
	}

	return trades, nil

}
