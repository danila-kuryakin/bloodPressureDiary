package postgres

import (
	"bp-service/internal/model"
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/lib/pq"
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

	rows, err := tx.QueryContext(ctx, `SELECT id FROM user_tags WHERE name = ANY($1)`, pq.Array(p.TagNames))
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println("Close Rows error:", err)
			err := tx.Rollback()
			if err != nil {
				return
			}
		}
	}(rows)

	var tagIds []int64
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			err := tx.Rollback()
			if err != nil {
				return err
			}
			return err
		}
		tagIds = append(tagIds, id)
	}

	err = tx.QueryRowContext(ctx, `
		INSERT INTO blood_pressure (user_id, systolic, diastolic, pulse)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at`,
		p.UserID, p.Systolic, p.Diastolic, p.Pulse,
	).Scan(&p.ID, &p.CreatedAt)
	if err != nil {
		log.Println("Create QueryRow error:", err)
		err := tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}

	for _, tagId := range tagIds {
		_, err = tx.ExecContext(ctx, `
			INSERT INTO blood_pressure_tag (pressure_id, tag_id)
			VALUES ($1, $2)`,
			p.ID, tagId)
		if err != nil {
			log.Println("Create Exec error:", err)
			err := tx.Rollback()
			if err != nil {
				return err
			}
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
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println("ListByUser Close Rows error:", err)
		}
	}(rows)

	var pressures []*model.BloodPressure
	var IDs []*int64
	for rows.Next() {
		p := &model.BloodPressure{}
		var id = new(int64)
		err := rows.Scan(&id, &p.UserID, &p.Systolic, &p.Diastolic, &p.Pulse, &p.CreatedAt)
		if err != nil {
			log.Println("ListByUser Rows Scan error:", err)
			return nil, err
		}
		p.ID = *id

		pressures = append(pressures, p)
		IDs = append(IDs, id)
	}

	rowsIDs, err := r.db.QueryContext(ctx, `
		select bpt.pressure_id, ut.name
			from blood_pressure_tag as bpt
			join user_tags ut on ut.id = bpt.tag_id
			where bpt.pressure_id = ANY($1)`, pq.Array(IDs))
	if err != nil {
		return nil, err
	}
	defer func(rowsIDs *sql.Rows) {
		err := rowsIDs.Close()
		if err != nil {
			log.Println("ListByUser Close Rows error:", err)
		}
	}(rows)

	tagsMap := make(map[int64][]string)
	for rowsIDs.Next() {
		retTag := &model.RetTag{}
		err := rowsIDs.Scan(&retTag.PressureId, &retTag.Tag)
		if err != nil {
			log.Println("ListByUser Rows Scan error:", err)
			return nil, err
		}
		fmt.Println("retTag", *retTag)

		if tagsMap[retTag.PressureId] == nil {
			tagsMap[retTag.PressureId] = []string{retTag.Tag}
		} else {
			tagsMap[retTag.PressureId] = append(tagsMap[retTag.PressureId], retTag.Tag)
		}
	}

	fmt.Println("tagsMap", tagsMap)

	for _, pressure := range pressures {
		pressure.TagNames = tagsMap[pressure.ID]
	}

	for _, prs := range pressures {
		fmt.Println("pressures:", prs)
	}

	return pressures, nil
}
