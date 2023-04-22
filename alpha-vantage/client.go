package alphavantage

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const (
	BASE_URL = "https://www.alphavantage.co/query"
)

type Client struct {
	http    *http.Client
	apiKey  string
	baseURL string
}

func NewClient(baseURL, apiKey string) *Client {
	httpClient := http.Client{
		Timeout: 5 * time.Second,
	}

	return &Client{
		http:    &httpClient,
		apiKey:  apiKey,
		baseURL: baseURL,
	}
}

func (c *Client) Do(queryParams url.Values) (*http.Response, error) {
	queryParams.Add("apikey", c.apiKey)

	req, err := http.NewRequest(http.MethodGet, c.baseURL, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %w", err)
	}

	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.URL.RawQuery = queryParams.Encode()
	fmt.Println(req.URL.String())

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code from response: %d", resp.StatusCode)
	}

	return resp, nil

}
