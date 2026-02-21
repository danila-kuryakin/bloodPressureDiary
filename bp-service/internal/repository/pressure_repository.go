package repository

import ()

type PressureRepository interface {
	Create(ctx context.Context, p *model.BloodPressure) error
	ListByUser(ctx context.Context, userID string) ([]*model.BloodPressure, error)
}
