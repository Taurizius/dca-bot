package scheduler

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/Zmey56/dca-bot/internal/binance"
//	"dca-bot/internal/binance"
	"github.com/robfig/cron/v3"
)

func Start(client binance.BinanceClient) {
	symbol := os.Getenv("SYMBOL")
	amountStr := os.Getenv("BUY_AMOUNT")
	cronExpr := os.Getenv("SCHEDULE") // example: "0 10 * * *"

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		log.Fatalf("❌ BUY_AMOUNT parsing error: %v", err)
	}

	c := cron.New()
	_, err = c.AddFunc(cronExpr, func() {


// bis ich alle unnötigen variablen raus habe

	symbol = "BTCUSDC"
	amount = 0

log.Printf("DEBUG  %.2f USDC", amount, symbol)

	//Vorbereitung: Sell All EUR
		symbol0 := "EURUSDC"

		log.Println("🕒 It's time to sell EUR!")
		ctx0, cancel0 := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel0()

		err0 := client.CreateSellAllOrder(ctx0, symbol0)
		if binance.HandleBinanceError(err) {
			err0 = client.CreateSellAllOrder(ctx0, symbol0)
		}
		if err0 != nil {
			log.Printf("❌ Buy error: %v", err0)
			return
		}

		log.Printf("✅ Sold  all EUR")
		
		time.Sleep(10 * time.Second) 

  //ERSTE ORDER BTC
		symbol1 := "BTCUSDC"
		amount1 := 1.00

		log.Println("🕒 It's time to buy BTC!")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := client.CreateMarketOrder(ctx, symbol1, amount1)
		if binance.HandleBinanceError(err) {
			err = client.CreateMarketOrder(ctx, symbol1, amount1)
		}
		if err != nil {
			log.Printf("❌ Buy error: %v", err)
			return
		}

		log.Printf("✅ Buy for  %.2f USDC", amount1)

  //ZWEITE ORDER WBETH
		symbol2 := "ETHUSDC"
		amount2 := 1.00

		log.Println("🕒 It's time to buy WBETH!")
		ctx2, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err2 := client.CreateMarketOrder(ctx2, symbol2, amount2)
		if binance.HandleBinanceError(err2) {
			err2 = client.CreateMarketOrder(ctx2, symbol2, amount2)
		}
		if err != nil {
			log.Printf("❌ Buy error: %v", err2)
			return
		}

		log.Printf("✅ Buy for  %.2f USDC", amount2)


	//Nachbereitung: SWAP all ETH

		time.Sleep(10 * time.Second) 

		log.Println("🕒 It's time to swap ETH!")
		ctx9, cancel9 := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel9()

		err9 := client.CreateSellAllOrder(ctx9, "WBETHETH")
		if binance.HandleBinanceError(err9) {
			err9 = client.CreateSellAllOrder(ctx9, "WBETHETH")
		}
		if err9 != nil {
			log.Printf("❌ Buy error: %v", err9)
			return
		}

		log.Printf("✅ 		Sold  all ETH")


	})
	if err != nil {
		log.Fatalf("Error when adding a cron task: %v", err)
	}

	log.Println("📅 The scheduler is running")
	c.Start()
	select {} //blocking main so that cron works
}
