package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"https:/github.com/danila-kuryakin/bloodPressureDiary/bp-service/internal/repository/postgres"
	"https:/github.com/danila-kuryakin/bloodPressureDiary/bp-service/internal/service"
)

func main() {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = "postgres://user:pass@localhost:5432/bp?sslmode=disable"
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	tagRepo := postgres.NewTagRepo(db)
	pressureRepo := postgres.NewPressureRepo(db)

	tagService := service.NewTagService(tagRepo)
	pressureService := service.NewPressureService(pressureRepo, tagRepo)

	mux := http.NewServeMux()
	httpTransport.NewTagHandler(tagService).Register(mux)
	httpTransport.NewPressureHandler(pressureService).Register(mux)

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
