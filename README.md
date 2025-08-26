# DCA Bot

DCA (Dollar-Cost Averaging) Bot is an automated trading bot for the Binance cryptocurrency exchange, implemented in Go.

## Description

DCA Bot automates the process of dollar-cost averaging when trading cryptocurrencies on the Binance exchange. The bot performs periodic purchases of a specified cryptocurrency for a fixed amount, helping to reduce the impact of market volatility on investments.

## Tech Stack

- Go 1.24.2
- Binance API (github.com/adshao/go-binance/v2)
- Cron for task scheduling (github.com/robfig/cron/v3)
- Mock for testing (go.uber.org/mock)
- Configuration via .env files (github.com/joho/godotenv)

## Project Structure

```
.
├── cmd/
│   └── dca-bot/        # Application entry point
├── internal/
│   ├── binance/        # Binance API integration
│   └── scheduler/      # Task scheduler
├── go.mod              # Project dependencies
├── go.sum              # Dependency checksums
└── dca.log            # Application log file
```

## Installation and Running

1. Clone the repository:
```bash
git clone https://github.com/Zmey56/dca-bot.git
cd dca-bot
```

2. Install dependencies:
```bash
go mod download
```

3. Create a .env file in the project root directory and add the required environment variables:
```env
BINANCE_API_KEY=your_api_key
BINANCE_SECRET_KEY=your_secret_key
```

4. Run the bot:
```bash
go run cmd/dca-bot/main.go
```

## Testing

To run tests, use the command:
```bash
go test ./...
```

---

## What’s Next?

This bot is an early version of a larger vision: a full-featured system that analyzes crypto markets and provides DCA signals — potentially through platforms like Telegram.

For now, the focus is on giving developers and traders a clean, extensible base to build and experiment with DCA strategies using the Binance API.

---

## 📚 Full Walkthrough Article

If you’d like a detailed, beginner-friendly explanation of how this bot works, including the rationale behind each part:

👉 [Step-by-Step Guide to Creating a DCA Bot in Go with Binance API](https://medium.com/@alsgladkikh/step-by-step-guide-on-how-to-create-a-dca-bot-on-go-using-the-binance-api-b4eada9c83b9)


## License

MIT

## Author

Zmey56 