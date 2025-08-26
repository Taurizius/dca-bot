package binance

import (
    "context"
    "errors"
    "testing"
)

// FakeBinanceClient — простая ручная заглушка для тестов
type FakeBinanceClient struct {
    Balance       float64
    BalanceErr    error
    OrderErr      error
    LastSymbol    string
    LastQuantity  float64
}

func (f *FakeBinanceClient) GetBalance(ctx context.Context, asset string) (float64, error) {
    return f.Balance, f.BalanceErr
}

func (f *FakeBinanceClient) CreateMarketOrder(ctx context.Context, symbol string, quantity float64) error {
    f.LastSymbol = symbol
    f.LastQuantity = quantity
    return f.OrderErr
}

func TestGetBalanceSuccess(t *testing.T) {
    fake := &FakeBinanceClient{Balance: 123.45}

    balance, err := fake.GetBalance(context.Background(), "USDT")
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if balance != 123.45 {
        t.Errorf("expected 123.45, got %v", balance)
    }
}

func TestCreateMarketOrder(t *testing.T) {
    fake := &FakeBinanceClient{}

    err := fake.CreateMarketOrder(context.Background(), "BTCUSDT", 0.001)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if fake.LastSymbol != "BTCUSDT" || fake.LastQuantity != 0.001 {
        t.Errorf("unexpected order data: %s %f", fake.LastSymbol, fake.LastQuantity)
    }
}

func TestGetBalanceWithError(t *testing.T) {
    fake := &FakeBinanceClient{BalanceErr: errors.New("balance error")}

    _, err := fake.GetBalance(context.Background(), "USDT")
    if err == nil {
        t.Fatal("expected error, got nil")
    }
}
