package api

import (
	"bloodPressureDiary/bp-telegram-bot/model"
	"encoding/json"
	"strconv"
)

func (c *Client) CreateTag(name string) error {
	_, err := c.do("POST", "/tags", map[string]string{"name": name})
	return err
}

func (c *Client) UpdateTag(id int, name string) error {
	_, err := c.do("PUT", "/tags/"+strconv.Itoa(id), map[string]string{"name": name})
	return err
}

func (c *Client) DeleteTag(id int) error {
	_, err := c.do("DELETE", "/tags/"+strconv.Itoa(id), nil)
	return err
}

func (c *Client) ListTags() ([]model.Tag, error) {
	resp, err := c.do("GET", "/tags", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var tags []model.Tag
	err = json.NewDecoder(resp.Body).Decode(&tags)
	return tags, err
}
