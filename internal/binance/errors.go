package binance

import (
	"log"
	"strings"
	"time"
)

func HandleBinanceError(err error) bool {
	if err == nil {
		return false
	}

	errMsg := err.Error()

	if strings.Contains(errMsg, "Too many requests") ||
		strings.Contains(errMsg, "-1003") {
		log.Println("⚠️ The limit of requests to the Binance API has been exceeded. Let's wait and try again...")
		time.Sleep(2 * time.Second)
		return true // Retry the request
	}

	log.Printf("❌ Error of Binance API: %s\n", errMsg)
	return false
}
