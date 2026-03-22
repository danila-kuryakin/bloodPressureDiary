package client_rest

import (
	"bp-telegram-bot/internal/model"
	"encoding/json"
	"io"
	"log"
	"strconv"
)

type PressureRest struct {
	rest *ClientRest
}

func NewPressureRest(rest *ClientRest) *PressureRest {
	return &PressureRest{
		rest: rest,
	}
}

func (c *PressureRest) CreatePressure(p model.BloodPressure) error {
	resp, err := c.rest.do("POST", "/pressure", p, p.UserID)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)
	return nil
}

func (c *PressureRest) ListPressure(userId string) ([]*model.BloodPressure, error) {
	resp, err := c.rest.do("GET", "/pressure", nil, userId)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	var press []*model.BloodPressure
	err = json.NewDecoder(resp.Body).Decode(&press)
	return press, err
}

func (c *PressureRest) UpdatePressure(p model.BloodPressure) error {
	resp, err := c.rest.do("PUT", "/pressure/"+strconv.FormatInt(p.ID, 10), p, p.UserID)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Error closing body:", err)
		}
	}(resp.Body)
	return nil
}

func (c *PressureRest) DeletePressure(id int64, userId string) error {
	resp, err := c.rest.do("DELETE", "/pressure/"+strconv.FormatInt(id, 10), nil, userId)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Error closing body:", err)
		}
	}(resp.Body)
	return nil
}
