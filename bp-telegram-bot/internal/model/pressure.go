package model

import "time"

type BloodPressure struct {
	ID        int64     `json:"id"`
	UserID    string    `json:"user_id"`
	Systolic  int32     `json:"systolic"`
	Diastolic int32     `json:"diastolic"`
	Pulse     int32     `json:"pulse"`
	TagNames  []string  `json:"tag_names"`
	CreatedAt time.Time `json:"created_at"`
}

type DeletePressure struct {
	ID int64 `json:"id"`
}
