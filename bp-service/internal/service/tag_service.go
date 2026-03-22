package service

import (
	"bp-service/internal/customErrors"
	"bp-service/internal/model"
	"bp-service/internal/repository"
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
		return customErrors.ErrInvalidTagName
	}
	return s.repo.TagRepository.Create(ctx, tag)
}

func (s *TagService) List(ctx context.Context, userID string) ([]*model.UserTag, error) {
	return s.repo.TagRepository.ListActiveByUser(ctx, userID)
}

func (s *TagService) Update(ctx context.Context, tagName, userID, name string) error {
	if name == "" {
		return customErrors.ErrInvalidTagName
	}
	return s.repo.TagRepository.Update(ctx, tagName, userID, name)
}

func (s *TagService) Delete(ctx context.Context, tagName, userID string) error {
	return s.repo.TagRepository.Delete(ctx, tagName, userID)
}
