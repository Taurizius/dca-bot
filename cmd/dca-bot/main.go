package main

import (
	"context"
	"github.com/Zmey56/dca-bot/internal/scheduler"
	"io"
	"log"
	"os"
	"time"

	"github.com/Zmey56/dca-bot/internal/binance"
	"github.com/joho/godotenv"
)

// func getBalance(symbol) {
// }

func main() {
	logFile, err := os.OpenFile("dca.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("❌  Couldn't open dca.log: %v", err)
	}
	log.SetOutput(io.MultiWriter(os.Stdout, logFile))

	// Load .env
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("❌ Couldn't load .env: %v", err)
	}

	// Initializing the client
	client := binance.NewClientWrapper(binance.NewBinanceClient())

	// The timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Trying to get a balance
	balance, err := client.GetBalance(ctx, "USDC")
	if binance.HandleBinanceError(err) {
		// can be repeated if necessary.
		balance, err = client.GetBalance(ctx, "USDC")
	}
	if err != nil {
		log.Fatalf("Error when receiving the balance: %v", err)
	}

	log.Printf("💰 Balance USDC: %.2f", balance)

	scheduler.Start(client)
}
