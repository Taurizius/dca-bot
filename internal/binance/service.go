package binance

import (
	"context"
	"fmt"
	"github.com/adshao/go-binance/v2"
	"strconv"
	"log"
	"math"
	
	// "github.com/Zmey56/dca-bot/internal/binance"
)

type ClientWrapper struct {
	api *binance.Client
}

func floorToStepSize(qty, step float64) float64 {
    return math.Floor(qty/step) * step
}

func NewClientWrapper(api *binance.Client) *ClientWrapper {
	return &ClientWrapper{api: api}
}

func (c *ClientWrapper) GetBalance(ctx context.Context, asset string) (float64, error) {
	acc, err := c.api.NewGetAccountService().Do(ctx)
	if err != nil {
		return 0, err
	}

	for _, b := range acc.Balances {
		if b.Asset == asset {
			return strconv.ParseFloat(b.Free, 64)
		}
	}

	return 0, fmt.Errorf("asset %s not found", asset)
}


func (c *ClientWrapper) CreateMarketOrder(ctx context.Context, symbol string, quantity float64) error {
    // 1. Preis abfragen
    prices, err := c.api.NewListPricesService().Symbol(symbol).Do(ctx)
    if err != nil {
        return fmt.Errorf("failed to get price: %w", err)
    }
    if len(prices) == 0 {
        return fmt.Errorf("no price found for symbol %s", symbol)
    }

    price, err := strconv.ParseFloat(prices[0].Price, 64)
    if err != nil {
        return fmt.Errorf("invalid price format: %w", err)
    }

    log.Printf("Aktueller Preis für %s: %.8f", symbol, price)

    // 2. Menge berechnen: BTC = USDC / Preis
    quantityBTC := quantity / price

    // Optional: runden auf 6 Stellen, da Binance meist 6 erlaubt, für ETH auf 5
	adjustedQty := 0.00
	if symbol == "BTCUSDC" {
   		 adjustedQty = floorToStepSize(quantityBTC, 0.00001)
	} 
	if symbol == "ETHUSDC" {
   		 adjustedQty = floorToStepSize(quantityBTC, 0.0001)
	}
	if symbol == "EURUSDC" {
   		 adjustedQty = floorToStepSize(quantityBTC, 0.1)
	}
//    adjustedQty := floorToStepSize(quantityBTC, 0.00001)
    quantityStr := fmt.Sprintf("%.6f", adjustedQty)


    log.Printf("Es werden gekauft %s: %s", symbol, quantityStr)

    // 3. Order platzieren
    _, err = c.api.NewCreateOrderService().
        Symbol(symbol).
        Side(binance.SideTypeBuy).
        Type(binance.OrderTypeMarket).
        Quantity(quantityStr).
        Do(ctx)

	return err
}




func (c *ClientWrapper) CreateSellAllOrder(ctx context.Context, symbol string) error {

	// Trying to get a balance
	
    // Initializing the client
	//client := NewClientWrapper(NewBinanceClient())


	balance, err := c.GetBalance(ctx, symbol)
	if HandleBinanceError(err) {
		// can be repeated if necessary.
		balance, err = c.GetBalance(ctx, symbol)
	}
	if err != nil {
		log.Fatalf("Error when receiving the balance: %v", err)
	}

	log.Printf("💰 Balance USDC: %.2f", symbol, balance)


    // 1. Preis abfragen
    prices, err := c.api.NewListPricesService().Symbol(symbol).Do(ctx)
    if err != nil {
        return fmt.Errorf("failed to get price: %w", err)
    }
    if len(prices) == 0 {
        return fmt.Errorf("no price found for symbol %s", symbol)
    }

    price, err := strconv.ParseFloat(prices[0].Price, 64)
    if err != nil {
        return fmt.Errorf("invalid price format: %w", err)
    }

    log.Printf("Aktueller Preis für %s: %.8f", symbol, price)



    // 2. Menge berechnen: 
    quantitySell := balance 

    // Optional: runden auf 6 Stellen, da Binance meist 6 erlaubt, für ETH auf 5
	adjustedQty := 0.00
	if symbol == "BTCUSDC" {
   		 adjustedQty = floorToStepSize(quantitySell, 0.00001)
	} 
	if symbol == "WBETHETH" {
   		 adjustedQty = floorToStepSize(quantitySell, 0.0001)
	}
	if symbol == "EURUSDC" {
   		 adjustedQty = floorToStepSize(quantitySell, 0.1)
	}
//    adjustedQty := floorToStepSize(quantityBTC, 0.00001)
    quantityStr := fmt.Sprintf("%.6f", adjustedQty)


    log.Printf("Es werden gekauft %s: %s", symbol, quantityStr)

    // 3. Order platzieren
    _, err = c.api.NewCreateOrderService().
        Symbol(symbol).
        Side(binance.SideTypeSell).
        Type(binance.OrderTypeMarket).
        Quantity(quantityStr).
        Do(ctx)

	return err
}
