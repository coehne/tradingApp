package integration

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Stock struct {
	AvgTotalVolume         int     `json:"avgTotalVolume"`
	CalculationPrice       string  `json:"calculationPrice"`
	Change                 float64 `json:"change"`
	ChangePercent          float64 `json:"changePercent"`
	Close                  float64 `json:"close"`
	CloseSource            string  `json:"closeSource"`
	CloseTime              int64   `json:"closeTime"`
	CompanyName            string  `json:"companyName"`
	Currency               string  `json:"currency"`
	DelayedPrice           float64 `json:"delayedPrice"`
	DelayedPriceTime       int64   `json:"delayedPriceTime"`
	ExtendedChange         float64 `json:"extendedChange"`
	ExtendedChangePercent  float64 `json:"extendedChangePercent"`
	ExtendedPrice          float64 `json:"extendedPrice"`
	ExtendedPriceTime      int64   `json:"extendedPriceTime"`
	High                   float64 `json:"high"`
	HighSource             string  `json:"highSource"`
	HighTime               int64   `json:"highTime"`
	IexAskPrice            float64 `json:"iexAskPrice"`
	IexAskSize             int     `json:"iexAskSize"`
	IexBidPrice            float64 `json:"iexBidPrice"`
	IexBidSize             int     `json:"iexBidSize"`
	IexClose               float64 `json:"iexClose"`
	IexCloseTime           int64   `json:"iexCloseTime"`
	IexLastUpdated         int64   `json:"iexLastUpdated"`
	IexMarketPercent       float64 `json:"iexMarketPercent"`
	IexOpen                float64 `json:"iexOpen"`
	IexOpenTime            int64   `json:"iexOpenTime"`
	IexRealtimePrice       float64 `json:"iexRealtimePrice"`
	IexRealtimeSize        int     `json:"iexRealtimeSize"`
	IexVolume              int     `json:"iexVolume"`
	LastTradeTime          int64   `json:"lastTradeTime"`
	LatestPrice            float64 `json:"latestPrice"`
	LatestSource           string  `json:"latestSource"`
	LatestTime             string  `json:"latestTime"`
	LatestUpdate           int64   `json:"latestUpdate"`
	LatestVolume           int     `json:"latestVolume"`
	Low                    float64 `json:"low"`
	LowSource              string  `json:"lowSource"`
	LowTime                int64   `json:"lowTime"`
	MarketCap              int64   `json:"marketCap"`
	OddLotDelayedPrice     float64 `json:"oddLotDelayedPrice"`
	OddLotDelayedPriceTime int64   `json:"oddLotDelayedPriceTime"`
	Open                   float64 `json:"open"`
	OpenTime               int64   `json:"openTime"`
	OpenSource             string  `json:"openSource"`
	PeRatio                float64 `json:"peRatio"`
	PreviousClose          float64 `json:"previousClose"`
	PreviousVolume         int     `json:"previousVolume"`
	PrimaryExchange        string  `json:"primaryExchange"`
	Symbol                 string  `json:"symbol"`
	Volume                 int     `json:"volume"`
	Week52High             float64 `json:"week52High"`
	Week52Low              float64 `json:"week52Low"`
	YtdChange              float64 `json:"ytdChange"`
	IsUSMarketOpen         bool    `json:"isUSMarketOpen"`
}

// TODO: Refactor into API Client structure
// Credits: https://www.youtube.com/watch?v=evorkFq3Y5k

func GetStockInfo(symbol string) (*Stock, error) {

	apiKey := os.Getenv("API_KEY")
	url := fmt.Sprintf("https://cloud.iexapis.com/stable/stock/%s/quote?token=%s", symbol, apiKey)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var stock Stock

	if err := json.NewDecoder(res.Body).Decode(&stock); err != nil {
		return nil, err
	}

	return &stock, nil
}
