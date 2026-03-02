package main

import (
	"bp-telegram-bot/internal/api"
	"bp-telegram-bot/internal/bot"
	"bp-telegram-bot/internal/config"
	"os"
)

func main() {

	config.LoadEnv(".env")
	token := os.Getenv("TG_TOKEN")
	apiURL := os.Getenv("BP_API_URL")

	apiClient := api.NewClient(apiURL)
	botTelegram := bot.NewBot(token, apiClient)

	botTelegram.Start()
}
