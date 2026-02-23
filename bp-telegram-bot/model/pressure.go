package model

type BloodPressureCreate struct {
	Systolic  int      `json:"systolic"`
	Diastolic int      `json:"diastolic"`
	Pulse     int      `json:"pulse"`
	Tags      []string `json:"tags"`
}
