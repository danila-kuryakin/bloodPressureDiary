package api

import (
	"bloodPressureDiary/bp-telegram-bot/internal/model"
	"io"
	"strconv"
)

func (c *Client) CreatePressure(p model.BloodPressure, userId string) error {
	resp, err := c.do("POST", "/pressure", p, userId)
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

func (c *Client) UpdatePressure(id int, p model.BloodPressure, userId string) error {
	resp, err := c.do("PUT", "/pressure/"+strconv.Itoa(id), p, userId)
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

func (c *Client) DeletePressure(id int, userId string) error {
	resp, err := c.do("DELETE", "/pressure/"+strconv.Itoa(id), nil, userId)
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
