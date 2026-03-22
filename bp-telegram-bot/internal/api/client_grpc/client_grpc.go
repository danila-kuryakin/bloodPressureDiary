package client_grpc

import (
	"bp-telegram-bot/internal/api"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewClientGRPC(baseURL string) (*api.ClientAPI, *grpc.ClientConn) {
	conn, err := grpc.NewClient(
		baseURL,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	//defer conn.Close()

	pressureGRPC := NewPressureGRPC(conn)
	tagGRPC := NewTagGRPC(conn)
	apiClient := api.NewAPI(pressureGRPC, tagGRPC)
	return apiClient, conn
}
