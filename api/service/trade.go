package service

import (
	"fmt"
	"strings"

	"github.com/dakicka/tradingApp/api/entity"
	"github.com/dakicka/tradingApp/api/integration"
	"github.com/dakicka/tradingApp/api/package/auth"
	"github.com/gofiber/fiber/v2"
)

func (s Service) CreateTrade(ctx *fiber.Ctx, qty int, symbol string) (entity.Trade, error) {
	userId, err := auth.GetUserIdByContext(ctx)
	if err != nil {
		return entity.Trade{}, fiber.NewError(fiber.ErrInternalServerError.Code, "could not get user id from token")
	}

	// Get current stock price
	stock, err := integration.GetStockInfo(symbol)
	if err != nil {
		fmt.Println(err.Error())
		return entity.Trade{}, err
	}

	// Build trade
	t := entity.Trade{
		UserID:      userId,
		Qty:         qty,
		Symbol:      strings.ToUpper(symbol),
		Price:       stock.LatestPrice,
		CompanyName: stock.CompanyName,
	}

	// Build transaction
	tx := entity.Transaction{
		UserID: userId,
		Amount: t.Price * float64(t.Qty) * (-1),
	}

	// Check if user has enough cash
	cash, err := s.GetCashByUserId(userId)
	if err != nil {
		return entity.Trade{}, fiber.NewError(fiber.StatusInternalServerError, "could not get balance of user")
	}

	if t.Qty > 0 && cash < t.Price*float64(t.Qty) {
		return entity.Trade{}, fiber.NewError(fiber.StatusBadRequest, "not enough cash")
	}

	// Check if user has enough stonks to sell
	if tx.Amount > 0 {

		// Check how many stonks of that symbol are in the depot
		depot, err := s.trades.GetDepotByUserId(userId)
		if err != nil {
			return entity.Trade{}, fiber.NewError(fiber.StatusInternalServerError, "could not get depot from repo")
		}

		var stocks int
		for _, stock := range depot {
			if t.Symbol == stock.Symbol {
				stocks += stock.Qty
			}
		}

		if stocks+t.Qty < 0 {
			return entity.Trade{}, fiber.NewError(fiber.StatusBadRequest, "not enough stonks")
		}

	}

	t, err = s.trades.Create(t)
	if err != nil {
		return entity.Trade{}, fiber.NewError(fiber.StatusInternalServerError, "could not send trade to repo")
	}

	tx.TradeID = &t.ID
	tx, err = s.transactions.Create(tx)

	return t, nil
}

func (s Service) GetTradesByUserId(ctx *fiber.Ctx) ([]entity.Trade, error) {
	userId, err := auth.GetUserIdByContext(ctx)
	if err != nil {
		return []entity.Trade{}, fiber.NewError(fiber.StatusInternalServerError, "could not get user id from token")
	}

	trades, err := s.trades.GetAllByUserId(userId)
	if err != nil {
		return []entity.Trade{}, fiber.NewError(fiber.StatusInternalServerError, "could not get trades from repo")
	}

	return trades, nil

}

func (s Service) GetTradeById(ctx *fiber.Ctx) (entity.Trade, error) {
	userId, err := auth.GetUserIdByContext(ctx)
	if err != nil {
		return entity.Trade{}, fiber.NewError(fiber.StatusInternalServerError, "could not get user id from token")
	}

	// Get tradeId from params
	tradeId, err := ctx.ParamsInt("id")
	if err != nil {
		return entity.Trade{}, ctx.Status(fiber.StatusBadRequest).JSON("Please ensure that :id is an integer")
	}

	trade, err := s.trades.GetById(userId, uint(tradeId))
	if err != nil {
		return entity.Trade{}, fiber.NewError(fiber.StatusForbidden, "access denied")
	}

	return trade, nil

}

func (s Service) GetTradesForDepot(ctx *fiber.Ctx) ([]entity.Trade, error) {
	userId, err := auth.GetUserIdByContext(ctx)
	if err != nil {
		return []entity.Trade{}, fiber.NewError(fiber.StatusInternalServerError, "could not get user id from token")
	}

	trades, err := s.trades.GetDepotByUserId(userId)
	if err != nil {
		return []entity.Trade{}, fiber.NewError(fiber.StatusInternalServerError, "could not get trades from repo")
	}

	var depoTrades []entity.Trade

	// Get current price and clean data
	for _, trade := range trades {
		if trade.Qty == 0 {
			continue
		} else {
			// Get current stock price
			stock, err := integration.GetStockInfo(trade.Symbol)
			if err != nil {
				return []entity.Trade{}, fiber.NewError(fiber.StatusInternalServerError, "internal server error -> could not get stock data from iexcloud, try again later")
			}

			trade.Price = stock.LatestPrice
			depoTrades = append(depoTrades, trade)
		}
	}

	return depoTrades, nil

}
