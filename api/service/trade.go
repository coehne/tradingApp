package service

import (
	"fmt"
	"strings"

	"github.com/dakicka/tradingApp/api/auth"
	"github.com/dakicka/tradingApp/api/entity"
	"github.com/dakicka/tradingApp/api/integration/iexcloud"

	"github.com/gofiber/fiber/v2"
)

func (s Service) CreateTrade(ctx *fiber.Ctx, qty int, symbol string) (entity.Trade, error) {
	userId, err := auth.GetUserIdFromToken(ctx)
	if err != nil {
		return entity.Trade{}, fiber.NewError(fiber.ErrInternalServerError.Code, "could not get user id from token")
	}

	// Get current stock price
	stock, err := iexcloud.Client.GetStock(nil, symbol)
	if err != nil {
		fmt.Println(err.Error())
		return entity.Trade{}, err
	}

	// build trade
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
	cash, err := s.GetCashForUserId(userId)
	if err != nil {
		return entity.Trade{}, fiber.NewError(fiber.StatusInternalServerError, "could not get balance of user")
	}
	if t.Qty > 0 && cash < tx.Amount {
		return entity.Trade{}, fiber.NewError(fiber.StatusBadRequest, "not enough cash")
	}

	// Check if user has enough stonks to sell
	if tx.Amount > 0 {
		// Check how many stonks of that symbol are in the depot
		depot, err := s.trades.GetDepot(userId)
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
	tx, err = s.transactions.CreateTransaction(tx)

	return t, nil
}

func (s Service) GetTradesByUserId(ctx *fiber.Ctx) ([]entity.Trade, error) {
	userId, err := auth.GetUserIdFromToken(ctx)
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
	userId, err := auth.GetUserIdFromToken(ctx)
	if err != nil {
		return entity.Trade{}, fiber.NewError(fiber.StatusInternalServerError, "could not get user id from token")
	}
	// Get tradeId from params
	tradeId, err := ctx.ParamsInt("id")
	if err != nil {
		return entity.Trade{}, ctx.Status(fiber.StatusBadRequest).JSON("Please ensure that :id is an integer")
	}

	// Get trade from repo
	trade, err := s.trades.GetById(userId, uint(tradeId))
	if err != nil {
		return entity.Trade{}, fiber.NewError(fiber.StatusForbidden, "access denied")
	}

	return trade, nil

}

func (s Service) GetTradesForDepot(ctx *fiber.Ctx) ([]entity.Trade, error) {
	userId, err := auth.GetUserIdFromToken(ctx)
	if err != nil {
		return []entity.Trade{}, fiber.NewError(fiber.StatusInternalServerError, "could not get user id from token")
	}

	trades, err := s.trades.GetDepot(userId)
	if err != nil {
		return []entity.Trade{}, fiber.NewError(fiber.StatusInternalServerError, "could not get trades from repo")
	}
	depoTrades := []entity.Trade{}
	// Get current price and clean data
	for _, trade := range trades {
		if trade.Qty == 0 {
			continue
		} else {
			// Get current stock price
			stock, err := iexcloud.Client.GetStock(nil, trade.Symbol)
			if err != nil {
				return []entity.Trade{}, fiber.NewError(fiber.StatusInternalServerError, "internal server error -> could not get stock data from iexcloud")
			}

			trade.Price = stock.LatestPrice
			depoTrades = append(depoTrades, trade)
		}
	}

	return depoTrades, nil

}
