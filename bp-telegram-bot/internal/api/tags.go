package api

import (
	"bloodPressureDiary/bp-telegram-bot/internal/model"
	"encoding/json"
	"io"
	"strconv"
)

func (c *Client) CreateTag(name model.UserTag, userId string) error {
	//fmt.Println("CreateTag ", name)
	_, err := c.do("POST", "/tags", name, userId)
	return err
}

func (c *Client) UpdateTag(id int, name string, userId string) error {
	_, err := c.do("PUT", "/tags/"+strconv.Itoa(id), map[string]string{"name": name}, userId)
	return err
}

func (c *Client) DeleteTag(name string, userId string) error {
	_, err := c.do("DELETE", "/tags/"+name, nil, userId)
	return err
}

func (c *Client) ListTags(userId string) ([]model.Tag, error) {
	resp, err := c.do("GET", "/tags", nil, userId)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	var tags []model.Tag
	err = json.NewDecoder(resp.Body).Decode(&tags)
	return tags, err
}
