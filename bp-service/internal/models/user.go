package models

import "time"

type BloodPressureRecord struct {
	UserId      string      `json:"user_id" db:"user_id"`
	Measurement Measurement `json:"measurement" db:"measurement"`
	Time        time.Time   `json:"time" db:"time"`
	Tags        []string    `json:"tags" db:"tags"`

	ID        int64     `json:"id"`
	Systolic  int       `json:"systolic"`
	Diastolic int       `json:"diastolic"`
	Pulse     int       `json:"pulse"`
	CreatedAt time.Time `json:"created_at"`
}

type Measurement struct {
	Systolic  int `json:"systolic" db:"systolic"`
	Diastolic int `json:"diastolic" db:"diastolic"`
	Pulse     int `json:"pulse" db:"pulse"`
}
