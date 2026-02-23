package repository

import (
	"bloodPressureDiary/bp-service/internal/repository/postgres"
	"database/sql"
)

type Repository struct {
	PressureRepository
	TagRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		PressureRepository: postgres.NewPressureRepo(db),
		TagRepository:      postgres.NewTagRepo(db),
	}
}
