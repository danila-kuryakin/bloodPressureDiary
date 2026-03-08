package api

import (
	"bp-telegram-bot/internal/model"
	"encoding/json"
	"io"
	"log"
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

func (c *Client) ListPressure(userId string) ([]model.BloodPressure, error) {
	resp, err := c.do("GET", "/pressure", nil, userId)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	var press []model.BloodPressure
	err = json.NewDecoder(resp.Body).Decode(&press)
	return press, err
}

func (c *Client) UpdatePressure(id int, p model.BloodPressure, userId string) error {
	resp, err := c.do("PUT", "/pressure/"+strconv.Itoa(id), p, userId)
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

func (c *Client) DeletePressure(id int, userId string) error {
	resp, err := c.do("DELETE", "/pressure/"+strconv.Itoa(id), nil, userId)
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
