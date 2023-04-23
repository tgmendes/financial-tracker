package main

import (
	"fmt"
	"github.com/joho/godotenv"
	alphavantage "github.com/tgmendes/financial-tracker/alpha-vantage"
	"os"
)

func main() {
	err := godotenv.Load("env.local")
	if err != nil {
		panic(err)
	}

	apiKey := os.Getenv("API_KEY")
	baseURL := alphavantage.BASE_URL

	avClient := alphavantage.NewClient(baseURL, apiKey)

	ts, err := avClient.DailyAdjusted("VWRP.L", "compact")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", ts)
}
