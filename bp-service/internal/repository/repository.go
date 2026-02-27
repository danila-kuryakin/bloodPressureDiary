package repository

import (
	"bloodPressureDiary/bp-service/internal/model"
	"bloodPressureDiary/bp-service/internal/repository/postgres"
	"context"
	"database/sql"
)

type PressureRepository interface {
	Create(ctx context.Context, p *model.BloodPressure) error
	ListByUser(ctx context.Context, userID string) ([]*model.BloodPressure, error)
}

type TagRepository interface {
	Create(ctx context.Context, tag *model.UserTag) error
	ListActiveByUser(ctx context.Context, userID string) ([]*model.UserTag, error)
	Rename(ctx context.Context, tagID int64, userID, name string) error
	Delete(ctx context.Context, tagID int64, userID string) error
	ExistsForUser(ctx context.Context, userID string, tagID int64) (bool, error)
	IsUsed(ctx context.Context, tagID int64) (bool, error)
}

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
