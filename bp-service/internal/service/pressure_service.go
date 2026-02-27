package service

import (
	"bloodPressureDiary/bp-service/internal/errors"
	"bloodPressureDiary/bp-service/internal/model"
	"bloodPressureDiary/bp-service/internal/repository"
	"context"
)

type PressureService struct {
	repo *repository.Repository
}

func NewPressureService(repo *repository.Repository) *PressureService {
	return &PressureService{repo: repo}
}

func (s *PressureService) Create(ctx context.Context, p *model.BloodPressure) error {
	for _, tagID := range p.TagIDs {
		ok, err := s.repo.TagRepository.ExistsForUser(ctx, p.UserID, tagID)
		if err != nil {
			return err
		}
		if !ok {
			return errors.ErrTagNotOwnedByUser
		}
	}
	return s.repo.PressureRepository.Create(ctx, p)
}

func (s *PressureService) List(ctx context.Context, userID string) ([]*model.BloodPressure, error) {
	return s.repo.PressureRepository.ListByUser(ctx, userID)
}
