package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

type Client struct {
	BaseURL string
	Client  *http.Client
}

func NewClient(baseURL string) *Client {
	return &Client{
		BaseURL: baseURL,
		Client:  &http.Client{},
	}
}

func (c *Client) do(method, path string, body any, userId string) (*http.Response, error) {
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

	//fmt.Println("====================")
	//fmt.Println(method)
	//fmt.Println(c.BaseURL + path)
	//fmt.Println(body)
	//fmt.Println(userId)
	//fmt.Println("====================\n")

	req.Header.Set("Content-Type", "application/json")
	return c.Client.Do(req)
}
