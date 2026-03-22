package postgres

import (
	"bp-service/internal/model"
	"context"
	"database/sql"
	"fmt"
	"log"
)

type TagRepo struct {
	db *sql.DB
}

func NewTagRepo(db *sql.DB) *TagRepo {
	return &TagRepo{db: db}
}

func (r *TagRepo) Create(ctx context.Context, tag *model.UserTag) error {
	//fmt.Println("Repo Create", tag)

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer func(tx *sql.Tx) {
		err := tx.Rollback()
		if err != nil {

		}
	}(tx)

	query := `INSERT INTO user_tags (user_id, name)
		VALUES ($1, $2)`

	_, err = tx.ExecContext(ctx, query, tag.UserID, tag.Name)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (r *TagRepo) CreateAll(ctx context.Context, tag *model.UserTag) error {
	//fmt.Println("Repo Create", tag)

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer func(tx *sql.Tx) {
		err := tx.Rollback()
		if err != nil {

		}
	}(tx)

	query := `INSERT INTO user_tags (user_id, name)
		VALUES ($1, $2)`

	for _, tagName := range tag.Name {
		_, err := tx.ExecContext(ctx, query, tag.UserID, tagName)
		if err != nil {
			return err
		}
	}
	return tx.Commit()
}

func (r *TagRepo) ListActiveByUser(ctx context.Context, userID string) ([]*model.UserTag, error) {
	fmt.Println("Repo ListActiveByUser", userID)
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, user_id, name, is_active, created_at
		FROM user_tags
		WHERE user_id=$1 AND is_active=true
		ORDER BY name`, userID)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println("Close ListActiveByUser:", err)
		}
	}(rows)

	var res []*model.UserTag
	for rows.Next() {
		t := &model.UserTag{}
		err := rows.Scan(&t.ID, &t.UserID, &t.Name, &t.IsActive, &t.CreatedAt)
		if err != nil {
			log.Println("Scan ListActiveByUser:", err)
			return nil, err
		}
		res = append(res, t)
	}
	return res, nil
}

func (r *TagRepo) ExistsForUser(ctx context.Context, userID string, tagName string) (bool, error) {
	var ok bool
	err := r.db.QueryRowContext(ctx, `
		SELECT EXISTS (
			SELECT 1 FROM user_tags
			WHERE name=$1 AND user_id=$2 AND is_active=true
		)`, tagName, userID).Scan(&ok)
	return ok, err
}

func (r *TagRepo) IsUsed(ctx context.Context, tagName string) (bool, error) {
	var used bool
	err := r.db.QueryRowContext(ctx, `
		SELECT EXISTS (
			SELECT 1 FROM blood_pressure_tag WHERE tag_id=$1
		)`, tagName).Scan(&used)

	if err != nil {
		log.Println("IsUsed:", err)
	}

	return used, err
}

func (r *TagRepo) Update(ctx context.Context, tagName, userID, name string) error {
	_, err := r.db.ExecContext(ctx, `
		UPDATE user_tags
		SET name=$1
		WHERE name=$2 AND user_id=$3 AND is_active=true`,
		name, tagName, userID)
	return err
}

func (r *TagRepo) Delete(ctx context.Context, tagName string, userID string) error {

	log.Println("Repo Delete", tagName, userID)

	//used, err := r.IsUsed(ctx, tagName)
	//if err != nil {
	//	return err
	//}
	//
	//log.Println("1")
	//
	//if used {
	//	_, err = r.db.ExecContext(ctx, `
	//		UPDATE user_tags
	//		SET is_active=false
	//		WHERE name=$1 AND user_id=$2`,
	//		tagName, userID)
	//	return err
	//}
	//
	//log.Println("1")

	_, err := r.db.ExecContext(ctx,
		`DELETE FROM user_tags WHERE name=$1 AND user_id=$2`,
		tagName, userID)

	return err
}
