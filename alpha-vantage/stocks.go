package alphavantage

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"
)

var (
	outputSizeOpts = map[string]struct{}{
		"compact": {},
		"full":    {},
	}
)

var ErrTooManyRequests = errors.New("too many API requests")

type TimeSeriesResponse struct {
	Metadata   Metadata              `json:"Meta Data"`
	TimeSeries map[string]TimeSeries `json:"Time Series (Daily)"`
	Note       string                `json:"Note,omitempty"`
}

type Metadata struct {
	Information   string `json:"1. Information,omitempty"`
	Symbol        string `json:"2. Symbol,omitempty"`
	LastRefreshed string `json:"3. Last Refreshed,omitempty"`
	OutputSize    string `json:"4. Output Size,omitempty"`
	TimeZone      string `json:"5. Time Zone,omitempty"`
}

type TimeSeries struct {
	Open             float64 `json:"1. open,string,omitempty"`
	High             float64 `json:"2. high,string,omitempty"`
	Low              float64 `json:"3. low,string,omitempty"`
	Close            float64 `json:"4. close,string,omitempty"`
	AdjustedClose    float64 `json:"5. adjusted close,string,omitempty"`
	Volume           int64   `json:"6. volume,string,omitempty"`
	DividendAmount   float64 `json:"7. dividend amount,string,omitempty"`
	SplitCoefficient float64 `json:"8. split coefficient,string,omitempty"`
}

func (c *Client) DailyAdjusted(symbol, outputSize string) (*TimeSeriesResponse, error) {
	if _, ok := outputSizeOpts[outputSize]; !ok {
		return nil, errors.New("invalid output size param")
	}

	q := url.Values{}
	q.Add("function", "TIME_SERIES_DAILY_ADJUSTED")
	q.Add("outputsize", outputSize)
	q.Add("symbol", symbol)

	resp, err := c.Do(q)
	if err != nil {
		return nil, fmt.Errorf("fetching daily adjusted series: %w", err)
	}

	var ts TimeSeriesResponse
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&ts); err != nil {
		return nil, fmt.Errorf("parsing daily adjusted series: %w", err)
	}

	if strings.Contains(ts.Note, "Thank you for using Alpha Vantage!") {
		return nil, ErrTooManyRequests
	}

	return &ts, nil
}
