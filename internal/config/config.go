package config

import (
	"os"
	"time"
)

type Config struct {
	APIKey       string
	APISecret    string
	BaseURL      string
	TradeAmount  float64
	PollInterval time.Duration
}

func LoadConfig() *Config {
	return &Config{
		APIKey:       os.Getenv("APCA_API_KEY_ID"),
		APISecret:    os.Getenv("APCA_API_SECRET_KEY"),
		BaseURL:      "https://paper-api.alpaca.markets",
		TradeAmount:  100.0,
		PollInterval: 5 * time.Minute,
	}
}
