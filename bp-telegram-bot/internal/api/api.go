package api

import (
	"bp-telegram-bot/internal/model"
)

type BloodPressureAPI interface {
	CreatePressure(p model.BloodPressure) error
	ListPressure(userId string) ([]*model.BloodPressure, error)
	UpdatePressure(p model.BloodPressure) error
	DeletePressure(id int64, userId string) error
}

type TagAPI interface {
	CreateTag(name model.UserTag, userId string) error
	ListTags(userId string) ([]*model.Tag, error)
	UpdateTag(id int64, name string, userId string) error
	DeleteTag(name string, userId string) error
}

type ClientAPI struct {
	BloodPressureAPI BloodPressureAPI
	TagAPI           TagAPI
}

func NewAPI(bpa BloodPressureAPI, tga TagAPI) *ClientAPI {
	return &ClientAPI{
		BloodPressureAPI: bpa,
		TagAPI:           tga,
	}
}
