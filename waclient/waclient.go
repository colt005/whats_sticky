package waclient

import (
	"net/http"

	"github.com/colt005/whats_sticky/config"
)

type HTTPClient struct {
	client *http.Client
}

//Do dispatches the HTTP request to the network
func (c *HTTPClient) Do(req *http.Request) (*http.Response, error) {
	resp, err := c.client.Do(req)
	resp.Header.Add("Authorization", "Bearer "+config.Config("BEARER_TOKEN"))
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
