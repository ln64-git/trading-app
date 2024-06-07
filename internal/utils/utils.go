package utils

import (
	"net/http"
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/v2/alpaca"
	"github.com/shopspring/decimal"
)

func InitClient(apiKey, apiSecret, baseURL string) alpaca.Client {
	return alpaca.NewClient(alpaca.ClientOpts{
		ApiKey:    apiKey,
		ApiSecret: apiSecret,
		BaseURL:   baseURL,
		HttpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	})
}

func GetAccount(client alpaca.Client) (*alpaca.Account, error) {
	return client.GetAccount()
}

func PlaceOrder(client alpaca.Client, symbol string, qty *decimal.Decimal, side alpaca.Side, orderType alpaca.OrderType, timeInForce alpaca.TimeInForce) (*alpaca.Order, error) {
	orderReq := alpaca.PlaceOrderRequest{
		AssetKey:    &symbol,
		Qty:         qty,
		Side:        side,
		Type:        orderType,
		TimeInForce: timeInForce,
	}
	return client.PlaceOrder(orderReq)
}

func GetPosition(client alpaca.Client, symbol string) (*alpaca.Position, error) {
	return client.GetPosition(symbol)
}

func ListPositions(client alpaca.Client) ([]alpaca.Position, error) {
	return client.ListPositions()
}
