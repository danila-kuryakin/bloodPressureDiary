package repository

import (
	"bp-service/internal/model"
	"bp-service/internal/repository/postgres"
	"context"
	"database/sql"
)

type PressureRepository interface {
	Create(ctx context.Context, p *model.BloodPressure) error
	ListByUser(ctx context.Context, userID string) ([]*model.BloodPressure, error)
}

type TagRepository interface {
	Create(ctx context.Context, tag *model.UserTag) error
	CreateAll(ctx context.Context, tag *model.UserTag) error
	ListActiveByUser(ctx context.Context, userID string) ([]*model.UserTag, error)
	Rename(ctx context.Context, tagName string, userID, name string) error
	Delete(ctx context.Context, tagName string, userID string) error
	ExistsForUser(ctx context.Context, userID string, tagName string) (bool, error)
	IsUsed(ctx context.Context, tagName string) (bool, error)
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
