package main

import (
	"bp-telegram-bot/internal/api/client_grpc"
	"bp-telegram-bot/internal/api/client_rest"
	"bp-telegram-bot/internal/bot"
	"bp-telegram-bot/internal/config"
	"log"
	"os"

	"google.golang.org/grpc"
)

func main() {

	config.LoadEnv(".env")
	token := os.Getenv("TG_TOKEN")
	apiURL := os.Getenv("BP_API_URL")
	protocolType := os.Getenv("BP_PROTOCOL_TYPE")

	log.Println(apiURL, protocolType)
	apiClient, conn := client_grpc.NewClientGRPC(apiURL)
	switch protocolType {
	case "rest":
		apiClient = client_rest.NewClientRest(apiURL)
	case "grpc":
		apiClient, conn = client_grpc.NewClientGRPC(apiURL)
		defer func(conn *grpc.ClientConn) {
			err := conn.Close()
			if err != nil {

			}
		}(conn)
	}
	botTelegram := bot.NewBot(token, apiClient)

	log.Println("Starting bot")
	botTelegram.Start()
}
