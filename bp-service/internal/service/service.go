package service

import (
	"bp-service/internal/model"
	"bp-service/internal/repository"
	"context"
)

type Pressure interface {
	Create(ctx context.Context, p *model.BloodPressure) error
	List(ctx context.Context, userID string) ([]*model.BloodPressure, error)
	Delete(ctx context.Context, id int64, userID string) error
}

type Tag interface {
	Create(ctx context.Context, tag *model.UserTag) error
	List(ctx context.Context, userID string) ([]*model.UserTag, error)
	Update(ctx context.Context, tagName, userID, name string) error
	Delete(ctx context.Context, tagName, userID string) error
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
