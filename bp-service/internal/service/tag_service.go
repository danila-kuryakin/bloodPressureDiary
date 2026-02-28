package service

import (
	"bloodPressureDiary/bp-service/internal/errors"
	"bloodPressureDiary/bp-service/internal/model"
	"bloodPressureDiary/bp-service/internal/repository"
	"context"
)

type TagService struct {
	repo *repository.Repository
}

func NewTagService(r *repository.Repository) *TagService {
	return &TagService{repo: r}
}

func (s *TagService) Create(ctx context.Context, tag *model.UserTag) error {
	if tag.UserID == "" {
		return errors.ErrInvalidTagName
	}
	return s.repo.TagRepository.Create(ctx, tag)
}

func (s *TagService) List(ctx context.Context, userID string) ([]*model.UserTag, error) {
	return s.repo.TagRepository.ListActiveByUser(ctx, userID)
}

func (s *TagService) Rename(ctx context.Context, tagName string, userID, name string) error {
	if name == "" {
		return errors.ErrInvalidTagName
	}
	return s.repo.TagRepository.Rename(ctx, tagName, userID, name)
}

func (s *TagService) Delete(ctx context.Context, tagName string, userID string) error {
	return s.repo.TagRepository.Delete(ctx, tagName, userID)
}
