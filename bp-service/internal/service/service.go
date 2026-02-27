package service

import (
	"bloodPressureDiary/bp-service/internal/model"
	"bloodPressureDiary/bp-service/internal/repository"
	"context"
)

type Pressure interface {
	Create(ctx context.Context, p *model.BloodPressure) error
	List(ctx context.Context, userID string) ([]*model.BloodPressure, error)
}

type Tag interface {
	Create(ctx context.Context, tag *model.UserTag) error
	List(ctx context.Context, userID string) ([]*model.UserTag, error)
	Rename(ctx context.Context, tagID int64, userID, name string) error
	Delete(ctx context.Context, tagID int64, userID string) error
}

type Service struct {
	Pressure
	Tag
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Pressure: NewPressureService(repo),
		Tag:      NewTagService(repo),
	}
}
