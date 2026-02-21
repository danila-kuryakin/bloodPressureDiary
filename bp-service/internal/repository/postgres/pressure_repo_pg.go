package postgres

import (
	"context"
	"database/sql"
)

type PressureRepo struct {
	db *sql.DB
}

func NewPressureRepo(db *sql.DB) *PressureRepo {
	return &PressureRepo{db: db}
}

func (r *PressureRepo) Create(ctx context.Context, p *model.BloodPressure) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = tx.QueryRowContext(ctx, `
		INSERT INTO blood_pressure (user_id, systolic, diastolic, pulse)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at`,
		p.UserID, p.Systolic, p.Diastolic, p.Pulse,
	).Scan(&p.ID, &p.CreatedAt)
	if err != nil {
		return err
	}

	for _, tagID := range p.TagIDs {
		_, err = tx.ExecContext(ctx, `
			INSERT INTO blood_pressure_tag (pressure_id, tag_id)
			VALUES ($1, $2)`,
			p.ID, tagID)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r *PressureRepo) ListByUser(ctx context.Context, userID string) ([]*model.BloodPressure, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, user_id, systolic, diastolic, pulse, created_at
		FROM blood_pressure
		WHERE user_id=$1
		ORDER BY created_at DESC`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []*model.BloodPressure
	for rows.Next() {
		p := &model.BloodPressure{}
		rows.Scan(&p.ID, &p.UserID, &p.Systolic, &p.Diastolic, &p.Pulse, &p.CreatedAt)
		res = append(res, p)
	}
	return res, nil
}
