package model

import "time"

type BloodPressure struct {
	ID        int64     `json:"id"`
	UserID    string    `json:"user_id"`
	Systolic  int       `json:"systolic"`
	Diastolic int       `json:"diastolic"`
	Pulse     int       `json:"pulse"`
	TagIDs    []string  `json:"tag_ids"`
	CreatedAt time.Time `json:"created_at"`
}
