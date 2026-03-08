package service

import (
	"bp-service/internal/errors"
	"bp-service/internal/model"
	"bp-service/internal/repository"
	"context"
	"log"
)

type PressureService struct {
	repo *repository.Repository
}

func NewPressureService(repo *repository.Repository) *PressureService {
	return &PressureService{repo: repo}
}

func (s *PressureService) Create(ctx context.Context, p *model.BloodPressure) error {
	for _, tagName := range p.TagNames {
		ok, err := s.repo.TagRepository.ExistsForUser(ctx, p.UserID, tagName)
		if err != nil {
			log.Println("Error checking tag exists: ", err)
			return err
		}
		if !ok {
			log.Println("Tag exists: ", tagName)
			return errors.ErrTagNotOwnedByUser
		}
	}
	return s.repo.PressureRepository.Create(ctx, p)
}

func (s *PressureService) List(ctx context.Context, userID string) ([]*model.BloodPressure, error) {
	return s.repo.PressureRepository.ListByUser(ctx, userID)
}

func (s *PressureService) Delete(ctx context.Context, id int64, userID string) error {
	return s.repo.PressureRepository.Delete(ctx, id, userID)
}
