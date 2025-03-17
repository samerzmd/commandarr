package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/samerzmd/commandarr/internal/config"
	"github.com/samerzmd/commandarr/internal/handlers"

	"github.com/joho/godotenv"
)

func main() {
	// Load env vars from .env for local development (optional)
	godotenv.Load()

	cfg := config.Load()

	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
		os.Exit(1)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		go handlers.HandleMessage(bot, update.Message, cfg)
	}
}
