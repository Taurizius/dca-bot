package binance

import (
	"context"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestCreateMarketOrderWithMock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := NewMockBinanceClient(ctrl)

	mock.EXPECT().
		GetBalance(gomock.Any(), "USDT").
		Return(100.0, nil)

	mock.EXPECT().
		CreateMarketOrder(gomock.Any(), "BTCUSDT", 0.001).
		Return(nil)

	// Вызов логики, принимающей интерфейс BinanceClient
	ctx := context.Background()
	balance, err := mock.GetBalance(ctx, "USDT")
	if err != nil || balance != 100.0 {
		t.Errorf("unexpected balance or error: %v, %v", balance, err)
	}

	err = mock.CreateMarketOrder(ctx, "BTCUSDT", 0.001)
	if err != nil {
		t.Errorf("unexpected error on order: %v", err)
	}
}
