package main

import (
	"bloodPressureDiary/bp-telegram-bot/internal/api"
	"bloodPressureDiary/bp-telegram-bot/internal/bot"
	"bloodPressureDiary/bp-telegram-bot/internal/config"
	"os"
)

func main() {

	config.LoadEnv("bp-telegram-bot/.env")
	token := os.Getenv("TG_TOKEN")
	apiURL := os.Getenv("BP_API_URL")

	apiClient := api.NewClient(apiURL)
	botTelegram := bot.NewBot(token, apiClient)

	botTelegram.Start()
}
