package client_rest

import (
	"bp-telegram-bot/internal/model"
	"encoding/json"
	"io"
	"log"
	"strconv"
)

type TagRest struct {
	rest *ClientRest
}

func NewTagRest(rest *ClientRest) *TagRest {
	return &TagRest{
		rest: rest,
	}
}

func (c *TagRest) CreateTag(name model.UserTag, userId string) error {
	_, err := c.rest.do("POST", "/tags", name, userId)
	return err
}

func (c *TagRest) UpdateTag(id int64, name string, userId string) error {
	_, err := c.rest.do("PUT", "/tags/"+strconv.FormatInt(id, 10), map[string]string{"name": name}, userId)
	return err
}

func (c *TagRest) DeleteTag(name string, userId string) error {
	_, err := c.rest.do("DELETE", "/tags/"+name, nil, userId)
	return err
}

func (c *TagRest) ListTags(userId string) ([]*model.Tag, error) {
	resp, err := c.rest.do("GET", "/tags", nil, userId)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Failed to close response body: %v", err)
		}
	}(resp.Body)

	var tags []*model.Tag
	err = json.NewDecoder(resp.Body).Decode(&tags)
	return tags, err
}
