package api

import (
	"bloodPressureDiary/bp-telegram-bot/model"
	"strconv"
)

func (c *Client) CreatePressure(p model.BloodPressureCreate) error {
	resp, err := c.do("POST", "/pressure", p)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (c *Client) UpdatePressure(id int, p model.BloodPressureCreate) error {
	resp, err := c.do("PUT", "/pressure/"+strconv.Itoa(id), p)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (c *Client) DeletePressure(id int) error {
	resp, err := c.do("DELETE", "/pressure/"+strconv.Itoa(id), nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
