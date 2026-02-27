package api

import (
	"bloodPressureDiary/bp-telegram-bot/model"
	"strconv"
)

func (c *Client) CreatePressure(p model.BloodPressureCreate, userId string) error {
	resp, err := c.do("POST", "/pressure", p, userId)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (c *Client) UpdatePressure(id int, p model.BloodPressureCreate, userId string) error {
	resp, err := c.do("PUT", "/pressure/"+strconv.Itoa(id), p, userId)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (c *Client) DeletePressure(id int, userId string) error {
	resp, err := c.do("DELETE", "/pressure/"+strconv.Itoa(id), nil, userId)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
