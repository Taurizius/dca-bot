package binance

import "context"

type BinanceClient interface {
	GetBalance(ctx context.Context, asset string) (float64, error)
	CreateMarketOrder(ctx context.Context, symbol string, quantity float64) error
}
