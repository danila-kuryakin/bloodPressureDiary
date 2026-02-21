package service

import (
	"context"
)

type TagService struct {
	repo repository.TagRepository
}

func NewTagService(r repository.TagRepository) *TagService {
	return &TagService{repo: r}
}

func (s *TagService) Create(ctx context.Context, tag *model.UserTag) error {
	if tag.Name == "" {
		return errors.ErrInvalidTagName
	}
	return s.repo.Create(ctx, tag)
}

func (s *TagService) List(ctx context.Context, userID string) ([]*model.UserTag, error) {
	return s.repo.ListActiveByUser(ctx, userID)
}

func (s *TagService) Rename(ctx context.Context, tagID int64, userID, name string) error {
	if name == "" {
		return errors.ErrInvalidTagName
	}
	return s.repo.Rename(ctx, tagID, userID, name)
}

func (s *TagService) Delete(ctx context.Context, tagID int64, userID string) error {
	return s.repo.Delete(ctx, tagID, userID)
}
