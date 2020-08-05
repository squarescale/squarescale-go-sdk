package api

import "net/http"

const defaultSquarescaleAddress = "https://www.squarescale.io/"

type Client struct {
	apiKey  string
	address string
}

func NewClient(apiKey string) *Client {
	return &Client{apiKey: apiKey, address: defaultSquarescaleAddress}
}

// SetCustomAddress is used to request non default
// Squarescale plateform.
func (c *Client) SetCustomAddress(address string) {
	c.address = address
}

func (c *Client) httpClient() *http.Client {
	return &http.Client{}
}
