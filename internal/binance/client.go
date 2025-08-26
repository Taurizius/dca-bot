package binance

import (
	"github.com/adshao/go-binance/v2"
	"os"
)

func NewBinanceClient() *binance.Client {
	return binance.NewClient(
		os.Getenv("BINANCE_API_KEY"),
		os.Getenv("BINANCE_SECRET_KEY"),
	)
}
