package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/tgmendes/financial-tracker/pkg/dao"
	"github.com/tgmendes/financial-tracker/pkg/handler"
	"os"
)

func main() {
	err := godotenv.Load("env.local")
	if err != nil {
		panic(err)
	}

	connStr := os.Getenv("POSTGRES_URL")
	apiKey := os.Getenv("API_KEY")

	ctx := context.Background()
	dbpool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	q := dao.New(dbpool)

	h := handler.NewStockFetcher(apiKey, q)
	err = h.AllTimeData(ctx)
	if err != nil {
		panic(err)
	}
}
