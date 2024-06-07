package strategy

import (
	"log"

	"trading-app/internal/config"
	"trading-app/internal/utils"

	"github.com/alpacahq/alpaca-trade-api-go/v2/alpaca"
	"github.com/shopspring/decimal"
)

func Trade(client alpaca.Client, config *config.Config) {
	account, err := utils.GetAccount(client)
	if err != nil {
		log.Fatalf("failed to get account: %v", err)
	}

	log.Printf("Current cash: %s", account.Cash)

	// Example strategy: Buy AAPL if cash > TradeAmount
	symbol := "AAPL"
	tradeAmount := config.TradeAmount

	if account.Cash.GreaterThanOrEqual(decimal.NewFromFloat(tradeAmount)) {
		qty := decimal.NewFromFloat(tradeAmount / 100.0) // assuming price is $100 for simplicity
		_, err := utils.PlaceOrder(client, symbol, &qty, alpaca.Buy, alpaca.Market, alpaca.GTC)
		if err != nil {
			log.Fatalf("failed to place order: %v", err)
		}
		log.Printf("Placed order to buy %f of %s", qty.InexactFloat64(), symbol)
	} else {
		log.Printf("Not enough cash to place order")
	}

	// Monitor and adjust positions
	positions, err := utils.ListPositions(client)
	if err != nil {
		log.Fatalf("failed to list positions: %v", err)
	}

	for _, position := range positions {
		// Simple strategy: sell if position has made a profit
		if position.UnrealizedPL.GreaterThan(decimal.NewFromFloat(1.0)) { // assuming $1 profit for simplicity
			_, err := utils.PlaceOrder(client, position.Symbol, &position.Qty, alpaca.Sell, alpaca.Market, alpaca.GTC)
			if err != nil {
				log.Fatalf("failed to place sell order: %v", err)
			}
			log.Printf("Placed order to sell %f of %s", position.Qty.InexactFloat64(), position.Symbol)
		}
	}
}
