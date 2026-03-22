package client_rest

import (
	"bp-telegram-bot/internal/api"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type ClientRest struct {
	BaseURL string
	Client  *http.Client
}

func NewClientRestBase(baseURL string) *ClientRest {
	return &ClientRest{
		BaseURL: baseURL,
		Client:  http.DefaultClient,
	}
}

func NewClientRest(baseURL string) *api.ClientAPI {
	fulBaseURL := fmt.Sprintf("http://%s", baseURL)
	client := NewClientRestBase(fulBaseURL)
	pressure := NewPressureRest(client)
	tag := NewTagRest(client)
	apiClient := api.NewAPI(pressure, tag)
	return apiClient
}

func (c *ClientRest) do(method, path string, body any, userId string) (*http.Response, error) {
	var buf *bytes.Buffer

	if body != nil {
		b, _ := json.Marshal(body)
		buf = bytes.NewBuffer(b)
	} else {
		buf = bytes.NewBuffer(nil)
	}

	req, err := http.NewRequest(method, c.BaseURL+path, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("user_id", userId)
	req.Header.Set("token_api", os.Getenv("SERVICE_TOKEN_API"))
	req.Header.Set("Content-Type", "application/json")
	return c.Client.Do(req)
}
