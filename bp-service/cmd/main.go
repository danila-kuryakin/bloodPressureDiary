package main

import (
	"bloodPressureDiary/bp-service/internal/config"
	"bloodPressureDiary/bp-service/internal/repository"
	"bloodPressureDiary/bp-service/internal/service"
	"bloodPressureDiary/bp-service/internal/transport/rest"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	config.LoadEnv("bp-service/.env")
	cfg := config.LoadConfig("bp-service/configs/config.yml")

	// Конфигурация и подключение к PostgreSQL
	postgresConf := repository.PostgresConfig{
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     cfg.Database.Host,
		Port:     strconv.Itoa(cfg.Database.Port),
		Name:     cfg.Database.Name,
		SSLMode:  cfg.Database.SSLMode,
	}
	db, err := repository.NewPostgresDB(postgresConf)
	if err != nil {
		log.Println(err)
	}

	repo := repository.NewRepository(db)
	svc := service.NewService(repo)
	hand := rest.NewHandler(svc)

	mux := http.NewServeMux()
	hand.Register(mux)

	// Создаем http сервер
	log.Println(fmt.Sprintf("Server started on: %s", cfg.Server.Port))
	if err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.Server.Port), mux); err != nil {
		return
	}
}
