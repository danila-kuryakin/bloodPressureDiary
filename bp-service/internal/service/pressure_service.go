package service

import (
	"bloodPressureDiary/bp-service/internal/errors"
	"bloodPressureDiary/bp-service/internal/model"
	"bloodPressureDiary/bp-service/internal/repository"
	"context"
)

type PressureService struct {
	repo    repository.PressureRepository
	tagRepo repository.TagRepository
}

func NewPressureService(
	repo repository.PressureRepository,
	tagRepo repository.TagRepository,
) *PressureService {
	return &PressureService{repo: repo, tagRepo: tagRepo}
}

func (s *PressureService) Create(ctx context.Context, p *model.BloodPressure) error {
	for _, tagID := range p.TagIDs {
		ok, err := s.tagRepo.ExistsForUser(ctx, p.UserID, tagID)
		if err != nil {
			return err
		}
		if !ok {
			return errors.ErrTagNotOwnedByUser
		}
	}
	return s.repo.Create(ctx, p)
}

func (s *PressureService) List(ctx context.Context, userID string) ([]*model.BloodPressure, error) {
	return s.repo.ListByUser(ctx, userID)
}
