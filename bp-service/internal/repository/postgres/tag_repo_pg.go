package postgres

import (
	"context"
	"database/sql"
)

type TagRepo struct {
	db *sql.DB
}

func NewTagRepo(db *sql.DB) *TagRepo {
	return &TagRepo{db: db}
}

func (r *TagRepo) Create(ctx context.Context, tag *model.UserTag) error {
	return r.db.QueryRowContext(ctx, `
		INSERT INTO user_tags (user_id, name)
		VALUES ($1, $2)
		RETURNING id, is_active, created_at`,
		tag.UserID, tag.Name,
	).Scan(&tag.ID, &tag.IsActive, &tag.CreatedAt)
}

func (r *TagRepo) ListActiveByUser(ctx context.Context, userID string) ([]*model.UserTag, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, user_id, name, is_active, created_at
		FROM user_tags
		WHERE user_id=$1 AND is_active=true
		ORDER BY name`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []*model.UserTag
	for rows.Next() {
		t := &model.UserTag{}
		rows.Scan(&t.ID, &t.UserID, &t.Name, &t.IsActive, &t.CreatedAt)
		res = append(res, t)
	}
	return res, nil
}

func (r *TagRepo) ExistsForUser(ctx context.Context, userID string, tagID int64) (bool, error) {
	var ok bool
	err := r.db.QueryRowContext(ctx, `
		SELECT EXISTS (
			SELECT 1 FROM user_tags
			WHERE id=$1 AND user_id=$2 AND is_active=true
		)`, tagID, userID).Scan(&ok)
	return ok, err
}

func (r *TagRepo) IsUsed(ctx context.Context, tagID int64) (bool, error) {
	var used bool
	err := r.db.QueryRowContext(ctx, `
		SELECT EXISTS (
			SELECT 1 FROM blood_pressure_tag WHERE tag_id=$1
		)`, tagID).Scan(&used)
	return used, err
}

func (r *TagRepo) Rename(ctx context.Context, tagID int64, userID, name string) error {
	_, err := r.db.ExecContext(ctx, `
		UPDATE user_tags
		SET name=$1
		WHERE id=$2 AND user_id=$3 AND is_active=true`,
		name, tagID, userID)
	return err
}

func (r *TagRepo) Delete(ctx context.Context, tagID int64, userID string) error {
	used, err := r.IsUsed(ctx, tagID)
	if err != nil {
		return err
	}

	if used {
		_, err = r.db.ExecContext(ctx, `
			UPDATE user_tags
			SET is_active=false
			WHERE id=$1 AND user_id=$2`,
			tagID, userID)
		return err
	}

	_, err = r.db.ExecContext(ctx,
		`DELETE FROM user_tags WHERE id=$1 AND user_id=$2`,
		tagID, userID)
	return err
}
