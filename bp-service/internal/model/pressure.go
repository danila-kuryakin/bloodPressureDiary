package model

import "time"

type BloodPressure struct {
	ID        int64     `json:"id"`
	UserID    string    `json:"user_id"`
	Systolic  int       `json:"systolic"`
	Diastolic int       `json:"diastolic"`
	Pulse     int       `json:"pulse"`
	TagNames  []string  `json:"tag_names"`
	CreatedAt time.Time `json:"created_at"`
}

type RetTag struct {
	PressureId int64  `json:"pressure_id"`
	Tag        string `json:"name"`
}
