package alphavantage

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApiResponse(t *testing.T) {
	sample := getTestData()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(sample)
	}))
	defer srv.Close()

	cl := NewClient(srv.URL, "some-key")
	ts, err := cl.DailyAdjusted("test", "compact")
	require.NoError(t, err)

	_, ok := ts.TimeSeries["2023-04-21"]
	assert.True(t, ok)
	assert.Equal(t, "IBM", ts.Metadata.Symbol)
}

func TestRateLimitReached(t *testing.T) {
	jsonResp := `
{
    "Note": "Thank you for using Alpha Vantage! Our standard API call frequency is 5 calls per minute and 500 calls per day. Please visit https://www.alphavantage.co/premium/ if you would like to target a higher API call frequency."
}
`
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(jsonResp))
	}))
	defer srv.Close()

	cl := NewClient(srv.URL, "some-key")
	ts, err := cl.DailyAdjusted("test", "compact")
	assert.Error(t, err)
	assert.True(t, errors.Is(err, ErrTooManyRequests))
	assert.Nil(t, ts)
}

func getTestData() []byte {
	testData := `
	{
  "Meta Data": {
    "1. Information": "Daily Time Series with Splits and Dividend Events",
    "2. Symbol": "IBM",
    "3. Last Refreshed": "2023-04-21",
    "4. Output Size": "Compact",
    "5. Time Zone": "US/Eastern"
  },
  "Time Series (Daily)": {
    "2023-04-21": {
      "1. open": "126.0",
      "2. high": "126.7",
      "3. low": "125.27",
      "4. close": "125.73",
      "5. adjusted close": "125.73",
      "6. volume": "6725426",
      "7. dividend amount": "0.0000",
      "8. split coefficient": "1.0"
    },
    "2023-04-20": {
      "1. open": "130.15",
      "2. high": "130.98",
      "3. low": "125.84",
      "4. close": "126.36",
      "5. adjusted close": "126.36",
      "6. volume": "9749618",
      "7. dividend amount": "0.0000",
      "8. split coefficient": "1.0"
    }
  }
}
`
	return []byte(testData)
}
