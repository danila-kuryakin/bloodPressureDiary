package api

import (
	model_tg "bloodPressureDiary/bp-telegram-bot/internal/model"
	"bloodPressureDiary/bp-telegram-bot/model"
	"encoding/json"
	"strconv"
)

func (c *Client) CreateTag(name model_tg.UserTag, userId string) error {
	//fmt.Println("CreateTag ", name)
	_, err := c.do("POST", "/tags", name, userId)
	return err
}

func (c *Client) UpdateTag(id int, name string, userId string) error {
	_, err := c.do("PUT", "/tags/"+strconv.Itoa(id), map[string]string{"name": name}, userId)
	return err
}

func (c *Client) DeleteTag(id int, userId string) error {
	_, err := c.do("DELETE", "/tags/"+strconv.Itoa(id), nil, userId)
	return err
}

func (c *Client) ListTags(userId string) ([]model.Tag, error) {
	resp, err := c.do("GET", "/tags", nil, userId)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var tags []model.Tag
	err = json.NewDecoder(resp.Body).Decode(&tags)
	return tags, err
}
