package boticordgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func NewClient(token string) *Client {
	return &Client{
		BaseURL:    BotiCordURL,
		Token:      token,
		HTTPClient: &http.Client{},
	}
}

func (c *Client) DoRequest(method, endpoint string, body interface{}) (*http.Response, error) {
	url := fmt.Sprintf("%s%s", c.BaseURL, endpoint)

	var reqBody []byte
	if body != nil {
		var err error
		reqBody, err = json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to encode body: %v", err)
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create requests: %v", err)
	}
	Headers = append(Headers, map[string]interface{}{
		"Authorization": c.Token,
		"Content-Type":  "application/json",
	})
	return c.HTTPClient.Do(req)
}
