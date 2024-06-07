package main

import (
	"time"

	"trading-app/internal/config"
	"trading-app/internal/strategy"
	"trading-app/internal/utils"
)

func main() {
	cfg := config.LoadConfig()
	client := utils.InitClient(cfg.APIKey, cfg.APISecret, cfg.BaseURL)

	for {
		strategy.Trade(client, cfg)
		time.Sleep(cfg.PollInterval)
	}
}
