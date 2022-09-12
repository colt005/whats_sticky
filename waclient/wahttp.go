package waclient

import (
	"net/http"

	"github.com/colt005/whats_sticky/config"
)

var httpClient *HTTPClient

//Do dispatches the HTTP request to the network
func (c *HTTPClient) Do(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", "Bearer "+config.Config("BEARER_TOKEN"))
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func NewClient() *HTTPClient {
	return &HTTPClient{
		client: &http.Client{},
	}
}
