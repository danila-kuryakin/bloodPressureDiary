package repository

import (
	"bloodPressureDiary/bp-service/internal/models"
	"database/sql"
)

type BloodPressureMeasurement interface {
	AddMeasurement(measurement models.Measurement) error
	GetMeasurements(userId string) ([]models.Measurement, error)
	DeleteMeasurement(id int64) error
}

type Repository struct {
	BloodPressureMeasurement
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		BloodPressureMeasurement: NewBloodPressureMeasurementPostgres(db),
	}
}
