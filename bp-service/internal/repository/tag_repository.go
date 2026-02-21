package repository

import (
	"context"
)

type TagRepository interface {
	Create(ctx context.Context, tag *model.UserTag) error
	ListActiveByUser(ctx context.Context, userID string) ([]*model.UserTag, error)
	Rename(ctx context.Context, tagID int64, userID, name string) error
	Delete(ctx context.Context, tagID int64, userID string) error
	ExistsForUser(ctx context.Context, userID string, tagID int64) (bool, error)
	IsUsed(ctx context.Context, tagID int64) (bool, error)
}
