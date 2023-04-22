package alphavantage

import (
	"net/http"
	"time"
)

const (
	BASE_URL = "https://www.alphavantage.co/query"
)

type Client struct {
	http   *http.Client
	apiKey string
}

func NewClient(apiKey string) *Client {
	httpClient := http.Client{
		Timeout: 5 * time.Second,
	}

	return &Client{
		http:   &httpClient,
		apiKey: apiKey,
	}
}

func (c *Client) Do() {

}
